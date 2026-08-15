package formgen

import "github.com/torabian/emi/lib/core"

// EnumOption is one selectable value of a WidgetSelect field.
type EnumOption struct {
	Value string
	Label string
}

// FieldPlan is one EmiField translated into form terms: which widget it
// needs, and - for the group/relation kinds - enough naming information for
// a renderer to emit a sub-component and wire it up.
type FieldPlan struct {
	// Field is the source EmiField this plan was built from. Kept around so a
	// renderer can reach anything this struct doesn't already surface
	// (Description, Tags, Recommended, ...).
	Field *core.EmiField

	Name     string
	Widget   WidgetKind
	Nullable bool

	// EnumOptions is populated for WidgetSelect.
	EnumOptions []EnumOption

	// Primitive is populated for WidgetPrimitiveList (EmiField.Primitive).
	Primitive string

	// MapKeyOf/MapPairOf are populated for WidgetMap.
	MapKeyOf  string
	MapPairOf string

	// Target/Module are populated for WidgetOneRelation/WidgetCollectionRelation
	// (EmiField.Target/Module) - the entity this field points at, which this
	// compiler was never given the fields of.
	Target string
	Module string

	// GroupName is populated for WidgetObjectGroup/WidgetArrayGroup: a
	// PascalCase identifier, unique within the dto, a renderer can use to name
	// the sub-form function/component/type it emits for this field's own
	// Children. For an array group, Children describe a single item of the
	// array, not the array itself.
	GroupName string
	Children  []*FieldPlan
}

// FormPlan is a whole EmiDto translated into form terms.
type FormPlan struct {
	// RootName is a PascalCase identifier for the dto itself (typically
	// EmiDto.GetClassName()), used as the prefix for every GroupName so
	// sibling dtos in the same package never collide.
	RootName string
	Fields   []*FieldPlan
}

// BuildFormPlan walks a dto's fields (and recursively, EmiField.Fields for
// object/array fields) into a FormPlan. rootName should be a PascalCase
// identifier for the dto, e.g. EmiDto.GetClassName().
func BuildFormPlan(rootName string, fields []*core.EmiField) *FormPlan {
	return &FormPlan{
		RootName: rootName,
		Fields:   buildFieldPlans(fields, rootName),
	}
}

func buildFieldPlans(fields []*core.EmiField, groupPrefix string) []*FieldPlan {
	plans := make([]*FieldPlan, 0, len(fields))

	for _, f := range fields {
		if f == nil {
			continue
		}

		kind := ResolveWidgetKind(f)

		plan := &FieldPlan{
			Field:     f,
			Name:      f.Name,
			Widget:    kind,
			Nullable:  IsNullableFieldType(f.Type),
			Primitive: f.Primitive,
			MapKeyOf:  f.MapKeyOf,
			MapPairOf: f.MapPairOf,
			Target:    f.Target,
			Module:    f.Module,
		}

		if kind == WidgetSelect {
			for _, of := range f.OfType {
				if of == nil {
					continue
				}
				// Value is the raw yaml key - the actual value that travels over
				// JSON - not EmiEnumInline.GetKey()'s Go-identifier form (that
				// normalisation, e.g. "admin" -> "Admin", is for generated Go code
				// only; see TsComputedField's own enum handling in lib/js for the
				// same convention on the TypeScript side).
				label := of.Description
				if label == "" {
					label = of.Key
				}
				plan.EnumOptions = append(plan.EnumOptions, EnumOption{
					Value: of.Key,
					Label: label,
				})
			}
		}

		if kind == WidgetObjectGroup || kind == WidgetArrayGroup {
			plan.GroupName = groupPrefix + core.ToUpper(f.Name)
			plan.Children = buildFieldPlans(f.Fields, plan.GroupName)
		}

		plans = append(plans, plan)
	}

	return plans
}
