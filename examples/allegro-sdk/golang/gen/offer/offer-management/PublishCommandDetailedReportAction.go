package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"test.com/emi/golang/emigo"
)

/**
* Action to communicate with the action PublishCommandDetailedReportAction
 */
func PublishCommandDetailedReportActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "PublishCommandDetailedReportAction",
		URL:    "https://api.{environment}/sale/offer-publication-commands/{commandId}/tasks",
		Method: "GET",
	}
}

// The base class definition for publishCommandDetailedReportActionRes
type PublishCommandDetailedReportActionRes struct {
	Tasks []PublishCommandDetailedReportActionResTasks `json:"tasks" yaml:"tasks"`
}

// The base class definition for tasks
type PublishCommandDetailedReportActionResTasks struct {
	Field   string                                             `json:"field" yaml:"field"`
	Message string                                             `json:"message" yaml:"message"`
	Offer   PublishCommandDetailedReportActionResTasksOffer    `json:"offer" yaml:"offer"`
	Status  string                                             `json:"status" yaml:"status"`
	Errors  []PublishCommandDetailedReportActionResTasksErrors `json:"errors" yaml:"errors"`
}

// The base class definition for offer
type PublishCommandDetailedReportActionResTasksOffer struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for errors
type PublishCommandDetailedReportActionResTasksErrors struct {
	Code        string                                                   `json:"code" yaml:"code"`
	Details     string                                                   `json:"details" yaml:"details"`
	Message     string                                                   `json:"message" yaml:"message"`
	Path        string                                                   `json:"path" yaml:"path"`
	UserMessage string                                                   `json:"userMessage" yaml:"userMessage"`
	Metadata    PublishCommandDetailedReportActionResTasksErrorsMetadata `json:"metadata" yaml:"metadata"`
}

// The base class definition for metadata
type PublishCommandDetailedReportActionResTasksErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}
type PublishCommandDetailedReportActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// PublishCommandDetailedReportActionRaw registers a raw Gin route for the PublishCommandDetailedReportAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PublishCommandDetailedReportActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PublishCommandDetailedReportActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // PublishCommandDetailedReportActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PublishCommandDetailedReportAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PublishCommandDetailedReportActionHandler(
	handler func(c PublishCommandDetailedReportActionRequest, gin *gin.Context) (*PublishCommandDetailedReportActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PublishCommandDetailedReportActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PublishCommandDetailedReportActionRequest{
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

// PublishCommandDetailedReportAction is a high-level convenience wrapper around PublishCommandDetailedReportActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PublishCommandDetailedReportAction(r gin.IRoutes, handler func(c PublishCommandDetailedReportActionRequest, gin *gin.Context) (*PublishCommandDetailedReportActionResponse, error)) {
	method, url, h := PublishCommandDetailedReportActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Publish command detailed reportAction
 */
// Query wrapper with private fields
type PublishCommandDetailedReportActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func PublishCommandDetailedReportActionQueryFromString(rawQuery string) PublishCommandDetailedReportActionQuery {
	v := PublishCommandDetailedReportActionQuery{}
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
func PublishCommandDetailedReportActionQueryFromGin(c *gin.Context) PublishCommandDetailedReportActionQuery {
	return PublishCommandDetailedReportActionQueryFromString(c.Request.URL.RawQuery)
}
func PublishCommandDetailedReportActionQueryFromHttp(r *http.Request) PublishCommandDetailedReportActionQuery {
	return PublishCommandDetailedReportActionQueryFromString(r.URL.RawQuery)
}
func (q PublishCommandDetailedReportActionQuery) Values() url.Values {
	return q.values
}
func (q PublishCommandDetailedReportActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PublishCommandDetailedReportActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PublishCommandDetailedReportActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PublishCommandDetailedReportActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type PublishCommandDetailedReportActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func PublishCommandDetailedReportActionCall(
	req PublishCommandDetailedReportActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PublishCommandDetailedReportActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := PublishCommandDetailedReportActionMeta()
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
	var result PublishCommandDetailedReportActionResult
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
