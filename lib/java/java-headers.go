package java

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type javaRenderedHeader struct {
	FieldName    string
	HeaderName   string
	Type         string // boxed Java type
	ParseExpr    string // expression parsing the raw string `v` into the typed value
	SerializeStr string // expression serializing the typed value `v` back to string
}

func javaHeaderType(headerType string) string {
	switch headerType {
	case "int64":
		return "Long"
	case "int32", "int":
		return "Integer"
	case "float64", "float32":
		return "Double"
	case "bool":
		return "Boolean"
	default:
		return "String"
	}
}

func renderJavaHeaders(columns []core.EmiHeader) []javaRenderedHeader {
	out := make([]javaRenderedHeader, 0, len(columns))
	for _, h := range columns {
		fieldName := core.ToLower(core.NormaliseKey(h.Name))
		typeHint := javaHeaderType(h.Type)
		if h.Complex != "" {
			typeHint = h.Complex
		}

		parseExpr, serializeStr := "v", "v"
		switch typeHint {
		case "Integer":
			parseExpr, serializeStr = "Integer.parseInt(v)", "String.valueOf(v)"
		case "Long":
			parseExpr, serializeStr = "Long.parseLong(v)", "String.valueOf(v)"
		case "Double":
			parseExpr, serializeStr = "Double.parseDouble(v)", "String.valueOf(v)"
		case "Boolean":
			parseExpr, serializeStr = `v.trim().equalsIgnoreCase("true") || v.trim().equals("1")`, `v ? "true" : "false"`
		}

		out = append(out, javaRenderedHeader{
			FieldName:    fieldName,
			HeaderName:   h.Name,
			Type:         typeHint,
			ParseExpr:    parseExpr,
			SerializeStr: serializeStr,
		})
	}
	return out
}

var javaHeaderTmplFuncs = template.FuncMap{"javaString": escapeDoubleQuoted}

var javaHeaderTmpl = template.Must(template.New("javaheaders").Funcs(javaHeaderTmplFuncs).Parse(`
public class {{ .className }} {
{{- range .headers }}
    public {{ .Type }} {{ .FieldName }};
{{- end }}

    public java.util.Map<String, String> toHttpHeaders() {
        var result = new java.util.HashMap<String, String>();
{{- range .headers }}
        if ({{ .FieldName }} != null) {
            var v = {{ .FieldName }};
            result.put({{ .HeaderName | javaString }}, {{ .SerializeStr }});
        }
{{- end }}
        return result;
    }

    public static {{ .className }} fromHttpHeaders(java.util.Map<String, String> headers) {
        var instance = new {{ .className }}();
{{- range .headers }}
        {
            var v = headers.get({{ .HeaderName | javaString }});
            if (v != null) {
                instance.{{ .FieldName }} = {{ .ParseExpr }};
            }
        }
{{- end }}
        return instance;
    }
}
`))

// JavaHeaderClass renders a (package-private) class representing a set of
// typed HTTP headers, with toHttpHeaders()/fromHttpHeaders() to convert
// to/from the raw string-keyed map used by an actual request/response.
func JavaHeaderClass(className string, columns []core.EmiHeader, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	headers := renderJavaHeaders(columns)

	var buf bytes.Buffer
	if err := javaHeaderTmpl.Execute(&buf, core.H{"className": className, "headers": headers}); err != nil {
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
