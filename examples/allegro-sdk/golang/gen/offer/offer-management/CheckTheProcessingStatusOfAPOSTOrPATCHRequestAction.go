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
* Action to communicate with the action CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction
*/
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction",
        URL:    "https://api.{environment}/sale/product-offers/{offerId}/operations/{operationId}",
        Method: "GET",
    }
}
  // The base class definition for checkTheProcessingStatusOfAPOSTOrPATCHRequestActionRes
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes struct {
		Offer  CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer `json:"offer" yaml:"offer"`
		Operation  CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation `json:"operation" yaml:"operation"`
}
  // The base class definition for offer
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for operation
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation struct {
		Id string `json:"id" yaml:"id"`
		Status string `json:"status" yaml:"status"`
		StartedAt string `json:"startedAt" yaml:"startedAt"`
}
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRaw registers a raw Gin route for the CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHandler(
	handler func(c CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest, gin *gin.Context) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest{
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
// CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction is a high-level convenience wrapper around CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction(r gin.IRoutes, handler func(c CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest, gin *gin.Context) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse, error),) {
	method, url, h := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for Check the processing status of a POST or PATCH requestAction
 */
// Query wrapper with private fields
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromString(rawQuery string) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery {
	v := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery{}
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
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromGin(c *gin.Context) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery {
	return CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromString(c.Request.URL.RawQuery)
}
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromHttp(r *http.Request) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery {
	return CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromString(r.URL.RawQuery)
}
func (q CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) Values() url.Values {
	return q.values
}
func (q CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest struct {
	QueryParams url.Values
	Headers http.Header
}
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionCall(
	req CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
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
	var result CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResult
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