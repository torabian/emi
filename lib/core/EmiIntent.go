package core

// EmiIntent is a model-facing tool signature - the unit Emi compiles into MCP
// (Model Context Protocol) server tool registrations (name/description/input schema/
// output schema/behavior hints), the same way EmiAction compiles into a REST route or
// CLI command. An intent either stands alone (In/Out declared inline, same shape as an
// EmiAction's) or is derived from an existing action via From, so exposing an action as
// a tool needs no field duplication - see preprocessIntents.
type EmiIntent struct {

	// Tool name exposed to the model, e.g. "searchInvoices". Defaults to From when From
	// is set and Name is left empty.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Tool name exposed to the model e.g. searchInvoices. Defaults to From when From is set and Name is left empty."`

	// Name of an existing top-level EmiAction this intent is derived from. When set,
	// Description/In/Out are copied from that action wherever the intent doesn't
	// already declare its own value - see preprocessIntents.
	From string `yaml:"from,omitempty" json:"from,omitempty" jsonschema:"description=Name of an existing EmiAction to derive description/in/out from. The intent's own Description/In/Out (if set) always take precedence over the action's."`

	// Human readable title, surfaced as the MCP tool's 'title' annotation in clients/UIs.
	Title string `yaml:"title,omitempty" json:"title,omitempty" jsonschema:"description=Human-readable title surfaced as the MCP tool's 'title' annotation in clients/UIs."`

	// Description sent verbatim to the model - the most load-bearing field, since this
	// is what a planning LLM reads to decide whether/how to call the tool.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Tool description sent verbatim to the model - write it for the model not for a developer. When From is set and this is empty the source action's Description is used instead."`

	// Request/arguments shape; becomes the tool's inputSchema.
	In *EmiActionBody `yaml:"in,omitempty" json:"in,omitempty" jsonschema:"description=Becomes the tool's inputSchema. When From is set and this is nil the source action's In is used instead."`

	// Response shape; becomes the tool's outputSchema/structuredContent.
	Out *EmiActionBody `yaml:"out,omitempty" json:"out,omitempty" jsonschema:"description=Becomes the tool's outputSchema/structuredContent shape. When From is set and this is nil the source action's Out is used instead."`

	// MCP tool annotation hints.
	Annotations *EmiIntentAnnotations `yaml:"annotations,omitempty" json:"annotations,omitempty" jsonschema:"description=MCP tool behavior-hint annotations passed straight through to the tool registration."`

	// Ties into the existing EmiPermission tree; unauthorized callers don't see the
	// tool listed.
	Permission string `yaml:"permission,omitempty" json:"permission,omitempty" jsonschema:"description=Permission key from this module's Permissions tree required to see/call this tool."`
}

// EmiIntentAnnotations mirrors the MCP spec's standard tool annotation hints
// (https://modelcontextprotocol.io) - purely advisory metadata a client may use to
// decide how to surface or gate the tool (e.g. asking for confirmation before a
// destructive call).
type EmiIntentAnnotations struct {
	// Tool does not modify its environment.
	ReadOnlyHint bool `yaml:"readOnlyHint,omitempty" json:"readOnlyHint,omitempty" jsonschema:"description=Tool does not modify its environment."`

	// Tool may perform destructive updates (only meaningful when ReadOnlyHint is false).
	DestructiveHint bool `yaml:"destructiveHint,omitempty" json:"destructiveHint,omitempty" jsonschema:"description=Tool may perform destructive updates (only meaningful when ReadOnlyHint is false)"`

	// Calling the tool repeatedly with the same arguments has no additional effect.
	IdempotentHint bool `yaml:"idempotentHint,omitempty" json:"idempotentHint,omitempty" jsonschema:"description=Calling the tool repeatedly with the same arguments has no additional effect."`

	// Tool interacts with an open-ended external world (e.g. web search) rather than a
	// closed enumerable domain.
	OpenWorldHint bool `yaml:"openWorldHint,omitempty" json:"openWorldHint,omitempty" jsonschema:"description=Tool interacts with an open-ended external world (e.g. web search) rather than a closed enumerable domain."`
}

// Upper returns the intent's name in the same UpperCamelCase convention every other
// generated Go identifier in this package uses (see core.ToUpper).
func (x *EmiIntent) Upper() string {
	if x == nil {
		return ""
	}
	return ToUpper(x.Name)
}

// GetName returns the generated-code identifier prefix for this intent, e.g.
// "SearchInvoicesIntent" - analogous to EmiAction.GetName's "...Action" suffix.
func (x *EmiIntent) GetName() string {
	return x.Upper() + "Intent"
}

// HasIn reports whether the intent declares a request/arguments body.
func (x *EmiIntent) HasIn() bool {
	return x != nil && x.In != nil
}

// HasOut reports whether the intent declares a response body.
func (x *EmiIntent) HasOut() bool {
	return x != nil && x.Out != nil
}
