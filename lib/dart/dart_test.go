package dart

import (
	"os"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestDartCommonObjectGenerator_SafeDefaults covers the core promise behind
// dartResolveField: every field - no matter its type or nullability - always
// gets a valid, compile-time-legal constructor entry, so `SomeDto()` is
// always constructible without any required-argument errors.
func TestDartCommonObjectGenerator_SafeDefaults(t *testing.T) {
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
	}

	chunk, err := DartCommonObjectGenerator(fields, core.MicroGenContext{}, DartCommonObjectContext{RootClassName: "Sample"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	cases := []string{
		"String requiredString;",
		"String? nullableString;",
		"int requiredInt;",
		"bool requiredBool;",
		"this.requiredString = '',",
		"List<String>? items,",
		"items = items ?? <String>[]",
		"SampleNested? nested,",
		"nested = nested ?? SampleNested()",
		// a `one` relation is never eagerly default-constructed, even when
		// required - always a plain nullable `this.` param.
		"OtherDto? relation;",
		"this.relation,",
	}
	for _, want := range cases {
		if !strings.Contains(script, want) {
			t.Errorf("expected generated script to contain %q, got:\n%v", want, script)
		}
	}

	nestedIdx := strings.Index(script, "class SampleNested {")
	rootIdx := strings.Index(script, "class Sample {")
	if nestedIdx == -1 || rootIdx == -1 || nestedIdx > rootIdx {
		t.Errorf("expected SampleNested to be declared before Sample, got:\n%v", script)
	}

	if chunk.SuggestedFileName != "sample" || chunk.SuggestedExtension != ".dart" {
		t.Errorf("unexpected file naming: %v%v", chunk.SuggestedFileName, chunk.SuggestedExtension)
	}
}

// TestDartCommonObjectGenerator_EmptyNestedObjectStillGetsAClass covers a
// degenerate but schema-legal case: an `object`/`array` field with no
// `fields:` at all. The type still has to name *something*, and it must
// actually be emitted (even as an empty class) or the reference fails to
// compile.
func TestDartCommonObjectGenerator_EmptyNestedObjectStillGetsAClass(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "meta", Type: core.FieldTypeObject},
		{Name: "tags", Type: core.FieldTypeArray},
	}

	chunk, err := DartCommonObjectGenerator(fields, core.MicroGenContext{}, DartCommonObjectContext{RootClassName: "Widget"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	for _, want := range []string{"class WidgetMeta {", "class WidgetTags {", "WidgetMeta meta;", "List<WidgetTags> tags;"} {
		if !strings.Contains(script, want) {
			t.Errorf("expected %q in generated script:\n%v", want, script)
		}
	}
}

// TestDartFieldTypeOnNestedClasses_MissingTargetFallsBackToDynamic covers a
// malformed (but not generator-crashing) `one`/`collection` field with no
// `target:` - it must never reference nestedClassName, since object/array
// are the only field types that ever actually get a flattened class
// generated for them.
func TestDartFieldTypeOnNestedClasses_MissingTargetFallsBackToDynamic(t *testing.T) {
	one := &core.EmiField{Name: "owner", Type: core.FieldTypeOne}
	if got := dartFieldType(one, "RootOwner", "Root"); got != "dynamic" {
		t.Errorf("expected targetless `one` to fall back to dynamic, got %q", got)
	}

	collection := &core.EmiField{Name: "items", Type: core.FieldTypeCollection}
	if got := dartFieldType(collection, "RootItems", "Root"); got != "List<dynamic>" {
		t.Errorf("expected targetless `collection` to fall back to List<dynamic>, got %q", got)
	}
}

// TestDartFieldTypeOnNestedClasses_GolangOnlyListAndClass covers `_list`/
// `class` - documented as "golang-only for now" and never special-cased by
// any other sibling generator either, but still part of
// core.GetEmiFieldTypeCatalog().DtoFieldTypes.
func TestDartFieldTypeOnNestedClasses_GolangOnlyListAndClass(t *testing.T) {
	list := &core.EmiField{Name: "items", Type: core.FieldTypeList}
	if got := dartFieldType(list, "RootItems", "Root"); got != "List<RootItems>" {
		t.Errorf("expected `_list` to render like `array`, got %q", got)
	}

	class := &core.EmiField{Name: "owner", Type: core.FieldTypeClass, Target: "UserDto"}
	if got := dartFieldType(class, "RootOwner", "Root"); got != "UserDto?" {
		t.Errorf("expected `class` to render like `one`, got %q", got)
	}
}

// TestDartEnumField covers every shape an `enum`/`enum?` field can take: a
// `target:` reference to a module-level enum, inline `of:` values (a
// generated companion enum), and neither (a plain String).
func TestDartEnumField(t *testing.T) {
	t.Run("target reference imports the enum class and defaults to its first value", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "status", Type: core.FieldTypeEnum, Target: "StatusEnum"}}
		chunk, err := DartCommonObjectGenerator(fields, core.MicroGenContext{}, DartCommonObjectContext{RootClassName: "Task"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "StatusEnum status;") {
			t.Errorf("expected status: StatusEnum status;, got:\n%v", script)
		}
		if !strings.Contains(script, "status = status ?? StatusEnum.values.first") {
			t.Errorf("expected a first-value default initializer, got:\n%v", script)
		}
		found := false
		for _, dep := range chunk.CodeChunkDependensies {
			if dep.Location == "status_enum.dart" {
				found = true
			}
		}
		if !found {
			t.Errorf("expected an import for status_enum.dart, got %+v", chunk.CodeChunkDependensies)
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
		chunk, err := DartCommonObjectGenerator(fields, core.MicroGenContext{}, DartCommonObjectContext{RootClassName: "Pet"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "enum PetKind {") {
			t.Errorf("expected a companion PetKind enum, got:\n%v", script)
		}
		if !strings.Contains(script, `catKind("cat")`) && !strings.Contains(script, `cat("cat")`) {
			// case identifiers are lowerCamelCase of the raw key
			if !strings.Contains(script, `cat("cat")`) {
				t.Errorf(`expected a cat("cat") enum case, got:`+"\n%v", script)
			}
		}
	})

	t.Run("neither target nor of: falls back to plain String", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "freeform", Type: core.FieldTypeEnum}}
		chunk, err := DartCommonObjectGenerator(fields, core.MicroGenContext{}, DartCommonObjectContext{RootClassName: "Loose"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "String freeform;") || !strings.Contains(script, "this.freeform = '',") {
			t.Errorf("expected freeform as a plain defaulted String, got:\n%v", script)
		}
	})
}

