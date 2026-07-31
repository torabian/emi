package golang

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type goHeaderCliField struct {
	CliKey      string // flag key without prefix, e.g. "header-x-api-key"
	HeaderName  string // canonical outgoing header name, e.g. "x-api-key"
	Type        string
	Description string
}

func headerFieldCliName(name string) string {
	return strings.ReplaceAll(core.ToSnakeCase(name), "_", "-")
}

// GoActionHeadersCli generates typed CLI flags for an action's typesafe request headers
// (in.headers), plus a {Action}HeadersFromCli(c *cli.Command) helper that builds the same
// http.Header the Gin/http transports populate {Action}Request.Headers with from a live
// *http.Request.
//
// Complex headers are intentionally skipped: on the wire every header is a string, so
// there is no generic, safe way to build an arbitrary complex value out of a single CLI
// flag. Developers needing that can still set it via the generic --header/-H flag that
// {Action}CliHandler always registers.
func GoActionHeadersCli(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	headers := action.GetRequestHeaders()
	if len(headers) == 0 {
		return nil, nil
	}

	fields := make([]goHeaderCliField, 0, len(headers))
	for _, h := range headers {
		if h.Type == "complex" || h.Type == "" {
			continue
		}
		fields = append(fields, goHeaderCliField{
			CliKey:      "header-" + headerFieldCliName(h.Name),
			HeaderName:  h.Name,
			Type:        h.Type,
			Description: h.Description,
		})
	}

	if len(fields) == 0 {
		return nil, nil
	}

	actionName := core.ToUpper(core.NormaliseKey(action.GetName()))
	f := GetCommonFlags(ctx)

	const tmpl = `
func Get{{ .ActionName }}HeaderCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{{ range .Fields }}
		{
			Name: prefix + "{{ .CliKey }}",
			Type: "{{ .Type }}",
			{{ if .Description }}
			Description: {{ escapeBackTick .Description }},
			{{ end }}
		},
		{{ end }}
	}
}

// {{ .ActionName }}HeadersFromCli builds the same http.Header the Gin/http transports
// populate {{ .ActionName }}Request.Headers with, but reads values off urfave v3 CLI flags.
func {{ .ActionName }}HeadersFromCli(c *cli.Command) http.Header {
	headers := http.Header{}

	{{ range .Fields }}
	if c.IsSet("{{ .CliKey }}") {
		{{ if eq .Type "bool" }}
		headers.Set("{{ .HeaderName }}", strconv.FormatBool(c.Bool("{{ .CliKey }}")))
		{{ else if eq .Type "int64" }}
		headers.Set("{{ .HeaderName }}", strconv.FormatInt(c.Int64("{{ .CliKey }}"), 10))
		{{ else if eq .Type "float64" }}
		headers.Set("{{ .HeaderName }}", strconv.FormatFloat(c.Float64("{{ .CliKey }}"), 'f', -1, 64))
		{{ else }}
		headers.Set("{{ .HeaderName }}", c.String("{{ .CliKey }}"))
		{{ end }}
	}
	{{ end }}

	return headers
}
`

	t := template.Must(template.New("headersCli").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, map[string]any{
		"ActionName": actionName,
		"Fields":     fields,
	}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript: buf.Bytes(),
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "github.com/urfave/cli/v3"},
			{Location: f.Emigo},
			{Location: "net/http"},
			{Location: "strconv"},
		},
	}, nil
}
