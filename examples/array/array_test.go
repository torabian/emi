package collection_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/goccy/go-json"

	"github.com/torabian/emi/emigo"
)

// Tag is a tiny item type used inside Array[Tag] for these tests.
type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Profile is a struct that embeds an Array field — this is the shape
// emi-generated DTOs use in practice, so we want to make sure JSON round-trips
// through a real struct (not just a bare Array) behave correctly.
type Profile struct {
	Name string           `json:"name"`
	Tags emigo.Array[Tag] `json:"tags"`
}

// --- console pretty-printing helpers ----------------------------------------

func banner(t *testing.T, title string) {
	t.Helper()
	bar := strings.Repeat("─", len(title)+4)
	fmt.Println("┌" + bar + "┐")
	fmt.Println("│  " + title + "  │")
	fmt.Println("└" + bar + "┘")
}

func step(t *testing.T, label string, value any) {
	t.Helper()
	fmt.Printf("    %-22s → %v\n", label, value)
}

func describe(t *testing.T, c emigo.Array[Tag]) {
	t.Helper()
	step(t, "IsSet", c.IsSet())
	step(t, "__Operation", fmt.Sprintf("%q", c.Operation))
	step(t, "Items (len)", len(c.Items))
	step(t, "Items", c.Items)
}

// ---------------------------------------------------------------------------
//  JSON → Go (unmarshal)
// ---------------------------------------------------------------------------

func TestUnmarshal_FieldOmitted(t *testing.T) {
	banner(t, "JSON has no `tags` key at all → Array stays unset")

	payload := `{"name":"ali"}`
	step(t, "payload", payload)

	var p Profile
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	describe(t, p.Tags)

	if p.Tags.IsSet() {
		t.Fatal("expected IsSet=false when the field is absent from JSON")
	}
	if p.Tags.Operation != "" || p.Tags.Items != nil {
		t.Fatal("expected zero Operation/Items when the field is absent")
	}
	fmt.Println("    ✓ undefined → IsSet=false")
	fmt.Println()
}

func TestUnmarshal_ExplicitNull(t *testing.T) {
	banner(t, "JSON has `\"tags\": null` → IsSet=true, Items=nil (explicit clear)")

	payload := `{"name":"ali","tags":null}`
	step(t, "payload", payload)

	var p Profile
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	describe(t, p.Tags)

	if !p.Tags.IsSet() {
		t.Fatal("expected IsSet=true for explicit null (caller spoke about the field)")
	}
	if p.Tags.Operation != "" {
		t.Fatalf("expected empty Operation for null, got %q", p.Tags.Operation)
	}
	if p.Tags.Items != nil {
		t.Fatalf("expected nil Items for null, got %v", p.Tags.Items)
	}
	fmt.Println("    ✓ explicit null → IsSet=true, Items=nil")
	fmt.Println()
}

func TestUnmarshal_BareEmptyArray(t *testing.T) {
	banner(t, "JSON has `\"tags\": []` → IsSet=true, Operation=\"replace\", Items=[]")

	payload := `{"name":"ali","tags":[]}`
	step(t, "payload", payload)

	var p Profile
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	describe(t, p.Tags)

	if !p.Tags.IsSet() {
		t.Fatal("expected IsSet=true for empty array")
	}
	if p.Tags.Operation != "replace" {
		t.Fatalf("expected Operation=replace, got %q", p.Tags.Operation)
	}
	if len(p.Tags.Items) != 0 {
		t.Fatalf("expected zero items, got %d", len(p.Tags.Items))
	}
	fmt.Println("    ✓ bare [] → Operation=replace, Items length 0")
	fmt.Println()
}

func TestUnmarshal_BareArrayWithItems(t *testing.T) {
	banner(t, "JSON has `\"tags\": [...]` → bare array shorthand becomes a `replace` op")

	payload := `{"name":"ali","tags":[{"key":"plan","value":"pro"},{"key":"region","value":"eu"}]}`
	step(t, "payload", payload)

	var p Profile
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	describe(t, p.Tags)

	if !p.Tags.IsSet() || p.Tags.Operation != "replace" || len(p.Tags.Items) != 2 {
		t.Fatalf("unexpected state: %+v", p.Tags)
	}
	if p.Tags.Items[0].Key != "plan" || p.Tags.Items[1].Value != "eu" {
		t.Fatalf("items did not unmarshal correctly: %+v", p.Tags.Items)
	}
	fmt.Println("    ✓ bare array with items → Operation=replace, 2 items decoded")
	fmt.Println()
}

