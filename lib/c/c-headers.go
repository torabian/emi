package c

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type cRenderedHeader struct {
	FieldName    string
	HeaderName   string
	Type         string // C scalar type, or "char*" for string
	IsString     bool
	ParseExpr    string // expression parsing the raw string `v` into the typed value (only used for string: assigns directly)
	SerializeFmt string // printf-style format + cast used to render the value into a header line
}

func cHeaderType(headerType string) string {
	switch headerType {
	case "int64":
		return "int64_t"
	case "int32", "int":
		return "int32_t"
	case "float64", "float32":
		return "double"
	case "bool":
		return "bool"
	default:
		return "char*"
	}
}

func renderCHeaders(columns []core.EmiHeader) []cRenderedHeader {
	out := make([]cRenderedHeader, 0, len(columns))
	for _, h := range columns {
		fieldName := core.ToLower(core.NormaliseKey(h.Name))
		typeHint := cHeaderType(h.Type)
		isString := typeHint == "char*"

		format := `"%d"`
		switch typeHint {
		case "int64_t":
			format = `"%lld"`
		case "double":
			format = `"%f"`
		case "bool":
			format = `"%s"`
		}

		out = append(out, cRenderedHeader{
			FieldName:    fieldName,
			HeaderName:   h.Name,
			Type:         typeHint,
			IsString:     isString,
			SerializeFmt: format,
		})
	}
	return out
}

func trimQuotes(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

var cHeaderTmplFuncs = template.FuncMap{"cString": escapeDoubleQuoted, "trimQuotes": trimQuotes}

var cHeaderTmpl = template.Must(template.New("cheaders").Funcs(cHeaderTmplFuncs).Parse(`
typedef struct {{ .className }} {
{{- range .headers }}
{{- if .IsString }}
    char* {{ .FieldName }}; /* NULL if unset */
{{- else }}
    bool {{ .FieldName }}_set;
    {{ .Type }} {{ .FieldName }};
{{- end }}
{{- end }}
} {{ .className }};

static inline void {{ .className }}_init({{ .className }}* self) {
    if (!self) return;
{{- range .headers }}
{{- if .IsString }}
    self->{{ .FieldName }} = NULL;
{{- else }}
    self->{{ .FieldName }}_set = false;
    self->{{ .FieldName }} = ({{ .Type }}) 0;
{{- end }}
{{- end }}
}

static inline void {{ .className }}_free_fields({{ .className }}* self) {
    if (!self) return;
{{- range .headers }}
{{- if .IsString }}
    free(self->{{ .FieldName }});
    self->{{ .FieldName }} = NULL;
{{- end }}
{{- end }}
}

/* Appends every set field as an HTTP header line onto (and returns) the
 * given curl_slist - pass NULL to start a new one. */
static inline struct curl_slist* {{ .className }}_to_curl_slist(const {{ .className }}* self, struct curl_slist* list) {
    if (!self) return list;
    char buf[256];
{{- range .headers }}
{{- if .IsString }}
    if (self->{{ .FieldName }}) {
        snprintf(buf, sizeof(buf), "{{ .HeaderName }}: %s", self->{{ .FieldName }});
        list = curl_slist_append(list, buf);
    }
{{- else }}
    if (self->{{ .FieldName }}_set) {
        snprintf(buf, sizeof(buf), "{{ .HeaderName }}: {{ .SerializeFmt | trimQuotes }}", self->{{ .FieldName }});
        list = curl_slist_append(list, buf);
    }
{{- end }}
{{- end }}
    return list;
}
`))

// CHeaderStruct renders a struct representing a set of typed HTTP headers,
// with `_to_curl_slist` to serialize the set ones onto a curl header list.
func CHeaderStruct(className string, columns []core.EmiHeader, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	headers := renderCHeaders(columns)

	var buf bytes.Buffer
	if err := cHeaderTmpl.Execute(&buf, core.H{"className": className, "headers": headers}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  core.ToSnakeCase(className),
		SuggestedExtension: ".h",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: className},
		},
	}, nil
}
