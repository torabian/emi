package external
import (
"bytes"
"encoding/json"
"fmt"
"github.com/gin-gonic/gin"
"io"
"net/http"
"net/url"
"test.com/emi/golang/emigo"
)
/**
* Action to communicate with the action BatchOfferPublishUnpublishAction
*/
func BatchOfferPublishUnpublishActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "BatchOfferPublishUnpublishAction",
        URL:    "https://api.{environment}/sale/offer-publication-commands/{commandId}",
        Method: "PUT",
    }
}
  // The base class definition for batchOfferPublishUnpublishActionReq
type BatchOfferPublishUnpublishActionReq struct {
		OfferCriteria []BatchOfferPublishUnpublishActionReqOfferCriteria `json:"offerCriteria" yaml:"offerCriteria"`
		Publication  BatchOfferPublishUnpublishActionReqPublication `json:"publication" yaml:"publication"`
}
  // The base class definition for offerCriteria
type BatchOfferPublishUnpublishActionReqOfferCriteria struct {
		Offers []BatchOfferPublishUnpublishActionReqOfferCriteriaOffers `json:"offers" yaml:"offers"`
		Type string `json:"type" yaml:"type"`
}
  // The base class definition for offers
type BatchOfferPublishUnpublishActionReqOfferCriteriaOffers struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for publication
type BatchOfferPublishUnpublishActionReqPublication struct {
		Action string `json:"action" yaml:"action"`
		ScheduledFor string `json:"scheduledFor" yaml:"scheduledFor"`
}
  // The base class definition for batchOfferPublishUnpublishActionRes
type BatchOfferPublishUnpublishActionRes struct {
		Id string `json:"id" yaml:"id"`
		CreatedAt string `json:"createdAt" yaml:"createdAt"`
		CompletedAt string `json:"completedAt" yaml:"completedAt"`
		TaskCount  BatchOfferPublishUnpublishActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}
  // The base class definition for taskCount
type BatchOfferPublishUnpublishActionResTaskCount struct {
		Failed int `json:"failed" yaml:"failed"`
		Success int `json:"success" yaml:"success"`
		Total int `json:"total" yaml:"total"`
}
type BatchOfferPublishUnpublishActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// BatchOfferPublishUnpublishActionRaw registers a raw Gin route for the BatchOfferPublishUnpublishAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func BatchOfferPublishUnpublishActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := BatchOfferPublishUnpublishActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// BatchOfferPublishUnpublishActionHandler returns the HTTP method, route URL, and a typed Gin handler for the BatchOfferPublishUnpublishAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func BatchOfferPublishUnpublishActionHandler(
	handler func(c BatchOfferPublishUnpublishActionRequest, gin *gin.Context) (*BatchOfferPublishUnpublishActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := BatchOfferPublishUnpublishActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body BatchOfferPublishUnpublishActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := BatchOfferPublishUnpublishActionRequest{
			Body:        body,
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
// BatchOfferPublishUnpublishAction is a high-level convenience wrapper around BatchOfferPublishUnpublishActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func BatchOfferPublishUnpublishAction(r gin.IRoutes, handler func(c BatchOfferPublishUnpublishActionRequest, gin *gin.Context) (*BatchOfferPublishUnpublishActionResponse, error),) {
	method, url, h := BatchOfferPublishUnpublishActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for Batch offer publish / unpublishAction
 */
// Query wrapper with private fields
type BatchOfferPublishUnpublishActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func BatchOfferPublishUnpublishActionQueryFromString(rawQuery string) BatchOfferPublishUnpublishActionQuery {
	v := BatchOfferPublishUnpublishActionQuery{}
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
func BatchOfferPublishUnpublishActionQueryFromGin(c *gin.Context) BatchOfferPublishUnpublishActionQuery {
	return BatchOfferPublishUnpublishActionQueryFromString(c.Request.URL.RawQuery)
}
func BatchOfferPublishUnpublishActionQueryFromHttp(r *http.Request) BatchOfferPublishUnpublishActionQuery {
	return BatchOfferPublishUnpublishActionQueryFromString(r.URL.RawQuery)
}
func (q BatchOfferPublishUnpublishActionQuery) Values() url.Values {
	return q.values
}
func (q BatchOfferPublishUnpublishActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *BatchOfferPublishUnpublishActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *BatchOfferPublishUnpublishActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type BatchOfferPublishUnpublishActionRequest struct {
	Body BatchOfferPublishUnpublishActionReq
	QueryParams url.Values
	Headers http.Header
}
type BatchOfferPublishUnpublishActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func BatchOfferPublishUnpublishActionCall(
	req BatchOfferPublishUnpublishActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*BatchOfferPublishUnpublishActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := BatchOfferPublishUnpublishActionMeta()
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
			bodyBytes, err := json.Marshal(req.Body)
			if err != nil {
				return nil, err
			}
			req0, err := http.NewRequest(meta.Method, u.String(), bytes.NewReader(bodyBytes))
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
	var result BatchOfferPublishUnpublishActionResult
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