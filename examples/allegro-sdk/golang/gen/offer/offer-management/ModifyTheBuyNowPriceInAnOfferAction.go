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
* Action to communicate with the action ModifyTheBuyNowPriceInAnOfferAction
*/
func ModifyTheBuyNowPriceInAnOfferActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "ModifyTheBuyNowPriceInAnOfferAction",
        URL:    "https://api.{environment}/offers/{offerId}/change-price-commands/{commandId}",
        Method: "PUT",
    }
}
  // The base class definition for modifyTheBuyNowPriceInAnOfferActionReq
type ModifyTheBuyNowPriceInAnOfferActionReq struct {
		Id string `json:"id" yaml:"id"`
		Input  ModifyTheBuyNowPriceInAnOfferActionReqInput `json:"input" yaml:"input"`
}
  // The base class definition for input
type ModifyTheBuyNowPriceInAnOfferActionReqInput struct {
		BuyNowPrice  ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice `json:"buyNowPrice" yaml:"buyNowPrice"`
}
  // The base class definition for buyNowPrice
type ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for modifyTheBuyNowPriceInAnOfferActionRes
type ModifyTheBuyNowPriceInAnOfferActionRes struct {
		Id string `json:"id" yaml:"id"`
		Input  ModifyTheBuyNowPriceInAnOfferActionResInput `json:"input" yaml:"input"`
		Output  ModifyTheBuyNowPriceInAnOfferActionResOutput `json:"output" yaml:"output"`
}
  // The base class definition for input
type ModifyTheBuyNowPriceInAnOfferActionResInput struct {
		BuyNowPrice  ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice `json:"buyNowPrice" yaml:"buyNowPrice"`
}
  // The base class definition for buyNowPrice
type ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for output
type ModifyTheBuyNowPriceInAnOfferActionResOutput struct {
		Status string `json:"status" yaml:"status"`
		Errors []ModifyTheBuyNowPriceInAnOfferActionResOutputErrors `json:"errors" yaml:"errors"`
}
  // The base class definition for errors
type ModifyTheBuyNowPriceInAnOfferActionResOutputErrors struct {
		Code string `json:"code" yaml:"code"`
		Details string `json:"details" yaml:"details"`
		Message string `json:"message" yaml:"message"`
		Path string `json:"path" yaml:"path"`
		UserMessage string `json:"userMessage" yaml:"userMessage"`
		Metadata  ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata `json:"metadata" yaml:"metadata"`
}
  // The base class definition for metadata
type ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata struct {
		ProductId string `json:"productId" yaml:"productId"`
}
type ModifyTheBuyNowPriceInAnOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// ModifyTheBuyNowPriceInAnOfferActionRaw registers a raw Gin route for the ModifyTheBuyNowPriceInAnOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModifyTheBuyNowPriceInAnOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// ModifyTheBuyNowPriceInAnOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModifyTheBuyNowPriceInAnOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModifyTheBuyNowPriceInAnOfferActionHandler(
	handler func(c ModifyTheBuyNowPriceInAnOfferActionRequest, gin *gin.Context) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ModifyTheBuyNowPriceInAnOfferActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ModifyTheBuyNowPriceInAnOfferActionRequest{
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
// ModifyTheBuyNowPriceInAnOfferAction is a high-level convenience wrapper around ModifyTheBuyNowPriceInAnOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModifyTheBuyNowPriceInAnOfferAction(r gin.IRoutes, handler func(c ModifyTheBuyNowPriceInAnOfferActionRequest, gin *gin.Context) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error),) {
	method, url, h := ModifyTheBuyNowPriceInAnOfferActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for Modify the Buy Now price in an offerAction
 */
// Query wrapper with private fields
type ModifyTheBuyNowPriceInAnOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func ModifyTheBuyNowPriceInAnOfferActionQueryFromString(rawQuery string) ModifyTheBuyNowPriceInAnOfferActionQuery {
	v := ModifyTheBuyNowPriceInAnOfferActionQuery{}
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
func ModifyTheBuyNowPriceInAnOfferActionQueryFromGin(c *gin.Context) ModifyTheBuyNowPriceInAnOfferActionQuery {
	return ModifyTheBuyNowPriceInAnOfferActionQueryFromString(c.Request.URL.RawQuery)
}
func ModifyTheBuyNowPriceInAnOfferActionQueryFromHttp(r *http.Request) ModifyTheBuyNowPriceInAnOfferActionQuery {
	return ModifyTheBuyNowPriceInAnOfferActionQueryFromString(r.URL.RawQuery)
}
func (q ModifyTheBuyNowPriceInAnOfferActionQuery) Values() url.Values {
	return q.values
}
func (q ModifyTheBuyNowPriceInAnOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModifyTheBuyNowPriceInAnOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModifyTheBuyNowPriceInAnOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type ModifyTheBuyNowPriceInAnOfferActionRequest struct {
	Body ModifyTheBuyNowPriceInAnOfferActionReq
	QueryParams url.Values
	Headers http.Header
}
type ModifyTheBuyNowPriceInAnOfferActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func ModifyTheBuyNowPriceInAnOfferActionCall(
	req ModifyTheBuyNowPriceInAnOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModifyTheBuyNowPriceInAnOfferActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
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
	var result ModifyTheBuyNowPriceInAnOfferActionResult
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