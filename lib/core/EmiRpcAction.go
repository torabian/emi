package core

// Use for both remotes and actions to generate the same code
type EmiRpcAction interface {
	GetName() string
	GetUrl() string
	GetMethod() string
	MethodUpper() string
	Upper() string

	GetQuery() []*EmiQueryField

	HasRequest() bool
	HasRequestHeaders() bool
	GetRequestHeaders() []EmiHeader

	// Convert the definition to json
	GetDefinition() string

	// Checks if the request has inline fields definition
	HasRequestFields() bool
	GetRequestFields() []*EmiField

	HasResponse() bool
	HasRequestDto() bool
	GetRequestDto() string
	HasResponseDto() bool
	GetResponseDto() string

	GetResponseEnvelopeClass() string

	HasResponseHeaders() bool
	GetResponseHeaders() []EmiHeader

	// Checks if the response has inline fields definition
	HasResponseFields() bool
	GetResponseFields() []*EmiField
}
