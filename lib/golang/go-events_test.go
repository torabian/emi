package golang

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// eventsTestModuleYaml is the exact module used by golang-events.mdx (the "Events"
// docs page, in the golang section) to show the yaml -> Go pairing side by side. Keep
// this and the yaml block on that page identical - TestEventsGenerate below asserts
// the compiler's actual output still contains every line the doc claims it does, so
// the two can't silently drift apart.
//
// It exercises the three ways an event's payload can be typed (see
// core.EmiEvent.Payload/HasPayloadFields/HasPayloadDto in lib/core/EmiEvent.go):
//   - postPublished declares payload fields inline, compiled into a real struct via
//     the same Common struct builder action bodies use (GoCommonStructGenerator), and
//     also declares permissions using the "with" shape (core.EmiEventPermission).
//   - commentAdded references an existing dto by name instead.
//   - postArchived declares no payload at all, so its generated payload type is a
//     plain interface{} alias.
const eventsTestModuleYaml = `
name: blog
namespace: blog
dtos:
  - name: comment
    fields:
      - name: commentId
        type: string
      - name: body
        type: text
events:
  - key: postPublished
    name:
      en: Post published
    description:
      en: Fires once a post moves from draft to published.
    permissions:
      - with: ["post.publish"]
      - with: ["post.*"]
    params:
      fields:
        - name: workspaceId
          type: string
    payload:
      fields:
        - name: postId
          type: string
        - name: title
          type: string
        - name: publishedAt
          type: string
  - key: commentAdded
    payload:
      dto: CommentDto
  - key: postArchived
`

// TestEventsGenerate compiles eventsTestModuleYaml's events: list end to end (through
// GoModuleFull, the same path the "go" primary action uses) and checks the generated
// Events.go against the exact shapes documented on golang-events.mdx: a payload
// struct compiled from fields (plus its Cli helpers, since those aren't disabled
// here), a payload type aliased to an existing dto, a payload aliased to interface{}
// when none was declared, the "with"-shaped permissions flattened into
// emigo.Event's [][]string, and AllEventsList tying every var together.
func TestEventsGenerate(t *testing.T) {
	module, err := core.StringToEmi(eventsTestModuleYaml)
	if err != nil {
		t.Fatalf("StringToEmi error: %v", err)
	}

	files, err := GoModuleFull(&module, core.MicroGenContext{})
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}

	var eventsFile *core.VirtualFile
	for i := range files {
		if files[i].Name == "BlogEvents" {
			eventsFile = &files[i]
			break
		}
	}
	if eventsFile == nil {
		t.Fatalf("expected an Events file, got: %+v", fileNames(files))
	}

	got := eventsFile.ActualScript

	wantContains := []string{
		// postPublished: payload fields compiled into a real struct - one field per
		// declared payload field, PascalCase, JSON/yaml tags matching the yaml name.
		"type PostPublishedEventParams struct {",
		"WorkspaceId string `json:\"workspaceId\" yaml:\"workspaceId\"`",
		"type PostArchivedEventParams = interface{}",
		"type PostPublishedEventPayload struct {",
		"PostId      string `json:\"postId\" yaml:\"postId\"`",
		"Title       string `json:\"title\" yaml:\"title\"`",
		"PublishedAt string `json:\"publishedAt\" yaml:\"publishedAt\"`",
		// ... plus its CLI helpers, generated the same way an action's request body's
		// would be (GoCommonStructGeneratorCli), since split-cli isn't set here.
		"func GetPostPublishedEventPayloadCliFlags(prefix string) []emigo.CliFlag {",
		"func CastPostPublishedEventPayloadFromCli(c emigo.CliCastable) PostPublishedEventPayload {",
		// permissions: "with" AND-sets, OR'd together, flattened to [][]string.
		`var PostPublishedEvent = emigo.Event{`,
		`Key:         "postPublished",`,
		`Name:        map[string]string{"en": "Post published"},`,
		`Description: map[string]string{"en": "Fires once a post moves from draft to published."},`,
		"Permissions: [][]string{",
		`{"post.publish"},`,
		`{"post.*"},`,
		// every event also gets a New<EventName> constructor that copies the base var
		// and sets Payload, typed through its own <EventName>Payload identifier.
		"func NewPostPublishedEvent(payload PostPublishedEventPayload) emigo.Event {",
		"evt := PostPublishedEvent",
		"evt.Payload = payload",
		"return evt",
		// commentAdded: payload aliased to the dto it references by name.
		"type CommentAddedEventPayload = CommentDto",
		`var CommentAddedEvent = emigo.Event{`,
		`Key:         "commentAdded",`,
		"func NewCommentAddedEvent(payload CommentAddedEventPayload) emigo.Event {",
		// postArchived: no payload declared, aliased straight to interface{}.
		"type PostArchivedEventPayload = interface{}",
		`var PostArchivedEvent = emigo.Event{`,
		`Key:         "postArchived",`,
		"func NewPostArchivedEvent(payload PostArchivedEventPayload) emigo.Event {",
		// every var tied together, in declaration order.
		"var AllEventsList = []emigo.Event{",
		"\tPostPublishedEvent,\n\tCommentAddedEvent,\n\tPostArchivedEvent,\n}",
	}

	for _, want := range wantContains {
		if !strings.Contains(got, want) {
			t.Errorf("expected generated Events file to contain:\n%s\n\ngot:\n%s", want, got)
		}
	}
}

