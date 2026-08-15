package js

import "github.com/torabian/emi/lib/formgen"

// DefaultRjsfFieldsImport is where the two custom react-jsonschema-form
// `fields` this compiler always references are assumed to live when
// --rjsf-fields isn't set. Unlike react-form-toolset.go's components, RJSF
// only needs custom code for what it cannot express as a plain JSON Schema
// type - everything else (text, number, checkbox, select, nested object,
// array of objects, even the map field via `additionalProperties`) already
// has a built-in RJSF widget, so this compiler never has to import one.
const DefaultRjsfFieldsImport = "./rjsf-fields"

// RjsfToolset is react-jsonschema-form's much smaller version of
// formgen.Toolset: RJSF's `fields`/`widgets` props are just string-keyed
// registries the *host app* fills in, so all this compiler needs to decide
// is which registry key ("ui:field"/"ui:widget" name) to reference per
// WidgetKind - never an import path or props shape.
type RjsfToolset struct {
	// Widgets optionally overrides RJSF's own default widget for a WidgetKind
	// via "ui:widget" (e.g. force a long string field to "textarea"). Leave a
	// kind unset to let RJSF pick its own default widget for that JSON Schema
	// type - which is enough for Text/Number/Checkbox/Select/PrimitiveList/Map
	// out of the box.
	Widgets map[formgen.WidgetKind]string

	// RelationFieldName is the "ui:field" registry key used for both
	// WidgetOneRelation and WidgetCollectionRelation - the host app registers
	// one React field component under this name (see the generated uiSchema's
	// "ui:options": {target, module, multiple} for how it tells that
	// component which entity/module to point its picker at).
	RelationFieldName string

	// JsonFieldName is the "ui:field" registry key used for WidgetJson
	// (any/complex fields), which have no fixed shape a plain JSON Schema
	// type can express.
	JsonFieldName string
}

// DefaultRjsfToolset returns the toolset ReactJsonSchemaFormGenerator uses
// when nothing overrides it: no ui:widget overrides (trust RJSF's built-ins),
// and the two custom field names the generated wrapper component imports
// from DefaultRjsfFieldsImport (or --rjsf-fields).
func DefaultRjsfToolset() RjsfToolset {
	return RjsfToolset{
		Widgets:           map[formgen.WidgetKind]string{},
		RelationFieldName: "emiRelationField",
		JsonFieldName:     "emiJsonField",
	}
}
