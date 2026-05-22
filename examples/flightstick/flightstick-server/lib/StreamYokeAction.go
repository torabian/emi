package unknownpackage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

/**
* Action to communicate with the action StreamYokeAction
 */
func StreamYokeActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "StreamYokeAction",
		URL:    "/stream-yoke",
		Method: "GET", // WebSocket handshake
	}
}

// StreamYokeActionRequest provides both the gin context and websocket connection
type StreamYokeActionRequest struct {
	C  *gin.Context
	WS *websocket.Conn
}

// StreamYokeActionHandler registers the WebSocket handler
func StreamYokeActionHandler(
	handler func(req StreamYokeActionRequest),
) (method, url string, h gin.HandlerFunc) {
	meta := StreamYokeActionMeta()
	return meta.Method, meta.URL, func(c *gin.Context) {
		var upgrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.String(http.StatusBadRequest, "Failed to upgrade: %v", err)
			return
		}
		defer conn.Close()
		handler(StreamYokeActionRequest{C: c, WS: conn})
	}
}

// StreamYokeAction automatically registers the route
func StreamYokeAction(r gin.IRoutes, handler func(StreamYokeActionRequest)) {
	method, url, h := StreamYokeActionHandler(handler)
	r.Handle(method, url, h)
}
