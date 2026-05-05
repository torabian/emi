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
* Action to communicate with the action GetSmartClassificationReportOfTheParticularOfferAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetSmartClassificationReportOfTheParticularOfferAction
func GetSmartClassificationReportOfTheParticularOfferAction(c GetSmartClassificationReportOfTheParticularOfferActionRequest) (*GetSmartClassificationReportOfTheParticularOfferActionResponse, error) {
	return &GetSmartClassificationReportOfTheParticularOfferActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetSmartClassificationReportOfTheParticularOfferActionMeta() struct {
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
		Name:        "GetSmartClassificationReportOfTheParticularOfferAction",
		CliName:     "get -smart! classification report of the particular offer-action",
		URL:         "https://api.{environment}/sale/offers/{offerId}/smart",
		Method:      "GET",
		Description: `Use this resource to get a full Smart! offer classification report of one of your offers. Please keep in mind you have to meet Smart! seller conditions first - for more details, use GET /sale/smart. To learn more about Smart! offer requirements, see our knowledge base article: PL / EN. Read more: PL / EN.`,
	}
}
func GetGetSmartClassificationReportOfTheParticularOfferActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "scheduled-for-reclassification",
			Type:        "bool",
			Description: "Indicates if offer is queued for reclassification",
		},
		{
			Name:        prefix + "classification",
			Type:        "object",
			Children:    GetGetSmartClassificationReportOfTheParticularOfferActionResClassificationCliFlags("classification-"),
			Description: "Offer classification status and last change date",
		},
		{
			Name:        prefix + "smart-delivery-methods",
			Type:        "array",
			Description: "List of smart delivery method identifiers",
		},
		{
			Name:        prefix + "conditions",
			Type:        "array",
			Description: "List of classification conditions with delivery method checks",
		},
	}
}
func CastGetSmartClassificationReportOfTheParticularOfferActionResFromCli(c emigo.CliCastable) GetSmartClassificationReportOfTheParticularOfferActionRes {
	data := GetSmartClassificationReportOfTheParticularOfferActionRes{}
	if c.IsSet("scheduled-for-reclassification") {
		data.ScheduledForReclassification = bool(c.Bool("scheduled-for-reclassification"))
	}
	if c.IsSet("classification") {
		data.Classification = CastGetSmartClassificationReportOfTheParticularOfferActionResClassificationFromCli(c)
	}
	if c.IsSet("smart-delivery-methods") {
		data.SmartDeliveryMethods = emigo.CapturePossibleArray(CastGetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethodsFromCli, "smart-delivery-methods", c)
	}
	if c.IsSet("conditions") {
		data.Conditions = emigo.CapturePossibleArray(CastGetSmartClassificationReportOfTheParticularOfferActionResConditionsFromCli, "conditions", c)
	}
	return data
}

// The base class definition for getSmartClassificationReportOfTheParticularOfferActionRes
type GetSmartClassificationReportOfTheParticularOfferActionRes struct {
	// Indicates if offer is queued for reclassification
	ScheduledForReclassification bool `json:"scheduledForReclassification" yaml:"scheduledForReclassification"`
	// Offer classification status and last change date
	Classification GetSmartClassificationReportOfTheParticularOfferActionResClassification `json:"classification" yaml:"classification"`
	// List of smart delivery method identifiers
	SmartDeliveryMethods []GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods `json:"smartDeliveryMethods" yaml:"smartDeliveryMethods"`
	// List of classification conditions with delivery method checks
	Conditions []GetSmartClassificationReportOfTheParticularOfferActionResConditions `json:"conditions" yaml:"conditions"`
}

func GetGetSmartClassificationReportOfTheParticularOfferActionResClassificationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "fulfilled",
			Type:        "bool",
			Description: "Whether the classification conditions are fulfilled",
		},
		{
			Name:        prefix + "last-changed",
			Type:        "string",
			Description: "ISO8601 timestamp of last classification change",
		},
	}
}
func CastGetSmartClassificationReportOfTheParticularOfferActionResClassificationFromCli(c emigo.CliCastable) GetSmartClassificationReportOfTheParticularOfferActionResClassification {
	data := GetSmartClassificationReportOfTheParticularOfferActionResClassification{}
	if c.IsSet("fulfilled") {
		data.Fulfilled = bool(c.Bool("fulfilled"))
	}
	if c.IsSet("last-changed") {
		data.LastChanged = c.String("last-changed")
	}
	return data
}

