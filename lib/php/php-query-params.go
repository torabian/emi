package php

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type phpQueryField struct {
	Name         string
	Type         string
	NullableType string // `?Type`, unless Type is already `mixed` (PHP forbids `?mixed`)
}

func phpNullableType(t string) string {
	if t == "mixed" {
		return t
	}
	return "?" + t
}

func phpQueryFieldType(query *core.EmiQueryField) string {
	switch query.Type {
	case core.FieldTypeArray, core.FieldTypeSlice:
		return "array"
	default:
		t := extractPrimitive(string(query.Type))
		if t == "" {
			t = "mixed"
		}
		return t
	}
}

var phpQueryTmplFuncs = template.FuncMap{"phpString": escapeDoubleQuoted}

var phpQueryTmpl = template.Must(template.New("phpqs").Funcs(phpQueryTmplFuncs).Parse(`
class {{ .className }}
{
{{- range .fields }}
    public {{ .NullableType }} ${{ .Name }} = null;
{{- end }}

    public function toQueryParams(): array
    {
        $result = [];
{{- range .fields }}
        if ($this->{{ .Name }} !== null) {
            $result[{{ .Name | phpString }}] = $this->{{ .Name }};
        }
{{- end }}
        return $result;
    }
}
`))

// PhpActionQueryClass renders a class out of an action's typed `qs:`
// definition, plus a toQueryParams() helper returning an associative array
// with unset (null) fields stripped, ready to be handed straight to the
// fetchx runtime's query encoder.
func PhpActionQueryClass(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	query := action.GetQuery()
	if len(query) == 0 {
		return nil, nil
	}

	fields := make([]phpQueryField, 0, len(query))
	for _, q := range query {
		t := phpQueryFieldType(q)
		fields = append(fields, phpQueryField{Name: q.Name, Type: t, NullableType: phpNullableType(t)})
	}

	className := fmt.Sprintf("%vQueryParams", core.ToUpper(core.NormaliseKey(action.GetName())))

	var buf bytes.Buffer
	if err := phpQueryTmpl.Execute(&buf, core.H{"className": className, "fields": fields}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  className,
		SuggestedExtension: ".php",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: className},
		},
	}, nil
}
