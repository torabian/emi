// Package formgen is the shared, renderer-agnostic half of Emi's form
// compiler. It answers one question - "given an EmiDto, what tree of form
// widgets does it describe?" - and nothing else. It never emits source code.
//
// Every concrete form compiler (lib/js's react-form-*.go today, a future
// lib/kotlin jetpack-compose-form-*.go, a future lib/swift swiftui-form-*.go)
// is expected to:
//  1. Call BuildFormPlan(dto) to get a FormPlan.
//  2. Provide its own Toolset (see toolset.go) mapping each WidgetKind to a
//     concrete component/import for that renderer.
//  3. Walk the FormPlan and emit its own source text.
//
// That keeps the "which abstract control does this field need" decision
// (this file) shared across languages, while "what does a TextField look
// like in React vs Compose vs SwiftUI" stays entirely in the per-language
// lib, where it belongs.
package formgen

import (
	"strings"

	"github.com/torabian/emi/lib/core"
)

// WidgetKind is an abstract, renderer-agnostic classification of what kind
// of control a field should be edited with. Renderers translate a WidgetKind
// into their own concrete component via a Toolset - this package never picks
// component names itself.
type WidgetKind string

const (
	// WidgetText edits a single-line string.
	WidgetText WidgetKind = "text"

	// WidgetNumber edits int/int32/int64/float32/float64 and their nullable
	// variants.
	WidgetNumber WidgetKind = "number"

	// WidgetCheckbox edits a bool.
	WidgetCheckbox WidgetKind = "checkbox"

	// WidgetSelect edits an enum field (EmiField.OfType).
	WidgetSelect WidgetKind = "select"

	// WidgetPrimitiveList edits a `slice` field - a flat list of primitives
	// (EmiField.Primitive), as opposed to WidgetArrayGroup which is a list of
	// sub-objects.
	WidgetPrimitiveList WidgetKind = "primitiveList"

	// WidgetMap edits a `map`/`map?` field - an open string/int-keyed bag.
	WidgetMap WidgetKind = "map"

	// WidgetJson is the catch-all for `any` and `complex` fields, whose shape
	// this compiler cannot know ahead of time.
	WidgetJson WidgetKind = "json"

	// WidgetObjectGroup is a nested `object`/`object?` field - rendered as a
	// separate sub-form function/component in the same file (per-field
	// EmiField.Fields).
	WidgetObjectGroup WidgetKind = "objectGroup"

	// WidgetArrayGroup is an `array`/`_list` field - a list of sub-objects,
	// rendered as a separate sub-form function/component that repeats the
	// object-group form for each item.
	WidgetArrayGroup WidgetKind = "arrayGroup"

	// WidgetOneRelation is a `one`/`class` field: a reference to a single row
	// of another entity/dto. The compiler does not know that target's fields,
	// so it renders a picker component imported from elsewhere (see
	// Toolset.RelationComponent/RelationImportFrom).
	WidgetOneRelation WidgetKind = "oneRelation"

	// WidgetCollectionRelation is a `collection` field: a reference to many
	// rows of another entity. Same "import from elsewhere" treatment as
	// WidgetOneRelation.
	WidgetCollectionRelation WidgetKind = "collectionRelation"

	// WidgetUnknown is returned only for a field type this package has not
	// been taught about yet. See widget_test.go's TestAllCatalogFieldTypesResolve
	// - it fails CI the moment core.GetEmiFieldTypeCatalog() grows a type this
	// switch does not cover, so this value should never actually reach a
	// renderer for a real dto.
	WidgetUnknown WidgetKind = "unknown"
)

// IsNullableFieldType reports whether an Emi field type carries the `?`
// nullable suffix (e.g. "string?", "one?").
func IsNullableFieldType(t core.FieldType) bool {
	return strings.HasSuffix(string(t), "?")
}

// ResolveWidgetKind maps a single EmiField's type onto the abstract widget it
// should be edited with. It covers every entry of core.GetEmiFieldTypeCatalog
// (nullable and non-nullable) - see widget_test.go for the coverage check.
func ResolveWidgetKind(field *core.EmiField) WidgetKind {
	if field == nil {
		return WidgetUnknown
	}

	// A complex type takes priority over the base Type, same convention core
	// and lib/js follow (see TsComputedField).
	if field.Complex != "" {
		return WidgetJson
	}

	switch field.Type {
	case core.FieldTypeString, core.FieldTypeStringNullable:
		return WidgetText

	case core.FieldTypeBool, core.FieldTypeBoolNullable:
		return WidgetCheckbox

	case core.FieldTypeInt, core.FieldTypeIntNullable,
		core.FieldTypeInt32, core.FieldTypeInt32Nullable,
		core.FieldTypeInt64, core.FieldTypeInt64Nullable,
		core.FieldTypeFloat32, core.FieldTypeFloat32Nullable,
		core.FieldTypeFloat64, core.FieldTypeFloat64Nullable:
		return WidgetNumber

	case core.FieldTypeEnum, core.FieldTypeEnumNullable:
		return WidgetSelect

	case core.FieldTypeSlice, core.FieldTypeSliceNullable:
		return WidgetPrimitiveList

	case core.FieldTypeMap, core.FieldTypeMapNullable:
		return WidgetMap

	case core.FieldTypeObject, core.FieldTypeObjectNullable:
		return WidgetObjectGroup

	case core.FieldTypeArray, core.FieldTypeArrayNullable,
		core.FieldTypeList, core.FieldTypeListNullable:
		return WidgetArrayGroup

	case core.FieldTypeOne, core.FieldTypeOneNullable,
		core.FieldTypeClass, core.FieldTypeClassNullable:
		return WidgetOneRelation

	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		return WidgetCollectionRelation

	case core.FieldTypeAny, core.FieldTypeComplex:
		return WidgetJson

	default:
		return WidgetUnknown
	}
}
