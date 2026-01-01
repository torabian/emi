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
* Action to communicate with the action PublishCommandSummaryAction
 */
func PublishCommandSummaryActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "PublishCommandSummaryAction",
		URL:    "https://api.{environment}/sale/offer-publication-commands/{commandId}",
		Method: "GET",
	}
}

// The base class definition for publishCommandSummaryActionRes
type PublishCommandSummaryActionRes struct {
	Id          string                                  `json:"id" yaml:"id"`
	CreatedAt   string                                  `json:"createdAt" yaml:"createdAt"`
	CompletedAt string                                  `json:"completedAt" yaml:"completedAt"`
	TaskCount   PublishCommandSummaryActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}

// The base class definition for taskCount
type PublishCommandSummaryActionResTaskCount struct {
	Failed  int `json:"failed" yaml:"failed"`
	Success int `json:"success" yaml:"success"`
	Total   int `json:"total" yaml:"total"`
}
type PublishCommandSummaryActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// PublishCommandSummaryActionRaw registers a raw Gin route for the PublishCommandSummaryAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PublishCommandSummaryActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PublishCommandSummaryActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // PublishCommandSummaryActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PublishCommandSummaryAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PublishCommandSummaryActionHandler(
	handler func(c PublishCommandSummaryActionRequest, gin *gin.Context) (*PublishCommandSummaryActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PublishCommandSummaryActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PublishCommandSummaryActionRequest{
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

// PublishCommandSummaryAction is a high-level convenience wrapper around PublishCommandSummaryActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PublishCommandSummaryAction(r gin.IRoutes, handler func(c PublishCommandSummaryActionRequest, gin *gin.Context) (*PublishCommandSummaryActionResponse, error)) {
	method, url, h := PublishCommandSummaryActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Publish command summaryAction
 */
// Query wrapper with private fields
type PublishCommandSummaryActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func PublishCommandSummaryActionQueryFromString(rawQuery string) PublishCommandSummaryActionQuery {
	v := PublishCommandSummaryActionQuery{}
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
func PublishCommandSummaryActionQueryFromGin(c *gin.Context) PublishCommandSummaryActionQuery {
	return PublishCommandSummaryActionQueryFromString(c.Request.URL.RawQuery)
}
func PublishCommandSummaryActionQueryFromHttp(r *http.Request) PublishCommandSummaryActionQuery {
	return PublishCommandSummaryActionQueryFromString(r.URL.RawQuery)
}
func (q PublishCommandSummaryActionQuery) Values() url.Values {
	return q.values
}
func (q PublishCommandSummaryActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PublishCommandSummaryActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PublishCommandSummaryActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PublishCommandSummaryActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type PublishCommandSummaryActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func PublishCommandSummaryActionCall(
	req PublishCommandSummaryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PublishCommandSummaryActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := PublishCommandSummaryActionMeta()
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
	var result PublishCommandSummaryActionResult
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
