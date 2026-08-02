package external

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

/**
* Action to communicate with the action Entity3GetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity3GetAction
func Entity3GetAction(c Entity3GetActionRequest) (*Entity3GetActionResponse, error) {
	return &Entity3GetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity3GetActionMeta() struct {
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
		Name:        "Entity3GetAction",
		CliName:     "entity3-get-action",
		URL:         "/entity3/:id",
		Method:      "GET",
		Description: `Looks up a single "entity3" row by uniqueId.`,
	}
}

type Entity3GetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity3GetActionResponse) SetContentType(contentType string) *Entity3GetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity3GetActionResponse) AsStream(r io.Reader, contentType string) *Entity3GetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity3GetActionResponse) AsJSON(payload any) *Entity3GetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity3GetActionResponse) WithIdeal(payload Entity3Dto) *Entity3GetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity3GetActionResponse) AsIdeal() (*Entity3Dto, error) {
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
func (x *Entity3GetActionResponse) AsHTML(payload string) *Entity3GetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity3GetActionResponse) AsBytes(payload []byte) *Entity3GetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity3GetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity3GetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity3GetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity3GetActionRequestSig = func(c Entity3GetActionRequest) (*Entity3GetActionResponse, error)

/**
 * Path parameters for Entity3GetAction
 */
type Entity3GetActionPathParameter struct {
	Id string
}

// Converts a placeholder url, and applies the parameters to it.
func Entity3GetActionPathParameterApply(params Entity3GetActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":id", fmt.Sprintf("%v", params.Id))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func Entity3GetActionPathParameterFromFn(fn func(key string) string) Entity3GetActionPathParameter {
	res := Entity3GetActionPathParameter{}
	res.Id = fn("id")
	return res
}

/**
 * Query parameters for Entity3GetAction
 */
// Query wrapper with private fields
type Entity3GetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func Entity3GetActionQueryFromString(rawQuery string) Entity3GetActionQuery {
	v := Entity3GetActionQuery{}
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
func Entity3GetActionQueryFromHttp(r *http.Request) Entity3GetActionQuery {
	return Entity3GetActionQueryFromString(r.URL.RawQuery)
}
func (q Entity3GetActionQuery) Values() url.Values {
	return q.values
}
func (q Entity3GetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity3GetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity3GetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity3GetActionRequest struct {
	Body        interface{}
	Params      Entity3GetActionPathParameter
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
func (x Entity3GetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity3GetActionClientCreateUrl(
	req Entity3GetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity3GetActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = Entity3GetActionPathParameterApply(req.Params, urlAddr)
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
func Entity3GetActionClientExecuteTyped(httpReq *http.Request) (*Entity3GetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity3GetActionResponse
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
func Entity3GetActionClientBuildRequest(req Entity3GetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity3GetActionMeta()
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
func Entity3GetActionCall(
	req Entity3GetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity3GetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity3GetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity3GetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity3GetActionClientExecuteTyped(r)
}
func GetEntity3GetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-id",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func Entity3GetActionPathParameterFromCli(c *cli.Command) Entity3GetActionPathParameter {
	return Entity3GetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x Entity3GetActionRequest) IsCli() bool {
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

// Entity3GetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity3GetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity3GetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEntity3GetActionPathParameterCliFlags(""))...)
	return flags
}

// Entity3GetActionCliHandler builds a full *cli.Command for the
// Entity3GetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity3GetActionRequest the same way
// Entity3GetActionHandler (Gin) and Entity3GetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity3GetActionCliHandler(
	handler func(c Entity3GetActionRequest) (*Entity3GetActionResponse, error),
) *cli.Command {
	meta := Entity3GetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity3GetActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity3GetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      Entity3GetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity3GetActionCli is a high-level convenience wrapper around
// Entity3GetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity3GetActionGin
// registers a route on a Gin engine.
func Entity3GetActionCli(
	app *cli.Command,
	handler func(c Entity3GetActionRequest) (*Entity3GetActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity3GetActionCliHandler(handler))
}

// Entity3GetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity3GetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity3GetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity3GetActionHttpHandler(
	handler func(c Entity3GetActionRequest) (*Entity3GetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity3GetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := Entity3GetActionRequest{
			Body: nil,
			Params: Entity3GetActionPathParameterFromFn(func(key string) string {
				return r.PathValue(key)
			}),
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

// Entity3GetActionHttp is a high-level convenience wrapper around
// Entity3GetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity3GetActionHttp(
	mux *http.ServeMux,
	handler func(c Entity3GetActionRequest) (*Entity3GetActionResponse, error),
) {
	method, pattern, h := Entity3GetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
