package cpp

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// CppCommonObjectContext configures a single class-generation pass - used for
// dtos, request/response bodies, and any inline object/array field. RootClassName
// is the *logical* name (e.g. "WidgetDto") - the unreal dialect's own `F`-prefix
// convention (see ueStructName) is applied at render time, never baked into
// RootClassName itself, so the same context works for either dialect.
type CppCommonObjectContext struct {
	Dialect             Dialect
	RootClassName       string
	RecognizedComplexes []RecognizedComplex
}

// CppCommonObjectGenerator generates one or more class/struct definitions (the
// root, plus any flattened nested object/array-of-object, plus any inline enum)
// out of a field tree, dispatching to the right dialect's own renderer. Used for
// dtos, action request/response bodies, and headers-adjacent inline shapes.
func CppCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, cctx CppCommonObjectContext) (*core.CodeChunkCompiled, error) {
	if cctx.Dialect == DialectUnreal {
		return cppUnrealCommonObjectGenerator(fields, cctx)
	}
	return cppGenericCommonObjectGenerator(fields, cctx)
}

// ==============================================================================
// generic dialect
// ==============================================================================

type cppRenderedGenericClass struct {
	ClassName     string
	Doc           string
	DeclLines     []string
	ToJsonLines   []string
	FromJsonLines []string
	InlineEnums   []*cppRenderedEnum
}

// renderCppGenericClasses walks the field tree, flattening every object/array-of-
// object field into its own top-level class, and every inline `enum: of:` field
// into its own top-level enum - both named by concatenating the parent chain,
// exactly like every sibling generator's own equivalent.
func renderCppGenericClasses(fields []*core.EmiField, className string, prefixName string, complexes []RecognizedComplex) []cppRenderedGenericClass {
	current := cppRenderedGenericClass{ClassName: prefixName, Doc: className}

	for _, f := range fields {
		if f == nil {
			continue
		}
		plan := cppGenericResolveField(f, prefixName, prefixName, complexes)
		current.DeclLines = append(current.DeclLines, plan.Decl...)
		current.ToJsonLines = append(current.ToJsonLines, plan.ToJson("")...)
		current.FromJsonLines = append(current.FromJsonLines, plan.FromJson("result.")...)
		if plan.InlineEnum != nil {
			current.InlineEnums = append(current.InlineEnums, plan.InlineEnum)
		}
	}

	out := []cppRenderedGenericClass{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		if f.Type == core.FieldTypeObject || f.Type == core.FieldTypeObjectNullable ||
			f.Type == core.FieldTypeArray || f.Type == core.FieldTypeArrayNullable ||
			f.Type == core.FieldTypeList || f.Type == core.FieldTypeListNullable {
			childPrefix := prefixName + core.ToUpper(f.Name)
			out = append(out, renderCppGenericClasses(f.Fields, core.ToUpper(f.Name), childPrefix, complexes)...)
		}
	}

	// Children before the parent - a class referencing another (by value, or as a
	// std::vector<T> element) needs that type's complete definition already
	// visible, or the compiler can't compute this class's own size.
	out = append(out, current)
	return out
}

var cppGenericClassTmplFuncs = template.FuncMap{"renderEnum": func(e *cppRenderedEnum) string { return renderCppGenericEnumDecl(e) }}

var cppGenericClassTmpl = template.Must(template.New("cppgenericclass").Funcs(cppGenericClassTmplFuncs).Parse(`
{{ range .classes }}
{{ range .InlineEnums }}
{{ renderEnum . }}
{{ end }}
// {{ .Doc }}
class {{ .ClassName }} {
public:
{{- range .DeclLines }}
    {{ . }}
{{- end }}

    {{ .ClassName }}() = default;

    cJSON* ToJson() const {
        cJSON* json = cJSON_CreateObject();
{{- range .ToJsonLines }}
        {{ . }}
{{- end }}
        return json;
    }

    static {{ .ClassName }} FromJson(const cJSON* json) {
        {{ .ClassName }} result;
        if (!json) return result;
{{- range .FromJsonLines }}
        {{ . }}
{{- end }}
        return result;
    }

    // Parses a full JSON document into a {{ .ClassName }} - a malformed document
    // simply leaves every field at its default (matches FromJson's own "always
    // safe, never throws" contract).
    static {{ .ClassName }} Parse(const std::string& text) {
        cJSON* root = cJSON_Parse(text.c_str());
        {{ .ClassName }} result = FromJson(root);
        cJSON_Delete(root);
        return result;
    }

    std::string Dump() const {
        cJSON* json = ToJson();
        char* rendered = cJSON_PrintUnformatted(json);
        std::string out = rendered ? rendered : "{}";
        cJSON_free(rendered);
        cJSON_Delete(json);
        return out;
    }
};
{{ end }}
`))

