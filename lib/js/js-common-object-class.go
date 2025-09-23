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
	LateInitFields       []jsRenderedField
	Signature            string
	JsDoc                string
	SubClasses           []jsRenderedDataClass
	ClassTypePath        string
	ClassNamePath        string
	ClassStaticFunctions []string

	// For typescript, it has the Partial types.
	DataSourceStatement string
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

	lateInitFields := []jsRenderedField{}
	fieldsRendered := jsRenderFieldsShallow(fields, treeLocation, fieldDepth, ctx, jsctx)
	for _, field := range fieldsRendered {
		if field.LateInitStatement != "" {
			lateInitFields = append(lateInitFields, field)
		}
	}

	currentClass := jsRenderedDataClass{
		ClassName:           core.ToUpper(className),
		Fields:              fieldsRendered,
		ClassNamePath:       treeLocation,
		ClassTypePath:       treeAsType(treeLocation),
		LateInitFields:      lateInitFields,
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

// when developer says target: AnotherEntity or target: AnotherDto, we need
// to import that. Here we search through the definition tree and extract them all
func CollectTargets(fields []*core.EmiField) []string {
	var result []string

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field.Target != "" {
				result = append(result, field.Target)
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
			if field.Type == core.FieldTypeArray || field.Type == core.FieldTypeMany2ManyNullable || field.Type == core.FieldTypeObject || field.Type == core.FieldTypeArrayNullable || field.Type == core.FieldTypeObjectNullable || field.Type == core.FieldTypeOne || field.Type == core.FieldTypeMany2Many {
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
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{},
	}

	if hasChildrenWithStaticFields {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
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

		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Objects:  []string{item},
			Location: location,
		})
	}

	collectTargets := CollectTargets(fields)
	for _, item := range collectTargets {
		m := castDtoNameToCodeChunk(item)
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, m.CodeChunkDependensies...)
	}

	renderedClasses := jsRenderDataClasses(fields, jsctx.RootClassName, jsctx.RootClassName, "", true, ctx, jsctx)
	if len(renderedClasses) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: renderedClasses[0].ClassName})
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_OBJ_CLASS, Value: renderedClasses[0].ClassName})
	}

	var abstractFactoryClass string
	if isTypeScript {
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
			{{ if .LateInitFields }}
				this.#lateInitFields();
			{{ end }}
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
		const g = globalThis
		const isBuffer =
			typeof g.Buffer !== "undefined" &&
			typeof g.Buffer.isBuffer === "function" &&
			g.Buffer.isBuffer(obj);

		const isBlob =
			typeof g.Blob !== "undefined" && obj instanceof g.Blob;

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

		{{ if .LateInitFields }}
		this.#lateInitFields(data)
		{{ end }}
	}
		
	{{ if .LateInitFields }}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		{{ .DataSourceStatement }}
		{{ range .LateInitFields }}
			{{ .LateInitStatement }}	
		{{ end }}
	}

	{{ end }}

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

	/**
	* Creates an instance of {{ .ClassNamePath }}, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: {{ .ClassTypePath }}) {
		return new {{ .ClassNamePath }}(possibleDtoObject);
	}

	/**
	* Creates an instance of {{ .ClassNamePath }}, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<{{ .ClassTypePath }}>) {
		return new {{ .ClassNamePath }}(partialDtoObject);
	}

	copyWith(partial: PartialDeep<{{ .ClassTypePath }}>): InstanceType<typeof {{ .ClassNamePath }}> {
		return new {{ .ClassNamePath }} ({ ...this.toJSON(), ...partial });
	}

	clone(): InstanceType<typeof {{ .ClassNamePath }}> {
		return new {{ .ClassNamePath }}(this.toJSON());
	}

}
{{ end }}

{{ range .renderedClasses }}
	{{ template "printClass" . }}
{{ end }}

{{ if .abstractFactoryClass }}
	{{ .abstractFactoryClass }}
{{ end }}

type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"isTypeScript":         isTypeScript,
		"renderedClasses":      renderedClasses,
		"abstractFactoryClass": abstractFactoryClass,
	}); err != nil {
		return nil, err
	}

	result := buf.String()

	if isTypeScript {
		result = strings.ReplaceAll(string(result), "#isJsonAppliable(obj) {", "#isJsonAppliable(obj: unknown) {")
		result = strings.ReplaceAll(string(result), "const g = globalThis", "const g = globalThis as unknown as { Buffer: any; Blob: any };")
		// result = strings.ReplaceAll(string(result), "const g = globalThis", "const g = globalThis as any")
		res.ActualScript = []byte(strings.ReplaceAll(result, "constructor(data)", "constructor(data: unknown = undefined)"))
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
