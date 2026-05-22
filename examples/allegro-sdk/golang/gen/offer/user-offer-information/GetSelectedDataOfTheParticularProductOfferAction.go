package external

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action GetSelectedDataOfTheParticularProductOfferAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetSelectedDataOfTheParticularProductOfferAction
func GetSelectedDataOfTheParticularProductOfferAction(c GetSelectedDataOfTheParticularProductOfferActionRequest) (*GetSelectedDataOfTheParticularProductOfferActionResponse, error) {
	return &GetSelectedDataOfTheParticularProductOfferActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetSelectedDataOfTheParticularProductOfferActionMeta() struct {
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
		Name:        "GetSelectedDataOfTheParticularProductOfferAction",
		CliName:     "get selected data of the particular product-offer-action",
		URL:         "https://api.{environment}/sale/product-offers/{offerId}/parts",
		Method:      "GET",
		Description: `Use this resource to retrieve selected data of the particular product-offer. The model and functionality is a subset of the full product offer get endpoint (GET /sale/product-offers/{offerId}), but it is faster and more reliable.`,
	}
}
func GetGetSelectedDataOfTheParticularProductOfferActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "string",
			Description: "Unique offer identifier",
		},
		{
			Name:     prefix + "stock",
			Type:     "object",
			Children: GetGetSelectedDataOfTheParticularProductOfferActionResStockCliFlags("stock-"),
		},
		{
			Name:     prefix + "selling-mode",
			Type:     "object",
			Children: GetGetSelectedDataOfTheParticularProductOfferActionResSellingModeCliFlags("selling-mode-"),
		},
		{
			Name:        prefix + "additional-marketplaces",
			Type:        "object",
			Children:    GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesCliFlags("additional-marketplaces-"),
			Description: "Marketplace-specific price information",
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResFromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionRes {
	data := GetSelectedDataOfTheParticularProductOfferActionRes{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("stock") {
		data.Stock = CastGetSelectedDataOfTheParticularProductOfferActionResStockFromCli(c)
	}
	if c.IsSet("selling-mode") {
		data.SellingMode = CastGetSelectedDataOfTheParticularProductOfferActionResSellingModeFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesFromCli(c)
	}
	return data
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

func GetGetSelectedDataOfTheParticularProductOfferActionResStockCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "available",
			Type:        "int",
			Description: "Number of available items in stock",
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResStockFromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResStock {
	data := GetSelectedDataOfTheParticularProductOfferActionResStock{}
	if c.IsSet("available") {
		data.Available = int(c.Int64("available"))
	}
	return data
}

// The base class definition for stock
type GetSelectedDataOfTheParticularProductOfferActionResStock struct {
	// Number of available items in stock
	Available int `json:"available" yaml:"available"`
}

func GetGetSelectedDataOfTheParticularProductOfferActionResSellingModeCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "price",
			Type:     "object",
			Children: GetGetSelectedDataOfTheParticularProductOfferActionResSellingModePriceCliFlags("price-"),
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResSellingModeFromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResSellingMode {
	data := GetSelectedDataOfTheParticularProductOfferActionResSellingMode{}
	if c.IsSet("price") {
		data.Price = CastGetSelectedDataOfTheParticularProductOfferActionResSellingModePriceFromCli(c)
	}
	return data
}

// The base class definition for sellingMode
type GetSelectedDataOfTheParticularProductOfferActionResSellingMode struct {
	Price GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice `json:"price" yaml:"price"`
}

func GetGetSelectedDataOfTheParticularProductOfferActionResSellingModePriceCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "amount",
			Type: "string",
		},
		{
			Name: prefix + "currency",
			Type: "string",
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResSellingModePriceFromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice {
	data := GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for price
type GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "marketplace-id1",
			Type:     "object",
			Children: GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1CliFlags("marketplace-id1-"),
		},
		{
			Name:     prefix + "marketplace-id2",
			Type:     "object",
			Children: GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2CliFlags("marketplace-id2-"),
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesFromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplaces {
	data := GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplaces{}
	if c.IsSet("marketplace-id1") {
		data.MarketplaceId1 = CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1FromCli(c)
	}
	if c.IsSet("marketplace-id2") {
		data.MarketplaceId2 = CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2FromCli(c)
	}
	return data
}

// The base class definition for additionalMarketplaces
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplaces struct {
	MarketplaceId1 GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1 `json:"marketplaceId1" yaml:"marketplaceId1"`
	MarketplaceId2 GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2 `json:"marketplaceId2" yaml:"marketplaceId2"`
}

func GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "selling-mode",
			Type:     "object",
			Children: GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModeCliFlags("selling-mode-"),
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1FromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1 {
	data := GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1{}
	if c.IsSet("selling-mode") {
		data.SellingMode = CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModeFromCli(c)
	}
	return data
}

// The base class definition for marketplaceId1
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1 struct {
	SellingMode GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode `json:"sellingMode" yaml:"sellingMode"`
}

func GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModeCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "price",
			Type:     "object",
			Children: GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePriceCliFlags("price-"),
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModeFromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode {
	data := GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode{}
	if c.IsSet("price") {
		data.Price = CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePriceFromCli(c)
	}
	return data
}

// The base class definition for sellingMode
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode struct {
	Price GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice `json:"price" yaml:"price"`
}

func GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePriceCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "amount",
			Type: "string",
		},
		{
			Name: prefix + "currency",
			Type: "string",
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePriceFromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice {
	data := GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for price
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "selling-mode",
			Type:     "object",
			Children: GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModeCliFlags("selling-mode-"),
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2FromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2 {
	data := GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2{}
	if c.IsSet("selling-mode") {
		data.SellingMode = CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModeFromCli(c)
	}
	return data
}

// The base class definition for marketplaceId2
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2 struct {
	SellingMode GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode `json:"sellingMode" yaml:"sellingMode"`
}

func GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModeCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "price",
			Type:     "object",
			Children: GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePriceCliFlags("price-"),
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModeFromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode {
	data := GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode{}
	if c.IsSet("price") {
		data.Price = CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePriceFromCli(c)
	}
	return data
}

// The base class definition for sellingMode
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode struct {
	Price GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice `json:"price" yaml:"price"`
}

func GetGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePriceCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "amount",
			Type: "string",
		},
		{
			Name: prefix + "currency",
			Type: "string",
		},
	}
}
func CastGetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePriceFromCli(c emigo.CliCastable) GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice {
	data := GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for price
type GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func (x *GetSelectedDataOfTheParticularProductOfferActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetSelectedDataOfTheParticularProductOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetSelectedDataOfTheParticularProductOfferActionResponse) SetContentType(contentType string) *GetSelectedDataOfTheParticularProductOfferActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetSelectedDataOfTheParticularProductOfferActionResponse) AsStream(r io.Reader, contentType string) *GetSelectedDataOfTheParticularProductOfferActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetSelectedDataOfTheParticularProductOfferActionResponse) AsJSON(payload any) *GetSelectedDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetSelectedDataOfTheParticularProductOfferActionResponse) WithIdeal(payload GetSelectedDataOfTheParticularProductOfferActionRes) *GetSelectedDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	return x
}
func (x *GetSelectedDataOfTheParticularProductOfferActionResponse) AsHTML(payload string) *GetSelectedDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetSelectedDataOfTheParticularProductOfferActionResponse) AsBytes(payload []byte) *GetSelectedDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetSelectedDataOfTheParticularProductOfferActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetSelectedDataOfTheParticularProductOfferActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetSelectedDataOfTheParticularProductOfferActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetSelectedDataOfTheParticularProductOfferActionRaw registers a raw Gin route for the GetSelectedDataOfTheParticularProductOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetSelectedDataOfTheParticularProductOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetSelectedDataOfTheParticularProductOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetSelectedDataOfTheParticularProductOfferActionRequestSig = func(c GetSelectedDataOfTheParticularProductOfferActionRequest) (*GetSelectedDataOfTheParticularProductOfferActionResponse, error)

// GetSelectedDataOfTheParticularProductOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetSelectedDataOfTheParticularProductOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetSelectedDataOfTheParticularProductOfferActionHandler(
	handler GetSelectedDataOfTheParticularProductOfferActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetSelectedDataOfTheParticularProductOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetSelectedDataOfTheParticularProductOfferActionRequest{
			Body:        nil,
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

// GetSelectedDataOfTheParticularProductOfferAction is a high-level convenience wrapper around GetSelectedDataOfTheParticularProductOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetSelectedDataOfTheParticularProductOfferActionGin(r gin.IRoutes, handler GetSelectedDataOfTheParticularProductOfferActionRequestSig) {
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
	Body        interface{}
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	GinCtx *gin.Context
	// Urfave context, per each request
	CliCtx *cli.Command
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func (x GetSelectedDataOfTheParticularProductOfferActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GetSelectedDataOfTheParticularProductOfferActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GetSelectedDataOfTheParticularProductOfferActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GetSelectedDataOfTheParticularProductOfferActionClientCreateUrl(
	req GetSelectedDataOfTheParticularProductOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetSelectedDataOfTheParticularProductOfferActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// Build final URL with query string
	u, err := url.Parse(urlAddr)
	if err != nil {
		return nil, err
	}
	// if UrlValues present, encode and append
	if len(req.QueryParams) > 0 {
		u.RawQuery = req.QueryParams.Encode()
	}
	return u, nil
}
func GetSelectedDataOfTheParticularProductOfferActionClientExecuteTyped(httpReq *http.Request) (*GetSelectedDataOfTheParticularProductOfferActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetSelectedDataOfTheParticularProductOfferActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetSelectedDataOfTheParticularProductOfferActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetSelectedDataOfTheParticularProductOfferActionResponse{Payload: result}, err
	}
	return &GetSelectedDataOfTheParticularProductOfferActionResponse{Payload: result}, nil
}
func GetSelectedDataOfTheParticularProductOfferActionClientBuildRequest(req GetSelectedDataOfTheParticularProductOfferActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetSelectedDataOfTheParticularProductOfferActionMeta()
	httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header = make(http.Header)
	// copy defaults
	for k, v := range config.Headers {
		for _, vv := range v {
			httpReq.Header.Add(k, vv)
		}
	}
	// override with request-specific headers
	for k, v := range req.Headers {
		httpReq.Header.Del(k) // ensure override, not duplicate
		for _, vv := range v {
			httpReq.Header.Add(k, vv)
		}
	}
	return httpReq, nil
}
func GetSelectedDataOfTheParticularProductOfferActionCall(
	req GetSelectedDataOfTheParticularProductOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetSelectedDataOfTheParticularProductOfferActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetSelectedDataOfTheParticularProductOfferActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetSelectedDataOfTheParticularProductOfferActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetSelectedDataOfTheParticularProductOfferActionClientExecuteTyped(r)
}
