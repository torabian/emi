package emigorm

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/h22rana/jsonlogic2sql"
	"gorm.io/gorm"
)

// QueryDSL is the input to a generated {Entity}BrowseFn: a JSON-logic filter
// (https://jsonlogic.com) transpiled to a SQL WHERE clause via jsonlogic2sql, plus a
// plain offset/limit pager, a raw sort expression, a keyset-pagination cursor, and a
// second, independent scope condition. Deliberately much leaner than fireback's own
// QueryDSL (no gin/cli/websocket/permission coupling) - emi's generated code has no
// HTTP framework opinion of its own; a caller's own handler layer is expected to build
// one of these from whatever request shape it has.
type QueryDSL struct {
	// Filter is a JSON-logic expression as JSON text, e.g.
	// `{"contains":[{"var":"title"},"hello"]}`. Empty means no filtering at all. This is
	// the *user*-controlled half of filtering - see Scope for the other half.
	Filter string `json:"filter"`

	// Scope is a second, independent SQL WHERE fragment (optionally with `?`
	// placeholders, values in ScopeArgs) - e.g. "workspace_id = ?" - that the *caller's
	// own handler* enforces, never something derived from end-user input. Applied via
	// ApplyQueryScope, as its own step separate from Filter/ApplyQueryFilter, so a
	// user-supplied Filter has no way to widen or bypass it. json:"-" deliberately:
	// unlike Filter, Scope must never be settable from a request body/querystring the
	// way an EmiAction's fields are - only a handler assembling QueryDSL by hand sets
	// it, after that caller's identity/permissions are already known.
	Scope     string        `json:"-"`
	ScopeArgs []interface{} `json:"-"`

	// Sort is passed straight to gorm's Order() - e.g. "created_at desc". Empty means
	// the database's own default order.
	Sort string `json:"sort"`

	// StartIndex/ItemsPerPage are a plain offset/limit pager. Values <= 0 mean that
	// part of paging isn't applied (no offset / no limit, respectively).
	StartIndex   int `json:"startIndex"`
	ItemsPerPage int `json:"itemsPerPage"`

	// Cursor resumes a previous page - always exactly what a prior call's
	// emigo.QueryResultMeta.Cursor returned (see ApplyQueryCursor/BuildQueryCursor);
	// empty means "start from the beginning" (or use StartIndex, if set, for a plain
	// offset instead).
	Cursor string `json:"cursor"`
}

// registerQueryOperators adds the custom operators emi's Query filter understands on
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

// ApplyQueryFilter transpiles dsl.Filter (if set) to SQL and applies it to tx via
// Where(), returning the resulting *gorm.DB. Deliberately does *not* apply
// Sort/StartIndex/ItemsPerPage - see ApplyQuerySort/ApplyQueryPage - since a caller
// almost always wants to run Count() against the filtered-but-unpaged query before
// paging the actual Find(), and gorm's Count() only ignores Limit/Offset reliably when
// they were never added to the statement being counted in the first place.
func ApplyQueryFilter(tx *gorm.DB, dsl QueryDSL) (*gorm.DB, error) {
	if dsl.Filter == "" {
		return tx, nil
	}

	tr, err := jsonlogic2sql.NewTranspiler(jsonlogic2sql.DialectPostgreSQL)
	if err != nil {
		return nil, err
	}
	registerQueryOperators(tr)

	sql, err := tr.TranspileCondition(dsl.Filter)
	if err != nil {
		return nil, err
	}
	if sql == "" {
		return tx, nil
	}
	return tx.Where(sql), nil
}

// ApplyQueryScope applies dsl.Scope (if set) via Where(), exactly like ApplyQueryFilter
// but for the condition a caller's own handler enforces rather than one derived from
// end-user input (see QueryDSL.Scope) - e.g. workspace/tenant isolation. Kept as its
// own step (rather than folded into ApplyQueryFilter) precisely so it can be applied
// unconditionally, with Filter having no way to interact with or override it.
func ApplyQueryScope(tx *gorm.DB, dsl QueryDSL) *gorm.DB {
	if dsl.Scope == "" {
		return tx
	}
	return tx.Where(dsl.Scope, dsl.ScopeArgs...)
}

// ApplyQuerySort applies dsl.Sort (if set) via Order().
func ApplyQuerySort(tx *gorm.DB, dsl QueryDSL) *gorm.DB {
	if dsl.Sort == "" {
		return tx
	}
	return tx.Order(dsl.Sort)
}

// ApplyQueryPage applies dsl.StartIndex/ItemsPerPage (if set) via Offset()/Limit().
func ApplyQueryPage(tx *gorm.DB, dsl QueryDSL) *gorm.DB {
	if dsl.StartIndex > 0 {
		tx = tx.Offset(dsl.StartIndex)
	}
	if dsl.ItemsPerPage > 0 {
		tx = tx.Limit(dsl.ItemsPerPage)
	}
	return tx
}

// cursorPattern matches the "id(<value>)" cursor format ApplyQueryCursor/
// BuildQueryCursor use - mirrors fireback's own cursor encoding exactly (see
// fireback's CrudCoreActions.go, parseCursor/QueryEntitiesPointer), minus the
// "+sort(...)" suffix fireback tacks on (the same information already travels
// separately as QueryDSL.Sort, so it doesn't need to be re-embedded in the cursor).
var cursorPattern = regexp.MustCompile(`^(\w+)\((\d+)\)$`)

// ApplyQueryCursor applies dsl.Cursor (if set and well-formed) as a keyset-pagination
// WHERE clause, picking up right after the last row a previous page returned - id > n,
// where n is whatever BuildQueryCursor encoded into the cursor a prior call returned.
// A malformed or empty cursor is silently ignored (same behavior ApplyQueryFilter has
// for a bad filter), returning tx unchanged.
//
// Like fireback's own implementation, the comparison operator is unconditionally ">" -
// correct for the common "ascending by id" case, but not adjusted for a descending
// dsl.Sort. Keyset pagination direction-awareness is a real gap in both, not something
// unique to this port.
func ApplyQueryCursor(tx *gorm.DB, dsl QueryDSL) *gorm.DB {
	if dsl.Cursor == "" {
		return tx
	}
	matches := cursorPattern.FindStringSubmatch(dsl.Cursor)
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
