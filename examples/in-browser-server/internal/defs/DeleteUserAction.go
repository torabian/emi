package defs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action DeleteUserAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of DeleteUserAction
func DeleteUserAction(c DeleteUserActionRequest) (*DeleteUserActionResponse, error) {
	return &DeleteUserActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func DeleteUserActionMeta() struct {
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
		Name:        "DeleteUserAction",
		CliName:     "delete-user-action",
		URL:         "/users",
		Method:      "DELETE",
		Description: `Delete a single user row by id.`,
	}
}

// The base class definition for deleteUserActionReq
type DeleteUserActionReq struct {
	// Id of the user to delete
	Id int `json:"id" yaml:"id"`
}

func (x *DeleteUserActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// The base class definition for deleteUserActionRes
type DeleteUserActionRes struct {
	Deleted bool `json:"deleted" yaml:"deleted"`
}

func (x *DeleteUserActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type DeleteUserActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *DeleteUserActionResponse) SetContentType(contentType string) *DeleteUserActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *DeleteUserActionResponse) AsStream(r io.Reader, contentType string) *DeleteUserActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *DeleteUserActionResponse) AsJSON(payload any) *DeleteUserActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *DeleteUserActionResponse) WithIdeal(payload DeleteUserActionRes) *DeleteUserActionResponse {
	x.Payload = payload
	return x
}
func (x *DeleteUserActionResponse) AsHTML(payload string) *DeleteUserActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *DeleteUserActionResponse) AsBytes(payload []byte) *DeleteUserActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x DeleteUserActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x DeleteUserActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x DeleteUserActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type DeleteUserActionRequestSig = func(c DeleteUserActionRequest) (*DeleteUserActionResponse, error)

/**
 * Query parameters for DeleteUserAction
 */
// Query wrapper with private fields
type DeleteUserActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func DeleteUserActionQueryFromString(rawQuery string) DeleteUserActionQuery {
	v := DeleteUserActionQuery{}
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
func DeleteUserActionQueryFromHttp(r *http.Request) DeleteUserActionQuery {
	return DeleteUserActionQueryFromString(r.URL.RawQuery)
}
func (q DeleteUserActionQuery) Values() url.Values {
	return q.values
}
func (q DeleteUserActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *DeleteUserActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *DeleteUserActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type DeleteUserActionRequest struct {
	Body        DeleteUserActionReq
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

// DeleteUserActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the DeleteUserAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *DeleteUserActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func DeleteUserActionHttpHandler(
	handler func(c DeleteUserActionRequest) (*DeleteUserActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := DeleteUserActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body DeleteUserActionReq
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
		req := DeleteUserActionRequest{
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

// DeleteUserActionHttp is a high-level convenience wrapper around
// DeleteUserActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func DeleteUserActionHttp(
	mux *http.ServeMux,
	handler func(c DeleteUserActionRequest) (*DeleteUserActionResponse, error),
) {
	method, pattern, h := DeleteUserActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
