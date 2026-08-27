package kotlin

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// KotlinPermissionsGenerate renders the module's `permissions:` tree into a single
// Permissions.kt:
//
//   - one exported object (or, for a childless root, a plain val) per root
//     permission - not a single object wrapping the whole forest - so IDEs can
//     autocomplete straight down it, e.g. ManagePostsPermission.CreatePost.key. A
//     node with children is a nested Kotlin `object` (Kotlin's idiomatic
//     compile-time singleton namespace): its own key/name/title/description sit
//     directly on it as vals, alongside one nested object/val per child (named by
//     the child's own identifier, PascalCase). A node with no children is instead a
//     plain `val` holding a Permission instance directly. The self attributes are
//     always lowercase (key, name, title, description) while every child is always
//     PascalCase, so the two can never collide the way JS's flattened object (where
//     both live in the same, single-case namespace) needs a "_" escape for -
//     nothing here needs one. The root's own name is always PascalCase with a
//     trailing "Permission" (see kotlinRootPermissionName), e.g. `managePosts` ->
//     `ManagePostsPermission`, so it reads unambiguously as a permission and can't
//     collide with an unrelated single-letter root (mirrors the Go/JS generators).
//   - alongside each root, a `val <Name>PermissionList: List<Permission>`: that root
//     permission and every one of its descendants, flattened into its own list, by
//     referencing each node's own `permission` val (its own key/name/title/
//     description composed into a Permission instance) rather than re-declaring
//     separate literals, so it can never drift out of sync with it. One list per
//     root, not a single list combining every root, so a module with several
//     unrelated permission groups doesn't force every caller to filter one combined
//     list back down to the group it actually cares about.
//
// The Permission data class itself only holds a single node's own attributes - no
// Children field - since the tree structure lives in the nested objects above, not
// in the data instances; the data class exists purely as the common type every
// node's `permission` val and every `<Name>PermissionList` entry share.
//
// core.ResolvePermissionFullKeys must have already run (Emi.Preprocess does this for
// every StringToEmi/StringToEmiForAction caller) so every node here already has
// FullKey populated. Returns (nil, nil) when the module declares no permissions.
func KotlinPermissionsGenerate(
	permissions []*core.EmiPermission,
	ctx core.MicroGenContext,
) (*core.CodeChunkCompiled, error) {

	if len(permissions) == 0 {
		return nil, nil
	}

	rootVars := &strings.Builder{}
	renderKotlinPermissionRootVars(rootVars, permissions)

	const tmpl = `/**
* Permission keys generated from the module's permissions tree.
*/

data class Permission(
	val key: String,
	val name: String,
	val title: Map<String, String>? = null,
	val description: Map<String, String>? = null
)

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
		SuggestedExtension: ".kt",
		ActualScript:       buf.Bytes(),
	}, nil
}

// kotlinPermissionIdentifier is the base identifier a permission is referred to by -
// its Name when set, falling back to its Key (mirrors permissionIdentifier in
// lib/golang/go-permissions.go and jsPermissionIdentifier in
// lib/js/js-permissions.go, so every compiler names a node the same way).
func kotlinPermissionIdentifier(p *core.EmiPermission) string {
	if p.Name != "" {
		return p.Name
	}
	return p.Key
}

// kotlinChildName turns a permission's own identifier into the PascalCase Kotlin
// object/val name it's nested under (e.g. "createPost" -> "CreatePost").
func kotlinChildName(p *core.EmiPermission) string {
	return core.ToUpper(kotlinPermissionIdentifier(p))
}

// kotlinRootPermissionName turns a root permission's own identifier into its
// exported name (and, with "List" appended, its flattened list's): PascalCase with
// a trailing "Permission" always appended (e.g. "managePosts" ->
// "ManagePostsPermission") - unlike a nested child, which stays plain PascalCase
// with no suffix. Mirrors goRootPermissionName in lib/golang/go-permissions.go and
// jsRootPermissionName in lib/js/js-permissions.go.
func kotlinRootPermissionName(p *core.EmiPermission) string {
	return kotlinChildName(p) + "Permission"
}

// kotlinPermissionLiteral renders a single node's own fields as a Permission(...)
// constructor call, each field on its own line for readability once Title/
// Description are multi-entry maps (mirrors goPermissionLiteral in
// lib/golang/go-permissions.go).
func kotlinPermissionLiteral(p *core.EmiPermission, indent string) string {
	return fmt.Sprintf(
		"Permission(\n%s\tkey = %q,\n%s\tname = %q,\n%s\ttitle = %s,\n%s\tdescription = %s\n%s)",
		indent, p.EffectiveKey(),
		indent, p.Name,
		indent, kotlinStringMapLiteral(p.Title),
		indent, kotlinStringMapLiteral(p.Description),
		indent,
	)
}

// renderKotlinPermissionNode renders a single node under the given name: a plain
// `val <name>: Permission = Permission(...)` for a leaf, or a nested
// `object <name> { ... }` - its own key/name/title/description as vals, a
// `permission` val composing them into a Permission instance, then one nested
// object/val per child - for a node with children.
func renderKotlinPermissionNode(w *strings.Builder, name string, p *core.EmiPermission, indent string) {
	if len(p.Children) == 0 {
		fmt.Fprintf(w, "%sval %s: Permission = %s\n", indent, name, kotlinPermissionLiteral(p, indent))
		return
	}

	fmt.Fprintf(w, "%sobject %s {\n", indent, name)

	childIndent := indent + "\t"
	fmt.Fprintf(w, "%sval key: String = %q\n", childIndent, p.EffectiveKey())
	fmt.Fprintf(w, "%sval name: String = %q\n", childIndent, p.Name)
	fmt.Fprintf(w, "%sval title: Map<String, String>? = %s\n", childIndent, kotlinStringMapLiteral(p.Title))
	fmt.Fprintf(w, "%sval description: Map<String, String>? = %s\n", childIndent, kotlinStringMapLiteral(p.Description))
	fmt.Fprintf(w, "%sval permission: Permission = Permission(key = key, name = name, title = title, description = description)\n", childIndent)

	for _, c := range p.Children {
		if c == nil {
			continue
		}
		w.WriteString("\n")
		renderKotlinPermissionNode(w, kotlinChildName(c), c, childIndent)
	}

	fmt.Fprintf(w, "%s}\n", indent)
}

// renderKotlinPermissionRootVars renders, for each root permission, one exported
// object (or plain val, for a childless root) - not a single object wrapping the
// whole forest - named PascalCase with a trailing "Permission" (see
// kotlinRootPermissionName). Right alongside it, a `<Name>PermissionList` val
// flattens that same root (and everything nested under it) into its own
// List<Permission>, by referencing each node's own `permission` val rather than
// re-declaring separate literals, so it can never drift out of sync with it.
func renderKotlinPermissionRootVars(w *strings.Builder, permissions []*core.EmiPermission) {
	for _, p := range permissions {
		if p == nil {
			continue
		}

		name := kotlinRootPermissionName(p)
		fmt.Fprintf(w, "// %s mirrors the %q permission node (and everything nested under it),\n", name, kotlinPermissionIdentifier(p))
		fmt.Fprintf(w, "// so it can be navigated directly, e.g. %s.key.\n", name)
		renderKotlinPermissionNode(w, name, p, "")
		w.WriteString("\n")

		fmt.Fprintf(w, "// %sList flattens %s (and everything nested under it) into a single list,\n", name, name)
		fmt.Fprintf(w, "// for anything that wants to walk them all at once (seeding an ACL table,\n")
		fmt.Fprintf(w, "// rendering a permissions picker, ...).\n")
		fmt.Fprintf(w, "val %sList: List<Permission> = listOf(\n", name)
		for _, path := range collectKotlinPermissionPaths(p, name) {
			fmt.Fprintf(w, "\t%s,\n", path)
		}
		w.WriteString(")\n\n")
	}
}

// collectKotlinPermissionPaths returns, for a single root permission accessed under
// varPath, the Kotlin expression that refers to *this* node's own Permission
// instance (self first, then recursively each descendant's, in the same order
// renderKotlinPermissionNode emits them). A leaf's own val already is that
// Permission instance; a branch's is its nested object's `permission` val. Mirrors
// collectPermissionPaths in lib/golang/go-permissions.go.
func collectKotlinPermissionPaths(p *core.EmiPermission, varPath string) []string {
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
		paths = append(paths, collectKotlinPermissionPaths(c, varPath+"."+kotlinChildName(c))...)
	}

	return paths
}

// kotlinStringMapLiteral renders a map[string]string as a Kotlin mapOf(...) literal,
// sorted by key so the generated file is deterministic across runs.
func kotlinStringMapLiteral(m map[string]string) string {
	if len(m) == 0 {
		return "null"
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	pairs := make([]string, 0, len(keys))
	for _, k := range keys {
		pairs = append(pairs, fmt.Sprintf("%q to %q", k, m[k]))
	}

	return "mapOf(" + strings.Join(pairs, ", ") + ")"
}
