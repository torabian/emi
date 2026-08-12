package php

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// doubledNullablePrefix matches a doubled `?`-type-prefix bug (`??string`),
// without false-positiving on the legitimate `??` null-coalescing operator
// (always followed by whitespace in generated code, e.g. `$ctx ?? \Foo`) or
// the nullsafe `?->` operator.
var doubledNullablePrefix = regexp.MustCompile(`\?\?[A-Za-z\\]`)

// TestPhpCommonObjectGenerator_SafeDefaults covers the core promise behind
// phpResolveField: every field - no matter its type or nullability - always
// has either a constant default or a constructor-assigned one, so
// `new SomeDto()` is always valid (PHP typed properties throw if accessed
// while uninitialized, so this isn't just a nicety here).
func TestPhpCommonObjectGenerator_SafeDefaults(t *testing.T) {
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

	chunk, err := PhpCommonObjectGenerator(fields, core.MicroGenContext{}, PhpCommonObjectContext{RootClassName: "Sample"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	cases := []string{
		`public string $requiredString = '';`,
		`public ?string $nullableString = null;`,
		`public int $requiredInt = 0;`,
		`public bool $requiredBool = false;`,
		`public array $items = [];`,
		// `new SampleNested()` isn't a constant expression - constructor-assigned.
		`public SampleNested $nested;`,
		`$this->nested = new SampleNested();`,
		// a `one` relation is never eagerly default-constructed, even when
		// required - always a plain nullable property.
		`public ?OtherDto $relation = null;`,
		// `any`/`complex` has no sensible non-null default either.
		`public mixed $raw = null;`,
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

	if chunk.SuggestedFileName != "Sample" || chunk.SuggestedExtension != ".php" {
		t.Errorf("unexpected file naming: %v%v", chunk.SuggestedFileName, chunk.SuggestedExtension)
	}
}

// TestPhpCommonObjectGenerator_NoDoubleNullablePrefix guards against a real
// bug this generator once had: phpFieldType already returns a `?`-prefixed
// type for nullable primitives, so the field-plan must never prepend a
// second `?` (PHP rejects `??string`, and separately rejects `?mixed`
// outright since `mixed` already includes null).
func TestPhpCommonObjectGenerator_NoDoubleNullablePrefix(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "nickname", Type: core.FieldTypeStringNullable},
	}
	chunk, err := PhpCommonObjectGenerator(fields, core.MicroGenContext{}, PhpCommonObjectContext{RootClassName: "Widget"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)
	if doubledNullablePrefix.MatchString(script) {
		t.Errorf("expected no double-`?` nullable prefix, got:\n%v", script)
	}
	if !strings.Contains(script, "public ?string $nickname = null;") {
		t.Errorf("expected a single `?` prefix, got:\n%v", script)
	}
}

// TestPhpCommonObjectGenerator_EmptyNestedObjectStillGetsAClass covers a
// degenerate but schema-legal case: an `object`/`array` field with no
// `fields:` at all - the type still has to name *something*, and it must
// actually be emitted or the reference fails to compile.
func TestPhpCommonObjectGenerator_EmptyNestedObjectStillGetsAClass(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "meta", Type: core.FieldTypeObject},
		{Name: "tags", Type: core.FieldTypeArray},
	}

	chunk, err := PhpCommonObjectGenerator(fields, core.MicroGenContext{}, PhpCommonObjectContext{RootClassName: "Widget"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	for _, want := range []string{"class WidgetMeta", "class WidgetTags", "public WidgetMeta $meta;", "public array $tags = [];"} {
		if !strings.Contains(script, want) {
			t.Errorf("expected %q in generated script:\n%v", want, script)
		}
	}
}

// TestPhpFieldType_MissingTargetFallsBackToMixed covers a malformed (but
// not generator-crashing) `one`/`collection` field with no `target:` - it
// must never reference nestedClassName, since object/array are the only
// field types that ever actually get a flattened class generated for them.
func TestPhpFieldType_MissingTargetFallsBackToMixed(t *testing.T) {
	one := &core.EmiField{Name: "owner", Type: core.FieldTypeOne}
	if got := phpFieldType(one, "RootOwner", "Root"); got != "mixed" {
		t.Errorf("expected targetless `one` to fall back to mixed, got %q", got)
	}

	collection := &core.EmiField{Name: "items", Type: core.FieldTypeCollection}
	if got := phpFieldType(collection, "RootItems", "Root"); got != "array" {
		t.Errorf("expected targetless `collection` to fall back to array, got %q", got)
	}
}

// TestPhpFieldType_GolangOnlyListAndClass covers `_list`/`class` -
// documented as "golang-only for now" and never special-cased by any other
// sibling generator either, but still part of
// core.GetEmiFieldTypeCatalog().DtoFieldTypes.
func TestPhpFieldType_GolangOnlyListAndClass(t *testing.T) {
	list := &core.EmiField{Name: "items", Type: core.FieldTypeList}
	if got := phpFieldType(list, "RootItems", "Root"); got != "array" {
		t.Errorf("expected `_list` to render like `array`, got %q", got)
	}

	class := &core.EmiField{Name: "owner", Type: core.FieldTypeClass, Target: "UserDto"}
	if got := phpFieldType(class, "RootOwner", "Root"); got != "?UserDto" {
		t.Errorf("expected `class` to render like `one`, got %q", got)
	}
}

// TestPhpEnumField covers every shape an `enum`/`enum?` field can take: a
// `target:` reference to a module-level backed enum (defaulting eagerly to
// its first case, since a backed enum - unlike an arbitrary object - is
// always safely enumerable), inline `of:` values (a generated companion
// backed enum), and neither (a plain string).
func TestPhpEnumField(t *testing.T) {
	t.Run("target reference defaults eagerly to the first case", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "status", Type: core.FieldTypeEnum, Target: "StatusEnum"}}
		chunk, err := PhpCommonObjectGenerator(fields, core.MicroGenContext{}, PhpCommonObjectContext{RootClassName: "Task"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "public StatusEnum $status;") {
			t.Errorf("expected status: StatusEnum $status;, got:\n%v", script)
		}
		if !strings.Contains(script, "$this->status = StatusEnum::cases()[0];") {
			t.Errorf("expected a first-case default in the constructor, got:\n%v", script)
		}
	})

	t.Run("inline of: values become a generated companion backed enum", func(t *testing.T) {
		fields := []*core.EmiField{{
			Name: "kind",
			Type: core.FieldTypeEnum,
			OfType: []*core.EmiEnumInline{
				{Key: "cat"},
				{Key: "dog"},
			},
		}}
		chunk, err := PhpCommonObjectGenerator(fields, core.MicroGenContext{}, PhpCommonObjectContext{RootClassName: "Pet"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "enum PetKind: string") {
			t.Errorf("expected a companion PetKind backed enum, got:\n%v", script)
		}
		if !strings.Contains(script, `case Cat = "cat";`) {
			t.Errorf(`expected a case Cat = "cat"; case, got:`+"\n%v", script)
		}
	})

	t.Run("neither target nor of: falls back to plain string", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "freeform", Type: core.FieldTypeEnum}}
		chunk, err := PhpCommonObjectGenerator(fields, core.MicroGenContext{}, PhpCommonObjectContext{RootClassName: "Loose"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, `public string $freeform = '';`) {
			t.Errorf("expected freeform as a plain defaulted string, got:\n%v", script)
		}
	})
}

