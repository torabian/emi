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
* Action to communicate with the action GetAllAvailableOfferPromotionPackagesAction
 */
func GetAllAvailableOfferPromotionPackagesActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "GetAllAvailableOfferPromotionPackagesAction",
		URL:    "https://api.{environment}/sale/offer-promotion-packages",
		Method: "GET",
	}
}

// The base class definition for getAllAvailableOfferPromotionPackagesActionRes
type GetAllAvailableOfferPromotionPackagesActionRes struct {
	MarketplaceId          string                                                                 `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackages           []GetAllAvailableOfferPromotionPackagesActionResBasePackages           `json:"basePackages" yaml:"basePackages"`
	ExtraPackages          []GetAllAvailableOfferPromotionPackagesActionResExtraPackages          `json:"extraPackages" yaml:"extraPackages"`
	AdditionalMarketplaces []GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for basePackages
type GetAllAvailableOfferPromotionPackagesActionResBasePackages struct {
	Id            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	CycleDuration string `json:"cycleDuration" yaml:"cycleDuration"`
}

// The base class definition for extraPackages
type GetAllAvailableOfferPromotionPackagesActionResExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	CycleDuration string `json:"cycleDuration" yaml:"cycleDuration"`
}

// The base class definition for additionalMarketplaces
type GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces struct {
	MarketplaceId string                                                                              `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackages  []GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages  `json:"basePackages" yaml:"basePackages"`
	ExtraPackages []GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages `json:"extraPackages" yaml:"extraPackages"`
}

// The base class definition for basePackages
type GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages struct {
	Id            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	CycleDuration string `json:"cycleDuration" yaml:"cycleDuration"`
}

// The base class definition for extraPackages
type GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	CycleDuration string `json:"cycleDuration" yaml:"cycleDuration"`
}
type GetAllAvailableOfferPromotionPackagesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// GetAllAvailableOfferPromotionPackagesActionRaw registers a raw Gin route for the GetAllAvailableOfferPromotionPackagesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetAllAvailableOfferPromotionPackagesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetAllAvailableOfferPromotionPackagesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // GetAllAvailableOfferPromotionPackagesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetAllAvailableOfferPromotionPackagesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetAllAvailableOfferPromotionPackagesActionHandler(
	handler func(c GetAllAvailableOfferPromotionPackagesActionRequest, gin *gin.Context) (*GetAllAvailableOfferPromotionPackagesActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetAllAvailableOfferPromotionPackagesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetAllAvailableOfferPromotionPackagesActionRequest{
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

// GetAllAvailableOfferPromotionPackagesAction is a high-level convenience wrapper around GetAllAvailableOfferPromotionPackagesActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetAllAvailableOfferPromotionPackagesAction(r gin.IRoutes, handler func(c GetAllAvailableOfferPromotionPackagesActionRequest, gin *gin.Context) (*GetAllAvailableOfferPromotionPackagesActionResponse, error)) {
	method, url, h := GetAllAvailableOfferPromotionPackagesActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get all available offer promotion packagesAction
 */
// Query wrapper with private fields
type GetAllAvailableOfferPromotionPackagesActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetAllAvailableOfferPromotionPackagesActionQueryFromString(rawQuery string) GetAllAvailableOfferPromotionPackagesActionQuery {
	v := GetAllAvailableOfferPromotionPackagesActionQuery{}
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
func GetAllAvailableOfferPromotionPackagesActionQueryFromGin(c *gin.Context) GetAllAvailableOfferPromotionPackagesActionQuery {
	return GetAllAvailableOfferPromotionPackagesActionQueryFromString(c.Request.URL.RawQuery)
}
func GetAllAvailableOfferPromotionPackagesActionQueryFromHttp(r *http.Request) GetAllAvailableOfferPromotionPackagesActionQuery {
	return GetAllAvailableOfferPromotionPackagesActionQueryFromString(r.URL.RawQuery)
}
func (q GetAllAvailableOfferPromotionPackagesActionQuery) Values() url.Values {
	return q.values
}
func (q GetAllAvailableOfferPromotionPackagesActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetAllAvailableOfferPromotionPackagesActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetAllAvailableOfferPromotionPackagesActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetAllAvailableOfferPromotionPackagesActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type GetAllAvailableOfferPromotionPackagesActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func GetAllAvailableOfferPromotionPackagesActionCall(
	req GetAllAvailableOfferPromotionPackagesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetAllAvailableOfferPromotionPackagesActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetAllAvailableOfferPromotionPackagesActionMeta()
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
	var result GetAllAvailableOfferPromotionPackagesActionResult
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
