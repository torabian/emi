#ifndef EMISDK_STATE_STREAM_ACTION_H
#define EMISDK_STATE_STREAM_ACTION_H
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
#error "StateStreamAction is a reactive action: enable CONFIG_HTTPD_WS_SUPPORT (menuconfig -> Component config -> HTTP Server -> WebSocket server support)"
#endif
typedef struct StateStreamActionPathParameters {
    char* room;
} StateStreamActionPathParameters;
static inline void StateStreamActionPathParameters_init(StateStreamActionPathParameters* self) {
    if (!self) return;
    self->room = NULL;
}
static inline void StateStreamActionPathParameters_free_fields(StateStreamActionPathParameters* self) {
    if (!self) return;
    free(self->room);
    self->room = NULL;
}
/* Returns a newly allocated url - the caller must free() it. */
static inline char* StateStreamActionPathParameters_apply(const StateStreamActionPathParameters* self, const char* template_url) {
    char* url = strdup(template_url);
    {
        char* next = fetchx_url_replace(url, ":room", self->room ? self->room : "");
        free(url);
        url = next;
    }
    return url;
}
typedef struct StateStreamActionQueryParams {
    bool interval_set;
    int32_t interval;
} StateStreamActionQueryParams;
static inline void StateStreamActionQueryParams_init(StateStreamActionQueryParams* self) {
    if (!self) return;
    self->interval_set = false;
    self->interval = (int32_t) 0;
}
static inline void StateStreamActionQueryParams_free_fields(StateStreamActionQueryParams* self) {
    if (!self) return;
}
/* Returns a newly allocated url with every set field appended as a query
 * parameter - the caller must free() it (url itself is never modified or
 * freed). */
static inline char* StateStreamActionQueryParams_apply(const StateStreamActionQueryParams* self, const char* url) {
    char* out = strdup(url);
    if (!self) return out;
    if (self->interval_set) {
        char value_buf[64];
        snprintf(value_buf, sizeof(value_buf), "%lld", (long long) self->interval);
        char* next = fetchx_url_append_query(out, "interval", value_buf);
        free(out);
        out = next;
    }
    return out;
}
/* StateStreamActionRequest */
typedef struct StateStreamActionRequest {
    float setPoint;
} StateStreamActionRequest;
/* Forward declarations: a self-referencing field (e.g. a relation whose
 * target: is this very struct) means _free_fields/_from_json_into below
 * can call _free/_from_json, which aren't defined until further down this
 * same file - static inline functions can be forward-declared exactly like
 * any other, so this resolves the ordering regardless of which fields end
 * up actually self-referencing. */
static inline void StateStreamActionRequest_init(StateStreamActionRequest* self);
static inline void StateStreamActionRequest_free_fields(StateStreamActionRequest* self);
static inline StateStreamActionRequest* StateStreamActionRequest_new(void);
static inline void StateStreamActionRequest_free(StateStreamActionRequest* self);
static inline void StateStreamActionRequest_from_json_into(StateStreamActionRequest* self, const cJSON* json);
static inline StateStreamActionRequest* StateStreamActionRequest_from_json(const cJSON* json);
static inline cJSON* StateStreamActionRequest_to_json(const StateStreamActionRequest* self);
static inline void StateStreamActionRequest_init(StateStreamActionRequest* self) {
    if (!self) return;
    self->setPoint = 0;
}
static inline void StateStreamActionRequest_free_fields(StateStreamActionRequest* self) {
    if (!self) return;
}
static inline StateStreamActionRequest* StateStreamActionRequest_new(void) {
    StateStreamActionRequest* self = (StateStreamActionRequest*) malloc(sizeof(StateStreamActionRequest));
    StateStreamActionRequest_init(self);
    return self;
}
static inline void StateStreamActionRequest_free(StateStreamActionRequest* self) {
    if (!self) return;
    StateStreamActionRequest_free_fields(self);
    free(self);
}
static inline void StateStreamActionRequest_from_json_into(StateStreamActionRequest* self, const cJSON* json) {
    if (!self || !json) return;
    {
        const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, "setPoint");
        if (cJSON_IsNumber(v)) { self->setPoint = (float) v->valuedouble; }
        else if (cJSON_IsBool(v)) { self->setPoint = (float) cJSON_IsTrue(v); }
    }
}
static inline StateStreamActionRequest* StateStreamActionRequest_from_json(const cJSON* json) {
    StateStreamActionRequest* self = StateStreamActionRequest_new();
    StateStreamActionRequest_from_json_into(self, json);
    return self;
}
static inline cJSON* StateStreamActionRequest_to_json(const StateStreamActionRequest* self) {
    cJSON* json = cJSON_CreateObject();
    if (!self) return json;
    {
        char float_buf[32];
        snprintf(float_buf, sizeof(float_buf), "%.7g", (double) self->setPoint);
        cJSON_AddNumberToObject(json, "setPoint", strtod(float_buf, NULL));
    }
    return json;
}
/* StateStreamActionResponse */
typedef struct StateStreamActionResponse {
    float inside;
    float outside;
} StateStreamActionResponse;
/* Forward declarations: a self-referencing field (e.g. a relation whose
 * target: is this very struct) means _free_fields/_from_json_into below
 * can call _free/_from_json, which aren't defined until further down this
 * same file - static inline functions can be forward-declared exactly like
 * any other, so this resolves the ordering regardless of which fields end
 * up actually self-referencing. */