func renderCppGenericClassScript(classes []cppRenderedGenericClass) (string, error) {
	var buf bytes.Buffer
	if err := cppGenericClassTmpl.Execute(&buf, core.H{"classes": classes}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func cppGenericCommonObjectGenerator(fields []*core.EmiField, cctx CppCommonObjectContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	for _, symbol := range collectComplexSymbols(fields) {
		location := findComplexLocation(symbol, cctx.RecognizedComplexes)
		if location == "" {
			continue
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{Location: location})
	}
	for _, target := range core.CollectTargets(fields, cctx.RootClassName) {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, CppHeaderIncludeDependency(target)...)
	}

	classes := renderCppGenericClasses(fields, cctx.RootClassName, cctx.RootClassName, cctx.RecognizedComplexes)
	if len(classes) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: cctx.RootClassName})
	}

	script, err := renderCppGenericClassScript(classes)
	if err != nil {
		return nil, err
	}

	res.ActualScript = []byte(script)
	res.SuggestedFileName = cctx.RootClassName
	res.SuggestedExtension = ".hpp"
	return res, nil
}

// CppHeaderIncludeDependency turns a `target: SomeDto` / `dto: SomeDto|OtherDto`
// style reference into #include dependencies pointing at the generated header for
// each name (generic-dialect files are always named after the class name, plain
// `.hpp`).
func CppHeaderIncludeDependency(dtoName string) []core.CodeChunkDependency {
	deps := []core.CodeChunkDependency{}
	for name := range strings.SplitSeq(dtoName, "|") {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		_, className := core.ParseDtoPath(name)
		deps = append(deps, core.CodeChunkDependency{Location: className + ".hpp"})
	}
	return deps
}

// ==============================================================================
// unreal dialect
// ==============================================================================

type cppRenderedUnrealClass struct {
	ClassName   string // already `F`-prefixed
	Doc         string
	DeclLines   []string
	InlineEnums []*cppRenderedEnum
}

func renderCppUnrealClasses(fields []*core.EmiField, className string, prefixName string, complexes []RecognizedComplex) []cppRenderedUnrealClass {
	selfTypeName := ueStructName(prefixName)
	current := cppRenderedUnrealClass{ClassName: selfTypeName, Doc: className}

	for _, f := range fields {
		if f == nil {
			continue
		}
		plan := cppUnrealResolveField(f, prefixName, selfTypeName, complexes)
		current.DeclLines = append(current.DeclLines, plan.Decl...)
		if plan.InlineEnum != nil {
			current.InlineEnums = append(current.InlineEnums, plan.InlineEnum)
		}
	}

	out := []cppRenderedUnrealClass{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		if f.Type == core.FieldTypeObject || f.Type == core.FieldTypeObjectNullable ||
			f.Type == core.FieldTypeArray || f.Type == core.FieldTypeArrayNullable ||
			f.Type == core.FieldTypeList || f.Type == core.FieldTypeListNullable {
			childPrefix := prefixName + core.ToUpper(f.Name)
			out = append(out, renderCppUnrealClasses(f.Fields, core.ToUpper(f.Name), childPrefix, complexes)...)
		}
	}

	out = append(out, current)
	return out
}

var cppUnrealClassTmplFuncs = template.FuncMap{"renderEnum": func(e *cppRenderedEnum) string { return renderCppUnrealEnumDecl(e) }}

var cppUnrealClassTmpl = template.Must(template.New("cppunrealclass").Funcs(cppUnrealClassTmplFuncs).Parse(`
{{ range .classes }}
{{ range .InlineEnums }}
{{ renderEnum . }}
{{ end }}
// {{ .Doc }}
USTRUCT(BlueprintType)
struct {{ .ClassName }} {
    GENERATED_BODY()

{{- range .DeclLines }}
    {{ . }}
{{- end }}
};
{{ end }}
`))

func renderCppUnrealClassScript(classes []cppRenderedUnrealClass) (string, error) {
	var buf bytes.Buffer
	if err := cppUnrealClassTmpl.Execute(&buf, core.H{"classes": classes}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func cppUnrealCommonObjectGenerator(fields []*core.EmiField, cctx CppCommonObjectContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}
	rootType := ueStructName(cctx.RootClassName)

	for _, symbol := range collectComplexSymbols(fields) {
		location := findComplexLocation(symbol, cctx.RecognizedComplexes)
		if location == "" {
			continue
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{Location: location})
	}
	for _, target := range core.CollectTargets(fields, cctx.RootClassName) {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{Location: ueStructName(target) + ".h"})
	}

	classes := renderCppUnrealClasses(fields, cctx.RootClassName, cctx.RootClassName, cctx.RecognizedComplexes)
	if len(classes) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: rootType})
	}

	script, err := renderCppUnrealClassScript(classes)
	if err != nil {
		return nil, err
	}

	res.ActualScript = []byte(script)
	res.SuggestedFileName = rootType
	res.SuggestedExtension = ".h"
	return res, nil
}
