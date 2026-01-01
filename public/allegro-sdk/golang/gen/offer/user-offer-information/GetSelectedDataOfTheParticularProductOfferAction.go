package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
)

/**
* Action to communicate with the action GetSelectedDataOfTheParticularProductOfferAction
 */
func GetSelectedDataOfTheParticularProductOfferActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "GetSelectedDataOfTheParticularProductOfferAction",
		URL:    "https://api.{environment}/sale/product-offers/{offerId}/parts",
		Method: "GET",
	}
}

// The base class definition for getSelectedDataOfTheParticularProductOfferActionRes
type GetSelectedDataOfTheParticularProductOfferActionRes struct {
	// Unique offer identifier
	Id          string                                                         `json:"id" yaml:"id"`
	Stock       GetSelectedDataOfTheParticularProductOfferActionResStock       `json:"stock" yaml:"stock"`
	SellingMode GetSelectedDataOfTheParticularProductOfferActionResSellingMode `json:"sellingMode" yaml:"sellingMode"`
	// Marketplace-specific price information
	AdditionalMarketplaces GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for stock
type GetSelectedDataOfTheParticularProductOfferActionResStock struct {
	// Number of available items in stock
	Available int `json:"available" yaml:"available"`
}

// The base class definition for sellingMode
type GetSelectedDataOfTheParticularProductOfferActionResSellingMode struct {
	Price GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice `json:"price" yaml:"price"`
}

// The base class definition for price
type GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for additionalMarketplaces
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplaces struct {
	MarketplaceId1 GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1 `json:"marketplaceId1" yaml:"marketplaceId1"`
	MarketplaceId2 GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2 `json:"marketplaceId2" yaml:"marketplaceId2"`
}

// The base class definition for marketplaceId1
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1 struct {
	SellingMode GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode `json:"sellingMode" yaml:"sellingMode"`
}

// The base class definition for sellingMode
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode struct {
	Price GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice `json:"price" yaml:"price"`
}

// The base class definition for price
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for marketplaceId2
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2 struct {
	SellingMode GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode `json:"sellingMode" yaml:"sellingMode"`
}

// The base class definition for sellingMode
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode struct {
	Price GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice `json:"price" yaml:"price"`
}

// The base class definition for price
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}
type GetSelectedDataOfTheParticularProductOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// GetSelectedDataOfTheParticularProductOfferActionRaw registers a raw Gin route for the GetSelectedDataOfTheParticularProductOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetSelectedDataOfTheParticularProductOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetSelectedDataOfTheParticularProductOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // GetSelectedDataOfTheParticularProductOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetSelectedDataOfTheParticularProductOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetSelectedDataOfTheParticularProductOfferActionHandler(
	handler func(c GetSelectedDataOfTheParticularProductOfferActionRequest, gin *gin.Context) (*GetSelectedDataOfTheParticularProductOfferActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetSelectedDataOfTheParticularProductOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetSelectedDataOfTheParticularProductOfferActionRequest{
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

// GetSelectedDataOfTheParticularProductOfferAction is a high-level convenience wrapper around GetSelectedDataOfTheParticularProductOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetSelectedDataOfTheParticularProductOfferAction(r gin.IRoutes, handler func(c GetSelectedDataOfTheParticularProductOfferActionRequest, gin *gin.Context) (*GetSelectedDataOfTheParticularProductOfferActionResponse, error)) {
	method, url, h := GetSelectedDataOfTheParticularProductOfferActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get selected data of the particular product-offerAction
 */
// Query wrapper with private fields
type GetSelectedDataOfTheParticularProductOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetSelectedDataOfTheParticularProductOfferActionQueryFromString(rawQuery string) GetSelectedDataOfTheParticularProductOfferActionQuery {
	v := GetSelectedDataOfTheParticularProductOfferActionQuery{}
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
func GetSelectedDataOfTheParticularProductOfferActionQueryFromGin(c *gin.Context) GetSelectedDataOfTheParticularProductOfferActionQuery {
	return GetSelectedDataOfTheParticularProductOfferActionQueryFromString(c.Request.URL.RawQuery)
}
func GetSelectedDataOfTheParticularProductOfferActionQueryFromHttp(r *http.Request) GetSelectedDataOfTheParticularProductOfferActionQuery {
	return GetSelectedDataOfTheParticularProductOfferActionQueryFromString(r.URL.RawQuery)
}
func (q GetSelectedDataOfTheParticularProductOfferActionQuery) Values() url.Values {
	return q.values
}
func (q GetSelectedDataOfTheParticularProductOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetSelectedDataOfTheParticularProductOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetSelectedDataOfTheParticularProductOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetSelectedDataOfTheParticularProductOfferActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type GetSelectedDataOfTheParticularProductOfferActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func GetSelectedDataOfTheParticularProductOfferActionCall(
	req GetSelectedDataOfTheParticularProductOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetSelectedDataOfTheParticularProductOfferActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetSelectedDataOfTheParticularProductOfferActionMeta()
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
	var result GetSelectedDataOfTheParticularProductOfferActionResult
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
