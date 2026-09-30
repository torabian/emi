package js

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// An emi interface becomes a TypeScript `interface` with one property per field, and every
// dto implementing it declares `implements <Interface>` - so tsc checks the contract and
// any function typed with the interface accepts all of them. Plain JavaScript has no
// interfaces: nothing is generated there (the dto still gets the fields).
//
// An interface declared in another module (JsProvider set) is imported from there, never
// generated a second time. A target that embeds its own copy of the SDK runtime can't
// satisfy an interface typed against another copy's classes (two `MOne` declarations
// are unrelated types to TypeScript), so it opts out with the no-interfaces tag.

// JsInterfaceRef is one interface a dto implements, as its class needs it.
type JsInterfaceRef struct {
	// Name is the interface's TypeScript type name.
	Name string
	// JsProvider is the import path of the file declaring it (including the file name)
	// when it lives elsewhere; empty for one generated next to the dto.
	JsProvider string
}

// JsInterfaceRefs turns the interfaces a dto implements into what its class needs.
func JsInterfaceRefs(module *core.Emi, ctx core.MicroGenContext, names []string) []JsInterfaceRef {
	refs := []JsInterfaceRef{}
	if ctx.HasTag(NoInterfaces) {
		return refs
	}
	for _, iface := range module.FindInterfaces(names) {
		refs = append(refs, JsInterfaceRef{Name: iface.GetClassName(), JsProvider: iface.JsProvider})
	}
	return refs
}

// jsInterfaceImports is the import of every interface a dto implements: a type-only
// import (the marker is required once verbatimModuleSyntax is on), from the sibling file
// or, for a referenced interface, from its JsProvider.
func jsInterfaceImports(refs []JsInterfaceRef) []core.CodeChunkDependency {
	deps := make([]core.CodeChunkDependency, 0, len(refs))
	for _, ref := range refs {
		if ref.JsProvider == "" {
			deps = append(deps, core.CodeChunkDependency{
				Objects:  []string{"type " + ref.Name},
				Location: "./" + ref.Name,
			})
			continue
		}
		directory, exported := parseDtoPath(ref.JsProvider)
		object := "type " + exported
		if exported != ref.Name {
			object += " as " + ref.Name
		}
		deps = append(deps, core.CodeChunkDependency{Objects: []string{object}, Location: directory})
	}
	return deps
}

// JsInterfaceGenerate renders a locally declared interface as its own TypeScript file,
// named after the type. Returns nil for plain JavaScript output.
func JsInterfaceGenerate(iface *core.EmiInterface, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error) {
	if !ctx.HasTag(Typescript) || ctx.HasTag(NoInterfaces) {
		return nil, nil
	}

	name := iface.GetClassName()
	var buf bytes.Buffer
	if iface.Description != "" {
		fmt.Fprintf(&buf, "/**\n * %s\n */\n", oneLine(iface.Description))
	}
	writeTsInterface(&buf, name, iface.Fields, iface.Mutable)

	chunk := &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  name,
		SuggestedExtension: ".ts",
	}

	for _, item := range CollectComplexClasses(iface.Fields) {
		if location := findComplexLocation(item, jsctx); location != "" {
			chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies, core.CodeChunkDependency{
				Objects:  []string{"type " + item},
				Location: location,
			})
		}
	}
	for _, item := range CollectTargets(iface.Fields) {
		chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies, entityTargetToCodeChunk(item.Target, item.JsProvider, true).CodeChunkDependensies...)
	}
	for _, operator := range []struct {
		name  string
		types []core.FieldType
	}{
		{"MCollection", []core.FieldType{core.FieldTypeCollection, core.FieldTypeCollectionNullable}},
		{"MArray", []core.FieldType{core.FieldTypeArray, core.FieldTypeArrayNullable}},
		{"MOne", []core.FieldType{core.FieldTypeOne, core.FieldTypeOneNullable}},
	} {
		if containsTypeDeep(iface.Fields, operator.types) {
			chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies, core.CodeChunkDependency{
				Objects:  []string{"type " + operator.name},
				Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "operators"),
			})
		}
	}

	return chunk, nil
}

