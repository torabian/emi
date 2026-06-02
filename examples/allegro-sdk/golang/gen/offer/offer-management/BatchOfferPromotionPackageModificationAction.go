package external

import (
	"bytes"
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action BatchOfferPromotionPackageModificationAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of BatchOfferPromotionPackageModificationAction
func BatchOfferPromotionPackageModificationAction(c BatchOfferPromotionPackageModificationActionRequest) (*BatchOfferPromotionPackageModificationActionResponse, error) {
	return &BatchOfferPromotionPackageModificationActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func BatchOfferPromotionPackageModificationActionMeta() struct {
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
		Name:        "BatchOfferPromotionPackageModificationAction",
		CliName:     "batch offer promotion package modification-action",
		URL:         "https://api.{environment}/sale/offers/promo-options-commands/{commandId}",
		Method:      "PUT",
		Description: `Use this resource to modify promotion packages on multiple offers at once. Read more: PL / EN.`,
	}
}
func GetBatchOfferPromotionPackageModificationActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "offer-criteria",
			Type: "array",
		},
		{
			Name:     prefix + "modification",
			Type:     "object",
			Children: GetBatchOfferPromotionPackageModificationActionReqModificationCliFlags("modification-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "array",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReq {
	data := BatchOfferPromotionPackageModificationActionReq{}
	if c.IsSet("offer-criteria") {
		data.OfferCriteria = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqOfferCriteriaFromCli, "offer-criteria", c)
	}
	if c.IsSet("modification") {
		data.Modification = CastBatchOfferPromotionPackageModificationActionReqModificationFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesFromCli, "additional-marketplaces", c)
	}
	return data
}

