package cpp

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type cppPathParam struct {
	WireName string // the raw `:name` placeholder in the url template
	Field    string
	CppType  string
	IsString bool
}

var cppGenericPathParamTmpl = template.Must(template.New("cppgenericpathparams").Parse(`
class {{ .TypeName }} {
public:
{{- range .Params }}
    {{ .CppType }} {{ .Field }};
{{- end }}

    std::string Apply(const std::string& templateUrl) const {
        std::string url = templateUrl;
{{- range .Params }}
{{- if .IsString }}
        url = emi::EmiUrlReplace(url, ":{{ .WireName }}", {{ .Field }});
{{- else }}
        url = emi::EmiUrlReplace(url, ":{{ .WireName }}", std::to_string({{ .Field }}));
{{- end }}
{{- end }}
        return url;
    }
};
`))

var cppUnrealPathParamTmpl = template.Must(template.New("cppunrealpathparams").Parse(`
class {{ .TypeName }} {
public:
{{- range .Params }}
    {{ .CppType }} {{ .Field }};
{{- end }}

    FString Apply(const FString& TemplateUrl) const {
        FString Url = TemplateUrl;
{{- range .Params }}
{{- if .IsString }}
        Url = Url.Replace(TEXT(":{{ .WireName }}"), *{{ .Field }});
{{- else if eq .CppType "bool" }}
        Url = Url.Replace(TEXT(":{{ .WireName }}"), {{ .Field }} ? TEXT("true") : TEXT("false"));
{{- else }}
        Url = Url.Replace(TEXT(":{{ .WireName }}"), *FString::SanitizeFloat((double) {{ .Field }}));
{{- end }}
{{- end }}
        return Url;
    }
};
`))

// CppActionPathParams extracts `/:id`-style placeholders out of an action's url
// and renders a small class to carry them, plus an Apply(url) method that
// substitutes them back into the template url. Returns (nil, nil) when the url
// has no placeholders at all.
func CppActionPathParams(action core.EmiRpcAction, dialect Dialect) (*core.CodeChunkCompiled, error) {
	placeholders := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders) == 0 {
		return nil, nil
	}

	params := make([]cppPathParam, 0, len(placeholders))
	for _, p := range placeholders {
		t := extractPrimitive(dialect, p.Type)
		// extractPrimitive resolves "string" to a real (non-empty) type here -
		// unlike the C target's own extractPrimitive, which has no "string" case
		// at all - so isString can't be inferred from t=="" the way c-path-
		// parameter.go does it; check the type string itself instead.
		isString := t == "" || p.Type == "string" || p.Type == ""
		if t == "" {
			if dialect == DialectUnreal {
				t = "FString"
			} else {
				t = "std::string"
			}
		}
		fieldName := p.Original
		if dialect == DialectUnreal {
			fieldName = core.ToUpper(core.NormaliseKey(p.Original))
		}
		params = append(params, cppPathParam{WireName: p.Original, Field: fieldName, CppType: t, IsString: isString})
	}

	className := core.ToUpper(core.NormaliseKey(action.GetName()))
	typeName := fmt.Sprintf("%vPathParameters", className)
	extension := ".hpp"
	tmpl := cppGenericPathParamTmpl
	if dialect == DialectUnreal {
		typeName = "F" + typeName
		extension = ".h"
		tmpl = cppUnrealPathParamTmpl
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, core.H{"TypeName": typeName, "Params": params}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  typeName,
		SuggestedExtension: extension,
		Tokens:             []core.GeneratedScriptToken{{Name: TOKEN_ROOT_CLASS, Value: typeName}},
	}, nil
}
