package emigorm

import (
	"strings"
	"testing"

	"github.com/h22rana/jsonlogic2sql"
	"gorm.io/gorm"
	gormtests "gorm.io/gorm/utils/tests"
)

func TestRegisterQueryOperators_ContainsTranspilesToCaseInsensitiveLike(t *testing.T) {
	tr, err := jsonlogic2sql.NewTranspiler(jsonlogic2sql.DialectPostgreSQL)
	if err != nil {
		t.Fatalf("NewTranspiler error: %v", err)
	}
	registerQueryOperators(tr)

	sql, err := tr.TranspileCondition(`{"contains":[{"var":"title"},"hello"]}`)
	if err != nil {
		t.Fatalf("TranspileCondition error: %v", err)
	}
	if !strings.Contains(sql, "ILIKE") || !strings.Contains(sql, "%hello%") {
		t.Fatalf("expected an ILIKE %%hello%% clause, got %q", sql)
	}
	// See registerQueryOperators' own doc comment: the column is always cast to
	// ::text first, so "contains" also works against non-text columns - most
	// notably a `complex: TString` field, stored as jsonb, which a bare
	// `column ILIKE '%...%'` fails against outright ("operator does not exist:
	// jsonb ~~* unknown").
	if !strings.Contains(sql, "title::text") {
		t.Fatalf("expected the column to be cast to ::text before ILIKE, got %q", sql)
	}
}

// TestRegisterQueryOperators_ContainsDoesNotDoubleEscapeQuotes guards against
// both a broken query and a straightforward SQL injection - and against
// re-introducing a specific regression: jsonlogic2sql already quotes *and*
// escapes a string literal argument before handing it to a custom operator
// (Parser.primitiveToSQL doubles embedded `'`, then wraps the result in
// `'...'`), so registerQueryOperators must only strip that outer quote pair
// back off, not escape the value a second time (which would double `''` into
// `''''`, corrupting the query).
func TestRegisterQueryOperators_ContainsDoesNotDoubleEscapeQuotes(t *testing.T) {
	tr, err := jsonlogic2sql.NewTranspiler(jsonlogic2sql.DialectPostgreSQL)
	if err != nil {
		t.Fatalf("NewTranspiler error: %v", err)
	}
	registerQueryOperators(tr)

	sql, err := tr.TranspileCondition(`{"contains":[{"var":"title"},"o'brien' OR '1'='1"]}`)
	if err != nil {
		t.Fatalf("TranspileCondition error: %v", err)
	}
	if !strings.Contains(sql, `o''brien'' OR ''1''=''1`) {
		t.Fatalf("expected every single quote doubled exactly once, got %q", sql)
	}
	if strings.Contains(sql, `''''`) {
		t.Fatalf("expected quotes not to be double-escaped, got %q", sql)
	}
}

func TestApplyQueryFilter_EmptyFilterReturnsTxUnchanged(t *testing.T) {
	// A zero-value *gorm.DB is enough here: ApplyQueryFilter must return early,
	// without ever touching tx, when there's no filter to apply - exercising the
	// Where() path needs a real, initialized session (see the integration-tagged
	// entity tests instead).
	tx := &gorm.DB{}

	out, err := ApplyQueryFilter(tx, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != tx {
		t.Fatal("expected the exact same *gorm.DB back when Filter is empty")
	}
}

// TestApplyQuerySort_DefaultsToIdAscWhenEmpty exercises real SQL generation
// (gorm's DummyDialector - no actual database needed) rather than the
// zero-value-*gorm.DB shortcut the other tests here use, since this is
// exactly the path that shortcut can't cover: whether Order() actually gets
// called when sort is empty. See ApplyQuerySort's own doc comment for why
// this default matters - ApplyQueryCursor/BuildQueryCursor's keyset paging
// is silently wrong without it.
func TestApplyQuerySort_DefaultsToIdAscWhenEmpty(t *testing.T) {
	type row struct {
		Id   int64
		Name string
	}

	db, err := gorm.Open(gormtests.DummyDialector{}, &gorm.Config{})
	if err != nil {
		t.Fatalf("unexpected error opening dummy dialector: %v", err)
	}

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		var out []row
		return ApplyQuerySort(tx.Model(&row{}), "").Find(&out)
	})
	if !strings.Contains(strings.ToLower(sql), "order by id asc") {
		t.Fatalf("expected an empty sort to default to \"ORDER BY id asc\", got: %s", sql)
	}

	// An explicit sort must still be honored as-is, not overridden.
	sql = db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		var out []row
		return ApplyQuerySort(tx.Model(&row{}), "name desc").Find(&out)
	})
	if !strings.Contains(strings.ToLower(sql), "order by name desc") {
		t.Fatalf("expected an explicit sort to be honored as-is, got: %s", sql)
	}
}

func TestApplyQueryCursor_EmptyOrMalformedCursorReturnsTxUnchanged(t *testing.T) {
	tx := &gorm.DB{}

	for _, cursor := range []string{"", "not-a-cursor", "sort(title)"} {
		out := ApplyQueryCursor(tx, cursor)
		if out != tx {
			t.Fatalf("cursor %q: expected the exact same *gorm.DB back, got a different one", cursor)
		}
	}
}

type queryCursorTestRow struct {
	Id       int64
	UniqueId string
}

func TestBuildQueryCursor_EncodesLastItemsIdOrNilWhenEmpty(t *testing.T) {
	if cursor := BuildQueryCursor([]*queryCursorTestRow{}); cursor != nil {
		t.Fatalf("expected nil cursor for an empty page, got %q", *cursor)
	}

	items := []*queryCursorTestRow{
		{Id: 41, UniqueId: "a"},
		{Id: 42, UniqueId: "b"},
	}
	cursor := BuildQueryCursor(items)
	if cursor == nil || *cursor != "id(42)" {
		t.Fatalf("expected cursor \"id(42)\" (the last item's Id), got %v", cursor)
	}
}
