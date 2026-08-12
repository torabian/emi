package c

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// RecognizedComplex mirrors the same concept in the sibling generators: a
// custom/complex data type that's importable from somewhere, so a field
// referencing it (via `complex: "+Vector3"`) can be resolved to a real type
// instead of falling back to a raw cJSON node. ImportLocation is a header
// to #include (e.g. `"vector3.h"` or `<mymath/vector3.h>`).
type RecognizedComplex struct {
	Symbol         string
	ImportLocation string
}

type cRenderedStruct struct {
	TypeName    string
	Doc         string
	Fields      []cFieldPlan
	InlineEnums []*cRenderedEnum
}

// CCommonObjectContext configures a single struct-generation pass - used
// for dtos, request/response bodies, and any inline object/array field.
type CCommonObjectContext struct {
	RootTypeName        string
	RecognizedComplexes []RecognizedComplex
}

// renderStructs walks the field tree, flattening every object/array-of-object
// field into its own companion struct, and every inline `enum: of:` field
// into its own companion enum - both named by concatenating the parent
// chain, exactly like the sibling generators.
func renderStructs(fields []*core.EmiField, typeName string, prefixName string, complexes []RecognizedComplex) []cRenderedStruct {
	fieldPlans := make([]cFieldPlan, 0, len(fields))
	inlineEnums := []*cRenderedEnum{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		plan := cResolveField(f, prefixName, prefixName, complexes)
		fieldPlans = append(fieldPlans, plan)
		if plan.InlineEnum != nil {
			inlineEnums = append(inlineEnums, plan.InlineEnum)
		}
	}

	current := cRenderedStruct{
		TypeName:    prefixName,
		Doc:         typeName,
		Fields:      fieldPlans,
		InlineEnums: inlineEnums,
	}

	out := []cRenderedStruct{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		if f.Type == core.FieldTypeObject || f.Type == core.FieldTypeObjectNullable ||
			f.Type == core.FieldTypeArray || f.Type == core.FieldTypeArrayNullable ||
			f.Type == core.FieldTypeList || f.Type == core.FieldTypeListNullable {
			childPrefix := prefixName + core.ToUpper(f.Name)
			out = append(out, renderStructs(f.Fields, core.ToUpper(f.Name), childPrefix, complexes)...)
		}
	}

	// Children before the parent - a struct referencing another (by
	// pointer or by-value array element) needs that type's definition
	// already visible for `sizeof`/member access to compile.
	out = append(out, current)
	return out
}

var cStructTmplFuncs = template.FuncMap{"renderEnumDecl": renderCEnumDecl, "renderEnumImpl": renderCEnumImpl}

