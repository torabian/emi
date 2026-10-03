package core

import (
	"fmt"
)

// EmiEvent describes a single fact a module can emit through its event system - e.g.
// "a new user was created in a workspace". It is purely declarative, like
// EmiPermission: the actual firing/dispatch is left to whatever app consumes the
// generated code, but the compiler still validates the definition and turns it into
// a typed catalog entry per target language (see lib/golang/go-events.go for Go).
//
// Unlike EmiPermission, events are not a tree - there is no Children/FullKey
// derivation, just a flat, uniquely-keyed list per module.
type EmiEvent struct {

	// Key is the unique, machine identifier of this event, used both as the wire-level
	// name (what actually gets sent as the event's Name at fire time) and to derive
	// the generated variable name for it (e.g. Key: "userCreated" generates
	// UserCreatedEvent). Must be unique across the whole module.
	Key string `yaml:"key,omitempty" json:"key,omitempty" jsonschema:"description=Unique machine identifier of this event, used as both the wire-level name and to derive the generated variable name (e.g. 'userCreated' generates UserCreatedEvent)."`

	// Name is the human readable, localized label of the event, keyed by locale (e.g.
	// en, fa) - what a notification-settings screen would show a user, not a
	// programmatic identifier (that's Key).
	Name map[string]string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Human readable label of the event, keyed by locale (e.g. en, fa)."`

	// Description explains, for a human, what this event represents, keyed by locale.
	Description map[string]string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Explains what this event represents, keyed by locale."`

	// Permissions is a list of permission combinations that grant visibility into this
	// event: the outer list is OR'd together, and each entry's With list of permission
	// full keys is AND'd - e.g. [{with: [A, B]}, {with: [C, D]}] means "(A and B) or
	// (C and D)". An identity needs to satisfy every key within at least one entry's
	// With to be considered allowed to know this event happened. Leave empty for an
	// event nobody needs a specific permission to know about.
	Permissions []EmiEventPermission `yaml:"permissions,omitempty" json:"permissions,omitempty" jsonschema:"description=Permission combinations that grant visibility into this event: the outer list is OR'd together, each entry's With list is AND'd together (e.g. [{with: [A, B]}, {with: [C, D]}] means (A and B) or (C and D))."`

	// Payload defines the shape of data this event carries when it fires. Compiled
	// through the same Common struct builder as an action's in/out body (see
	// EmiActionBody): Fields becomes a typed DTO nested under the event, Dto
	// references an existing shape by name. Left empty, the event's payload is
	// untyped (interface{}) in the generated code.
	Payload *EmiActionBody `yaml:"payload,omitempty" json:"payload,omitempty" jsonschema:"description=Shape of data this event carries when it fires. Compiled the same way an action's request/response body is: fields becomes a typed DTO, dto references an existing shape. Untyped (interface{}) in generated code if omitted."`

	// Params defines the parameters that scope or qualify this event - e.g. the
	// workspaceId a userCreated event fired in - as opposed to Payload, which is the
	// data the event carries. Compiled through the same Common struct builder as
	// Payload (and an action's in/out body) in Go, JS/TS, Kotlin and Swift: Fields
	// becomes a typed DTO named <EventName>Params, Dto makes <EventName>Params an alias
	// of an existing dto (use one of them, like an action's in/out).
	Params *EmiActionBody `yaml:"params,omitempty" json:"params,omitempty" jsonschema:"description=Parameters that scope or qualify this event (e.g. the workspaceId it fired in), as opposed to the payload it carries. Compiled with the common struct generator in go, js/ts, kotlin and swift: fields becomes a typed DTO named <EventName>Params, dto references an existing shape."`
}

// EmiEventPermission is one AND'd permission combination: With lists every permission
// full key that must all be held together for this combination to grant visibility
// into the event it belongs to. EmiEvent.Permissions is a list of these, OR'd
// together - see its doc comment.
type EmiEventPermission struct {
	With []string `yaml:"with,omitempty" json:"with,omitempty" jsonschema:"description=Permission full keys that must all be held together (AND'd) for this combination to grant visibility."`
}

// HasPayload reports whether this event declared a payload shape at all.
func (x EmiEvent) HasPayload() bool {
	return x.Payload != nil
}

// HasPayloadFields reports whether the event's payload is declared inline via fields,
// to be compiled into a typed DTO - mirrors EmiAction.HasRequestFields.
func (x EmiEvent) HasPayloadFields() bool {
	return x.HasPayload() && len(x.Payload.Fields) > 0
}

// GetPayloadFields returns the event's payload fields, or an empty slice if none were
// declared - mirrors EmiAction.GetRequestFields.
func (x EmiEvent) GetPayloadFields() []*EmiField {
	if !x.HasPayloadFields() {
		return []*EmiField{}
	}
	return x.Payload.Fields
}

// HasPayloadDto reports whether the event's payload references an existing dto by
// name rather than declaring fields inline - mirrors EmiAction.HasRequestDto.
func (x EmiEvent) HasPayloadDto() bool {
	return x.HasPayload() && x.Payload.Dto != ""
}