func TestUnmarshal_TaggedAppend(t *testing.T) {
	banner(t, "JSON has `{operation:\"append\", items:[...]}` → explicit append op")

	payload := `{"name":"ali","tags":{"__operation":"append","items":[{"key":"role","value":"admin"}]}}`
	step(t, "payload", payload)

	var p Profile
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	describe(t, p.Tags)

	if !p.Tags.IsSet() || p.Tags.Operation != "append" || len(p.Tags.Items) != 1 {
		t.Fatalf("unexpected state: %+v", p.Tags)
	}
	if p.Tags.Items[0].Key != "role" || p.Tags.Items[0].Value != "admin" {
		t.Fatalf("item not decoded: %+v", p.Tags.Items[0])
	}
	fmt.Println("    ✓ tagged object with operation=append → preserved")
	fmt.Println()
}

func TestUnmarshal_TaggedReplaceEmpty(t *testing.T) {
	banner(t, "JSON has `{operation:\"replace\", items:[]}` → explicit clear-the-set op")

	payload := `{"name":"ali","tags":{"__operation":"replace","items":[]}}`
	step(t, "payload", payload)

	var p Profile
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	describe(t, p.Tags)

	if !p.Tags.IsSet() || p.Tags.Operation != "replace" || len(p.Tags.Items) != 0 {
		t.Fatalf("unexpected state: %+v", p.Tags)
	}
	fmt.Println("    ✓ tagged replace with empty items → distinct from null")
	fmt.Println()
}

func TestUnmarshal_TaggedNoOperation(t *testing.T) {
	banner(t, "JSON has `{items:[...]}` (no operation key) → defaults to `replace`")

	payload := `{"name":"ali","tags":{"items":[{"key":"k","value":"v"}]}}`
	step(t, "payload", payload)

	var p Profile
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	describe(t, p.Tags)

	if p.Tags.Operation != "replace" {
		t.Fatalf("expected Operation to default to 'replace', got %q", p.Tags.Operation)
	}
	fmt.Println("    ✓ object without operation defaults to replace")
	fmt.Println()
}

// ---------------------------------------------------------------------------
//  Go → JSON (marshal)
// ---------------------------------------------------------------------------

func TestMarshal_UnsetEmitsNull(t *testing.T) {
	banner(t, "Unset Array (zero value) marshals as `null`")

	p := Profile{Name: "ali"} // Tags untouched
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	out := string(data)
	step(t, "output", out)

	if !strings.Contains(out, `"tags":null`) {
		t.Fatalf("expected tags:null, got: %s", out)
	}
	fmt.Println("    ✓ zero-value Array → null on the wire")
	fmt.Println()
}

func TestMarshal_ReplaceWithItems(t *testing.T) {
	banner(t, "ArrayReplace(items) marshals as tagged {operation:\"replace\", items:[...]}")

	p := Profile{
		Name: "ali",
		Tags: emigo.ArrayReplace([]Tag{{Key: "plan", Value: "pro"}}),
	}
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	out := string(data)
	step(t, "output", out)

	if !strings.Contains(out, `"__operation":"replace"`) || !strings.Contains(out, `"key":"plan"`) {
		t.Fatalf("missing expected fields in: %s", out)
	}
	fmt.Println("    ✓ replace with values → tagged object on the wire")
	fmt.Println()
}

func TestMarshal_AppendWithItems(t *testing.T) {
	banner(t, "ArrayAppend(items) marshals as tagged {__operation:\"append\", items:[...]}")

	p := Profile{
		Name: "ali",
		Tags: emigo.ArrayAppend([]Tag{{Key: "role", Value: "admin"}}),
	}
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	out := string(data)
	step(t, "output", out)

	if !strings.Contains(out, `"__operation":"append"`) || !strings.Contains(out, `"key":"role"`) {
		t.Fatalf("missing expected fields in: %s", out)
	}
	fmt.Println("    ✓ append with values → tagged object on the wire")
	fmt.Println()
}

func TestMarshal_EmptyReplace(t *testing.T) {
	banner(t, "ArrayReplace(nil items) is still a set value (clears the child rows)")

	p := Profile{
		Name: "ali",
		Tags: emigo.ArrayReplace[Tag](nil),
	}
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	out := string(data)
	step(t, "output", out)

	if !strings.Contains(out, `"__operation":"replace"`) {
		t.Fatalf("expected __operation=replace, got: %s", out)
	}
	fmt.Println("    ✓ explicit clear-the-set is wire-distinct from absent/null")
	fmt.Println()
}

// ---------------------------------------------------------------------------
//  Round-trip & API surface
// ---------------------------------------------------------------------------

