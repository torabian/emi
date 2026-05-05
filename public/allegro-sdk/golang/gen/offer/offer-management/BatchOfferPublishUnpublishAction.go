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
* Action to communicate with the action BatchOfferPublishUnpublishAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of BatchOfferPublishUnpublishAction
func BatchOfferPublishUnpublishAction(c BatchOfferPublishUnpublishActionRequest) (*BatchOfferPublishUnpublishActionResponse, error) {
	return &BatchOfferPublishUnpublishActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func BatchOfferPublishUnpublishActionMeta() struct {
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
		Name:        "BatchOfferPublishUnpublishAction",
		CliName:     "batch offer publish / unpublish-action",
		URL:         "https://api.{environment}/sale/offer-publication-commands/{commandId}",
		Method:      "PUT",
		Description: `Use this resource to modify multiple offers publication at once. Read more: PL / EN. This resource is rate limited to 250 000 offer changes per hour or 9000 offer changes per minute.`,
	}
}
func GetBatchOfferPublishUnpublishActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "offer-criteria",
			Type: "array",
		},
		{
			Name:     prefix + "publication",
			Type:     "object",
			Children: GetBatchOfferPublishUnpublishActionReqPublicationCliFlags("publication-"),
		},
	}
}
func CastBatchOfferPublishUnpublishActionReqFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionReq {
	data := BatchOfferPublishUnpublishActionReq{}
	if c.IsSet("offer-criteria") {
		data.OfferCriteria = emigo.CapturePossibleArray(CastBatchOfferPublishUnpublishActionReqOfferCriteriaFromCli, "offer-criteria", c)
	}
	if c.IsSet("publication") {
		data.Publication = CastBatchOfferPublishUnpublishActionReqPublicationFromCli(c)
	}
	return data
}

// The base class definition for batchOfferPublishUnpublishActionReq
type BatchOfferPublishUnpublishActionReq struct {
	OfferCriteria []BatchOfferPublishUnpublishActionReqOfferCriteria `json:"offerCriteria" yaml:"offerCriteria"`
	Publication   BatchOfferPublishUnpublishActionReqPublication     `json:"publication" yaml:"publication"`
}

func GetBatchOfferPublishUnpublishActionReqOfferCriteriaCliFlags(prefix string) []emigo.CliFlag {
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
func CastBatchOfferPublishUnpublishActionReqOfferCriteriaFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionReqOfferCriteria {
	data := BatchOfferPublishUnpublishActionReqOfferCriteria{}
	if c.IsSet("offers") {
		data.Offers = emigo.CapturePossibleArray(CastBatchOfferPublishUnpublishActionReqOfferCriteriaOffersFromCli, "offers", c)
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for offerCriteria
type BatchOfferPublishUnpublishActionReqOfferCriteria struct {
	Offers []BatchOfferPublishUnpublishActionReqOfferCriteriaOffers `json:"offers" yaml:"offers"`
	Type   string                                                   `json:"type" yaml:"type"`
}

func GetBatchOfferPublishUnpublishActionReqOfferCriteriaOffersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPublishUnpublishActionReqOfferCriteriaOffersFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionReqOfferCriteriaOffers {
	data := BatchOfferPublishUnpublishActionReqOfferCriteriaOffers{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for offers
type BatchOfferPublishUnpublishActionReqOfferCriteriaOffers struct {
	Id string `json:"id" yaml:"id"`
}

func GetBatchOfferPublishUnpublishActionReqPublicationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "action",
			Type: "string",
		},
		{
			Name: prefix + "scheduled-for",
			Type: "string",
		},
	}
}
func CastBatchOfferPublishUnpublishActionReqPublicationFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionReqPublication {
	data := BatchOfferPublishUnpublishActionReqPublication{}
	if c.IsSet("action") {
		data.Action = c.String("action")
	}
	if c.IsSet("scheduled-for") {
		data.ScheduledFor = c.String("scheduled-for")
	}
	return data
}

// The base class definition for publication
type BatchOfferPublishUnpublishActionReqPublication struct {
	Action       string `json:"action" yaml:"action"`
	ScheduledFor string `json:"scheduledFor" yaml:"scheduledFor"`
}

func (x *BatchOfferPublishUnpublishActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetBatchOfferPublishUnpublishActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "created-at",
			Type: "string",
		},
		{
			Name: prefix + "completed-at",
			Type: "string",
		},
		{
			Name:     prefix + "task-count",
			Type:     "object",
			Children: GetBatchOfferPublishUnpublishActionResTaskCountCliFlags("task-count-"),
		},
	}
}
func CastBatchOfferPublishUnpublishActionResFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionRes {
	data := BatchOfferPublishUnpublishActionRes{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("created-at") {
		data.CreatedAt = c.String("created-at")
	}
	if c.IsSet("completed-at") {
		data.CompletedAt = c.String("completed-at")
	}
	if c.IsSet("task-count") {
		data.TaskCount = CastBatchOfferPublishUnpublishActionResTaskCountFromCli(c)
	}
	return data
}

// The base class definition for batchOfferPublishUnpublishActionRes
type BatchOfferPublishUnpublishActionRes struct {
	Id          string                                       `json:"id" yaml:"id"`
	CreatedAt   string                                       `json:"createdAt" yaml:"createdAt"`
	CompletedAt string                                       `json:"completedAt" yaml:"completedAt"`
	TaskCount   BatchOfferPublishUnpublishActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}

func GetBatchOfferPublishUnpublishActionResTaskCountCliFlags(prefix string) []emigo.CliFlag {
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
func CastBatchOfferPublishUnpublishActionResTaskCountFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionResTaskCount {
	data := BatchOfferPublishUnpublishActionResTaskCount{}
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
type BatchOfferPublishUnpublishActionResTaskCount struct {
	Failed  int `json:"failed" yaml:"failed"`
	Success int `json:"success" yaml:"success"`
	Total   int `json:"total" yaml:"total"`
}

func (x *BatchOfferPublishUnpublishActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type BatchOfferPublishUnpublishActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *BatchOfferPublishUnpublishActionResponse) SetContentType(contentType string) *BatchOfferPublishUnpublishActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *BatchOfferPublishUnpublishActionResponse) AsStream(r io.Reader, contentType string) *BatchOfferPublishUnpublishActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *BatchOfferPublishUnpublishActionResponse) AsJSON(payload any) *BatchOfferPublishUnpublishActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *BatchOfferPublishUnpublishActionResponse) WithIdeal(payload BatchOfferPublishUnpublishActionRes) *BatchOfferPublishUnpublishActionResponse {
	x.Payload = payload
	return x
}
func (x *BatchOfferPublishUnpublishActionResponse) AsHTML(payload string) *BatchOfferPublishUnpublishActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *BatchOfferPublishUnpublishActionResponse) AsBytes(payload []byte) *BatchOfferPublishUnpublishActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x BatchOfferPublishUnpublishActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x BatchOfferPublishUnpublishActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x BatchOfferPublishUnpublishActionResponse) GetPayload() interface{} {
	return x.Payload
}

