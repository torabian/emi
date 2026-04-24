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
* Action to communicate with the action ModificationCommandSummaryAction
 */
func ModificationCommandSummaryActionMeta() struct {
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
		Name:        "ModificationCommandSummaryAction",
		CliName:     "modification command summary-action",
		URL:         "https://api.{environment}/sale/offers/promo-options-commands/{commandId}",
		Method:      "GET",
		Description: `Use this resource to find out how many offers were edited within one {commandId}. You will receive a summary with a number of successfully edited offers and errors. Read more: PL / EN.`,
	}
}
func GetModificationCommandSummaryActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name:     prefix + "task-count",
			Type:     "object",
			Children: GetModificationCommandSummaryActionResTaskCountCliFlags("task-count-"),
		},
	}
}
func CastModificationCommandSummaryActionResFromCli(c emigo.CliCastable) ModificationCommandSummaryActionRes {
	data := ModificationCommandSummaryActionRes{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("task-count") {
		data.TaskCount = CastModificationCommandSummaryActionResTaskCountFromCli(c)
	}
	return data
}

// The base class definition for modificationCommandSummaryActionRes
type ModificationCommandSummaryActionRes struct {
	Id        string                                       `json:"id" yaml:"id"`
	TaskCount ModificationCommandSummaryActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}

func GetModificationCommandSummaryActionResTaskCountCliFlags(prefix string) []emigo.CliFlag {
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
func CastModificationCommandSummaryActionResTaskCountFromCli(c emigo.CliCastable) ModificationCommandSummaryActionResTaskCount {
	data := ModificationCommandSummaryActionResTaskCount{}
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
type ModificationCommandSummaryActionResTaskCount struct {
	Failed  int `json:"failed" yaml:"failed"`
	Success int `json:"success" yaml:"success"`
	Total   int `json:"total" yaml:"total"`
}

func (x *ModificationCommandSummaryActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ModificationCommandSummaryActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *ModificationCommandSummaryActionResponse) SetContentType(contentType string) *ModificationCommandSummaryActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ModificationCommandSummaryActionResponse) AsStream(r io.Reader, contentType string) *ModificationCommandSummaryActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ModificationCommandSummaryActionResponse) AsJSON(payload any) *ModificationCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *ModificationCommandSummaryActionResponse) AsHTML(payload string) *ModificationCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ModificationCommandSummaryActionResponse) AsBytes(payload []byte) *ModificationCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ModificationCommandSummaryActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ModificationCommandSummaryActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ModificationCommandSummaryActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ModificationCommandSummaryActionRaw registers a raw Gin route for the ModificationCommandSummaryAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModificationCommandSummaryActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModificationCommandSummaryActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ModificationCommandSummaryActionRequestSig = func(c ModificationCommandSummaryActionRequest) (*ModificationCommandSummaryActionResponse, error)

// ModificationCommandSummaryActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModificationCommandSummaryAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModificationCommandSummaryActionHandler(
	handler ModificationCommandSummaryActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ModificationCommandSummaryActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := ModificationCommandSummaryActionRequest{
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

// ModificationCommandSummaryAction is a high-level convenience wrapper around ModificationCommandSummaryActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModificationCommandSummaryActionGin(r gin.IRoutes, handler ModificationCommandSummaryActionRequestSig) {
	method, url, h := ModificationCommandSummaryActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Modification command summaryAction
 */
// Query wrapper with private fields
type ModificationCommandSummaryActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModificationCommandSummaryActionQueryFromString(rawQuery string) ModificationCommandSummaryActionQuery {
	v := ModificationCommandSummaryActionQuery{}
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
func ModificationCommandSummaryActionQueryFromGin(c *gin.Context) ModificationCommandSummaryActionQuery {
	return ModificationCommandSummaryActionQueryFromString(c.Request.URL.RawQuery)
}
func ModificationCommandSummaryActionQueryFromHttp(r *http.Request) ModificationCommandSummaryActionQuery {
	return ModificationCommandSummaryActionQueryFromString(r.URL.RawQuery)
}
func (q ModificationCommandSummaryActionQuery) Values() url.Values {
	return q.values
}
func (q ModificationCommandSummaryActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModificationCommandSummaryActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModificationCommandSummaryActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModificationCommandSummaryActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type ModificationCommandSummaryActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ModificationCommandSummaryActionCall(
	req ModificationCommandSummaryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModificationCommandSummaryActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ModificationCommandSummaryActionMeta()
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
	var result ModificationCommandSummaryActionResult
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
