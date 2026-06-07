package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoActionRenderReactive(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) ([]*core.CodeChunkCompiled, error) {

	realms, deps, err := GoActionReactiveRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  core.TOKEN_ORIGINAL_NAME,
				Value: realms.ActionName,
			},
		},
	}

	const tmpl = `/**
* Action to communicate with the action {{ .realms.ActionName }}
*/
 

func {{ .realms.ActionName }}Meta() struct {
    Name   string
    URL    string
    Method string
	CliName string
	Description string
} {
    return struct {
        Name   string
        URL    string
        Method string
		CliName string
		Description string
    }{
        Name:   "{{ .realms.ActionName }}",
        URL:    "{{ .realms.SafeUrl }}",
        Method: "{{ UPPER .action.Method }}",
		CliName: "{{ .action.CliName }}",
		Description: "{{ .action.Description }}",
    }
}

{{ if .realms.PathParameter }}
	{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}

{{ if .realms.QueryParams }}
	{{ b2s .realms.QueryParams.ActualScript }}
{{ end }}



type {{ .realms.ActionName }}Message struct {
	Raw []byte
	// Conn *websocket.Conn	
	Conn interface{}
	MessageType int
	Error error
	{{ if .realms.PathParameter }}
	PathParams {{ .realms.ActionName }}PathParameter
	{{ end}}
}

// Developer handler type
type {{ .realms.ActionName }}Handler func(msg {{ .realms.ActionName }}Message ) error

type {{ .realms.ActionName }}Session struct {
	// Ctx    *gin.Context
	// Socket *websocket.Conn
	Ctx    interface{}
	Socket interface{}
	Done   chan bool
	Read   chan {{ .realms.ActionName }}ReadChan
	QueryParams {{ .realms.ActionName }}Query
}

type {{ .realms.ActionName }}HandlerDuplex func(*{{ .realms.ActionName }}Session)

type {{ .realms.ActionName }}ReadChan struct {
	Data  []byte
	Error error
	MessageType int
}
 


// {{ .realms.ActionName }}ClientSession is the client-side mirror of
// {{ .realms.ActionName }}Session. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type {{ .realms.ActionName }}ClientSession struct {
	// Socket *websocket.Conn
	Socket interface{}
	Done   chan bool
	Read   chan {{ .realms.ActionName }}ReadChan
	Write  chan []byte
}

`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":       action,
		"realms":       realms,
		"shouldExport": true,
	}); err != nil {
		return nil, err
	}

	outputs := []*core.CodeChunkCompiled{}
	splitCli := ctx.HasTag(SplitCli)
	enabledGin := !ctx.HasTag(SkipGin)

	if enabledGin {

		reactiveGin, ginErr := GoActionRenderReactiveGin(action, ctx, complexes)
		if ginErr != nil {
			return nil, ginErr
		}
		if reactiveGin != nil {
			outputs = append(outputs, reactiveGin)
		}
	}

	reactiveCli, ginErr := GoActionReactiveCliRender(action, ctx, realms)
	if ginErr != nil {
		return nil, ginErr
	}

	if splitCli {
		if reactiveCli != nil {
			outputs = append(outputs, reactiveCli)
		}
	} else {
		res.ActualScript = append(res.ActualScript, reactiveCli.ActualScript...)
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, reactiveCli.CodeChunkDependensies...)
	}

	reactiveWasm, ginErr := GoActionRenderReactiveWasm(action, ctx, complexes)
	if ginErr != nil {
		return nil, ginErr
	}
	if reactiveWasm != nil {
		outputs = append(outputs, reactiveWasm)
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".go"
	res.CodeChunkDependensies = deps

	outputs = append(outputs, res)

	return outputs, nil
}
