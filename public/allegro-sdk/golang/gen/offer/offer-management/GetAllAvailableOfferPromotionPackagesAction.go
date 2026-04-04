package external

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action GetAllAvailableOfferPromotionPackagesAction
 */
func GetAllAvailableOfferPromotionPackagesActionMeta() struct {
	Name        string
	CliName     string
	URL         string
	Method      string
	Description string
} {
	return struct {
		Name        string
		CliName     string
		URL         string
		Method      string
		Description string
	}{
		Name:        "GetAllAvailableOfferPromotionPackagesAction",
		CliName:     "get all available offer promotion packages-action",
		URL:         "https://api.{environment}/sale/offer-promotion-packages",
		Method:      "GET",
		Description: `Use this resource to retrieve all available offer promotion packages. Read more: PL / EN.`,
	}
}
func GetGetAllAvailableOfferPromotionPackagesActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name: prefix + "base-packages",
			Type: "array",
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "array",
		},
	}
}
func CastGetAllAvailableOfferPromotionPackagesActionResFromCli(c emigo.CliCastable) GetAllAvailableOfferPromotionPackagesActionRes {
	data := GetAllAvailableOfferPromotionPackagesActionRes{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("base-packages") {
		data.BasePackages = emigo.CapturePossibleArray(CastGetAllAvailableOfferPromotionPackagesActionResBasePackagesFromCli, "base-packages", c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastGetAllAvailableOfferPromotionPackagesActionResExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = emigo.CapturePossibleArray(CastGetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesFromCli, "additional-marketplaces", c)
	}
	return data
}

// The base class definition for getAllAvailableOfferPromotionPackagesActionRes
type GetAllAvailableOfferPromotionPackagesActionRes struct {
	MarketplaceId          string                                                                 `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackages           []GetAllAvailableOfferPromotionPackagesActionResBasePackages           `json:"basePackages" yaml:"basePackages"`
	ExtraPackages          []GetAllAvailableOfferPromotionPackagesActionResExtraPackages          `json:"extraPackages" yaml:"extraPackages"`
	AdditionalMarketplaces []GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

func GetGetAllAvailableOfferPromotionPackagesActionResBasePackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "cycle-duration",
			Type: "string",
		},
	}
}
func CastGetAllAvailableOfferPromotionPackagesActionResBasePackagesFromCli(c emigo.CliCastable) GetAllAvailableOfferPromotionPackagesActionResBasePackages {
	data := GetAllAvailableOfferPromotionPackagesActionResBasePackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("cycle-duration") {
		data.CycleDuration = c.String("cycle-duration")
	}
	return data
}

// The base class definition for basePackages
type GetAllAvailableOfferPromotionPackagesActionResBasePackages struct {
	Id            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	CycleDuration string `json:"cycleDuration" yaml:"cycleDuration"`
}

func GetGetAllAvailableOfferPromotionPackagesActionResExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "cycle-duration",
			Type: "string",
		},
	}
}
func CastGetAllAvailableOfferPromotionPackagesActionResExtraPackagesFromCli(c emigo.CliCastable) GetAllAvailableOfferPromotionPackagesActionResExtraPackages {
	data := GetAllAvailableOfferPromotionPackagesActionResExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("cycle-duration") {
		data.CycleDuration = c.String("cycle-duration")
	}
	return data
}

// The base class definition for extraPackages
type GetAllAvailableOfferPromotionPackagesActionResExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	CycleDuration string `json:"cycleDuration" yaml:"cycleDuration"`
}

func GetGetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name: prefix + "base-packages",
			Type: "array",
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
	}
}
func CastGetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesFromCli(c emigo.CliCastable) GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces {
	data := GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("base-packages") {
		data.BasePackages = emigo.CapturePossibleArray(CastGetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackagesFromCli, "base-packages", c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastGetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackagesFromCli, "extra-packages", c)
	}
	return data
}

// The base class definition for additionalMarketplaces
type GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces struct {
	MarketplaceId string                                                                              `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackages  []GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages  `json:"basePackages" yaml:"basePackages"`
	ExtraPackages []GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages `json:"extraPackages" yaml:"extraPackages"`
}

func GetGetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "cycle-duration",
			Type: "string",
		},
	}
}
func CastGetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackagesFromCli(c emigo.CliCastable) GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages {
	data := GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("cycle-duration") {
		data.CycleDuration = c.String("cycle-duration")
	}
	return data
}

// The base class definition for basePackages
type GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages struct {
	Id            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	CycleDuration string `json:"cycleDuration" yaml:"cycleDuration"`
}

func GetGetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "cycle-duration",
			Type: "string",
		},
	}
}
func CastGetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackagesFromCli(c emigo.CliCastable) GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages {
	data := GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("cycle-duration") {
		data.CycleDuration = c.String("cycle-duration")
	}
	return data
}

// The base class definition for extraPackages
type GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	CycleDuration string `json:"cycleDuration" yaml:"cycleDuration"`
}

func (x *GetAllAvailableOfferPromotionPackagesActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetAllAvailableOfferPromotionPackagesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *GetAllAvailableOfferPromotionPackagesActionResponse) SetContentType(contentType string) *GetAllAvailableOfferPromotionPackagesActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetAllAvailableOfferPromotionPackagesActionResponse) AsStream(r io.Reader, contentType string) *GetAllAvailableOfferPromotionPackagesActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetAllAvailableOfferPromotionPackagesActionResponse) AsJSON(payload any) *GetAllAvailableOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *GetAllAvailableOfferPromotionPackagesActionResponse) AsHTML(payload string) *GetAllAvailableOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetAllAvailableOfferPromotionPackagesActionResponse) AsBytes(payload []byte) *GetAllAvailableOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetAllAvailableOfferPromotionPackagesActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetAllAvailableOfferPromotionPackagesActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetAllAvailableOfferPromotionPackagesActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetAllAvailableOfferPromotionPackagesActionRaw registers a raw Gin route for the GetAllAvailableOfferPromotionPackagesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetAllAvailableOfferPromotionPackagesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetAllAvailableOfferPromotionPackagesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetAllAvailableOfferPromotionPackagesActionRequestSig = func(c GetAllAvailableOfferPromotionPackagesActionRequest) (*GetAllAvailableOfferPromotionPackagesActionResponse, error)

// GetAllAvailableOfferPromotionPackagesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetAllAvailableOfferPromotionPackagesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetAllAvailableOfferPromotionPackagesActionHandler(
	handler GetAllAvailableOfferPromotionPackagesActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetAllAvailableOfferPromotionPackagesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetAllAvailableOfferPromotionPackagesActionRequest{
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx:      m,
		}
		resp, err := handler(req)
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
func GetAllAvailableOfferPromotionPackagesActionGin(r gin.IRoutes, handler GetAllAvailableOfferPromotionPackagesActionRequestSig) {
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
	GinCtx      *gin.Context
	CliCtx      *cli.Context
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
