package golang

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type RecognizedComplex struct {
	Symbol string

	ImportLocation string

	// Useful on languages such as golang or java
	Namespace string
}

type goRenderedStruct struct {
	ClassName       string
	Fields          []goRenderedField
	LateInitFields  []goRenderedField
	Signature       string
	GoDoc           string
	SubClasses      []goRenderedStruct
	ClassTypePath   string
	FullClassName   string
	ClassNamePath   string
	CliCastFunction string
}

// Is is a bit weird function I am adding to capture type.
// basically, Name.Object.Item will become NameType.ObjectType.ItemType
func treeAsType(treeLocation string) string {
	parts := strings.Split(treeLocation, ".")
	for i, p := range parts {
		parts[i] = p + "Type"
	}
	return strings.Join(parts, ".")
}

type GoCommonStructContext struct {

	// The class name which will be used to generate nested classes,
	// in case of array or object
	RootClassName string

	// List of allowed complexes type to be used on fields
	RecognizedComplexes []RecognizedComplex

	// the package location of the emi runtime.
	// If the project wants to copy that and override we use this
	EmiLocation string
}

func goRenderStructs(fields []*core.EmiField, className, treeLocation string, fieldDepth string, prefixName string, ctx core.MicroGenContext, goctx GoCommonStructContext) []goRenderedStruct {
	if len(fields) == 0 {
		return nil
	}

	GoDoc := NewGoDoc("  ").Add(fmt.Sprintf("The base class definition for %v", core.ToLower(className)))
	signature := fmt.Sprintf("type %v struct", prefixName)

	fieldsRendered := goRenderFieldsShallow(fields, treeLocation, goctx)

	currentClass := goRenderedStruct{
		ClassName:       core.ToUpper(className),
		Fields:          fieldsRendered,
		ClassNamePath:   treeLocation,
		CliCastFunction: fmt.Sprintf("func Cast%vFromCli(c emigo.CliCastable) %v", prefixName, prefixName),
		ClassTypePath:   treeAsType(treeLocation),
		GoDoc:           GoDoc.String(),
		FullClassName:   prefixName,
		Signature:       signature,
	}

	for _, f := range fields {
		if f != nil && (f.Type == core.FieldTypeObject || f.Type == core.FieldTypeObjectNullable || f.Type == core.FieldTypeArray || f.Type == core.FieldTypeArrayNullable) {
			childName := core.ToUpper(f.Name)
			newDepth := fieldDepth + f.Name
			if fieldDepth == "" {
				newDepth = f.Name
			}
			currentClass.SubClasses = append(currentClass.SubClasses, goRenderStructs(f.Fields, childName, treeLocation+childName, newDepth, prefixName+childName, ctx, goctx)...)
		}
	}

	return []goRenderedStruct{currentClass}
}

func CollectComplexClasses(fields []*core.EmiField) []string {
	var result []string

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field.Complex != "" && field.Type == core.FieldTypeComplex {
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

// Detects if inside fields, complex type has been used.
func DetectIfComplexIsUsed(fields []*core.EmiField) bool {
	var result bool

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field.Complex != "" && field.Type == core.FieldTypeComplex {
				result = true
			}

			if !result {
				if len(field.Fields) > 0 {
					walk(field.Fields)
				}
			}
		}
	}

	walk(fields)
	return result
}

func DetectIfEmiGoIsUsed(fields []*core.EmiField) bool {

	// As of recent changes, since we are having cli tools,
	// always emigo is included, it seems.
	if len(fields) > 0 {
		return true
	}

	var result bool = false

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			// At least in this case, a nullable value has been used.
			if strings.Contains(string(field.Type), "?") {
				result = true
				break
			}
			if len(field.Fields) > 0 {
				walk(field.Fields)
			}
		}
	}

	walk(fields)
	return result
}

func findComplexGlobalDefinition(complexName string, goctx GoCommonStructContext) *RecognizedComplex {

	for _, item := range goctx.RecognizedComplexes {
		if item.Symbol == complexName {
			return &item
		}
	}

	return nil
}

func findComplexLocation(complexName string, goctx GoCommonStructContext) string {

	for _, item := range goctx.RecognizedComplexes {
		if item.Symbol == complexName {
			return item.ImportLocation
		}
	}

	// If no location, then basically skip it.
	return ""
}

// In golang, only if the external package name is required we collect target.
// By default, the target is supposed to be inside the same package.
func CollectTargets(fields []*core.EmiField, currentName string) []string {
	var result []string

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field.Provider != "" {
				if field.Provider != currentName {
					result = append(result, field.Provider)
				}
			}

			if len(field.Fields) > 0 {
				walk(field.Fields)
			}
		}
	}

	walk(fields)
	return result
}

