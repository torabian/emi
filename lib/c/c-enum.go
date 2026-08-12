package c

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type cEnumCase struct {
	Identifier  string // e.g. WIDGET_DTO_KIND_CAT_KIND - C enum constants share one flat, global namespace, so every case is prefixed with its own enum's name
	Value       string
	Description string
}

type cRenderedEnum struct {
	Name  string
	Cases []cEnumCase
}

func lastCase(i int, cases []cEnumCase) bool {
	return i == len(cases)-1
}

var cEnumTmplFuncs = template.FuncMap{
	"cString": escapeDoubleQuoted,
	"last":    lastCase,
}

var cEnumTmpl = template.Must(template.New("cenum").Funcs(cEnumTmplFuncs).Parse(`
typedef enum {{ .Name }} {
{{- range $i, $c := .Cases }}
    {{ $c.Identifier }} = {{ $i }}{{ if not (last $i $.Cases) }},{{ end }}{{ if $c.Description }} /* {{ $c.Description }} */{{ end }}
{{- end }}
} {{ .Name }};
`))

// Every function is `static inline`, the standard C idiom for a
// header-only library (matching vendored cjson/cJSON.h's own style) - each
// generated .h is meant to be #include-d directly, from as many
// translation units as needed, without a matching .c file or any risk of
// duplicate-symbol link errors.
var cEnumImplTmpl = template.Must(template.New("cenumimpl").Funcs(cEnumTmplFuncs).Parse(`
static inline const char* {{ .Name }}_to_string({{ .Name }} value) {
    switch (value) {
{{- range .Cases }}
        case {{ .Identifier }}: return {{ .Value | cString }};
{{- end }}
        default: return "";
    }
}

static inline {{ .Name }} {{ .Name }}_from_string(const char* value) {
    if (value != NULL) {
{{- range .Cases }}
        if (strcmp(value, {{ .Value | cString }}) == 0) { return {{ .Identifier }}; }
{{- end }}
    }
    return ({{ .Name }}) 0;
}
`))

func renderCEnumDecl(e *cRenderedEnum) string {
	var buf bytes.Buffer
	if err := cEnumTmpl.Execute(&buf, e); err != nil {
		return ""
	}
	return buf.String()
}

func renderCEnumImpl(e *cRenderedEnum) string {
	var buf bytes.Buffer
	if err := cEnumImplTmpl.Execute(&buf, e); err != nil {
		return ""
	}
	return buf.String()
}

// cEnumBaseType resolves an `enum`/`enum?` field: a `target:` reference to a
// module-level `enums:` entry, or inline `of:` values (a flattened,
// generated companion enum named `<prefixName><FieldName>`). Falls back to
// a synthetic marker meaning "no real enum type" is handled by the caller
// (there's no plain-string fallback for C the way there is for the other
// targets, since a C enum always needs an actual declared type - an
// enum-typed field with neither `target:` nor `of:` is empty/degenerate and
// still gets its own, empty enum).
func cEnumBaseType(field *core.EmiField, prefixName string) string {
	if field.Target != "" {
		return field.Target
	}
	return prefixName + core.ToUpper(field.Name)
}

// cRenderEnumFromInline builds the companion enum for an inline `of:` enum
// field (no `target:`), named after the flattened prefix (see
// cEnumBaseType). Also used (with zero cases) for the "neither target nor
// of:" degenerate case.
func cRenderEnumFromInline(name string, ofType []*core.EmiEnumInline) *cRenderedEnum {
	prefix := core.ToSnakeUpper(name)
	cases := make([]cEnumCase, 0, len(ofType))
	for _, item := range ofType {
		if item == nil {
			continue
		}
		cases = append(cases, cEnumCase{
			Identifier:  prefix + "_" + core.ToSnakeUpper(core.NormaliseKey(item.Key)),
			Value:       item.Key,
			Description: item.Description,
		})
	}
	if len(cases) == 0 {
		// Unlike every other target, a C enum can never be declared with
		// zero enumerators (`enum {}` is a syntax error) - a field with
		// neither `target:` nor `of:` values still needs one to exist.
		cases = append(cases, cEnumCase{
			Identifier: prefix + "_UNSPECIFIED",
			Value:      "",
		})
	}
	return &cRenderedEnum{Name: name, Cases: cases}
}

// CStandaloneEnum renders a module-level `enums:` entry.
func CStandaloneEnum(enum core.EmiEnum, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	rendered := cRenderEnumFromInline(enum.GetName(), func() []*core.EmiEnumInline {
		out := make([]*core.EmiEnumInline, len(enum.Fields))
		for i := range enum.Fields {
			out[i] = &enum.Fields[i]
		}
		return out
	}())

	script := renderCEnumDecl(rendered) + "\n" + renderCEnumImpl(rendered)

	return &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: enum.GetName()},
			{Name: TOKEN_ROOT_CLASS, Value: enum.GetName()},
		},
		ActualScript: []byte(script),
		// Must match HeaderIncludeDependency's naming convention (every
		// #include of a target/dto reference is spelled
		// `ToSnakeCase(name)+".h"`), or a field of this enum's type in
		// another struct would #include a file that doesn't exist.
		SuggestedFileName:  core.ToSnakeCase(enum.GetName()),
		SuggestedExtension: ".h",
	}, nil
}
