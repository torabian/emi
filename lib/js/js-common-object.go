// Renders the common object, such as entities, dtos.

package js

import (
	"bytes"
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
	Name       string
	Type       string
	Children   []jsRenderedField
	Output     string
	GetterFunc string
	SetterFunc string
}

// Some field types such as array and object,
// need to have the correct generated class to be assigned to them.
func jsFieldTypeOnNestedClasses(field *core.Module3Field, parentChain string) string {

	if field.Type == core.FIELD_TYPE_ARRAY || field.Type == core.FIELD_TYPE_OBJECT {
		return core.ToUpper(parentChain) + "." + core.ToUpper(field.Name)
	}

	return TsComputedField(field, false)
}

func tsFieldTypeOnNestedClasses(field *core.Module3Field, parentChain string) string {

	if field.Type == core.FIELD_TYPE_ARRAY || field.Type == core.FIELD_TYPE_OBJECT {
		return fmt.Sprintf("InstanceType<typeof %v>", core.ToUpper(parentChain)+"."+core.ToUpper(field.Name))
	}

	return TsComputedField(field, false)
}

func jsRenderField(field *core.Module3Field, parentChain string, ctx core.MicroGenContext) jsRenderedField {

	jsFieldType := jsFieldTypeOnNestedClasses(field, parentChain)
	tsFieldType := tsFieldTypeOnNestedClasses(field, parentChain)

	jsdoc := NewJsDoc("  ")
	jsdoc.Add(fmt.Sprintf("@type {%v}", jsFieldType))
	jsdoc.Add(fmt.Sprintf("@description %v", field.Description))

	// Javascript field definition
	output := fmt.Sprintf("%v %v;", jsdoc.String(), field.PrivateName())
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	if isTypeScript {
		output = fmt.Sprintf("%v %v: %v;", jsdoc.String(), field.PrivateName(), tsFieldType)
	}

	getterjsdoc := NewJsDoc("  ")
	getterjsdoc.Add(fmt.Sprintf("@returns {%v}", jsFieldType))
	getterjsdoc.Add(fmt.Sprintf("@description %v", field.Description))
	getterFunc := getterjsdoc.String() + fmt.Sprintf("get%v () { return this[`%v`] }", core.ToUpper(field.Name), field.Name)

	setterjsdoc := NewJsDoc("  ")
	setterjsdoc.Add(fmt.Sprintf("@param {%v}", jsFieldType))
	setterjsdoc.Add(fmt.Sprintf("@description %v", field.Description))
	setterFunc := setterjsdoc.String() + fmt.Sprintf("set%v (value) { this[`%v`] = value; return this; }", core.ToUpper(field.Name), field.Name)

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
		Name:       field.Name,
		Type:       field.Type,
		Output:     output,
		SetterFunc: setterFunc,
		GetterFunc: getterFunc,
	}
}

func jsRenderFieldsShallow(fields []*core.Module3Field, parentChain string, ctx core.MicroGenContext) []jsRenderedField {

	output := []jsRenderedField{}

	for _, field := range fields {
		item := jsRenderField(field, parentChain, ctx)
		output = append(output, item)
	}

	return output
}

func jsRenderDataClasses(fields []*core.Module3Field, className string, treeLocation string, isFirst bool, ctx core.MicroGenContext) []jsRenderedDataClass {
	var content []jsRenderedDataClass

	jsdoc := NewJsDoc("  ").Add(fmt.Sprintf("@decription The base class definition for %v", core.ToLower(className)))
	signature := fmt.Sprintf("export class %v", core.ToUpper(className))

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

	// then descend into object/array fields
	for _, field := range fields {
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

var TOKEN_ROOT_CLASS = "root.class"

// This function can be used in different locations of the code generation,
// creates dtos, entities for actions or other specs.
func JsCommonObjectGenerator(fields []*core.Module3Field, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

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
{{ define "printClass" }}
{{ .JsDoc }}
{{ .Signature  }} {

	constructor(data) {
		// This probably doesn't cover the nested objects
		Object.assign(this, data)
	}

	{{ range .Fields }} 
		{{ .Output }}
		{{ .GetterFunc }}
		{{ .SetterFunc }} 
	{{ end }}

	{{ range .SubClasses }}
		{{ template "printClass" . }}
	{{ end }}

	/** a placeholder for WebRequestX auto patching the json content to the object **/
	static __jsonParsable;

	{{ range .ClassStaticFunctions }}
		{{ . }}
	{{ end }}
}
{{ end }}


{{ range .renderedClasses }}
	{{ template "printClass" . }}
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

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"shouldExport":    true,
		"nestjsDecorator": nestJsDecorator,
		"renderedClasses": renderedClasses,
		"fields":          fields,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()

	return res, nil
}

func TsComputedField(field *core.Module3Field, isWorkspace bool) string {
	switch field.Type {
	case "string", "text":
		return "string"
	case "one":
		return field.Target
	case "daterange":
		return "any"
	case "enum":
		items := []string{}
		for _, item := range field.OfType {
			items = append(items, "\""+item.Key+"\"")
		}
		return strings.Join(items, " | ")
	case "json":
		return TsCalcJsonField(field)
	case "many2many":
		return field.Target + "[]"
	case "int64?", "int32?", "int?", "float64?", "float32?":
		return "number"
	case "bool?":
		return "boolean"
	case "array":
		return field.PublicName() + "[]"
	case "arrayP":
		return TsPrimitve(field.Primitive) + "[]"
	case "html":
		return "string"
	case "int64", "int32", "int":
		return "number"
	case "float64", "float32", "float":
		return "number"
	case "bool":
		return "boolean"
	case "Timestamp", "datenano":
		return "string"
	case "date":
		return "Date"
	case "double":
		return "number"
	case "object", "embed":
		return field.PublicName()
	case "money?":
		return "{amount: number, currency: string, formatted?: string}"
	default:
		return "string"
		// return field.Type
	}
}

func TsPrimitve(primitive string) string {
	switch primitive {
	case "string", "text":
		return "string"
	case "string?", "text?":
		return "string"
	case "int64", "int32", "int", "float64", "float32":
		return "number"
	case "int64?", "int32?", "int?", "float64?", "float32?":
		return "number"
	case "bool":
		return "boolean"
	case "bool?":
		return "boolean"
	default:
		return "unknown"
	}
}

func TsCalcJsonField(field *core.Module3Field) string {
	t := []string{}

	if len(field.Matches) > 0 {

		for _, match := range field.Matches {
			if match.Dto != nil {
				t = append(t, match.PublicName())
			}
		}

	} else {
		t = append(t, "any")
	}

	return strings.Join(t, "|")
}
