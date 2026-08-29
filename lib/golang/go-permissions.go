package golang

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// GoPermissionsGenerate renders the module's `permissions:` tree into a single
// Permissions.go:
//
//   - one exported var per root permission - not a single var wrapping the whole
//     forest - each a fully typed, nested anonymous struct. Every node embeds
//     emigo.Permission and has one field per child, so an IDE can autocomplete
//     straight down it (e.g. ManagePostsPermission.CreatePost.Key) with no
//     map/string lookup and thus no risk of a typo'd key only failing at runtime -
//     which is also why there's no flat per-node constant here the way JS/Kotlin/
//     Swift have: with everything already reachable this way, a parallel
//     `CreatePostPermission = "..."` would just be the same string under a
//     second, disconnected name to keep in sync. The root var's own name is always
//     PascalCase with a trailing "Permission" (see goRootPermissionName) - e.g.
//     `managePosts` -> `ManagePostsPermission` - so it reads unambiguously as a
//     permission and two unrelated single-letter roots can't collide.
//   - alongside each root var, a <Name>PermissionList var: that root permission and
//     every one of its descendants, flattened into its own []emigo.Permission (e.g.
//     ManagePostsPermissionList, next to ManagePostsPermission) by referencing the
//     root var's own fields (e.g. ManagePostsPermission.Permission,
//     ManagePostsPermission.CreatePost) - not re-declared as separate literals - so
//     it can never drift from them, for anything that wants to walk just that
//     root's permissions at once (seeding an ACL table, rendering a permissions
//     picker, ...). One list per root, not a single list combining every root, so a
//     module with several unrelated permission groups doesn't force every caller to
//     filter one combined list back down to the group it actually cares about.
//
// The Permission type itself is emigo.Permission (see emigo/Permission.go), not
// redefined here - every generated Permissions.go across every module shares the
// same type, instead of each one declaring its own incompatible copy.
//
// core.ResolvePermissionFullKeys must have already run (Emi.Preprocess does this for
// every StringToEmi/StringToEmiForAction caller) so every node here already has
// FullKey populated. Returns (nil, nil) when the module declares no permissions.
func GoPermissionsGenerate(
	permissions []*core.EmiPermission,
	ctx core.MicroGenContext,
	emigoImportPath string,
) (*core.CodeChunkCompiled, error) {

	if len(permissions) == 0 {
		return nil, nil
	}

	rootVars := &strings.Builder{}
	renderPermissionRootVars(rootVars, permissions)

	const tmpl = `/**
* Permission keys generated from the module's permissions tree.
*/

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
		SuggestedExtension: ".go",
		ActualScript:       buf.Bytes(),
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: emigoImportPath},
		},
	}, nil
}

// permissionIdentifier is the base identifier a permission is referred to by - its
// Name when set, falling back to its Key (guaranteed unique among siblings, unlike
// Name which is optional). Used to derive its struct/var field name.
func permissionIdentifier(p *core.EmiPermission) string {
	if p.Name != "" {
		return p.Name
	}
	return p.Key
}

// permissionFieldName turns a permission's own identifier (its Name, falling back
// to its Key) into the PascalCase Go struct field name for it (e.g. "createPost" ->
// "CreatePost").
func permissionFieldName(p *core.EmiPermission) string {
	return core.ToUpper(permissionIdentifier(p))
}

// goPermissionLiteral renders a single node's own fields (Key/Name/Title/
// Description) as an emigo.Permission{...} composite literal, with no Children -
// used both as the leaf-node field value and embedded (as the promoted Permission
// field) inside a branch node's own anonymous struct value. Each field gets its own
// line rather than one long comma-separated line - gofmt only reflows a composite
// literal onto multiple aligned lines when the source already breaks it that way,
// so with Title/Description often being multi-entry maps, keeping this on one line
// would leave some of these unreadably long once AsFullDocument's gofmt pass runs
// over the file.
func goPermissionLiteral(p *core.EmiPermission) string {
	return fmt.Sprintf(
		"emigo.Permission{\nKey: %q,\nName: %q,\nTitle: %s,\nDescription: %s,\n}",
		p.EffectiveKey(), p.Name, goStringMapLiteral(p.Title), goStringMapLiteral(p.Description),
	)
}

// selfPermissionFieldName returns the field name a node's own emigo.Permission
// value is stored under within its generated struct: "" to embed it anonymously
// (promoting Key/Name/... straight onto the node - the common case), or the
// explicit name "Permission_" when one of its children happens to have that exact
// field name ("Permission", the name Go gives an anonymous emigo.Permission field)
// - embedding both would otherwise declare/initialize the same field name twice.
func selfPermissionFieldName(p *core.EmiPermission) string {
	for _, c := range p.Children {
		if c != nil && permissionFieldName(c) == "Permission" {
			return "Permission_"
		}
	}
	return ""
}

// goPermissionNodeType returns the Go type of a permission node: plain
// emigo.Permission for a leaf, or - for a node with children - a struct holding
// emigo.Permission (so its own Key/... are promoted, unless selfPermissionFieldName
// says a child's field name collides with that promotion, in which case it's given
// an explicit, non-embedded name instead) plus one field per child, typed the same
// way recursively. The struct itself is anonymous rather than named so two
// permissions that happen to share a Name in different branches of the tree (an
// alarm you can't get from a top-level type name, since Go requires those to be
// package-unique) can never collide - there is no name to collide over.
func goPermissionNodeType(p *core.EmiPermission, indent string) string {
	if len(p.Children) == 0 {
		return "emigo.Permission"
	}

	var b strings.Builder
	b.WriteString("struct {\n")
	if self := selfPermissionFieldName(p); self != "" {
		fmt.Fprintf(&b, "%s\t%s emigo.Permission\n", indent+"\t", self)
	} else {
		fmt.Fprintf(&b, "%s\temigo.Permission\n", indent+"\t")
	}
	for _, c := range p.Children {
		if c == nil {
			continue
		}
		fmt.Fprintf(&b, "%s\t%s %s\n", indent+"\t", permissionFieldName(c), goPermissionNodeType(c, indent+"\t"))
	}
	fmt.Fprintf(&b, "%s}", indent)

	return b.String()
}

// goPermissionNodeValue renders the composite literal for a permission node,
// matching whatever goPermissionNodeType returned for it.
func goPermissionNodeValue(p *core.EmiPermission, indent string) string {
	if len(p.Children) == 0 {
		return goPermissionLiteral(p)
	}

	selfField := "Permission"
	if self := selfPermissionFieldName(p); self != "" {
		selfField = self
	}

	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", goPermissionNodeType(p, indent))
	fmt.Fprintf(&b, "%s\t%s: %s,\n", indent+"\t", selfField, goPermissionLiteral(p))
	for _, c := range p.Children {
		if c == nil {
			continue
		}
		fmt.Fprintf(&b, "%s\t%s: %s,\n", indent+"\t", permissionFieldName(c), goPermissionNodeValue(c, indent+"\t"))
	}
	fmt.Fprintf(&b, "%s}", indent)

	return b.String()
}

// goRootPermissionName turns a root permission's own identifier into its exported
// var name (and, with "List" appended, its flattened slice's): PascalCase with a
// trailing "Permission" always appended (e.g. "ali" -> "AliPermission",
// "managePosts" -> "ManagePostsPermission") - unlike a nested child field, which
// stays plain PascalCase with no suffix. Mirrors jsRootPermissionName in
// lib/js/js-permissions.go, so a root permission reads unambiguously as one at a
// glance in both languages.
func goRootPermissionName(p *core.EmiPermission) string {
	return permissionFieldName(p) + "Permission"
}

// renderPermissionRootVars renders, for each root permission, one exported var -
// not a single wrapping var for the whole forest - a fully typed nested struct (or,
// for a childless root, a plain emigo.Permission) so an IDE can autocomplete a path
// straight down it, e.g. ManagePostsPermission.CreatePost.Key. Its name is always
// PascalCase with a trailing "Permission" (see goRootPermissionName), so it reads
// unambiguously as a permission and can't collide with an unrelated single-letter
// root. Right alongside it, a <Name>PermissionList var flattens that same root (and
// everything nested under it) into its own []emigo.Permission, by referencing the
// var's own fields rather than re-declaring separate literals, so it can never
// drift out of sync with it.
func renderPermissionRootVars(w *strings.Builder, permissions []*core.EmiPermission) {
	for _, p := range permissions {
		if p == nil {
			continue
		}

		name := goRootPermissionName(p)
		fmt.Fprintf(w, "// %s mirrors the %q permission node (and everything nested under it),\n", name, permissionIdentifier(p))
		fmt.Fprintf(w, "// so it can be navigated directly, e.g. %s.Key.\n", name)
		fmt.Fprintf(w, "var %s = %s\n\n", name, goPermissionNodeValue(p, ""))

		fmt.Fprintf(w, "// %sList flattens %s (and everything nested under it) into a single slice,\n", name, name)
		fmt.Fprintf(w, "// for anything that wants to walk them all at once (seeding an ACL table,\n")
		fmt.Fprintf(w, "// rendering a permissions picker, ...).\n")
		fmt.Fprintf(w, "var %sList = []emigo.Permission{\n", name)
		for _, path := range collectPermissionPaths(p, name) {
			fmt.Fprintf(w, "\t%s,\n", path)
		}
		w.WriteString("}\n\n")
	}
}

// collectPermissionPaths returns, for a single root permission var accessed under
// varPath, the Go expression that refers to *this* node's own emigo.Permission
// value (self first, in the same order goPermissionNodeValue emits fields),
// followed recursively by every descendant's. A leaf is the var path itself (its
// field is already emigo.Permission); a branch is varPath plus whatever field name
// its own value was stored under (see selfPermissionFieldName).
func collectPermissionPaths(p *core.EmiPermission, varPath string) []string {
	if p == nil {
		return nil
	}

	if len(p.Children) == 0 {
		return []string{varPath}
	}

	selfField := "Permission"
	if self := selfPermissionFieldName(p); self != "" {
		selfField = self
	}

	paths := []string{varPath + "." + selfField}
	for _, c := range p.Children {
		if c == nil {
			continue
		}
		paths = append(paths, collectPermissionPaths(c, varPath+"."+permissionFieldName(c))...)
	}

	return paths
}

// goStringMapLiteral renders a map[string]string as a Go literal, sorted by key so
// the generated file is deterministic across runs.
func goStringMapLiteral(m map[string]string) string {
	if len(m) == 0 {
		return "nil"
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	b.WriteString("map[string]string{")
	for _, k := range keys {
		fmt.Fprintf(&b, "%q: %q, ", k, m[k])
	}
	b.WriteString("}")

	return b.String()
}
