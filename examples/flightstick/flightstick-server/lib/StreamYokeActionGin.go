package external

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"unicode/utf8"
)

/**
* Action to communicate with the action StreamYokeAction
 */
func StreamYokeActionPathParameterFromGin(g *gin.Context) StreamYokeActionPathParameter {
	return StreamYokeActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// WebSocket upgrader
var upgraderStreamYokeAction = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (x *StreamYokeActionMessage) Connection() *websocket.Conn {
	return x.Conn.(*websocket.Conn)
}
func (x *StreamYokeActionClientSession) Connection() *websocket.Conn {
	return x.Socket.(*websocket.Conn)
}
func (x *StreamYokeActionSession) GinCtx() *gin.Context {
	return x.Ctx.(*gin.Context)
}
func (x *StreamYokeActionSession) GetSocket() *websocket.Conn {
	return x.Socket.(*websocket.Conn)
}

// Generated handler
func StreamYokeActionGin(r *gin.Engine, factory func(
	session StreamYokeActionSession,
) (chan []byte, error)) {
	meta := StreamYokeActionMeta()
	r.GET(meta.URL, StreamYokeActionReactiveHandler(factory))
}
func StreamYokeActionReactiveHandler(factory func(
	session StreamYokeActionSession,
) (chan []byte, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		read := make(chan StreamYokeActionReadChan)
		done := make(chan bool)
		c, err := upgraderStreamYokeAction.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			c.Close()
			return
		}
		session := StreamYokeActionSession{
			Ctx:    ctx,
			Socket: c,
			Done:   done,
			Read:   read,
		}
		session.QueryParams = StreamYokeActionQueryFromHttp(ctx.Request)
		write, err := factory(session)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
		go func() {
			for {
				_, data, err := c.ReadMessage()
				read <- StreamYokeActionReadChan{
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

// StreamYokeActionClientOptions configures a client dial. All fields are
// optional — pass a zero value for a plaintext, header-less ws:// connection.
//
// TLSConfig governs the TLS handshake when dialing wss://. Set Certificates
// and RootCAs (and optionally ServerName) for mTLS; leave nil for ws://.
type StreamYokeActionClientOptions struct {
	Query     url.Values
	TLSConfig *tls.Config
	Headers   http.Header
}

// StreamYokeActionClient dials the StreamYokeAction endpoint at baseURL
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
//	sess, err := StreamYokeActionClient(
//	    "wss://hub.example.com",
//	    StreamYokeActionClientOptions{TLSConfig: tlsCfg},
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
func StreamYokeActionClient(baseURL string, opts *StreamYokeActionClientOptions) (*StreamYokeActionClientSession, error) {
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
	u.Path = StreamYokeActionMeta().URL
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
	session := &StreamYokeActionClientSession{
		Socket: c,
		Done:   make(chan bool, 1),
		Read:   make(chan StreamYokeActionReadChan),
		Write:  make(chan []byte, 16),
	}
	// Reader goroutine: pumps frames from the socket into Read. On error it
	// forwards the error frame, signals Done, and exits.
	go func() {
		for {
			_, data, err := c.ReadMessage()
			session.Read <- StreamYokeActionReadChan{
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
