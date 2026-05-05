package external

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action GetOfferTranslationsAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetOfferTranslationsAction
func GetOfferTranslationsAction(c GetOfferTranslationsActionRequest) (*GetOfferTranslationsActionResponse, error) {
	return &GetOfferTranslationsActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetOfferTranslationsActionMeta() struct {
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
		Name:        "GetOfferTranslationsAction",
		CliName:     "get offer translations-action",
		URL:         "https://api.{environment}/sale/offers/{offerId}/translations",
		Method:      "GET",
		Description: `Get offer translation for given language or all present. Read more: PL / EN.`,
	}
}
func GetGetOfferTranslationsActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "translations",
			Type: "array",
		},
	}
}
func CastGetOfferTranslationsActionResFromCli(c emigo.CliCastable) GetOfferTranslationsActionRes {
	data := GetOfferTranslationsActionRes{}
	if c.IsSet("translations") {
		data.Translations = emigo.CapturePossibleArray(CastGetOfferTranslationsActionResTranslationsFromCli, "translations", c)
	}
	return data
}

// The base class definition for getOfferTranslationsActionRes
type GetOfferTranslationsActionRes struct {
	Translations []GetOfferTranslationsActionResTranslations `json:"translations" yaml:"translations"`
}

func GetGetOfferTranslationsActionResTranslationsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "language",
			Type: "string",
		},
		{
			Name:     prefix + "title",
			Type:     "object",
			Children: GetGetOfferTranslationsActionResTranslationsTitleCliFlags("title-"),
		},
		{
			Name:     prefix + "description",
			Type:     "object",
			Children: GetGetOfferTranslationsActionResTranslationsDescriptionCliFlags("description-"),
		},
		{
			Name:     prefix + "safety-information",
			Type:     "object",
			Children: GetGetOfferTranslationsActionResTranslationsSafetyInformationCliFlags("safety-information-"),
		},
	}
}
func CastGetOfferTranslationsActionResTranslationsFromCli(c emigo.CliCastable) GetOfferTranslationsActionResTranslations {
	data := GetOfferTranslationsActionResTranslations{}
	if c.IsSet("language") {
		data.Language = c.String("language")
	}
	if c.IsSet("title") {
		data.Title = CastGetOfferTranslationsActionResTranslationsTitleFromCli(c)
	}
	if c.IsSet("description") {
		data.Description = CastGetOfferTranslationsActionResTranslationsDescriptionFromCli(c)
	}
	if c.IsSet("safety-information") {
		data.SafetyInformation = CastGetOfferTranslationsActionResTranslationsSafetyInformationFromCli(c)
	}
	return data
}

// The base class definition for translations
type GetOfferTranslationsActionResTranslations struct {
	Language          string                                                     `json:"language" yaml:"language"`
	Title             GetOfferTranslationsActionResTranslationsTitle             `json:"title" yaml:"title"`
	Description       GetOfferTranslationsActionResTranslationsDescription       `json:"description" yaml:"description"`
	SafetyInformation GetOfferTranslationsActionResTranslationsSafetyInformation `json:"safetyInformation" yaml:"safetyInformation"`
}

func GetGetOfferTranslationsActionResTranslationsTitleCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "translation",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastGetOfferTranslationsActionResTranslationsTitleFromCli(c emigo.CliCastable) GetOfferTranslationsActionResTranslationsTitle {
	data := GetOfferTranslationsActionResTranslationsTitle{}
	if c.IsSet("translation") {
		data.Translation = c.String("translation")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for title
type GetOfferTranslationsActionResTranslationsTitle struct {
	Translation string `json:"translation" yaml:"translation"`
	Type        string `json:"type" yaml:"type"`
}

func GetGetOfferTranslationsActionResTranslationsDescriptionCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "translation",
			Type:     "object",
			Children: GetGetOfferTranslationsActionResTranslationsDescriptionTranslationCliFlags("translation-"),
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastGetOfferTranslationsActionResTranslationsDescriptionFromCli(c emigo.CliCastable) GetOfferTranslationsActionResTranslationsDescription {
	data := GetOfferTranslationsActionResTranslationsDescription{}
	if c.IsSet("translation") {
		data.Translation = CastGetOfferTranslationsActionResTranslationsDescriptionTranslationFromCli(c)
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for description
type GetOfferTranslationsActionResTranslationsDescription struct {
	Translation GetOfferTranslationsActionResTranslationsDescriptionTranslation `json:"translation" yaml:"translation"`
	Type        string                                                          `json:"type" yaml:"type"`
}

func GetGetOfferTranslationsActionResTranslationsDescriptionTranslationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sections",
			Type: "array",
		},
	}
}
func CastGetOfferTranslationsActionResTranslationsDescriptionTranslationFromCli(c emigo.CliCastable) GetOfferTranslationsActionResTranslationsDescriptionTranslation {
	data := GetOfferTranslationsActionResTranslationsDescriptionTranslation{}
	if c.IsSet("sections") {
		data.Sections = emigo.CapturePossibleArray(CastGetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsFromCli, "sections", c)
	}
	return data
}

// The base class definition for translation
type GetOfferTranslationsActionResTranslationsDescriptionTranslation struct {
	Sections []GetOfferTranslationsActionResTranslationsDescriptionTranslationSections `json:"sections" yaml:"sections"`
}

func GetGetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
	}
}
func CastGetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsFromCli(c emigo.CliCastable) GetOfferTranslationsActionResTranslationsDescriptionTranslationSections {
	data := GetOfferTranslationsActionResTranslationsDescriptionTranslationSections{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastGetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItemsFromCli, "items", c)
	}
	return data
}

