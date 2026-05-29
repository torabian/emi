//go:build !wasm

package defs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// DeleteUserActionRaw registers a raw Gin route for the DeleteUserAction action.
// This gives the developer full control over middleware, handlers, and response handling.
//
func DeleteUserActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := DeleteUserActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// DeleteUserActionHandler returns the HTTP method, route URL, and a typed Gin handler for the DeleteUserAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func DeleteUserActionHandler(
	handler func(c DeleteUserActionRequest) (*DeleteUserActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := DeleteUserActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body DeleteUserActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := DeleteUserActionRequest{
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

// DeleteUserActionGin is a high-level convenience wrapper around DeleteUserActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func DeleteUserActionGin(r gin.IRoutes, handler func(c DeleteUserActionRequest) (*DeleteUserActionResponse, error)) {
	method, url, h := DeleteUserActionHandler(handler)
	r.Handle(method, url, h)
}
func (x DeleteUserActionRequest) IsGin() bool {
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
func DeleteUserActionQueryFromGin(c *gin.Context) DeleteUserActionQuery {
	return DeleteUserActionQueryFromString(c.Request.URL.RawQuery)
}
