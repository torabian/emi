// Renders common JS/TS objects such as entities and DTOs.

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

	// For typescript, it has the Partial types.
	DataSourceStatement string
}

func jsRenderDataClasses(fields []*core.EmiField, className, treeLocation string, fieldDepth string, isFirst bool, ctx core.MicroGenContext, jsctx JsCommonObjectContext) []jsRenderedDataClass {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	if len(fields) == 0 {
		return nil
	}

	jsdoc := NewJsDoc("  ").Add(fmt.Sprintf("The base class definition for %v", core.ToLower(className)))
	signature := fmt.Sprintf("export class %v", core.ToUpper(className))
	if !isFirst {
		signature = fmt.Sprintf("static %v = class %v", className, className)
	}

	currentClass := jsRenderedDataClass{
		ClassName:           core.ToUpper(className),
		Fields:              jsRenderFieldsShallow(fields, treeLocation, fieldDepth, ctx, jsctx),
		JsDoc:               jsdoc.String(),
		Signature:           signature,
		DataSourceStatement: "const d = data;",
	}

	if isTypeScript {
		currentClass.DataSourceStatement = fmt.Sprintf("const d = data as Partial<%v>;", currentClass.ClassName)
	}

	for _, f := range fields {
		if f != nil && (f.Type == core.FieldTypeObject || f.Type == core.FieldTypeObjectNullable || f.Type == core.FieldTypeArray || f.Type == core.FieldTypeArrayNullable) {
			childName := core.ToUpper(f.Name)
			newDepth := fieldDepth + "." + f.Name
			if fieldDepth == "" {
				newDepth = f.Name
			}
			currentClass.SubClasses = append(currentClass.SubClasses, jsRenderDataClasses(f.Fields, childName, treeLocation+"."+childName, newDepth, false, ctx, jsctx)...)
		}
	}

	return []jsRenderedDataClass{currentClass}
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

func hasClassesAsChildren(fields []*core.EmiField) bool {
	var result = false

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field.Type == core.FieldTypeArray || field.Type == core.FieldTypeObject || field.Type == core.FieldTypeArrayNullable || field.Type == core.FieldTypeObjectNullable || field.Type == core.FieldTypeOne || field.Type == core.FieldTypeMany2Many {
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

func findComplexLocation(complexName string, jsctx JsCommonObjectContext) string {

	for _, item := range jsctx.RecognizedComplexes {
		if item.Symbol == complexName {
			return item.ImportLocation
		}
	}

	// If no location, then basically skip it.
	return ""
}

// Generates a class with getters and setters.
func JsCommonObjectClassGenerator(fields []*core.EmiField, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error) {

	hasChildrenWithStaticFields := hasClassesAsChildren(fields)
	isTS := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	res := &core.CodeChunkCompiled{
		CodeChunkDependenies: []core.CodeChunkDependency{},
	}

	if hasChildrenWithStaticFields {
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
			Objects:  []string{"withPrefix"},
			Location: INTERNAL_SDK_JS_LOCATION + "/withPrefix",
		})
	}

	usedComplexes := CollectComplexClasses(fields)
	for _, item := range usedComplexes {

		location := findComplexLocation(item, jsctx)
		if location == "" {
			continue
		}

		res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
			Objects:  []string{item},
			Location: location,
		})
	}

	renderedClasses := jsRenderDataClasses(fields, jsctx.RootClassName, jsctx.RootClassName, "", true, ctx, jsctx)
	if len(renderedClasses) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: renderedClasses[0].ClassName})
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_OBJ_CLASS, Value: renderedClasses[0].ClassName})
	}

	var abstractFactoryClass string
	if isTS {
		abstractFactoryClass = fmt.Sprintf(`
export abstract class %vFactory {
	abstract create(data: unknown): %v;
}`, core.ToUpper(jsctx.RootClassName), core.ToUpper(jsctx.RootClassName))
	}

	const tmpl = `
{{ define "printClass" }}
{{ .JsDoc }}
{{ .Signature  }} {
	{{ range .Fields }}
		{{ .PrivateField }}
		{{ .GetterFunc }}
		{{ .SetterFunc }}
	{{ end }}

	{{ range .SubClasses }}
		{{ template "printClass" . }}
	{{ end }}

	{{ range .ClassStaticFunctions }}
		{{ . }}
	{{ end }}

	constructor(data) {
		if (data === null || data === undefined) {
			return;
		}
		if (typeof data === "string") {
			this.applyFromObject(JSON.parse(data));
		} else if (this.#isJsonAppliable(data)) {
			this.applyFromObject(data);
		} else {
			throw new Error("Instance cannot be created on an unknown value, check the content being passed. got: "  + typeof data);
		}
	}

	#isJsonAppliable(obj) {
		const isBuffer =
			typeof globalThis.Buffer !== "undefined" &&
			typeof globalThis.Buffer.isBuffer === "function" &&
			globalThis.Buffer.isBuffer(obj);

		const isBlob =
			typeof globalThis.Blob !== "undefined" && obj instanceof globalThis.Blob;

		return (
			obj &&
			typeof obj === "object" &&
			!Array.isArray(obj) &&
			!isBuffer &&
			!(obj instanceof ArrayBuffer) &&
			!isBlob
		);
	}


	/**
	* casts the fields of a javascript object into the class properties one by one
	**/
	applyFromObject(data = {}) {
		{{ .DataSourceStatement }}
		{{ range .Fields }}
			{{ .SetterCallInConstructor }}
		{{ end }}
	}

	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
			{{ range .Fields }}
				{{ .Name }}: this.#{{ .Name }},
			{{ end }}
		};
  	}

	toString() {
		return JSON.stringify(this);
	}

	static get Fields() {
      return {
		{{ range .Fields }}
			{{ .StaticFieldItem }}
		{{ end }}
	  }
	}
}
{{ end }}

{{ range .renderedClasses }}
	{{ template "printClass" . }}
{{ end }}

{{ if .abstractFactoryClass }}
	{{ .abstractFactoryClass }}
{{ end }}
 
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"renderedClasses":      renderedClasses,
		"abstractFactoryClass": abstractFactoryClass,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	if isTS {
		res.ActualScript = []byte(strings.ReplaceAll(string(res.ActualScript), "constructor(data = {})", "constructor(data: unknown = {})"))
		res.SuggestedExtension = ".ts"
	} else {
		res.SuggestedExtension = ".js"
	}
	res.SuggestedFileName = jsctx.RootClassName

	return res, nil
}

func jsFieldTypeOnNestedClasses(field *core.EmiField, parentChain string) string {
	value := ""
	if field == nil {
		value = ""
	} else if field.Type == core.FieldTypeArray || field.Type == core.FieldTypeObject || field.Type == core.FieldTypeArrayNullable || field.Type == core.FieldTypeObjectNullable {
		value = core.ToUpper(parentChain) + "." + core.ToUpper(field.Name)
	} else {
		value = TsComputedField(field, false)
	}

	// For the complex fields
	return strings.ReplaceAll(value, "+", "")
}

func tsFieldTypeOnNestedClasses(field *core.EmiField, parentChain string) string {
	if field == nil {
		return ""
	}
	prefix := core.ToUpper(parentChain) + "." + core.ToUpper(field.Name)
	switch field.Type {
	case core.FieldTypeObject:
		return fmt.Sprintf("InstanceType<typeof %v>", prefix)
	case core.FieldTypeArray:
		return fmt.Sprintf("InstanceType<typeof %v>[]", prefix)
	case core.FieldTypeObjectNullable:
		return fmt.Sprintf("InstanceType<typeof %v> | null | undefined", prefix)
	case core.FieldTypeArrayNullable:
		return fmt.Sprintf("InstanceType<typeof %v>[] | null | undefined", prefix)
	default:
		return TsComputedField(field, false)
	}
}
