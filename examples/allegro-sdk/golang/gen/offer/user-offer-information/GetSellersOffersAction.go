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
* Action to communicate with the action GetSellersOffersAction
*/
func GetSellersOffersActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "GetSellersOffersAction",
        URL:    "https://api.{environment}/sale/offers",
        Method: "GET",
    }
}
  // The base class definition for getSellersOffersActionRes
type GetSellersOffersActionRes struct {
		  // Number of offers in this page
 Count int `json:"count" yaml:"count"`
		  // Total number of offers available
 TotalCount int `json:"totalCount" yaml:"totalCount"`
		Offers []GetSellersOffersActionResOffers `json:"offers" yaml:"offers"`
}
  // The base class definition for offers
type GetSellersOffersActionResOffers struct {
		  // Offer identifier
 Id string `json:"id" yaml:"id"`
		  // Offer name or title
 Name string `json:"name" yaml:"name"`
		Category  GetSellersOffersActionResOffersCategory `json:"category" yaml:"category"`
		PrimaryImage  GetSellersOffersActionResOffersPrimaryImage `json:"primaryImage" yaml:"primaryImage"`
		SellingMode  GetSellersOffersActionResOffersSellingMode `json:"sellingMode" yaml:"sellingMode"`
		SaleInfo  GetSellersOffersActionResOffersSaleInfo `json:"saleInfo" yaml:"saleInfo"`
		Stock  GetSellersOffersActionResOffersStock `json:"stock" yaml:"stock"`
		Stats  GetSellersOffersActionResOffersStats `json:"stats" yaml:"stats"`
		Publication  GetSellersOffersActionResOffersPublication `json:"publication" yaml:"publication"`
		AfterSalesServices  GetSellersOffersActionResOffersAfterSalesServices `json:"afterSalesServices" yaml:"afterSalesServices"`
		AdditionalServices  GetSellersOffersActionResOffersAdditionalServices `json:"additionalServices" yaml:"additionalServices"`
		External  GetSellersOffersActionResOffersExternal `json:"external" yaml:"external"`
		Delivery  GetSellersOffersActionResOffersDelivery `json:"delivery" yaml:"delivery"`
		B2b  GetSellersOffersActionResOffersB2b `json:"b2b" yaml:"b2b"`
		FundraisingCampaign  GetSellersOffersActionResOffersFundraisingCampaign `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
		  // Marketplace-specific extensions for offer
 AdditionalMarketplaces emigo.Nullable[interface{}] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}
  // The base class definition for category
type GetSellersOffersActionResOffersCategory struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for primaryImage
type GetSellersOffersActionResOffersPrimaryImage struct {
		Url string `json:"url" yaml:"url"`
}
  // The base class definition for sellingMode
type GetSellersOffersActionResOffersSellingMode struct {
		Format string `json:"format" yaml:"format"`
		Price  GetSellersOffersActionResOffersSellingModePrice `json:"price" yaml:"price"`
		PriceAutomation  GetSellersOffersActionResOffersSellingModePriceAutomation `json:"priceAutomation" yaml:"priceAutomation"`
		MinimalPrice  GetSellersOffersActionResOffersSellingModeMinimalPrice `json:"minimalPrice" yaml:"minimalPrice"`
		StartingPrice  GetSellersOffersActionResOffersSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}
  // The base class definition for price
type GetSellersOffersActionResOffersSellingModePrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for priceAutomation
type GetSellersOffersActionResOffersSellingModePriceAutomation struct {
		Rule  GetSellersOffersActionResOffersSellingModePriceAutomationRule `json:"rule" yaml:"rule"`
}
  // The base class definition for rule
type GetSellersOffersActionResOffersSellingModePriceAutomationRule struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for minimalPrice
type GetSellersOffersActionResOffersSellingModeMinimalPrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for startingPrice
type GetSellersOffersActionResOffersSellingModeStartingPrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for saleInfo
type GetSellersOffersActionResOffersSaleInfo struct {
		CurrentPrice  GetSellersOffersActionResOffersSaleInfoCurrentPrice `json:"currentPrice" yaml:"currentPrice"`
		BiddersCount int `json:"biddersCount" yaml:"biddersCount"`
}
  // The base class definition for currentPrice
type GetSellersOffersActionResOffersSaleInfoCurrentPrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for stock
type GetSellersOffersActionResOffersStock struct {
		Available int `json:"available" yaml:"available"`
		Sold int `json:"sold" yaml:"sold"`
}
  // The base class definition for stats
type GetSellersOffersActionResOffersStats struct {
		WatchersCount int `json:"watchersCount" yaml:"watchersCount"`
		VisitsCount int `json:"visitsCount" yaml:"visitsCount"`
}
  // The base class definition for publication
type GetSellersOffersActionResOffersPublication struct {
		Status string `json:"status" yaml:"status"`
		StartingAt string `json:"startingAt" yaml:"startingAt"`
		StartedAt string `json:"startedAt" yaml:"startedAt"`
		EndingAt string `json:"endingAt" yaml:"endingAt"`
		EndedAt string `json:"endedAt" yaml:"endedAt"`
		Marketplaces  GetSellersOffersActionResOffersPublicationMarketplaces `json:"marketplaces" yaml:"marketplaces"`
}
  // The base class definition for marketplaces
type GetSellersOffersActionResOffersPublicationMarketplaces struct {
		Base  GetSellersOffersActionResOffersPublicationMarketplacesBase `json:"base" yaml:"base"`
		Additional []GetSellersOffersActionResOffersPublicationMarketplacesAdditional `json:"additional" yaml:"additional"`
}
  // The base class definition for base
type GetSellersOffersActionResOffersPublicationMarketplacesBase struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for additional
type GetSellersOffersActionResOffersPublicationMarketplacesAdditional struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for afterSalesServices
type GetSellersOffersActionResOffersAfterSalesServices struct {
		ImpliedWarranty  GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
		ReturnPolicy  GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy `json:"returnPolicy" yaml:"returnPolicy"`
		Warranty  GetSellersOffersActionResOffersAfterSalesServicesWarranty `json:"warranty" yaml:"warranty"`
}
  // The base class definition for impliedWarranty
type GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for returnPolicy
type GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for warranty
type GetSellersOffersActionResOffersAfterSalesServicesWarranty struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for additionalServices
type GetSellersOffersActionResOffersAdditionalServices struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for external
type GetSellersOffersActionResOffersExternal struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for delivery
type GetSellersOffersActionResOffersDelivery struct {
		ShippingRates  GetSellersOffersActionResOffersDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}
  // The base class definition for shippingRates
type GetSellersOffersActionResOffersDeliveryShippingRates struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for b2b
type GetSellersOffersActionResOffersB2b struct {
		BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}
  // The base class definition for fundraisingCampaign
type GetSellersOffersActionResOffersFundraisingCampaign struct {
		Id string `json:"id" yaml:"id"`
}
type GetSellersOffersActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// GetSellersOffersActionRaw registers a raw Gin route for the GetSellersOffersAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetSellersOffersActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetSellersOffersActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// GetSellersOffersActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetSellersOffersAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetSellersOffersActionHandler(
	handler func(c GetSellersOffersActionRequest, gin *gin.Context) (*GetSellersOffersActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetSellersOffersActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetSellersOffersActionRequest{
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
// GetSellersOffersAction is a high-level convenience wrapper around GetSellersOffersActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetSellersOffersAction(r gin.IRoutes, handler func(c GetSellersOffersActionRequest, gin *gin.Context) (*GetSellersOffersActionResponse, error),) {
	method, url, h := GetSellersOffersActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for Get sellers offersAction
 */
// Query wrapper with private fields
type GetSellersOffersActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func GetSellersOffersActionQueryFromString(rawQuery string) GetSellersOffersActionQuery {
	v := GetSellersOffersActionQuery{}
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
func GetSellersOffersActionQueryFromGin(c *gin.Context) GetSellersOffersActionQuery {
	return GetSellersOffersActionQueryFromString(c.Request.URL.RawQuery)
}
func GetSellersOffersActionQueryFromHttp(r *http.Request) GetSellersOffersActionQuery {
	return GetSellersOffersActionQueryFromString(r.URL.RawQuery)
}
func (q GetSellersOffersActionQuery) Values() url.Values {
	return q.values
}
func (q GetSellersOffersActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetSellersOffersActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetSellersOffersActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type GetSellersOffersActionRequest struct {
	QueryParams url.Values
	Headers http.Header
}
type GetSellersOffersActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func GetSellersOffersActionCall(
	req GetSellersOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetSellersOffersActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetSellersOffersActionMeta()
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
	var result GetSellersOffersActionResult
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