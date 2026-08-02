package core

// entityOptionalDtoName returns the *declared* Name for the "every field optional" dto
// Emi.Preprocess synthesizes for an entity - not its Go (or any other language's) class
// name. EmiDto.GetClassName() always appends "Dto" to whatever it's given, so this is
// engineered so that suffix lands on "OptionalDto": entity.Name is already lowerCamel
// (e.g. "entity1"), and GetClassName() upper-cases only the first rune, so "entity1" +
// "Optional" -> GetClassName() -> "Entity1OptionalDto". Deliberately does *not* go
// through entity.GetClassName() (which would add its own "Entity" suffix, landing on
// the noisier "Entity1EntityOptionalDto") - this is the one place that name is built, so
// anything referencing it (lib/golang/go-entity-actions.go's entityOptionalDtoClassName,
// preprocess-entity-actions.go's Update/Browse actions) has to match this formula
// exactly.
func entityOptionalDtoName(entity *Module3Entity) string {
	return entity.Name + "Optional"
}

// nullableFieldType returns the "?" counterpart of a field type. Types with no
// nullable counterpart (any, complex) are returned unchanged - a caller has no
// portable way to say "leave this one alone", so the optional dto just always carries
// them.
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

// shallowCloneField copies a field (not its Tags - a clone destined for a synthesized
// dto isn't tied to whatever persistence tags the source field carries) without
// touching its own Fields slice, letting the caller decide whether/how to recurse.
func shallowCloneField(f *EmiField) *EmiField {
	if f == nil {
		return nil
	}
	clone := *f
	clone.Tags = nil
	return &clone
}

// cloneEntityOptionalField builds one field of an entity's optional dto - every field
// portable (plain EmiField types only, no compiler-specific escape hatches), so any
// compiler (js, kotlin, swift, ...) can generate this dto exactly like any other one
// declared directly in the module, with zero awareness that an entity produced it.
//
//   - plain scalars (string, int, enum, map, slice, ...) become their "?" counterpart,
//     so a caller can tell "not provided" from "explicitly set to the zero value" -
//     even for fields that aren't nullable on the entity itself.
//   - object/object? recurses, uniformly wrapped "?" regardless of the source field's
//     own nullability (the whole point of this dto is that everything is optional).
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
//     Update always carries them; there's no way to say "leave this one untouched".
func cloneEntityOptionalField(field *EmiField) *EmiField {
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
			nested = append(nested, cloneEntityOptionalField(f))
		}
		clone.Fields = nested

	default:
		clone.Type = nullableFieldType(field.Type)
	}

	return clone
}

// BuildEntityOptionalDto derives the "every field optional" dto for an entity, even
// fields that aren't nullable on the entity itself. It's a plain EmiDto - no
// gorm/relation concerns, no table/migration of its own - used two ways:
//
//   - as Update's request body (lib/golang/go-entity-actions.go's buildUpdateFn uses it
//     to build a partial change-set - only fields explicitly set are applied), and
//   - as the item shape inside Browse's response list (see
//     preprocess-entity-actions.go) - a browse/list query may only ever select or
//     return some fields, not all of them, so its items can't be described by the
//     strict, every-field-required entity dto (see BuildEntityDto) the same way a
//     single Create/Get/Update response can.
//
// Carries a synthetic, nullable uniqueId field, same as BuildEntityDto (entity.Fields
// itself never declares one - see PrependEntityDefaultFields) - without it, a Browse
// response item would have no portable way to say which row it actually is.
//
// Must be called with the entity's Fields exactly as declared - i.e. as part of
// Emi.Preprocess(), before any language-specific compiler (which may mutate an
// entity's Fields in place - see ApplyEntityGormTags) ever sees it.
func BuildEntityOptionalDto(entity *Module3Entity) *EmiDto {
	fields := make([]*EmiField, 0, len(entity.Fields)+1)
	fields = append(fields, &EmiField{Name: "uniqueId", Type: FieldTypeStringNullable})
	for _, f := range entity.Fields {
		if f == nil {
			continue
		}
		fields = append(fields, cloneEntityOptionalField(f))
	}

	return &EmiDto{
		Name:        entityOptionalDtoName(entity),
		Description: "Every field of the \"" + entity.Name + "\" entity, but optional - used both as Update's partial input and as Browse's response item shape.",
		Fields:      fields,
	}
}

