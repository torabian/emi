package external

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action Entity1CreateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity1CreateAction
func Entity1CreateAction(c Entity1CreateActionRequest) (*Entity1CreateActionResponse, error) {
	return &Entity1CreateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity1CreateActionMeta() struct {
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
		Name:        "Entity1CreateAction",
		CliName:     "entity1-create-action",
		URL:         "/entity1",
		Method:      "POST",
		Description: `Creates a new "entity1" row.`,
	}
}

type Entity1CreateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity1CreateActionResponse) SetContentType(contentType string) *Entity1CreateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity1CreateActionResponse) AsStream(r io.Reader, contentType string) *Entity1CreateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity1CreateActionResponse) AsJSON(payload any) *Entity1CreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity1CreateActionResponse) WithIdeal(payload Entity1Dto) *Entity1CreateActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity1CreateActionResponse) AsIdeal() (*Entity1Dto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res Entity1Dto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *Entity1CreateActionResponse) AsHTML(payload string) *Entity1CreateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity1CreateActionResponse) AsBytes(payload []byte) *Entity1CreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity1CreateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity1CreateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity1CreateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity1CreateActionRequestSig = func(c Entity1CreateActionRequest) (*Entity1CreateActionResponse, error)

/**
 * Query parameters for Entity1CreateAction
 */
// Query wrapper with private fields
type Entity1CreateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func Entity1CreateActionQueryFromString(rawQuery string) Entity1CreateActionQuery {
	v := Entity1CreateActionQuery{}
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
func Entity1CreateActionQueryFromHttp(r *http.Request) Entity1CreateActionQuery {
	return Entity1CreateActionQueryFromString(r.URL.RawQuery)
}
func (q Entity1CreateActionQuery) Values() url.Values {
	return q.values
}
func (q Entity1CreateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity1CreateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity1CreateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity1CreateActionRequest struct {
	Body        Entity1Dto
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
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

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x Entity1CreateActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity1CreateActionClientCreateUrl(
	req Entity1CreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity1CreateActionMeta()
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
func Entity1CreateActionClientExecuteTyped(httpReq *http.Request) (*Entity1CreateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity1CreateActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &result, err
	}
	return &result, nil
}
func Entity1CreateActionClientBuildRequest(req Entity1CreateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity1CreateActionMeta()
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
func Entity1CreateActionCall(
	req Entity1CreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity1CreateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity1CreateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity1CreateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity1CreateActionClientExecuteTyped(r)
}
func (x Entity1CreateActionRequest) IsCli() bool {
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

// Entity1CreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity1CreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity1CreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// Entity1CreateActionCliHandler builds a full *cli.Command for the
// Entity1CreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity1CreateActionRequest the same way
// Entity1CreateActionHandler (Gin) and Entity1CreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity1CreateActionCliHandler(
	handler func(c Entity1CreateActionRequest) (*Entity1CreateActionResponse, error),
) *cli.Command {
	meta := Entity1CreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity1CreateActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity1CreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity1CreateActionCli is a high-level convenience wrapper around
// Entity1CreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity1CreateActionGin
// registers a route on a Gin engine.
func Entity1CreateActionCli(
	app *cli.Command,
	handler func(c Entity1CreateActionRequest) (*Entity1CreateActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity1CreateActionCliHandler(handler))
}

// Entity1CreateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity1CreateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity1CreateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity1CreateActionHttpHandler(
	handler func(c Entity1CreateActionRequest) (*Entity1CreateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity1CreateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body Entity1Dto
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
		req := Entity1CreateActionRequest{
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

// Entity1CreateActionHttp is a high-level convenience wrapper around
// Entity1CreateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity1CreateActionHttp(
	mux *http.ServeMux,
	handler func(c Entity1CreateActionRequest) (*Entity1CreateActionResponse, error),
) {
	method, pattern, h := Entity1CreateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
