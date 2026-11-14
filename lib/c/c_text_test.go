package c

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// The program drives the text transport without a socket: the same routes the http server has, called with lines.
const textHostMain = `
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "set_temperature_action.h"
#include "get_state_action.h"
#include "state_stream_action.h"

static SetTemperatureActionResponse* on_set(const SetTemperatureActionCall* c, int* status) {
    if (c->body->target > 30) { emihttpd_set_error("target above 30"); *status = 422; return NULL; }
    SetTemperatureActionResponse* r = SetTemperatureActionResponse_new();
    r->ok = true;
    r->applied = c->body->target + (float) c->path->zone;
    return r;
}
static GetStateActionResponse* on_state(const GetStateActionCall* c, int* status) {
    (void) status;
    if (!c->query->verbose_set) return NULL; /* 204 */
    GetStateActionResponse* r = GetStateActionResponse_new();
    r->roomSetPoint = 21.5f;
    if (c->query->prefix) r->inside = (float) strlen(c->query->prefix);
    return r;
}

static StateStreamActionBinding* g_stream;
static int on_open(const StateStreamActionSession* s) { return strcmp(s->path->room, "locked") == 0; }
static void on_msg(const StateStreamActionSession* s, const StateStreamActionRequest* m) {
    StateStreamActionResponse* r = StateStreamActionResponse_new();
    r->inside = m->setPoint * 2;
    state_stream_action_send(g_stream, s->conn, r);
    StateStreamActionResponse_free(r);
}
static char wsout[1024];
static void sink(emihttpd_conn_t conn, bool binary, const uint8_t* d, size_t n, void* user) {
    (void) binary; (void) user;
    size_t at = strlen(wsout);
    snprintf(wsout + at, sizeof(wsout) - at, "[%d %.*s]", (int) (conn > 0), (int) n, (const char*) d);
}

static char out[8192];
static size_t pieces;
static size_t longest;
static int collect(const uint8_t* d, size_t n, void* user) {
    (void) user;
    pieces++;
    if (n > longest) longest = n;
    strncat(out, (const char*) d, n);
    return 0;
}
static void call(EmiHttpdRouter* r, const char* line) {
    char* f = emihttpd_text_call(r, line);
    printf("call> %s", f ? f : "(null)\n");
    free(f);
}
int main(void) {
    mg_log_set(MG_LL_NONE);
    struct mg_mgr mgr; mg_mgr_init(&mgr);
    EmiHttpdRouter router; emihttpd_router_init(&router, &mgr);
    set_temperature_action_register(&router, on_set, NULL);
    get_state_action_register(&router, on_state, NULL);
    StateStreamActionCallbacks cb = {.on_open = on_open, .on_message = on_msg};
    g_stream = state_stream_action_register(&router, &cb);

    call(&router, "get /api/state?verbose=true&prefix=ab%20c");
    call(&router, ":::GET /api/state");
    call(&router, "post#7 /api/ot/set-temp/7 {\"target\":22.5}");
    call(&router, "post#e-1 /api/ot/set-temp/7 {\"target\":50}");
    call(&router, "post /api/ot/set-temp/7 {broken");
    call(&router, "get /api/ot/set-temp/7");
    call(&router, "get /nothing/here");
    call(&router, "get /ws/state/kitchen");
    call(&router, "fetch /api/state");
    call(&router, "get api/state");
    call(&router, "get#bad!id /api/state");
    call(&router, "routes");

    /* a stream: log noise, a request cut at every byte, a request after a partial line, an answer in pieces of 8 bytes */
    EmiTextLink link;
    emihttpd_text_link_init(&link, &router, collect, NULL, 8);
    const char* stream = "boot: wifi up\r\n:::get#a /api/state?verbose=1\n:: not ours\nxx:::get /api/state\n:::post#b /api/ot/set-temp/2 {\"target\":1}\r\n:::\n";
    for (size_t i = 0; stream[i]; i++) emitext_link_feed(&link, (const uint8_t*) stream + i, 1);
    printf("stream> %s", out);
    printf("pieces> %d longest %d\n", (int) (pieces > 3), (int) longest);

    out[0] = '\0';
    const char* msg = ":::get#m /api/state?verbose=1";
    emitext_link_message(&link, (const uint8_t*) msg, strlen(msg));
    msg = "get /api/state";
    emitext_link_message(&link, (const uint8_t*) msg, strlen(msg));
    printf("message> %s", out);

    /* a line longer than the limit is answered 413 and the next request is read */
    out[0] = '\0';
    size_t big = EMITEXT_MAX_LINE + 100;
    char* line = (char*) malloc(big + 64);
    strcpy(line, ":::post /api/ot/set-temp/1 {\"target\":\"");
    memset(line + strlen(line), 'x', big);
    strcpy(line + strlen(line), "\"}\n:::get#n /api/state?verbose=1\n");
    emitext_link_feed(&link, (const uint8_t*) line, strlen(line));
    free(line);
    printf("long> %s", out);

    /* a websocket action without a socket: the client is the program, the frames come out of the sink */
    router.ws_sink = sink;
    emihttpd_conn_t c = emihttpd_inproc_ws_open(&router, "/ws/state/kitchen");
    emihttpd_conn_t locked = emihttpd_inproc_ws_open(&router, "/ws/state/locked");
    emihttpd_conn_t none = emihttpd_inproc_ws_open(&router, "/ws/nothing");
    const char* hello = "{\"setPoint\":21}";
    int sent = emihttpd_inproc_ws_message(&router, c, false, (const uint8_t*) hello, strlen(hello));
    emihttpd_inproc_ws_close(&router, c);
    int after = emihttpd_inproc_ws_message(&router, c, false, (const uint8_t*) hello, strlen(hello));
    printf("inproc> open=%d locked=%d none=%d sent=%d after=%d frames=%s\n", (int) (c > 0), (int) locked, (int) none, sent, after, wsout);

    emitext_link_deinit(&link);
    mg_mgr_free(&mgr);
    return 0;
}
`

