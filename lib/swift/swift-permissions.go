package swift

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// SwiftPermissionsGenerate renders the module's `permissions:` tree into a single
// Permissions.swift:
//
//   - one exported enum namespace (or, for a childless root, a plain let) per root
//     permission - not a single enum wrapping the whole forest - so Xcode can
//     autocomplete straight down it, e.g. ManagePostsPermission.CreatePost.key. A
//     node with children is a nested Swift `enum` with no cases - Swift's idiomatic
//     uninstantiable namespace - holding its own key/name/title/description as
//     `static let`s, alongside one nested enum/let per child (named by the child's
//     own identifier, PascalCase). A node with no children is instead a plain `let`
//     holding a Permission instance directly. The self attributes are always
//     lowercase (key, name, title, description) while every child is always
//     PascalCase, so the two can never collide the way JS's flattened object (where
//     both live in the same, single-case namespace) needs a "_" escape for -
//     nothing here needs one. The root's own name is always PascalCase with a
//     trailing "Permission" (see swiftRootPermissionName), e.g. `managePosts` ->
//     `ManagePostsPermission`, so it reads unambiguously as a permission and can't
//     collide with an unrelated single-letter root (mirrors the Go/Kotlin/JS
//     generators).
//   - alongside each root, a `let <Name>PermissionList: [Permission]`: that root
//     permission and every one of its descendants, flattened into its own array, by
//     referencing each node's own `permission` static let (its own key/name/title/
//     description composed into a Permission instance) rather than re-declaring
//     separate literals, so it can never drift out of sync with it. One list per
//     root, not a single list combining every root, so a module with several
//     unrelated permission groups doesn't force every caller to filter one combined
//     list back down to the group it actually cares about.
//
// The Permission struct itself only holds a single node's own attributes - no
// children property - since the tree structure lives in the nested enums above, not
// in the data instances; the struct exists purely as the common type every node's
// `permission` let and every `<Name>PermissionList` entry share.
//
// core.ResolvePermissionFullKeys must have already run (Emi.Preprocess does this for
// every StringToEmi/StringToEmiForAction caller) so every node here already has
// FullKey populated. Returns (nil, nil) when the module declares no permissions.
func SwiftPermissionsGenerate(
	permissions []*core.EmiPermission,
	ctx core.MicroGenContext,
) (*core.CodeChunkCompiled, error) {

	if len(permissions) == 0 {
		return nil, nil
	}

	rootVars := &strings.Builder{}
	renderSwiftPermissionRootVars(rootVars, permissions)

	const tmpl = `/**
* Permission keys generated from the module's permissions tree.
*/

struct Permission: Codable {
	let key: String
	let name: String
	let title: [String: String]?
	let description: [String: String]?
}

{{ .rootVars }}`

	t := template.Must(template.New("permissions").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"rootVars": rootVars.String(),
	}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		SuggestedFileName:  "Permissions",
		SuggestedExtension: ".swift",
		ActualScript:       buf.Bytes(),
	}, nil
}

// swiftPermissionIdentifier is the base identifier a permission is referred to by -
// its Name when set, falling back to its Key (mirrors permissionIdentifier in
// lib/golang/go-permissions.go and jsPermissionIdentifier in
// lib/js/js-permissions.go, so every compiler names a node the same way).
func swiftPermissionIdentifier(p *core.EmiPermission) string {
	if p.Name != "" {
		return p.Name
	}
	return p.Key
}

// swiftChildName turns a permission's own identifier into the PascalCase Swift
// enum/let name it's nested under (e.g. "createPost" -> "CreatePost").
func swiftChildName(p *core.EmiPermission) string {
	return core.ToUpper(swiftPermissionIdentifier(p))
}

// swiftRootPermissionName turns a root permission's own identifier into its
// exported name (and, with "List" appended, its flattened array's): PascalCase with
// a trailing "Permission" always appended (e.g. "managePosts" ->
// "ManagePostsPermission") - unlike a nested child, which stays plain PascalCase
// with no suffix. Mirrors goRootPermissionName in lib/golang/go-permissions.go and
// jsRootPermissionName in lib/js/js-permissions.go.
func swiftRootPermissionName(p *core.EmiPermission) string {
	return swiftChildName(p) + "Permission"
}

// swiftPermissionLiteral renders a single node's own fields as a Permission(...)
// initializer call, each field on its own line for readability once title/
// description are multi-entry dictionaries (mirrors goPermissionLiteral in
// lib/golang/go-permissions.go).
func swiftPermissionLiteral(p *core.EmiPermission, indent string) string {
	return fmt.Sprintf(
		"Permission(\n%s\tkey: %q,\n%s\tname: %q,\n%s\ttitle: %s,\n%s\tdescription: %s\n%s)",
		indent, p.EffectiveKey(),
		indent, p.Name,
		indent, swiftStringMapLiteral(p.Title),
		indent, swiftStringMapLiteral(p.Description),
		indent,
	)
}

