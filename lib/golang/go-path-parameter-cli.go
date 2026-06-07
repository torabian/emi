package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoActionPathParamsCli(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	realms, err := GoActionPathParamsRealms(action, ctx)
	if err != nil {
		return nil, err
	}

	if realms == nil {
		return nil, nil
	}

	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: realms.Dependencies,
	}
	f := GetCommonFlags(ctx)

	if len(realms.Params) > 0 {
		res.CodeChunkDependensies = append(
			res.CodeChunkDependensies,
			core.CodeChunkDependency{
				Location: "github.com/urfave/cli/v3",
			},
			core.CodeChunkDependency{
				Location: f.Emigo,
			},
		)
	}

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: realms.TypeName})
	const tmpl = `


func Get{{ .TypeName }}CliFlags(prefix string) []emigo.CliFlag {

	return []emigo.CliFlag{
		{{ range .Params }}
		{
			Name: prefix + "pp-{{ .CliName }}",
			Type: "{{ .GolangType }}",
			Required: true,
		},
		{{ end }}
	}
}

// Extracts the path parameter from a urfave v3 cli.
func {{ .TypeName }}FromCli(c *cli.Command) {{ .TypeName }} {
	return {{ .TypeName }}FromFn(func (key string) string {

		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key) 
	})
}

`

	t := template.Must(template.New("pathParams").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, map[string]any{
		"TypeName": realms.TypeName,
		"Params":   realms.Params,
		"realms":   realms,
	}); err != nil {
		return nil, err
	}
	res.ActualScript = buf.Bytes()

	return res, nil
}
