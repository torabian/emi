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
		core.FieldTypeList, core.FieldTypeListNullable,
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
			// walkCreateFields only ever runs against an *entity's* fields (called
			// from buildCreateFn, called from GoEntityActionsRender), and always on
			// the *original*, pre-ApplyEntityGormTags copy (see the ordering note on
			// GoEntityActionsRender) - so field.Type still reads "array"/"array?"
			// here even though, by the time this generated code actually runs, the
			// real Entity1Entity struct's own field has already been converted to
			// _list/_list? (see go-entity-gorm.go): a real, gorm-native has-many
			// (foreignKey:LinkerId), not the Operation-wrapped emigo.Array shape.
			// There's nothing to reconcile against on a fresh Create (no pre-existing
			// rows), so no explicit emigorm call is needed here at all: gorm's own
			// Create() cascades has-many children automatically, stamping LinkerId in
			// from dto.Id as it goes, purely from this now being a real association.

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

		case core.FieldTypeArray, core.FieldTypeArrayNullable, core.FieldTypeList, core.FieldTypeListNullable:
			// field.Type here is entity.Fields' own type - which ApplyEntityGormTags
			// has, for an entity's array/array? field, already converted to _list/
			// _list? (see go-entity-gorm.go) - but input (the update dto) was built
			// from the *pristine* fields before that conversion (core.Emi.Preprocess
			// runs first - see BuildEntityOptionalDto in preprocess-entities.go), so
			// input.{Field} is always still the portable, Operation-wrapped
			// emigo.ArrayNullable shape regardless of what the entity's own field
			// looks like now. Same code either way.
			//
			// The update dto's array item type is its *own* struct (a plain, portable
			// array? field), not the real entity's child row type, so each item has to
			// be copied field-by-field into a fresh one. UniqueId is the synthetic
			// field the preprocessor prepends to every array item - set means "this is
			// an existing row, patch it", matching how the entity's own child rows
			// work (see emigorm.ReconcileHasMany).
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

// buildUpdateFn renders {Entity}UpdateFn(tx *gorm.DB, uniqueId string, input {Entity}UpdateDto) (*{Entity}Entity, error).
// The parameter is named uniqueId (not id) to match the ":uniqueId" path parameter
// preprocess-entity-actions.go's Update action declares - emi's action codegen derives
// {Action}PathParameter's field name directly from the URL segment, so the two have to
// agree.
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
// against the given uniqueId - all regardless of how deeply the field is nested inside
// object/object? containers, see walkUpdateFields.
func buildUpdateFn(className string, updateDtoClassName string, fields []*core.EmiField) string {
	var scalarChanges, oneResolve, afterUpdate strings.Builder
	walkUpdateFields(fields, "input.", className, &scalarChanges, &oneResolve, &afterUpdate)

	return fmt.Sprintf(`
// %[1]sUpdateFn applies a partial update to the %[1]s row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// %[1]sCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func %[1]sUpdateFn(tx *gorm.DB, uniqueId string, input %[2]s) (*%[1]s, error) {
	var entity %[1]s
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
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
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}
`, className, updateDtoClassName, oneResolve.String(), scalarChanges.String(), afterUpdate.String())
}

// buildGetFn renders {Entity}GetFn(tx *gorm.DB, uniqueId string) (*{Entity}, error) - a
// plain single-row lookup by the public uniqueId (the same identity Update/AwareDelete
// use, never the internal auto-increment id). Named uniqueId, not id, to match the
// ":uniqueId" path parameter preprocess-entity-actions.go's Get action declares.
func buildGetFn(className string) string {
	return fmt.Sprintf(`
// %[1]sGetFn looks up a single %[1]s row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func %[1]sGetFn(tx *gorm.DB, uniqueId string) (*%[1]s, error) {
	var entity %[1]s
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}
`, className)
}

// entityBrowseActionQueryClassName computes the same class name emi's own action
// codegen (lib/golang/go-query-params.go's GoActionQueryParams) generates for the
// Browse action's querystring-bound struct: {ActionName}Query, where ActionName is
// EmiAction.GetName() run through core.ToUpper+core.NormaliseKey - and GetName() itself
// already appends "Action" to the declared name (see core.EmiAction.GetName). For our
// synthesized Browse action (declared name "{entity}Browse"), that's
// "Entity1Browse" -> GetName() -> "Entity1BrowseAction" -> struct "Entity1BrowseActionQuery".
// This has to match that formula exactly, since {Entity}BrowseFn takes this
// already-generated struct as its own qs parameter (see buildBrowseFn) rather than
// building a second, separately-maintained one.
func entityBrowseActionQueryClassName(entity *core.Module3Entity) string {
	return core.ToUpper(entity.Name+"Browse") + "ActionQuery"
}

