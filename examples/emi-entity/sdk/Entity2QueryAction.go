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
* Action to communicate with the action Entity2QueryAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity2QueryAction
func Entity2QueryAction(c Entity2QueryActionRequest) (*Entity2QueryActionResponse, error) {
	return &Entity2QueryActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity2QueryActionMeta() struct {
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
		Name:        "Entity2QueryAction",
		CliName:     "entity2-query-action",
		URL:         "/entity2/query",
		Method:      "POST",
		Description: `Returns "entity2" rows matching a JSON-logic filter, sorted/paged (see emigorm.QueryDSL).`,
	}
}

// The base class definition for entity2QueryActionReq
type Entity2QueryActionReq struct {
	Filter       emigo.Nullable[string] `json:"filter" yaml:"filter"`
	Sort         emigo.Nullable[string] `json:"sort" yaml:"sort"`
	StartIndex   emigo.Nullable[int]    `json:"startIndex" yaml:"startIndex"`
	ItemsPerPage emigo.Nullable[int]    `json:"itemsPerPage" yaml:"itemsPerPage"`
}

func (x *Entity2QueryActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity2QueryActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "filter",
			Type: "string?",
		},
		{
			Name: prefix + "sort",
			Type: "string?",
		},
		{
			Name: prefix + "start-index",
			Type: "int?",
		},
		{
			Name: prefix + "items-per-page",
			Type: "int?",
		},
	}
}
func CastEntity2QueryActionReqFromCli(c emigo.CliCastable) Entity2QueryActionReq {
	data := Entity2QueryActionReq{}
	if c.IsSet("filter") {
		emigo.ParseNullable(c.String("filter"), &data.Filter)
	}
	if c.IsSet("sort") {
		emigo.ParseNullable(c.String("sort"), &data.Sort)
	}
	if c.IsSet("start-index") {
		emigo.ParseNullable(c.String("start-index"), &data.StartIndex)
	}
	if c.IsSet("items-per-page") {
		emigo.ParseNullable(c.String("items-per-page"), &data.ItemsPerPage)
	}
	return data
}

// The base class definition for entity2QueryActionRes
type Entity2QueryActionRes struct {
	Items emigo.Array[Entity2QueryActionResItems] `json:"items" yaml:"items"`
	Total int64                                   `json:"total" yaml:"total"`
}

// The base class definition for items
type Entity2QueryActionResItems struct {
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	Label2   string `json:"label2" yaml:"label2"`
}

func (x *Entity2QueryActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity2QueryActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
		{
			Name: prefix + "total",
			Type: "int64",
		},
	}
}
func CastEntity2QueryActionResFromCli(c emigo.CliCastable) Entity2QueryActionRes {
	data := Entity2QueryActionRes{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastEntity2QueryActionResItemsFromCli, "items", c)
	}
	if c.IsSet("total") {
		data.Total = int64(c.Int64("total"))
	}
	return data
}
func GetEntity2QueryActionResItemsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEntity2QueryActionResItemsFromCli(c emigo.CliCastable) Entity2QueryActionResItems {
	data := Entity2QueryActionResItems{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("label2") {
		data.Label2 = c.String("label2")
	}
	return data
}

type Entity2QueryActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity2QueryActionResponse) SetContentType(contentType string) *Entity2QueryActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity2QueryActionResponse) AsStream(r io.Reader, contentType string) *Entity2QueryActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity2QueryActionResponse) AsJSON(payload any) *Entity2QueryActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity2QueryActionResponse) WithIdeal(payload Entity2QueryActionRes) *Entity2QueryActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity2QueryActionResponse) AsIdeal() (*Entity2QueryActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res Entity2QueryActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *Entity2QueryActionResponse) AsHTML(payload string) *Entity2QueryActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity2QueryActionResponse) AsBytes(payload []byte) *Entity2QueryActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity2QueryActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity2QueryActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity2QueryActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity2QueryActionRequestSig = func(c Entity2QueryActionRequest) (*Entity2QueryActionResponse, error)

/**
 * Query parameters for Entity2QueryAction
 */
// Query wrapper with private fields
type Entity2QueryActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func Entity2QueryActionQueryFromString(rawQuery string) Entity2QueryActionQuery {
	v := Entity2QueryActionQuery{}
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
func Entity2QueryActionQueryFromHttp(r *http.Request) Entity2QueryActionQuery {
	return Entity2QueryActionQueryFromString(r.URL.RawQuery)
}
func (q Entity2QueryActionQuery) Values() url.Values {
	return q.values
}
func (q Entity2QueryActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity2QueryActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity2QueryActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity2QueryActionRequest struct {
	Body        Entity2QueryActionReq
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
func (x Entity2QueryActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity2QueryActionClientCreateUrl(
	req Entity2QueryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity2QueryActionMeta()
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
func Entity2QueryActionClientExecuteTyped(httpReq *http.Request) (*Entity2QueryActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity2QueryActionResponse
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
func Entity2QueryActionClientBuildRequest(req Entity2QueryActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity2QueryActionMeta()
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
func Entity2QueryActionCall(
	req Entity2QueryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity2QueryActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity2QueryActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity2QueryActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity2QueryActionClientExecuteTyped(r)
}
func (x Entity2QueryActionRequest) IsCli() bool {
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

// Entity2QueryActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity2QueryAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity2QueryActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEntity2QueryActionReqCliFlags(""))...)
	return flags
}

// Entity2QueryActionCliHandler builds a full *cli.Command for the
// Entity2QueryAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity2QueryActionRequest the same way
// Entity2QueryActionHandler (Gin) and Entity2QueryActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity2QueryActionCliHandler(
	handler func(c Entity2QueryActionRequest) (*Entity2QueryActionResponse, error),
) *cli.Command {
	meta := Entity2QueryActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity2QueryActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity2QueryActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastEntity2QueryActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity2QueryActionCli is a high-level convenience wrapper around
// Entity2QueryActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity2QueryActionGin
// registers a route on a Gin engine.
func Entity2QueryActionCli(
	app *cli.Command,
	handler func(c Entity2QueryActionRequest) (*Entity2QueryActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity2QueryActionCliHandler(handler))
}

// Entity2QueryActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity2QueryAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity2QueryActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity2QueryActionHttpHandler(
	handler func(c Entity2QueryActionRequest) (*Entity2QueryActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity2QueryActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body Entity2QueryActionReq
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
		req := Entity2QueryActionRequest{
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

// Entity2QueryActionHttp is a high-level convenience wrapper around
// Entity2QueryActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity2QueryActionHttp(
	mux *http.ServeMux,
	handler func(c Entity2QueryActionRequest) (*Entity2QueryActionResponse, error),
) {
	method, pattern, h := Entity2QueryActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
