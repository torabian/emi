//go:build !wasm

package defs

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"reflect"
)

// ListUsersActionRaw registers a raw Gin route for the ListUsersAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ListUsersActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ListUsersActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// ListUsersActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ListUsersAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. Body binding (JSON/YAML/XML/form), headers,
// errors, and the success response are all handled by emigo - see BindGinRequestBody,
// RenderGinError and RenderGinResult in github.com/torabian/emi/emigo.
func ListUsersActionHandler(
	handler func(c ListUsersActionRequest) (*ListUsersActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ListUsersActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := ListUsersActionRequest{
			Body:        nil,
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx:      m,
		}
		resp, err := handler(req)
		if err != nil {
			emigo.RenderGinError(m, err)
			return
		}
		// If the handler returned nil (and no error), it means the response was handled manually.
		if resp == nil {
			return
		}
		emigo.RenderGinResult(m, resp)
	}
}

// ListUsersActionGin is a high-level convenience wrapper around ListUsersActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ListUsersActionGin(r gin.IRoutes, handler func(c ListUsersActionRequest) (*ListUsersActionResponse, error)) {
	method, url, h := ListUsersActionHandler(handler)
	r.Handle(method, url, h)
}
func (x ListUsersActionRequest) IsGin() bool {
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
func ListUsersActionQueryFromGin(c *gin.Context) ListUsersActionQuery {
	return ListUsersActionQueryFromString(c.Request.URL.RawQuery)
}
