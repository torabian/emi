package c

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestCCommonObjectGenerator_SafeDefaults covers the core promise behind
// cResolveField: every field - no matter its type or nullability - always
// has an `_init` line, so `Type_new()` is always immediately safe to use and
// `Type_free()` is always safe to call on the result.
func TestCCommonObjectGenerator_SafeDefaults(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "requiredString", Type: core.FieldTypeString},
		{Name: "nullableString", Type: core.FieldTypeStringNullable},
		{Name: "requiredInt", Type: core.FieldTypeInt},
		{Name: "requiredBool", Type: core.FieldTypeBool},
		{Name: "items", Type: core.FieldTypeSlice, Primitive: "string"},
		{
			Name: "nested",
			Type: core.FieldTypeObject,
			Fields: []*core.EmiField{
				{Name: "innerValue", Type: core.FieldTypeInt},
			},
		},
		{Name: "relation", Type: core.FieldTypeOne, Target: "OtherDto"},
		{Name: "raw", Type: core.FieldTypeAny},
	}

	chunk, err := CCommonObjectGenerator(fields, core.MicroGenContext{}, CCommonObjectContext{RootTypeName: "Sample"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	cases := []string{
		`char* requiredString;`,
		`self->requiredString = strdup("");`,
		`char* nullableString;`,
		`self->nullableString = NULL;`,
		`int32_t requiredInt;`,
		`self->requiredInt = 0;`,
		`bool requiredBool;`,
		`self->requiredBool = false;`,
		`char** items;`,
		// a locally-nested, non-nullable, non-self-referencing `object`
		// field is eagerly default-constructed.
		`SampleNested* nested;`,
		`self->nested = SampleNested_new();`,
		// a `one` relation is never eagerly default-constructed, even when
		// pointing at a real target - always NULL.
		`OtherDto* relation;`,
		`self->relation = NULL;`,
		// `any`/unresolved falls back to a raw cJSON node, also NULL by
		// default.
		`cJSON* raw;`,
		`self->raw = NULL;`,
	}
	for _, want := range cases {
		if !strings.Contains(script, want) {
			t.Errorf("expected generated script to contain %q, got:\n%v", want, script)
		}
	}

	nestedIdx := strings.Index(script, "typedef struct SampleNested")
	rootIdx := strings.Index(script, "typedef struct Sample {")
	if nestedIdx == -1 || rootIdx == -1 || nestedIdx > rootIdx {
		t.Errorf("expected SampleNested to be declared before Sample, got:\n%v", script)
	}

	if chunk.SuggestedFileName != "sample" || chunk.SuggestedExtension != ".h" {
		t.Errorf("unexpected file naming: %v%v", chunk.SuggestedFileName, chunk.SuggestedExtension)
	}
}

// TestCStringField_NullableRoundTrips guards against a real bug this
// generator once had: cStringField's nullable parameter was silently
// discarded, so every `string?` field always defaulted to `strdup("")` and
// always serialized as a JSON string, never able to represent/round-trip
// `null`.
func TestCStringField_NullableRoundTrips(t *testing.T) {
	fields := []*core.EmiField{{Name: "nickname", Type: core.FieldTypeStringNullable}}
	chunk, err := CCommonObjectGenerator(fields, core.MicroGenContext{}, CCommonObjectContext{RootTypeName: "Widget"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	if !strings.Contains(script, "self->nickname = NULL;") {
		t.Errorf("expected a nullable string to default to NULL, got:\n%v", script)
	}
	if strings.Contains(script, `self->nickname = strdup("");`) {
		t.Errorf("expected a nullable string to never default to an empty string, got:\n%v", script)
	}
	if !strings.Contains(script, "if (self->nickname) { cJSON_AddStringToObject(json, \"nickname\", self->nickname); } else { cJSON_AddNullToObject(json, \"nickname\"); }") {
		t.Errorf("expected nullable string serialization to emit JSON null when unset, got:\n%v", script)
	}
}

// TestCCommonObjectGenerator_EmptyNestedObjectStillGetsAClass covers a
// degenerate but schema-legal case: an `object`/`array` field with no
// `fields:` at all - the type still has to name *something*, and it must
// actually be emitted (with the zero-sized-struct guard) or the reference
// fails to compile.
func TestCCommonObjectGenerator_EmptyNestedObjectStillGetsAClass(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "meta", Type: core.FieldTypeObject},
		{Name: "tags", Type: core.FieldTypeArray},
	}

	chunk, err := CCommonObjectGenerator(fields, core.MicroGenContext{}, CCommonObjectContext{RootTypeName: "Widget"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	for _, want := range []string{
		"typedef struct WidgetMeta",
		"char _reserved;", // the zero-sized-struct guard
		"typedef struct WidgetTags",
		"WidgetMeta* meta;",
		"WidgetTags* tags;",
		"size_t tags_count;",
	} {
		if !strings.Contains(script, want) {
			t.Errorf("expected %q in generated script:\n%v", want, script)
		}
	}
}

// TestCSelfReferencingStruct_UsesStructTag guards against a real compile
// error this generator once had: a field whose target is the very struct
// being declared (a self-referencing relation or collection) referenced the
// bare typedef alias, which doesn't exist yet inside the struct's own body -
// it must use the `struct X` tag form instead.
func TestCSelfReferencingStruct_UsesStructTag(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "parent", Type: core.FieldTypeOne, Target: "Node"},
		{Name: "children", Type: core.FieldTypeCollection, Target: "Node"},
	}
	chunk, err := CCommonObjectGenerator(fields, core.MicroGenContext{}, CCommonObjectContext{RootTypeName: "Node"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	if !strings.Contains(script, "struct Node* parent;") {
		t.Errorf("expected a self-referencing `one` field to use the `struct X` tag, got:\n%v", script)
	}
	if !strings.Contains(script, "struct Node* children;") {
		t.Errorf("expected a self-referencing `collection` field to use the `struct X` tag, got:\n%v", script)
	}
	// forward declarations must exist so _free_fields/_from_json_into can
	// call _free/_from_json before they're defined further down the file.
	if !strings.Contains(script, "static inline void Node_free(Node* self);") {
		t.Errorf("expected a forward declaration for Node_free, got:\n%v", script)
	}
}

// TestCEnumField covers every shape an `enum`/`enum?` field can take: a
// `target:` reference to a module-level enum, inline `of:` values (a
// flattened companion enum), and neither - which, unlike every other
// sibling language, still needs a real (if degenerate) declared enum type,
// since C has no plain-string fallback for an enum-typed field.
func TestCEnumField(t *testing.T) {
	t.Run("target reference defaults to ordinal 0", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "status", Type: core.FieldTypeEnum, Target: "StatusEnum"}}
		chunk, err := CCommonObjectGenerator(fields, core.MicroGenContext{}, CCommonObjectContext{RootTypeName: "Task"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "StatusEnum status;") {
			t.Errorf("expected a StatusEnum status; member, got:\n%v", script)
		}
		if !strings.Contains(script, "self->status = (StatusEnum) 0;") {
			t.Errorf("expected an ordinal-0 default, got:\n%v", script)
		}
	})

	t.Run("inline of: values become a generated companion enum", func(t *testing.T) {
		fields := []*core.EmiField{{
			Name: "kind",
			Type: core.FieldTypeEnum,
			OfType: []*core.EmiEnumInline{
				{Key: "cat"},
				{Key: "dog"},
			},
		}}
		chunk, err := CCommonObjectGenerator(fields, core.MicroGenContext{}, CCommonObjectContext{RootTypeName: "Pet"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "typedef enum PetKind {") {
			t.Errorf("expected a companion PetKind enum, got:\n%v", script)
		}
		if !strings.Contains(script, "PET_KIND_CAT = 0") {
			t.Errorf("expected a screaming-snake-case, self-prefixed case identifier, got:\n%v", script)
		}
	})

	t.Run("neither target nor of: still gets a real (degenerate) enum", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "freeform", Type: core.FieldTypeEnum}}
		chunk, err := CCommonObjectGenerator(fields, core.MicroGenContext{}, CCommonObjectContext{RootTypeName: "Loose"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "typedef enum LooseFreeform {") {
			t.Errorf("expected a synthetic LooseFreeform enum, got:\n%v", script)
		}
		if !strings.Contains(script, "LOOSE_FREEFORM_UNSPECIFIED = 0") {
			t.Errorf("expected a single _UNSPECIFIED placeholder case (C forbids empty enums), got:\n%v", script)
		}
	})
}

