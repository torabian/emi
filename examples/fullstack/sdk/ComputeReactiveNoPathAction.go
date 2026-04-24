package external

import (
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
