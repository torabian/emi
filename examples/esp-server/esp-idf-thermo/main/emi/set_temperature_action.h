#ifndef EMISDK_SET_TEMPERATURE_ACTION_H
#define EMISDK_SET_TEMPERATURE_ACTION_H
#include "cjson/cJSON.h"
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "emihttpd_esp.h"
#include "fetchx.h"
typedef struct SetTemperatureActionPathParameters {
    int32_t zone;
} SetTemperatureActionPathParameters;
static inline void SetTemperatureActionPathParameters_init(SetTemperatureActionPathParameters* self) {
    if (!self) return;
    self->zone = (int32_t) 0;
}
static inline void SetTemperatureActionPathParameters_free_fields(SetTemperatureActionPathParameters* self) {
    if (!self) return;
}
/* Returns a newly allocated url - the caller must free() it. */
static inline char* SetTemperatureActionPathParameters_apply(const SetTemperatureActionPathParameters* self, const char* template_url) {
    char* url = strdup(template_url);
    {
        char value_buf[64];
        snprintf(value_buf, sizeof(value_buf), "%lld", (long long) self->zone);
        char* next = fetchx_url_replace(url, ":zone", value_buf);
        free(url);
        url = next;
    }
    return url;
}
/* SetTemperatureActionRequest */
typedef struct SetTemperatureActionRequest {
    float target;
} SetTemperatureActionRequest;
/* Forward declarations: a self-referencing field (e.g. a relation whose
 * target: is this very struct) means _free_fields/_from_json_into below
 * can call _free/_from_json, which aren't defined until further down this
 * same file - static inline functions can be forward-declared exactly like
 * any other, so this resolves the ordering regardless of which fields end
 * up actually self-referencing. */
static inline void SetTemperatureActionRequest_init(SetTemperatureActionRequest* self);
static inline void SetTemperatureActionRequest_free_fields(SetTemperatureActionRequest* self);
static inline SetTemperatureActionRequest* SetTemperatureActionRequest_new(void);
static inline void SetTemperatureActionRequest_free(SetTemperatureActionRequest* self);
static inline void SetTemperatureActionRequest_from_json_into(SetTemperatureActionRequest* self, const cJSON* json);
static inline SetTemperatureActionRequest* SetTemperatureActionRequest_from_json(const cJSON* json);
static inline cJSON* SetTemperatureActionRequest_to_json(const SetTemperatureActionRequest* self);
static inline void SetTemperatureActionRequest_init(SetTemperatureActionRequest* self) {
    if (!self) return;
    self->target = 0;
}
static inline void SetTemperatureActionRequest_free_fields(SetTemperatureActionRequest* self) {
    if (!self) return;
}
static inline SetTemperatureActionRequest* SetTemperatureActionRequest_new(void) {
    SetTemperatureActionRequest* self = (SetTemperatureActionRequest*) malloc(sizeof(SetTemperatureActionRequest));
    SetTemperatureActionRequest_init(self);
    return self;
}
static inline void SetTemperatureActionRequest_free(SetTemperatureActionRequest* self) {
    if (!self) return;
    SetTemperatureActionRequest_free_fields(self);
    free(self);
}
static inline void SetTemperatureActionRequest_from_json_into(SetTemperatureActionRequest* self, const cJSON* json) {
    if (!self || !json) return;
    {
        const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, "target");
        if (cJSON_IsNumber(v)) { self->target = (float) v->valuedouble; }
        else if (cJSON_IsBool(v)) { self->target = (float) cJSON_IsTrue(v); }
    }
}
static inline SetTemperatureActionRequest* SetTemperatureActionRequest_from_json(const cJSON* json) {
    SetTemperatureActionRequest* self = SetTemperatureActionRequest_new();
    SetTemperatureActionRequest_from_json_into(self, json);
    return self;
}
static inline cJSON* SetTemperatureActionRequest_to_json(const SetTemperatureActionRequest* self) {
    cJSON* json = cJSON_CreateObject();
    if (!self) return json;
    {
        char float_buf[32];
        snprintf(float_buf, sizeof(float_buf), "%.7g", (double) self->target);
        cJSON_AddNumberToObject(json, "target", strtod(float_buf, NULL));
    }
    return json;
}
/* SetTemperatureActionResponse */
typedef struct SetTemperatureActionResponse {
    bool ok;
    float applied;
} SetTemperatureActionResponse;
/* Forward declarations: a self-referencing field (e.g. a relation whose
 * target: is this very struct) means _free_fields/_from_json_into below
 * can call _free/_from_json, which aren't defined until further down this
 * same file - static inline functions can be forward-declared exactly like
 * any other, so this resolves the ordering regardless of which fields end
 * up actually self-referencing. */
