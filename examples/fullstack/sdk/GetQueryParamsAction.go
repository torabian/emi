package external

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/**
* Action to communicate with the action GetQueryParamsAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetQueryParamsAction
func GetQueryParamsAction(c GetQueryParamsActionRequest) (*GetQueryParamsActionResponse, error) {
	return &GetQueryParamsActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetQueryParamsActionMeta() struct {
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
		Name:        "GetQueryParamsAction",
		CliName:     "get-query-params-action",
		URL:         "/stream/:addres1/:addressName2/:count",
		Method:      "GET",
		Description: `The goal is to check that if query params also become available in the cli params`,
	}
}
func GetGetQueryParamsActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastGetQueryParamsActionResFromCli(c emigo.CliCastable) GetQueryParamsActionRes {
	data := GetQueryParamsActionRes{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for getQueryParamsActionRes
type GetQueryParamsActionRes struct {
	Name string `json:"name" yaml:"name"`
}

func (x *GetQueryParamsActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetQueryParamsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetQueryParamsActionResponse) SetContentType(contentType string) *GetQueryParamsActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetQueryParamsActionResponse) AsStream(r io.Reader, contentType string) *GetQueryParamsActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetQueryParamsActionResponse) AsJSON(payload any) *GetQueryParamsActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetQueryParamsActionResponse) WithIdeal(payload GetQueryParamsActionRes) *GetQueryParamsActionResponse {
	x.Payload = payload
	return x
}
func (x *GetQueryParamsActionResponse) AsHTML(payload string) *GetQueryParamsActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetQueryParamsActionResponse) AsBytes(payload []byte) *GetQueryParamsActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetQueryParamsActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetQueryParamsActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetQueryParamsActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetQueryParamsActionRaw registers a raw Gin route for the GetQueryParamsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetQueryParamsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetQueryParamsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetQueryParamsActionRequestSig = func(c GetQueryParamsActionRequest) (*GetQueryParamsActionResponse, error)

// GetQueryParamsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetQueryParamsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetQueryParamsActionHandler(
	handler GetQueryParamsActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetQueryParamsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetQueryParamsActionRequest{
			Body:        nil,
			Params:      GetQueryParamsActionPathParameterFromGin(m),
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

// GetQueryParamsAction is a high-level convenience wrapper around GetQueryParamsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetQueryParamsActionGin(r gin.IRoutes, handler GetQueryParamsActionRequestSig) {
	method, url, h := GetQueryParamsActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Path parameters for GetQueryParamsAction
 */
type GetQueryParamsActionPathParameter struct {
	Addres1      string
	AddressName2 string
	Count        int
}

func GetGetQueryParamsActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-addres1",
			Type:     "string",
			Required: true,
		},
		{
			Name:     prefix + "pp-addressName2",
			Type:     "string",
			Required: true,
		},
		{
			Name:     prefix + "pp-count",
			Type:     "int",
			Required: true,
		},
	}
}

// Converts a placeholder url, and applies the parameters to it.
func GetQueryParamsActionPathParameterApply(params GetQueryParamsActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":addres1", fmt.Sprintf("%v", params.Addres1))
	templateUrl = strings.ReplaceAll(templateUrl, ":addressName2", fmt.Sprintf("%v", params.AddressName2))
	templateUrl = strings.ReplaceAll(templateUrl, ":count", fmt.Sprintf("%v", params.Count))
	return templateUrl
}

// Extracts the path parameter from a gin request context
func GetQueryParamsActionPathParameterFromGin(g *gin.Context) GetQueryParamsActionPathParameter {
	return GetQueryParamsActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// Extracts the path parameter from a urfave v3 cli.
func GetQueryParamsActionPathParameterFromCli(c *cli.Command) GetQueryParamsActionPathParameter {
	return GetQueryParamsActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}

// General purpose to extract the value and cast based on type.
func GetQueryParamsActionPathParameterFromFn(fn func(key string) string) GetQueryParamsActionPathParameter {
	res := GetQueryParamsActionPathParameter{}
	res.Addres1 = fn("addres1")
	res.AddressName2 = fn("addressName2")
	if v := fn("count"); v != "" {
		res.Count, _ = strconv.Atoi(v)
	}
	return res
}

/**
 * Query parameters for GetQueryParamsAction
 */
// Query wrapper with private fields
type GetQueryParamsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetQueryParamsActionQueryFromString(rawQuery string) GetQueryParamsActionQuery {
	v := GetQueryParamsActionQuery{}
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
func GetQueryParamsActionQueryFromGin(c *gin.Context) GetQueryParamsActionQuery {
	return GetQueryParamsActionQueryFromString(c.Request.URL.RawQuery)
}
func GetQueryParamsActionQueryFromHttp(r *http.Request) GetQueryParamsActionQuery {
	return GetQueryParamsActionQueryFromString(r.URL.RawQuery)
}
func (q GetQueryParamsActionQuery) Values() url.Values {
	return q.values
}
func (q GetQueryParamsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetQueryParamsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetQueryParamsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetQueryParamsActionRequest struct {
	Body        interface{}
	Params      GetQueryParamsActionPathParameter
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

func (x GetQueryParamsActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GetQueryParamsActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GetQueryParamsActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GetQueryParamsActionClientCreateUrl(
	req GetQueryParamsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetQueryParamsActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = GetQueryParamsActionPathParameterApply(req.Params, urlAddr)
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
func GetQueryParamsActionClientExecuteTyped(httpReq *http.Request) (*GetQueryParamsActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetQueryParamsActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetQueryParamsActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetQueryParamsActionResponse{Payload: result}, err
	}
	return &GetQueryParamsActionResponse{Payload: result}, nil
}
func GetQueryParamsActionClientBuildRequest(req GetQueryParamsActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetQueryParamsActionMeta()
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
func GetQueryParamsActionCall(
	req GetQueryParamsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetQueryParamsActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetQueryParamsActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetQueryParamsActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetQueryParamsActionClientExecuteTyped(r)
}
