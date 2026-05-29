package internal

// Manual CRUD implementation for the user actions. These functions are the
// shared business logic behind the generated emi handlers (CreateUserAction,
// ListUsersAction, DeleteUserAction) and are wired into the wasm net/http
// router in cmd/wasm/main.go. They talk to Postgres exclusively through the
// Database interface, so the same code runs against a real *pgx.Conn or the
// in-browser pglite-backed WasmDatabase.

import (
	"context"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/examples/in-browser-server/internal/defs"
)

// UserModule bundles the database handle so the action handlers can be
// expressed as methods with the exact signatures emi expects.
type UserModule struct {
	DB Database
}

// EnsureSchema creates the users table if it does not already exist. birth_date
// is stored as TEXT so the ISO "YYYY-MM-DD" value round-trips verbatim through
// the JS bridge (a real DATE column would come back as a timestamp string).
func (m *UserModule) EnsureSchema(ctx context.Context) error {
	_, err := m.DB.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id         SERIAL PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name  TEXT NOT NULL,
			birth_date TEXT NOT NULL
		)
	`)
	return err
}

// Create inserts a new user and returns the freshly assigned id alongside the
// submitted fields.
func (m *UserModule) Create(c defs.CreateUserActionRequest) (*defs.CreateUserActionResponse, error) {
	ctx := context.Background()

	var id int
	row := m.DB.QueryRow(ctx,
		`INSERT INTO users (first_name, last_name, birth_date)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		c.Body.FirstName, c.Body.LastName, c.Body.BirthDate,
	)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &defs.CreateUserActionResponse{
		Payload: defs.CreateUserActionRes{
			Id:        id,
			FirstName: c.Body.FirstName,
			LastName:  c.Body.LastName,
			BirthDate: c.Body.BirthDate,
		},
	}, nil
}

// List returns every user row ordered by id.
func (m *UserModule) List(c defs.ListUsersActionRequest) (*defs.ListUsersActionResponse, error) {
	ctx := context.Background()

	rows, err := m.DB.Query(ctx,
		`SELECT id, first_name, last_name, birth_date FROM users ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []defs.ListUsersActionResUsers{}
	for rows.Next() {
		var u defs.ListUsersActionResUsers
		if err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.BirthDate); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &defs.ListUsersActionResponse{
		Payload: defs.ListUsersActionRes{
			Users: emigo.NewArray(users...),
		},
	}, nil
}

// Delete removes a single user by id and reports whether a row was affected.
func (m *UserModule) Delete(c defs.DeleteUserActionRequest) (*defs.DeleteUserActionResponse, error) {
	ctx := context.Background()

	tag, err := m.DB.Exec(ctx, `DELETE FROM users WHERE id = $1`, c.Body.Id)
	if err != nil {
		return nil, err
	}

	return &defs.DeleteUserActionResponse{
		Payload: defs.DeleteUserActionRes{
			Deleted: tag.RowsAffected() > 0,
		},
	}, nil
}