static inline void SetTemperatureActionResponse_init(SetTemperatureActionResponse* self);
static inline void SetTemperatureActionResponse_free_fields(SetTemperatureActionResponse* self);
static inline SetTemperatureActionResponse* SetTemperatureActionResponse_new(void);
static inline void SetTemperatureActionResponse_free(SetTemperatureActionResponse* self);
static inline void SetTemperatureActionResponse_from_json_into(SetTemperatureActionResponse* self, const cJSON* json);
static inline SetTemperatureActionResponse* SetTemperatureActionResponse_from_json(const cJSON* json);
static inline cJSON* SetTemperatureActionResponse_to_json(const SetTemperatureActionResponse* self);
static inline void SetTemperatureActionResponse_init(SetTemperatureActionResponse* self) {
    if (!self) return;
    self->ok = false;
    self->applied = 0;
}
static inline void SetTemperatureActionResponse_free_fields(SetTemperatureActionResponse* self) {
    if (!self) return;
}
static inline SetTemperatureActionResponse* SetTemperatureActionResponse_new(void) {
    SetTemperatureActionResponse* self = (SetTemperatureActionResponse*) malloc(sizeof(SetTemperatureActionResponse));
    SetTemperatureActionResponse_init(self);
    return self;
}
static inline void SetTemperatureActionResponse_free(SetTemperatureActionResponse* self) {
    if (!self) return;
    SetTemperatureActionResponse_free_fields(self);
    free(self);
}
static inline void SetTemperatureActionResponse_from_json_into(SetTemperatureActionResponse* self, const cJSON* json) {
    if (!self || !json) return;
    {
        const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, "ok");
        if (cJSON_IsNumber(v)) { self->ok = (bool) v->valuedouble; }
        else if (cJSON_IsBool(v)) { self->ok = (bool) cJSON_IsTrue(v); }
    }
    {
        const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, "applied");
        if (cJSON_IsNumber(v)) { self->applied = (float) v->valuedouble; }
        else if (cJSON_IsBool(v)) { self->applied = (float) cJSON_IsTrue(v); }
    }
}
static inline SetTemperatureActionResponse* SetTemperatureActionResponse_from_json(const cJSON* json) {
    SetTemperatureActionResponse* self = SetTemperatureActionResponse_new();
    SetTemperatureActionResponse_from_json_into(self, json);
    return self;
}
static inline cJSON* SetTemperatureActionResponse_to_json(const SetTemperatureActionResponse* self) {
    cJSON* json = cJSON_CreateObject();
    if (!self) return json;
    cJSON_AddBoolToObject(json, "ok", self->ok);
    {
        char float_buf[32];
        snprintf(float_buf, sizeof(float_buf), "%.7g", (double) self->applied);
        cJSON_AddNumberToObject(json, "applied", strtod(float_buf, NULL));
    }
    return json;
}
/* Serves SetTemperatureAction (POST /api/ot/set-temp/:zone).
 *
 * Implement SetTemperatureActionHandler, then register it once the http server is running:
 *
 *   SetTemperatureActionBinding* b = set_temperature_action_register(server, my_handler, my_user_data);
 *
 * The handler returns a newly allocated SetTemperatureActionResponse* (freed for you with SetTemperatureActionResponse_free),
 * or NULL. *status starts at 0: leave it for 200 (204 when NULL is returned)
 * or set any code - with NULL and status >= 400 an error body is sent. */
