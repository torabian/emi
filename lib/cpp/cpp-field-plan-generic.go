// Generic-dialect field resolution: computes, for one EmiField, its class member
// declaration plus the ToJson()/FromJson() statements that (de)serialize it
// through cJSON (vendored, same library the C target uses - see
// cpp-include/generic/cjson). C++ gets no reflection-based shortcut the way
// lib/csharp or the unreal dialect (cpp-field-plan-unreal.go) do, so this mirrors
// lib/c/c-field-plan.go's approach, translated to member-function bodies instead
// of `Type_to_json`/`Type_from_json_into` free functions.
//
// `owner` is the field-access prefix the caller wants ("" from inside ToJson(),
// a const member function where fields are unqualified; "result." from inside the
// static FromJson() factory, which builds into a local `result` instance) - see
// cpp-class-generator.go.
package cpp

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

type cppFieldPlan struct {
	Decl       []string                    // class member declaration line(s)
	ToJson     func(owner string) []string // statements appending into the local `cJSON* json` inside ToJson()
	FromJson   func(owner string) []string // statements reading from the local `const cJSON* json` inside FromJson()
	InlineEnum *cppRenderedEnum
}

func jsonKey(name string) string {
	return escapeDoubleQuoted(name)
}

// cppGenericResolveField computes everything the generic-dialect class generator
// needs for one field.
func cppGenericResolveField(field *core.EmiField, prefixName string, selfTypeName string, complexes []RecognizedComplex) cppFieldPlan {
	name := field.Name
	key := jsonKey(name)
	nestedType := cppFieldNestedTypeName(field, prefixName)

	if field.Complex != "" {
		symbol, resolved := resolveComplexSymbol(field, complexes)
		if resolved {
			return cppGenericRelationField(name, key, symbol)
		}
		return cppGenericRawJsonField(name, key)
	}

	switch {
	case field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable:
		base := cppEnumBaseType(field, prefixName)
		plan := cppGenericEnumField(name, key, base, core.IsNullable(string(field.Type)))
		if field.Target == "" {
			plan.InlineEnum = cppRenderEnumFromInline(base, field.OfType)
		}
		return plan

	case field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass ||
		field.Type == core.FieldTypeOneNullable || field.Type == core.FieldTypeClassNullable:
		if field.Target == "" {
			return cppGenericRawJsonField(name, key)
		}
		return cppGenericRelationField(name, key, field.Target)

	case field.Type == core.FieldTypeCollection || field.Type == core.FieldTypeCollectionNullable:
		if field.Target == "" {
			return cppGenericRawJsonField(name, key)
		}
		return cppGenericObjectArrayField(name, key, field.Target, core.IsNullable(string(field.Type)))

	case field.Type == core.FieldTypeArray || field.Type == core.FieldTypeArrayNullable ||
		field.Type == core.FieldTypeList || field.Type == core.FieldTypeListNullable:
		return cppGenericObjectArrayField(name, key, nestedType, core.IsNullable(string(field.Type)))

	case field.Type == core.FieldTypeSlice || field.Type == core.FieldTypeSliceNullable:
		primitive := extractPrimitiveGeneric(field.Primitive)
		if field.Primitive == "string" {
			return cppGenericScalarArrayField(name, key, "std::string", true, core.IsNullable(string(field.Type)))
		}
		if primitive == "" {
			return cppGenericRawJsonField(name, key)
		}
		return cppGenericScalarArrayField(name, key, primitive, false, core.IsNullable(string(field.Type)))

	case field.Type == core.FieldTypeObject || field.Type == core.FieldTypeObjectNullable:
		selfReferencing := nestedType == selfTypeName
		nullable := core.IsNullable(string(field.Type))
		if nullable || selfReferencing {
			return cppGenericRelationField(name, key, nestedType)
		}
		return cppGenericValueObjectField(name, key, nestedType)

	case field.Type == core.FieldTypeMap || field.Type == core.FieldTypeMapNullable:
		return cppGenericMapField(field, name, key)

	case field.Type == core.FieldTypeString:
		return cppGenericStringField(name, key, false)
	case field.Type == core.FieldTypeStringNullable:
		return cppGenericStringField(name, key, true)

	default:
		primitive := extractPrimitiveGeneric(string(field.Type))
		if primitive == "" {
			return cppGenericRawJsonField(name, key)
		}
		return cppGenericScalarField(name, key, primitive, core.IsNullable(string(field.Type)))
	}
}

