// Renders the common object, such as entities, dtos.

package js

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type jsRenderedDataClass struct {
	ClassName            string
	Fields               []jsRenderedField
	Signature            string
	JsDoc                string
	SubClasses           []jsRenderedDataClass
	ClassStaticFunctions []string
}

// Each field when rendered, becomes like this
type jsRenderedField struct {
	Name                    string
	Type                    string
	Children                []jsRenderedField
	Output                  string
	GetterFunc              string
	SetterCallInConstructor string
	SetterFunc              string
}

// Some field types such as array and object,
// need to have the correct generated class to be assigned to them.
func jsFieldTypeOnNestedClasses(field *core.Module3Field, parentChain string) string {
	if field == nil {
		return ""
	}
	if field.Type == core.FIELD_TYPE_ARRAY || field.Type == core.FIELD_TYPE_OBJECT {
		return core.ToUpper(parentChain) + "." + core.ToUpper(field.Name)
	}

	return TsComputedField(field, false)
}

func tsFieldTypeOnNestedClasses(field *core.Module3Field, parentChain string) string {
	if field == nil {
		return ""
	}
	if field.Type == core.FIELD_TYPE_OBJECT {
		return fmt.Sprintf("InstanceType<typeof %v>", core.ToUpper(parentChain)+"."+core.ToUpper(field.Name))
	}
	if field.Type == core.FIELD_TYPE_ARRAY {
		return fmt.Sprintf("InstanceType<typeof %v>[]", core.ToUpper(parentChain)+"."+core.ToUpper(field.Name))
	}

	return TsComputedField(field, false)
}

var setterTemplate = template.Must(template.New("setter").Parse(`
{{.Jsdoc}}
set{{.UpperName}}(value{{if .FieldType}}: {{.FieldType}}{{end}}) {
	{{if eq .Type "array"}}
	// Only accept array types
	if (!Array.isArray(value) && value !== null && value !== undefined) {
		return this
	}

	// If the value is array, we need to check first item, if is instance of the class
	if (value.length > 0 && value[0] instanceof {{.ArrayClass}}) {
		this["{{.Name}}"] = value
	} else {
		this["{{.Name}}"] = value.map(item => new {{.ArrayClass}}(item))
	}
	{{else}}
	this["{{.Name}}"] = value
	{{end}}
	return this
}
`))

func jsClassSetterFunction(jsFieldType string, field *core.Module3Field) (*JsdocComment, string) {
	setterjsdoc := NewJsDoc("  ")
	setterjsdoc.Add(fmt.Sprintf("@param {%v}", jsFieldType))
	setterjsdoc.Add(fmt.Sprintf("@description %v", field.Description))

	data := map[string]any{
		"Jsdoc":      setterjsdoc.String(),
		"UpperName":  core.ToUpper(field.Name),
		"Name":       field.Name,
		"FieldType":  jsFieldType,
		"Type":       field.Type,
		"ArrayClass": "GetSinglePostRes.Histories", // maybe dynamic later
	}

	var buf bytes.Buffer
	if err := setterTemplate.Execute(&buf, data); err != nil {
		panic(err)
	}

	return setterjsdoc, buf.String()
}
func getNullableDefaultValue(field *core.Module3Field) string {
	// In case the feild has default value, we need to extract it.
	if field.Default != nil {
		switch v := field.Default.(type) {
		case string:
			return fmt.Sprintf(" = %q", v)
		case int, int64, float64, bool:
			return fmt.Sprintf(" = %v", v)
		default:
			b, _ := json.Marshal(v)
			return " = '" + string(b) + "'"
		}
	}

	return "null"
}

func jsGetSafeFieldValue(field *core.Module3Field) string {
	if field == nil {
		return "null"
	}

	// return default if explicitly defined
	if field.Default != nil {
		switch v := field.Default.(type) {
		case string:
			return fmt.Sprintf("%q", v)
		case int, int64, float64, bool:
			return fmt.Sprintf("%v", v)
		default:
			b, _ := json.Marshal(v)
			return string(b)
		}
	}

	switch field.Type {
	case "array", "arrayP", "many2many":
		return "[]"
	case "json", "object", "embed":
		return "{}"
	case "string", "text":
		return "\"\""
	case "float", "float32", "float64", "float?", "float32?", "float64?":
		return "0.0"
	case "int", "int32", "int64", "int?", "int32?", "int64?":
		return "0"
	case "money?":
		return `{"amount":0,"currency":"USD"}`
	case "computed":
		return "null" // maybe skip or undefined?
	case "any":
		return "null"
	default:
		return "null"
	}
}

