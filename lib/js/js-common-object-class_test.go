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
func TestEntityTargetToCodeChunk(t *testing.T) {
	t.Run("Entity-suffixed target aliases the sibling Dto, not itself", func(t *testing.T) {
		chunk := entityTargetToCodeChunk("PassportEntity", "")

		if len(chunk.CodeChunkDependensies) != 1 {
			t.Fatalf("expected exactly one dependency, got %d: %+v", len(chunk.CodeChunkDependensies), chunk.CodeChunkDependensies)
		}

		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "./PassportDto" {
			t.Errorf("expected the import location to be the sibling Dto module \"./PassportDto\", got %q (would dangle on a nonexistent \"./PassportEntity\")", dep.Location)
		}
		if len(dep.Objects) != 1 || dep.Objects[0] != "PassportDto as PassportEntity" {
			t.Errorf(`expected the aliased import "PassportDto as PassportEntity", got %v`, dep.Objects)
		}
	})

	t.Run("multi-word entity name", func(t *testing.T) {
		chunk := entityTargetToCodeChunk("UserWorkspaceEntity", "")
		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "./UserWorkspaceDto" || dep.Objects[0] != "UserWorkspaceDto as UserWorkspaceEntity" {
			t.Errorf("expected UserWorkspaceDto aliased as UserWorkspaceEntity from ./UserWorkspaceDto, got %+v", dep)
		}
	})

	t.Run("non-Entity-suffixed target falls back to the normal same-named-file resolution", func(t *testing.T) {
		chunk := entityTargetToCodeChunk("SomeActionResDto", "")
		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "./SomeActionResDto" || dep.Objects[0] != "SomeActionResDto" {
			t.Errorf("expected the untouched castDtoNameToCodeChunk behavior for a non-Entity target, got %+v", dep)
		}
	})

	t.Run("literal \"Entity\" with nothing to strip falls back too", func(t *testing.T) {
		chunk := entityTargetToCodeChunk("Entity", "")
		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "./Entity" || dep.Objects[0] != "Entity" {
			t.Errorf("expected the untouched castDtoNameToCodeChunk behavior for the literal \"Entity\", got %+v", dep)
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
		chunk := entityTargetToCodeChunk("WalletEntity", "@fireback/wallet/sdk/WalletEntity")
		dep := chunk.CodeChunkDependensies[0]
		if dep.Location != "@fireback/wallet/sdk/WalletDto" {
			t.Errorf("expected the import location to come from jsProvider (\"@fireback/wallet/sdk/WalletDto\"), got %q", dep.Location)
		}
		if len(dep.Objects) != 1 || dep.Objects[0] != "WalletDto as WalletEntity" {
			t.Errorf(`expected the aliased import "WalletDto as WalletEntity" (the short target name, not jsProvider's path), got %v`, dep.Objects)
		}
	})
}
