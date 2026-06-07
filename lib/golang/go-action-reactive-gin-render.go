package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoActionRenderReactiveGin(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {

	realms, deps, err := GoActionReactiveGinRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	// For socket, we need some extra dependencies
	deps = append(deps, core.CodeChunkDependency{
		Location: "github.com/gorilla/websocket",
	},

		core.CodeChunkDependency{
			Location: "crypto/tls",
		},

		core.CodeChunkDependency{
			Location: "unicode/utf8",
		})

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  core.TOKEN_ORIGINAL_NAME,
				Value: realms.ActionName,
			},
		},
	}

	const tmpl = `
{{ if .realms.SkipGinWasm }}
//go:build !wasm
{{ end }}
	
/**
* Action to communicate with the action {{ .realms.ActionName }}
*/

{{ if .realms.PathParameterGin }}
	{{ b2s .realms.PathParameterGin.ActualScript }}
{{ end }}

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


// WebSocket upgrader
var upgrader{{ .realms.ActionName }} = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (x *{{ .realms.ActionName }}Message) Connection() *websocket.Conn {
	return x.Conn.(*websocket.Conn)
}

func (x *{{ .realms.ActionName }}ClientSession) Connection() *websocket.Conn {
	return x.Socket.(*websocket.Conn)
}

func (x *{{ .realms.ActionName }}Session) GinCtx() *gin.Context {
	return x.Ctx.(*gin.Context)
}

func (x *{{ .realms.ActionName }}Session) GetSocket() *websocket.Conn {
	return x.Socket.(*websocket.Conn)
}

// Generated handler
func {{ .realms.ActionName }}Gin(r *gin.Engine, factory func(
	session {{ .realms.ActionName }}Session,
) (chan []byte, error)) {
	meta := {{ .realms.ActionName }}Meta()
	r.GET(meta.URL, {{ .realms.ActionName }}ReactiveHandler(factory))
}


func {{ .realms.ActionName }}ReactiveHandler(factory func(
	session {{ .realms.ActionName }}Session,
) (chan []byte, error)) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		read := make(chan {{ .realms.ActionName }}ReadChan)
		done := make(chan bool)

		c, err := upgrader{{ .realms.ActionName }}.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))

			c.Close()
			return
		}

		session := {{ .realms.ActionName }}Session{
			Ctx:    ctx,
			Socket: c,
			Done:   done,
			Read:   read,
		}
		session.QueryParams = {{ .realms.ActionName }}QueryFromHttp(ctx.Request)

		write, err := factory(session)

		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}

		go func() {
			for {
				_, data, err := c.ReadMessage()
				read <- {{ .realms.ActionName }}ReadChan{
					Data:  data,
					Error: err,
				}

				if err != nil {
					return
				}
			}
		}()

		go func() {
			for {
				select {
				case msg, ok := <-write:
					if !ok {
						// Channel closed; shutdown
						c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
						done <- true
						return
					}
					msgType := websocket.TextMessage
					if !utf8.Valid(msg) {
						msgType = websocket.BinaryMessage
					}
					err := c.WriteMessage(msgType, msg)

					if err != nil {
						// Optionally log the error or send to a logger
						done <- true
						return
					}
				case <-done:
					c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					return
				}
			}
		}()
	}
}

// {{ .realms.ActionName }}ClientOptions configures a client dial. All fields are
// optional — pass a zero value for a plaintext, header-less ws:// connection.
//
// TLSConfig governs the TLS handshake when dialing wss://. Set Certificates
// and RootCAs (and optionally ServerName) for mTLS; leave nil for ws://.
type {{ .realms.ActionName }}ClientOptions struct {
	Query     url.Values
	TLSConfig *tls.Config
	Headers   http.Header
}

// {{ .realms.ActionName }}Client dials the {{ .realms.ActionName }} endpoint at baseURL
// (e.g. "ws://localhost:8080" or "https://host" — http/https are auto-rewritten
// to ws/wss) and returns a session whose channels behave like the server's.
//
// Connect, then read and write concurrently. The Read goroutine bails on
// msg.Error (server closed, network drop, etc.); the Done channel fires once
// for any disconnect — use it to unblock main or trigger reconnect logic.
//
// For mTLS, build a *tls.Config with the client keypair + server CA and pass
// it via opts.TLSConfig. The handshake completes before the websocket upgrade
// is sent, so a bad cert fails this call rather than later on Read/Write.
//
//	sess, err := {{ .realms.ActionName }}Client(
//	    "wss://hub.example.com",
//	    {{ .realms.ActionName }}ClientOptions{TLSConfig: tlsCfg},
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Reader: pull frames off sess.Read until error.
//	go func() {
//	    for {
//	        msg := <-sess.Read
//	        if msg.Error != nil {
//	            log.Println("read error:", msg.Error)
//	            return
//	        }
//	        log.Printf("server sent %d bytes: %s", len(msg.Data), msg.Data)
//	    }
//	}()
//
//	// Writer: send frames whenever you have something to say. Bytes that
//	// aren't valid UTF-8 are sent as binary frames automatically.
//	sess.Write <- []byte("hello server")
//
//	// Block until the connection drops, then exit. Alternatively, close
//	// sess.Write to initiate a clean shutdown from the client side:
//	//   close(sess.Write)
//	<-sess.Done
func {{ .realms.ActionName }}Client(baseURL string, opts *{{ .realms.ActionName }}ClientOptions) (*{{ .realms.ActionName }}ClientSession, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http":
		u.Scheme = "ws"
	case "https":
		u.Scheme = "wss"
	}
	u.Path = {{ .realms.ActionName }}Meta().URL
	if opts != nil && opts.Query != nil {
		u.RawQuery = opts.Query.Encode()
	}

	var headers http.Header

	if opts != nil {
		headers = opts.Headers
	}

	dialer := &websocket.Dialer{}

	if opts != nil {
		dialer.TLSClientConfig = opts.TLSConfig
	}

	c, _, err := dialer.Dial(u.String(), headers)
	if err != nil {
		return nil, err
	}

	session := &{{ .realms.ActionName }}ClientSession{
		Socket: c,
		Done:   make(chan bool, 1),
		Read:   make(chan {{ .realms.ActionName }}ReadChan),
		Write:  make(chan []byte, 16),
	}

	// Reader goroutine: pumps frames from the socket into Read. On error it
	// forwards the error frame, signals Done, and exits.
	go func() {
		for {
			_, data, err := c.ReadMessage()
			session.Read <- {{ .realms.ActionName }}ReadChan{
				Data:  data,
				Error: err,
			}
			if err != nil {
				select {
				case session.Done <- true:
				default:
				}
				return
			}
		}
	}()

	// Writer goroutine: drains Write to the socket. Closing Write triggers a
	// clean close handshake; an error or Done signal closes the socket.
	go func() {
		defer c.Close()
		for {
			select {
			case msg, ok := <-session.Write:
				if !ok {
					c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					select {
					case session.Done <- true:
					default:
					}
					return
				}
				msgType := websocket.TextMessage
				if !utf8.Valid(msg) {
					msgType = websocket.BinaryMessage
				}
				if err := c.WriteMessage(msgType, msg); err != nil {
					select {
					case session.Done <- true:
					default:
					}
					return
				}
			case <-session.Done:
				c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				return
			}
		}
	}()

	return session, nil
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
	res.SuggestedFileName = realms.ActionName + "Gin"
	res.SuggestedExtension = ".go"
	res.CodeChunkDependensies = deps

	return res, nil
}
