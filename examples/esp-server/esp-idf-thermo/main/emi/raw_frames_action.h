#ifndef EMISDK_RAW_FRAMES_ACTION_H
#define EMISDK_RAW_FRAMES_ACTION_H
#include "cjson/cJSON.h"
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "emihttpd_esp.h"
#include "fetchx.h"
#ifdef EMIHTTPD_NO_WS
#error "RawFramesAction is a reactive action: enable CONFIG_HTTPD_WS_SUPPORT (menuconfig -> Component config -> HTTP Server -> WebSocket server support)"
#endif
/* Serves the RawFramesAction websocket (/ws/raw).
 *
 * Fill RawFramesActionCallbacks, then register it once the http server is running:
 *
 *   RawFramesActionBinding* b = raw_frames_action_register(server, &callbacks);
 *   ...
 *   raw_frames_action_broadcast(b, &update);   // push to every connected client, from any task
 *
 * With esp_http_server set config.uri_match_fn = emihttpd_uri_match when the url has :params.
 * Clients that vanish are dropped (on_close fires): on the next failed send with
 * esp_http_server, when the connection closes with Mongoose. */
typedef struct RawFramesActionClientState {
    emihttpd_conn_t conn;
} RawFramesActionClientState;
typedef struct RawFramesActionSession {
    emihttpd_server_t server;
    emihttpd_conn_t conn; /* use with raw_frames_action_send*() */
    void* user;
} RawFramesActionSession;
typedef struct RawFramesActionCallbacks {
    /* Called after the handshake; return non-zero to reject the client. */
    int (*on_open)(const RawFramesActionSession* session);
    /* A binary frame; data is only valid during the call. */
    void (*on_binary)(const RawFramesActionSession* session, const uint8_t* data, size_t len);
    /* The client closed or was dropped. */
    void (*on_close)(const RawFramesActionSession* session);
    void* user;
} RawFramesActionCallbacks;
typedef struct RawFramesActionBinding {
    RawFramesActionCallbacks cb;
    EmiHttpdClients clients;
    emihttpd_server_t server;
    void* route;
} RawFramesActionBinding;
static inline void raw_frames_action_fill_session(RawFramesActionBinding* b, RawFramesActionClientState* st, RawFramesActionSession* s) {
    s->server = b->server;
    s->conn = st->conn;
    s->user = b->cb.user;
}
static inline void raw_frames_action_free_state(RawFramesActionClientState* st) {
    if (!st) return;
    free(st);
}
/* Forgets a client (calls on_close). Safe to call for unknown connections. */
static inline void raw_frames_action_drop_client(RawFramesActionBinding* b, emihttpd_conn_t conn) {
    RawFramesActionClientState* st = (RawFramesActionClientState*) emihttpd_clients_remove(&b->clients, conn);
    if (!st) return;
    if (b->cb.on_close) {
        RawFramesActionSession s;
        raw_frames_action_fill_session(b, st, &s);
        b->cb.on_close(&s);
    }
    raw_frames_action_free_state(st);
}
/* Number of clients currently connected. */
static inline size_t raw_frames_action_client_count(RawFramesActionBinding* b) {
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    return emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
}
static inline int raw_frames_action_write(RawFramesActionBinding* b, emihttpd_conn_t conn, bool binary, const uint8_t* data, size_t len) {
    emihttpd_mutex_lock(&b->clients.send_lock);
    int rc = emihttpd_ws_write(b->server, conn, binary, data, len);
    emihttpd_mutex_unlock(&b->clients.send_lock);
    if (rc != 0) { raw_frames_action_drop_client(b, conn); return -1; }
    return 0;
}
/* Sends a text frame to one client; 0 on success (queued, with Mongoose), the client is dropped on failure. */
static inline int raw_frames_action_send_text(RawFramesActionBinding* b, emihttpd_conn_t conn, const char* text) {
    return raw_frames_action_write(b, conn, false, (const uint8_t*) text, strlen(text));
}
/* Sends a text frame to every client; returns how many received it. */
static inline int raw_frames_action_broadcast_text(RawFramesActionBinding* b, const char* text) {
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    size_t n = emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
    int delivered = 0;
    for (size_t i = 0; i < n; i++) {
        if (raw_frames_action_send_text(b, conns[i], text) == 0) delivered++;
    }
    return delivered;
}
/* Sends a binary frame to one client; 0 on success, the client is dropped on failure. */
static inline int raw_frames_action_send_binary(RawFramesActionBinding* b, emihttpd_conn_t conn, const uint8_t* data, size_t len) {
    return raw_frames_action_write(b, conn, true, data, len);
}
/* Sends a binary frame to every client; returns how many received it. */
static inline int raw_frames_action_broadcast_binary(RawFramesActionBinding* b, const uint8_t* data, size_t len) {
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    size_t n = emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
    int delivered = 0;
    for (size_t i = 0; i < n; i++) {
        if (raw_frames_action_send_binary(b, conns[i], data, len) == 0) delivered++;
    }
    return delivered;
}
/* The handshake: decode path/query, remember the client, ask on_open. */
static inline int raw_frames_action_ws_open(void* request, emihttpd_conn_t conn, void* ctx) {
    RawFramesActionBinding* b = (RawFramesActionBinding*) ctx;
    emihttpd_req_t* req = (emihttpd_req_t*) request;
    RawFramesActionClientState* st = (RawFramesActionClientState*) calloc(1, sizeof(*st));
    if (!st) return -1;
    st->conn = conn;
    (void) req;
    raw_frames_action_drop_client(b, conn); /* a stale entry for a reused connection id */
    if (emihttpd_clients_add(&b->clients, conn, st) != 0) {
        raw_frames_action_free_state(st);
        return -1;
    }
    if (b->cb.on_open) {
        RawFramesActionSession s;
        raw_frames_action_fill_session(b, st, &s);
        if (b->cb.on_open(&s) != 0) {
            emihttpd_clients_remove(&b->clients, conn);
            raw_frames_action_free_state(st);
            return -1;
        }
    }
    return 0;
}
/* One inbound frame; data is NUL terminated. */
static inline void raw_frames_action_ws_message(emihttpd_conn_t conn, bool binary, const uint8_t* data, size_t len, void* ctx) {
    RawFramesActionBinding* b = (RawFramesActionBinding*) ctx;
    RawFramesActionClientState* st = (RawFramesActionClientState*) emihttpd_clients_find(&b->clients, conn);
    if (!st) return;
    if (binary) {
        if (b->cb.on_binary) {
            RawFramesActionSession s;
            raw_frames_action_fill_session(b, st, &s);
            b->cb.on_binary(&s, data, len);
        }
        return;
    }
    (void) len;
}
static inline void raw_frames_action_ws_close(emihttpd_conn_t conn, void* ctx) {
    raw_frames_action_drop_client((RawFramesActionBinding*) ctx, conn);
}
/* Registers the websocket route; returns the binding or NULL on failure. */
static inline RawFramesActionBinding* raw_frames_action_register(emihttpd_server_t server, const RawFramesActionCallbacks* callbacks) {
    RawFramesActionBinding* b = (RawFramesActionBinding*) calloc(1, sizeof(*b));
    if (!b) return NULL;
    if (callbacks) b->cb = *callbacks;
    b->server = server;
    if (emihttpd_clients_init(&b->clients) != 0) {
        free(b);
        return NULL;
    }
    EmiHttpdWsHooks hooks = {raw_frames_action_ws_open, raw_frames_action_ws_message, raw_frames_action_ws_close};
    b->route = emihttpd_add_ws_route(server, "/ws/raw", &hooks, b);
    if (!b->route) {
        emihttpd_clients_deinit(&b->clients);
        free(b);
        return NULL;
    }
    return b;
}
/* Unregisters the route, drops every client (on_close is called) and frees the binding. */
static inline void raw_frames_action_unregister(emihttpd_server_t server, RawFramesActionBinding* b) {
    if (!b) return;
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    emihttpd_remove_route(server, b->route);
    size_t n = emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
    for (size_t i = 0; i < n; i++) {
        raw_frames_action_drop_client(b, conns[i]);
    }
    emihttpd_clients_deinit(&b->clients);
    free(b);
}
#endif /* EMISDK_RAW_FRAMES_ACTION_H */
