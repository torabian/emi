package external
import (
"fmt"
"github.com/gin-gonic/gin"
"github.com/gorilla/websocket"
"github.com/torabian/emi/examples/fullstack/emigo"
"net/http"
"net/url"
"strconv"
"strings"
)
/**
* Action to communicate with the action ComputeReactiveAction
*/
func ComputeReactiveActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "ComputeReactiveAction",
        URL:    "/compute/reactive/:id/:age",
        Method: "REACTIVE",
    }
}
	/**
 * Path parameters for ComputeReactiveAction
 */
type ComputeReactiveActionPathParameter struct {
	Id int32
	Age int32
}
// Converts a placeholder url, and applies the parameters to it.
func ComputeReactiveActionPathParameterApply(params ComputeReactiveActionPathParameter, templateUrl string) string {
		templateUrl = strings.ReplaceAll(templateUrl, "id", fmt.Sprintf("%v", params.Id))
		templateUrl = strings.ReplaceAll(templateUrl, "age", fmt.Sprintf("%v", params.Age))
	return templateUrl
}
// Creates the parameters from the gin
// Creates the parameters from the gin
func ComputeReactiveActionPathParameterFromGin(g *gin.Context) ComputeReactiveActionPathParameter {
	res := ComputeReactiveActionPathParameter{}
			if v := g.Param("id"); v != "" {
					t, _ := strconv.ParseInt(v, 10, 32)
					res.Id = int32(t)
			}
			if v := g.Param("age"); v != "" {
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
			QueryParam1 string `json:"queryParam1"`
			SecurityToken string `json:"securityToken"`
			Object1 struct {
			Field1 string `json:"field1"`
			Field2 string `json:"field2"`
			} `json:"object1"`
			IntSlice []int `json:"intSlice"`
			InlineArray [] struct {
			Slice1num int `json:"slice1num"`
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
	Raw []byte
	Conn *websocket.Conn	
	MessageType int
	Error error
	PathParams ComputeReactiveActionPathParameter
	QueryParams ComputeReactiveActionQuery
}
// Developer handler type
type ComputeReactiveActionHandler func(msg ComputeReactiveActionMessage ) error
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
				Conn: ws,
				Raw: raw,
				Error: err,
				MessageType: mt,
			}
			msg.PathParams = pathParams
			msg.QueryParams = ComputeReactiveActionQueryFromGin(c)
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
	In   <-chan ComputeReactiveActionMessage
	Out  chan<- ComputeReactiveActionMessage
	Done <-chan struct{}
	Close func(err error)
	PathParams ComputeReactiveActionPathParameter
	QueryParams ComputeReactiveActionQuery
}
type ComputeReactiveActionHandlerDuplex func(*ComputeReactiveActionSession)
// ComputeReactiveActionDuplex upgrades the HTTP connection to a WebSocket and
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
//	external.ComputeReactiveActionDuplex(r, func(ctx *external.ComputeReactiveActionSession) {
//		for {
//			select {
//			case msg, ok := <-ctx.In:
//				if !ok {
//					return // client disconnected
//				}
//				ctx.Out <- external.ComputeReactiveActionMessage{
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
func ComputeReactiveActionDuplex(r *gin.Engine, handler ComputeReactiveActionHandlerDuplex) {
	meta := ComputeReactiveActionMeta()
	// The actual callback is extracted, in case you need to handle multiple handlers or customize, use it directly.
	r.GET(meta.URL, func(ctx *gin.Context) {
		ComputeReactiveActionDuplexGinHandler(ctx, handler)
	})
}
func ComputeReactiveActionDuplexGinHandler(c *gin.Context, handler ComputeReactiveActionHandlerDuplex) {
	pathParams := ComputeReactiveActionPathParameterFromGin(c)
	ws, err := upgraderComputeReactiveAction.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
		return
	}
	in := make(chan ComputeReactiveActionMessage)
	out := make(chan ComputeReactiveActionMessage)
	done := make(chan struct{})
	session := &ComputeReactiveActionSession{
		In:   in,
		Out:  out,
		Done: done,
		Close: func(err error) {
			close(done)
			ws.Close()
		},
	}
		session.PathParams = pathParams
	session.QueryParams = ComputeReactiveActionQueryFromGin(c)
	// Read loop
	go func() {
		defer close(in)
		for {
			mt, raw, err := ws.ReadMessage()
			in <- ComputeReactiveActionMessage{MessageType: mt, Raw: raw, Error: err}
		}
	}()
	// Write loop
	go func() {
		for msg := range out {
			if err := ws.WriteMessage(msg.MessageType, msg.Raw); err != nil {
				// When message is -1, means it's internal error coming out
				in <- ComputeReactiveActionMessage{MessageType: -1, Error: err}
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