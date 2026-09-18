package js

import (
	"slices"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestTsFieldTypeOneAndCollectionUseTargetType covers the fix: a one/
// collection field's exported *Type declaration must reference the target's
// own plain Type (e.g. "PassportEntityType"), never the class
// (TsComputedField's "PassportEntity" - correct for the class body's own
// getter/setter annotations, but wrong here). Before this, tsFieldType fell
// through to TsComputedField's default case for these two field types,
// mentioning a class inside what's supposed to be a plain, class-free type.
func TestTsFieldTypeOneAndCollectionUseTargetType(t *testing.T) {
	cases := []struct {
		name  string
		field *core.EmiField
		want  string
	}{
		{
			name:  "one",
			field: &core.EmiField{Name: "passport", Type: core.FieldTypeOne, Target: "PassportEntity"},
			want:  "PassportEntityType",
		},
		{
			name:  "one?",
			field: &core.EmiField{Name: "passport", Type: core.FieldTypeOneNullable, Target: "PassportEntity"},
			want:  "PassportEntityType",
		},
		{
			name:  "collection",
			field: &core.EmiField{Name: "passports", Type: core.FieldTypeCollection, Target: "PassportEntity"},
			want:  "PassportEntityType[]",
		},
		{
			name:  "collection?",
			field: &core.EmiField{Name: "passports", Type: core.FieldTypeCollectionNullable, Target: "PassportEntity"},
			want:  "PassportEntityType[]",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := tsFieldType(c.field, "UserSessionDto")
			if got != c.want {
				t.Fatalf("got %q, want %q", got, c.want)
			}
		})
	}
}

// TestJsCommonObjectGeneratorOneFieldTypeDeclaration is the end-to-end
// version: generates a real dto class with a `one` field and checks both
// halves of the fix land in the same file - the exported *Type uses
// PassportDtoType, and the import pulls in both the class (still needed by
// the class body itself) and its Type companion.
func TestJsCommonObjectGeneratorOneFieldTypeDeclaration(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "token", Type: core.FieldTypeString},
		{Name: "passport", Type: core.FieldTypeOne, Target: "PassportDto"},
	}

	ctx := core.MicroGenContext{Tags: "typescript"}
	chunk, err := JsCommonObjectGenerator(fields, ctx, JsCommonObjectContext{RootClassName: "UserSession"})
	if err != nil {
		t.Fatalf("JsCommonObjectGenerator: %v", err)
	}
	src := string(chunk.ActualScript)

	if !strings.Contains(src, "PassportDtoType;") {
		t.Fatalf("expected the exported Type to reference PassportDtoType, got:\n%s", src)
	}
	if strings.Contains(src, "passport : PassportDto;") || strings.Contains(src, "passport: PassportDto;") {
		t.Fatalf("did not expect the exported Type to reference the class PassportDto, got:\n%s", src)
	}

	foundClassImport, foundTypeImport := false, false
	for _, dep := range chunk.CodeChunkDependensies {
		if dep.Location != "./PassportDto" {
			continue
		}
		for _, obj := range dep.Objects {
			if obj == "PassportDto" {
				foundClassImport = true
			}
			if obj == "type PassportDtoType" {
				foundTypeImport = true
			}
		}
	}
	if !foundClassImport {
		t.Fatalf("expected PassportDto to still be imported for the class body, deps: %+v", chunk.CodeChunkDependensies)
	}
	if !foundTypeImport {
		t.Fatalf("expected PassportDtoType to also be imported for the plain Type, deps: %+v", chunk.CodeChunkDependensies)
	}
}

// TestTsFieldTypeSelfReferenceDoesNotDoubleType is the regression test for a
// real bug: a self-referencing collection/one field (`target: self@...`, a
// dto pointing at its own shape - e.g. a tree node's "children") rendered as
// "RootTypeType[]" instead of "RootType[]". getSelfReferencingField resolves
// via parentChain's root segment, which on the type-rendering side is
// already Type-suffixed (tsRenderTypes seeds treeLocation as
// RootTypeName+"Type" from its very first call) - tsTargetTypeName used to
// append "Type" again on top of that.
func TestTsFieldTypeSelfReferenceDoesNotDoubleType(t *testing.T) {
	field := &core.EmiField{Name: "children", Type: core.FieldTypeCollection, Target: "self@"}

	// parentChain shaped exactly like tsRenderTypes' real root call
	// (RootTypeName + "Type"), not the bare class name - that's the whole
	// point being tested here.
	got := tsFieldType(field, "CapabilityInfoDtoType")
	want := "CapabilityInfoDtoType[]"
	if got != want {
		t.Fatalf("got %q, want %q (self-reference must not double the \"Type\" suffix)", got, want)
	}
}

