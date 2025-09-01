// Renders the common object, such as entities, dtos, but as *types* instead of classes.

package js

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type tsRenderedType struct {
	TypeName  string
	Fields    []tsRenderedField
	SubTypes  []tsRenderedType
	JsDoc     string
	Signature string
}

type tsRenderedField struct {
	Name   string
	Type   string
	Output string
}

func tsFieldType(field *core.EmiField, parentChain string) string {
	if field == nil {
		return ""
	}

	fieldName := core.ToUpper(field.Name) + "Type"

	if field.Type == core.FieldTypeArray {
		if len(field.Fields) > 0 {
			return fmt.Sprintf("%v[]", core.ToUpper(parentChain)+"."+fieldName)
		}
		return "any[]"
	}
	if field.Type == core.FieldTypeObject {
		return core.ToUpper(parentChain) + "." + fieldName
	}
	return TsComputedField(field, false)
}

func tsRenderField(field *core.EmiField, parentChain string) tsRenderedField {
	tsFieldTypeStr := tsFieldType(field, parentChain)

	jsdoc := NewJsDoc("  ")
	jsdoc.Add(fmt.Sprintf("@type {%v}", tsFieldTypeStr))
	jsdoc.Add(fmt.Sprintf("@description %v", field.Description))

	output := fmt.Sprintf("%v %v?: %v;", jsdoc.String(), field.PrivateName(), tsFieldTypeStr)

	return tsRenderedField{
		Name:   field.Name,
		Type:   tsFieldTypeStr,
		Output: output,
	}
}

func tsRenderFields(fields []*core.EmiField, parentChain string) []tsRenderedField {
	output := []tsRenderedField{}
	for _, field := range fields {
		if field == nil {
			continue
		}
		output = append(output, tsRenderField(field, parentChain))
	}
	return output
}

func tsRenderTypes(fields []*core.EmiField, typeName string, treeLocation string, isFirst bool) []tsRenderedType {
	if len(fields) == 0 {
		return []tsRenderedType{}
	}

	typeNameUpper := core.ToUpper(typeName)

	jsdoc := NewJsDoc("  ").Add(fmt.Sprintf("@description The base type definition for %v", core.ToLower(typeName)))
	typeNameFirst := typeNameUpper
	if isFirst {
		typeNameFirst = typeNameUpper + "Type"
	}
	signature := fmt.Sprintf("export type %v = ", typeNameFirst)

	currentType := tsRenderedType{
		TypeName:  typeNameUpper,
		Fields:    tsRenderFields(fields, treeLocation),
		JsDoc:     jsdoc.String(),
		Signature: signature,
	}

	// recurse for object/array subtypes
	for _, field := range fields {
		if field == nil {
			continue
		}
		if field.Type == core.FieldTypeObject || field.Type == core.FieldTypeArray {
			childName := core.ToUpper(field.Name) + "Type"
			currentType.SubTypes = append(
				currentType.SubTypes,
				tsRenderTypes(field.Fields, (childName), treeLocation+"."+(childName), false)...,
			)
		}
	}

	return []tsRenderedType{currentType}
}

type TsCommonObjectContext struct {
	RootTypeName string
}

func TsCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, tsctx TsCommonObjectContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	renderedTypes := tsRenderTypes(fields, tsctx.RootTypeName, tsctx.RootTypeName+"Type", true)

	if len(renderedTypes) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
			Name:  TOKEN_ROOT_CLASS,
			Value: renderedTypes[0].TypeName,
		})
	}

	const tmpl = `

{{ define "printClassItSelf" }}
	{{ .JsDoc }}
	{{ .Signature }} {
		{{ range .Fields }}
			{{ .Output }}
		{{ end }}
	}
{{ end }}
{{ define "printType" }}
{{ template "printClassItSelf" .}}

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace {{ .TypeName }} {
	{{ range .SubTypes }}
		{{ template "printType" . }}
	{{ end }}
}
{{ end }}

{{ template "printClassItSelf" .mainInterface }}

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace {{ .namespaceName }} {
	{{ range .renderedTypes }}
		{{ template "printType" . }}
	{{ end }}
}
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"renderedTypes": renderedTypes[0].SubTypes,
		"mainInterface": renderedTypes[0],
		"namespaceName": renderedTypes[0].TypeName + "Type",
		"fields":        fields,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	return res, nil
}