// Generates a class with getters and setters.
func GoCommonStructGenerator(fields []*core.EmiField, ctx core.MicroGenContext, goctx GoCommonStructContext) (*core.CodeChunkCompiled, error) {

	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Location: "encoding/json",
			},
		},
	}

	// Complexes are having a weird casting mechanism
	if DetectIfComplexIsUsed(fields) {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Location: "encoding",
		})
	}

	includeEmiGo := DetectIfEmiGoIsUsed(fields)
	emiLocation := ""

	if includeEmiGo {
		emiLocation = goctx.EmiLocation
	}

	usedComplexes := CollectComplexClasses(fields)
	for _, item := range usedComplexes {

		location := findComplexLocation(item, goctx)
		if location == "" {
			continue
		}

		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Objects:  []string{},
			Location: location,
		})
	}

	// For golang actually, the core collect target doesn't work
	collectTargets := CollectTargets(fields, goctx.RootClassName)

	for _, item := range collectTargets {
		m := castDtoNameToCodeChunk(item)
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, m.CodeChunkDependensies...)
	}

	renderedClasses := goRenderStructs(fields, goctx.RootClassName, goctx.RootClassName, "", core.ToUpper(goctx.RootClassName), ctx, goctx)
	if len(renderedClasses) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: renderedClasses[0].ClassName})
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_OBJ_CLASS, Value: renderedClasses[0].ClassName})
	}

	const tmpl = `

{{ if .emiRuntimeLocation }}
	import emigo "{{ .emiRuntimeLocation }}"
{{ end }}


{{ define "printClass" }}

func Get{{ .FullClassName }}CliFlags(prefix string) []emigo.CliFlag {

	return []emigo.CliFlag{
		{{ range .Fields }}
		{
			Name: prefix + "{{ .CliName }}",
			Type: "{{ .Type }}",
			{{ if eq .Type "object" }}
			Children: Get{{ $.FullClassName }}{{upper .Name}}CliFlags("{{ .CliName }}-"),
			{{ end }}
		},
		{{ end }}
	}
}

{{ .CliCastFunction }} {
	data := {{ .FullClassName }}{}

	{{ range .Fields }}
		{{ if .CliCaptureStatement }}
			{{ .CliCaptureStatement}}
		{{ end }}
	{{ end }}

	return data
}

{{ .GoDoc }}
{{ .Signature  }} {
	{{ range .Fields }}
		{{ .PrivateField }}
	{{ end }}
}



	{{ range .SubClasses }}
		{{ template "printClass" . }}
	{{ end }}


{{ end }}
{{ range .renderedClasses }}
	{{ template "printClass" . }}
{{ end }}
func (x *{{ .rootClass }}) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

  
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"renderedClasses":    renderedClasses,
		"emiRuntimeLocation": emiLocation,
		"rootClass":          core.ToUpper(goctx.RootClassName),
	}); err != nil {
		return nil, err
	}

	result := buf.String()

	res.ActualScript = []byte(result)
	res.SuggestedFileName = goctx.RootClassName + ".go"

	return res, nil
}

func goPrimitiveDetect(fieldType string) string {
	prefix := "emigo."
	switch fieldType {

	case "string", "text", "html", "enum":
		return "string"
	case "string?", "text?", "html?", "enum?":
		return prefix + "Nullable[string]"
	case "int64", "int32", "int", "float64", "float32", "bool", "uint", "uint32", "uint64", "uint8", "uint16":
		return string(fieldType)
	case "int64?", "int32?", "int?", "float64?", "float32?", "bool?", "uint?", "uint32?", "uint64?", "uint8?", "uint16?":
		return prefix + "Nullable[" + strings.ReplaceAll(core.ToLower(string(fieldType)), "?", "") + "]"
	}

	return ""
}

func goListAndObjectTypes(field fieldLike) string {

	prefix := "emigo."
	switch field.GetType() {

	case core.FieldTypeOne:
		if field.GetModule() != "" {
			return field.GetModule() + "." + field.GetTarget()
		}
		return field.GetTarget()
	case core.FieldTypeArray:
		return field.PublicName()
	case core.FieldTypeCollection:
		if field.GetModule() != "" {
			return "[]" + field.GetModule() + "." + field.GetTarget()
		}
		return "[]" + field.GetTarget()
	case core.FieldTypeSlice:
		return "[]" + goPrimitiveDetect(field.GetPrimitive())
	case core.FieldTypeSliceNullable:
		return prefix + "Nullable[[]" + goPrimitiveDetect(field.GetPrimitive()) + "]"
	case core.FieldTypeObject:
		return field.PublicName()
	default:
		return "interface{}"
	}
}

func goComputedField(field fieldLike) string {

	if goprimitive := goPrimitiveDetect(string(field.GetType())); goprimitive != "" {
		return goprimitive
	}

	if computedType := goListAndObjectTypes(field); computedType != "" {
		if core.IsNullable(string(field.GetType())) {
			return fmt.Sprintf("emigo.Nullable[%v]", computedType)
		} else {
			return computedType
		}
	}

	return "interface{}"
}

func goFieldTypeOnNestedClasses(
	field fieldLike,
	parentChain string,
	goctx GoCommonStructContext,
) string {
	if field == nil {
		return ""
	}
	prefix := core.ToUpper(parentChain) + core.ToUpper(field.GetName())
	switch field.GetType() {
	case core.FieldTypeComplex:

		globalDef := findComplexGlobalDefinition(field.GetComplex(), goctx)
		if globalDef != nil && globalDef.Namespace != "" {
			return fmt.Sprintf("%v.%v", globalDef.Namespace, field.GetComplex())
		}

		return fmt.Sprintf(" %v", field.GetComplex())
	case core.FieldTypeObject:
		return fmt.Sprintf(" %v", prefix)
	case core.FieldTypeArray:
		return fmt.Sprintf("[]%v", prefix)
	case core.FieldTypeObjectNullable:
		return fmt.Sprintf("emigo.Nullable[%v]", prefix)
	case core.FieldTypeArrayNullable:
		return fmt.Sprintf("emigo.Nullable[[]%v]", prefix)
	default:
		return goComputedField(field)
	}
}
