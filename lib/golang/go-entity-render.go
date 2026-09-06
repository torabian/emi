package golang

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

//go:embed go-entity.tpl go-entity-shared.tpl
var goEntityTplFS embed.FS

// loadGoEntityTemplate parses the shared entity template definitions first, then the
// main entity template on top of it, mirroring how lib/querypredict loads its .tpl
// files, so the entity template can be edited on disk instead of as an inline Go string.
func loadGoEntityTemplate() (*template.Template, error) {
	t := template.New("go-entity.tpl").Funcs(core.CommonMap)

	if _, err := t.ParseFS(goEntityTplFS, "go-entity-shared.tpl"); err != nil {
		return nil, err
	}

	if _, err := t.ParseFS(goEntityTplFS, "go-entity.tpl"); err != nil {
		return nil, err
	}

	return t, nil
}

// GoEntityRender generates the golang struct for an entity. It reuses
// GoCommonStructGenerator - the same struct generator dtos/actions use - so entities get
// the exact same field-type handling (including "one"/"collection" relation wrappers,
// nested classes, CLI helpers, etc.) for free. Before doing so, ApplyEntityGormTags fills
// in default gorm relation tags on array/array?/collection/collection? fields, and
// returns every struct name (the entity itself, plus each array/array? child struct,
// at any nesting depth) that owns its own id/uniqueId pair. The
// go-entity.tpl/go-entity-shared.tpl templates append entity-only extras after the
// struct GoCommonStructGenerator produced: TableName(), plus one BeforeCreate hook per
// identity struct that assigns uniqueId in Go (via emigo.NewUUIDv4) when the caller
// hasn't already set it - see go-entity-default-fields.go's doc comment for why this
// replaced a DB-level column default.
func GoEntityRender(
	entity *core.Module3Entity,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*StructGenerationParticles, error) {

	identityStructs := ApplyEntityGormTags(entity)

	f := GetCommonFlags(ctx)

	particles, err := GoCommonStructGenerator(entity.Fields, ctx, GoCommonStructContext{
		RootClassName:       entity.GetClassName(),
		RecognizedComplexes: complexes,
		EmiLocation:         f.Emigo,
	})
	if err != nil {
		return nil, err
	}

	t, err := loadGoEntityTemplate()
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"ClassName":       entity.GetClassName(),
		"Table":           entity.Table,
		"IdentityStructs": identityStructs,
	}); err != nil {
		return nil, err
	}

	particles.MainClass.ActualScript = append(particles.MainClass.ActualScript, buf.Bytes()...)

	// The BeforeCreate hooks just appended above always reference emigo.NewUUIDv4() and
	// take a *gorm.DB, regardless of whether entity.Fields happens to use any other
	// emigo-typed shape (Nullable/Collection/etc, which is what would otherwise pull
	// emigo in) or whether Create/Update actions (the only other thing that would
	// otherwise pull gorm in - see go-entity-actions.go - and can be turned off per
	// entity.Features) are even generated for this entity. So both have to be added
	// unconditionally here rather than relying on either's own field/feature-driven
	// dependency collection. CombineGoImports dedupes by import statement text, so this
	// is a no-op if something else already added the same one.
	particles.MainClass.CodeChunkDependensies = append(particles.MainClass.CodeChunkDependensies,
		core.CodeChunkDependency{Location: f.Emigo},
		core.CodeChunkDependency{Location: "gorm.io/gorm"},
	)

	return particles, nil
}
