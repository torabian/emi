package golang

import (
	"github.com/torabian/emi/lib/core"
)

// ApplyEntityGormTags prepares an entity's fields so the struct GoCommonStructGenerator
// renders is gorm-relation-ready, and can actually be handed to gorm.AutoMigrate.
//
// It assumes PrependEntityDefaultFields has already added id (int64, the real gorm
// primary key) and uniqueId (the public identifier) to entity.Fields - see that
// function's comment for why joins/foreign keys use id rather than uniqueId.
//
// The DTO-style relation fields (emigo.One/Collection/Array) stay exactly as they are -
// they're the API/CLI shape, used for PATCH-style request bodies - but they get
// gorm:"-", since gorm cannot map them (they don't implement driver.Valuer/sql.Scanner
// and aren't a slice-of-struct/pointer-to-struct shape gorm recognizes as an
// association; verified empirically via gorm.io/gorm/schema.Parse, which errors on them
// otherwise). Real gorm-shaped sibling fields are injected alongside them:
//
//   - one/one?   -> {field}Id int64 (the FK column, referencing the target's id) +
//     {field}Row *Target, tagged foreignKey:{field}Id;references:Id (a real "belongs
//     to"). emigorm.ReconcileOne resolves the target's id from whatever identity the
//     DTO field's Selector/inline value actually carries (its own uniqueId).
//   - collection/collection? -> {field}Row []*Target, tagged
//     many2many:{entity}_{field};foreignKey:Id;references:Id.
//   - array/array? -> {field}Row []*{ChildStruct} (the same auto-generated nested
//     struct as the DTO field, e.g. Entity1EntityItems), tagged
//     foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE (a real "has
//     many"). The child struct itself gets the same id/uniqueId pair every entity gets
//     (so it has a real identity of its own - e.g. for grandchildren later), plus
//     LinkerId int64 (the foreign key back to this entity's id).
//
// {field}Id/{field}Row are marked json:"-" yaml:"-": they're purely a persistence-layer
// concern, so the existing DTO-shaped field (owner, items3, ...) remains the only public
// JSON/YAML surface for that relation.
//
// object gets gorm:"embedded" directly (inlined into this entity's own row) -
// including when nested inside another object any number of levels deep: gorm's own
// embedding correctly walks a relation field through the chain and still recognizes it
// as a real association, keyed by its bare field name (e.g. "ItemsRow", not
// "Outer.Inner.ItemsRow") - verified directly against gorm.io/gorm/schema.Parse and a
// real sqlite run. What ApplyEntityGormTags has to get right for this to work is naming
// an array field's child struct with the *same* prefix chain
// GoCommonStructGenerator's own nested-struct naming uses (Entity1EntityOuterInnerItems,
// not the flat Entity1EntityItems) - that's what the childStructPrefix parameter is for.
//
// object? can't use the same direct gorm:"embedded" trick - see the FieldTypeObjectNullable
// case below for why, and how it's fixed the same way the relation fields are (a hidden,
// gorm-visible sibling field).
//
// map/slice (non-nullable only) get gorm:"serializer:json", since they render as a
// plain Go map/slice with no Scan/Value of their own; any gets the same treatment.
// map?/slice? are left alone - they render wrapped in emigo.Nullable[T], which already
// implements Value()/Scan() with its own JSON fallback, so adding serializer:json on top
// would make gorm bypass that pair and serialize the wrapper's internal fields instead.
// enum/enum? need no tag: they already resolve to plain string/emigo.Nullable[string].
//
// A "gorm" tag the developer already set on a field (tags: { gorm: ... }) always wins
// and is left untouched.
func ApplyEntityGormTags(entity *core.Module3Entity) {
	entity.Fields = append(entity.Fields, applyEntityGormTags(entity, entity.GetClassName(), entity.Fields)...)
}

// hiddenSibling builds a synthetic field that only exists for gorm's benefit: it never
// shows up in JSON/YAML, and carries the given gorm tag unconditionally (there's no
// developer-authored field to respect an override on, since it's generated on the fly).
func hiddenSibling(name string, fieldType core.FieldType, complex string, gormTag string) *core.EmiField {
	return &core.EmiField{
		Name:    name,
		Type:    fieldType,
		Complex: complex,
		Tags: map[string]string{
			"gorm": gormTag,
			"json": "-",
			"yaml": "-",
		},
	}
}

