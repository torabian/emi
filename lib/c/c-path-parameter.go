package c

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type cPathParam struct {
	WireName string // the raw `:name` placeholder in the url template
	Field    string
	Type     string // "char*" or a numeric C type
	IsString bool
	Format   string // printf-style format used to render a numeric value into a replacement string
}

var cPathParamTmpl = template.Must(template.New("cpathparams").Parse(`
typedef struct {{ .TypeName }} {
{{- range .Params }}
    {{ .Type }} {{ .Field }};
{{- end }}
} {{ .TypeName }};

static inline void {{ .TypeName }}_init({{ .TypeName }}* self) {
    if (!self) return;
{{- range .Params }}
{{- if .IsString }}
    self->{{ .Field }} = NULL;
{{- else }}
    self->{{ .Field }} = ({{ .Type }}) 0;
{{- end }}
{{- end }}
}

static inline void {{ .TypeName }}_free_fields({{ .TypeName }}* self) {
    if (!self) return;
{{- range .Params }}
{{- if .IsString }}
    free(self->{{ .Field }});
    self->{{ .Field }} = NULL;
{{- end }}
{{- end }}
}

/* Returns a newly allocated url - the caller must free() it. */
static inline char* {{ .TypeName }}_apply(const {{ .TypeName }}* self, const char* template_url) {
    char* url = strdup(template_url);
{{- range .Params }}
    {
{{- if .IsString }}
        char* next = fetchx_url_replace(url, ":{{ .WireName }}", self->{{ .Field }} ? self->{{ .Field }} : "");
{{- else }}
        char value_buf[64];
        snprintf(value_buf, sizeof(value_buf), {{ .Format }}, self->{{ .Field }});
        char* next = fetchx_url_replace(url, ":{{ .WireName }}", value_buf);
{{- end }}
        free(url);
        url = next;
    }
{{- end }}
    return url;
}
`))

// CActionPathParams extracts `/:id`-style placeholders out of an action's
// url and renders a small struct to carry them, plus an `_apply` function
// that substitutes them back into the template url (returning a newly
// allocated string). Returns (nil, nil) when the url has no placeholders at
// all.
func CActionPathParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {
	placeholders := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders) == 0 {
		return nil, nil
	}

	params := make([]cPathParam, 0, len(placeholders))
	for _, p := range placeholders {
		t := extractPrimitive(p.Type)
		isString := t == ""
		format := `"%d"`
		switch t {
		case "int64_t":
			format = `"%lld"`
		case "float", "double":
			format = `"%f"`
		case "bool":
			format = `"%d"`
		}
		if isString {
			t = "char*"
		}
		params = append(params, cPathParam{
			WireName: p.Original,
			Field:    core.ToLower(core.NormaliseKey(p.Original)),
			Type:     t,
			IsString: isString,
			Format:   format,
		})
	}

	className := core.ToUpper(core.NormaliseKey(action.GetName()))
	typeName := fmt.Sprintf("%vPathParameters", className)

	var buf bytes.Buffer
	if err := cPathParamTmpl.Execute(&buf, core.H{"TypeName": typeName, "Params": params}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  core.ToSnakeCase(typeName),
		SuggestedExtension: ".h",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: typeName},
		},
	}, nil
}
