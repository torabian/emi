package cpp

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type cppEnumCase struct {
	Identifier  string
	Value       string
	Description string
}

type cppRenderedEnum struct {
	Name  string
	Cases []cppEnumCase
}

// ---- generic dialect ---------------------------------------------------------
//
// A plain `enum class Name : int32_t`, plus free `<Name>ToString`/
// `<Name>FromString` functions doing the wire-value mapping - C++ scoped enums
// have no reflection of their own, so (unlike Unreal's UENUM, or C#'s
// [JsonConverter]) this is hand-generated the same way the C target does it.

var cppGenericEnumTmplFuncs = template.FuncMap{"cppString": escapeDoubleQuoted}

var cppGenericEnumTmpl = template.Must(template.New("cppgenericenum").Funcs(cppGenericEnumTmplFuncs).Parse(`
enum class {{ .Name }} : int32_t {
{{- range $i, $c := .Cases }}
    {{ $c.Identifier }} = {{ $i }},{{ if $c.Description }} // {{ $c.Description }}{{ end }}
{{- end }}
};

inline const char* {{ .Name }}ToString({{ .Name }} value) {
    switch (value) {
{{- range .Cases }}
        case {{ $.Name }}::{{ .Identifier }}: return {{ .Value | cppString }};
{{- end }}
        default: return "";
    }
}

inline {{ .Name }} {{ .Name }}FromString(const std::string& value) {
{{- range .Cases }}
    if (value == {{ .Value | cppString }}) return {{ $.Name }}::{{ .Identifier }};
{{- end }}
    return static_cast<{{ .Name }}>(0);
}
`))

func renderCppGenericEnumDecl(e *cppRenderedEnum) string {
	var buf bytes.Buffer
	if err := cppGenericEnumTmpl.Execute(&buf, e); err != nil {
		return ""
	}
	return buf.String()
}

// ---- unreal dialect ------------------------------------------------------------
//
// A `UENUM(BlueprintType)` - visible to Blueprint and the editor for free, and
// (de)serialized through the exact same UPROPERTY reflection every generated
// USTRUCT field already relies on (see cpp-class-generator.go), so no hand-written
// to/from-string glue is needed at all here, unlike the generic dialect above.

var cppUnrealEnumTmpl = template.Must(template.New("cppunrealenum").Parse(`
UENUM(BlueprintType)
enum class {{ .Name }} : uint8 {
{{- range .Cases }}
    {{ .Identifier }} UMETA(DisplayName = "{{ .Value }}"),{{ if .Description }} // {{ .Description }}{{ end }}
{{- end }}
};
`))

func renderCppUnrealEnumDecl(e *cppRenderedEnum) string {
	var buf bytes.Buffer
	if err := cppUnrealEnumTmpl.Execute(&buf, e); err != nil {
		return ""
	}
	return buf.String()
}

func renderCppEnumDecl(dialect Dialect, e *cppRenderedEnum) string {
	if dialect == DialectUnreal {
		return renderCppUnrealEnumDecl(e)
	}
	return renderCppGenericEnumDecl(e)
}

// cppRenderEnumFromInline builds the companion enum for an inline `of:` enum
// field (no `target:`), named after the flattened prefix (see cppEnumBaseType).
func cppRenderEnumFromInline(name string, ofType []*core.EmiEnumInline) *cppRenderedEnum {
	cases := make([]cppEnumCase, 0, len(ofType))
	for _, item := range ofType {
		if item == nil {
			continue
		}
		cases = append(cases, cppEnumCase{
			Identifier:  core.NormaliseKey(item.Key),
			Value:       item.Key,
			Description: item.Description,
		})
	}
	if len(cases) == 0 {
		// A scoped enum with zero enumerators is legal C++, but degenerate and
		// useless - give it one placeholder member the same way the C target's
		// (mandatory, since a plain C enum forbids the empty case entirely) does.
		cases = append(cases, cppEnumCase{Identifier: "Unspecified"})
	}
	return &cppRenderedEnum{Name: name, Cases: cases}
}

// CppStandaloneEnum renders a module-level `enums:` entry.
func CppStandaloneEnum(enum core.EmiEnum, dialect Dialect, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	rendered := &cppRenderedEnum{Name: enum.GetName()}
	for _, item := range enum.Fields {
		rendered.Cases = append(rendered.Cases, cppEnumCase{
			Identifier:  core.NormaliseKey(item.Key),
			Value:       item.Key,
			Description: item.Description,
		})
	}
	if len(rendered.Cases) == 0 {
		rendered.Cases = append(rendered.Cases, cppEnumCase{Identifier: "Unspecified"})
	}

	return &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: enum.GetName()},
			{Name: TOKEN_ROOT_CLASS, Value: enum.GetName()},
		},
		ActualScript:       []byte(renderCppEnumDecl(dialect, rendered)),
		SuggestedFileName:  enum.GetName(),
		SuggestedExtension: ".h",
	}, nil
}
