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
// in default gorm relation tags on array/array?/collection/collection? fields. The
// go-entity.tpl/go-entity-shared.tpl templates only append entity-only extras (currently
// just TableName()) after the struct GoCommonStructGenerator produced.
func GoEntityRender(
	entity *core.Module3Entity,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*StructGenerationParticles, error) {

	ApplyEntityGormTags(entity)

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
		"ClassName": entity.GetClassName(),
		"Table":     entity.Table,
	}); err != nil {
		return nil, err
	}

	particles.MainClass.ActualScript = append(particles.MainClass.ActualScript, buf.Bytes()...)

	return particles, nil
}
