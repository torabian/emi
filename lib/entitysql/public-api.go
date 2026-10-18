package entitysql

import (
	"errors"

	"github.com/torabian/emi/lib/core"
)

// EntitySqlAction compiles the entities of a module into sql. The database is picked by a tag,
// `sqlite` or `postgres` (the default), from the cli (--tags sqlite) and from a target of
// `emi compile` alike (tags: [sqlite]).
var EntitySqlAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "entity-sql",
		Description:      "Compiles the entities of a module into sql for postgres (default) or sqlite (--tags sqlite): idempotent create table / add column statements",
		WasmFunctionName: "entitySqlGen",
		Flags:            []core.FlagDef{},
	},
	Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
		type_, err := core.DetectEmiStringContentType(ctx.Content)
		if err != nil {
			return nil, err
		}
		if type_ != "module" {
			return nil, errors.New("entity-sql compiles modules containing entities, got type: " + type_)
		}

		dialect, err := DialectFromTags(ctx.Tags)
		if err != nil {
			return nil, err
		}

		m, err := core.StringToEmiWithPath(ctx.Content, ctx.Path)
		if err != nil {
			return nil, err
		}

		return ModuleToSQL(m, Options{Dialect: dialect})
	},
}

func GetEntitySqlPublicActions() core.PublicAPIActions {
	return core.PublicAPIActions{
		FileActions: []core.ActionFile{EntitySqlAction},
	}
}
