package external

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ComputeApiSseAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ComputeApiSseAction
func ComputeApiSseAction(c ComputeApiSseActionRequest) (*ComputeApiSseActionResponse, error) {
	return &ComputeApiSseActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ComputeApiSseActionMeta() struct {
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
		Name:        "ComputeApiSseAction",
		CliName:     "compute-api-sse-action",
		URL:         "/compute/sse",
		Method:      "POST",
		Description: `The same compute api, but it would return the response as SSE.`,
	}
}
func GetComputeApiSseActionReqCliFlags(prefix string) []emigo.CliFlag {
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
func CastComputeApiSseActionReqFromCli(c emigo.CliCastable) ComputeApiSseActionReq {
	data := ComputeApiSseActionReq{}
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

// The base class definition for computeApiSseActionReq
type ComputeApiSseActionReq struct {
	InitialVector1 []int                  `json:"initialVector1" yaml:"initialVector1"`
	Value          emigo.Nullable[string] `json:"value" yaml:"value"`
	InitialVector2 []int                  `json:"initialVector2" yaml:"initialVector2"`
}

func (x *ComputeApiSseActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetComputeApiSseActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "output-vector",
			Type: "slice",
		},
	}
}
func CastComputeApiSseActionResFromCli(c emigo.CliCastable) ComputeApiSseActionRes {
	data := ComputeApiSseActionRes{}
	if c.IsSet("output-vector") {
		emigo.InflatePossibleSlice(c.String("output-vector"), &data.OutputVector)
	}
	return data
}

// The base class definition for computeApiSseActionRes
type ComputeApiSseActionRes struct {
	OutputVector []int `json:"outputVector" yaml:"outputVector"`
}

func (x *ComputeApiSseActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ComputeApiSseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ComputeApiSseActionResponse) SetContentType(contentType string) *ComputeApiSseActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ComputeApiSseActionResponse) AsStream(r io.Reader, contentType string) *ComputeApiSseActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ComputeApiSseActionResponse) AsJSON(payload any) *ComputeApiSseActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ComputeApiSseActionResponse) WithIdeal(payload ComputeApiSseActionRes) *ComputeApiSseActionResponse {
	x.Payload = payload
	return x
}
func (x *ComputeApiSseActionResponse) AsHTML(payload string) *ComputeApiSseActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ComputeApiSseActionResponse) AsBytes(payload []byte) *ComputeApiSseActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ComputeApiSseActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ComputeApiSseActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ComputeApiSseActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ComputeApiSseActionRaw registers a raw Gin route for the ComputeApiSseAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ComputeApiSseActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ComputeApiSseActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ComputeApiSseActionRequestSig = func(c ComputeApiSseActionRequest) (*ComputeApiSseActionResponse, error)

// ComputeApiSseActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ComputeApiSseAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ComputeApiSseActionHandler(
	handler ComputeApiSseActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ComputeApiSseActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ComputeApiSseActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ComputeApiSseActionRequest{
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

// ComputeApiSseAction is a high-level convenience wrapper around ComputeApiSseActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ComputeApiSseActionGin(r gin.IRoutes, handler ComputeApiSseActionRequestSig) {
	method, url, h := ComputeApiSseActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for ComputeApiSseAction
 */
// Query wrapper with private fields
type ComputeApiSseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ComputeApiSseActionQueryFromString(rawQuery string) ComputeApiSseActionQuery {
	v := ComputeApiSseActionQuery{}
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
func ComputeApiSseActionQueryFromGin(c *gin.Context) ComputeApiSseActionQuery {
	return ComputeApiSseActionQueryFromString(c.Request.URL.RawQuery)
}
func ComputeApiSseActionQueryFromHttp(r *http.Request) ComputeApiSseActionQuery {
	return ComputeApiSseActionQueryFromString(r.URL.RawQuery)
}
func (q ComputeApiSseActionQuery) Values() url.Values {
	return q.values
}
func (q ComputeApiSseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ComputeApiSseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ComputeApiSseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ComputeApiSseActionRequest struct {
	Body        ComputeApiSseActionReq
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	GinCtx *gin.Context
	// Urfave context, per each request
	CliCtx *cli.Context
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func (x ComputeApiSseActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x ComputeApiSseActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type ComputeApiSseActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func ComputeApiSseActionClientCreateUrl(
	req ComputeApiSseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ComputeApiSseActionMeta()
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
func ComputeApiSseActionClientExecuteTyped(httpReq *http.Request) (*ComputeApiSseActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ComputeApiSseActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ComputeApiSseActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ComputeApiSseActionResponse{Payload: result}, err
	}
	return &ComputeApiSseActionResponse{Payload: result}, nil
}
func ComputeApiSseActionClientBuildRequest(req ComputeApiSseActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ComputeApiSseActionMeta()
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
func ComputeApiSseActionCall(
	req ComputeApiSseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ComputeApiSseActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ComputeApiSseActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ComputeApiSseActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ComputeApiSseActionClientExecuteTyped(r)
}
