package core

// This file turns each entity's generated operations (Create/Update/Get/Browse/
// AwareDelete - see lib/golang/go-entity-actions.go and go-entity-delete.go) into
// regular, portable EmiAction definitions appended to m.Actions, so they're exposed via
// the exact same mechanism as any hand-declared action - HTTP route, gin/cli/http
// envelopes, request/response types - through the existing action codegen, completely
// unmodified. This is deliberately the plainest possible default setup: no query-string
// binding beyond what a feature strictly needs, no permissions/headers/custom
// envelopes - a caller wanting more can always declare their own EmiAction by hand
// instead (or in addition; see preprocessEntityActions's dedup-by-name below).
//
// The stub these produce still needs a real Go function body wired up by hand (see the
// "quick function implementation" comment GoActionRender emits) - this only gets the
// action's shape (name/method/url/in/out) generated automatically, it does not
// automatically call {Entity}CreateFn/{Entity}UpdateFn/etc for you.

// entityActionName is the {Name} in {Name}Create/{Name}Update/etc - entity.Name is
// already lowerCamel, matching every other Emi action's naming convention.
func entityActionName(entity *Module3Entity, suffix string) string {
	return entity.Name + suffix
}

// Returns only the affix for action name, so user doesn't have to type constantly.
func entityActionCliName(entity *Module3Entity, word string) string {
	return word
}

// entityActionCliShort is the action's CLI alias (EmiAction.CliShort -> cmd.Aliases,
// see lib/golang/go-action-cli-render.go) - just the bare per-operation short code
// ("c"/"u"/"g"/"b"/"d"/"dp"), matching CliName ("create"/"update"/...). Entity actions
// are registered as subcommands nested under their own entity's CLI node (e.g. "msg
// email provider create"/"c"), not in one flat/global command list, so there's no
// cross-entity collision to guard against by prefixing with entity.Name here.
func entityActionCliShort(entity *Module3Entity, short string) string {
	return short
}

// entityActionResponseEnvelope is the envelope every entity action's Out uses -
// "GResponse", the same google-json-style-guide envelope class hand-declared actions
// elsewhere in this repo opt into by hand (see lib/js/js-envelopes and
// examples/fullstack/definitions.emi.yml) - wraps a single item as data.item, a list as
// data.items, transparently, on the JS side (see lib/js/js-action-class-realms.go's
// GetResponseEnvelopeClass handling; Go/Kotlin/Swift/openapi don't consume Envelope at
// all today, so this is a no-op for them).
const entityActionResponseEnvelope = "GResponse"

func buildEntityCreateAction(entity *Module3Entity) *EmiAction {
	dtoClassName := BuildEntityDto(entity).GetClassName()
	return &EmiAction{
		Name:        entityActionName(entity, "Create"),
		CliName:     entityActionCliName(entity, "create"),
		CliShort:    entityActionCliShort(entity, "c"),
		Description: "Creates a new \"" + entity.Name + "\" row.",
		Method:      "post",
		Url:         "/" + entity.Name,
		In:          &EmiActionBody{Dto: dtoClassName},
		Out:         &EmiActionBody{Dto: dtoClassName, Envelope: entityActionResponseEnvelope},
	}
}

func buildEntityUpdateAction(entity *Module3Entity) *EmiAction {
	return &EmiAction{
		Name:        entityActionName(entity, "Update"),
		CliName:     entityActionCliName(entity, "update"),
		CliShort:    entityActionCliShort(entity, "u"),
		Description: "Applies a partial update to a \"" + entity.Name + "\" row by uniqueId.",
		Method:      "patch",
		Url:         "/" + entity.Name + "/:uniqueId string",
		In:          &EmiActionBody{Dto: BuildEntityOptionalDto(entity).GetClassName()},
		Out:         &EmiActionBody{Dto: BuildEntityDto(entity).GetClassName(), Envelope: entityActionResponseEnvelope},
	}
}

func buildEntityGetAction(entity *Module3Entity) *EmiAction {
	return &EmiAction{
		Name:        entityActionName(entity, "Get"),
		CliName:     entityActionCliName(entity, "get"),
		CliShort:    entityActionCliShort(entity, "g"),
		Description: "Looks up a single \"" + entity.Name + "\" row by uniqueId.",
		Method:      "get",
		Url:         "/" + entity.Name + "/:uniqueId string",
		Out:         &EmiActionBody{Dto: BuildEntityDto(entity).GetClassName(), Envelope: entityActionResponseEnvelope},
	}
}

// buildEntityBrowseAction renders the entity's list/filter/page operation. Named
// "browse" (not "query") because it's deliberately the simple case - a straightforward
// filtered/sorted/paged list; "query" is reserved for a more advanced mechanism later.
//
// items uses the entity's *optional* dto (BuildEntityOptionalDto), not the strict one
// (BuildEntityDto) a single Create/Get/Update response uses - a browse/list result may
// only select or return some fields, not the complete row every time, so its items need
// to be able to say "this field wasn't included" the same way Update's input can say
// "this field wasn't set".
// BuildEntityBrowseAction is buildEntityBrowseAction exposed for callers outside this
// package (namely lib/golang's GoEntityActionsRender) that need the exact same synthetic
// EmiAction shape even when features.browseAction (or the master features.actions)
// suppresses it from ever being added to m.Actions - {Entity}BrowseFn's own generated Go
// code still needs the qs struct/parsing helpers this action's Query fields drive (see
// GoActionQueryParams), regardless of whether it's exposed as a portable action.
func BuildEntityBrowseAction(entity *Module3Entity) *EmiAction {
	return buildEntityBrowseAction(entity)
}

