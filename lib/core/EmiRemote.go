package core

import (
	"encoding/json"
	"strings"
)

// This is a Emi remote definition, you can make the external API calls typesafe using
// definitions. This feature is documented in docs/remotes.md
type EmiRemote struct {
	// Remote action name, it will become the Golang function that you will call
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Remote action name, it will become the Golang function that you will call"`

	// Standard HTTP methods
	Method string `yaml:"method,omitempty" json:"method,omitempty" jsonschema:"enum=get,enum=post,enum=put,enum=delete,enum=patch,enum=options,enum=head,description=Standard HTTP methods"`

	// The url which will be requested. You need to add full url here, but maybe you could add a prefix
	// also in the client from your Go code - There might be a prefix for remotes later version of Emi
	Url string `yaml:"url,omitempty" json:"url,omitempty" jsonschema:"description=The url which will be requested. You need to add full url here, but maybe you could add a prefix also in the client from your Go code - There might be a prefix for remotes later version of Emi"`

	// Standard EmiActionBody object. Could have fields, entity, dto as content and you
	// can define the output to cast automatically into them. If the response could be different objects, add them all
	// and create custom dtos and manually map them
	Out *EmiActionBody `yaml:"out,omitempty" json:"out,omitempty" jsonschema:"description=Standard EmiActionBody object. Could have fields, entity, dto as content and you can define the output to cast automatically into them. If the response could be different objects, add them all and create custom dtos and manually map them."`

	// Standard EmiActionBody object. Could have fields, entity, dto as content and you
	// can define the input parameters as struct in Go and Emi will convert it into
	// json.
	In *EmiActionBody `yaml:"in,omitempty" json:"in,omitempty" jsonschema:"description=Standard EmiActionBody object. Could have fields entity dto as content and you can define the input parameters as struct in Go and Emi will convert it into json."`

	// Query params for the address if you want to define them in Golang dynamically instead of URL
	Query []*EmiField `yaml:"query,omitempty" json:"query,omitempty" jsonschema:"description=Query params for the address if you want to define them in Golang dynamically instead of URL."`
}

func (x EmiRemote) GetName() string {
	return ToUpper(x.Name) + "Remote"
}

func (x EmiRemote) GetQuery() []*EmiField {
	return x.Query
}

func (x EmiRemote) GetUrl() string {
	return x.Url
}

func (x EmiRemote) GetMethod() string {
	return x.Method
}

func (x EmiRemote) HasRequest() bool {
	return x.In != nil
}

func (x EmiRemote) HasRequestHeaders() bool {
	return x.HasRequest() && len(x.In.Headers) > 0
}

func (x EmiRemote) HasRequestFields() bool {
	return x.HasRequest() && len(x.In.Fields) > 0
}

func (x EmiRemote) GetRequestFields() []*EmiField {
	if !x.HasRequestFields() {
		return []*EmiField{}
	}
	return x.In.Fields
}

func (x EmiRemote) GetRequestHeaders() []EmiHeader {
	if x.HasRequestHeaders() {
		return x.In.Headers
	}

	return []EmiHeader{}
}

func (x EmiRemote) HasResponse() bool {
	return x.Out != nil
}

func (x EmiRemote) HasResponseHeaders() bool {
	return x.HasResponse() && len(x.Out.Headers) > 0
}

func (x EmiRemote) GetResponseHeaders() []EmiHeader {
	if x.HasResponseHeaders() {
		return x.Out.Headers
	}

	return []EmiHeader{}
}

func (x EmiRemote) HasResponseDto() bool {
	return x.HasResponse() && x.Out.Dto != ""
}

func (x EmiRemote) GetResponseDto() string {
	return x.Out.Dto
}

func (x EmiRemote) GetResponseEnvelopeClass() string {
	if !x.HasResponse() {
		return ""
	}

	return x.Out.Envelope
}

func (x EmiRemote) HasRequestDto() bool {
	return x.HasRequest() && x.In.Dto != ""
}

func (x EmiRemote) GetRequestDto() string {
	return x.In.Dto
}

func (x EmiRemote) HasResponseFields() bool {
	return x.HasResponse() && len(x.Out.Fields) > 0
}

func (x EmiRemote) GetResponseFields() []*EmiField {
	if !x.HasResponseFields() {
		return []*EmiField{}
	}
	return x.Out.Fields
}

func (x EmiRemote) MethodUpper() string {
	return strings.ToUpper(x.Method)
}

func (x EmiRemote) Upper() string {
	return ToUpper(x.Name)
}

func (x EmiRemote) GetDefinition() string {
	m, _ := json.MarshalIndent(x, "", "  ")
	return string(m)
}
