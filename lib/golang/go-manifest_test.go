package golang

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// mcpManifestTestModuleYaml exercises a "go-mcp" manifest: editInvoice is a
// From-derived intent (so GoIntentsGenerate generates a real
// RegisterEditInvoiceIntentIntentTool for it - see go-intent-render_test.go's own
// module, reused here in spirit), searchInvoices is a standalone intent with no From
// (so no Register...IntentTool exists for it at all, and the manifest must skip it),
// and the manifest's own includes: excludes searchInvoices from consideration by name
// - both routes to the same "must not appear" outcome, so together they cover a
// standalone intent being skipped for lack of a resolved action *and* an
// explicitly-excluded one being skipped by pattern.
const mcpManifestTestModuleYaml = `
name: billing
namespace: billing
dtos:
  - name: invoicePatch
    fields:
      - name: amount
        type: float64?
actions:
  - name: editInvoice
    url: /invoices/:uniqueId
    in:
      dto: InvoicePatchDto
intents:
  - name: searchInvoices
    description: "Search invoices by free text query."
    in:
      fields:
        - name: query
          type: string
  - name: editInvoiceIntent
    from: editInvoice
manifests:
  - name: billing
    package: billingdefs
    types:
      - go-mcp
    excludes:
      - searchInvoices
`

// TestGoManifest_McpBundlesIntents checks that a "go-mcp" manifest type combines a
// module's own From-derived intents into one <Name>McpToolManifest(reg) function -
// the manifest-level equivalent of <Name>CliManifest for CLI commands - skipping any
// intent with no resolved action (nothing to call) or matched by the manifest's own
// excludes, and that the generated file is syntactically valid Go.
func TestGoManifest_McpBundlesIntents(t *testing.T) {
	module, err := core.StringToEmi(mcpManifestTestModuleYaml)
	if err != nil {
		t.Fatalf("StringToEmi error: %v", err)
	}

	files, err := GoModuleFull(&module, core.MicroGenContext{})
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}

	var manifestFile *core.VirtualFile
	for i := range files {
		if files[i].Name == "BillingManifest" {
			manifestFile = &files[i]
			break
		}
	}
	if manifestFile == nil {
		t.Fatalf("expected a BillingManifest file, got: %+v", fileNames(files))
	}
	got := string(manifestFile.ActualScript)

	wantContains := []string{
		"func BillingMcpToolManifest(reg *emigo.ToolRegistry) {",
		"RegisterEditInvoiceIntentIntentTool(reg, EditInvoiceAction)",
	}
	for _, want := range wantContains {
		if !strings.Contains(got, want) {
			t.Errorf("expected BillingManifest.go to contain %q, got:\n%s", want, got)
		}
	}

	// searchInvoices has no From (no resolved action, no Register...IntentTool ever
	// generated for it) - it must never be bundled, regardless of the manifest's own
	// excludes list also happening to name it.
	if strings.Contains(got, "SearchInvoices") {
		t.Errorf("standalone intent with no From must not appear in the mcp manifest, got:\n%s", got)
	}

	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "BillingManifest.go", got, parser.AllErrors); err != nil {
		t.Fatalf("generated BillingManifest.go is not syntactically valid Go: %v\n\n%s", err, got)
	}
}
