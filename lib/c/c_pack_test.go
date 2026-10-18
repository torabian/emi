package c

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// buildPackProgram generates the entities of yml with the given tags, compiles main.c next to them
// (with the address and undefined behaviour sanitizers) and returns the binary. It never links SQLite:
// the pack tag has to work without it.
func buildPackProgram(t *testing.T, yml, tags, main string) (bin string, generated map[string]string) {
	t.Helper()
	cc, err := exec.LookPath("cc")
	if err != nil {
		t.Skip("no C compiler available")
	}
	module, err := core.StringToEmi(yml)
	if err != nil {
		t.Fatal(err)
	}
	ctx := core.MicroGenContext{Content: yml, Flags: map[string]string{}, Tags: tags}
	files, err := CModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	generated = map[string]string{}
	for _, f := range files {
		full := filepath.Join(dir, f.Location, f.Name+f.Extension)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(f.ActualScript), 0644); err != nil {
			t.Fatal(err)
		}
		generated[f.Name+f.Extension] = f.ActualScript
	}
	if err := os.WriteFile(filepath.Join(dir, "main.c"), []byte(main), 0644); err != nil {
		t.Fatal(err)
	}
	bin = filepath.Join(dir, "prog")
	out, err := exec.Command(cc, "-std=c99", "-D_POSIX_C_SOURCE=200809L", "-Wall", "-Wextra", "-Werror",
		"-fsanitize=address,undefined", "-fno-sanitize-recover=undefined", "-g",
		"-I", dir, filepath.Join(dir, "main.c"), filepath.Join(dir, "cjson", "cJSON.c"), "-o", bin, "-lm").CombinedOutput()
	if err != nil {
		t.Fatalf("generated C failed to compile:\n%s", out)
	}
	return bin, generated
}

const packMain = `
#include <stdio.h>
#include <string.h>
#include "order_entity.h"

#define CHECK(cond) do { if (!(cond)) { printf("FAILED line %d: %s\n", __LINE__, #cond); return 1; } } while (0)

static char* json_of(const OrderEntity* o) {
    cJSON* j = OrderEntity_to_json(o);
    char* s = cJSON_PrintUnformatted(j);
    cJSON_Delete(j);
    return s;
}

static OrderEntity* from_text(const char* text) {
    cJSON* in = cJSON_Parse(text);
    OrderEntity* o = OrderEntity_from_json(in);
    cJSON_Delete(in);
    return o;
}

/* packs, unpacks, and checks the json of both is the same; returns the packed size or -1 */
static int roundtrip(const char* text) {
    OrderEntity* o = from_text(text);
    uint8_t* blob = NULL;
    size_t len = 0;
    if (OrderEntity_pack(o, &blob, &len) != 0) return -1;
    OrderEntity* back = OrderEntity_unpack(blob, len);
    if (!back) return -1;
    char* a = json_of(o);
    char* b = json_of(back);
    int same = strcmp(a, b) == 0;
    if (!same) printf("MISMATCH\n  %s\n  %s\n", a, b);
    free(a); free(b);
    OrderEntity_free(o);
    OrderEntity_free(back);
    free(blob);
    return same ? (int) len : -1;
}

int main(void) {
    const char* full = "{\"title\":\"first\",\"quantity\":3,\"price\":9.5,\"paid\":false,\"status\":\"open\","
        "\"address\":{\"street\":\"Main 1\",\"geo\":{\"lat\":52.25}},"
        "\"lines\":[{\"sku\":\"a\",\"notes\":[{\"text\":\"n1\"},{\"text\":\"n2\"}]},{\"sku\":\"b\"}]}";
    int n = roundtrip(full);
    CHECK(n > 0);
    printf("full entity: %d bytes (%d as json)\n", n, (int) strlen(full));
    CHECK(n < (int) strlen(full) / 2);

    CHECK(roundtrip("{}") > 0);                                           /* nothing set: defaults */
    CHECK(roundtrip("{\"paid\":true}") > 0);                              /* a bool lives in the bits */
    CHECK(roundtrip("{\"quantity\":-123456789}") > 0);                    /* negative int, zigzag */
    CHECK(roundtrip("{\"quantity\":2147483647}") > 0);
    CHECK(roundtrip("{\"price\":0.1}") > 0);                              /* not a float32: 8 bytes */
    CHECK(roundtrip("{\"price\":-1234.5678901234}") > 0);
    CHECK(roundtrip("{\"title\":\"zażółć gęślą jaźń \\u2603 \\\"q\\\" \\n\"}") > 0);
    CHECK(roundtrip("{\"lines\":[{\"sku\":\"x\"},{\"sku\":\"y\",\"notes\":[]}]}") > 0);
    CHECK(roundtrip("{\"address\":{\"geo\":{\"lat\":-0.5}}}") > 0);

    /* small stays small: a float32 is 4 bytes, a bool costs no byte */
    int small = roundtrip("{\"price\":9.5,\"paid\":true}");
    CHECK(small > 0 && small < 20);
    printf("price + paid: %d bytes\n", small);

    /* every cut-off prefix is refused, and nothing reads out of bounds (sanitizers are on) */
    OrderEntity* o = from_text(full);
    uint8_t* blob = NULL;
    size_t len = 0;
    CHECK(OrderEntity_pack(o, &blob, &len) == 0);
    for (size_t cut = 0; cut < len; cut++) {
        uint8_t* copy = malloc(cut ? cut : 1); /* exact size, so a read past the end is caught */
        memcpy(copy, blob, cut);
        OrderEntity* bad = OrderEntity_unpack(copy, cut);
        CHECK(bad == NULL);
        free(copy);
    }
    /* trailing bytes are refused by unpack, reported by unpack_prefix */
    uint8_t* two = malloc(len * 2);
    memcpy(two, blob, len);
    memcpy(two + len, blob, len);
    CHECK(OrderEntity_unpack(two, len * 2) == NULL);
    size_t used = 0;
    OrderEntity* first = OrderEntity_unpack_prefix(two, len * 2, &used);
    CHECK(first != NULL && used == len);
    OrderEntity* second = OrderEntity_unpack_prefix(two + used, len * 2 - used, &used);
    CHECK(second != NULL && used == len);
    char* a = json_of(first);
    char* b = json_of(second);
    CHECK(strcmp(a, b) == 0);
    free(a); free(b);
    OrderEntity_free(first); OrderEntity_free(second);
    free(two);

    /* garbage never crashes */
    uint8_t junk[64];
    for (int seed = 0; seed < 2000; seed++) {
        for (int i = 0; i < 64; i++) junk[i] = (uint8_t) ((seed * 131 + i * 17 + (i >> 1) * seed) >> 2);
        OrderEntity* g = OrderEntity_unpack(junk, 1 + (size_t) (seed % 63));
        OrderEntity_free(g);
    }

    /* a damaged copy of a real blob: flip each byte */
    for (size_t i = 0; i < len; i++) {
        uint8_t* copy = malloc(len);
        memcpy(copy, blob, len);
        copy[i] ^= 0xFF;
        OrderEntity* g = OrderEntity_unpack(copy, len);
        OrderEntity_free(g);
        free(copy);
    }

    /* text form */
    char* text = emipack_to_base64(blob, len);
    uint8_t* raw = NULL;
    size_t rawlen = 0;
    CHECK(emipack_from_base64(text, strlen(text), &raw, &rawlen) == 0);
    CHECK(rawlen == len && memcmp(raw, blob, len) == 0);
    free(text); free(raw);
    CHECK(emipack_from_base64("not base64!", 11, &raw, &rawlen) != 0);

    OrderEntity_free(o);
    free(blob);
    printf("ok\n");
    return 0;
}
`

