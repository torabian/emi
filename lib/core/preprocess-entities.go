package core

// entityUpdateDtoName returns the *declared* Name for the update dto
// Emi.Preprocess synthesizes for an entity - not its Go (or any other language's)
// class name. EmiDto.GetClassName() always appends "Dto" to whatever it's given, so
// this is engineered so that suffix lands on "UpdateDto": entity.Name is already
// lowerCamel (e.g. "entity1"), and GetClassName() upper-cases only the first rune, so
// "entity1" + "EntityUpdate" -> GetClassName() -> "Entity1EntityUpdateDto".
func entityUpdateDtoName(entity *Module3Entity) string {
	return entity.Name + "EntityUpdate"
}

// nullableFieldType returns the "?" counterpart of a field type. Types with no
// nullable counterpart (any, complex) are returned unchanged - a caller has no
// portable way to say "leave this one alone", so update dtos just always carry them.
func nullableFieldType(t FieldType) FieldType {
	switch t {
	case FieldTypeString:
		return FieldTypeStringNullable
	case FieldTypeBool:
		return FieldTypeBoolNullable
	case FieldTypeInt:
		return FieldTypeIntNullable
	case FieldTypeInt32:
		return FieldTypeInt32Nullable
	case FieldTypeInt64:
		return FieldTypeInt64Nullable
	case FieldTypeFloat32:
		return FieldTypeFloat32Nullable
	case FieldTypeFloat64:
		return FieldTypeFloat64Nullable
	case FieldTypeEnum:
		return FieldTypeEnumNullable
	case FieldTypeObject:
		return FieldTypeObjectNullable
	case FieldTypeArray:
		return FieldTypeArrayNullable
	case FieldTypeOne:
		return FieldTypeOneNullable
	case FieldTypeCollection:
		return FieldTypeCollectionNullable
	case FieldTypeMap:
		return FieldTypeMapNullable
	case FieldTypeSlice:
		return FieldTypeSliceNullable
	default:
		return t
	}
}

// shallowCloneField copies a field (not its Tags - a clone destined for the update dto
// isn't tied to whatever persistence tags the source field carries) without touching
// its own Fields slice, letting the caller decide whether/how to recurse.
func shallowCloneField(f *EmiField) *EmiField {
	if f == nil {
		return nil
	}
	clone := *f
	clone.Tags = nil
	return &clone
}

// cloneEntityUpdateField builds one field of an entity's update dto - every field
// portable (plain EmiField types only, no compiler-specific escape hatches), so any
// compiler (js, kotlin, swift, ...) can generate this dto exactly like any other one
// declared directly in the module, with zero awareness that an entity produced it.
//
//   - plain scalars (string, int, enum, map, slice, ...) become their "?" counterpart,
//     so a caller can tell "not provided" from "explicitly set to the zero value" -
//     even for fields that aren't nullable on the entity itself.
//   - object/object? recurses, uniformly wrapped "?" regardless of the source field's
//     own nullability (the whole point of an update dto is that everything is
//     optional).
//   - array/array? becomes array?, keeping the item's own declared fields exactly as
//     they are (an item given in a replace/append batch still has to be given in
//     full - see lib/golang/go-entity-actions.go's reconcile helpers, which persist
//     each item as a single unit) plus a synthetic, portable uniqueId field prepended
//     so a caller can reference an *existing* item to patch it, the same way the
//     entity's own array children get an id the Go/gorm layer can key off of (see
//     lib/golang/go-entity-default-fields.go) - just meaningful to every compiler, not
//     only Go's.
//   - collection/collection? and one/one? become their "?" counterpart; Target is
//     already a portable reference to another entity/dto, unaffected by any of this.
//   - complex/any have no portable "?" form, so they're left exactly as declared -
//     Update always carries them; there's no away to say "leave this one untouched".
func cloneEntityUpdateField(field *EmiField) *EmiField {
	if field == nil {
		return nil
	}

	clone := shallowCloneField(field)

	switch field.Type {
	case FieldTypeArray, FieldTypeArrayNullable:
		clone.Type = FieldTypeArrayNullable
		nested := make([]*EmiField, 0, len(field.Fields)+1)
		nested = append(nested, &EmiField{Name: "uniqueId", Type: FieldTypeStringNullable})
		for _, f := range field.Fields {
			nested = append(nested, shallowCloneField(f))
		}
		clone.Fields = nested

	case FieldTypeCollection, FieldTypeCollectionNullable:
		clone.Type = FieldTypeCollectionNullable
		clone.Fields = nil

	case FieldTypeOne, FieldTypeOneNullable:
		clone.Type = FieldTypeOneNullable
		clone.Fields = nil

	case FieldTypeObject, FieldTypeObjectNullable:
		clone.Type = FieldTypeObjectNullable
		nested := make([]*EmiField, 0, len(field.Fields))
		for _, f := range field.Fields {
			nested = append(nested, cloneEntityUpdateField(f))
		}
		clone.Fields = nested

	default:
		clone.Type = nullableFieldType(field.Type)
	}

	return clone
}

// BuildEntityUpdateDto derives the update dto for an entity: every field optional,
// even ones that aren't nullable on the entity itself, so a caller can patch a single
// field without resending every other one. It's a plain EmiDto - no gorm/relation
// concerns, no table/migration of its own - that maps onto the *same* table as the
// source entity once the Go-specific layer (lib/golang/go-entity-actions.go) uses it to
// build a partial change-set.
//
// Must be called with the entity's Fields exactly as declared - i.e. as part of
// Emi.Preprocess(), before any language-specific compiler (which may mutate an
// entity's Fields in place - see ApplyEntityGormTags) ever sees it.
func BuildEntityUpdateDto(entity *Module3Entity) *EmiDto {
	fields := make([]*EmiField, 0, len(entity.Fields))
	for _, f := range entity.Fields {
		if f == nil {
			continue
		}
		fields = append(fields, cloneEntityUpdateField(f))
	}

	return &EmiDto{
		Name:        entityUpdateDtoName(entity),
		Description: "Update input for the \"" + entity.Name + "\" entity - every field is optional; only fields explicitly set are applied.",
		Fields:      fields,
	}
}

// preprocessEntityUpdateDtos synthesizes each entity's update dto (see
// BuildEntityUpdateDto) and appends it to m.Dto, so every compiler picks it up exactly
// like any other declared dto. Idempotent: a second Preprocess() call (or a module that
// happens to already declare a dto by that exact name) is a no-op for that entity.
func (m *Emi) preprocessEntityUpdateDtos() {
	existing := make(map[string]bool, len(m.Dto)+len(m.Entities))
	for _, d := range m.Dto {
		existing[d.Name] = true
	}

	for _, e := range m.Entities {
		if e == nil || e.Name == "" {
			continue
		}
		name := entityUpdateDtoName(e)
		if existing[name] {
			continue
		}
		m.Dto = append(m.Dto, *BuildEntityUpdateDto(e))
		existing[name] = true
	}
}