// ---- string --------------------------------------------------------------

func cppGenericStringField(name, key string, nullable bool) cppFieldPlan {
	cppType := "std::string"
	if nullable {
		cppType = cppNullable(cppType)
	}
	return cppFieldPlan{
		Decl: []string{fmt.Sprintf("%v %v;", cppType, name)},
		ToJson: func(owner string) []string {
			if nullable {
				return []string{
					fmt.Sprintf("if (%v%v.has_value()) { cJSON_AddStringToObject(json, %v, %v%v->c_str()); } else { cJSON_AddNullToObject(json, %v); }", owner, name, key, owner, name, key),
				}
			}
			return []string{fmt.Sprintf("cJSON_AddStringToObject(json, %v, %v%v.c_str());", key, owner, name)}
		},
		FromJson: func(owner string) []string {
			lines := []string{
				"{",
				fmt.Sprintf("    const cJSON* v = cJSON_GetObjectItemCaseSensitive(json, %v);", key),
				"    if (cJSON_IsString(v) && v->valuestring) {",
			}
			if nullable {
				lines = append(lines, fmt.Sprintf("        %v%v = std::string(v->valuestring);", owner, name))
			} else {
				lines = append(lines, fmt.Sprintf("        %v%v = v->valuestring;", owner, name))
			}
			lines = append(lines, "    }", "}")
			return lines
		},
	}
}

// ---- scalar (int/float/bool) ----------------------------------------------

func cppGenericScalarField(name, key, cppType string, nullable bool) cppFieldPlan {
	fullType := cppType
	if nullable {
		fullType = cppNullable(cppType)
	}
	return cppFieldPlan{
		Decl: []string{fmt.Sprintf("%v %v;", fullType, name)},
		ToJson: func(owner string) []string {
			if nullable {
				return []string{
					fmt.Sprintf("if (%v%v.has_value()) { cJSON_AddNumberToObject(json, %v, (double) *%v%v); } else { cJSON_AddNullToObject(json, %v); }", owner, name, key, owner, name, key),
				}
			}
			return []string{fmt.Sprintf("cJSON_AddNumberToObject(json, %v, (double) %v%v);", key, owner, name)}
		},
		FromJson: func(owner string) []string {
			return []string{
				"{",
				fmt.Sprintf("    const cJSON* v = cJSON_GetObjectItemCaseSensitive(json, %v);", key),
				fmt.Sprintf("    if (cJSON_IsNumber(v)) { %v%v = (%v) v->valuedouble; }", owner, name, cppType),
				fmt.Sprintf("    else if (cJSON_IsBool(v)) { %v%v = (%v) cJSON_IsTrue(v); }", owner, name, cppType),
				"}",
			}
		},
	}
}

// ---- enum ---------------------------------------------------------------

func cppGenericEnumField(name, key, enumType string, nullable bool) cppFieldPlan {
	fullType := enumType
	if nullable {
		fullType = cppNullable(enumType)
	}
	return cppFieldPlan{
		Decl: []string{fmt.Sprintf("%v %v;", fullType, name)},
		ToJson: func(owner string) []string {
			if nullable {
				return []string{
					fmt.Sprintf("if (%v%v.has_value()) { cJSON_AddStringToObject(json, %v, %vToString(*%v%v)); } else { cJSON_AddNullToObject(json, %v); }", owner, name, key, enumType, owner, name, key),
				}
			}
			return []string{fmt.Sprintf("cJSON_AddStringToObject(json, %v, %vToString(%v%v));", key, enumType, owner, name)}
		},
		FromJson: func(owner string) []string {
			return []string{
				"{",
				fmt.Sprintf("    const cJSON* v = cJSON_GetObjectItemCaseSensitive(json, %v);", key),
				fmt.Sprintf("    if (cJSON_IsString(v) && v->valuestring) { %v%v = %vFromString(v->valuestring); }", owner, name, enumType),
				"}",
			}
		},
	}
}

