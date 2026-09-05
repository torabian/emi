package core

import (
	"fmt"
	"regexp"
)

// EmiPermission describes a single node in a module's permission tree. It is purely
// declarative — the actual authorization check is left to whatever app consumes the
// generated code — but the compiler still resolves FullKey for every node so the
// generated Go/JS output exposes one stable, dotted identifier per permission that
// downstream code (ACL/ABAC checks, seed scripts, a permissions picker UI, ...) can
// reference directly, instead of everyone hand-rolling their own string concatenation.
type EmiPermission struct {

	// Name is the programmatic identifier of this permission, used by the compiler to
	// name the generated constant/variable for it (e.g. Name: "name" generates
	// NamePermission). Unlike Key, it does not need to be unique among siblings only -
	// keep it unique across the whole module if you want every permission to get its
	// own generated identifier.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Programmatic identifier of this permission used to name the generated constant/variable (e.g. Name: 'name' generates NamePermission)."`

	// Title is the human readable label of the permission, keyed by locale (e.g. en, fa).
	Title map[string]string `yaml:"title,omitempty" json:"title,omitempty" jsonschema:"description=Human readable label of the permission, keyed by locale (e.g. en, fa)."`

	// Description explains what granting this permission allows, keyed by locale.
	Description map[string]string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Explains what granting this permission allows, keyed by locale."`

	// Key is the short identifier of this permission, unique among its siblings.
	Key string `yaml:"key,omitempty" json:"key,omitempty" jsonschema:"description=Short identifier of this permission, unique among its siblings."`

	// FullKey is the fully qualified, dot separated identifier of this permission. If left
	// empty, the compiler computes it while preprocessing the module, as the parent's
	// FullKey plus this Key (or just this Key for a root permission).
	FullKey string `yaml:"fullKey,omitempty" json:"fullKey,omitempty" jsonschema:"description=Fully qualified dot separated identifier. If left empty it is computed as the parent FullKey plus this Key while the module is preprocessed."`

	// Children are permissions nested under this one, inheriting its FullKey as their prefix.
	Children []*EmiPermission `yaml:"children,omitempty" json:"children,omitempty" jsonschema:"description=Permissions nested under this one, inheriting its FullKey as their prefix."`

	// autoFullKey records whether FullKey was left for the compiler to derive (true),
	// as opposed to set explicitly in yaml (false). Set by ResolvePermissionFullKeys,
	// consumed by EffectiveKey's wildcard-suffix decision. Unexported: internal
	// bookkeeping, not part of the module schema.
	autoFullKey bool

	// autoKey records whether Key was left empty in yaml and normalized from Name
	// (true), as opposed to set explicitly (false). Set by ResolvePermissionFullKeys,
	// consumed by EffectiveKey's wildcard-suffix decision for leaf permissions.
	// Unexported: internal bookkeeping, not part of the module schema.
	autoKey bool
}

// ResolvePermissionFullKeys walks a permission tree in place and fills in FullKey for
// every node that doesn't already define one explicitly, deriving it from the
// parent's already-resolved FullKey and the node's own Key. Call with parentFullKey ""
// for a module's root permissions (see Emi.Preprocess, which does exactly that).
//
// A node left with no Key at all (Key: "" in yaml) gets one normalized from its Name
// first (e.g. Name: "managePosts" -> Key: "managePosts"), so FullKey always has
// something to build on even when the author only bothered to set Name.
//
// Safe to call more than once: a node whose FullKey is already set (explicitly, or
// resolved by an earlier pass) is left untouched.
func ResolvePermissionFullKeys(permissions []*EmiPermission, parentFullKey string) {
	for _, p := range permissions {
		if p == nil {
			continue
		}

		if p.Key == "" && p.Name != "" {
			p.Key = ToLower(p.Name)
			p.autoKey = true
		}

		if p.FullKey == "" {
			p.autoFullKey = true
			if parentFullKey == "" {
				p.FullKey = p.Key
			} else {
				p.FullKey = parentFullKey + "." + p.Key
			}
		}

		ResolvePermissionFullKeys(p.Children, p.FullKey)
	}
}

// permissionIdentifierPattern is what every generator's identifier for a permission
// node - a Go struct field, a JS/TS property, a Kotlin/Swift constant, ... - is
// required to match: a plain ASCII identifier, starting with a letter or
// underscore and containing only letters, digits, and underscores. Nothing else
// (spaces, dots, "@", a leading digit, ...) is guaranteed to compile in every
// target language, and some of it (e.g. "@") would silently compile into a
// different name than the one written, in whichever language happens to tolerate
// it syntactically.
var permissionIdentifierPattern = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

// ValidatePermissionIdentifiers walks a permission tree and rejects any node whose
// identifier - its Name, or its Key when Name is empty, since that's whichever the
// generators actually use to derive a field/const/var name from (see
// permissionIdentifier in lib/golang/go-permissions.go, jsPermissionIdentifier in
// lib/js/js-permissions.go, and their Kotlin/Swift equivalents) - doesn't match
// permissionIdentifierPattern. Call after ResolvePermissionFullKeys, so a node that
// only set Name (no Key) has already had Key normalized from it and isn't flagged
// twice for the same problem.
func ValidatePermissionIdentifiers(permissions []*EmiPermission) error {
	for _, p := range permissions {
		if p == nil {
			continue
		}

		id := p.Name
		if id == "" {
			id = p.Key
		}

		if id != "" && !permissionIdentifierPattern.MatchString(id) {
			return fmt.Errorf(
				"permission %q (fullKey %q) has an invalid identifier %q: a permission's name (or its key, when name is left empty) must start with a letter or underscore and contain only letters, digits, and underscores",
				id, p.FullKey, id,
			)
		}

		if err := ValidatePermissionIdentifiers(p.Children); err != nil {
			return err
		}
	}

	return nil
}

// EffectiveKey returns the single identifier generated code exposes for this
// permission: FullKey normally, or FullKey with a trailing ".*" when its FullKey
// was left for the compiler to derive (rather than set explicitly in yaml) and
// either it groups children - signaling that the key covers this permission and
// everything below it - or it's a leaf whose Key was itself left for the compiler
// to infer from Name, since there's then no author-chosen exact string to honor
// and a bare inferred key is easy to mistake for the whole subtree. A FullKey set
// explicitly in yaml is always used as-is, whether or not the node has children,
// since the author chose that exact string on purpose.
//
// core.ResolvePermissionFullKeys must have already run (Emi.Preprocess does this for
// every StringToEmi/StringToEmiForAction caller) so FullKey is populated.
func (p *EmiPermission) EffectiveKey() string {
	if p == nil {
		return ""
	}
	if p.autoFullKey && (len(p.Children) > 0 || p.autoKey) {
		return p.FullKey + ".*"
	}
	return p.FullKey
}

// Flatten returns this permission and every descendant as a single flat slice, in
// depth-first order (self, then children left to right).
func (x *EmiPermission) Flatten() []*EmiPermission {
	if x == nil {
		return nil
	}

	items := []*EmiPermission{x}
	for _, c := range x.Children {
		items = append(items, c.Flatten()...)
	}

	return items
}

// FlattenPermissions runs Flatten across a whole forest of root permissions - the
// shape module.Permissions actually is.
func FlattenPermissions(permissions []*EmiPermission) []*EmiPermission {
	items := []*EmiPermission{}
	for _, p := range permissions {
		items = append(items, p.Flatten()...)
	}
	return items
}
