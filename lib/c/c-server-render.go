// Renders the ESP-IDF *server* flavour of an action (`--tags esp-server`):
// the same dto/path/query structs the client gets, plus typed handler
// glue on top of esp_http_server - a handler type the developer implements,
// a register function, and for `method: reactive` websocket sessions with
// per-client send and broadcast. Mirrors what lib/golang does for gin.
package c

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// serverUrlPath turns an action url into the path esp_http_server routes on:
// a scheme/host prefix (`ws://localhost:8081/x`) is dropped, a missing leading
// slash is added and any query part is cut off.
func serverUrlPath(url string) string {
	if i := strings.Index(url, "://"); i >= 0 {
		rest := url[i+3:]
		if j := strings.Index(rest, "/"); j >= 0 {
			url = rest[j:]
		} else {
			url = "/"
		}
	}
	if i := strings.IndexAny(url, "?#"); i >= 0 {
		url = url[:i]
	}
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}
	return url
}

// serverBackendHeader is the runtime header of the selected server backend;
// "" when the compiler is not in server mode.
func serverBackendHeader(ctx core.MicroGenContext) string {
	switch {
	case ctx.HasTag(MongooseServer):
		return "emihttpd_mongoose.h"
	case ctx.HasTag(EspServer):
		return "emihttpd_esp.h"
	}
	return ""
}

func serverMethod(method string, hasBody bool) (string, error) {
	switch method {
	case "":
		if hasBody {
			return "EMIHTTPD_POST", nil
		}
		return "EMIHTTPD_GET", nil
	case "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS":
		return "EMIHTTPD_" + method, nil
	}
	return "", fmt.Errorf("server mode: http method %q is not supported", method)
}

// cNumericParse returns the C expression turning the C string `src` into the
// scalar type `ctype`.
func cNumericParse(ctype, src string) string {
	switch ctype {
	case "float", "double":
		return fmt.Sprintf("(%s) strtod(%s, NULL)", ctype, src)
	case "bool":
		return fmt.Sprintf("emihttpd_parse_bool(%s)", src)
	default:
		return fmt.Sprintf("(%s) strtoll(%s, NULL, 10)", ctype, src)
	}
}

type cServerQueryField struct {
	cQueryField
	Parse string
}

type cServerPathField struct {
	cPathParam
	Parse string
}

