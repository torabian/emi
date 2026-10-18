package golang

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// Temporary verification test for the split-cli / "encoding" import bug.
func TestSplitCliEncodingDependencyPlacement(t *testing.T) {
	fields := []*core.EmiField{
		{
			Name:    "Payload",
			Type:    core.FieldTypeComplex,
			Complex: "MyComplex",
		},
	}

	ctx := core.MicroGenContext{Tags: "split-cli"}
	goctx := GoCommonStructContext{RootClassName: "Sample"}

	particles, err := GoCommonStructGenerator(fields, ctx, goctx)
	if err != nil {
		t.Fatalf("GoCommonStructGenerator error: %v", err)
	}

	if particles.CliHelpers == nil {
		t.Fatalf("expected CliHelpers to be split out")
	}

	mainHasEncoding := depHasLocation(particles.MainClass.CodeChunkDependensies, "encoding")
	cliHasEncoding := depHasLocation(particles.CliHelpers.CodeChunkDependensies, "encoding")
	cliUsesEncoding := strings.Contains(string(particles.CliHelpers.ActualScript), "encoding.TextUnmarshaler")
	mainUsesEncoding := strings.Contains(string(particles.MainClass.ActualScript), "encoding.TextUnmarshaler")

	t.Logf("MainClass deps: %+v", particles.MainClass.CodeChunkDependensies)
	t.Logf("CliHelpers deps: %+v", particles.CliHelpers.CodeChunkDependensies)

	if mainUsesEncoding {
		t.Fatalf("MainClass script unexpectedly uses encoding.TextUnmarshaler")
	}
	if !cliUsesEncoding {
		t.Fatalf("CliHelpers script expected to use encoding.TextUnmarshaler")
	}
	if mainHasEncoding {
		t.Errorf(`BUG: MainClass declares "encoding" import but never uses it (unused import)`)
	}
	if !cliHasEncoding {
		t.Errorf(`BUG: CliHelpers uses encoding.TextUnmarshaler but does not declare the "encoding" import (missing import)`)
	}
}

func depHasLocation(deps []core.CodeChunkDependency, loc string) bool {
	for _, d := range deps {
		if d.Location == loc {
			return true
		}
	}
	return false
}

// A field typed by another go package must bring that import into the split *Cli.go file,
// because Cast<Type>FromCli is called package-qualified there.
func TestSplitCliExternalProviderImport(t *testing.T) {
	fields := []*core.EmiField{
		{
			Name:     "User",
			Type:     core.FieldTypeOneNullable,
			Provider: "github.com/x/user/defs",
			Module:   "UserEntity",
		},
	}
	ctx := core.MicroGenContext{Tags: "split-cli"}
	goctx := GoCommonStructContext{RootClassName: "Sample"}

	particles, err := GoCommonStructGenerator(fields, ctx, goctx)
	if err != nil {
		t.Fatal(err)
	}
	if particles.CliHelpers == nil {
		t.Fatalf("expected CliHelpers to be split out")
	}
	t.Logf("cli script: %s", particles.CliHelpers.ActualScript)
	t.Logf("cli deps: %+v", particles.CliHelpers.CodeChunkDependensies)
	if !depHasLocation(particles.CliHelpers.CodeChunkDependensies, "github.com/x/user/defs") {
		t.Errorf("BUG: CliHelpers misses the import of the package providing UserEntity")
	}
}