// Every function is `static inline` for the same header-only-library reason
// noted in c-enum.go. `_init`/`_free_fields` operate on an already-placed
// `self` (heap or a by-value array element); `_new`/`_free` add the
// malloc/free of `self` itself; `_from_json_into` populates an already-`
// _init`-ed `self`; `_from_json`/`_to_json` are the two convenient
// allocate-and-populate / build-and-return entry points.
var cStructTmpl = template.Must(template.New("cstruct").Funcs(cStructTmplFuncs).Parse(`
{{ range .structs }}
{{ range .InlineEnums }}
{{ renderEnumDecl . }}
{{ renderEnumImpl . }}
{{ end }}
/* {{ .Doc }} */
typedef struct {{ .TypeName }} {
{{- range .Fields }}
{{- range .Decl }}
    {{ . }}
{{- end }}
{{- end }}
{{- if not .Fields }}
    char _reserved; /* avoids a zero-sized struct - this type has no fields */
{{- end }}
} {{ .TypeName }};

/* Forward declarations: a self-referencing field (e.g. a relation whose
 * target: is this very struct) means _free_fields/_from_json_into below
 * can call _free/_from_json, which aren't defined until further down this
 * same file - static inline functions can be forward-declared exactly like
 * any other, so this resolves the ordering regardless of which fields end
 * up actually self-referencing. */
static inline void {{ .TypeName }}_init({{ .TypeName }}* self);
static inline void {{ .TypeName }}_free_fields({{ .TypeName }}* self);
static inline {{ .TypeName }}* {{ .TypeName }}_new(void);
static inline void {{ .TypeName }}_free({{ .TypeName }}* self);
static inline void {{ .TypeName }}_from_json_into({{ .TypeName }}* self, const cJSON* json);
static inline {{ .TypeName }}* {{ .TypeName }}_from_json(const cJSON* json);
static inline cJSON* {{ .TypeName }}_to_json(const {{ .TypeName }}* self);

static inline void {{ .TypeName }}_init({{ .TypeName }}* self) {
    if (!self) return;
{{- range .Fields }}
{{- range .Init }}
    {{ . }}
{{- end }}
{{- end }}
}

static inline void {{ .TypeName }}_free_fields({{ .TypeName }}* self) {
    if (!self) return;
{{- range .Fields }}
{{- range .FreeFields }}
    {{ . }}
{{- end }}
{{- end }}
}

static inline {{ .TypeName }}* {{ .TypeName }}_new(void) {
    {{ .TypeName }}* self = ({{ .TypeName }}*) malloc(sizeof({{ .TypeName }}));
    {{ .TypeName }}_init(self);
    return self;
}

static inline void {{ .TypeName }}_free({{ .TypeName }}* self) {
    if (!self) return;
    {{ .TypeName }}_free_fields(self);
    free(self);
}

static inline void {{ .TypeName }}_from_json_into({{ .TypeName }}* self, const cJSON* json) {
    if (!self || !json) return;
{{- range .Fields }}
{{- range .FromJson }}
    {{ . }}
{{- end }}
{{- end }}
}

static inline {{ .TypeName }}* {{ .TypeName }}_from_json(const cJSON* json) {
    {{ .TypeName }}* self = {{ .TypeName }}_new();
    {{ .TypeName }}_from_json_into(self, json);
    return self;
}

static inline cJSON* {{ .TypeName }}_to_json(const {{ .TypeName }}* self) {
    cJSON* json = cJSON_CreateObject();
    if (!self) return json;
{{- range .Fields }}
{{- range .ToJson }}
    {{ . }}
{{- end }}
{{- end }}
    return json;
}
{{ end }}
`))

func renderCStructs(structs []cRenderedStruct) (string, error) {
	var buf bytes.Buffer
	if err := cStructTmpl.Execute(&buf, core.H{"structs": structs}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// CCommonObjectGenerator generates one or more struct definitions (the
// root, plus any flattened nested object/array-of-object, plus any inline
// enum) out of a field tree. Used for dtos, action request/response
// bodies, and headers-adjacent inline shapes.
func CCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, cctx CCommonObjectContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	for _, symbol := range collectComplexSymbols(fields) {
		location := findComplexLocation(symbol, cctx.RecognizedComplexes)
		if location == "" {
			continue
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{Location: location})
	}
	for _, target := range core.CollectTargets(fields, cctx.RootTypeName) {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, HeaderIncludeDependency(target)...)
	}

	structs := renderStructs(fields, cctx.RootTypeName, cctx.RootTypeName, cctx.RecognizedComplexes)
	if len(structs) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: cctx.RootTypeName})
	}

	script, err := renderCStructs(structs)
	if err != nil {
		return nil, err
	}

	res.ActualScript = []byte(script)
	res.SuggestedFileName = core.ToSnakeCase(cctx.RootTypeName)
	res.SuggestedExtension = ".h"
	return res, nil
}

// HeaderIncludeDependency turns a `target: SomeDto` / `dto: SomeDto|OtherDto`
// style reference into #include dependencies pointing at the generated
// header for each name (files are always named after
// `ToSnakeCase(TypeName)+".h"`).
func HeaderIncludeDependency(dtoName string) []core.CodeChunkDependency {
	deps := []core.CodeChunkDependency{}
	for name := range strings.SplitSeq(dtoName, "|") {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		_, className := core.ParseDtoPath(name)
		deps = append(deps, core.CodeChunkDependency{Location: core.ToSnakeCase(className) + ".h"})
	}
	return deps
}

func collectComplexSymbols(fields []*core.EmiField) []string {
	var result []string
	var walk func([]*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field == nil {
				continue
			}
			if strings.Contains(field.Complex, "+") {
				result = append(result, strings.ReplaceAll(field.Complex, "+", ""))
			}
			if len(field.Fields) > 0 {
				walk(field.Fields)
			}
		}
	}
	walk(fields)
	return result
}

func findComplexLocation(complexName string, complexes []RecognizedComplex) string {
	for _, item := range complexes {
		if item.Symbol == complexName {
			return item.ImportLocation
		}
	}
	return ""
}

var TOKEN_ROOT_CLASS = "root.class"
var TOKEN_ORIGINAL_NAME = core.TOKEN_ORIGINAL_NAME
