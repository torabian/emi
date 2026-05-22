package external

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ModifyTheBuyNowPriceInAnOfferAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ModifyTheBuyNowPriceInAnOfferAction
func ModifyTheBuyNowPriceInAnOfferAction(c ModifyTheBuyNowPriceInAnOfferActionRequest) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error) {
	return &ModifyTheBuyNowPriceInAnOfferActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ModifyTheBuyNowPriceInAnOfferActionMeta() struct {
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
		Name:        "ModifyTheBuyNowPriceInAnOfferAction",
		CliName:     "modify the -buy -now price in an offer-action",
		URL:         "https://api.{environment}/offers/{offerId}/change-price-commands/{commandId}",
		Method:      "PUT",
		Description: `Use this resource to change the Buy Now price in a single offer. Read more: PL / EN.`,
	}
}
func GetModifyTheBuyNowPriceInAnOfferActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name:     prefix + "input",
			Type:     "object",
			Children: GetModifyTheBuyNowPriceInAnOfferActionReqInputCliFlags("input-"),
		},
	}
}
func CastModifyTheBuyNowPriceInAnOfferActionReqFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionReq {
	data := ModifyTheBuyNowPriceInAnOfferActionReq{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("input") {
		data.Input = CastModifyTheBuyNowPriceInAnOfferActionReqInputFromCli(c)
	}
	return data
}

// The base class definition for modifyTheBuyNowPriceInAnOfferActionReq
type ModifyTheBuyNowPriceInAnOfferActionReq struct {
	Id    string                                      `json:"id" yaml:"id"`
	Input ModifyTheBuyNowPriceInAnOfferActionReqInput `json:"input" yaml:"input"`
}

func GetModifyTheBuyNowPriceInAnOfferActionReqInputCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "buy-now-price",
			Type:     "object",
			Children: GetModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPriceCliFlags("buy-now-price-"),
		},
	}
}
func CastModifyTheBuyNowPriceInAnOfferActionReqInputFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionReqInput {
	data := ModifyTheBuyNowPriceInAnOfferActionReqInput{}
	if c.IsSet("buy-now-price") {
		data.BuyNowPrice = CastModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPriceFromCli(c)
	}
	return data
}

// The base class definition for input
type ModifyTheBuyNowPriceInAnOfferActionReqInput struct {
	BuyNowPrice ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice `json:"buyNowPrice" yaml:"buyNowPrice"`
}

func GetModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPriceFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice {
	data := ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for buyNowPrice
type ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func (x *ModifyTheBuyNowPriceInAnOfferActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetModifyTheBuyNowPriceInAnOfferActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name:     prefix + "input",
			Type:     "object",
			Children: GetModifyTheBuyNowPriceInAnOfferActionResInputCliFlags("input-"),
		},
		{
			Name:     prefix + "output",
			Type:     "object",
			Children: GetModifyTheBuyNowPriceInAnOfferActionResOutputCliFlags("output-"),
		},
	}
}
func CastModifyTheBuyNowPriceInAnOfferActionResFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionRes {
	data := ModifyTheBuyNowPriceInAnOfferActionRes{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("input") {
		data.Input = CastModifyTheBuyNowPriceInAnOfferActionResInputFromCli(c)
	}
	if c.IsSet("output") {
		data.Output = CastModifyTheBuyNowPriceInAnOfferActionResOutputFromCli(c)
	}
	return data
}

// The base class definition for modifyTheBuyNowPriceInAnOfferActionRes
type ModifyTheBuyNowPriceInAnOfferActionRes struct {
	Id     string                                       `json:"id" yaml:"id"`
	Input  ModifyTheBuyNowPriceInAnOfferActionResInput  `json:"input" yaml:"input"`
	Output ModifyTheBuyNowPriceInAnOfferActionResOutput `json:"output" yaml:"output"`
}

func GetModifyTheBuyNowPriceInAnOfferActionResInputCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "buy-now-price",
			Type:     "object",
			Children: GetModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPriceCliFlags("buy-now-price-"),
		},
	}
}
func CastModifyTheBuyNowPriceInAnOfferActionResInputFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionResInput {
	data := ModifyTheBuyNowPriceInAnOfferActionResInput{}
	if c.IsSet("buy-now-price") {
		data.BuyNowPrice = CastModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPriceFromCli(c)
	}
	return data
}

