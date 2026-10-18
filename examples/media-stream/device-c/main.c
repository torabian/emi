/* Host build of the plain-C client (the same calls the ESP-IDF project in
 * ../esp-idf-mic makes): a fake 440 Hz microphone streamed as PCM16 frames.
 *
 *   make device-c && ./device-c/media-device-c [host] [port] [seconds]
 */
#include "emi/stream_media_action.h"

#include <math.h>
#include <stdlib.h>
#include <time.h>
#include <unistd.h>

#define SAMPLE_RATE 16000
#define FRAME_MS 20
#define SAMPLES_PER_FRAME (SAMPLE_RATE * FRAME_MS / 1000)

static uint32_t now_ms(void) {
    struct timespec ts;
    clock_gettime(CLOCK_MONOTONIC, &ts);
    return (uint32_t) (ts.tv_sec * 1000u + ts.tv_nsec / 1000000u);
}

int main(int argc, char** argv) {
    const char* host = argc > 1 ? argv[1] : "127.0.0.1";
    const char* port = argc > 2 ? argv[2] : "8080";
    int seconds = argc > 3 ? atoi(argv[3]) : 3;

    char base[128];
    snprintf(base, sizeof(base), "http://%s:%s", host, port);
    FetchxContext* ctx = FetchxContext_new();
    FetchxContext_set_base_url(ctx, base);

    /* Typed path + query parameters straight from media.emi.yml. */
    StreamMediaActionPathParameters path;
    StreamMediaActionPathParameters_init(&path);
    path.deviceId = strdup("esp32-c-mic");
    StreamMediaActionQueryParams query;
    StreamMediaActionQueryParams_init(&query);
    query.sampleRate_set = true; query.sampleRate = SAMPLE_RATE;
    query.channels_set = true;   query.channels = 1;
    query.audioCodec = strdup("pcm16");

    EmiWebSocket ws;
    EmiPosixStream state;
    EmiStream stream;
    emiws_init(&ws);
    emiws_posix_stream_init(&stream, &state);

    if (stream_media_action(&ws, &stream, ctx, &path, &query) != 0) {
        fprintf(stderr, "could not open the media websocket\n");
        return 1;
    }

    EmiMediaSender sender;
    emimedia_sender_init(&sender, &ws);

    int16_t pcm[SAMPLES_PER_FRAME];
    double phase = 0;
    uint32_t start = now_ms();
    int frames = 0;
    while (ws.open && now_ms() - start < (uint32_t) seconds * 1000u) {
        for (int i = 0; i < SAMPLES_PER_FRAME; i++) { /* ESP-IDF: i2s_channel_read() */
            pcm[i] = (int16_t) (8000 * sin(phase));
            phase += 2 * M_PI * 440.0 / SAMPLE_RATE;
        }
        /* PCM16 is little endian on the wire = native order on x86/arm/ESP32. */
        if (emimedia_send_audio(&sender, (const uint8_t*) pcm, sizeof(pcm), now_ms() - start) != 0) break;
        frames++;
        usleep(FRAME_MS * 1000);
    }

    emiws_close(&ws);
    printf("sent %d audio frames\n", frames);
    StreamMediaActionPathParameters_free_fields(&path);
    StreamMediaActionQueryParams_free_fields(&query);
    FetchxContext_free(ctx);
    return 0;
}
