package emigo_test

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/torabian/emi/emigo"
	"gopkg.in/yaml.v3"
)

func TestNullable_SetAndGet(t *testing.T) {
	var n emigo.Nullable[int]

	if n.IsSet() {
		t.Fatal("Expected IsSet false initially")
	}

	// Set a value
	n.Set(emigo.Ptr(42))
	if !n.IsSet() {
		t.Fatal("Expected IsSet true after Set")
	}
	if n.IsNull() {
		t.Fatal("Expected IsNull false after Set to value")
	}
	v, ok := n.Get()
	if !ok || *v != 42 {
		t.Fatalf("Expected value 42, got %v", v)
	}

	// Set to nil explicitly
	n.Set(nil)
	if !n.IsSet() {
		t.Fatal("Expected IsSet true after Set(nil)")
	}
	if !n.IsNull() {
		t.Fatal("Expected IsNull true after Set(nil)")
	}
	v, ok = n.Get()
	if ok && v != nil {
		t.Fatalf("Expected value nil, got %v", v)
	}

	// Reset
	n.Reset()
	if n.IsSet() {
		t.Fatal("Expected IsSet false after Reset")
	}
	if n.IsNull() {
		t.Fatal("Expected IsNull false after Reset")
	}
}

func TestNullable_JSON(t *testing.T) {
	// Value
	n := emigo.NullableOf(123)
	data, err := json.Marshal(n)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(data) != "123" {
		t.Fatalf("Expected 123, got %s", string(data))
	}

	// Explicit null
	n.Set(nil)
	data, _ = json.Marshal(n)
	if string(data) != "null" {
		t.Fatalf("Expected null, got %s", string(data))
	}

	// Unmarshal value
	var n2 emigo.Nullable[int]
	err = json.Unmarshal([]byte("456"), &n2)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if !n2.IsSet() || *n2.Ptr() != 456 {
		t.Fatalf("Expected 456, got %v", n2.Ptr())
	}

	// Unmarshal null
	err = json.Unmarshal([]byte("null"), &n2)
	if err != nil {
		t.Fatalf("Unmarshal null error: %v", err)
	}
	if !n2.IsSet() || !n2.IsNull() {
		t.Fatalf("Expected IsNull true, got IsNull=%v", n2.IsNull())
	}
}

func TestNullable_YAML(t *testing.T) {
	// Value
	n := emigo.NullableOf("hello")
	data, err := yaml.Marshal(n)
	if err != nil {
		t.Fatalf("YAML Marshal error: %v", err)
	}
	if string(data) != "hello\n" {
		t.Fatalf("Expected 'hello', got %s", string(data))
	}

	// Explicit null
	n.Set(nil)
	data, _ = yaml.Marshal(n)
	if string(data) != "null\n" {
		t.Fatalf("Expected 'null', got %s", string(data))
	}

	// Unmarshal value
	var n2 emigo.Nullable[string]
	err = yaml.Unmarshal([]byte("world"), &n2)
	if err != nil {
		t.Fatalf("YAML Unmarshal error: %v", err)
	}
	if !n2.IsSet() || *n2.Ptr() != "world" {
		t.Fatalf("Expected 'world', got %v", n2.Ptr())
	}

}

func TestNullable_OrDefaultAndMap(t *testing.T) {
	n := emigo.NullableOf(10)

	// OrDefault
	if n.OrDefault(100) != 10 {
		t.Fatalf("Expected OrDefault 10, got %v", n.OrDefault(100))
	}

	// Map
	n.Map(func(x int) int { return x * 3 })
	if *n.Ptr() != 30 {
		t.Fatalf("Expected 30 after Map, got %v", *n.Ptr())
	}

	// OrDefault when nil
	n.Set(nil)
	if n.OrDefault(999) != 999 {
		t.Fatalf("Expected OrDefault 999 for nil, got %v", n.OrDefault(999))
	}
}

type UserInfo struct {
	FirstName  emigo.Nullable[string] `json:"first_name,omitempty"`
	LastName   emigo.Nullable[string] `json:"last_name,omitempty"`
	LoginCount emigo.Nullable[int]    `json:"login_count,omitempty,omitempty"`
	IsOnline   emigo.Nullable[bool]   `json:"is_online,omitempty"`
}

func TestNullable_StructJSON(t *testing.T) {
	// Create a struct with mixed Nullable fields
	user := UserInfo{
		FirstName:  emigo.NullableOf("Alice"),
		LastName:   emigo.NullableOfPtr[string](nil), // Explicit null
		LoginCount: emigo.NullableOf(10),
		IsOnline:   emigo.Nullable[bool]{},
	}

	// Marshal to YAML
	data, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		t.Fatalf("JSON Marshal error: %v", err)
	}

	// Print JSON output for verification (optional)
	t.Logf("JSON output:\n%s", string(data))

	// Expected YAML should include:
	// first_name: Alice
	// last_name: null
	// login_count: 10
	// is_online: null  (or omitted depending on Marshal implementation)
	// We accept null for unset as YAML does not distinguish undefined

	// Unmarshal back
	var user2 UserInfo
	err = json.Unmarshal(data, &user2)
	if err != nil {
		t.Fatalf("YAML Unmarshal error: %v", err)
	}

	// Check values
	if !user2.FirstName.IsSet() || *user2.FirstName.Ptr() != "Alice" {
		t.Errorf("Expected FirstName=Alice, got %v", user2.FirstName.Ptr())
	}

	if !user2.LastName.IsSet() || !user2.LastName.IsNull() {
		t.Errorf("Expected LastName=null, got %v", user2.LastName.Ptr())
	}

	if !user2.LoginCount.IsSet() || *user2.LoginCount.Ptr() != 10 {
		t.Errorf("Expected LoginCount=10, got %v", user2.LoginCount.Ptr())
	}

}

func TestUserInfo_JSON_Unmarshal_FirstNameUndefined(t *testing.T) {
	jsonStr := `{
		"last_name": "Doe",
		"login_count": 5,
		"is_online": true
	}`

	var user UserInfo
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// FirstName is undefined (not present in JSON)
	if user.FirstName.IsSet() {
		t.Errorf("Expected FirstName to be undefined (IsSet=false), got IsSet=true")
	}

	if user.FirstName.IsNull() {
		t.Errorf("Field should not be null, because the value is not set yet.")
	}
}
