package golang

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestEntitySplitCliOwnFile verifies that with the "split-cli" tag set, an entity's own
// CLI helpers (Get{Entity}CliFlags/Cast{Entity}FromCli, rendered via
// GoCommonStructGeneratorCli through GoEntityRender) land in their own *Cli.go virtual
// file - not folded into the combined entity file behind the main struct's script. If
// they were, go/format (via AsFullDocument's FormatGoCode) would hoist the CLI chunk's
// leading `//go:build !wasm` comment to the very top of the *combined* file - gofmt
// promotes a `//go:build` comment found anywhere in a file to file scope, regardless of
// where it originally sat (verified directly against `gofmt`) - silently gating the
// entity's struct and all its Create/Update/Get/Browse/Delete functions behind !wasm
// too, not just the CLI helpers.
func TestEntitySplitCliOwnFile(t *testing.T) {
	module := newSplitCliTestModule(t)

	ctx := core.MicroGenContext{Tags: "split-cli"}

	files, err := GoModuleFull(module, ctx)
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}

	var mainFile, cliFile *core.VirtualFile
	for i := range files {
		switch files[i].Name {
		case "WidgetEntity.go":
			mainFile = &files[i]
		case "WidgetEntityCli":
			cliFile = &files[i]
		}
	}

	if mainFile == nil {
		t.Fatalf("expected a WidgetEntity.go file, got: %+v", fileNames(files))
	}
	if cliFile == nil {
		t.Fatalf("expected split-cli to produce a separate WidgetEntityCli file, got: %+v", fileNames(files))
	}

	if !strings.HasPrefix(strings.TrimLeft(cliFile.ActualScript, "\r\n\t "), "//go:build !wasm") {
		t.Errorf("expected WidgetEntityCli to start with a //go:build !wasm constraint, got:\n%s", cliFile.ActualScript)
	}
	if !strings.Contains(cliFile.ActualScript, "GetWidgetEntityCliFlags") {
		t.Errorf("expected WidgetEntityCli to contain the entity's CLI flags helper, got:\n%s", cliFile.ActualScript)
	}

	if strings.Contains(mainFile.ActualScript, "GetWidgetEntityCliFlags") {
		t.Errorf("expected the entity's CLI flags helper to be split out of the main WidgetEntity file, but found it there:\n%s", mainFile.ActualScript)
	}
	if strings.Contains(mainFile.ActualScript, "//go:build") {
		t.Errorf("expected no stray //go:build constraint in the main WidgetEntity file, got:\n%s", mainFile.ActualScript)
	}
}

// TestEntityNoSplitCliStaysMerged is TestEntitySplitCliOwnFile's counterpart for the
// default (no "split-cli" tag) case: the entity's CLI helpers must stay exactly where
// they always have - appended inside the combined WidgetEntity.go file, with no separate
// *Cli file and no //go:build constraint anywhere. The split above must be strictly
// opt-in.
func TestEntityNoSplitCliStaysMerged(t *testing.T) {
	module := newSplitCliTestModule(t)

	files, err := GoModuleFull(module, core.MicroGenContext{})
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}

	var mainFile *core.VirtualFile
	for i := range files {
		switch files[i].Name {
		case "WidgetEntity.go":
			mainFile = &files[i]
		case "WidgetEntityCli":
			t.Fatalf("did not expect a separate WidgetEntityCli file without the split-cli tag, got: %+v", fileNames(files))
		}
	}

	if mainFile == nil {
		t.Fatalf("expected a WidgetEntity.go file, got: %+v", fileNames(files))
	}
	if !strings.Contains(mainFile.ActualScript, "GetWidgetEntityCliFlags") {
		t.Errorf("expected the entity's CLI flags helper to stay merged into WidgetEntity.go, got:\n%s", mainFile.ActualScript)
	}
	if strings.Contains(mainFile.ActualScript, "//go:build") {
		t.Errorf("did not expect any //go:build constraint without the split-cli tag, got:\n%s", mainFile.ActualScript)
	}
}

// newSplitCliTestModule builds the single-entity module TestEntitySplitCliOwnFile and
// TestEntityNoSplitCliStaysMerged both generate from, differing only in the tag passed
// to GoModuleFull.
func newSplitCliTestModule(t *testing.T) *core.Emi {
	t.Helper()

	// Registers the entity preprocess hooks (PreprocessEntityOptionalDtos/Dtos/Actions)
	// used below - idempotent, mirrors what any real caller does before generating Go.
	GetGolangPublicActions()

	module := &core.Emi{
		Name: "sample",
		Entities: []*core.Module3Entity{
			{
				Name:  "widget",
				Table: "widgets",
				Fields: []*core.EmiField{
					{Name: "title", Type: core.FieldTypeString},
				},
			},
		},
	}

	if err := module.PreprocessForAction("go"); err != nil {
		t.Fatalf("Preprocess error: %v", err)
	}

	return module
}

func fileNames(files []core.VirtualFile) []string {
	names := make([]string, len(files))
	for i, f := range files {
		names[i] = f.Name
	}
	return names
}
