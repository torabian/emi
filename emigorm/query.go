package emigorm

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/h22rana/jsonlogic2sql"
	"gorm.io/gorm"
)

// registerQueryOperators adds the custom operators emi's Browse filter understands on
// top of jsonlogic2sql's own built-ins (==, !=, >, <, and, or, in, ...). "contains" -
// substring match via LIKE - mirrors fireback's own JsonQueryTools.go exactly, since
// json-logic has no native substring operator.
func registerQueryOperators(tr *jsonlogic2sql.Transpiler) {
	tr.RegisterOperatorFunc("contains", func(op string, args []interface{}) (string, error) {
		if len(args) != 2 {
			return "", fmt.Errorf("contains requires 2 arguments")
		}
		column, _ := args[0].(string)
		value, _ := args[1].(string)
		value = strings.Trim(value, `"'`)
		return fmt.Sprintf("%s LIKE '%%%s%%'", column, value), nil
	})
}

// ApplyQueryFilter transpiles filter (a JSON-logic expression, https://jsonlogic.com -
// e.g. `{"contains":[{"var":"title"},"hello"]}` - empty means no filtering at all) to
// SQL via jsonlogic2sql and applies it to tx via Where(), returning the resulting
// *gorm.DB. filter is the *user*-controlled half of filtering - see ApplyQueryScope for
// the other half.
//
// Deliberately does *not* apply sort/paging - see ApplyQuerySort/ApplyQueryPage - since
// a caller almost always wants to run Count() against the filtered-but-unpaged query
// before paging the actual Find(), and gorm's Count() only ignores Limit/Offset
// reliably when they were never added to the statement being counted in the first
// place.
func ApplyQueryFilter(tx *gorm.DB, filter string) (*gorm.DB, error) {
	if filter == "" {
		return tx, nil
	}

	tr, err := jsonlogic2sql.NewTranspiler(jsonlogic2sql.DialectPostgreSQL)
	if err != nil {
		return nil, err
	}
	registerQueryOperators(tr)

	sql, err := tr.TranspileCondition(filter)
	if err != nil {
		return nil, err
	}
	if sql == "" {
		return tx, nil
	}
	return tx.Where(sql), nil
}

// ApplyQueryScope applies scope (a raw SQL WHERE fragment, optionally with `?`
// placeholders, values in scopeArgs - e.g. "workspace_id = ?" - empty means no scoping
// at all) via Where(), exactly like ApplyQueryFilter but for a condition the *caller's
// own handler* enforces - e.g. workspace/tenant isolation - rather than one derived
// from end-user input. Kept as its own step (rather than folded into ApplyQueryFilter)
// precisely so it can be applied unconditionally, with filter having no way to interact
// with or override it.
func ApplyQueryScope(tx *gorm.DB, scope string, scopeArgs ...interface{}) *gorm.DB {
	if scope == "" {
		return tx
	}
	return tx.Where(scope, scopeArgs...)
}

// ApplyQuerySort applies sort (if non-empty) via Order() - e.g. "created_at desc".
func ApplyQuerySort(tx *gorm.DB, sort string) *gorm.DB {
	if sort == "" {
		return tx
	}
	return tx.Order(sort)
}

// ApplyQueryPage applies startIndex/itemsPerPage (values <= 0 mean that part of paging
// isn't applied - no offset / no limit, respectively) via Offset()/Limit().
func ApplyQueryPage(tx *gorm.DB, startIndex int, itemsPerPage int) *gorm.DB {
	if startIndex > 0 {
		tx = tx.Offset(startIndex)
	}
	if itemsPerPage > 0 {
		tx = tx.Limit(itemsPerPage)
	}
	return tx
}

// cursorPattern matches the "id(<value>)" cursor format ApplyQueryCursor/
// BuildQueryCursor use - mirrors fireback's own cursor encoding exactly (see
// fireback's CrudCoreActions.go, parseCursor/QueryEntitiesPointer), minus the
// "+sort(...)" suffix fireback tacks on (the same information already travels
// separately as the sort value, so it doesn't need to be re-embedded in the cursor).
var cursorPattern = regexp.MustCompile(`^(\w+)\((\d+)\)$`)

// ApplyQueryCursor applies cursor (if non-empty and well-formed) as a
// keyset-pagination WHERE clause, picking up right after the last row a previous page
// returned - id > n, where n is whatever BuildQueryCursor encoded into the cursor a
// prior call returned. Resumes a previous page - cursor is always exactly what a prior
// call's own emigo.QueryResultMeta.Cursor returned; empty means "start from the
// beginning" (or use startIndex, if set, for a plain offset instead). A malformed or
// empty cursor is silently ignored (same behavior ApplyQueryFilter has for a bad
// filter), returning tx unchanged.
//
// Like fireback's own implementation, the comparison operator is unconditionally ">" -
// correct for the common "ascending by id" case, but not adjusted for a descending
// sort. Keyset pagination direction-awareness is a real gap in both, not something
// unique to this port.
func ApplyQueryCursor(tx *gorm.DB, cursor string) *gorm.DB {
	if cursor == "" {
		return tx
	}
	matches := cursorPattern.FindStringSubmatch(cursor)
	if len(matches) != 3 || matches[1] != "id" {
		return tx
	}
	return tx.Where("id > ?", matches[2])
}

// BuildQueryCursor derives the outgoing cursor for a page of results: the last item's
// internal Id (see idOf - the real auto-increment primary key, never UniqueId), encoded
// the same "id(<value>)" way ApplyQueryCursor parses. nil once items is empty, since
// there's nothing to resume from.
func BuildQueryCursor[T any](items []*T) *string {
	if len(items) == 0 {
		return nil
	}
	cursor := "id(" + strconv.FormatInt(idOf(items[len(items)-1]), 10) + ")"
	return &cursor
}
