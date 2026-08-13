package cpp

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type cppRenderedHeaderField struct {
	FieldName  string // generic: camelCase member name; unreal: PascalCase member name
	HeaderName string // the raw wire header name
	CppType    string // scalar type, without any nullable wrapper
}

func cppHeaderType(dialect Dialect, headerType string) string {
	primitive := extractPrimitive(dialect, headerType)
	if primitive == "" {
		if dialect == DialectUnreal {
			return "FString"
		}
		return "std::string"
	}
	return primitive
}

// ---- generic dialect ---------------------------------------------------------

var cppGenericHeaderTmpl = template.Must(template.New("cppgenericheaders").Funcs(template.FuncMap{"cppString": escapeDoubleQuoted}).Parse(`
class {{ .className }} {
public:
{{- range .headers }}
    std::optional<{{ .CppType }}> {{ .FieldName }};
{{- end }}

    std::vector<emi::EmiHttpHeader> ToHeaders() const {
        std::vector<emi::EmiHttpHeader> result;
{{- range .headers }}
        if ({{ .FieldName }}.has_value()) {
{{- if eq .CppType "std::string" }}
            result.push_back({ {{ .HeaderName | cppString }}, *{{ .FieldName }} });
{{- else if eq .CppType "bool" }}
            result.push_back({ {{ .HeaderName | cppString }}, *{{ .FieldName }} ? "true" : "false" });
{{- else }}
            result.push_back({ {{ .HeaderName | cppString }}, std::to_string(*{{ .FieldName }}) });
{{- end }}
        }
{{- end }}
        return result;
    }

    static {{ .className }} FromHeaders(const std::vector<emi::EmiHttpHeader>& headers) {
        {{ .className }} result;
        for (const auto& h : headers) {
{{- range .headers }}
            if (h.name == {{ .HeaderName | cppString }}) {
{{- if eq .CppType "std::string" }}
                result.{{ .FieldName }} = h.value;
{{- else if eq .CppType "bool" }}
                result.{{ .FieldName }} = (h.value == "true" || h.value == "1");
{{- else if eq .CppType "int32_t" }}
                result.{{ .FieldName }} = (int32_t) std::stol(h.value);
{{- else if eq .CppType "int64_t" }}
                result.{{ .FieldName }} = (int64_t) std::stoll(h.value);
{{- else }}
                result.{{ .FieldName }} = ({{ .CppType }}) std::stod(h.value);
{{- end }}
            }
{{- end }}
        }
        return result;
    }
};
`))

// CppGenericHeaderClass renders a class representing a set of typed HTTP
// headers, with ToHeaders()/FromHeaders() to convert to/from the transport-level
// emi::EmiHttpHeader list.
func CppGenericHeaderClass(className string, columns []core.EmiHeader, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	headers := make([]cppRenderedHeaderField, 0, len(columns))
	for _, h := range columns {
		headers = append(headers, cppRenderedHeaderField{
			FieldName:  h.Name,
			HeaderName: h.Name,
			CppType:    cppHeaderType(DialectGeneric, h.Type),
		})
	}

	var buf bytes.Buffer
	if err := cppGenericHeaderTmpl.Execute(&buf, core.H{"className": className, "headers": headers}); err != nil {
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

var cppUnrealHeaderTmpl = template.Must(template.New("cppunrealheaders").Parse(`
class {{ .className }} {
public:
{{- range .headers }}
    TOptional<{{ .CppType }}> {{ .FieldName }};
{{- end }}

    TMap<FString, FString> ToHeaders() const {
        TMap<FString, FString> Result;
{{- range .headers }}
        if ({{ .FieldName }}.IsSet()) {
{{- if eq .CppType "FString" }}
            Result.Add(TEXT("{{ .HeaderName }}"), {{ .FieldName }}.GetValue());
{{- else if eq .CppType "bool" }}
            Result.Add(TEXT("{{ .HeaderName }}"), {{ .FieldName }}.GetValue() ? TEXT("true") : TEXT("false"));
{{- else }}
            Result.Add(TEXT("{{ .HeaderName }}"), FString::SanitizeFloat((double) {{ .FieldName }}.GetValue()));
{{- end }}
        }
{{- end }}
        return Result;
    }

    static {{ .className }} FromHeaders(const TMap<FString, FString>& Headers) {
        {{ .className }} Result;
{{- range .headers }}
        if (const FString* Found = Headers.Find(TEXT("{{ .HeaderName }}"))) {
{{- if eq .CppType "FString" }}
            Result.{{ .FieldName }} = *Found;
{{- else if eq .CppType "bool" }}
            Result.{{ .FieldName }} = (*Found == TEXT("true") || *Found == TEXT("1"));
{{- else }}
            Result.{{ .FieldName }} = ({{ .CppType }}) FCString::Atod(**Found);
{{- end }}
        }
{{- end }}
        return Result;
    }
};
`))

// CppUnrealHeaderClass mirrors CppGenericHeaderClass for the unreal dialect.
func CppUnrealHeaderClass(className string, columns []core.EmiHeader, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	headers := make([]cppRenderedHeaderField, 0, len(columns))
	for _, h := range columns {
		headers = append(headers, cppRenderedHeaderField{
			FieldName:  core.ToUpper(h.Name),
			HeaderName: h.Name,
			CppType:    cppHeaderType(DialectUnreal, h.Type),
		})
	}

	var buf bytes.Buffer
	if err := cppUnrealHeaderTmpl.Execute(&buf, core.H{"className": className, "headers": headers}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  className,
		SuggestedExtension: ".h",
		Tokens:             []core.GeneratedScriptToken{{Name: TOKEN_ROOT_CLASS, Value: className}},
	}, nil
}

// CppHeaderClass dispatches to the right dialect's header class renderer.
func CppHeaderClass(className string, columns []core.EmiHeader, dialect Dialect, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	if dialect == DialectUnreal {
		return CppUnrealHeaderClass(className, columns, ctx)
	}
	return CppGenericHeaderClass(className, columns, ctx)
}
