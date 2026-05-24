package external

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"strings"
	"test/emi/emigo"
)

/**
* Action to communicate with the action GetAsGiantsAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetAsGiantsAction
func GetAsGiantsAction(c GetAsGiantsActionRequest) (*GetAsGiantsActionResponse, error) {
	return &GetAsGiantsActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetAsGiantsActionMeta() struct {
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
		Name:        "GetAsGiantsAction",
		CliName:     "get-as-giants-action",
		URL:         "/get/giant/:id",
		Method:      "GET",
		Description: ``,
	}
}

type GetAsGiantsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetAsGiantsActionResponse) SetContentType(contentType string) *GetAsGiantsActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetAsGiantsActionResponse) AsStream(r io.Reader, contentType string) *GetAsGiantsActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetAsGiantsActionResponse) AsJSON(payload any) *GetAsGiantsActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *GetAsGiantsActionResponse) AsHTML(payload string) *GetAsGiantsActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetAsGiantsActionResponse) AsBytes(payload []byte) *GetAsGiantsActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetAsGiantsActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetAsGiantsActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetAsGiantsActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetAsGiantsActionRaw registers a raw Gin route for the GetAsGiantsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetAsGiantsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetAsGiantsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetAsGiantsActionRequestSig = func(c GetAsGiantsActionRequest) (*GetAsGiantsActionResponse, error)

// GetAsGiantsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetAsGiantsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetAsGiantsActionHandler(
	handler GetAsGiantsActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetAsGiantsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetAsGiantsActionRequest{
			Body:        nil,
			Params:      GetAsGiantsActionPathParameterFromGin(m),
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

// GetAsGiantsAction is a high-level convenience wrapper around GetAsGiantsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetAsGiantsActionGin(r gin.IRoutes, handler GetAsGiantsActionRequestSig) {
	method, url, h := GetAsGiantsActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Path parameters for GetAsGiantsAction
 */
type GetAsGiantsActionPathParameter struct {
	Id string
}

func GetGetAsGiantsActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-id",
			Type:     "string",
			Required: true,
		},
	}
}

// Converts a placeholder url, and applies the parameters to it.
func GetAsGiantsActionPathParameterApply(params GetAsGiantsActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":id", fmt.Sprintf("%v", params.Id))
	return templateUrl
}

// Extracts the path parameter from a gin request context
func GetAsGiantsActionPathParameterFromGin(g *gin.Context) GetAsGiantsActionPathParameter {
	return GetAsGiantsActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// Extracts the path parameter from a urfave v3 cli.
func GetAsGiantsActionPathParameterFromCli(c *cli.Command) GetAsGiantsActionPathParameter {
	return GetAsGiantsActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}

// General purpose to extract the value and cast based on type.
func GetAsGiantsActionPathParameterFromFn(fn func(key string) string) GetAsGiantsActionPathParameter {
	res := GetAsGiantsActionPathParameter{}
	res.Id = fn("id")
	return res
}

/**
 * Query parameters for GetAsGiantsAction
 */
// Query wrapper with private fields
type GetAsGiantsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetAsGiantsActionQueryFromString(rawQuery string) GetAsGiantsActionQuery {
	v := GetAsGiantsActionQuery{}
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
func GetAsGiantsActionQueryFromGin(c *gin.Context) GetAsGiantsActionQuery {
	return GetAsGiantsActionQueryFromString(c.Request.URL.RawQuery)
}
func GetAsGiantsActionQueryFromHttp(r *http.Request) GetAsGiantsActionQuery {
	return GetAsGiantsActionQueryFromString(r.URL.RawQuery)
}
func (q GetAsGiantsActionQuery) Values() url.Values {
	return q.values
}
func (q GetAsGiantsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetAsGiantsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetAsGiantsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetAsGiantsActionRequest struct {
	Body        interface{}
	Params      GetAsGiantsActionPathParameter
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

func (x GetAsGiantsActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GetAsGiantsActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GetAsGiantsActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GetAsGiantsActionClientCreateUrl(
	req GetAsGiantsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetAsGiantsActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = GetAsGiantsActionPathParameterApply(req.Params, urlAddr)
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
func GetAsGiantsActionClientExecuteTyped(httpReq *http.Request) (*GetAsGiantsActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetAsGiantsActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetAsGiantsActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetAsGiantsActionResponse{Payload: result}, err
	}
	return &GetAsGiantsActionResponse{Payload: result}, nil
}
func GetAsGiantsActionClientBuildRequest(req GetAsGiantsActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetAsGiantsActionMeta()
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
func GetAsGiantsActionCall(
	req GetAsGiantsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetAsGiantsActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetAsGiantsActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetAsGiantsActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetAsGiantsActionClientExecuteTyped(r)
}
