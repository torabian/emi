package emigorm

import (
	"strings"
	"testing"

	"github.com/h22rana/jsonlogic2sql"
	"gorm.io/gorm"
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
