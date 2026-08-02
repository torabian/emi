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
* Action to communicate with the action Entity3CreateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity3CreateAction
func Entity3CreateAction(c Entity3CreateActionRequest) (*Entity3CreateActionResponse, error) {
	return &Entity3CreateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity3CreateActionMeta() struct {
	Name        string
	CliName     string
	CliShort    string
	URL         string
	Method      string
	Description string
} {
	return struct {
		Name        string
		CliName     string
		CliShort    string
		URL         string
		Method      string
		Description string
	}{
		Name:        "Entity3CreateAction",
		CliName:     "entity3-create-action",
		CliShort:    "entity3-c",
		URL:         "/entity3",
		Method:      "POST",
		Description: `Creates a new "entity3" row.`,
	}
}

type Entity3CreateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity3CreateActionResponse) SetContentType(contentType string) *Entity3CreateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity3CreateActionResponse) AsStream(r io.Reader, contentType string) *Entity3CreateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity3CreateActionResponse) AsJSON(payload any) *Entity3CreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity3CreateActionResponse) WithIdeal(payload Entity3Dto) *Entity3CreateActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity3CreateActionResponse) AsIdeal() (*Entity3Dto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res Entity3Dto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *Entity3CreateActionResponse) AsHTML(payload string) *Entity3CreateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity3CreateActionResponse) AsBytes(payload []byte) *Entity3CreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity3CreateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity3CreateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity3CreateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity3CreateActionRequestSig = func(c Entity3CreateActionRequest) (*Entity3CreateActionResponse, error)

/**
 * Query parameters for Entity3CreateAction
 */
// Query wrapper with private fields
type Entity3CreateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func Entity3CreateActionQueryFromString(rawQuery string) Entity3CreateActionQuery {
	v := Entity3CreateActionQuery{}
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
func Entity3CreateActionQueryFromHttp(r *http.Request) Entity3CreateActionQuery {
	return Entity3CreateActionQueryFromString(r.URL.RawQuery)
}
func (q Entity3CreateActionQuery) Values() url.Values {
	return q.values
}
func (q Entity3CreateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity3CreateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity3CreateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity3CreateActionRequest struct {
	Body        Entity3Dto
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
func (x Entity3CreateActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity3CreateActionClientCreateUrl(
	req Entity3CreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity3CreateActionMeta()
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
func Entity3CreateActionClientExecuteTyped(httpReq *http.Request) (*Entity3CreateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity3CreateActionResponse
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
func Entity3CreateActionClientBuildRequest(req Entity3CreateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity3CreateActionMeta()
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
func Entity3CreateActionCall(
	req Entity3CreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity3CreateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity3CreateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity3CreateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity3CreateActionClientExecuteTyped(r)
}
func (x Entity3CreateActionRequest) IsCli() bool {
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

// Entity3CreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity3CreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity3CreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// Entity3CreateActionCliHandler builds a full *cli.Command for the
// Entity3CreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity3CreateActionRequest the same way
// Entity3CreateActionHandler (Gin) and Entity3CreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity3CreateActionCliHandler(
	handler func(c Entity3CreateActionRequest) (*Entity3CreateActionResponse, error),
) *cli.Command {
	meta := Entity3CreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity3CreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity3CreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity3CreateActionCli is a high-level convenience wrapper around
// Entity3CreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity3CreateActionGin
// registers a route on a Gin engine.
func Entity3CreateActionCli(
	app *cli.Command,
	handler func(c Entity3CreateActionRequest) (*Entity3CreateActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity3CreateActionCliHandler(handler))
}

// Entity3CreateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity3CreateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity3CreateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity3CreateActionHttpHandler(
	handler func(c Entity3CreateActionRequest) (*Entity3CreateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity3CreateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body Entity3Dto
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
		req := Entity3CreateActionRequest{
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

// Entity3CreateActionHttp is a high-level convenience wrapper around
// Entity3CreateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity3CreateActionHttp(
	mux *http.ServeMux,
	handler func(c Entity3CreateActionRequest) (*Entity3CreateActionResponse, error),
) {
	method, pattern, h := Entity3CreateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