typedef struct SetTemperatureActionCall {
    emihttpd_req_t* req; /* the backend's request: headers, connection, raw response if needed */
    const SetTemperatureActionPathParameters* path;
    const SetTemperatureActionRequest* body;
    void* user;
} SetTemperatureActionCall;
typedef SetTemperatureActionResponse* (*SetTemperatureActionHandler)(const SetTemperatureActionCall* call, int* status);
typedef struct SetTemperatureActionBinding {
    SetTemperatureActionHandler handler;
    void* user;
    void* route;
} SetTemperatureActionBinding;
/* Takes ownership of the strings captured by emihttpd_path_capture. */
static inline void SetTemperatureActionPathParameters_from_captures(SetTemperatureActionPathParameters* self, char** caps, size_t n) {
    if (n > 0 && caps[0]) {
        self->zone = (int32_t) strtoll(caps[0], NULL, 10);
    }
    for (size_t i = 0; i < n; i++) free(caps[i]);
}
static inline int set_temperature_action_serve(emihttpd_req_t* req, void* ctx) {
    SetTemperatureActionBinding* binding = (SetTemperatureActionBinding*) ctx;
    SetTemperatureActionCall call;
    memset(&call, 0, sizeof(call));
    call.req = req;
    call.user = binding->user;
    int rc = 0;
    int status = 0;
    SetTemperatureActionPathParameters path;
    SetTemperatureActionPathParameters_init(&path);
    {
        char* caps[EMIHTTPD_MAX_PARAMS] = {0};
        size_t ncaps = 0;
        if (!emihttpd_req_capture(req, "/api/ot/set-temp/:zone", caps, EMIHTTPD_MAX_PARAMS, &ncaps)) {
            return emihttpd_reply_error(req, 404, "not found");
        }
        SetTemperatureActionPathParameters_from_captures(&path, caps, ncaps);
    }
    call.path = &path;
    SetTemperatureActionRequest* body = NULL;
    {
        char* raw = NULL;
        int br = emihttpd_req_body(req, &raw);
        if (br != 0) {
            rc = emihttpd_reply_error(req, br == -1 ? 413 : 400, br == -1 ? "body too large" : "cannot read body");
            goto done;
        }
        if (raw[0] == '\0') {
            body = SetTemperatureActionRequest_new();
        } else {
            cJSON* root = cJSON_Parse(raw);
            if (root) body = SetTemperatureActionRequest_from_json(root);
            cJSON_Delete(root);
        }
        free(raw);
        if (!body) {
            rc = emihttpd_reply_error(req, 400, "invalid json body");
            goto done;
        }
    }
    call.body = body;
    {
        SetTemperatureActionResponse* res = binding->handler(&call, &status);
        if (!res) {
            rc = (status >= 400) ? emihttpd_reply_error(req, status, "request failed") : emihttpd_reply_json(req, 204, NULL);
        } else {
            cJSON* out = SetTemperatureActionResponse_to_json(res);
            char* text = out ? cJSON_PrintUnformatted(out) : NULL;
            cJSON_Delete(out);
            rc = text ? emihttpd_reply_json(req, status ? status : 200, text) : emihttpd_reply_error(req, 500, "cannot serialize response");
            free(text);
            SetTemperatureActionResponse_free(res);
        }
    }
done:
    SetTemperatureActionRequest_free(body);
    SetTemperatureActionPathParameters_free_fields(&path);
    return rc;
}
/* Registers the route on the backend's server (esp_http_server handle / Mongoose router).
 * With esp_http_server set config.uri_match_fn = emihttpd_uri_match when the url has :params.
 * Returns the binding (free it with set_temperature_action_unregister) or NULL on failure. */
static inline SetTemperatureActionBinding* set_temperature_action_register(emihttpd_server_t server, SetTemperatureActionHandler handler, void* user) {
    SetTemperatureActionBinding* binding = (SetTemperatureActionBinding*) calloc(1, sizeof(*binding));
    if (!binding) return NULL;
    binding->handler = handler;
    binding->user = user;
    binding->route = emihttpd_add_route(server, EMIHTTPD_POST, "/api/ot/set-temp/:zone", set_temperature_action_serve, binding);
    if (!binding->route) {
        free(binding);
        return NULL;
    }
    return binding;
}
static inline void set_temperature_action_unregister(emihttpd_server_t server, SetTemperatureActionBinding* binding) {
    if (!binding) return;
    emihttpd_remove_route(server, binding->route);
    free(binding);
}
#endif /* EMISDK_SET_TEMPERATURE_ACTION_H */
