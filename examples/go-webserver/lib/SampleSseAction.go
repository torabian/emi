package unknownpackage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* Action to communicate with the action SampleSseAction
 */
func SampleSseActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "SampleSseAction",
		URL:    "/stream",
		Method: "GET",
	}
}

// SampleSseActionRequest wraps the current Gin context.
// It can be extended later to include typed request data (headers, body, params, etc.)
// so developers can access both the raw context and strongly typed fields.
type SampleSseActionRequest struct {
	C *gin.Context
}
type SampleSseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// SampleSseActionRaw registers a raw Gin route for the SampleSseAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func SampleSseActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := SampleSseActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// SampleSseActionHandler returns the HTTP method, route URL, and a typed Gin handler for the SampleSseAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func SampleSseActionHandler(
	handler func(c SampleSseActionRequest) (*SampleSseActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := SampleSseActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		req := SampleSseActionRequest{C: m}
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

// SampleSseAction is a high-level convenience wrapper around SampleSseActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func SampleSseAction(r gin.IRoutes, handler func(SampleSseActionRequest) (*SampleSseActionResponse, error)) {
	method, url, h := SampleSseActionHandler(handler)
	r.Handle(method, url, h)
}
