package golang

import (
	"github.com/torabian/emi/lib/core"
)

// nullableFieldType returns the "?" counterpart of a field type. Fields with no
// nullable counterpart (any, complex, and anything already nullable) are returned
// unchanged - complex is special-cased separately in cloneFieldAsOptional (turned into
// a pointer instead), any has no available fix yet and is left as a known limitation.
// array/collection/one never reach this: cloneFieldAsOptional intercepts them first.
func nullableFieldType(t core.FieldType) core.FieldType {
	switch t {
	case core.FieldTypeString:
		return core.FieldTypeStringNullable
	case core.FieldTypeBool:
		return core.FieldTypeBoolNullable
	case core.FieldTypeInt:
		return core.FieldTypeIntNullable
	case core.FieldTypeInt32:
		return core.FieldTypeInt32Nullable
	case core.FieldTypeInt64:
		return core.FieldTypeInt64Nullable
	case core.FieldTypeFloat32:
		return core.FieldTypeFloat32Nullable
	case core.FieldTypeFloat64:
		return core.FieldTypeFloat64Nullable
	case core.FieldTypeEnum:
		return core.FieldTypeEnumNullable
	case core.FieldTypeObject:
		return core.FieldTypeObjectNullable
	case core.FieldTypeMap:
		return core.FieldTypeMapNullable
	case core.FieldTypeSlice:
		return core.FieldTypeSliceNullable
	default:
		return t
	}
}

// cloneFieldAsOptional handles one field for BuildEntityUpdateInput, at *any* nesting
// depth (called for the entity's direct fields, and recursively for object/object?
// children). prefix is the class-name prefix an array field's own child struct is
// named with, matching GoCommonStructGenerator's nested-struct naming exactly - the
// entity's class name at the top level, {parentPrefix}{ParentFieldName} once recursed
// into an object.
//
// array/array?, collection/collection?, and one/one? are rewritten into a "complex"
// field whose type is literally the matching emigo.*Nullable[...] wrapper around the
// *real*, already-generated row/target type (Entity1EntityItems, the collection/one
// Target) - not a fresh, separately-optional nested struct. This matters because
// ReconcileHasMany/ReconcileManyToMany persist each item with a plain Save(), which
// writes *every* field of the item: if the update input's array items were their own
// "doubly optional" clone (each sub-field individually Nullable), an item where the
// caller only set one sub-field would silently null out the rest on save. Reusing the
// real row type means a replace/append item has to be given in full, exactly like
// Create - only whether the *field itself* was touched (IsSet()) is optional.
//
// Everything else gets the plain "?" treatment via nullableFieldType, recursing into
// object/object?'s own Fields with prefix extended by the field's own name.
func cloneFieldAsOptional(entity *core.Module3Entity, prefix string, field *core.EmiField) *core.EmiField {
	if field == nil {
		return nil
	}

	relationField := func(complexType string) *core.EmiField {
		return &core.EmiField{
			Name:        field.Name,
			CliName:     field.CliName,
			Description: field.Description,
			Type:        core.FieldTypeComplex,
			Complex:     complexType,
		}
	}

	switch field.Type {
	case core.FieldTypeArray, core.FieldTypeArrayNullable:
		childStruct := prefix + core.ToUpper(field.Name)
		return relationField("emigo.ArrayNullable[" + childStruct + "]")
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		return relationField("emigo.CollectionNullable[" + field.Target + "]")
	case core.FieldTypeOne, core.FieldTypeOneNullable:
		return relationField("emigo.OneNullable[" + field.Target + "]")
	}

	clone := *field
	clone.Type = nullableFieldType(field.Type)

	// complex has no "?" counterpart in the type system; a pointer gives the same
	// "nil means untouched" signal instead.
	if field.Type == core.FieldTypeComplex && len(field.Complex) > 0 && field.Complex[0] != '*' {
		clone.Complex = "*" + field.Complex
	}

	if len(field.Fields) > 0 {
		newPrefix := prefix + core.ToUpper(field.Name)
		nested := make([]*core.EmiField, 0, len(field.Fields))
		for _, f := range field.Fields {
			nested = append(nested, cloneFieldAsOptional(entity, newPrefix, f))
		}
		clone.Fields = nested
	}

	// This clone isn't a gorm-mapped model of its own - it's a plain struct read back
	// by the Update function to build a partial change-set - so none of the base
	// entity's own tags (gorm:"-", gorm:"embedded", ...) apply here.
	clone.Tags = nil

	return &clone
}

