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
* Action to communicate with the action Entity2CreateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity2CreateAction
func Entity2CreateAction(c Entity2CreateActionRequest) (*Entity2CreateActionResponse, error) {
	return &Entity2CreateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity2CreateActionMeta() struct {
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
		Name:        "Entity2CreateAction",
		CliName:     "entity2-create-action",
		URL:         "/entity2",
		Method:      "POST",
		Description: `Creates a new "entity2" row.`,
	}
}

// The base class definition for entity2CreateActionRes
type Entity2CreateActionRes struct {
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	Label2   string `json:"label2" yaml:"label2"`
}

func (x *Entity2CreateActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity2CreateActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "label2",
			Type: "string",
		},
	}
}
func CastEntity2CreateActionResFromCli(c emigo.CliCastable) Entity2CreateActionRes {
	data := Entity2CreateActionRes{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("label2") {
		data.Label2 = c.String("label2")
	}
	return data
}

type Entity2CreateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity2CreateActionResponse) SetContentType(contentType string) *Entity2CreateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity2CreateActionResponse) AsStream(r io.Reader, contentType string) *Entity2CreateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity2CreateActionResponse) AsJSON(payload any) *Entity2CreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity2CreateActionResponse) WithIdeal(payload Entity2CreateActionRes) *Entity2CreateActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity2CreateActionResponse) AsIdeal() (*Entity2CreateActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res Entity2CreateActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *Entity2CreateActionResponse) AsHTML(payload string) *Entity2CreateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity2CreateActionResponse) AsBytes(payload []byte) *Entity2CreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity2CreateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity2CreateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity2CreateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity2CreateActionRequestSig = func(c Entity2CreateActionRequest) (*Entity2CreateActionResponse, error)

/**
 * Query parameters for Entity2CreateAction
 */
// Query wrapper with private fields
type Entity2CreateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func Entity2CreateActionQueryFromString(rawQuery string) Entity2CreateActionQuery {
	v := Entity2CreateActionQuery{}
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
func Entity2CreateActionQueryFromHttp(r *http.Request) Entity2CreateActionQuery {
	return Entity2CreateActionQueryFromString(r.URL.RawQuery)
}
func (q Entity2CreateActionQuery) Values() url.Values {
	return q.values
}
func (q Entity2CreateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity2CreateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity2CreateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity2CreateActionRequest struct {
	Body        Entity2Dto
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
func (x Entity2CreateActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity2CreateActionClientCreateUrl(
	req Entity2CreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity2CreateActionMeta()
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
func Entity2CreateActionClientExecuteTyped(httpReq *http.Request) (*Entity2CreateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity2CreateActionResponse
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
func Entity2CreateActionClientBuildRequest(req Entity2CreateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity2CreateActionMeta()
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
func Entity2CreateActionCall(
	req Entity2CreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity2CreateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity2CreateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity2CreateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity2CreateActionClientExecuteTyped(r)
}
func (x Entity2CreateActionRequest) IsCli() bool {
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

// Entity2CreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity2CreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity2CreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// Entity2CreateActionCliHandler builds a full *cli.Command for the
// Entity2CreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity2CreateActionRequest the same way
// Entity2CreateActionHandler (Gin) and Entity2CreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity2CreateActionCliHandler(
	handler func(c Entity2CreateActionRequest) (*Entity2CreateActionResponse, error),
) *cli.Command {
	meta := Entity2CreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity2CreateActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity2CreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity2CreateActionCli is a high-level convenience wrapper around
// Entity2CreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity2CreateActionGin
// registers a route on a Gin engine.
func Entity2CreateActionCli(
	app *cli.Command,
	handler func(c Entity2CreateActionRequest) (*Entity2CreateActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity2CreateActionCliHandler(handler))
}

// Entity2CreateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity2CreateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity2CreateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity2CreateActionHttpHandler(
	handler func(c Entity2CreateActionRequest) (*Entity2CreateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity2CreateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body Entity2Dto
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
		req := Entity2CreateActionRequest{
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

// Entity2CreateActionHttp is a high-level convenience wrapper around
// Entity2CreateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity2CreateActionHttp(
	mux *http.ServeMux,
	handler func(c Entity2CreateActionRequest) (*Entity2CreateActionResponse, error),
) {
	method, pattern, h := Entity2CreateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
