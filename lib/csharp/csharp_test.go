package csharp

import (
	"os"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestCSharpCommonObjectGenerator_SafeDefaults covers the core promise
// behind csResolveField: every property - no matter its type or nullability
// - always has either a safe initializer or is nullable, so
// `new SomeDto()` is always valid and never triggers a
// nullable-reference-types (CS8618) warning.
func TestCSharpCommonObjectGenerator_SafeDefaults(t *testing.T) {
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

	chunk, err := CSharpCommonObjectGenerator(fields, core.MicroGenContext{}, CSharpCommonObjectContext{RootClassName: "Sample"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	cases := []string{
		`public string RequiredString { get; set; } = "";`,
		`public string? NullableString { get; set; }`,
		`public int RequiredInt { get; set; } = 0;`,
		`public bool RequiredBool { get; set; } = false;`,
		`public List<string> Items { get; set; } = new();`,
		`public SampleNested Nested { get; set; } = new();`,
		// a `one` relation is never eagerly default-constructed, even when
		// required - always a plain nullable property.
		`public OtherDto? Relation { get; set; }`,
		// `any`/`complex` has no sensible non-null default under NRT.
		`public object? Raw { get; set; }`,
	}
	for _, want := range cases {
		if !strings.Contains(script, want) {
			t.Errorf("expected generated script to contain %q, got:\n%v", want, script)
		}
	}

	nestedIdx := strings.Index(script, "class SampleNested")
	rootIdx := strings.Index(script, "class Sample\n")
	if nestedIdx == -1 || rootIdx == -1 || nestedIdx > rootIdx {
		t.Errorf("expected SampleNested to be declared before Sample, got:\n%v", script)
	}

	if chunk.SuggestedFileName != "Sample" || chunk.SuggestedExtension != ".cs" {
		t.Errorf("unexpected file naming: %v%v", chunk.SuggestedFileName, chunk.SuggestedExtension)
	}
}

// TestCSharpCommonObjectGenerator_EmptyNestedObjectStillGetsAClass covers a
// degenerate but schema-legal case: an `object`/`array` field with no
// `fields:` at all - the type still has to name *something*, and it must
// actually be emitted or the reference fails to compile.
func TestCSharpCommonObjectGenerator_EmptyNestedObjectStillGetsAClass(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "meta", Type: core.FieldTypeObject},
		{Name: "tags", Type: core.FieldTypeArray},
	}

	chunk, err := CSharpCommonObjectGenerator(fields, core.MicroGenContext{}, CSharpCommonObjectContext{RootClassName: "Widget"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	for _, want := range []string{"class WidgetMeta", "class WidgetTags", "public WidgetMeta Meta", "public List<WidgetTags> Tags"} {
		if !strings.Contains(script, want) {
			t.Errorf("expected %q in generated script:\n%v", want, script)
		}
	}
}

// TestCSharpFieldType_MissingTargetFallsBackToObject covers a malformed
// (but not generator-crashing) `one`/`collection` field with no `target:` -
// it must never reference nestedClassName, since object/array are the only
// field types that ever actually get a flattened class generated for them.
func TestCSharpFieldType_MissingTargetFallsBackToObject(t *testing.T) {
	one := &core.EmiField{Name: "owner", Type: core.FieldTypeOne}
	if got := csFieldType(one, "RootOwner", "Root"); got != "object?" {
		t.Errorf("expected targetless `one` to fall back to object?, got %q", got)
	}

	collection := &core.EmiField{Name: "items", Type: core.FieldTypeCollection}
	if got := csFieldType(collection, "RootItems", "Root"); got != "List<object?>" {
		t.Errorf("expected targetless `collection` to fall back to List<object?>, got %q", got)
	}
}

// TestCSharpFieldType_GolangOnlyListAndClass covers `_list`/`class` -
// documented as "golang-only for now" and never special-cased by any other
// sibling generator either, but still part of
// core.GetEmiFieldTypeCatalog().DtoFieldTypes.
func TestCSharpFieldType_GolangOnlyListAndClass(t *testing.T) {
	list := &core.EmiField{Name: "items", Type: core.FieldTypeList}
	if got := csFieldType(list, "RootItems", "Root"); got != "List<RootItems>" {
		t.Errorf("expected `_list` to render like `array`, got %q", got)
	}

	class := &core.EmiField{Name: "owner", Type: core.FieldTypeClass, Target: "UserDto"}
	if got := csFieldType(class, "RootOwner", "Root"); got != "UserDto?" {
		t.Errorf("expected `class` to render like `one`, got %q", got)
	}
}

// TestCSharpEnumField covers every shape an `enum`/`enum?` field can take: a
// `target:` reference to a module-level enum, inline `of:` values (a
// generated companion enum with its own JsonConverter), and neither (a
// plain string).
func TestCSharpEnumField(t *testing.T) {
	t.Run("target reference is a real enum property with no initializer needed", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "status", Type: core.FieldTypeEnum, Target: "StatusEnum"}}
		chunk, err := CSharpCommonObjectGenerator(fields, core.MicroGenContext{}, CSharpCommonObjectContext{RootClassName: "Task"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "public StatusEnum Status { get; set; }") {
			t.Errorf("expected a plain (non-nullable, no-initializer) StatusEnum property, got:\n%v", script)
		}
	})

	t.Run("inline of: values become a generated companion enum with its own converter", func(t *testing.T) {
		fields := []*core.EmiField{{
			Name: "kind",
			Type: core.FieldTypeEnum,
			OfType: []*core.EmiEnumInline{
				{Key: "cat"},
				{Key: "dog"},
			},
		}}
		chunk, err := CSharpCommonObjectGenerator(fields, core.MicroGenContext{}, CSharpCommonObjectContext{RootClassName: "Pet"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "public enum PetKind") {
			t.Errorf("expected a companion PetKind enum, got:\n%v", script)
		}
		if !strings.Contains(script, "class PetKindConverter") {
			t.Errorf("expected a PetKindConverter class, got:\n%v", script)
		}
		if !strings.Contains(script, `["cat"] = PetKind.Cat`) {
			t.Errorf(`expected the converter to map "cat" to PetKind.Cat, got:`+"\n%v", script)
		}
	})

	t.Run("neither target nor of: falls back to plain string", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "freeform", Type: core.FieldTypeEnum}}
		chunk, err := CSharpCommonObjectGenerator(fields, core.MicroGenContext{}, CSharpCommonObjectContext{RootClassName: "Loose"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, `public string Freeform { get; set; } = "";`) {
			t.Errorf("expected freeform as a plain defaulted string, got:\n%v", script)
		}
	})
}

