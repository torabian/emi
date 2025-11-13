package unknownpackage
import (
"github.com/gin-gonic/gin"
"net/http"
)
/**
* Action to communicate with the action WebSocketOrgEchoAction
*/
func WebSocketOrgEchoActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "WebSocketOrgEchoAction",
        URL:    "wss://echo.websocket.org/.ws",
        Method: "reactive",
    }
}
// WebSocketOrgEchoActionRequest wraps the current Gin context.
// It can be extended later to include typed request data (headers, body, params, etc.)
// so developers can access both the raw context and strongly typed fields.
type WebSocketOrgEchoActionRequest struct {
	c *gin.Context
}
type WebSocketOrgEchoActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// WebSocketOrgEchoActionRaw registers a raw Gin route for the WebSocketOrgEchoAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func WebSocketOrgEchoActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := WebSocketOrgEchoActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}
// WebSocketOrgEchoActionHandler returns the HTTP method, route URL, and a typed Gin handler for the WebSocketOrgEchoAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func WebSocketOrgEchoActionHandler(
	handler func(c WebSocketOrgEchoActionRequest) (*WebSocketOrgEchoActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := WebSocketOrgEchoActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		req := WebSocketOrgEchoActionRequest{c: m}
		resp, err := handler(req)
		if err != nil {
			m.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// If the handler returned nil (and no error), it means the response was handled manually.
		if resp == nil {
			return
		}
		// Apply headers
		for k, v := range resp.Headers {
			m.Header(k, v)
		}
		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}
		if resp.Payload != nil {
			m.JSON(status, resp.Payload)
		} else {
			m.Status(status)
		}
	}
}
// WebSocketOrgEchoAction is a high-level convenience wrapper around WebSocketOrgEchoActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func WebSocketOrgEchoAction(r gin.IRoutes, handler func(WebSocketOrgEchoActionRequest) (*WebSocketOrgEchoActionResponse, error)) {
	method, url, h := WebSocketOrgEchoActionHandler(handler)
	r.Handle(method, url, h)
}