// The base class definition for input
type ModifyTheBuyNowPriceInAnOfferActionResInput struct {
	BuyNowPrice ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice `json:"buyNowPrice" yaml:"buyNowPrice"`
}

func GetModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPriceFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice {
	data := ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for buyNowPrice
type ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetModifyTheBuyNowPriceInAnOfferActionResOutputCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "status",
			Type: "string",
		},
		{
			Name: prefix + "errors",
			Type: "array",
		},
	}
}
func CastModifyTheBuyNowPriceInAnOfferActionResOutputFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionResOutput {
	data := ModifyTheBuyNowPriceInAnOfferActionResOutput{}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("errors") {
		data.Errors = emigo.CapturePossibleArray(CastModifyTheBuyNowPriceInAnOfferActionResOutputErrorsFromCli, "errors", c)
	}
	return data
}

// The base class definition for output
type ModifyTheBuyNowPriceInAnOfferActionResOutput struct {
	Status string                                               `json:"status" yaml:"status"`
	Errors []ModifyTheBuyNowPriceInAnOfferActionResOutputErrors `json:"errors" yaml:"errors"`
}

func GetModifyTheBuyNowPriceInAnOfferActionResOutputErrorsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "code",
			Type: "string",
		},
		{
			Name: prefix + "details",
			Type: "string",
		},
		{
			Name: prefix + "message",
			Type: "string",
		},
		{
			Name: prefix + "path",
			Type: "string",
		},
		{
			Name: prefix + "user-message",
			Type: "string",
		},
		{
			Name:     prefix + "metadata",
			Type:     "object",
			Children: GetModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadataCliFlags("metadata-"),
		},
	}
}
func CastModifyTheBuyNowPriceInAnOfferActionResOutputErrorsFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionResOutputErrors {
	data := ModifyTheBuyNowPriceInAnOfferActionResOutputErrors{}
	if c.IsSet("code") {
		data.Code = c.String("code")
	}
	if c.IsSet("details") {
		data.Details = c.String("details")
	}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("path") {
		data.Path = c.String("path")
	}
	if c.IsSet("user-message") {
		data.UserMessage = c.String("user-message")
	}
	if c.IsSet("metadata") {
		data.Metadata = CastModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadataFromCli(c)
	}
	return data
}

// The base class definition for errors
type ModifyTheBuyNowPriceInAnOfferActionResOutputErrors struct {
	Code        string                                                     `json:"code" yaml:"code"`
	Details     string                                                     `json:"details" yaml:"details"`
	Message     string                                                     `json:"message" yaml:"message"`
	Path        string                                                     `json:"path" yaml:"path"`
	UserMessage string                                                     `json:"userMessage" yaml:"userMessage"`
	Metadata    ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata `json:"metadata" yaml:"metadata"`
}

func GetModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadataCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "product-id",
			Type: "string",
		},
	}
}
func CastModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadataFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata {
	data := ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata{}
	if c.IsSet("product-id") {
		data.ProductId = c.String("product-id")
	}
	return data
}

// The base class definition for metadata
type ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

