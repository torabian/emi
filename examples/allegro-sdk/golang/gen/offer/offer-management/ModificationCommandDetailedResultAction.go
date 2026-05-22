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
* Action to communicate with the action ModificationCommandDetailedResultAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ModificationCommandDetailedResultAction
func ModificationCommandDetailedResultAction(c ModificationCommandDetailedResultActionRequest) (*ModificationCommandDetailedResultActionResponse, error) {
	return &ModificationCommandDetailedResultActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ModificationCommandDetailedResultActionMeta() struct {
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
		Name:        "ModificationCommandDetailedResultAction",
		CliName:     "modification command detailed result-action",
		URL:         "https://api.{environment}/sale/offers/promo-options-commands/{commandId}/tasks",
		Method:      "GET",
		Description: ``,
	}
}
func GetModificationCommandDetailedResultActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "tasks",
			Type: "array",
		},
		{
			Name:     prefix + "modification",
			Type:     "object",
			Children: GetModificationCommandDetailedResultActionResModificationCliFlags("modification-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "array",
		},
	}
}
func CastModificationCommandDetailedResultActionResFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionRes {
	data := ModificationCommandDetailedResultActionRes{}
	if c.IsSet("tasks") {
		data.Tasks = emigo.CapturePossibleArray(CastModificationCommandDetailedResultActionResTasksFromCli, "tasks", c)
	}
	if c.IsSet("modification") {
		data.Modification = CastModificationCommandDetailedResultActionResModificationFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = emigo.CapturePossibleArray(CastModificationCommandDetailedResultActionResAdditionalMarketplacesFromCli, "additional-marketplaces", c)
	}
	return data
}

// The base class definition for modificationCommandDetailedResultActionRes
type ModificationCommandDetailedResultActionRes struct {
	Tasks                  []ModificationCommandDetailedResultActionResTasks                  `json:"tasks" yaml:"tasks"`
	Modification           ModificationCommandDetailedResultActionResModification             `json:"modification" yaml:"modification"`
	AdditionalMarketplaces []ModificationCommandDetailedResultActionResAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

func GetModificationCommandDetailedResultActionResTasksCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "offer",
			Type:     "object",
			Children: GetModificationCommandDetailedResultActionResTasksOfferCliFlags("offer-"),
		},
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name: prefix + "scheduled-at",
			Type: "string",
		},
		{
			Name: prefix + "finished-at",
			Type: "string",
		},
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
func CastModificationCommandDetailedResultActionResTasksFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResTasks {
	data := ModificationCommandDetailedResultActionResTasks{}
	if c.IsSet("offer") {
		data.Offer = CastModificationCommandDetailedResultActionResTasksOfferFromCli(c)
	}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("scheduled-at") {
		data.ScheduledAt = c.String("scheduled-at")
	}
	if c.IsSet("finished-at") {
		data.FinishedAt = c.String("finished-at")
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("errors") {
		data.Errors = emigo.CapturePossibleArray(CastModificationCommandDetailedResultActionResTasksErrorsFromCli, "errors", c)
	}
	return data
}

// The base class definition for tasks
type ModificationCommandDetailedResultActionResTasks struct {
	Offer         ModificationCommandDetailedResultActionResTasksOffer    `json:"offer" yaml:"offer"`
	MarketplaceId string                                                  `json:"marketplaceId" yaml:"marketplaceId"`
	ScheduledAt   string                                                  `json:"scheduledAt" yaml:"scheduledAt"`
	FinishedAt    string                                                  `json:"finishedAt" yaml:"finishedAt"`
	Status        string                                                  `json:"status" yaml:"status"`
	Errors        []ModificationCommandDetailedResultActionResTasksErrors `json:"errors" yaml:"errors"`
}

func GetModificationCommandDetailedResultActionResTasksOfferCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastModificationCommandDetailedResultActionResTasksOfferFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResTasksOffer {
	data := ModificationCommandDetailedResultActionResTasksOffer{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for offer
type ModificationCommandDetailedResultActionResTasksOffer struct {
	Id string `json:"id" yaml:"id"`
}

func GetModificationCommandDetailedResultActionResTasksErrorsCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetModificationCommandDetailedResultActionResTasksErrorsMetadataCliFlags("metadata-"),
		},
	}
}
func CastModificationCommandDetailedResultActionResTasksErrorsFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResTasksErrors {
	data := ModificationCommandDetailedResultActionResTasksErrors{}
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
		data.Metadata = CastModificationCommandDetailedResultActionResTasksErrorsMetadataFromCli(c)
	}
	return data
}

