package kotlin

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// An emi interface becomes a Kotlin `interface` with one property per field, and every
// dto implementing it lists it as a supertype and marks the matching constructor
// properties `override` - so a function typed with the interface accepts all of them.
//
// An interface declared in another module (Provider set) is not generated again; with
// Module set it is imported from that package instead.

// interfaceFieldOverrides maps the name of every field the interfaces declare to
// whether the interface wants it mutable (a `var`).
func interfaceFieldOverrides(ifaces []*core.EmiInterface) map[string]bool {
	overrides := map[string]bool{}
	for _, iface := range ifaces {
		for _, f := range iface.Fields {
			if f != nil {
				overrides[f.Name] = overrides[f.Name] || iface.Mutable
			}
		}
	}
	return overrides
}

// interfaceSupertypes is the body of the `: A, B` clause.
func interfaceSupertypes(ifaces []*core.EmiInterface) string {
	names := make([]string, 0, len(ifaces))
	for _, iface := range ifaces {
		names = append(names, iface.GetClassName())
	}
	return strings.Join(names, ", ")
}

// kotlinInterfaceImports imports the interfaces declared in another package.
func kotlinInterfaceImports(ifaces []*core.EmiInterface) []core.CodeChunkDependency {
	deps := []core.CodeChunkDependency{}
	for _, iface := range ifaces {
		if iface.IsGoReference() && iface.Module != "" {
			deps = append(deps, core.CodeChunkDependency{Location: iface.Module + "." + iface.GetClassName()})
		}
	}
	return deps
}

// KotlinInterfaceGenerate renders a locally declared interface as its own Kotlin file.
func KotlinInterfaceGenerate(iface *core.EmiInterface, ctx core.MicroGenContext, goctx commonClassContext) (*core.CodeChunkCompiled, error) {
	name := iface.GetClassName()
	goctx.RootClassName = name

	keyword := "val"
	if iface.Mutable {
		keyword = "var"
	}

	var buf bytes.Buffer
	if iface.Description != "" {
		fmt.Fprintf(&buf, "/** %s */\n", strings.Join(strings.Fields(iface.Description), " "))
	}
	fmt.Fprintf(&buf, "interface %s {\n", name)
	for _, f := range iface.Fields {
		if f == nil {
			continue
		}
		t := strings.TrimSpace(strings.ReplaceAll(goFieldTypeOnNestedClasses(f, name, name), "+", ""))
		fmt.Fprintf(&buf, "    %s %s: %s\n", keyword, core.ToLower(f.PublicName()), t)
	}
	buf.WriteString("}\n")

	chunk := &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  name,
		SuggestedExtension: ".kt",
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "emikot.MaybeField"},
			{Location: "emikot.Maybe"},
		},
	}
	for _, item := range CollectComplexClasses(iface.Fields) {
		if location := findComplexLocation(item, goctx); location != "" {
			chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies, core.CodeChunkDependency{Objects: []string{}, Location: location})
		}
	}
	chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies, kotlinCollectTargetDeps(iface.Fields, name)...)

	return chunk, nil
}
