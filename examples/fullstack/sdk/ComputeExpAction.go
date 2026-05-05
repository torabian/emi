package external

import (
	"bytes"
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"
)

/**
* Action to communicate with the action ComputeExpAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ComputeExpAction
func ComputeExpAction(c ComputeExpActionRequest) (*ComputeExpActionResponse, error) {
	return &ComputeExpActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ComputeExpActionMeta() struct {
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
		Name:        "ComputeExpAction",
		CliName:     "compute-exp-action",
		URL:         "/big/exp/:first/:second",
		Method:      "",
		Description: `Computes the exp value using big integer`,
	}
}
func GetComputeExpActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "base",
			Type: "complex",
		},
		{
			Name: prefix + "exponent",
			Type: "complex",
		},
	}
}
func CastComputeExpActionReqFromCli(c emigo.CliCastable) ComputeExpActionReq {
	data := ComputeExpActionReq{}
	if c.IsSet("base") {
		if u, ok := any(&data.Base).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("base")))
		}
	}
	if c.IsSet("exponent") {
		if u, ok := any(&data.Exponent).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("exponent")))
		}
	}
	return data
}

// The base class definition for computeExpActionReq
type ComputeExpActionReq struct {
	Base     big.Int `json:"base" yaml:"base"`
	Exponent big.Int `json:"exponent" yaml:"exponent"`
}

func (x *ComputeExpActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetComputeExpActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "result",
			Type: "complex",
		},
	}
}
func CastComputeExpActionResFromCli(c emigo.CliCastable) ComputeExpActionRes {
	data := ComputeExpActionRes{}
	if c.IsSet("result") {
		if u, ok := any(&data.Result).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("result")))
		}
	}
	return data
}

// The base class definition for computeExpActionRes
type ComputeExpActionRes struct {
	Result big.Int `json:"result" yaml:"result"`
}

func (x *ComputeExpActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ComputeExpActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ComputeExpActionResponse) SetContentType(contentType string) *ComputeExpActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ComputeExpActionResponse) AsStream(r io.Reader, contentType string) *ComputeExpActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ComputeExpActionResponse) AsJSON(payload any) *ComputeExpActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ComputeExpActionResponse) WithIdeal(payload ComputeExpActionRes) *ComputeExpActionResponse {
	x.Payload = payload
	return x
}
func (x *ComputeExpActionResponse) AsHTML(payload string) *ComputeExpActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ComputeExpActionResponse) AsBytes(payload []byte) *ComputeExpActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ComputeExpActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ComputeExpActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ComputeExpActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ComputeExpActionRaw registers a raw Gin route for the ComputeExpAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ComputeExpActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ComputeExpActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ComputeExpActionRequestSig = func(c ComputeExpActionRequest) (*ComputeExpActionResponse, error)

// ComputeExpActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ComputeExpAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ComputeExpActionHandler(
	handler ComputeExpActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ComputeExpActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ComputeExpActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ComputeExpActionRequest{
			Body:        body,
			Params:      ComputeExpActionPathParameterFromGin(m),
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

// ComputeExpAction is a high-level convenience wrapper around ComputeExpActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ComputeExpActionGin(r gin.IRoutes, handler ComputeExpActionRequestSig) {
	method, url, h := ComputeExpActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Path parameters for ComputeExpAction
 */
type ComputeExpActionPathParameter struct {
	First  string
	Second string
}

func GetComputeExpActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-first",
			Type:     "string",
			Required: true,
		},
		{
			Name:     prefix + "pp-second",
			Type:     "string",
			Required: true,
		},
	}
}

// Converts a placeholder url, and applies the parameters to it.
func ComputeExpActionPathParameterApply(params ComputeExpActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":first", fmt.Sprintf("%v", params.First))
	templateUrl = strings.ReplaceAll(templateUrl, ":second", fmt.Sprintf("%v", params.Second))
	return templateUrl
}

// Extracts the path parameter from a gin request context
func ComputeExpActionPathParameterFromGin(g *gin.Context) ComputeExpActionPathParameter {
	return ComputeExpActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// Extracts the path parameter from a urfave v3 cli.
func ComputeExpActionPathParameterFromCli(c *cli.Command) ComputeExpActionPathParameter {
	return ComputeExpActionPathParameterFromFn(func(key string) string {
		return c.String(key)
	})
}

// General purpose to extract the value and cast based on type.
func ComputeExpActionPathParameterFromFn(fn func(key string) string) ComputeExpActionPathParameter {
	res := ComputeExpActionPathParameter{}
	res.First = fn("first")
	res.Second = fn("second")
	return res
}

/**
 * Query parameters for ComputeExpAction
 */
// Query wrapper with private fields
type ComputeExpActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ComputeExpActionQueryFromString(rawQuery string) ComputeExpActionQuery {
	v := ComputeExpActionQuery{}
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
func ComputeExpActionQueryFromGin(c *gin.Context) ComputeExpActionQuery {
	return ComputeExpActionQueryFromString(c.Request.URL.RawQuery)
}
func ComputeExpActionQueryFromHttp(r *http.Request) ComputeExpActionQuery {
	return ComputeExpActionQueryFromString(r.URL.RawQuery)
}
func (q ComputeExpActionQuery) Values() url.Values {
	return q.values
}
func (q ComputeExpActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ComputeExpActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ComputeExpActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ComputeExpActionRequest struct {
	Body        ComputeExpActionReq
	Params      ComputeExpActionPathParameter
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

func (x ComputeExpActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x ComputeExpActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type ComputeExpActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func ComputeExpActionClientCreateUrl(
	req ComputeExpActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ComputeExpActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = ComputeExpActionPathParameterApply(req.Params, urlAddr)
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
func ComputeExpActionClientExecuteTyped(httpReq *http.Request) (*ComputeExpActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ComputeExpActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ComputeExpActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ComputeExpActionResponse{Payload: result}, err
	}
	return &ComputeExpActionResponse{Payload: result}, nil
}
func ComputeExpActionClientBuildRequest(req ComputeExpActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ComputeExpActionMeta()
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
func ComputeExpActionCall(
	req ComputeExpActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ComputeExpActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ComputeExpActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ComputeExpActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ComputeExpActionClientExecuteTyped(r)
}
