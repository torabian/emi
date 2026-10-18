// Renders the entities of a module as C for SQLite (`--tags sqlite`): the
// struct (the same one a dto gets, with to_json / from_json), the layout of its
// tables as a const table, and typed insert / get / update / delete / list
// functions on top of the emisql.h runtime. The tables and the layout come from
// lib/entitysql, so the code that reads and writes the rows and the sql that
// creates them can not disagree.
package c

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/entitysql"
)

func cIdent(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '_' {
			sb.WriteRune(r)
		} else {
			sb.WriteRune('_')
		}
	}
	return sb.String()
}

func cKind(k entitysql.Kind) string {
	switch k {
	case entitysql.KindInt:
		return "EMISQL_INT"
	case entitysql.KindFloat:
		return "EMISQL_FLOAT"
	case entitysql.KindBool:
		return "EMISQL_BOOL"
	case entitysql.KindJSON:
		return "EMISQL_JSON"
	case entitysql.KindTime:
		return "EMISQL_TIME"
	case entitysql.KindDate:
		return "EMISQL_DATE"
	}
	return "EMISQL_STRING"
}

func layoutVar(l *entitysql.Layout) string { return "emisql_layout_" + cIdent(l.Table) }

// cPathArray declares a const array holding a json path and returns its name.
func cPathArray(sb *strings.Builder, name string, path []string) {
	items := make([]string, len(path))
	for i, p := range path {
		items[i] = escapeDoubleQuoted(p)
	}
	fmt.Fprintf(sb, "static const char* const %s[] = {%s};\n", name, strings.Join(items, ", "))
}

// cLayoutTables writes the const tables of a layout, children first (a table
// can only point at what is declared above it).
func cLayoutTables(sb *strings.Builder, l *entitysql.Layout, done map[string]bool) {
	if done[l.Table] {
		return
	}
	done[l.Table] = true
	for _, c := range l.Children {
		cLayoutTables(sb, c.Layout, done)
	}
	v := layoutVar(l)

	for i, col := range l.Columns {
		cPathArray(sb, fmt.Sprintf("%s_p%d", v, i), col.Path)
	}
	for i, c := range l.Children {
		cPathArray(sb, fmt.Sprintf("%s_c%d_path", v, i), c.Path)
	}

	if len(l.Columns) > 0 {
		fmt.Fprintf(sb, "static const EmiSqlColumn %s_columns[] = {\n", v)
		for i, col := range l.Columns {
			fmt.Fprintf(sb, "    {%s, %s, %s_p%d, %d},\n", escapeDoubleQuoted(col.Column), cKind(col.Kind), v, i, len(col.Path))
		}
		sb.WriteString("};\n")
	}
	if len(l.Children) > 0 {
		fmt.Fprintf(sb, "static const EmiSqlChild %s_children[] = {\n", v)
		for i, c := range l.Children {
			fmt.Fprintf(sb, "    {%s_c%d_path, %d, &%s},\n", v, i, len(c.Path), layoutVar(c.Layout))
		}
		sb.WriteString("};\n")
	}

	cols, ncols := "NULL", "0"
	if len(l.Columns) > 0 {
		cols, ncols = v+"_columns", fmt.Sprint(len(l.Columns))
	}
	kids, nkids := "NULL", "0"
	if len(l.Children) > 0 {
		kids, nkids = v+"_children", fmt.Sprint(len(l.Children))
	}
	fmt.Fprintf(sb, "static const EmiSqlLayout %s = {%s, %s, %s, %s, %s};\n\n", v, escapeDoubleQuoted(l.Table), cols, ncols, kids, nkids)
}

// cEntityLayoutTmpl is what both storages share: the layout tables of the entity.
const cEntityLayoutTmpl = `
/* ---- layout of table %[1]s ---- */
%[2]s

/* Several entities with their ids, whatever stored them (SQLite, a file). Free with %[3]sList_free. */
typedef struct %[3]sList {
    size_t count;
    int64_t* ids;
    %[3]s** items;
    int64_t total; /* rows matching the filter without paging (a query), the rows listed (a list) */
} %[3]sList;

static inline void %[3]sList_free(%[3]sList* list) {
    if (!list) return;
    for (size_t i = 0; i < list->count; i++) %[3]s_free(list->items[i]);
    free(list->items);
    free(list->ids);
    free(list);
}
`

