package c

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// A just-enough fake of esp_http_server (+ the FreeRTOS mutex) so the code
// generated with --tags esp-server can run on the host: handlers are stored
// at registration, and the test program dispatches fake requests to them.
const fakeEspHttpServerH = `
#ifndef FAKE_ESP_HTTP_SERVER_H
#define FAKE_ESP_HTTP_SERVER_H
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#define CONFIG_HTTPD_WS_SUPPORT 1
typedef int esp_err_t;
#define ESP_OK 0
#define ESP_FAIL -1
#define ESP_ERR_NO_MEM 0x101
#define HTTPD_RESP_USE_STRLEN -1
#define HTTPD_SOCK_ERR_TIMEOUT -3

typedef void* httpd_handle_t;
typedef enum { HTTP_GET = 1, HTTP_POST = 3, HTTP_PUT = 4, HTTP_PATCH = 28, HTTP_DELETE = 0, HTTP_HEAD = 2, HTTP_OPTIONS = 6 } httpd_method_t;

typedef struct httpd_req {
    httpd_method_t method;
    char uri[256];
    size_t content_len;
    void* user_ctx;
    /* fake internals */
    const char* body;
    size_t body_pos;
    int fd;
} httpd_req_t;

typedef struct httpd_uri {
    const char* uri;
    httpd_method_t method;
    esp_err_t (*handler)(httpd_req_t* r);
    void* user_ctx;
    bool is_websocket;
    bool handle_ws_control_frames;
} httpd_uri_t;

typedef enum { HTTPD_WS_TYPE_CONTINUE = 0, HTTPD_WS_TYPE_TEXT = 1, HTTPD_WS_TYPE_BINARY = 2, HTTPD_WS_TYPE_CLOSE = 8, HTTPD_WS_TYPE_PING = 9, HTTPD_WS_TYPE_PONG = 10 } httpd_ws_type_t;
typedef struct httpd_ws_frame {
    bool final; bool fragmented; httpd_ws_type_t type; uint8_t* payload; size_t len;
} httpd_ws_frame_t;

typedef void* SemaphoreHandle_t;
#define portMAX_DELAY 0
static inline SemaphoreHandle_t xSemaphoreCreateMutex(void) { return malloc(1); }
static inline void vSemaphoreDelete(SemaphoreHandle_t s) { free(s); }
#define xSemaphoreTake(s, t) ((void) (s), (void) (t), 1)
#define xSemaphoreGive(s) ((void) (s))

/* ---- the fake server: filled by the generated *_register() ---- */
#define FAKE_MAX_HANDLERS 16
static httpd_uri_t fake_handlers[FAKE_MAX_HANDLERS];
static size_t fake_handler_count = 0;

static inline esp_err_t httpd_register_uri_handler(httpd_handle_t h, const httpd_uri_t* u) {
    (void) h;
    if (fake_handler_count >= FAKE_MAX_HANDLERS) return ESP_FAIL;
    fake_handlers[fake_handler_count++] = *u;
    return ESP_OK;
}
static inline esp_err_t httpd_unregister_uri_handler(httpd_handle_t h, const char* uri, httpd_method_t m) {
    (void) h; (void) uri; (void) m;
    return ESP_OK;
}

/* ---- request side ---- */
static inline size_t httpd_req_get_url_query_len(httpd_req_t* r) {
    const char* q = strchr(r->uri, '?');
    return q ? strlen(q + 1) : 0;
}
static inline esp_err_t httpd_req_get_url_query_str(httpd_req_t* r, char* buf, size_t len) {
    const char* q = strchr(r->uri, '?');
    if (!q) return ESP_FAIL;
    snprintf(buf, len, "%s", q + 1);
    return ESP_OK;
}
static inline int httpd_req_recv(httpd_req_t* r, char* buf, size_t len) {
    size_t left = r->content_len - r->body_pos;
    size_t n = left < len ? left : len;
    memcpy(buf, r->body + r->body_pos, n);
    r->body_pos += n;
    return (int) n;
}
static inline int httpd_req_to_sockfd(httpd_req_t* r) { return r->fd; }

/* ---- response side: captured for the assertions ---- */
static char fake_status[64];
static char fake_type[64];
static char fake_resp[2048];
static inline esp_err_t httpd_resp_set_status(httpd_req_t* r, const char* s) { (void) r; snprintf(fake_status, sizeof(fake_status), "%s", s); return ESP_OK; }
static inline esp_err_t httpd_resp_set_type(httpd_req_t* r, const char* t) { (void) r; snprintf(fake_type, sizeof(fake_type), "%s", t); return ESP_OK; }
static inline esp_err_t httpd_resp_send(httpd_req_t* r, const char* b, long len) {
    (void) r;
    if (len < 0) len = (long) strlen(b);
    memcpy(fake_resp, b, (size_t) len);
    fake_resp[len] = 0;
    return ESP_OK;
}

/* ---- websocket: one inbound frame at a time, outbound frames recorded ---- */
static httpd_ws_frame_t fake_inbound;
static int fake_sent_count = 0;
static int fake_sent_fds[16];
static char fake_sent_text[16][256];
static httpd_ws_type_t fake_sent_type[16];
static int fake_fail_fd = -1; /* sends to this fd fail, like a vanished client */

static inline esp_err_t httpd_ws_recv_frame(httpd_req_t* r, httpd_ws_frame_t* f, size_t max_len) {
    (void) r;
    f->type = fake_inbound.type;
    f->len = fake_inbound.len;
    if (max_len > 0 && f->payload) memcpy(f->payload, fake_inbound.payload, fake_inbound.len);
    return ESP_OK;
}
static inline esp_err_t httpd_ws_send_frame(httpd_req_t* r, httpd_ws_frame_t* f) { (void) r; (void) f; return ESP_OK; }
static inline esp_err_t httpd_ws_send_frame_async(httpd_handle_t h, int fd, httpd_ws_frame_t* f) {
    (void) h;
    if (fd == fake_fail_fd) return ESP_FAIL;
    fake_sent_fds[fake_sent_count] = fd;
    fake_sent_type[fake_sent_count] = f->type;
    memcpy(fake_sent_text[fake_sent_count], f->payload, f->len);
    fake_sent_text[fake_sent_count][f->len] = 0;
    fake_sent_count++;
    return ESP_OK;
}
#endif
`

