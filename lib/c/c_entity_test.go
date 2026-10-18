package c

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

const entityYml = `name: shop
entities:
  - name: order
    fields:
      - name: title
        type: string
      - name: quantity
        type: int?
      - name: price
        type: float64
      - name: paid
        type: bool
        default: true
      - name: status
        type: enum
        of:
          - k: open
          - k: closed
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
`

const entityHostMain = `
#include <stdio.h>
#include <string.h>
#include "order_entity.h"
#include "shop_schema.h"

#define CHECK(cond) do { if (!(cond)) { printf("FAILED line %d: %s\n", __LINE__, #cond); return 1; } } while (0)

int main(void) {
    sqlite3* db;
    CHECK(sqlite3_open(":memory:", &db) == SQLITE_OK);
    CHECK(shop_schema_apply(db) == 0);
    CHECK(shop_schema_apply(db) == 0); /* the schema can be applied on every start */

    /* insert: scalars, a flattened object, an array with a nested array */
    cJSON* in = cJSON_Parse("{\"title\":\"first\",\"quantity\":3,\"price\":9.5,\"paid\":false,\"status\":\"open\","
        "\"address\":{\"street\":\"Main 1\",\"geo\":{\"lat\":52.25}},"
        "\"lines\":[{\"sku\":\"a\",\"notes\":[{\"text\":\"n1\"},{\"text\":\"n2\"}]},{\"sku\":\"b\"}]}");
    OrderEntity* o = OrderEntity_from_json(in);
    cJSON_Delete(in);
    int64_t id = 0;
    CHECK(OrderEntity_insert(db, o, &id) == 0 && id == 1);
    OrderEntity_free(o);

    /* the rows are where the sql says: a flattened object is columns, an array is a table */
    sqlite3_stmt* s;
    CHECK(sqlite3_prepare_v2(db, "SELECT address_street, address_geo_lat, paid, quantity FROM order_entities WHERE id = 1", -1, &s, NULL) == SQLITE_OK);
    CHECK(sqlite3_step(s) == SQLITE_ROW);
    CHECK(strcmp((const char*) sqlite3_column_text(s, 0), "Main 1") == 0);
    CHECK(sqlite3_column_double(s, 1) == 52.25);
    CHECK(sqlite3_column_int(s, 2) == 0);
    CHECK(sqlite3_column_int(s, 3) == 3);
    sqlite3_finalize(s);
    CHECK(sqlite3_prepare_v2(db, "SELECT count(*) FROM order_entity_lines", -1, &s, NULL) == SQLITE_OK);
    CHECK(sqlite3_step(s) == SQLITE_ROW && sqlite3_column_int(s, 0) == 2);
    sqlite3_finalize(s);

    /* get: back to the struct, arrays in order */
    OrderEntity* got = OrderEntity_get(db, id);
    CHECK(got != NULL);
    CHECK(strcmp(got->title, "first") == 0 && got->quantity == 3 && got->price == 9.5 && got->paid == false);
    CHECK(strcmp(got->address->street, "Main 1") == 0 && got->address->geo && got->address->geo->lat == 52.25);
    cJSON* out = OrderEntity_to_json(got);
    char* text = cJSON_PrintUnformatted(out);
    CHECK(strstr(text, "\"sku\":\"a\"") && strstr(text, "\"text\":\"n2\"") && strstr(text, "\"sku\":\"b\""));
    free(text);
    cJSON_Delete(out);

    /* update rewrites the arrays */
    free(got->title);
    got->title = strdup("renamed");
    got->paid = true;
    CHECK(OrderEntity_update(db, id, got) == 0);
    OrderEntity_free(got);
    got = OrderEntity_get(db, id);
    CHECK(strcmp(got->title, "renamed") == 0 && got->paid == true);
    OrderEntity_free(got);
    CHECK(sqlite3_prepare_v2(db, "SELECT count(*) FROM order_entity_lines", -1, &s, NULL) == SQLITE_OK);
    CHECK(sqlite3_step(s) == SQLITE_ROW && sqlite3_column_int(s, 0) == 2); /* rewritten, not doubled */
    sqlite3_finalize(s);

    /* list, limit and offset */
    OrderEntity* second = OrderEntity_new();
    free(second->title);
    second->title = strdup("second");
    int64_t id2 = 0;
    CHECK(OrderEntity_insert(db, second, &id2) == 0 && id2 == 2);
    OrderEntity_free(second);
    OrderEntityList* all = OrderEntity_list(db, -1, 0);
    CHECK(all && all->count == 2 && all->ids[1] == 2 && strcmp(all->items[1]->title, "second") == 0);
    OrderEntityList_free(all);
    OrderEntityList* page = OrderEntity_list(db, 1, 1);
    CHECK(page && page->count == 1 && page->ids[0] == 2);
    OrderEntityList_free(page);

    /* query: filter (json-logic), sort, paging, total */
    EmiSqlQuery q;
    OrderEntityList* r;
    memset(&q, 0, sizeof(q));
    q.filter = "{\"==\":[{\"var\":\"title\"},\"second\"]}";
    r = OrderEntity_query(db, &q);
    CHECK(r && r->count == 1 && r->ids[0] == 2 && r->total == 1);
    OrderEntityList_free(r);
    q.filter = "{\"and\":[{\">\":[{\"var\":\"price\"},5]},{\"==\":[{\"var\":\"paid\"},true]}]}";
    r = OrderEntity_query(db, &q);
    CHECK(r && r->count == 1 && r->ids[0] == 1);
    OrderEntityList_free(r);
    q.filter = "{\"contains\":[{\"var\":\"title\"},\"SEC\"]}"; /* case does not matter */
    r = OrderEntity_query(db, &q);
    CHECK(r && r->count == 1 && r->ids[0] == 2);
    OrderEntityList_free(r);
    q.filter = "{\"==\":[{\"var\":\"address.street\"},\"Main 1\"]}"; /* a flattened object */
    r = OrderEntity_query(db, &q);
    CHECK(r && r->count == 1 && r->ids[0] == 1);
    OrderEntityList_free(r);
    q.filter = "{\"in\":[{\"var\":\"title\"},[\"second\",\"zzz\"]]}";
    r = OrderEntity_query(db, &q);
    CHECK(r && r->count == 1 && r->ids[0] == 2);
    OrderEntityList_free(r);
    q.filter = "{\"or\":[{\"!\":{\"==\":[{\"var\":\"paid\"},true]}},{\"==\":[{\"var\":\"title\"},\"nothing\"]}]}";
    r = OrderEntity_query(db, &q);
    CHECK(r && r->count == 1 && r->ids[0] == 2);
    OrderEntityList_free(r);
    q.filter = "{\"==\":[{\"var\":\"title\"},\"x' OR '1'='1\"]}"; /* a value is never sql */
    r = OrderEntity_query(db, &q);
    CHECK(r && r->count == 0 && r->total == 0);
    OrderEntityList_free(r);
    q.filter = NULL;
    q.sort = "title desc";
    r = OrderEntity_query(db, &q);
    CHECK(r && r->count == 2 && r->ids[0] == 2 && r->ids[1] == 1);
    OrderEntityList_free(r);
    q.sort = "title";
    q.start_index = 1;
    q.items_per_page = 1;
    r = OrderEntity_query(db, &q);
    CHECK(r && r->count == 1 && r->ids[0] == 2 && r->total == 2); /* total ignores the paging */
    OrderEntityList_free(r);
    memset(&q, 0, sizeof(q));
    q.filter = "{\"==\":[{\"var\":\"nope\"},1]}"; /* an unknown field */
    CHECK(OrderEntity_query(db, &q) == NULL);
    q.filter = "not json";
    CHECK(OrderEntity_query(db, &q) == NULL);
    q.filter = "{\"=\":[1,1]}";
    CHECK(OrderEntity_query(db, &q) == NULL);
    q.filter = NULL;
    q.sort = "title; DROP TABLE order_entities";
    CHECK(OrderEntity_query(db, &q) == NULL);
    r = OrderEntity_query(db, NULL); /* no query: everything */
    CHECK(r && r->count == 2 && r->total == 2);
    OrderEntityList_free(r);

    /* patch: only the fields given change, an object is merged, an array is replaced */
    cJSON* patch = cJSON_Parse("{\"title\":\"patched\",\"address\":{\"street\":\"Side 2\"}}");
    CHECK(OrderEntity_patch(db, 1, patch) == 0);
    cJSON_Delete(patch);
    got = OrderEntity_get(db, 1);
    CHECK(strcmp(got->title, "patched") == 0 && strcmp(got->address->street, "Side 2") == 0);
    CHECK(got->address->geo && got->address->geo->lat == 52.25 && got->price == 9.5 && got->paid == true);
    OrderEntity_free(got);
    CHECK(sqlite3_prepare_v2(db, "SELECT count(*) FROM order_entity_lines", -1, &s, NULL) == SQLITE_OK);
    CHECK(sqlite3_step(s) == SQLITE_ROW && sqlite3_column_int(s, 0) == 2); /* the lines were not touched */
    sqlite3_finalize(s);
    patch = cJSON_Parse("{\"lines\":[{\"sku\":\"z\"}]}");
    CHECK(OrderEntity_patch(db, 1, patch) == 0);
    cJSON_Delete(patch);
    CHECK(sqlite3_prepare_v2(db, "SELECT count(*) FROM order_entity_lines", -1, &s, NULL) == SQLITE_OK);
    CHECK(sqlite3_step(s) == SQLITE_ROW && sqlite3_column_int(s, 0) == 1);
    sqlite3_finalize(s);
    patch = cJSON_Parse("{\"title\":\"x\"}");
    CHECK(OrderEntity_patch(db, 99, patch) == SQLITE_NOTFOUND);
    cJSON_Delete(patch);
    got = OrderEntity_get(db, 2);
    CHECK(strcmp(got->title, "second") == 0); /* the other row is as it was */
    OrderEntity_free(got);

    /* delete takes the arrays with it */
    CHECK(OrderEntity_delete(db, id) == 0);
    CHECK(OrderEntity_get(db, id) == NULL);
    CHECK(OrderEntity_delete(db, id) == SQLITE_NOTFOUND);
    CHECK(OrderEntity_update(db, 99, OrderEntity_new()) == SQLITE_NOTFOUND);
    CHECK(sqlite3_prepare_v2(db, "SELECT count(*) FROM order_entity_lines", -1, &s, NULL) == SQLITE_OK);
    CHECK(sqlite3_step(s) == SQLITE_ROW && sqlite3_column_int(s, 0) == 0);
    sqlite3_finalize(s);

    sqlite3_close(db);
    printf("ok\n");
    return 0;
}
`

