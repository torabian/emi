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
* Action to communicate with the action ModificationCommandSummaryAction
 */
func ModificationCommandSummaryActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "ModificationCommandSummaryAction",
		URL:    "https://api.{environment}/sale/offers/promo-options-commands/{commandId}",
		Method: "GET",
	}
}

// The base class definition for modificationCommandSummaryActionRes
type ModificationCommandSummaryActionRes struct {
	Id        string                                       `json:"id" yaml:"id"`
	TaskCount ModificationCommandSummaryActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}

// The base class definition for taskCount
type ModificationCommandSummaryActionResTaskCount struct {
	Failed  int `json:"failed" yaml:"failed"`
	Success int `json:"success" yaml:"success"`
	Total   int `json:"total" yaml:"total"`
}
type ModificationCommandSummaryActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// ModificationCommandSummaryActionRaw registers a raw Gin route for the ModificationCommandSummaryAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModificationCommandSummaryActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModificationCommandSummaryActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // ModificationCommandSummaryActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModificationCommandSummaryAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModificationCommandSummaryActionHandler(
	handler func(c ModificationCommandSummaryActionRequest, gin *gin.Context) (*ModificationCommandSummaryActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ModificationCommandSummaryActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := ModificationCommandSummaryActionRequest{
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

// ModificationCommandSummaryAction is a high-level convenience wrapper around ModificationCommandSummaryActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModificationCommandSummaryAction(r gin.IRoutes, handler func(c ModificationCommandSummaryActionRequest, gin *gin.Context) (*ModificationCommandSummaryActionResponse, error)) {
	method, url, h := ModificationCommandSummaryActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Modification command summaryAction
 */
// Query wrapper with private fields
type ModificationCommandSummaryActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModificationCommandSummaryActionQueryFromString(rawQuery string) ModificationCommandSummaryActionQuery {
	v := ModificationCommandSummaryActionQuery{}
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
func ModificationCommandSummaryActionQueryFromGin(c *gin.Context) ModificationCommandSummaryActionQuery {
	return ModificationCommandSummaryActionQueryFromString(c.Request.URL.RawQuery)
}
func ModificationCommandSummaryActionQueryFromHttp(r *http.Request) ModificationCommandSummaryActionQuery {
	return ModificationCommandSummaryActionQueryFromString(r.URL.RawQuery)
}
func (q ModificationCommandSummaryActionQuery) Values() url.Values {
	return q.values
}
func (q ModificationCommandSummaryActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModificationCommandSummaryActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModificationCommandSummaryActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModificationCommandSummaryActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type ModificationCommandSummaryActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ModificationCommandSummaryActionCall(
	req ModificationCommandSummaryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModificationCommandSummaryActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ModificationCommandSummaryActionMeta()
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
	var result ModificationCommandSummaryActionResult
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
