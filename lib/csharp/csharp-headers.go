package csharp

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type csRenderedHeader struct {
	FieldName    string
	HeaderName   string
	Type         string // C# type, without `?`
	ParseExpr    string // expression parsing the raw string `v` into the typed value
	SerializeStr string // expression serializing the typed value `v` back to string
}

func csHeaderType(headerType string) string {
	switch headerType {
	case "int64", "int32", "int":
		return "int"
	case "float64", "float32":
		return "double"
	case "bool":
		return "bool"
	default:
		return "string"
	}
}

func renderCsHeaders(columns []core.EmiHeader) []csRenderedHeader {
	out := make([]csRenderedHeader, 0, len(columns))
	for _, h := range columns {
		fieldName := core.NormaliseKey(h.Name)
		typeHint := csHeaderType(h.Type)
		if h.Complex != "" {
			typeHint = h.Complex
		}

		parseExpr, serializeStr := "v", "v"
		switch typeHint {
		case "int":
			parseExpr, serializeStr = "int.Parse(v)", "v.ToString()"
		case "double":
			parseExpr, serializeStr = "double.Parse(v)", "v.ToString()"
		case "bool":
			parseExpr, serializeStr = `v.Trim().ToLowerInvariant() == "true" || v.Trim() == "1"`, `v ? "true" : "false"`
		}

		out = append(out, csRenderedHeader{
			FieldName:    fieldName,
			HeaderName:   h.Name,
			Type:         typeHint,
			ParseExpr:    parseExpr,
			SerializeStr: serializeStr,
		})
	}
	return out
}

var csHeaderTmplFuncs = template.FuncMap{"csString": escapeDoubleQuoted, "lowerFirst": core.ToLower}

var csHeaderTmpl = template.Must(template.New("csheaders").Funcs(csHeaderTmplFuncs).Parse(`
public class {{ .className }}
{
{{- range .headers }}
    public {{ .Type }}? {{ .FieldName }} { get; set; }
{{- end }}

    public Dictionary<string, string> ToHttpHeaders()
    {
        var result = new Dictionary<string, string>();
{{- range .headers }}
        if ({{ .FieldName }} != null)
        {
            var v = {{ .FieldName }}!.Value;
            result[{{ .HeaderName | csString }}] = {{ .SerializeStr }};
        }
{{- end }}
        return result;
    }

    public static {{ .className }} FromHttpHeaders(IDictionary<string, string> headers)
    {
        var instance = new {{ .className }}();
{{- range .headers }}
        if (headers.TryGetValue({{ .HeaderName | csString }}, out var {{ .FieldName | lowerFirst }}Raw))
        {
            var v = {{ .FieldName | lowerFirst }}Raw;
            instance.{{ .FieldName }} = {{ .ParseExpr }};
        }
{{- end }}
        return instance;
    }
}
`))

// CSharpHeaderClass renders a class representing a set of typed HTTP
// headers, with ToHttpHeaders()/FromHttpHeaders() to convert to/from the raw
// string-keyed dictionary used by an actual request/response.
func CSharpHeaderClass(className string, columns []core.EmiHeader, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	headers := renderCsHeaders(columns)

	var buf bytes.Buffer
	if err := csHeaderTmpl.Execute(&buf, core.H{"className": className, "headers": headers}); err != nil {
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
