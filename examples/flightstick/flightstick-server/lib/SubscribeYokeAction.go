package unknownpackage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

/**
* Action to communicate with the action SubscribeYokeAction
 */
func SubscribeYokeActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "SubscribeYokeAction",
		URL:    "/subscribe-yoke",
		Method: "GET", // WebSocket handshake must be GET
	}
}

// WebSocket upgrader (you can reuse a global one)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // you can restrict origin later
}

// SubscribeYokeActionRequest provides access to both gin context and websocket connection
type SubscribeYokeActionRequest struct {
	C  *gin.Context
	WS *websocket.Conn
}

// SubscribeYokeActionHandler sets up a typed WebSocket handler
func SubscribeYokeActionHandler(
	handler func(req SubscribeYokeActionRequest),
) (method, url string, h gin.HandlerFunc) {
	meta := SubscribeYokeActionMeta()
	return meta.Method, meta.URL, func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.String(http.StatusBadRequest, "Failed to upgrade: %v", err)
			return
		}
		defer conn.Close()

		handler(SubscribeYokeActionRequest{C: c, WS: conn})
	}
}

// SubscribeYokeAction auto-registers the WebSocket route
func SubscribeYokeAction(r gin.IRoutes, handler func(SubscribeYokeActionRequest)) {
	method, url, h := SubscribeYokeActionHandler(handler)
	r.Handle(method, url, h)
}
