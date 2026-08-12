// Renders one field of a generated C struct: the member declaration(s) (an
// array/relation-list field is *two* members - a pointer plus a `_count`),
// and the four function bodies every struct gets (see c-struct-generator.go):
// `_init` (safe defaults, no allocation of `self` itself), `_free_fields`
// (frees everything owned, without freeing `self`), `_to_json` (build a
// cJSON tree) and `_from_json_into` (populate `self` from a cJSON tree).
//
// Every struct always gets *some* default for every field (see the sibling
// generators' PythonSafeDefaultValue/dartResolveField/etc.) - here that's
// what `_init` guarantees, so `Type_new()` is always immediately usable and
// `Type_free()` is always safe to call on it.
package c

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type cFieldPlan struct {
	Decl       []string // struct member declaration line(s)
	Init       []string // `_init` function body line(s)
	FreeFields []string // `_free_fields` function body line(s)
	ToJson     []string // `_to_json` function body line(s) - builds into local `cJSON *json`
	FromJson   []string // `_from_json_into` function body line(s) - reads from local `const cJSON *json`
	InlineEnum *cRenderedEnum
}

func cFieldNestedTypeName(field *core.EmiField, prefixName string) string {
	return prefixName + core.ToUpper(field.Name)
}

// cStructTypeRef spells a struct type name for use inside a *member
// declaration*. A field pointing back at the struct it's declared in (a
// self-referencing relation/collection) needs the `struct` tag - at that
// point in the source, `typedef struct X { ... } X;` hasn't finished yet,
// so the bare typedef alias `X` doesn't exist as a type yet, only the tag
// `struct X` does. Every other reference (used inside a function body,
// which always comes after the complete typedef) can use the bare alias.
func cStructTypeRef(typeName, selfTypeName string) string {
	if typeName == selfTypeName {
		return "struct " + typeName
	}
	return typeName
}

func jsonKey(name string) string {
	b, _ := jsonMarshal(name)
	return b
}

// jsonMarshal is a tiny wrapper so this file doesn't need its own
// encoding/json import just for one call site.
func jsonMarshal(s string) (string, error) {
	return escapeDoubleQuoted(s), nil
}

// cResolveField computes everything the struct generator needs for one
// field: its declaration, and its `_init`/`_free_fields`/`_to_json`/
// `_from_json_into` contributions.
func cResolveField(field *core.EmiField, prefixName string, selfTypeName string, complexes []RecognizedComplex) cFieldPlan {
	name := field.Name
	key := jsonKey(name)
	nestedType := cFieldNestedTypeName(field, prefixName)

	if field.Complex != "" {
		symbol, resolved := resolveComplexSymbol(field, complexes)
		if resolved {
			return cRelationField(name, key, symbol, selfTypeName)
		}
		return cRawJsonField(name, key)
	}

	switch {
	case field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable:
		base := cEnumBaseType(field, prefixName)
		plan := cEnumField(name, key, base)
		if field.Target == "" {
			// Unlike every other target, a C enum always needs an actual
			// declared type - there's no plain-string fallback to fall
			// back to - so even a field with neither `target:` nor `of:`
			// gets its own (empty) companion enum.
			plan.InlineEnum = cRenderEnumFromInline(base, field.OfType)
		}
		return plan

	case field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass ||
		field.Type == core.FieldTypeOneNullable || field.Type == core.FieldTypeClassNullable:
		if field.Target == "" {
			return cRawJsonField(name, key)
		}
		return cRelationField(name, key, field.Target, selfTypeName)

	case field.Type == core.FieldTypeCollection || field.Type == core.FieldTypeCollectionNullable:
		if field.Target == "" {
			return cRawJsonField(name, key)
		}
		return cObjectArrayField(name, key, field.Target, selfTypeName)

	case field.Type == core.FieldTypeArray || field.Type == core.FieldTypeArrayNullable ||
		field.Type == core.FieldTypeList || field.Type == core.FieldTypeListNullable:
		return cObjectArrayField(name, key, nestedType, selfTypeName)

	case field.Type == core.FieldTypeSlice || field.Type == core.FieldTypeSliceNullable:
		primitive := extractPrimitive(field.Primitive)
		if field.Primitive == "string" {
			return cStringArrayField(name, key)
		}
		if primitive == "" {
			return cRawJsonField(name, key)
		}
		return cScalarArrayField(name, key, primitive)

	case field.Type == core.FieldTypeObject || field.Type == core.FieldTypeObjectNullable:
		selfReferencing := nestedType == selfTypeName
		nullable := core.IsNullable(string(field.Type))
		return cObjectField(name, key, nestedType, nullable || selfReferencing)

	case field.Type == core.FieldTypeMap || field.Type == core.FieldTypeMapNullable:
		return cRawJsonField(name, key)

	case field.Type == core.FieldTypeString:
		return cStringField(name, key, false)
	case field.Type == core.FieldTypeStringNullable:
		return cStringField(name, key, true)

	default:
		primitive := extractPrimitive(string(field.Type))
		if primitive == "" {
			return cRawJsonField(name, key)
		}
		return cScalarField(name, key, primitive)
	}
}

