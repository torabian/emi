// Unreal-dialect field resolution. Unlike the generic dialect
// (cpp-field-plan-generic.go), no per-field (de)serialization code is generated
// here at all - once every field is a real UPROPERTY, Unreal's own reflection
// system ((de)serializes any USTRUCT via FJsonObjectConverter::UStructToJsonObjectString
// / FJsonObjectConverter::JsonObjectStringToUStruct<T>()) does that job, the same
// reason lib/csharp never hand-generates ToJson/FromJson either (System.Text.Json
// does it via reflection there). So this file only ever computes *declarations*.
package cpp

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type cppUnrealFieldPlan struct {
	Decl       []string
	InlineEnum *cppRenderedEnum
}

// ueStructName spells the USTRUCT name Emi generates for a dto/nested-object/
// target name, following Unreal's `F` value-type prefix convention. Left alone
// if already prefixed (a `target:`/complex symbol the developer already spelled
// in Unreal style, e.g. `target: FInventoryItem`).
func ueStructName(name string) string {
	if name == "" || strings.HasPrefix(name, "F") {
		return name
	}
	return "F" + name
}

// ueEnumName mirrors ueStructName for Unreal's `E` enum prefix convention.
func ueEnumName(name string) string {
	if name == "" || strings.HasPrefix(name, "E") {
		return name
	}
	return "E" + name
}

// cppUnrealResolveField computes one field's UPROPERTY + declaration line(s).
func cppUnrealResolveField(field *core.EmiField, prefixName string, selfTypeName string, complexes []RecognizedComplex) cppUnrealFieldPlan {
	name := core.ToUpper(field.Name)
	nestedType := ueStructName(cppFieldNestedTypeName(field, prefixName))

	if field.Complex != "" {
		symbol, resolved := resolveComplexSymbol(field, complexes)
		if resolved {
			return ueValueField(symbol, name)
		}
		return ueAnyField(name)
	}

	switch {
	case field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable:
		base := ueEnumName(cppEnumBaseType(field, prefixName))
		plan := ueFieldMaybeNullable(base, name, core.IsNullable(string(field.Type)))
		if field.Target == "" {
			plan.InlineEnum = cppRenderEnumFromInline(base, field.OfType)
		}
		return plan

	case field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass ||
		field.Type == core.FieldTypeOneNullable || field.Type == core.FieldTypeClassNullable:
		if field.Target == "" {
			return ueAnyField(name)
		}
		target := ueStructName(field.Target)
		if target == selfTypeName {
			return ueSelfReferencingField(target, name)
		}
		return ueValueField(target, name)

	case field.Type == core.FieldTypeCollection || field.Type == core.FieldTypeCollectionNullable:
		if field.Target == "" {
			return ueAnyField(name)
		}
		return ueArrayField(ueStructName(field.Target), name)

	case field.Type == core.FieldTypeArray || field.Type == core.FieldTypeArrayNullable ||
		field.Type == core.FieldTypeList || field.Type == core.FieldTypeListNullable:
		return ueArrayField(nestedType, name)

	case field.Type == core.FieldTypeSlice || field.Type == core.FieldTypeSliceNullable:
		primitive := extractPrimitiveUnreal(field.Primitive)
		if primitive == "" {
			return ueAnyField(name)
		}
		return ueArrayField(primitive, name)

	case field.Type == core.FieldTypeObject || field.Type == core.FieldTypeObjectNullable:
		selfReferencing := nestedType == selfTypeName
		nullable := core.IsNullable(string(field.Type))
		if selfReferencing {
			return ueSelfReferencingField(nestedType, name)
		}
		if nullable {
			return ueFieldMaybeNullable(nestedType, name, true)
		}
		return ueValueField(nestedType, name)

	case field.Type == core.FieldTypeMap || field.Type == core.FieldTypeMapNullable:
		return ueMapField(field, name)

	default:
		primitive := extractPrimitiveUnreal(string(field.Type))
		if primitive == "" {
			return ueAnyField(name)
		}
		return ueFieldMaybeNullable(primitive, name, core.IsNullable(string(field.Type)))
	}
}

func ueScalarField(ueType, name string) cppUnrealFieldPlan {
	return cppUnrealFieldPlan{
		Decl: []string{
			`UPROPERTY(BlueprintReadWrite, EditAnywhere, Category = "Emi")`,
			fmt.Sprintf("%v %v;", ueType, name),
		},
	}
}

