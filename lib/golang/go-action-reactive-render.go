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
	CliName string
	Description string
} {
    return struct {
        Name   string
        URL    string
        Method string
		CliName string
		Description string
    }{
        Name:   "{{ .realms.ActionName }}",
        URL:    "{{ .realms.SafeUrl }}",
        Method: "{{ UPPER .action.Method }}",
		CliName: "{{ .action.CliName }}",
		Description: "{{ .action.Description }}",
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
	Ctx    *gin.Context
	Socket *websocket.Conn
	Done   chan bool
	Read   chan {{ .realms.ActionName }}ReadChan
	QueryParams {{ .realms.ActionName }}Query
}

type {{ .realms.ActionName }}HandlerDuplex func(*{{ .realms.ActionName }}Session)


type {{ .realms.ActionName }}ReadChan struct {
	Data  []byte
	Error error
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
		session.QueryParams = {{ .realms.ActionName }}QueryFromGin(ctx)

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



// {{ .realms.ActionName }}ClientSession is the client-side mirror of
// {{ .realms.ActionName }}Session. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type {{ .realms.ActionName }}ClientSession struct {
	Socket *websocket.Conn
	Done   chan bool
	Read   chan {{ .realms.ActionName }}ReadChan
	Write  chan []byte
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
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".go"
	res.CodeChunkDependensies = deps

	return res, nil
}
