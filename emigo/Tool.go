package emigo

import (
	"context"
	"encoding/json"
	"fmt"
)

// Tool is one function a model may call - the transport-agnostic shape every
// tool-calling backend (a provider's native "function calling"/"tool use" feature,
// a real MCP server, ...) gets built from/adapted to. Generated per EmiIntent (see
// GoIntentsGenerate in the golang compiler backend) rather than hand-written, the
// same way an HTTP route or CLI command already is - InputSchema/OutputSchema are
// JSON Schema objects.
type Tool struct {
	Name         string
	Description  string
	InputSchema  json.RawMessage
	OutputSchema json.RawMessage
}

// ToolCall is one invocation a model asked for. ID is whatever caller-assigned
// identifier ties this call to its eventual result (a provider's own tool_use id,
// or an MCP JSON-RPC request id) - ToolRegistry itself never needs it, it's carried
// through purely for the handler/caller's own bookkeeping.
type ToolCall struct {
	ID    string
	Name  string
	Input json.RawMessage
}

// ToolHandlerFn executes one ToolCall, given the same EmiRequestContexts the
// surrounding action/connection was already authorized through - so a handler can
// call straight into an existing generated action (e.g. the real Go function behind
// a browse/create/update route) exactly as it would over that action's own HTTP
// route, with no separate auth path of its own. Returns the tool result content (a
// JSON-encoded string is fine) and whether it represents a failed call rather than a
// real answer.
type ToolHandlerFn func(ctx context.Context, req EmiRequestContexts, call ToolCall) (content string, isError bool)

// ToolRegistry collects Tool/ToolHandlerFn pairs. The per-intent
// Register<Name>IntentTool functions GoIntentsGenerate emits each take one of these
// and register themselves into it - so a project wires up its whole tool list with
// one call per intent (mirroring how a generated <Action>Gin/<Action>Cli function
// wires up one HTTP route/CLI command per action, taking the real handler function
// as its only argument) instead of hand-assembling a []Tool slice, a JSON-unmarshal
// step per tool, and a switch statement dispatching by name.
type ToolRegistry struct {
	tools    []Tool
	handlers map[string]ToolHandlerFn
}

// NewToolRegistry returns an empty ToolRegistry, ready for Register calls.
func NewToolRegistry() *ToolRegistry {
	return &ToolRegistry{handlers: map[string]ToolHandlerFn{}}
}

// Register adds tool to the registry's tool list and binds handler to its Name -
// registering a second tool under the same Name replaces the first one's handler
// (and appends a second, now-shadowed entry to Tools() - callers shouldn't rely on
// this beyond "the last registration for a given name wins on Execute").
func (r *ToolRegistry) Register(tool Tool, handler ToolHandlerFn) {
	if r == nil {
		return
	}
	r.tools = append(r.tools, tool)
	if r.handlers == nil {
		r.handlers = map[string]ToolHandlerFn{}
	}
	r.handlers[tool.Name] = handler
}

// Tools returns every tool registered so far, in registration order - typically
// passed straight through to whatever config field a tool-calling module exposes
// (e.g. an ai module's AiModuleConfig.Tools).
func (r *ToolRegistry) Tools() []Tool {
	if r == nil {
		return nil
	}
	return r.tools
}

// Execute dispatches call to whichever handler was registered under call.Name -
// matches ToolHandlerFn's own signature, so a *ToolRegistry can be passed directly
// wherever a single ToolHandlerFn is expected (e.g. an ai module's
// AiModuleConfig.ExecuteTool). An unregistered call.Name is reported back as a
// failed tool result (isError true) rather than a panic or Go error, consistent
// with every other handler-side failure this returns.
func (r *ToolRegistry) Execute(ctx context.Context, req EmiRequestContexts, call ToolCall) (string, bool) {
	if r == nil {
		return fmt.Sprintf("unknown tool %q", call.Name), true
	}
	h, ok := r.handlers[call.Name]
	if !ok {
		return fmt.Sprintf("unknown tool %q", call.Name), true
	}
	return h(ctx, req, call)
}