// TestPackEntities: Entity_pack / Entity_unpack with only the pack tag - no SQLite anywhere.
func TestPackEntities(t *testing.T) {
	bin, files := buildPackProgram(t, entityYml, string(Pack), packMain)
	for name, script := range files {
		if !strings.HasSuffix(name, ".md") && strings.Contains(script, "sqlite3.h") {
			t.Errorf("%s includes sqlite3.h, the pack tag must not need it", name)
		}
	}
	if _, ok := files["emisql.h"]; ok {
		t.Error("emisql.h shipped without the sqlite tag")
	}
	for _, want := range []string{"emipack.h", "emilayout.h"} {
		if _, ok := files[want]; !ok {
			t.Errorf("%s not shipped with the pack tag", want)
		}
	}
	if _, ok := files["shop_schema.h"]; ok {
		t.Error("the sql schema is only for the sqlite tag")
	}
	if out, err := exec.Command(bin).CombinedOutput(); err != nil {
		t.Fatalf("pack checks failed: %v\n%s", err, out)
	} else {
		t.Logf("\n%s", out)
	}
}

const itemV1 = `name: shop
entities:
  - name: item
    fields:
      - name: title
        type: string
      - name: count
        type: int
`

// a later version only appended fields: a column and an array
const itemV2 = itemV1 + `      - name: price
        type: float64
      - name: tags
        type: array
        fields:
          - name: name
            type: string
`

const itemWrite = `
#include <stdio.h>
#include "item_entity.h"
int main(int argc, char** argv) {
    (void) argc;
    cJSON* in = cJSON_Parse("{\"title\":\"kettle\",\"count\":-5}");
    ItemEntity* e = ItemEntity_from_json(in);
    cJSON_Delete(in);
    uint8_t* blob; size_t len;
    if (ItemEntity_pack(e, &blob, &len) != 0) return 1;
    FILE* f = fopen(argv[1], "wb");
    fwrite(blob, 1, len, f);
    fclose(f);
    ItemEntity_free(e);
    free(blob);
    return 0;
}
`