// ueFieldMaybeNullable renders a scalar/string/enum/nested-object field,
// nullable or not. A nullable field gets a real value UPROPERTY plus a paired
// `b<Name>IsSet` bool UPROPERTY, rather than `UPROPERTY() TOptional<T>` - both
// are real, Blueprint-visible, editor-visible UPROPERTYs, reflectable (and thus
// round-trippable through FJsonObjectConverter) on every Unreal engine version
// since UE4. UPROPERTY-reflecting a bare TOptional<T> is a comparatively recent,
// version-gated Unreal Header Tool feature (not supported on UE4, and
// inconsistently supported across UE5 point releases before it), so this
// compiler deliberately never emits it as a UPROPERTY at all - this two-field
// pattern is also the idiomatic way Unreal codebases represented "optional"
// long before reflected TOptional existed.
func ueFieldMaybeNullable(ueType, name string, nullable bool) cppUnrealFieldPlan {
	if !nullable {
		return ueScalarField(ueType, name)
	}
	return cppUnrealFieldPlan{
		Decl: []string{
			`UPROPERTY(BlueprintReadWrite, EditAnywhere, Category = "Emi")`,
			fmt.Sprintf("bool b%vIsSet = false;", name),
			``,
			`UPROPERTY(BlueprintReadWrite, EditAnywhere, Category = "Emi")`,
			fmt.Sprintf("%v %v;", ueType, name),
		},
	}
}

func ueValueField(ueType, name string) cppUnrealFieldPlan {
	return cppUnrealFieldPlan{
		Decl: []string{
			`UPROPERTY(BlueprintReadWrite, EditAnywhere, Category = "Emi")`,
			fmt.Sprintf("%v %v;", ueType, name),
		},
	}
}

func ueArrayField(elemType, name string) cppUnrealFieldPlan {
	return cppUnrealFieldPlan{
		Decl: []string{
			`UPROPERTY(BlueprintReadWrite, EditAnywhere, Category = "Emi")`,
			fmt.Sprintf("TArray<%v> %v;", elemType, name),
		},
	}
}

func ueMapField(field *core.EmiField, name string) cppUnrealFieldPlan {
	keyType := extractPrimitiveUnreal(field.MapKeyOf)
	if keyType == "" {
		keyType = "FString"
	}
	valueType := extractPrimitiveUnreal(field.MapPairOf)
	if valueType == "" {
		return ueAnyField(name)
	}
	return cppUnrealFieldPlan{
		Decl: []string{
			`UPROPERTY(BlueprintReadWrite, EditAnywhere, Category = "Emi")`,
			fmt.Sprintf("TMap<%v, %v> %v;", keyType, valueType, name),
		},
	}
}

// ueSelfReferencingField covers a `one`/`class`/`object` field whose target is
// the very USTRUCT being declared - an incomplete type can never be a by-value
// UPROPERTY member (nor is it UPROPERTY-reflectable through a raw/shared
// pointer), so this field is intentionally plain C++ (still fully usable, just
// outside Blueprint/JSON reflection) - the same documented boundary the generic
// dialect resolves with std::unique_ptr<T> (cpp-field-plan-generic.go).
func ueSelfReferencingField(ueType, name string) cppUnrealFieldPlan {
	return cppUnrealFieldPlan{
		Decl: []string{
			"// Self-referencing - not reflected (UPROPERTY can't take an incomplete type; see the",
			"// doc comment on ueSelfReferencingField in the Emi compiler for why).",
			fmt.Sprintf("TSharedPtr<%v> %v;", ueType, name),
		},
	}
}

// ueAnyField covers `any`/an unresolved `complex`/unresolved `one`/`collection`/
// `map` - Unreal's own dynamic JSON value type (from the Json module). Also not
// UPROPERTY-reflected, for the same reason as ueSelfReferencingField.
func ueAnyField(name string) cppUnrealFieldPlan {
	return cppUnrealFieldPlan{
		Decl: []string{
			"// Dynamic/unresolved shape - not reflected (TSharedPtr<FJsonValue> isn't a",
			"// UPROPERTY-compatible type).",
			fmt.Sprintf("TSharedPtr<FJsonValue> %v;", name),
		},
	}
}
