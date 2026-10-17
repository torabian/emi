#ifndef EMISDK_GET_STATE_ACTION_H
#define EMISDK_GET_STATE_ACTION_H
#include "cjson/cJSON.h"
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "emihttpd_esp.h"
#include "fetchx.h"
typedef struct GetStateActionQueryParams {
    bool verbose_set;
    bool verbose;
    char* prefix; /* NULL if unset */
} GetStateActionQueryParams;
static inline void GetStateActionQueryParams_init(GetStateActionQueryParams* self) {
    if (!self) return;
    self->verbose_set = false;
    self->verbose = (bool) 0;
    self->prefix = NULL;
}
static inline void GetStateActionQueryParams_free_fields(GetStateActionQueryParams* self) {
    if (!self) return;
    free(self->prefix);
    self->prefix = NULL;
}
/* Returns a newly allocated url with every set field appended as a query
 * parameter - the caller must free() it (url itself is never modified or
 * freed). */
static inline char* GetStateActionQueryParams_apply(const GetStateActionQueryParams* self, const char* url) {
    char* out = strdup(url);
    if (!self) return out;
    if (self->verbose_set) {
        char value_buf[64];
        snprintf(value_buf, sizeof(value_buf), "%d", self->verbose);
        char* next = fetchx_url_append_query(out, "verbose", value_buf);
        free(out);
        out = next;
    }
    if (self->prefix) {
        char* next = fetchx_url_append_query(out, "prefix", self->prefix);
        free(out);
        out = next;
    }
    return out;
}
/* GetStateActionResponse */
typedef struct GetStateActionResponse {
    float roomSetPoint;
    float inside;
} GetStateActionResponse;
/* Forward declarations: a self-referencing field (e.g. a relation whose
 * target: is this very struct) means _free_fields/_from_json_into below
 * can call _free/_from_json, which aren't defined until further down this
 * same file - static inline functions can be forward-declared exactly like
 * any other, so this resolves the ordering regardless of which fields end
 * up actually self-referencing. */
static inline void GetStateActionResponse_init(GetStateActionResponse* self);
static inline void GetStateActionResponse_free_fields(GetStateActionResponse* self);
static inline GetStateActionResponse* GetStateActionResponse_new(void);
static inline void GetStateActionResponse_free(GetStateActionResponse* self);
static inline void GetStateActionResponse_from_json_into(GetStateActionResponse* self, const cJSON* json);
static inline GetStateActionResponse* GetStateActionResponse_from_json(const cJSON* json);
static inline cJSON* GetStateActionResponse_to_json(const GetStateActionResponse* self);
static inline void GetStateActionResponse_init(GetStateActionResponse* self) {
    if (!self) return;
    self->roomSetPoint = 0;
    self->inside = 0;
}
static inline void GetStateActionResponse_free_fields(GetStateActionResponse* self) {
    if (!self) return;
}
static inline GetStateActionResponse* GetStateActionResponse_new(void) {
    GetStateActionResponse* self = (GetStateActionResponse*) malloc(sizeof(GetStateActionResponse));
    GetStateActionResponse_init(self);
    return self;
}
static inline void GetStateActionResponse_free(GetStateActionResponse* self) {
    if (!self) return;
    GetStateActionResponse_free_fields(self);
    free(self);
}
static inline void GetStateActionResponse_from_json_into(GetStateActionResponse* self, const cJSON* json) {
    if (!self || !json) return;
    {
        const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, "roomSetPoint");
        if (cJSON_IsNumber(v)) { self->roomSetPoint = (float) v->valuedouble; }
        else if (cJSON_IsBool(v)) { self->roomSetPoint = (float) cJSON_IsTrue(v); }
    }
    {
        const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, "inside");
        if (cJSON_IsNumber(v)) { self->inside = (float) v->valuedouble; }
        else if (cJSON_IsBool(v)) { self->inside = (float) cJSON_IsTrue(v); }
    }
}
static inline GetStateActionResponse* GetStateActionResponse_from_json(const cJSON* json) {
    GetStateActionResponse* self = GetStateActionResponse_new();
    GetStateActionResponse_from_json_into(self, json);
    return self;
}
static inline cJSON* GetStateActionResponse_to_json(const GetStateActionResponse* self) {
    cJSON* json = cJSON_CreateObject();
    if (!self) return json;
    {
        char float_buf[32];
        snprintf(float_buf, sizeof(float_buf), "%.7g", (double) self->roomSetPoint);
        cJSON_AddNumberToObject(json, "roomSetPoint", strtod(float_buf, NULL));
    }
    {
        char float_buf[32];
        snprintf(float_buf, sizeof(float_buf), "%.7g", (double) self->inside);
        cJSON_AddNumberToObject(json, "inside", strtod(float_buf, NULL));
    }
    return json;
}
/* Serves GetStateAction (GET /api/state).
 *
 * Implement GetStateActionHandler, then register it once the http server is running:
 *
 *   GetStateActionBinding* b = get_state_action_register(server, my_handler, my_user_data);
 *
 * The handler returns a newly allocated GetStateActionResponse* (freed for you with GetStateActionResponse_free),
 * or NULL. *status starts at 0: leave it for 200 (204 when NULL is returned)
 * or set any code - with NULL and status >= 400 an error body is sent. */