// The base class definition for sections
type GetOfferTranslationsActionResTranslationsDescriptionTranslationSections struct {
	Items []GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems `json:"items" yaml:"items"`
}

func GetGetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastGetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItemsFromCli(c emigo.CliCastable) GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems {
	data := GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for items
type GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

func GetGetOfferTranslationsActionResTranslationsSafetyInformationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "products",
			Type: "array",
		},
	}
}
func CastGetOfferTranslationsActionResTranslationsSafetyInformationFromCli(c emigo.CliCastable) GetOfferTranslationsActionResTranslationsSafetyInformation {
	data := GetOfferTranslationsActionResTranslationsSafetyInformation{}
	if c.IsSet("products") {
		data.Products = emigo.CapturePossibleArray(CastGetOfferTranslationsActionResTranslationsSafetyInformationProductsFromCli, "products", c)
	}
	return data
}

// The base class definition for safetyInformation
type GetOfferTranslationsActionResTranslationsSafetyInformation struct {
	Products []GetOfferTranslationsActionResTranslationsSafetyInformationProducts `json:"products" yaml:"products"`
}

func GetGetOfferTranslationsActionResTranslationsSafetyInformationProductsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "translation",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastGetOfferTranslationsActionResTranslationsSafetyInformationProductsFromCli(c emigo.CliCastable) GetOfferTranslationsActionResTranslationsSafetyInformationProducts {
	data := GetOfferTranslationsActionResTranslationsSafetyInformationProducts{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("translation") {
		data.Translation = c.String("translation")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for products
type GetOfferTranslationsActionResTranslationsSafetyInformationProducts struct {
	Id          string `json:"id" yaml:"id"`
	Translation string `json:"translation" yaml:"translation"`
	Type        string `json:"type" yaml:"type"`
}

func (x *GetOfferTranslationsActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetOfferTranslationsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetOfferTranslationsActionResponse) SetContentType(contentType string) *GetOfferTranslationsActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetOfferTranslationsActionResponse) AsStream(r io.Reader, contentType string) *GetOfferTranslationsActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetOfferTranslationsActionResponse) AsJSON(payload any) *GetOfferTranslationsActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetOfferTranslationsActionResponse) WithIdeal(payload GetOfferTranslationsActionRes) *GetOfferTranslationsActionResponse {
	x.Payload = payload
	return x
}
func (x *GetOfferTranslationsActionResponse) AsHTML(payload string) *GetOfferTranslationsActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetOfferTranslationsActionResponse) AsBytes(payload []byte) *GetOfferTranslationsActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetOfferTranslationsActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetOfferTranslationsActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetOfferTranslationsActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetOfferTranslationsActionRaw registers a raw Gin route for the GetOfferTranslationsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetOfferTranslationsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetOfferTranslationsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetOfferTranslationsActionRequestSig = func(c GetOfferTranslationsActionRequest) (*GetOfferTranslationsActionResponse, error)

// GetOfferTranslationsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetOfferTranslationsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetOfferTranslationsActionHandler(
	handler GetOfferTranslationsActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetOfferTranslationsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetOfferTranslationsActionRequest{
			Body:        nil,
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx:      m,
		}
		resp, err := handler(req)
		if err != nil {
			m.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// If the handler returned nil (and no error), it means the response was handled manually.
		if resp == nil {
			return
		}
		// Apply headers
		for k, v := range resp.Headers {
			m.Header(k, v)
		}
		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}
		if resp.Payload != nil {
			m.JSON(status, resp.Payload)
		} else {
			m.Status(status)
		}
	}
}

// GetOfferTranslationsAction is a high-level convenience wrapper around GetOfferTranslationsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetOfferTranslationsActionGin(r gin.IRoutes, handler GetOfferTranslationsActionRequestSig) {
	method, url, h := GetOfferTranslationsActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get offer translationsAction
 */
// Query wrapper with private fields
type GetOfferTranslationsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetOfferTranslationsActionQueryFromString(rawQuery string) GetOfferTranslationsActionQuery {
	v := GetOfferTranslationsActionQuery{}
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
func GetOfferTranslationsActionQueryFromGin(c *gin.Context) GetOfferTranslationsActionQuery {
	return GetOfferTranslationsActionQueryFromString(c.Request.URL.RawQuery)
}
func GetOfferTranslationsActionQueryFromHttp(r *http.Request) GetOfferTranslationsActionQuery {
	return GetOfferTranslationsActionQueryFromString(r.URL.RawQuery)
}
func (q GetOfferTranslationsActionQuery) Values() url.Values {
	return q.values
}
func (q GetOfferTranslationsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetOfferTranslationsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetOfferTranslationsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetOfferTranslationsActionRequest struct {
	Body        interface{}
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	GinCtx *gin.Context
	// Urfave context, per each request
	CliCtx *cli.Command
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func (x GetOfferTranslationsActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GetOfferTranslationsActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GetOfferTranslationsActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GetOfferTranslationsActionClientCreateUrl(
	req GetOfferTranslationsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetOfferTranslationsActionMeta()
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
func GetOfferTranslationsActionClientExecuteTyped(httpReq *http.Request) (*GetOfferTranslationsActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetOfferTranslationsActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetOfferTranslationsActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetOfferTranslationsActionResponse{Payload: result}, err
	}
	return &GetOfferTranslationsActionResponse{Payload: result}, nil
}
func GetOfferTranslationsActionClientBuildRequest(req GetOfferTranslationsActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetOfferTranslationsActionMeta()
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
func GetOfferTranslationsActionCall(
	req GetOfferTranslationsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetOfferTranslationsActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetOfferTranslationsActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetOfferTranslationsActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetOfferTranslationsActionClientExecuteTyped(r)
}
