package golang

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// entityRowFieldName / entityIdFieldName mirror the exact sibling-naming formulas
// ApplyEntityGormTags uses (go-entity-gorm.go) - they're always the *bare* Go field
// name (e.g. "OwnerId", "Items3Row"), never a dotted path, even when the relation field
// is nested inside an object: gorm looks associations up by that bare name in its own
// schema regardless of how deep the field lives inside an embedded chain (verified
// directly against gorm.io/gorm/schema.Parse and a real sqlite run), and embedding
// requires field names to be unique across the whole entity anyway. Only the *value
// access path* (dto.Outer.Inner.Owner, entity.Outer.Inner.OwnerId, ...) needs to track
// nesting - see accessPrefix in walkCreateFields/walkUpdateFields below.
func entityRowFieldName(field *core.EmiField) string {
	return core.ToUpper(field.Name + "Row")
}

func entityIdFieldName(field *core.EmiField) string {
	return core.ToUpper(field.Name + "Id")
}

// isScalarLikeFieldType reports whether a field is a plain column on the entity's own
// row that Create/Update can read/write directly (as opposed to array/collection/one,
// which need the reconcile helpers, or object, which needs flattening/recursing into).
func isScalarLikeFieldType(t core.FieldType) bool {
	switch t {
	case core.FieldTypeArray, core.FieldTypeArrayNullable,
		core.FieldTypeCollection, core.FieldTypeCollectionNullable,
		core.FieldTypeOne, core.FieldTypeOneNullable,
		core.FieldTypeObject, core.FieldTypeObjectNullable:
		return false
	default:
		return true
	}
}

// walkCreateFields recurses through fields - including into object/object? containers,
// at any depth - collecting the one-resolve and post-Create reconcile statements
// buildCreateFn needs.
//
// accessPrefix is the Go value access path to reach these fields at runtime
// ("dto." at the top level, "dto.Outer." once recursed into a plain object field).
// structPrefix is the class-name prefix an array field's own child struct is named
// with, matching GoCommonStructGenerator's nested-struct naming exactly (the entity's
// class name at the top level, {parentPrefix}{ParentFieldName} once recursed).
//
// object? containers need an extra step: the field itself is emigo.Nullable[T], so
// reaching its children means unwrapping it first via Get() (which returns the same
// pointer Nullable[T] stores internally - mutating through it, e.g. writing a nested
// one field's resolved *Id, correctly lands back in dto before Create runs).
//
// KNOWN LIMITATION: this does not recurse into an array field's own item fields - a
// relation declared on an array *item* (as opposed to on the entity or one of its
// object containers) isn't picked up here; each array item is still persisted as a
// single flat unit via Save().
func walkCreateFields(fields []*core.EmiField, accessPrefix string, structPrefix string, oneResolve, afterCreate *strings.Builder) {
	for _, field := range fields {
		if field == nil {
			continue
		}

		goName := core.ToUpper(field.Name)
		accessPath := accessPrefix + goName

		switch field.Type {
		case core.FieldTypeOne, core.FieldTypeOneNullable:
			idPath := accessPrefix + entityIdFieldName(field)
			fmt.Fprintf(oneResolve, `
	if %[1]s.IsSet() {
		var selectorId string
		if %[1]s.Operation == "select" {
			if s, ok := %[1]s.Selector.(string); ok {
				selectorId = s
			}
		}
		var item *%[3]s
		if %[1]s.Operation != "select" {
			item = &%[1]s.Item
		}
		resolvedId, err := emigorm.ReconcileOne(tx, %[1]s.Operation, selectorId, item)
		if err != nil {
			return err
		}
		%[2]s = resolvedId
	}
`, accessPath, idPath, field.Target)

		case core.FieldTypeArray, core.FieldTypeArrayNullable:
			childStruct := structPrefix + goName
			fmt.Fprintf(afterCreate, `
	if %[1]s.IsSet() {
		items := make([]*%[2]s, len(%[1]s.Items))
		for i := range %[1]s.Items {
			items[i] = &%[1]s.Items[i]
		}
		if err := emigorm.ReconcileHasMany(tx, "linker_id", dto.Id, %[1]s.Operation, items); err != nil {
			return err
		}
	}
`, accessPath, childStruct)

		case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
			rowField := entityRowFieldName(field)
			fmt.Fprintf(afterCreate, `
	if %[1]s.IsSet() {
		items := make([]*%[3]s, len(%[1]s.Items))
		for i := range %[1]s.Items {
			items[i] = &%[1]s.Items[i]
		}
		if err := emigorm.ReconcileManyToMany(tx, dto, "%[2]s", %[1]s.Operation, items); err != nil {
			return err
		}
	}
`, accessPath, rowField, field.Target)

		case core.FieldTypeObject:
			walkCreateFields(field.Fields, accessPath+".", structPrefix+goName, oneResolve, afterCreate)

		case core.FieldTypeObjectNullable:
			// object? can't get gorm:"embedded" directly (see ApplyEntityGormTags) -
			// sync the DTO field's content into its hidden {field}Row sibling before
			// Create touches it. Get() returns the same pointer Nullable[T] stores
			// internally, so any nested writes below (relations resolved inside this
			// container) land in the exact same memory gorm reads from at Create time.
			rowField := entityRowFieldName(field)
			fmt.Fprintf(oneResolve, "\n\tif v, ok := %[1]s.Get(); ok && v != nil {\n\t\t%[2]s%[3]s = v\n\t}\n", accessPath, accessPrefix, rowField)

			var subOne, subAfter strings.Builder
			walkCreateFields(field.Fields, "v.", structPrefix+goName, &subOne, &subAfter)
			if subOne.Len() > 0 {
				fmt.Fprintf(oneResolve, "\n\tif v, ok := %s.Get(); ok && v != nil {%s\t}\n", accessPath, subOne.String())
			}
			if subAfter.Len() > 0 {
				fmt.Fprintf(afterCreate, "\n\tif v, ok := %s.Get(); ok && v != nil {%s\t}\n", accessPath, subAfter.String())
			}
		}
	}
}