// TestCModuleFullVirtualFiles_Fixture is a broad smoke test: running the
// full module compiler over a real, feature-rich fixture (nested objects,
// slices, one/collection relations, complex fields, path/query params,
// reactive actions, envelopes...) must never error, and must produce
// non-empty, header-guarded C for every dto/action/enum plus the embedded
// runtime (vendored cJSON + fetchx.h) and README.
func TestCModuleFullVirtualFiles_Fixture(t *testing.T) {
	content, err := os.ReadFile("../../examples/fullstack/definitions.emi.yml")
	if err != nil {
		t.Fatal(err)
	}

	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}

	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}}

	files, err := CModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatalf("unexpected error compiling the fixture module: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected at least one generated file")
	}

	sawCjson, sawFetchx, sawReadme := false, false, false
	for _, f := range files {
		if strings.Contains(f.Name, "cJSON") {
			sawCjson = true
		}
		if strings.Contains(f.Name, "fetchx") {
			sawFetchx = true
		}
		if f.Name == "README" && f.Extension == ".md" {
			sawReadme = true
		}
		// only the module compiler's own dto/enum/action outputs go
		// through AsFullDocument (header guard + common includes) - the
		// embedded runtime files (cjson/*, fetchx.h) are vendored/hand-
		// written as-is (SuggestedExtension is empty; the extension is
		// already part of their embedded file name) and carry their own
		// guards already.
		if f.Extension == ".h" {
			if strings.TrimSpace(f.ActualScript) == "" {
				t.Errorf("generated empty header: %v/%v", f.Location, f.Name)
			}
			if !strings.Contains(f.ActualScript, "#ifndef EMISDK_") || !strings.Contains(f.ActualScript, "#define EMISDK_") {
				t.Errorf("expected %v.h to have a header guard, got:\n%v", f.Name, f.ActualScript)
			}
			if !strings.Contains(f.ActualScript, `#include "cjson/cJSON.h"`) {
				t.Errorf("expected %v.h to include cJSON.h, got:\n%v", f.Name, f.ActualScript)
			}
		}
	}
	if !sawCjson {
		t.Error("expected the vendored cJSON library to be included")
	}
	if !sawFetchx {
		t.Error("expected the embedded fetchx.h runtime to be included")
	}
	if !sawReadme {
		t.Error("expected a README.md at the package root")
	}
}

