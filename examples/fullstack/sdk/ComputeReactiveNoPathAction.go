package external
import (
"fmt"
"github.com/gin-gonic/gin"
"github.com/gorilla/websocket"
"github.com/torabian/emi/examples/fullstack/emigo"
"net/http"
"net/url"
)
/**
* Action to communicate with the action ComputeReactiveNoPathAction
*/
func ComputeReactiveNoPathActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "ComputeReactiveNoPathAction",
        URL:    "/compute/reactive",
        Method: "REACTIVE",
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
			QueryParam1 string `json:"queryParam1"`
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
	Raw []byte
	Conn *websocket.Conn	
	MessageType int
	Error error
	QueryParams ComputeReactiveNoPathActionQuery
}
// Developer handler type
type ComputeReactiveNoPathActionHandler func(msg ComputeReactiveNoPathActionMessage ) error
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
				Conn: ws,
				Raw: raw,
				Error: err,
				MessageType: mt,
			}
			msg.QueryParams = ComputeReactiveNoPathActionQueryFromGin(c)
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
	In   <-chan ComputeReactiveNoPathActionMessage
	Out  chan<- ComputeReactiveNoPathActionMessage
	Done <-chan struct{}
	Close func(err error)
	QueryParams ComputeReactiveNoPathActionQuery
}
type ComputeReactiveNoPathActionHandlerDuplex func(*ComputeReactiveNoPathActionSession)
// ComputeReactiveNoPathActionDuplex upgrades the HTTP connection to a WebSocket and
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
//	external.ComputeReactiveNoPathActionDuplex(r, func(ctx *external.ComputeReactiveNoPathActionSession) {
//		for {
//			select {
//			case msg, ok := <-ctx.In:
//				if !ok {
//					return // client disconnected
//				}
//				ctx.Out <- external.ComputeReactiveNoPathActionMessage{
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
func ComputeReactiveNoPathActionDuplex(r *gin.Engine, handler ComputeReactiveNoPathActionHandlerDuplex) {
	meta := ComputeReactiveNoPathActionMeta()
	// The actual callback is extracted, in case you need to handle multiple handlers or customize, use it directly.
	r.GET(meta.URL, func(ctx *gin.Context) {
		ComputeReactiveNoPathActionDuplexGinHandler(ctx, handler)
	})
}
func ComputeReactiveNoPathActionDuplexGinHandler(c *gin.Context, handler ComputeReactiveNoPathActionHandlerDuplex) {
	ws, err := upgraderComputeReactiveNoPathAction.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
		return
	}
	in := make(chan ComputeReactiveNoPathActionMessage)
	out := make(chan ComputeReactiveNoPathActionMessage)
	done := make(chan struct{})
	session := &ComputeReactiveNoPathActionSession{
		In:   in,
		Out:  out,
		Done: done,
		Close: func(err error) {
			close(done)
			ws.Close()
		},
	}
	session.QueryParams = ComputeReactiveNoPathActionQueryFromGin(c)
	// Read loop
	go func() {
		defer close(in)
		for {
			mt, raw, err := ws.ReadMessage()
			in <- ComputeReactiveNoPathActionMessage{MessageType: mt, Raw: raw, Error: err}
		}
	}()
	// Write loop
	go func() {
		for msg := range out {
			if err := ws.WriteMessage(msg.MessageType, msg.Raw); err != nil {
				// When message is -1, means it's internal error coming out
				in <- ComputeReactiveNoPathActionMessage{MessageType: -1, Error: err}
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