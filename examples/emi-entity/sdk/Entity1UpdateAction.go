package external

import (
	"bytes"
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
* Action to communicate with the action Entity1UpdateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of Entity1UpdateAction
func Entity1UpdateAction(c Entity1UpdateActionRequest) (*Entity1UpdateActionResponse, error) {
	return &Entity1UpdateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func Entity1UpdateActionMeta() struct {
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
		Name:        "Entity1UpdateAction",
		CliName:     "entity1-update-action",
		CliShort:    "entity1-u",
		URL:         "/entity1/:uniqueId",
		Method:      "PATCH",
		Description: `Applies a partial update to a "entity1" row by uniqueId.`,
	}
}

type Entity1UpdateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *Entity1UpdateActionResponse) SetContentType(contentType string) *Entity1UpdateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *Entity1UpdateActionResponse) AsStream(r io.Reader, contentType string) *Entity1UpdateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *Entity1UpdateActionResponse) AsJSON(payload any) *Entity1UpdateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *Entity1UpdateActionResponse) WithIdeal(payload Entity1Dto) *Entity1UpdateActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *Entity1UpdateActionResponse) AsIdeal() (*Entity1Dto, error) {
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
func (x *Entity1UpdateActionResponse) AsHTML(payload string) *Entity1UpdateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *Entity1UpdateActionResponse) AsBytes(payload []byte) *Entity1UpdateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x Entity1UpdateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x Entity1UpdateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x Entity1UpdateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type Entity1UpdateActionRequestSig = func(c Entity1UpdateActionRequest) (*Entity1UpdateActionResponse, error)

/**
 * Path parameters for Entity1UpdateAction
 */
type Entity1UpdateActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func Entity1UpdateActionPathParameterApply(params Entity1UpdateActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func Entity1UpdateActionPathParameterFromFn(fn func(key string) string) Entity1UpdateActionPathParameter {
	res := Entity1UpdateActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for Entity1UpdateAction
 */
// Query wrapper with private fields
type Entity1UpdateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func Entity1UpdateActionQueryFromString(rawQuery string) Entity1UpdateActionQuery {
	v := Entity1UpdateActionQuery{}
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
func Entity1UpdateActionQueryFromHttp(r *http.Request) Entity1UpdateActionQuery {
	return Entity1UpdateActionQueryFromString(r.URL.RawQuery)
}
func (q Entity1UpdateActionQuery) Values() url.Values {
	return q.values
}
func (q Entity1UpdateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *Entity1UpdateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *Entity1UpdateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type Entity1UpdateActionRequest struct {
	Body        Entity1OptionalDto
	Params      Entity1UpdateActionPathParameter
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
func (x Entity1UpdateActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func Entity1UpdateActionClientCreateUrl(
	req Entity1UpdateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := Entity1UpdateActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = Entity1UpdateActionPathParameterApply(req.Params, urlAddr)
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
func Entity1UpdateActionClientExecuteTyped(httpReq *http.Request) (*Entity1UpdateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result Entity1UpdateActionResponse
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
func Entity1UpdateActionClientBuildRequest(req Entity1UpdateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := Entity1UpdateActionMeta()
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
func Entity1UpdateActionCall(
	req Entity1UpdateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*Entity1UpdateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := Entity1UpdateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := Entity1UpdateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return Entity1UpdateActionClientExecuteTyped(r)
}
func GetEntity1UpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func Entity1UpdateActionPathParameterFromCli(c *cli.Command) Entity1UpdateActionPathParameter {
	return Entity1UpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x Entity1UpdateActionRequest) IsCli() bool {
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

// Entity1UpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the Entity1UpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func Entity1UpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEntity1UpdateActionPathParameterCliFlags(""))...)
	return flags
}

// Entity1UpdateActionCliHandler builds a full *cli.Command for the
// Entity1UpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a Entity1UpdateActionRequest the same way
// Entity1UpdateActionHandler (Gin) and Entity1UpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func Entity1UpdateActionCliHandler(
	handler func(c Entity1UpdateActionRequest) (*Entity1UpdateActionResponse, error),
) *cli.Command {
	meta := Entity1UpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: Entity1UpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := Entity1UpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      Entity1UpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// Entity1UpdateActionCli is a high-level convenience wrapper around
// Entity1UpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way Entity1UpdateActionGin
// registers a route on a Gin engine.
func Entity1UpdateActionCli(
	app *cli.Command,
	handler func(c Entity1UpdateActionRequest) (*Entity1UpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, Entity1UpdateActionCliHandler(handler))
}

// Entity1UpdateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the Entity1UpdateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *Entity1UpdateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func Entity1UpdateActionHttpHandler(
	handler func(c Entity1UpdateActionRequest) (*Entity1UpdateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := Entity1UpdateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body Entity1OptionalDto
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
		req := Entity1UpdateActionRequest{
			Body: body,
			Params: Entity1UpdateActionPathParameterFromFn(func(key string) string {
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

// Entity1UpdateActionHttp is a high-level convenience wrapper around
// Entity1UpdateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func Entity1UpdateActionHttp(
	mux *http.ServeMux,
	handler func(c Entity1UpdateActionRequest) (*Entity1UpdateActionResponse, error),
) {
	method, pattern, h := Entity1UpdateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
