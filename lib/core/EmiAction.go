package core

import "strings"

type EmiAction struct {

	// Type of the emi content
	Emi string `jsonschema:"description=Type of the emi content.;enum=action" json:"emi" yaml:"emi"`

	// General name of the action used for generating code and CLI commands.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=General name of the action used for generating code and CLI commands"`

	// Overrides the CLI action name if specified otherwise defaults to Name.
	CliName string `yaml:"cliName,omitempty" json:"cliName,omitempty" jsonschema:"description=Overrides the CLI action name if specified otherwise defaults to Name"`

	// CLI command aliases for shorter action names.
	ActionAliases []string `yaml:"actionAliases,omitempty" json:"actionAliases,omitempty" jsonschema:"description=CLI command aliases for shorter action names"`

	// HTTP route of the action; if not specified the action is CLI-only.
	Url string `yaml:"url,omitempty" json:"url,omitempty" jsonschema:"description=HTTP route of the action; if not specified the action is CLI-only"`

	// HTTP method type including standard and Emi-specific methods.
	Method string `yaml:"method,omitempty" json:"method,omitempty" jsonschema:"enum=post,enum=patch,enum=put,enum=get,enum=delete,enum=webrtc,enum=reactive,enum=sse,description=HTTP method type including standard and Emi-specific methods"`

	// Text by default for websocket, can be changed to arraybuffer or blob.
	BinaryType string `yaml:"binaryType,omitempty" json:"binaryType,omitempty" jsonschema:"enum=text,enum=arraybuffer,enum=blob,description=Text by default for websocket, can be changed to arraybuffer or blob"`

	// Type-safe query strings for action
	Query []*EmiField `yaml:"qs,omitempty" json:"qs,omitempty" jsonschema:"description=Type-safe query parameters for CLI and HTTP requests"`

	// Data channels in a typesafe mode in case of webrtc
	DataChannels []EmiWebRtcDataChannel `yaml:"dataChannels,omitempty" json:"dataChannels,omitempty" jsonschema:"description=Data channels in a typesafe mode in case of webrtc"`

	// Action description used in API specs and documentation.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Action description used in API specs and documentation"`

	// Higher-level request format such as POST_ONE PATCH_ONE, QUERY, and PATCH_BULK.
	Format string `yaml:"format,omitempty" json:"format,omitempty" jsonschema:"enum=reactive,enum=query,description=Higher-level request format such as POST_ONE PATCH_ONE, QUERY, and PATCH_BULK"`

	// Request body definition similar to HTTP request body.
	In *EmiActionBody `yaml:"in,omitempty" json:"in,omitempty" jsonschema:"description=Request body definition similar to HTTP request body"`

	// Response body definition similar to HTTP response body.
	Out *EmiActionBody `yaml:"out,omitempty" json:"out,omitempty" jsonschema:"description=Response body definition similar to HTTP response body"`

	// Defines access control similar to middleware checking permissions, tokens, and roles.
	SecurityModel *SecurityModel `yaml:"security,omitempty" json:"security,omitempty" jsonschema:"description=Defines access control similar to middleware checking permissions, tokens, and roles"`

	// External function name used in generated code.
	ExternFuncName string `yaml:"-" json:"-" jsonschema:"-"`

	// Struct representing the request body used for generating RPC code.
	RequestEntity any `yaml:"-" json:"-" jsonschema:"-"`

	// Struct representing the response body used for generating RPC code.
	ResponseEntity any `yaml:"-" json:"-" jsonschema:"-"`

	// Function representing the action's implementation.
	Action any `yaml:"-" json:"-" jsonschema:"-"`

	// Pointer to the struct representing the entity being operated on.
	TargetEntity any `yaml:"-" json:"-" jsonschema:"-"`

	// Internal metadata for code generation.
	RootModule *Emi `yaml:"-" json:"-" jsonschema:"-"`

	// How the qs is being handled.
	QsMode string `yaml:"qsMode,omitempty" json:"qsMode,omitempty" jsonschema:"enum=reflect"`
}

func (x EmiAction) MethodUpper() string {
	return strings.ToUpper(x.Method)
}

func (x EmiAction) QSFields() []*EmiField {

	// At the moment, the query string fields is only supported for query
	if x.Format != "query" {
		return nil
	}

	if x.QsMode == "reflect" && x.Out != nil && x.Out.Fields != nil {
		return x.Out.Fields
	}

	if x.In != nil && x.In.Fields != nil {
		return x.In.Fields
	}

	return nil
}

func (x EmiAction) GetName() string {
	return ToUpper(x.Name) + "Action"
}

func (x EmiAction) GetQuery() []*EmiField {
	return x.Query
}

func (x EmiAction) GetUrl() string {
	return x.Url
}

func (x EmiAction) GetMethod() string {
	return x.Method
}

func (x EmiAction) HasRequest() bool {
	return x.In != nil
}

func (x EmiAction) HasRequestHeaders() bool {
	return x.HasRequest() && len(x.In.Headers) > 0
}

func (x EmiAction) HasRequestFields() bool {
	return x.HasRequest() && len(x.In.Fields) > 0
}

func (x EmiAction) GetRequestFields() []*EmiField {
	if !x.HasRequestFields() {
		return []*EmiField{}
	}
	return x.In.Fields
}

func (x EmiAction) GetRequestHeaders() []EmiHeader {
	if x.HasRequestHeaders() {
		return x.In.Headers
	}

	return []EmiHeader{}
}

func (x EmiAction) HasResponse() bool {
	return x.Out != nil
}

func (x EmiAction) HasResponseDto() bool {
	return x.HasResponse() && x.Out.Dto != ""
}

func (x EmiAction) GetResponseDto() string {
	return x.Out.Dto
}

func (x EmiAction) HasRequestDto() bool {
	return x.HasRequest() && x.In.Dto != ""
}

func (x EmiAction) GetRequestDto() string {
	return x.In.Dto
}

func (x EmiAction) HasResponseHeaders() bool {
	return x.HasResponse() && len(x.Out.Headers) > 0
}

func (x EmiAction) GetResponseEnvelopeClass() string {
	if !x.HasResponse() {
		return ""
	}

	return x.Out.Envelope
}

func (x EmiAction) GetResponseHeaders() []EmiHeader {
	if x.HasResponseHeaders() {
		return x.Out.Headers
	}

	return []EmiHeader{}
}

func (x EmiAction) HasResponseFields() bool {
	return x.HasResponse() && len(x.Out.Fields) > 0
}

func (x EmiAction) GetResponseFields() []*EmiField {
	if !x.HasResponseFields() {
		return []*EmiField{}
	}
	return x.Out.Fields
}
