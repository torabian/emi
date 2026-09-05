// Renders common JS/TS objects such as entities and DTOs.

package js

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/formgen"
)

type jsRenderedDataClass struct {
	ClassName            string
	Fields               []jsRenderedField
	LateInitFields       []jsRenderedField
	Signature            string
	JsDoc                string
	IsTypeScript         bool
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
	isTypeScript := ctx.HasTag(Typescript)

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
		IsTypeScript:        isTypeScript,
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

// Finds every complex class referenced by fields (e.g. "TString", "XDate")
// so the caller can import each one - see JsCommonObjectClassGenerator's own
// usedComplexes loop, which is the only caller.
//
// Bug fix: this used to only collect field.Complex when it contained a
// literal "+" (`strings.Contains(field.Complex, "+")`) - a convention no
// actual .emi.yml in either this repo or any consumer project's modules
// ever uses (a plain `complex: TString` field, the normal and only way
// every complex type is actually declared, has no "+" in it at all). So a
// complex field's generated class property was always correctly *typed* as
// e.g. `TString` (js-data-types.go's ComputedFieldType reads field.Complex
// directly, unaffected by this function), but the corresponding `import {
// TString } from "..."` was silently never added - every generated dto
// using a complex type failed to type-check (`Cannot find name 'TString'`)
// even though nothing about it looked wrong at a glance, since JS runtimes
// erase type annotations and never noticed the missing import either.
// strings.ReplaceAll below still strips a "+" if one somehow is present, so
// that hypothetical old convention keeps working exactly as before for
// anyone who was actually relying on it - this just stops requiring it.
func CollectComplexClasses(fields []*core.EmiField) []string {
	var result []string

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field == nil {
				continue
			}

			if field.Complex != "" {
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

var SELF_FIELD = "self@"

// EntityTargetRef is one `target:` reference CollectTargets found - the
// short name (`field.Target`, e.g. "WalletEntity") always used as the
// generated alias/type name everywhere else in this file, kept separate from
// JsProvider (e.g. "@fireback/wallet/sdk/WalletEntity", empty for a
// same-module reference) so entityTargetToCodeChunk never has to guess which
// one a single collapsed string was - see its own doc comment for the bug
// that produced (`import { WalletDto as @fireback/wallet/sdk/WalletEntity }`,
// the full import path used as the alias) before Target/JsProvider were
// tracked separately here.
type EntityTargetRef struct {
	Target     string
	JsProvider string
}

// when developer says target: AnotherEntity or target: AnotherDto, we need
// to import that. Here we search through the definition tree and extract them all
func CollectTargets(fields []*core.EmiField) []EntityTargetRef {
	var result []EntityTargetRef

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field == nil {
				continue
			}

			isSelf, _ := getSelfReferencingField(field, "")
			if field.Target != "" && !isSelf {
				result = append(result, EntityTargetRef{Target: field.Target, JsProvider: field.JsProvider})
			}

			if len(field.Fields) > 0 {
				walk(field.Fields)
			}
		}
	}

	walk(fields)
	return result
}

// entityTargetToCodeChunk resolves a relation field's `target: XEntity` (the
// documented convention for referencing another entity by its Go/generic name - PascalCase
// plus "Entity" suffix, e.g. `target: PassportEntity`) to an import of the sibling,
// actually-generated `XDto` class, aliased back to the `XEntity` name the field (and
// every other generated reference to it, since field.Target is used verbatim as both the
// TS type and the runtime constructor class name elsewhere in this file) expects.
//
// The JS/TS entity compiler only ever emits Dto/OptionalDto/action files for an entity -
// never a literal "XEntity.ts" - so importing `target` as a same-named sibling file (what
// castDtoNameToCodeChunk does for e.g. action request/response Dto names, which really are
// real sibling files) would always be a dangling import for entity relation targets. This
// used to be worked around with hand-written `export { XDto as XEntity } from "./XDto"`
// shim files sitting in the generated output directory - which then got deleted on every
// `clean: true` regeneration, since nothing there marks them as hand-written. Resolving
// the alias at generation time instead means no physical shim file is needed at all.
//
// jsProvider (EmiField.JsProvider - empty for a same-module/same-output-directory
// reference) only ever changes *where* the sibling XDto is imported *from* - the
// alias importers see (`target`, e.g. "WalletEntity") is always the field's own
// short target name, never jsProvider's full path, even when jsProvider is set.
//
// target isn't always Entity-shaped: a dto field can also reference another dto
// directly by its own Dto-suffixed name (e.g. `target: MusicalWorkDto`, no
// Entity/Dto suffix swap needed since it already *is* the Dto's class name). That
// shape used to fall straight through to castDtoNameToCodeChunk below, which has
// no notion of jsProvider at all and always assumes a same-directory sibling file
// - so a cross-module Dto-shaped target silently ignored jsProvider and produced a
// dangling "./XDto" import. Handled explicitly first so jsProvider is honored for
// both shapes.
func entityTargetToCodeChunk(target string, jsProvider string) *core.CodeChunkCompiled {
	if !strings.HasSuffix(target, "Entity") || target == "Entity" {
		if jsProvider == "" {
			return castDtoNameToCodeChunk(target)
		}

		directory, dtoClassName := parseDtoPath(jsProvider)
		objectRef := dtoClassName
		if dtoClassName != target {
			objectRef = dtoClassName + " as " + target
		}

		return &core.CodeChunkCompiled{
			ActualScript: []byte(""),
			Tokens: []core.GeneratedScriptToken{
				{Name: TOKEN_OBJ_CLASS, Value: target},
				{Name: TOKEN_ROOT_CLASS, Value: target},
			},
			CodeChunkDependensies: []core.CodeChunkDependency{
				{
					Objects:  []string{objectRef},
					Location: directory,
				},
			},
		}
	}

	source := target
	if jsProvider != "" {
		source = jsProvider
	}

	dtoName := strings.TrimSuffix(source, "Entity") + "Dto"
	directory, dtoClassName := parseDtoPath(dtoName)

	return &core.CodeChunkCompiled{
		ActualScript: []byte(""),
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_OBJ_CLASS, Value: target},
			{Name: TOKEN_ROOT_CLASS, Value: target},
		},
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Objects:  []string{dtoClassName + " as " + target},
				Location: directory,
			},
		},
	}
}

