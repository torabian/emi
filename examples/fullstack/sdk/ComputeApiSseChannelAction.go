package external

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ComputeApiSseChannelAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ComputeApiSseChannelAction
func ComputeApiSseChannelAction(c ComputeApiSseChannelActionRequest) (*ComputeApiSseChannelActionResponse, error) {
	return &ComputeApiSseChannelActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ComputeApiSseChannelActionMeta() struct {
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
		Name:        "ComputeApiSseChannelAction",
		CliName:     "compute-api-sse-channel-action",
		URL:         "/compute/sse/ch",
		Method:      "GET",
		Description: `We use a helper in order to send SSE, instead of pure code in gin.`,
	}
}
func GetComputeApiSseChannelActionReqCliFlags(prefix string) []emigo.CliFlag {
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
func CastComputeApiSseChannelActionReqFromCli(c emigo.CliCastable) ComputeApiSseChannelActionReq {
	data := ComputeApiSseChannelActionReq{}
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

// The base class definition for computeApiSseChannelActionReq
type ComputeApiSseChannelActionReq struct {
	InitialVector1 []int                  `json:"initialVector1" yaml:"initialVector1"`
	Value          emigo.Nullable[string] `json:"value" yaml:"value"`
	InitialVector2 []int                  `json:"initialVector2" yaml:"initialVector2"`
}

func (x *ComputeApiSseChannelActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetComputeApiSseChannelActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "output-vector",
			Type: "slice",
		},
	}
}
func CastComputeApiSseChannelActionResFromCli(c emigo.CliCastable) ComputeApiSseChannelActionRes {
	data := ComputeApiSseChannelActionRes{}
	if c.IsSet("output-vector") {
		emigo.InflatePossibleSlice(c.String("output-vector"), &data.OutputVector)
	}
	return data
}

// The base class definition for computeApiSseChannelActionRes
type ComputeApiSseChannelActionRes struct {
	OutputVector []int `json:"outputVector" yaml:"outputVector"`
}

func (x *ComputeApiSseChannelActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ComputeApiSseChannelActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ComputeApiSseChannelActionResponse) SetContentType(contentType string) *ComputeApiSseChannelActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ComputeApiSseChannelActionResponse) AsStream(r io.Reader, contentType string) *ComputeApiSseChannelActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ComputeApiSseChannelActionResponse) AsJSON(payload any) *ComputeApiSseChannelActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ComputeApiSseChannelActionResponse) WithIdeal(payload ComputeApiSseChannelActionRes) *ComputeApiSseChannelActionResponse {
	x.Payload = payload
	return x
}
func (x *ComputeApiSseChannelActionResponse) AsHTML(payload string) *ComputeApiSseChannelActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ComputeApiSseChannelActionResponse) AsBytes(payload []byte) *ComputeApiSseChannelActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ComputeApiSseChannelActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ComputeApiSseChannelActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ComputeApiSseChannelActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ComputeApiSseChannelActionRaw registers a raw Gin route for the ComputeApiSseChannelAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ComputeApiSseChannelActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ComputeApiSseChannelActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ComputeApiSseChannelActionRequestSig = func(c ComputeApiSseChannelActionRequest) (*ComputeApiSseChannelActionResponse, error)

// ComputeApiSseChannelActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ComputeApiSseChannelAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ComputeApiSseChannelActionHandler(
	handler ComputeApiSseChannelActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ComputeApiSseChannelActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ComputeApiSseChannelActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ComputeApiSseChannelActionRequest{
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

// ComputeApiSseChannelAction is a high-level convenience wrapper around ComputeApiSseChannelActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ComputeApiSseChannelActionGin(r gin.IRoutes, handler ComputeApiSseChannelActionRequestSig) {
	method, url, h := ComputeApiSseChannelActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for ComputeApiSseChannelAction
 */
// Query wrapper with private fields
type ComputeApiSseChannelActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ComputeApiSseChannelActionQueryFromString(rawQuery string) ComputeApiSseChannelActionQuery {
	v := ComputeApiSseChannelActionQuery{}
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
func ComputeApiSseChannelActionQueryFromGin(c *gin.Context) ComputeApiSseChannelActionQuery {
	return ComputeApiSseChannelActionQueryFromString(c.Request.URL.RawQuery)
}
func ComputeApiSseChannelActionQueryFromHttp(r *http.Request) ComputeApiSseChannelActionQuery {
	return ComputeApiSseChannelActionQueryFromString(r.URL.RawQuery)
}
func (q ComputeApiSseChannelActionQuery) Values() url.Values {
	return q.values
}
func (q ComputeApiSseChannelActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ComputeApiSseChannelActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ComputeApiSseChannelActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ComputeApiSseChannelActionRequest struct {
	Body        ComputeApiSseChannelActionReq
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

func (x ComputeApiSseChannelActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x ComputeApiSseChannelActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type ComputeApiSseChannelActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func ComputeApiSseChannelActionClientCreateUrl(
	req ComputeApiSseChannelActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ComputeApiSseChannelActionMeta()
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
func ComputeApiSseChannelActionClientExecuteTyped(httpReq *http.Request) (*ComputeApiSseChannelActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ComputeApiSseChannelActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ComputeApiSseChannelActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ComputeApiSseChannelActionResponse{Payload: result}, err
	}
	return &ComputeApiSseChannelActionResponse{Payload: result}, nil
}
func ComputeApiSseChannelActionClientBuildRequest(req ComputeApiSseChannelActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ComputeApiSseChannelActionMeta()
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
func ComputeApiSseChannelActionCall(
	req ComputeApiSseChannelActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ComputeApiSseChannelActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ComputeApiSseChannelActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ComputeApiSseChannelActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ComputeApiSseChannelActionClientExecuteTyped(r)
}
