package js

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// jsPermissionIdentifier is the base identifier a permission is referred to by -
// its Name when set, falling back to its Key (mirrors the Go generator's
// permissionIdentifier in lib/golang/go-permissions.go, so both compilers name a
// node the same way).
func jsPermissionIdentifier(p *core.EmiPermission) string {
	if p.Name != "" {
		return p.Name
	}
	return p.Key
}

// jsPermissionFieldName turns a permission's own identifier into the camelCase
// object-property name it's nested under in PermissionsTree (e.g. "createPost").
func jsPermissionFieldName(p *core.EmiPermission) string {
	return core.ToLower(jsPermissionIdentifier(p))
}

// jsRootPermissionName turns a root permission's own identifier into the exported
// const name for its tree (and, with "List" appended, its flattened array): PascalCase
// with a trailing "Permission" always appended (e.g. "managePosts" -> "ManagePostsPermission",
// "a" -> "APermission"), so every root is unambiguously a permission at a glance and
// two single-letter roots like "a" and "A" can't collide the way bare jsPermissionFieldName
// output could.
func jsRootPermissionName(p *core.EmiPermission) string {
	return core.ToUpper(jsPermissionIdentifier(p)) + "Permission"
}

// jsSelfPropertyName returns the property name a node's own reserved self
// attribute ("key", "name", "title", or "description") is emitted under: the
// reserved name itself, or that name with a trailing "_" when a child happens to
// have the exact same field name - which would otherwise silently overwrite it,
// since a node's own attributes and its children are flattened onto one object.
// Mirrors selfPermissionFieldName in lib/golang/go-permissions.go.
func jsSelfPropertyName(reserved string, children []*core.EmiPermission) string {
	for _, c := range children {
		if c != nil && jsPermissionFieldName(c) == reserved {
			return reserved + "_"
		}
	}
	return reserved
}

// renderJsPermissionChildren renders one indent level of a permissions forest as
// `<fieldName>: {...},` entries - used recursively for each node's own children
// (the root level instead gets its own `export const <fieldName> = {...}`, one per
// root permission, rendered by renderJsPermissionRootVars).
func renderJsPermissionChildren(w *strings.Builder, permissions []*core.EmiPermission, indent string) {
	for _, p := range permissions {
		if p == nil {
			continue
		}

		fmt.Fprintf(w, "%s%s: {\n", indent, jsPermissionFieldName(p))
		renderJsPermissionNodeBody(w, p, indent+"\t")
		fmt.Fprintf(w, "%s},\n", indent)
	}
}

// renderJsPermissionNodeBody renders a single node's own key/name/title/description
// (its "self" attributes) followed by one property per child - flattened onto the
// same object, so managePosts.key and managePosts.createPost.key are both valid
// paths off the same exported const.
func renderJsPermissionNodeBody(w *strings.Builder, p *core.EmiPermission, indent string) {
	fmt.Fprintf(w, "%s%s: %q,\n", indent, jsSelfPropertyName("key", p.Children), p.EffectiveKey())

	if p.Name != "" {
		fmt.Fprintf(w, "%s%s: %q,\n", indent, jsSelfPropertyName("name", p.Children), p.Name)
	}
	if len(p.Title) > 0 {
		fmt.Fprintf(w, "%s%s: %s,\n", indent, jsSelfPropertyName("title", p.Children), jsStringMapLiteral(p.Title))
	}
	if len(p.Description) > 0 {
		fmt.Fprintf(w, "%s%s: %s,\n", indent, jsSelfPropertyName("description", p.Children), jsStringMapLiteral(p.Description))
	}

	renderJsPermissionChildren(w, p.Children, indent)
}

