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
* Action to communicate with the action EnvelopeExampleAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of EnvelopeExampleAction
func EnvelopeExampleAction(c EnvelopeExampleActionRequest) (*EnvelopeExampleActionResponse, error) {
	return &EnvelopeExampleActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func EnvelopeExampleActionMeta() struct {
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
		Name:        "EnvelopeExampleAction",
		CliName:     "envelope-example-action",
		URL:         "/response/with/envelop",
		Method:      "GET",
		Description: ``,
	}
}
func GetEnvelopeExampleActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "content",
			Type: "string",
		},
	}
}
func CastEnvelopeExampleActionResFromCli(c emigo.CliCastable) EnvelopeExampleActionRes {
	data := EnvelopeExampleActionRes{}
	if c.IsSet("content") {
		data.Content = c.String("content")
	}
	return data
}

// The base class definition for envelopeExampleActionRes
type EnvelopeExampleActionRes struct {
	Content string `json:"content" yaml:"content"`
}

func (x *EnvelopeExampleActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type EnvelopeExampleActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *EnvelopeExampleActionResponse) SetContentType(contentType string) *EnvelopeExampleActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *EnvelopeExampleActionResponse) AsStream(r io.Reader, contentType string) *EnvelopeExampleActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *EnvelopeExampleActionResponse) AsJSON(payload any) *EnvelopeExampleActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *EnvelopeExampleActionResponse) WithIdeal(payload EnvelopeExampleActionRes) *EnvelopeExampleActionResponse {
	x.Payload = payload
	return x
}
func (x *EnvelopeExampleActionResponse) AsHTML(payload string) *EnvelopeExampleActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *EnvelopeExampleActionResponse) AsBytes(payload []byte) *EnvelopeExampleActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x EnvelopeExampleActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x EnvelopeExampleActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x EnvelopeExampleActionResponse) GetPayload() interface{} {
	return x.Payload
}

// EnvelopeExampleActionRaw registers a raw Gin route for the EnvelopeExampleAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func EnvelopeExampleActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := EnvelopeExampleActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type EnvelopeExampleActionRequestSig = func(c EnvelopeExampleActionRequest) (*EnvelopeExampleActionResponse, error)

// EnvelopeExampleActionHandler returns the HTTP method, route URL, and a typed Gin handler for the EnvelopeExampleAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func EnvelopeExampleActionHandler(
	handler EnvelopeExampleActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := EnvelopeExampleActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := EnvelopeExampleActionRequest{
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

// EnvelopeExampleAction is a high-level convenience wrapper around EnvelopeExampleActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func EnvelopeExampleActionGin(r gin.IRoutes, handler EnvelopeExampleActionRequestSig) {
	method, url, h := EnvelopeExampleActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for EnvelopeExampleAction
 */
// Query wrapper with private fields
type EnvelopeExampleActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func EnvelopeExampleActionQueryFromString(rawQuery string) EnvelopeExampleActionQuery {
	v := EnvelopeExampleActionQuery{}
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
func EnvelopeExampleActionQueryFromGin(c *gin.Context) EnvelopeExampleActionQuery {
	return EnvelopeExampleActionQueryFromString(c.Request.URL.RawQuery)
}
func EnvelopeExampleActionQueryFromHttp(r *http.Request) EnvelopeExampleActionQuery {
	return EnvelopeExampleActionQueryFromString(r.URL.RawQuery)
}
func (q EnvelopeExampleActionQuery) Values() url.Values {
	return q.values
}
func (q EnvelopeExampleActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *EnvelopeExampleActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *EnvelopeExampleActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type EnvelopeExampleActionRequest struct {
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

func (x EnvelopeExampleActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x EnvelopeExampleActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type EnvelopeExampleActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func EnvelopeExampleActionClientCreateUrl(
	req EnvelopeExampleActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := EnvelopeExampleActionMeta()
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
func EnvelopeExampleActionClientExecuteTyped(httpReq *http.Request) (*EnvelopeExampleActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result EnvelopeExampleActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &EnvelopeExampleActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &EnvelopeExampleActionResponse{Payload: result}, err
	}
	return &EnvelopeExampleActionResponse{Payload: result}, nil
}
func EnvelopeExampleActionClientBuildRequest(req EnvelopeExampleActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := EnvelopeExampleActionMeta()
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
func EnvelopeExampleActionCall(
	req EnvelopeExampleActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*EnvelopeExampleActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := EnvelopeExampleActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := EnvelopeExampleActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return EnvelopeExampleActionClientExecuteTyped(r)
}
