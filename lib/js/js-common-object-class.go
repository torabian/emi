// Renders common JS/TS objects such as entities and DTOs.

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

	// For typescript, it has the Partial types.
	DataSourceStatement string
}

type jsRenderedField struct {
	Name                    string
	Type                    string
	Children                []jsRenderedField
	PrivateField            string
	GetterFunc              string
	SetterCallInConstructor string
	StaticFieldItem         string
	SetterFunc              string
}

func getNullableDefaultValue(field *core.EmiField) string {
	if field.Default == nil {
		return "undefined"
	}
	switch v := field.Default.(type) {
	case string:
		return fmt.Sprintf("%q", v)
	case int, int64, float64, bool:
		return fmt.Sprintf("%v", v)
	default:
		b, _ := json.Marshal(v)
		return "'" + string(b) + "'"
	}
}

func jsGetSafeFieldValue(field *core.EmiField) string {
	if field == nil {
		return "null"
	}
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
	case "json", "object", "embed", "computed", "any":
		return "null"
	case "string", "text":
		return `""`
	case "float", "float32", "float64", "float?", "float32?", "float64?":
		return "0.0"
	case "int", "int32", "int64", "int?", "int32?", "int64?":
		return "0"
	case "money?":
		return `{"amount":0,"currency":"USD"}`
	default:
		return "null"
	}
}

type jsFieldVariable struct {
	Name                 string
	NullableDefaultValue string
	SafeDefaultValue     string
	Type                 string
	ConstructorClass     string
	ComputedType         string
	IsNullable           bool
	IsNumeric            bool
	JsDoc                string
	Modifier             string
}

func (x jsFieldVariable) Upper() string {
	return core.ToUpper(x.Name)
}

func (x jsFieldVariable) CreateSetterFunction(ctx core.MicroGenContext) string {

	claims := []core.JsFnArgument{
		{
			Key: "arg.value",
			Ts:  "value: " + x.ComputedType,
			Js:  "value",
		},
	}
	claimsRendered := core.ClaimRender(claims, ctx)

	var setterTemplate = template.Must(template.New("setter").Parse(`
{{.ctx.JsDoc}}
set {{ .ctx.Name }} (|@arg.value|) {
	{{ if or (eq .ctx.Type "string") }}
	 	const correctType = typeof value === 'string';
		this.#{{.ctx.Name}} = correctType ? value : ('' + value);
	{{ end }}

	{{ if or (eq .ctx.Type "string?") }}
	 	const correctType = typeof value === 'string' || value === undefined || value === null
		this.#{{.ctx.Name}} = correctType ? value : ('' + value);
	{{ end }}

	{{ if and (eq .ctx.IsNumeric true) (eq .ctx.IsNullable true) }}
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)

		if (!Number.isNaN(parsedValue)) {
			this.#{{.ctx.Name}} = parsedValue;
		}
	{{ end }}
	
	{{ if and (eq .ctx.IsNumeric true) (eq .ctx.IsNullable false) }}
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)

		if (!Number.isNaN(parsedValue)) {
			this.#{{.ctx.Name}} = parsedValue;
		}
	{{ end }}
	
	{{ if and (eq .ctx.Type "bool")}}
	 	const correctType = value === true || value === false
		this.#{{.ctx.Name}} = correctType ? value : Boolean(value);
	{{ end }}
	
	{{ if and (eq .ctx.Type "bool?")}}
	 	const correctType = value === true || value === false || value === undefined || value === null
		this.#{{.ctx.Name}} = correctType ? value : Boolean(value);
	{{ end }}
	
	{{ if and (eq .ctx.Type "array")}}
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof {{.ctx.ConstructorClass}}) {
			this.#{{.ctx.Name}} = value
		} else {
			this.#{{.ctx.Name}} = value.map(item => new {{.ctx.ConstructorClass}}(item))
		}
 	{{ end }}

	{{ if and (eq .ctx.Type "object")}}
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof {{.ctx.ConstructorClass}}) {
			this.#{{.ctx.Name}} = value
		} else {
			this.#{{.ctx.Name}} = new {{.ctx.ConstructorClass}}(value)
		}
 	{{ end }}
}

set{{ .ctx.Upper }} (|@arg.value|) {
	this.{{.ctx.Name}} = value

	return this
}
`))
	setterjsdoc := NewJsDoc("  ")
	setterjsdoc.Add(fmt.Sprintf("@param {%v}", x.ComputedType))

	data := map[string]any{
		"ctx": x,
	}

	var buf bytes.Buffer
	if err := setterTemplate.Execute(&buf, data); err != nil {
		panic(err)
	}

	templateResult := buf.String()
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	return templateResult
}

// generates a field variable section, such as public #fieldName: string = null;
func (x jsFieldVariable) Compile(isTypeScript bool) string {
	sequence := []string{}

	if x.JsDoc != "" {
		sequence = append(sequence, x.JsDoc)
	}

	if x.Modifier != "" && x.Modifier != "private" {
		sequence = append(sequence, x.Modifier)
	}

	varName := x.Name
	if x.Modifier == "private" {
		varName = "#" + varName
	}

	sequence = append(sequence, varName)

	if x.IsNullable {
		sequence = append(sequence, "?")
	}

	if isTypeScript {
		sequence = append(sequence, ": "+x.ComputedType)

		if x.IsNullable {
			sequence = append(sequence, " | null")
		}
	}

	sequence = append(sequence, " = ")

	if x.IsNullable {
		if x.NullableDefaultValue != "" {
			sequence = append(sequence, x.NullableDefaultValue)
		} else {
			sequence = append(sequence, "undefined")
		}
	} else if !x.IsNullable {
		if x.SafeDefaultValue != "" {
			sequence = append(sequence, x.SafeDefaultValue)
		} else {
			sequence = append(sequence, "undefined")
		}
	}

	return strings.Join(sequence, " ")
}