// ---- object (nested, by value) --------------------------------------------

// cppGenericValueObjectField covers a genuinely local, nested, non-nullable,
// non-self-referencing `object` field - held by value (default-constructed, so
// FromJson()'s starting `result` already has one, just like every other field).
func cppGenericValueObjectField(name, key, nestedType string) cppFieldPlan {
	return cppFieldPlan{
		Decl: []string{fmt.Sprintf("%v %v;", nestedType, name)},
		ToJson: func(owner string) []string {
			return []string{fmt.Sprintf("cJSON_AddItemToObject(json, %v, %v%v.ToJson());", key, owner, name)}
		},
		FromJson: func(owner string) []string {
			return []string{
				"{",
				fmt.Sprintf("    const cJSON* v = cJSON_GetObjectItemCaseSensitive(json, %v);", key),
				fmt.Sprintf("    if (cJSON_IsObject(v)) { %v%v = %v::FromJson(v); }", owner, name, nestedType),
				"}",
			}
		},
	}
}

// ---- one/class relation, and a nullable/self-referencing object -----------
//
// std::unique_ptr<T> is used uniformly for both cases: a single relation is
// never eagerly constructed regardless of what the schema claims (matching every
// sibling generator's own rule), and a self-referencing object member can only
// ever be held through a pointer in the first place (an incomplete type can't be
// a by-value member).

func cppGenericRelationField(name, key, targetType string) cppFieldPlan {
	cppType := fmt.Sprintf("std::unique_ptr<%v>", targetType)
	return cppFieldPlan{
		Decl: []string{fmt.Sprintf("%v %v;", cppType, name)},
		ToJson: func(owner string) []string {
			return []string{
				fmt.Sprintf("if (%v%v) { cJSON_AddItemToObject(json, %v, %v%v->ToJson()); }", owner, name, key, owner, name),
			}
		},
		FromJson: func(owner string) []string {
			return []string{
				"{",
				fmt.Sprintf("    const cJSON* v = cJSON_GetObjectItemCaseSensitive(json, %v);", key),
				fmt.Sprintf("    if (cJSON_IsObject(v)) { %v%v = std::make_unique<%v>(%v::FromJson(v)); }", owner, name, targetType, targetType),
				"}",
			}
		},
	}
}

// ---- array/_list (locally nested elements) and collection (target elements) -

func cppGenericObjectArrayField(name, key, elemType string, nullable bool) cppFieldPlan {
	vecType := fmt.Sprintf("std::vector<%v>", elemType)
	cppType := vecType
	if nullable {
		cppType = cppNullable(vecType)
	}
	return cppFieldPlan{
		Decl: []string{fmt.Sprintf("%v %v;", cppType, name)},
		ToJson: func(owner string) []string {
			itemsExpr := owner + name
			if nullable {
				itemsExpr = "(*" + owner + name + ")"
			}
			lines := []string{}
			if nullable {
				lines = append(lines, fmt.Sprintf("if (%v%v.has_value()) {", owner, name))
			}
			lines = append(lines,
				"{",
				fmt.Sprintf("    cJSON* arr = cJSON_AddArrayToObject(json, %v);", key),
				fmt.Sprintf("    for (const auto& item : %v) { cJSON_AddItemToArray(arr, item.ToJson()); }", itemsExpr),
				"}",
			)
			if nullable {
				lines = append(lines, "} else {", fmt.Sprintf("    cJSON_AddNullToObject(json, %v);", key), "}")
			}
			return lines
		},
		FromJson: func(owner string) []string {
			assignTarget := owner + name
			lines := []string{
				"{",
				fmt.Sprintf("    const cJSON* v = cJSON_GetObjectItemCaseSensitive(json, %v);", key),
				"    if (cJSON_IsArray(v)) {",
				fmt.Sprintf("        %v items;", vecType),
				"        const cJSON* item = NULL;",
				"        cJSON_ArrayForEach(item, v) {",
				fmt.Sprintf("            items.push_back(%v::FromJson(item));", elemType),
				"        }",
			}
			if nullable {
				lines = append(lines, fmt.Sprintf("        %v = std::move(items);", assignTarget))
			} else {
				lines = append(lines, fmt.Sprintf("        %v = std::move(items);", assignTarget))
			}
			lines = append(lines, "    }", "}")
			return lines
		},
	}
}