// buildCreateFn renders {Entity}CreateFn(tx *gorm.DB, dto *{Entity}Entity) (*{Entity}Entity, error).
//
// Order of operations (validated against a real gorm+sqlite run):
//  1. Resolve one/one? fields into their {field}Id FK column *before* the initial
//     Create - belongs-to doesn't depend on the parent's own id.
//  2. tx.Create(dto): id and uniqueId are left at their Go zero values, so gorm omits
//     them from the INSERT and lets the database assign them (autoIncrement for id,
//     the gen_random_uuid() column default for uniqueId - see
//     lib/golang/go-entity-default-fields.go); gorm reads dto.Id back afterwards.
//     array/collection Row siblings are left nil at this point too, so gorm doesn't
//     attempt to cascade-create them itself (which we don't want - we handle those
//     explicitly next, since gorm's own cascade create for an already-existing inline
//     "one" target, or for array/collection children, isn't what the reconcile
//     helpers below assume).
//  3. Reconcile array/array? and collection/collection? fields now that dto.Id is
//     known, via emigorm.ReconcileHasMany / emigorm.ReconcileManyToMany. This step
//     (and the one/one? resolve above) happens no matter how deeply the relation field
//     is nested inside object/object? containers - see walkCreateFields.
func buildCreateFn(className string, fields []*core.EmiField) string {
	var oneResolve, afterCreate strings.Builder
	walkCreateFields(fields, "dto.", className, &oneResolve, &afterCreate)

	return fmt.Sprintf(`
// %[1]sCreateFn creates a new %[1]s row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func %[1]sCreateFn(tx *gorm.DB, dto *%[1]s) (*%[1]s, error) {
	err := tx.Transaction(func(tx *gorm.DB) error {
%[2]s
		if err := tx.Create(dto).Error; err != nil {
			return err
		}
%[3]s
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dto, nil
}
`, className, oneResolve.String(), afterCreate.String())
}