// The base class definition for classification
type GetSmartClassificationReportOfTheParticularOfferActionResClassification struct {
	// Whether the classification conditions are fulfilled
	Fulfilled bool `json:"fulfilled" yaml:"fulfilled"`
	// ISO8601 timestamp of last classification change
	LastChanged string `json:"lastChanged" yaml:"lastChanged"`
}

func GetGetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethodsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethodsFromCli(c emigo.CliCastable) GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods {
	data := GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for smartDeliveryMethods
type GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSmartClassificationReportOfTheParticularOfferActionResConditionsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "code",
			Type:        "string",
			Description: "Condition code identifier",
		},
		{
			Name:        prefix + "name",
			Type:        "string",
			Description: "Human-readable condition name",
		},
		{
			Name:        prefix + "description",
			Type:        "string",
			Description: "Detailed condition description",
		},
		{
			Name:        prefix + "fulfilled",
			Type:        "bool",
			Description: "Indicates if this condition is fulfilled",
		},
		{
			Name:        prefix + "passed-delivery-methods",
			Type:        "array",
			Description: "Delivery methods that passed validation for this condition",
		},
		{
			Name:        prefix + "failed-delivery-methods",
			Type:        "array",
			Description: "Delivery methods that failed validation for this condition",
		},
	}
}
func CastGetSmartClassificationReportOfTheParticularOfferActionResConditionsFromCli(c emigo.CliCastable) GetSmartClassificationReportOfTheParticularOfferActionResConditions {
	data := GetSmartClassificationReportOfTheParticularOfferActionResConditions{}
	if c.IsSet("code") {
		data.Code = c.String("code")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	if c.IsSet("fulfilled") {
		data.Fulfilled = bool(c.Bool("fulfilled"))
	}
	if c.IsSet("passed-delivery-methods") {
		data.PassedDeliveryMethods = emigo.CapturePossibleArray(CastGetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethodsFromCli, "passed-delivery-methods", c)
	}
	if c.IsSet("failed-delivery-methods") {
		data.FailedDeliveryMethods = emigo.CapturePossibleArray(CastGetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethodsFromCli, "failed-delivery-methods", c)
	}
	return data
}

// The base class definition for conditions
type GetSmartClassificationReportOfTheParticularOfferActionResConditions struct {
	// Condition code identifier
	Code string `json:"code" yaml:"code"`
	// Human-readable condition name
	Name string `json:"name" yaml:"name"`
	// Detailed condition description
	Description string `json:"description" yaml:"description"`
	// Indicates if this condition is fulfilled
	Fulfilled bool `json:"fulfilled" yaml:"fulfilled"`
	// Delivery methods that passed validation for this condition
	PassedDeliveryMethods []GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods `json:"passedDeliveryMethods" yaml:"passedDeliveryMethods"`
	// Delivery methods that failed validation for this condition
	FailedDeliveryMethods []GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods `json:"failedDeliveryMethods" yaml:"failedDeliveryMethods"`
}

func GetGetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethodsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethodsFromCli(c emigo.CliCastable) GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods {
	data := GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for passedDeliveryMethods
type GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethodsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethodsFromCli(c emigo.CliCastable) GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods {
	data := GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for failedDeliveryMethods
type GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods struct {
	Id string `json:"id" yaml:"id"`
}

func (x *GetSmartClassificationReportOfTheParticularOfferActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetSmartClassificationReportOfTheParticularOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetSmartClassificationReportOfTheParticularOfferActionResponse) SetContentType(contentType string) *GetSmartClassificationReportOfTheParticularOfferActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetSmartClassificationReportOfTheParticularOfferActionResponse) AsStream(r io.Reader, contentType string) *GetSmartClassificationReportOfTheParticularOfferActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetSmartClassificationReportOfTheParticularOfferActionResponse) AsJSON(payload any) *GetSmartClassificationReportOfTheParticularOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetSmartClassificationReportOfTheParticularOfferActionResponse) WithIdeal(payload GetSmartClassificationReportOfTheParticularOfferActionRes) *GetSmartClassificationReportOfTheParticularOfferActionResponse {
	x.Payload = payload
	return x
}
func (x *GetSmartClassificationReportOfTheParticularOfferActionResponse) AsHTML(payload string) *GetSmartClassificationReportOfTheParticularOfferActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetSmartClassificationReportOfTheParticularOfferActionResponse) AsBytes(payload []byte) *GetSmartClassificationReportOfTheParticularOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetSmartClassificationReportOfTheParticularOfferActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetSmartClassificationReportOfTheParticularOfferActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetSmartClassificationReportOfTheParticularOfferActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetSmartClassificationReportOfTheParticularOfferActionRaw registers a raw Gin route for the GetSmartClassificationReportOfTheParticularOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetSmartClassificationReportOfTheParticularOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetSmartClassificationReportOfTheParticularOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetSmartClassificationReportOfTheParticularOfferActionRequestSig = func(c GetSmartClassificationReportOfTheParticularOfferActionRequest) (*GetSmartClassificationReportOfTheParticularOfferActionResponse, error)

// GetSmartClassificationReportOfTheParticularOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetSmartClassificationReportOfTheParticularOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetSmartClassificationReportOfTheParticularOfferActionHandler(
	handler GetSmartClassificationReportOfTheParticularOfferActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetSmartClassificationReportOfTheParticularOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetSmartClassificationReportOfTheParticularOfferActionRequest{
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

// GetSmartClassificationReportOfTheParticularOfferAction is a high-level convenience wrapper around GetSmartClassificationReportOfTheParticularOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetSmartClassificationReportOfTheParticularOfferActionGin(r gin.IRoutes, handler GetSmartClassificationReportOfTheParticularOfferActionRequestSig) {
	method, url, h := GetSmartClassificationReportOfTheParticularOfferActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get Smart! classification report of the particular offerAction
 */
// Query wrapper with private fields
type GetSmartClassificationReportOfTheParticularOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetSmartClassificationReportOfTheParticularOfferActionQueryFromString(rawQuery string) GetSmartClassificationReportOfTheParticularOfferActionQuery {
	v := GetSmartClassificationReportOfTheParticularOfferActionQuery{}
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
func GetSmartClassificationReportOfTheParticularOfferActionQueryFromGin(c *gin.Context) GetSmartClassificationReportOfTheParticularOfferActionQuery {
	return GetSmartClassificationReportOfTheParticularOfferActionQueryFromString(c.Request.URL.RawQuery)
}
func GetSmartClassificationReportOfTheParticularOfferActionQueryFromHttp(r *http.Request) GetSmartClassificationReportOfTheParticularOfferActionQuery {
	return GetSmartClassificationReportOfTheParticularOfferActionQueryFromString(r.URL.RawQuery)
}
func (q GetSmartClassificationReportOfTheParticularOfferActionQuery) Values() url.Values {
	return q.values
}
func (q GetSmartClassificationReportOfTheParticularOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetSmartClassificationReportOfTheParticularOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetSmartClassificationReportOfTheParticularOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetSmartClassificationReportOfTheParticularOfferActionRequest struct {
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

func (x GetSmartClassificationReportOfTheParticularOfferActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GetSmartClassificationReportOfTheParticularOfferActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GetSmartClassificationReportOfTheParticularOfferActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GetSmartClassificationReportOfTheParticularOfferActionClientCreateUrl(
	req GetSmartClassificationReportOfTheParticularOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetSmartClassificationReportOfTheParticularOfferActionMeta()
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
func GetSmartClassificationReportOfTheParticularOfferActionClientExecuteTyped(httpReq *http.Request) (*GetSmartClassificationReportOfTheParticularOfferActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetSmartClassificationReportOfTheParticularOfferActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetSmartClassificationReportOfTheParticularOfferActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetSmartClassificationReportOfTheParticularOfferActionResponse{Payload: result}, err
	}
	return &GetSmartClassificationReportOfTheParticularOfferActionResponse{Payload: result}, nil
}
func GetSmartClassificationReportOfTheParticularOfferActionClientBuildRequest(req GetSmartClassificationReportOfTheParticularOfferActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetSmartClassificationReportOfTheParticularOfferActionMeta()
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
func GetSmartClassificationReportOfTheParticularOfferActionCall(
	req GetSmartClassificationReportOfTheParticularOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetSmartClassificationReportOfTheParticularOfferActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetSmartClassificationReportOfTheParticularOfferActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetSmartClassificationReportOfTheParticularOfferActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetSmartClassificationReportOfTheParticularOfferActionClientExecuteTyped(r)
}
