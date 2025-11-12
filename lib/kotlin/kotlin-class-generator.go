package kotlin

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
}

type renderedClass struct {
	ClassName      string
	Fields         []renderedField
	LateInitFields []renderedField
	Signature      string
	GoDoc          string
	SubClasses     []renderedClass
	ClassTypePath  string
	ClassNamePath  string
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

type commonClassContext struct {

	// The class name which will be used to generate nested classes,
	// in case of array or object
	RootClassName string

	// List of allowed complexes type to be used on fields
	RecognizedComplexes []RecognizedComplex

	// the package location of the emi runtime.
	// If the project wants to copy that and override we use this
	EmiLocation string
}

func renderClasses(fields []*core.EmiField, className, treeLocation string, fieldDepth string, prefixName string, ctx core.MicroGenContext, goctx commonClassContext) []renderedClass {
	if len(fields) == 0 {
		return nil
	}

	GoDoc := NewGoDoc("  ").Add(fmt.Sprintf("The base class definition for %v", core.ToLower(className)))
	signature := fmt.Sprintf("@Serializable\r\ndata class %v", prefixName)

	fieldsRendered := renderFieldsShallow(fields, treeLocation, fieldDepth, ctx, goctx)

	currentClass := renderedClass{
		ClassName:     core.ToUpper(className),
		Fields:        fieldsRendered,
		ClassNamePath: treeLocation,
		ClassTypePath: treeAsType(treeLocation),
		GoDoc:         GoDoc.String(),
		Signature:     signature,
	}

	for _, f := range fields {
		if f != nil && (f.Type == core.FieldTypeObject || f.Type == core.FieldTypeObjectNullable || f.Type == core.FieldTypeArray || f.Type == core.FieldTypeArrayNullable) {
			childName := core.ToUpper(f.Name)
			newDepth := fieldDepth + "." + f.Name
			if fieldDepth == "" {
				newDepth = f.Name
			}
			currentClass.SubClasses = append(currentClass.SubClasses, renderClasses(f.Fields, childName, treeLocation+"."+childName, newDepth, prefixName+childName, ctx, goctx)...)
		}
	}

	return []renderedClass{currentClass}
}

// Only finds the complex classes, those have + prefix or affix
func CollectComplexClasses(fields []*core.EmiField) []string {
	var result []string

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
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

func findComplexLocation(complexName string, goctx commonClassContext) string {

	for _, item := range goctx.RecognizedComplexes {
		if item.Symbol == complexName {
			return item.ImportLocation
		}
	}

	// If no location, then basically skip it.
	return ""
}

// Generates a class with getters and setters.
func KotlinCommonStructGenerator(fields []*core.EmiField, ctx core.MicroGenContext, goctx commonClassContext) (*core.CodeChunkCompiled, error) {

	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{},
	}

	usedComplexes := CollectComplexClasses(fields)
	for _, item := range usedComplexes {

		location := findComplexLocation(item, goctx)
		if location == "" {
			continue
		}

		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Objects:  []string{item},
			Location: location,
		})
	}

	collectTargets := core.CollectTargets(fields)
	for _, item := range collectTargets {
		m := castDtoNameToCodeChunk(item)
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, m.CodeChunkDependensies...)
	}

	renderedClasses := renderClasses(fields, goctx.RootClassName, goctx.RootClassName, "", core.ToUpper(goctx.RootClassName), ctx, goctx)
	if len(renderedClasses) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: renderedClasses[0].ClassName})
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_OBJ_CLASS, Value: renderedClasses[0].ClassName})
	}

	const tmpl = `

{{ if .emiRuntimeLocation }}
	import emi "{{ .emiRuntimeLocation }}"
{{ end }}

{{ define "printClass" }}
{{ .GoDoc }}
{{ .Signature  }} (
	{{ range .Fields }}
		{{ .PrivateField }},
	{{ end }}
)

	{{ range .SubClasses }}
		{{ template "printClass" . }}
	{{ end }}

	

{{ end }}
{{ range .renderedClasses }}
	{{ template "printClass" . }}
{{ end }}

  
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"renderedClasses":    renderedClasses,
		"emiRuntimeLocation": goctx.EmiLocation,
	}); err != nil {
		return nil, err
	}

	result := buf.String()

	res.ActualScript = []byte(result)
	res.SuggestedFileName = goctx.RootClassName + ".kt"

	return res, nil
}

func goComputedField(field *core.EmiField) string {
	switch field.Type {

	case "string", "text", "html", "enum":
		return "string"
	case "string?", "text?", "html?", "enum?":
		return "String?"
	case "one":
		if field.Module != "" {
			return field.Module + "." + field.Target
		}
		return field.Target
	case "array":
		return field.PublicName()
	case "any":
		return "interface{}"
	case "collection":
		if field.Module != "" {
			return field.Module + "." + field.Target
		}
		return field.Target
	case "slice":
		return fmt.Sprintf("List<%v>", core.ToUpper(field.Primitive))
	case "int64", "int32", "int", "float64", "float32", "bool":
		return string(field.Type)
	case "int64?", "int32?", "int?", "float64?", "float32?", "bool?":
		return strings.ReplaceAll(core.ToLower(string(field.Type)), "?", "") + "?"
	case "Timestamp":
		return "*string"
	case "datenano":
		return "int64"
	case "boolean":
		return "*bool"
	case "double":
		return "*float64"
	case "object", "embed":
		return field.PublicName()
	case "json":
		return "JSON"
	default:
		return "Any"
	}
}

func goFieldTypeOnNestedClasses(field *core.EmiField, parentChain string) string {
	if field == nil {
		return ""
	}
	prefix := core.ToUpper(parentChain) + core.ToUpper(field.Name)
	switch field.Type {
	case core.FieldTypeObject:
		return fmt.Sprintf(" %v", prefix)
	case core.FieldTypeArray:
		return fmt.Sprintf("List<%v>", prefix)
	case core.FieldTypeObjectNullable:
		return fmt.Sprintf("List<%v>", prefix)
	case core.FieldTypeArrayNullable:
		return fmt.Sprintf("List<%v>", prefix)
	default:
		return goComputedField(field)
	}
}
