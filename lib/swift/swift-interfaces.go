package swift

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// An emi interface becomes a Swift `protocol` with one property requirement per field
// ({ get }, or { get set } when the interface is mutable), and every dto implementing it
// declares conformance - so a function typed with the protocol accepts all of them.
//
// An interface declared in another module (Provider set) is not generated again.

// protocolConformances is the `, A, B` tail appended to a struct's inheritance list.
func protocolConformances(ifaces []*core.EmiInterface) string {
	var b strings.Builder
	for _, iface := range ifaces {
		b.WriteString(", " + iface.GetClassName())
	}
	return b.String()
}

// interfaceMutableFields names every field a mutable interface requires a setter for.
func interfaceMutableFields(ifaces []*core.EmiInterface) map[string]bool {
	mutables := map[string]bool{}
	for _, iface := range ifaces {
		if !iface.Mutable {
			continue
		}
		for _, f := range iface.Fields {
			if f != nil {
				mutables[f.Name] = true
			}
		}
	}
	return mutables
}

// SwiftInterfaceGenerate renders a locally declared interface as its own Swift file.
// Property types come from the same resolver the dtos use, so they always match.
func SwiftInterfaceGenerate(iface *core.EmiInterface, ctx core.MicroGenContext, goctx commonClassContext) (*core.CodeChunkCompiled, error) {
	name := iface.GetClassName()
	goctx.RootClassName = name

	access := "{ get }"
	if iface.Mutable {
		access = "{ get set }"
	}

	var buf bytes.Buffer
	if iface.Description != "" {
		fmt.Fprintf(&buf, "/// %s\n", strings.Join(strings.Fields(iface.Description), " "))
	}
	fmt.Fprintf(&buf, "protocol %s {\n", name)
	for _, f := range iface.Fields {
		if f == nil {
			continue
		}
		t := strings.TrimSpace(strings.ReplaceAll(goFieldTypeOnNestedClasses(f, name, name), "+", ""))
		fmt.Fprintf(&buf, "    var %s: %s %s\n", core.ToLower(f.PublicName()), t, access)
	}
	buf.WriteString("}\n")

	chunk := &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		SuggestedFileName:     name + ".swift",
		CodeChunkDependensies: []core.CodeChunkDependency{},
	}
	for _, item := range CollectComplexClasses(iface.Fields) {
		if location := findComplexLocation(item, goctx); location != "" {
			chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies, core.CodeChunkDependency{Objects: []string{item}, Location: location})
		}
	}
	for _, item := range core.CollectTargets(iface.Fields, name) {
		chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies, castDtoNameToCodeChunk(item).CodeChunkDependensies...)
	}

	return chunk, nil
}
