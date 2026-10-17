/*
 * A thermostat web API served from the device, through the code emi generated
 * from thermo.emi.yml (`make gen` regenerates main/emi/):
 *
 *   POST /api/ot/set-temp/:zone   set_temperature_action_register()
 *   GET  /api/state               get_state_action_register()
 *   WS   /ws/state/:room          state_stream_action_register()  (push on change)
 *   WS   /ws/raw                  raw_frames_action_register()    (binary frames)
 *
 * Emi owns routing, path/query/body decoding and JSON; this file is only the
 * behaviour behind each action. Target: ESP-IDF v5.x.
 */
#include <string.h>

#include "esp_event.h"
#include "esp_http_server.h"
#include "esp_log.h"
#include "esp_netif.h"
#include "esp_timer.h"
#include "esp_wifi.h"
#include "freertos/FreeRTOS.h"
#include "freertos/task.h"
#include "nvs_flash.h"

#include "emi/get_state_action.h"
#include "emi/raw_frames_action.h"
#include "emi/set_temperature_action.h"
#include "emi/state_stream_action.h"

static const char* TAG = "emi-thermo";

/* The device state the actions read and change. */
static float s_set_point = 21.0f;
static float s_inside = 20.5f;
static float s_outside = 5.0f;

static StateStreamActionBinding* s_state_stream;
static RawFramesActionBinding* s_raw_frames;

/* ---- POST /api/ot/set-temp/:zone ---------------------------------------------- */

static SetTemperatureActionResponse* on_set_temperature(const SetTemperatureActionCall* call, int* status) {
    if (call->body->target < 5.0f || call->body->target > 30.0f) {
        *status = 422;
        return NULL;
    }
    ESP_LOGI(TAG, "zone %d -> %.1f", (int) call->path->zone, (double) call->body->target);
    s_set_point = call->body->target;

    SetTemperatureActionResponse* res = SetTemperatureActionResponse_new();
    res->ok = true;
    res->applied = s_set_point;
    return res;
}

/* ---- GET /api/state?verbose=true --------------------------------------------- */

static GetStateActionResponse* on_get_state(const GetStateActionCall* call, int* status) {
    (void) status;
    if (call->query->verbose) ESP_LOGI(TAG, "verbose state requested");
    GetStateActionResponse* res = GetStateActionResponse_new();
    res->roomSetPoint = s_set_point;
    res->inside = s_inside;
    return res;
}

/* ---- WS /ws/state/:room --------------------------------------------------------- */

static int on_state_open(const StateStreamActionSession* s) {
    ESP_LOGI(TAG, "state client %ld joined room %s", (long) s->conn, s->path->room ? s->path->room : "?");
    return 0; /* non-zero would reject the client */
}

static void on_state_message(const StateStreamActionSession* s, const StateStreamActionRequest* msg) {
    ESP_LOGI(TAG, "client %ld wants set point %.1f", (long) s->conn, (double) msg->setPoint);
    s_set_point = msg->setPoint;
}

static void on_state_close(const StateStreamActionSession* s) {
    ESP_LOGI(TAG, "state client %ld left", (long) s->conn);
}

/* ---- WS /ws/raw ------------------------------------------------------------------- */

static void on_raw_binary(const RawFramesActionSession* s, const uint8_t* data, size_t len) {
    ESP_LOGI(TAG, "client %ld sent %u bytes, echoing", (long) s->conn, (unsigned) len);
    raw_frames_action_send_binary(s_raw_frames, s->conn, data, len);
}

/* The reactive part: whenever the simulated sensors move, every connected
 * client gets the new reading without asking. */
static void sensor_task(void* arg) {
    (void) arg;
    for (;;) {
        s_inside += (s_set_point - s_inside) * 0.05f;
        s_outside += 0.1f;
        if (s_outside > 8.0f) s_outside = 5.0f;

        if (s_state_stream && state_stream_action_client_count(s_state_stream) > 0) {
            StateStreamActionResponse* update = StateStreamActionResponse_new();
            update->inside = s_inside;
            update->outside = s_outside;
            state_stream_action_broadcast(s_state_stream, update);
            StateStreamActionResponse_free(update);
        }
        vTaskDelay(pdMS_TO_TICKS(1000));
    }
}

static void wifi_start_softap(void) {
    ESP_ERROR_CHECK(esp_netif_init());
    ESP_ERROR_CHECK(esp_event_loop_create_default());
    esp_netif_create_default_wifi_ap();
    wifi_init_config_t cfg = WIFI_INIT_CONFIG_DEFAULT();
    ESP_ERROR_CHECK(esp_wifi_init(&cfg));

    wifi_config_t ap = {0};
    strcpy((char*) ap.ap.ssid, "emi-thermo");
    ap.ap.ssid_len = strlen("emi-thermo");
    strcpy((char*) ap.ap.password, "12345678");
    ap.ap.max_connection = 4;
    ap.ap.authmode = WIFI_AUTH_WPA2_PSK;
    ESP_ERROR_CHECK(esp_wifi_set_mode(WIFI_MODE_AP));
    ESP_ERROR_CHECK(esp_wifi_set_config(WIFI_IF_AP, &ap));
    ESP_ERROR_CHECK(esp_wifi_start());
}

void app_main(void) {
    esp_err_t err = nvs_flash_init();
    if (err == ESP_ERR_NVS_NO_FREE_PAGES || err == ESP_ERR_NVS_NEW_VERSION_FOUND) {
        ESP_ERROR_CHECK(nvs_flash_erase());
        err = nvs_flash_init();
    }
    ESP_ERROR_CHECK(err);
    wifi_start_softap();

    httpd_config_t config = HTTPD_DEFAULT_CONFIG();
    config.uri_match_fn = emihttpd_uri_match; /* so "/api/ot/set-temp/:zone" matches */
    config.max_uri_handlers = 16;
    httpd_handle_t server = NULL;
    ESP_ERROR_CHECK(httpd_start(&server, &config));

    set_temperature_action_register(server, on_set_temperature, NULL);
    get_state_action_register(server, on_get_state, NULL);

    StateStreamActionCallbacks state_cb = {
        .on_open = on_state_open,
        .on_message = on_state_message,
        .on_close = on_state_close,
    };
    s_state_stream = state_stream_action_register(server, &state_cb);

    RawFramesActionCallbacks raw_cb = {.on_binary = on_raw_binary};
    s_raw_frames = raw_frames_action_register(server, &raw_cb);

    xTaskCreate(sensor_task, "sensors", 4096, NULL, 5, NULL);
    ESP_LOGI(TAG, "ready: join Wi-Fi emi-thermo, then http://192.168.4.1/api/state");
}