// TestSqliteEntities compiles the entity code generated with --tags sqlite
// against the system SQLite and drives it: schema (twice), insert, get with a
// flattened object and nested arrays, update, list, query (filter, sort, paging), patch, delete.
func TestSqliteEntities(t *testing.T) {
	cc, err := exec.LookPath("cc")
	if err != nil {
		t.Skip("no C compiler available")
	}
	module, err := core.StringToEmi(entityYml)
	if err != nil {
		t.Fatal(err)
	}
	ctx := core.MicroGenContext{Content: entityYml, Flags: map[string]string{}, Tags: string(Sqlite)}
	files, err := CModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	for _, f := range files {
		full := filepath.Join(dir, f.Location, f.Name+f.Extension)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(f.ActualScript), 0644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(dir, "main.c"), []byte(entityHostMain), 0644); err != nil {
		t.Fatal(err)
	}
	bin := filepath.Join(dir, "ent")
	out, err := exec.Command(cc, "-std=c99", "-D_POSIX_C_SOURCE=200809L", "-Wall", "-Wextra",
		"-I", dir, filepath.Join(dir, "main.c"), filepath.Join(dir, "cjson", "cJSON.c"), "-o", bin, "-lsqlite3", "-lm").CombinedOutput()
	if err != nil {
		if strings.Contains(string(out), "sqlite3.h") || strings.Contains(string(out), "lsqlite3") {
			t.Skipf("no sqlite3 development files: %s", out)
		}
		t.Fatalf("generated entity C failed to compile:\n%s", out)
	}
	if out, err := exec.Command(bin).CombinedOutput(); err != nil {
		t.Fatalf("entity checks failed: %v\n%s", err, out)
	}
}