// tsInterfacePropertyType is the type a dto class gives the property of this field. That is
// the *setter's* parameter type, not the plain field type: the generated getter has no
// return annotation, and TypeScript types an accessor pair by the setter's parameter when
// the getter is unannotated - so it is wider than the field (a nullable field also takes
// null/undefined, a one/collection/array also takes a plain instance or array). Using
// anything narrower makes `implements` fail. Mirrors CreateSetterFunction.
func tsInterfacePropertyType(f *core.EmiField, parentChain string) string {
	// An inline object/array is typed by the interface's own nested interface, which the
	// implementer's nested class satisfies structurally (see writeTsInterface).
	if core.IsInlineNestedType(f.Type) {
		nested := parentChain + core.ToUpper(f.Name)
		switch f.Type {
		case core.FieldTypeObject:
			return nested
		case core.FieldTypeObjectNullable:
			return nested + " | null | undefined"
		case core.FieldTypeArray:
			return "MArray<" + nested + "> | " + nested + "[]"
		default: // array?
			return "MArray<" + nested + "> | " + nested + "[] | null | undefined"
		}
	}

	// The nullable suffix is added once, at the end, whatever the base type already carries.
	t := strings.TrimSuffix(strings.ReplaceAll(tsFieldTypeOnNestedClasses(f, parentChain), "+", ""), " | null | undefined")

	constructorClass := core.ToUpper(parentChain) + "." + core.ToUpper(f.Name)
	if isSelf, value := getSelfReferencingField(f, parentChain); isSelf {
		constructorClass = value
	} else if f.Target != "" {
		constructorClass = f.Target
	}

	switch f.Type {
	case core.FieldTypeOne, core.FieldTypeOneNullable:
		t += " | InstanceType<typeof " + constructorClass + ">"
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable, core.FieldTypeArray, core.FieldTypeArrayNullable:
		t += " | InstanceType<typeof " + constructorClass + ">[]"
	}
	if IsNullable(string(f.Type)) {
		t += " | null | undefined"
	}
	return t
}

// writeTsInterface writes `export interface <name>` for fields, then one interface per
// inline object/array field (named <name><Field>, recursively) - the interface owns the
// nested types, and each implementer's own nested class satisfies them structurally.
func writeTsInterface(buf *bytes.Buffer, name string, fields []*core.EmiField, mutable bool) {
	type nested struct {
		name   string
		fields []*core.EmiField
	}
	var children []nested

	fmt.Fprintf(buf, "export interface %s {\n", name)
	for _, f := range fields {
		if f == nil {
			continue
		}
		if f.Description != "" {
			fmt.Fprintf(buf, "  /** %s */\n", oneLine(f.Description))
		}
		readonly := "readonly "
		if mutable {
			readonly = ""
		}
		fmt.Fprintf(buf, "  %s%s: %s;\n", readonly, f.Name, tsInterfacePropertyType(f, name))
		if core.IsInlineNestedType(f.Type) {
			children = append(children, nested{name: name + core.ToUpper(f.Name), fields: f.Fields})
		}
	}
	buf.WriteString("}\n")

	for _, child := range children {
		buf.WriteString("\n")
		writeTsInterface(buf, child.name, child.fields, mutable)
	}
}

// containsTypeDeep is core.ContainsAnyOfTypes looking through nested children too - a
// nested interface can need MArray/MOne/... even when no top-level field does.
func containsTypeDeep(fields []*core.EmiField, types []core.FieldType) bool {
	if core.ContainsAnyOfTypes(fields, types) {
		return true
	}
	for _, f := range fields {
		if f != nil && containsTypeDeep(f.Fields, types) {
			return true
		}
	}
	return false
}