// childStructPrefix is the class-name prefix an array field's own child struct is
// named with (matching GoCommonStructGenerator's nested-struct naming exactly): the
// entity's class name at the top level, or {parentPrefix}{ParentFieldName} once
// recursed into an object/object? field.
func applyEntityGormTags(entity *core.Module3Entity, childStructPrefix string, fields []*core.EmiField) []*core.EmiField {
	var extra []*core.EmiField

	for _, field := range fields {
		if field == nil {
			continue
		}

		if field.Tags == nil {
			field.Tags = map[string]string{}
		}

		hasOverride := field.Tags["gorm"] != ""

		switch field.Type {
		case core.FieldTypeArray, core.FieldTypeArrayNullable:
			// Converted to _list/_list? - only here, only for the entity's own
			// persisted struct (see core.FieldTypeList's doc comment in
			// EmiFieldType.go) - so the primary field itself becomes a real,
			// gorm-native has-many: no more gorm:"-", no hidden {field}Row shadow
			// sibling; gorm's own reflection-based schema builder can finally see it
			// directly, since a plain slice is exactly the shape it requires.
			//
			// DTOs (BuildEntityDto/BuildEntityOptionalDto in preprocess-entities.go)
			// never see this conversion - they're built from entity.Fields *before*
			// ApplyEntityGormTags ever mutates it (core.Emi.Preprocess runs first),
			// so Create/Update's own request bodies keep the portable,
			// Operation-wrapped array/array? shape they still need.
			childStruct := childStructPrefix + core.ToUpper(field.Name)

			field.Fields = append(field.Fields, cloneEntityDefaultFields()...)
			field.Fields = append(field.Fields, &core.EmiField{
				Name: "linkerId",
				Type: core.FieldTypeInt64,
				Tags: map[string]string{"gorm": "index"},
			})

			field.Target = childStruct
			if field.Type == core.FieldTypeArray {
				field.Type = core.FieldTypeList
			} else {
				field.Type = core.FieldTypeListNullable
			}

			if !hasOverride {
				field.Tags["gorm"] = "foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE"
			}

		case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
			if !hasOverride {
				field.Tags["gorm"] = "-"
			}

			extra = append(extra, hiddenSibling(
				field.Name+"Row",
				core.FieldTypeComplex,
				"[]*"+field.Target,
				"many2many:"+entity.Name+"_"+field.Name+";foreignKey:Id;references:Id",
			))

		case core.FieldTypeObject:
			if !hasOverride {
				field.Tags["gorm"] = "embedded"
			}

		case core.FieldTypeObjectNullable:
			// Unlike "object", this can't just get gorm:"embedded" directly: the DTO
			// field renders as emigo.Nullable[T], and gorm's embedding only ever sees
			// the *wrapper's own* fields (value, isSet - both unexported), never T's -
			// so none of T's columns (or any relation nested inside it) are visible to
			// gorm at all (verified directly: a column that should exist from this
			// field's own nested fields doesn't show up in gorm.io/gorm/schema.Parse).
			// Same fix as the relation fields above: leave the DTO field alone but
			// gorm:"-", and add a hidden *ChildStruct sibling gorm can actually see
			// through (a plain nilable pointer is a well-supported gorm embedding
			// shape). buildCreateFn syncs {field}Row from the DTO field's Get() before
			// Create; Update's existing scalar-flattening already writes each nested
			// field by its own bare column name, unaffected by which Go field
			// physically holds it.
			if !hasOverride {
				field.Tags["gorm"] = "-"
			}

			childStruct := childStructPrefix + core.ToUpper(field.Name)
			extra = append(extra, hiddenSibling(
				field.Name+"Row",
				core.FieldTypeComplex,
				"*"+childStruct,
				"embedded",
			))

		case core.FieldTypeOne, core.FieldTypeOneNullable:
			if !hasOverride {
				field.Tags["gorm"] = "-"
			}

			idField := field.Name + "Id"
			extra = append(extra, hiddenSibling(idField, core.FieldTypeInt64, "", "index"))
			extra = append(extra, hiddenSibling(
				field.Name+"Row",
				core.FieldTypeComplex,
				"*"+field.Target,
				"foreignKey:"+core.ToUpper(idField)+";references:Id",
			))

		case core.FieldTypeMap, core.FieldTypeSlice, core.FieldTypeAny:
			if !hasOverride {
				field.Tags["gorm"] = "serializer:json"
			}
		}

		if len(field.Fields) > 0 {
			newPrefix := childStructPrefix + core.ToUpper(field.Name)
			field.Fields = append(field.Fields, applyEntityGormTags(entity, newPrefix, field.Fields)...)
		}
	}

	return extra
}