// cStringField covers `string`/`string?`. Unlike int/float/bool/enum (see
// extractPrimitive's doc comment), a string is already a pointer, so there's
// no extra cost to making it properly nullable - `string?` really does
// default to (and round-trip) NULL/JSON-null, not "".
func cStringField(name, key string, nullable bool) cFieldPlan {
	init := fmt.Sprintf(`self->%s = strdup("");`, name)
	toJSON := fmt.Sprintf(`cJSON_AddStringToObject(json, %s, self->%s ? self->%s : "");`, key, name, name)
	if nullable {
		init = fmt.Sprintf("self->%s = NULL;", name)
		toJSON = fmt.Sprintf(
			"if (self->%s) { cJSON_AddStringToObject(json, %s, self->%s); } else { cJSON_AddNullToObject(json, %s); }",
			name, key, name, key,
		)
	}

	return cFieldPlan{
		Decl: []string{fmt.Sprintf("char* %s;", name)},
		Init: []string{init},
		FreeFields: []string{
			fmt.Sprintf("free(self->%s);", name),
			fmt.Sprintf("self->%s = NULL;", name),
		},
		ToJson: []string{toJSON},
		FromJson: []string{
			"{",
			fmt.Sprintf("    const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, %s);", key),
			"    if (cJSON_IsString(v) && v->valuestring) {",
			fmt.Sprintf("        free(self->%s);", name),
			fmt.Sprintf("        self->%s = strdup(v->valuestring);", name),
			"    }",
			"}",
		},
	}
}

func cScalarField(name, key, ctype string) cFieldPlan {
	zero := "0"
	if ctype == "bool" {
		zero = "false"
	}
	return cFieldPlan{
		Decl: []string{fmt.Sprintf("%s %s;", ctype, name)},
		Init: []string{fmt.Sprintf("self->%s = %s;", name, zero)},
		ToJson: []string{
			fmt.Sprintf("cJSON_AddNumberToObject(json, %s, (double) self->%s);", key, name),
		},
		FromJson: []string{
			"{",
			fmt.Sprintf("    const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, %s);", key),
			fmt.Sprintf("    if (cJSON_IsNumber(v)) { self->%s = (%s) v->valuedouble; }", name, ctype),
			fmt.Sprintf("    else if (cJSON_IsBool(v)) { self->%s = (%s) cJSON_IsTrue(v); }", name, ctype),
			"}",
		},
	}
}

// cRelationField covers `one`/`class` and a resolved `complex:` field: a
// single pointer to another generated struct, never eagerly constructed
// (see cObjectField for the contrast with a genuinely local nested object).
func cRelationField(name, key, targetType, selfTypeName string) cFieldPlan {
	return cFieldPlan{
		Decl: []string{fmt.Sprintf("%s* %s;", cStructTypeRef(targetType, selfTypeName), name)},
		Init: []string{fmt.Sprintf("self->%s = NULL;", name)},
		FreeFields: []string{
			fmt.Sprintf("%s_free(self->%s);", targetType, name),
			fmt.Sprintf("self->%s = NULL;", name),
		},
		ToJson: []string{
			fmt.Sprintf("if (self->%s) { cJSON_AddItemToObject(json, %s, %s_to_json(self->%s)); }", name, key, targetType, name),
		},
		FromJson: []string{
			"{",
			fmt.Sprintf("    const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, %s);", key),
			"    if (cJSON_IsObject(v)) {",
			fmt.Sprintf("        %s_free(self->%s);", targetType, name),
			fmt.Sprintf("        self->%s = %s_from_json(v);", name, targetType),
			"    }",
			"}",
		},
	}
}