// TestTsFieldTypeComplexUsesPlainOf covers deriving a complex field's
// exported *Type from PlainOf<Complex> instead of the bare complex name
// (TsComputedField's own class-side handling, correct there but wrong for a
// plain type declaration that's supposed to be class-free) - see
// tsPlainOfDeclaration's own doc comment for why a generic ReturnType<>-based
// conditional, not a generated "<Complex>Type" name, is the right mechanism
// here: Emi has no way to know a hand-authored complex type's own plain
// shape the way it does for its own dto/entity targets.
func TestTsFieldTypeComplexUsesPlainOf(t *testing.T) {
	field := &core.EmiField{Name: "name", Type: core.FieldTypeComplex, Complex: "TString"}
	got := tsFieldType(field, "CapabilityInfoDtoType")
	if got != "PlainOf<TString>" {
		t.Fatalf("got %q, want %q", got, "PlainOf<TString>")
	}
}

// TestTsFieldTypeComplexStripsAutoInstantiatePrefix covers the "+ClassName"
// auto-instantiate convention (see docs/emi-complex-types) - the "+" must
// not leak into the PlainOf<...> wrapper.
func TestTsFieldTypeComplexStripsAutoInstantiatePrefix(t *testing.T) {
	field := &core.EmiField{Name: "date", Type: core.FieldTypeComplex, Complex: "+Date"}
	got := tsFieldType(field, "SomeDtoType")
	if got != "PlainOf<Date>" {
		t.Fatalf("got %q, want %q", got, "PlainOf<Date>")
	}
}

// TestTsCommonObjectGeneratorEmitsPlainOfOnlyWhenNeeded covers the
// per-file import: PlainOf must be imported from the shared SDK
// (sdk/common/fetchx, alongside PartialDeep - see the fix that moved this
// from an inline per-file `type PlainOf<T> = ...` declaration to a real
// import once the user added PlainOf to fetchx.ts) exactly when any field
// (including nested ones) is complex, and not at all otherwise - it would
// be a dangling, unused import in a file with nothing to apply it to.
func TestTsCommonObjectGeneratorEmitsPlainOfOnlyWhenNeeded(t *testing.T) {
	findPlainOfImport := func(deps []core.CodeChunkDependency) bool {
		for _, dep := range deps {
			if slices.Contains(dep.Objects, "type PlainOf") {
				return true
			}
		}
		return false
	}

	t.Run("with a complex field", func(t *testing.T) {
		fields := []*core.EmiField{
			{Name: "name", Type: core.FieldTypeComplex, Complex: "TString"},
		}
		chunk, err := TsCommonObjectGenerator(fields, core.MicroGenContext{}, TsCommonObjectContext{RootTypeName: "capabilityInfo"})
		if err != nil {
			t.Fatalf("TsCommonObjectGenerator: %v", err)
		}
		src := string(chunk.ActualScript)
		if strings.Contains(src, "type PlainOf<T>") {
			t.Fatalf("did not expect an inline PlainOf<T> declaration, it should be imported instead, got:\n%s", src)
		}
		if !findPlainOfImport(chunk.CodeChunkDependensies) {
			t.Fatalf("expected a \"type PlainOf\" import, deps: %+v", chunk.CodeChunkDependensies)
		}
	})

	t.Run("without any complex field", func(t *testing.T) {
		fields := []*core.EmiField{
			{Name: "id", Type: core.FieldTypeInt64},
		}
		chunk, err := TsCommonObjectGenerator(fields, core.MicroGenContext{}, TsCommonObjectContext{RootTypeName: "simple"})
		if err != nil {
			t.Fatalf("TsCommonObjectGenerator: %v", err)
		}
		src := string(chunk.ActualScript)
		if strings.Contains(src, "PlainOf") {
			t.Fatalf("did not expect PlainOf when nothing needs it, got:\n%s", src)
		}
		if findPlainOfImport(chunk.CodeChunkDependensies) {
			t.Fatalf("did not expect a PlainOf import when nothing needs it, deps: %+v", chunk.CodeChunkDependensies)
		}
	})
}
