package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
)

/**
* Action to communicate with the action UpdateOfferTranslationAction
 */
func UpdateOfferTranslationActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "UpdateOfferTranslationAction",
		URL:    "https://api.{environment}/sale/offers/{offerId}/translations/{language}",
		Method: "PATCH",
	}
}

// The base class definition for updateOfferTranslationActionReq
type UpdateOfferTranslationActionReq struct {
	Description       UpdateOfferTranslationActionReqDescription       `json:"description" yaml:"description"`
	Title             UpdateOfferTranslationActionReqTitle             `json:"title" yaml:"title"`
	SafetyInformation UpdateOfferTranslationActionReqSafetyInformation `json:"safetyInformation" yaml:"safetyInformation"`
}

// The base class definition for description
type UpdateOfferTranslationActionReqDescription struct {
	Translation UpdateOfferTranslationActionReqDescriptionTranslation `json:"translation" yaml:"translation"`
}

// The base class definition for translation
type UpdateOfferTranslationActionReqDescriptionTranslation struct {
	Sections []UpdateOfferTranslationActionReqDescriptionTranslationSections `json:"sections" yaml:"sections"`
}

// The base class definition for sections
type UpdateOfferTranslationActionReqDescriptionTranslationSections struct {
	Items []UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems `json:"items" yaml:"items"`
}

// The base class definition for items
type UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

// The base class definition for title
type UpdateOfferTranslationActionReqTitle struct {
	Translation string `json:"translation" yaml:"translation"`
}

// The base class definition for safetyInformation
type UpdateOfferTranslationActionReqSafetyInformation struct {
	Products []UpdateOfferTranslationActionReqSafetyInformationProducts `json:"products" yaml:"products"`
}

// The base class definition for products
type UpdateOfferTranslationActionReqSafetyInformationProducts struct {
	Id          string `json:"id" yaml:"id"`
	Translation string `json:"translation" yaml:"translation"`
}
type UpdateOfferTranslationActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// UpdateOfferTranslationActionRaw registers a raw Gin route for the UpdateOfferTranslationAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func UpdateOfferTranslationActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := UpdateOfferTranslationActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // UpdateOfferTranslationActionHandler returns the HTTP method, route URL, and a typed Gin handler for the UpdateOfferTranslationAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func UpdateOfferTranslationActionHandler(
	handler func(c UpdateOfferTranslationActionRequest, gin *gin.Context) (*UpdateOfferTranslationActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := UpdateOfferTranslationActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body UpdateOfferTranslationActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := UpdateOfferTranslationActionRequest{
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

// UpdateOfferTranslationAction is a high-level convenience wrapper around UpdateOfferTranslationActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func UpdateOfferTranslationAction(r gin.IRoutes, handler func(c UpdateOfferTranslationActionRequest, gin *gin.Context) (*UpdateOfferTranslationActionResponse, error)) {
	method, url, h := UpdateOfferTranslationActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Update offer translationAction
 */
// Query wrapper with private fields
type UpdateOfferTranslationActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func UpdateOfferTranslationActionQueryFromString(rawQuery string) UpdateOfferTranslationActionQuery {
	v := UpdateOfferTranslationActionQuery{}
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
func UpdateOfferTranslationActionQueryFromGin(c *gin.Context) UpdateOfferTranslationActionQuery {
	return UpdateOfferTranslationActionQueryFromString(c.Request.URL.RawQuery)
}
func UpdateOfferTranslationActionQueryFromHttp(r *http.Request) UpdateOfferTranslationActionQuery {
	return UpdateOfferTranslationActionQueryFromString(r.URL.RawQuery)
}
func (q UpdateOfferTranslationActionQuery) Values() url.Values {
	return q.values
}
func (q UpdateOfferTranslationActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *UpdateOfferTranslationActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *UpdateOfferTranslationActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type UpdateOfferTranslationActionRequest struct {
	Body        UpdateOfferTranslationActionReq
	QueryParams url.Values
	Headers     http.Header
}
type UpdateOfferTranslationActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func UpdateOfferTranslationActionCall(
	req UpdateOfferTranslationActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*UpdateOfferTranslationActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := UpdateOfferTranslationActionMeta()
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
	var result UpdateOfferTranslationActionResult
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
