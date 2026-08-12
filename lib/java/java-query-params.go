package java

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type javaQueryField struct {
	Name string
	Type string
}

func javaQueryFieldType(query *core.EmiQueryField) string {
	switch query.Type {
	case core.FieldTypeArray, core.FieldTypeSlice:
		primitive := extractPrimitive(query.Primitive)
		if primitive == "" {
			primitive = "Object"
		}
		return fmt.Sprintf("List<%v>", primitive)
	default:
		t := extractPrimitive(string(query.Type))
		if t == "" {
			t = "Object"
		}
		return t
	}
}

var javaQueryTmplFuncs = template.FuncMap{"javaString": escapeDoubleQuoted}

var javaQueryTmpl = template.Must(template.New("javaqs").Funcs(javaQueryTmplFuncs).Parse(`
public class {{ .className }} {
{{- range .fields }}
    public {{ .Type }} {{ .Name }};
{{- end }}

    public java.util.Map<String, Object> toQueryParams() {
        var result = new java.util.HashMap<String, Object>();
{{- range .fields }}
        if ({{ .Name }} != null) {
            result.put({{ .Name | javaString }}, {{ .Name }});
        }
{{- end }}
        return result;
    }
}
`))

// JavaActionQueryClass renders a (package-private) class out of an action's
// typed `qs:` definition, plus a toQueryParams() helper returning a
// Map<String, Object> with unset (null) fields stripped, ready to be handed
// straight to the fetchx runtime's query encoder.
func JavaActionQueryClass(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	query := action.GetQuery()
	if len(query) == 0 {
		return nil, nil
	}

	fields := make([]javaQueryField, 0, len(query))
	for _, q := range query {
		fields = append(fields, javaQueryField{Name: q.Name, Type: javaQueryFieldType(q)})
	}

	className := fmt.Sprintf("%vQueryParams", core.ToUpper(core.NormaliseKey(action.GetName())))

	var buf bytes.Buffer
	if err := javaQueryTmpl.Execute(&buf, core.H{"className": className, "fields": fields}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  className,
		SuggestedExtension: ".java",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: className},
		},
	}, nil
}
