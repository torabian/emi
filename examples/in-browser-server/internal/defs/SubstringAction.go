package defs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action SubstringAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of SubstringAction
func SubstringAction(c SubstringActionRequest) (*SubstringActionResponse, error) {
	return &SubstringActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func SubstringActionMeta() struct {
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
		Name:        "SubstringAction",
		CliName:     "substring-action",
		URL:         "/",
		Method:      "POST",
		Description: ``,
	}
}

// The base class definition for substringActionReq
type SubstringActionReq struct {
	// The string you want to do substring
	Input string `json:"input" yaml:"input"`
	// Start position
	Start int `json:"start" yaml:"start"`
	// End position
	End int `json:"end" yaml:"end"`
}

func (x *SubstringActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// The base class definition for substringActionRes
type SubstringActionRes struct {
	Result string `json:"result" yaml:"result"`
}

func (x *SubstringActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type SubstringActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *SubstringActionResponse) SetContentType(contentType string) *SubstringActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *SubstringActionResponse) AsStream(r io.Reader, contentType string) *SubstringActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *SubstringActionResponse) AsJSON(payload any) *SubstringActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *SubstringActionResponse) WithIdeal(payload SubstringActionRes) *SubstringActionResponse {
	x.Payload = payload
	return x
}
func (x *SubstringActionResponse) AsHTML(payload string) *SubstringActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *SubstringActionResponse) AsBytes(payload []byte) *SubstringActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x SubstringActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x SubstringActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x SubstringActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type SubstringActionRequestSig = func(c SubstringActionRequest) (*SubstringActionResponse, error)

/**
 * Query parameters for SubstringAction
 */
// Query wrapper with private fields
type SubstringActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func SubstringActionQueryFromString(rawQuery string) SubstringActionQuery {
	v := SubstringActionQuery{}
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
func SubstringActionQueryFromHttp(r *http.Request) SubstringActionQuery {
	return SubstringActionQueryFromString(r.URL.RawQuery)
}
func (q SubstringActionQuery) Values() url.Values {
	return q.values
}
func (q SubstringActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *SubstringActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *SubstringActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type SubstringActionRequest struct {
	Body        SubstringActionReq
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
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

// SubstringActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the SubstringAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *SubstringActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func SubstringActionHttpHandler(
	handler func(c SubstringActionRequest) (*SubstringActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := SubstringActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body SubstringActionReq
		if r.Body != nil {
			defer r.Body.Close()
			if data, _ := io.ReadAll(r.Body); len(data) > 0 {
				if err := json.Unmarshal(data, &body); err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON: " + err.Error()})
					return
				}
			}
		}
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := SubstringActionRequest{
			Body:        body,
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

// SubstringActionHttp is a high-level convenience wrapper around
// SubstringActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func SubstringActionHttp(
	mux *http.ServeMux,
	handler func(c SubstringActionRequest) (*SubstringActionResponse, error),
) {
	method, pattern, h := SubstringActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
