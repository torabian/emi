package cpp

import (
	"os"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestCppGenericCommonObjectGenerator_SafeDefaults covers the core promise
// behind cppGenericResolveField: every field gets a real declaration and both a
// ToJson() and a FromJson() contribution, with nullability represented as
// std::optional<T> (never silently collapsed to a zero value the way the C
// target's own non-nullable-int scope boundary works).
func TestCppGenericCommonObjectGenerator_SafeDefaults(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "requiredString", Type: core.FieldTypeString},
		{Name: "nullableString", Type: core.FieldTypeStringNullable},
		{Name: "requiredInt", Type: core.FieldTypeInt},
		{Name: "nullableInt", Type: core.FieldTypeIntNullable},
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

	chunk, err := CppCommonObjectGenerator(fields, core.MicroGenContext{}, CppCommonObjectContext{
		Dialect:       DialectGeneric,
		RootClassName: "Sample",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	cases := []string{
		`std::string requiredString;`,
		`cJSON_AddStringToObject(json, "requiredString", requiredString.c_str());`,
		`std::optional<std::string> nullableString;`,
		`int32_t requiredInt;`,
		`std::optional<int32_t> nullableInt;`,
		`bool requiredBool;`,
		`std::vector<std::string> items;`,
		// a locally-nested, non-nullable, non-self-referencing `object` field is
		// held by value.
		`SampleNested nested;`,
		// a `one` relation is never eagerly constructed, even pointing at a real
		// target - always a null std::unique_ptr.
		`std::unique_ptr<OtherDto> relation;`,
		// `any`/unresolved falls back to the dynamic emi::EmiJson wrapper.
		`emi::EmiJson raw;`,
	}
	for _, want := range cases {
		if !strings.Contains(script, want) {
			t.Errorf("expected generated script to contain %q, got:\n%v", want, script)
		}
	}

	nestedIdx := strings.Index(script, "class SampleNested")
	rootIdx := strings.Index(script, "class Sample {")
	if nestedIdx == -1 || rootIdx == -1 || nestedIdx > rootIdx {
		t.Errorf("expected SampleNested to be declared before Sample, got:\n%v", script)
	}

	if chunk.SuggestedFileName != "Sample" || chunk.SuggestedExtension != ".hpp" {
		t.Errorf("unexpected file naming: %v%v", chunk.SuggestedFileName, chunk.SuggestedExtension)
	}
}

// TestCppGenericNullableString_RoundTripsViaOptional guards against the same
// class of bug the C target's own TestCStringField_NullableRoundTrips test
// documents: a nullable field must actually round-trip absence (JSON null),
// not silently collapse to a present-but-empty value.
func TestCppGenericNullableString_RoundTripsViaOptional(t *testing.T) {
	fields := []*core.EmiField{{Name: "nickname", Type: core.FieldTypeStringNullable}}
	chunk, err := CppCommonObjectGenerator(fields, core.MicroGenContext{}, CppCommonObjectContext{
		Dialect:       DialectGeneric,
		RootClassName: "Widget",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	if !strings.Contains(script, "std::optional<std::string> nickname;") {
		t.Errorf("expected a nullable string field, got:\n%v", script)
	}
	if !strings.Contains(script, `if (nickname.has_value()) { cJSON_AddStringToObject(json, "nickname", nickname->c_str()); } else { cJSON_AddNullToObject(json, "nickname"); }`) {
		t.Errorf("expected nullable string serialization to emit JSON null when unset, got:\n%v", script)
	}
}

// TestCppSelfReferencingClass_UsesUniquePtr guards against a real class of bug
// the C target's own TestCSelfReferencingStruct_UsesStructTag test documents: a
// field whose target is the very class being declared can't be a by-value
// member (an incomplete type), so it must go through indirection instead.
func TestCppSelfReferencingClass_UsesUniquePtr(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "parent", Type: core.FieldTypeOne, Target: "Node"},
		{Name: "children", Type: core.FieldTypeCollection, Target: "Node"},
	}
	chunk, err := CppCommonObjectGenerator(fields, core.MicroGenContext{}, CppCommonObjectContext{
		Dialect:       DialectGeneric,
		RootClassName: "Node",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	if !strings.Contains(script, "std::unique_ptr<Node> parent;") {
		t.Errorf("expected a self-referencing `one` field to use std::unique_ptr, got:\n%v", script)
	}
	// A self-referencing `collection` field is fine as a plain std::vector<Node>
	// (C++17 guarantees vector/list/forward_list tolerate an incomplete element
	// type at declaration point, same trick a hand-written recursive tree struct
	// uses) - no indirection needed there, unlike `parent` above.
	if !strings.Contains(script, "std::vector<Node> children;") {
		t.Errorf("expected a self-referencing `collection` field to be a plain std::vector<Node>, got:\n%v", script)
	}
}

// TestCppEnumField covers every shape an `enum`/`enum?` field can take: a
// `target:` reference to a module-level enum, and inline `of:` values (a
// flattened companion enum).
func TestCppEnumField(t *testing.T) {
	t.Run("target reference", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "status", Type: core.FieldTypeEnum, Target: "StatusEnum"}}
		chunk, err := CppCommonObjectGenerator(fields, core.MicroGenContext{}, CppCommonObjectContext{
			Dialect: DialectGeneric, RootClassName: "Task",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "StatusEnum status;") {
			t.Errorf("expected a StatusEnum status; member, got:\n%v", script)
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
		chunk, err := CppCommonObjectGenerator(fields, core.MicroGenContext{}, CppCommonObjectContext{
			Dialect: DialectGeneric, RootClassName: "Pet",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "enum class PetKind : int32_t {") {
			t.Errorf("expected a companion PetKind enum, got:\n%v", script)
		}
		if !strings.Contains(script, "Cat = 0,") {
			t.Errorf("expected a PascalCase case identifier, got:\n%v", script)
		}
		if !strings.Contains(script, "inline const char* PetKindToString(PetKind value)") ||
			!strings.Contains(script, "inline PetKind PetKindFromString(const std::string& value)") {
			t.Errorf("expected hand-generated to/from-string helpers, got:\n%v", script)
		}
	})
}

// TestCppUnrealDialect_GeneratesReflectedUStruct covers the unreal dialect's
// core promise: every ordinary field becomes a real UPROPERTY on an
// `F`-prefixed USTRUCT(BlueprintType), with no hand-generated
// (de)serialization code at all (unlike the generic dialect) - Unreal's own
// reflection does that job (see the package doc comment on
// cpp-field-plan-unreal.go).
func TestCppUnrealDialect_GeneratesReflectedUStruct(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "title", Type: core.FieldTypeString},
		{Name: "score", Type: core.FieldTypeFloat64Nullable},
		{Name: "tags", Type: core.FieldTypeSlice, Primitive: "string"},
		{Name: "target", Type: core.FieldTypeOne, Target: "OtherDto"},
	}

	chunk, err := CppCommonObjectGenerator(fields, core.MicroGenContext{}, CppCommonObjectContext{
		Dialect:       DialectUnreal,
		RootClassName: "WidgetDto",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	for _, want := range []string{
		"USTRUCT(BlueprintType)",
		"struct FWidgetDto {",
		"GENERATED_BODY()",
		`UPROPERTY(BlueprintReadWrite, EditAnywhere, Category = "Emi")`,
		"FString Title;",
		// A nullable field is a paired value + bool flag (both real UPROPERTYs,
		// UE4-compatible), never a UPROPERTY-reflected TOptional<T> - see
		// ueFieldMaybeNullable's doc comment.
		"bool bScoreIsSet = false;",
		"double Score;",
		"TArray<FString> Tags;",
		// a `one` relation embeds the target USTRUCT by value, the idiomatic
		// Unreal shape (unlike the generic dialect's std::unique_ptr) - see
		// ueValueField's doc comment.
		"FOtherDto Target;",
	} {
		if !strings.Contains(script, want) {
			t.Errorf("expected %q in generated script:\n%v", want, script)
		}
	}
	if strings.Contains(script, "TOptional") {
		t.Errorf("did not expect a UPROPERTY-reflected TOptional<T> anywhere (UE4-incompatible), got:\n%v", script)
	}
	// No hand-generated (de)serialization glue - Unreal's own reflection does it.
	if strings.Contains(script, "ToJson()") || strings.Contains(script, "cJSON") {
		t.Errorf("did not expect any hand-generated JSON glue in the unreal dialect, got:\n%v", script)
	}

	if chunk.SuggestedFileName != "FWidgetDto" || chunk.SuggestedExtension != ".h" {
		t.Errorf("unexpected file naming: %v%v", chunk.SuggestedFileName, chunk.SuggestedExtension)
	}
}

// TestCppUnrealAsFullDocument_GeneratedHeaderIsLast covers Unreal Header Tool's
// own hard requirement: `#include "<FileName>.generated.h"` must be the very
// last line of the header, after every USTRUCT/UENUM/UPROPERTY it reflects.
func TestCppUnrealAsFullDocument_GeneratedHeaderIsLast(t *testing.T) {
	chunk, err := CppCommonObjectGenerator([]*core.EmiField{{Name: "value", Type: core.FieldTypeInt}},
		core.MicroGenContext{}, CppCommonObjectContext{Dialect: DialectUnreal, RootClassName: "Sample"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	doc := CppUnrealAsFullDocument(chunk, chunk.SuggestedFileName)

	trimmed := strings.TrimRight(doc, "\n")
	if !strings.HasSuffix(trimmed, `#include "FSample.generated.h"`) {
		t.Errorf("expected #include \"FSample.generated.h\" to be the last line, got:\n%v", doc)
	}
	if !strings.Contains(doc, "#pragma once") || !strings.Contains(doc, `#include "CoreMinimal.h"`) {
		t.Errorf("expected #pragma once and CoreMinimal.h, got:\n%v", doc)
	}
}

// TestCppModuleFullVirtualFiles_Fixture is a broad smoke test: running the full
// module compiler over a real, feature-rich fixture (nested objects, slices,
// one/collection relations, complex fields, path/query params, reactive
// actions...) must never error, for either dialect, and must produce non-empty
// output plus the embedded runtime and README.
func TestCppModuleFullVirtualFiles_Fixture(t *testing.T) {
	content, err := os.ReadFile("../../examples/fullstack/definitions.emi.yml")
	if err != nil {
		t.Fatal(err)
	}
	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}

	for _, dialect := range []Dialect{DialectGeneric, DialectUnreal} {
		t.Run(string(dialect), func(t *testing.T) {
			ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{"dialect": string(dialect)}}
			files, err := CppModuleFullVirtualFiles(&module, ctx)
			if err != nil {
				t.Fatalf("unexpected error compiling the fixture module: %v", err)
			}
			if len(files) == 0 {
				t.Fatal("expected at least one generated file")
			}

			sawReadme := false
			sawRuntime := false
			for _, f := range files {
				if f.Name == "README" && f.Extension == ".md" {
					sawReadme = true
				}
				if strings.Contains(f.Name, "EmiHttpTransport") || strings.Contains(f.Name, "EmiHttpClient") {
					sawRuntime = true
				}
				if strings.TrimSpace(f.ActualScript) == "" {
					t.Errorf("generated empty file: %v/%v%v", f.Location, f.Name, f.Extension)
				}
			}
			if !sawReadme {
				t.Error("expected a README.md at the package root")
			}
			if !sawRuntime {
				t.Error("expected the embedded runtime to be included")
			}
		})
	}
}

// TestResolveDialect covers the --dialect flag and the `unreal` compiler tag,
// and that neither being set defaults to the generic dialect.
func TestResolveDialect(t *testing.T) {
	cases := []struct {
		name string
		ctx  core.MicroGenContext
		want Dialect
	}{
		{"default", core.MicroGenContext{}, DialectGeneric},
		{"dialect flag", core.MicroGenContext{Flags: map[string]string{"dialect": "unreal"}}, DialectUnreal},
		{"ue5 alias", core.MicroGenContext{Flags: map[string]string{"dialect": "ue5"}}, DialectUnreal},
		{"tag", core.MicroGenContext{Tags: "unreal"}, DialectUnreal},
		{"arduino alias stays generic", core.MicroGenContext{Flags: map[string]string{"dialect": "arduino"}}, DialectGeneric},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ResolveDialect(c.ctx); got != c.want {
				t.Errorf("ResolveDialect() = %v, want %v", got, c.want)
			}
		})
	}
}