// preprocessEntityOptionalDtos synthesizes each entity's optional dto (see
// BuildEntityOptionalDto) and appends it to m.Dto, so every compiler picks it up
// exactly like any other declared dto. Idempotent: a second Preprocess() call (or a
// module that happens to already declare a dto by that exact name) is a no-op for that
// entity.
//
// Gated on Update/Browse rather than being unconditional: those are this dto's only
// consumers (see preprocess-entity-actions.go) - an entity with both disabled has
// nothing that would ever reference it, and leaving one lying around in module.Dto
// would make every other compiler (openapi, postman, ...) advertise capabilities that
// don't actually exist.
func (m *Emi) preprocessEntityOptionalDtos() {
	existing := make(map[string]bool, len(m.Dto)+len(m.Entities))
	for _, d := range m.Dto {
		existing[d.Name] = true
	}

	for _, e := range m.Entities {
		if e == nil || e.Name == "" {
			continue
		}
		if !(e.Features.UpdateEnabled() || e.Features.BrowseEnabled()) {
			continue
		}
		name := entityOptionalDtoName(e)
		if existing[name] {
			continue
		}
		m.Dto = append(m.Dto, *BuildEntityOptionalDto(e))
		existing[name] = true
	}
}

// PreprocessEntityOptionalDtos is preprocessEntityOptionalDtos exposed as a
// core.PreprocessHook. core itself never registers it - a compiler backend that
// actually needs entities' optional dtos synthesized (currently just lib/golang, from
// GetGolangPublicActions) registers it explicitly via RegisterPreprocessHook. Since
// registration is global, once one backend does that, every other Preprocess() caller
// (openapi, postman, md, ...) picks the dto up too, without needing any entity-specific
// code of their own.
func PreprocessEntityOptionalDtos(m *Emi) error {
	m.preprocessEntityOptionalDtos()
	return nil
}

// entityDtoName returns the *declared* Name for the plain dto Emi.Preprocess
// synthesizes for an entity's own shape - entity.Name itself, unchanged, so
// EmiDto.GetClassName() lands on "{Entity}Dto" (e.g. "entity2" -> "Entity2Dto"),
// distinct from both the entity's own generated struct ("Entity2Entity") and its
// optional dto ("Entity2OptionalDto" - see entityOptionalDtoName).
func entityDtoName(entity *Module3Entity) string {
	return entity.Name
}

// BuildEntityDto derives a plain dto mirroring entity's own declared fields exactly -
// no optional-wrapping, no relation-type rewriting (contrast BuildEntityOptionalDto,
// which does both) - plus a synthetic uniqueId (entity.Fields itself never carries one;
// that's added later, Go-specifically, by PrependEntityDefaultFields, see
// lib/golang/go-entity-default-fields.go), so it doubles as both the request body for
// Create (a caller may supply its own uniqueId, e.g. a client-generated one - or leave
// it unset and let the database assign one) and the response body for
// Create/Update/Get (see preprocess-entity-actions.go) - a single, complete row, one
// dto instead of a separate ad-hoc field list per action.
func BuildEntityDto(entity *Module3Entity) *EmiDto {
	fields := make([]*EmiField, 0, len(entity.Fields)+1)
	fields = append(fields, &EmiField{Name: "uniqueId", Type: FieldTypeStringNullable})
	for _, f := range entity.Fields {
		if f == nil {
			continue
		}
		fields = append(fields, shallowCloneField(f))
	}

	return &EmiDto{
		Name:        entityDtoName(entity),
		Description: "Plain dto mirroring the \"" + entity.Name + "\" entity's own fields.",
		Fields:      fields,
	}
}

// preprocessEntityDtos synthesizes each entity's plain dto (see BuildEntityDto) and
// appends it to m.Dto. Idempotent, and yields to a dto the module already declares
// under that exact name - same convention as preprocessEntityOptionalDtos.
//
// Gated on Create/Update/Get rather than being unconditional: those are this dto's only
// consumers (see preprocess-entity-actions.go; Browse uses the optional dto instead,
// see BuildEntityOptionalDto) - an entity with all three disabled has nothing that
// would ever reference it.
func (m *Emi) preprocessEntityDtos() {
	existing := make(map[string]bool, len(m.Dto)+len(m.Entities))
	for _, d := range m.Dto {
		existing[d.Name] = true
	}

	for _, e := range m.Entities {
		if e == nil || e.Name == "" {
			continue
		}
		f := e.Features
		if !(f.CreateEnabled() || f.UpdateEnabled() || f.GetEnabled()) {
			continue
		}
		name := entityDtoName(e)
		if existing[name] {
			continue
		}
		m.Dto = append(m.Dto, *BuildEntityDto(e))
		existing[name] = true
	}
}

// PreprocessEntityDtos is preprocessEntityDtos exposed as a core.PreprocessHook. Same
// opt-in convention as PreprocessEntityOptionalDtos - core never registers this itself.
func PreprocessEntityDtos(m *Emi) error {
	m.preprocessEntityDtos()
	return nil
}
