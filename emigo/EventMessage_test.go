package emigo

import (
	"strings"
	"testing"
)

type invitePayload struct {
	FirstName string `json:"firstName"`
	Workspace string `json:"workspaceName"`
}

func TestRenderMessageUsesPayloadFieldsByTheirDefinitionName(t *testing.T) {
	e := Event{
		Key:     "invited",
		Message: map[string]string{"en": "{{.firstName}} was invited to {{.workspaceName}}.", "pl": "{{.firstName}} zaproszono do {{.workspaceName}}."},
		Payload: invitePayload{FirstName: "Ali", Workspace: "Acme"},
	}

	got, err := e.RenderMessage(nil)
	if err != nil {
		t.Fatal(err)
	}
	if got["en"] != "Ali was invited to Acme." || got["pl"] != "Ali zaproszono do Acme." {
		t.Fatalf("rendered = %v", got)
	}
	if e.Message["en"] != "{{.firstName}} was invited to {{.workspaceName}}." {
		t.Fatalf("the declared message must not change: %v", e.Message)
	}
}

func TestRenderMessageParamsWinOverPayload(t *testing.T) {
	e := Event{
		Message: map[string]string{"en": "{{.inviter}} invited {{.firstName}}."},
		Payload: map[string]any{"firstName": "Ali", "inviter": "from payload"},
	}
	got, err := e.RenderMessage(map[string]any{"inviter": "Sara"})
	if err != nil || got["en"] != "Sara invited Ali." {
		t.Fatalf("rendered = %v, err = %v", got, err)
	}
}

func TestRenderMessageWithoutTemplatesIsUntouched(t *testing.T) {
	e := Event{Message: map[string]string{"en": "A member joined."}}
	got, err := e.RenderMessage(map[string]any{"x": 1})
	if err != nil || got["en"] != "A member joined." {
		t.Fatalf("rendered = %v, err = %v", got, err)
	}
}

func TestRenderMessageMissingKeyIsEmptyAndDefaultGivesAFallback(t *testing.T) {
	e := Event{
		Message: map[string]string{"en": "[{{.nobody}}] {{default \"Someone\" .inviter}} joined, {{upper .team}}."},
		Payload: map[string]any{"team": "red"},
	}
	got, err := e.RenderMessage(nil)
	if err != nil {
		t.Fatal(err)
	}
	if got["en"] != "[] Someone joined, RED." {
		t.Fatalf("rendered = %q", got["en"])
	}
}

func TestRenderMessageBrokenLocaleKeepsItsTextAndOthersStillRender(t *testing.T) {
	e := Event{
		Message: map[string]string{"en": "Hello {{.name}}.", "pl": "Cześć {{.name"},
		Payload: map[string]any{"name": "Ali"},
	}
	got, err := e.RenderMessage(nil)
	if got["en"] != "Hello Ali." {
		t.Fatalf("the healthy locale must render: %v", got)
	}
	if got["pl"] != "Cześć {{.name" {
		t.Fatalf("the broken locale keeps its text: %v", got)
	}
	if err == nil || !strings.Contains(err.Error(), `"pl"`) {
		t.Fatalf("the problem is reported with its locale: %v", err)
	}
}

func TestRenderMessageNonObjectPayloadContributesNothing(t *testing.T) {
	e := Event{Message: map[string]string{"en": "Done {{.what}}."}, Payload: "just a string"}
	got, err := e.RenderMessage(map[string]any{"what": "well"})
	if err != nil || got["en"] != "Done well." {
		t.Fatalf("rendered = %v, err = %v", got, err)
	}
}