// walkUpdateFields is walkCreateFields's counterpart for buildUpdateFn: besides the
// same one-resolve/reconcile recursion (now writing into changes/afterUpdate instead of
// directly onto dto), it also has to recursively flatten *every* scalar field - at any
// object/object? nesting depth - into the same changes map, since Update's whole point
// is "only touch what the caller actually set" (input.{Field}.IsSet()), which a
// struct-based gorm Update can't do on its own.
//
// changes map keys are always the *bare* Go field name (e.g. "Item1"), never a dotted
// path: gorm's map-based Updates() resolves keys against the model's schema by field
// name, and embedding already requires those names to be unique across the whole
// entity - same reasoning as entityRowFieldName/entityIdFieldName above.
func walkUpdateFields(fields []*core.EmiField, accessPrefix string, structPrefix string, scalarChanges, oneResolve, afterUpdate *strings.Builder) {
	for _, field := range fields {
		if field == nil {
			continue
		}

		// id/uniqueId are the row's immutable identity - BuildEntityUpdateInput
		// leaves them out of the update input entirely, so there's nothing to read
		// here either. (These only ever occur at the true top level in practice -
		// PrependEntityDefaultFields adds them once, to the entity itself.)
		if field.Name == "id" || field.Name == "uniqueId" {
			continue
		}

		goName := core.ToUpper(field.Name)
		accessPath := accessPrefix + goName

		switch field.Type {
		case core.FieldTypeObject, core.FieldTypeObjectNullable:
			// Unlike the real entity struct (where plain "object" is direct and only
			// "object?" is Nullable-wrapped), BuildEntityUpdateInput wraps *every*
			// object field in emigo.Nullable[T] regardless of its original
			// nullability - the whole update input exists to make everything
			// optional. So both cases need the same Get()-unwrap here.
			var subScalar, subOne, subAfter strings.Builder
			walkUpdateFields(field.Fields, "v.", structPrefix+goName, &subScalar, &subOne, &subAfter)
			wrap := func(dst *strings.Builder, body string) {
				if body == "" {
					return
				}
				fmt.Fprintf(dst, "\n\tif %[1]s.IsSet() {\n\t\tif v, ok := %[1]s.Get(); ok && v != nil {%[2]s\t\t}\n\t}\n", accessPath, body)
			}
			wrap(scalarChanges, subScalar.String())
			wrap(oneResolve, subOne.String())
			wrap(afterUpdate, subAfter.String())

		case core.FieldTypeOne, core.FieldTypeOneNullable:
			idField := entityIdFieldName(field)
			fmt.Fprintf(oneResolve, `
	if %[1]s.IsSet() {
		var selectorId string
		if %[1]s.Operation == "select" {
			if s, ok := %[1]s.Selector.(string); ok {
				selectorId = s
			}
		}
		var item *%[3]s
		if %[1]s.Operation != "select" {
			item = &%[1]s.Item
		}
		resolvedId, err := emigorm.ReconcileOne(tx, %[1]s.Operation, selectorId, item)
		if err != nil {
			return err
		}
		changes["%[2]s"] = resolvedId
	}
`, accessPath, idField, field.Target)

		case core.FieldTypeArray, core.FieldTypeArrayNullable:
			// The update dto's array item type is its *own* struct (core.Emi.Preprocess
			// built it as a plain, portable array? field - see
			// lib/core/preprocess-entities.go), not the real entity's child row type,
			// so each item has to be copied field-by-field into a fresh one. UniqueId
			// is the synthetic field the preprocessor prepends to every array item -
			// set means "this is an existing row, patch it", matching how the entity's
			// own child rows work (see emigorm.ReconcileHasMany).
			childStruct := structPrefix + goName
			var copyFields strings.Builder
			fmt.Fprintf(&copyFields, "\t\t\t\tUniqueId: src.UniqueId.OrDefault(\"\"),\n")
			for _, sub := range field.Fields {
				if sub == nil {
					continue
				}
				subGoName := core.ToUpper(sub.Name)
				fmt.Fprintf(&copyFields, "\t\t\t\t%[1]s: src.%[1]s,\n", subGoName)
			}
			fmt.Fprintf(afterUpdate, `
	if %[1]s.IsSet() {
		items := make([]*%[2]s, len(%[1]s.Items))
		for i := range %[1]s.Items {
			src := %[1]s.Items[i]
			items[i] = &%[2]s{
%[3]s			}
		}
		if err := emigorm.ReconcileHasMany(tx, "linker_id", entity.Id, %[1]s.Operation, items); err != nil {
			return err
		}
	}
`, accessPath, childStruct, copyFields.String())

		case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
			// collection/one reference the same real Target type directly (Target was
			// always a portable reference to another entity/dto, unaffected by any of
			// this), so no conversion is needed here the way array needs one.
			rowField := entityRowFieldName(field)
			fmt.Fprintf(afterUpdate, `
	if %[1]s.IsSet() {
		items := make([]*%[3]s, len(%[1]s.Items))
		for i := range %[1]s.Items {
			items[i] = &%[1]s.Items[i]
		}
		if err := emigorm.ReconcileManyToMany(tx, &entity, "%[2]s", %[1]s.Operation, items); err != nil {
			return err
		}
	}
`, accessPath, rowField, field.Target)

		case core.FieldTypeComplex:
			// complex has no portable "?" counterpart (see
			// lib/core/preprocess-entities.go), so it's carried over unchanged - a
			// plain value, same as on the entity itself. There's no way to say
			// "leave this one untouched": Update always includes it.
			fmt.Fprintf(scalarChanges, `
	changes["%[1]s"] = %[2]s
`, goName, accessPath)

		case core.FieldTypeAny:
			// any is a bare interface{} on the update dto - no IsSet(), a plain nil
			// check is the closest thing to an "untouched" signal it has.
			fmt.Fprintf(scalarChanges, `
	if %[2]s != nil {
		changes["%[1]s"] = %[2]s
	}
`, goName, accessPath)

		default:
			if isScalarLikeFieldType(field.Type) {
				fmt.Fprintf(scalarChanges, `
	if %[2]s.IsSet() {
		changes["%[1]s"] = %[2]s
	}
`, goName, accessPath)
			}
		}
	}
}