// ---- slice (array of a primitive) ------------------------------------------

func cppGenericScalarArrayField(name, key, elemType string, isString bool, nullable bool) cppFieldPlan {
	vecType := fmt.Sprintf("std::vector<%v>", elemType)
	cppType := vecType
	if nullable {
		cppType = cppNullable(vecType)
	}

	addItem := "cJSON_AddItemToArray(arr, cJSON_CreateNumber((double) item));"
	if isString {
		addItem = "cJSON_AddItemToArray(arr, cJSON_CreateString(item.c_str()));"
	}
	readItem := fmt.Sprintf("items.push_back(cJSON_IsNumber(item) ? (%v) item->valuedouble : (%v) 0);", elemType, elemType)
	if isString {
		readItem = `items.push_back(cJSON_IsString(item) && item->valuestring ? item->valuestring : "");`
	}

	return cppFieldPlan{
		Decl: []string{fmt.Sprintf("%v %v;", cppType, name)},
		ToJson: func(owner string) []string {
			itemsExpr := owner + name
			if nullable {
				itemsExpr = "(*" + owner + name + ")"
			}
			lines := []string{}
			if nullable {
				lines = append(lines, fmt.Sprintf("if (%v%v.has_value()) {", owner, name))
			}
			lines = append(lines,
				"{",
				fmt.Sprintf("    cJSON* arr = cJSON_AddArrayToObject(json, %v);", key),
				fmt.Sprintf("    for (const auto& item : %v) { %v }", itemsExpr, addItem),
				"}",
			)
			if nullable {
				lines = append(lines, "} else {", fmt.Sprintf("    cJSON_AddNullToObject(json, %v);", key), "}")
			}
			return lines
		},
		FromJson: func(owner string) []string {
			return []string{
				"{",
				fmt.Sprintf("    const cJSON* v = cJSON_GetObjectItemCaseSensitive(json, %v);", key),
				"    if (cJSON_IsArray(v)) {",
				fmt.Sprintf("        %v items;", vecType),
				"        const cJSON* item = NULL;",
				"        cJSON_ArrayForEach(item, v) {",
				fmt.Sprintf("            %v", readItem),
				"        }",
				fmt.Sprintf("        %v%v = std::move(items);", owner, name),
				"    }",
				"}",
			}
		},
	}
}

// ---- map -------------------------------------------------------------------
//
// Only a primitive-keyed, primitive-or-object-valued map resolves to a real
// std::map<K, V> - anything else (an unresolved value shape) falls back to
// emi::EmiJson, same as an unresolved `any`/`complex` field.

