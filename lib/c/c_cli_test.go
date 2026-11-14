package c

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

const cliHostMain = `
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "set_temperature_action.h"
#include "get_state_action.h"
#include "state_stream_action.h"
#include "emi_cli.h"

static SetTemperatureActionResponse* on_set(const SetTemperatureActionCall* c, int* status) {
    if (c->body->target > 30) { emihttpd_set_error("target above 30"); *status = 422; return NULL; }
    SetTemperatureActionResponse* r = SetTemperatureActionResponse_new();
    r->ok = true;
    r->applied = c->body->target + (float) c->path->zone;
    return r;
}
static GetStateActionResponse* on_state(const GetStateActionCall* c, int* status) {
    (void) status;
    if (!c->query->verbose_set) return NULL;
    GetStateActionResponse* r = GetStateActionResponse_new();
    r->roomSetPoint = c->query->verbose ? 1.5f : 0.5f;
    if (c->query->filter.zone) r->inside = (c->query->filter.page.n_set ? c->query->filter.page.n : -1) * 100 + c->query->filter.page.size + (float) strlen(c->query->filter.zone);
    return r;
}
int main(int argc, const char** argv) {
    mg_log_set(MG_LL_NONE);
    struct mg_mgr mgr; mg_mgr_init(&mgr);
    EmiHttpdRouter router; emihttpd_router_init(&router, &mgr);
    set_temperature_action_register(&router, on_set, NULL);
    get_state_action_register(&router, on_state, NULL);
    StateStreamActionCallbacks cb = {0};
    state_stream_action_register(&router, &cb);
    int rc = emi_cli_main(&router, argc, argv);
    mg_mgr_free(&mgr);
    return rc;
}
`

func TestCli(t *testing.T) {
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
	if _, err := CModuleFullVirtualFiles(&module, core.MicroGenContext{Content: string(content), Flags: map[string]string{}, Tags: string(Cli)}); err == nil {
		t.Error("cli without mongoose-server must be refused")
	}
	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}, Tags: string(MongooseServer) + "," + string(Cli)}
	files, err := CModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	seen := map[string]bool{}
	for _, f := range files {
		full := filepath.Join(dir, f.Location, f.Name+f.Extension)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(f.ActualScript), 0644); err != nil {
			t.Fatal(err)
		}
		seen[filepath.Join(f.Location, f.Name+f.Extension)] = true
	}
	for _, want := range []string{"emi_cli.h", "emicli.h", "emitext.h", "argparse/argparse.c", "argparse/argparse.h", "argparse/LICENSE"} {
		if !seen[want] {
			t.Errorf("%s not shipped (got %v)", want, seen)
		}
	}
	if err := os.WriteFile(filepath.Join(dir, "main.c"), []byte(cliHostMain), 0644); err != nil {
		t.Fatal(err)
	}
	bin := filepath.Join(dir, "cli")
	out, err := exec.Command(cc, "-std=c99", "-D_POSIX_C_SOURCE=200809L", "-D_DARWIN_C_SOURCE", "-Wall", "-Wextra", "-Wno-unused-parameter",
		"-I", dir, "-I", mg, filepath.Join(dir, "main.c"), filepath.Join(dir, "cjson", "cJSON.c"), filepath.Join(dir, "argparse", "argparse.c"),
		filepath.Join(mg, "mongoose.c"), "-o", bin, "-lm", "-lpthread").CombinedOutput()
	if err != nil {
		t.Fatalf("generated cli failed to compile:\n%s", out)
	}
	if strings.Contains(string(out), "warning") {
		t.Errorf("warnings:\n%s", out)
	}

	run := func(args ...string) (string, string, int) {
		cmd := exec.Command(bin, args...)
		var so, se strings.Builder
		cmd.Stdout, cmd.Stderr = &so, &se
		err := cmd.Run()
		code := 0
		if ee, ok := err.(*exec.ExitError); ok {
			code = ee.ExitCode()
		} else if err != nil {
			t.Fatal(err)
		}
		return so.String(), se.String(), code
	}

	if so, _, code := run("set-temperature", "--zone", "7", "--target", "22.5"); code != 0 || so != `{"ok":true,"applied":29.5}`+"\n" {
		t.Errorf("set-temperature: %d %q", code, so)
	}
	if so, _, code := run("set-temperature", "--zone", "7", "--body", `{"target":1}`, "--pretty"); code != 0 || !strings.Contains(so, "{\n\t\"ok\":\ttrue,\n\t\"applied\":\t8\n}") {
		t.Errorf("body + pretty: %d %q", code, so)
	}
	if so, _, code := run("set-temperature", "--zone", "7", "--body", `{"target":1}`, "--target", "3"); code != 0 || !strings.Contains(so, `"applied":10`) {
		t.Errorf("a flag overrides --body: %d %q", code, so)
	}
	if _, se, code := run("set-temperature", "--zone", "7", "--target", "50"); code != 1 || se != `error: 422 {"error":"target above 30"}`+"\n" {
		t.Errorf("handler error: %d %q", code, se)
	}
	if _, se, code := run("set-temperature", "--target", "5"); code != 2 || !strings.Contains(se, "--zone is required") {
		t.Errorf("missing path parameter: %d %q", code, se)
	}
	if _, se, code := run("set-temperature", "--zone", "x", "--target", "5"); code != 2 || !strings.Contains(se, "--zone must be int") {
		t.Errorf("bad type: %d %q", code, se)
	}
	if _, se, code := run("set-temperature", "--zone", "1", "--target", "x"); code != 2 || !strings.Contains(se, "--target must be number") {
		t.Errorf("bad body type: %d %q", code, se)
	}
	if so, _, code := run("get-state", "--verbose", "yes", "--filter-zone", "ab", "--filter-page-n", "3", "--filter-page-size", "7"); code != 0 || !strings.Contains(so, `"roomSetPoint":1.5`) || !strings.Contains(so, `"inside":309`) {
		t.Errorf("nested query: %d %q", code, so)
	}
	if so, se, code := run("get-state"); code != 0 || so != "" || se != "" {
		t.Errorf("204 prints nothing: %d %q %q", code, so, se)
	}
	if so, _, code := run("request", "get", "/api/state?verbose=1&filter%5Bzone%5D=a%20bc"); code != 0 || !strings.Contains(so, `"inside":-96`) {
		t.Errorf("request: %d %q", code, so)
	}
	if _, se, code := run("request", "get", "/nothing"); code != 1 || !strings.HasPrefix(se, "error: 404") {
		t.Errorf("request 404: %d %q", code, se)
	}
	if so, _, code := run("help"); code != 0 || !strings.Contains(so, "set-temperature") || !strings.Contains(so, "/api/ot/set-temp/:zone") || strings.Contains(so, "state-stream") {
		t.Errorf("help (a websocket action is not a command): %d %q", code, so)
	}
	if so, _, code := run("set-temperature", "--help"); code != 0 || !strings.Contains(so, "--zone") || !strings.Contains(so, "--target") || !strings.Contains(so, "--body") {
		t.Errorf("command help: %d %q", code, so)
	}
	if _, _, code := run("nope"); code != 2 {
		t.Errorf("unknown command: %d", code)
	}
}
