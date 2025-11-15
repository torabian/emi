package kotlin

import (
	"encoding/json"
	"fmt"

	"github.com/torabian/emi/lib/core"
)

type DefaultValueHandler interface {
	StringNullable(field *core.EmiField) string
	String(field *core.EmiField) string
	Array(field *core.EmiField) string
	Slice(field *core.EmiField) string
}

func getDefault(field *core.EmiField) string {
	switch v := field.Default.(type) {
	case string:
		return fmt.Sprintf("%q", v)
	case int, int64, float64, bool:
		return fmt.Sprintf("%v", v)
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

func ResolveDefaultField(content DefaultValueHandler, field *core.EmiField) string {
	switch field.Type {

	case core.FieldTypeArray:
		return content.Array(field)
	case core.FieldTypeSlice:
		return content.Slice(field)

	case core.FieldTypeStringNullable:
		return content.StringNullable(field)
	case core.FieldTypeString:
		return content.String(field)
	}

	return ""
}

type KotlinFieldResolver struct{}

func (x KotlinFieldResolver) StringNullable(field *core.EmiField) string {
	s, ok := field.Default.(string)
	if ok && s != "" {
		return fmt.Sprintf("MaybeField(Maybe.Value(\"%v\"))", s)
	}

	return `MaybeField(Maybe.Absent)`
}

func (x KotlinFieldResolver) String(field *core.EmiField) string {
	s, ok := field.Default.(string)
	if ok && s != "" {
		return fmt.Sprintf("\"%v\"", s)
	}

	return `""`
}

func (x KotlinFieldResolver) Array(field *core.EmiField) string {
	return "emptyList()"
}

func (x KotlinFieldResolver) Slice(field *core.EmiField) string {
	return "emptyList()"
}

func KotlinSafeDefaultValue(field *core.EmiField) string {

	m := KotlinFieldResolver{}

	if viaResolver := ResolveDefaultField(m, field); viaResolver != "" {
		return viaResolver
	}

	if field == nil {
		return "null"
	}

	switch field.Type {

	case core.FieldTypeAny:
		return ""

	case core.FieldTypeString:
		if field.Default != "" {
			return `"` + getDefault(field) + `"`
		}

		return `""`
	}

	if field.Default != nil {
		switch v := field.Default.(type) {
		case string:
			return fmt.Sprintf("%q", v)
		case int, int64, float64, bool:
			return fmt.Sprintf("%v", v)
		default:
			b, _ := json.Marshal(v)
			return string(b)
		}
	}

	switch field.Type {
	case "array", "slice", "collection":
		return "emptyList()"
	case "object?":
		return "MaybeField(Maybe.Absent)"
	case "string", "text":
		return `""`
	case "float", "float32": //"float?", "float32?", "float64?":
		return "0.0f"
	case core.FieldTypeFloat64:
		return "0.0"
	case "int", "int32", "int64": // "int?", "int32?", "int64?":
		return "0"
	default:
		return ""
	}
}
