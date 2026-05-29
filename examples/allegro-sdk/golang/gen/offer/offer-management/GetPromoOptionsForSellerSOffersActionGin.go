package external

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// GetPromoOptionsForSellerSOffersActionRaw registers a raw Gin route for the GetPromoOptionsForSellerSOffersAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetPromoOptionsForSellerSOffersActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetPromoOptionsForSellerSOffersActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// GetPromoOptionsForSellerSOffersActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetPromoOptionsForSellerSOffersAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetPromoOptionsForSellerSOffersActionHandler(
	handler func(c GetPromoOptionsForSellerSOffersActionRequest) (*GetPromoOptionsForSellerSOffersActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetPromoOptionsForSellerSOffersActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetPromoOptionsForSellerSOffersActionRequest{
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

// GetPromoOptionsForSellerSOffersActionGin is a high-level convenience wrapper around GetPromoOptionsForSellerSOffersActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetPromoOptionsForSellerSOffersActionGin(r gin.IRoutes, handler func(c GetPromoOptionsForSellerSOffersActionRequest) (*GetPromoOptionsForSellerSOffersActionResponse, error)) {
	method, url, h := GetPromoOptionsForSellerSOffersActionHandler(handler)
	r.Handle(method, url, h)
}
func (x GetPromoOptionsForSellerSOffersActionRequest) IsGin() bool {
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
func GetPromoOptionsForSellerSOffersActionQueryFromGin(c *gin.Context) GetPromoOptionsForSellerSOffersActionQuery {
	return GetPromoOptionsForSellerSOffersActionQueryFromString(c.Request.URL.RawQuery)
}
