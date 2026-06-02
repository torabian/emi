package external

import (
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction(c CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse, error) {
	return &CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta() struct {
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
		Name:        "CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction",
		CliName:     "check the processing status of a post or patch request-action",
		URL:         "https://api.{environment}/sale/product-offers/{offerId}/operations/{operationId}",
		Method:      "GET",
		Description: `The URI for the resource given by Location header of POST /sale/product-offers and PATCH /sale/product-offers/{offerId}. Use this resource to check processing status of a POST or PATCH request. Read more: PL / EN.`,
	}
}
func GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "offer",
			Type:     "object",
			Children: GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOfferCliFlags("offer-"),
		},
		{
			Name:     prefix + "operation",
			Type:     "object",
			Children: GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperationCliFlags("operation-"),
		},
	}
}
func CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResFromCli(c emigo.CliCastable) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes {
	data := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes{}
	if c.IsSet("offer") {
		data.Offer = CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOfferFromCli(c)
	}
	if c.IsSet("operation") {
		data.Operation = CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperationFromCli(c)
	}
	return data
}

// The base class definition for checkTheProcessingStatusOfAPOSTOrPATCHRequestActionRes
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes struct {
	Offer     CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer     `json:"offer" yaml:"offer"`
	Operation CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation `json:"operation" yaml:"operation"`
}

func GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOfferCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOfferFromCli(c emigo.CliCastable) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer {
	data := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for offer
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer struct {
	Id string `json:"id" yaml:"id"`
}

func GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "status",
			Type: "string",
		},
		{
			Name: prefix + "started-at",
			Type: "string",
		},
	}
}
func CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperationFromCli(c emigo.CliCastable) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation {
	data := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("started-at") {
		data.StartedAt = c.String("started-at")
	}
	return data
}

// The base class definition for operation
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation struct {
	Id        string `json:"id" yaml:"id"`
	Status    string `json:"status" yaml:"status"`
	StartedAt string `json:"startedAt" yaml:"startedAt"`
}

func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) SetContentType(contentType string) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) AsStream(r io.Reader, contentType string) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) AsJSON(payload any) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) WithIdeal(payload CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	x.Payload = payload
	return x
}
func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) AsHTML(payload string) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) AsBytes(payload []byte) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequestSig = func(c CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse, error)

/**
 * Query parameters for Check the processing status of a POST or PATCH requestAction
 */
// Query wrapper with private fields
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromString(rawQuery string) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery {
	v := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery{}
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
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromHttp(r *http.Request) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery {
	return CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromString(r.URL.RawQuery)
}
func (q CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) Values() url.Values {
	return q.values
}
func (q CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest struct {
	Body        interface{}
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	// Now it's interface, so the code gen doesn't depend on the instance
	// or gin package. Make sure you cast is later into *gin.Context, or whatever
	// your framework is passing when creating a request.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	GinCtx interface{}
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

func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionClientCreateUrl(
	req CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
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
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionClientExecuteTyped(httpReq *http.Request) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse{Payload: result}, err
	}
	return &CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse{Payload: result}, nil
}
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionClientBuildRequest(req CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
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
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionCall(
	req CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionClientExecuteTyped(r)
}
func (x CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest) IsCli() bool {
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

// CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHttpHandler(
	handler func(c CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest{
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

// CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHttp is a high-level convenience wrapper around
// CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHttp(
	mux *http.ServeMux,
	handler func(c CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse, error),
) {
	method, pattern, h := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
