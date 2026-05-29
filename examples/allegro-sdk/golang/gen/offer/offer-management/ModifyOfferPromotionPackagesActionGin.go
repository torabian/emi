package external

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// ModifyOfferPromotionPackagesActionRaw registers a raw Gin route for the ModifyOfferPromotionPackagesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModifyOfferPromotionPackagesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModifyOfferPromotionPackagesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// ModifyOfferPromotionPackagesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModifyOfferPromotionPackagesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModifyOfferPromotionPackagesActionHandler(
	handler func(c ModifyOfferPromotionPackagesActionRequest) (*ModifyOfferPromotionPackagesActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ModifyOfferPromotionPackagesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ModifyOfferPromotionPackagesActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ModifyOfferPromotionPackagesActionRequest{
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

// ModifyOfferPromotionPackagesActionGin is a high-level convenience wrapper around ModifyOfferPromotionPackagesActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModifyOfferPromotionPackagesActionGin(r gin.IRoutes, handler func(c ModifyOfferPromotionPackagesActionRequest) (*ModifyOfferPromotionPackagesActionResponse, error)) {
	method, url, h := ModifyOfferPromotionPackagesActionHandler(handler)
	r.Handle(method, url, h)
}
func (x ModifyOfferPromotionPackagesActionRequest) IsGin() bool {
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
func ModifyOfferPromotionPackagesActionQueryFromGin(c *gin.Context) ModifyOfferPromotionPackagesActionQuery {
	return ModifyOfferPromotionPackagesActionQueryFromString(c.Request.URL.RawQuery)
}