// TestBinaryReactiveWebsocket compiles the websocket flavour of a reactive
// action (`binaryType: arraybuffer`) with a plain C compiler and no libcurl
// (-DFETCHX_NO_CURL, as on ESP-IDF), then checks the pieces a device depends on:
// the RFC 6455 accept-key vector, the media frame layout shared with the Go,
// C++ and JS sides, typed query/path url building and the portable int format.
func TestBinaryReactiveWebsocket(t *testing.T) {
	cc, err := exec.LookPath("cc")
	if err != nil {
		t.Skip("no C compiler available")
	}

	module := &core.Emi{Actions: []*core.EmiAction{{
		Name: "streamMedia", Url: "/media/stream/:deviceId", Method: "reactive", BinaryType: "arraybuffer",
		Query: []*core.EmiQueryField{
			{Name: "sampleRate", Type: core.FieldTypeInt},
			{Name: "audioCodec", Type: core.FieldTypeString},
		},
	}}}
	files, err := CModuleFullVirtualFiles(module, core.MicroGenContext{Flags: map[string]string{}})
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	var header string
	for _, f := range files {
		full := filepath.Join(dir, f.Location, f.Name+f.Extension)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(f.ActualScript), 0644); err != nil {
			t.Fatal(err)
		}
		if f.Name == "stream_media_action" {
			header = string(f.ActualScript)
		}
	}
	for _, want := range []string{"EmiWebSocket* ws, EmiStream* stream", `#include "emiws.h"`, "emiws_connect", `"%lld", (long long) self->sampleRate`} {
		if !strings.Contains(header, want) {
			t.Errorf("generated header is missing %q", want)
		}
	}
	if strings.Contains(header, "FetchxEventList") {
		t.Errorf("binary reactive action must not use the SSE/libcurl path")
	}

	program := `
#include "stream_media_action.h"
#include <stdio.h>
int main(void) {
    /* RFC 6455 section 1.3 example */
    uint8_t digest[20]; char accept[32];
    const char* in = "dGhlIHNhbXBsZSBub25jZQ==258EAFA5-E914-47DA-95CA-C5AB0DC85B11";
    emiws_sha1((const uint8_t*) in, strlen(in), digest);
    emiws_base64(digest, 20, accept);
    if (strcmp(accept, "s3pPLMBiTxaQ9kYGzzhZRbK+xOo=") != 0) { printf("accept %s\n", accept); return 1; }

    EmiMediaHeader h = { EMIMEDIA_VIDEO, EMIMEDIA_KEYFRAME, 0x01020304u, 0xAABBCCDDu };
    uint8_t b[EMIMEDIA_HEADER_SIZE];
    const uint8_t want[12] = {2, 1, 0, 0, 1, 2, 3, 4, 0xAA, 0xBB, 0xCC, 0xDD};
    emimedia_header_write(&h, b);
    if (memcmp(b, want, 12) != 0) return 2;
    EmiMediaHeader r;
    if (emimedia_header_read(b, 12, &r) != 0 || r.seq != h.seq || r.timestamp_ms != h.timestamp_ms) return 3;
    if (emimedia_header_read(b, 11, &r) == 0) return 4;

    StreamMediaActionPathParameters p; StreamMediaActionPathParameters_init(&p); p.deviceId = strdup("dev1");
    StreamMediaActionQueryParams q; StreamMediaActionQueryParams_init(&q);
    q.sampleRate_set = true; q.sampleRate = 16000; q.audioCodec = strdup("pcm16");
    char* u = fetchx_build_url("http://h:1", "/media/stream/:deviceId");
    char* u1 = StreamMediaActionPathParameters_apply(&p, u);
    char* u2 = StreamMediaActionQueryParams_apply(&q, u1);
    char* ws = emiws_http_to_ws(u2);
    if (strcmp(ws, "ws://h:1/media/stream/dev1?sampleRate=16000&audioCodec=pcm16") != 0) { printf("%s\n", ws); return 5; }
    return 0;
}
`
	if err := os.WriteFile(filepath.Join(dir, "main.c"), []byte(program), 0644); err != nil {
		t.Fatal(err)
	}
	bin := filepath.Join(dir, "ws")
	out, err := exec.Command(cc, "-std=c99", "-D_POSIX_C_SOURCE=200809L", "-DFETCHX_NO_CURL", "-Wall", "-Wextra", "-Werror",
		"-I", dir, filepath.Join(dir, "main.c"), filepath.Join(dir, "cjson", "cJSON.c"), "-o", bin).CombinedOutput()
	if err != nil {
		t.Fatalf("generated C failed to compile:\n%s", out)
	}
	if out, err := exec.Command(bin).CombinedOutput(); err != nil {
		t.Fatalf("websocket/media checks failed: %v\n%s", err, out)
	}
}
