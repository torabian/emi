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
		return fmt.Sprintf("InstanceType<typeof %v> | null", prefix)
	case core.FieldTypeArray:
		return fmt.Sprintf("InstanceType<typeof %v>[]", prefix)
	default:
		return TsComputedField(field, false)
	}
}
