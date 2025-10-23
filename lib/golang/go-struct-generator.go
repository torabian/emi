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
}

type goRenderedStruct struct {
	ClassName      string
	Fields         []goRenderedField
	LateInitFields []goRenderedField
	Signature      string
	GoDoc          string
	SubClasses     []goRenderedStruct
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

type GoCommonStructContext struct {

	// The class name which will be used to generate nested classes,
	// in case of array or object
	RootClassName string

	// List of allowed complexes type to be used on fields
	RecognizedComplexes []RecognizedComplex
}

func goRenderStructs(fields []*core.EmiField, className, treeLocation string, fieldDepth string, prefixName string, ctx core.MicroGenContext, goctx GoCommonStructContext) []goRenderedStruct {
	if len(fields) == 0 {
		return nil
	}

	GoDoc := NewGoDoc("  ").Add(fmt.Sprintf("The base class definition for %v", core.ToLower(className)))
	signature := fmt.Sprintf("type %v struct", prefixName)

	fieldsRendered := goRenderFieldsShallow(fields, treeLocation, fieldDepth, ctx, goctx)

	currentClass := goRenderedStruct{
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
			currentClass.SubClasses = append(currentClass.SubClasses, goRenderStructs(f.Fields, childName, treeLocation+"."+childName, newDepth, prefixName+childName, ctx, goctx)...)
		}
	}

	return []goRenderedStruct{currentClass}
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

func findComplexLocation(complexName string, goctx GoCommonStructContext) string {

	for _, item := range goctx.RecognizedComplexes {
		if item.Symbol == complexName {
			return item.ImportLocation
		}
	}

	// If no location, then basically skip it.
	return ""
}

// Generates a class with getters and setters.
func GoCommonStructGenerator(fields []*core.EmiField, ctx core.MicroGenContext, goctx GoCommonStructContext) (*core.CodeChunkCompiled, error) {

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

	renderedClasses := goRenderStructs(fields, goctx.RootClassName, goctx.RootClassName, "", core.ToUpper(goctx.RootClassName), ctx, goctx)
	if len(renderedClasses) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: renderedClasses[0].ClassName})
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_OBJ_CLASS, Value: renderedClasses[0].ClassName})
	}

	const tmpl = `

	package main

	import emi "test/goruntime"

{{ define "printClass" }}
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

  
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"renderedClasses": renderedClasses,
	}); err != nil {
		return nil, err
	}

	result := buf.String()

	res.ActualScript = []byte(result)
	res.SuggestedFileName = goctx.RootClassName

	return res, nil
}

func goComputedField(field *core.EmiField) string {
	prefix := "emi."
	switch field.Type {

	case "string", "text", "html", "enum":
		return "string"
	case "string?", "text?", "html?", "enum?":
		return prefix + "Nullable[string]"
		// return prefix + "String"
	case "duration?":
		return prefix + "Duration"

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
		return "[]" + field.Primitive
	case "int64", "int32", "int", "float64", "float32", "bool":
		return string(field.Type)
	case "int64?", "int32?", "int?", "float64?", "float32?", "bool?":
		return prefix + "Nullable[" + strings.ReplaceAll(core.ToLower(string(field.Type)), "?", "") + "]"
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
	case "date":
		return prefix + "XDate"
	case "datetime":
		return "*" + prefix + "XDateTime"
	default:
		return "interface{}"
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
		return fmt.Sprintf("[]%v", prefix)
	case core.FieldTypeObjectNullable:
		return fmt.Sprintf("emi.Nullable[%v]", prefix)
	case core.FieldTypeArrayNullable:
		return fmt.Sprintf("emi.Nullable[[]%v]", prefix)
	default:
		return goComputedField(field)
	}
}
