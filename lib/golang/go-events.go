package golang

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// GoEventsGenerate renders the module's `events:` list into a single Events.go: one
// exported var per event (emigo.Event), named PascalCase(Key) + "Event" - e.g.
// Key: "userCreated" -> UserCreatedEvent - so every event is reachable directly by
// name, mirroring how GoPermissionsGenerate exposes each permission root. Alongside
// each var, a <EventName>Payload type is generated - a real struct when the event
// declares payload fields (compiled through the same Common struct builder actions'
// in/out bodies use), an alias to an existing dto when it references one by name, or
// an alias to interface{} when no payload was declared at all - so there is always
// exactly one identifier to reach for regardless of how (or whether) the event typed
// its payload. A New<EventName>(payload <EventName>Payload) emigo.Event constructor is
// generated next to each var too - it copies the var and sets .Payload, so callers
// firing the event don't have to repeat that copy-and-set by hand at every call site.
// AllEventsList collects every event into a single []emigo.Event, for anything that
// wants to walk them all at once (a notification-settings catalog, seeding a
// NotificationType table, ...).
//
// core.ValidateEventIdentifiers must already have run (Emi.Preprocess does this for
// every StringToEmi/StringToEmiForAction caller) so every event's Key is guaranteed
// unique and a valid identifier basis. Returns (nil, nil) when the module declares no
// events.
func GoEventsGenerate(
	events []*core.EmiEvent,
	ctx core.MicroGenContext,
	emigoImportPath string,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {

	if len(events) == 0 {
		return nil, nil
	}

	vars := &strings.Builder{}
	deps, err := renderEventVars(vars, events, ctx, complexes)
	if err != nil {
		return nil, err
	}

	const tmpl = `/**
* Event keys generated from the module's events list.
*/

{{ .vars }}`

	t := template.Must(template.New("events").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"vars": vars.String(),
	}); err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		SuggestedFileName:  "Events",
		SuggestedExtension: ".go",
		ActualScript:       buf.Bytes(),
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: emigoImportPath},
		},
	}
	res.CodeChunkDependensies = append(res.CodeChunkDependensies, deps...)

	return res, nil
}

// goEventVarName turns an event's Key into its exported var name: UpperCamelCase
// (via core.EventGoVarName, not a bare first-letter capitalisation - "user_created"
// and "user-created" both come out "UserCreated" the same as "userCreated" would)
// with a trailing "Event" always appended (e.g. "userCreated" -> "UserCreatedEvent"),
// mirroring goRootPermissionName in go-permissions.go so an event reads unambiguously
// as one at a glance and can't collide with an unrelated identifier.
// core.ValidateEventIdentifiers already rejects two events whose Key normalises to
// the same name, so this can't collide across events within a module.
func goEventVarName(e *core.EmiEvent) string {
	return core.EventGoVarName(e) + "Event"
}

// goEventPermissionSetsLiteral renders Permissions ([]core.EmiEventPermission) as a Go
// [][]string composite literal - one inner slice per line (pulled from each entry's
// With), so a multi-combination event stays readable once gofmt runs over the
// generated file. emigo.Event.Permissions stays [][]string at runtime - With is only
// the yaml-level spelling of each AND'd set.
func goEventPermissionSetsLiteral(sets []core.EmiEventPermission) string {
	if len(sets) == 0 {
		return "nil"
	}

	var b strings.Builder
	b.WriteString("[][]string{\n")
	for _, set := range sets {
		b.WriteString("\t{")
		for i, key := range set.With {
			if i > 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(&b, "%q", key)
		}
		b.WriteString("},\n")
	}
	b.WriteString("}")

	return b.String()
}

// goEventLiteral renders a single event as an emigo.Event{...} composite literal.
// Each field gets its own line rather than one long comma-separated line - same
// reasoning as goPermissionLiteral: Name/Description are often multi-entry locale
// maps, and gofmt only reflows a composite literal onto multiple lines when the
// source already breaks it that way. Payload is deliberately left unset here - it's
// runtime data, not metadata; see goEventPayloadTypeName for the generated
// <EventName>Payload type callers assign/cast it through.
func goEventLiteral(e *core.EmiEvent) string {
	return fmt.Sprintf(
		"emigo.Event{\nKey: %q,\nName: %s,\nDescription: %s,\nPermissions: %s,\n}",
		e.Key, goStringMapLiteral(e.Name), goStringMapLiteral(e.Description), goEventPermissionSetsLiteral(e.Permissions),
	)
}

