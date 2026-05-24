package external

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"test/emi/emigo"
)

/**
* Action to communicate with the action ModifyGiantDtoAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ModifyGiantDtoAction
func ModifyGiantDtoAction(c ModifyGiantDtoActionRequest) (*ModifyGiantDtoActionResponse, error) {
	return &ModifyGiantDtoActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ModifyGiantDtoActionMeta() struct {
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
		Name:        "ModifyGiantDtoAction",
		CliName:     "modify-giant-dto-action",
		URL:         "/modify/dto",
		Method:      "POST",
		Description: ``,
	}
}

type ModifyGiantDtoActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ModifyGiantDtoActionResponse) SetContentType(contentType string) *ModifyGiantDtoActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ModifyGiantDtoActionResponse) AsStream(r io.Reader, contentType string) *ModifyGiantDtoActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ModifyGiantDtoActionResponse) AsJSON(payload any) *ModifyGiantDtoActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *ModifyGiantDtoActionResponse) AsHTML(payload string) *ModifyGiantDtoActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ModifyGiantDtoActionResponse) AsBytes(payload []byte) *ModifyGiantDtoActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ModifyGiantDtoActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ModifyGiantDtoActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ModifyGiantDtoActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ModifyGiantDtoActionRaw registers a raw Gin route for the ModifyGiantDtoAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModifyGiantDtoActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModifyGiantDtoActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ModifyGiantDtoActionRequestSig = func(c ModifyGiantDtoActionRequest) (*ModifyGiantDtoActionResponse, error)

// ModifyGiantDtoActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModifyGiantDtoAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModifyGiantDtoActionHandler(
	handler ModifyGiantDtoActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ModifyGiantDtoActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body GiantDto
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ModifyGiantDtoActionRequest{
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

// ModifyGiantDtoAction is a high-level convenience wrapper around ModifyGiantDtoActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModifyGiantDtoActionGin(r gin.IRoutes, handler ModifyGiantDtoActionRequestSig) {
	method, url, h := ModifyGiantDtoActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for ModifyGiantDtoAction
 */
// Query wrapper with private fields
type ModifyGiantDtoActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModifyGiantDtoActionQueryFromString(rawQuery string) ModifyGiantDtoActionQuery {
	v := ModifyGiantDtoActionQuery{}
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
func ModifyGiantDtoActionQueryFromGin(c *gin.Context) ModifyGiantDtoActionQuery {
	return ModifyGiantDtoActionQueryFromString(c.Request.URL.RawQuery)
}
func ModifyGiantDtoActionQueryFromHttp(r *http.Request) ModifyGiantDtoActionQuery {
	return ModifyGiantDtoActionQueryFromString(r.URL.RawQuery)
}
func (q ModifyGiantDtoActionQuery) Values() url.Values {
	return q.values
}
func (q ModifyGiantDtoActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModifyGiantDtoActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModifyGiantDtoActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModifyGiantDtoActionRequest struct {
	Body        GiantDto
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

func (x ModifyGiantDtoActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x ModifyGiantDtoActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type ModifyGiantDtoActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func ModifyGiantDtoActionClientCreateUrl(
	req ModifyGiantDtoActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ModifyGiantDtoActionMeta()
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
func ModifyGiantDtoActionClientExecuteTyped(httpReq *http.Request) (*ModifyGiantDtoActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ModifyGiantDtoActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ModifyGiantDtoActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ModifyGiantDtoActionResponse{Payload: result}, err
	}
	return &ModifyGiantDtoActionResponse{Payload: result}, nil
}
func ModifyGiantDtoActionClientBuildRequest(req ModifyGiantDtoActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ModifyGiantDtoActionMeta()
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
func ModifyGiantDtoActionCall(
	req ModifyGiantDtoActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModifyGiantDtoActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ModifyGiantDtoActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ModifyGiantDtoActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ModifyGiantDtoActionClientExecuteTyped(r)
}