const cServerHttpTmpl = `
{{ if .realms.PathParameter }}
{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}
{{ if .realms.QueryStruct }}
{{ b2s .realms.QueryStruct.ActualScript }}
{{ end }}
{{ if .realms.RequestStruct }}
{{ b2s .realms.RequestStruct.ActualScript }}
{{ end }}
{{ if .realms.ResponseStruct }}
{{ b2s .realms.ResponseStruct.ActualScript }}
{{ end }}
{{ $A := .realms.ActionName }}
{{ $fn := .realms.FunctionName }}
/* Serves {{ $A }} ({{ .realms.HttpMethod }} {{ .path }}).
 *
 * Implement {{ $A }}Handler, then register it once the http server is running:
 *
 *   {{ $A }}Binding* b = {{ $fn }}_register(server, my_handler, my_user_data);
 *
 * The handler returns {{ if .returnsTyped }}a newly allocated {{ .responseType }}* (freed for you with {{ .responseType }}_free){{ else }}a malloc'd JSON string (freed for you){{ end }},
 * or NULL. *status starts at 0: leave it for 200 (204 when NULL is returned)
 * or set any code - with NULL and status >= 400 an error body is sent. */
typedef struct {{ $A }}Call {
    emihttpd_req_t* req; /* the backend's request: headers, connection, raw response if needed */
{{- if .realms.PathParameter }}
    const {{ $A }}PathParameters* path;
{{- end }}
{{- if .realms.QueryStruct }}
    const {{ $A }}QueryParams* query;
{{- end }}
{{- if .hasBody }}
    const {{ .requestType }}* body;
{{- end }}
    void* user;
} {{ $A }}Call;

typedef {{ if .returnsTyped }}{{ .responseType }}*{{ else }}char*{{ end }} (*{{ $A }}Handler)(const {{ $A }}Call* call, int* status);

typedef struct {{ $A }}Binding {
    {{ $A }}Handler handler;
    void* user;
    void* route;
} {{ $A }}Binding;

{{ if .realms.PathParameter }}
/* Takes ownership of the strings captured by emihttpd_path_capture. */
static inline void {{ $A }}PathParameters_from_captures({{ $A }}PathParameters* self, char** caps, size_t n) {
{{- range $i, $p := .pathFields }}
    if (n > {{ $i }} && caps[{{ $i }}]) {
{{- if $p.IsString }}
        free(self->{{ $p.Field }});
        self->{{ $p.Field }} = caps[{{ $i }}];
        caps[{{ $i }}] = NULL;
{{- else }}
        self->{{ $p.Field }} = {{ $p.Parse }};
{{- end }}
    }
{{- end }}
    for (size_t i = 0; i < n; i++) free(caps[i]);
}
{{ end }}
{{ if .realms.QueryStruct }}
static inline void {{ $A }}QueryParams_parse({{ $A }}QueryParams* self, const char* qs) {
    char* v;
{{- range .queryFields }}
    v = emihttpd_query_get(qs, "{{ .Key }}");
    if (v) {
{{- if .IsString }}
        free(self->{{ .Path }});
        self->{{ .Path }} = v;
{{- else }}
        self->{{ .Path }}_set = true;
        self->{{ .Path }} = {{ .Parse }};
        free(v);
{{- end }}
    }
{{- end }}
}
{{ end }}

static inline int {{ $fn }}_serve(emihttpd_req_t* req, void* ctx) {
    {{ $A }}Binding* binding = ({{ $A }}Binding*) ctx;
    {{ $A }}Call call;
    memset(&call, 0, sizeof(call));
    call.req = req;
    call.user = binding->user;
    int rc = 0;
    int status = 0;
{{- if .realms.PathParameter }}
    {{ $A }}PathParameters path;
    {{ $A }}PathParameters_init(&path);
    {
        char* caps[EMIHTTPD_MAX_PARAMS] = {0};
        size_t ncaps = 0;
        if (!emihttpd_req_capture(req, {{ .path | cString }}, caps, EMIHTTPD_MAX_PARAMS, &ncaps)) {
            return emihttpd_reply_error(req, 404, "not found");
        }
        {{ $A }}PathParameters_from_captures(&path, caps, ncaps);
    }
    call.path = &path;
{{- end }}
{{- if .realms.QueryStruct }}
    {{ $A }}QueryParams query;
    {{ $A }}QueryParams_init(&query);
    {
        char* qs = emihttpd_req_query(req);
        {{ $A }}QueryParams_parse(&query, qs);
        free(qs);
    }
    call.query = &query;
{{- end }}
{{- if .hasBody }}
    {{ .requestType }}* body = NULL;
    {
        char* raw = NULL;
        int br = emihttpd_req_body(req, &raw);
        if (br != 0) {
            rc = emihttpd_reply_error(req, br == -1 ? 413 : 400, br == -1 ? "body too large" : "cannot read body");
            goto done;
        }
        if (raw[0] == '\0') {
            body = {{ .requestType }}_new();
        } else {
            cJSON* root = cJSON_Parse(raw);
            if (root) body = {{ .requestType }}_from_json(root);
            cJSON_Delete(root);
        }
        free(raw);
        if (!body) {
            rc = emihttpd_reply_error(req, 400, "invalid json body");
            goto done;
        }
    }
    call.body = body;
{{- end }}

    {
{{- if .returnsTyped }}
        {{ .responseType }}* res = binding->handler(&call, &status);
        if (!res) {
            rc = (status >= 400) ? emihttpd_reply_error(req, status, emihttpd_take_error("request failed")) : emihttpd_reply_json(req, 204, NULL);
        } else {
            cJSON* out = {{ .responseType }}_to_json(res);
            char* text = out ? cJSON_PrintUnformatted(out) : NULL;
            cJSON_Delete(out);
            rc = text ? emihttpd_reply_json(req, status ? status : 200, text) : emihttpd_reply_error(req, 500, "cannot serialize response");
            free(text);
            {{ .responseType }}_free(res);
        }
{{- else }}
        char* text = binding->handler(&call, &status);
        if (!text) {
            rc = (status >= 400) ? emihttpd_reply_error(req, status, emihttpd_take_error("request failed")) : emihttpd_reply_json(req, 204, NULL);
        } else {
            rc = emihttpd_reply_json(req, status ? status : 200, text);
            free(text);
        }
{{- end }}
    }
{{ if .hasBody }}
done:
    {{ .requestType }}_free(body);
{{- end }}
{{- if .realms.PathParameter }}
    {{ $A }}PathParameters_free_fields(&path);
{{- end }}
{{- if .realms.QueryStruct }}
    {{ $A }}QueryParams_free_fields(&query);
{{- end }}
    return rc;
}

/* Registers the route on the backend's server (esp_http_server handle / Mongoose router).
 * With esp_http_server set config.uri_match_fn = emihttpd_uri_match when the url has :params.
 * Returns the binding (free it with {{ $fn }}_unregister) or NULL on failure. */
static inline {{ $A }}Binding* {{ $fn }}_register(emihttpd_server_t server, {{ $A }}Handler handler, void* user) {
    {{ $A }}Binding* binding = ({{ $A }}Binding*) calloc(1, sizeof(*binding));
    if (!binding) return NULL;
    binding->handler = handler;
    binding->user = user;
    binding->route = emihttpd_add_route(server, {{ .espMethod }}, {{ .path | cString }}, {{ $fn }}_serve, binding);
    if (!binding->route) {
        free(binding);
        return NULL;
    }
    return binding;
}

static inline void {{ $fn }}_unregister(emihttpd_server_t server, {{ $A }}Binding* binding) {
    if (!binding) return;
    emihttpd_remove_route(server, binding->route);
    free(binding);
}
`