// cObjectField covers a genuinely local, nested `object` field - a pointer,
// eagerly allocated by `_init` unless it's nullable or self-referencing (in
// which case eager construction would either be pointless or recurse
// forever), exactly the same reasoning the other generators apply.
func cObjectField(name, key, nestedType string, alwaysNull bool) cFieldPlan {
	initLine := fmt.Sprintf("self->%s = %s_new();", name, nestedType)
	if alwaysNull {
		initLine = fmt.Sprintf("self->%s = NULL;", name)
	}
	return cFieldPlan{
		Decl: []string{fmt.Sprintf("%s* %s;", nestedType, name)},
		Init: []string{initLine},
		FreeFields: []string{
			fmt.Sprintf("%s_free(self->%s);", nestedType, name),
			fmt.Sprintf("self->%s = NULL;", name),
		},
		ToJson: []string{
			fmt.Sprintf("if (self->%s) { cJSON_AddItemToObject(json, %s, %s_to_json(self->%s)); }", name, key, nestedType, name),
		},
		FromJson: []string{
			"{",
			fmt.Sprintf("    const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, %s);", key),
			"    if (cJSON_IsObject(v)) {",
			fmt.Sprintf("        %s_free(self->%s);", nestedType, name),
			fmt.Sprintf("        self->%s = %s_from_json(v);", name, nestedType),
			"    }",
			"}",
		},
	}
}

// cObjectArrayField covers `array`/`_list` (locally nested elements) and
// `collection` (elements of another generated struct) - a contiguous array
// of structs *by value* (not an array of pointers), stored alongside its
// own `_count`.
func cObjectArrayField(name, key, elemType, selfTypeName string) cFieldPlan {
	countField := name + "_count"
	return cFieldPlan{
		Decl: []string{
			fmt.Sprintf("%s* %s;", cStructTypeRef(elemType, selfTypeName), name),
			fmt.Sprintf("size_t %s;", countField),
		},
		Init: []string{
			fmt.Sprintf("self->%s = NULL;", name),
			fmt.Sprintf("self->%s = 0;", countField),
		},
		FreeFields: []string{
			fmt.Sprintf("for (size_t i = 0; i < self->%s; i++) { %s_free_fields(&self->%s[i]); }", countField, elemType, name),
			fmt.Sprintf("free(self->%s);", name),
			fmt.Sprintf("self->%s = NULL;", name),
			fmt.Sprintf("self->%s = 0;", countField),
		},
		ToJson: []string{
			"{",
			fmt.Sprintf("    cJSON *arr = cJSON_AddArrayToObject(json, %s);", key),
			fmt.Sprintf("    for (size_t i = 0; i < self->%s; i++) { cJSON_AddItemToArray(arr, %s_to_json(&self->%s[i])); }", countField, elemType, name),
			"}",
		},
		FromJson: []string{
			"{",
			fmt.Sprintf("    const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, %s);", key),
			"    if (cJSON_IsArray(v)) {",
			fmt.Sprintf("        for (size_t i = 0; i < self->%s; i++) { %s_free_fields(&self->%s[i]); }", countField, elemType, name),
			fmt.Sprintf("        free(self->%s);", name),
			"        size_t n = (size_t) cJSON_GetArraySize(v);",
			fmt.Sprintf("        self->%s = n ? (%s*) calloc(n, sizeof(%s)) : NULL;", name, elemType, elemType),
			fmt.Sprintf("        self->%s = n;", countField),
			"        size_t i = 0;",
			"        const cJSON *item = NULL;",
			"        cJSON_ArrayForEach(item, v) {",
			fmt.Sprintf("            %s_init(&self->%s[i]);", elemType, name),
			fmt.Sprintf("            %s_from_json_into(&self->%s[i], item);", elemType, name),
			"            i++;",
			"        }",
			"    }",
			"}",
		},
	}
}

func cScalarArrayField(name, key, ctype string) cFieldPlan {
	countField := name + "_count"
	return cFieldPlan{
		Decl: []string{
			fmt.Sprintf("%s* %s;", ctype, name),
			fmt.Sprintf("size_t %s;", countField),
		},
		Init: []string{
			fmt.Sprintf("self->%s = NULL;", name),
			fmt.Sprintf("self->%s = 0;", countField),
		},
		FreeFields: []string{
			fmt.Sprintf("free(self->%s);", name),
			fmt.Sprintf("self->%s = NULL;", name),
			fmt.Sprintf("self->%s = 0;", countField),
		},
		ToJson: []string{
			"{",
			fmt.Sprintf("    cJSON *arr = cJSON_AddArrayToObject(json, %s);", key),
			fmt.Sprintf("    for (size_t i = 0; i < self->%s; i++) { cJSON_AddItemToArray(arr, cJSON_CreateNumber((double) self->%s[i])); }", countField, name),
			"}",
		},
		FromJson: []string{
			"{",
			fmt.Sprintf("    const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, %s);", key),
			"    if (cJSON_IsArray(v)) {",
			fmt.Sprintf("        free(self->%s);", name),
			"        size_t n = (size_t) cJSON_GetArraySize(v);",
			fmt.Sprintf("        self->%s = n ? (%s*) calloc(n, sizeof(%s)) : NULL;", name, ctype, ctype),
			fmt.Sprintf("        self->%s = n;", countField),
			"        size_t i = 0;",
			"        const cJSON *item = NULL;",
			"        cJSON_ArrayForEach(item, v) {",
			fmt.Sprintf("            self->%s[i] = cJSON_IsNumber(item) ? (%s) item->valuedouble : (%s) 0;", name, ctype, ctype),
			"            i++;",
			"        }",
			"    }",
			"}",
		},
	}
}