// The base class definition for errors
type ModificationCommandDetailedResultActionResTasksErrors struct {
	Code        string                                                        `json:"code" yaml:"code"`
	Details     string                                                        `json:"details" yaml:"details"`
	Message     string                                                        `json:"message" yaml:"message"`
	Path        string                                                        `json:"path" yaml:"path"`
	UserMessage string                                                        `json:"userMessage" yaml:"userMessage"`
	Metadata    ModificationCommandDetailedResultActionResTasksErrorsMetadata `json:"metadata" yaml:"metadata"`
}

func GetModificationCommandDetailedResultActionResTasksErrorsMetadataCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "product-id",
			Type: "string",
		},
	}
}
func CastModificationCommandDetailedResultActionResTasksErrorsMetadataFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResTasksErrorsMetadata {
	data := ModificationCommandDetailedResultActionResTasksErrorsMetadata{}
	if c.IsSet("product-id") {
		data.ProductId = c.String("product-id")
	}
	return data
}

// The base class definition for metadata
type ModificationCommandDetailedResultActionResTasksErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

func GetModificationCommandDetailedResultActionResModificationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetModificationCommandDetailedResultActionResModificationBasePackageCliFlags("base-package-"),
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
func CastModificationCommandDetailedResultActionResModificationFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResModification {
	data := ModificationCommandDetailedResultActionResModification{}
	if c.IsSet("base-package") {
		data.BasePackage = CastModificationCommandDetailedResultActionResModificationBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastModificationCommandDetailedResultActionResModificationExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("modification-time") {
		data.ModificationTime = c.String("modification-time")
	}
	return data
}

// The base class definition for modification
type ModificationCommandDetailedResultActionResModification struct {
	BasePackage      ModificationCommandDetailedResultActionResModificationBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages    []ModificationCommandDetailedResultActionResModificationExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	ModificationTime string                                                                `json:"modificationTime" yaml:"modificationTime"`
}

func GetModificationCommandDetailedResultActionResModificationBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastModificationCommandDetailedResultActionResModificationBasePackageFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResModificationBasePackage {
	data := ModificationCommandDetailedResultActionResModificationBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for basePackage
type ModificationCommandDetailedResultActionResModificationBasePackage struct {
	Id string `json:"id" yaml:"id"`
}

func GetModificationCommandDetailedResultActionResModificationExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastModificationCommandDetailedResultActionResModificationExtraPackagesFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResModificationExtraPackages {
	data := ModificationCommandDetailedResultActionResModificationExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for extraPackages
type ModificationCommandDetailedResultActionResModificationExtraPackages struct {
	Id string `json:"id" yaml:"id"`
}

func GetModificationCommandDetailedResultActionResAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name:     prefix + "modification",
			Type:     "object",
			Children: GetModificationCommandDetailedResultActionResAdditionalMarketplacesModificationCliFlags("modification-"),
		},
	}
}
func CastModificationCommandDetailedResultActionResAdditionalMarketplacesFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResAdditionalMarketplaces {
	data := ModificationCommandDetailedResultActionResAdditionalMarketplaces{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("modification") {
		data.Modification = CastModificationCommandDetailedResultActionResAdditionalMarketplacesModificationFromCli(c)
	}
	return data
}

// The base class definition for additionalMarketplaces
type ModificationCommandDetailedResultActionResAdditionalMarketplaces struct {
	MarketplaceId string                                                                       `json:"marketplaceId" yaml:"marketplaceId"`
	Modification  ModificationCommandDetailedResultActionResAdditionalMarketplacesModification `json:"modification" yaml:"modification"`
}

func GetModificationCommandDetailedResultActionResAdditionalMarketplacesModificationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackageCliFlags("base-package-"),
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
func CastModificationCommandDetailedResultActionResAdditionalMarketplacesModificationFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResAdditionalMarketplacesModification {
	data := ModificationCommandDetailedResultActionResAdditionalMarketplacesModification{}
	if c.IsSet("base-package") {
		data.BasePackage = CastModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("modification-time") {
		data.ModificationTime = c.String("modification-time")
	}
	return data
}

// The base class definition for modification
type ModificationCommandDetailedResultActionResAdditionalMarketplacesModification struct {
	BasePackage      ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages    []ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	ModificationTime string                                                                                      `json:"modificationTime" yaml:"modificationTime"`
}

func GetModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackageFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage {
	data := ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for basePackage
type ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage struct {
	Id string `json:"id" yaml:"id"`
}

func GetModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackagesFromCli(c emigo.CliCastable) ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages {
	data := ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for extraPackages
type ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages struct {
	Id string `json:"id" yaml:"id"`
}

func (x *ModificationCommandDetailedResultActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ModificationCommandDetailedResultActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ModificationCommandDetailedResultActionResponse) SetContentType(contentType string) *ModificationCommandDetailedResultActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ModificationCommandDetailedResultActionResponse) AsStream(r io.Reader, contentType string) *ModificationCommandDetailedResultActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ModificationCommandDetailedResultActionResponse) AsJSON(payload any) *ModificationCommandDetailedResultActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ModificationCommandDetailedResultActionResponse) WithIdeal(payload ModificationCommandDetailedResultActionRes) *ModificationCommandDetailedResultActionResponse {
	x.Payload = payload
	return x
}
func (x *ModificationCommandDetailedResultActionResponse) AsHTML(payload string) *ModificationCommandDetailedResultActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ModificationCommandDetailedResultActionResponse) AsBytes(payload []byte) *ModificationCommandDetailedResultActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ModificationCommandDetailedResultActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ModificationCommandDetailedResultActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ModificationCommandDetailedResultActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ModificationCommandDetailedResultActionRaw registers a raw Gin route for the ModificationCommandDetailedResultAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModificationCommandDetailedResultActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModificationCommandDetailedResultActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ModificationCommandDetailedResultActionRequestSig = func(c ModificationCommandDetailedResultActionRequest) (*ModificationCommandDetailedResultActionResponse, error)

// ModificationCommandDetailedResultActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModificationCommandDetailedResultAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModificationCommandDetailedResultActionHandler(
	handler ModificationCommandDetailedResultActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ModificationCommandDetailedResultActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := ModificationCommandDetailedResultActionRequest{
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

// ModificationCommandDetailedResultAction is a high-level convenience wrapper around ModificationCommandDetailedResultActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModificationCommandDetailedResultActionGin(r gin.IRoutes, handler ModificationCommandDetailedResultActionRequestSig) {
	method, url, h := ModificationCommandDetailedResultActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Modification command detailed resultAction
 */
// Query wrapper with private fields
type ModificationCommandDetailedResultActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModificationCommandDetailedResultActionQueryFromString(rawQuery string) ModificationCommandDetailedResultActionQuery {
	v := ModificationCommandDetailedResultActionQuery{}
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
func ModificationCommandDetailedResultActionQueryFromGin(c *gin.Context) ModificationCommandDetailedResultActionQuery {
	return ModificationCommandDetailedResultActionQueryFromString(c.Request.URL.RawQuery)
}
func ModificationCommandDetailedResultActionQueryFromHttp(r *http.Request) ModificationCommandDetailedResultActionQuery {
	return ModificationCommandDetailedResultActionQueryFromString(r.URL.RawQuery)
}
func (q ModificationCommandDetailedResultActionQuery) Values() url.Values {
	return q.values
}
func (q ModificationCommandDetailedResultActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModificationCommandDetailedResultActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModificationCommandDetailedResultActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModificationCommandDetailedResultActionRequest struct {
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

func (x ModificationCommandDetailedResultActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x ModificationCommandDetailedResultActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type ModificationCommandDetailedResultActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func ModificationCommandDetailedResultActionClientCreateUrl(
	req ModificationCommandDetailedResultActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ModificationCommandDetailedResultActionMeta()
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
func ModificationCommandDetailedResultActionClientExecuteTyped(httpReq *http.Request) (*ModificationCommandDetailedResultActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ModificationCommandDetailedResultActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ModificationCommandDetailedResultActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ModificationCommandDetailedResultActionResponse{Payload: result}, err
	}
	return &ModificationCommandDetailedResultActionResponse{Payload: result}, nil
}
func ModificationCommandDetailedResultActionClientBuildRequest(req ModificationCommandDetailedResultActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ModificationCommandDetailedResultActionMeta()
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
func ModificationCommandDetailedResultActionCall(
	req ModificationCommandDetailedResultActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModificationCommandDetailedResultActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ModificationCommandDetailedResultActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ModificationCommandDetailedResultActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ModificationCommandDetailedResultActionClientExecuteTyped(r)
}
