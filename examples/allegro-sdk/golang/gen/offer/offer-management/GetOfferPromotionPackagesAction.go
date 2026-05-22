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
* Action to communicate with the action GetOfferPromotionPackagesAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetOfferPromotionPackagesAction
func GetOfferPromotionPackagesAction(c GetOfferPromotionPackagesActionRequest) (*GetOfferPromotionPackagesActionResponse, error) {
	return &GetOfferPromotionPackagesActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetOfferPromotionPackagesActionMeta() struct {
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
		Name:        "GetOfferPromotionPackagesAction",
		CliName:     "get offer promotion packages-action",
		URL:         "https://api.{environment}/sale/offers/{offerId}/promo-options",
		Method:      "GET",
		Description: ``,
	}
}
func GetGetOfferPromotionPackagesActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "offer-id",
			Type: "string",
		},
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetGetOfferPromotionPackagesActionResBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name:     prefix + "pending-changes",
			Type:     "object",
			Children: GetGetOfferPromotionPackagesActionResPendingChangesCliFlags("pending-changes-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "array",
		},
	}
}
func CastGetOfferPromotionPackagesActionResFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionRes {
	data := GetOfferPromotionPackagesActionRes{}
	if c.IsSet("offer-id") {
		data.OfferId = c.String("offer-id")
	}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("base-package") {
		data.BasePackage = CastGetOfferPromotionPackagesActionResBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastGetOfferPromotionPackagesActionResExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("pending-changes") {
		data.PendingChanges = CastGetOfferPromotionPackagesActionResPendingChangesFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = emigo.CapturePossibleArray(CastGetOfferPromotionPackagesActionResAdditionalMarketplacesFromCli, "additional-marketplaces", c)
	}
	return data
}

// The base class definition for getOfferPromotionPackagesActionRes
type GetOfferPromotionPackagesActionRes struct {
	OfferId                string                                                     `json:"offerId" yaml:"offerId"`
	MarketplaceId          string                                                     `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage            GetOfferPromotionPackagesActionResBasePackage              `json:"basePackage" yaml:"basePackage"`
	ExtraPackages          []GetOfferPromotionPackagesActionResExtraPackages          `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges         GetOfferPromotionPackagesActionResPendingChanges           `json:"pendingChanges" yaml:"pendingChanges"`
	AdditionalMarketplaces []GetOfferPromotionPackagesActionResAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

func GetGetOfferPromotionPackagesActionResBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetOfferPromotionPackagesActionResBasePackageFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionResBasePackage {
	data := GetOfferPromotionPackagesActionResBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for basePackage
type GetOfferPromotionPackagesActionResBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetOfferPromotionPackagesActionResExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetOfferPromotionPackagesActionResExtraPackagesFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionResExtraPackages {
	data := GetOfferPromotionPackagesActionResExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for extraPackages
type GetOfferPromotionPackagesActionResExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetOfferPromotionPackagesActionResPendingChangesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetGetOfferPromotionPackagesActionResPendingChangesBasePackageCliFlags("base-package-"),
		},
	}
}
func CastGetOfferPromotionPackagesActionResPendingChangesFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionResPendingChanges {
	data := GetOfferPromotionPackagesActionResPendingChanges{}
	if c.IsSet("base-package") {
		data.BasePackage = CastGetOfferPromotionPackagesActionResPendingChangesBasePackageFromCli(c)
	}
	return data
}

// The base class definition for pendingChanges
type GetOfferPromotionPackagesActionResPendingChanges struct {
	BasePackage GetOfferPromotionPackagesActionResPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

func GetGetOfferPromotionPackagesActionResPendingChangesBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetOfferPromotionPackagesActionResPendingChangesBasePackageFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionResPendingChangesBasePackage {
	data := GetOfferPromotionPackagesActionResPendingChangesBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for basePackage
type GetOfferPromotionPackagesActionResPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetOfferPromotionPackagesActionResAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetGetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name:     prefix + "pending-changes",
			Type:     "object",
			Children: GetGetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesCliFlags("pending-changes-"),
		},
	}
}
func CastGetOfferPromotionPackagesActionResAdditionalMarketplacesFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionResAdditionalMarketplaces {
	data := GetOfferPromotionPackagesActionResAdditionalMarketplaces{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("base-package") {
		data.BasePackage = CastGetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastGetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("pending-changes") {
		data.PendingChanges = CastGetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesFromCli(c)
	}
	return data
}

// The base class definition for additionalMarketplaces
type GetOfferPromotionPackagesActionResAdditionalMarketplaces struct {
	MarketplaceId  string                                                                  `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage    GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages  []GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges  `json:"pendingChanges" yaml:"pendingChanges"`
}

func GetGetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackageFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage {
	data := GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for basePackage
type GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackagesFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages {
	data := GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for extraPackages
type GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetGetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackageCliFlags("base-package-"),
		},
	}
}
func CastGetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges {
	data := GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges{}
	if c.IsSet("base-package") {
		data.BasePackage = CastGetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackageFromCli(c)
	}
	return data
}

// The base class definition for pendingChanges
type GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges struct {
	BasePackage GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

func GetGetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackageFromCli(c emigo.CliCastable) GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage {
	data := GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for basePackage
type GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func (x *GetOfferPromotionPackagesActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetOfferPromotionPackagesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetOfferPromotionPackagesActionResponse) SetContentType(contentType string) *GetOfferPromotionPackagesActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetOfferPromotionPackagesActionResponse) AsStream(r io.Reader, contentType string) *GetOfferPromotionPackagesActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetOfferPromotionPackagesActionResponse) AsJSON(payload any) *GetOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetOfferPromotionPackagesActionResponse) WithIdeal(payload GetOfferPromotionPackagesActionRes) *GetOfferPromotionPackagesActionResponse {
	x.Payload = payload
	return x
}
func (x *GetOfferPromotionPackagesActionResponse) AsHTML(payload string) *GetOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetOfferPromotionPackagesActionResponse) AsBytes(payload []byte) *GetOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetOfferPromotionPackagesActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetOfferPromotionPackagesActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetOfferPromotionPackagesActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetOfferPromotionPackagesActionRaw registers a raw Gin route for the GetOfferPromotionPackagesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetOfferPromotionPackagesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetOfferPromotionPackagesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetOfferPromotionPackagesActionRequestSig = func(c GetOfferPromotionPackagesActionRequest) (*GetOfferPromotionPackagesActionResponse, error)

// GetOfferPromotionPackagesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetOfferPromotionPackagesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetOfferPromotionPackagesActionHandler(
	handler GetOfferPromotionPackagesActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetOfferPromotionPackagesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetOfferPromotionPackagesActionRequest{
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

// GetOfferPromotionPackagesAction is a high-level convenience wrapper around GetOfferPromotionPackagesActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetOfferPromotionPackagesActionGin(r gin.IRoutes, handler GetOfferPromotionPackagesActionRequestSig) {
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

func (x GetOfferPromotionPackagesActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GetOfferPromotionPackagesActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GetOfferPromotionPackagesActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GetOfferPromotionPackagesActionClientCreateUrl(
	req GetOfferPromotionPackagesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetOfferPromotionPackagesActionMeta()
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
func GetOfferPromotionPackagesActionClientExecuteTyped(httpReq *http.Request) (*GetOfferPromotionPackagesActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetOfferPromotionPackagesActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetOfferPromotionPackagesActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetOfferPromotionPackagesActionResponse{Payload: result}, err
	}
	return &GetOfferPromotionPackagesActionResponse{Payload: result}, nil
}
func GetOfferPromotionPackagesActionClientBuildRequest(req GetOfferPromotionPackagesActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetOfferPromotionPackagesActionMeta()
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
func GetOfferPromotionPackagesActionCall(
	req GetOfferPromotionPackagesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetOfferPromotionPackagesActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetOfferPromotionPackagesActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetOfferPromotionPackagesActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetOfferPromotionPackagesActionClientExecuteTyped(r)
}
