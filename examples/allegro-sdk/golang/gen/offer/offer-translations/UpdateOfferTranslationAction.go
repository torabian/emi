package external

import (
	"bytes"
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action UpdateOfferTranslationAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of UpdateOfferTranslationAction
func UpdateOfferTranslationAction(c UpdateOfferTranslationActionRequest) (*UpdateOfferTranslationActionResponse, error) {
	return &UpdateOfferTranslationActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func UpdateOfferTranslationActionMeta() struct {
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
		Name:        "UpdateOfferTranslationAction",
		CliName:     "update offer translation-action",
		URL:         "https://api.{environment}/sale/offers/{offerId}/translations/{language}",
		Method:      "PATCH",
		Description: `Update manual translation for offer. Read more: PL / EN.`,
	}
}
func GetUpdateOfferTranslationActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "description",
			Type:     "object",
			Children: GetUpdateOfferTranslationActionReqDescriptionCliFlags("description-"),
		},
		{
			Name:     prefix + "title",
			Type:     "object",
			Children: GetUpdateOfferTranslationActionReqTitleCliFlags("title-"),
		},
		{
			Name:     prefix + "safety-information",
			Type:     "object",
			Children: GetUpdateOfferTranslationActionReqSafetyInformationCliFlags("safety-information-"),
		},
	}
}
func CastUpdateOfferTranslationActionReqFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReq {
	data := UpdateOfferTranslationActionReq{}
	if c.IsSet("description") {
		data.Description = CastUpdateOfferTranslationActionReqDescriptionFromCli(c)
	}
	if c.IsSet("title") {
		data.Title = CastUpdateOfferTranslationActionReqTitleFromCli(c)
	}
	if c.IsSet("safety-information") {
		data.SafetyInformation = CastUpdateOfferTranslationActionReqSafetyInformationFromCli(c)
	}
	return data
}

// The base class definition for updateOfferTranslationActionReq
type UpdateOfferTranslationActionReq struct {
	Description       UpdateOfferTranslationActionReqDescription       `json:"description" yaml:"description"`
	Title             UpdateOfferTranslationActionReqTitle             `json:"title" yaml:"title"`
	SafetyInformation UpdateOfferTranslationActionReqSafetyInformation `json:"safetyInformation" yaml:"safetyInformation"`
}

func GetUpdateOfferTranslationActionReqDescriptionCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "translation",
			Type:     "object",
			Children: GetUpdateOfferTranslationActionReqDescriptionTranslationCliFlags("translation-"),
		},
	}
}
func CastUpdateOfferTranslationActionReqDescriptionFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqDescription {
	data := UpdateOfferTranslationActionReqDescription{}
	if c.IsSet("translation") {
		data.Translation = CastUpdateOfferTranslationActionReqDescriptionTranslationFromCli(c)
	}
	return data
}

// The base class definition for description
type UpdateOfferTranslationActionReqDescription struct {
	Translation UpdateOfferTranslationActionReqDescriptionTranslation `json:"translation" yaml:"translation"`
}

func GetUpdateOfferTranslationActionReqDescriptionTranslationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sections",
			Type: "array",
		},
	}
}
func CastUpdateOfferTranslationActionReqDescriptionTranslationFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqDescriptionTranslation {
	data := UpdateOfferTranslationActionReqDescriptionTranslation{}
	if c.IsSet("sections") {
		data.Sections = emigo.CapturePossibleArray(CastUpdateOfferTranslationActionReqDescriptionTranslationSectionsFromCli, "sections", c)
	}
	return data
}

// The base class definition for translation
type UpdateOfferTranslationActionReqDescriptionTranslation struct {
	Sections emigo.Array[UpdateOfferTranslationActionReqDescriptionTranslationSections] `json:"sections" yaml:"sections"`
}

func GetUpdateOfferTranslationActionReqDescriptionTranslationSectionsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
	}
}
func CastUpdateOfferTranslationActionReqDescriptionTranslationSectionsFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqDescriptionTranslationSections {
	data := UpdateOfferTranslationActionReqDescriptionTranslationSections{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastUpdateOfferTranslationActionReqDescriptionTranslationSectionsItemsFromCli, "items", c)
	}
	return data
}

// The base class definition for sections
type UpdateOfferTranslationActionReqDescriptionTranslationSections struct {
	Items emigo.Array[UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems] `json:"items" yaml:"items"`
}

func GetUpdateOfferTranslationActionReqDescriptionTranslationSectionsItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastUpdateOfferTranslationActionReqDescriptionTranslationSectionsItemsFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems {
	data := UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for items
type UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

func GetUpdateOfferTranslationActionReqTitleCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "translation",
			Type: "string",
		},
	}
}
func CastUpdateOfferTranslationActionReqTitleFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqTitle {
	data := UpdateOfferTranslationActionReqTitle{}
	if c.IsSet("translation") {
		data.Translation = c.String("translation")
	}
	return data
}

