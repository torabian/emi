package external

import (
	"context"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

/**
* Action to communicate with the action Entity4BrowseAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity4BrowseAction
func Entity4BrowseAction(c Entity4BrowseActionRequest) (*Entity4BrowseActionResponse, error) {
	return &Entity4BrowseActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity4BrowseActionMeta() struct {
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
		Name:        "Entity4BrowseAction",
		CliName:     "entity4-browse-action",
		URL:         "/entity4/browse",
		Method:      "GET",
		Description: `Returns "entity4" rows matching a filter, sorted/paged (see emigorm.QueryDSL).`,
	}
}

type Entity4BrowseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity4BrowseActionResponse) SetContentType(contentType string) *Entity4BrowseActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity4BrowseActionResponse) AsStream(r io.Reader, contentType string) *Entity4BrowseActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity4BrowseActionResponse) AsJSON(payload any) *Entity4BrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity4BrowseActionResponse) WithIdeal(payload Entity4OptionalDto) *Entity4BrowseActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity4BrowseActionResponse) AsIdeal() (*Entity4OptionalDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res Entity4OptionalDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *Entity4BrowseActionResponse) AsHTML(payload string) *Entity4BrowseActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity4BrowseActionResponse) AsBytes(payload []byte) *Entity4BrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity4BrowseActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity4BrowseActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity4BrowseActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity4BrowseActionRequestSig = func(c Entity4BrowseActionRequest) (*Entity4BrowseActionResponse, error)

/**
 * Query parameters for Entity4BrowseAction
 */
// Query wrapper with private fields
type Entity4BrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func Entity4BrowseActionQueryFromString(rawQuery string) Entity4BrowseActionQuery {
	v := Entity4BrowseActionQuery{}
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
func Entity4BrowseActionQueryFromHttp(r *http.Request) Entity4BrowseActionQuery {
	return Entity4BrowseActionQueryFromString(r.URL.RawQuery)
}
func (q Entity4BrowseActionQuery) Values() url.Values {
	return q.values
}
func (q Entity4BrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity4BrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity4BrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity4BrowseActionRequest struct {
	Body        interface{}
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
func (x Entity4BrowseActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity4BrowseActionClientCreateUrl(
	req Entity4BrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity4BrowseActionMeta()
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
func Entity4BrowseActionClientExecuteTyped(httpReq *http.Request) (*Entity4BrowseActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity4BrowseActionResponse
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
func Entity4BrowseActionClientBuildRequest(req Entity4BrowseActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity4BrowseActionMeta()
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
func Entity4BrowseActionCall(
	req Entity4BrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity4BrowseActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity4BrowseActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity4BrowseActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity4BrowseActionClientExecuteTyped(r)
}
func GetEntity4BrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-filter",
			Type: "string",
		},
		{
			Name: prefix + "qs-sort",
			Type: "string",
		},
		{
			Name: prefix + "qs-start-index",
			Type: "int",
		},
		{
			Name: prefix + "qs-items-per-page",
			Type: "int",
		},
		{
			Name: prefix + "qs-cursor",
			Type: "string",
		},
	}
}

// Entity4BrowseActionQueryFromCli extracts and casts query parameters the same way
// Entity4BrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func Entity4BrowseActionQueryFromCli(c *cli.Command) Entity4BrowseActionQuery {
	data := Entity4BrowseActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-filter") {
		data.Filter = c.String("qs-filter")
		values.Set("filter", data.Filter)
	}
	if c.IsSet("qs-sort") {
		data.Sort = c.String("qs-sort")
		values.Set("sort", data.Sort)
	}
	if c.IsSet("qs-start-index") {
		data.StartIndex = int(c.Int64("qs-start-index"))
		values.Set("startIndex", strconv.FormatInt(int64(data.StartIndex), 10))
	}
	if c.IsSet("qs-items-per-page") {
		data.ItemsPerPage = int(c.Int64("qs-items-per-page"))
		values.Set("itemsPerPage", strconv.FormatInt(int64(data.ItemsPerPage), 10))
	}
	if c.IsSet("qs-cursor") {
		data.Cursor = c.String("qs-cursor")
		values.Set("cursor", data.Cursor)
	}
	data.SetValues(values)
	return data
}
func (x Entity4BrowseActionRequest) IsCli() bool {
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

// Entity4BrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity4BrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity4BrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEntity4BrowseActionQueryCliFlags(""))...)
	return flags
}

// Entity4BrowseActionCliHandler builds a full *cli.Command for the
// Entity4BrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity4BrowseActionRequest the same way
// Entity4BrowseActionHandler (Gin) and Entity4BrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity4BrowseActionCliHandler(
	handler func(c Entity4BrowseActionRequest) (*Entity4BrowseActionResponse, error),
) *cli.Command {
	meta := Entity4BrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity4BrowseActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity4BrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = Entity4BrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity4BrowseActionCli is a high-level convenience wrapper around
// Entity4BrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity4BrowseActionGin
// registers a route on a Gin engine.
func Entity4BrowseActionCli(
	app *cli.Command,
	handler func(c Entity4BrowseActionRequest) (*Entity4BrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity4BrowseActionCliHandler(handler))
}

// Entity4BrowseActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity4BrowseAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity4BrowseActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity4BrowseActionHttpHandler(
	handler func(c Entity4BrowseActionRequest) (*Entity4BrowseActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity4BrowseActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := Entity4BrowseActionRequest{
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

// Entity4BrowseActionHttp is a high-level convenience wrapper around
// Entity4BrowseActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity4BrowseActionHttp(
	mux *http.ServeMux,
	handler func(c Entity4BrowseActionRequest) (*Entity4BrowseActionResponse, error),
) {
	method, pattern, h := Entity4BrowseActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