// GetPayloadDto returns the name of the dto the event's payload references.
func (x EmiEvent) GetPayloadDto() string {
	return x.Payload.Dto
}

// HasParams reports whether this event declared a params shape at all.
func (x EmiEvent) HasParams() bool {
	return x.Params != nil
}

// HasParamsFields reports whether the event's params are declared inline via fields,
// to be compiled into a typed DTO - mirrors HasPayloadFields.
func (x EmiEvent) HasParamsFields() bool {
	return x.HasParams() && len(x.Params.Fields) > 0
}

// GetParamsFields returns the event's params fields, or an empty slice if none were
// declared - mirrors GetPayloadFields.
func (x EmiEvent) GetParamsFields() []*EmiField {
	if !x.HasParamsFields() {
		return []*EmiField{}
	}
	return x.Params.Fields
}

// HasParamsDto reports whether the event's params reference an existing dto by name
// rather than declaring fields inline - mirrors HasPayloadDto.
func (x EmiEvent) HasParamsDto() bool {
	return x.HasParams() && x.Params.Dto != ""
}

// GetParamsDto returns the name of the dto the event's params reference.
func (x EmiEvent) GetParamsDto() string {
	return x.Params.Dto
}

// EventParamsClassName is the generated class/struct name of an event's inline params
// shape in every target language (e.g. Key: "userCreated" -> UserCreatedEventParams),
// derived from EventGoVarName so all backends agree on it.
func EventParamsClassName(e *EmiEvent) string {
	return EventGoVarName(e) + "EventParams"
}

// EventGoVarName is the single source of truth for an event's generated variable
// name basis: Key run through NormaliseKey (the same UpperCamelCase normalisation
// EmiEnumInline.GetKey uses) rather than a bare first-letter capitalisation, so a
// key like "user_created" or "user-created" comes out "UserCreated" exactly like
// "userCreated" would - every generator backend derives its event identifier from
// this, never from Key directly, so they can't drift apart from one another or from
// the duplicate check ValidateEventIdentifiers runs against this same value.
func EventGoVarName(e *EmiEvent) string {
	return NormaliseKey(e.Key)
}

// ValidateEventIdentifiers walks a module's events list and rejects any event whose
// Key isn't a valid identifier basis - the same rule EmiPermission's Name/Key are held
// to (see permissionIdentifierPattern in EmiPermission.go, reused here since both
// ultimately have to survive being turned into a Go/JS/... identifier by a generator).
// It also rejects a duplicate Key, since two events sharing one would collide on the
// same generated variable name, and - separately - a duplicate normalised name, since
// two distinct keys can still collide once run through EventGoVarName (e.g.
// "user_created" and "userCreated" both normalise to "UserCreated").
func ValidateEventIdentifiers(events []*EmiEvent) error {
	seen := map[string]bool{}
	seenNames := map[string]string{}

	for _, e := range events {
		if e == nil {
			continue
		}

		if e.Key == "" {
			return fmt.Errorf("event has no key: an event's key is required (used as both its wire name and its generated variable name)")
		}

		if !permissionIdentifierPattern.MatchString(e.Key) {
			return fmt.Errorf(
				"event %q has an invalid key: an event's key must start with a letter or underscore and contain only letters, digits, and underscores",
				e.Key,
			)
		}

		if seen[e.Key] {
			return fmt.Errorf("duplicate event key %q: event keys must be unique within a module", e.Key)
		}
		seen[e.Key] = true

		name := EventGoVarName(e)
		if otherKey, exists := seenNames[name]; exists {
			return fmt.Errorf(
				"event %q and event %q both normalise to the same generated variable name %q: event keys must be distinct once normalised to UpperCamelCase",
				otherKey, e.Key, name,
			)
		}
		seenNames[name] = e.Key
	}

	return nil
}

// ParamsDtoAlias is an event/permission whose params reference an existing dto by name
// (params.dto) instead of declaring fields: the generated <Class> is an alias of Dto.
type ParamsDtoAlias struct {
	// Class is the generated params type name (EventParamsClassName / PermissionParamsClassName).
	Class string
	// Dto is the dto name exactly as written in `params.dto`.
	Dto string
}

// ParamsDtoAliases lists every event and permission of the module whose params use
// `dto:`, so a language backend can declare the alias types. Go declares its own
// aliases inline (next to the event / in PermissionParams); the other languages
// declare all of them together in one file.
func (m *Emi) ParamsDtoAliases() []ParamsDtoAlias {
	var out []ParamsDtoAlias
	for _, e := range m.Events {
		if e != nil && e.HasParamsDto() {
			out = append(out, ParamsDtoAlias{Class: EventParamsClassName(e), Dto: e.GetParamsDto()})
		}
	}
	for _, p := range FlattenPermissions(m.Permissions) {
		if p.HasParamsDto() {
			out = append(out, ParamsDtoAlias{Class: PermissionParamsClassName(p), Dto: p.GetParamsDto()})
		}
	}
	return out
}
