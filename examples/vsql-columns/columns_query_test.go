// Exercises InsertUsersVsqlColumnsQuery - building the column picker
// straight from a URL query string (or an *http.Request) instead of
// constructing InsertUsersVsqlColumns by hand in Go. Useful once a vsql is
// exposed over HTTP: a client's "?email=true&firstName=true" becomes the
// exact same picker NewInsertUsersVsqlColumns() + manual field flips would
// have produced.
package vsqlcolumns_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	external "github.com/torabian/emi/examples/vsql-columns/sdkgen"
)

func TestColumnsQueryFromStringOverridesOnlyMentionedKeys(t *testing.T) {
	q := external.InsertUsersVsqlColumnsQueryFromString("firstName=true&email=false")
	cols := q.Columns()

	if !cols.FirstName.Selected {
		t.Fatalf("expected firstName=true from the query string to select FirstName")
	}
	if cols.Email.Selected {
		t.Fatalf("expected email=false from the query string to deselect Email")
	}

	// Every column the query string never mentioned keeps its own declared
	// default (see batch_insert_users.emi.yml) rather than becoming false.
	defaults := external.NewInsertUsersVsqlColumns()
	if cols.Id.Selected != defaults.Id.Selected {
		t.Fatalf("expected untouched Id to keep its default, got %+v", cols.Id)
	}
	if cols.Status.Selected != defaults.Status.Selected {
		t.Fatalf("expected untouched Status to keep its default, got %+v", cols.Status)
	}
	if cols.Money.Selected != defaults.Money.Selected {
		t.Fatalf("expected untouched Money to keep its default, got %+v", cols.Money)
	}
}

func TestColumnsQueryFromStringEmptyKeepsAllDefaults(t *testing.T) {
	got := external.InsertUsersVsqlColumnsQueryFromString("").Columns()
	want := external.NewInsertUsersVsqlColumns()
	if got != want {
		t.Fatalf("expected an empty query string to reproduce the plain defaults exactly, got %+v want %+v", got, want)
	}
}

func TestColumnsQueryFromHttpReadsRequestQueryString(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users?money=true&status=false", nil)

	cols := external.InsertUsersVsqlColumnsQueryFromHttp(req).Columns()

	if !cols.Money.Selected {
		t.Fatalf("expected money=true on the request to select Money")
	}
	if cols.Status.Selected {
		t.Fatalf("expected status=false on the request to deselect Status")
	}
}

// TestColumnsQueryThenPrepare is the end-to-end shape: a caller parses an
// incoming query string straight into the picker Prepare needs, with no
// manual field-by-field translation in between.
func TestColumnsQueryThenPrepare(t *testing.T) {
	columns := external.InsertUsersVsqlColumnsQueryFromString("firstName=true").Columns()
	params := external.InsertUsersVsqlParams{}

	query, args := external.PrepareInsertUsersVsql(params, columns)
	sql, err := render(query, args)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if !strings.Contains(sql, "first_name") {
		t.Fatalf("expected first_name to join the RETURNING projection, got:\n%s", sql)
	}
}