func TestPermissionParamsGenerate(t *testing.T) {
	module, err := core.StringToEmi(`
name: blog
namespace: blog
permissions:
  - name: post
    key: post
    children:
      - name: publish
        key: publish
        params:
          fields:
            - name: workspaceId
              type: string
`)
	if err != nil {
		t.Fatalf("StringToEmi error: %v", err)
	}

	files, err := GoModuleFull(&module, core.MicroGenContext{})
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}

	for _, f := range files {
		if f.Name == "BlogPermissionParams" {
			if !strings.Contains(f.ActualScript, "type PostPermissionsPublishParams struct {") {
				t.Fatalf("missing params struct:\n%s", f.ActualScript)
			}
			return
		}
	}
	t.Fatalf("expected a PermissionParams file, got: %+v", fileNames(files))
}

func TestParamsDtoAliasesGo(t *testing.T) {
	module, err := core.StringToEmi(`
name: blog
namespace: blog
dtos:
  - name: scope
    fields:
      - name: workspaceId
        type: string
permissions:
  - name: post
    key: post
    params:
      dto: ScopeDto
events:
  - key: postPublished
    params:
      dto: ScopeDto
`)
	if err != nil {
		t.Fatalf("StringToEmi error: %v", err)
	}
	files, err := GoModuleFull(&module, core.MicroGenContext{})
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}
	got := map[string]string{}
	for _, f := range files {
		got[f.Name] = f.ActualScript
	}
	if !strings.Contains(got["BlogPermissionParams"], "type PostPermissionsParams = ScopeDto") {
		t.Errorf("permission params alias missing:\n%s", got["BlogPermissionParams"])
	}
	if !strings.Contains(got["BlogEvents"], "type PostPublishedEventParams = ScopeDto") {
		t.Errorf("event params alias missing:\n%s", got["BlogEvents"])
	}
}

// A permission's params are also embedded as JSON Schema text on the emigo.Permission
// literal (ParamsSchema), for inline fields and for dto references, and absent otherwise.
func TestPermissionParamsSchemaEmbedded(t *testing.T) {
	module, err := core.StringToEmi(`
name: blog
namespace: blog
dtos:
  - name: scope
    fields:
      - name: channel
        type: string
permissions:
  - name: post
    key: post
    children:
      - name: publish
        key: publish
        params:
          fields:
            - name: workspaceId
              type: string
      - name: archive
        key: archive
        params:
          dto: ScopeDto
      - name: view
        key: view
`)
	if err != nil {
		t.Fatalf("StringToEmi error: %v", err)
	}

	files, err := GoModuleFull(&module, core.MicroGenContext{})
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}

	for _, f := range files {
		if f.Name != "BlogPermissions" {
			continue
		}
		if got := strings.Count(f.ActualScript, "ParamsSchema:"); got != 2 {
			t.Fatalf("want ParamsSchema on exactly the 2 permissions with params, got %d:\n%s", got, f.ActualScript)
		}
		for _, want := range []string{`workspaceId`, `channel`} {
			if !strings.Contains(f.ActualScript, want) {
				t.Fatalf("schema lacks %q:\n%s", want, f.ActualScript)
			}
		}
		return
	}
	t.Fatalf("no permissions file in %+v", fileNames(files))
}
