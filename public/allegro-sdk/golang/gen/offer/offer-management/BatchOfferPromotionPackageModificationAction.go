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
* Action to communicate with the action BatchOfferPromotionPackageModificationAction
 */
func BatchOfferPromotionPackageModificationActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "BatchOfferPromotionPackageModificationAction",
		URL:    "https://api.{environment}/sale/offers/promo-options-commands/{commandId}",
		Method: "PUT",
	}
}

// The base class definition for batchOfferPromotionPackageModificationActionReq
type BatchOfferPromotionPackageModificationActionReq struct {
	OfferCriteria          []BatchOfferPromotionPackageModificationActionReqOfferCriteria          `json:"offerCriteria" yaml:"offerCriteria"`
	Modification           BatchOfferPromotionPackageModificationActionReqModification             `json:"modification" yaml:"modification"`
	AdditionalMarketplaces []BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for offerCriteria
type BatchOfferPromotionPackageModificationActionReqOfferCriteria struct {
	Offers []BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers `json:"offers" yaml:"offers"`
	Type   string                                                               `json:"type" yaml:"type"`
}

// The base class definition for offers
type BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for modification
type BatchOfferPromotionPackageModificationActionReqModification struct {
	BasePackage      BatchOfferPromotionPackageModificationActionReqModificationBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages    []BatchOfferPromotionPackageModificationActionReqModificationExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	ModificationTime string                                                                     `json:"modificationTime" yaml:"modificationTime"`
}

// The base class definition for basePackage
type BatchOfferPromotionPackageModificationActionReqModificationBasePackage struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for extraPackages
type BatchOfferPromotionPackageModificationActionReqModificationExtraPackages struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for additionalMarketplaces
type BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces struct {
	MarketplaceId string                                                                            `json:"marketplaceId" yaml:"marketplaceId"`
	Modification  BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification `json:"modification" yaml:"modification"`
}

// The base class definition for modification
type BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification struct {
	BasePackage      BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages    []BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	ModificationTime string                                                                                           `json:"modificationTime" yaml:"modificationTime"`
}

// The base class definition for basePackage
type BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for extraPackages
type BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for batchOfferPromotionPackageModificationActionRes
type BatchOfferPromotionPackageModificationActionRes struct {
	Id        string                                                   `json:"id" yaml:"id"`
	TaskCount BatchOfferPromotionPackageModificationActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}

// The base class definition for taskCount
type BatchOfferPromotionPackageModificationActionResTaskCount struct {
	Failed  int `json:"failed" yaml:"failed"`
	Success int `json:"success" yaml:"success"`
	Total   int `json:"total" yaml:"total"`
}
type BatchOfferPromotionPackageModificationActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// BatchOfferPromotionPackageModificationActionRaw registers a raw Gin route for the BatchOfferPromotionPackageModificationAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func BatchOfferPromotionPackageModificationActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := BatchOfferPromotionPackageModificationActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // BatchOfferPromotionPackageModificationActionHandler returns the HTTP method, route URL, and a typed Gin handler for the BatchOfferPromotionPackageModificationAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func BatchOfferPromotionPackageModificationActionHandler(
	handler func(c BatchOfferPromotionPackageModificationActionRequest, gin *gin.Context) (*BatchOfferPromotionPackageModificationActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := BatchOfferPromotionPackageModificationActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body BatchOfferPromotionPackageModificationActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := BatchOfferPromotionPackageModificationActionRequest{
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

// BatchOfferPromotionPackageModificationAction is a high-level convenience wrapper around BatchOfferPromotionPackageModificationActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func BatchOfferPromotionPackageModificationAction(r gin.IRoutes, handler func(c BatchOfferPromotionPackageModificationActionRequest, gin *gin.Context) (*BatchOfferPromotionPackageModificationActionResponse, error)) {
	method, url, h := BatchOfferPromotionPackageModificationActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Batch offer promotion package modificationAction
 */
// Query wrapper with private fields
type BatchOfferPromotionPackageModificationActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func BatchOfferPromotionPackageModificationActionQueryFromString(rawQuery string) BatchOfferPromotionPackageModificationActionQuery {
	v := BatchOfferPromotionPackageModificationActionQuery{}
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
func BatchOfferPromotionPackageModificationActionQueryFromGin(c *gin.Context) BatchOfferPromotionPackageModificationActionQuery {
	return BatchOfferPromotionPackageModificationActionQueryFromString(c.Request.URL.RawQuery)
}
func BatchOfferPromotionPackageModificationActionQueryFromHttp(r *http.Request) BatchOfferPromotionPackageModificationActionQuery {
	return BatchOfferPromotionPackageModificationActionQueryFromString(r.URL.RawQuery)
}
func (q BatchOfferPromotionPackageModificationActionQuery) Values() url.Values {
	return q.values
}
func (q BatchOfferPromotionPackageModificationActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *BatchOfferPromotionPackageModificationActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *BatchOfferPromotionPackageModificationActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type BatchOfferPromotionPackageModificationActionRequest struct {
	Body        BatchOfferPromotionPackageModificationActionReq
	QueryParams url.Values
	Headers     http.Header
}
type BatchOfferPromotionPackageModificationActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func BatchOfferPromotionPackageModificationActionCall(
	req BatchOfferPromotionPackageModificationActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*BatchOfferPromotionPackageModificationActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := BatchOfferPromotionPackageModificationActionMeta()
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
	var result BatchOfferPromotionPackageModificationActionResult
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