// TestCombineDartImports covers the import-rendering: every Location is
// already a full relative-import path, deduplicated and sorted.
func TestCombineDartImports(t *testing.T) {
	chunk := core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "user_dto.dart"},
			{Location: "user_dto.dart"},
			{Location: "runtime/fetchx.dart"},
		},
	}

	got := CombineDartImports(chunk)
	want := "import 'runtime/fetchx.dart';\nimport 'user_dto.dart';"
	if got != want {
		t.Errorf("unexpected import block:\ngot:  %q\nwant: %q", got, want)
	}
}

// TestDartModuleFullVirtualFiles_Fixture is a broad smoke test: running the
// full module compiler over a real, feature-rich fixture (nested objects,
// slices, one/collection relations, complex fields, path/query params,
// reactive actions, envelopes...) must never error, and must produce
// non-empty dart for every dto/action plus the embedded runtime and
// pubspec.yaml, laid out as a real dart package (dto/action/runtime under
// lib/, pubspec.yaml at the root).
func TestDartModuleFullVirtualFiles_Fixture(t *testing.T) {
	content, err := os.ReadFile("../../examples/fullstack/definitions.emi.yml")
	if err != nil {
		t.Fatal(err)
	}

	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}

	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}}

	files, err := DartModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatalf("unexpected error compiling the fixture module: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected at least one generated file")
	}

	sawRuntime, sawPubspec := false, false
	for _, f := range files {
		if strings.Contains(f.Name, "runtime/fetchx") && f.Location == "lib" {
			sawRuntime = true
		}
		if f.Name == "pubspec" && f.Location == "" {
			sawPubspec = true
		}
		if f.Extension == ".dart" && f.Location != "lib" {
			t.Errorf("expected every .dart file under lib/, got Location=%q for %v", f.Location, f.Name)
		}
		if f.Extension == ".dart" && strings.TrimSpace(f.ActualScript) == "" {
			t.Errorf("generated empty dart file: %v/%v", f.Location, f.Name)
		}
	}
	if !sawRuntime {
		t.Error("expected the embedded runtime (fetchx) to be included under lib/runtime")
	}
	if !sawPubspec {
		t.Error("expected pubspec.yaml at the package root")
	}
}
