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
* Action to communicate with the action Entity1AwareDeleteAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity1AwareDeleteAction
func Entity1AwareDeleteAction(c Entity1AwareDeleteActionRequest) (*Entity1AwareDeleteActionResponse, error) {
	return &Entity1AwareDeleteActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity1AwareDeleteActionMeta() struct {
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
		Name:        "Entity1AwareDeleteAction",
		CliName:     "entity1-aware-delete-action",
		CliShort:    "entity1-d",
		URL:         "/entity1/delete",
		Method:      "POST",
		Description: `Deletes the given "entity1" uniqueIds, along with everything entity1AwareDeletePreview reports.`,
	}
}

// The base class definition for entity1AwareDeleteActionReq
type Entity1AwareDeleteActionReq struct {
	UniqueIds []string `json:"uniqueIds" yaml:"uniqueIds"`
}

func (x *Entity1AwareDeleteActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity1AwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastEntity1AwareDeleteActionReqFromCli(c emigo.CliCastable) Entity1AwareDeleteActionReq {
	data := Entity1AwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}

type Entity1AwareDeleteActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity1AwareDeleteActionResponse) SetContentType(contentType string) *Entity1AwareDeleteActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity1AwareDeleteActionResponse) AsStream(r io.Reader, contentType string) *Entity1AwareDeleteActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity1AwareDeleteActionResponse) AsJSON(payload any) *Entity1AwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *Entity1AwareDeleteActionResponse) AsHTML(payload string) *Entity1AwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity1AwareDeleteActionResponse) AsBytes(payload []byte) *Entity1AwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity1AwareDeleteActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity1AwareDeleteActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity1AwareDeleteActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity1AwareDeleteActionRequestSig = func(c Entity1AwareDeleteActionRequest) (*Entity1AwareDeleteActionResponse, error)

/**
 * Query parameters for Entity1AwareDeleteAction
 */
// Query wrapper with private fields
type Entity1AwareDeleteActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func Entity1AwareDeleteActionQueryFromString(rawQuery string) Entity1AwareDeleteActionQuery {
	v := Entity1AwareDeleteActionQuery{}
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
func Entity1AwareDeleteActionQueryFromHttp(r *http.Request) Entity1AwareDeleteActionQuery {
	return Entity1AwareDeleteActionQueryFromString(r.URL.RawQuery)
}
func (q Entity1AwareDeleteActionQuery) Values() url.Values {
	return q.values
}
func (q Entity1AwareDeleteActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity1AwareDeleteActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity1AwareDeleteActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity1AwareDeleteActionRequest struct {
	Body        Entity1AwareDeleteActionReq
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
func (x Entity1AwareDeleteActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity1AwareDeleteActionClientCreateUrl(
	req Entity1AwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity1AwareDeleteActionMeta()
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
func Entity1AwareDeleteActionClientExecuteTyped(httpReq *http.Request) (*Entity1AwareDeleteActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity1AwareDeleteActionResponse
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
func Entity1AwareDeleteActionClientBuildRequest(req Entity1AwareDeleteActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity1AwareDeleteActionMeta()
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
func Entity1AwareDeleteActionCall(
	req Entity1AwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity1AwareDeleteActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity1AwareDeleteActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity1AwareDeleteActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity1AwareDeleteActionClientExecuteTyped(r)
}
func (x Entity1AwareDeleteActionRequest) IsCli() bool {
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

// Entity1AwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity1AwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity1AwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEntity1AwareDeleteActionReqCliFlags(""))...)
	return flags
}

// Entity1AwareDeleteActionCliHandler builds a full *cli.Command for the
// Entity1AwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity1AwareDeleteActionRequest the same way
// Entity1AwareDeleteActionHandler (Gin) and Entity1AwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity1AwareDeleteActionCliHandler(
	handler func(c Entity1AwareDeleteActionRequest) (*Entity1AwareDeleteActionResponse, error),
) *cli.Command {
	meta := Entity1AwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity1AwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity1AwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastEntity1AwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity1AwareDeleteActionCli is a high-level convenience wrapper around
// Entity1AwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity1AwareDeleteActionGin
// registers a route on a Gin engine.
func Entity1AwareDeleteActionCli(
	app *cli.Command,
	handler func(c Entity1AwareDeleteActionRequest) (*Entity1AwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity1AwareDeleteActionCliHandler(handler))
}

// Entity1AwareDeleteActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity1AwareDeleteAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity1AwareDeleteActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity1AwareDeleteActionHttpHandler(
	handler func(c Entity1AwareDeleteActionRequest) (*Entity1AwareDeleteActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity1AwareDeleteActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body Entity1AwareDeleteActionReq
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
		req := Entity1AwareDeleteActionRequest{
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

// Entity1AwareDeleteActionHttp is a high-level convenience wrapper around
// Entity1AwareDeleteActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity1AwareDeleteActionHttp(
	mux *http.ServeMux,
	handler func(c Entity1AwareDeleteActionRequest) (*Entity1AwareDeleteActionResponse, error),
) {
	method, pattern, h := Entity1AwareDeleteActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