static inline void StateStreamActionResponse_init(StateStreamActionResponse* self);
static inline void StateStreamActionResponse_free_fields(StateStreamActionResponse* self);
static inline StateStreamActionResponse* StateStreamActionResponse_new(void);
static inline void StateStreamActionResponse_free(StateStreamActionResponse* self);
static inline void StateStreamActionResponse_from_json_into(StateStreamActionResponse* self, const cJSON* json);
static inline StateStreamActionResponse* StateStreamActionResponse_from_json(const cJSON* json);
static inline cJSON* StateStreamActionResponse_to_json(const StateStreamActionResponse* self);
static inline void StateStreamActionResponse_init(StateStreamActionResponse* self) {
    if (!self) return;
    self->inside = 0;
    self->outside = 0;
}
static inline void StateStreamActionResponse_free_fields(StateStreamActionResponse* self) {
    if (!self) return;
}
static inline StateStreamActionResponse* StateStreamActionResponse_new(void) {
    StateStreamActionResponse* self = (StateStreamActionResponse*) malloc(sizeof(StateStreamActionResponse));
    StateStreamActionResponse_init(self);
    return self;
}
static inline void StateStreamActionResponse_free(StateStreamActionResponse* self) {
    if (!self) return;
    StateStreamActionResponse_free_fields(self);
    free(self);
}
static inline void StateStreamActionResponse_from_json_into(StateStreamActionResponse* self, const cJSON* json) {
    if (!self || !json) return;
    {
        const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, "inside");
        if (cJSON_IsNumber(v)) { self->inside = (float) v->valuedouble; }
        else if (cJSON_IsBool(v)) { self->inside = (float) cJSON_IsTrue(v); }
    }
    {
        const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, "outside");
        if (cJSON_IsNumber(v)) { self->outside = (float) v->valuedouble; }
        else if (cJSON_IsBool(v)) { self->outside = (float) cJSON_IsTrue(v); }
    }
}
static inline StateStreamActionResponse* StateStreamActionResponse_from_json(const cJSON* json) {
    StateStreamActionResponse* self = StateStreamActionResponse_new();
    StateStreamActionResponse_from_json_into(self, json);
    return self;
}
static inline cJSON* StateStreamActionResponse_to_json(const StateStreamActionResponse* self) {
    cJSON* json = cJSON_CreateObject();
    if (!self) return json;
    {
        char float_buf[32];
        snprintf(float_buf, sizeof(float_buf), "%.7g", (double) self->inside);
        cJSON_AddNumberToObject(json, "inside", strtod(float_buf, NULL));
    }
    {
        char float_buf[32];
        snprintf(float_buf, sizeof(float_buf), "%.7g", (double) self->outside);
        cJSON_AddNumberToObject(json, "outside", strtod(float_buf, NULL));
    }
    return json;
}
/* Serves the StateStreamAction websocket (/ws/state/:room).
 *
 * Fill StateStreamActionCallbacks, then register it once the http server is running:
 *
 *   StateStreamActionBinding* b = state_stream_action_register(server, &callbacks);
 *   ...
 *   state_stream_action_broadcast(b, &update);   // push to every connected client, from any task
 *
 * With esp_http_server set config.uri_match_fn = emihttpd_uri_match when the url has :params.
 * Clients that vanish are dropped (on_close fires): on the next failed send with
 * esp_http_server, when the connection closes with Mongoose. */