const cServerReactiveTmpl = `
#ifdef EMIHTTPD_NO_WS
#error "{{ .realms.ActionName }} is a reactive action: enable CONFIG_HTTPD_WS_SUPPORT (menuconfig -> Component config -> HTTP Server -> WebSocket server support)"
#endif
{{ if .realms.PathParameter }}
{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}
{{ if .realms.QueryStruct }}
{{ b2s .realms.QueryStruct.ActualScript }}
{{ end }}
{{ if .realms.RequestStruct }}
{{ b2s .realms.RequestStruct.ActualScript }}
{{ end }}
{{ if .realms.ResponseStruct }}
{{ b2s .realms.ResponseStruct.ActualScript }}
{{ end }}
{{ $A := .realms.ActionName }}
{{ $fn := .realms.FunctionName }}
/* Serves the {{ $A }} websocket ({{ .path }}).
 *
 * Fill {{ $A }}Callbacks, then register it once the http server is running:
 *
 *   {{ $A }}Binding* b = {{ $fn }}_register(server, &callbacks);
 *   ...
 *   {{ $fn }}_broadcast(b, &update);   // push to every connected client, from any task
 *
 * With esp_http_server set config.uri_match_fn = emihttpd_uri_match when the url has :params.
 * Clients that vanish are dropped (on_close fires): on the next failed send with
 * esp_http_server, when the connection closes with Mongoose. */
typedef struct {{ $A }}ClientState {
    emihttpd_conn_t conn;
{{- if .realms.PathParameter }}
    {{ $A }}PathParameters path;
{{- end }}
{{- if .realms.QueryStruct }}
    {{ $A }}QueryParams query;
{{- end }}
} {{ $A }}ClientState;

typedef struct {{ $A }}Session {
    emihttpd_server_t server;
    emihttpd_conn_t conn; /* use with {{ $fn }}_send*() */
{{- if .realms.PathParameter }}
    const {{ $A }}PathParameters* path;
{{- end }}
{{- if .realms.QueryStruct }}
    const {{ $A }}QueryParams* query;
{{- end }}
    void* user;
} {{ $A }}Session;

typedef struct {{ $A }}Callbacks {
    /* Called after the handshake; return non-zero to reject the client. */
    int (*on_open)(const {{ $A }}Session* session);
{{- if .hasBody }}
    /* A text frame, parsed as {{ .requestType }}; msg is freed after the call. */
    void (*on_message)(const {{ $A }}Session* session, const {{ .requestType }}* msg);
{{- else if not .binary }}
    /* A text frame, as received (NUL terminated). */
    void (*on_text)(const {{ $A }}Session* session, const char* text);
{{- end }}
{{- if .binary }}
    /* A binary frame; data is only valid during the call. */
    void (*on_binary)(const {{ $A }}Session* session, const uint8_t* data, size_t len);
{{- end }}
    /* The client closed or was dropped. */
    void (*on_close)(const {{ $A }}Session* session);
    void* user;
} {{ $A }}Callbacks;

typedef struct {{ $A }}Binding {
    {{ $A }}Callbacks cb;
    EmiHttpdClients clients;
    emihttpd_server_t server;
    void* route;
} {{ $A }}Binding;

{{ if .realms.PathParameter }}
/* Takes ownership of the strings captured by emihttpd_path_capture. */
static inline void {{ $A }}PathParameters_from_captures({{ $A }}PathParameters* self, char** caps, size_t n) {
{{- range $i, $p := .pathFields }}
    if (n > {{ $i }} && caps[{{ $i }}]) {
{{- if $p.IsString }}
        free(self->{{ $p.Field }});
        self->{{ $p.Field }} = caps[{{ $i }}];
        caps[{{ $i }}] = NULL;
{{- else }}
        self->{{ $p.Field }} = {{ $p.Parse }};
{{- end }}
    }
{{- end }}
    for (size_t i = 0; i < n; i++) free(caps[i]);
}
{{ end }}
{{ if .realms.QueryStruct }}
static inline void {{ $A }}QueryParams_parse({{ $A }}QueryParams* self, const char* qs) {
    char* v;
{{- range .queryFields }}
    v = emihttpd_query_get(qs, "{{ .Key }}");
    if (v) {
{{- if .IsString }}
        free(self->{{ .Path }});
        self->{{ .Path }} = v;
{{- else }}
        self->{{ .Path }}_set = true;
        self->{{ .Path }} = {{ .Parse }};
        free(v);
{{- end }}
    }
{{- end }}
}
{{ end }}

static inline void {{ $fn }}_fill_session({{ $A }}Binding* b, {{ $A }}ClientState* st, {{ $A }}Session* s) {
    s->server = b->server;
    s->conn = st->conn;
{{- if .realms.PathParameter }}
    s->path = &st->path;
{{- end }}
{{- if .realms.QueryStruct }}
    s->query = &st->query;
{{- end }}
    s->user = b->cb.user;
}

static inline void {{ $fn }}_free_state({{ $A }}ClientState* st) {
    if (!st) return;
{{- if .realms.PathParameter }}
    {{ $A }}PathParameters_free_fields(&st->path);
{{- end }}
{{- if .realms.QueryStruct }}
    {{ $A }}QueryParams_free_fields(&st->query);
{{- end }}
    free(st);
}

/* Forgets a client (calls on_close). Safe to call for unknown connections. */
static inline void {{ $fn }}_drop_client({{ $A }}Binding* b, emihttpd_conn_t conn) {
    {{ $A }}ClientState* st = ({{ $A }}ClientState*) emihttpd_clients_remove(&b->clients, conn);
    if (!st) return;
    if (b->cb.on_close) {
        {{ $A }}Session s;
        {{ $fn }}_fill_session(b, st, &s);
        b->cb.on_close(&s);
    }
    {{ $fn }}_free_state(st);
}

/* Number of clients currently connected. */
static inline size_t {{ $fn }}_client_count({{ $A }}Binding* b) {
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    return emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
}

static inline int {{ $fn }}_write({{ $A }}Binding* b, emihttpd_conn_t conn, bool binary, const uint8_t* data, size_t len) {
    emihttpd_mutex_lock(&b->clients.send_lock);
    int rc = emihttpd_ws_write(b->server, conn, binary, data, len);
    emihttpd_mutex_unlock(&b->clients.send_lock);
    if (rc != 0) { {{ $fn }}_drop_client(b, conn); return -1; }
    return 0;
}

/* Sends a text frame to one client; 0 on success (queued, with Mongoose), the client is dropped on failure. */
static inline int {{ $fn }}_send_text({{ $A }}Binding* b, emihttpd_conn_t conn, const char* text) {
    return {{ $fn }}_write(b, conn, false, (const uint8_t*) text, strlen(text));
}

/* Sends a text frame to every client; returns how many received it. */
static inline int {{ $fn }}_broadcast_text({{ $A }}Binding* b, const char* text) {
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    size_t n = emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
    int delivered = 0;
    for (size_t i = 0; i < n; i++) {
        if ({{ $fn }}_send_text(b, conns[i], text) == 0) delivered++;
    }
    return delivered;
}
{{ if .returnsTyped }}
/* Sends a {{ .responseType }} as a JSON text frame to one client. */
static inline int {{ $fn }}_send({{ $A }}Binding* b, emihttpd_conn_t conn, const {{ .responseType }}* res) {
    cJSON* json = {{ .responseType }}_to_json(res);
    char* text = json ? cJSON_PrintUnformatted(json) : NULL;
    cJSON_Delete(json);
    if (!text) return -1;
    int rc = {{ $fn }}_send_text(b, conn, text);
    free(text);
    return rc;
}

/* Sends a {{ .responseType }} to every client (serialized once); returns how many received it. */
static inline int {{ $fn }}_broadcast({{ $A }}Binding* b, const {{ .responseType }}* res) {
    cJSON* json = {{ .responseType }}_to_json(res);
    char* text = json ? cJSON_PrintUnformatted(json) : NULL;
    cJSON_Delete(json);
    if (!text) return 0;
    int delivered = {{ $fn }}_broadcast_text(b, text);
    free(text);
    return delivered;
}
{{ end }}
{{ if .binary }}
/* Sends a binary frame to one client; 0 on success, the client is dropped on failure. */
static inline int {{ $fn }}_send_binary({{ $A }}Binding* b, emihttpd_conn_t conn, const uint8_t* data, size_t len) {
    return {{ $fn }}_write(b, conn, true, data, len);
}

/* Sends a binary frame to every client; returns how many received it. */
static inline int {{ $fn }}_broadcast_binary({{ $A }}Binding* b, const uint8_t* data, size_t len) {
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    size_t n = emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
    int delivered = 0;
    for (size_t i = 0; i < n; i++) {
        if ({{ $fn }}_send_binary(b, conns[i], data, len) == 0) delivered++;
    }
    return delivered;
}
{{ end }}

/* The handshake: decode path/query, remember the client, ask on_open. */
static inline int {{ $fn }}_ws_open(void* request, emihttpd_conn_t conn, void* ctx) {
    {{ $A }}Binding* b = ({{ $A }}Binding*) ctx;
    emihttpd_req_t* req = (emihttpd_req_t*) request;
    {{ $A }}ClientState* st = ({{ $A }}ClientState*) calloc(1, sizeof(*st));
    if (!st) return -1;
    st->conn = conn;
{{- if .realms.PathParameter }}
    {{ $A }}PathParameters_init(&st->path);
    {
        char* caps[EMIHTTPD_MAX_PARAMS] = {0};
        size_t ncaps = 0;
        if (!emihttpd_req_capture(req, {{ .path | cString }}, caps, EMIHTTPD_MAX_PARAMS, &ncaps)) {
            {{ $fn }}_free_state(st);
            return -1;
        }
        {{ $A }}PathParameters_from_captures(&st->path, caps, ncaps);
    }
{{- end }}
{{- if .realms.QueryStruct }}
    {{ $A }}QueryParams_init(&st->query);
    {
        char* qs = emihttpd_req_query(req);
        {{ $A }}QueryParams_parse(&st->query, qs);
        free(qs);
    }
{{- end }}
{{- if not (or .realms.PathParameter .realms.QueryStruct) }}
    (void) req;
{{- end }}
    {{ $fn }}_drop_client(b, conn); /* a stale entry for a reused connection id */
    if (emihttpd_clients_add(&b->clients, conn, st) != 0) {
        {{ $fn }}_free_state(st);
        return -1;
    }
    if (b->cb.on_open) {
        {{ $A }}Session s;
        {{ $fn }}_fill_session(b, st, &s);
        if (b->cb.on_open(&s) != 0) {
            emihttpd_clients_remove(&b->clients, conn);
            {{ $fn }}_free_state(st);
            return -1;
        }
    }
    return 0;
}

/* One inbound frame; data is NUL terminated. */
static inline void {{ $fn }}_ws_message(emihttpd_conn_t conn, bool binary, const uint8_t* data, size_t len, void* ctx) {
    {{ $A }}Binding* b = ({{ $A }}Binding*) ctx;
    {{ $A }}ClientState* st = ({{ $A }}ClientState*) emihttpd_clients_find(&b->clients, conn);
    if (!st) return;
{{- if .binary }}
    if (binary) {
        if (b->cb.on_binary) {
            {{ $A }}Session s;
            {{ $fn }}_fill_session(b, st, &s);
            b->cb.on_binary(&s, data, len);
        }
        return;
    }
{{- else }}
    if (binary) return;
{{- end }}
    (void) len;
{{- if .hasBody }}
    if (b->cb.on_message) {
        cJSON* root = cJSON_Parse((const char*) data);
        {{ .requestType }}* msg = root ? {{ .requestType }}_from_json(root) : NULL;
        cJSON_Delete(root);
        if (msg) {
            {{ $A }}Session s;
            {{ $fn }}_fill_session(b, st, &s);
            b->cb.on_message(&s, msg);
            {{ .requestType }}_free(msg);
        }
    }
{{- else if not .binary }}
    if (b->cb.on_text) {
        {{ $A }}Session s;
        {{ $fn }}_fill_session(b, st, &s);
        b->cb.on_text(&s, (const char*) data);
    }
{{- end }}
}

static inline void {{ $fn }}_ws_close(emihttpd_conn_t conn, void* ctx) {
    {{ $fn }}_drop_client(({{ $A }}Binding*) ctx, conn);
}

/* Registers the websocket route; returns the binding or NULL on failure. */
static inline {{ $A }}Binding* {{ $fn }}_register(emihttpd_server_t server, const {{ $A }}Callbacks* callbacks) {
    {{ $A }}Binding* b = ({{ $A }}Binding*) calloc(1, sizeof(*b));
    if (!b) return NULL;
    if (callbacks) b->cb = *callbacks;
    b->server = server;
    if (emihttpd_clients_init(&b->clients) != 0) {
        free(b);
        return NULL;
    }
    EmiHttpdWsHooks hooks = {{ "{" }}{{ $fn }}_ws_open, {{ $fn }}_ws_message, {{ $fn }}_ws_close{{ "}" }};
    b->route = emihttpd_add_ws_route(server, {{ .path | cString }}, &hooks, b);
    if (!b->route) {
        emihttpd_clients_deinit(&b->clients);
        free(b);
        return NULL;
    }
    return b;
}

/* Unregisters the route, drops every client (on_close is called) and frees the binding. */
static inline void {{ $fn }}_unregister(emihttpd_server_t server, {{ $A }}Binding* b) {
    if (!b) return;
    emihttpd_conn_t conns[EMIHTTPD_MAX_CLIENTS];
    emihttpd_remove_route(server, b->route);
    size_t n = emihttpd_clients_snapshot(&b->clients, conns, EMIHTTPD_MAX_CLIENTS);
    for (size_t i = 0; i < n; i++) {
        {{ $fn }}_drop_client(b, conns[i]);
    }
    emihttpd_clients_deinit(&b->clients);
    free(b);
}
`