// buildUpdateFn renders {Entity}UpdateFn(tx *gorm.DB, id string, input {Entity}UpdateDto) (*{Entity}Entity, error).
//
// input's type is the dto core.Emi.Preprocess synthesized for this entity (see
// lib/core/preprocess-entities.go) - a plain, portable dto, not something built here.
//
// Scalars (including map/slice/any/enum, and object/object? fields flattened
// recursively at any nesting depth into the same column set) are only written into the
// update map when input.{Field}.IsSet() - so a field the caller didn't mention is left
// completely untouched, which a struct-based gorm Update can't do (it silently drops
// zero values, verified unreliable for this). one/one? are resolved into their
// {field}Id FK column the same way Create does. array/array? and
// collection/collection? are reconciled via the same emigorm helpers Create uses,
// against the given id - all regardless of how deeply the field is nested inside
// object/object? containers, see walkUpdateFields.
func buildUpdateFn(className string, updateDtoClassName string, fields []*core.EmiField) string {
	var scalarChanges, oneResolve, afterUpdate strings.Builder
	walkUpdateFields(fields, "input.", className, &scalarChanges, &oneResolve, &afterUpdate)

	return fmt.Sprintf(`
// %[1]sUpdateFn applies a partial update to the %[1]s row identified by id (its public
// uniqueId, e.g. from an API path parameter - never the internal auto-increment id).
// Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// %[1]sCreateFn uses, against entity.Id (the row's real primary key, resolved from id
// up front - gorm's Association API and the has-many reconcile both join on it, not on
// uniqueId).
func %[1]sUpdateFn(tx *gorm.DB, id string, input %[2]s) (*%[1]s, error) {
	var entity %[1]s
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", id).Error; err != nil {
			return err
		}

		changes := map[string]interface{}{}
%[3]s
%[4]s
		if len(changes) > 0 {
			if err := tx.Model(&entity).Updates(changes).Error; err != nil {
				return err
			}
		}
%[5]s
		return nil
	})
	if err != nil {
		return nil, err
	}

	var updated %[1]s
	if err := tx.First(&updated, "unique_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}
`, className, updateDtoClassName, oneResolve.String(), scalarChanges.String(), afterUpdate.String())
}

