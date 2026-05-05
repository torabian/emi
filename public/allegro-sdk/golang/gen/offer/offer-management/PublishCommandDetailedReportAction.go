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
* Action to communicate with the action PublishCommandDetailedReportAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PublishCommandDetailedReportAction
func PublishCommandDetailedReportAction(c PublishCommandDetailedReportActionRequest) (*PublishCommandDetailedReportActionResponse, error) {
	return &PublishCommandDetailedReportActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PublishCommandDetailedReportActionMeta() struct {
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
		Name:        "PublishCommandDetailedReportAction",
		CliName:     "publish command detailed report-action",
		URL:         "https://api.{environment}/sale/offer-publication-commands/{commandId}/tasks",
		Method:      "GET",
		Description: `Use this resource to retrieve information about the offer statuses on the site (Defaults: limit = 100, offset = 0). Read more: PL / EN. This resource is rate limited to retrieving information about 270 000 offer changes per minute.`,
	}
}
func GetPublishCommandDetailedReportActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "tasks",
			Type: "array",
		},
	}
}
func CastPublishCommandDetailedReportActionResFromCli(c emigo.CliCastable) PublishCommandDetailedReportActionRes {
	data := PublishCommandDetailedReportActionRes{}
	if c.IsSet("tasks") {
		data.Tasks = emigo.CapturePossibleArray(CastPublishCommandDetailedReportActionResTasksFromCli, "tasks", c)
	}
	return data
}

// The base class definition for publishCommandDetailedReportActionRes
type PublishCommandDetailedReportActionRes struct {
	Tasks []PublishCommandDetailedReportActionResTasks `json:"tasks" yaml:"tasks"`
}

func GetPublishCommandDetailedReportActionResTasksCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "field",
			Type: "string",
		},
		{
			Name: prefix + "message",
			Type: "string",
		},
		{
			Name:     prefix + "offer",
			Type:     "object",
			Children: GetPublishCommandDetailedReportActionResTasksOfferCliFlags("offer-"),
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
func CastPublishCommandDetailedReportActionResTasksFromCli(c emigo.CliCastable) PublishCommandDetailedReportActionResTasks {
	data := PublishCommandDetailedReportActionResTasks{}
	if c.IsSet("field") {
		data.Field = c.String("field")
	}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("offer") {
		data.Offer = CastPublishCommandDetailedReportActionResTasksOfferFromCli(c)
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("errors") {
		data.Errors = emigo.CapturePossibleArray(CastPublishCommandDetailedReportActionResTasksErrorsFromCli, "errors", c)
	}
	return data
}

// The base class definition for tasks
type PublishCommandDetailedReportActionResTasks struct {
	Field   string                                             `json:"field" yaml:"field"`
	Message string                                             `json:"message" yaml:"message"`
	Offer   PublishCommandDetailedReportActionResTasksOffer    `json:"offer" yaml:"offer"`
	Status  string                                             `json:"status" yaml:"status"`
	Errors  []PublishCommandDetailedReportActionResTasksErrors `json:"errors" yaml:"errors"`
}

func GetPublishCommandDetailedReportActionResTasksOfferCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastPublishCommandDetailedReportActionResTasksOfferFromCli(c emigo.CliCastable) PublishCommandDetailedReportActionResTasksOffer {
	data := PublishCommandDetailedReportActionResTasksOffer{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for offer
type PublishCommandDetailedReportActionResTasksOffer struct {
	Id string `json:"id" yaml:"id"`
}

func GetPublishCommandDetailedReportActionResTasksErrorsCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetPublishCommandDetailedReportActionResTasksErrorsMetadataCliFlags("metadata-"),
		},
	}
}
func CastPublishCommandDetailedReportActionResTasksErrorsFromCli(c emigo.CliCastable) PublishCommandDetailedReportActionResTasksErrors {
	data := PublishCommandDetailedReportActionResTasksErrors{}
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
		data.Metadata = CastPublishCommandDetailedReportActionResTasksErrorsMetadataFromCli(c)
	}
	return data
}

// The base class definition for errors
type PublishCommandDetailedReportActionResTasksErrors struct {
	Code        string                                                   `json:"code" yaml:"code"`
	Details     string                                                   `json:"details" yaml:"details"`
	Message     string                                                   `json:"message" yaml:"message"`
	Path        string                                                   `json:"path" yaml:"path"`
	UserMessage string                                                   `json:"userMessage" yaml:"userMessage"`
	Metadata    PublishCommandDetailedReportActionResTasksErrorsMetadata `json:"metadata" yaml:"metadata"`
}