// goEventPayloadTypeName resolves the Go type an event's payload should be reached
// through and, when the payload is declared inline via fields, the struct source to
// render for it. Three cases, same precedence EmiAction's request/response bodies
// follow (see GoActionRealms): fields compiled through the Common struct builder into
// a real struct, an existing dto referenced by name, or - when the event declares no
// payload at all - a plain interface{}. The returned deps are only ever populated in
// the fields case, mirroring how RequestClass's dependencies are collected there.
func goEventPayloadTypeName(
	e *core.EmiEvent,
	payloadClassName string,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (typeName string, structScript string, deps []core.CodeChunkDependency, err error) {
	switch {
	case e.HasPayloadFields():
		fields, genErr := GoCommonStructGenerator(e.GetPayloadFields(), ctx, GoCommonStructContext{
			RootClassName:       payloadClassName,
			RecognizedComplexes: complexes,
		})
		if genErr != nil {
			return "", "", nil, genErr
		}
		if fields.MainClass == nil {
			return "interface{}", "", nil, nil
		}
		return payloadClassName, string(fields.MainClass.ActualScript), fields.MainClass.CodeChunkDependensies, nil

	case e.HasPayloadDto():
		typeName := e.GetPayloadDto()
		chunk := castDtoNameToCodeChunk(typeName)
		if token := core.FindTokenByName(chunk.Tokens, TOKEN_ROOT_CLASS); token != nil {
			typeName = token.Value
		}
		return typeName, "", nil, nil

	default:
		return "interface{}", "", nil, nil
	}
}

// renderEventVars renders, for each event, one exported var holding its emigo.Event
// value - not a single wrapping slice - so it can be referenced directly by name
// (e.g. UserCreatedEvent.Key). Alongside it, <EventName>Payload is always declared:
// a real struct when the event compiled payload fields, otherwise a type alias to the
// referenced dto or to interface{} - always exactly one name to reach for regardless
// of how the event typed (or didn't type) its payload. A New<EventName> constructor is
// also emitted right after the var - it copies it and sets .Payload to the given
// <EventName>Payload, so firing the event is one call instead of a copy-and-assign at
// every call site. AllEventsList is emitted alongside them, referencing each var
// rather than re-declaring separate literals, so it can never drift out of sync with
// them - same pattern as <Name>PermissionList in go-permissions.go.
func renderEventVars(
	w *strings.Builder,
	events []*core.EmiEvent,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) ([]core.CodeChunkDependency, error) {
	names := make([]string, 0, len(events))
	var deps []core.CodeChunkDependency

	for _, e := range events {
		if e == nil {
			continue
		}

		name := goEventVarName(e)
		names = append(names, name)

		payloadClassName := name + "Payload"
		typeName, structScript, payloadDeps, err := goEventPayloadTypeName(e, payloadClassName, ctx, complexes)
		if err != nil {
			return nil, err
		}
		deps = append(deps, payloadDeps...)

		if structScript != "" {
			w.WriteString(structScript)
			w.WriteString("\n")
		} else {
			fmt.Fprintf(w, "// %s is the payload type for the %q event.\n", payloadClassName, e.Key)
			fmt.Fprintf(w, "type %s = %s\n\n", payloadClassName, typeName)
		}

		fmt.Fprintf(w, "// %s mirrors the %q event declared in this module's events list.\n", name, e.Key)
		fmt.Fprintf(w, "var %s = %s\n\n", name, goEventLiteral(e))

		ctorName := "New" + name
		fmt.Fprintf(w, "// %s returns a copy of %s with Payload set to payload - the usual way to\n", ctorName, name)
		fmt.Fprintf(w, "// build the event instance a caller actually fires, without repeating its\n")
		fmt.Fprintf(w, "// Key/Name/Description/Permissions metadata at every call site.\n")
		fmt.Fprintf(w, "func %s(payload %s) emigo.Event {\n", ctorName, payloadClassName)
		fmt.Fprintf(w, "\tevt := %s\n", name)
		fmt.Fprintf(w, "\tevt.Payload = payload\n")
		fmt.Fprintf(w, "\treturn evt\n")
		fmt.Fprintf(w, "}\n\n")
	}

	w.WriteString("// AllEventsList collects every event declared by this module, for anything that\n")
	w.WriteString("// wants to walk them all at once (a notification-settings catalog, seeding a\n")
	w.WriteString("// NotificationType table, ...) - referencing each var above rather than\n")
	w.WriteString("// re-declaring separate literals, so it can never drift out of sync with them.\n")
	w.WriteString("var AllEventsList = []emigo.Event{\n")
	for _, name := range names {
		fmt.Fprintf(w, "\t%s,\n", name)
	}
	w.WriteString("}\n")

	return deps, nil
}