// entityUpdateDtoClassName computes the same class name
// core.Emi.Preprocess's synthesized update dto resolves to (entityUpdateDtoName in
// lib/core/preprocess-entities.go, run through EmiDto.GetClassName()'s "+Dto" suffix) -
// deterministic from entity.GetClassName() alone, so the actions codegen can reference
// the dto's type name in a signature without needing to look it up in module.Dto (the
// dto itself is rendered separately, through the plain generic per-dto pipeline in
// go-public-api.go, since it's just an ordinary dto by the time Go codegen sees it).
func entityUpdateDtoClassName(entity *core.Module3Entity) string {
	return entity.GetClassName() + "UpdateDto"
}

// hasRelationField reports whether fields contains an array/collection/one at any
// nesting depth (recursing into object/object? children).
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

// GoEntityActionsRender renders {Entity}CreateFn, {Entity}UpdateFn, and a
// {Entity}ActionsSig bundle (mirroring fireback's own {Entity}ActionsSig /
// {Entity}Actions pattern) with Create/Update wired to those two functions by default.
//
// IMPORTANT: like GoEntityUpdateDtoRender, this must run against the entity's
// *original* Fields, before ApplyEntityGormTags/GoEntityRender mutate them in place.
func GoEntityActionsRender(
	entity *core.Module3Entity,
	ctx core.MicroGenContext,
) (*core.CodeChunkCompiled, error) {

	className := entity.GetClassName()
	updateDtoClassName := entityUpdateDtoClassName(entity)

	var buf strings.Builder
	buf.WriteString(buildCreateFn(className, entity.Fields))
	buf.WriteString(buildUpdateFn(className, updateDtoClassName, entity.Fields))

	fmt.Fprintf(&buf, `
// %[1]sActionsSig bundles the actions available for %[1]s. Extend this (and
// %[1]sActions below) with more fields as more actions are generated - Create/Update
// are wired to %[1]sCreateFn/%[1]sUpdateFn by default, but callers can swap either out
// (e.g. in tests, or to layer extra validation/side effects around them).
type %[1]sActionsSig struct {
	Create func(tx *gorm.DB, dto *%[1]s) (*%[1]s, error)
	Update func(tx *gorm.DB, id string, input %[2]s) (*%[1]s, error)
}

var %[1]sActions %[1]sActionsSig = %[1]sActionsSig{
	Create: %[1]sCreateFn,
	Update: %[1]sUpdateFn,
}
`, className, updateDtoClassName)

	deps := []core.CodeChunkDependency{
		{Location: "gorm.io/gorm"},
	}
	// Only entities with an array/collection/one field (at any nesting depth) actually
	// call into emigorm - a purely-scalar entity's Create/Update never reference it.
	if hasRelationField(entity.Fields) {
		deps = append(deps, core.CodeChunkDependency{Location: "github.com/torabian/emi/emigorm"})
	}

	return &core.CodeChunkCompiled{
		ActualScript:          []byte(buf.String()),
		SuggestedFileName:     className + "Actions",
		SuggestedExtension:    ".go",
		CodeChunkDependensies: deps,
	}, nil
}
