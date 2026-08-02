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
)

/**
* Action to communicate with the action Entity3AwareDeletePreviewAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity3AwareDeletePreviewAction
func Entity3AwareDeletePreviewAction(c Entity3AwareDeletePreviewActionRequest) (*Entity3AwareDeletePreviewActionResponse, error) {
	return &Entity3AwareDeletePreviewActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity3AwareDeletePreviewActionMeta() struct {
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
		Name:        "Entity3AwareDeletePreviewAction",
		CliName:     "entity3-aware-delete-preview-action",
		CliShort:    "entity3-dp",
		URL:         "/entity3/delete-preview",
		Method:      "GET",
		Description: `Reports what deleting the given "entity3" uniqueIds would affect, without deleting anything.`,
	}
}

// The base class definition for entity3AwareDeletePreviewActionRes
type Entity3AwareDeletePreviewActionRes struct {
	Message  string                                                  `json:"message" yaml:"message"`
	Affected emigo.Array[Entity3AwareDeletePreviewActionResAffected] `json:"affected" yaml:"affected"`
}

// The base class definition for affected
type Entity3AwareDeletePreviewActionResAffected struct {
	Relation string `json:"relation" yaml:"relation"`
	Count    int64  `json:"count" yaml:"count"`
}

func (x *Entity3AwareDeletePreviewActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity3AwareDeletePreviewActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "message",
			Type: "string",
		},
		{
			Name: prefix + "affected",
			Type: "array",
		},
	}
}
func CastEntity3AwareDeletePreviewActionResFromCli(c emigo.CliCastable) Entity3AwareDeletePreviewActionRes {
	data := Entity3AwareDeletePreviewActionRes{}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("affected") {
		data.Affected = emigo.CapturePossibleArray(CastEntity3AwareDeletePreviewActionResAffectedFromCli, "affected", c)
	}
	return data
}
func GetEntity3AwareDeletePreviewActionResAffectedCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "relation",
			Type: "string",
		},
		{
			Name: prefix + "count",
			Type: "int64",
		},
	}
}
func CastEntity3AwareDeletePreviewActionResAffectedFromCli(c emigo.CliCastable) Entity3AwareDeletePreviewActionResAffected {
	data := Entity3AwareDeletePreviewActionResAffected{}
	if c.IsSet("relation") {
		data.Relation = c.String("relation")
	}
	if c.IsSet("count") {
		data.Count = int64(c.Int64("count"))
	}
	return data
}

type Entity3AwareDeletePreviewActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity3AwareDeletePreviewActionResponse) SetContentType(contentType string) *Entity3AwareDeletePreviewActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity3AwareDeletePreviewActionResponse) AsStream(r io.Reader, contentType string) *Entity3AwareDeletePreviewActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity3AwareDeletePreviewActionResponse) AsJSON(payload any) *Entity3AwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity3AwareDeletePreviewActionResponse) WithIdeal(payload Entity3AwareDeletePreviewActionRes) *Entity3AwareDeletePreviewActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity3AwareDeletePreviewActionResponse) AsIdeal() (*Entity3AwareDeletePreviewActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res Entity3AwareDeletePreviewActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *Entity3AwareDeletePreviewActionResponse) AsHTML(payload string) *Entity3AwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity3AwareDeletePreviewActionResponse) AsBytes(payload []byte) *Entity3AwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity3AwareDeletePreviewActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity3AwareDeletePreviewActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity3AwareDeletePreviewActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity3AwareDeletePreviewActionRequestSig = func(c Entity3AwareDeletePreviewActionRequest) (*Entity3AwareDeletePreviewActionResponse, error)

/**
 * Query parameters for Entity3AwareDeletePreviewAction
 */
// Query wrapper with private fields
type Entity3AwareDeletePreviewActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	UniqueIds []string `json:"uniqueIds"`
}

func Entity3AwareDeletePreviewActionQueryFromString(rawQuery string) Entity3AwareDeletePreviewActionQuery {
	v := Entity3AwareDeletePreviewActionQuery{}
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
func Entity3AwareDeletePreviewActionQueryFromHttp(r *http.Request) Entity3AwareDeletePreviewActionQuery {
	return Entity3AwareDeletePreviewActionQueryFromString(r.URL.RawQuery)
}
func (q Entity3AwareDeletePreviewActionQuery) Values() url.Values {
	return q.values
}
func (q Entity3AwareDeletePreviewActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity3AwareDeletePreviewActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity3AwareDeletePreviewActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity3AwareDeletePreviewActionRequest struct {
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
func (x Entity3AwareDeletePreviewActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity3AwareDeletePreviewActionClientCreateUrl(
	req Entity3AwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity3AwareDeletePreviewActionMeta()
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
func Entity3AwareDeletePreviewActionClientExecuteTyped(httpReq *http.Request) (*Entity3AwareDeletePreviewActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity3AwareDeletePreviewActionResponse
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
func Entity3AwareDeletePreviewActionClientBuildRequest(req Entity3AwareDeletePreviewActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity3AwareDeletePreviewActionMeta()
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
func Entity3AwareDeletePreviewActionCall(
	req Entity3AwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity3AwareDeletePreviewActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity3AwareDeletePreviewActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity3AwareDeletePreviewActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity3AwareDeletePreviewActionClientExecuteTyped(r)
}
func GetEntity3AwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// Entity3AwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// Entity3AwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func Entity3AwareDeletePreviewActionQueryFromCli(c *cli.Command) Entity3AwareDeletePreviewActionQuery {
	data := Entity3AwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x Entity3AwareDeletePreviewActionRequest) IsCli() bool {
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

// Entity3AwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity3AwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity3AwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEntity3AwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// Entity3AwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// Entity3AwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity3AwareDeletePreviewActionRequest the same way
// Entity3AwareDeletePreviewActionHandler (Gin) and Entity3AwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity3AwareDeletePreviewActionCliHandler(
	handler func(c Entity3AwareDeletePreviewActionRequest) (*Entity3AwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := Entity3AwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity3AwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity3AwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = Entity3AwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity3AwareDeletePreviewActionCli is a high-level convenience wrapper around
// Entity3AwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity3AwareDeletePreviewActionGin
// registers a route on a Gin engine.
func Entity3AwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c Entity3AwareDeletePreviewActionRequest) (*Entity3AwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity3AwareDeletePreviewActionCliHandler(handler))
}

// Entity3AwareDeletePreviewActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity3AwareDeletePreviewAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity3AwareDeletePreviewActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity3AwareDeletePreviewActionHttpHandler(
	handler func(c Entity3AwareDeletePreviewActionRequest) (*Entity3AwareDeletePreviewActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity3AwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := Entity3AwareDeletePreviewActionRequest{
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

// Entity3AwareDeletePreviewActionHttp is a high-level convenience wrapper around
// Entity3AwareDeletePreviewActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity3AwareDeletePreviewActionHttp(
	mux *http.ServeMux,
	handler func(c Entity3AwareDeletePreviewActionRequest) (*Entity3AwareDeletePreviewActionResponse, error),
) {
	method, pattern, h := Entity3AwareDeletePreviewActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
