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
* Action to communicate with the action GetEventsAboutTheSellerSOffersAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetEventsAboutTheSellerSOffersAction
func GetEventsAboutTheSellerSOffersAction(c GetEventsAboutTheSellerSOffersActionRequest) (*GetEventsAboutTheSellerSOffersActionResponse, error) {
	return &GetEventsAboutTheSellerSOffersActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetEventsAboutTheSellerSOffersActionMeta() struct {
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
		Name:    "GetEventsAboutTheSellerSOffersAction",
		CliName: "get events about the seller's offers-action",
		URL:     "https://api.{environment}/sale/offer-events",
		Method:  "GET",
		Description: `Use this endpoint to get events from the last 24 hours concerning changes in the authorized seller's offers. At present we support the following events:
OFFER_ACTIVATED - offer is visible on site and available for purchase, occurs when offer status changes from ACTIVATING to ACTIVE. OFFER_CHANGED - occurs when offer's fields has been changed e.g. description or photos. OFFER_ENDED - offer is no longer available for purchase, occurs when offer status changes from ACTIVE to ENDED. OFFER_STOCK_CHANGED - stock in an offer was changed either via purchase or by seller. OFFER_PRICE_CHANGED - occurs when price in an offer was changed. OFFER_ARCHIVED - offer is no longer available on listing and has been archived. OFFER_BID_PLACED - bid was placed on the offer. OFFER_BID_CANCELED - bid for offer was canceled. OFFER_TRANSLATION_UPDATED - translation of offer was updated. OFFER_VISIBILITY_CHANGED - visibility of offer was changed on marketplaces. OFFER_DELIVERY_COUNTRIES_BLOCKED - the offer has been blocked in selected countries. Returned events may occur by actions made via browser or API. The resource allows you to get events concerning active offers and offers scheduled for activation (status ACTIVE and ACTIVATING). Returned events do not concern offers in INACTIVE and ENDED status (the exception is OFFER_ARCHIVED event). External id is returned for all event types except OFFER_BID_PLACED and OFFER_BID_CANCELED. Please note that one change may result in more than one event. Read more: PL / EN.`,
	}
}
func GetGetEventsAboutTheSellerSOffersActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "offer-events",
			Type:        "array",
			Description: "List of events related to offer state changes",
		},
	}
}
func CastGetEventsAboutTheSellerSOffersActionResFromCli(c emigo.CliCastable) GetEventsAboutTheSellerSOffersActionRes {
	data := GetEventsAboutTheSellerSOffersActionRes{}
	if c.IsSet("offer-events") {
		data.OfferEvents = emigo.CapturePossibleArray(CastGetEventsAboutTheSellerSOffersActionResOfferEventsFromCli, "offer-events", c)
	}
	return data
}

// The base class definition for getEventsAboutTheSellerSOffersActionRes
type GetEventsAboutTheSellerSOffersActionRes struct {
	// List of events related to offer state changes
	OfferEvents []GetEventsAboutTheSellerSOffersActionResOfferEvents `json:"offerEvents" yaml:"offerEvents"`
}

func GetGetEventsAboutTheSellerSOffersActionResOfferEventsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "string",
			Description: "Unique event identifier (base64 encoded)",
		},
		{
			Name:        prefix + "occurred-at",
			Type:        "string",
			Description: "ISO8601 timestamp when the event occurred",
		},
		{
			Name:        prefix + "type",
			Type:        "string",
			Description: "Event type (e.g., OFFER_ACTIVATED, OFFER_ENDED, etc.)",
		},
		{
			Name:        prefix + "offer",
			Type:        "object",
			Children:    GetGetEventsAboutTheSellerSOffersActionResOfferEventsOfferCliFlags("offer-"),
			Description: "Basic offer information for which event occurred",
		},
	}
}
func CastGetEventsAboutTheSellerSOffersActionResOfferEventsFromCli(c emigo.CliCastable) GetEventsAboutTheSellerSOffersActionResOfferEvents {
	data := GetEventsAboutTheSellerSOffersActionResOfferEvents{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("occurred-at") {
		data.OccurredAt = c.String("occurred-at")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("offer") {
		data.Offer = CastGetEventsAboutTheSellerSOffersActionResOfferEventsOfferFromCli(c)
	}
	return data
}

// The base class definition for offerEvents
type GetEventsAboutTheSellerSOffersActionResOfferEvents struct {
	// Unique event identifier (base64 encoded)
	Id string `json:"id" yaml:"id"`
	// ISO8601 timestamp when the event occurred
	OccurredAt string `json:"occurredAt" yaml:"occurredAt"`
	// Event type (e.g., OFFER_ACTIVATED, OFFER_ENDED, etc.)
	Type string `json:"type" yaml:"type"`
	// Basic offer information for which event occurred
	Offer GetEventsAboutTheSellerSOffersActionResOfferEventsOffer `json:"offer" yaml:"offer"`
}

func GetGetEventsAboutTheSellerSOffersActionResOfferEventsOfferCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name:     prefix + "external",
			Type:     "object",
			Children: GetGetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternalCliFlags("external-"),
		},
	}
}
func CastGetEventsAboutTheSellerSOffersActionResOfferEventsOfferFromCli(c emigo.CliCastable) GetEventsAboutTheSellerSOffersActionResOfferEventsOffer {
	data := GetEventsAboutTheSellerSOffersActionResOfferEventsOffer{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("external") {
		data.External = CastGetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternalFromCli(c)
	}
	return data
}

// The base class definition for offer
type GetEventsAboutTheSellerSOffersActionResOfferEventsOffer struct {
	Id       string                                                          `json:"id" yaml:"id"`
	External GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal `json:"external" yaml:"external"`
}

func GetGetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternalCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternalFromCli(c emigo.CliCastable) GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal {
	data := GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for external
type GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal struct {
	Id string `json:"id" yaml:"id"`
}

func (x *GetEventsAboutTheSellerSOffersActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetEventsAboutTheSellerSOffersActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetEventsAboutTheSellerSOffersActionResponse) SetContentType(contentType string) *GetEventsAboutTheSellerSOffersActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetEventsAboutTheSellerSOffersActionResponse) AsStream(r io.Reader, contentType string) *GetEventsAboutTheSellerSOffersActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetEventsAboutTheSellerSOffersActionResponse) AsJSON(payload any) *GetEventsAboutTheSellerSOffersActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetEventsAboutTheSellerSOffersActionResponse) WithIdeal(payload GetEventsAboutTheSellerSOffersActionRes) *GetEventsAboutTheSellerSOffersActionResponse {
	x.Payload = payload
	return x
}
func (x *GetEventsAboutTheSellerSOffersActionResponse) AsHTML(payload string) *GetEventsAboutTheSellerSOffersActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetEventsAboutTheSellerSOffersActionResponse) AsBytes(payload []byte) *GetEventsAboutTheSellerSOffersActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetEventsAboutTheSellerSOffersActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetEventsAboutTheSellerSOffersActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetEventsAboutTheSellerSOffersActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetEventsAboutTheSellerSOffersActionRaw registers a raw Gin route for the GetEventsAboutTheSellerSOffersAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetEventsAboutTheSellerSOffersActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetEventsAboutTheSellerSOffersActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetEventsAboutTheSellerSOffersActionRequestSig = func(c GetEventsAboutTheSellerSOffersActionRequest) (*GetEventsAboutTheSellerSOffersActionResponse, error)

// GetEventsAboutTheSellerSOffersActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetEventsAboutTheSellerSOffersAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetEventsAboutTheSellerSOffersActionHandler(
	handler GetEventsAboutTheSellerSOffersActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetEventsAboutTheSellerSOffersActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetEventsAboutTheSellerSOffersActionRequest{
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

// GetEventsAboutTheSellerSOffersAction is a high-level convenience wrapper around GetEventsAboutTheSellerSOffersActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetEventsAboutTheSellerSOffersActionGin(r gin.IRoutes, handler GetEventsAboutTheSellerSOffersActionRequestSig) {
	method, url, h := GetEventsAboutTheSellerSOffersActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get events about the seller's offersAction
 */
// Query wrapper with private fields
type GetEventsAboutTheSellerSOffersActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetEventsAboutTheSellerSOffersActionQueryFromString(rawQuery string) GetEventsAboutTheSellerSOffersActionQuery {
	v := GetEventsAboutTheSellerSOffersActionQuery{}
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
func GetEventsAboutTheSellerSOffersActionQueryFromGin(c *gin.Context) GetEventsAboutTheSellerSOffersActionQuery {
	return GetEventsAboutTheSellerSOffersActionQueryFromString(c.Request.URL.RawQuery)
}
func GetEventsAboutTheSellerSOffersActionQueryFromHttp(r *http.Request) GetEventsAboutTheSellerSOffersActionQuery {
	return GetEventsAboutTheSellerSOffersActionQueryFromString(r.URL.RawQuery)
}
func (q GetEventsAboutTheSellerSOffersActionQuery) Values() url.Values {
	return q.values
}
func (q GetEventsAboutTheSellerSOffersActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetEventsAboutTheSellerSOffersActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetEventsAboutTheSellerSOffersActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetEventsAboutTheSellerSOffersActionRequest struct {
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

func (x GetEventsAboutTheSellerSOffersActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GetEventsAboutTheSellerSOffersActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GetEventsAboutTheSellerSOffersActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GetEventsAboutTheSellerSOffersActionClientCreateUrl(
	req GetEventsAboutTheSellerSOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetEventsAboutTheSellerSOffersActionMeta()
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
func GetEventsAboutTheSellerSOffersActionClientExecuteTyped(httpReq *http.Request) (*GetEventsAboutTheSellerSOffersActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetEventsAboutTheSellerSOffersActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetEventsAboutTheSellerSOffersActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetEventsAboutTheSellerSOffersActionResponse{Payload: result}, err
	}
	return &GetEventsAboutTheSellerSOffersActionResponse{Payload: result}, nil
}
func GetEventsAboutTheSellerSOffersActionClientBuildRequest(req GetEventsAboutTheSellerSOffersActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetEventsAboutTheSellerSOffersActionMeta()
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
func GetEventsAboutTheSellerSOffersActionCall(
	req GetEventsAboutTheSellerSOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetEventsAboutTheSellerSOffersActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetEventsAboutTheSellerSOffersActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetEventsAboutTheSellerSOffersActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetEventsAboutTheSellerSOffersActionClientExecuteTyped(r)
}
