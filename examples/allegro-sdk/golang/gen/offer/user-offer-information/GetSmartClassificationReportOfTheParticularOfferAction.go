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
* Action to communicate with the action GetSmartClassificationReportOfTheParticularOfferAction
*/
func GetSmartClassificationReportOfTheParticularOfferActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "GetSmartClassificationReportOfTheParticularOfferAction",
        URL:    "https://api.{environment}/sale/offers/{offerId}/smart",
        Method: "GET",
    }
}
  // The base class definition for getSmartClassificationReportOfTheParticularOfferActionRes
type GetSmartClassificationReportOfTheParticularOfferActionRes struct {
		  // Indicates if offer is queued for reclassification
 ScheduledForReclassification bool `json:"scheduledForReclassification" yaml:"scheduledForReclassification"`
		  // Offer classification status and last change date
 Classification  GetSmartClassificationReportOfTheParticularOfferActionResClassification `json:"classification" yaml:"classification"`
		  // List of smart delivery method identifiers
 SmartDeliveryMethods []GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods `json:"smartDeliveryMethods" yaml:"smartDeliveryMethods"`
		  // List of classification conditions with delivery method checks
 Conditions []GetSmartClassificationReportOfTheParticularOfferActionResConditions `json:"conditions" yaml:"conditions"`
}
  // The base class definition for classification
type GetSmartClassificationReportOfTheParticularOfferActionResClassification struct {
		  // Whether the classification conditions are fulfilled
 Fulfilled bool `json:"fulfilled" yaml:"fulfilled"`
		  // ISO8601 timestamp of last classification change
 LastChanged string `json:"lastChanged" yaml:"lastChanged"`
}
  // The base class definition for smartDeliveryMethods
type GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for conditions
type GetSmartClassificationReportOfTheParticularOfferActionResConditions struct {
		  // Condition code identifier
 Code string `json:"code" yaml:"code"`
		  // Human-readable condition name
 Name string `json:"name" yaml:"name"`
		  // Detailed condition description
 Description string `json:"description" yaml:"description"`
		  // Indicates if this condition is fulfilled
 Fulfilled bool `json:"fulfilled" yaml:"fulfilled"`
		  // Delivery methods that passed validation for this condition
 PassedDeliveryMethods []GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods `json:"passedDeliveryMethods" yaml:"passedDeliveryMethods"`
		  // Delivery methods that failed validation for this condition
 FailedDeliveryMethods []GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods `json:"failedDeliveryMethods" yaml:"failedDeliveryMethods"`
}
  // The base class definition for passedDeliveryMethods
type GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for failedDeliveryMethods
type GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods struct {
		Id string `json:"id" yaml:"id"`
}
type GetSmartClassificationReportOfTheParticularOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// GetSmartClassificationReportOfTheParticularOfferActionRaw registers a raw Gin route for the GetSmartClassificationReportOfTheParticularOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetSmartClassificationReportOfTheParticularOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetSmartClassificationReportOfTheParticularOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// GetSmartClassificationReportOfTheParticularOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetSmartClassificationReportOfTheParticularOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetSmartClassificationReportOfTheParticularOfferActionHandler(
	handler func(c GetSmartClassificationReportOfTheParticularOfferActionRequest, gin *gin.Context) (*GetSmartClassificationReportOfTheParticularOfferActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetSmartClassificationReportOfTheParticularOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetSmartClassificationReportOfTheParticularOfferActionRequest{
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
// GetSmartClassificationReportOfTheParticularOfferAction is a high-level convenience wrapper around GetSmartClassificationReportOfTheParticularOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetSmartClassificationReportOfTheParticularOfferAction(r gin.IRoutes, handler func(c GetSmartClassificationReportOfTheParticularOfferActionRequest, gin *gin.Context) (*GetSmartClassificationReportOfTheParticularOfferActionResponse, error),) {
	method, url, h := GetSmartClassificationReportOfTheParticularOfferActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for Get Smart! classification report of the particular offerAction
 */
// Query wrapper with private fields
type GetSmartClassificationReportOfTheParticularOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func GetSmartClassificationReportOfTheParticularOfferActionQueryFromString(rawQuery string) GetSmartClassificationReportOfTheParticularOfferActionQuery {
	v := GetSmartClassificationReportOfTheParticularOfferActionQuery{}
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
func GetSmartClassificationReportOfTheParticularOfferActionQueryFromGin(c *gin.Context) GetSmartClassificationReportOfTheParticularOfferActionQuery {
	return GetSmartClassificationReportOfTheParticularOfferActionQueryFromString(c.Request.URL.RawQuery)
}
func GetSmartClassificationReportOfTheParticularOfferActionQueryFromHttp(r *http.Request) GetSmartClassificationReportOfTheParticularOfferActionQuery {
	return GetSmartClassificationReportOfTheParticularOfferActionQueryFromString(r.URL.RawQuery)
}
func (q GetSmartClassificationReportOfTheParticularOfferActionQuery) Values() url.Values {
	return q.values
}
func (q GetSmartClassificationReportOfTheParticularOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetSmartClassificationReportOfTheParticularOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetSmartClassificationReportOfTheParticularOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type GetSmartClassificationReportOfTheParticularOfferActionRequest struct {
	QueryParams url.Values
	Headers http.Header
}
type GetSmartClassificationReportOfTheParticularOfferActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func GetSmartClassificationReportOfTheParticularOfferActionCall(
	req GetSmartClassificationReportOfTheParticularOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetSmartClassificationReportOfTheParticularOfferActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetSmartClassificationReportOfTheParticularOfferActionMeta()
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
	var result GetSmartClassificationReportOfTheParticularOfferActionResult
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