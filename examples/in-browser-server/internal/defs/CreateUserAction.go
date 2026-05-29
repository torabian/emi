package defs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action CreateUserAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CreateUserAction
func CreateUserAction(c CreateUserActionRequest) (*CreateUserActionResponse, error) {
	return &CreateUserActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CreateUserActionMeta() struct {
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
		Name:        "CreateUserAction",
		CliName:     "create-user-action",
		URL:         "/users",
		Method:      "POST",
		Description: `Insert a new user row and return the created record.`,
	}
}

// The base class definition for createUserActionReq
type CreateUserActionReq struct {
	// User's first name
	FirstName string `json:"firstName" yaml:"firstName"`
	// User's last name
	LastName string `json:"lastName" yaml:"lastName"`
	// User's birth date, ISO format YYYY-MM-DD
	BirthDate string `json:"birthDate" yaml:"birthDate"`
}

func (x *CreateUserActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// The base class definition for createUserActionRes
type CreateUserActionRes struct {
	Id        int    `json:"id" yaml:"id"`
	FirstName string `json:"firstName" yaml:"firstName"`
	LastName  string `json:"lastName" yaml:"lastName"`
	BirthDate string `json:"birthDate" yaml:"birthDate"`
}

func (x *CreateUserActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CreateUserActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CreateUserActionResponse) SetContentType(contentType string) *CreateUserActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CreateUserActionResponse) AsStream(r io.Reader, contentType string) *CreateUserActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CreateUserActionResponse) AsJSON(payload any) *CreateUserActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CreateUserActionResponse) WithIdeal(payload CreateUserActionRes) *CreateUserActionResponse {
	x.Payload = payload
	return x
}
func (x *CreateUserActionResponse) AsHTML(payload string) *CreateUserActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CreateUserActionResponse) AsBytes(payload []byte) *CreateUserActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CreateUserActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CreateUserActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CreateUserActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type CreateUserActionRequestSig = func(c CreateUserActionRequest) (*CreateUserActionResponse, error)

/**
 * Query parameters for CreateUserAction
 */
// Query wrapper with private fields
type CreateUserActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CreateUserActionQueryFromString(rawQuery string) CreateUserActionQuery {
	v := CreateUserActionQuery{}
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
func CreateUserActionQueryFromHttp(r *http.Request) CreateUserActionQuery {
	return CreateUserActionQueryFromString(r.URL.RawQuery)
}
func (q CreateUserActionQuery) Values() url.Values {
	return q.values
}
func (q CreateUserActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CreateUserActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CreateUserActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CreateUserActionRequest struct {
	Body        CreateUserActionReq
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

// CreateUserActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the CreateUserAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *CreateUserActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func CreateUserActionHttpHandler(
	handler func(c CreateUserActionRequest) (*CreateUserActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := CreateUserActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body CreateUserActionReq
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
		req := CreateUserActionRequest{
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

// CreateUserActionHttp is a high-level convenience wrapper around
// CreateUserActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func CreateUserActionHttp(
	mux *http.ServeMux,
	handler func(c CreateUserActionRequest) (*CreateUserActionResponse, error),
) {
	method, pattern, h := CreateUserActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
