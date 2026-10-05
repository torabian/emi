package entitysql

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

const sample = `name: shop
entities:
  - name: category
    table: categories_custom
    fields:
      - name: title
        type: string
  - name: orderItem
    fields:
      - name: title
        type: string
      - name: quantity
        type: int?
      - name: paid
        type: bool
        default: true
      - name: status
        type: enum
        of:
          - k: open
          - k: closed
      - name: category
        type: one
        target: CategoryEntity
      - name: owner
        type: one?
        target: UserEntity
      - name: address
        type: object
        fields:
          - name: street
            type: string
          - name: geo
            type: object?
            fields:
              - name: lat
                type: float64
      - name: lines
        type: array
        fields:
          - name: sku
            type: string
          - name: notes
            type: array?
            fields:
              - name: text
                type: string
      - name: tags
        type: collection
        target: TagEntity
      - name: related
        type: collection?
        target: OrderItemEntity
`

func render(t *testing.T) string {
	t.Helper()
	m, err := core.StringToEmi(sample)
	if err != nil {
		t.Fatal(err)
	}
	files, err := ModuleToPostgres(m)
	if err != nil {
		t.Fatal(err)
	}
	return files[0].ActualScript
}

func TestModuleToPostgres(t *testing.T) {
	sql := render(t)
	for _, want := range []string{
		// scalar
		`CREATE TABLE IF NOT EXISTS "order_item_entities"`,
		`ADD COLUMN IF NOT EXISTS "title" text NOT NULL DEFAULT '';`,
		`ADD COLUMN IF NOT EXISTS "quantity" integer;`,
		`ADD COLUMN IF NOT EXISTS "paid" boolean NOT NULL DEFAULT true;`,
		`CHECK ("status" IN ('open', 'closed'))`,
		// one: real fk to the custom named table; one?: column and index only
		`ADD COLUMN IF NOT EXISTS "category_id" bigint;`,
		`FOREIGN KEY ("category_id") REFERENCES "categories_custom" ("id")`,
		`ADD COLUMN IF NOT EXISTS "owner_id" bigint;`,
		`CREATE INDEX IF NOT EXISTS "idx_order_item_entities_owner_id"`,
		// object flattened, nested object? nullable
		`ADD COLUMN IF NOT EXISTS "address_street" text NOT NULL DEFAULT '';`,
		`ADD COLUMN IF NOT EXISTS "address_geo_lat" double precision;`,
		// array: child tables, recursive
		`CREATE TABLE IF NOT EXISTS "order_item_entity_lines"`,
		`ALTER TABLE "order_item_entity_lines" ADD COLUMN IF NOT EXISTS "linker_id" bigint;`,
		`FOREIGN KEY ("linker_id") REFERENCES "order_item_entities" ("id") ON DELETE CASCADE`,
		`CREATE TABLE IF NOT EXISTS "order_item_entity_lines_notes"`,
		`FOREIGN KEY ("linker_id") REFERENCES "order_item_entity_lines" ("id") ON DELETE CASCADE`,
		// collection: join tables, self reference
		`CREATE TABLE IF NOT EXISTS "orderItem_tags"`,
		`"order_item_entity_id" bigint NOT NULL`,
		`"tag_entity_id" bigint NOT NULL`,
		`REFERENCES "tag_entities" ("id") ON DELETE CASCADE`,
		`"reference_order_item_entity_id" bigint NOT NULL`,
	} {
		if !strings.Contains(sql, want) {
			t.Errorf("missing %q in:\n%s", want, sql)
		}
	}
	if strings.Contains(sql, `"owner_id") REFERENCES`) {
		t.Errorf("one? must not have a foreign key constraint:\n%s", sql)
	}
}

func TestPhasesOrder(t *testing.T) {
	sql := render(t)
	last := -1
	for _, title := range []string{"-- Tables", "-- Columns", "-- Indexes", "-- Constraints"} {
		i := strings.Index(sql, title)
		if i <= last {
			t.Fatalf("%s out of order", title)
		}
		last = i
	}
}

func TestDeterministic(t *testing.T) {
	if render(t) != render(t) {
		t.Fatal("output differs between runs")
	}
}