const serverHostTestMain = `
#include <stdio.h>
#include "fake/esp_http_server.h"
#include "fake_override.h"
#include "set_temperature_action.h"
#include "get_state_action.h"
#include "state_stream_action.h"
#include "raw_frames_action.h"

#define CHECK(cond) do { if (!(cond)) { printf("FAILED line %d: %s\n", __LINE__, #cond); return 1; } } while (0)

/* Dispatches like esp_http_server: first registered handler whose method and
 * uri (through emihttpd_uri_match) fit. */
static esp_err_t dispatch(httpd_method_t method, const char* uri, const char* body, int fd) {
    for (size_t i = 0; i < fake_handler_count; i++) {
        if (fake_handlers[i].method != method) continue;
        if (!emihttpd_uri_match(fake_handlers[i].uri, uri, strlen(uri))) continue;
        httpd_req_t req;
        memset(&req, 0, sizeof(req));
        req.method = method;
        snprintf(req.uri, sizeof(req.uri), "%s", uri);
        req.body = body ? body : "";
        req.content_len = body ? strlen(body) : 0;
        req.user_ctx = fake_handlers[i].user_ctx;
        req.fd = fd;
        fake_status[0] = fake_type[0] = fake_resp[0] = 0;
        return fake_handlers[i].handler(&req);
    }
    snprintf(fake_status, sizeof(fake_status), "NO ROUTE");
    return ESP_FAIL;
}

static void ws_frame(httpd_ws_type_t t, const char* text) {
    fake_inbound.type = t;
    fake_inbound.payload = (uint8_t*) text;
    fake_inbound.len = strlen(text);
}

static float g_set_point = 0;
static int g_zone = -1;

static SetTemperatureActionResponse* on_set(const SetTemperatureActionCall* c, int* status) {
    g_zone = c->path->zone;
    if (c->body->target > 30) { *status = 422; return NULL; }
    g_set_point = c->body->target;
    SetTemperatureActionResponse* r = SetTemperatureActionResponse_new();
    r->ok = true;
    r->applied = c->body->target;
    if (c->body->target == 11) *status = 201;
    return r;
}

static char g_prefix[64];
static bool g_verbose;
static GetStateActionResponse* on_state(const GetStateActionCall* c, int* status) {
    (void) status;
    g_verbose = c->query->verbose;
    snprintf(g_prefix, sizeof(g_prefix), "%s", c->query->prefix ? c->query->prefix : "(none)");
    if (g_verbose == false && c->query->verbose_set == false && c->query->prefix == NULL) return NULL; /* -> 204 */
    GetStateActionResponse* r = GetStateActionResponse_new();
    r->roomSetPoint = 21.5f;
    r->inside = 19.0f;
    return r;
}

static int g_open_count, g_close_count, g_msgs;
static char g_room[32];
static int g_interval;
static float g_msg_set_point;
static int on_open(const StateStreamActionSession* s) {
    g_open_count++;
    snprintf(g_room, sizeof(g_room), "%s", s->path->room);
    g_interval = s->query->interval;
    return s->query->interval == 99 ? 1 : 0; /* 99 -> reject */
}
static void on_msg(const StateStreamActionSession* s, const StateStreamActionRequest* m) { (void) s; g_msgs++; g_msg_set_point = m->setPoint; }
static void on_close(const StateStreamActionSession* s) { (void) s; g_close_count++; }

static size_t g_bin_len;
static void on_bin(const RawFramesActionSession* s, const uint8_t* d, size_t len) { (void) s; (void) d; g_bin_len = len; }

int main(void) {
    /* ---- uri matching ---- */
    CHECK(emihttpd_uri_match("/api/ot/set-temp/:zone", "/api/ot/set-temp/7?x=1", 40));
    CHECK(!emihttpd_uri_match("/api/ot/set-temp/:zone", "/api/ot/set-temp/", 20));
    CHECK(!emihttpd_uri_match("/api/ot/set-temp/:zone", "/api/ot/set-temp/7/extra", 30));
    CHECK(!emihttpd_uri_match("/api/state", "/api/states", 11));
    CHECK(emihttpd_uri_match("/api/state", "/api/state", 10));
    {
        char* caps[4] = {0}; size_t n = 0;
        CHECK(emihttpd_path_capture("/a/:x/b/:y", "/a/hello%20world/b/42?q=1", caps, 4, &n));
        CHECK(n == 2 && strcmp(caps[0], "hello world") == 0 && strcmp(caps[1], "42") == 0);
        free(caps[0]); free(caps[1]);
        char* q = emihttpd_query_get("a=1&name=x%20y+z&flag", "name");
        CHECK(q && strcmp(q, "x y z") == 0); free(q);
        q = emihttpd_query_get("a=1&name=x&flag", "flag");
        CHECK(q && q[0] == 0); free(q);
        CHECK(emihttpd_query_get("a=1", "b") == NULL);
        CHECK(emihttpd_query_get("aa=1", "a") == NULL);
    }

    /* ---- plain http actions ---- */
    CHECK(set_temperature_action_register(NULL, on_set, NULL) != NULL);
    CHECK(get_state_action_register(NULL, on_state, NULL) != NULL);

    CHECK(dispatch(HTTP_POST, "/api/ot/set-temp/7", "{\"target\":22.5}", 0) == ESP_OK);
    CHECK(g_zone == 7 && g_set_point == 22.5f);
    CHECK(strcmp(fake_status, "200 OK") == 0);
    CHECK(strcmp(fake_type, "application/json") == 0);
    CHECK(strstr(fake_resp, "\"ok\":true") != NULL);       /* a real JSON boolean */
    CHECK(strstr(fake_resp, "\"applied\":22.5") != NULL);

    dispatch(HTTP_POST, "/api/ot/set-temp/7", "{\"target\":11}", 0);
    CHECK(strcmp(fake_status, "201 Created") == 0);        /* handler picked the status */

    dispatch(HTTP_POST, "/api/ot/set-temp/7", "{\"target\":50}", 0);
    CHECK(strncmp(fake_status, "422", 3) == 0);            /* NULL + status >= 400 */
    CHECK(strstr(fake_resp, "error") != NULL);

    dispatch(HTTP_POST, "/api/ot/set-temp/7", "{not json", 0);
    CHECK(strncmp(fake_status, "400", 3) == 0);
    dispatch(HTTP_GET, "/api/ot/set-temp/7", NULL, 0);     /* wrong method: not routed */
    CHECK(strcmp(fake_status, "NO ROUTE") == 0);

    dispatch(HTTP_GET, "/api/state?verbose=true&prefix=a%20b", NULL, 0);
    CHECK(g_verbose && strcmp(g_prefix, "a b") == 0);
    CHECK(strstr(fake_resp, "\"roomSetPoint\":21.5") != NULL);
    dispatch(HTTP_GET, "/api/state", NULL, 0);
    CHECK(strncmp(fake_status, "204", 3) == 0);            /* NULL, no status -> 204 */

    /* ---- websocket actions ---- */
    StateStreamActionCallbacks cb = {.on_open = on_open, .on_message = on_msg, .on_close = on_close};
    StateStreamActionBinding* b = state_stream_action_register(NULL, &cb);
    CHECK(b != NULL);
    RawFramesActionCallbacks rcb = {.on_binary = on_bin};
    RawFramesActionBinding* rb = raw_frames_action_register(NULL, &rcb);
    CHECK(rb != NULL);
    CHECK(fake_handlers[fake_handler_count - 2].is_websocket);

    CHECK(dispatch(HTTP_GET, "/ws/state/kitchen?interval=5", NULL, 10) == ESP_OK);   /* handshake */
    CHECK(g_open_count == 1 && strcmp(g_room, "kitchen") == 0 && g_interval == 5);
    CHECK(dispatch(HTTP_GET, "/ws/state/hall", NULL, 11) == ESP_OK);
    CHECK(state_stream_action_client_count(b) == 2);
    CHECK(dispatch(HTTP_GET, "/ws/state/bad?interval=99", NULL, 12) == ESP_FAIL);    /* on_open rejected */
    CHECK(state_stream_action_client_count(b) == 2);

    /* an incoming text frame arrives as the typed request */
    ws_frame(HTTPD_WS_TYPE_TEXT, "{\"setPoint\":18.5}");
    for (size_t i = 0; i < fake_handler_count; i++) if (fake_handlers[i].is_websocket && strstr(fake_handlers[i].uri, "state")) {
        httpd_req_t req; memset(&req, 0, sizeof(req));
        req.method = HTTP_POST; /* anything but GET marks a frame */
        req.user_ctx = fake_handlers[i].user_ctx; req.fd = 10;
        CHECK(fake_handlers[i].handler(&req) == ESP_OK);
    }
    CHECK(g_msgs == 1 && g_msg_set_point == 18.5f);

    /* broadcast reaches everyone, and serializes once */
    StateStreamActionResponse* upd = StateStreamActionResponse_new();
    upd->inside = 20.25f; upd->outside = 3.5f;
    fake_sent_count = 0;
    CHECK(state_stream_action_broadcast(b, upd) == 2);
    CHECK(fake_sent_count == 2 && strstr(fake_sent_text[0], "\"inside\":20.25") != NULL);
    CHECK(fake_sent_type[0] == HTTPD_WS_TYPE_TEXT);

    /* a client that vanished is dropped on the failed send, on_close fires */
    fake_fail_fd = 11; fake_sent_count = 0;
    CHECK(state_stream_action_broadcast(b, upd) == 1);
    CHECK(g_close_count == 1 && state_stream_action_client_count(b) == 1);
    fake_fail_fd = -1;
    StateStreamActionResponse_free(upd);

    /* a close frame removes the client */
    ws_frame(HTTPD_WS_TYPE_CLOSE, "");
    for (size_t i = 0; i < fake_handler_count; i++) if (fake_handlers[i].is_websocket && strstr(fake_handlers[i].uri, "state")) {
        httpd_req_t req; memset(&req, 0, sizeof(req));
        req.method = HTTP_POST; req.user_ctx = fake_handlers[i].user_ctx; req.fd = 10;
        fake_handlers[i].handler(&req);
    }
    CHECK(g_close_count == 2 && state_stream_action_client_count(b) == 0);

    /* binary action: frames in, frames out */
    CHECK(dispatch(HTTP_GET, "/ws/raw", NULL, 20) == ESP_OK);
    fake_inbound.type = HTTPD_WS_TYPE_BINARY;
    static uint8_t blob[5] = {1, 2, 3, 4, 5};
    fake_inbound.payload = blob; fake_inbound.len = 5;
    for (size_t i = 0; i < fake_handler_count; i++) if (fake_handlers[i].is_websocket && strstr(fake_handlers[i].uri, "raw")) {
        httpd_req_t req; memset(&req, 0, sizeof(req));
        req.method = HTTP_POST; req.user_ctx = fake_handlers[i].user_ctx; req.fd = 20;
        CHECK(fake_handlers[i].handler(&req) == ESP_OK);
    }
    CHECK(g_bin_len == 5);
    fake_sent_count = 0;
    CHECK(raw_frames_action_broadcast_binary(rb, blob, 5) == 1);
    CHECK(fake_sent_type[0] == HTTPD_WS_TYPE_BINARY);

    printf("ok\n");
    return 0;
}
`