func TestTextRouter(t *testing.T) {
	mg := os.Getenv("MONGOOSE_DIR")
	if mg == "" {
		t.Skip("set MONGOOSE_DIR to a directory with mongoose.c/mongoose.h to run")
	}
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

	// the text transport is a mongoose feature
	if _, err := CModuleFullVirtualFiles(&module, core.MicroGenContext{Content: string(content), Flags: map[string]string{}, Tags: string(TextRouter)}); err == nil {
		t.Error("text-router without mongoose-server must be refused")
	}

	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}, Tags: string(MongooseServer) + "," + string(TextRouter)}
	files, err := CModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	sawText := false
	for _, f := range files {
		full := filepath.Join(dir, f.Location, f.Name+f.Extension)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(f.ActualScript), 0644); err != nil {
			t.Fatal(err)
		}
		sawText = sawText || f.Name+f.Extension == "emitext.h"
		if strings.HasPrefix(f.Name, "argparse/") || f.Name+f.Extension == "emicli.h" || f.Name+f.Extension == "emi_cli.h" {
			t.Errorf("%s is the command line, which was not asked for", f.Name)
		}
	}
	if !sawText {
		t.Fatal("emitext.h not shipped")
	}
	if err := os.WriteFile(filepath.Join(dir, "main.c"), []byte(textHostMain), 0644); err != nil {
		t.Fatal(err)
	}
	bin := filepath.Join(dir, "text")
	out, err := exec.Command(cc, "-std=c99", "-D_POSIX_C_SOURCE=200809L", "-D_DARWIN_C_SOURCE", "-Wall", "-Wextra",
		"-I", dir, "-I", mg, filepath.Join(dir, "main.c"), filepath.Join(dir, "cjson", "cJSON.c"),
		filepath.Join(mg, "mongoose.c"), "-o", bin, "-lm", "-lpthread").CombinedOutput()
	if err != nil {
		t.Fatalf("text router failed to compile:\n%s", out)
	}
	res, err := exec.Command(bin).CombinedOutput()
	if err != nil {
		t.Fatalf("run: %v\n%s", err, res)
	}
	got := string(res)

	for _, want := range []string{
		`call> :::200 {"roomSetPoint":21.5,"inside":4}`,
		"call> :::204\n",
		`call> :::200#7 {"ok":true,"applied":29.5}`,
		`call> :::422#e-1 {"error":"target above 30"}`,
		`call> :::400 {"error":"invalid json body"}`,
		`call> :::405 {"error":"method not allowed"}`,
		`call> :::404 {"error":"not found"}`,
		`call> :::501 {"error":"a websocket action is not served over text"}`,
		`call> :::400 {"error":"unreadable request"}`,
		`{"method":"POST","path":"/api/ot/set-temp/:zone"}`,
		// the stream: the log and the broken lines are skipped, the answers carry the ids, a request in the middle of a line is not one
		`stream> :::200#a {"roomSetPoint":21.5,"inside":0}` + "\n" + `:::200#b {"ok":true,"applied":3}` + "\n" + `:::400 {"error":"unreadable request"}` + "\n",
		"pieces> 1 longest 8\n",
		`message> :::200#m {"roomSetPoint":21.5,"inside":0}` + "\n" + ":::204\n",
		// a websocket action in process: opened, refused by the hook (locked), no such action; a message is answered through the sink; after close it is no connection
		`inproc> open=1 locked=0 none=0 sent=0 after=-1 frames=[1 {"inside":42`,
		`long> :::413 {"error":"request too long"}` + "\n" + `:::200#n {"roomSetPoint":21.5,"inside":0}` + "\n",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in:\n%s", want, got)
		}
	}
	if strings.Count(got, "call> :::400") != 4 {
		t.Errorf("expected four 400 answers (broken json, bad method, no slash, bad id):\n%s", got)
	}
}
