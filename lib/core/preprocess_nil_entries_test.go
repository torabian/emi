package core

import (
	"strings"
	"testing"
)

func TestStringToEmiRejectsNilEntries(t *testing.T) {
	cases := map[string]struct {
		yaml string
		want string
	}{
		"bare dash in actions": {"actions:\n  - name: ok\n  - \n", "actions[1] is empty"},
		"null action":          {"actions:\n  - null\n", "actions[0] is empty"},
		"bare dash in remotes": {"remotes:\n  - \n", "remotes[0] is empty"},
		"bare dash in events":  {"events:\n  - \n", "events[0] is empty"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := StringToEmi(tc.yaml)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("expected error containing %q, got %v", tc.want, err)
			}
		})
	}
}
