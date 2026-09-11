package emigo

// Event is a single declared event, as described via a module's `events:` yaml block
// and compiled by the emi go compiler (see lib/golang/go-events.go). It lives here,
// in the shared runtime, instead of being redefined at the top of every generated
// Events.go - the same reasoning as Permission living here for Permissions.go.
type Event struct {

	// Key is the fully qualified, unique identifier of this event - both its
	// wire-level name (what a firing call actually sends as the event's name) and the
	// basis its generated variable name was derived from (e.g. "userCreated" ->
	// UserCreatedEvent).
	Key string

	// Name is the human readable, localized label of the event, keyed by locale.
	Name map[string]string

	// Description explains, for a human, what this event represents, keyed by locale.
	Description map[string]string

	// Permissions is a list of permission combinations that grant visibility into this
	// event: the outer slice is OR'd together, and each inner slice of permission keys
	// is AND'd - e.g. [[A, B], [C, D]] means "(A and B) or (C and D)". Empty means no
	// specific permission is required to know this event happened.
	Permissions [][]string

	// Payload carries this event's data at fire time. Declared untyped here since one
	// shared Event struct backs every event in a module; the generated
	// <EventName>Payload type alongside each event's var (see lib/golang/go-events.go)
	// gives the concrete shape to assert/cast against - interface{} itself when the
	// event declared no payload.
	Payload interface{}
}

func (x Event) GetKey() string {
	return x.Key
}

func (x Event) GetName() map[string]string {
	return x.Name
}

func (x Event) GetDescription() map[string]string {
	return x.Description
}

func (x Event) GetPermissions() [][]string {
	return x.Permissions
}

// MeetsPermissions reports whether granted (the flat set of permission keys an
// identity holds) satisfies at least one of Event's OR'd permission combinations -
// i.e. every key within at least one inner AND-set is present in granted. An event
// declaring no Permissions at all is met by anyone.
func (x Event) MeetsPermissions(granted []string) bool {
	if len(x.Permissions) == 0 {
		return true
	}

	held := make(map[string]bool, len(granted))
	for _, g := range granted {
		held[g] = true
	}

	for _, set := range x.Permissions {
		if len(set) == 0 {
			continue
		}

		allHeld := true
		for _, key := range set {
			if !held[key] {
				allHeld = false
				break
			}
		}
		if allHeld {
			return true
		}
	}

	return false
}
