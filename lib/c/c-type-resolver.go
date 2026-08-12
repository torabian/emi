package c

// extractPrimitive maps a leaf emi primitive type onto its C scalar type.
// Returns "" when the field isn't a primitive, so callers fall through to
// structural resolution (array, object, one, collection...).
//
// C has no native nullable value types without wrapping every scalar in a
// pointer or a tagged union - unlike every other target here, a `?` on
// int/float/bool/enum is not distinguished from its zero value (0/0.0/false
// /first-enum-member) in the generated struct. string/object/array/
// relation fields *are* still properly nullable (they're already pointers),
// which covers the cases that matter most in practice; this is a
// deliberate, documented scope boundary for the C target rather than an
// oversight.
func extractPrimitive(fieldType string) string {
	switch fieldType {
	case "int", "int32", "int32?", "int?":
		return "int32_t"
	case "int64", "int64?":
		return "int64_t"
	case "float32", "float32?":
		return "float"
	case "float64", "float64?":
		return "double"
	case "bool", "bool?":
		return "bool"
	default:
		return ""
	}
}