func TestRoundTrip_NullPreservesIsSet(t *testing.T) {
	banner(t, "Round-trip: null in → IsSet stays true (caller's null is preserved as state)")

	var p Profile
	if err := json.Unmarshal([]byte(`{"name":"ali","tags":null}`), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	step(t, "IsSet after null", p.Tags.IsSet())
	step(t, "IsSet via Get()", func() bool {
		_, ok := p.Tags.Get()
		return ok
	}())

	if !p.Tags.IsSet() {
		t.Fatal("expected IsSet=true after null unmarshal")
	}
	// Note: re-marshalling a null-sourced Array emits the tagged object
	// form ({"__operation":"","items":null}), not "null". Confirm that explicitly
	// so future readers know the round-trip is asymmetric by design.
	data, err := json.Marshal(p.Tags)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	step(t, "re-marshalled", string(data))
	if string(data) == "null" {
		t.Fatal("documented behavior: a null-unmarshalled Array should re-marshal as a tagged object")
	}
	fmt.Println("    ✓ null is preserved as a set state, re-emitted as tagged object")
	fmt.Println()
}

func TestRoundTrip_ValuesIn_ValuesOut(t *testing.T) {
	banner(t, "Round-trip: values → JSON → values (identity through the wire)")

	original := Profile{
		Name: "ali",
		Tags: emigo.ArrayReplace([]Tag{{Key: "a", Value: "1"}, {Key: "b", Value: "2"}}),
	}
	wire, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	step(t, "wire", string(wire))

	var back Profile
	if err := json.Unmarshal(wire, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	describe(t, back.Tags)

	if back.Tags.Operation != "replace" || len(back.Tags.Items) != 2 {
		t.Fatalf("round-trip lost data: %+v", back.Tags)
	}
	if back.Tags.Items[0] != original.Tags.Items[0] || back.Tags.Items[1] != original.Tags.Items[1] {
		t.Fatal("items differ after round-trip")
	}
	fmt.Println("    ✓ items survive a marshal → unmarshal trip intact")
	fmt.Println()
}

func TestAPI_SetAndReset(t *testing.T) {
	banner(t, "Direct Go API: Set / Reset / IsSet / Get behave as documented")

	var c emigo.Array[Tag]
	step(t, "IsSet (zero value)", c.IsSet())
	if c.IsSet() {
		t.Fatal("zero value must report IsSet=false")
	}

	c.Set("append", []Tag{{Key: "x", Value: "y"}})
	step(t, "after Set(append,1)", fmt.Sprintf("IsSet=%v op=%q items=%d", c.IsSet(), c.Operation, len(c.Items)))
	if !c.IsSet() || c.Operation != "append" || len(c.Items) != 1 {
		t.Fatalf("unexpected: %+v", c)
	}

	ptr, ok := c.Get()
	step(t, "Get() returns", fmt.Sprintf("ptr!=nil=%v ok=%v", ptr != nil, ok))
	if !ok || ptr == nil {
		t.Fatal("Get() should return (ptr, true) for a set Array")
	}

	c.Reset()
	step(t, "after Reset()", fmt.Sprintf("IsSet=%v op=%q items=%d", c.IsSet(), c.Operation, len(c.Items)))
	if c.IsSet() || c.Operation != "" || c.Items != nil {
		t.Fatalf("Reset did not clear: %+v", c)
	}

	_, ok = c.Get()
	if ok {
		t.Fatal("Get() should return ok=false after Reset")
	}
	fmt.Println("    ✓ Set / Reset / Get cycle behaves correctly")
	fmt.Println()
}

func TestAPI_SQLValue(t *testing.T) {
	banner(t, "SQLValue: unset is skipped, set value emits a quoted jsonb literal")

	var unset emigo.Array[Tag]
	v, ok := unset.SQLValue()
	step(t, "unset → (value, ok)", fmt.Sprintf("%q, %v", v, ok))
	if ok {
		t.Fatal("unset Array must opt out of SQL column emission")
	}

	c := emigo.ArrayReplace([]Tag{{Key: "plan", Value: "pro"}})
	v, ok = c.SQLValue()
	step(t, "set → (value, ok)", fmt.Sprintf("%s, %v", v, ok))
	if !ok {
		t.Fatal("set Array must opt into SQL column emission")
	}
	if !strings.HasPrefix(v, "'") || !strings.HasSuffix(v, "'::jsonb") {
		t.Fatalf("expected '...'::jsonb literal, got %s", v)
	}
	if !strings.Contains(v, `"__operation":"replace"`) {
		t.Fatalf("expected __operation=replace inside the literal, got %s", v)
	}
	fmt.Println("    ✓ SQLValue opts out when unset, emits jsonb literal when set")
	fmt.Println()
}

// TestMain prints a summary header so the test output reads as a small report.
func TestMain(m *testing.M) {
	fmt.Println()
	fmt.Println("════════════════════════════════════════════════════════════════")
	fmt.Println("  emigo.Array[T]  —  JSON round-trip & API tests")
	fmt.Println("════════════════════════════════════════════════════════════════")
	fmt.Println()
	code := m.Run()
	fmt.Println("════════════════════════════════════════════════════════════════")
	if code == 0 {
		fmt.Println("  ✓ all Array cases passed")
	} else {
		fmt.Println("  ✗ at least one Array case failed (exit", code, ")")
	}
	fmt.Println("════════════════════════════════════════════════════════════════")
	fmt.Println()
	os.Exit(code)
}
