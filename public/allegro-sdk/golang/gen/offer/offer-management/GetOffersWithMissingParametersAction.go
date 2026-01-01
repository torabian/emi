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
* Action to communicate with the action GetOffersWithMissingParametersAction
 */
func GetOffersWithMissingParametersActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "GetOffersWithMissingParametersAction",
		URL:    "https://api.{environment}/sale/offers/unfilled-parameters",
		Method: "GET",
	}
}

// The base class definition for getOffersWithMissingParametersActionRes
type GetOffersWithMissingParametersActionRes struct {
	Offers     []GetOffersWithMissingParametersActionResOffers `json:"offers" yaml:"offers"`
	Count      int                                             `json:"count" yaml:"count"`
	TotalCount int                                             `json:"totalCount" yaml:"totalCount"`
}

// The base class definition for offers
type GetOffersWithMissingParametersActionResOffers struct {
	Id         string                                                    `json:"id" yaml:"id"`
	Parameters []GetOffersWithMissingParametersActionResOffersParameters `json:"parameters" yaml:"parameters"`
	Category   GetOffersWithMissingParametersActionResOffersCategory     `json:"category" yaml:"category"`
}

// The base class definition for parameters
type GetOffersWithMissingParametersActionResOffersParameters struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for category
type GetOffersWithMissingParametersActionResOffersCategory struct {
	Id string `json:"id" yaml:"id"`
}
type GetOffersWithMissingParametersActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// GetOffersWithMissingParametersActionRaw registers a raw Gin route for the GetOffersWithMissingParametersAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetOffersWithMissingParametersActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetOffersWithMissingParametersActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // GetOffersWithMissingParametersActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetOffersWithMissingParametersAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetOffersWithMissingParametersActionHandler(
	handler func(c GetOffersWithMissingParametersActionRequest, gin *gin.Context) (*GetOffersWithMissingParametersActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetOffersWithMissingParametersActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetOffersWithMissingParametersActionRequest{
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

// GetOffersWithMissingParametersAction is a high-level convenience wrapper around GetOffersWithMissingParametersActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetOffersWithMissingParametersAction(r gin.IRoutes, handler func(c GetOffersWithMissingParametersActionRequest, gin *gin.Context) (*GetOffersWithMissingParametersActionResponse, error)) {
	method, url, h := GetOffersWithMissingParametersActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get offers with missing parametersAction
 */
// Query wrapper with private fields
type GetOffersWithMissingParametersActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetOffersWithMissingParametersActionQueryFromString(rawQuery string) GetOffersWithMissingParametersActionQuery {
	v := GetOffersWithMissingParametersActionQuery{}
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
func GetOffersWithMissingParametersActionQueryFromGin(c *gin.Context) GetOffersWithMissingParametersActionQuery {
	return GetOffersWithMissingParametersActionQueryFromString(c.Request.URL.RawQuery)
}
func GetOffersWithMissingParametersActionQueryFromHttp(r *http.Request) GetOffersWithMissingParametersActionQuery {
	return GetOffersWithMissingParametersActionQueryFromString(r.URL.RawQuery)
}
func (q GetOffersWithMissingParametersActionQuery) Values() url.Values {
	return q.values
}
func (q GetOffersWithMissingParametersActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetOffersWithMissingParametersActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetOffersWithMissingParametersActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetOffersWithMissingParametersActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type GetOffersWithMissingParametersActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func GetOffersWithMissingParametersActionCall(
	req GetOffersWithMissingParametersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetOffersWithMissingParametersActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetOffersWithMissingParametersActionMeta()
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
	var result GetOffersWithMissingParametersActionResult
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
