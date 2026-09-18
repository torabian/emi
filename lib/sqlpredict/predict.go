// Package sqlpredict detects the projected columns of a SELECT statement by
// parsing it - the same column-detection mechanism lib/querypredict has long
// used to generate typed row structs from hand-written SQL, extracted here
// (dependency-free of lib/core) so lib/core's vsql preprocessing can use it
// too, under EmiVsql.Predict, without an import cycle (lib/querypredict
// already depends on lib/core). The intent is for this package to outlive
// lib/querypredict, not sit alongside it forever.
package sqlpredict

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/xwb1989/sqlparser"
)

// Column is one projected SELECT expression: what to call it (Name), the
// literal SQL text that produces it (Source), its emi/Go type, and whether
// it's optional (nullable).
type Column struct {
	// Name is a Go-identifier-safe field name: the third field() argument
	// verbatim if given, otherwise derived from Source (e.g. "u.user_name"
	// or "user_name" both become "UserName").
	Name string

	// Source is the literal SQL expression to project - a plain column
	// name, a qualified one, or an arbitrary expression (aggregate,
	// function call, ...), field()'s wrapping stripped off.
	Source string

	// Type is an emi/Go type name (string, int64, float64, bool, ...) -
	// field()'s second argument verbatim if given, otherwise "string".
	Type string

	// Optional is field()'s fourth argument - true/false or a nonzero
	// int literal.
	Optional bool
}

var templateMarkup = regexp.MustCompile(`\{\{[^{}]*\}\}`)

// stripTemplateMarkup replaces every {{ ... }} block (vsql's Go-template
// syntax) with the SQL literal NULL, so a query mixing template markup into
// its WHERE/VALUES/RETURNING clauses still parses as syntactically valid SQL
// for column-detection purposes - only the SELECT list actually needs to be
// real SQL here; everything replaced is discarded anyway; column detection
// never executes this text.
func stripTemplateMarkup(query string) string {
	return templateMarkup.ReplaceAllString(query, "NULL")
}

// DetectSelectColumns parses query (after stripTemplateMarkup) and returns
// one Column per entry in its SELECT list, in order.
//
// Only a genuine SELECT is supported, matching lib/querypredict's own
// long-standing scope: the underlying parser implements MySQL grammar and
// has no notion of RETURNING, so an INSERT/UPDATE/DELETE ... RETURNING query
// - the shape most vsql examples in this repo actually use - can't be
// predicted this way. Declare Columns by hand for those; Predict is for a
// query whose whole job is a SELECT's projection.
//
// A SELECT list containing "*" (bare, or table-qualified like "u.*")
// returns (nil, nil) - no error, no columns - rather than a guess: what "*"
// expands to depends on the live table schema, and inspecting that is
// exactly what this project's SQL philosophy avoids (see
// sql_design_philosophy). A caller (predictVsqlColumns) treats that the same
// as "nothing to populate" - the vsql simply ends up with no Columns/no row
// DTO, same as never having set predict: true at all. This applies even to
// a select list that mixes "*" with named columns: a partial, silently
// incomplete column list would be worse than declaring defeat outright.
func DetectSelectColumns(query string) ([]Column, error) {
	stripped := stripTemplateMarkup(query)

	stmt, err := sqlparser.Parse(stripped)
	if err != nil {
		return nil, fmt.Errorf("parse query for column prediction: %w", err)
	}
	sel, ok := stmt.(*sqlparser.Select)
	if !ok {
		return nil, fmt.Errorf("predict only supports a plain SELECT statement, got %T", stmt)
	}

	for _, expr := range sel.SelectExprs {
		if _, ok := expr.(*sqlparser.StarExpr); ok {
			return nil, nil
		}
	}

	var cols []Column
	for _, expr := range sel.SelectExprs {
		if e, ok := expr.(*sqlparser.AliasedExpr); ok {
			cols = append(cols, columnFromAliasedExpr(e))
		}
	}

	if dup := firstDuplicateName(cols); dup != "" {
		return nil, fmt.Errorf(
			"two projected columns both resolve to the Go field name %q - "+
				"alias one of them (\"... AS someName\") or give it an explicit "+
				"field(expr, 'type', 'goName') name to disambiguate", dup)
	}

	return cols, nil
}

