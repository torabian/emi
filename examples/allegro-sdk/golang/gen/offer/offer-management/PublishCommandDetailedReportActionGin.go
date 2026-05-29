package external

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// PublishCommandDetailedReportActionRaw registers a raw Gin route for the PublishCommandDetailedReportAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PublishCommandDetailedReportActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PublishCommandDetailedReportActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// PublishCommandDetailedReportActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PublishCommandDetailedReportAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PublishCommandDetailedReportActionHandler(
	handler func(c PublishCommandDetailedReportActionRequest) (*PublishCommandDetailedReportActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PublishCommandDetailedReportActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PublishCommandDetailedReportActionRequest{
			Body:        nil,
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx:      m,
		}
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

// PublishCommandDetailedReportActionGin is a high-level convenience wrapper around PublishCommandDetailedReportActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PublishCommandDetailedReportActionGin(r gin.IRoutes, handler func(c PublishCommandDetailedReportActionRequest) (*PublishCommandDetailedReportActionResponse, error)) {
	method, url, h := PublishCommandDetailedReportActionHandler(handler)
	r.Handle(method, url, h)
}
func (x PublishCommandDetailedReportActionRequest) IsGin() bool {
	if x.GinCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.GinCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}
func PublishCommandDetailedReportActionQueryFromGin(c *gin.Context) PublishCommandDetailedReportActionQuery {
	return PublishCommandDetailedReportActionQueryFromString(c.Request.URL.RawQuery)
}
