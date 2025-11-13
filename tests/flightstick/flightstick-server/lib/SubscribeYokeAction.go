package unknownpackage

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Meta info for introspection
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Request object
type SubscribeYokeActionRequest struct {
	C  *gin.Context
	WS *websocket.Conn
}

// Core handler wrapper
func SubscribeYokeActionHandler(
	handler func(req SubscribeYokeActionRequest),
) (method, url string, h gin.HandlerFunc) {
	meta := SubscribeYokeActionMeta()
	return meta.Method, meta.URL, func(c *gin.Context) {
		fmt.Println("[SubscribeYokeAction] Incoming connection from:", c.ClientIP())

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("[SubscribeYokeAction] Upgrade failed:", err)
			c.String(http.StatusBadRequest, "WebSocket upgrade failed")
			return
		}
		fmt.Println("[SubscribeYokeAction] WebSocket connection established")

		defer func() {
			conn.Close()
			fmt.Println("[SubscribeYokeAction] Connection closed")
		}()

		fmt.Println("[SubscribeYokeAction] Invoking user handler")
		handler(SubscribeYokeActionRequest{C: c, WS: conn})
		fmt.Println("[SubscribeYokeAction] Handler returned cleanly")
	}
}

// Route registration helper
func SubscribeYokeAction(r gin.IRoutes, handler func(SubscribeYokeActionRequest)) {
	method, url, h := SubscribeYokeActionHandler(handler)
	fmt.Printf("[SubscribeYokeAction] Route registered: %s %s\n", method, url)
	r.Handle(method, url, h)
}
