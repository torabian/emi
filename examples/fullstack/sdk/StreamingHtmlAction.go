package external

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action StreamingHtmlAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of StreamingHtmlAction
func StreamingHtmlAction(c StreamingHtmlActionRequest) (*StreamingHtmlActionResponse, error) {
	return &StreamingHtmlActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func StreamingHtmlActionMeta() struct {
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
		Name:        "StreamingHtmlAction",
		CliName:     "streaming-html-action",
		URL:         "/stream/html",
		Method:      "GET",
		Description: ``,
	}
}

type StreamingHtmlActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *StreamingHtmlActionResponse) SetContentType(contentType string) *StreamingHtmlActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *StreamingHtmlActionResponse) AsStream(r io.Reader, contentType string) *StreamingHtmlActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *StreamingHtmlActionResponse) AsJSON(payload any) *StreamingHtmlActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *StreamingHtmlActionResponse) AsHTML(payload string) *StreamingHtmlActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *StreamingHtmlActionResponse) AsBytes(payload []byte) *StreamingHtmlActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x StreamingHtmlActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x StreamingHtmlActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x StreamingHtmlActionResponse) GetPayload() interface{} {
	return x.Payload
}

// StreamingHtmlActionRaw registers a raw Gin route for the StreamingHtmlAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func StreamingHtmlActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := StreamingHtmlActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type StreamingHtmlActionRequestSig = func(c StreamingHtmlActionRequest) (*StreamingHtmlActionResponse, error)

// StreamingHtmlActionHandler returns the HTTP method, route URL, and a typed Gin handler for the StreamingHtmlAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func StreamingHtmlActionHandler(
	handler StreamingHtmlActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := StreamingHtmlActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := StreamingHtmlActionRequest{
			Body:        nil,
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

// StreamingHtmlAction is a high-level convenience wrapper around StreamingHtmlActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func StreamingHtmlActionGin(r gin.IRoutes, handler StreamingHtmlActionRequestSig) {
	method, url, h := StreamingHtmlActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for StreamingHtmlAction
 */
// Query wrapper with private fields
type StreamingHtmlActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func StreamingHtmlActionQueryFromString(rawQuery string) StreamingHtmlActionQuery {
	v := StreamingHtmlActionQuery{}
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
func StreamingHtmlActionQueryFromGin(c *gin.Context) StreamingHtmlActionQuery {
	return StreamingHtmlActionQueryFromString(c.Request.URL.RawQuery)
}
func StreamingHtmlActionQueryFromHttp(r *http.Request) StreamingHtmlActionQuery {
	return StreamingHtmlActionQueryFromString(r.URL.RawQuery)
}
func (q StreamingHtmlActionQuery) Values() url.Values {
	return q.values
}
func (q StreamingHtmlActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *StreamingHtmlActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *StreamingHtmlActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type StreamingHtmlActionRequest struct {
	Body        interface{}
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

func (x StreamingHtmlActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x StreamingHtmlActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type StreamingHtmlActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func StreamingHtmlActionClientCreateUrl(
	req StreamingHtmlActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := StreamingHtmlActionMeta()
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
func StreamingHtmlActionClientExecuteTyped(httpReq *http.Request) (*StreamingHtmlActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result StreamingHtmlActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &StreamingHtmlActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &StreamingHtmlActionResponse{Payload: result}, err
	}
	return &StreamingHtmlActionResponse{Payload: result}, nil
}
func StreamingHtmlActionClientBuildRequest(req StreamingHtmlActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := StreamingHtmlActionMeta()
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
func StreamingHtmlActionCall(
	req StreamingHtmlActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*StreamingHtmlActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := StreamingHtmlActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := StreamingHtmlActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return StreamingHtmlActionClientExecuteTyped(r)
}