// BatchOfferPublishUnpublishActionRaw registers a raw Gin route for the BatchOfferPublishUnpublishAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func BatchOfferPublishUnpublishActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := BatchOfferPublishUnpublishActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type BatchOfferPublishUnpublishActionRequestSig = func(c BatchOfferPublishUnpublishActionRequest) (*BatchOfferPublishUnpublishActionResponse, error)

// BatchOfferPublishUnpublishActionHandler returns the HTTP method, route URL, and a typed Gin handler for the BatchOfferPublishUnpublishAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func BatchOfferPublishUnpublishActionHandler(
	handler BatchOfferPublishUnpublishActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := BatchOfferPublishUnpublishActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body BatchOfferPublishUnpublishActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := BatchOfferPublishUnpublishActionRequest{
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

// BatchOfferPublishUnpublishAction is a high-level convenience wrapper around BatchOfferPublishUnpublishActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func BatchOfferPublishUnpublishActionGin(r gin.IRoutes, handler BatchOfferPublishUnpublishActionRequestSig) {
	method, url, h := BatchOfferPublishUnpublishActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Batch offer publish / unpublishAction
 */
// Query wrapper with private fields
type BatchOfferPublishUnpublishActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func BatchOfferPublishUnpublishActionQueryFromString(rawQuery string) BatchOfferPublishUnpublishActionQuery {
	v := BatchOfferPublishUnpublishActionQuery{}
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
func BatchOfferPublishUnpublishActionQueryFromGin(c *gin.Context) BatchOfferPublishUnpublishActionQuery {
	return BatchOfferPublishUnpublishActionQueryFromString(c.Request.URL.RawQuery)
}
func BatchOfferPublishUnpublishActionQueryFromHttp(r *http.Request) BatchOfferPublishUnpublishActionQuery {
	return BatchOfferPublishUnpublishActionQueryFromString(r.URL.RawQuery)
}
func (q BatchOfferPublishUnpublishActionQuery) Values() url.Values {
	return q.values
}
func (q BatchOfferPublishUnpublishActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *BatchOfferPublishUnpublishActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *BatchOfferPublishUnpublishActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type BatchOfferPublishUnpublishActionRequest struct {
	Body        BatchOfferPublishUnpublishActionReq
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

func (x BatchOfferPublishUnpublishActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x BatchOfferPublishUnpublishActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type BatchOfferPublishUnpublishActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func BatchOfferPublishUnpublishActionClientCreateUrl(
	req BatchOfferPublishUnpublishActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := BatchOfferPublishUnpublishActionMeta()
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
func BatchOfferPublishUnpublishActionClientExecuteTyped(httpReq *http.Request) (*BatchOfferPublishUnpublishActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result BatchOfferPublishUnpublishActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &BatchOfferPublishUnpublishActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &BatchOfferPublishUnpublishActionResponse{Payload: result}, err
	}
	return &BatchOfferPublishUnpublishActionResponse{Payload: result}, nil
}
func BatchOfferPublishUnpublishActionClientBuildRequest(req BatchOfferPublishUnpublishActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := BatchOfferPublishUnpublishActionMeta()
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
func BatchOfferPublishUnpublishActionCall(
	req BatchOfferPublishUnpublishActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*BatchOfferPublishUnpublishActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := BatchOfferPublishUnpublishActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := BatchOfferPublishUnpublishActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return BatchOfferPublishUnpublishActionClientExecuteTyped(r)
}