// hasRelationField reports whether fields contains an array/collection/one at any
// nesting depth (recursing into object/object? children) - i.e. whether
// BuildEntityUpdateInput's output is guaranteed to reference emigo.*Nullable[...]
// itself (as opposed to only via a plain "?" scalar field, which
// GoCommonStructGenerator's own dependency detection already catches).
func hasRelationField(fields []*core.EmiField) bool {
	for _, f := range fields {
		if f == nil {
			continue
		}
		switch f.Type {
		case core.FieldTypeArray, core.FieldTypeArrayNullable,
			core.FieldTypeCollection, core.FieldTypeCollectionNullable,
			core.FieldTypeOne, core.FieldTypeOneNullable:
			return true
		case core.FieldTypeObject, core.FieldTypeObjectNullable:
			if hasRelationField(f.Fields) {
				return true
			}
		}
	}
	return false
}

// BuildEntityUpdateInput derives a new entity definition, {Name}UpdateInput, from the
// given one: every field is optional, even ones that aren't nullable on the entity
// itself, so a caller can patch a single field without resending every other one. It
// carries no gorm tags/primary key/relations of its own and creates no table or
// migration - it's a plain struct that maps onto the *same* table as the source entity,
// used purely as the Update function's input to build a partial change-set.
//
// id/uniqueId (added by PrependEntityDefaultFields) are deliberately left out entirely -
// they're the row's immutable identity, not something an Update call patches.
//
// IMPORTANT: this must be called with the entity's *original*, not-yet-mutated Fields -
// i.e. before ApplyEntityGormTags/GoEntityRender run on the same *Module3Entity, since
// those mutate entity.Fields (and the individual *EmiField values) in place, adding
// Row-sibling/LinkerId fields that have no place in an update-input struct.
func BuildEntityUpdateInput(entity *core.Module3Entity) *core.Module3Entity {
	fields := make([]*core.EmiField, 0, len(entity.Fields))
	for _, f := range entity.Fields {
		if f != nil && (f.Name == "id" || f.Name == "uniqueId") {
			continue
		}
		fields = append(fields, cloneFieldAsOptional(entity, entity.GetClassName(), f))
	}

	return &core.Module3Entity{
		Name:        entity.Name + "UpdateInput",
		Description: "Update input for the \"" + entity.Name + "\" entity - every field is optional; only fields explicitly set are applied.",
		Fields:      fields,
	}
}

// GoEntityUpdateInputRender renders {Entity}UpdateInput via the plain
// GoCommonStructGenerator - the same generator dtos use - since this struct has no
// gorm/relation concerns of its own (see BuildEntityUpdateInput).
func GoEntityUpdateInputRender(
	entity *core.Module3Entity,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*StructGenerationParticles, error) {

	updateInput := BuildEntityUpdateInput(entity)
	f := GetCommonFlags(ctx)

	particles, err := GoCommonStructGenerator(updateInput.Fields, ctx, GoCommonStructContext{
		RootClassName:       entity.GetClassName() + "UpdateInput",
		RecognizedComplexes: complexes,
		EmiLocation:         f.Emigo,
	})
	if err != nil {
		return nil, err
	}

	// Relation fields are rewritten to a raw "complex" type above (see
	// cloneFieldAsOptional), so GoCommonStructGenerator's own DetectIfEmiGoIsUsed2
	// heuristic (which looks at field.Type, not field.Complex) won't necessarily catch
	// that emigo.*Nullable[...] is used - make sure it's imported whenever that's the
	// case regardless.
	if hasRelationField(entity.Fields) {
		particles.MainClass.CodeChunkDependensies = append(particles.MainClass.CodeChunkDependensies, core.CodeChunkDependency{
			Location: f.Emigo,
		})
	}

	return particles, nil
}
