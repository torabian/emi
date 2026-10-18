package c

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

const envelopeYml = `
name: widgets
actions:
  - name: listWidgets
    url: /widgets
    method: get
    out:
      envelope: GResponse
      list: true
      fields:
        - {name: id, type: int}
        - {name: label, type: string}
  - name: getWidget
    url: /widgets/:id int
    method: get
    out:
      envelope: GResponse
      fields:
        - {name: id, type: int}
        - {name: label, type: string}
  - name: plainWidget
    url: /plain
    method: get
    out:
      fields:
        - {name: id, type: int}
`

const envelopeHostMain = `
#include <stdio.h>
#include <string.h>
#include "list_widgets_action.h"
#include "get_widget_action.h"
#include "plain_widget_action.h"

static ListWidgetsActionResult* on_list(const ListWidgetsActionCall* c, int* status) {
    (void) c; (void) status;
    ListWidgetsActionResult* r = (ListWidgetsActionResult*) calloc(1, sizeof(*r));
    r->items_count = 2;
    r->items = (ListWidgetsActionResponse**) calloc(2, sizeof(*r->items));
    for (int i = 0; i < 2; i++) {
        r->items[i] = ListWidgetsActionResponse_new();
        r->items[i]->id = i + 1;
        free(r->items[i]->label);
        r->items[i]->label = strdup(i ? "b" : "a");
    }
    r->total_items = 7;
    r->start_index = 3;
    r->items_per_page = 2;
    return r;
}
static GetWidgetActionResponse* on_get(const GetWidgetActionCall* c, int* status) {
    if (c->path->id != 1) { emihttpd_set_error("no such widget"); *status = 404; return NULL; }
    GetWidgetActionResponse* r = GetWidgetActionResponse_new();
    r->id = 1;
    return r;
}
static PlainWidgetActionResponse* on_plain(const PlainWidgetActionCall* c, int* status) {
    (void) c; (void) status;
    PlainWidgetActionResponse* r = PlainWidgetActionResponse_new();
    r->id = 5;
    return r;
}
int main(int argc, char** argv) {
    struct mg_mgr mgr; mg_mgr_init(&mgr);
    EmiHttpdRouter router; emihttpd_router_init(&router, &mgr);
    list_widgets_action_register(&router, on_list, NULL);
    get_widget_action_register(&router, on_get, NULL);
    plain_widget_action_register(&router, on_plain, NULL);
    if (!emihttpd_listen(&router, argv[1])) return 1;
    printf("ready\n"); fflush(stdout);
    for (;;) mg_mgr_poll(&mgr, 20);
}
`

// TestMongooseEnvelope: an action with envelope GResponse is answered as {"data":{"item":..}}, a list action as
// {"data":{"items":[..],"totalItems":..}}, an error as {"error":{"code":..,"message":..}}; one without an envelope
// keeps the bare body. Needs MONGOOSE_DIR like TestMongooseServer.
func TestMongooseEnvelope(t *testing.T) {
	mg := os.Getenv("MONGOOSE_DIR")
	if mg == "" {
		t.Skip("set MONGOOSE_DIR to a directory with mongoose.c/mongoose.h to run")
	}
	cc, err := exec.LookPath("cc")
	if err != nil {
		t.Skip("no C compiler available")
	}
	module, err := core.StringToEmi(envelopeYml)
	if err != nil {
		t.Fatal(err)
	}
	ctx := core.MicroGenContext{Content: envelopeYml, Flags: map[string]string{}, Tags: string(MongooseServer)}
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
	if err := os.WriteFile(filepath.Join(dir, "main.c"), []byte(envelopeHostMain), 0644); err != nil {
		t.Fatal(err)
	}
	bin := filepath.Join(dir, "srv")
	out, err := exec.Command(cc, "-std=c99", "-D_POSIX_C_SOURCE=200809L", "-D_DARWIN_C_SOURCE", "-Wall", "-Wextra",
		"-I", dir, "-I", mg, filepath.Join(dir, "main.c"), filepath.Join(dir, "cjson", "cJSON.c"),
		filepath.Join(mg, "mongoose.c"), "-o", bin, "-lm", "-lpthread").CombinedOutput()
	if err != nil {
		t.Fatalf("generated mongoose server failed to compile:\n%s", out)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := ln.Addr().String()
	_ = ln.Close()
	cmd := exec.Command(bin, "http://"+addr)
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = cmd.Process.Kill(); _ = cmd.Wait() }()
	lines := bufio.NewScanner(stdout)
	ready := false
	for lines.Scan() {
		if lines.Text() == "ready" {
			ready = true
			break
		}
	}
	if !ready {
		t.Fatal("server did not start")
	}
	go func() {
		for lines.Scan() {
		}
	}()

	get := func(path string) (int, string) {
		resp, err := http.Get("http://" + addr + path)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		return resp.StatusCode, string(b)
	}
	if code, body := get("/widgets"); code != 200 ||
		!strings.Contains(body, `"data":{"items":[{"id":1,"label":"a"},{"id":2,"label":"b"}]`) ||
		!strings.Contains(body, `"totalItems":7`) || !strings.Contains(body, `"startIndex":3`) || !strings.Contains(body, `"itemsPerPage":2`) {
		t.Errorf("list: %d %s", code, body)
	}
	if code, body := get("/widgets/1"); code != 200 || !strings.HasPrefix(body, `{"data":{"item":{"id":1`) {
		t.Errorf("single: %d %s", code, body)
	}
	if code, body := get("/widgets/9"); code != 404 || body != `{"error":{"code":404,"message":"no such widget"}}` {
		t.Errorf("error: %d %s", code, body)
	}
	if code, body := get("/plain"); code != 200 || body != `{"id":5}` {
		t.Errorf("no envelope must stay bare: %d %s", code, body)
	}
}
