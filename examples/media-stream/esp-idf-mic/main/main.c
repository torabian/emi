/*
 * Streams an I2S microphone (INMP441 / SPH0645 / any I2S MEMS mic) to the Go
 * server of examples/media-stream as raw PCM16 frames, through the code emi
 * generated from media.emi.yml:
 *
 *   - stream_media_action(): typed path/query parameters -> websocket handshake
 *   - EmiMediaSender:        12 byte frame header + payload, one websocket frame
 *
 * Target: ESP-IDF v5.x (uses the i2s_std driver). `make gen` in the parent
 * directory (re)generates main/emi/.
 */
#include <string.h>

#include "esp_event.h"
#include "esp_log.h"
#include "esp_netif.h"
#include "esp_random.h"
#include "esp_timer.h"
#include "esp_wifi.h"
#include "freertos/FreeRTOS.h"
#include "freertos/event_groups.h"
#include "freertos/semphr.h"
#include "freertos/task.h"
#include "driver/i2s_std.h"
#include "nvs_flash.h"

/* Websocket mask keys / nonce from the hardware RNG instead of rand(). */
#define EMIWS_RANDOM32() esp_random()
#include "emi/stream_media_action.h"

static const char* TAG = "emi-mic";

#define SAMPLE_RATE CONFIG_EMI_SAMPLE_RATE
#define FRAME_SAMPLES (SAMPLE_RATE / 50) /* 20 ms */

/* ---- Wi-Fi ----------------------------------------------------------------- */

static EventGroupHandle_t s_wifi_events;
#define WIFI_CONNECTED_BIT BIT0

static void on_wifi_event(void* arg, esp_event_base_t base, int32_t id, void* data) {
    (void) arg; (void) data;
    if (base == WIFI_EVENT && id == WIFI_EVENT_STA_START) {
        esp_wifi_connect();
    } else if (base == WIFI_EVENT && id == WIFI_EVENT_STA_DISCONNECTED) {
        xEventGroupClearBits(s_wifi_events, WIFI_CONNECTED_BIT);
        esp_wifi_connect(); /* keep retrying; the stream loop below reconnects the socket */
    } else if (base == IP_EVENT && id == IP_EVENT_STA_GOT_IP) {
        xEventGroupSetBits(s_wifi_events, WIFI_CONNECTED_BIT);
    }
}

static void wifi_start(void) {
    s_wifi_events = xEventGroupCreate();
    ESP_ERROR_CHECK(esp_netif_init());
    ESP_ERROR_CHECK(esp_event_loop_create_default());
    esp_netif_create_default_wifi_sta();

    wifi_init_config_t init = WIFI_INIT_CONFIG_DEFAULT();
    ESP_ERROR_CHECK(esp_wifi_init(&init));
    ESP_ERROR_CHECK(esp_event_handler_register(WIFI_EVENT, ESP_EVENT_ANY_ID, on_wifi_event, NULL));
    ESP_ERROR_CHECK(esp_event_handler_register(IP_EVENT, IP_EVENT_STA_GOT_IP, on_wifi_event, NULL));

    wifi_config_t cfg = { 0 };
    strlcpy((char*) cfg.sta.ssid, CONFIG_EMI_WIFI_SSID, sizeof(cfg.sta.ssid));
    strlcpy((char*) cfg.sta.password, CONFIG_EMI_WIFI_PASSWORD, sizeof(cfg.sta.password));
    ESP_ERROR_CHECK(esp_wifi_set_mode(WIFI_MODE_STA));
    ESP_ERROR_CHECK(esp_wifi_set_config(WIFI_IF_STA, &cfg));
    ESP_ERROR_CHECK(esp_wifi_start());
    esp_wifi_set_ps(WIFI_PS_NONE); /* power save adds tens of ms of latency to every frame */
}

/* ---- microphone -------------------------------------------------------------- */

static i2s_chan_handle_t s_mic;

static void mic_start(void) {
    i2s_chan_config_t chan = I2S_CHANNEL_DEFAULT_CONFIG(I2S_NUM_AUTO, I2S_ROLE_MASTER);
    ESP_ERROR_CHECK(i2s_new_channel(&chan, NULL, &s_mic));

    i2s_std_config_t cfg = {
        .clk_cfg = I2S_STD_CLK_DEFAULT_CONFIG(SAMPLE_RATE),
        .slot_cfg = I2S_STD_PHILIPS_SLOT_DEFAULT_CONFIG(I2S_DATA_BIT_WIDTH_32BIT, I2S_SLOT_MODE_MONO),
        .gpio_cfg = {
            .mclk = I2S_GPIO_UNUSED,
            .bclk = CONFIG_EMI_I2S_BCLK,
            .ws = CONFIG_EMI_I2S_WS,
            .dout = I2S_GPIO_UNUSED,
            .din = CONFIG_EMI_I2S_DIN,
            .invert_flags = { 0 },
        },
    };
    cfg.slot_cfg.slot_mask = I2S_STD_SLOT_LEFT; /* mic with L/R tied to GND */
    ESP_ERROR_CHECK(i2s_channel_init_std_mode(s_mic, &cfg));
    ESP_ERROR_CHECK(i2s_channel_enable(s_mic));
}