// renderJsPermissionRootVars renders one exported const per root permission - not a
// single wrapping const for the whole forest - each one the node's own flattened
// key/name/title/description plus one property per child, so consumers can
// navigate e.g. ManagePostsPermission.createPost.key directly. Its name is always
// PascalCase with a trailing "Permission" (see jsRootPermissionName) - not the plain
// camelCase used for a nested child field - so a root is recognizable as a
// permission on sight and can't collide with an unrelated single-letter root. In
// TypeScript the literal is asserted `as const` so its shape is inferred precisely;
// in plain JS (no such assertion exists) it's wrapped in Object.freeze(...) instead,
// as a runtime guard against the tree being mutated after the fact.
//
// Right alongside each root const, a <Name>PermissionList const flattens that same
// root (and everything nested under it) into its own array, by referencing the root
// const's own paths (e.g. ManagePostsPermission.createPost) rather than re-declaring
// separate literals, so it can never drift out of sync with it - mirroring the
// <Name>List vars lib/golang/go-permissions.go generates next to each root var. One
// list per root, not a single list combining every root, so a module with several
// unrelated permission groups doesn't force every caller to filter one combined list
// back down to the group it actually cares about.
func renderJsPermissionRootVars(w *strings.Builder, permissions []*core.EmiPermission, isTypeScript bool) {
	for _, p := range permissions {
		if p == nil {
			continue
		}

		name := jsRootPermissionName(p)
		if isTypeScript {
			fmt.Fprintf(w, "export const %s = {\n", name)
			renderJsPermissionNodeBody(w, p, "\t")
			w.WriteString("} as const;\n\n")
		} else {
			fmt.Fprintf(w, "export const %s = Object.freeze({\n", name)
			renderJsPermissionNodeBody(w, p, "\t")
			w.WriteString("});\n\n")
		}

		paths := collectJsPermissionPaths(p, name)
		if isTypeScript {
			fmt.Fprintf(w, "export const %sList = [\n", name)
			for _, path := range paths {
				fmt.Fprintf(w, "\t%s,\n", path)
			}
			w.WriteString("] as const;\n\n")
		} else {
			fmt.Fprintf(w, "export const %sList = Object.freeze([\n", name)
			for _, path := range paths {
				fmt.Fprintf(w, "\t%s,\n", path)
			}
			w.WriteString("]);\n\n")
		}
	}
}

// collectJsPermissionPaths returns, for a single root permission const accessed
// under varPath, the JS expression that refers to *this* node itself (its own
// key/name/title/description already sit directly on it, unlike Go's separate
// embedded Permission field, so there's no extra suffix for the self entry),
// followed recursively by every descendant's, in the same order
// renderJsPermissionNodeBody emits properties.
func collectJsPermissionPaths(p *core.EmiPermission, varPath string) []string {
	if p == nil {
		return nil
	}

	paths := []string{varPath}
	for _, c := range p.Children {
		if c == nil {
			continue
		}
		paths = append(paths, collectJsPermissionPaths(c, varPath+"."+jsPermissionFieldName(c))...)
	}

	return paths
}

// jsStringMapLiteral renders a map[string]string as a JS object literal, sorted by
// key so the generated file is deterministic across runs.
func jsStringMapLiteral(m map[string]string) string {
	if len(m) == 0 {
		return "{}"
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	b.WriteString("{ ")
	for _, k := range keys {
		fmt.Fprintf(&b, "%q: %q, ", k, m[k])
	}
	b.WriteString("}")

	return b.String()
}

// JsStandalonePermissions renders the module's `permissions:` tree into a single
// standalone JS/TS file:
//
//   - one exported const per root permission - not a single const wrapping the
//     whole forest - each one a flattened, nested object literal: the node's own
//     key/name/title/description sit directly alongside one property per child
//     (named by the child's own identifier), so consumers can navigate e.g.
//     ManagePostsPermission.createPost.key with full IDE autocomplete. Its name is
//     always PascalCase with a trailing "Permission" (see jsRootPermissionName), so
//     a root reads unambiguously as a permission and can't collide with an
//     unrelated single-letter root. In TypeScript each literal is asserted
//     `as const`, so its whole shape is inferred directly - no separate type needs
//     to be hand maintained in step with the tree (mirrors the per-root vars in
//     lib/golang/go-permissions.go, translated to what TS's structural type system
//     does for free instead of Go's explicit anonymous structs)
//   - alongside each root const, a <Name>PermissionList const: that root permission
//     and every one of its descendants, flattened into its own array (e.g.
//     ManagePostsPermissionList, next to ManagePostsPermission) by referencing the
//     root const's own paths (e.g. ManagePostsPermission, ManagePostsPermission.createPost)
//   - not re-declared as separate literals - so it can never drift from it. One
//     list per root, not a single list combining every root (mirrors <Name>List in
//     lib/golang/go-permissions.go). This is the only flattened form generated -
//     there is no separate flat constant per node, since one would just be the same
//     string reachable a second, disconnected way.
//
// core.ResolvePermissionFullKeys must have already run (Emi.Preprocess does this for
// every StringToEmi caller) so every node here already has its key populated.
func JsStandalonePermissions(
	permissions []*core.EmiPermission,
	ctx core.MicroGenContext,
) (*core.CodeChunkCompiled, error) {

	isTypeScript := ctx.HasTag(Typescript)

	rootVars := &strings.Builder{}
	renderJsPermissionRootVars(rootVars, permissions, isTypeScript)

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

	res := &core.CodeChunkCompiled{
		SuggestedFileName: "Permissions",
		ActualScript:      buf.Bytes(),
	}

	res.SuggestedExtension = ".js"
	if isTypeScript {
		res.SuggestedExtension = ".ts"
	}

	return res, nil
}
