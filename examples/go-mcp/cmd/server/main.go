// Command gomcp-server is a minimal MCP (Model Context Protocol) server built on the
// emi `intents` section: internal/defs/Intents.go was generated straight from
// ../../gomcp.emi.yml (see the "def" Makefile target - `make def`) and gives us the
// name/description/annotations/typed Args+Result for each tool; this file is the
// hand-written half - the actual math and the wiring into a real *mcp.Server - the
// same division of labor an emi action gets between its generated file and a
// developer-written Impl (see examples/in-browser-server/internal/ChatActionImpl.go).
package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/torabian/emi/examples/go-mcp/internal/defs"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "gomcp-example",
		Version: "1.0.0",
	}, nil)

	registerAdd(server)
	registerIsPrime(server)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}

// registerAdd wires the generated "add" intent into the server: AddIntentMeta gives
// the name/title/description/annotations, AddIntentArgs/AddIntentResult give the
// typed shapes the SDK infers a JSON Schema from - only the actual addition below is
// something emi could never have generated.
func registerAdd(server *mcp.Server) {
	meta := defs.AddIntentMeta()

	mcp.AddTool(server, &mcp.Tool{
		Name:        meta.Name,
		Title:       meta.Title,
		Description: meta.Description,
		Annotations: &mcp.ToolAnnotations{
			DestructiveHint: &meta.DestructiveHint,
			ReadOnlyHint:    meta.ReadOnlyHint,
			IdempotentHint:  meta.IdempotentHint,
			OpenWorldHint:   &meta.OpenWorldHint,
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest, args defs.AddIntentArgs) (*mcp.CallToolResult, defs.AddIntentResult, error) {
		// A is optional (declared "float64?" in gomcp.emi.yml, hence emigo.Nullable) -
		// missing means "add zero".
		result := defs.AddIntentResult{Sum: args.A.OrDefault(0) + args.B}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: result.Json()}},
		}, result, nil
	})
}

// registerIsPrime wires the generated "isPrime" intent the same way - trial division
// is plenty for a smoke test; a real module would swap in whatever it actually needs.
func registerIsPrime(server *mcp.Server) {
	meta := defs.IsPrimeIntentMeta()

	mcp.AddTool(server, &mcp.Tool{
		Name:        meta.Name,
		Title:       meta.Title,
		Description: meta.Description,
		Annotations: &mcp.ToolAnnotations{
			DestructiveHint: &meta.DestructiveHint,
			ReadOnlyHint:    meta.ReadOnlyHint,
			IdempotentHint:  meta.IdempotentHint,
			OpenWorldHint:   &meta.OpenWorldHint,
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest, args defs.IsPrimeIntentArgs) (*mcp.CallToolResult, defs.IsPrimeIntentResult, error) {
		result := defs.IsPrimeIntentResult{Prime: isPrime(args.N)}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: result.Json()}},
		}, result, nil
	})
}

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