// buildBrowseFn renders {Entity}BrowseFn(tx *gorm.DB, qs {Entity}BrowseActionQuery,
// scope string, scopeArgs ...interface{}) ([]*{Entity}, *emigo.QueryResultMeta, error).
// Named "Browse" (not "Query") because it's deliberately the simple case - a
// straightforward filtered/sorted/paged list; "query" is reserved for a more advanced
// mechanism later.
//
// qs is exactly the struct emi's own action codegen already generates for reading the
// Browse action's querystring (see entityBrowseActionQueryClassName) - passed straight
// through rather than re-declared as individual parameters, since a caller reading a
// real request's query params already has one of these on hand (see
// {Entity}BrowseActionQueryFromHttp/FromString/FromGin in {Entity}BrowseAction.go).
// scope/scopeArgs never come from qs at all - only ever set by a caller's own handler
// code.
//
// qs.Filter (if set) is transpiled from JSON-logic to SQL via emigorm.ApplyQueryFilter
// (github.com/h22rana/jsonlogic2sql under the hood - mirrors fireback's own
// modules/fireback/JsonQueryTools.go mechanism); scope/scopeArgs (if set) are applied
// separately via emigorm.ApplyQueryScope - a second, independent condition a caller's
// own handler enforces (e.g. workspace isolation), kept apart from qs.Filter so
// end-user input can never widen or bypass it. Both apply *before* Count(), so the
// returned total reflects them but ignores paging; qs.Sort/StartIndex/ItemsPerPage/
// Cursor are only applied afterwards, to the rows actually fetched - see
// emigorm.ApplyQueryCursor/BuildQueryCursor for how the keyset-pagination cursor itself
// works (mirrors fireback's own cursor mechanism, see fireback's CrudCoreActions.go).
func buildBrowseFn(className string, queryClassName string) string {
	return fmt.Sprintf(`
// %[1]sBrowseFn returns %[1]s rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func %[1]sBrowseFn(tx *gorm.DB, qs %[2]s, scope string, scopeArgs ...interface{}) ([]*%[1]s, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&%[1]s{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)

	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}

	var items []*%[1]s
	paged := emigorm.ApplyQueryPage(emigorm.ApplyQueryCursor(emigorm.ApplyQuerySort(filtered, qs.Sort), qs.Cursor), qs.StartIndex, qs.ItemsPerPage)
	if err := paged.Find(&items).Error; err != nil {
		return nil, nil, err
	}

	meta := &emigo.QueryResultMeta{
		TotalItems: total,
		Cursor:     emigorm.BuildQueryCursor(items),
	}

	return items, meta, nil
}
`, className, queryClassName)
}