// TestPhpModuleFullVirtualFiles_Fixture is a broad smoke test: running the
// full module compiler over a real, feature-rich fixture (nested objects,
// slices, one/collection relations, complex fields, path/query params,
// reactive actions, envelopes...) must never error, and must produce
// non-empty PHP for every dto/action plus the embedded runtime and
// composer.json, laid out as a real PSR-4 package (src/*.php,
// src/Runtime/*.php).
func TestPhpModuleFullVirtualFiles_Fixture(t *testing.T) {
	content, err := os.ReadFile("../../examples/fullstack/definitions.emi.yml")
	if err != nil {
		t.Fatal(err)
	}

	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}

	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}}

	files, err := PhpModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatalf("unexpected error compiling the fixture module: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected at least one generated file")
	}

	sawRuntime, sawComposer := false, false
	for _, f := range files {
		if strings.Contains(f.Name, "Fetchx") {
			sawRuntime = true
		}
		if f.Name == "composer" && f.Extension == ".json" {
			sawComposer = true
		}
		if f.Extension == ".php" {
			if f.Location != "src" {
				t.Errorf("expected every .php file under src/, got Location=%q for %v", f.Location, f.Name)
			}
			if strings.TrimSpace(f.ActualScript) == "" {
				t.Errorf("generated empty php file: %v/%v", f.Location, f.Name)
			}
			if !strings.HasPrefix(f.ActualScript, "<?php") {
				t.Errorf("expected %v.php to start with <?php, got:\n%v", f.Name, f.ActualScript)
			}
			if doubledNullablePrefix.MatchString(f.ActualScript) {
				t.Errorf("expected no double-`?` nullable prefix anywhere in %v.php, got:\n%v", f.Name, f.ActualScript)
			}
			if strings.Contains(f.ActualScript, "?mixed") {
				t.Errorf("expected no `?mixed` (PHP forbids it) anywhere in %v.php, got:\n%v", f.Name, f.ActualScript)
			}
		}
	}
	if !sawRuntime {
		t.Error("expected the embedded runtime (Fetchx) to be included under src/Runtime")
	}
	if !sawComposer {
		t.Error("expected composer.json at the package root")
	}
}
