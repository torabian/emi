package golang

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// intentsTestModuleYaml exercises every source an intent's combined InputSchema can
// draw from at once: searchInvoices is a standalone intent with its own inline In
// (no From), editInvoice is derived (From) from an action whose body is a dto
// reference, has a URL path param (:uniqueId), a query param (verbose), and a typed
// request header (X-Idempotency-Key) - so InputSchema must merge body+query+path+
// header, while OutputSchema reflects Out plus its own typed response header
// (X-Request-Id), never the query/path/request-header sources.
const intentsTestModuleYaml = `
name: billing
namespace: billing
dtos:
  - name: invoice
    fields:
      - name: uniqueId
        type: string
      - name: amount
        type: float64
  - name: invoicePatch
    fields:
      - name: amount
        type: float64?
actions:
  - name: editInvoice
    url: /invoices/:uniqueId
    qs:
      - name: verbose
        type: bool
    in:
      dto: InvoicePatchDto
      headers:
        - name: X-Idempotency-Key
          type: string
          description: "Client-generated key to make retries safe."
    out:
      dto: InvoiceDto
      headers:
        - name: X-Request-Id
          type: string
intents:
  - name: searchInvoices
    description: "Search invoices by free text query."
    in:
      fields:
        - name: query
          type: string
          description: "Free text search query."
        - name: limit
          type: int?
  - name: editInvoiceIntent
    from: editInvoice
    description: "Edit an existing invoice."
`

// TestIntentsGenerate_MergesBodyQueryAndPathParams checks that GoIntentsGenerate
// builds one flat InputSchema per intent covering everything a tool caller actually
// needs to supply - inline fields for a standalone intent, and for a From-derived
// intent, its dto-referenced body merged with the source action's own query
// parameter, URL path parameter, and typed request header - while OutputSchema
// reflects Out plus the source action's own typed response header.
func TestIntentsGenerate_MergesBodyQueryAndPathParams(t *testing.T) {
	module, err := core.StringToEmi(intentsTestModuleYaml)
	if err != nil {
		t.Fatalf("StringToEmi error: %v", err)
	}

	files, err := GoModuleFull(&module, core.MicroGenContext{})
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}

	var intentsFile *core.VirtualFile
	for i := range files {
		if files[i].Name == "Intents" {
			intentsFile = &files[i]
			break
		}
	}
	if intentsFile == nil {
		t.Fatalf("expected an Intents file, got: %+v", fileNames(files))
	}

	got := intentsFile.ActualScript

	wantContains := []string{
		// searchInvoices: standalone intent, InputSchema built straight from its own
		// inline In fields - no From, so no query/path merge applies.
		`Name:            "searchInvoices",`,
		`"query":{"description":"Free text search query.","type":"string"}`,
		`"limit":{"type":"integer"}`,
		`"required":["query"]`,

		// editInvoiceIntent: derived from editInvoice. Its In is a dto reference
		// (InvoicePatchDto -> "amount"), merged with editInvoice's own query param
		// (verbose) and URL path param (uniqueId).
		`Name:            "editInvoiceIntent",`,
		`"amount":{"type":"number"}`,
		`"verbose":{"type":"boolean"}`,
		`"uniqueId":{"type":"string"}`,
		// uniqueId came from the URL path, so it's required even though nothing in
		// the dto/query marked it that way; amount is optional (float64? in the dto).
		`"required":["uniqueId"]`,

		// editInvoice's own typed request header is folded into InputSchema too,
		// always optional, with its description prefixed so a model can tell it
		// travels as a header rather than a body/query/path field.
		`"X-Idempotency-Key":{"description":"HTTP header. Client-generated key to make retries safe.","type":"string"}`,

		// ...and its typed response header is folded into OutputSchema.
		`"X-Request-Id":{"description":"HTTP header.","type":"string"}`,
	}

	for _, want := range wantContains {
		if !strings.Contains(string(got), want) {
			t.Errorf("expected generated Intents.go to contain %q, got:\n%s", want, got)
		}
	}

	// searchInvoices has no From, so it has no query/path params to merge - its
	// InputSchema must not mention verbose/uniqueId anywhere.
	if strings.Contains(string(got), `"query":{"description":"Free text search query.","type":"string"},"verbose"`) {
		t.Errorf("searchInvoices' schema must not pick up editInvoice's query/path params")
	}
}

