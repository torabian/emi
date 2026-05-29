//go:build wasm

package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"syscall/js"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBError struct {
	Message  string
	Code     string
	Severity string
}

func (e *DBError) Error() string {
	return e.Message
}

type WasmQueryResult struct {
	Rows         []map[string]any `json:"rows"`
	Fields       []WasmField      `json:"fields"`
	AffectedRows int64            `json:"affectedRows"`
	Error        *DBError         `json:"error"`
}

type WasmField struct {
	Name       string `json:"name"`
	DataTypeID int    `json:"dataTypeID"`
}

type WasmDatabase struct {
	queryFunc js.Value
}

type Row interface {
	Scan(dest ...any) error
}

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close()
	Err() error
}

func NewWasmDatabase(
	queryFunc js.Value,
) *WasmDatabase {

	return &WasmDatabase{
		queryFunc: queryFunc,
	}
}

type wasmCopyFromSource struct {
	columns []string
	rows    [][]any
	index   int
}

func (s *wasmCopyFromSource) Next() bool {
	return s.index < len(s.rows)
}
func (s *wasmCopyFromSource) Values() ([]any, error) {
	row := s.rows[s.index]
	s.index++
	return row, nil
}
func (s *wasmCopyFromSource) Err() error {
	return nil
}
func (s *wasmCopyFromSource) ColumnNames() []string {
	return s.columns
}

func (db *WasmDatabase) CopyFrom(
	ctx context.Context,
	table pgx.Identifier,
	columns []string,
	src pgx.CopyFromSource,
) (int64, error) {

	var rows [][]any

	for src.Next() {

		values, err := src.Values()
		if err != nil {
			return 0, err
		}

		rows = append(rows, values)
	}

	payload := map[string]any{
		"table":   table,
		"columns": columns,
		"rows":    rows,
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	promise := db.queryFunc.Invoke(
		"COPY",
		string(jsonBytes),
	)

	resultBytes, err := awaitPromise(promise)
	if err != nil {
		return 0, err
	}

	var result struct {
		AffectedRows int64    `json:"affectedRows"`
		Error        *DBError `json:"error"`
	}

	if err := json.Unmarshal(resultBytes, &result); err != nil {
		return 0, err
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return result.AffectedRows, nil
}
func (r *WasmRows) CommandTag() pgconn.CommandTag {
	return pgconn.CommandTag{}
}
func (r *WasmRows) Conn() *pgx.Conn {
	// fake connection — NEVER used in WASM logic
	return &pgx.Conn{}
}
func (r *WasmRows) FieldDescriptions() []pgconn.FieldDescription {
	return nil
}
func (r *WasmRows) RawValues() [][]byte {
	return nil
}
func (r *WasmRows) Values() ([]any, error) {
	if r.index >= len(r.rows) {
		return nil, fmt.Errorf("out of rows")
	}

	row := r.rows[r.index]

	values := make([]any, 0, len(row))
	for _, v := range row {
		values = append(values, v)
	}

	return values, nil
}
func (db *WasmDatabase) Query(
	ctx context.Context,
	query string,
	args ...any,
) (pgx.Rows, error) {

	result, err := db.executeQuery(
		query,
		args...,
	)

	if err != nil {
		return nil, err
	}

	return &WasmRows{
		rows:   result.Rows,
		fields: result.Fields,
	}, nil
}

type sliceCopySource struct {
	columns []string
	rows    [][]any
	index   int
}

func (db *WasmDatabase) QueryRow(
	ctx context.Context,
	query string,
	args ...any,
) pgx.Row {

	result, err := db.executeQuery(
		query,
		args...,
	)

	if err != nil {
		return &WasmRow{
			err: err,
		}
	}

	if len(result.Rows) == 0 {
		return &WasmRow{
			err: errors.New("no rows"),
		}
	}

	return &WasmRow{
		data:   result.Rows[0],
		fields: result.Fields,
	}
}

func (db *WasmDatabase) Exec(
	ctx context.Context,
	query string,
	args ...any,
) (pgconn.CommandTag, error) {

	result, err := db.executeQuery(
		query,
		args...,
	)
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	// if JS returned DB error
	if result.Error != nil {
		return pgconn.CommandTag{}, result.Error
	}

	// build Postgres-like command tag
	tag := pgconn.NewCommandTag(
		fmt.Sprintf("EXEC %d", result.AffectedRows),
	)

	return tag, nil
}

func (db *WasmDatabase) executeQuery(
	query string,
	args ...any,
) (*WasmQueryResult, error) {

	promise := db.queryFunc.Invoke(
		query,
		js.ValueOf(args),
	)

	jsonBytes, err := awaitPromise(promise)
	if err != nil {
		return nil, err
	}

	var result WasmQueryResult

	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}

	// 🔥 JS-level DB error (Postgres-like)
	if result.Error != nil {
		return nil, result.Error
	}

	return &result, nil
}

type WasmRow struct {
	data   map[string]any
	fields []WasmField
	err    error
}

func (r *WasmRow) Scan(dest ...any) error {

	if r.err != nil {
		return r.err
	}

	return scanMap(
		r.data,
		r.fields,
		dest,
	)
}

//
// ROWS
//

type WasmRows struct {
	rows   []map[string]any
	fields []WasmField
	index  int
}

func (r *WasmRows) Next() bool {
	return r.index < len(r.rows)
}

func (r *WasmRows) Scan(dest ...any) error {

	err := scanMap(
		r.rows[r.index],
		r.fields,
		dest,
	)

	r.index++

	return err
}

func (r *WasmRows) Close() {}

func (r *WasmRows) Err() error {
	return nil
}

//
// SCAN
//

func scanMap(
	row map[string]any,
	fields []WasmField,
	dest []any,
) error {

	// Scan in column order. Go map iteration is randomized, so ranging over
	// `row` directly would assign values to the wrong dest pointers as soon as
	// there is more than one column. The field metadata preserves the SELECT
	// order; fall back to map iteration only when it is unavailable.
	names := make([]string, 0, len(row))
	if len(fields) > 0 {
		for _, f := range fields {
			names = append(names, f.Name)
		}
	} else {
		for name := range row {
			names = append(names, name)
		}
	}

	for i, name := range names {

		if i >= len(dest) {
			break
		}

		value := row[name]

		switch d := dest[i].(type) {

		case *string:
			if v, ok := value.(string); ok {
				*d = v
			}

		case *int:
			switch v := value.(type) {
			case float64:
				*d = int(v)
			case int:
				*d = v
			}

		case *int64:
			switch v := value.(type) {
			case float64:
				*d = int64(v)
			case int64:
				*d = v
			}
		}
	}

	return nil
}

func awaitPromise(promise js.Value) ([]byte, error) {

	resultCh := make(chan []byte)

	then := js.FuncOf(func(this js.Value, args []js.Value) any {
		resultCh <- []byte(args[0].String())
		return nil
	})

	defer then.Release()

	promise.Call("then", then)

	// IMPORTANT: no catch at all

	select {
	case res := <-resultCh:
		return res, nil
	}
}
