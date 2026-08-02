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
* Action to communicate with the action Entity2BrowseAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity2BrowseAction
func Entity2BrowseAction(c Entity2BrowseActionRequest) (*Entity2BrowseActionResponse, error) {
	return &Entity2BrowseActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity2BrowseActionMeta() struct {
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
		Name:        "Entity2BrowseAction",
		CliName:     "entity2-browse-action",
		CliShort:    "entity2-b",
		URL:         "/entity2/browse",
		Method:      "GET",
		Description: `Returns "entity2" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).`,
	}
}

type Entity2BrowseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity2BrowseActionResponse) SetContentType(contentType string) *Entity2BrowseActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity2BrowseActionResponse) AsStream(r io.Reader, contentType string) *Entity2BrowseActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity2BrowseActionResponse) AsJSON(payload any) *Entity2BrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity2BrowseActionResponse) WithIdeal(payload Entity2OptionalDto) *Entity2BrowseActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity2BrowseActionResponse) AsIdeal() (*Entity2OptionalDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res Entity2OptionalDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *Entity2BrowseActionResponse) AsHTML(payload string) *Entity2BrowseActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity2BrowseActionResponse) AsBytes(payload []byte) *Entity2BrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity2BrowseActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity2BrowseActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity2BrowseActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity2BrowseActionRequestSig = func(c Entity2BrowseActionRequest) (*Entity2BrowseActionResponse, error)

/**
 * Query parameters for Entity2BrowseAction
 */
// Query wrapper with private fields
type Entity2BrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func Entity2BrowseActionQueryFromString(rawQuery string) Entity2BrowseActionQuery {
	v := Entity2BrowseActionQuery{}
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
func Entity2BrowseActionQueryFromHttp(r *http.Request) Entity2BrowseActionQuery {
	return Entity2BrowseActionQueryFromString(r.URL.RawQuery)
}
func (q Entity2BrowseActionQuery) Values() url.Values {
	return q.values
}
func (q Entity2BrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity2BrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity2BrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity2BrowseActionRequest struct {
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
func (x Entity2BrowseActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity2BrowseActionClientCreateUrl(
	req Entity2BrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity2BrowseActionMeta()
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
func Entity2BrowseActionClientExecuteTyped(httpReq *http.Request) (*Entity2BrowseActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity2BrowseActionResponse
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
func Entity2BrowseActionClientBuildRequest(req Entity2BrowseActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity2BrowseActionMeta()
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
func Entity2BrowseActionCall(
	req Entity2BrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity2BrowseActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity2BrowseActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity2BrowseActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity2BrowseActionClientExecuteTyped(r)
}
func GetEntity2BrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
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

// Entity2BrowseActionQueryFromCli extracts and casts query parameters the same way
// Entity2BrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func Entity2BrowseActionQueryFromCli(c *cli.Command) Entity2BrowseActionQuery {
	data := Entity2BrowseActionQuery{}
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
func (x Entity2BrowseActionRequest) IsCli() bool {
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

// Entity2BrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity2BrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity2BrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEntity2BrowseActionQueryCliFlags(""))...)
	return flags
}

// Entity2BrowseActionCliHandler builds a full *cli.Command for the
// Entity2BrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity2BrowseActionRequest the same way
// Entity2BrowseActionHandler (Gin) and Entity2BrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity2BrowseActionCliHandler(
	handler func(c Entity2BrowseActionRequest) (*Entity2BrowseActionResponse, error),
) *cli.Command {
	meta := Entity2BrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity2BrowseActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity2BrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = Entity2BrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity2BrowseActionCli is a high-level convenience wrapper around
// Entity2BrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity2BrowseActionGin
// registers a route on a Gin engine.
func Entity2BrowseActionCli(
	app *cli.Command,
	handler func(c Entity2BrowseActionRequest) (*Entity2BrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity2BrowseActionCliHandler(handler))
}

// Entity2BrowseActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity2BrowseAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity2BrowseActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity2BrowseActionHttpHandler(
	handler func(c Entity2BrowseActionRequest) (*Entity2BrowseActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity2BrowseActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := Entity2BrowseActionRequest{
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

// Entity2BrowseActionHttp is a high-level convenience wrapper around
// Entity2BrowseActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity2BrowseActionHttp(
	mux *http.ServeMux,
	handler func(c Entity2BrowseActionRequest) (*Entity2BrowseActionResponse, error),
) {
	method, pattern, h := Entity2BrowseActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