func cppGenericMapField(field *core.EmiField, name, key string) cppFieldPlan {
	nullable := core.IsNullable(string(field.Type))
	keyType := extractPrimitiveGeneric(field.MapKeyOf)
	if keyType == "" || field.MapKeyOf == "" {
		keyType = "std::string"
	}
	valueType := extractPrimitiveGeneric(field.MapPairOf)
	if valueType == "" {
		return cppGenericRawJsonField(name, key)
	}

	mapType := fmt.Sprintf("std::map<%v, %v>", keyType, valueType)
	cppType := mapType
	if nullable {
		cppType = cppNullable(mapType)
	}

	return cppFieldPlan{
		Decl: []string{fmt.Sprintf("%v %v;", cppType, name)},
		ToJson: func(owner string) []string {
			itemsExpr := owner + name
			if nullable {
				itemsExpr = "(*" + owner + name + ")"
			}
			lines := []string{}
			if nullable {
				lines = append(lines, fmt.Sprintf("if (%v%v.has_value()) {", owner, name))
			}
			lines = append(lines,
				"{",
				fmt.Sprintf("    cJSON* obj = cJSON_AddObjectToObject(json, %v);", key),
				fmt.Sprintf("    for (const auto& kv : %v) {", itemsExpr),
				fmt.Sprintf("        std::string mapKey = %v;", mapKeyToStringExpr(keyType)),
				fmt.Sprintf("        cJSON_AddItemToObject(obj, mapKey.c_str(), %v);", mapValueToJsonExpr(valueType)),
				"    }",
				"}",
			)
			if nullable {
				lines = append(lines, "} else {", fmt.Sprintf("    cJSON_AddNullToObject(json, %v);", key), "}")
			}
			return lines
		},
		FromJson: func(owner string) []string {
			return []string{
				"{",
				fmt.Sprintf("    const cJSON* v = cJSON_GetObjectItemCaseSensitive(json, %v);", key),
				"    if (cJSON_IsObject(v)) {",
				fmt.Sprintf("        %v items;", mapType),
				"        const cJSON* entry = NULL;",
				"        cJSON_ArrayForEach(entry, v) {",
				fmt.Sprintf("            items[%v] = %v;", mapKeyFromStringExpr(keyType, "entry->string"), mapValueFromJsonExpr(valueType, "entry")),
				"        }",
				fmt.Sprintf("        %v%v = std::move(items);", owner, name),
				"    }",
				"}",
			}
		},
	}
}

func mapKeyToStringExpr(keyType string) string {
	if keyType == "std::string" {
		return "kv.first"
	}
	return "std::to_string(kv.first)"
}

func mapKeyFromStringExpr(keyType, raw string) string {
	if keyType == "std::string" {
		return raw
	}
	return fmt.Sprintf("(%v) std::stoll(%v)", keyType, raw)
}

func mapValueToJsonExpr(valueType string) string {
	if valueType == "std::string" {
		return "cJSON_CreateString(kv.second.c_str())"
	}
	if valueType == "bool" {
		return "cJSON_CreateBool(kv.second)"
	}
	return "cJSON_CreateNumber((double) kv.second)"
}

func mapValueFromJsonExpr(valueType, entryVar string) string {
	if valueType == "std::string" {
		return fmt.Sprintf("(cJSON_IsString(%v) && %v->valuestring ? %v->valuestring : \"\")", entryVar, entryVar, entryVar)
	}
	if valueType == "bool" {
		return fmt.Sprintf("cJSON_IsTrue(%v)", entryVar)
	}
	return fmt.Sprintf("(%v) (cJSON_IsNumber(%v) ? %v->valuedouble : 0)", valueType, entryVar, entryVar)
}

// ---- any/complex (unresolved), and every unresolvable structural fallback --

func cppGenericRawJsonField(name, key string) cppFieldPlan {
	return cppFieldPlan{
		Decl: []string{fmt.Sprintf("emi::EmiJson %v;", name)},
		ToJson: func(owner string) []string {
			return []string{fmt.Sprintf("cJSON_AddItemToObject(json, %v, %v%v.Clone());", key, owner, name)}
		},
		FromJson: func(owner string) []string {
			return []string{
				"{",
				fmt.Sprintf("    const cJSON* v = cJSON_GetObjectItemCaseSensitive(json, %v);", key),
				fmt.Sprintf("    if (v) { %v%v = emi::EmiJson(cJSON_Duplicate(v, true)); }", owner, name),
				"}",
			}
		},
	}
}
