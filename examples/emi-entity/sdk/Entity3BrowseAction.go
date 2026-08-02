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
* Action to communicate with the action Entity3BrowseAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity3BrowseAction
func Entity3BrowseAction(c Entity3BrowseActionRequest) (*Entity3BrowseActionResponse, error) {
	return &Entity3BrowseActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity3BrowseActionMeta() struct {
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
		Name:        "Entity3BrowseAction",
		CliName:     "entity3-browse-action",
		CliShort:    "entity3-b",
		URL:         "/entity3/browse",
		Method:      "GET",
		Description: `Returns "entity3" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).`,
	}
}

type Entity3BrowseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity3BrowseActionResponse) SetContentType(contentType string) *Entity3BrowseActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity3BrowseActionResponse) AsStream(r io.Reader, contentType string) *Entity3BrowseActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity3BrowseActionResponse) AsJSON(payload any) *Entity3BrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity3BrowseActionResponse) WithIdeal(payload Entity3OptionalDto) *Entity3BrowseActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity3BrowseActionResponse) AsIdeal() (*Entity3OptionalDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res Entity3OptionalDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *Entity3BrowseActionResponse) AsHTML(payload string) *Entity3BrowseActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity3BrowseActionResponse) AsBytes(payload []byte) *Entity3BrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity3BrowseActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity3BrowseActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity3BrowseActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity3BrowseActionRequestSig = func(c Entity3BrowseActionRequest) (*Entity3BrowseActionResponse, error)

/**
 * Query parameters for Entity3BrowseAction
 */
// Query wrapper with private fields
type Entity3BrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func Entity3BrowseActionQueryFromString(rawQuery string) Entity3BrowseActionQuery {
	v := Entity3BrowseActionQuery{}
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
func Entity3BrowseActionQueryFromHttp(r *http.Request) Entity3BrowseActionQuery {
	return Entity3BrowseActionQueryFromString(r.URL.RawQuery)
}
func (q Entity3BrowseActionQuery) Values() url.Values {
	return q.values
}
func (q Entity3BrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity3BrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity3BrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity3BrowseActionRequest struct {
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
func (x Entity3BrowseActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity3BrowseActionClientCreateUrl(
	req Entity3BrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity3BrowseActionMeta()
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
func Entity3BrowseActionClientExecuteTyped(httpReq *http.Request) (*Entity3BrowseActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity3BrowseActionResponse
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
func Entity3BrowseActionClientBuildRequest(req Entity3BrowseActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity3BrowseActionMeta()
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
func Entity3BrowseActionCall(
	req Entity3BrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity3BrowseActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity3BrowseActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity3BrowseActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity3BrowseActionClientExecuteTyped(r)
}
func GetEntity3BrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
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

// Entity3BrowseActionQueryFromCli extracts and casts query parameters the same way
// Entity3BrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func Entity3BrowseActionQueryFromCli(c *cli.Command) Entity3BrowseActionQuery {
	data := Entity3BrowseActionQuery{}
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
func (x Entity3BrowseActionRequest) IsCli() bool {
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

// Entity3BrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity3BrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity3BrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEntity3BrowseActionQueryCliFlags(""))...)
	return flags
}

// Entity3BrowseActionCliHandler builds a full *cli.Command for the
// Entity3BrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity3BrowseActionRequest the same way
// Entity3BrowseActionHandler (Gin) and Entity3BrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity3BrowseActionCliHandler(
	handler func(c Entity3BrowseActionRequest) (*Entity3BrowseActionResponse, error),
) *cli.Command {
	meta := Entity3BrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity3BrowseActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity3BrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = Entity3BrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity3BrowseActionCli is a high-level convenience wrapper around
// Entity3BrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity3BrowseActionGin
// registers a route on a Gin engine.
func Entity3BrowseActionCli(
	app *cli.Command,
	handler func(c Entity3BrowseActionRequest) (*Entity3BrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity3BrowseActionCliHandler(handler))
}

// Entity3BrowseActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity3BrowseAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity3BrowseActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity3BrowseActionHttpHandler(
	handler func(c Entity3BrowseActionRequest) (*Entity3BrowseActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity3BrowseActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := Entity3BrowseActionRequest{
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

// Entity3BrowseActionHttp is a high-level convenience wrapper around
// Entity3BrowseActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity3BrowseActionHttp(
	mux *http.ServeMux,
	handler func(c Entity3BrowseActionRequest) (*Entity3BrowseActionResponse, error),
) {
	method, pattern, h := Entity3BrowseActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
