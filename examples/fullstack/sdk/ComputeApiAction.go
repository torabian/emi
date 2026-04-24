package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli"
)

/**
* Action to communicate with the action ComputeApiAction
 */
func ComputeApiActionMeta() struct {
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
		Name:        "ComputeApiAction",
		CliName:     "compute-api-action",
		URL:         "/compute/api",
		Method:      "POST",
		Description: `An API, which computes two vectors`,
	}
}
func GetComputeApiActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "initial-vector1",
			Type: "slice",
		},
		{
			Name: prefix + "value",
			Type: "string?",
		},
		{
			Name: prefix + "initial-vector2",
			Type: "slice",
		},
	}
}
func CastComputeApiActionReqFromCli(c emigo.CliCastable) ComputeApiActionReq {
	data := ComputeApiActionReq{}
	if c.IsSet("initial-vector1") {
		emigo.InflatePossibleSlice(c.String("initial-vector1"), &data.InitialVector1)
	}
	if c.IsSet("value") {
		emigo.ParseNullable(c.String("value"), &data.Value)
	}
	if c.IsSet("initial-vector2") {
		emigo.InflatePossibleSlice(c.String("initial-vector2"), &data.InitialVector2)
	}
	return data
}

// The base class definition for computeApiActionReq
type ComputeApiActionReq struct {
	InitialVector1 []int                  `json:"initialVector1" yaml:"initialVector1"`
	Value          emigo.Nullable[string] `json:"value" yaml:"value"`
	InitialVector2 []int                  `json:"initialVector2" yaml:"initialVector2"`
}

func (x *ComputeApiActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetComputeApiActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "output-vector",
			Type: "slice",
		},
	}
}
func CastComputeApiActionResFromCli(c emigo.CliCastable) ComputeApiActionRes {
	data := ComputeApiActionRes{}
	if c.IsSet("output-vector") {
		emigo.InflatePossibleSlice(c.String("output-vector"), &data.OutputVector)
	}
	return data
}

// The base class definition for computeApiActionRes
type ComputeApiActionRes struct {
	OutputVector []int `json:"outputVector" yaml:"outputVector"`
}

func (x *ComputeApiActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ComputeApiActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *ComputeApiActionResponse) SetContentType(contentType string) *ComputeApiActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ComputeApiActionResponse) AsStream(r io.Reader, contentType string) *ComputeApiActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ComputeApiActionResponse) AsJSON(payload any) *ComputeApiActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ComputeApiActionResponse) WithIdeal(payload ComputeApiActionRes) *ComputeApiActionResponse {
	x.Payload = payload
	return x
}
func (x *ComputeApiActionResponse) AsHTML(payload string) *ComputeApiActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ComputeApiActionResponse) AsBytes(payload []byte) *ComputeApiActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ComputeApiActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ComputeApiActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ComputeApiActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ComputeApiActionRaw registers a raw Gin route for the ComputeApiAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ComputeApiActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ComputeApiActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ComputeApiActionRequestSig = func(c ComputeApiActionRequest) (*ComputeApiActionResponse, error)

// ComputeApiActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ComputeApiAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ComputeApiActionHandler(
	handler ComputeApiActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ComputeApiActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ComputeApiActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ComputeApiActionRequest{
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

// ComputeApiAction is a high-level convenience wrapper around ComputeApiActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ComputeApiActionGin(r gin.IRoutes, handler ComputeApiActionRequestSig) {
	method, url, h := ComputeApiActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for ComputeApiAction
 */
// Query wrapper with private fields
type ComputeApiActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Action     string `json:"action"`
	SnapPoints []int  `json:"snapPoints"`
}

func ComputeApiActionQueryFromString(rawQuery string) ComputeApiActionQuery {
	v := ComputeApiActionQuery{}
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
func ComputeApiActionQueryFromGin(c *gin.Context) ComputeApiActionQuery {
	return ComputeApiActionQueryFromString(c.Request.URL.RawQuery)
}
func ComputeApiActionQueryFromHttp(r *http.Request) ComputeApiActionQuery {
	return ComputeApiActionQueryFromString(r.URL.RawQuery)
}
func (q ComputeApiActionQuery) Values() url.Values {
	return q.values
}
func (q ComputeApiActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ComputeApiActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ComputeApiActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ComputeApiActionRequest struct {
	Body        ComputeApiActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type ComputeApiActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ComputeApiActionCall(
	req ComputeApiActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ComputeApiActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ComputeApiActionMeta()
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
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		req0, err := http.NewRequest(meta.Method, u.String(), bytes.NewReader(bodyBytes))
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
	var result ComputeApiActionResult
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
