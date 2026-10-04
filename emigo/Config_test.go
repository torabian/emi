//go:build !wasm

package emigo

import (
	"os"
	"path/filepath"
	"testing"
)

type envTestConfigA struct {
	Dsn  string `envconfig:"TEST_DSN"`
	Port int64  `envconfig:"TEST_PORT"`
}

type envTestConfigB struct {
	BaseUrl string `envconfig:"TEST_BASE_URL"`
}

// Saving one config struct must keep the keys another struct (or a person) put in the same
// file, and an empty value must unset only its own key.
func TestSaveEnvFileMergesIntoExistingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".env")

	if err := SaveEnvFile(&envTestConfigA{Dsn: "host=x", Port: 4500}, path); err != nil {
		t.Fatal(err)
	}
	if err := SaveEnvFile(&envTestConfigB{BaseUrl: "http://localhost"}, path); err != nil {
		t.Fatal(err)
	}

	got, _ := os.ReadFile(path)
	want := "TEST_BASE_URL=http://localhost\nTEST_DSN=host=x\nTEST_PORT=4500\n"
	if string(got) != want {
		t.Fatalf("other struct's keys were lost:\n got %q\nwant %q", got, want)
	}

	// An empty value clears just that key.
	if err := SaveEnvFile(&envTestConfigB{BaseUrl: ""}, path); err != nil {
		t.Fatal(err)
	}
	got, _ = os.ReadFile(path)
	if string(got) != "TEST_DSN=host=x\nTEST_PORT=4500\n" {
		t.Fatalf("unexpected after clearing: %q", got)
	}
}
