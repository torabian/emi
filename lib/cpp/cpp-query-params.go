package cpp

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type cppQueryField struct {
	Name     string
	WireName string
	CppType  string
	IsArray  string // "" (scalar), or the element type when this is an array/slice of a primitive
}

func cppQueryFieldType(dialect Dialect, query *core.EmiQueryField) (cppType string, elemType string) {
	switch query.Type {
	case core.FieldTypeArray, core.FieldTypeSlice:
		primitive := extractPrimitive(dialect, query.Primitive)
		if primitive == "" {
			if dialect == DialectUnreal {
				primitive = "FString"
			} else {
				primitive = "std::string"
			}
		}
		return "", primitive
	default:
		t := extractPrimitive(dialect, string(query.Type))
		if t == "" {
			if dialect == DialectUnreal {
				t = "FString"
			} else {
				t = "std::string"
			}
		}
		return t, ""
	}
}

// ---- generic dialect ---------------------------------------------------------

var cppGenericQueryTmplFuncs = template.FuncMap{"cppString": escapeDoubleQuoted}

var cppGenericQueryTmpl = template.Must(template.New("cppgenericqs").Funcs(cppGenericQueryTmplFuncs).Parse(`
class {{ .className }} {
public:
{{- range .fields }}
{{- if .IsArray }}
    std::optional<std::vector<{{ .IsArray }}>> {{ .Name }};
{{- else }}
    std::optional<{{ .CppType }}> {{ .Name }};
{{- end }}
{{- end }}

    // Appends every set field onto the given url as a query parameter.
    std::string Apply(const std::string& url) const {
        std::string out = url;
{{- range .fields }}
{{- if .IsArray }}
        if ({{ .Name }}.has_value()) {
            for (const auto& item : *{{ .Name }}) {
{{- if eq .IsArray "std::string" }}
                out = emi::EmiUrlAppendQuery(out, {{ .WireName | cppString }}, item);
{{- else }}
                out = emi::EmiUrlAppendQuery(out, {{ .WireName | cppString }}, std::to_string(item));
{{- end }}
            }
        }
{{- else if eq .CppType "std::string" }}
        if ({{ .Name }}.has_value()) { out = emi::EmiUrlAppendQuery(out, {{ .WireName | cppString }}, *{{ .Name }}); }
{{- else if eq .CppType "bool" }}
        if ({{ .Name }}.has_value()) { out = emi::EmiUrlAppendQuery(out, {{ .WireName | cppString }}, *{{ .Name }} ? "true" : "false"); }
{{- else }}
        if ({{ .Name }}.has_value()) { out = emi::EmiUrlAppendQuery(out, {{ .WireName | cppString }}, std::to_string(*{{ .Name }})); }
{{- end }}
{{- end }}
        return out;
    }
};
`))

// CppGenericActionQueryClass renders a class out of an action's typed `qs:`
// definition, plus an Apply(url) method that appends every set field onto a url
// as a query string.
func CppGenericActionQueryClass(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	query := action.GetQuery()
	if len(query) == 0 {
		return nil, nil
	}

	fields := make([]cppQueryField, 0, len(query))
	for _, q := range query {
		cppType, elemType := cppQueryFieldType(DialectGeneric, q)
		fields = append(fields, cppQueryField{Name: q.Name, WireName: q.Name, CppType: cppType, IsArray: elemType})
	}

	className := fmt.Sprintf("%vQueryParams", core.ToUpper(core.NormaliseKey(action.GetName())))

	var buf bytes.Buffer
	if err := cppGenericQueryTmpl.Execute(&buf, core.H{"className": className, "fields": fields}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  className,
		SuggestedExtension: ".hpp",
		Tokens:             []core.GeneratedScriptToken{{Name: TOKEN_ROOT_CLASS, Value: className}},
	}, nil
}

// ---- unreal dialect ------------------------------------------------------------

var cppUnrealQueryTmpl = template.Must(template.New("cppunrealqs").Parse(`
class {{ .className }} {
public:
{{- range .fields }}
{{- if .IsArray }}
    TOptional<TArray<{{ .IsArray }}>> {{ .Name }};
{{- else }}
    TOptional<{{ .CppType }}> {{ .Name }};
{{- end }}
{{- end }}

    FString Apply(const FString& Url) const {
        FString Out = Url;
{{- range .fields }}
{{- if .IsArray }}
        if ({{ .Name }}.IsSet()) {
            for (const auto& Item : {{ .Name }}.GetValue()) {
{{- if eq .IsArray "FString" }}
                Out = Out + (Out.Contains(TEXT("?")) ? TEXT("&") : TEXT("?")) + TEXT("{{ .WireName }}=") + Item;
{{- else }}
                Out = Out + (Out.Contains(TEXT("?")) ? TEXT("&") : TEXT("?")) + TEXT("{{ .WireName }}=") + FString::SanitizeFloat((double) Item);
{{- end }}
            }
        }
{{- else if eq .CppType "FString" }}
        if ({{ .Name }}.IsSet()) { Out = Out + (Out.Contains(TEXT("?")) ? TEXT("&") : TEXT("?")) + TEXT("{{ .WireName }}=") + {{ .Name }}.GetValue(); }
{{- else if eq .CppType "bool" }}
        if ({{ .Name }}.IsSet()) { Out = Out + (Out.Contains(TEXT("?")) ? TEXT("&") : TEXT("?")) + TEXT("{{ .WireName }}=") + ({{ .Name }}.GetValue() ? TEXT("true") : TEXT("false")); }
{{- else }}
        if ({{ .Name }}.IsSet()) { Out = Out + (Out.Contains(TEXT("?")) ? TEXT("&") : TEXT("?")) + TEXT("{{ .WireName }}=") + FString::SanitizeFloat((double) {{ .Name }}.GetValue()); }
{{- end }}
{{- end }}
        return Out;
    }
};
`))

// CppUnrealActionQueryClass mirrors CppGenericActionQueryClass for the unreal
// dialect.
func CppUnrealActionQueryClass(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	query := action.GetQuery()
	if len(query) == 0 {
		return nil, nil
	}

	fields := make([]cppQueryField, 0, len(query))
	for _, q := range query {
		cppType, elemType := cppQueryFieldType(DialectUnreal, q)
		fields = append(fields, cppQueryField{Name: core.ToUpper(q.Name), WireName: q.Name, CppType: cppType, IsArray: elemType})
	}

	className := "F" + fmt.Sprintf("%vQueryParams", core.ToUpper(core.NormaliseKey(action.GetName())))

	var buf bytes.Buffer
	if err := cppUnrealQueryTmpl.Execute(&buf, core.H{"className": className, "fields": fields}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  className,
		SuggestedExtension: ".h",
		Tokens:             []core.GeneratedScriptToken{{Name: TOKEN_ROOT_CLASS, Value: className}},
	}, nil
}

// CppActionQueryClass dispatches to the right dialect's query class renderer.
func CppActionQueryClass(action core.EmiRpcAction, dialect Dialect, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	if dialect == DialectUnreal {
		return CppUnrealActionQueryClass(action, ctx)
	}
	return CppGenericActionQueryClass(action, ctx)
}
