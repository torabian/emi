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
* Action to communicate with the action ModifyOfferPromotionPackagesAction
 */
func ModifyOfferPromotionPackagesActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "ModifyOfferPromotionPackagesAction",
		URL:    "https://api.{environment}/sale/offers/{offerId}/promo-options-modification",
		Method: "POST",
	}
}

// The base class definition for modifyOfferPromotionPackagesActionReq
type ModifyOfferPromotionPackagesActionReq struct {
	Modifications          []ModifyOfferPromotionPackagesActionReqModifications          `json:"modifications" yaml:"modifications"`
	AdditionalMarketplaces []ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for modifications
type ModifyOfferPromotionPackagesActionReqModifications struct {
	ModificationType string `json:"modificationType" yaml:"modificationType"`
	PackageType      string `json:"packageType" yaml:"packageType"`
	PackageId        string `json:"packageId" yaml:"packageId"`
}

// The base class definition for additionalMarketplaces
type ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces struct {
	MarketplaceId string                                                                     `json:"marketplaceId" yaml:"marketplaceId"`
	Modifications []ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications `json:"modifications" yaml:"modifications"`
}

// The base class definition for modifications
type ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications struct {
	ModificationType string `json:"modificationType" yaml:"modificationType"`
	PackageType      string `json:"packageType" yaml:"packageType"`
	PackageId        string `json:"packageId" yaml:"packageId"`
}

// The base class definition for modifyOfferPromotionPackagesActionRes
type ModifyOfferPromotionPackagesActionRes struct {
	OfferId                string                                                        `json:"offerId" yaml:"offerId"`
	MarketplaceId          string                                                        `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage            ModifyOfferPromotionPackagesActionResBasePackage              `json:"basePackage" yaml:"basePackage"`
	ExtraPackages          []ModifyOfferPromotionPackagesActionResExtraPackages          `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges         ModifyOfferPromotionPackagesActionResPendingChanges           `json:"pendingChanges" yaml:"pendingChanges"`
	AdditionalMarketplaces []ModifyOfferPromotionPackagesActionResAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for basePackage
type ModifyOfferPromotionPackagesActionResBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for extraPackages
type ModifyOfferPromotionPackagesActionResExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for pendingChanges
type ModifyOfferPromotionPackagesActionResPendingChanges struct {
	BasePackage ModifyOfferPromotionPackagesActionResPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

// The base class definition for basePackage
type ModifyOfferPromotionPackagesActionResPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for additionalMarketplaces
type ModifyOfferPromotionPackagesActionResAdditionalMarketplaces struct {
	MarketplaceId  string                                                                     `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage    ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages  []ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges  `json:"pendingChanges" yaml:"pendingChanges"`
}

// The base class definition for basePackage
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for extraPackages
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for pendingChanges
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges struct {
	BasePackage ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

// The base class definition for basePackage
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
type ModifyOfferPromotionPackagesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// ModifyOfferPromotionPackagesActionRaw registers a raw Gin route for the ModifyOfferPromotionPackagesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModifyOfferPromotionPackagesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModifyOfferPromotionPackagesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // ModifyOfferPromotionPackagesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModifyOfferPromotionPackagesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModifyOfferPromotionPackagesActionHandler(
	handler func(c ModifyOfferPromotionPackagesActionRequest, gin *gin.Context) (*ModifyOfferPromotionPackagesActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ModifyOfferPromotionPackagesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ModifyOfferPromotionPackagesActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ModifyOfferPromotionPackagesActionRequest{
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

// ModifyOfferPromotionPackagesAction is a high-level convenience wrapper around ModifyOfferPromotionPackagesActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModifyOfferPromotionPackagesAction(r gin.IRoutes, handler func(c ModifyOfferPromotionPackagesActionRequest, gin *gin.Context) (*ModifyOfferPromotionPackagesActionResponse, error)) {
	method, url, h := ModifyOfferPromotionPackagesActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Modify offer promotion packagesAction
 */
// Query wrapper with private fields
type ModifyOfferPromotionPackagesActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModifyOfferPromotionPackagesActionQueryFromString(rawQuery string) ModifyOfferPromotionPackagesActionQuery {
	v := ModifyOfferPromotionPackagesActionQuery{}
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
func ModifyOfferPromotionPackagesActionQueryFromGin(c *gin.Context) ModifyOfferPromotionPackagesActionQuery {
	return ModifyOfferPromotionPackagesActionQueryFromString(c.Request.URL.RawQuery)
}
func ModifyOfferPromotionPackagesActionQueryFromHttp(r *http.Request) ModifyOfferPromotionPackagesActionQuery {
	return ModifyOfferPromotionPackagesActionQueryFromString(r.URL.RawQuery)
}
func (q ModifyOfferPromotionPackagesActionQuery) Values() url.Values {
	return q.values
}
func (q ModifyOfferPromotionPackagesActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModifyOfferPromotionPackagesActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModifyOfferPromotionPackagesActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModifyOfferPromotionPackagesActionRequest struct {
	Body        ModifyOfferPromotionPackagesActionReq
	QueryParams url.Values
	Headers     http.Header
}
type ModifyOfferPromotionPackagesActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ModifyOfferPromotionPackagesActionCall(
	req ModifyOfferPromotionPackagesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModifyOfferPromotionPackagesActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ModifyOfferPromotionPackagesActionMeta()
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
	var result ModifyOfferPromotionPackagesActionResult
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
