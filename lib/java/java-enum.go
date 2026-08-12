package java

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type javaEnumCase struct {
	Identifier  string
	Value       string
	Description string
}

type javaRenderedEnum struct {
	Name     string
	IsPublic bool
	Cases    []javaEnumCase
}

var javaEnumTmplFuncs = template.FuncMap{"javaString": escapeDoubleQuoted}

// Every generated enum carries its own wire-value mapping via @JsonValue
// (serialize) and @JsonCreator (deserialize) - the standard Jackson idiom
// for an enum whose wire representation isn't (or might not be) a valid
// Java identifier, independent of any shared runtime helper.
var javaEnumTmpl = template.Must(template.New("javaenum").Funcs(core.CommonMap).Funcs(javaEnumTmplFuncs).Parse(`
{{ if .IsPublic }}public {{ end }}enum {{ .Name }} {
{{- range $i, $c := .Cases }}
    {{ $c.Identifier }}({{ $c.Value | javaString }}){{ if last $i $.Cases }};{{ else }},{{ end }}{{ if $c.Description }} // {{ $c.Description }}{{ end }}
{{- end }}

    private final String wireValue;

    {{ .Name }}(String wireValue) {
        this.wireValue = wireValue;
    }

    @JsonValue
    public String toWireValue() {
        return wireValue;
    }

    @JsonCreator
    public static {{ .Name }} fromWireValue(String value) {
        for ({{ .Name }} candidate : values()) {
            if (candidate.wireValue.equals(value)) {
                return candidate;
            }
        }
        return values()[0];
    }
}
`))

func renderJavaEnumDecl(e *javaRenderedEnum) string {
	var buf bytes.Buffer
	if err := javaEnumTmpl.Execute(&buf, core.H{"Name": e.Name, "IsPublic": e.IsPublic, "Cases": e.Cases}); err != nil {
		return ""
	}
	return buf.String()
}

// javaRenderEnumFromInline builds the (package-private, since it lives
// alongside another public type in the same file) companion enum for an
// inline `of:` enum field (no `target:`), named after the flattened prefix
// (see javaEnumBaseType).
func javaRenderEnumFromInline(name string, ofType []*core.EmiEnumInline) *javaRenderedEnum {
	cases := make([]javaEnumCase, 0, len(ofType))
	for _, item := range ofType {
		if item == nil {
			continue
		}
		cases = append(cases, javaEnumCase{
			Identifier:  core.ToSnakeUpper(core.NormaliseKey(item.Key)),
			Value:       item.Key,
			Description: item.Description,
		})
	}
	return &javaRenderedEnum{Name: name, Cases: cases}
}

// JavaStandaloneEnum renders a module-level `enums:` entry - the one public
// type in its own file.
func JavaStandaloneEnum(enum core.EmiEnum, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	rendered := &javaRenderedEnum{Name: enum.GetName(), IsPublic: true}
	for _, item := range enum.Fields {
		rendered.Cases = append(rendered.Cases, javaEnumCase{
			Identifier:  core.ToSnakeUpper(core.NormaliseKey(item.Key)),
			Value:       item.Key,
			Description: item.Description,
		})
	}

	return &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: enum.GetName()},
			{Name: TOKEN_ROOT_CLASS, Value: enum.GetName()},
		},
		ActualScript:       []byte(renderJavaEnumDecl(rendered)),
		SuggestedFileName:  enum.GetName(),
		SuggestedExtension: ".java",
	}, nil
}
