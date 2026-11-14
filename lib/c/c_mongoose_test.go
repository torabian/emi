package c

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/torabian/emi/lib/core"
)

const mongooseHostMain = `
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "set_temperature_action.h"
#include "get_state_action.h"
#include "state_stream_action.h"

static SetTemperatureActionResponse* on_set(const SetTemperatureActionCall* c, int* status) {
    if (c->body->target > 30) { emihttpd_set_error("target above 30, \"max\""); *status = 422; return NULL; }
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
    /* nested qs: filter[zone], filter[page][n], filter[page][size] */
    if (c->query->filter.zone) r->inside = (c->query->filter.page.n_set ? c->query->filter.page.n : -1) * 100 + c->query->filter.page.size + (float) strlen(c->query->filter.zone);
    return r;
}
static StateStreamActionBinding* g_stream;
static int on_open(const StateStreamActionSession* s) {
    printf("open %s\n", s->path->room); fflush(stdout);
    return strcmp(s->path->room, "locked") == 0;
}
static void on_msg(const StateStreamActionSession* s, const StateStreamActionRequest* m) {
    StateStreamActionResponse* r = StateStreamActionResponse_new();
    r->inside = m->setPoint * 2;
    state_stream_action_send(g_stream, s->conn, r);
    StateStreamActionResponse_free(r);
}
static void on_request(void* user) { (void) user; printf("request\n"); fflush(stdout); }
int main(int argc, char** argv) {
    struct mg_mgr mgr; mg_mgr_init(&mgr);
    EmiHttpdRouter router; emihttpd_router_init(&router, &mgr);
    router.on_request = on_request; /* told of every HTTP request, whichever route (or none) answers it */
    set_temperature_action_register(&router, on_set, NULL);
    get_state_action_register(&router, on_state, NULL);
    StateStreamActionCallbacks cb = {.on_open = on_open, .on_message = on_msg};
    g_stream = state_stream_action_register(&router, &cb);
    if (!emihttpd_listen(&router, argv[1])) return 1;
    printf("ready\n"); fflush(stdout);
    for (;;) mg_mgr_poll(&mgr, 20);
}
`

func wsHandshake(t *testing.T, addr, path string) (net.Conn, *bufio.Reader, string) {
	t.Helper()
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	_ = conn.SetDeadline(time.Now().Add(5 * time.Second))
	key := make([]byte, 16)
	_, _ = rand.Read(key)
	fmt.Fprintf(conn, "GET %s HTTP/1.1\r\nHost: %s\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Key: %s\r\nSec-WebSocket-Version: 13\r\n\r\n",
		path, addr, base64.StdEncoding.EncodeToString(key))
	br := bufio.NewReader(conn)
	status, err := br.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			t.Fatal(err)
		}
		if line == "\r\n" {
			break
		}
	}
	return conn, br, status
}

func wsReadText(t *testing.T, br *bufio.Reader) string {
	t.Helper()
	var h [2]byte
	if _, err := io.ReadFull(br, h[:]); err != nil {
		t.Fatal(err)
	}
	n := int(h[1] & 0x7f)
	switch n {
	case 126:
		var b [2]byte
		_, _ = io.ReadFull(br, b[:])
		n = int(binary.BigEndian.Uint16(b[:]))
	case 127:
		var b [8]byte
		_, _ = io.ReadFull(br, b[:])
		n = int(binary.BigEndian.Uint64(b[:]))
	}
	payload := make([]byte, n)
	if _, err := io.ReadFull(br, payload); err != nil {
		t.Fatal(err)
	}
	return string(payload)
}

func wsSendText(conn net.Conn, text string) {
	mask := []byte{1, 2, 3, 4}
	data := []byte(text)
	frame := append([]byte{0x81, 0x80 | byte(len(data))}, mask...)
	for i, b := range data {
		frame = append(frame, b^mask[i%4])
	}
	_, _ = conn.Write(frame)
}

