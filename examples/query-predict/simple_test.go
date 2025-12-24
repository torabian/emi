package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/h22rana/jsonlogic2sql"
	_ "github.com/mattn/go-sqlite3"
	xxx "github.com/torabian/emi/examples/query-predict/output"
)

// func TestSimple(t *testing.T) {
// 	// Open in-memory SQLite database
// 	db, err := sql.Open("sqlite3", ":memory:")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Create table and insert test data
// 	db.Exec(`CREATE TABLE users (id INTEGER, name TEXT, email TEXT)`)
// 	db.Exec(`INSERT INTO users (id, name, email) VALUES (1, 'Alice', 'alice@example.com')`)
// 	db.Exec(`INSERT INTO users (id, name, email) VALUES (2, 'Bob', 'bob@example.com')`)

// 	// Run the generated query
// 	users, err := artifacts.GetUsers(db, xxx{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, u := range users {
// 		t.Logf("%+v", u) // logs appear only in verbose mode
// 	}
// }

// func TestUsersWithOrders(t *testing.T) {
// 	db, err := sql.Open("sqlite3", ":memory:")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Create tables
// 	db.Exec(`CREATE TABLE users (id INTEGER, name TEXT, email TEXT)`)
// 	db.Exec(`CREATE TABLE orders (id INTEGER, user_id INTEGER, total REAL)`)

// 	// Insert test data
// 	db.Exec(`INSERT INTO users (id, name, email) VALUES (1, 'Alice', 'alice@example.com')`)
// 	db.Exec(`INSERT INTO users (id, name, email) VALUES (2, 'Bob', 'bob@example.com')`)
// 	db.Exec(`INSERT INTO orders (id, user_id, total) VALUES (1, 1, 100.5)`)
// 	db.Exec(`INSERT INTO orders (id, user_id, total) VALUES (2, 1, 50.0)`)

// 	m := xxx.Manifest{
// 		DB: db,
// 		FilterResolver: func(s string) (string, error) {
// 			result, err := jsonlogic2sql.Transpile(s)
// 			if err != nil {
// 				return "", err
// 			}

// 			result = result[6:]
// 			return result, nil
// 		},
// 	}

// 	rows, err := m.GetUsersWithOrders(artifacts.GetUsersWithOrdersContext{
// 		Filter: `{"and": [{">": [{"var": "order_id"}, 1]}]}`,
// 	})

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	var results []xxx.GetUsersWithOrdersRow
// 	for _, r := range rows {
// 		results = append(results, r)
// 	}

// 	// Print results
// 	for _, r := range results {
// 		t.Logf("%+v", r)
// 	}
// }

// func TestVirtualTableWhereHaving(t *testing.T) {
// 	db, err := sql.Open("sqlite3", ":memory:")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()

// 	m := xxx.Manifest{
// 		DB: db,
// 		FilterResolver: func(s string) (string, error) {
// 			result, err := jsonlogic2sql.Transpile(s)
// 			if err != nil {
// 				return "", err
// 			}

// 			result = result[6:]
// 			return result, nil
// 		},
// 	}

// 	rows, err := m.VirtualUserOrder(xxx.VirtualUserOrderContext{
// 		Filter: `{"and": [{"!=": [{"var": "u.user_name"}, "Alicex"]}]}`,
// 		Params: map[string]interface{}{
// 			"limit": 2,
// 		},
// 	})

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	var results []xxx.VirtualUserOrderRow
// 	for _, r := range rows {
// 		results = append(results, r)
// 	}

// 	// Print results
// 	for _, r := range results {
// 		t.Logf("%+v", r)
// 	}
// }

func TestUserInsertion(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	m := xxx.Manifest{
		DB: db,
		FilterResolver: func(s string) (string, error) {
			result, err := jsonlogic2sql.Transpile(s)
			if err != nil {
				return "", err
			}

			result = result[6:]
			return result, nil
		},
	}

	// Let's create the table first.
	if _, err := xxx.CreateUserTable(db, xxx.CreateUserTableContext{}); err != nil {
		fmt.Printf("%v", err)
		return
	}

	/// Let's insert 2 users into it.
	if _, err := xxx.Transaction(db, xxx.TransactionContext{}); err != nil {
		fmt.Printf("%v", err)
		return
	}

	// Let's query back how many users are inside.
	if result, err := xxx.CountUsers(db, xxx.CountUsersContext{}); err != nil {
		return
	} else {
		if result[0].Count != "2" {
			t.Errorf("Result count is not 2.")
		}
	}

	// Let's print all users, using manifest, so the filter resolver is automatically applied.
	if result, err := m.SelectAllUsers(xxx.SelectAllUsersContext{
		Filter: `{"and": [{"==": [{"var": "id"}, 1]}]}`,
	}); err != nil {
		return
	} else {
		data, _ := json.MarshalIndent(result, "", "  ")

		fmt.Println(5, string(data))
	}

}