func GetPublishCommandDetailedReportActionResTasksErrorsMetadataCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "product-id",
			Type: "string",
		},
	}
}
func CastPublishCommandDetailedReportActionResTasksErrorsMetadataFromCli(c emigo.CliCastable) PublishCommandDetailedReportActionResTasksErrorsMetadata {
	data := PublishCommandDetailedReportActionResTasksErrorsMetadata{}
	if c.IsSet("product-id") {
		data.ProductId = c.String("product-id")
	}
	return data
}

// The base class definition for metadata
type PublishCommandDetailedReportActionResTasksErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

func (x *PublishCommandDetailedReportActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type PublishCommandDetailedReportActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PublishCommandDetailedReportActionResponse) SetContentType(contentType string) *PublishCommandDetailedReportActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PublishCommandDetailedReportActionResponse) AsStream(r io.Reader, contentType string) *PublishCommandDetailedReportActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PublishCommandDetailedReportActionResponse) AsJSON(payload any) *PublishCommandDetailedReportActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *PublishCommandDetailedReportActionResponse) WithIdeal(payload PublishCommandDetailedReportActionRes) *PublishCommandDetailedReportActionResponse {
	x.Payload = payload
	return x
}
func (x *PublishCommandDetailedReportActionResponse) AsHTML(payload string) *PublishCommandDetailedReportActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PublishCommandDetailedReportActionResponse) AsBytes(payload []byte) *PublishCommandDetailedReportActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PublishCommandDetailedReportActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PublishCommandDetailedReportActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PublishCommandDetailedReportActionResponse) GetPayload() interface{} {
	return x.Payload
}

// PublishCommandDetailedReportActionRaw registers a raw Gin route for the PublishCommandDetailedReportAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PublishCommandDetailedReportActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PublishCommandDetailedReportActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type PublishCommandDetailedReportActionRequestSig = func(c PublishCommandDetailedReportActionRequest) (*PublishCommandDetailedReportActionResponse, error)

// PublishCommandDetailedReportActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PublishCommandDetailedReportAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PublishCommandDetailedReportActionHandler(
	handler PublishCommandDetailedReportActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := PublishCommandDetailedReportActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PublishCommandDetailedReportActionRequest{
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

// PublishCommandDetailedReportAction is a high-level convenience wrapper around PublishCommandDetailedReportActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PublishCommandDetailedReportActionGin(r gin.IRoutes, handler PublishCommandDetailedReportActionRequestSig) {
	method, url, h := PublishCommandDetailedReportActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Publish command detailed reportAction
 */
// Query wrapper with private fields
type PublishCommandDetailedReportActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func PublishCommandDetailedReportActionQueryFromString(rawQuery string) PublishCommandDetailedReportActionQuery {
	v := PublishCommandDetailedReportActionQuery{}
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
func PublishCommandDetailedReportActionQueryFromGin(c *gin.Context) PublishCommandDetailedReportActionQuery {
	return PublishCommandDetailedReportActionQueryFromString(c.Request.URL.RawQuery)
}
func PublishCommandDetailedReportActionQueryFromHttp(r *http.Request) PublishCommandDetailedReportActionQuery {
	return PublishCommandDetailedReportActionQueryFromString(r.URL.RawQuery)
}
func (q PublishCommandDetailedReportActionQuery) Values() url.Values {
	return q.values
}
func (q PublishCommandDetailedReportActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PublishCommandDetailedReportActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PublishCommandDetailedReportActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PublishCommandDetailedReportActionRequest struct {
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

func (x PublishCommandDetailedReportActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x PublishCommandDetailedReportActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type PublishCommandDetailedReportActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func PublishCommandDetailedReportActionClientCreateUrl(
	req PublishCommandDetailedReportActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PublishCommandDetailedReportActionMeta()
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
func PublishCommandDetailedReportActionClientExecuteTyped(httpReq *http.Request) (*PublishCommandDetailedReportActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PublishCommandDetailedReportActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &PublishCommandDetailedReportActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &PublishCommandDetailedReportActionResponse{Payload: result}, err
	}
	return &PublishCommandDetailedReportActionResponse{Payload: result}, nil
}
func PublishCommandDetailedReportActionClientBuildRequest(req PublishCommandDetailedReportActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PublishCommandDetailedReportActionMeta()
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
func PublishCommandDetailedReportActionCall(
	req PublishCommandDetailedReportActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PublishCommandDetailedReportActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PublishCommandDetailedReportActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PublishCommandDetailedReportActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PublishCommandDetailedReportActionClientExecuteTyped(r)
}
