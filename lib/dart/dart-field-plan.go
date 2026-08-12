// Dart has no runtime reflection in production (Flutter/AOT) builds, so
// unlike lib/python's generic to_dict/from_dict, every dart class needs its
// own hand-generated fromJson/toJson. This file computes, per field, every
// piece those three places (field declaration, constructor parameter,
// fromJson body, toJson body) need - once, in one place - so the class
// generator just assembles them.
//
// One Dart-specific wrinkle drives most of the shape below: constructor
// *default parameter values* must be compile-time constants. `= const []`
// would satisfy that but hands back an unmodifiable list the first time a
// caller tries `.add()` on it - a real footgun. So any field whose "safe
// default" isn't a literal (list/map/nested-object/enum) is declared as a
// plain nullable parameter with NO default in the parameter list, and
// resolved instead through a constructor initializer-list entry
// (`: field = field ?? <expr>`), which runs at call time and always
// produces a fresh, mutable value.
package dart

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type dartFieldPlan struct {
	Name     string
	TypeHint string

	FieldDecl string // e.g. `String title;`
	CtorParam string // e.g. `this.title = '',` or `List<int>? scores,`

	// Only set when CtorParam does NOT use `this.` - the field is instead
	// bound through a constructor initializer-list entry:
	// `field = <InitializerExpr>`.
	InitializerExpr string

	FromJson   string // RHS expression, may reference `json['name']`
	ToJson     string // RHS expression, may reference `name`
	InlineEnum *dartRenderedEnum
}

func (p dartFieldPlan) NeedsInitializer() bool {
	return p.InitializerExpr != ""
}

func jsonKeyExpr(field *core.EmiField) string {
	b, _ := json.Marshal(field.Name)
	return fmt.Sprintf("json[%v]", string(b))
}

// dartPrimitiveCast casts a single decoded JSON value (already extracted,
// bound to `expr`) into the given dart primitive type.
func dartPrimitiveCast(dartType string, expr string) string {
	switch dartType {
	case "int":
		return fmt.Sprintf("(%v as num).toInt()", expr)
	case "double":
		return fmt.Sprintf("(%v as num).toDouble()", expr)
	case "bool":
		return fmt.Sprintf("%v as bool", expr)
	case "String":
		return fmt.Sprintf("%v as String", expr)
	default:
		return expr
	}
}

func dartFieldNestedClassName(field *core.EmiField, prefixName string) string {
	return prefixName + core.ToUpper(field.Name)
}

// literalParam builds a `this.x = <const>,` (or `this.x,` if no default)
// parameter - used for the cases where the default really is a compile-time
// constant.
func literalParam(name string, constDefault string) string {
	if constDefault == "" {
		return fmt.Sprintf("this.%v,", name)
	}
	return fmt.Sprintf("this.%v = %v,", name, constDefault)
}

// computedParam builds a bare nullable parameter (no `this.`) paired with an
// initializer-list expression - used whenever the default isn't a
// compile-time constant.
func computedParam(typeWithoutOptional string, name string) string {
	return fmt.Sprintf("%v? %v,", typeWithoutOptional, name)
}

