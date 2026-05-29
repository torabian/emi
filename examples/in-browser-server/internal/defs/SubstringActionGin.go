//go:build !wasm

package defs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// SubstringActionRaw registers a raw Gin route for the SubstringAction action.
// This gives the developer full control over middleware, handlers, and response handling.
//
func SubstringActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := SubstringActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// SubstringActionHandler returns the HTTP method, route URL, and a typed Gin handler for the SubstringAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func SubstringActionHandler(
	handler func(c SubstringActionRequest) (*SubstringActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := SubstringActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body SubstringActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := SubstringActionRequest{
			Body:        body,
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

// SubstringActionGin is a high-level convenience wrapper around SubstringActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func SubstringActionGin(r gin.IRoutes, handler func(c SubstringActionRequest) (*SubstringActionResponse, error)) {
	method, url, h := SubstringActionHandler(handler)
	r.Handle(method, url, h)
}
func (x SubstringActionRequest) IsGin() bool {
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
func SubstringActionQueryFromGin(c *gin.Context) SubstringActionQuery {
	return SubstringActionQueryFromString(c.Request.URL.RawQuery)
}
