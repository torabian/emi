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
* Action to communicate with the action PublishCommandSummaryAction
 */
func PublishCommandSummaryActionMeta() struct {
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
		Name:        "PublishCommandSummaryAction",
		CliName:     "publish command summary-action",
		URL:         "https://api.{environment}/sale/offer-publication-commands/{commandId}",
		Method:      "GET",
		Description: `Use this resource to retrieve information about the offer listing statuses.  You will receive a summary with a number of correctly listed offers and errors.  Read more: PL / EN. This resource is rate limited to retrieving information about 270 000 offer changes per minute.`,
	}
}
func GetPublishCommandSummaryActionResCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetPublishCommandSummaryActionResTaskCountCliFlags("task-count-"),
		},
	}
}
func CastPublishCommandSummaryActionResFromCli(c emigo.CliCastable) PublishCommandSummaryActionRes {
	data := PublishCommandSummaryActionRes{}
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
		data.TaskCount = CastPublishCommandSummaryActionResTaskCountFromCli(c)
	}
	return data
}

// The base class definition for publishCommandSummaryActionRes
type PublishCommandSummaryActionRes struct {
	Id          string                                  `json:"id" yaml:"id"`
	CreatedAt   string                                  `json:"createdAt" yaml:"createdAt"`
	CompletedAt string                                  `json:"completedAt" yaml:"completedAt"`
	TaskCount   PublishCommandSummaryActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}

func GetPublishCommandSummaryActionResTaskCountCliFlags(prefix string) []emigo.CliFlag {
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
func CastPublishCommandSummaryActionResTaskCountFromCli(c emigo.CliCastable) PublishCommandSummaryActionResTaskCount {
	data := PublishCommandSummaryActionResTaskCount{}
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
type PublishCommandSummaryActionResTaskCount struct {
	Failed  int `json:"failed" yaml:"failed"`
	Success int `json:"success" yaml:"success"`
	Total   int `json:"total" yaml:"total"`
}

func (x *PublishCommandSummaryActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type PublishCommandSummaryActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *PublishCommandSummaryActionResponse) SetContentType(contentType string) *PublishCommandSummaryActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PublishCommandSummaryActionResponse) AsStream(r io.Reader, contentType string) *PublishCommandSummaryActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PublishCommandSummaryActionResponse) AsJSON(payload any) *PublishCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *PublishCommandSummaryActionResponse) AsHTML(payload string) *PublishCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PublishCommandSummaryActionResponse) AsBytes(payload []byte) *PublishCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PublishCommandSummaryActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PublishCommandSummaryActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PublishCommandSummaryActionResponse) GetPayload() interface{} {
	return x.Payload
}

// PublishCommandSummaryActionRaw registers a raw Gin route for the PublishCommandSummaryAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PublishCommandSummaryActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PublishCommandSummaryActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type PublishCommandSummaryActionRequestSig = func(c PublishCommandSummaryActionRequest) (*PublishCommandSummaryActionResponse, error)

// PublishCommandSummaryActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PublishCommandSummaryAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PublishCommandSummaryActionHandler(
	handler PublishCommandSummaryActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := PublishCommandSummaryActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PublishCommandSummaryActionRequest{
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

// PublishCommandSummaryAction is a high-level convenience wrapper around PublishCommandSummaryActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PublishCommandSummaryActionGin(r gin.IRoutes, handler PublishCommandSummaryActionRequestSig) {
	method, url, h := PublishCommandSummaryActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Publish command summaryAction
 */
// Query wrapper with private fields
type PublishCommandSummaryActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func PublishCommandSummaryActionQueryFromString(rawQuery string) PublishCommandSummaryActionQuery {
	v := PublishCommandSummaryActionQuery{}
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
func PublishCommandSummaryActionQueryFromGin(c *gin.Context) PublishCommandSummaryActionQuery {
	return PublishCommandSummaryActionQueryFromString(c.Request.URL.RawQuery)
}
func PublishCommandSummaryActionQueryFromHttp(r *http.Request) PublishCommandSummaryActionQuery {
	return PublishCommandSummaryActionQueryFromString(r.URL.RawQuery)
}
func (q PublishCommandSummaryActionQuery) Values() url.Values {
	return q.values
}
func (q PublishCommandSummaryActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PublishCommandSummaryActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PublishCommandSummaryActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PublishCommandSummaryActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type PublishCommandSummaryActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func PublishCommandSummaryActionCall(
	req PublishCommandSummaryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PublishCommandSummaryActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := PublishCommandSummaryActionMeta()
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
	var result PublishCommandSummaryActionResult
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
