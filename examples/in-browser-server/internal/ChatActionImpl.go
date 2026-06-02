package internal

// This file is the manual, actual implementation of the reactive ChatAction.
// Like SubstringActionImpl, it is transport-agnostic: the very same factory is
// wired into the gin server (via defs.ChatActionReactiveHandler) and into the
// wasm build (via defs.ChatActionReactiveHandlerWasm).

import (
	"fmt"

	"github.com/torabian/emi/examples/in-browser-server/internal/defs"
)

// ChatHandler echoes back the length of every message the client types. It
// returns the write channel the transport drains; closing that channel (here on
// session.Done) initiates a clean close.
func ChatHandler(session defs.ChatActionSession) (chan []byte, error) {
	out := make(chan []byte, 16)

	go func() {
		defer close(out)
		for {
			select {
			case msg := <-session.Read:
				if msg.Error != nil {
					return
				}
				out <- []byte(fmt.Sprintf("you typed %d chars: %q", len(msg.Data), string(msg.Data)))
			case <-session.Done:
				return
			}
		}
	}()

	return out, nil
}
