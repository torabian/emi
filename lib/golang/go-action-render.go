package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	realms, deps, err := GoActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  core.TOKEN_ORIGINAL_NAME,
				Value: action.GetName(),
			},
		},
	}

	const tmpl = `/**
* Action to communicate with the action {{ .realms.Name }}
*/

type {{ .realms.Name }} struct {}

func ({{ .realms.Name}}) Meta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "{{ .realms.Name }}",
        URL:    "{{ .action.Url }}",
        Method: "{{ .action.Method }}",
    }
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

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = action.GetName()
	res.SuggestedExtension = ".js"
	res.CodeChunkDependensies = deps

	return res, nil
}
