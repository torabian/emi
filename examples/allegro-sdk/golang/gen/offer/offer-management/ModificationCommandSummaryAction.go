package external

import (
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ModificationCommandSummaryAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ModificationCommandSummaryAction
func ModificationCommandSummaryAction(c ModificationCommandSummaryActionRequest) (*ModificationCommandSummaryActionResponse, error) {
	return &ModificationCommandSummaryActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ModificationCommandSummaryActionMeta() struct {
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
		Name:        "ModificationCommandSummaryAction",
		CliName:     "modification command summary-action",
		URL:         "https://api.{environment}/sale/offers/promo-options-commands/{commandId}",
		Method:      "GET",
		Description: `Use this resource to find out how many offers were edited within one {commandId}. You will receive a summary with a number of successfully edited offers and errors. Read more: PL / EN.`,
	}
}

// The base class definition for modificationCommandSummaryActionRes
type ModificationCommandSummaryActionRes struct {
	Id        string                                       `json:"id" yaml:"id"`
	TaskCount ModificationCommandSummaryActionResTaskCount `json:"taskCount" yaml:"taskCount"`
}

// The base class definition for taskCount
type ModificationCommandSummaryActionResTaskCount struct {
	Failed  int `json:"failed" yaml:"failed"`
	Success int `json:"success" yaml:"success"`
	Total   int `json:"total" yaml:"total"`
}

func (x *ModificationCommandSummaryActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ModificationCommandSummaryActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ModificationCommandSummaryActionResponse) SetContentType(contentType string) *ModificationCommandSummaryActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ModificationCommandSummaryActionResponse) AsStream(r io.Reader, contentType string) *ModificationCommandSummaryActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ModificationCommandSummaryActionResponse) AsJSON(payload any) *ModificationCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ModificationCommandSummaryActionResponse) WithIdeal(payload ModificationCommandSummaryActionRes) *ModificationCommandSummaryActionResponse {
	x.Payload = payload
	return x
}
func (x *ModificationCommandSummaryActionResponse) AsHTML(payload string) *ModificationCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ModificationCommandSummaryActionResponse) AsBytes(payload []byte) *ModificationCommandSummaryActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ModificationCommandSummaryActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ModificationCommandSummaryActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ModificationCommandSummaryActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type ModificationCommandSummaryActionRequestSig = func(c ModificationCommandSummaryActionRequest) (*ModificationCommandSummaryActionResponse, error)

/**
 * Query parameters for Modification command summaryAction
 */
// Query wrapper with private fields
type ModificationCommandSummaryActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModificationCommandSummaryActionQueryFromString(rawQuery string) ModificationCommandSummaryActionQuery {
	v := ModificationCommandSummaryActionQuery{}
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
func ModificationCommandSummaryActionQueryFromHttp(r *http.Request) ModificationCommandSummaryActionQuery {
	return ModificationCommandSummaryActionQueryFromString(r.URL.RawQuery)
}
func (q ModificationCommandSummaryActionQuery) Values() url.Values {
	return q.values
}
func (q ModificationCommandSummaryActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModificationCommandSummaryActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModificationCommandSummaryActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModificationCommandSummaryActionRequest struct {
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

func ModificationCommandSummaryActionClientCreateUrl(
	req ModificationCommandSummaryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ModificationCommandSummaryActionMeta()
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
func ModificationCommandSummaryActionClientExecuteTyped(httpReq *http.Request) (*ModificationCommandSummaryActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ModificationCommandSummaryActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ModificationCommandSummaryActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ModificationCommandSummaryActionResponse{Payload: result}, err
	}
	return &ModificationCommandSummaryActionResponse{Payload: result}, nil
}
func ModificationCommandSummaryActionClientBuildRequest(req ModificationCommandSummaryActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ModificationCommandSummaryActionMeta()
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
func ModificationCommandSummaryActionCall(
	req ModificationCommandSummaryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModificationCommandSummaryActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ModificationCommandSummaryActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ModificationCommandSummaryActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ModificationCommandSummaryActionClientExecuteTyped(r)
}

// ModificationCommandSummaryActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the ModificationCommandSummaryAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *ModificationCommandSummaryActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func ModificationCommandSummaryActionHttpHandler(
	handler func(c ModificationCommandSummaryActionRequest) (*ModificationCommandSummaryActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := ModificationCommandSummaryActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := ModificationCommandSummaryActionRequest{
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

// ModificationCommandSummaryActionHttp is a high-level convenience wrapper around
// ModificationCommandSummaryActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func ModificationCommandSummaryActionHttp(
	mux *http.ServeMux,
	handler func(c ModificationCommandSummaryActionRequest) (*ModificationCommandSummaryActionResponse, error),
) {
	method, pattern, h := ModificationCommandSummaryActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