typedef struct GetStateActionCall {
    emihttpd_req_t* req; /* the backend's request: headers, connection, raw response if needed */
    const GetStateActionQueryParams* query;
    void* user;
} GetStateActionCall;
typedef GetStateActionResponse* (*GetStateActionHandler)(const GetStateActionCall* call, int* status);
typedef struct GetStateActionBinding {
    GetStateActionHandler handler;
    void* user;
    void* route;
} GetStateActionBinding;
static inline void GetStateActionQueryParams_parse(GetStateActionQueryParams* self, const char* qs) {
    char* v;
    v = emihttpd_query_get(qs, "verbose");
    if (v) {
        self->verbose_set = true;
        self->verbose = emihttpd_parse_bool(v);
        free(v);
    }
    v = emihttpd_query_get(qs, "prefix");
    if (v) {
        free(self->prefix);
        self->prefix = v;
    }
}
static inline int get_state_action_serve(emihttpd_req_t* req, void* ctx) {
    GetStateActionBinding* binding = (GetStateActionBinding*) ctx;
    GetStateActionCall call;
    memset(&call, 0, sizeof(call));
    call.req = req;
    call.user = binding->user;
    int rc = 0;
    int status = 0;
    GetStateActionQueryParams query;
    GetStateActionQueryParams_init(&query);
    {
        char* qs = emihttpd_req_query(req);
        GetStateActionQueryParams_parse(&query, qs);
        free(qs);
    }
    call.query = &query;
    {
        GetStateActionResponse* res = binding->handler(&call, &status);
        if (!res) {
            rc = (status >= 400) ? emihttpd_reply_error(req, status, "request failed") : emihttpd_reply_json(req, 204, NULL);
        } else {
            cJSON* out = GetStateActionResponse_to_json(res);
            char* text = out ? cJSON_PrintUnformatted(out) : NULL;
            cJSON_Delete(out);
            rc = text ? emihttpd_reply_json(req, status ? status : 200, text) : emihttpd_reply_error(req, 500, "cannot serialize response");
            free(text);
            GetStateActionResponse_free(res);
        }
    }
    GetStateActionQueryParams_free_fields(&query);
    return rc;
}
/* Registers the route on the backend's server (esp_http_server handle / Mongoose router).
 * With esp_http_server set config.uri_match_fn = emihttpd_uri_match when the url has :params.
 * Returns the binding (free it with get_state_action_unregister) or NULL on failure. */
static inline GetStateActionBinding* get_state_action_register(emihttpd_server_t server, GetStateActionHandler handler, void* user) {
    GetStateActionBinding* binding = (GetStateActionBinding*) calloc(1, sizeof(*binding));
    if (!binding) return NULL;
    binding->handler = handler;
    binding->user = user;
    binding->route = emihttpd_add_route(server, EMIHTTPD_GET, "/api/state", get_state_action_serve, binding);
    if (!binding->route) {
        free(binding);
        return NULL;
    }
    return binding;
}
static inline void get_state_action_unregister(emihttpd_server_t server, GetStateActionBinding* binding) {
    if (!binding) return;
    emihttpd_remove_route(server, binding->route);
    free(binding);
}
#endif /* EMISDK_GET_STATE_ACTION_H */
