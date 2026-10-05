package entitysql

import (
	"github.com/torabian/emi/lib/core"
)

// ModuleToPostgres renders every entity of the module into one idempotent postgres sql file.
func ModuleToPostgres(m core.Emi) ([]core.VirtualFile, error) {
	return ModuleToSQL(m, Options{Dialect: Postgres})
}

// ModuleToSQL renders every entity of the module into one sql file, for the database in
// opts.Dialect. The postgres file is <module>.sql, any other <module>.<dialect>.sql, so the files
// of several databases can sit in one folder.
func ModuleToSQL(m core.Emi, opts Options) ([]core.VirtualFile, error) {
	if opts.Dialect == "" {
		opts.Dialect = Postgres
	}
	sql, err := EntitiesToSQL(m.Entities, opts)
	if err != nil {
		return nil, err
	}

	name := m.Name
	if name == "" {
		name = "entities"
	}
	if opts.Dialect != Postgres {
		name += "." + string(opts.Dialect)
	}

	return []core.VirtualFile{{
		Name:         name,
		MimeType:     "application/sql",
		Extension:    ".sql",
		ActualScript: sql,
	}}, nil
}