func jsRenderField(field *core.EmiField, parentChain string, fieldDepth string, ctx core.MicroGenContext) jsRenderedField {
	jsFieldType := jsFieldTypeOnNestedClasses(field, parentChain)
	tsFieldType := tsFieldTypeOnNestedClasses(field, parentChain)
	isFieldNullable := IsNullable(string(field.Type))
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	jsdoc := NewJsDoc("  ")
	jsdoc.Add(field.Description)
	jsdoc.Add(fmt.Sprintf("@type {%v}", jsFieldType))

	privateFieldToken := jsFieldVariable{
		Modifier:             "private",
		Name:                 field.PrivateName(),
		JsDoc:                jsdoc.String(),
		NullableDefaultValue: getNullableDefaultValue(field),
		SafeDefaultValue:     jsGetSafeFieldValue(field),
		IsNullable:           isFieldNullable,
		Type:                 string(field.Type),
		ComputedType:         tsFieldType,
		IsNumeric:            IsNumericDataType(string(field.Type)),
		ConstructorClass:     core.ToUpper(parentChain) + "." + core.ToUpper(field.Name),
	}

	privateField := privateFieldToken.Compile(isTypeScript)

	getterjsdoc := NewJsDoc("  ")
	getterjsdoc.Add(field.Description)
	getterjsdoc.Add(fmt.Sprintf("@returns {%v}", jsFieldType))
	getterFunc := getterjsdoc.String() + fmt.Sprintf("get %v () { return this.#%v }", field.Name, field.Name)

	setterCallInConstructor := fmt.Sprintf("if (d.%v !== undefined) { this.%v = d.%v }", field.Name, field.Name, field.Name)

	staticVariables := []string{}

	if field.Type == core.FieldTypeArrayP || field.Type == core.FieldTypeArray || field.Type == core.FieldTypeMany2Many || field.Type == core.FieldTypeEmbed || field.Type == core.FieldTypeObject || field.Type == core.FieldTypeOne {
		staticVariables = append(
			staticVariables,
			fmt.Sprintf("%v$: '%v',", field.Name, field.Name),
		)

		withArrayIndex := ""

		if field.Type == core.FieldTypeArrayP || field.Type == core.FieldTypeArray || field.Type == core.FieldTypeMany2Many {
			withArrayIndex = "[:i]"
		}

		newDepth := fieldDepth + "." + field.Name
		if fieldDepth == "" {
			newDepth = field.Name
		}
		staticVariables = append(
			staticVariables,
			fmt.Sprintf(`get %v() {
				return withPrefix(
					"%v%v",
					%v.Fields
				);
			},`, field.Name, newDepth, withArrayIndex, core.ToUpper(parentChain)+"."+core.ToUpper(field.Name)),
		)
	} else {
		// For default types the string name is enough.
		staticVariables = append(
			staticVariables,
			fmt.Sprintf("%v: '%v',", field.Name, field.Name),
		)
	}

	staticFieldItem := strings.Join(staticVariables, "\r\n")
	setterFunc := privateFieldToken.CreateSetterFunction(ctx)

	return jsRenderedField{
		Name:                    field.Name,
		Type:                    string(field.Type),
		PrivateField:            privateField,
		SetterFunc:              setterFunc,
		SetterCallInConstructor: setterCallInConstructor,
		StaticFieldItem:         staticFieldItem,
		GetterFunc:              getterFunc,
	}
}

func jsRenderFieldsShallow(fields []*core.EmiField, parentChain string, fieldDepth string, ctx core.MicroGenContext) []jsRenderedField {
	out := make([]jsRenderedField, 0, len(fields))
	for _, f := range fields {
		if f != nil {
			out = append(out, jsRenderField(f, parentChain, fieldDepth, ctx))
		}
	}
	return out
}

func jsRenderDataClasses(fields []*core.EmiField, className, treeLocation string, fieldDepth string, isFirst bool, ctx core.MicroGenContext) []jsRenderedDataClass {
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
		Fields:              jsRenderFieldsShallow(fields, treeLocation, fieldDepth, ctx),
		JsDoc:               jsdoc.String(),
		Signature:           signature,
		DataSourceStatement: "const d = data;",
	}

	if isTypeScript {
		currentClass.DataSourceStatement = fmt.Sprintf("const d = data as Partial<%v>;", currentClass.ClassName)
	}

	for _, f := range fields {
		if f != nil && (f.Type == core.FieldTypeObject || f.Type == core.FieldTypeArray) {
			childName := core.ToUpper(f.Name)
			newDepth := fieldDepth + "." + f.Name
			if fieldDepth == "" {
				newDepth = f.Name
			}
			currentClass.SubClasses = append(currentClass.SubClasses, jsRenderDataClasses(f.Fields, childName, treeLocation+"."+childName, newDepth, false, ctx)...)
		}
	}

	return []jsRenderedDataClass{currentClass}
}

// Generates a class with getters and setters.
func JsCommonObjectClassGenerator(fields []*core.EmiField, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error) {
	isTS := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	res := &core.CodeChunkCompiled{}

	renderedClasses := jsRenderDataClasses(fields, jsctx.RootClassName, jsctx.RootClassName, "", true, ctx)
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
		} else if (isPlausibleObject(data)) {
			this.applyFromObject(data);
		} else {
			throw new Error("Instance is not implemented.");
		}
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
	if field == nil {
		return ""
	}
	if field.Type == core.FieldTypeArray || field.Type == core.FieldTypeObject {
		return core.ToUpper(parentChain) + "." + core.ToUpper(field.Name)
	}
	return TsComputedField(field, false)
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
	default:
		return TsComputedField(field, false)
	}
}