func (x *ModifyTheBuyNowPriceInAnOfferActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ModifyTheBuyNowPriceInAnOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) SetContentType(contentType string) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) AsStream(r io.Reader, contentType string) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) AsJSON(payload any) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) WithIdeal(payload ModifyTheBuyNowPriceInAnOfferActionRes) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = payload
	return x
}
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) AsHTML(payload string) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) AsBytes(payload []byte) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ModifyTheBuyNowPriceInAnOfferActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ModifyTheBuyNowPriceInAnOfferActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ModifyTheBuyNowPriceInAnOfferActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ModifyTheBuyNowPriceInAnOfferActionRaw registers a raw Gin route for the ModifyTheBuyNowPriceInAnOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModifyTheBuyNowPriceInAnOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ModifyTheBuyNowPriceInAnOfferActionRequestSig = func(c ModifyTheBuyNowPriceInAnOfferActionRequest) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error)

// ModifyTheBuyNowPriceInAnOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModifyTheBuyNowPriceInAnOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModifyTheBuyNowPriceInAnOfferActionHandler(
	handler ModifyTheBuyNowPriceInAnOfferActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ModifyTheBuyNowPriceInAnOfferActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ModifyTheBuyNowPriceInAnOfferActionRequest{
			Body:        body,
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

// ModifyTheBuyNowPriceInAnOfferAction is a high-level convenience wrapper around ModifyTheBuyNowPriceInAnOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModifyTheBuyNowPriceInAnOfferActionGin(r gin.IRoutes, handler ModifyTheBuyNowPriceInAnOfferActionRequestSig) {
	method, url, h := ModifyTheBuyNowPriceInAnOfferActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Modify the Buy Now price in an offerAction
 */
// Query wrapper with private fields
type ModifyTheBuyNowPriceInAnOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModifyTheBuyNowPriceInAnOfferActionQueryFromString(rawQuery string) ModifyTheBuyNowPriceInAnOfferActionQuery {
	v := ModifyTheBuyNowPriceInAnOfferActionQuery{}
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
func ModifyTheBuyNowPriceInAnOfferActionQueryFromGin(c *gin.Context) ModifyTheBuyNowPriceInAnOfferActionQuery {
	return ModifyTheBuyNowPriceInAnOfferActionQueryFromString(c.Request.URL.RawQuery)
}
func ModifyTheBuyNowPriceInAnOfferActionQueryFromHttp(r *http.Request) ModifyTheBuyNowPriceInAnOfferActionQuery {
	return ModifyTheBuyNowPriceInAnOfferActionQueryFromString(r.URL.RawQuery)
}
func (q ModifyTheBuyNowPriceInAnOfferActionQuery) Values() url.Values {
	return q.values
}
func (q ModifyTheBuyNowPriceInAnOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModifyTheBuyNowPriceInAnOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModifyTheBuyNowPriceInAnOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModifyTheBuyNowPriceInAnOfferActionRequest struct {
	Body        ModifyTheBuyNowPriceInAnOfferActionReq
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

func (x ModifyTheBuyNowPriceInAnOfferActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x ModifyTheBuyNowPriceInAnOfferActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type ModifyTheBuyNowPriceInAnOfferActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func ModifyTheBuyNowPriceInAnOfferActionClientCreateUrl(
	req ModifyTheBuyNowPriceInAnOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
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
func ModifyTheBuyNowPriceInAnOfferActionClientExecuteTyped(httpReq *http.Request) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ModifyTheBuyNowPriceInAnOfferActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ModifyTheBuyNowPriceInAnOfferActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ModifyTheBuyNowPriceInAnOfferActionResponse{Payload: result}, err
	}
	return &ModifyTheBuyNowPriceInAnOfferActionResponse{Payload: result}, nil
}
func ModifyTheBuyNowPriceInAnOfferActionClientBuildRequest(req ModifyTheBuyNowPriceInAnOfferActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
	bodyBytes, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), bytes.NewReader(bodyBytes))
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
func ModifyTheBuyNowPriceInAnOfferActionCall(
	req ModifyTheBuyNowPriceInAnOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ModifyTheBuyNowPriceInAnOfferActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ModifyTheBuyNowPriceInAnOfferActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ModifyTheBuyNowPriceInAnOfferActionClientExecuteTyped(r)
}