const itemRead = `
#include <stdio.h>
#include <string.h>
#include "item_entity.h"
int main(int argc, char** argv) {
    (void) argc;
    uint8_t buf[256];
    FILE* f = fopen(argv[1], "rb");
    size_t len = fread(buf, 1, sizeof(buf), f);
    fclose(f);
    ItemEntity* e = ItemEntity_unpack(buf, len);
    if (!e) { printf("older blob refused\n"); return 1; }
    cJSON* j = ItemEntity_to_json(e);
    char* s = cJSON_PrintUnformatted(j);
    printf("%s\n", s);
    int ok = strstr(s, "\"title\":\"kettle\"") && strstr(s, "\"count\":-5") && strstr(s, "\"price\":0");
    free(s);
    cJSON_Delete(j);
    ItemEntity_free(e);
    return ok ? 0 : 1;
}
`

// TestPackOlderBlob: a blob written by version 1 of an entity is read by version 2, which has
// appended fields; and the other way round - a blob with more than the reader knows - is refused.
func TestPackOlderBlob(t *testing.T) {
	blob := filepath.Join(t.TempDir(), "item.bin")
	write1, _ := buildPackProgram(t, itemV1, string(Pack), itemWrite)
	if out, err := exec.Command(write1, blob).CombinedOutput(); err != nil {
		t.Fatalf("write v1: %v\n%s", err, out)
	}
	read2, _ := buildPackProgram(t, itemV2, string(Pack), itemRead)
	out, err := exec.Command(read2, blob).CombinedOutput()
	if err != nil {
		t.Fatalf("v2 could not read a v1 blob: %v\n%s", err, out)
	}
	t.Logf("v1 blob read by v2: %s", out)

	write2, _ := buildPackProgram(t, itemV2, string(Pack), itemWrite)
	if out, err := exec.Command(write2, blob).CombinedOutput(); err != nil {
		t.Fatalf("write v2: %v\n%s", err, out)
	}
	read1, _ := buildPackProgram(t, itemV1, string(Pack), strings.Replace(itemRead, `&& strstr(s, "\"price\":0")`, "", 1))
	if out, err := exec.Command(read1, blob).CombinedOutput(); err == nil {
		t.Fatalf("v1 read a v2 blob, it should refuse: %s", out)
	}
}

// TestPackWithSqlite: both tags together still give one header which has the sql functions and the
// binary ones, and a packed entity equals the one stored in SQLite.
func TestPackWithSqlite(t *testing.T) {
	// with both tags the header still compiles where there is no sqlite3.h: -DEMI_NO_SQLITE leaves the sql out
	cc, err := exec.LookPath("cc")
	if err != nil {
		t.Skip("no C compiler available")
	}
	module0, err := core.StringToEmi(entityYml)
	if err != nil {
		t.Fatal(err)
	}
	files0, err := CModuleFullVirtualFiles(&module0, core.MicroGenContext{Content: entityYml, Flags: map[string]string{}, Tags: string(Sqlite) + "," + string(Pack)})
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	for _, f := range files0 {
		full := filepath.Join(dir, f.Location, f.Name+f.Extension)
		_ = os.MkdirAll(filepath.Dir(full), 0755)
		_ = os.WriteFile(full, []byte(f.ActualScript), 0644)
	}
	trap := filepath.Join(t.TempDir(), "trap")
	_ = os.MkdirAll(trap, 0755)
	_ = os.WriteFile(filepath.Join(trap, "sqlite3.h"), []byte("#error \"sqlite3.h must not be included\"\n"), 0644)
	_ = os.WriteFile(filepath.Join(dir, "main.c"), []byte("#define EMI_NO_SQLITE\n#include \"order_entity.h\"\nint main(void){OrderEntity*o=OrderEntity_new();uint8_t*b;size_t n;int rc=OrderEntity_pack(o,&b,&n);free(b);OrderEntity_free(o);return rc;}\n"), 0644)
	if out, err := exec.Command(cc, "-std=c99", "-Wall", "-Wextra", "-Werror", "-I", trap, "-I", dir, filepath.Join(dir, "main.c"),
		filepath.Join(dir, "cjson", "cJSON.c"), "-o", filepath.Join(dir, "prog"), "-lm").CombinedOutput(); err != nil {
		t.Fatalf("EMI_NO_SQLITE still needs sqlite3.h:\n%s", out)
	}

	module, err := core.StringToEmi(entityYml)
	if err != nil {
		t.Fatal(err)
	}
	ctx := core.MicroGenContext{Content: entityYml, Flags: map[string]string{}, Tags: string(Sqlite) + "," + string(Pack)}
	files, err := CModuleFullVirtualFiles(&module, ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if f.Name == "order_entity" {
			for _, want := range []string{"OrderEntity_pack", "OrderEntity_unpack", "OrderEntity_insert", "emipack.h", "emisql.h"} {
				if !strings.Contains(f.ActualScript, want) {
					t.Errorf("order_entity.h lacks %s", want)
				}
			}
		}
	}
}