// The base class definition for batchOfferPromotionPackageModificationActionReq
type BatchOfferPromotionPackageModificationActionReq struct {
	OfferCriteria          emigo.Array[BatchOfferPromotionPackageModificationActionReqOfferCriteria]          `json:"offerCriteria" yaml:"offerCriteria"`
	Modification           BatchOfferPromotionPackageModificationActionReqModification                        `json:"modification" yaml:"modification"`
	AdditionalMarketplaces emigo.Array[BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

func GetBatchOfferPromotionPackageModificationActionReqOfferCriteriaCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "offers",
			Type: "array",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqOfferCriteriaFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqOfferCriteria {
	data := BatchOfferPromotionPackageModificationActionReqOfferCriteria{}
	if c.IsSet("offers") {
		data.Offers = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqOfferCriteriaOffersFromCli, "offers", c)
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for offerCriteria
type BatchOfferPromotionPackageModificationActionReqOfferCriteria struct {
	Offers emigo.Array[BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers] `json:"offers" yaml:"offers"`
	Type   string                                                                          `json:"type" yaml:"type"`
}

func GetBatchOfferPromotionPackageModificationActionReqOfferCriteriaOffersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqOfferCriteriaOffersFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers {
	data := BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for offers
type BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers struct {
	Id string `json:"id" yaml:"id"`
}

func GetBatchOfferPromotionPackageModificationActionReqModificationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetBatchOfferPromotionPackageModificationActionReqModificationBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name: prefix + "modification-time",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqModificationFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqModification {
	data := BatchOfferPromotionPackageModificationActionReqModification{}
	if c.IsSet("base-package") {
		data.BasePackage = CastBatchOfferPromotionPackageModificationActionReqModificationBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqModificationExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("modification-time") {
		data.ModificationTime = c.String("modification-time")
	}
	return data
}

// The base class definition for modification
type BatchOfferPromotionPackageModificationActionReqModification struct {
	BasePackage      BatchOfferPromotionPackageModificationActionReqModificationBasePackage                `json:"basePackage" yaml:"basePackage"`
	ExtraPackages    emigo.Array[BatchOfferPromotionPackageModificationActionReqModificationExtraPackages] `json:"extraPackages" yaml:"extraPackages"`
	ModificationTime string                                                                                `json:"modificationTime" yaml:"modificationTime"`
}

func GetBatchOfferPromotionPackageModificationActionReqModificationBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqModificationBasePackageFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqModificationBasePackage {
	data := BatchOfferPromotionPackageModificationActionReqModificationBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for basePackage
type BatchOfferPromotionPackageModificationActionReqModificationBasePackage struct {
	Id string `json:"id" yaml:"id"`
}

func GetBatchOfferPromotionPackageModificationActionReqModificationExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqModificationExtraPackagesFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqModificationExtraPackages {
	data := BatchOfferPromotionPackageModificationActionReqModificationExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for extraPackages
type BatchOfferPromotionPackageModificationActionReqModificationExtraPackages struct {
	Id string `json:"id" yaml:"id"`
}

func GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name:     prefix + "modification",
			Type:     "object",
			Children: GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationCliFlags("modification-"),
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces {
	data := BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("modification") {
		data.Modification = CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationFromCli(c)
	}
	return data
}

// The base class definition for additionalMarketplaces
type BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces struct {
	MarketplaceId string                                                                            `json:"marketplaceId" yaml:"marketplaceId"`
	Modification  BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification `json:"modification" yaml:"modification"`
}

func GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name: prefix + "modification-time",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification {
	data := BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification{}
	if c.IsSet("base-package") {
		data.BasePackage = CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("modification-time") {
		data.ModificationTime = c.String("modification-time")
	}
	return data
}

// The base class definition for modification
type BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification struct {
	BasePackage      BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage                `json:"basePackage" yaml:"basePackage"`
	ExtraPackages    emigo.Array[BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages] `json:"extraPackages" yaml:"extraPackages"`
	ModificationTime string                                                                                                      `json:"modificationTime" yaml:"modificationTime"`
}

func GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackageFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage {
	data := BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for basePackage
type BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage struct {
	Id string `json:"id" yaml:"id"`
}

func GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackagesFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages {
	data := BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for extraPackages
type BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages struct {
	Id string `json:"id" yaml:"id"`
}

func (x *BatchOfferPromotionPackageModificationActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetBatchOfferPromotionPackageModificationActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name:     prefix + "task-count",
			Type:     "object",
			Children: GetBatchOfferPromotionPackageModificationActionResTaskCountCliFlags("task-count-"),
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionResFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionRes {
	data := BatchOfferPromotionPackageModificationActionRes{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("task-count") {
		data.TaskCount = CastBatchOfferPromotionPackageModificationActionResTaskCountFromCli(c)
	}
	return data
}

// The base class definition for batchOfferPromotionPackageModificationActionRes
type BatchOfferPromotionPackageModificationActionRes struct {
	Id        string                                                   `json:"id" yaml:"id"`
	TaskCount BatchOfferPromotionPackageModificationActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}

func GetBatchOfferPromotionPackageModificationActionResTaskCountCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "failed",
			Type: "int",
		},
		{
			Name: prefix + "success",
			Type: "int",
		},
		{
			Name: prefix + "total",
			Type: "int",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionResTaskCountFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionResTaskCount {
	data := BatchOfferPromotionPackageModificationActionResTaskCount{}
	if c.IsSet("failed") {
		data.Failed = int(c.Int64("failed"))
	}
	if c.IsSet("success") {
		data.Success = int(c.Int64("success"))
	}
	if c.IsSet("total") {
		data.Total = int(c.Int64("total"))
	}
	return data
}

// The base class definition for taskCount
type BatchOfferPromotionPackageModificationActionResTaskCount struct {
	Failed  int `json:"failed" yaml:"failed"`
	Success int `json:"success" yaml:"success"`
	Total   int `json:"total" yaml:"total"`
}

func (x *BatchOfferPromotionPackageModificationActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type BatchOfferPromotionPackageModificationActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *BatchOfferPromotionPackageModificationActionResponse) SetContentType(contentType string) *BatchOfferPromotionPackageModificationActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *BatchOfferPromotionPackageModificationActionResponse) AsStream(r io.Reader, contentType string) *BatchOfferPromotionPackageModificationActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *BatchOfferPromotionPackageModificationActionResponse) AsJSON(payload any) *BatchOfferPromotionPackageModificationActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *BatchOfferPromotionPackageModificationActionResponse) WithIdeal(payload BatchOfferPromotionPackageModificationActionRes) *BatchOfferPromotionPackageModificationActionResponse {
	x.Payload = payload
	return x
}
func (x *BatchOfferPromotionPackageModificationActionResponse) AsHTML(payload string) *BatchOfferPromotionPackageModificationActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *BatchOfferPromotionPackageModificationActionResponse) AsBytes(payload []byte) *BatchOfferPromotionPackageModificationActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x BatchOfferPromotionPackageModificationActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x BatchOfferPromotionPackageModificationActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x BatchOfferPromotionPackageModificationActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type BatchOfferPromotionPackageModificationActionRequestSig = func(c BatchOfferPromotionPackageModificationActionRequest) (*BatchOfferPromotionPackageModificationActionResponse, error)

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
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	// Now it's interface, so the code gen doesn't depend on the instance
	// or gin package. Make sure you cast is later into *gin.Context, or whatever
	// your framework is passing when creating a request.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	GinCtx interface{}
	// Cli library helper (urfave) by default. The instance is interface{}, and you
	// need to manually cast it to the *cli.Command, so gives you freedom and independence
	// of external library.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	CliCtx interface{}
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func BatchOfferPromotionPackageModificationActionClientCreateUrl(
	req BatchOfferPromotionPackageModificationActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := BatchOfferPromotionPackageModificationActionMeta()
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
func BatchOfferPromotionPackageModificationActionClientExecuteTyped(httpReq *http.Request) (*BatchOfferPromotionPackageModificationActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result BatchOfferPromotionPackageModificationActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &BatchOfferPromotionPackageModificationActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &BatchOfferPromotionPackageModificationActionResponse{Payload: result}, err
	}
	return &BatchOfferPromotionPackageModificationActionResponse{Payload: result}, nil
}
func BatchOfferPromotionPackageModificationActionClientBuildRequest(req BatchOfferPromotionPackageModificationActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := BatchOfferPromotionPackageModificationActionMeta()
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
func BatchOfferPromotionPackageModificationActionCall(
	req BatchOfferPromotionPackageModificationActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*BatchOfferPromotionPackageModificationActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := BatchOfferPromotionPackageModificationActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := BatchOfferPromotionPackageModificationActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return BatchOfferPromotionPackageModificationActionClientExecuteTyped(r)
}
func (x BatchOfferPromotionPackageModificationActionRequest) IsCli() bool {
	if x.CliCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.CliCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}

// BatchOfferPromotionPackageModificationActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the BatchOfferPromotionPackageModificationAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *BatchOfferPromotionPackageModificationActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func BatchOfferPromotionPackageModificationActionHttpHandler(
	handler func(c BatchOfferPromotionPackageModificationActionRequest) (*BatchOfferPromotionPackageModificationActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := BatchOfferPromotionPackageModificationActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body BatchOfferPromotionPackageModificationActionReq
		if r.Body != nil {
			defer r.Body.Close()
			if data, _ := io.ReadAll(r.Body); len(data) > 0 {
				if err := json.Unmarshal(data, &body); err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON: " + err.Error()})
					return
				}
			}
		}
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := BatchOfferPromotionPackageModificationActionRequest{
			Body:        body,
			QueryParams: r.URL.Query(),
			Headers:     r.Header,
		}
		resp, err := handler(req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		// If the handler returned nil (and no error), the response was handled
		// manually.
		if resp == nil {
			return
		}
		// Apply headers
		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}
		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}
		if resp.Payload != nil {
			if w.Header().Get("Content-Type") == "" {
				w.Header().Set("Content-Type", "application/json")
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(resp.Payload)
		} else {
			w.WriteHeader(status)
		}
	}
}

// BatchOfferPromotionPackageModificationActionHttp is a high-level convenience wrapper around
// BatchOfferPromotionPackageModificationActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func BatchOfferPromotionPackageModificationActionHttp(
	mux *http.ServeMux,
	handler func(c BatchOfferPromotionPackageModificationActionRequest) (*BatchOfferPromotionPackageModificationActionResponse, error),
) {
	method, pattern, h := BatchOfferPromotionPackageModificationActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
