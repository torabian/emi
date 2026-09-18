// Runs all ten vsql queries generated from queries.emi.yml against a real,
// in-memory SQLite database - not just checking the rendered SQL text (as
// examples/vsql-columns does), but actually executing it and asserting on
// what comes back. Subtests run in declaration order against one shared
// connection, since later queries depend on rows earlier ones inserted.
//
// Rendering uses examples/vsql/vsql.Render, the reference renderer, so
// Nullable[T] params/columns get its real reflection-based handling (an
// unset Nullable[T] is quietly omitted or must be IsSet-guarded - see that
// package's doc comment) rather than a bespoke minimal one.
package vsqlsqlite_test

import (
	"database/sql"
	"encoding/json"
	"strings"
	"testing"
	"testing/fstest"

	_ "github.com/mattn/go-sqlite3"

	"github.com/torabian/emi/emigo"
	external "github.com/torabian/emi/examples/vsql-sqlite/sdkgen"
	"github.com/torabian/emi/examples/vsql-sqlite/sqlfiles"
	"github.com/torabian/emi/examples/vsql/vsql"
)

// render wraps vsql.Render so an inline query string (returned directly by
// most Prepare*Vsql helpers) can go through the same renderer as
// queryName-based ones, which already come as an fs.FS + file name.
func render(t *testing.T, query string, data any) string {
	t.Helper()
	fsys := fstest.MapFS{"q.sql": &fstest.MapFile{Data: []byte(query)}}
	out, err := vsql.Render(fsys, "q.sql", data)
	if err != nil {
		t.Fatalf("render: %v\nquery:\n%s", err, query)
	}
	return out
}