typedef struct StateStreamActionClientState {
    emihttpd_conn_t conn;
    StateStreamActionPathParameters path;
    StateStreamActionQueryParams query;
} StateStreamActionClientState;
typedef struct StateStreamActionSession {
    emihttpd_server_t server;
    emihttpd_conn_t conn; /* use with state_stream_action_send*() */
    const StateStreamActionPathParameters* path;
    const StateStreamActionQueryParams* query;
    void* user;
} StateStreamActionSession;
typedef struct StateStreamActionCallbacks {
    /* Called after the handshake; return non-zero to reject the client. */
    int (*on_open)(const StateStreamActionSession* session);
    /* A text frame, parsed as StateStreamActionRequest; msg is freed after the call. */
    void (*on_message)(const StateStreamActionSession* session, const StateStreamActionRequest* msg);
    /* The client closed or was dropped. */
    void (*on_close)(const StateStreamActionSession* session);
    void* user;
} StateStreamActionCallbacks;
typedef struct StateStreamActionBinding {
    StateStreamActionCallbacks cb;
    EmiHttpdClients clients;
    emihttpd_server_t server;
    void* route;
} StateStreamActionBinding;
/* Takes ownership of the strings captured by emihttpd_path_capture. */
static inline void StateStreamActionPathParameters_from_captures(StateStreamActionPathParameters* self, char** caps, size_t n) {
    if (n > 0 && caps[0]) {
        free(self->room);
        self->room = caps[0];
        caps[0] = NULL;
    }
    for (size_t i = 0; i < n; i++) free(caps[i]);
}
static inline void StateStreamActionQueryParams_parse(StateStreamActionQueryParams* self, const char* qs) {
    char* v;
    v = emihttpd_query_get(qs, "interval");
    if (v) {
        self->interval_set = true;
        self->interval = (int32_t) strtoll(v, NULL, 10);
        free(v);
    }
}
static inline void state_stream_action_fill_session(StateStreamActionBinding* b, StateStreamActionClientState* st, StateStreamActionSession* s) {
    s->server = b->server;
    s->conn = st->conn;
    s->path = &st->path;
    s->query = &st->query;
    s->user = b->cb.user;
}
static inline void state_stream_action_free_state(StateStreamActionClientState* st) {
    if (!st) return;
    StateStreamActionPathParameters_free_fields(&st->path);
    StateStreamActionQueryParams_free_fields(&st->query);
    free(st);
}
/* Forgets a client (calls on_close). Safe to call for unknown connections. */
static inline void state_stream_action_drop_client(StateStreamActionBinding* b, emihttpd_conn_t conn) {
    StateStreamActionClientState* st = (StateStreamActionClientState*) emihttpd_clients_remove(&b->clients, conn);
    if (!st) return;
    if (b->cb.on_close) {
        StateStreamActionSession s;
        state_stream_action_fill_session(b, st, &s);
        b->cb.on_close(&s);
    }
    state_stream_action_free_state(st);
}
/* Number of clients currently connected. */
static inline size_t state_stream_action_client_count(StateStreamActionBinding* b) {
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    return emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
}
static inline int state_stream_action_write(StateStreamActionBinding* b, emihttpd_conn_t conn, bool binary, const uint8_t* data, size_t len) {
    emihttpd_mutex_lock(&b->clients.send_lock);
    int rc = emihttpd_ws_write(b->server, conn, binary, data, len);
    emihttpd_mutex_unlock(&b->clients.send_lock);
    if (rc != 0) { state_stream_action_drop_client(b, conn); return -1; }
    return 0;
}
/* Sends a text frame to one client; 0 on success (queued, with Mongoose), the client is dropped on failure. */
static inline int state_stream_action_send_text(StateStreamActionBinding* b, emihttpd_conn_t conn, const char* text) {
    return state_stream_action_write(b, conn, false, (const uint8_t*) text, strlen(text));
}
/* Sends a text frame to every client; returns how many received it. */
static inline int state_stream_action_broadcast_text(StateStreamActionBinding* b, const char* text) {
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    size_t n = emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
    int delivered = 0;
    for (size_t i = 0; i < n; i++) {
        if (state_stream_action_send_text(b, conns[i], text) == 0) delivered++;
    }
    return delivered;
}
/* Sends a StateStreamActionResponse as a JSON text frame to one client. */
static inline int state_stream_action_send(StateStreamActionBinding* b, emihttpd_conn_t conn, const StateStreamActionResponse* res) {
    cJSON* json = StateStreamActionResponse_to_json(res);
    char* text = json ? cJSON_PrintUnformatted(json) : NULL;
    cJSON_Delete(json);
    if (!text) return -1;
    int rc = state_stream_action_send_text(b, conn, text);
    free(text);
    return rc;
}
/* Sends a StateStreamActionResponse to every client (serialized once); returns how many received it. */
static inline int state_stream_action_broadcast(StateStreamActionBinding* b, const StateStreamActionResponse* res) {
    cJSON* json = StateStreamActionResponse_to_json(res);
    char* text = json ? cJSON_PrintUnformatted(json) : NULL;
    cJSON_Delete(json);
    if (!text) return 0;
    int delivered = state_stream_action_broadcast_text(b, text);
    free(text);
    return delivered;
}
/* The handshake: decode path/query, remember the client, ask on_open. */
static inline int state_stream_action_ws_open(void* request, emihttpd_conn_t conn, void* ctx) {
    StateStreamActionBinding* b = (StateStreamActionBinding*) ctx;
    emihttpd_req_t* req = (emihttpd_req_t*) request;
    StateStreamActionClientState* st = (StateStreamActionClientState*) calloc(1, sizeof(*st));
    if (!st) return -1;
    st->conn = conn;
    StateStreamActionPathParameters_init(&st->path);
    {
        char* caps[EMIHTTPD_MAX_PARAMS] = {0};
        size_t ncaps = 0;
        if (!emihttpd_req_capture(req, "/ws/state/:room", caps, EMIHTTPD_MAX_PARAMS, &ncaps)) {
            state_stream_action_free_state(st);
            return -1;
        }
        StateStreamActionPathParameters_from_captures(&st->path, caps, ncaps);
    }
    StateStreamActionQueryParams_init(&st->query);
    {
        char* qs = emihttpd_req_query(req);
        StateStreamActionQueryParams_parse(&st->query, qs);
        free(qs);
    }
    state_stream_action_drop_client(b, conn); /* a stale entry for a reused connection id */
    if (emihttpd_clients_add(&b->clients, conn, st) != 0) {
        state_stream_action_free_state(st);
        return -1;
    }
    if (b->cb.on_open) {
        StateStreamActionSession s;
        state_stream_action_fill_session(b, st, &s);
        if (b->cb.on_open(&s) != 0) {
            emihttpd_clients_remove(&b->clients, conn);
            state_stream_action_free_state(st);
            return -1;
        }
    }
    return 0;
}
/* One inbound frame; data is NUL terminated. */
static inline void state_stream_action_ws_message(emihttpd_conn_t conn, bool binary, const uint8_t* data, size_t len, void* ctx) {
    StateStreamActionBinding* b = (StateStreamActionBinding*) ctx;
    StateStreamActionClientState* st = (StateStreamActionClientState*) emihttpd_clients_find(&b->clients, conn);
    if (!st) return;
    if (binary) return;
    (void) len;
    if (b->cb.on_message) {
        cJSON* root = cJSON_Parse((const char*) data);
        StateStreamActionRequest* msg = root ? StateStreamActionRequest_from_json(root) : NULL;
        cJSON_Delete(root);
        if (msg) {
            StateStreamActionSession s;
            state_stream_action_fill_session(b, st, &s);
            b->cb.on_message(&s, msg);
            StateStreamActionRequest_free(msg);
        }
    }
}
static inline void state_stream_action_ws_close(emihttpd_conn_t conn, void* ctx) {
    state_stream_action_drop_client((StateStreamActionBinding*) ctx, conn);
}
/* Registers the websocket route; returns the binding or NULL on failure. */
static inline StateStreamActionBinding* state_stream_action_register(emihttpd_server_t server, const StateStreamActionCallbacks* callbacks) {
    StateStreamActionBinding* b = (StateStreamActionBinding*) calloc(1, sizeof(*b));
    if (!b) return NULL;
    if (callbacks) b->cb = *callbacks;
    b->server = server;
    if (emihttpd_clients_init(&b->clients) != 0) {
        free(b);
        return NULL;
    }
    EmiHttpdWsHooks hooks = {state_stream_action_ws_open, state_stream_action_ws_message, state_stream_action_ws_close};
    b->route = emihttpd_add_ws_route(server, "/ws/state/:room", &hooks, b);
    if (!b->route) {
        emihttpd_clients_deinit(&b->clients);
        free(b);
        return NULL;
    }
    return b;
}
/* Unregisters the route, drops every client (on_close is called) and frees the binding. */
static inline void state_stream_action_unregister(emihttpd_server_t server, StateStreamActionBinding* b) {
    if (!b) return;
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    emihttpd_remove_route(server, b->route);
    size_t n = emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
    for (size_t i = 0; i < n; i++) {
        state_stream_action_drop_client(b, conns[i]);
    }
    emihttpd_clients_deinit(&b->clients);
    free(b);
}
#endif /* EMISDK_STATE_STREAM_ACTION_H */
