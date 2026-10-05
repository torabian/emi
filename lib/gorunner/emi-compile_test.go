package gorunner

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/urfave/cli/v3"
)

// runCompile invokes the real `compile --path <path>` CLI command the same way a
// user would from the terminal (see EntryPoint), so these tests exercise the actual
// flag wiring (createCliContext, target.Output resolution, ...) rather than calling
// internals directly.
func runCompile(t *testing.T, path string) {
	t.Helper()
	app := &cli.Command{
		Name:     "emi",
		Commands: []*cli.Command{&CompileCommand},
	}
	if err := app.Run(context.Background(), []string{"emi", "compile", "--path", path}); err != nil {
		t.Fatalf("compile failed: %v", err)
	}
}

// writeCleanTestFixture writes a minimal, single-target emi definition using the
// "preprocessor" compiler - the simplest target type (it just writes the
// preprocessed definition back out as preprocessed.yml, with no per-language
// codegen involved), which keeps these tests focused on Clean's own file-system
// behavior rather than any particular compiler's output.
func writeCleanTestFixture(t *testing.T, dir string, clean bool) string {
	t.Helper()
	fixture := `name: cleantest
targets:
  - compiler: preprocessor
    output: ./out
    clean: ` + boolYaml(clean) + `
dtos:
  - name: sample
    fields:
      - name: title
        type: string
`
	path := filepath.Join(dir, "cleantest.emi.yml")
	if err := os.WriteFile(path, []byte(fixture), 0644); err != nil {
		t.Fatalf("writing fixture: %v", err)
	}
	return path
}

func boolYaml(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func writeStaleFile(t *testing.T, outDir string) string {
	t.Helper()
	if err := os.MkdirAll(outDir, 0755); err != nil {
		t.Fatalf("mkdir %v: %v", outDir, err)
	}
	stale := filepath.Join(outDir, "StaleAction.ts")
	if err := os.WriteFile(stale, []byte("// leftover from a since-removed action"), 0644); err != nil {
		t.Fatalf("writing stale file: %v", err)
	}
	return stale
}

func TestCompile_Clean_RemovesStaleFilesBeforeWriting(t *testing.T) {
	dir := t.TempDir()
	path := writeCleanTestFixture(t, dir, true)
	outDir := filepath.Join(dir, "out")
	stale := writeStaleFile(t, outDir)

	runCompile(t, path)

	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatalf("expected the stale file to be removed by clean: true, stat err = %v", err)
	}
	if _, err := os.Stat(filepath.Join(outDir, "preprocessed.yml")); err != nil {
		t.Fatalf("expected preprocessed.yml to have been (re)generated: %v", err)
	}
}

func TestCompile_NoClean_KeepsStaleFiles(t *testing.T) {
	dir := t.TempDir()
	path := writeCleanTestFixture(t, dir, false)
	outDir := filepath.Join(dir, "out")
	stale := writeStaleFile(t, outDir)

	runCompile(t, path)

	if _, err := os.Stat(stale); err != nil {
		t.Fatalf("expected the stale file to survive without clean: true, stat err = %v", err)
	}
	if _, err := os.Stat(filepath.Join(outDir, "preprocessed.yml")); err != nil {
		t.Fatalf("expected preprocessed.yml to have been generated: %v", err)
	}
}

// TestCompile_Clean_OnlyTouchesItsOwnTargetDirectory verifies clean: true on one
// target never reaches into a sibling target's Output, even when both targets run
// in the same compile pass.
func TestCompile_Clean_OnlyTouchesItsOwnTargetDirectory(t *testing.T) {
	dir := t.TempDir()
	fixture := `name: cleantest
targets:
  - compiler: preprocessor
    output: ./out-a
    clean: true
  - compiler: preprocessor
    output: ./out-b
dtos:
  - name: sample
    fields:
      - name: title
        type: string
`
	path := filepath.Join(dir, "cleantest.emi.yml")
	if err := os.WriteFile(path, []byte(fixture), 0644); err != nil {
		t.Fatalf("writing fixture: %v", err)
	}

	outB := filepath.Join(dir, "out-b")
	untouched := writeStaleFile(t, outB)

	runCompile(t, path)

	if _, err := os.Stat(untouched); err != nil {
		t.Fatalf("expected the sibling (non-clean) target's own file to be left alone, stat err = %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "out-a", "preprocessed.yml")); err != nil {
		t.Fatalf("expected out-a/preprocessed.yml to have been generated: %v", err)
	}
	if _, err := os.Stat(filepath.Join(outB, "preprocessed.yml")); err != nil {
		t.Fatalf("expected out-b/preprocessed.yml to have been generated too: %v", err)
	}
}

// A module compiles to the sql of both databases through its own targets: the same compiler,
// told which database by the tags of the target.
func TestCompile_EntitySqlTargetPicksTheDatabaseByTag(t *testing.T) {
	dir := t.TempDir()
	fixture := `name: shop
targets:
  - compiler: entity-sql
    output: ./sql
  - compiler: entity-sql
    output: ./sql
    tags:
      - sqlite
entities:
  - name: product
    fields:
      - name: title
        type: string
      - name: tags
        type: array
        fields:
          - name: label
            type: string
`
	path := filepath.Join(dir, "shop.emi.yml")
	if err := os.WriteFile(path, []byte(fixture), 0644); err != nil {
		t.Fatal(err)
	}

	runCompile(t, path)

	pg, err := os.ReadFile(filepath.Join(dir, "sql", "shop.sql"))
	if err != nil {
		t.Fatalf("the postgres file: %v", err)
	}
	lite, err := os.ReadFile(filepath.Join(dir, "sql", "shop.sqlite.sql"))
	if err != nil {
		t.Fatalf("the sqlite file: %v", err)
	}
	if !strings.Contains(string(pg), "bigserial") || !strings.Contains(string(pg), "ADD COLUMN IF NOT EXISTS") {
		t.Errorf("not postgres:\n%s", pg)
	}
	if !strings.Contains(string(lite), "AUTOINCREMENT") || strings.Contains(string(lite), "bigserial") {
		t.Errorf("not sqlite:\n%s", lite)
	}
}
