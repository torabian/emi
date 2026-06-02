package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type reactiveWasmRealms struct {
	ActionName string
}

func GoActionRenderReactiveWasm(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	realms := &reactiveWasmRealms{
		ActionName: core.ToUpper(core.NormaliseKey(action.GetName())),
	}

	f := GetCommonFlags(ctx)

	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Location: f.Emigo,
			},
		},
	}

	const tmpl = `
//go:build wasm

// {{ .realms.ActionName }}ReactiveHandlerWasm is the in-browser counterpart of
// {{ .realms.ActionName }}ReactiveHandler. It registers the same developer factory against an
// emigo.WasmReactor instead of a gin engine, so the business logic is shared
// verbatim across the real server and the wasm build.
//
// There is no gorilla, no gin, and no socket here: the reactor bridges the
// session's channels straight to the WebSocketWasm JS class.
func {{ .realms.ActionName }}ReactiveHandlerWasm(
	reactor *emigo.WasmReactor,
	factory func(session {{ .realms.ActionName }}Session) (chan []byte, error),
) {
	reactor.Handle({{ .realms.ActionName }}Meta().URL, func(conn *emigo.WasmReactiveConn) error {
		session := {{ .realms.ActionName }}Session{
			Done:        conn.Done,
			Read:        make(chan {{ .realms.ActionName }}ReadChan),
			QueryParams: {{ .realms.ActionName }}QueryFromString(conn.Query.Encode()),
			// Socket and Ctx stay nil — meaningless in the browser.
		}

		// client -> server: adapt raw []byte frames into the typed read channel.
		go func() {
			for {
				select {
				case data, ok := <-conn.Read:
					if !ok {
						return
					}
					session.Read <- {{ .realms.ActionName }}ReadChan{Data: data}
				case <-conn.Done:
					return
				}
			}
		}()

		write, err := factory(session)
		if err != nil {
			return err
		}

		// server -> client: forward the factory's write channel to the bridge.
		// Closing write signals a server-initiated close to the reactor.
		go func() {
			for {
				select {
				case msg, ok := <-write:
					if !ok {
						close(conn.Write)
						return
					}
					conn.Write <- msg
				case <-conn.Done:
					return
				}
			}
		}()

		return nil
	})
}


`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":       action,
		"realms":       realms,
		"shouldExport": true,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName + "Wasm"
	res.SuggestedExtension = ".go"

	return res, nil
}
