package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoActionRenderReactive(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {

	realms, deps, err := GoActionReactiveRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	// For socket, we need some extra dependencies
	deps = append(deps, core.CodeChunkDependency{
		Location: "github.com/gorilla/websocket",
	})

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  core.TOKEN_ORIGINAL_NAME,
				Value: realms.ActionName,
			},
		},
	}

	const tmpl = `/**
* Action to communicate with the action {{ .realms.ActionName }}
*/

{{ define "upgradeSequence" }}

	{{ if .realms.PathParameter }}
	pathParams := {{ .realms.ActionName }}PathParameterFromGin(c)
	{{ end }}

	ws, err := upgrader{{ .realms.ActionName }}.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
		return
	}
{{ end }}

func {{ .realms.ActionName }}Meta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "{{ .realms.ActionName }}",
        URL:    "{{ .realms.SafeUrl }}",
        Method: "{{ UPPER .action.Method }}",
    }
}

{{ if .realms.PathParameter }}
	{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}

{{ if .realms.QueryParams }}
	{{ b2s .realms.QueryParams.ActualScript }}
{{ end }}


// WebSocket upgrader
var upgrader{{ .realms.ActionName }} = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}


type {{ .realms.ActionName }}Message struct {
	Raw []byte
	Conn *websocket.Conn	
	MessageType int
	Error error
	{{ if .realms.PathParameter }}
	PathParams {{ .realms.ActionName }}PathParameter
	{{ end}}

	QueryParams {{ .realms.ActionName }}Query
}

// Developer handler type
type {{ .realms.ActionName }}Handler func(msg {{ .realms.ActionName }}Message ) error

// Generated handler
func {{ .realms.ActionName }}(r *gin.Engine, handler {{ .realms.ActionName }}Handler) {
	meta := {{ .realms.ActionName }}Meta()
	r.GET(meta.URL, func(c *gin.Context) {
		{{ template "upgradeSequence" . }}

		defer ws.Close()

		for {
			mt, raw, err := ws.ReadMessage()
			msg := {{ .realms.ActionName }}Message{
				Conn: ws,
				Raw: raw,
				Error: err,
				MessageType: mt,
			}

			{{ if .realms.PathParameter }}
			msg.PathParams = pathParams
			{{ end}}

			msg.QueryParams = {{ .realms.ActionName }}QueryFromGin(c)
			
			// Provide raw message to developer handler
			if err := handler(msg); err != nil {
				errMsg := fmt.Sprintf("handler error: %v", err)
				if writeErr := ws.WriteMessage(mt, []byte(errMsg)); writeErr != nil {
					break
				}
			}
		}
	})
}


type {{ .realms.ActionName }}Session struct {
	In   <-chan {{ .realms.ActionName }}Message
	Out  chan<- {{ .realms.ActionName }}Message
	Done <-chan struct{}

	Close func(err error)

	{{ if .realms.PathParameter }}
	PathParams {{ .realms.ActionName }}PathParameter
	{{ end}}

	QueryParams {{ .realms.ActionName }}Query

}

type {{ .realms.ActionName }}HandlerDuplex func(*{{ .realms.ActionName }}Session)


// {{ .realms.ActionName }}Duplex upgrades the HTTP connection to a WebSocket and
// exposes it as a full-duplex, blocking session.
//
// The provided handler owns the lifetime of the connection.
// The WebSocket remains open as long as the handler is running.
// Returning from the handler will close the connection.
//
// Session channels:
//   - ctx.In   : incoming messages from the client (closed on disconnect)
//   - ctx.Out  : outgoing messages to the client (blocking send)
//   - ctx.Done : closed when the server terminates the session
//
// Usage pattern:
//
//	{{ .realms.PackageName }}.{{ .realms.ActionName }}Duplex(r, func(ctx *{{ .realms.PackageName }}.{{ .realms.ActionName }}Session) {
//		for {
//			select {
//			case msg, ok := <-ctx.In:
//				if !ok {
//					return // client disconnected
//				}
//				ctx.Out <- {{ .realms.PackageName }}.{{ .realms.ActionName }}Message{
//					MessageType: websocket.TextMessage,
//					Raw:         msg.Raw,
//				}
//
//			case <-ctx.Done:
//				return // server-side close
//			}
//		}
//	})
//
// Important:
//   - Always read the generated code, don't use blindly.
//   - If there is an error on write, you'll get a message back, with message type -1 (instead of default websocket message type int.)
//   - The handler MUST block (typically via a loop).
//   - Returning from the handler closes the WebSocket.
//   - Do not treat this as a per-message callback.
func {{ .realms.ActionName }}Duplex(r *gin.Engine, handler {{ .realms.ActionName }}HandlerDuplex) {
	meta := {{ .realms.ActionName }}Meta()

	// The actual callback is extracted, in case you need to handle multiple handlers or customize, use it directly.
	r.GET(meta.URL, func(ctx *gin.Context) {
		{{ .realms.ActionName }}DuplexGinHandler(ctx, handler)
	})
}

func {{ .realms.ActionName }}DuplexGinHandler(c *gin.Context, handler {{ .realms.ActionName }}HandlerDuplex) {
	{{ template "upgradeSequence" . }}

	in := make(chan {{ .realms.ActionName }}Message)
	out := make(chan {{ .realms.ActionName }}Message)
	done := make(chan struct{})

	session := &{{ .realms.ActionName }}Session{
		In:   in,
		Out:  out,
		Done: done,
		Close: func(err error) {
			close(done)
			ws.Close()
		},
	}

	{{ if .realms.PathParameter }}
		session.PathParams = pathParams
	{{ end }}

	session.QueryParams = {{ .realms.ActionName }}QueryFromGin(c)

	// Read loop
	go func() {
		defer close(in)
		for {
			mt, raw, err := ws.ReadMessage()
			in <- {{ .realms.ActionName }}Message{MessageType: mt, Raw: raw, Error: err}
		}
	}()

	// Write loop
	go func() {
		for msg := range out {
			if err := ws.WriteMessage(msg.MessageType, msg.Raw); err != nil {
				// When message is -1, means it's internal error coming out
				in <- {{ .realms.ActionName }}Message{MessageType: -1, Error: err}
				return
			}
		}
	}()

	// Run developer code (blocking)
	handler(session)

	// Cleanup
	close(out)
	ws.Close()
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
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".go"
	res.CodeChunkDependensies = deps

	return res, nil
}
