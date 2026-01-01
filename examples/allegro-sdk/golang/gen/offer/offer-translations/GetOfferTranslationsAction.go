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
* Action to communicate with the action GetOfferTranslationsAction
*/
func GetOfferTranslationsActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "GetOfferTranslationsAction",
        URL:    "https://api.{environment}/sale/offers/{offerId}/translations",
        Method: "GET",
    }
}
  // The base class definition for getOfferTranslationsActionRes
type GetOfferTranslationsActionRes struct {
		Translations []GetOfferTranslationsActionResTranslations `json:"translations" yaml:"translations"`
}
  // The base class definition for translations
type GetOfferTranslationsActionResTranslations struct {
		Language string `json:"language" yaml:"language"`
		Title  GetOfferTranslationsActionResTranslationsTitle `json:"title" yaml:"title"`
		Description  GetOfferTranslationsActionResTranslationsDescription `json:"description" yaml:"description"`
		SafetyInformation  GetOfferTranslationsActionResTranslationsSafetyInformation `json:"safetyInformation" yaml:"safetyInformation"`
}
  // The base class definition for title
type GetOfferTranslationsActionResTranslationsTitle struct {
		Translation string `json:"translation" yaml:"translation"`
		Type string `json:"type" yaml:"type"`
}
  // The base class definition for description
type GetOfferTranslationsActionResTranslationsDescription struct {
		Translation  GetOfferTranslationsActionResTranslationsDescriptionTranslation `json:"translation" yaml:"translation"`
		Type string `json:"type" yaml:"type"`
}
  // The base class definition for translation
type GetOfferTranslationsActionResTranslationsDescriptionTranslation struct {
		Sections []GetOfferTranslationsActionResTranslationsDescriptionTranslationSections `json:"sections" yaml:"sections"`
}
  // The base class definition for sections
type GetOfferTranslationsActionResTranslationsDescriptionTranslationSections struct {
		Items []GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems `json:"items" yaml:"items"`
}
  // The base class definition for items
type GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems struct {
		Type string `json:"type" yaml:"type"`
}
  // The base class definition for safetyInformation
type GetOfferTranslationsActionResTranslationsSafetyInformation struct {
		Products []GetOfferTranslationsActionResTranslationsSafetyInformationProducts `json:"products" yaml:"products"`
}
  // The base class definition for products
type GetOfferTranslationsActionResTranslationsSafetyInformationProducts struct {
		Id string `json:"id" yaml:"id"`
		Translation string `json:"translation" yaml:"translation"`
		Type string `json:"type" yaml:"type"`
}
type GetOfferTranslationsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// GetOfferTranslationsActionRaw registers a raw Gin route for the GetOfferTranslationsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetOfferTranslationsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetOfferTranslationsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// GetOfferTranslationsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetOfferTranslationsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetOfferTranslationsActionHandler(
	handler func(c GetOfferTranslationsActionRequest, gin *gin.Context) (*GetOfferTranslationsActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetOfferTranslationsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetOfferTranslationsActionRequest{
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
// GetOfferTranslationsAction is a high-level convenience wrapper around GetOfferTranslationsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetOfferTranslationsAction(r gin.IRoutes, handler func(c GetOfferTranslationsActionRequest, gin *gin.Context) (*GetOfferTranslationsActionResponse, error),) {
	method, url, h := GetOfferTranslationsActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for Get offer translationsAction
 */
// Query wrapper with private fields
type GetOfferTranslationsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func GetOfferTranslationsActionQueryFromString(rawQuery string) GetOfferTranslationsActionQuery {
	v := GetOfferTranslationsActionQuery{}
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
func GetOfferTranslationsActionQueryFromGin(c *gin.Context) GetOfferTranslationsActionQuery {
	return GetOfferTranslationsActionQueryFromString(c.Request.URL.RawQuery)
}
func GetOfferTranslationsActionQueryFromHttp(r *http.Request) GetOfferTranslationsActionQuery {
	return GetOfferTranslationsActionQueryFromString(r.URL.RawQuery)
}
func (q GetOfferTranslationsActionQuery) Values() url.Values {
	return q.values
}
func (q GetOfferTranslationsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetOfferTranslationsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetOfferTranslationsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type GetOfferTranslationsActionRequest struct {
	QueryParams url.Values
	Headers http.Header
}
type GetOfferTranslationsActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func GetOfferTranslationsActionCall(
	req GetOfferTranslationsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetOfferTranslationsActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetOfferTranslationsActionMeta()
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
	var result GetOfferTranslationsActionResult
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