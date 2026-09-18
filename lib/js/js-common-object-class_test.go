package js

import "testing"

// TestEntityTargetToCodeChunk covers the fix for relation fields declared as
// `target: XEntity` (the documented convention for pointing at another entity, e.g.
// `target: PassportEntity`): the JS/TS compiler never emits a literal "XEntity.ts" file
// for any entity - only Dto/OptionalDto/action classes - so importing target as a
// same-named sibling file (castDtoNameToCodeChunk's normal behavior) is always a
// dangling import for these. entityTargetToCodeChunk must instead resolve it to the
// real, generated "XDto" module, aliased back to "XEntity" so both the TS type position
// and the runtime `new XEntity(...)` constructor call (both of which use field.Target
// verbatim elsewhere in this package) keep working without a hand-written shim file.
//
// All TypeScript-mode cases below also assert the plain-type companion import
// (<target>Type) travels alongside the class - see tsTargetTypeName,
// js-common-object-types.go - as a type-only specifier ("type " prefix,
// CombineImportsJsWorld), required once a consuming tsconfig turns on
// verbatimModuleSyntax. TestEntityTargetToCodeChunkPlainJS covers the other
// half: that companion must be entirely absent from plain JS output, not just
// unprefixed - `import { type X }` isn't valid JavaScript grammar at all, so
// even a bundler merely parsing a .js file as plain JS rejects it outright
// (a real, reproduced regression - see that test's own doc comment).
func TestEntityTargetToCodeChunk(t *testing.T) {
	t.Run("Entity-suffixed target aliases the sibling Dto, not itself", func(t *testing.T) {
		chunk := entityTargetToCodeChunk("PassportEntity", "", true)

		if len(chunk.CodeChunkDependensies) != 1 {
			t.Fatalf("expected exactly one dependency, got %d: %+v", len(chunk.CodeChunkDependensies), chunk.CodeChunkDependensies)
		}

		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "./PassportDto" {
			t.Errorf("expected the import location to be the sibling Dto module \"./PassportDto\", got %q (would dangle on a nonexistent \"./PassportEntity\")", dep.Location)
		}
		if len(dep.Objects) != 2 ||
			dep.Objects[0] != "PassportDto as PassportEntity" ||
			dep.Objects[1] != "type PassportDtoType as PassportEntityType" {
			t.Errorf(`expected both "PassportDto as PassportEntity" and "type PassportDtoType as PassportEntityType", got %v`, dep.Objects)
		}
	})

	t.Run("multi-word entity name", func(t *testing.T) {
		chunk := entityTargetToCodeChunk("UserWorkspaceEntity", "", true)
		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "./UserWorkspaceDto" || dep.Objects[0] != "UserWorkspaceDto as UserWorkspaceEntity" {
			t.Errorf("expected UserWorkspaceDto aliased as UserWorkspaceEntity from ./UserWorkspaceDto, got %+v", dep)
		}
		if len(dep.Objects) != 2 || dep.Objects[1] != "type UserWorkspaceDtoType as UserWorkspaceEntityType" {
			t.Errorf("expected the Type companion aliased alongside it, got %+v", dep)
		}
	})

	t.Run("non-Entity-suffixed target falls back to the normal same-named-file resolution", func(t *testing.T) {
		chunk := entityTargetToCodeChunk("SomeActionResDto", "", true)
		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "./SomeActionResDto" {
			t.Errorf("expected the untouched castDtoNameToCodeChunk location behavior for a non-Entity target, got %+v", dep)
		}
		// No alias needed here (dtoName already equals the class name), but
		// the Type companion still has to be imported alongside it.
		if len(dep.Objects) != 2 || dep.Objects[0] != "SomeActionResDto" || dep.Objects[1] != "type SomeActionResDtoType" {
			t.Errorf("expected SomeActionResDto plus its type-only Type companion, got %+v", dep)
		}
	})

	t.Run("literal \"Entity\" with nothing to strip falls back too", func(t *testing.T) {
		chunk := entityTargetToCodeChunk("Entity", "", true)
		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "./Entity" || dep.Objects[0] != "Entity" || len(dep.Objects) != 2 || dep.Objects[1] != "type EntityType" {
			t.Errorf("expected the untouched castDtoNameToCodeChunk behavior (plus its type-only Type companion) for the literal \"Entity\", got %+v", dep)
		}
	})

	// A cross-module target (field.JsProvider set - e.g. a standalone dto
	// compiled via js:dto:class, referencing an entity from a module it was
	// never generated alongside) must still alias back to the short target
	// name ("WalletEntity"), never to jsProvider's own full path - regression
	// test for exactly that bug (the alias briefly became the literal string
	// "@fireback/wallet/sdk/WalletEntity" before Target/JsProvider were
	// tracked separately in EntityTargetRef).
	t.Run("jsProvider changes the import location, never the alias", func(t *testing.T) {
		chunk := entityTargetToCodeChunk("WalletEntity", "@fireback/wallet/sdk/WalletEntity", true)
		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "@fireback/wallet/sdk/WalletDto" {
			t.Errorf("expected the import location to come from jsProvider (\"@fireback/wallet/sdk/WalletDto\"), got %q", dep.Location)
		}
		if len(dep.Objects) != 2 ||
			dep.Objects[0] != "WalletDto as WalletEntity" ||
			dep.Objects[1] != "type WalletDtoType as WalletEntityType" {
			t.Errorf(`expected "WalletDto as WalletEntity" and "type WalletDtoType as WalletEntityType" (the short target name, not jsProvider's path), got %v`, dep.Objects)
		}
	})
}

// TestEntityTargetToCodeChunkPlainJS is the regression test for a real bug:
// isTypeScript: false used to still add the "type <X>Type" object (the fix
// only made the specifier type-only-*prefixed*, not conditional on the
// output language at all). Under plain JS output that broke a real consumer
// at the parser level, not just tsc: `import { User, type UserType } from
// "./User"` inside a .js file made Rollup/esbuild fail outright ("Parse
// failure: Expected ',', got 'UserType'") - `type` as an import-specifier
// modifier is TypeScript-only grammar, invalid even to a bundler merely
// parsing the file as plain JavaScript, not something that only tsc would
// catch. Every branch entityTargetToCodeChunk/castDtoNameToCodeChunk can
// take must therefore drop the Type companion entirely (not just its
// prefix) when isTypeScript is false - there is no plain type to import in
// JS output anyway, since a JS consumer only ever gets the class.
func TestEntityTargetToCodeChunkPlainJS(t *testing.T) {
	cases := []struct {
		name       string
		target     string
		jsProvider string
	}{
		{"Entity-suffixed, same module", "PassportEntity", ""},
		{"Entity-suffixed, cross-module via jsProvider", "WalletEntity", "@fireback/wallet/sdk/WalletEntity"},
		{"non-Entity-suffixed (plain Dto target)", "SomeActionResDto", ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			chunk := entityTargetToCodeChunk(c.target, c.jsProvider, false)
			dep := chunk.CodeChunkDependensies[0]

			if len(dep.Objects) != 1 {
				t.Fatalf("expected exactly the class import, no Type companion, in plain JS mode, got %v", dep.Objects)
			}
			for _, obj := range dep.Objects {
				if obj == "type" || len(obj) >= 5 && obj[:5] == "type " {
					t.Fatalf("did not expect any type-only specifier in plain JS output, got %q", obj)
				}
			}
		})
	}
}
