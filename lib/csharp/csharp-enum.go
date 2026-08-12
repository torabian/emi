package csharp

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type csEnumCase struct {
	Identifier  string
	Value       string
	Description string
}

type csRenderedEnum struct {
	Name  string
	Cases []csEnumCase
}

var csEnumTmplFuncs = template.FuncMap{"csString": escapeDoubleQuoted}

// Each generated enum carries its own tiny JsonConverter, mapping the C#
// member back and forth to the exact raw wire string (which might not even
// be a valid C# identifier) - independent of any reflection/attribute
// mechanism, so it never depends on anything beyond System.Text.Json itself.
var csEnumTmpl = template.Must(template.New("csenum").Funcs(csEnumTmplFuncs).Parse(`
[JsonConverter(typeof({{ .Name }}Converter))]
public enum {{ .Name }}
{
{{- range .Cases }}
    {{ .Identifier }},{{ if .Description }} // {{ .Description }}{{ end }}
{{- end }}
}

public class {{ .Name }}Converter : JsonConverter<{{ .Name }}>
{
    private static readonly Dictionary<string, {{ .Name }}> FromWire = new()
    {
{{- range .Cases }}
        [{{ .Value | csString }}] = {{ $.Name }}.{{ .Identifier }},
{{- end }}
    };

    private static readonly Dictionary<{{ .Name }}, string> ToWire = new()
    {
{{- range .Cases }}
        [{{ $.Name }}.{{ .Identifier }}] = {{ .Value | csString }},
{{- end }}
    };

    public override {{ .Name }} Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
    {
        var raw = reader.GetString();
        return raw != null && FromWire.TryGetValue(raw, out var value) ? value : default;
    }

    public override void Write(Utf8JsonWriter writer, {{ .Name }} value, JsonSerializerOptions options)
    {
        writer.WriteStringValue(ToWire.TryGetValue(value, out var raw) ? raw : value.ToString());
    }
}
`))

func renderCsEnumDecl(e *csRenderedEnum) string {
	var buf bytes.Buffer
	if err := csEnumTmpl.Execute(&buf, e); err != nil {
		return ""
	}
	return buf.String()
}

// csRenderEnumFromInline builds the companion enum for an inline `of:` enum
// field (no `target:`), named after the flattened prefix (see
// csEnumBaseType).
func csRenderEnumFromInline(name string, ofType []*core.EmiEnumInline) *csRenderedEnum {
	cases := make([]csEnumCase, 0, len(ofType))
	for _, item := range ofType {
		if item == nil {
			continue
		}
		cases = append(cases, csEnumCase{
			Identifier:  core.NormaliseKey(item.Key),
			Value:       item.Key,
			Description: item.Description,
		})
	}
	return &csRenderedEnum{Name: name, Cases: cases}
}

// CSharpStandaloneEnum renders a module-level `enums:` entry.
func CSharpStandaloneEnum(enum core.EmiEnum, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	rendered := &csRenderedEnum{Name: enum.GetName()}
	for _, item := range enum.Fields {
		rendered.Cases = append(rendered.Cases, csEnumCase{
			Identifier:  core.NormaliseKey(item.Key),
			Value:       item.Key,
			Description: item.Description,
		})
	}

	return &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: enum.GetName()},
			{Name: TOKEN_ROOT_CLASS, Value: enum.GetName()},
		},
		ActualScript:       []byte(renderCsEnumDecl(rendered)),
		SuggestedFileName:  enum.GetName(),
		SuggestedExtension: ".cs",
	}, nil
}
