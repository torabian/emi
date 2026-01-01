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
* Action to communicate with the action GetPromoOptionsForSellerSOffersAction
*/
func GetPromoOptionsForSellerSOffersActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "GetPromoOptionsForSellerSOffersAction",
        URL:    "https://api.{environment}/sale/offers/promo-options",
        Method: "GET",
    }
}
  // The base class definition for getPromoOptionsForSellerSOffersActionRes
type GetPromoOptionsForSellerSOffersActionRes struct {
		PromoOptions []GetPromoOptionsForSellerSOffersActionResPromoOptions `json:"promoOptions" yaml:"promoOptions"`
		Count int `json:"count" yaml:"count"`
		TotalCount int `json:"totalCount" yaml:"totalCount"`
}
  // The base class definition for promoOptions
type GetPromoOptionsForSellerSOffersActionResPromoOptions struct {
		OfferId string `json:"offerId" yaml:"offerId"`
		MarketplaceId string `json:"marketplaceId" yaml:"marketplaceId"`
		BasePackage  GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage `json:"basePackage" yaml:"basePackage"`
		ExtraPackages []GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages `json:"extraPackages" yaml:"extraPackages"`
		PendingChanges  GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges `json:"pendingChanges" yaml:"pendingChanges"`
		AdditionalMarketplaces []GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}
  // The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for extraPackages
type GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for pendingChanges
type GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges struct {
		BasePackage  GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}
  // The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for additionalMarketplaces
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces struct {
		MarketplaceId string `json:"marketplaceId" yaml:"marketplaceId"`
		BasePackage  GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage `json:"basePackage" yaml:"basePackage"`
		ExtraPackages []GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages `json:"extraPackages" yaml:"extraPackages"`
		PendingChanges  GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges `json:"pendingChanges" yaml:"pendingChanges"`
}
  // The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for extraPackages
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for pendingChanges
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges struct {
		BasePackage  GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}
  // The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
type GetPromoOptionsForSellerSOffersActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// GetPromoOptionsForSellerSOffersActionRaw registers a raw Gin route for the GetPromoOptionsForSellerSOffersAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetPromoOptionsForSellerSOffersActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetPromoOptionsForSellerSOffersActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// GetPromoOptionsForSellerSOffersActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetPromoOptionsForSellerSOffersAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetPromoOptionsForSellerSOffersActionHandler(
	handler func(c GetPromoOptionsForSellerSOffersActionRequest, gin *gin.Context) (*GetPromoOptionsForSellerSOffersActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetPromoOptionsForSellerSOffersActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetPromoOptionsForSellerSOffersActionRequest{
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
// GetPromoOptionsForSellerSOffersAction is a high-level convenience wrapper around GetPromoOptionsForSellerSOffersActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetPromoOptionsForSellerSOffersAction(r gin.IRoutes, handler func(c GetPromoOptionsForSellerSOffersActionRequest, gin *gin.Context) (*GetPromoOptionsForSellerSOffersActionResponse, error),) {
	method, url, h := GetPromoOptionsForSellerSOffersActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for Get promo options for seller's offersAction
 */
// Query wrapper with private fields
type GetPromoOptionsForSellerSOffersActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func GetPromoOptionsForSellerSOffersActionQueryFromString(rawQuery string) GetPromoOptionsForSellerSOffersActionQuery {
	v := GetPromoOptionsForSellerSOffersActionQuery{}
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
func GetPromoOptionsForSellerSOffersActionQueryFromGin(c *gin.Context) GetPromoOptionsForSellerSOffersActionQuery {
	return GetPromoOptionsForSellerSOffersActionQueryFromString(c.Request.URL.RawQuery)
}
func GetPromoOptionsForSellerSOffersActionQueryFromHttp(r *http.Request) GetPromoOptionsForSellerSOffersActionQuery {
	return GetPromoOptionsForSellerSOffersActionQueryFromString(r.URL.RawQuery)
}
func (q GetPromoOptionsForSellerSOffersActionQuery) Values() url.Values {
	return q.values
}
func (q GetPromoOptionsForSellerSOffersActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetPromoOptionsForSellerSOffersActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetPromoOptionsForSellerSOffersActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type GetPromoOptionsForSellerSOffersActionRequest struct {
	QueryParams url.Values
	Headers http.Header
}
type GetPromoOptionsForSellerSOffersActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func GetPromoOptionsForSellerSOffersActionCall(
	req GetPromoOptionsForSellerSOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetPromoOptionsForSellerSOffersActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetPromoOptionsForSellerSOffersActionMeta()
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
	var result GetPromoOptionsForSellerSOffersActionResult
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