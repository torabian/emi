package js

import (
	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/formgen"
)

// DefaultFormComponentsImport is where the primitive form widgets
// (TextField, NumberField, ...) are assumed to live when --form-components
// isn't set. This compiler only ever imports from it - it never generates
// these components itself (see the package doc on react-form-writer.go for
// why).
const DefaultFormComponentsImport = "./form-controls"

// DefaultFormRelationsImport is where one/collection relation picker
// components are assumed to live when --form-relations isn't set.
const DefaultFormRelationsImport = "./relations"

// DefaultReactFormToolset builds the out-of-the-box formgen.Toolset used by
// ReactFormGenerator. componentsImport/relationsImport override the two
// default locations above; pass "" to keep the default for either.
//
// This is the "internal configuration: which component, which props, for
// each field type" the form compiler needs - swap any entry out (or build an
// entirely different formgen.Toolset) to target a different component
// library without touching the writer.
func DefaultReactFormToolset(componentsImport, relationsImport string) formgen.Toolset {
	if componentsImport == "" {
		componentsImport = DefaultFormComponentsImport
	}
	if relationsImport == "" {
		relationsImport = DefaultFormRelationsImport
	}

	toolset := formgen.NewToolset()

	toolset.Set(formgen.WidgetText, formgen.WidgetComponent{
		Component: "TextField", ImportFrom: componentsImport, Named: true,
	})
	toolset.Set(formgen.WidgetNumber, formgen.WidgetComponent{
		Component: "NumberField", ImportFrom: componentsImport, Named: true,
	})
	toolset.Set(formgen.WidgetCheckbox, formgen.WidgetComponent{
		Component: "CheckboxField", ImportFrom: componentsImport, Named: true,
	})
	toolset.Set(formgen.WidgetSelect, formgen.WidgetComponent{
		Component: "SelectField", ImportFrom: componentsImport, Named: true,
	})
	toolset.Set(formgen.WidgetPrimitiveList, formgen.WidgetComponent{
		Component: "ListField", ImportFrom: componentsImport, Named: true,
	})
	toolset.Set(formgen.WidgetMap, formgen.WidgetComponent{
		Component: "MapField", ImportFrom: componentsImport, Named: true,
	})
	toolset.Set(formgen.WidgetJson, formgen.WidgetComponent{
		Component: "JsonField", ImportFrom: componentsImport, Named: true,
	})

	toolset.RelationComponent = func(target string) string {
		return core.ToUpper(target) + "Picker"
	}
	toolset.RelationImportFrom = func(target, module string) string {
		if module != "" {
			return relationsImport + "/" + module + "/" + target
		}
		return relationsImport + "/" + target
	}

	return toolset
}
