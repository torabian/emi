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
