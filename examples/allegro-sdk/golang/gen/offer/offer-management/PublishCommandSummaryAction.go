package external

import (
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action PublishCommandSummaryAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PublishCommandSummaryAction
func PublishCommandSummaryAction(c PublishCommandSummaryActionRequest) (*PublishCommandSummaryActionResponse, error) {
	return &PublishCommandSummaryActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PublishCommandSummaryActionMeta() struct {
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
		Name:        "PublishCommandSummaryAction",
		CliName:     "publish command summary-action",
		URL:         "https://api.{environment}/sale/offer-publication-commands/{commandId}",
		Method:      "GET",
		Description: `Use this resource to retrieve information about the offer listing statuses.  You will receive a summary with a number of correctly listed offers and errors.  Read more: PL / EN. This resource is rate limited to retrieving information about 270 000 offer changes per minute.`,
	}
}
func GetPublishCommandSummaryActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "created-at",
			Type: "string",
		},
		{
			Name: prefix + "completed-at",
			Type: "string",
		},
		{
			Name:     prefix + "task-count",
			Type:     "object",
			Children: GetPublishCommandSummaryActionResTaskCountCliFlags("task-count-"),
		},
	}
}
func CastPublishCommandSummaryActionResFromCli(c emigo.CliCastable) PublishCommandSummaryActionRes {
	data := PublishCommandSummaryActionRes{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("created-at") {
		data.CreatedAt = c.String("created-at")
	}
	if c.IsSet("completed-at") {
		data.CompletedAt = c.String("completed-at")
	}
	if c.IsSet("task-count") {
		data.TaskCount = CastPublishCommandSummaryActionResTaskCountFromCli(c)
	}
	return data
}

// The base class definition for publishCommandSummaryActionRes
type PublishCommandSummaryActionRes struct {
	Id          string                                  `json:"id" yaml:"id"`
	CreatedAt   string                                  `json:"createdAt" yaml:"createdAt"`
	CompletedAt string                                  `json:"completedAt" yaml:"completedAt"`
	TaskCount   PublishCommandSummaryActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}

func GetPublishCommandSummaryActionResTaskCountCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "failed",
			Type: "int",
		},
		{
			Name: prefix + "success",
			Type: "int",
		},
		{
			Name: prefix + "total",
			Type: "int",
		},
	}
}
func CastPublishCommandSummaryActionResTaskCountFromCli(c emigo.CliCastable) PublishCommandSummaryActionResTaskCount {
	data := PublishCommandSummaryActionResTaskCount{}
	if c.IsSet("failed") {
		data.Failed = int(c.Int64("failed"))
	}
	if c.IsSet("success") {
		data.Success = int(c.Int64("success"))
	}
	if c.IsSet("total") {
		data.Total = int(c.Int64("total"))
	}
	return data
}

// The base class definition for taskCount
type PublishCommandSummaryActionResTaskCount struct {
	Failed  int `json:"failed" yaml:"failed"`
	Success int `json:"success" yaml:"success"`
	Total   int `json:"total" yaml:"total"`
}

func (x *PublishCommandSummaryActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type PublishCommandSummaryActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PublishCommandSummaryActionResponse) SetContentType(contentType string) *PublishCommandSummaryActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PublishCommandSummaryActionResponse) AsStream(r io.Reader, contentType string) *PublishCommandSummaryActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PublishCommandSummaryActionResponse) AsJSON(payload any) *PublishCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *PublishCommandSummaryActionResponse) WithIdeal(payload PublishCommandSummaryActionRes) *PublishCommandSummaryActionResponse {
	x.Payload = payload
	return x
}
func (x *PublishCommandSummaryActionResponse) AsHTML(payload string) *PublishCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PublishCommandSummaryActionResponse) AsBytes(payload []byte) *PublishCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PublishCommandSummaryActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PublishCommandSummaryActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PublishCommandSummaryActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type PublishCommandSummaryActionRequestSig = func(c PublishCommandSummaryActionRequest) (*PublishCommandSummaryActionResponse, error)

/**
 * Query parameters for Publish command summaryAction
 */
// Query wrapper with private fields
type PublishCommandSummaryActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func PublishCommandSummaryActionQueryFromString(rawQuery string) PublishCommandSummaryActionQuery {
	v := PublishCommandSummaryActionQuery{}
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
func PublishCommandSummaryActionQueryFromHttp(r *http.Request) PublishCommandSummaryActionQuery {
	return PublishCommandSummaryActionQueryFromString(r.URL.RawQuery)
}
func (q PublishCommandSummaryActionQuery) Values() url.Values {
	return q.values
}
func (q PublishCommandSummaryActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PublishCommandSummaryActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PublishCommandSummaryActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PublishCommandSummaryActionRequest struct {
	Body        interface{}
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	// Now it's interface, so the code gen doesn't depend on the instance
	// or gin package. Make sure you cast is later into *gin.Context, or whatever
	// your framework is passing when creating a request.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	GinCtx interface{}
	// Cli library helper (urfave) by default. The instance is interface{}, and you
	// need to manually cast it to the *cli.Command, so gives you freedom and independence
	// of external library.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	CliCtx interface{}
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func PublishCommandSummaryActionClientCreateUrl(
	req PublishCommandSummaryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PublishCommandSummaryActionMeta()
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
func PublishCommandSummaryActionClientExecuteTyped(httpReq *http.Request) (*PublishCommandSummaryActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PublishCommandSummaryActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &PublishCommandSummaryActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &PublishCommandSummaryActionResponse{Payload: result}, err
	}
	return &PublishCommandSummaryActionResponse{Payload: result}, nil
}
func PublishCommandSummaryActionClientBuildRequest(req PublishCommandSummaryActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PublishCommandSummaryActionMeta()
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
func PublishCommandSummaryActionCall(
	req PublishCommandSummaryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PublishCommandSummaryActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PublishCommandSummaryActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PublishCommandSummaryActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PublishCommandSummaryActionClientExecuteTyped(r)
}
func (x PublishCommandSummaryActionRequest) IsCli() bool {
	if x.CliCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.CliCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}

// PublishCommandSummaryActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the PublishCommandSummaryAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *PublishCommandSummaryActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func PublishCommandSummaryActionHttpHandler(
	handler func(c PublishCommandSummaryActionRequest) (*PublishCommandSummaryActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := PublishCommandSummaryActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := PublishCommandSummaryActionRequest{
			Body:        nil,
			QueryParams: r.URL.Query(),
			Headers:     r.Header,
		}
		resp, err := handler(req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		// If the handler returned nil (and no error), the response was handled
		// manually.
		if resp == nil {
			return
		}
		// Apply headers
		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}
		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}
		if resp.Payload != nil {
			if w.Header().Get("Content-Type") == "" {
				w.Header().Set("Content-Type", "application/json")
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(resp.Payload)
		} else {
			w.WriteHeader(status)
		}
	}
}

// PublishCommandSummaryActionHttp is a high-level convenience wrapper around
// PublishCommandSummaryActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func PublishCommandSummaryActionHttp(
	mux *http.ServeMux,
	handler func(c PublishCommandSummaryActionRequest) (*PublishCommandSummaryActionResponse, error),
) {
	method, pattern, h := PublishCommandSummaryActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
