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
* Action to communicate with the action GetOfferPromotionPackagesAction
*/
func GetOfferPromotionPackagesActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "GetOfferPromotionPackagesAction",
        URL:    "https://api.{environment}/sale/offers/{offerId}/promo-options",
        Method: "GET",
    }
}
  // The base class definition for getOfferPromotionPackagesActionRes
type GetOfferPromotionPackagesActionRes struct {
		OfferId string `json:"offerId" yaml:"offerId"`
		MarketplaceId string `json:"marketplaceId" yaml:"marketplaceId"`
		BasePackage  GetOfferPromotionPackagesActionResBasePackage `json:"basePackage" yaml:"basePackage"`
		ExtraPackages []GetOfferPromotionPackagesActionResExtraPackages `json:"extraPackages" yaml:"extraPackages"`
		PendingChanges  GetOfferPromotionPackagesActionResPendingChanges `json:"pendingChanges" yaml:"pendingChanges"`
		AdditionalMarketplaces []GetOfferPromotionPackagesActionResAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}
  // The base class definition for basePackage
type GetOfferPromotionPackagesActionResBasePackage struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for extraPackages
type GetOfferPromotionPackagesActionResExtraPackages struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for pendingChanges
type GetOfferPromotionPackagesActionResPendingChanges struct {
		BasePackage  GetOfferPromotionPackagesActionResPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}
  // The base class definition for basePackage
type GetOfferPromotionPackagesActionResPendingChangesBasePackage struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for additionalMarketplaces
type GetOfferPromotionPackagesActionResAdditionalMarketplaces struct {
		MarketplaceId string `json:"marketplaceId" yaml:"marketplaceId"`
		BasePackage  GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage `json:"basePackage" yaml:"basePackage"`
		ExtraPackages []GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages `json:"extraPackages" yaml:"extraPackages"`
		PendingChanges  GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges `json:"pendingChanges" yaml:"pendingChanges"`
}
  // The base class definition for basePackage
type GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for extraPackages
type GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
  // The base class definition for pendingChanges
type GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges struct {
		BasePackage  GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}
  // The base class definition for basePackage
type GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage struct {
		Id string `json:"id" yaml:"id"`
		ValidFrom string `json:"validFrom" yaml:"validFrom"`
		ValidTo string `json:"validTo" yaml:"validTo"`
		NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}
type GetOfferPromotionPackagesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// GetOfferPromotionPackagesActionRaw registers a raw Gin route for the GetOfferPromotionPackagesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetOfferPromotionPackagesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetOfferPromotionPackagesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// GetOfferPromotionPackagesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetOfferPromotionPackagesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetOfferPromotionPackagesActionHandler(
	handler func(c GetOfferPromotionPackagesActionRequest, gin *gin.Context) (*GetOfferPromotionPackagesActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetOfferPromotionPackagesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetOfferPromotionPackagesActionRequest{
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
// GetOfferPromotionPackagesAction is a high-level convenience wrapper around GetOfferPromotionPackagesActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetOfferPromotionPackagesAction(r gin.IRoutes, handler func(c GetOfferPromotionPackagesActionRequest, gin *gin.Context) (*GetOfferPromotionPackagesActionResponse, error),) {
	method, url, h := GetOfferPromotionPackagesActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for Get offer promotion packagesAction
 */
// Query wrapper with private fields
type GetOfferPromotionPackagesActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func GetOfferPromotionPackagesActionQueryFromString(rawQuery string) GetOfferPromotionPackagesActionQuery {
	v := GetOfferPromotionPackagesActionQuery{}
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
func GetOfferPromotionPackagesActionQueryFromGin(c *gin.Context) GetOfferPromotionPackagesActionQuery {
	return GetOfferPromotionPackagesActionQueryFromString(c.Request.URL.RawQuery)
}
func GetOfferPromotionPackagesActionQueryFromHttp(r *http.Request) GetOfferPromotionPackagesActionQuery {
	return GetOfferPromotionPackagesActionQueryFromString(r.URL.RawQuery)
}
func (q GetOfferPromotionPackagesActionQuery) Values() url.Values {
	return q.values
}
func (q GetOfferPromotionPackagesActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetOfferPromotionPackagesActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetOfferPromotionPackagesActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type GetOfferPromotionPackagesActionRequest struct {
	QueryParams url.Values
	Headers http.Header
}
type GetOfferPromotionPackagesActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func GetOfferPromotionPackagesActionCall(
	req GetOfferPromotionPackagesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetOfferPromotionPackagesActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetOfferPromotionPackagesActionMeta()
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
	var result GetOfferPromotionPackagesActionResult
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