// entityOptionalDtoClassName computes the same class name core.Emi.Preprocess's
// synthesized optional dto resolves to: entity.Name + "Optional" (see
// entityOptionalDtoName in lib/core/preprocess-entities.go), then
// EmiDto.GetClassName()'s "+Dto" suffix - "entity1" -> "Entity1OptionalDto".
// Deliberately built from entity.Name directly, *not* entity.GetClassName() (which
// would add its own "Entity" suffix and land on "Entity1EntityOptionalDto" instead) -
// has to match preprocess-entities.go's formula exactly, since this lets the actions
// codegen reference the dto's type name in a signature without needing to look it up in
// module.Dto (the dto itself is rendered
// separately, through the plain generic per-dto pipeline in go-public-api.go, since
// it's just an ordinary dto by the time Go codegen sees it).
func entityOptionalDtoClassName(entity *core.Module3Entity) string {
	return core.ToUpper(entity.Name) + "OptionalDto"
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

// GoEntityActionsRender renders {Entity}CreateFn, {Entity}UpdateFn, {Entity}GetFn,
// {Entity}BrowseFn, {Entity}AwareDeletePreviewFn/{Entity}AwareDeleteFn, and a
// {Entity}ActionsSig bundle (mirroring fireback's own {Entity}ActionsSig /
// {Entity}Actions pattern) wiring all of them up by default. Any of them can be turned
// off via entity.Features (see Module3EntityFeatures) - all default to enabled, so an
// entity with no features block gets exactly what's described above.
//
// IMPORTANT: like GoEntityUpdateDtoRender, this must run against the entity's
// *original* Fields, before ApplyEntityGormTags/GoEntityRender mutate them in place.
func GoEntityActionsRender(
	entity *core.Module3Entity,
	ctx core.MicroGenContext,
) (*core.CodeChunkCompiled, error) {

	className := entity.GetClassName()
	createEnabled := entity.Features.CreateEnabled()
	updateEnabled := entity.Features.UpdateEnabled()
	getEnabled := entity.Features.GetEnabled()
	browseEnabled := entity.Features.BrowseEnabled()
	deleteEnabled := entity.Features.DeleteEnabled()

	var buf strings.Builder
	if createEnabled {
		buf.WriteString(buildCreateFn(className, entity.Fields))
	}

	var optionalDtoClassName string
	if updateEnabled {
		// core.Emi.Preprocess only synthesizes this dto when features.update or
		// features.browse is enabled too (see preprocessEntityOptionalDtos in
		// preprocess-entities.go) - the two have to agree, or this would reference a
		// type that was never generated.
		optionalDtoClassName = entityOptionalDtoClassName(entity)
		buf.WriteString(buildUpdateFn(className, optionalDtoClassName, entity.Fields))
	}

	if getEnabled {
		buf.WriteString(buildGetFn(className))
	}

	var browseQueryClassName string
	if browseEnabled {
		browseQueryClassName = entityBrowseActionQueryClassName(entity)
		buf.WriteString(buildBrowseFn(className, browseQueryClassName))
	}

	if deleteEnabled {
		buf.WriteString(buildAwareDeleteFns(className, entity.Fields))
	}

	var sigFields, sigInit strings.Builder
	if createEnabled {
		fmt.Fprintf(&sigFields, "\tCreate func(tx *gorm.DB, dto *%[1]s) (*%[1]s, error)\n", className)
		fmt.Fprintf(&sigInit, "\tCreate: %[1]sCreateFn,\n", className)
	}
	if updateEnabled {
		fmt.Fprintf(&sigFields, "\tUpdate func(tx *gorm.DB, uniqueId string, input %[1]s) (*%[2]s, error)\n", optionalDtoClassName, className)
		fmt.Fprintf(&sigInit, "\tUpdate: %[1]sUpdateFn,\n", className)
	}
	if getEnabled {
		fmt.Fprintf(&sigFields, "\tGet func(tx *gorm.DB, uniqueId string) (*%[1]s, error)\n", className)
		fmt.Fprintf(&sigInit, "\tGet: %[1]sGetFn,\n", className)
	}
	if browseEnabled {
		fmt.Fprintf(&sigFields, "\tBrowse func(tx *gorm.DB, qs %[1]s, scope string, scopeArgs ...interface{}) ([]*%[2]s, *emigo.QueryResultMeta, error)\n", browseQueryClassName, className)
		fmt.Fprintf(&sigInit, "\tBrowse: %[1]sBrowseFn,\n", className)
	}
	if deleteEnabled {
		fmt.Fprintf(&sigFields, "\tAwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*%[1]sAwareDeletePreview, error)\n", className)
		fmt.Fprintf(&sigFields, "\tAwareDelete func(tx *gorm.DB, uniqueIds []string) error\n")
		fmt.Fprintf(&sigInit, "\tAwareDeletePreview: %[1]sAwareDeletePreviewFn,\n", className)
		fmt.Fprintf(&sigInit, "\tAwareDelete: %[1]sAwareDeleteFn,\n", className)
	}

	fmt.Fprintf(&buf, `
// %[1]sActionsSig bundles the actions available for %[1]s. Extend this (and
// %[1]sActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type %[1]sActionsSig struct {
%[2]s}

var %[1]sActions %[1]sActionsSig = %[1]sActionsSig{
%[3]s}
`, className, sigFields.String(), sigInit.String())

	deps := []core.CodeChunkDependency{}
	if createEnabled || updateEnabled || getEnabled || browseEnabled || deleteEnabled {
		deps = append(deps, core.CodeChunkDependency{Location: "gorm.io/gorm"})
	}
	// Only Create/Update actually call into emigorm's reconcile helpers, and only when
	// the entity has an array/collection/one field (at any nesting depth) to reconcile
	// in the first place; Browse always needs it for ApplyQueryFilter/etc, and
	// AwareDelete never does - it works directly off tx.Where/Delete/Association.
	if browseEnabled || ((createEnabled || updateEnabled) && hasRelationField(entity.Fields)) {
		deps = append(deps, core.CodeChunkDependency{Location: "github.com/torabian/emi/emigorm"})
	}
	// Browse's meta return value is a emigo.QueryResultMeta.
	if browseEnabled {
		deps = append(deps, core.CodeChunkDependency{Location: "github.com/torabian/emi/emigo"})
	}
	// Only AwareDelete's preview builds a message via fmt.Sprintf in the generated
	// code itself.
	if deleteEnabled {
		deps = append(deps, core.CodeChunkDependency{Location: "fmt"})
	}

	return &core.CodeChunkCompiled{
		ActualScript:          []byte(buf.String()),
		SuggestedFileName:     className + "Actions",
		SuggestedExtension:    ".go",
		CodeChunkDependensies: deps,
	}, nil
}
