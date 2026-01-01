package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
)

/**
* Action to communicate with the action GetEventsAboutTheSellerSOffersAction
 */
func GetEventsAboutTheSellerSOffersActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "GetEventsAboutTheSellerSOffersAction",
		URL:    "https://api.{environment}/sale/offer-events",
		Method: "GET",
	}
}

// The base class definition for getEventsAboutTheSellerSOffersActionRes
type GetEventsAboutTheSellerSOffersActionRes struct {
	// List of events related to offer state changes
	OfferEvents []GetEventsAboutTheSellerSOffersActionResOfferEvents `json:"offerEvents" yaml:"offerEvents"`
}

// The base class definition for offerEvents
type GetEventsAboutTheSellerSOffersActionResOfferEvents struct {
	// Unique event identifier (base64 encoded)
	Id string `json:"id" yaml:"id"`
	// ISO8601 timestamp when the event occurred
	OccurredAt string `json:"occurredAt" yaml:"occurredAt"`
	// Event type (e.g., OFFER_ACTIVATED, OFFER_ENDED, etc.)
	Type string `json:"type" yaml:"type"`
	// Basic offer information for which event occurred
	Offer GetEventsAboutTheSellerSOffersActionResOfferEventsOffer `json:"offer" yaml:"offer"`
}

// The base class definition for offer
type GetEventsAboutTheSellerSOffersActionResOfferEventsOffer struct {
	Id       string                                                          `json:"id" yaml:"id"`
	External GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal `json:"external" yaml:"external"`
}

// The base class definition for external
type GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal struct {
	Id string `json:"id" yaml:"id"`
}
type GetEventsAboutTheSellerSOffersActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// GetEventsAboutTheSellerSOffersActionRaw registers a raw Gin route for the GetEventsAboutTheSellerSOffersAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetEventsAboutTheSellerSOffersActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetEventsAboutTheSellerSOffersActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // GetEventsAboutTheSellerSOffersActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetEventsAboutTheSellerSOffersAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetEventsAboutTheSellerSOffersActionHandler(
	handler func(c GetEventsAboutTheSellerSOffersActionRequest, gin *gin.Context) (*GetEventsAboutTheSellerSOffersActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetEventsAboutTheSellerSOffersActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetEventsAboutTheSellerSOffersActionRequest{
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
		}
		resp, err := handler(req, m)
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

// GetEventsAboutTheSellerSOffersAction is a high-level convenience wrapper around GetEventsAboutTheSellerSOffersActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetEventsAboutTheSellerSOffersAction(r gin.IRoutes, handler func(c GetEventsAboutTheSellerSOffersActionRequest, gin *gin.Context) (*GetEventsAboutTheSellerSOffersActionResponse, error)) {
	method, url, h := GetEventsAboutTheSellerSOffersActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get events about the seller's offersAction
 */
// Query wrapper with private fields
type GetEventsAboutTheSellerSOffersActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetEventsAboutTheSellerSOffersActionQueryFromString(rawQuery string) GetEventsAboutTheSellerSOffersActionQuery {
	v := GetEventsAboutTheSellerSOffersActionQuery{}
	values, _ := url.ParseQuery(rawQuery)
	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}
	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json", // reuse json tags
		WeaklyTypedInput: true,   // "1" -> int, "true" -> bool
		Result:           &v,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}
	v.values = values
	v.mapped = mapped
	return v
}
func GetEventsAboutTheSellerSOffersActionQueryFromGin(c *gin.Context) GetEventsAboutTheSellerSOffersActionQuery {
	return GetEventsAboutTheSellerSOffersActionQueryFromString(c.Request.URL.RawQuery)
}
func GetEventsAboutTheSellerSOffersActionQueryFromHttp(r *http.Request) GetEventsAboutTheSellerSOffersActionQuery {
	return GetEventsAboutTheSellerSOffersActionQueryFromString(r.URL.RawQuery)
}
func (q GetEventsAboutTheSellerSOffersActionQuery) Values() url.Values {
	return q.values
}
func (q GetEventsAboutTheSellerSOffersActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetEventsAboutTheSellerSOffersActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetEventsAboutTheSellerSOffersActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetEventsAboutTheSellerSOffersActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type GetEventsAboutTheSellerSOffersActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func GetEventsAboutTheSellerSOffersActionCall(
	req GetEventsAboutTheSellerSOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetEventsAboutTheSellerSOffersActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetEventsAboutTheSellerSOffersActionMeta()
		baseURL := meta.URL
		// Build final URL with query string
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}
		// if UrlValues present, encode and append
		if len(req.QueryParams) > 0 {
			u.RawQuery = req.QueryParams.Encode()
		}
		req0, err := http.NewRequest(meta.Method, u.String(), nil)
		if err != nil {
			return nil, err
		}
		httpReq = req0
	} else {
		httpReq = config.Httpr
	}
	httpReq.Header = req.Headers
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var result GetEventsAboutTheSellerSOffersActionResult
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	if resp.StatusCode >= 400 {
		return &result, fmt.Errorf("request failed: %s", respBody)
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &result, err
	}
	return &result, nil
}
