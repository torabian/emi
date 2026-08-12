package python

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type pyQueryField struct {
	Name      string
	Type      string
	FieldLine string
}

func pythonQueryFieldType(query *core.EmiQueryField) string {
	switch query.Type {
	case core.FieldTypeArray, core.FieldTypeSlice:
		primitive := extractPrimitive(query.Primitive)
		if primitive == "" {
			primitive = "Any"
		}
		return fmt.Sprintf("List[%v]", primitive)
	default:
		t := extractPrimitive(string(query.Type))
		if t == "" {
			t = "Any"
		}
		return t
	}
}

// PythonActionQueryClass renders a dataclass out of an action's typed `qs:`
// definition, plus a `to_query_params()` helper returning a
// `Dict[str, Any]` with unset (`None`) fields stripped, ready to be handed
// straight to httpx's `params=`.
func PythonActionQueryClass(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	query := action.GetQuery()
	if len(query) == 0 {
		return nil, nil
	}

	fields := make([]pyQueryField, 0, len(query))
	for _, q := range query {
		typeHint := pythonQueryFieldType(q)
		fields = append(fields, pyQueryField{
			Name:      q.Name,
			Type:      typeHint,
			FieldLine: fmt.Sprintf("%v: Optional[%v] = None", q.Name, typeHint),
		})
	}

	className := fmt.Sprintf("%vQueryParams", core.ToUpper(core.NormaliseKey(action.GetName())))

	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "dataclasses", Objects: []string{"dataclass"}},
			{Location: "typing", Objects: []string{"Optional", "List", "Dict", "Any"}},
		},
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: className},
		},
	}

	const tmpl = `
@dataclass
class {{ .className }}:
{{- range .fields }}
    {{ .FieldLine }}
{{- end }}

    def to_query_params(self) -> Dict[str, Any]:
        result: Dict[str, Any] = {}
{{- range .fields }}
        if self.{{ .Name }} is not None:
            result["{{ .Name }}"] = self.{{ .Name }}
{{- end }}
        return result
`

	t := template.Must(template.New("pyqs").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{"className": className, "fields": fields}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToSnakeCase(className)
	res.SuggestedExtension = ".py"

	return res, nil
}
