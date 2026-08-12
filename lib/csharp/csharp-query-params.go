package csharp

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type csQueryField struct {
	Name         string // PascalCase C# property name
	WireName     string // the raw field.Name used in the query string
	Type         string // C# type, without `?`
	NullableType string
}

func csQueryFieldType(query *core.EmiQueryField) string {
	switch query.Type {
	case core.FieldTypeArray, core.FieldTypeSlice:
		primitive := extractPrimitive(query.Primitive)
		if primitive == "" {
			primitive = "string"
		}
		return fmt.Sprintf("List<%v>", primitive)
	default:
		t := extractPrimitive(string(query.Type))
		if t == "" {
			t = "string"
		}
		return t
	}
}

var csQueryTmplFuncs = template.FuncMap{"csString": escapeDoubleQuoted}

var csQueryTmpl = template.Must(template.New("csqs").Funcs(csQueryTmplFuncs).Parse(`
public class {{ .className }}
{
{{- range .fields }}
    public {{ .NullableType }} {{ .Name }} { get; set; }
{{- end }}

    public Dictionary<string, object> ToQueryParams()
    {
        var result = new Dictionary<string, object>();
{{- range .fields }}
        if ({{ .Name }} != null)
        {
            result[{{ .WireName | csString }}] = {{ .Name }};
        }
{{- end }}
        return result;
    }
}
`))

// CSharpActionQueryClass renders a class out of an action's typed `qs:`
// definition, plus a ToQueryParams() helper returning a Dictionary<string,
// object> with unset (null) fields stripped, ready to be handed straight to
// the fetchx runtime's query encoder.
func CSharpActionQueryClass(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	query := action.GetQuery()
	if len(query) == 0 {
		return nil, nil
	}

	fields := make([]csQueryField, 0, len(query))
	for _, q := range query {
		t := csQueryFieldType(q)
		fields = append(fields, csQueryField{Name: core.NormaliseKey(q.Name), WireName: q.Name, Type: t, NullableType: csNullable(t)})
	}

	className := fmt.Sprintf("%vQueryParams", core.ToUpper(core.NormaliseKey(action.GetName())))

	var buf bytes.Buffer
	if err := csQueryTmpl.Execute(&buf, core.H{"className": className, "fields": fields}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  className,
		SuggestedExtension: ".cs",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: className},
		},
	}, nil
}