func jsRenderField(field *core.Module3Field, parentChain string, ctx core.MicroGenContext) jsRenderedField {

	jsFieldType := jsFieldTypeOnNestedClasses(field, parentChain)
	tsFieldType := tsFieldTypeOnNestedClasses(field, parentChain)
	isFieldNullable := strings.Contains(field.Type, "?")

	jsdoc := NewJsDoc("  ")
	jsdoc.Add(fmt.Sprintf("@type {%v}", jsFieldType))
	jsdoc.Add(fmt.Sprintf("@description %v", field.Description))

	// Javascript field definition
	output := fmt.Sprintf("%v %v;", jsdoc.String(), field.PrivateName())
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	// In typescript, we need to react
	if isTypeScript {
		if isFieldNullable {
			nullableDefaultValue := getNullableDefaultValue(field)
			output = fmt.Sprintf("%v %v?: %v | null %v", jsdoc.String(), field.PrivateName(), tsFieldType, nullableDefaultValue)
		} else {
			output = fmt.Sprintf("%v %v: %v = %v", jsdoc.String(), field.PrivateName(), tsFieldType, jsGetSafeFieldValue(field))
		}

	}

	getterjsdoc := NewJsDoc("  ")
	getterjsdoc.Add(fmt.Sprintf("@returns {%v}", jsFieldType))
	getterjsdoc.Add(fmt.Sprintf("@description %v", field.Description))
	getterFunc := getterjsdoc.String() + fmt.Sprintf("get%v () { return this[`%v`] }", core.ToUpper(field.Name), field.Name)

	setterjsdoc, setterFunc := jsClassSetterFunction(jsFieldType, field)

	setterCallInConstructor := fmt.Sprintf("if (d[`%v`] !== undefined) { \r\n this.set%v (d[`%v`]) \r\n}", field.Name, core.ToUpper(field.Name), field.Name)

	if isTypeScript {
		setterFunc = setterjsdoc.String() +
			fmt.Sprintf(
				"set%v (value: %v) { this[`%v`] = value; return this; }",
				core.ToUpper(field.Name),
				tsFieldType,
				field.Name,
			)
	}

	return jsRenderedField{
		Name:                    field.Name,
		Type:                    field.Type,
		Output:                  output,
		SetterFunc:              setterFunc,
		SetterCallInConstructor: setterCallInConstructor,
		GetterFunc:              getterFunc,
	}
}

func jsRenderFieldsShallow(fields []*core.Module3Field, parentChain string, ctx core.MicroGenContext) []jsRenderedField {

	output := []jsRenderedField{}

	for _, field := range fields {
		if field == nil {
			continue
		}
		item := jsRenderField(field, parentChain, ctx)
		output = append(output, item)
	}

	return output
}

func jsRenderDataClasses(fields []*core.Module3Field, className string, treeLocation string, isFirst bool, ctx core.MicroGenContext) []jsRenderedDataClass {
	if len(fields) == 0 {
		return []jsRenderedDataClass{}
	}

	var content []jsRenderedDataClass

	jsdoc := NewJsDoc("  ").Add(fmt.Sprintf("@decription The base class definition for %v", core.ToLower(className)))
	signature := fmt.Sprintf("export class %v implements %vType", core.ToUpper(className), core.ToUpper(className))

	// When it's first one, we use class. For children, signature is a bit different since they go inside.
	if !isFirst {
		signature = fmt.Sprintf("static %v = class %v", className, className)
	}

	currentClass := jsRenderedDataClass{
		ClassName: core.ToUpper(className),
		Fields:    jsRenderFieldsShallow(fields, treeLocation, ctx),
		JsDoc:     jsdoc.String(),
		Signature: signature,
	}

	for _, field := range fields {
		if field == nil {
			continue
		}
		if field.Type == core.FIELD_TYPE_OBJECT || field.Type == core.FIELD_TYPE_ARRAY {
			childName := core.ToUpper(field.Name)
			currentClass.SubClasses = append(currentClass.SubClasses, jsRenderDataClasses(field.Fields, core.ToUpper(childName), treeLocation+"."+core.ToUpper(childName), false, ctx)...)
		}
	}

	content = append(content, currentClass)

	return content
}

