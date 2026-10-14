#ifndef EMISDK_STREAM_MEDIA_ACTION_H
#define EMISDK_STREAM_MEDIA_ACTION_H
#include "cjson/cJSON.h"
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "fetchx.h"
#include "emiws.h"
#include "emimedia.h"
typedef struct StreamMediaActionPathParameters {
    char* deviceId;
} StreamMediaActionPathParameters;
static inline void StreamMediaActionPathParameters_init(StreamMediaActionPathParameters* self) {
    if (!self) return;
    self->deviceId = NULL;
}
static inline void StreamMediaActionPathParameters_free_fields(StreamMediaActionPathParameters* self) {
    if (!self) return;
    free(self->deviceId);
    self->deviceId = NULL;
}
/* Returns a newly allocated url - the caller must free() it. */
static inline char* StreamMediaActionPathParameters_apply(const StreamMediaActionPathParameters* self, const char* template_url) {
    char* url = strdup(template_url);
    {
        char* next = fetchx_url_replace(url, ":deviceId", self->deviceId ? self->deviceId : "");
        free(url);
        url = next;
    }
    return url;
}
typedef struct StreamMediaActionQueryParams {
    bool sampleRate_set;
    int32_t sampleRate;
    bool channels_set;
    int32_t channels;
    char* audioCodec; /* NULL if unset */
    char* videoCodec; /* NULL if unset */
} StreamMediaActionQueryParams;
static inline void StreamMediaActionQueryParams_init(StreamMediaActionQueryParams* self) {
    if (!self) return;
    self->sampleRate_set = false;
    self->sampleRate = (int32_t) 0;
    self->channels_set = false;
    self->channels = (int32_t) 0;
    self->audioCodec = NULL;
    self->videoCodec = NULL;
}
static inline void StreamMediaActionQueryParams_free_fields(StreamMediaActionQueryParams* self) {
    if (!self) return;
    free(self->audioCodec);
    self->audioCodec = NULL;
    free(self->videoCodec);
    self->videoCodec = NULL;
}
/* Returns a newly allocated url with every set field appended as a query
 * parameter - the caller must free() it (url itself is never modified or
 * freed). */
static inline char* StreamMediaActionQueryParams_apply(const StreamMediaActionQueryParams* self, const char* url) {
    char* out = strdup(url);
    if (!self) return out;
    if (self->sampleRate_set) {
        char value_buf[64];
        snprintf(value_buf, sizeof(value_buf), "%lld", (long long) self->sampleRate);
        char* next = fetchx_url_append_query(out, "sampleRate", value_buf);
        free(out);
        out = next;
    }
    if (self->channels_set) {
        char value_buf[64];
        snprintf(value_buf, sizeof(value_buf), "%lld", (long long) self->channels);
        char* next = fetchx_url_append_query(out, "channels", value_buf);
        free(out);
        out = next;
    }
    if (self->audioCodec) {
        char* next = fetchx_url_append_query(out, "audioCodec", self->audioCodec);
        free(out);
        out = next;
    }
    if (self->videoCodec) {
        char* next = fetchx_url_append_query(out, "videoCodec", self->videoCodec);
        free(out);
        out = next;
    }
    return out;
}
/* Communicates with StreamMediaAction (REACTIVE /media/stream/:deviceId) */
/* Opens the StreamMediaAction websocket (binary frames, see emiws.h / emimedia.h).
 *
 *   EmiWebSocket ws; emiws_init(&ws);              // then set ws.on_binary / ws.on_text / ws.user
 *   EmiPosixStream state; EmiStream stream; emiws_posix_stream_init(&stream, &state);
 *   if (stream_media_action(&ws, &stream, ctx, &path, &query) == 0) {
 *       while (ws.open) emiws_poll(&ws);              // or only send, e.g. with EmiMediaSender
 *   }
 *
 * Returns 0 once the handshake completed. Close with emiws_close(&ws). */
static inline int stream_media_action(EmiWebSocket* ws, EmiStream* stream, const FetchxContext* ctx, const StreamMediaActionPathParameters* path, const StreamMediaActionQueryParams* query) {
    char* base = fetchx_build_url(ctx ? ctx->base_url : "", "/media/stream/:deviceId");
    char* url1 = StreamMediaActionPathParameters_apply(path, base);
    free(base);
    base = url1;
    char* url2 = StreamMediaActionQueryParams_apply(query, base);
    free(base);
    base = url2;
    char* ws_url = emiws_http_to_ws(base);
    free(base);
    int rc = ws_url ? emiws_connect(ws, stream, ws_url) : -1;
    free(ws_url);
    return rc;
}
#endif /* EMISDK_STREAM_MEDIA_ACTION_H */