// TestMongooseServer builds the code generated with --tags mongoose-server
// against a real Mongoose and drives it over sockets: HTTP routing, path,
// query and body decoding, status codes, and a websocket session. Mongoose is
// GPLv2/commercial and not vendored: point MONGOOSE_DIR at a directory with
// mongoose.c and mongoose.h to run it.
func TestMongooseServer(t *testing.T) {
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
	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}, Tags: string(MongooseServer)}
	files, err := CModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	sawMongoose, sawEsp := false, false
	for _, f := range files {
		full := filepath.Join(dir, f.Location, f.Name+f.Extension)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(f.ActualScript), 0644); err != nil {
			t.Fatal(err)
		}
		sawMongoose = sawMongoose || strings.HasPrefix(f.Name, "emihttpd_mongoose")
		sawEsp = sawEsp || strings.HasPrefix(f.Name, "emihttpd_esp")
	}
	if !sawMongoose || sawEsp {
		t.Fatalf("expected only the mongoose backend to be shipped (mongoose=%v esp=%v)", sawMongoose, sawEsp)
	}
	if err := os.WriteFile(filepath.Join(dir, "main.c"), []byte(mongooseHostMain), 0644); err != nil {
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
	// Mongoose logs to stdout too: wait for our own "ready" line, then keep draining.
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
	var heard, asked atomic.Int32 // requests the server's on_request hook reported / requests this test made
	go func() {
		for lines.Scan() {
			if lines.Text() == "request" {
				heard.Add(1)
			}
		}
	}()

	do := func(method, path, body string) (int, string, http.Header) {
		asked.Add(1)
		req, _ := http.NewRequest(method, "http://"+addr+path, strings.NewReader(body))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		return resp.StatusCode, string(b), resp.Header
	}

	code, body, hdr := do("POST", "/api/ot/set-temp/7", `{"target":22.5}`)
	if code != 200 || !strings.Contains(body, `"ok":true`) || !strings.Contains(body, `"applied":29.5`) {
		t.Errorf("set-temp: %d %s", code, body)
	}
	if hdr.Get("Content-Type") != "application/json" {
		t.Errorf("content type %q", hdr.Get("Content-Type"))
	}
	if hdr.Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("an answer must allow any origin, got %q", hdr.Get("Access-Control-Allow-Origin"))
	}
	// the preflight of a cross origin request: no route handles OPTIONS, the runtime answers it
	pre, _ := http.NewRequest("OPTIONS", "http://"+addr+"/api/ot/set-temp/7", nil)
	pre.Header.Set("Access-Control-Request-Headers", "workspace-id, content-type")
	if resp, err := http.DefaultClient.Do(pre); err != nil || resp.Header.Get("Access-Control-Allow-Headers") != "workspace-id, content-type" {
		t.Errorf("a preflight must allow the headers it asks for: %v %v", err, resp)
	}
	code, _, hdr = do("OPTIONS", "/api/ot/set-temp/7", "")
	if code != 204 || hdr.Get("Access-Control-Allow-Origin") != "*" || !strings.Contains(hdr.Get("Access-Control-Allow-Methods"), "PATCH") || !strings.Contains(hdr.Get("Access-Control-Allow-Headers"), "Content-Type") {
		t.Errorf("preflight: %d %v", code, hdr)
	}
	if code, body, _ = do("POST", "/api/ot/set-temp/7", `{"target":50}`); code != 422 {
		t.Errorf("handler status not used: %d", code)
	} else if body != `{"error":"target above 30, 'max'"}` {
		t.Errorf("handler error text not used (and made json safe): %s", body)
	}
	if code, body, _ = do("POST", "/api/ot/set-temp/7", `{broken`); code != 400 || strings.Contains(body, "target above") {
		t.Errorf("an error text must be used once: %d %s", code, body)
	}
	if code, _, _ = do("POST", "/api/ot/set-temp/7", `{broken`); code != 400 {
		t.Errorf("bad json: %d", code)
	}
	if code, _, _ = do("GET", "/api/ot/set-temp/7", ""); code != 404 {
		t.Errorf("wrong method should not route: %d", code)
	}
	if code, body, _ = do("GET", "/api/state?verbose=true", ""); code != 200 || !strings.Contains(body, `"roomSetPoint":21.5`) {
		t.Errorf("state: %d %s", code, body)
	}
	if code, body, _ = do("GET", "/api/state?verbose=1&filter%5Bzone%5D=ab&filter[page][n]=3&filter%5Bpage%5D%5Bsize%5D=7", ""); code != 200 || !strings.Contains(body, `"inside":309`) {
		t.Errorf("nested query: %d %s", code, body)
	}
	if code, _, hdr = do("GET", "/api/state", ""); code != 204 {
		t.Errorf("NULL response should be 204: %d", code)
	} else if hdr.Get("Content-Type") != "" {
		t.Errorf("an empty answer is not json: Content-Type %q", hdr.Get("Content-Type"))
	}

	// websocket: handshake with path param, a typed message answered by the handler
	conn, br, status := wsHandshake(t, addr, "/ws/state/kitchen?interval=5")
	defer conn.Close()
	if !strings.Contains(status, "101") {
		t.Fatalf("no upgrade: %s", status)
	}
	wsSendText(conn, `{"setPoint":10.5}`)
	if got := wsReadText(t, br); !strings.Contains(got, `"inside":21`) {
		t.Errorf("ws reply: %s", got)
	}
	// on_open can reject
	rc, _, rstatus := wsHandshake(t, addr, "/ws/state/locked")
	defer rc.Close()
	if !strings.Contains(rstatus, "403") {
		t.Errorf("rejected client should get 403: %s", rstatus)
	}

	// on_request heard every plain request (routes, 404s) and the websocket upgrades; give the output pipe a moment
	deadline := time.Now().Add(2 * time.Second)
	for heard.Load() < asked.Load() && time.Now().Before(deadline) {
		time.Sleep(20 * time.Millisecond)
	}
	if heard.Load() < asked.Load() || asked.Load() == 0 {
		t.Errorf("on_request heard %d of %d requests", heard.Load(), asked.Load())
	}
}

