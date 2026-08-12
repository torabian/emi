package php

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type phpRenderedHeader struct {
	FieldName    string
	HeaderName   string
	Type         string // PHP type, without `?`
	ParseExpr    string // expression parsing the raw string `$v` into the typed value
	SerializeStr string // expression serializing the typed value `$v` back to string
}

func phpHeaderType(headerType string) string {
	switch headerType {
	case "int64", "int32", "int":
		return "int"
	case "float64", "float32":
		return "float"
	case "bool":
		return "bool"
	default:
		return "string"
	}
}

func renderPhpHeaders(columns []core.EmiHeader) []phpRenderedHeader {
	out := make([]phpRenderedHeader, 0, len(columns))
	for _, h := range columns {
		fieldName := core.ToLower(core.NormaliseKey(h.Name))
		typeHint := phpHeaderType(h.Type)
		if h.Complex != "" {
			typeHint = h.Complex
		}

		parseExpr, serializeStr := "$v", "$v"
		switch typeHint {
		case "int":
			parseExpr, serializeStr = "(int) $v", "(string) $v"
		case "float":
			parseExpr, serializeStr = "(float) $v", "(string) $v"
		case "bool":
			parseExpr, serializeStr = `in_array(strtolower(trim($v)), ["true", "1"], true)`, `$v ? "true" : "false"`
		}

		out = append(out, phpRenderedHeader{
			FieldName:    fieldName,
			HeaderName:   h.Name,
			Type:         typeHint,
			ParseExpr:    parseExpr,
			SerializeStr: serializeStr,
		})
	}
	return out
}

var phpHeaderTmplFuncs = template.FuncMap{"phpString": escapeDoubleQuoted}

var phpHeaderTmpl = template.Must(template.New("phpheaders").Funcs(phpHeaderTmplFuncs).Parse(`
class {{ .className }}
{
{{- range .headers }}
    public ?{{ .Type }} ${{ .FieldName }} = null;
{{- end }}

    public function toHttpHeaders(): array
    {
        $result = [];
{{- range .headers }}
        if ($this->{{ .FieldName }} !== null) {
            $v = $this->{{ .FieldName }};
            $result[{{ .HeaderName | phpString }}] = {{ .SerializeStr }};
        }
{{- end }}
        return $result;
    }

    public static function fromHttpHeaders(array $headers): self
    {
        $instance = new self();
{{- range .headers }}
        if (array_key_exists({{ .HeaderName | phpString }}, $headers)) {
            $v = $headers[{{ .HeaderName | phpString }}];
            $instance->{{ .FieldName }} = {{ .ParseExpr }};
        }
{{- end }}
        return $instance;
    }
}
`))

// PhpHeaderClass renders a class representing a set of typed HTTP headers,
// with toHttpHeaders()/fromHttpHeaders() to convert to/from the raw
// string-keyed array used by an actual request/response.
func PhpHeaderClass(className string, columns []core.EmiHeader, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	headers := renderPhpHeaders(columns)

	var buf bytes.Buffer
	if err := phpHeaderTmpl.Execute(&buf, core.H{"className": className, "headers": headers}); err != nil {
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
