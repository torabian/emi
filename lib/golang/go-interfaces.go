package golang

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// An emi interface becomes a Go interface of accessors: Go interfaces can't hold fields,
// so every field turns into a Get<Field>() method (and Set<Field>(v) when the interface
// is mutable), and each dto implementing it gets those methods generated on its struct.
// Any function taking the interface then accepts every one of those dtos.
//
// A dto in another package satisfies it just the same - Go interfaces are structural -
// so an interface declared in one module (Provider empty there) is only imported, never
// generated again, by the modules that implement it (Provider set) - see
// core.EmiInterface.IsGoReference.

type goInterfaceField struct {
	Field *core.EmiField
	Type  string
}

func goInterfaceFields(iface *core.EmiInterface, goctx GoCommonStructContext) []goInterfaceField {
	return goInterfaceFieldsOf(goInterfaceTypedFields(iface), iface.GetClassName(), goctx)
}

// goInterfaceTypedFields are the fields the interface's Go types (and so its imports) are
// derived from: as declared for a dto, as the entity's own struct types them for
// `context: entity`.
func goInterfaceTypedFields(iface *core.EmiInterface) []*core.EmiField {
	if iface.Context == "entity" {
		return goEntityFieldViews(iface.Fields)
	}
	return iface.Fields
}

// goEntityFieldViews are the fields as an entity's own Go struct types them. The entity
// storage layer turns one/one? into a plain struct (class/class?) - see applyEntityGormTags -
// while everything an interface can hold besides that is typed the same as in a dto. An
// interface declared for entities (context: entity) must be typed the entity way, or no
// entity could satisfy it.
func goEntityFieldViews(fields []*core.EmiField) []*core.EmiField {
	views := make([]*core.EmiField, 0, len(fields))
	for _, f := range fields {
		if f == nil {
			continue
		}
		view := *f
		switch f.Type {
		case core.FieldTypeOne:
			view.Type = core.FieldTypeClass
		case core.FieldTypeOneNullable:
			view.Type = core.FieldTypeClassNullable
		}
		views = append(views, &view)
	}
	return views
}

// goInterfaceFieldsOf pairs each field with its Go type, exactly as the struct generator
// types it - a dto's struct field and the accessor returning it can never disagree.
func goInterfaceFieldsOf(fields []*core.EmiField, parentChain string, goctx GoCommonStructContext) []goInterfaceField {
	out := make([]goInterfaceField, 0, len(fields))
	for _, f := range fields {
		if f == nil {
			continue
		}
		// same "+" cleanup goRenderField applies
		goType := strings.TrimSpace(strings.ReplaceAll(goFieldTypeOnNestedClasses(f, parentChain, goctx), "+", ""))
		out = append(out, goInterfaceField{Field: f, Type: goType})
	}
	return out
}

// GoInterfaceGenerate renders the Go interface type of a locally declared interface, as
// its own file named after the type.
func GoInterfaceGenerate(iface *core.EmiInterface, ctx core.MicroGenContext, goctx GoCommonStructContext) (*core.CodeChunkCompiled, error) {
	name := iface.GetClassName()
	fields := goInterfaceFields(iface, goctx)

	var buf bytes.Buffer
	if iface.Description != "" {
		fmt.Fprintf(&buf, "// %s %s\n", name, oneLine(iface.Description))
	} else {
		fmt.Fprintf(&buf, "// %s is implemented by every dto that lists %q under implements.\n", name, iface.Name)
	}
	fmt.Fprintf(&buf, "type %s interface {\n", name)
	for _, f := range fields {
		if f.Field.Description != "" {
			fmt.Fprintf(&buf, "\t// %s\n", oneLine(f.Field.Description))
		}
		fmt.Fprintf(&buf, "\tGet%s() %s\n", f.Field.PublicName(), f.Type)
		if iface.Mutable {
			fmt.Fprintf(&buf, "\tSet%s(v %s)\n", f.Field.PublicName(), f.Type)
		}
	}
	buf.WriteString("}\n")

	// The nested types of object/array fields belong to the interface: declared once here,
	// aliased by every implementer's own nested type (see goInterfaceAliases).
	for _, f := range iface.Fields {
		if f == nil || !isNestedClassField(f) {
			continue
		}
		fieldClass := name + core.ToUpper(f.Name)
		for _, st := range goRenderStructs(f.Fields, core.ToUpper(f.Name), fieldClass, f.Name, fieldClass, ctx, goctx) {
			writeGoStructTree(&buf, st)
		}
	}

	chunk := &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		SuggestedFileName:     name,
		SuggestedExtension:    ".go",
		CodeChunkDependensies: goInterfaceDependencies(goInterfaceTypedFields(iface), name, goctx, buf.String()),
	}
	return chunk, nil
}