var cServerHttpTemplate = template.Must(template.New("cserverhttp").Funcs(core.CommonMap).Funcs(template.FuncMap{"cString": escapeDoubleQuoted}).Parse(cServerHttpTmpl))
var cServerReactiveTemplate = template.Must(template.New("cserverws").Funcs(core.CommonMap).Funcs(template.FuncMap{"cString": escapeDoubleQuoted}).Parse(cServerReactiveTmpl))

// cActionRenderServer renders one action as ESP-IDF server code, see the
// package comment above. Remotes (outbound calls to somebody else's server)
// are not served and never reach this function.
func cActionRenderServer(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	realms cActionRealms,
) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: realms.ActionName},
			{Name: TOKEN_ROOT_CLASS, Value: realms.FunctionName},
		},
	}

	// the backend header first: it defines FETCHX_NO_CURL before fetchx.h is
	// seen, and fetchx.h still provides the url helpers the path/query structs use.
	res.CodeChunkDependensies = append(res.CodeChunkDependensies,
		core.CodeChunkDependency{Location: serverBackendHeader(ctx)},
		core.CodeChunkDependency{Location: "fetchx.h"},
	)
	if realms.RequestStruct != nil {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, realms.RequestStruct.CodeChunkDependensies...)
	} else if realms.RequestTypeName != "" {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, HeaderIncludeDependency(realms.RequestTypeName)...)
	}
	if realms.ResponseStruct != nil {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, realms.ResponseStruct.CodeChunkDependensies...)
	} else if realms.ResponseTypeName != "" {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, HeaderIncludeDependency(realms.ResponseTypeName)...)
	}

	hasBody := realms.RequestTypeName != "" && !realms.RequestIsAmbiguous
	requestType := realms.RequestTypeName
	responseType := realms.ResponseTypeName
	returnsTyped := realms.HasResponseShape && responseType != "" && !realms.ResponseIsAmbiguous

	if !realms.HasUrl {
		return nil, fmt.Errorf("esp-server: action %q has no url to serve", realms.ActionName)
	}
	path := serverUrlPath(realms.Url)

	pathFields := []cServerPathField{}
	for _, p := range cBuildPathParams(action) {
		// the capture at the placeholder's position, see the template
		pathFields = append(pathFields, cServerPathField{cPathParam: p, Parse: cNumericParse(p.Type, fmt.Sprintf("caps[%d]", len(pathFields)))})
	}

	queryFields := []cServerQueryField{}
	_, leaves := cQueryTree(action.GetQuery())
	for _, l := range leaves {
		queryFields = append(queryFields, cServerQueryField{cQueryField: *l, Parse: cNumericParse(l.Type, "v")})
	}

	data := core.H{
		"realms":       realms,
		"path":         path,
		"hasBody":      hasBody,
		"requestType":  requestType,
		"responseType": responseType,
		"returnsTyped": returnsTyped,
		"pathFields":   pathFields,
		"queryFields":  queryFields,
		"binary":       realms.IsBinaryReactive,
	}

	tmpl := cServerHttpTemplate
	if realms.IsReactive {
		tmpl = cServerReactiveTemplate
	} else {
		method, err := serverMethod(realms.HttpMethod, hasBody)
		if err != nil {
			return nil, err
		}
		data["espMethod"] = method
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToSnakeCase(realms.ActionName)
	res.SuggestedExtension = ".h"
	res.Realms = realms
	return res, nil
}
