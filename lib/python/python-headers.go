package python

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type pyRenderedHeader struct {
	FieldName    string // python attribute name (safe identifier)
	HeaderName   string // actual wire header name, e.g. "X-Request-Id"
	Type         string // python type hint (without Optional wrapper)
	FieldLine    string
	ParseExpr    string // expression to parse os. header string `v` into typed value
	SerializeStr string // expression to serialize the typed value back to string
}

func pythonHeaderType(headerType string) string {
	switch headerType {
	case "int64", "int32", "int":
		return "int"
	case "float64", "float32":
		return "float"
	case "bool":
		return "bool"
	case "string", "":
		return "str"
	default:
		return "str"
	}
}

func renderPyHeaders(columns []core.EmiHeader) []pyRenderedHeader {
	out := []pyRenderedHeader{}
	for _, h := range columns {
		fieldName := headerNameNormalize(h.Name)
		typeHint := pythonHeaderType(h.Type)

		if h.Complex != "" {
			typeHint = h.Complex
		}

		parseExpr := "v"
		serializeStr := "v"
		switch typeHint {
		case "int":
			parseExpr = "int(v)"
			serializeStr = "str(v)"
		case "float":
			parseExpr = "float(v)"
			serializeStr = "str(v)"
		case "bool":
			parseExpr = `v.strip().lower() in ("1", "true", "yes")`
			serializeStr = `"true" if v else "false"`
		}

		out = append(out, pyRenderedHeader{
			FieldName:    fieldName,
			HeaderName:   h.Name,
			Type:         typeHint,
			FieldLine:    fmt.Sprintf("%v: Optional[%v] = None", fieldName, typeHint),
			ParseExpr:    parseExpr,
			SerializeStr: serializeStr,
		})
	}
	return out
}

// headerNameNormalize turns an arbitrary header name (e.g. "X-Request-Id")
// into a valid, camelCase-preserving python identifier ("xRequestId"),
// re-using the exact same normalization used across sibling generators so
// header field names line up between languages.
func headerNameNormalize(s string) string {
	return core.NormaliseKey(s)
}

// PythonHeaderClass renders a dataclass representing a set of typed HTTP
// headers, with `to_http_headers` / `from_http_headers` helpers to convert
// to/from the raw string-keyed mapping used by an actual request/response.
func PythonHeaderClass(className string, columns []core.EmiHeader, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "dataclasses", Objects: []string{"dataclass"}},
			{Location: "typing", Objects: []string{"Optional", "Dict", "Mapping"}},
		},
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: className},
		},
	}

	headers := renderPyHeaders(columns)

	const tmpl = `
@dataclass
class {{ .className }}:
{{ if not .headers }}    pass
{{ end -}}
{{- range .headers }}
    {{ .FieldLine }}
{{ end }}
    def to_http_headers(self) -> Dict[str, str]:
        result: Dict[str, str] = {}
{{- range .headers }}
        if self.{{ .FieldName }} is not None:
            v = self.{{ .FieldName }}
            result["{{ .HeaderName }}"] = {{ .SerializeStr }}
{{- end }}
        return result

    @classmethod
    def from_http_headers(cls, headers: Mapping[str, str]) -> "{{ .className }}":
        instance = cls()
{{- range .headers }}
        v = headers.get("{{ .HeaderName }}")
        if v is not None:
            instance.{{ .FieldName }} = {{ .ParseExpr }}
{{- end }}
        return instance
`

	t := template.Must(template.New("pyheaders").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{"className": className, "headers": headers}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToSnakeCase(className)
	res.SuggestedExtension = ".py"

	return res, nil
}
