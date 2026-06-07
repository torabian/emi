package external

import (
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action PublishCommandDetailedReportAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PublishCommandDetailedReportAction
func PublishCommandDetailedReportAction(c PublishCommandDetailedReportActionRequest) (*PublishCommandDetailedReportActionResponse, error) {
	return &PublishCommandDetailedReportActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PublishCommandDetailedReportActionMeta() struct {
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
		Name:        "PublishCommandDetailedReportAction",
		CliName:     "publish command detailed report-action",
		URL:         "https://api.{environment}/sale/offer-publication-commands/{commandId}/tasks",
		Method:      "GET",
		Description: `Use this resource to retrieve information about the offer statuses on the site (Defaults: limit = 100, offset = 0). Read more: PL / EN. This resource is rate limited to retrieving information about 270 000 offer changes per minute.`,
	}
}

// The base class definition for publishCommandDetailedReportActionRes
type PublishCommandDetailedReportActionRes struct {
	Tasks emigo.Array[PublishCommandDetailedReportActionResTasks] `json:"tasks" yaml:"tasks"`
}

// The base class definition for tasks
type PublishCommandDetailedReportActionResTasks struct {
	Field   string                                                        `json:"field" yaml:"field"`
	Message string                                                        `json:"message" yaml:"message"`
	Offer   PublishCommandDetailedReportActionResTasksOffer               `json:"offer" yaml:"offer"`
	Status  string                                                        `json:"status" yaml:"status"`
	Errors  emigo.Array[PublishCommandDetailedReportActionResTasksErrors] `json:"errors" yaml:"errors"`
}

// The base class definition for offer
type PublishCommandDetailedReportActionResTasksOffer struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for errors
type PublishCommandDetailedReportActionResTasksErrors struct {
	Code        string                                                   `json:"code" yaml:"code"`
	Details     string                                                   `json:"details" yaml:"details"`
	Message     string                                                   `json:"message" yaml:"message"`
	Path        string                                                   `json:"path" yaml:"path"`
	UserMessage string                                                   `json:"userMessage" yaml:"userMessage"`
	Metadata    PublishCommandDetailedReportActionResTasksErrorsMetadata `json:"metadata" yaml:"metadata"`
}

// The base class definition for metadata
type PublishCommandDetailedReportActionResTasksErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

func (x *PublishCommandDetailedReportActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type PublishCommandDetailedReportActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PublishCommandDetailedReportActionResponse) SetContentType(contentType string) *PublishCommandDetailedReportActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PublishCommandDetailedReportActionResponse) AsStream(r io.Reader, contentType string) *PublishCommandDetailedReportActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PublishCommandDetailedReportActionResponse) AsJSON(payload any) *PublishCommandDetailedReportActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *PublishCommandDetailedReportActionResponse) WithIdeal(payload PublishCommandDetailedReportActionRes) *PublishCommandDetailedReportActionResponse {
	x.Payload = payload
	return x
}
func (x *PublishCommandDetailedReportActionResponse) AsHTML(payload string) *PublishCommandDetailedReportActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PublishCommandDetailedReportActionResponse) AsBytes(payload []byte) *PublishCommandDetailedReportActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PublishCommandDetailedReportActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PublishCommandDetailedReportActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PublishCommandDetailedReportActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type PublishCommandDetailedReportActionRequestSig = func(c PublishCommandDetailedReportActionRequest) (*PublishCommandDetailedReportActionResponse, error)

/**
 * Query parameters for Publish command detailed reportAction
 */
// Query wrapper with private fields
type PublishCommandDetailedReportActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func PublishCommandDetailedReportActionQueryFromString(rawQuery string) PublishCommandDetailedReportActionQuery {
	v := PublishCommandDetailedReportActionQuery{}
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
func PublishCommandDetailedReportActionQueryFromHttp(r *http.Request) PublishCommandDetailedReportActionQuery {
	return PublishCommandDetailedReportActionQueryFromString(r.URL.RawQuery)
}
func (q PublishCommandDetailedReportActionQuery) Values() url.Values {
	return q.values
}
func (q PublishCommandDetailedReportActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PublishCommandDetailedReportActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PublishCommandDetailedReportActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PublishCommandDetailedReportActionRequest struct {
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

func PublishCommandDetailedReportActionClientCreateUrl(
	req PublishCommandDetailedReportActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PublishCommandDetailedReportActionMeta()
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
func PublishCommandDetailedReportActionClientExecuteTyped(httpReq *http.Request) (*PublishCommandDetailedReportActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PublishCommandDetailedReportActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &PublishCommandDetailedReportActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &PublishCommandDetailedReportActionResponse{Payload: result}, err
	}
	return &PublishCommandDetailedReportActionResponse{Payload: result}, nil
}
func PublishCommandDetailedReportActionClientBuildRequest(req PublishCommandDetailedReportActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PublishCommandDetailedReportActionMeta()
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
func PublishCommandDetailedReportActionCall(
	req PublishCommandDetailedReportActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PublishCommandDetailedReportActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PublishCommandDetailedReportActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PublishCommandDetailedReportActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PublishCommandDetailedReportActionClientExecuteTyped(r)
}

// PublishCommandDetailedReportActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the PublishCommandDetailedReportAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *PublishCommandDetailedReportActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func PublishCommandDetailedReportActionHttpHandler(
	handler func(c PublishCommandDetailedReportActionRequest) (*PublishCommandDetailedReportActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := PublishCommandDetailedReportActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := PublishCommandDetailedReportActionRequest{
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

// PublishCommandDetailedReportActionHttp is a high-level convenience wrapper around
// PublishCommandDetailedReportActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func PublishCommandDetailedReportActionHttp(
	mux *http.ServeMux,
	handler func(c PublishCommandDetailedReportActionRequest) (*PublishCommandDetailedReportActionResponse, error),
) {
	method, pattern, h := PublishCommandDetailedReportActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
