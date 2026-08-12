package dart

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type dartRenderedHeader struct {
	FieldName    string
	HeaderName   string
	Type         string // dart type, without `?`
	ParseExpr    string // expression parsing the raw string `v` into the typed value
	SerializeStr string // expression serializing the typed value `v` back to string
}

func dartHeaderType(headerType string) string {
	switch headerType {
	case "int64", "int32", "int":
		return "int"
	case "float64", "float32":
		return "double"
	case "bool":
		return "bool"
	default:
		return "String"
	}
}

func renderDartHeaders(columns []core.EmiHeader) []dartRenderedHeader {
	out := make([]dartRenderedHeader, 0, len(columns))
	for _, h := range columns {
		fieldName := core.NormaliseKey(h.Name)
		fieldName = core.ToLower(fieldName)
		typeHint := dartHeaderType(h.Type)
		if h.Complex != "" {
			typeHint = h.Complex
		}

		parseExpr, serializeStr := "v", "v"
		switch typeHint {
		case "int":
			parseExpr, serializeStr = "int.parse(v)", "v.toString()"
		case "double":
			parseExpr, serializeStr = "double.parse(v)", "v.toString()"
		case "bool":
			parseExpr, serializeStr = `v.trim().toLowerCase() == 'true' || v.trim() == '1'`, `v ? 'true' : 'false'`
		}

		out = append(out, dartRenderedHeader{
			FieldName:    fieldName,
			HeaderName:   h.Name,
			Type:         typeHint,
			ParseExpr:    parseExpr,
			SerializeStr: serializeStr,
		})
	}
	return out
}

var dartHeaderTmplFuncs = template.FuncMap{"dartJsonKey": escapeDoubleQuoted}

var dartHeaderTmpl = template.Must(template.New("dartheaders").Funcs(core.CommonMap).Funcs(dartHeaderTmplFuncs).Parse(`
class {{ .className }} {
{{- range .headers }}
  {{ .Type }}? {{ .FieldName }};
{{- end }}

  {{ .className }}({
{{- range .headers }}
    this.{{ .FieldName }},
{{- end }}
  });

  Map<String, String> toHttpHeaders() {
    final result = <String, String>{};
{{- range .headers }}
    if ({{ .FieldName }} != null) {
      final v = {{ .FieldName }}!;
      result[{{ .HeaderName | dartJsonKey }}] = {{ .SerializeStr }};
    }
{{- end }}
    return result;
  }

  factory {{ .className }}.fromHttpHeaders(Map<String, String> headers) {
    final instance = {{ .className }}();
{{- range .headers }}
    {
      final v = headers[{{ .HeaderName | dartJsonKey }}];
      if (v != null) {
        instance.{{ .FieldName }} = {{ .ParseExpr }};
      }
    }
{{- end }}
    return instance;
  }
}
`))

// DartHeaderClass renders a class representing a set of typed HTTP headers,
// with toHttpHeaders()/fromHttpHeaders() to convert to/from the raw
// string-keyed map used by an actual request/response.
func DartHeaderClass(className string, columns []core.EmiHeader, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	headers := renderDartHeaders(columns)

	var buf bytes.Buffer
	if err := dartHeaderTmpl.Execute(&buf, core.H{"className": className, "headers": headers}); err != nil {
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
