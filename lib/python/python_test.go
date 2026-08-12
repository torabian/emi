package python

import (
	"os"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestPythonCommonObjectGenerator_SafeDefaults covers the core promise behind
// PythonSafeDefaultValue/pyRenderField: every dataclass field - no matter its
// type or nullability - always gets *some* default, so `SomeDto()` is always
// constructible and field declaration order never has to be worried about
// (python dataclasses reject a required field after a defaulted one).
func TestPythonCommonObjectGenerator_SafeDefaults(t *testing.T) {
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

	// Real callers always pass an already-PascalCase class name (e.g.
	// EmiDto.GetClassName() applies core.ToUpper itself) - the generator
	// doesn't re-capitalize its root, only the nested/flattened children.
	chunk, err := PythonCommonObjectGenerator(fields, core.MicroGenContext{}, PyCommonObjectContext{RootClassName: "Sample"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	script := string(chunk.ActualScript)

	cases := []string{
		`requiredString: str = ""`,
		`nullableString: Optional[str] = None`,
		`requiredInt: int = 0`,
		`requiredBool: bool = False`,
		`items: List[str] = field(default_factory=list)`,
		`nested: SampleNested = field(default_factory=lambda: SampleNested())`,
		// a `one` relation is never eagerly default-constructed, even when
		// required, and its type-hint must say Optional to match.
		`relation: Optional[OtherDto] = None`,
	}
	for _, want := range cases {
		if !strings.Contains(script, want) {
			t.Errorf("expected generated script to contain %q, got:\n%v", want, script)
		}
	}

	// the nested class must be emitted before the class that references it,
	// so a plain top-to-bottom read never hits a forward reference.
	nestedIdx := strings.Index(script, "class SampleNested:")
	rootIdx := strings.Index(script, "class Sample:")
	if nestedIdx == -1 || rootIdx == -1 || nestedIdx > rootIdx {
		t.Errorf("expected SampleNested to be declared before Sample, got:\n%v", script)
	}

	if chunk.SuggestedFileName != "sample" || chunk.SuggestedExtension != ".py" {
		t.Errorf("unexpected file naming: %v%v", chunk.SuggestedFileName, chunk.SuggestedExtension)
	}
}

// TestPythonCommonObjectGenerator_UnresolvedComplexFallsBackToAny guards
// against a real crash: `typing.get_type_hints()` (called by the generated
// runtime's `from_dict`) blows up with a NameError on any type-hint
// referencing a name that isn't actually imported - which is exactly what a
// `complex:` field with no registered RecognizedComplex would produce if
// used verbatim as the type-hint.
func TestPythonCommonObjectGenerator_UnresolvedComplexFallsBackToAny(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "amount", Type: core.FieldTypeObject, Complex: "Money"},
		{Name: "goValue", Type: core.FieldTypeObject, Complex: "big.Int"},
	}

	chunk, err := PythonCommonObjectGenerator(fields, core.MicroGenContext{}, PyCommonObjectContext{RootClassName: "priced"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	script := string(chunk.ActualScript)
	if strings.Contains(script, "amount: Money") || strings.Contains(script, ": big.Int") {
		t.Errorf("unresolved complex types must never be used as a bare type-hint (would NameError under typing.get_type_hints), got:\n%v", script)
	}
	if !strings.Contains(script, "amount: Optional[Any] = None") {
		t.Errorf("expected the unresolved complex field to fall back to Optional[Any], got:\n%v", script)
	}
}

// TestPythonCommonObjectGenerator_RecognizedComplexImportsAndTypes covers the
// happy path of the above: a `+`-prefixed complex that *is* registered in
// RecognizedComplexes gets both a real import and its own type-hint.
func TestPythonCommonObjectGenerator_RecognizedComplexImportsAndTypes(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "position", Type: core.FieldTypeObject, Complex: "+Vector3"},
	}

	chunk, err := PythonCommonObjectGenerator(fields, core.MicroGenContext{}, PyCommonObjectContext{
		RootClassName: "entity",
		RecognizedComplexes: []RecognizedComplex{
			{Symbol: "Vector3", ImportLocation: "mymath.vector3"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found := false
	for _, dep := range chunk.CodeChunkDependensies {
		if dep.Location == "mymath.vector3" && len(dep.Objects) == 1 && dep.Objects[0] == "Vector3" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a dependency importing Vector3 from mymath.vector3, got %+v", chunk.CodeChunkDependensies)
	}

	if !strings.Contains(string(chunk.ActualScript), "position: Optional[Vector3] = None") {
		t.Errorf("expected position field typed as Optional[Vector3], got:\n%v", chunk.ActualScript)
	}
}

// TestPythonCommonObjectGenerator_EmptyNestedObjectStillGetsAClass covers a
// degenerate but schema-legal case: an `object`/`array` field with no `fields:`
// at all. The type-hint still has to name *something*, and that something
// must actually be emitted (even as an empty `pass`-bodied dataclass) or the
// reference dangles and blows up `typing.get_type_hints()`.
func TestPythonCommonObjectGenerator_EmptyNestedObjectStillGetsAClass(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "meta", Type: core.FieldTypeObject},
		{Name: "tags", Type: core.FieldTypeArray},
	}

	chunk, err := PythonCommonObjectGenerator(fields, core.MicroGenContext{}, PyCommonObjectContext{RootClassName: "Widget"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	for _, want := range []string{"class WidgetMeta:", "class WidgetTags:", "meta: WidgetMeta", "tags: List[WidgetTags]"} {
		if !strings.Contains(script, want) {
			t.Errorf("expected %q in generated script:\n%v", want, script)
		}
	}
}

// TestPythonDataStructureType_MissingTargetFallsBackToAny covers a malformed
// (but not generator-crashing) `one`/`collection` field with no `target:` -
// it must never reference nestedClassName, since object/array are the only
// field types that ever actually get a flattened subclass generated for them.
func TestPythonDataStructureType_MissingTargetFallsBackToAny(t *testing.T) {
	// `one` is a single relation, which is always Optional regardless of
	// nullability (see pythonFieldTypeOnNestedClasses) - on top of that, a
	// missing target falls back to Any, hence Optional[Any].
	one := &core.EmiField{Name: "owner", Type: core.FieldTypeOne}
	if got := pythonFieldTypeOnNestedClasses(one, "RootOwner"); got != "Optional[Any]" {
		t.Errorf("expected targetless `one` to fall back to Optional[Any], got %q", got)
	}

	collection := &core.EmiField{Name: "items", Type: core.FieldTypeCollection}
	if got := pythonFieldTypeOnNestedClasses(collection, "RootItems"); got != "List[Any]" {
		t.Errorf("expected targetless `collection` to fall back to List[Any], got %q", got)
	}
}

// TestPythonFieldTypeOnNestedClasses_GolangOnlyListAndClass covers `_list`/
// `class` - documented as "golang-only for now" (EmiFieldType.go) and never
// special-cased by any other sibling generator either, but still part of
// core.GetEmiFieldTypeCatalog().DtoFieldTypes, so a stray `type: _list` or
// `type: class` in a dto must render sensibly rather than fall through to a
// silent `Any`/blank type.
func TestPythonFieldTypeOnNestedClasses_GolangOnlyListAndClass(t *testing.T) {
	list := &core.EmiField{Name: "items", Type: core.FieldTypeList}
	if got := pythonFieldTypeOnNestedClasses(list, "RootItems"); got != "List[RootItems]" {
		t.Errorf("expected `_list` to render like `array`, got %q", got)
	}

	class := &core.EmiField{Name: "owner", Type: core.FieldTypeClass, Target: "UserDto"}
	if got := pythonFieldTypeOnNestedClasses(class, "RootOwner"); got != "Optional[UserDto]" {
		t.Errorf("expected `class` to render like `one`, got %q", got)
	}
}

// TestPythonEnumField covers every shape an `enum`/`enum?` field can take:
// a `target:` reference to a module-level enum, inline `of:` values (a
// Literal[...]), and neither (a plain str) - mirroring
// lib/js/js-data-types.go#TsComputedField's "enum" case.
func TestPythonEnumField(t *testing.T) {
	t.Run("target reference is always optional and imports the enum class", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "status", Type: core.FieldTypeEnum, Target: "StatusEnum"}}
		chunk, err := PythonCommonObjectGenerator(fields, core.MicroGenContext{}, PyCommonObjectContext{RootClassName: "Task"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, "status: Optional[StatusEnum] = None") {
			t.Errorf("expected status: Optional[StatusEnum] = None, got:\n%v", script)
		}
		found := false
		for _, dep := range chunk.CodeChunkDependensies {
			if dep.Location == ".status_enum" && len(dep.Objects) == 1 && dep.Objects[0] == "StatusEnum" {
				found = true
			}
		}
		if !found {
			t.Errorf("expected an import for StatusEnum from .status_enum, got %+v", chunk.CodeChunkDependensies)
		}
	})

	t.Run("inline of: values become a Literal with a valid default", func(t *testing.T) {
		fields := []*core.EmiField{{
			Name: "kind",
			Type: core.FieldTypeEnum,
			OfType: []*core.EmiEnumInline{
				{Key: "cat"},
				{Key: "dog"},
			},
		}}
		chunk, err := PythonCommonObjectGenerator(fields, core.MicroGenContext{}, PyCommonObjectContext{RootClassName: "Pet"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, `kind: Literal["cat", "dog"] = "cat"`) {
			t.Errorf(`expected kind: Literal["cat", "dog"] = "cat", got:`+"\n%v", script)
		}
		literalImported := false
		for _, dep := range chunk.CodeChunkDependensies {
			if dep.Location == "typing" {
				for _, obj := range dep.Objects {
					if obj == "Literal" {
						literalImported = true
					}
				}
			}
		}
		if !literalImported {
			t.Errorf("expected typing.Literal to be imported, got %+v", chunk.CodeChunkDependensies)
		}
	})

	t.Run("neither target nor of: falls back to plain str", func(t *testing.T) {
		fields := []*core.EmiField{{Name: "freeform", Type: core.FieldTypeEnum}}
		chunk, err := PythonCommonObjectGenerator(fields, core.MicroGenContext{}, PyCommonObjectContext{RootClassName: "Loose"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, `freeform: str = ""`) {
			t.Errorf(`expected freeform: str = "", got:`+"\n%v", script)
		}
	})

	t.Run("nullable inline enum defaults to None, not a Literal member", func(t *testing.T) {
		fields := []*core.EmiField{{
			Name: "kind",
			Type: core.FieldTypeEnumNullable,
			OfType: []*core.EmiEnumInline{
				{Key: "cat"},
			},
		}}
		chunk, err := PythonCommonObjectGenerator(fields, core.MicroGenContext{}, PyCommonObjectContext{RootClassName: "Pet"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if !strings.Contains(script, `kind: Optional[Literal["cat"]] = None`) {
			t.Errorf(`expected kind: Optional[Literal["cat"]] = None, got:`+"\n%v", script)
		}
	})
}

// TestCombinePythonImports covers the relative-vs-absolute import rendering:
// a Location starting with "." is a package-relative generated module, and
// anything else (stdlib/third-party) renders as a plain `import x`/`from x
// import y` statement, both deduplicated and sorted.
func TestCombinePythonImports(t *testing.T) {
	chunk := core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "typing", Objects: []string{"Optional"}},
			{Location: "typing", Objects: []string{"Any", "Optional"}},
			{Location: ".user_dto", Objects: []string{"UserDto"}},
			{Location: "httpx"},
		},
	}

	got := CombinePythonImports(chunk)
	want := "from .user_dto import UserDto\nfrom typing import Any, Optional\nimport httpx"
	if got != want {
		t.Errorf("unexpected import block:\ngot:  %q\nwant: %q", got, want)
	}
}

// TestPythonModuleFullVirtualFiles_Fixture is a broad smoke test: running the
// full module compiler over a real, feature-rich fixture (nested objects,
// slices, one/collection relations, complex fields, path/query params,
// reactive actions, envelopes...) must never error, and must produce valid,
// non-empty python for every dto/action plus the embedded runtime.
func TestPythonModuleFullVirtualFiles_Fixture(t *testing.T) {
	content, err := os.ReadFile("../../examples/fullstack/definitions.emi.yml")
	if err != nil {
		t.Fatal(err)
	}

	module, err := core.StringToEmi(string(content))
	if err != nil {
		t.Fatal(err)
	}

	ctx := core.MicroGenContext{Content: string(content), Flags: map[string]string{}}

	files, err := PythonModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatalf("unexpected error compiling the fixture module: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected at least one generated file")
	}

	sawRuntime := false
	for _, f := range files {
		if strings.Contains(f.Name, "runtime/fetchx") {
			sawRuntime = true
		}
		if f.Extension == ".py" && strings.TrimSpace(f.ActualScript) == "" {
			t.Errorf("generated empty python file: %v/%v", f.Location, f.Name)
		}
	}
	if !sawRuntime {
		t.Error("expected the embedded runtime (fetchx/serialization) to be included in the output")
	}
}
