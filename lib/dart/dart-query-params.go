package dart

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type dartQueryField struct {
	Name         string
	Type         string // dart type, without `?`
	NullableType string // dart type, with a trailing `?` unless already `dynamic`
}

func dartNullableType(t string) string {
	if t == "dynamic" {
		return t
	}
	return t + "?"
}

func dartQueryFieldType(query *core.EmiQueryField) string {
	switch query.Type {
	case core.FieldTypeArray, core.FieldTypeSlice:
		primitive := extractPrimitive(query.Primitive)
		if primitive == "" {
			primitive = "dynamic"
		}
		return fmt.Sprintf("List<%v>", primitive)
	default:
		t := extractPrimitive(string(query.Type))
		if t == "" {
			t = "dynamic"
		}
		return t
	}
}

var dartQueryTmplFuncs = template.FuncMap{"dartJsonKey": escapeDoubleQuoted}

var dartQueryTmpl = template.Must(template.New("dartqs").Funcs(dartQueryTmplFuncs).Parse(`
class {{ .className }} {
{{- range .fields }}
  {{ .NullableType }} {{ .Name }};
{{- end }}

  {{ .className }}({
{{- range .fields }}
    this.{{ .Name }},
{{- end }}
  });

  Map<String, dynamic> toQueryParams() {
    final result = <String, dynamic>{};
{{- range .fields }}
    if ({{ .Name }} != null) {
      result[{{ .Name | dartJsonKey }}] = {{ .Name }};
    }
{{- end }}
    return result;
  }
}
`))

// DartActionQueryClass renders a class out of an action's typed `qs:`
// definition, plus a `toQueryParams()` helper returning a `Map<String,
// dynamic>` with unset (null) fields stripped, ready to be handed straight
// to the fetchx runtime's query encoder.
func DartActionQueryClass(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	query := action.GetQuery()
	if len(query) == 0 {
		return nil, nil
	}

	fields := make([]dartQueryField, 0, len(query))
	for _, q := range query {
		t := dartQueryFieldType(q)
		fields = append(fields, dartQueryField{Name: q.Name, Type: t, NullableType: dartNullableType(t)})
	}

	className := fmt.Sprintf("%vQueryParams", core.ToUpper(core.NormaliseKey(action.GetName())))

	var buf bytes.Buffer
	if err := dartQueryTmpl.Execute(&buf, core.H{"className": className, "fields": fields}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  core.ToSnakeCase(className),
		SuggestedExtension: ".dart",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: className},
		},
	}, nil
}
