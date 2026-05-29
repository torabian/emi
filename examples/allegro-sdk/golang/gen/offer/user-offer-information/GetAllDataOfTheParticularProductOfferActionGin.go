package external

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// GetAllDataOfTheParticularProductOfferActionRaw registers a raw Gin route for the GetAllDataOfTheParticularProductOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetAllDataOfTheParticularProductOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// GetAllDataOfTheParticularProductOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetAllDataOfTheParticularProductOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetAllDataOfTheParticularProductOfferActionHandler(
	handler func(c GetAllDataOfTheParticularProductOfferActionRequest) (*GetAllDataOfTheParticularProductOfferActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetAllDataOfTheParticularProductOfferActionRequest{
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

// GetAllDataOfTheParticularProductOfferActionGin is a high-level convenience wrapper around GetAllDataOfTheParticularProductOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetAllDataOfTheParticularProductOfferActionGin(r gin.IRoutes, handler func(c GetAllDataOfTheParticularProductOfferActionRequest) (*GetAllDataOfTheParticularProductOfferActionResponse, error)) {
	method, url, h := GetAllDataOfTheParticularProductOfferActionHandler(handler)
	r.Handle(method, url, h)
}
func (x GetAllDataOfTheParticularProductOfferActionRequest) IsGin() bool {
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
func GetAllDataOfTheParticularProductOfferActionQueryFromGin(c *gin.Context) GetAllDataOfTheParticularProductOfferActionQuery {
	return GetAllDataOfTheParticularProductOfferActionQueryFromString(c.Request.URL.RawQuery)
}