// renderSwiftPermissionNode renders a single node under the given name: for a leaf,
// a `let <name>: Permission = Permission(...)` (root/file scope) or
// `static let <name>: Permission = Permission(...)` (nested inside an enum
// namespace - Swift enums can't hold instance-stored properties); or, for a node
// with children, a nested `enum <name> { ... }` - its own key/name/title/
// description as `static let`s, a `permission` static let composing them into a
// Permission instance, then one nested enum/let per child.
func renderSwiftPermissionNode(w *strings.Builder, name string, p *core.EmiPermission, indent string) {
	if len(p.Children) == 0 {
		// A file-scope root (indent == "") is a plain `let`; nested inside an enum
		// namespace it must be `static let` - Swift enums can't hold instance-stored
		// properties, only static ones.
		keyword := "let"
		if indent != "" {
			keyword = "static let"
		}
		fmt.Fprintf(w, "%s%s %s: Permission = %s\n", indent, keyword, name, swiftPermissionLiteral(p, indent))
		return
	}

	fmt.Fprintf(w, "%senum %s {\n", indent, name)

	childIndent := indent + "\t"
	fmt.Fprintf(w, "%sstatic let key: String = %q\n", childIndent, p.EffectiveKey())
	fmt.Fprintf(w, "%sstatic let name: String = %q\n", childIndent, p.Name)
	fmt.Fprintf(w, "%sstatic let title: [String: String]? = %s\n", childIndent, swiftStringMapLiteral(p.Title))
	fmt.Fprintf(w, "%sstatic let description: [String: String]? = %s\n", childIndent, swiftStringMapLiteral(p.Description))
	fmt.Fprintf(w, "%sstatic let permission: Permission = Permission(key: key, name: name, title: title, description: description)\n", childIndent)

	for _, c := range p.Children {
		if c == nil {
			continue
		}
		w.WriteString("\n")
		renderSwiftPermissionNode(w, swiftChildName(c), c, childIndent)
	}

	fmt.Fprintf(w, "%s}\n", indent)
}

// renderSwiftPermissionRootVars renders, for each root permission, one exported enum
// (or plain let, for a childless root) - not a single enum wrapping the whole
// forest - named PascalCase with a trailing "Permission" (see
// swiftRootPermissionName). Right alongside it, a `<Name>PermissionList` let
// flattens that same root (and everything nested under it) into its own
// [Permission], by referencing each node's own `permission` static let rather than
// re-declaring separate literals, so it can never drift out of sync with it.
func renderSwiftPermissionRootVars(w *strings.Builder, permissions []*core.EmiPermission) {
	for _, p := range permissions {
		if p == nil {
			continue
		}

		name := swiftRootPermissionName(p)
		fmt.Fprintf(w, "// %s mirrors the %q permission node (and everything nested under it),\n", name, swiftPermissionIdentifier(p))
		fmt.Fprintf(w, "// so it can be navigated directly, e.g. %s.key.\n", name)
		renderSwiftPermissionNode(w, name, p, "")
		w.WriteString("\n")

		fmt.Fprintf(w, "// %sList flattens %s (and everything nested under it) into a single array,\n", name, name)
		fmt.Fprintf(w, "// for anything that wants to walk them all at once (seeding an ACL table,\n")
		fmt.Fprintf(w, "// rendering a permissions picker, ...).\n")
		fmt.Fprintf(w, "let %sList: [Permission] = [\n", name)
		for _, path := range collectSwiftPermissionPaths(p, name) {
			fmt.Fprintf(w, "\t%s,\n", path)
		}
		w.WriteString("]\n\n")
	}
}

// collectSwiftPermissionPaths returns, for a single root permission accessed under
// varPath, the Swift expression that refers to *this* node's own Permission
// instance (self first, then recursively each descendant's, in the same order
// renderSwiftPermissionNode emits them). A leaf's own let already is that
// Permission instance; a branch's is its nested enum's `permission` static let.
// Mirrors collectPermissionPaths in lib/golang/go-permissions.go.
func collectSwiftPermissionPaths(p *core.EmiPermission, varPath string) []string {
	if p == nil {
		return nil
	}

	if len(p.Children) == 0 {
		return []string{varPath}
	}

	paths := []string{varPath + ".permission"}
	for _, c := range p.Children {
		if c == nil {
			continue
		}
		paths = append(paths, collectSwiftPermissionPaths(c, varPath+"."+swiftChildName(c))...)
	}

	return paths
}

// swiftStringMapLiteral renders a map[string]string as a Swift dictionary literal,
// sorted by key so the generated file is deterministic across runs.
func swiftStringMapLiteral(m map[string]string) string {
	if len(m) == 0 {
		return "nil"
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	pairs := make([]string, 0, len(keys))
	for _, k := range keys {
		pairs = append(pairs, fmt.Sprintf("%q: %q", k, m[k]))
	}

	return "[" + strings.Join(pairs, ", ") + "]"
}