/* Reads one 20 ms frame as PCM16. Returns the number of samples, 0 on timeout. */
static size_t mic_read_frame(int16_t* out) {
    static int32_t raw[FRAME_SAMPLES];
    size_t got = 0;
    if (i2s_channel_read(s_mic, raw, sizeof(raw), &got, 1000) != ESP_OK) return 0;
    size_t n = got / sizeof(int32_t);
    for (size_t i = 0; i < n; i++) {
        int32_t v = raw[i] >> CONFIG_EMI_MIC_SHIFT;
        out[i] = (int16_t) (v > 32767 ? 32767 : (v < -32768 ? -32768 : v));
    }
    return n; /* PCM16 little endian on the wire = the ESP32's native byte order */
}

/* ---- websocket ------------------------------------------------------------------ */

static SemaphoreHandle_t s_send_lock;
static void ws_lock(void* ctx) { xSemaphoreTake((SemaphoreHandle_t) ctx, portMAX_DELAY); }
static void ws_unlock(void* ctx) { xSemaphoreGive((SemaphoreHandle_t) ctx); }

/* Reads (and answers pings / notices close) so a dead connection is noticed. */
static void ws_reader_task(void* arg) {
    EmiWebSocket* ws = (EmiWebSocket*) arg;
    while (ws->open) emiws_poll(ws);
    vTaskDelete(NULL);
}

static void on_server_text(void* user, const char* text, size_t len) {
    (void) user;
    ESP_LOGI(TAG, "server: %.*s", (int) len, text);
}

/* One connected session; returns when the connection is lost. */
static void stream_once(void) {
    static EmiWebSocket ws;
    static EmiPosixStream state;
    static EmiStream stream;

    FetchxContext* ctx = FetchxContext_new();
    FetchxContext_set_base_url(ctx, CONFIG_EMI_SERVER_URL);

    /* Typed parameters generated from media.emi.yml. */
    StreamMediaActionPathParameters path;
    StreamMediaActionPathParameters_init(&path);
    path.deviceId = strdup(CONFIG_EMI_DEVICE_ID);
    StreamMediaActionQueryParams query;
    StreamMediaActionQueryParams_init(&query);
    query.sampleRate_set = true; query.sampleRate = SAMPLE_RATE;
    query.channels_set = true;   query.channels = 1;
    query.audioCodec = strdup("pcm16");

    emiws_init(&ws);
    ws.on_text = on_server_text;
    ws.lock = ws_lock;
    ws.unlock = ws_unlock;
    ws.lock_ctx = s_send_lock;
    emiws_posix_stream_init(&stream, &state);

    int rc = stream_media_action(&ws, &stream, ctx, &path, &query);
    StreamMediaActionPathParameters_free_fields(&path);
    StreamMediaActionQueryParams_free_fields(&query);
    FetchxContext_free(ctx);
    if (rc != 0) {
        ESP_LOGW(TAG, "connect failed");
        return;
    }
    ESP_LOGI(TAG, "streaming %d Hz mono PCM16 to %s", SAMPLE_RATE, CONFIG_EMI_SERVER_URL);

    xTaskCreate(ws_reader_task, "ws_rx", 4096, &ws, 5, NULL);

    EmiMediaSender sender;
    emimedia_sender_init(&sender, &ws);
    static int16_t pcm[FRAME_SAMPLES];
    int64_t start_us = esp_timer_get_time();

    while (ws.open) {
        size_t n = mic_read_frame(pcm);
        if (n == 0) continue;
        uint32_t ts_ms = (uint32_t) ((esp_timer_get_time() - start_us) / 1000);
        if (emimedia_send_audio(&sender, (const uint8_t*) pcm, n * sizeof(int16_t), ts_ms) != 0) break;
    }

    emiws_close(&ws);
    vTaskDelay(pdMS_TO_TICKS(100)); /* let the reader task observe ws.open == 0 before ws is reused */
}

void app_main(void) {
    esp_err_t err = nvs_flash_init();
    if (err == ESP_ERR_NVS_NO_FREE_PAGES || err == ESP_ERR_NVS_NEW_VERSION_FOUND) {
        ESP_ERROR_CHECK(nvs_flash_erase());
        err = nvs_flash_init();
    }
    ESP_ERROR_CHECK(err);

    s_send_lock = xSemaphoreCreateMutex();
    wifi_start();
    mic_start();

    while (1) {
        xEventGroupWaitBits(s_wifi_events, WIFI_CONNECTED_BIT, pdFALSE, pdTRUE, portMAX_DELAY);
        stream_once();
        ESP_LOGW(TAG, "disconnected, retrying in 2 s");
        vTaskDelay(pdMS_TO_TICKS(2000));
    }
}
