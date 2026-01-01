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
* Action to communicate with the action DeleteADraftOfferAction
*/
func DeleteADraftOfferActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "DeleteADraftOfferAction",
        URL:    "https://api.{environment}/sale/offers/{offerId}",
        Method: "DELETE",
    }
}
type DeleteADraftOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// DeleteADraftOfferActionRaw registers a raw Gin route for the DeleteADraftOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func DeleteADraftOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := DeleteADraftOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// DeleteADraftOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the DeleteADraftOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func DeleteADraftOfferActionHandler(
	handler func(c DeleteADraftOfferActionRequest, gin *gin.Context) (*DeleteADraftOfferActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := DeleteADraftOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := DeleteADraftOfferActionRequest{
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
// DeleteADraftOfferAction is a high-level convenience wrapper around DeleteADraftOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func DeleteADraftOfferAction(r gin.IRoutes, handler func(c DeleteADraftOfferActionRequest, gin *gin.Context) (*DeleteADraftOfferActionResponse, error),) {
	method, url, h := DeleteADraftOfferActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for Delete a draft offerAction
 */
// Query wrapper with private fields
type DeleteADraftOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func DeleteADraftOfferActionQueryFromString(rawQuery string) DeleteADraftOfferActionQuery {
	v := DeleteADraftOfferActionQuery{}
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
func DeleteADraftOfferActionQueryFromGin(c *gin.Context) DeleteADraftOfferActionQuery {
	return DeleteADraftOfferActionQueryFromString(c.Request.URL.RawQuery)
}
func DeleteADraftOfferActionQueryFromHttp(r *http.Request) DeleteADraftOfferActionQuery {
	return DeleteADraftOfferActionQueryFromString(r.URL.RawQuery)
}
func (q DeleteADraftOfferActionQuery) Values() url.Values {
	return q.values
}
func (q DeleteADraftOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *DeleteADraftOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *DeleteADraftOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type DeleteADraftOfferActionRequest struct {
	QueryParams url.Values
	Headers http.Header
}
type DeleteADraftOfferActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func DeleteADraftOfferActionCall(
	req DeleteADraftOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*DeleteADraftOfferActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := DeleteADraftOfferActionMeta()
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
	var result DeleteADraftOfferActionResult
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