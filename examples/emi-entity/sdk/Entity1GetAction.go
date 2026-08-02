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
* Action to communicate with the action Entity1GetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity1GetAction
func Entity1GetAction(c Entity1GetActionRequest) (*Entity1GetActionResponse, error) {
	return &Entity1GetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity1GetActionMeta() struct {
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
		Name:        "Entity1GetAction",
		CliName:     "entity1-get-action",
		URL:         "/entity1/:uniqueId",
		Method:      "GET",
		Description: `Looks up a single "entity1" row by uniqueId.`,
	}
}

type Entity1GetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity1GetActionResponse) SetContentType(contentType string) *Entity1GetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity1GetActionResponse) AsStream(r io.Reader, contentType string) *Entity1GetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity1GetActionResponse) AsJSON(payload any) *Entity1GetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity1GetActionResponse) WithIdeal(payload Entity1Dto) *Entity1GetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity1GetActionResponse) AsIdeal() (*Entity1Dto, error) {
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
func (x *Entity1GetActionResponse) AsHTML(payload string) *Entity1GetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity1GetActionResponse) AsBytes(payload []byte) *Entity1GetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity1GetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity1GetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity1GetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity1GetActionRequestSig = func(c Entity1GetActionRequest) (*Entity1GetActionResponse, error)

/**
 * Path parameters for Entity1GetAction
 */
type Entity1GetActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func Entity1GetActionPathParameterApply(params Entity1GetActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func Entity1GetActionPathParameterFromFn(fn func(key string) string) Entity1GetActionPathParameter {
	res := Entity1GetActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for Entity1GetAction
 */
// Query wrapper with private fields
type Entity1GetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func Entity1GetActionQueryFromString(rawQuery string) Entity1GetActionQuery {
	v := Entity1GetActionQuery{}
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
func Entity1GetActionQueryFromHttp(r *http.Request) Entity1GetActionQuery {
	return Entity1GetActionQueryFromString(r.URL.RawQuery)
}
func (q Entity1GetActionQuery) Values() url.Values {
	return q.values
}
func (q Entity1GetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity1GetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity1GetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity1GetActionRequest struct {
	Body        interface{}
	Params      Entity1GetActionPathParameter
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
func (x Entity1GetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity1GetActionClientCreateUrl(
	req Entity1GetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity1GetActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = Entity1GetActionPathParameterApply(req.Params, urlAddr)
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
func Entity1GetActionClientExecuteTyped(httpReq *http.Request) (*Entity1GetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity1GetActionResponse
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
func Entity1GetActionClientBuildRequest(req Entity1GetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity1GetActionMeta()
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
func Entity1GetActionCall(
	req Entity1GetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity1GetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity1GetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity1GetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity1GetActionClientExecuteTyped(r)
}
func GetEntity1GetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func Entity1GetActionPathParameterFromCli(c *cli.Command) Entity1GetActionPathParameter {
	return Entity1GetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x Entity1GetActionRequest) IsCli() bool {
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

// Entity1GetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity1GetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity1GetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEntity1GetActionPathParameterCliFlags(""))...)
	return flags
}

// Entity1GetActionCliHandler builds a full *cli.Command for the
// Entity1GetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity1GetActionRequest the same way
// Entity1GetActionHandler (Gin) and Entity1GetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity1GetActionCliHandler(
	handler func(c Entity1GetActionRequest) (*Entity1GetActionResponse, error),
) *cli.Command {
	meta := Entity1GetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity1GetActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity1GetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      Entity1GetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity1GetActionCli is a high-level convenience wrapper around
// Entity1GetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity1GetActionGin
// registers a route on a Gin engine.
func Entity1GetActionCli(
	app *cli.Command,
	handler func(c Entity1GetActionRequest) (*Entity1GetActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity1GetActionCliHandler(handler))
}

// Entity1GetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity1GetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity1GetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity1GetActionHttpHandler(
	handler func(c Entity1GetActionRequest) (*Entity1GetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity1GetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := Entity1GetActionRequest{
			Body: nil,
			Params: Entity1GetActionPathParameterFromFn(func(key string) string {
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

// Entity1GetActionHttp is a high-level convenience wrapper around
// Entity1GetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity1GetActionHttp(
	mux *http.ServeMux,
	handler func(c Entity1GetActionRequest) (*Entity1GetActionResponse, error),
) {
	method, pattern, h := Entity1GetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
