package external

import (
	"bytes"
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ModifyTheBuyNowPriceInAnOfferAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ModifyTheBuyNowPriceInAnOfferAction
func ModifyTheBuyNowPriceInAnOfferAction(c ModifyTheBuyNowPriceInAnOfferActionRequest) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error) {
	return &ModifyTheBuyNowPriceInAnOfferActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ModifyTheBuyNowPriceInAnOfferActionMeta() struct {
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
		Name:        "ModifyTheBuyNowPriceInAnOfferAction",
		CliName:     "modify the -buy -now price in an offer-action",
		URL:         "https://api.{environment}/offers/{offerId}/change-price-commands/{commandId}",
		Method:      "PUT",
		Description: `Use this resource to change the Buy Now price in a single offer. Read more: PL / EN.`,
	}
}

// The base class definition for modifyTheBuyNowPriceInAnOfferActionReq
type ModifyTheBuyNowPriceInAnOfferActionReq struct {
	Id    string                                      `json:"id" yaml:"id"`
	Input ModifyTheBuyNowPriceInAnOfferActionReqInput `json:"input" yaml:"input"`
}

// The base class definition for input
type ModifyTheBuyNowPriceInAnOfferActionReqInput struct {
	BuyNowPrice ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice `json:"buyNowPrice" yaml:"buyNowPrice"`
}

// The base class definition for buyNowPrice
type ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func (x *ModifyTheBuyNowPriceInAnOfferActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// The base class definition for modifyTheBuyNowPriceInAnOfferActionRes
type ModifyTheBuyNowPriceInAnOfferActionRes struct {
	Id     string                                       `json:"id" yaml:"id"`
	Input  ModifyTheBuyNowPriceInAnOfferActionResInput  `json:"input" yaml:"input"`
	Output ModifyTheBuyNowPriceInAnOfferActionResOutput `json:"output" yaml:"output"`
}

// The base class definition for input
type ModifyTheBuyNowPriceInAnOfferActionResInput struct {
	BuyNowPrice ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice `json:"buyNowPrice" yaml:"buyNowPrice"`
}

// The base class definition for buyNowPrice
type ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for output
type ModifyTheBuyNowPriceInAnOfferActionResOutput struct {
	Status string                                                          `json:"status" yaml:"status"`
	Errors emigo.Array[ModifyTheBuyNowPriceInAnOfferActionResOutputErrors] `json:"errors" yaml:"errors"`
}

// The base class definition for errors
type ModifyTheBuyNowPriceInAnOfferActionResOutputErrors struct {
	Code        string                                                     `json:"code" yaml:"code"`
	Details     string                                                     `json:"details" yaml:"details"`
	Message     string                                                     `json:"message" yaml:"message"`
	Path        string                                                     `json:"path" yaml:"path"`
	UserMessage string                                                     `json:"userMessage" yaml:"userMessage"`
	Metadata    ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata `json:"metadata" yaml:"metadata"`
}

// The base class definition for metadata
type ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

func (x *ModifyTheBuyNowPriceInAnOfferActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ModifyTheBuyNowPriceInAnOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) SetContentType(contentType string) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) AsStream(r io.Reader, contentType string) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) AsJSON(payload any) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) WithIdeal(payload ModifyTheBuyNowPriceInAnOfferActionRes) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = payload
	return x
}
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) AsHTML(payload string) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ModifyTheBuyNowPriceInAnOfferActionResponse) AsBytes(payload []byte) *ModifyTheBuyNowPriceInAnOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ModifyTheBuyNowPriceInAnOfferActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ModifyTheBuyNowPriceInAnOfferActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ModifyTheBuyNowPriceInAnOfferActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type ModifyTheBuyNowPriceInAnOfferActionRequestSig = func(c ModifyTheBuyNowPriceInAnOfferActionRequest) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error)

/**
 * Query parameters for Modify the Buy Now price in an offerAction
 */
// Query wrapper with private fields
type ModifyTheBuyNowPriceInAnOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModifyTheBuyNowPriceInAnOfferActionQueryFromString(rawQuery string) ModifyTheBuyNowPriceInAnOfferActionQuery {
	v := ModifyTheBuyNowPriceInAnOfferActionQuery{}
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
func ModifyTheBuyNowPriceInAnOfferActionQueryFromHttp(r *http.Request) ModifyTheBuyNowPriceInAnOfferActionQuery {
	return ModifyTheBuyNowPriceInAnOfferActionQueryFromString(r.URL.RawQuery)
}
func (q ModifyTheBuyNowPriceInAnOfferActionQuery) Values() url.Values {
	return q.values
}
func (q ModifyTheBuyNowPriceInAnOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModifyTheBuyNowPriceInAnOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModifyTheBuyNowPriceInAnOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModifyTheBuyNowPriceInAnOfferActionRequest struct {
	Body        ModifyTheBuyNowPriceInAnOfferActionReq
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

func ModifyTheBuyNowPriceInAnOfferActionClientCreateUrl(
	req ModifyTheBuyNowPriceInAnOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
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
func ModifyTheBuyNowPriceInAnOfferActionClientExecuteTyped(httpReq *http.Request) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ModifyTheBuyNowPriceInAnOfferActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ModifyTheBuyNowPriceInAnOfferActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ModifyTheBuyNowPriceInAnOfferActionResponse{Payload: result}, err
	}
	return &ModifyTheBuyNowPriceInAnOfferActionResponse{Payload: result}, nil
}
func ModifyTheBuyNowPriceInAnOfferActionClientBuildRequest(req ModifyTheBuyNowPriceInAnOfferActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
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
func ModifyTheBuyNowPriceInAnOfferActionCall(
	req ModifyTheBuyNowPriceInAnOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ModifyTheBuyNowPriceInAnOfferActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ModifyTheBuyNowPriceInAnOfferActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ModifyTheBuyNowPriceInAnOfferActionClientExecuteTyped(r)
}

// ModifyTheBuyNowPriceInAnOfferActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the ModifyTheBuyNowPriceInAnOfferAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *ModifyTheBuyNowPriceInAnOfferActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func ModifyTheBuyNowPriceInAnOfferActionHttpHandler(
	handler func(c ModifyTheBuyNowPriceInAnOfferActionRequest) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := ModifyTheBuyNowPriceInAnOfferActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body ModifyTheBuyNowPriceInAnOfferActionReq
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
		req := ModifyTheBuyNowPriceInAnOfferActionRequest{
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

// ModifyTheBuyNowPriceInAnOfferActionHttp is a high-level convenience wrapper around
// ModifyTheBuyNowPriceInAnOfferActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func ModifyTheBuyNowPriceInAnOfferActionHttp(
	mux *http.ServeMux,
	handler func(c ModifyTheBuyNowPriceInAnOfferActionRequest) (*ModifyTheBuyNowPriceInAnOfferActionResponse, error),
) {
	method, pattern, h := ModifyTheBuyNowPriceInAnOfferActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
