package c

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type cQueryField struct {
	Name     string
	Type     string
	IsString bool
	Format   string
}

func cQueryFieldType(query *core.EmiQueryField) (ctype string, isString bool, format string) {
	switch query.Type {
	case core.FieldTypeArray, core.FieldTypeSlice:
		// Arrays in a query string are a niche, non-scalar shape (bracket-
		// or repeat-encoding conventions vary by server) - out of scope for
		// this runtime's simple key=value appender; represented as a raw
		// (already-encoded-by-the-caller) string the same way an
		// unresolved type falls back below.
		return "char*", true, ""
	default:
		t := extractPrimitive(string(query.Type))
		if t == "" {
			return "char*", true, ""
		}
		f := `"%d"`
		switch t {
		case "int64_t":
			f = `"%lld"`
		case "float", "double":
			f = `"%f"`
		case "bool":
			f = `"%d"`
		}
		return t, false, f
	}
}

var cQueryTmpl = template.Must(template.New("cqs").Parse(`
typedef struct {{ .className }} {
{{- range .fields }}
{{- if .IsString }}
    char* {{ .Name }}; /* NULL if unset */
{{- else }}
    bool {{ .Name }}_set;
    {{ .Type }} {{ .Name }};
{{- end }}
{{- end }}
} {{ .className }};

static inline void {{ .className }}_init({{ .className }}* self) {
    if (!self) return;
{{- range .fields }}
{{- if .IsString }}
    self->{{ .Name }} = NULL;
{{- else }}
    self->{{ .Name }}_set = false;
    self->{{ .Name }} = ({{ .Type }}) 0;
{{- end }}
{{- end }}
}

static inline void {{ .className }}_free_fields({{ .className }}* self) {
    if (!self) return;
{{- range .fields }}
{{- if .IsString }}
    free(self->{{ .Name }});
    self->{{ .Name }} = NULL;
{{- end }}
{{- end }}
}

/* Returns a newly allocated url with every set field appended as a query
 * parameter - the caller must free() it (url itself is never modified or
 * freed). */
static inline char* {{ .className }}_apply(const {{ .className }}* self, const char* url) {
    char* out = strdup(url);
    if (!self) return out;
{{- range .fields }}
{{- if .IsString }}
    if (self->{{ .Name }}) {
        char* next = fetchx_url_append_query(out, "{{ .Name }}", self->{{ .Name }});
        free(out);
        out = next;
    }
{{- else }}
    if (self->{{ .Name }}_set) {
        char value_buf[64];
        snprintf(value_buf, sizeof(value_buf), {{ .Format }}, self->{{ .Name }});
        char* next = fetchx_url_append_query(out, "{{ .Name }}", value_buf);
        free(out);
        out = next;
    }
{{- end }}
{{- end }}
    return out;
}
`))

// CActionQueryStruct renders a struct out of an action's typed `qs:`
// definition, plus an `_apply` function that appends every set field onto a
// url as a query string.
func CActionQueryStruct(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	query := action.GetQuery()
	if len(query) == 0 {
		return nil, nil
	}

	fields := make([]cQueryField, 0, len(query))
	for _, q := range query {
		t, isString, format := cQueryFieldType(q)
		fields = append(fields, cQueryField{Name: q.Name, Type: t, IsString: isString, Format: format})
	}

	className := fmt.Sprintf("%vQueryParams", core.ToUpper(core.NormaliseKey(action.GetName())))

	var buf bytes.Buffer
	if err := cQueryTmpl.Execute(&buf, core.H{"className": className, "fields": fields}); err != nil {
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