// TestIntentsGenerate_ToolRegistration checks that GoIntentsGenerate emits a
// Register<Name>IntentTool function - the *ActionGin/*ActionCli-style helper meant to
// replace a project's own hand-written AiTools.go - only for editInvoiceIntent (From:
// editInvoice, so there's a real handler function/Request/Response/PathParameter
// shape to build against), never for searchInvoices (a standalone intent with no
// backing action, so there's nothing a generated registration could safely call).
// It also checks the generated file is syntactically valid Go, since a template typo
// in the hand-built registration source (go-intent-render.go's
// renderIntentToolRegistration) would otherwise only surface as a downstream `go
// build` failure in whatever project happens to declare a From-derived intent next.
func TestIntentsGenerate_ToolRegistration(t *testing.T) {
	module, err := core.StringToEmi(intentsTestModuleYaml)
	if err != nil {
		t.Fatalf("StringToEmi error: %v", err)
	}

	files, err := GoModuleFull(&module, core.MicroGenContext{})
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}

	var intentsFile *core.VirtualFile
	for i := range files {
		if files[i].Name == "Intents" {
			intentsFile = &files[i]
			break
		}
	}
	if intentsFile == nil {
		t.Fatalf("expected an Intents file, got: %+v", fileNames(files))
	}
	got := intentsFile.ActualScript

	wantContains := []string{
		// Signature: takes the real handler function unchanged (same shape
		// abac.UserBrowseAction etc already have), mirroring how EditInvoiceActionGin
		// itself takes a `func(EditInvoiceActionRequest) (*EditInvoiceActionResponse, error)`.
		`func RegisterEditInvoiceIntentIntentTool(reg *emigo.ToolRegistry, handler func(EditInvoiceActionRequest) (*EditInvoiceActionResponse, error)) {`,
		// Query field (verbose) extracted from the call's own arguments into url.Values,
		// then stripped so it never leaks into the body unmarshal below.
		`if v, ok := args["verbose"]; ok {`,
		`queryParams.Set("verbose", fmt.Sprintf("%v", v))`,
		`delete(args, "verbose")`,
		// Path field (uniqueId) extracted into the real EditInvoiceActionPathParameter -
		// string-typed here (":uniqueId" has no type annotation), so a plain
		// fmt.Sprintf conversion, not a strconv parse.
		`var pathParams EditInvoiceActionPathParameter`,
		`pathParams.UniqueId = fmt.Sprintf("%v", v)`,
		// Whatever's left of args (the dto-shaped fields) becomes the real Body,
		// unmarshaled into the exact same type buildEntityUpdateAction-style dto
		// reference already names.
		`var body InvoicePatchDto`,
		// The real handler is called with a genuine EditInvoiceActionRequest, and its
		// response's Payload (not some intent-specific shape) is what gets returned.
		`resp, err := handler(EditInvoiceActionRequest{`,
		`b, err := json.Marshal(resp.Payload)`,
	}
	for _, want := range wantContains {
		if !strings.Contains(string(got), want) {
			t.Errorf("expected generated Intents.go to contain %q, got:\n%s", want, got)
		}
	}

	// searchInvoices has no From - there's no real Go function it could dispatch to,
	// so no registration helper should be generated for it at all.
	if strings.Contains(string(got), "RegisterSearchInvoicesIntentTool") {
		t.Errorf("searchInvoices has no From/backing action - it must not get a Register...IntentTool function")
	}

	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "Intents.go", string(got), parser.AllErrors); err != nil {
		t.Fatalf("generated Intents.go is not syntactically valid Go: %v\n\n%s", err, got)
	}
}
