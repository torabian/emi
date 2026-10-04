package js

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// Event and permission params go through JsCommonObjectGenerator like dtos, so with
// --tags json-schema they must carry static JsonSchema and DefaultTranslations.
func TestEventAndPermissionParamsJsonSchema(t *testing.T) {
	m, err := core.StringToEmi(`
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
events:
  - key: postPublished
    params:
      fields:
        - name: workspaceId
          type: string
`)
	if err != nil {
		t.Fatal(err)
	}

	for _, tags := range []string{"json-schema", "json-schema,typescript"} {
		files, err := JsModuleFullVirtualFiles(&m, core.MicroGenContext{Tags: tags})
		if err != nil {
			t.Fatal(err)
		}
		for _, class := range []string{"PostPublishedEventParams", "PostPermissionsPublishParams"} {
			found := false
			for _, f := range files {
				if strings.Contains(f.ActualScript, "class "+class) {
					found = true
					for _, want := range []string{"static JsonSchema =", "static DefaultTranslations"} {
						if !strings.Contains(f.ActualScript, want) {
							t.Errorf("[%s] %s missing %q:\n%s", tags, class, want, f.ActualScript)
						}
					}
				}
			}
			if !found {
				t.Errorf("[%s] no class %s generated", tags, class)
			}
		}
	}
}

// Params implementing an interface get its fields and an `implements` clause.
func TestParamsImplementInterfacesJs(t *testing.T) {
	m, err := core.StringToEmi(`
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
`)
	if err != nil {
		t.Fatal(err)
	}

	files, err := JsModuleFullVirtualFiles(&m, core.MicroGenContext{Tags: "typescript,json-schema"})
	if err != nil {
		t.Fatal(err)
	}
	for _, class := range []string{"PostPermissionsParams", "PostPublishedEventParams"} {
		found := false
		for _, f := range files {
			if strings.Contains(f.ActualScript, "class "+class+" implements Scoped") {
				found = true
				if !strings.Contains(f.ActualScript, "workspaceId") || !strings.Contains(f.ActualScript, "workspace_id_title") {
					t.Errorf("%s lacks the interface field / schema:\n%s", class, f.ActualScript)
				}
			}
		}
		if !found {
			t.Errorf("no %s implementing Scoped", class)
		}
	}
}

func TestParamsDtoAliasesJs(t *testing.T) {
	m, err := core.StringToEmi(`
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
		t.Fatal(err)
	}
	files, err := JsModuleFullVirtualFiles(&m, core.MicroGenContext{Tags: "typescript"})
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if f.Name == "BlogParamsAliases" {
			for _, want := range []string{`import { ScopeDto as _ScopeDto } from "./ScopeDto"`, "export { _ScopeDto as PostPermissionsParams };", "export { _ScopeDto as PostPublishedEventParams };"} {
				if !strings.Contains(f.ActualScript, want) {
					t.Errorf("missing %q:\n%s", want, f.ActualScript)
				}
			}
			return
		}
	}
	t.Fatal("no BlogParamsAliases file")
}