// goInterfaceDependencies are the imports the interface's field types need - the same
// sources PrepareStruct uses for a struct, without its unconditional encoding/json (an
// interface never calls it, and an unused import doesn't compile).
func goInterfaceDependencies(fields []*core.EmiField, className string, goctx GoCommonStructContext, script string) []core.CodeChunkDependency {
	var deps []core.CodeChunkDependency

	for _, item := range CollectComplexClasses(fields) {
		if location := findComplexLocation(item, goctx); location != "" {
			deps = append(deps, core.CodeChunkDependency{Objects: []string{}, Location: location})
		}
	}
	for _, item := range CollectTargets(fields, className) {
		deps = append(deps, castDtoNameToCodeChunk(item).CodeChunkDependensies...)
	}
	// Decided from the code actually written, not from field types: a class pointer (an
	// entity relation) is a "?" type without touching emigo, and an unused import doesn't compile.
	if strings.Contains(script, "emigo.") {
		deps = append(deps, core.CodeChunkDependency{Location: goctx.EmiLocation})
	}
	return deps
}

// GoInterfaceImplementation returns the accessor methods (plus a compile-time assertion)
// that make ownerClass satisfy each of ifaces, and the import a referenced interface
// needs. ownerFields are the owner's own fields - already including the interface's - so
// each accessor returns exactly what the struct field is typed as.
func GoInterfaceImplementation(ownerClass string, ownerFields []*core.EmiField, ifaces []*core.EmiInterface, goctx GoCommonStructContext) (string, []core.CodeChunkDependency, error) {
	byName := map[string]goInterfaceField{}
	for _, f := range goInterfaceFieldsOf(ownerFields, ownerClass, goctx) {
		byName[f.Field.Name] = f
	}

	var buf bytes.Buffer
	var deps []core.CodeChunkDependency
	generated := map[string]bool{}

	for _, iface := range ifaces {
		qualified := iface.GetClassName()
		if iface.IsGoReference() {
			if iface.Module == "" {
				return "", nil, fmt.Errorf("interface %q sets provider but no module (the Go package alias to import it as)", iface.Name)
			}
			qualified = iface.Module + "." + iface.GetClassName()
			deps = append(deps, core.CodeChunkDependency{Objects: []string{iface.Module}, Location: iface.Provider})
		}

		for _, want := range iface.Fields {
			f, ok := byName[want.Name]
			if !ok {
				return "", nil, fmt.Errorf("%s implements %q but has no field %q", ownerClass, iface.Name, want.Name)
			}
			method := f.Field.PublicName()
			// two interfaces sharing a field need the accessor once
			if !generated["Get"+method] {
				generated["Get"+method] = true
				fmt.Fprintf(&buf, "func (x *%s) Get%s() %s { return x.%s }\n", ownerClass, method, f.Type, method)
			}
			if iface.Mutable && !generated["Set"+method] {
				generated["Set"+method] = true
				fmt.Fprintf(&buf, "func (x *%s) Set%s(v %s) { x.%s = v }\n", ownerClass, method, f.Type, method)
			}
		}
		fmt.Fprintf(&buf, "var _ %s = (*%s)(nil)\n", qualified, ownerClass)
	}

	return buf.String(), deps, nil
}

func oneLine(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// isNestedClassField reports whether a field renders its own nested struct.
func isNestedClassField(f *core.EmiField) bool {
	return core.IsInlineNestedType(f.Type)
}

// writeGoStructTree prints a rendered struct and, after it, its nested structs - the same
// shape the dto template prints.
func writeGoStructTree(buf *bytes.Buffer, st goRenderedStruct) {
	buf.WriteString(st.GoDoc + "\n")
	fmt.Fprintf(buf, "%s {\n", st.Signature)
	for _, f := range st.Fields {
		buf.WriteString(f.PrivateField + "\n")
	}
	buf.WriteString("}\n\n")
	for _, sub := range st.SubClasses {
		writeGoStructTree(buf, sub)
	}
}

// goInterfaceAliases maps the nested classes of a dto's fields that come from ifaces to the
// interface's shared nested types. The dto keeps generating (and calling) its own
// ClothDtoContent, ClothDtoContentParts, ... - each just becomes an alias of the
// interface's TitlableContent, TitlableContentParts - so the field's Go type is exactly
// what the interface's accessor returns, and all the generated CLI/cast code keeps
// compiling untouched.
func goInterfaceAliases(ownerClass string, ifaces []*core.EmiInterface) map[string]string {
	aliases := map[string]string{}
	owner := core.ToUpper(ownerClass)

	for _, iface := range ifaces {
		qualifier := ""
		if iface.IsGoReference() {
			qualifier = iface.Module + "."
		}
		shared := qualifier + iface.GetClassName()

		var walk func(fields []*core.EmiField, rel string)
		walk = func(fields []*core.EmiField, rel string) {
			for _, f := range fields {
				if f == nil || !isNestedClassField(f) {
					continue
				}
				childRel := rel + core.ToUpper(f.Name)
				aliases[owner+childRel] = shared + childRel
				walk(f.Fields, childRel)
			}
		}
		walk(iface.Fields, "")
	}
	return aliases
}