// A small device has few sockets and a browser keeps the connections of a page open for minutes: with -DEMIHTTPD_IDLE_MS a plain HTTP connection which has asked for nothing for that
// long is closed by the server, one in use is not, and a new one works.
func TestMongooseIdleConnectionsAreClosed(t *testing.T) {
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
	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}, Tags: string(MongooseServer)}
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
	if err := os.WriteFile(filepath.Join(dir, "main.c"), []byte(mongooseHostMain), 0644); err != nil {
		t.Fatal(err)
	}
	bin := filepath.Join(dir, "srv")
	if out, err := exec.Command(cc, "-std=c99", "-D_POSIX_C_SOURCE=200809L", "-D_DARWIN_C_SOURCE", "-Wall", "-Wextra", "-DEMIHTTPD_IDLE_MS=400",
		"-I", dir, "-I", mg, filepath.Join(dir, "main.c"), filepath.Join(dir, "cjson", "cJSON.c"),
		filepath.Join(mg, "mongoose.c"), "-o", bin, "-lm", "-lpthread").CombinedOutput(); err != nil {
		t.Fatalf("generated mongoose server with an idle limit failed to compile:\n%s", out)
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

	const request = "GET /nothing-here HTTP/1.1\r\nHost: x\r\n\r\n"
	// closedWithin: whether the server closes the connection (a read ends with EOF) within the time, without our sending anything
	closedWithin := func(c net.Conn, d time.Duration) bool {
		_ = c.SetReadDeadline(time.Now().Add(d))
		buf := make([]byte, 4096)
		for {
			if _, err := c.Read(buf); err != nil {
				ne, isNet := err.(net.Error)
				return !(isNet && ne.Timeout())
			}
		}
	}

	// a keep-alive connection which was used and is left alone: closed
	c1, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c1.Close()
	_, _ = c1.Write([]byte(request))
	_ = c1.SetReadDeadline(time.Now().Add(2 * time.Second))
	if n, err := c1.Read(make([]byte, 4096)); err != nil || n == 0 {
		t.Fatalf("no answer to the request: %v", err)
	}
	if !closedWithin(c1, 3*time.Second) {
		t.Error("an idle keep-alive connection must be closed by the server")
	}

	// one that never asked for anything is closed too
	c2, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c2.Close()
	if !closedWithin(c2, 3*time.Second) {
		t.Error("a connection which asks for nothing must be closed by the server")
	}

	// one in use is not closed: a request every 150 ms for 1.2 s (three idle limits) is answered every time
	c3, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c3.Close()
	for i := 0; i < 8; i++ {
		_, _ = c3.Write([]byte(request))
		_ = c3.SetReadDeadline(time.Now().Add(2 * time.Second))
		if n, err := c3.Read(make([]byte, 4096)); err != nil || n == 0 {
			t.Fatalf("request %d on a connection in use was not answered: %v", i, err)
		}
		time.Sleep(150 * time.Millisecond)
	}

	// and a new connection works after all that
	resp, err := http.Get("http://" + addr + "/nothing-here")
	if err != nil {
		t.Fatalf("a new connection must work: %v", err)
	}
	_ = resp.Body.Close()
}