func hasClassesAsChildren(fields []*core.EmiField) bool {
	var result = false

	var walk func(f []*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field == nil {
				continue
			}

			if field.Type == core.FieldTypeArray || field.Type == core.FieldTypeCollectionNullable || field.Type == core.FieldTypeObject || field.Type == core.FieldTypeArrayNullable || field.Type == core.FieldTypeObjectNullable || field.Type == core.FieldTypeOne || field.Type == core.FieldTypeCollection {
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
	isTypeScript := ctx.HasTag(Typescript)
	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{},
	}

	if hasChildrenWithStaticFields {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Objects:  []string{"withPrefix"},
			Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "withPrefix"),
		})
	}

	// Not sure under which condition need to import partial deep.
	if isTypeScript {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Objects:  []string{"type PartialDeep"},
			Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "fetchx"),
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
		m := entityTargetToCodeChunk(item.Target, item.JsProvider)
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, m.CodeChunkDependensies...)
	}

	var translationTypeDeclaration string

	renderedClasses := jsRenderDataClasses(fields, jsctx.RootClassName, jsctx.RootClassName, "", true, ctx, jsctx)
	if len(renderedClasses) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: renderedClasses[0].ClassName})
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_OBJ_CLASS, Value: renderedClasses[0].ClassName})

		// --tags json-schema (opt-in, off by default): embed the dto's JSON Schema
		// directly on its own class as `static JsonSchema = {...}`, the same way
		// action classes already carry `static Definition = {...}` (see
		// js-action-main-class.go) - one generated file per dto, no sibling
		// .schema.json to keep in sync. Root class only:
		// formgen.BuildJSONSchemaWithTranslationKeys already inlines nested
		// object/array children into the same schema tree (properties/items), so
		// a nested SubClass repeating it would just be a redundant, deeply-nested
		// copy of the same document.
		//
		// JsonSchema's own title/description/enum-option strings are translation
		// *keys* here (see formgen.BuildJSONSchemaWithTranslationKeys' own doc
		// comment for the exact naming), not literal text - `static
		// DefaultTranslations` right below resolves every one of them back to
		// the dto's own source-language text. A locale-aware consumer overlays
		// its own translated object (typed `{ClassName}Translations` - see
		// translationTypeDeclaration below, generated once TypeScript is
		// requested - so a translation missing a key is a compile error, not a
		// blank label discovered at runtime) instead of DefaultTranslations, and
		// walks the schema resolving each key through whichever one is active.
		if ctx.HasTag(JsonSchema) {
			schema, translations := formgen.BuildJSONSchemaWithTranslationKeys(jsctx.RootClassName, jsctx.Description, fields)
			schemaJSON, err := json.MarshalIndent(schema, "", "  ")
			if err != nil {
				return nil, fmt.Errorf("js: failed marshaling JsonSchema for %v: %w", jsctx.RootClassName, err)
			}
			renderedClasses[0].ClassStaticFunctions = append(renderedClasses[0].ClassStaticFunctions,
				fmt.Sprintf("static JsonSchema = %s", schemaJSON))

			translationsJSON, err := json.MarshalIndent(translations, "", "  ")
			if err != nil {
				return nil, fmt.Errorf("js: failed marshaling DefaultTranslations for %v: %w", jsctx.RootClassName, err)
			}
			defaultTranslationsSuffix := ""
			if isTypeScript {
				// `as const` keeps every key/value a literal string type instead of
				// widening to `string`, which is what makes `keyof typeof
				// {ClassName}.DefaultTranslations` in translationTypeDeclaration
				// below an exact union of the real keys instead of just `string`.
				defaultTranslationsSuffix = " as const"
			}
			renderedClasses[0].ClassStaticFunctions = append(renderedClasses[0].ClassStaticFunctions,
				fmt.Sprintf("static DefaultTranslations = %s%s", translationsJSON, defaultTranslationsSuffix))

			if isTypeScript {
				className := core.ToUpper(jsctx.RootClassName)
				translationTypeDeclaration = fmt.Sprintf(
					"export type %sTranslationKey = keyof typeof %s.DefaultTranslations;\nexport type %sTranslations = Record<%sTranslationKey, string>;",
					className, className, className, className,
				)
			}
		}
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
	
	{{ if .IsTypeScript }}
	static from(possibleDtoObject: {{ .ClassTypePath }}) {
	{{ else }}
	static from(possibleDtoObject) {
	{{ end }}
		return new {{ .ClassNamePath }}(possibleDtoObject);
	}

	/**
	* Creates an instance of {{ .ClassNamePath }}, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/

	{{ if .IsTypeScript }}
	static with(partialDtoObject: PartialDeep<{{ .ClassTypePath }}>) {
	{{ else }}
	static with(partialDtoObject) {
	{{ end }}
		return new {{ .ClassNamePath }}(partialDtoObject);
	}

	{{ if .IsTypeScript }}
	copyWith(partial: PartialDeep<{{ .ClassTypePath }}>): InstanceType<typeof {{ .ClassNamePath }}> {
	{{ else }}
	copyWith(partial) {
	{{ end }}
		return new {{ .ClassNamePath }} ({ ...this.toJSON(), ...partial });
	}

	{{ if .IsTypeScript }}
	clone(): InstanceType<typeof {{ .ClassNamePath }}> {
	{{ else }}
	clone() {
	{{ end }}
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

{{ if .translationTypeDeclaration }}
	{{ .translationTypeDeclaration }}
{{ end }}

`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"isTypeScript":               isTypeScript,
		"renderedClasses":            renderedClasses,
		"abstractFactoryClass":       abstractFactoryClass,
		"translationTypeDeclaration": translationTypeDeclaration,
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

		res.ActualScript = []byte(result)
		res.SuggestedExtension = ".js"
		res.ActualScript = []byte(result)
	}
	res.SuggestedFileName = jsctx.RootClassName

	return res, nil
}

func jsFieldTypeOnNestedClasses(field *core.EmiField, parentChain string) string {
	value := ""
	if field == nil {
		value = ""
	} else if isSelf, valuex := getSelfReferencingField(field, parentChain); isSelf {
		return valuex
	} else if field.Type == core.FieldTypeArray || field.Type == core.FieldTypeObject || field.Type == core.FieldTypeArrayNullable || field.Type == core.FieldTypeObjectNullable {
		value = core.ToUpper(parentChain) + "." + core.ToUpper(field.Name)
	} else {
		value = TsComputedField(field, false, parentChain)
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
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		target := field.Target

		isSelf, value := getSelfReferencingField(field, parentChain)

		if isSelf {
			target = value
		}

		return "MCollection<" + target + ">"
	case core.FieldTypeOne, core.FieldTypeOneNullable:

		target := field.Target

		isSelf, value := getSelfReferencingField(field, parentChain)
		if isSelf {
			target = value
		}

		// FieldTypeOneNullable needs "| null | undefined" the same way
		// FieldTypeArrayNullable/FieldTypeObjectNullable already get it below -
		// without it, the generated setter's own null-handling branch (see
		// js-setter-function.go's "one?" case) assigns a value TypeScript
		// doesn't consider valid for the field's declared type. (The private
		// field declaration's own "| null" - see jsFieldVariable.Compile in
		// js-common-fields.go - is deduped against this by that function
		// checking whether ComputedType already mentions "null" before
		// adding its own, so this doesn't end up doubled there.)
		if field.Type == core.FieldTypeOneNullable {
			return "MOne<" + target + "> | null | undefined"
		}

		return "MOne<" + target + ">"

	case core.FieldTypeArray:
		return fmt.Sprintf("MArray<InstanceType<typeof %v>>", prefix)
	case core.FieldTypeArrayNullable:
		return fmt.Sprintf("MArray<InstanceType<typeof %v>> | null | undefined", prefix)
	case core.FieldTypeObject:
		return fmt.Sprintf("InstanceType<typeof %v>", prefix)
	case core.FieldTypeObjectNullable:
		return fmt.Sprintf("InstanceType<typeof %v> | null | undefined", prefix)
	default:
		return TsComputedField(field, false, parentChain)
	}
}