func cStringArrayField(name, key string) cFieldPlan {
	countField := name + "_count"
	return cFieldPlan{
		Decl: []string{
			fmt.Sprintf("char** %s;", name),
			fmt.Sprintf("size_t %s;", countField),
		},
		Init: []string{
			fmt.Sprintf("self->%s = NULL;", name),
			fmt.Sprintf("self->%s = 0;", countField),
		},
		FreeFields: []string{
			fmt.Sprintf("for (size_t i = 0; i < self->%s; i++) { free(self->%s[i]); }", countField, name),
			fmt.Sprintf("free(self->%s);", name),
			fmt.Sprintf("self->%s = NULL;", name),
			fmt.Sprintf("self->%s = 0;", countField),
		},
		ToJson: []string{
			"{",
			fmt.Sprintf("    cJSON *arr = cJSON_AddArrayToObject(json, %s);", key),
			fmt.Sprintf("    for (size_t i = 0; i < self->%s; i++) { cJSON_AddItemToArray(arr, cJSON_CreateString(self->%s[i] ? self->%s[i] : \"\")); }", countField, name, name),
			"}",
		},
		FromJson: []string{
			"{",
			fmt.Sprintf("    const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, %s);", key),
			"    if (cJSON_IsArray(v)) {",
			fmt.Sprintf("        for (size_t i = 0; i < self->%s; i++) { free(self->%s[i]); }", countField, name),
			fmt.Sprintf("        free(self->%s);", name),
			"        size_t n = (size_t) cJSON_GetArraySize(v);",
			fmt.Sprintf("        self->%s = n ? (char**) calloc(n, sizeof(char*)) : NULL;", name),
			fmt.Sprintf("        self->%s = n;", countField),
			"        size_t i = 0;",
			"        const cJSON *item = NULL;",
			"        cJSON_ArrayForEach(item, v) {",
			fmt.Sprintf("            self->%s[i] = strdup(cJSON_IsString(item) && item->valuestring ? item->valuestring : \"\");", name),
			"            i++;",
			"        }",
			"    }",
			"}",
		},
	}
}

// cRawJsonField covers `any`/`complex` (unresolved), and every unresolvable
// structural fallback (a targetless relation/collection, an unresolved
// `map`) - the field just keeps a deep copy of whatever cJSON node was
// there, which is always representable regardless of shape.
func cRawJsonField(name, key string) cFieldPlan {
	return cFieldPlan{
		Decl: []string{fmt.Sprintf("cJSON* %s;", name)},
		Init: []string{fmt.Sprintf("self->%s = NULL;", name)},
		FreeFields: []string{
			fmt.Sprintf("cJSON_Delete(self->%s);", name),
			fmt.Sprintf("self->%s = NULL;", name),
		},
		ToJson: []string{
			fmt.Sprintf("if (self->%s) { cJSON_AddItemToObject(json, %s, cJSON_Duplicate(self->%s, cJSON_True)); }", name, key, name),
		},
		FromJson: []string{
			"{",
			fmt.Sprintf("    const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, %s);", key),
			"    if (v) {",
			fmt.Sprintf("        cJSON_Delete(self->%s);", name),
			fmt.Sprintf("        self->%s = cJSON_Duplicate(v, cJSON_True);", name),
			"    }",
			"}",
		},
	}
}

// cEnumField covers `enum`/`enum?` - a plain int-backed C enum value
// (always defaulting to ordinal 0, its first declared member - see
// c-enum.go), with `_to_string`/`_from_string` doing the wire-value
// mapping.
func cEnumField(name, key, enumType string) cFieldPlan {
	return cFieldPlan{
		Decl: []string{fmt.Sprintf("%s %s;", enumType, name)},
		Init: []string{fmt.Sprintf("self->%s = (%s) 0;", name, enumType)},
		ToJson: []string{
			fmt.Sprintf("cJSON_AddStringToObject(json, %s, %s_to_string(self->%s));", key, enumType, name),
		},
		FromJson: []string{
			"{",
			fmt.Sprintf("    const cJSON *v = cJSON_GetObjectItemCaseSensitive(json, %s);", key),
			fmt.Sprintf("    if (cJSON_IsString(v) && v->valuestring) { self->%s = %s_from_string(v->valuestring); }", name, enumType),
			"}",
		},
	}
}

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