type JsCommonObjectContext struct {

	// The class name which will be used to generate nested classes,
	// in case of array or object
	RootClassName string
}

// This function can be used in different locations of the code generation,
// creates dtos, entities for actions or other specs.
func JsCommonObjectGenerator(fields []*core.Module3Field, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	res := &core.CodeChunkCompiled{}

	tsTypes, tsTypeError := TsCommonObjectGenerator(fields, ctx, TsCommonObjectContext{
		RootTypeName: jsctx.RootClassName,
	})

	if tsTypeError != nil {
		return nil, tsTypeError
	}

	renderedClasses := jsRenderDataClasses(fields, jsctx.RootClassName, jsctx.RootClassName, true, ctx)

	if len(renderedClasses) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
			Name:  TOKEN_ROOT_CLASS,
			Value: renderedClasses[0].ClassName,
		})
	}

	if nestJsDecorator := strings.Contains(ctx.Tags, GEN_NEST_JS_COMPATIBILITY); nestJsDecorator && len(renderedClasses) > 0 {
		// If nest.js decorator is needed, what we are gonna do is to add the static Nest function
		// to the object. The important thing is we only add static Nest to first class, not children

		nestjsStaticDecorator, err := JsNestJsStaticDecorator(NestJsStaticDecoratorContext{
			ClassInstance:               renderedClasses[0].ClassName,
			NestJsStaticFunctionUseCase: RequestBody,
		}, ctx)

		if err != nil {
			return nil, err
		}

		// Make sure to add dependencies to render tree
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, nestjsStaticDecorator.CodeChunkDependenies...)

		// Add the static function to the class bottom
		renderedClasses[0].ClassStaticFunctions = append(
			renderedClasses[0].ClassStaticFunctions,
			string(nestjsStaticDecorator.ActualScript),
		)

	}

	const tmpl = `
{{ .tsInterface }}
{{ define "printClass" }}
{{ .JsDoc }}
{{ .Signature  }} {

	 
	constructor(data) {
		// This probably doesn't cover the nested objects
		const d = data as Partial<{{ .ClassName }}>;

		{{ range .Fields }}
			{{ .SetterCallInConstructor }}
		{{ end }}
	}

	{{ range .Fields }} 
		{{ .Output }}
		{{ .GetterFunc }}
		{{ .SetterFunc }} 
	{{ end }}

	{{ range .SubClasses }}
		{{ template "printClass" . }}
	{{ end }}

	{{ range .ClassStaticFunctions }}
		{{ . }}
	{{ end }}
}
{{ end }}


{{ range .renderedClasses }}
	{{ template "printClass" . }}
{{ end }}

{{ if .abstractFactoryClass }}
	{{ .abstractFactoryClass }}
{{ end }}

{{ define "matches" }}
  {{ range .Matches }}
    get {{$.Name}}As{{ .PublicName }}(): {{ .PublicName }} | null {
      return this.{{$.Name}} as any;
    }
  {{ end }}
{{ end }}
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))
	nestJsDecorator := strings.Contains(ctx.Tags, GEN_NEST_JS_COMPATIBILITY)

	// Abstract class factory is only useful then
	abstractFactoryClass := ""
	if isTypeScript {
		abstractFactoryClass = fmt.Sprintf(`
export abstract class %vFactory {
	abstract create(data: unknown): %v;
}
		`, core.ToUpper(jsctx.RootClassName), core.ToUpper(jsctx.RootClassName))
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"shouldExport":         true,
		"nestjsDecorator":      nestJsDecorator,
		"isTypeScript":         isTypeScript,
		"tsInterface":          string(tsTypes.ActualScript),
		"abstractFactoryClass": abstractFactoryClass,
		"renderedClasses":      renderedClasses,
		"fields":               fields,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()

	if isTypeScript {
		res.ActualScript = []byte(strings.ReplaceAll(string(res.ActualScript), "constructor(data)", "constructor(data: unknown)"))
	}

	return res, nil
}
