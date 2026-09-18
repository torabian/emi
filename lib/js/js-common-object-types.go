// Renders the common object, such as entities, dtos, but as *types* instead of classes.

package js

import (
	"bytes"
	"fmt"
	"strings"
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

// tsTargetTypeName resolves a `one`/`collection` field's `target` to its
// companion plain-type name (e.g. "PassportEntity" -> "PassportEntityType"),
// rather than the class name TsComputedField gives the class body's own
// getter/setter annotations. The exported *Type declaration must never
// mention a class - the entire point of also generating a plain type is so
// a caller who doesn't want the class can consume the type instead - so a
// one/collection field's own nested reference has to resolve to *that*
// target's plain type too. entityTargetToCodeChunk imports it (aliased to
// this same `<target>Type` name) alongside the class it already imports, so
// the name this returns is always in scope.
//
// A self-referencing target (field.Target == "self@...", a dto pointing at
// its own shape - e.g. a tree node's "children") resolves within the
// current file instead of needing any import at all - but NOT the same way
// TsComputedField's class-side self-reference does: getSelfReferencingField
// resolves via parentChain's own root segment, and on this (type-rendering)
// side that root segment is *already* Type-suffixed - tsRenderTypes' very
// first call seeds treeLocation as `RootTypeName + "Type"`, specifically so
// nested object/array subtypes come out namespaced under it correctly (e.g.
// `RootType.ChildType`). Appending "Type" again here for a self-reference
// would double it ("RootTypeType") - a real bug this once was: a
// self-referencing `collection`/`one` field (`target: self@`, e.g. a tree
// node's own "children") rendered as `RootTypeType[]` instead of
// `RootType[]`. A genuine external target has no such pre-suffixed value to
// worry about, so it still needs the plain "+ Type" append.
func tsTargetTypeName(field *core.EmiField, parentChain string) string {
	if isSelf, value := getSelfReferencingField(field, parentChain); isSelf {
		return value
	}
	return field.Target + "Type"
}

// PlainOf<T> itself lives in the shared SDK (sdk/common/fetchx.ts, alongside
// PartialDeep - see TsCommonObjectGenerator, which imports it exactly the
// same way, rather than emitting it inline here per file). A conditional
// type, not a plain "ComplexNameType" name, on purpose: unlike a dto/entity
// target (tsTargetTypeName), Emi has no way to know a hand-authored complex
// type's own plain shape - there's no "ComplexNameType" companion it could
// generate or assume exists, since a complex type isn't something Emi
// generates at all. What Emi *can* rely on is the same duck-typed contract
// the runtime side already uses everywhere (see #toPlainJSON,
// js-common-object-class.go): if the complex type declares its own
// toJSON(), that method's return type is the true plain shape it serializes
// to - PlainOf<T> reads that off at the type level using nothing but
// TypeScript's built-in ReturnType<>. A complex type with no toJSON() falls
// back to T itself unchanged (today's behavior, not a regression) - there
// is no plain type to reach for otherwise, so the class reference is still
// the best available answer for that field.

// tsPlainOfComplex wraps a complex field's class name (already stripped of
// the leading "+" some complex fields carry - see CollectComplexClasses'
// same ReplaceAll) in PlainOf<...> for the exported *Type declaration.
func tsPlainOfComplex(complex string) string {
	return fmt.Sprintf("PlainOf<%s>", strings.ReplaceAll(complex, "+", ""))
}

func tsFieldType(field *core.EmiField, parentChain string) string {

	value := ""
	if field != nil {

		fieldName := core.ToUpper(field.Name) + "Type"

		// Checked ahead of field.Type, mirroring TsComputedField's own
		// precedence (a Complex field's class-side type is field.Complex
		// regardless of what Type says) - see tsPlainOfComplex's own doc
		// comment for why the exported *Type wraps it in PlainOf<...>
		// instead of using the bare complex name the class body gets.
		if field.Complex != "" {
			return tsPlainOfComplex(field.Complex)
		}

		switch field.Type {
		case core.FieldTypeArray, core.FieldTypeArrayNullable:
			if len(field.Fields) > 0 {
				return fmt.Sprintf("%v[]", core.ToUpper(parentChain)+"."+fieldName)
			}
			value = "any[]"
		case core.FieldTypeObject, core.FieldTypeObjectNullable:
			value = core.ToUpper(parentChain) + "." + fieldName
		case core.FieldTypeOne, core.FieldTypeOneNullable:
			value = tsTargetTypeName(field, parentChain)
		case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
			value = tsTargetTypeName(field, parentChain) + "[]"
		default:
			value = TsComputedField(field, false, parentChain)
		}
	}

	return strings.ReplaceAll(value, "+", "")
}

func tsRenderField(field *core.EmiField, parentChain string) tsRenderedField {
	tsFieldTypeStr := tsFieldType(field, parentChain)

	jsdoc := NewJsDoc("  ")
	jsdoc.Add(field.Description)
	jsdoc.Add(fmt.Sprintf("@type {%v}", tsFieldTypeStr))

	// Add ? only for nullable fields.
	mark := ""
	if IsNullable(string(field.Type)) {
		mark = "?"
	}
	output := fmt.Sprintf("%v %v %v: %v;", jsdoc.String(), field.PrivateName(), mark, tsFieldTypeStr)

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

	typeNameUpper := core.ToUpper(typeName)

	jsdoc := NewJsDoc("  ").Add(fmt.Sprintf("The base type definition for %v", core.ToLower(typeName)))
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
		if field.Type == core.FieldTypeObject || field.Type == core.FieldTypeArray || field.Type == core.FieldTypeObjectNullable || field.Type == core.FieldTypeArrayNullable {
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
	var mainInterface *tsRenderedType
	var renderedSubTypes []tsRenderedType = []tsRenderedType{}

	if len(renderedTypes) > 0 {
		mainInterface = &renderedTypes[0]
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
			Name:  TOKEN_ROOT_CLASS,
			Value: renderedTypes[0].TypeName,
		})

		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
			Name:  TOKEN_OBJ_TYPE,
			Value: renderedTypes[0].TypeName,
		})

		// The real, `Type`-suffixed identifier of the standalone `export type`
		// declaration below - see TOKEN_TYPEDEF_NAME's own doc comment
		// (js-tokens.go) for why this needs to be a separate token from
		// TOKEN_OBJ_TYPE above (which intentionally stays the bare class name).
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
			Name:  TOKEN_TYPEDEF_NAME,
			Value: core.ToUpper(tsctx.RootTypeName) + "Type",
		})

		renderedSubTypes = renderedTypes[0].SubTypes
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

	if mainInterface != nil {

		if len(CollectComplexClasses(fields)) > 0 {
			res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
				Objects:  []string{"type PlainOf"},
				Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "fetchx"),
			})
		}

		var buf bytes.Buffer
		if err := t.Execute(&buf, core.H{
			"renderedTypes": renderedSubTypes,
			"mainInterface": mainInterface,
			"namespaceName": renderedTypes[0].TypeName + "Type",
			"fields":        fields,
		}); err != nil {
			return nil, err
		}

		res.ActualScript = buf.Bytes()
	}
	return res, nil
}
