package js

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// A reactive action with `binaryType: arraybuffer` must hand back a socket that
// already receives binary frames as ArrayBuffers (what MediaStreamer expects);
// a plain reactive action keeps the browser default.
func TestReactiveBinaryType(t *testing.T) {
	create := func(binaryType string) string {
		action := &core.EmiAction{Name: "streamMedia", Url: "/media/stream", Method: "reactive", BinaryType: binaryType}
		realms, _, err := JsActionManifestRealms(action, core.MicroGenContext{Tags: "typescript"}, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		return string(realms.FetchMetaClass.ActualScript)
	}

	if got := create("arraybuffer"); !strings.Contains(got, `socket.binaryType = "arraybuffer"`) {
		t.Errorf("expected binaryType to be applied to the socket, got:\n%v", got)
	}
	if got := create(""); strings.Contains(got, "binaryType") {
		t.Errorf("did not expect binaryType for a text reactive action, got:\n%v", got)
	}
}