func TestVsqlSqliteFullSuite(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	defer db.Close()

	// ── 1. createSchema: no params, no columns ──────────────────────────
	t.Run("createSchema", func(t *testing.T) {
		query, params := external.PrepareCreateSchemaVsql(external.CreateSchemaVsqlParams{})
		sqlText := render(t, query, params)
		if _, err := db.Exec(sqlText); err != nil {
			t.Fatalf("exec: %v\nsql:\n%s", err, sqlText)
		}
	})

	// ── 2. insertUser: scalar params (one nullable) + columns/RETURNING,
	//      one of which (isActive) defaults unselected. ─────────────────
	var aliceID int64
	t.Run("insertUser", func(t *testing.T) {
		params := external.InsertUserVsqlParams{
			Email:     "alice@example.com",
			FirstName: "Alice",
			// LastName left unset on purpose - exercises the {{ if
			// .LastName.IsSet }} guard against a genuinely-unset Nullable.
			Role:     "Admin",
			IsActive: true,
		}
		columns := external.NewInsertUserVsqlColumns() // id, email, role selected; isActive not

		query, args := external.PrepareInsertUserVsql(params, columns)
		sqlText := render(t, query, args)

		row := db.QueryRow(sqlText)
		var got external.InsertUserVsqlRow
		if err := row.Scan(&got.Id, &got.Email, &got.Role); err != nil {
			t.Fatalf("scan: %v\nsql:\n%s", err, sqlText)
		}
		aliceID = got.Id

		if got.Email != "alice@example.com" || got.Role != "Admin" {
			t.Fatalf("unexpected row: %+v", got)
		}
		// isActive was never selected, so its column never reached the
		// RETURNING list at all - nothing to scan for it. That's the
		// point: IsActive's Nullable[bool] type exists for callers that DO
		// select it (query 6 does), not this one.
	})

	t.Run("insertUser_withLastName", func(t *testing.T) {
		params := external.InsertUserVsqlParams{
			Email:     "bob@example.com",
			FirstName: "Bob",
			LastName:  emigo.NullableOf("Marley"),
			Role:      "Member",
			IsActive:  false,
		}
		columns := external.NewInsertUserVsqlColumns()

		query, args := external.PrepareInsertUserVsql(params, columns)
		sqlText := render(t, query, args)

		if !containsQuoted(sqlText, "Marley") {
			t.Fatalf("expected a set LastName to render as a quoted literal, got:\n%s", sqlText)
		}
		if _, err := db.Exec(sqlText); err != nil {
			t.Fatalf("exec: %v\nsql:\n%s", err, sqlText)
		}
	})

	// ── 3. insertUsersBatch: array param + complex Money column spanning
	//      two physical columns. ────────────────────────────────────────
	t.Run("insertUsersBatch", func(t *testing.T) {
		params := external.InsertUsersBatchVsqlParams{
			Users: emigo.ArrayReplace([]external.InsertUsersBatchVsqlParamsUsers{
				{Email: "carol@example.com", FirstName: "Carol", LastName: "Danvers", MoneyAmountCents: 5000, MoneyCurrency: "USD"},
				{Email: "dave@example.com", FirstName: "Dave", LastName: "Grohl", MoneyAmountCents: 0, MoneyCurrency: "USD"},
			}),
		}
		columns := external.NewInsertUsersBatchVsqlColumns()
		columns.Money.Selected = true // off by default - turn it on for this check

		query, args := external.PrepareInsertUsersBatchVsql(params, columns)
		sqlText := render(t, query, args)

		rows, err := db.Query(sqlText)
		if err != nil {
			t.Fatalf("query: %v\nsql:\n%s", err, sqlText)
		}
		defer rows.Close()

		count := 0
		for rows.Next() {
			var row external.InsertUsersBatchVsqlRow
			var amountCents int64
			var currency string
			if err := rows.Scan(&row.Id, &row.Email, &amountCents, &currency); err != nil {
				t.Fatalf("scan: %v", err)
			}
			row.Money = external.Money{AmountCents: amountCents, Currency: currency}
			count++
		}
		if count != 2 {
			t.Fatalf("expected 2 rows back from the batch insert, got %d", count)
		}
	})

	// ── 4. searchUsersByFilter: fields sourced entirely via captures. ───
	t.Run("searchUsersByFilter_byEmail", func(t *testing.T) {
		params := external.SearchUsersByFilterVsqlParams{
			Email: emigo.NullableOf("alice@example.com"),
			// Role left unset - only the email filter should apply.
		}
		query, args := external.PrepareSearchUsersByFilterVsql(params)
		sqlText := render(t, query, args)

		rows, err := db.Query(sqlText)
		if err != nil {
			t.Fatalf("query: %v\nsql:\n%s", err, sqlText)
		}
		defer rows.Close()

		var emails []string
		for rows.Next() {
			var id int64
			var email string
			if err := rows.Scan(&id, &email); err != nil {
				t.Fatalf("scan: %v", err)
			}
			emails = append(emails, email)
		}
		if len(emails) != 1 || emails[0] != "alice@example.com" {
			t.Fatalf("expected exactly alice@example.com, got %v", emails)
		}
	})

	t.Run("searchUsersByFilter_noFilter", func(t *testing.T) {
		// Both fields unset - every previously-inserted row should come
		// back (alice, bob, carol, dave = 4).
		query, args := external.PrepareSearchUsersByFilterVsql(external.SearchUsersByFilterVsqlParams{})
		sqlText := render(t, query, args)

		wrapped := "SELECT count(*) FROM (" + strings.TrimRight(sqlText, "; \n\t") + ")"
		var count int
		if err := db.QueryRow(wrapped).Scan(&count); err != nil {
			t.Fatalf("count wrapper query: %v\nsql:\n%s", err, wrapped)
		}
		if count != 4 {
			t.Fatalf("expected all 4 rows with no filter set, got %d", count)
		}
	})

	// ── 5. listUsersSorted: `in.dto` (orderBy) merged with inline params
	//      (limit/offset). ──────────────────────────────────────────────
	t.Run("listUsersSorted", func(t *testing.T) {
		params := external.ListUsersSortedVsqlParams{
			OrderBy: "email",
			Limit:   2,
			Offset:  0,
		}
		query, args := external.PrepareListUsersSortedVsql(params)
		sqlText := render(t, query, args)

		rows, err := db.Query(sqlText)
		if err != nil {
			t.Fatalf("query: %v\nsql:\n%s", err, sqlText)
		}
		defer rows.Close()

		var emails []string
		for rows.Next() {
			var id int64
			var email, first string
			var last sql.NullString // alice's last_name is genuinely NULL - never set in insertUser
			if err := rows.Scan(&id, &email, &first, &last); err != nil {
				t.Fatalf("scan: %v", err)
			}
			emails = append(emails, email)
		}
		if len(emails) != 2 || emails[0] != "alice@example.com" || emails[1] != "bob@example.com" {
			t.Fatalf("expected the first 2 emails alphabetically, got %v", emails)
		}
	})

	// ── 6. selectUserById: full nullability + object column via
	//      json_object(). ───────────────────────────────────────────────
	t.Run("selectUserById", func(t *testing.T) {
		params := external.SelectUserByIdVsqlParams{Id: aliceID}
		columns := external.NewSelectUserByIdVsqlColumns()
		columns.IsActive.Selected = true
		columns.Address.Selected = true

		query, args := external.PrepareSelectUserByIdVsql(params, columns)
		sqlText := render(t, query, args)

		var row external.SelectUserByIdVsqlRow
		var addressJSON sql.NullString
		if err := db.QueryRow(sqlText).Scan(&row.Id, &row.Email, &row.Role, &row.IsActive, &addressJSON); err != nil {
			t.Fatalf("scan: %v\nsql:\n%s", err, sqlText)
		}

		if row.Email != "alice@example.com" {
			t.Fatalf("unexpected email: %s", row.Email)
		}
		if v, ok := row.IsActive.Get(); !ok || v == nil || *v != true {
			t.Fatalf("expected isActive to scan as set/true, got %+v", row.IsActive)
		}
		// address was never set at insert time - json_object() still
		// returns a real (all-null) JSON object, not SQL NULL, since
		// street/city are selected columns of the row, just empty.
		if !addressJSON.Valid {
			t.Fatalf("expected json_object() to produce a value even for empty street/city")
		}
		var addr external.SelectUserByIdVsqlRowAddress
		if err := json.Unmarshal([]byte(addressJSON.String), &addr); err != nil {
			t.Fatalf("unmarshal address json %q: %v", addressJSON.String, err)
		}
	})

	// ── 7. updateUserPartial: Columns/Selected reused to gate SET clauses,
	//      not a projection. ────────────────────────────────────────────
	t.Run("updateUserPartial", func(t *testing.T) {
		params := external.UpdateUserPartialVsqlParams{
			Id:           aliceID,
			BalanceCents: emigo.NullableOf(int64(1234)),
			// Email/FirstName/LastName left unset - their columns stay
			// unselected below, so COALESCE keeps the existing values.
		}
		columns := external.NewUpdateUserPartialVsqlColumns()
		columns.BalanceCents.Selected = true

		query, args := external.PrepareUpdateUserPartialVsql(params, columns)
		sqlText := render(t, query, args)

		// UpdateUserPartialVsqlRow (built from Columns) isn't the right
		// shape to scan this into - Columns means "which SET clause to
		// apply" here, not "which columns RETURNING selects" (that's
		// hand-written below), so its fields don't line up with what's
		// actually returned. Plain locals instead.
		var id int64
		var email, firstName string
		var lastName sql.NullString // alice's is genuinely NULL - never set in insertUser
		var balance sql.NullInt64
		if err := db.QueryRow(sqlText).Scan(&id, &email, &firstName, &lastName, &balance); err != nil {
			t.Fatalf("scan: %v\nsql:\n%s", err, sqlText)
		}

		if email != "alice@example.com" || firstName != "Alice" {
			t.Fatalf("expected email/first_name to be untouched, got email=%q firstName=%q", email, firstName)
		}
		if !balance.Valid || balance.Int64 != 1234 {
			t.Fatalf("expected balance_cents to be updated to 1234, got %+v", balance)
		}
	})

	// ── 8. deleteUserById: RETURNING on a DELETE. ────────────────────────
	t.Run("deleteUserById", func(t *testing.T) {
		params := external.DeleteUserByIdVsqlParams{Id: aliceID}
		columns := external.NewDeleteUserByIdVsqlColumns()

		query, args := external.PrepareDeleteUserByIdVsql(params, columns)
		sqlText := render(t, query, args)

		var row external.DeleteUserByIdVsqlRow
		if err := db.QueryRow(sqlText).Scan(&row.Id, &row.Email); err != nil {
			t.Fatalf("scan: %v\nsql:\n%s", err, sqlText)
		}
		if row.Email != "alice@example.com" {
			t.Fatalf("unexpected deleted row: %+v", row)
		}

		var remaining int
		if err := db.QueryRow("SELECT count(*) FROM users WHERE id = ?", aliceID).Scan(&remaining); err != nil {
			t.Fatalf("count: %v", err)
		}
		if remaining != 0 {
			t.Fatalf("expected the row to actually be gone, found %d", remaining)
		}
	})

	// ── 9. countUsersByRole: aggregate, no columns, one optional filter. ─
	t.Run("countUsersByRole_filtered", func(t *testing.T) {
		params := external.CountUsersByRoleVsqlParams{Role: emigo.NullableOf("Member")}
		query, args := external.PrepareCountUsersByRoleVsql(params)
		sqlText := render(t, query, args)

		var total int
		if err := db.QueryRow(sqlText).Scan(&total); err != nil {
			t.Fatalf("scan: %v\nsql:\n%s", err, sqlText)
		}
		// bob is Member; carol/dave never got a role (NULL); alice was
		// deleted in the previous subtest.
		if total != 1 {
			t.Fatalf("expected exactly 1 Member, got %d", total)
		}
	})

	t.Run("countUsersByRole_unfiltered", func(t *testing.T) {
		query, args := external.PrepareCountUsersByRoleVsql(external.CountUsersByRoleVsqlParams{})
		sqlText := render(t, query, args)

		var total int
		if err := db.QueryRow(sqlText).Scan(&total); err != nil {
			t.Fatalf("scan: %v\nsql:\n%s", err, sqlText)
		}
		// bob, carol, dave remain (alice was deleted).
		if total != 3 {
			t.Fatalf("expected 3 remaining users with no filter, got %d", total)
		}
	})

	// ── 10. selectUsersPaged: queryName - SQL read from an embed.FS. ────
	t.Run("selectUsersPaged", func(t *testing.T) {
		params := external.SelectUsersPagedVsqlParams{Limit: 10, Offset: 0}
		columns := external.NewSelectUsersPagedVsqlColumns()

		query, args, err := external.PrepareSelectUsersPagedVsql(sqlfiles.Files, params, columns)
		if err != nil {
			t.Fatalf("PrepareSelectUsersPagedVsql: %v", err)
		}
		sqlText := render(t, query, args)

		rows, err := db.Query(sqlText)
		if err != nil {
			t.Fatalf("query: %v\nsql:\n%s", err, sqlText)
		}
		defer rows.Close()

		count := 0
		for rows.Next() {
			var id int64
			var email string
			if err := rows.Scan(&id, &email); err != nil {
				t.Fatalf("scan: %v", err)
			}
			count++
		}
		if count != 3 {
			t.Fatalf("expected the 3 remaining users, got %d", count)
		}
	})
}

func containsQuoted(s, substr string) bool {
	return strings.Contains(s, "'"+substr+"'")
}