// The base class definition for title
type UpdateOfferTranslationActionReqTitle struct {
	Translation string `json:"translation" yaml:"translation"`
}

func GetUpdateOfferTranslationActionReqSafetyInformationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "products",
			Type: "array",
		},
	}
}
func CastUpdateOfferTranslationActionReqSafetyInformationFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqSafetyInformation {
	data := UpdateOfferTranslationActionReqSafetyInformation{}
	if c.IsSet("products") {
		data.Products = emigo.CapturePossibleArray(CastUpdateOfferTranslationActionReqSafetyInformationProductsFromCli, "products", c)
	}
	return data
}

// The base class definition for safetyInformation
type UpdateOfferTranslationActionReqSafetyInformation struct {
	Products emigo.Array[UpdateOfferTranslationActionReqSafetyInformationProducts] `json:"products" yaml:"products"`
}

func GetUpdateOfferTranslationActionReqSafetyInformationProductsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "translation",
			Type: "string",
		},
	}
}
func CastUpdateOfferTranslationActionReqSafetyInformationProductsFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqSafetyInformationProducts {
	data := UpdateOfferTranslationActionReqSafetyInformationProducts{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("translation") {
		data.Translation = c.String("translation")
	}
	return data
}

// The base class definition for products
type UpdateOfferTranslationActionReqSafetyInformationProducts struct {
	Id          string `json:"id" yaml:"id"`
	Translation string `json:"translation" yaml:"translation"`
}

func (x *UpdateOfferTranslationActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type UpdateOfferTranslationActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *UpdateOfferTranslationActionResponse) SetContentType(contentType string) *UpdateOfferTranslationActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *UpdateOfferTranslationActionResponse) AsStream(r io.Reader, contentType string) *UpdateOfferTranslationActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *UpdateOfferTranslationActionResponse) AsJSON(payload any) *UpdateOfferTranslationActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *UpdateOfferTranslationActionResponse) AsHTML(payload string) *UpdateOfferTranslationActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *UpdateOfferTranslationActionResponse) AsBytes(payload []byte) *UpdateOfferTranslationActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x UpdateOfferTranslationActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x UpdateOfferTranslationActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x UpdateOfferTranslationActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type UpdateOfferTranslationActionRequestSig = func(c UpdateOfferTranslationActionRequest) (*UpdateOfferTranslationActionResponse, error)

/**
 * Query parameters for Update offer translationAction
 */
// Query wrapper with private fields
type UpdateOfferTranslationActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func UpdateOfferTranslationActionQueryFromString(rawQuery string) UpdateOfferTranslationActionQuery {
	v := UpdateOfferTranslationActionQuery{}
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
func UpdateOfferTranslationActionQueryFromHttp(r *http.Request) UpdateOfferTranslationActionQuery {
	return UpdateOfferTranslationActionQueryFromString(r.URL.RawQuery)
}
func (q UpdateOfferTranslationActionQuery) Values() url.Values {
	return q.values
}
func (q UpdateOfferTranslationActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *UpdateOfferTranslationActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *UpdateOfferTranslationActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type UpdateOfferTranslationActionRequest struct {
	Body        UpdateOfferTranslationActionReq
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

func UpdateOfferTranslationActionClientCreateUrl(
	req UpdateOfferTranslationActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := UpdateOfferTranslationActionMeta()
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
func UpdateOfferTranslationActionClientExecuteTyped(httpReq *http.Request) (*UpdateOfferTranslationActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result UpdateOfferTranslationActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &UpdateOfferTranslationActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &UpdateOfferTranslationActionResponse{Payload: result}, err
	}
	return &UpdateOfferTranslationActionResponse{Payload: result}, nil
}
func UpdateOfferTranslationActionClientBuildRequest(req UpdateOfferTranslationActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := UpdateOfferTranslationActionMeta()
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
func UpdateOfferTranslationActionCall(
	req UpdateOfferTranslationActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*UpdateOfferTranslationActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := UpdateOfferTranslationActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := UpdateOfferTranslationActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return UpdateOfferTranslationActionClientExecuteTyped(r)
}
func (x UpdateOfferTranslationActionRequest) IsCli() bool {
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

// UpdateOfferTranslationActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the UpdateOfferTranslationAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *UpdateOfferTranslationActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func UpdateOfferTranslationActionHttpHandler(
	handler func(c UpdateOfferTranslationActionRequest) (*UpdateOfferTranslationActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := UpdateOfferTranslationActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body UpdateOfferTranslationActionReq
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
		req := UpdateOfferTranslationActionRequest{
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

// UpdateOfferTranslationActionHttp is a high-level convenience wrapper around
// UpdateOfferTranslationActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func UpdateOfferTranslationActionHttp(
	mux *http.ServeMux,
	handler func(c UpdateOfferTranslationActionRequest) (*UpdateOfferTranslationActionResponse, error),
) {
	method, pattern, h := UpdateOfferTranslationActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
