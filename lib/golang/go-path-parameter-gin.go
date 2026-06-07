package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoActionPathParamsGin(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
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
	// f := GetCommonFlags(ctx)

	// if len(realms.Params) > 0 {
	// 	res.CodeChunkDependensies = append(
	// 		res.CodeChunkDependensies,
	// 		core.CodeChunkDependency{
	// 			Location: "github.com/urfave/cli/v3",
	// 		},
	// 		core.CodeChunkDependency{
	// 			Location: f.Emigo,
	// 		},
	// 	)
	// }

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: realms.TypeName})
	const tmpl = `

func {{ .TypeName }}FromGin(g *gin.Context) {{ .TypeName }} {
	return {{ .TypeName }}FromFn(func (key string) string {
		return g.Param(key) 
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