// cEntityPackTmpl is the binary form (--tags pack), see emipack.h.
const cEntityPackTmpl = `
/* ---- compact binary form (see emipack.h) ---- */

/* Packs the entity into a new buffer (*out, free() it) of *len bytes. Returns 0, non-zero on failure.
 * Only the values are written, no field names: append new fields at the end of the entity, never
 * reorder or remove stored ones. */
static inline int %[1]s_pack(const %[1]s* self, uint8_t** out, size_t* len) {
    cJSON* json = %[1]s_to_json(self);
    if (!json) return -1;
    int rc = emipack_encode(&%[2]s, json, out, len);
    cJSON_Delete(json);
    return rc;
}

/* The entity a packed buffer holds (free it with %[1]s_free); NULL if the buffer is damaged or was
 * packed by a newer version of the entity. A buffer from an older version reads fine: the fields
 * it does not have keep their defaults. */
static inline %[1]s* %[1]s_unpack(const uint8_t* data, size_t len) {
    cJSON* json = emipack_decode(&%[2]s, data, len);
    if (!json) return NULL;
    %[1]s* self = %[1]s_from_json(json);
    cJSON_Delete(json);
    return self;
}

/* Like %[1]s_unpack for a buffer which continues after the entity (several packed entities in a row):
 * *used is how many bytes this one took. */
static inline %[1]s* %[1]s_unpack_prefix(const uint8_t* data, size_t len, size_t* used) {
    cJSON* json = emipack_decode_prefix(&%[2]s, data, len, used);
    if (!json) return NULL;
    %[1]s* self = %[1]s_from_json(json);
    cJSON_Delete(json);
    return self;
}
`

const cEntityStoreTmpl = `
/* ---- storage: table %[1]s ----
 * Compile with -DEMI_NO_SQLITE where sqlite3.h is not available (a file with only the binary form). */
#ifndef EMI_NO_SQLITE
#include <sqlite3.h>
#include "emisql.h"

/* Stores the entity; its new id goes to *out_id (may be NULL). Returns 0 or an SQLITE_* code. */
static inline int %[3]s_insert(sqlite3* db, const %[3]s* self, int64_t* out_id) {
    cJSON* json = %[3]s_to_json(self);
    if (!json) return SQLITE_NOMEM;
    int rc = emisql_insert(db, &%[4]s, json, out_id);
    cJSON_Delete(json);
    return rc;
}

/* The entity with this id (free it with %[3]s_free), NULL if there is none. */
static inline %[3]s* %[3]s_get(sqlite3* db, int64_t id) {
    cJSON* json = emisql_select(db, &%[4]s, id);
    if (!json) return NULL;
    %[3]s* self = %[3]s_from_json(json);
    cJSON_Delete(json);
    return self;
}

/* Replaces the entity with this id. SQLITE_NOTFOUND if there is no such id. */
static inline int %[3]s_update(sqlite3* db, int64_t id, const %[3]s* self) {
    cJSON* json = %[3]s_to_json(self);
    if (!json) return SQLITE_NOMEM;
    int rc = emisql_update(db, &%[4]s, id, json);
    cJSON_Delete(json);
    return rc;
}

/* Deletes the entity and the rows of its arrays. SQLITE_NOTFOUND if there is no such id. */
static inline int %[3]s_delete(sqlite3* db, int64_t id) {
    return emisql_delete(db, &%[4]s, id);
}

/* The entities matching q (NULL: all), filtered, sorted and paged as EmiSqlQuery says - see emisql.h.
 * list->total is how many match without the paging. Free with %[3]sList_free.
 * NULL on error, and also for a filter/sort which is not valid or names an unknown field. */
static inline %[3]sList* %[3]s_query(sqlite3* db, const EmiSqlQuery* q) {
    %[3]sList* list = (%[3]sList*) calloc(1, sizeof(*list));
    if (!list) return NULL;
    int64_t* ids = NULL;
    size_t count = 0;
    if (emisql_query_ids(db, &%[4]s, q, &ids, &count, &list->total) != 0) { free(list); return NULL; }
    list->ids = ids;
    list->items = (%[3]s**) calloc(count ? count : 1, sizeof(*list->items));
    if (!list->items) { %[3]sList_free(list); return NULL; }
    for (size_t i = 0; i < count; i++) {
        list->items[i] = %[3]s_get(db, ids[i]);
        list->count = i + 1;
        if (!list->items[i]) { %[3]sList_free(list); return NULL; }
    }
    return list;
}

/* The entities in id order; limit < 0 lists all. Free with %[3]sList_free. NULL on error. */
static inline %[3]sList* %[3]s_list(sqlite3* db, int64_t limit, int64_t offset) {
    EmiSqlQuery q = {NULL, NULL, offset, limit < 0 ? 0 : limit};
    if (limit == 0) { /* nothing asked for: an empty list, as before */
        %[3]sList* none = (%[3]sList*) calloc(1, sizeof(*none));
        if (none) none->items = (%[3]s**) calloc(1, sizeof(*none->items));
        if (none && !none->items) { free(none); return NULL; }
        return none;
    }
    return %[3]s_query(db, &q);
}

/* Changes only the fields in patch (json, e.g. the request body) of the entity with this id.
 * Objects are merged field by field, a value or an array in patch replaces the stored one.
 * SQLITE_NOTFOUND if there is no such id. */
static inline int %[3]s_patch(sqlite3* db, int64_t id, const cJSON* patch) {
    return emisql_patch(db, &%[4]s, id, patch);
}
#endif /* EMI_NO_SQLITE */
`

