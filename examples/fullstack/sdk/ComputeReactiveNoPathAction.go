package external

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"net/http"
	"net/url"
	"unicode/utf8"
)

/**
* Action to communicate with the action ComputeReactiveNoPathAction
 */
func ComputeReactiveNoPathActionMeta() struct {
	Name        string
	URL         string
	Method      string
	CliName     string
	Description string
} {
	return struct {
		Name        string
		URL         string
		Method      string
		CliName     string
		Description string
	}{
		Name:        "ComputeReactiveNoPathAction",
		URL:         "/compute/reactive",
		Method:      "REACTIVE",
		CliName:     "",
		Description: "Reactive compute elsasements.",
	}
}

/**
 * Query parameters for ComputeReactiveNoPathAction
 */
// Query wrapper with private fields
type ComputeReactiveNoPathActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	QueryParam1   string `json:"queryParam1"`
	SecurityToken string `json:"securityToken"`
}

func ComputeReactiveNoPathActionQueryFromString(rawQuery string) ComputeReactiveNoPathActionQuery {
	v := ComputeReactiveNoPathActionQuery{}
	values, _ := url.ParseQuery(rawQuery)
	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}
	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json", // reuse json tags
		WeaklyTypedInput: true,   // "1" -> int, "true" -> bool
		Result:           &v,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}
	v.values = values
	v.mapped = mapped
	return v
}
func ComputeReactiveNoPathActionQueryFromGin(c *gin.Context) ComputeReactiveNoPathActionQuery {
	return ComputeReactiveNoPathActionQueryFromString(c.Request.URL.RawQuery)
}
func ComputeReactiveNoPathActionQueryFromHttp(r *http.Request) ComputeReactiveNoPathActionQuery {
	return ComputeReactiveNoPathActionQueryFromString(r.URL.RawQuery)
}
func (q ComputeReactiveNoPathActionQuery) Values() url.Values {
	return q.values
}
func (q ComputeReactiveNoPathActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ComputeReactiveNoPathActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ComputeReactiveNoPathActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

// WebSocket upgrader
var upgraderComputeReactiveNoPathAction = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type ComputeReactiveNoPathActionMessage struct {
	Raw         []byte
	Conn        *websocket.Conn
	MessageType int
	Error       error
}

// Developer handler type
type ComputeReactiveNoPathActionHandler func(msg ComputeReactiveNoPathActionMessage) error

// Generated handler
func ComputeReactiveNoPathAction(r *gin.Engine, handler ComputeReactiveNoPathActionHandler) {
	meta := ComputeReactiveNoPathActionMeta()
	r.GET(meta.URL, func(c *gin.Context) {
		ws, err := upgraderComputeReactiveNoPathAction.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
			return
		}
		defer ws.Close()
		for {
			mt, raw, err := ws.ReadMessage()
			msg := ComputeReactiveNoPathActionMessage{
				Conn:        ws,
				Raw:         raw,
				Error:       err,
				MessageType: mt,
			}
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

type ComputeReactiveNoPathActionSession struct {
	Ctx         *gin.Context
	Socket      *websocket.Conn
	Done        chan bool
	Read        chan ComputeReactiveNoPathActionReadChan
	QueryParams ComputeReactiveNoPathActionQuery
}
type ComputeReactiveNoPathActionHandlerDuplex func(*ComputeReactiveNoPathActionSession)
type ComputeReactiveNoPathActionReadChan struct {
	Data  []byte
	Error error
}

func ComputeReactiveNoPathActionReactiveHandler(factory func(
	session ComputeReactiveNoPathActionSession,
) (chan []byte, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		read := make(chan ComputeReactiveNoPathActionReadChan)
		done := make(chan bool)
		c, err := upgraderComputeReactiveNoPathAction.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			c.Close()
			return
		}
		session := ComputeReactiveNoPathActionSession{
			Ctx:    ctx,
			Socket: c,
			Done:   done,
			Read:   read,
		}
		session.QueryParams = ComputeReactiveNoPathActionQueryFromGin(ctx)
		write, err := factory(session)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
		go func() {
			for {
				_, data, err := c.ReadMessage()
				read <- ComputeReactiveNoPathActionReadChan{
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

// ComputeReactiveNoPathActionClientSession is the client-side mirror of
// ComputeReactiveNoPathActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type ComputeReactiveNoPathActionClientSession struct {
	Socket *websocket.Conn
	Done   chan bool
	Read   chan ComputeReactiveNoPathActionReadChan
	Write  chan []byte
}

// ComputeReactiveNoPathActionClientOptions configures a client dial. All fields are
// optional — pass a zero value for a plaintext, header-less ws:// connection.
//
// TLSConfig governs the TLS handshake when dialing wss://. Set Certificates
// and RootCAs (and optionally ServerName) for mTLS; leave nil for ws://.
type ComputeReactiveNoPathActionClientOptions struct {
	Query     url.Values
	TLSConfig *tls.Config
	Headers   http.Header
}

// ComputeReactiveNoPathActionClient dials the ComputeReactiveNoPathAction endpoint at baseURL
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
//	sess, err := edgehubdefs.ComputeReactiveNoPathActionClient(
//	    "wss://hub.example.com",
//	    edgehubdefs.ComputeReactiveNoPathActionClientOptions{TLSConfig: tlsCfg},
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
func ComputeReactiveNoPathActionClient(baseURL string, opts *ComputeReactiveNoPathActionClientOptions) (*ComputeReactiveNoPathActionClientSession, error) {
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
	u.Path = ComputeReactiveNoPathActionMeta().URL
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
	session := &ComputeReactiveNoPathActionClientSession{
		Socket: c,
		Done:   make(chan bool, 1),
		Read:   make(chan ComputeReactiveNoPathActionReadChan),
		Write:  make(chan []byte, 16),
	}
	// Reader goroutine: pumps frames from the socket into Read. On error it
	// forwards the error frame, signals Done, and exits.
	go func() {
		for {
			_, data, err := c.ReadMessage()
			session.Read <- ComputeReactiveNoPathActionReadChan{
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
