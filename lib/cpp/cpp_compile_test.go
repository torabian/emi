package cpp

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestSmokeGenericCompiles renders a feature-rich dto through the generic
// dialect and actually invokes clang++ -fsyntax-only against it, wired up with
// the real vendored runtime headers - a much stronger guarantee than any Go-side
// string assertion could give for generated C++. Skips (rather than fails) when
// clang++ isn't available, same as every other environment-dependent check.
func TestSmokeGenericCompiles(t *testing.T) {
	if _, err := exec.LookPath("clang++"); err != nil {
		t.Skip("clang++ not available")
	}

	fields := []*core.EmiField{
		{Name: "title", Type: core.FieldTypeString},
		{Name: "nickname", Type: core.FieldTypeStringNullable},
		{Name: "age", Type: core.FieldTypeInt},
		{Name: "score", Type: core.FieldTypeFloat64Nullable},
		{Name: "active", Type: core.FieldTypeBool},
		{Name: "tags", Type: core.FieldTypeSlice, Primitive: "string"},
		{Name: "scores", Type: core.FieldTypeSlice, Primitive: "int"},
		{
			Name: "address",
			Type: core.FieldTypeObject,
			Fields: []*core.EmiField{
				{Name: "city", Type: core.FieldTypeString},
				{Name: "zip", Type: core.FieldTypeStringNullable},
			},
		},
		{
			Name: "history",
			Type: core.FieldTypeArray,
			Fields: []*core.EmiField{
				{Name: "note", Type: core.FieldTypeString},
			},
		},
		{Name: "status", Type: core.FieldTypeEnum, OfType: []*core.EmiEnumInline{{Key: "active"}, {Key: "inactive"}}},
		{Name: "parent", Type: core.FieldTypeOne, Target: "SmokeWidget"},
		{Name: "children", Type: core.FieldTypeCollection, Target: "SmokeWidget"},
		{Name: "meta", Type: core.FieldTypeMap, MapKeyOf: "string", MapPairOf: "string"},
		{Name: "raw", Type: core.FieldTypeAny},
	}

	chunk, err := CppCommonObjectGenerator(fields, core.MicroGenContext{}, CppCommonObjectContext{
		Dialect:       DialectGeneric,
		RootClassName: "SmokeWidget",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	doc := CppGenericAsFullDocument(chunk)

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "SmokeWidget.hpp"), []byte(doc), 0644); err != nil {
		t.Fatal(err)
	}

	main := `#include "SmokeWidget.hpp"
int main() {
    SmokeWidget w;
    std::string s = w.Dump();
    SmokeWidget w2 = SmokeWidget::FromJson(nullptr);
    (void) s; (void) w2;
    return 0;
}
`
	if err := os.WriteFile(filepath.Join(dir, "main.cpp"), []byte(main), 0644); err != nil {
		t.Fatal(err)
	}

	includeDir, _ := filepath.Abs("cpp-include/generic")
	cmd := exec.Command("clang++", "-std=c++17", "-Wall", "-Wextra", "-fsyntax-only",
		"-I", dir, "-I", includeDir, filepath.Join(dir, "main.cpp"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generated C++ failed to compile:\n%v\n\n--- generated file ---\n%v", string(out), doc)
	}
	if len(out) > 0 {
		t.Logf("compiler output:\n%v", string(out))
	}
}

// TestSmokeGenericFixtureCompiles runs the full module compiler (dtos, enums,
// actions - classic and reactive, headers, path/query params) over the same
// feature-rich fixture the sibling generators' own module tests use, then
// actually compiles every generated header with clang++ - a much stronger
// guarantee than any Go-side string assertion could give.
func TestSmokeGenericFixtureCompiles(t *testing.T) {
	if _, err := exec.LookPath("clang++"); err != nil {
		t.Skip("clang++ not available")
	}

	content, err := os.ReadFile("../../examples/fullstack/definitions.emi.yml")
	if err != nil {
		t.Fatal(err)
	}
	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}

	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}}
	files, err := CppModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatalf("unexpected error compiling the fixture module: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected at least one generated file")
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

	var sb strings.Builder
	sb.WriteString("#include <cstdio>\n")
	for _, f := range files {
		if f.Extension != ".hpp" {
			continue
		}
		fmt.Fprintf(&sb, "#include %q\n", f.Name+f.Extension)
	}
	sb.WriteString("int main() { return 0; }\n")
	if err := os.WriteFile(filepath.Join(dir, "main.cpp"), []byte(sb.String()), 0644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("clang++", "-std=c++17", "-Wall", "-Wextra", "-fsyntax-only",
		"-I", dir, "-I", dir+"/cjson", filepath.Join(dir, "main.cpp"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generated C++ failed to compile:\n%v", string(out))
	}
	if len(out) > 0 {
		t.Logf("compiler output:\n%v", string(out))
	}
}

// TestMediaStreamExampleCompilesAndFramesMatchGo compiles the media-stream
// example's definition (a binary reactive action with path + query params) and
// runs a small program asserting the C++ EmiMediaFrameHeader writes the exact
// bytes emigo.MediaFrame expects (see emigo/MediaFrame_test.go for the Go side).
func TestMediaStreamExampleCompilesAndFramesMatchGo(t *testing.T) {
	if _, err := exec.LookPath("clang++"); err != nil {
		t.Skip("clang++ not available")
	}

	content, err := os.ReadFile("../../examples/media-stream/media.emi.yml")
	if err != nil {
		t.Fatal(err)
	}
	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}
	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}}
	files, err := CppModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	var action string
	for _, f := range files {
		full := filepath.Join(dir, f.Location, f.Name+f.Extension)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(f.ActualScript), 0644); err != nil {
			t.Fatal(err)
		}
		if f.Name == "StreamMediaAction" {
			action = string(f.ActualScript)
		}
	}
	for _, want := range []string{"StreamMediaActionQueryParams", "url = query.Apply(url);", "const StreamMediaActionQueryParams& query", "SendBinary"} {
		if !strings.Contains(action, want) {
			t.Errorf("generated StreamMediaAction.hpp is missing %q", want)
		}
	}

	program := `
#include "StreamMediaAction.hpp"
#include "EmiMediaFrame.hpp"
#include <cstdio>
int main() {
    emi::EmiMediaFrameHeader h;
    h.kind = emi::EmiMediaVideo; h.flags = emi::EmiMediaKeyframe;
    h.seq = 0x01020304; h.timestampMs = 0xAABBCCDD;
    uint8_t b[emi::kEmiMediaFrameHeaderSize];
    h.Write(b);
    const uint8_t want[12] = {2, 1, 0, 0, 1, 2, 3, 4, 0xAA, 0xBB, 0xCC, 0xDD};
    for (int i = 0; i < 12; i++) if (b[i] != want[i]) { std::printf("byte %d: %02x\n", i, b[i]); return 1; }
    emi::EmiMediaFrameHeader r;
    if (!emi::EmiMediaFrameHeader::Read(b, 12, r) || r.seq != h.seq || r.timestampMs != h.timestampMs || r.kind != h.kind) return 2;
    if (emi::EmiMediaFrameHeader::Read(b, 11, r)) return 3;
    StreamMediaActionQueryParams q; q.sampleRate = 16000; q.audioCodec = "pcm16";
    std::string u = q.Apply("ws://h/x");
    if (u != "ws://h/x?sampleRate=16000&audioCodec=pcm16") { std::printf("%s\n", u.c_str()); return 4; }
    return 0;
}
`
	if err := os.WriteFile(filepath.Join(dir, "main.cpp"), []byte(program), 0644); err != nil {
		t.Fatal(err)
	}
	bin := filepath.Join(dir, "frames")
	out, err := exec.Command("clang++", "-std=c++17", "-Wall", "-Wextra", "-I", dir, "-I", dir+"/cjson",
		filepath.Join(dir, "main.cpp"), filepath.Join(dir, "cjson", "cJSON.c"), "-o", bin).CombinedOutput()
	if err != nil {
		t.Fatalf("media example failed to compile:\n%v", string(out))
	}
	if out, err := exec.Command(bin).CombinedOutput(); err != nil {
		t.Fatalf("frame layout check failed: %v\n%s", err, out)
	}
}
