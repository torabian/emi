package external

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action AllDataAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of AllDataAction
func AllDataAction(c AllDataActionRequest) (*AllDataActionResponse, error) {
	return &AllDataActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func AllDataActionMeta() struct {
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
		Name:        "AllDataAction",
		CliName:     "all-data-action",
		URL:         "/res/22",
		Method:      "",
		Description: ``,
	}
}
func GetAllDataActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "string-type",
			Type: "string",
		},
		{
			Name: prefix + "string-type-null",
			Type: "string?",
		},
		{
			Name: prefix + "collection",
			Type: "collection",
		},
	}
}
func CastAllDataActionResFromCli(c emigo.CliCastable) AllDataActionRes {
	data := AllDataActionRes{}
	if c.IsSet("string-type") {
		data.StringType = c.String("string-type")
	}
	if c.IsSet("string-type-null") {
		emigo.ParseNullable(c.String("string-type-null"), &data.StringTypeNull)
	}
	return data
}

// The base class definition for allDataActionRes
type AllDataActionRes struct {
	StringType     string                    `json:"stringType" yaml:"stringType"`
	StringTypeNull emigo.Nullable[string]    `json:"stringTypeNull" yaml:"stringTypeNull"`
	Collection     []CommonVectorResponseDto `json:"collection" yaml:"collection"`
}

func (x *AllDataActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type AllDataActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *AllDataActionResponse) SetContentType(contentType string) *AllDataActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *AllDataActionResponse) AsStream(r io.Reader, contentType string) *AllDataActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *AllDataActionResponse) AsJSON(payload any) *AllDataActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *AllDataActionResponse) WithIdeal(payload AllDataActionRes) *AllDataActionResponse {
	x.Payload = payload
	return x
}
func (x *AllDataActionResponse) AsHTML(payload string) *AllDataActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *AllDataActionResponse) AsBytes(payload []byte) *AllDataActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x AllDataActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x AllDataActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x AllDataActionResponse) GetPayload() interface{} {
	return x.Payload
}

// AllDataActionRaw registers a raw Gin route for the AllDataAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func AllDataActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := AllDataActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type AllDataActionRequestSig = func(c AllDataActionRequest) (*AllDataActionResponse, error)

// AllDataActionHandler returns the HTTP method, route URL, and a typed Gin handler for the AllDataAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func AllDataActionHandler(
	handler AllDataActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := AllDataActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := AllDataActionRequest{
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

// AllDataAction is a high-level convenience wrapper around AllDataActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func AllDataActionGin(r gin.IRoutes, handler AllDataActionRequestSig) {
	method, url, h := AllDataActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for AllDataAction
 */
// Query wrapper with private fields
type AllDataActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func AllDataActionQueryFromString(rawQuery string) AllDataActionQuery {
	v := AllDataActionQuery{}
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
func AllDataActionQueryFromGin(c *gin.Context) AllDataActionQuery {
	return AllDataActionQueryFromString(c.Request.URL.RawQuery)
}
func AllDataActionQueryFromHttp(r *http.Request) AllDataActionQuery {
	return AllDataActionQueryFromString(r.URL.RawQuery)
}
func (q AllDataActionQuery) Values() url.Values {
	return q.values
}
func (q AllDataActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *AllDataActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *AllDataActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type AllDataActionRequest struct {
	Body        interface{}
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

func (x AllDataActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x AllDataActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type AllDataActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func AllDataActionClientCreateUrl(
	req AllDataActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := AllDataActionMeta()
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
func AllDataActionClientExecuteTyped(httpReq *http.Request) (*AllDataActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result AllDataActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &AllDataActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &AllDataActionResponse{Payload: result}, err
	}
	return &AllDataActionResponse{Payload: result}, nil
}
func AllDataActionClientBuildRequest(req AllDataActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := AllDataActionMeta()
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
func AllDataActionCall(
	req AllDataActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*AllDataActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := AllDataActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := AllDataActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return AllDataActionClientExecuteTyped(r)
}
