package golang

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

const paramsInterfacesYaml = `
name: blog
namespace: blog
interfaces:
  - name: scoped
    fields:
      - name: workspaceId
        type: string
permissions:
  - name: post
    key: post
    children:
      - name: publish
        key: publish
        params:
          implements: [scoped]
          fields:
            - name: label
              type: string
events:
  - key: postPublished
    params:
      implements: [scoped]
      fields:
        - name: channel
          type: string
    payload:
      implements: [scoped]
      fields:
        - name: postId
          type: string
`

// Event params/payload and permission params honor `implements` like dtos do: the
// interface's fields are included and the struct satisfies the Go interface.
func TestEventAndPermissionParamsImplementInterfaces(t *testing.T) {
	module, err := core.StringToEmi(paramsInterfacesYaml)
	if err != nil {
		t.Fatalf("StringToEmi error: %v", err)
	}

	files, err := GoModuleFull(&module, core.MicroGenContext{})
	if err != nil {
		t.Fatalf("GoModuleFull error: %v", err)
	}

	byName := map[string]string{}
	for _, f := range files {
		byName[f.Name] = f.ActualScript
	}

	checks := map[string][]string{
		"BlogEvents": {
			"type PostPublishedEventParams struct {",
			"WorkspaceId string",
			"func (x *PostPublishedEventParams) GetWorkspaceId()",
			"type PostPublishedEventPayload struct {",
			"func (x *PostPublishedEventPayload) GetWorkspaceId()",
		},
		"BlogPermissionParams": {
			"type PostPermissionsPublishParams struct {",
			"func (x *PostPermissionsPublishParams) GetWorkspaceId()",
		},
	}
	for file, wants := range checks {
		got, ok := byName[file]
		if !ok {
			t.Fatalf("missing %s file", file)
		}
		for _, want := range wants {
			if !strings.Contains(got, want) {
				t.Errorf("%s missing %q:\n%s", file, want, got)
			}
		}
	}
}

func TestImplementsWithoutInterfacesOnEventFails(t *testing.T) {
	_, err := core.StringToEmi("name: a\nnamespace: a\nevents:\n  - key: x\n    params:\n      implements: [nope]\n      fields:\n        - name: a\n          type: string\n")
	if err == nil {
		t.Fatal("expected an error for implements without declared interfaces")
	}
}