// TestEspServer compiles the code generated with --tags esp-server on the host
// against a fake esp_http_server and drives it: routing and path/query/body
// decoding, status handling, typed JSON, and websocket session bookkeeping.
func TestEspServer(t *testing.T) {
	cc, err := exec.LookPath("cc")
	if err != nil {
		t.Skip("no C compiler available")
	}

	content, err := os.ReadFile("../../examples/esp-server/thermo.emi.yml")
	if err != nil {
		t.Fatal(err)
	}
	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}
	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}, Tags: string(EspServer)}
	files, err := CModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	for _, f := range files {
		full := filepath.Join(dir, f.Location, f.Name+f.Extension)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(f.ActualScript), 0644); err != nil {
			t.Fatal(err)
		}
	}
	for _, name := range []string{"set_temperature_action.h", "get_state_action.h", "state_stream_action.h", "raw_frames_action.h"} {
		b, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Fatalf("missing generated %s: %v", name, err)
		}
		if strings.Contains(string(b), "curl_slist") || strings.Contains(string(b), "fetchx_request") {
			t.Errorf("%s: server code must not pull in the libcurl client", name)
		}
	}

	// emihttpd.h includes esp_http_server.h / freertos - both come from the fake.
	if err := os.MkdirAll(filepath.Join(dir, "fake"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "freertos"), 0755); err != nil {
		t.Fatal(err)
	}
	for path, body := range map[string]string{
		"esp_http_server.h":      fakeEspHttpServerH,
		"fake/esp_http_server.h": `#include "../esp_http_server.h"`,
		"freertos/FreeRTOS.h":    "",
		"freertos/semphr.h":      "",
		"fake_override.h":        "",
		"main.c":                 serverHostTestMain,
	} {
		if err := os.WriteFile(filepath.Join(dir, path), []byte(body), 0644); err != nil {
			t.Fatal(err)
		}
	}

	bin := filepath.Join(dir, "srv")
	out, err := exec.Command(cc, "-std=c99", "-D_POSIX_C_SOURCE=200809L", "-Wall", "-Wextra", "-Werror",
		"-I", dir, filepath.Join(dir, "main.c"), filepath.Join(dir, "cjson", "cJSON.c"), "-o", bin).CombinedOutput()
	if err != nil {
		t.Fatalf("generated server C failed to compile:\n%s", out)
	}
	if out, err := exec.Command(bin).CombinedOutput(); err != nil {
		t.Fatalf("server behaviour checks failed: %v\n%s", err, out)
	}
}