// TestCSharpModuleFullVirtualFiles_Fixture is a broad smoke test: running
// the full module compiler over a real, feature-rich fixture (nested
// objects, slices, one/collection relations, complex fields, path/query
// params, reactive actions, envelopes...) must never error, and must
// produce non-empty C# for every dto/action plus the embedded runtime and
// .csproj.
func TestCSharpModuleFullVirtualFiles_Fixture(t *testing.T) {
	content, err := os.ReadFile("../../examples/fullstack/definitions.emi.yml")
	if err != nil {
		t.Fatal(err)
	}

	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}

	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}}

	files, err := CSharpModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatalf("unexpected error compiling the fixture module: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected at least one generated file")
	}

	sawRuntime, sawCsproj := false, false
	for _, f := range files {
		if strings.Contains(f.Name, "Fetchx") {
			sawRuntime = true
		}
		if f.Extension == ".csproj" {
			sawCsproj = true
		}
		if f.Extension == ".cs" && strings.TrimSpace(f.ActualScript) == "" {
			t.Errorf("generated empty C# file: %v/%v", f.Location, f.Name)
		}
	}
	if !sawRuntime {
		t.Error("expected the embedded runtime (Fetchx) to be included in the output")
	}
	if !sawCsproj {
		t.Error("expected a .csproj to be included in the output")
	}
}