// firstDuplicateName returns the first Column.Name that appears more than
// once, or "" if every name is unique. A join selecting the same bare column
// name from two tables (e.g. "SELECT u.id, o.id ...") is the common way this
// happens if the query doesn't alias one of them - Source keeps the table
// qualifier (see columnFromAliasedExpr), so that case is already handled;
// this is the backstop for whatever it doesn't catch.
func firstDuplicateName(cols []Column) string {
	seen := make(map[string]bool, len(cols))
	for _, c := range cols {
		if seen[c.Name] {
			return c.Name
		}
		seen[c.Name] = true
	}
	return ""
}

// columnFromAliasedExpr mirrors lib/querypredict/sql-ast.go's handleExpr +
// handleFuncExpr, minus the SelectColumn/AST-node bookkeeping this package
// has no use for - with one deliberate departure: Source keeps a column's
// table qualifier (sqlparser.String renders "u.id", not just "id"), where
// lib/querypredict discarded it. Dropping it silently produced ambiguous SQL
// and duplicate Go field names for any join selecting the same column name
// from two tables (e.g. "SELECT u.id, o.id FROM users u JOIN orders o ...").
func columnFromAliasedExpr(e *sqlparser.AliasedExpr) Column {
	col := Column{Type: "string"}

	isField := false
	if fn, ok := e.Expr.(*sqlparser.FuncExpr); ok && strings.ToLower(fn.Name.Lowered()) == "field" {
		col = fieldFuncColumn(fn)
		isField = true
	} else {
		col.Source = sqlparser.String(e.Expr)
	}

	if alias := e.As.String(); alias != "" {
		if col.Name == "" {
			col.Name = MakeValidGoField(alias)
		}
		if !isField {
			// Reproject exactly as written - a bare alias isn't necessarily a
			// real column of the underlying table(s) on its own, so the
			// original expression has to travel with it.
			col.Source += " AS " + alias
		}
	}

	if col.Name == "" {
		col.Name = MakeValidGoField(col.Source)
	}
	return col
}

// fieldFuncColumn parses field(sqlExpr, 'type'?, 'goName'?, optional?).
func fieldFuncColumn(fn *sqlparser.FuncExpr) Column {
	col := Column{Type: "string"}

	arg := func(i int) (sqlparser.Expr, bool) {
		if len(fn.Exprs) <= i {
			return nil, false
		}
		ae, ok := fn.Exprs[i].(*sqlparser.AliasedExpr)
		if !ok {
			return nil, false
		}
		return ae.Expr, true
	}

	if e, ok := arg(0); ok {
		// Keep any table qualifier (see columnFromAliasedExpr's doc comment) -
		// sqlparser.String renders it whether e is a bare ColName or anything
		// else.
		col.Source = sqlparser.String(e)
	}
	if e, ok := arg(1); ok {
		if sv, ok := e.(*sqlparser.SQLVal); ok && sv.Type == sqlparser.StrVal {
			col.Type = string(sv.Val)
		}
	}
	if e, ok := arg(2); ok {
		if sv, ok := e.(*sqlparser.SQLVal); ok && sv.Type == sqlparser.StrVal {
			col.Name = string(sv.Val)
		}
	}
	if e, ok := arg(3); ok {
		switch v := e.(type) {
		case *sqlparser.SQLVal:
			if v.Type == sqlparser.StrVal {
				col.Optional = strings.EqualFold(string(v.Val), "true")
			} else if v.Type == sqlparser.IntVal {
				col.Optional = string(v.Val) != "0"
			}
		case sqlparser.BoolVal:
			col.Optional = bool(v)
		}
	}

	return col
}

// MakeValidGoField turns an arbitrary SQL expression/column name into a
// PascalCase Go identifier, e.g. "u.user_name" -> "UserName",
// "count(*)" -> "Count". Ported unchanged from
// lib/querypredict/generator.go so predicted names match what Query Predict
// itself would have produced for the same query.
func MakeValidGoField(col string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	clean := re.ReplaceAllString(col, " ")

	parts := strings.Fields(clean)
	for i, p := range parts {
		parts[i] = strings.Title(p) //nolint:staticcheck // matches lib/querypredict's own behavior exactly
	}
	field := strings.Join(parts, "")

	if field == "" {
		return "Col"
	}
	runes := []rune(field)
	if !unicode.IsUpper(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}
