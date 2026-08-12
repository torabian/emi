package java

import (
	"os"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// scriptOf concatenates every chunk's script - handy for tests that just
// want to grep the combined output for a substring across the (now
// per-type-file) chunks JavaCommonObjectGenerator returns.
func scriptOf(chunks []*core.CodeChunkCompiled) string {
	var sb strings.Builder
	for _, c := range chunks {
		sb.Write(c.ActualScript)
		sb.WriteString("\n")
	}
	return sb.String()
}

// TestJavaCommonObjectGenerator_SafeDefaults covers the core promise behind
// javaResolveField: every field - no matter its type or nullability -
// always has either a safe initializer or is left at Java's implicit `null`,
// so `new SomeDto()` is always valid.
func TestJavaCommonObjectGenerator_SafeDefaults(t *testing.T) {
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

	chunks, err := JavaCommonObjectGenerator(fields, core.MicroGenContext{}, JavaCommonObjectContext{RootClassName: "Sample"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := scriptOf(chunks)

	cases := []string{
		`public String requiredString = "";`,
		`public String nullableString;`,
		`public Integer requiredInt = 0;`,
		`public Boolean requiredBool = false;`,
		`public List<String> items = new java.util.ArrayList<>();`,
		`public SampleNested nested = new SampleNested();`,
		// a `one` relation is never eagerly default-constructed, even when
		// required - always left at Java's implicit null.
		`public OtherDto relation;`,
		// `any`/`complex` has no sensible non-null default either.
		`public Object raw;`,
	}
	for _, want := range cases {
		if !strings.Contains(script, want) {
			t.Errorf("expected generated script to contain %q, got:\n%v", want, script)
		}
	}

	// every flattened type must be its own file - and each its own public
	// type, named to match, since Java requires the file name to match its
	// public type.
	names := map[string]bool{}
	for _, c := range chunks {
		names[c.SuggestedFileName] = true
		if c.SuggestedExtension != ".java" {
			t.Errorf("expected .java extension, got %q for %v", c.SuggestedExtension, c.SuggestedFileName)
		}
		if !strings.Contains(string(c.ActualScript), "public class "+c.SuggestedFileName) &&
			!strings.Contains(string(c.ActualScript), "public enum "+c.SuggestedFileName) {
			t.Errorf("expected %v.java to declare a public type named %v, got:\n%v", c.SuggestedFileName, c.SuggestedFileName, string(c.ActualScript))
		}
	}
	if !names["Sample"] || !names["SampleNested"] {
		t.Errorf("expected both Sample and SampleNested as separate files, got %v", names)
	}
}

// TestJavaCommonObjectGenerator_EmptyNestedObjectStillGetsAClass covers a
// degenerate but schema-legal case: an `object`/`array` field with no
// `fields:` at all - the type still has to name *something*, and it must
// actually be emitted or the reference fails to compile. It also produces a
// property-less class, which Jackson treats as an error unless
// FAIL_ON_EMPTY_BEANS is disabled (see the runtime's Fetchx.MAPPER).
func TestJavaCommonObjectGenerator_EmptyNestedObjectStillGetsAClass(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "meta", Type: core.FieldTypeObject},
		{Name: "tags", Type: core.FieldTypeArray},
	}

	chunks, err := JavaCommonObjectGenerator(fields, core.MicroGenContext{}, JavaCommonObjectContext{RootClassName: "Widget"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := scriptOf(chunks)

	for _, want := range []string{"public class WidgetMeta", "public class WidgetTags", "public WidgetMeta meta", "public List<WidgetTags> tags"} {
		if !strings.Contains(script, want) {
			t.Errorf("expected %q in generated script:\n%v", want, script)
		}
	}
}

// TestJavaFieldType_MissingTargetFallsBackToObject covers a malformed (but
// not generator-crashing) `one`/`collection` field with no `target:` - it
// must never reference nestedClassName, since object/array are the only
// field types that ever actually get a flattened class generated for them.
func TestJavaFieldType_MissingTargetFallsBackToObject(t *testing.T) {
	one := &core.EmiField{Name: "owner", Type: core.FieldTypeOne}
	if got := javaFieldType(one, "RootOwner", "Root"); got != "Object" {
		t.Errorf("expected targetless `one` to fall back to Object, got %q", got)
	}

	collection := &core.EmiField{Name: "items", Type: core.FieldTypeCollection}
	if got := javaFieldType(collection, "RootItems", "Root"); got != "List<Object>" {
		t.Errorf("expected targetless `collection` to fall back to List<Object>, got %q", got)
	}
}

// TestJavaFieldType_GolangOnlyListAndClass covers `_list`/`class` -
// documented as "golang-only for now" and never special-cased by any other
// sibling generator either, but still part of
// core.GetEmiFieldTypeCatalog().DtoFieldTypes.
func TestJavaFieldType_GolangOnlyListAndClass(t *testing.T) {
	list := &core.EmiField{Name: "items", Type: core.FieldTypeList}
	if got := javaFieldType(list, "RootItems", "Root"); got != "List<RootItems>" {
		t.Errorf("expected `_list` to render like `array`, got %q", got)
	}

	class := &core.EmiField{Name: "owner", Type: core.FieldTypeClass, Target: "UserDto"}
	if got := javaFieldType(class, "RootOwner", "Root"); got != "UserDto" {
		t.Errorf("expected `class` to render like `one`, got %q", got)
	}
}

// TestJavaEnumField covers every shape an `enum`/`enum?` field can take: a
// `target:` reference to a module-level enum, inline `of:` values (a
// generated companion enum in its own public file, with its own
// @JsonValue/@JsonCreator wire-value mapping), and neither (a plain string).
func TestJavaEnumField(t *testing.T) {
	t.Run("target reference is a plain enum property defaulting to the first member", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "status", Type: core.FieldTypeEnum, Target: "StatusEnum"}}
		chunks, err := JavaCommonObjectGenerator(fields, core.MicroGenContext{}, JavaCommonObjectContext{RootClassName: "Task"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := scriptOf(chunks)
		if !strings.Contains(script, "public StatusEnum status = StatusEnum.values()[0];") {
			t.Errorf("expected status defaulting to the first StatusEnum member, got:\n%v", script)
		}
	})

	t.Run("inline of: values become their own public companion enum file", func(t *testing.T) {
		fields := []*core.EmiField{{
			Name: "kind",
			Type: core.FieldTypeEnum,
			OfType: []*core.EmiEnumInline{
				{Key: "cat"},
				{Key: "dog"},
			},
		}}
		chunks, err := JavaCommonObjectGenerator(fields, core.MicroGenContext{}, JavaCommonObjectContext{RootClassName: "Pet"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var enumChunk *core.CodeChunkCompiled
		for _, c := range chunks {
			if c.SuggestedFileName == "PetKind" {
				enumChunk = c
			}
		}
		if enumChunk == nil {
			t.Fatalf("expected a separate PetKind.java file, got %+v", chunks)
		}
		enumScript := string(enumChunk.ActualScript)
		if !strings.Contains(enumScript, "public enum PetKind") {
			t.Errorf("expected PetKind to be public, got:\n%v", enumScript)
		}
		if !strings.Contains(enumScript, `CAT("cat")`) {
			t.Errorf(`expected a CAT("cat") case, got:`+"\n%v", enumScript)
		}
		if !strings.Contains(enumScript, "@JsonValue") || !strings.Contains(enumScript, "@JsonCreator") {
			t.Errorf("expected @JsonValue/@JsonCreator wire-value mapping, got:\n%v", enumScript)
		}
	})

	t.Run("neither target nor of: falls back to plain String", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "freeform", Type: core.FieldTypeEnum}}
		chunks, err := JavaCommonObjectGenerator(fields, core.MicroGenContext{}, JavaCommonObjectContext{RootClassName: "Loose"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := scriptOf(chunks)
		if !strings.Contains(script, `public String freeform = "";`) {
			t.Errorf("expected freeform as a plain defaulted String, got:\n%v", script)
		}
	})
}

// TestJavaModuleFullVirtualFiles_Fixture is a broad smoke test: running the
// full module compiler over a real, feature-rich fixture (nested objects,
// slices, one/collection relations, complex fields, path/query params,
// reactive actions, envelopes...) must never error, and must produce
// non-empty Java for every dto/action/enum plus the embedded runtime and
// pom.xml, laid out as a real Maven-shaped tree
// (src/main/java/emisdk/*.java).
func TestJavaModuleFullVirtualFiles_Fixture(t *testing.T) {
	content, err := os.ReadFile("../../examples/fullstack/definitions.emi.yml")
	if err != nil {
		t.Fatal(err)
	}

	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}

	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}}

	files, err := JavaModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatalf("unexpected error compiling the fixture module: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected at least one generated file")
	}

	sawRuntime, sawPom := false, false
	for _, f := range files {
		if strings.Contains(f.Name, "Fetchx") {
			sawRuntime = true
		}
		if f.Extension == ".xml" && f.Name == "pom" {
			sawPom = true
		}
		if f.Extension == ".java" {
			if f.Location != "src/main/java/emisdk" {
				t.Errorf("expected every .java file under src/main/java/emisdk, got Location=%q for %v", f.Location, f.Name)
			}
			if strings.TrimSpace(f.ActualScript) == "" {
				t.Errorf("generated empty java file: %v/%v", f.Location, f.Name)
			}
			if !strings.Contains(f.ActualScript, "public class "+f.Name) && !strings.Contains(f.ActualScript, "public enum "+f.Name) {
				t.Errorf("expected %v.java to declare a public type matching its filename, got:\n%v", f.Name, f.ActualScript)
			}
		}
	}
	if !sawRuntime {
		t.Error("expected the embedded runtime (Fetchx) to be included in the output")
	}
	if !sawPom {
		t.Error("expected pom.xml at the package root")
	}
}
