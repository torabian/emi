package external

import (
	"bytes"
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
* Action to communicate with the action PatchGiantAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PatchGiantAction
func PatchGiantAction(c PatchGiantActionRequest) (*PatchGiantActionResponse, error) {
	return &PatchGiantActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PatchGiantActionMeta() struct {
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
		Name:        "PatchGiantAction",
		CliName:     "patch-giant-action",
		URL:         "/get/giant/:id",
		Method:      "PATCH",
		Description: ``,
	}
}

type PatchGiantActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PatchGiantActionResponse) SetContentType(contentType string) *PatchGiantActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PatchGiantActionResponse) AsStream(r io.Reader, contentType string) *PatchGiantActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PatchGiantActionResponse) AsJSON(payload any) *PatchGiantActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *PatchGiantActionResponse) AsHTML(payload string) *PatchGiantActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PatchGiantActionResponse) AsBytes(payload []byte) *PatchGiantActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PatchGiantActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PatchGiantActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PatchGiantActionResponse) GetPayload() interface{} {
	return x.Payload
}

// PatchGiantActionRaw registers a raw Gin route for the PatchGiantAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PatchGiantActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PatchGiantActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type PatchGiantActionRequestSig = func(c PatchGiantActionRequest) (*PatchGiantActionResponse, error)

// PatchGiantActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PatchGiantAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PatchGiantActionHandler(
	handler PatchGiantActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := PatchGiantActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body GiantDto
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := PatchGiantActionRequest{
			Body:        body,
			Params:      PatchGiantActionPathParameterFromGin(m),
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

// PatchGiantAction is a high-level convenience wrapper around PatchGiantActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PatchGiantActionGin(r gin.IRoutes, handler PatchGiantActionRequestSig) {
	method, url, h := PatchGiantActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Path parameters for PatchGiantAction
 */
type PatchGiantActionPathParameter struct {
	Id string
}

func GetPatchGiantActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-id",
			Type:     "string",
			Required: true,
		},
	}
}

// Converts a placeholder url, and applies the parameters to it.
func PatchGiantActionPathParameterApply(params PatchGiantActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":id", fmt.Sprintf("%v", params.Id))
	return templateUrl
}

// Extracts the path parameter from a gin request context
func PatchGiantActionPathParameterFromGin(g *gin.Context) PatchGiantActionPathParameter {
	return PatchGiantActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// Extracts the path parameter from a urfave v3 cli.
func PatchGiantActionPathParameterFromCli(c *cli.Command) PatchGiantActionPathParameter {
	return PatchGiantActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}

// General purpose to extract the value and cast based on type.
func PatchGiantActionPathParameterFromFn(fn func(key string) string) PatchGiantActionPathParameter {
	res := PatchGiantActionPathParameter{}
	res.Id = fn("id")
	return res
}

/**
 * Query parameters for PatchGiantAction
 */
// Query wrapper with private fields
type PatchGiantActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func PatchGiantActionQueryFromString(rawQuery string) PatchGiantActionQuery {
	v := PatchGiantActionQuery{}
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
func PatchGiantActionQueryFromGin(c *gin.Context) PatchGiantActionQuery {
	return PatchGiantActionQueryFromString(c.Request.URL.RawQuery)
}
func PatchGiantActionQueryFromHttp(r *http.Request) PatchGiantActionQuery {
	return PatchGiantActionQueryFromString(r.URL.RawQuery)
}
func (q PatchGiantActionQuery) Values() url.Values {
	return q.values
}
func (q PatchGiantActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PatchGiantActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PatchGiantActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PatchGiantActionRequest struct {
	Body        GiantDto
	Params      PatchGiantActionPathParameter
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

func (x PatchGiantActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x PatchGiantActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type PatchGiantActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func PatchGiantActionClientCreateUrl(
	req PatchGiantActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PatchGiantActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = PatchGiantActionPathParameterApply(req.Params, urlAddr)
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
func PatchGiantActionClientExecuteTyped(httpReq *http.Request) (*PatchGiantActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PatchGiantActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &PatchGiantActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &PatchGiantActionResponse{Payload: result}, err
	}
	return &PatchGiantActionResponse{Payload: result}, nil
}
func PatchGiantActionClientBuildRequest(req PatchGiantActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PatchGiantActionMeta()
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
func PatchGiantActionCall(
	req PatchGiantActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PatchGiantActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PatchGiantActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PatchGiantActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PatchGiantActionClientExecuteTyped(r)
}