// CEntityVirtualFiles renders every entity of the module, plus the schema
// header that creates (and later extends) its tables.
func CEntityVirtualFiles(module *core.Emi, ctx core.MicroGenContext, complexes []RecognizedComplex) ([]core.VirtualFile, error) {
	if len(module.Entities) == 0 {
		return nil, nil
	}
	withSqlite, withPack := ctx.HasTag(Sqlite), ctx.HasTag(Pack)
	result, err := entitysql.Generate(module.Entities, entitysql.Options{Dialect: entitysql.SQLite})
	if err != nil {
		return nil, err
	}
	layouts := map[string]*entitysql.Layout{}
	for _, l := range result.Layouts {
		layouts[l.Table] = l
	}

	files := []core.VirtualFile{}
	for _, entity := range module.Entities {
		if entity == nil {
			continue
		}
		layout := layouts[entitysql.TableName(entity)]
		if layout == nil {
			return nil, fmt.Errorf("entity %s: no table layout", entity.Name)
		}
		chunk, err := CCommonObjectGenerator(entity.Fields, ctx, CCommonObjectContext{
			RootTypeName:        entity.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}

		var tables strings.Builder
		cLayoutTables(&tables, layout, map[string]bool{})
		extra := fmt.Sprintf(cEntityLayoutTmpl, layout.Table, tables.String(), entity.GetClassName())
		chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies, core.CodeChunkDependency{Location: "emilayout.h"})
		if withPack {
			extra += fmt.Sprintf(cEntityPackTmpl, entity.GetClassName(), layoutVar(layout))
			chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies, core.CodeChunkDependency{Location: "emipack.h"})
		}
		if withSqlite {
			extra += fmt.Sprintf(cEntityStoreTmpl, layout.Table, "", entity.GetClassName(), layoutVar(layout))
		}
		chunk.ActualScript = append(chunk.ActualScript, []byte(extra)...)
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	if !withSqlite {
		return files, nil // the schema is for SQLite only
	}
	name := cIdent(core.ToSnakeCase(module.Name))
	if name == "" {
		name = "entities"
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "static const EmiSqlStatement %s_schema[] = {\n", name)
	for _, s := range result.Statements {
		table, column := "NULL", "NULL"
		if s.AddsColumn() {
			table, column = escapeDoubleQuoted(s.Table), escapeDoubleQuoted(s.Column)
		}
		fmt.Fprintf(&sb, "    {%s, %s, %s},\n", escapeDoubleQuoted(s.SQL), table, column)
	}
	sb.WriteString("};\n\n")
	fmt.Fprintf(&sb, "/* Creates the tables, and adds the columns a newer schema has; run it on every start. Returns 0 or an SQLITE_* code. */\n")
	fmt.Fprintf(&sb, "static inline int %s_schema_apply(sqlite3* db) {\n    return emisql_apply(db, %s_schema, sizeof(%s_schema) / sizeof(%s_schema[0]));\n}\n", name, name, name, name)

	schema := &core.CodeChunkCompiled{
		ActualScript:       []byte(sb.String()),
		SuggestedFileName:  name + "_schema",
		SuggestedExtension: ".h",
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "<sqlite3.h>"},
			{Location: "emisql.h"},
		},
	}
	files = append(files, core.VirtualFile{
		Name:         schema.SuggestedFileName,
		Extension:    schema.SuggestedExtension,
		ActualScript: AsFullDocument(schema),
	})
	return files, nil
}