func buildEntityBrowseAction(entity *Module3Entity) *EmiAction {
	return &EmiAction{
		Name:        entityActionName(entity, "Browse"),
		CliName:     entityActionCliName(entity, "browse"),
		CliShort:    entityActionCliShort(entity, "b"),
		Description: "Returns \"" + entity.Name + "\" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).",
		Method:      "get",
		Url:         "/" + entity.Name + "/browse",
		Query: []*EmiQueryField{
			{Name: "filter", Type: FieldTypeString},
			{Name: "sort", Type: FieldTypeString},
			{Name: "startIndex", Type: FieldTypeInt},
			{Name: "itemsPerPage", Type: FieldTypeInt},
			// Resumes a previous page - always exactly what a prior call's own
			// "cursor" response field returned (see emigorm.ApplyQueryCursor/
			// BuildQueryCursor); empty means "start from the beginning".
			{Name: "cursor", Type: FieldTypeString},
		},
		// Out is the entity's own optional dto (BuildEntityOptionalDto) - not a
		// hand-rolled {items, total, cursor} wrapper - so it's the exact same type
		// Update's request body already is, referenced by name like any other dto
		// response (see Create/Get/Update's Out above). Envelope: GResponse is what
		// actually carries the list-ness (data.items, see lib/js/js-envelopes) and is
		// where a real total/cursor belongs once a caller wires an actual handler up -
		// {Entity}BrowseFn itself still returns the full ([]*Entity, *emigo.QueryResultMeta,
		// error) shape (see lib/golang/go-entity-actions.go's buildBrowseFn), this is
		// only about what the portable action *declares*.
		//
		// Note: emigorm.QueryDSL.Scope (a second, handler-enforced filter - e.g.
		// workspace isolation) deliberately has no qs field here at all - it's
		// json:"-", only ever set by a caller's own handler code, never from a request.
		Out: &EmiActionBody{Dto: BuildEntityOptionalDto(entity).GetClassName(), Envelope: entityActionResponseEnvelope},
	}
}

func buildEntityAwareDeletePreviewAction(entity *Module3Entity) *EmiAction {
	return &EmiAction{
		Name:        entityActionName(entity, "AwareDeletePreview"),
		CliName:     entityActionCliName(entity, "delete-preview"),
		CliShort:    entityActionCliShort(entity, "dp"),
		Description: "Reports what deleting the given \"" + entity.Name + "\" uniqueIds would affect, without deleting anything.",
		Method:      "get",
		Url:         "/" + entity.Name + "/delete-preview",
		Query: []*EmiQueryField{
			{Name: "uniqueIds", Type: FieldTypeSlice, Primitive: "string"},
		},
		Out: &EmiActionBody{
			Envelope: entityActionResponseEnvelope,
			Fields: []*EmiField{
				{Name: "message", Type: FieldTypeString},
				{Name: "affected", Type: FieldTypeArray, Fields: []*EmiField{
					{Name: "relation", Type: FieldTypeString},
					{Name: "count", Type: FieldTypeInt64},
				}},
			},
		},
	}
}

func buildEntityAwareDeleteAction(entity *Module3Entity) *EmiAction {
	return &EmiAction{
		Name:        entityActionName(entity, "AwareDelete"),
		CliName:     entityActionCliName(entity, "delete"),
		CliShort:    entityActionCliShort(entity, "d"),
		Description: "Deletes the given \"" + entity.Name + "\" uniqueIds, along with everything " + entityActionName(entity, "AwareDeletePreview") + " reports.",
		Method:      "post",
		Url:         "/" + entity.Name + "/delete",
		In: &EmiActionBody{Fields: []*EmiField{
			{Name: "uniqueIds", Type: FieldTypeSlice, Primitive: "string"},
		}},
	}
}

// preprocessEntityActions synthesizes each enabled entity operation (see
// Module3EntityFeatures) into a plain EmiAction and appends it to m.Actions.
// Idempotent, and yields to any action a module already declares under the same name -
// same convention as preprocessEntityOptionalDtos.
func (m *Emi) preprocessEntityActions() {
	existing := make(map[string]bool, len(m.Actions)+len(m.Entities)*4)
	for _, a := range m.Actions {
		if a != nil {
			existing[a.Name] = true
		}
	}

	add := func(a *EmiAction) {
		if a == nil || existing[a.Name] {
			return
		}
		m.Actions = append(m.Actions, a)
		existing[a.Name] = true
	}

	for _, e := range m.Entities {
		if e == nil || e.Name == "" {
			continue
		}
		if e.Features.CreateEnabled() && e.Features.CreateActionEnabled() {
			add(buildEntityCreateAction(e))
		}
		if e.Features.UpdateEnabled() && e.Features.UpdateActionEnabled() {
			add(buildEntityUpdateAction(e))
		}
		if e.Features.GetEnabled() && e.Features.GetActionEnabled() {
			add(buildEntityGetAction(e))
		}
		if e.Features.BrowseEnabled() && e.Features.BrowseActionEnabled() {
			add(buildEntityBrowseAction(e))
		}
		if e.Features.DeleteEnabled() && e.Features.DeleteActionEnabled() {
			add(buildEntityAwareDeletePreviewAction(e))
			add(buildEntityAwareDeleteAction(e))
		}
	}
}

// PreprocessEntityActions is preprocessEntityActions exposed as a core.PreprocessHook.
// Same opt-in convention as PreprocessEntityOptionalDtos - core never registers this
// itself; a compiler backend that actually wants entities exposed as EmiActions
// (currently just lib/golang, from GetGolangPublicActions) registers it explicitly via
// RegisterPreprocessHook.
func PreprocessEntityActions(m *Emi) error {
	m.preprocessEntityActions()
	return nil
}
