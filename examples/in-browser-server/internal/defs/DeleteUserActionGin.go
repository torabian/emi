//go:build !wasm

package defs

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"reflect"
)

// DeleteUserActionRaw registers a raw Gin route for the DeleteUserAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func DeleteUserActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := DeleteUserActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// DeleteUserActionHandler returns the HTTP method, route URL, and a typed Gin handler for the DeleteUserAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. Body binding (JSON/YAML/XML/form), headers,
// errors, and the success response are all handled by emigo - see BindGinRequestBody,
// RenderGinError and RenderGinResult in github.com/torabian/emi/emigo.
func DeleteUserActionHandler(
	handler func(c DeleteUserActionRequest) (*DeleteUserActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := DeleteUserActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body DeleteUserActionReq
		if err := emigo.BindGinRequestBody(m, &body); err != nil {
			emigo.RenderGinError(m, err)
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
