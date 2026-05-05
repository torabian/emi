package external

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli/v3"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"
)

/**
* Action to communicate with the action ComputeReactiveAction
 */
func ComputeReactiveActionMeta() struct {
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
		Name:        "ComputeReactiveAction",
		URL:         "/compute/reactive/:id/:age",
		Method:      "REACTIVE",
		CliName:     "",
		Description: "Reactive compute elsasements.",
	}
}

/**
 * Path parameters for ComputeReactiveAction
 */
type ComputeReactiveActionPathParameter struct {
	Id  int32
	Age int32
}

func GetComputeReactiveActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-id",
			Type:     "int32",
			Required: true,
		},
		{
			Name:     prefix + "pp-age",
			Type:     "int32",
			Required: true,
		},
	}
}

// Converts a placeholder url, and applies the parameters to it.
func ComputeReactiveActionPathParameterApply(params ComputeReactiveActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":id", fmt.Sprintf("%v", params.Id))
	templateUrl = strings.ReplaceAll(templateUrl, ":age", fmt.Sprintf("%v", params.Age))
	return templateUrl
}

// Extracts the path parameter from a gin request context
func ComputeReactiveActionPathParameterFromGin(g *gin.Context) ComputeReactiveActionPathParameter {
	return ComputeReactiveActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// Extracts the path parameter from a urfave v3 cli.
func ComputeReactiveActionPathParameterFromCli(c *cli.Command) ComputeReactiveActionPathParameter {
	return ComputeReactiveActionPathParameterFromFn(func(key string) string {
		return c.String(key)
	})
}

// General purpose to extract the value and cast based on type.
func ComputeReactiveActionPathParameterFromFn(fn func(key string) string) ComputeReactiveActionPathParameter {
	res := ComputeReactiveActionPathParameter{}
	if v := fn("id"); v != "" {
		t, _ := strconv.ParseInt(v, 10, 32)
		res.Id = int32(t)
	}
	if v := fn("age"); v != "" {
		t, _ := strconv.ParseInt(v, 10, 32)
		res.Age = int32(t)
	}
	return res
}

/**
 * Query parameters for ComputeReactiveAction
 */
// Query wrapper with private fields
type ComputeReactiveActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	QueryParam1   string `json:"queryParam1"`
	SecurityToken string `json:"securityToken"`
	Object1       struct {
		Field1 string `json:"field1"`
		Field2 string `json:"field2"`
	} `json:"object1"`
	IntSlice    []int `json:"intSlice"`
	InlineArray []struct {
		Slice1num  int       `json:"slice1num"`
		InnerSlice []float64 `json:"innerSlice"`
	} `json:"inlineArray"`
}

func ComputeReactiveActionQueryFromString(rawQuery string) ComputeReactiveActionQuery {
	v := ComputeReactiveActionQuery{}
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
func ComputeReactiveActionQueryFromGin(c *gin.Context) ComputeReactiveActionQuery {
	return ComputeReactiveActionQueryFromString(c.Request.URL.RawQuery)
}
func ComputeReactiveActionQueryFromHttp(r *http.Request) ComputeReactiveActionQuery {
	return ComputeReactiveActionQueryFromString(r.URL.RawQuery)
}
func (q ComputeReactiveActionQuery) Values() url.Values {
	return q.values
}
func (q ComputeReactiveActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ComputeReactiveActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ComputeReactiveActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

// WebSocket upgrader
var upgraderComputeReactiveAction = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type ComputeReactiveActionMessage struct {
	Raw         []byte
	Conn        *websocket.Conn
	MessageType int
	Error       error
	PathParams  ComputeReactiveActionPathParameter
}

// Developer handler type
type ComputeReactiveActionHandler func(msg ComputeReactiveActionMessage) error

// Generated handler
func ComputeReactiveAction(r *gin.Engine, handler ComputeReactiveActionHandler) {
	meta := ComputeReactiveActionMeta()
	r.GET(meta.URL, func(c *gin.Context) {
		pathParams := ComputeReactiveActionPathParameterFromGin(c)
		ws, err := upgraderComputeReactiveAction.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
			return
		}
		defer ws.Close()
		for {
			mt, raw, err := ws.ReadMessage()
			msg := ComputeReactiveActionMessage{
				Conn:        ws,
				Raw:         raw,
				Error:       err,
				MessageType: mt,
			}
			msg.PathParams = pathParams
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

type ComputeReactiveActionSession struct {
	Ctx         *gin.Context
	Socket      *websocket.Conn
	Done        chan bool
	Read        chan ComputeReactiveActionReadChan
	QueryParams ComputeReactiveActionQuery
}
type ComputeReactiveActionHandlerDuplex func(*ComputeReactiveActionSession)
type ComputeReactiveActionReadChan struct {
	Data  []byte
	Error error
}

func ComputeReactiveActionReactiveHandler(factory func(
	session ComputeReactiveActionSession,
) (chan []byte, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		read := make(chan ComputeReactiveActionReadChan)
		done := make(chan bool)
		c, err := upgraderComputeReactiveAction.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			c.Close()
			return
		}
		session := ComputeReactiveActionSession{
			Ctx:    ctx,
			Socket: c,
			Done:   done,
			Read:   read,
		}
		session.QueryParams = ComputeReactiveActionQueryFromGin(ctx)
		write, err := factory(session)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
		go func() {
			for {
				_, data, err := c.ReadMessage()
				read <- ComputeReactiveActionReadChan{
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

// ComputeReactiveActionClientSession is the client-side mirror of
// ComputeReactiveActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type ComputeReactiveActionClientSession struct {
	Socket *websocket.Conn
	Done   chan bool
	Read   chan ComputeReactiveActionReadChan
	Write  chan []byte
}

// ComputeReactiveActionClientOptions configures a client dial. All fields are
// optional — pass a zero value for a plaintext, header-less ws:// connection.
//
// TLSConfig governs the TLS handshake when dialing wss://. Set Certificates
// and RootCAs (and optionally ServerName) for mTLS; leave nil for ws://.
type ComputeReactiveActionClientOptions struct {
	Query     url.Values
	TLSConfig *tls.Config
	Headers   http.Header
}

// ComputeReactiveActionClient dials the ComputeReactiveAction endpoint at baseURL
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
//	sess, err := ComputeReactiveActionClient(
//	    "wss://hub.example.com",
//	    ComputeReactiveActionClientOptions{TLSConfig: tlsCfg},
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
func ComputeReactiveActionClient(baseURL string, opts *ComputeReactiveActionClientOptions) (*ComputeReactiveActionClientSession, error) {
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
	u.Path = ComputeReactiveActionMeta().URL
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
	session := &ComputeReactiveActionClientSession{
		Socket: c,
		Done:   make(chan bool, 1),
		Read:   make(chan ComputeReactiveActionReadChan),
		Write:  make(chan []byte, 16),
	}
	// Reader goroutine: pumps frames from the socket into Read. On error it
	// forwards the error frame, signals Done, and exits.
	go func() {
		for {
			_, data, err := c.ReadMessage()
			session.Read <- ComputeReactiveActionReadChan{
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