// dartResolveField computes everything the class generator needs to render
// one field: its declaration, constructor parameter (+ optional
// initializer), and fromJson/toJson expressions, given the already-decided
// nested/self class names and the complexes recognized in this module.
func dartResolveField(field *core.EmiField, prefixName string, selfClassName string, complexes []RecognizedComplex) dartFieldPlan {
	nestedClassName := dartFieldNestedClassName(field, prefixName)
	typeHint := dartFieldType(field, nestedClassName, prefixName)
	key := jsonKeyExpr(field)
	nullable := strings.HasSuffix(typeHint, "?") || typeHint == "dynamic"
	bareType := strings.TrimSuffix(typeHint, "?")

	plan := dartFieldPlan{Name: field.Name, TypeHint: typeHint, FieldDecl: fmt.Sprintf("%v %v;", typeHint, field.Name)}

	// complex fields (e.g. `complex: "+Vector3"`) - only ever trusted as a
	// real generated type when resolvable to an import; otherwise `dynamic`.
	if field.Complex != "" {
		symbol, resolved := resolveComplexSymbol(field, complexes)
		if resolved {
			plan.TypeHint = symbol + "?"
			plan.FieldDecl = fmt.Sprintf("%v? %v;", symbol, field.Name)
			plan.CtorParam = literalParam(field.Name, "")
			plan.FromJson = fmt.Sprintf("%v == null ? null : %v.fromJson(%v as Map<String, dynamic>)", key, symbol, key)
			plan.ToJson = fmt.Sprintf("%v?.toJson()", field.Name)
		} else {
			plan.TypeHint = "dynamic"
			plan.FieldDecl = fmt.Sprintf("dynamic %v;", field.Name)
			plan.CtorParam = literalParam(field.Name, "")
			plan.FromJson = key
			plan.ToJson = field.Name
		}
		return plan
	}

	switch {
	case field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable:
		base := dartEnumBaseType(field, prefixName)
		if base == "String" {
			plan.CtorParam = literalParam(field.Name, ifStr(!nullable, "''"))
			if nullable {
				plan.FromJson = fmt.Sprintf("%v as String?", key)
			} else {
				plan.FromJson = fmt.Sprintf("%v as String? ?? ''", key)
			}
			plan.ToJson = field.Name
			return plan
		}

		if field.Target == "" {
			plan.InlineEnum = dartRenderEnumFromInline(base, field.OfType)
		}

		if nullable {
			plan.CtorParam = literalParam(field.Name, "")
			plan.FromJson = fmt.Sprintf("%v.fromValueOrNull(%v as String?)", base, key)
			plan.ToJson = fmt.Sprintf("%v?.value", field.Name)
		} else {
			// `X.values.first` isn't a compile-time constant.
			plan.CtorParam = computedParam(base, field.Name)
			plan.InitializerExpr = fmt.Sprintf("%v ?? %v.values.first", field.Name, base)
			plan.FromJson = fmt.Sprintf("%v.fromValue(%v as String? ?? '')", base, key)
			plan.ToJson = fmt.Sprintf("%v.value", field.Name)
		}
		return plan

	case field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass ||
		field.Type == core.FieldTypeOneNullable || field.Type == core.FieldTypeClassNullable:
		target := field.Target
		if target == "" {
			plan.FieldDecl = fmt.Sprintf("dynamic %v;", field.Name)
			plan.CtorParam = literalParam(field.Name, "")
			plan.FromJson = key
			plan.ToJson = field.Name
			return plan
		}
		// A single relation is never eagerly default-constructed, so it's
		// always a plain nullable `this.` param - no initializer needed.
		plan.CtorParam = literalParam(field.Name, "")
		plan.FromJson = fmt.Sprintf("%v == null ? null : %v.fromJson(%v as Map<String, dynamic>)", key, target, key)
		plan.ToJson = fmt.Sprintf("%v?.toJson()", field.Name)
		return plan

	case field.Type == core.FieldTypeCollection || field.Type == core.FieldTypeCollectionNullable:
		target := field.Target
		elemType := "dynamic"
		var elemFrom string
		if target != "" {
			elemType = target
			elemFrom = fmt.Sprintf("%v.fromJson(e as Map<String, dynamic>)", target)
		}

		if nullable {
			plan.CtorParam = literalParam(field.Name, "")
			if target != "" {
				plan.FromJson = fmt.Sprintf("(%v as List<dynamic>?)?.map((e) => %v).toList()", key, elemFrom)
				plan.ToJson = fmt.Sprintf("%v?.map((e) => e.toJson()).toList()", field.Name)
			} else {
				plan.FromJson = fmt.Sprintf("%v as List<dynamic>?", key)
				plan.ToJson = field.Name
			}
			return plan
		}

		plan.CtorParam = computedParam(fmt.Sprintf("List<%v>", elemType), field.Name)
		plan.InitializerExpr = fmt.Sprintf("%v ?? <%v>[]", field.Name, elemType)
		if target != "" {
			plan.FromJson = fmt.Sprintf("(%v as List<dynamic>?)?.map((e) => %v).toList() ?? <%v>[]", key, elemFrom, elemType)
			plan.ToJson = fmt.Sprintf("%v.map((e) => e.toJson()).toList()", field.Name)
		} else {
			plan.FromJson = fmt.Sprintf("(%v as List<dynamic>?) ?? <dynamic>[]", key)
			plan.ToJson = field.Name
		}
		return plan

	case field.Type == core.FieldTypeArray || field.Type == core.FieldTypeArrayNullable ||
		field.Type == core.FieldTypeList || field.Type == core.FieldTypeListNullable:
		elemFrom := fmt.Sprintf("%v.fromJson(e as Map<String, dynamic>)", nestedClassName)
		if nullable {
			plan.CtorParam = literalParam(field.Name, "")
			plan.FromJson = fmt.Sprintf("(%v as List<dynamic>?)?.map((e) => %v).toList()", key, elemFrom)
			plan.ToJson = fmt.Sprintf("%v?.map((e) => e.toJson()).toList()", field.Name)
			return plan
		}
		plan.CtorParam = computedParam(fmt.Sprintf("List<%v>", nestedClassName), field.Name)
		plan.InitializerExpr = fmt.Sprintf("%v ?? <%v>[]", field.Name, nestedClassName)
		plan.FromJson = fmt.Sprintf("(%v as List<dynamic>?)?.map((e) => %v).toList() ?? <%v>[]", key, elemFrom, nestedClassName)
		plan.ToJson = fmt.Sprintf("%v.map((e) => e.toJson()).toList()", field.Name)
		return plan

	case field.Type == core.FieldTypeSlice || field.Type == core.FieldTypeSliceNullable:
		primitive := extractPrimitive(field.Primitive)
		if primitive == "" {
			primitive = "dynamic"
		}
		elemFrom := dartPrimitiveCast(primitive, "e")
		if nullable {
			plan.CtorParam = literalParam(field.Name, "")
			plan.FromJson = fmt.Sprintf("(%v as List<dynamic>?)?.map((e) => %v).toList()", key, elemFrom)
			plan.ToJson = field.Name
			return plan
		}
		plan.CtorParam = computedParam(fmt.Sprintf("List<%v>", primitive), field.Name)
		plan.InitializerExpr = fmt.Sprintf("%v ?? <%v>[]", field.Name, primitive)
		plan.FromJson = fmt.Sprintf("(%v as List<dynamic>?)?.map((e) => %v).toList() ?? <%v>[]", key, elemFrom, primitive)
		plan.ToJson = field.Name
		return plan

	case field.Type == core.FieldTypeObject || field.Type == core.FieldTypeObjectNullable:
		fromExpr := fmt.Sprintf("%v.fromJson(%v as Map<String, dynamic>)", nestedClassName, key)
		selfReferencing := nestedClassName == selfClassName

		if nullable || selfReferencing {
			// Self-referencing (infinite recursion risk) or genuinely
			// optional: always plain nullable, never eagerly constructed.
			plan.TypeHint = nestedClassName + "?"
			plan.FieldDecl = fmt.Sprintf("%v? %v;", nestedClassName, field.Name)
			plan.CtorParam = literalParam(field.Name, "")
			plan.FromJson = fmt.Sprintf("%v == null ? null : %v", key, fromExpr)
			plan.ToJson = fmt.Sprintf("%v?.toJson()", field.Name)
			return plan
		}

		plan.CtorParam = computedParam(nestedClassName, field.Name)
		plan.InitializerExpr = fmt.Sprintf("%v ?? %v()", field.Name, nestedClassName)
		plan.FromJson = fmt.Sprintf("%v == null ? %v() : %v", key, nestedClassName, fromExpr)
		plan.ToJson = fmt.Sprintf("%v.toJson()", field.Name)
		return plan

	case field.Type == core.FieldTypeMap || field.Type == core.FieldTypeMapNullable:
		valueType, ok := strings.CutPrefix(bareType, "Map<String, ")
		if !ok {
			valueType = "dynamic"
		} else {
			valueType = strings.TrimSuffix(valueType, ">")
		}

		// A `Map<String, dynamic>` is not assignable to e.g. `Map<String,
		// int>` in dart (generics are invariant here) - every non-dynamic
		// value has to be individually cast/coerced, exactly like a List
		// element above.
		var fromMapExpr string
		if valueType == "dynamic" {
			fromMapExpr = fmt.Sprintf("(%v as Map<String, dynamic>?)", key)
		} else {
			elemCast := dartPrimitiveCast(valueType, "v")
			fromMapExpr = fmt.Sprintf("(%v as Map<String, dynamic>?)?.map((k, v) => MapEntry(k, %v))", key, elemCast)
		}

		if nullable {
			plan.CtorParam = literalParam(field.Name, "")
			plan.FromJson = fromMapExpr
			plan.ToJson = field.Name
			return plan
		}

		plan.CtorParam = computedParam(fmt.Sprintf("Map<String, %v>", valueType), field.Name)
		plan.InitializerExpr = fmt.Sprintf("%v ?? <String, %v>{}", field.Name, valueType)
		plan.FromJson = fmt.Sprintf("%v ?? <String, %v>{}", fromMapExpr, valueType)
		plan.ToJson = field.Name
		return plan

	default:
		// Primitive (string/int/float/bool/any).
		plan.ToJson = field.Name
		if typeHint == "dynamic" {
			plan.CtorParam = literalParam(field.Name, "")
			plan.FromJson = key
			return plan
		}
		if nullable {
			plan.CtorParam = literalParam(field.Name, "")
			switch typeHint {
			case "int?":
				plan.FromJson = fmt.Sprintf("(%v as num?)?.toInt()", key)
			case "double?":
				plan.FromJson = fmt.Sprintf("(%v as num?)?.toDouble()", key)
			default:
				plan.FromJson = fmt.Sprintf("%v as %v", key, typeHint)
			}
			return plan
		}

		zero := dartZeroValue(typeHint)
		plan.CtorParam = literalParam(field.Name, zero)
		switch typeHint {
		case "int":
			plan.FromJson = fmt.Sprintf("(%v as num?)?.toInt() ?? %v", key, zero)
		case "double":
			plan.FromJson = fmt.Sprintf("(%v as num?)?.toDouble() ?? %v", key, zero)
		default:
			plan.FromJson = fmt.Sprintf("%v as %v? ?? %v", key, typeHint, zero)
		}
		return plan
	}
}

func ifStr(cond bool, s string) string {
	if cond {
		return s
	}
	return ""
}

func dartZeroValue(dartType string) string {
	switch dartType {
	case "String":
		return "''"
	case "int":
		return "0"
	case "double":
		return "0.0"
	case "bool":
		return "false"
	default:
		return "null"
	}
}

// resolveComplexSymbol mirrors lib/python's resolveComplexTypeHint: a
// `complex:` field is only ever trusted as a real generated type when it's
// `+`-prefixed AND resolvable via RecognizedComplexes (which is how we know
// there's actually an import for it) - anything else falls back to
// `dynamic`.
func resolveComplexSymbol(field *core.EmiField, complexes []RecognizedComplex) (symbol string, resolved bool) {
	symbol = strings.ReplaceAll(field.Complex, "+", "")
	if !strings.Contains(field.Complex, "+") {
		return "", false
	}
	for _, item := range complexes {
		if item.Symbol == symbol {
			return symbol, item.ImportLocation != ""
		}
	}
	return "", false
}
