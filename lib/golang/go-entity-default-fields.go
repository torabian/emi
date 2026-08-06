package golang

import (
	"maps"

	"github.com/torabian/emi/lib/core"
)

// EntityDefaultFields are prepended to every entity's Fields, once, before any other
// processing sees them - BuildEntityUpdateInput, GoEntityActionsRender,
// ApplyEntityGormTags, and finally GoCommonStructGenerator all end up treating id/
// uniqueId exactly as if the developer had declared them directly. This mirrors
// fireback's own base entity columns ("defaultgofields" in its GoEntity template),
// trimmed down to just what's actually wired up on this side so far.
//
//   - id is the database-native primary key: a plain auto-incrementing integer. It's
//     never exposed over JSON/CLI (json:"-"/yaml:"-") - it exists purely so joins/
//     foreign keys (LinkerId on array children, the foreignKey/references on one/
//     collection) can use a small integer instead of the UUID below, which keeps
//     indexes smaller and joins faster. See PrependEntityDefaultFields.
//
//   - uniqueId is the public identifier every API/CLI surface, and the reconcile-diff
//     logic in emigorm, actually works with. Its default value is a UUID, generated
//     *natively by Postgres itself* (gen_random_uuid(), from the pgcrypto extension)
//     via a column default, rather than in application code - one less thing to get
//     wrong, and one less roundtrip. The column itself is a plain sized string
//     (varchar), not a native uuid column: applications regularly need to assign their
//     own non-UUID-shaped identifiers here (a fixed sentinel row's id, a slug, a value
//     that doubles as a natural key elsewhere) and a native uuid column would reject
//     every one of those with a type error - the whole point of uniqueId being a
//     plain string default type is that gen_random_uuid() only supplies the *default*,
//     it never constrains what a caller can put there instead. On a non-Postgres
//     database the default clause is simply ignored, or errors if the equivalent
//     function doesn't exist - the column itself still works everywhere, and the
//     application can always still set it explicitly itself if the database default
//     isn't available.
var EntityDefaultFields = []*core.EmiField{
	{
		Name: "id",
		Type: core.FieldTypeInt64,
		Tags: map[string]string{
			"gorm": "primaryKey;autoIncrement",
			"json": "-",
			"yaml": "-",
		},
	},
	{
		Name: "uniqueId",
		Type: core.FieldTypeString,
		Tags: map[string]string{
			"gorm": "type:varchar(100);default:gen_random_uuid();unique",
		},
	},
}

// cloneEntityDefaultFields deep-copies EntityDefaultFields (including each field's Tags
// map) so callers can prepend them to different entities/child structs without sharing
// the same *EmiField pointers - later mutation (ApplyEntityGormTags sets field.Tags
// entries in place all over this package) would otherwise leak across entities.
func cloneEntityDefaultFields() []*core.EmiField {
	cloned := make([]*core.EmiField, len(EntityDefaultFields))
	for i, f := range EntityDefaultFields {
		clone := *f
		tags := make(map[string]string, len(f.Tags))
		maps.Copy(tags, f.Tags)
		clone.Tags = tags
		cloned[i] = &clone
	}
	return cloned
}

// PrependEntityDefaultFields adds a fresh copy of EntityDefaultFields to the front of
// entity.Fields. Must run exactly once per entity, before GoEntityUpdateInputRender,
// GoEntityActionsRender, and GoEntityRender/ApplyEntityGormTags all see it - so every
// one of them has the same, single, consistent view of id/uniqueId as ordinary declared
// fields, rather than each renderer injecting its own version differently.
func PrependEntityDefaultFields(entity *core.Module3Entity) {
	entity.Fields = append(cloneEntityDefaultFields(), entity.Fields...)
}