const dynamicSample = `name: x
entities:
  - name: formdata
    table: formdata_abc
    fields:
      - name: lastName
        type: string?
        tags:
          json: last name
      - name: age
        type: int64?
      - name: score
        type: float64?
      - name: agreed
        type: bool?
      - name: born
        type: complex
        complex: XDate
      - name: seen
        type: complex
        complex: XDateTime
      - name: tags
        type: slice
        primitive: string
      - name: address
        type: object?
        fields:
          - name: street
            type: string?
          - name: geo
            type: object?
            fields:
              - name: lat
                type: float64?
      - name: phones
        type: array?
        fields:
          - name: number
            type: string?
          - name: notes
            type: array?
            fields:
              - name: text
                type: string?
      - name: workspaceId
        type: string?
        tags:
          json: "-"
      - name: createdAt
        type: complex
        complex: PlainTime
        tags:
          json: "-"
`

func dynamicLayout(t *testing.T) (*Layout, []string) {
	t.Helper()
	m, err := core.StringToEmi(dynamicSample)
	if err != nil {
		t.Fatal(err)
	}
	statements, layouts, err := PostgresStatements(m.Entities, Options{ChildTablesFromParent: true})
	if err != nil {
		t.Fatal(err)
	}
	return layouts[0], statements
}

func TestLayoutNamesTheJsonPathOfEveryColumn(t *testing.T) {
	layout, _ := dynamicLayout(t)

	if layout.Table != "formdata_abc" {
		t.Fatalf("table %q", layout.Table)
	}
	got := map[string]string{}
	for _, c := range layout.Columns {
		got[strings.Join(c.Path, ".")] = c.Column + " " + c.SqlType + " " + string(c.Kind)
	}
	for path, want := range map[string]string{
		"last name":       "last_name text string",
		"age":             "age bigint int",
		"score":           "score double precision float",
		"agreed":          "agreed boolean bool",
		"born":            "born date date",
		"seen":            "seen timestamptz time",
		"tags":            "tags jsonb json",
		"address.street":  "address_street text string",
		"address.geo.lat": "address_geo_lat double precision float",
	} {
		if got[path] != want {
			t.Errorf("path %q: got %q, want %q", path, got[path], want)
		}
	}
	// fields tagged json:"-" have a column, but no place in the layout
	for path := range got {
		if path == "workspaceId" || path == "createdAt" || path == "" {
			t.Errorf("hidden field in the layout: %q", path)
		}
	}
}

func TestLayoutChildrenAreTheirOwnTables(t *testing.T) {
	layout, statements := dynamicLayout(t)

	if len(layout.Children) != 1 || strings.Join(layout.Children[0].Path, ".") != "phones" {
		t.Fatalf("children: %+v", layout.Children)
	}
	phones := layout.Children[0].Layout
	if phones.Table != "formdata_abc_phones" {
		t.Fatalf("child table named after its parent, got %q", phones.Table)
	}
	if len(phones.Columns) != 1 || phones.Columns[0].Column != "number" {
		t.Fatalf("child columns: %+v", phones.Columns)
	}
	// and the array inside the array, one level deeper
	if len(phones.Children) != 1 || phones.Children[0].Layout.Table != "formdata_abc_phones_notes" {
		t.Fatalf("grandchildren: %+v", phones.Children)
	}

	all := strings.Join(statements, "\n")
	for _, want := range []string{
		`CREATE TABLE IF NOT EXISTS "formdata_abc_phones"`,
		`ADD COLUMN IF NOT EXISTS "created_at" timestamptz;`,
		`ADD COLUMN IF NOT EXISTS "workspace_id" text;`,
		`ADD COLUMN IF NOT EXISTS "born" date;`,
		`ADD COLUMN IF NOT EXISTS "last_name" text;`,
	} {
		if !strings.Contains(all, want) {
			t.Errorf("missing %q in\n%s", want, all)
		}
	}
}

func TestStatementsRunOneByOneInPhaseOrder(t *testing.T) {
	_, statements := dynamicLayout(t)
	phase := func(s string) int {
		switch {
		case strings.HasPrefix(s, "CREATE TABLE"):
			return 0
		case strings.HasPrefix(s, "ALTER TABLE"):
			return 1
		case strings.HasPrefix(s, "CREATE INDEX"):
			return 2
		}
		return 3
	}
	last := 0
	for _, s := range statements {
		if p := phase(s); p < last {
			t.Fatalf("statement out of phase order: %s", s)
		} else {
			last = p
		}
	}
}
