package external

import (
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action GetSellersOffersAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetSellersOffersAction
func GetSellersOffersAction(c GetSellersOffersActionRequest) (*GetSellersOffersActionResponse, error) {
	return &GetSellersOffersActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetSellersOffersActionMeta() struct {
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
		Name:        "GetSellersOffersAction",
		CliName:     "get sellers offers-action",
		URL:         "https://api.{environment}/sale/offers",
		Method:      "GET",
		Description: `Use this resource to get the list of the seller's offers. You can use different query parameters to filter the list. Read more: PL / EN.`,
	}
}

// The base class definition for getSellersOffersActionRes
type GetSellersOffersActionRes struct {
	// Number of offers in this page
	Count int `json:"count" yaml:"count"`
	// Total number of offers available
	TotalCount int                                          `json:"totalCount" yaml:"totalCount"`
	Offers     emigo.Array[GetSellersOffersActionResOffers] `json:"offers" yaml:"offers"`
}

// The base class definition for offers
type GetSellersOffersActionResOffers struct {
	// Offer identifier
	Id string `json:"id" yaml:"id"`
	// Offer name or title
	Name                string                                             `json:"name" yaml:"name"`
	Category            GetSellersOffersActionResOffersCategory            `json:"category" yaml:"category"`
	PrimaryImage        GetSellersOffersActionResOffersPrimaryImage        `json:"primaryImage" yaml:"primaryImage"`
	SellingMode         GetSellersOffersActionResOffersSellingMode         `json:"sellingMode" yaml:"sellingMode"`
	SaleInfo            GetSellersOffersActionResOffersSaleInfo            `json:"saleInfo" yaml:"saleInfo"`
	Stock               GetSellersOffersActionResOffersStock               `json:"stock" yaml:"stock"`
	Stats               GetSellersOffersActionResOffersStats               `json:"stats" yaml:"stats"`
	Publication         GetSellersOffersActionResOffersPublication         `json:"publication" yaml:"publication"`
	AfterSalesServices  GetSellersOffersActionResOffersAfterSalesServices  `json:"afterSalesServices" yaml:"afterSalesServices"`
	AdditionalServices  GetSellersOffersActionResOffersAdditionalServices  `json:"additionalServices" yaml:"additionalServices"`
	External            GetSellersOffersActionResOffersExternal            `json:"external" yaml:"external"`
	Delivery            GetSellersOffersActionResOffersDelivery            `json:"delivery" yaml:"delivery"`
	B2b                 GetSellersOffersActionResOffersB2b                 `json:"b2b" yaml:"b2b"`
	FundraisingCampaign GetSellersOffersActionResOffersFundraisingCampaign `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	// Marketplace-specific extensions for offer
	AdditionalMarketplaces emigo.Nullable[map[any]any] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for category
type GetSellersOffersActionResOffersCategory struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for primaryImage
type GetSellersOffersActionResOffersPrimaryImage struct {
	Url string `json:"url" yaml:"url"`
}

// The base class definition for sellingMode
type GetSellersOffersActionResOffersSellingMode struct {
	Format          string                                                    `json:"format" yaml:"format"`
	Price           GetSellersOffersActionResOffersSellingModePrice           `json:"price" yaml:"price"`
	PriceAutomation GetSellersOffersActionResOffersSellingModePriceAutomation `json:"priceAutomation" yaml:"priceAutomation"`
	MinimalPrice    GetSellersOffersActionResOffersSellingModeMinimalPrice    `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice   GetSellersOffersActionResOffersSellingModeStartingPrice   `json:"startingPrice" yaml:"startingPrice"`
}

// The base class definition for price
type GetSellersOffersActionResOffersSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for priceAutomation
type GetSellersOffersActionResOffersSellingModePriceAutomation struct {
	Rule GetSellersOffersActionResOffersSellingModePriceAutomationRule `json:"rule" yaml:"rule"`
}

// The base class definition for rule
type GetSellersOffersActionResOffersSellingModePriceAutomationRule struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for minimalPrice
type GetSellersOffersActionResOffersSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for startingPrice
type GetSellersOffersActionResOffersSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for saleInfo
type GetSellersOffersActionResOffersSaleInfo struct {
	CurrentPrice GetSellersOffersActionResOffersSaleInfoCurrentPrice `json:"currentPrice" yaml:"currentPrice"`
	BiddersCount int                                                 `json:"biddersCount" yaml:"biddersCount"`
}

// The base class definition for currentPrice
type GetSellersOffersActionResOffersSaleInfoCurrentPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for stock
type GetSellersOffersActionResOffersStock struct {
	Available int `json:"available" yaml:"available"`
	Sold      int `json:"sold" yaml:"sold"`
}

// The base class definition for stats
type GetSellersOffersActionResOffersStats struct {
	WatchersCount int `json:"watchersCount" yaml:"watchersCount"`
	VisitsCount   int `json:"visitsCount" yaml:"visitsCount"`
}

// The base class definition for publication
type GetSellersOffersActionResOffersPublication struct {
	Status       string                                                 `json:"status" yaml:"status"`
	StartingAt   string                                                 `json:"startingAt" yaml:"startingAt"`
	StartedAt    string                                                 `json:"startedAt" yaml:"startedAt"`
	EndingAt     string                                                 `json:"endingAt" yaml:"endingAt"`
	EndedAt      string                                                 `json:"endedAt" yaml:"endedAt"`
	Marketplaces GetSellersOffersActionResOffersPublicationMarketplaces `json:"marketplaces" yaml:"marketplaces"`
}

// The base class definition for marketplaces
type GetSellersOffersActionResOffersPublicationMarketplaces struct {
	Base       GetSellersOffersActionResOffersPublicationMarketplacesBase                    `json:"base" yaml:"base"`
	Additional emigo.Array[GetSellersOffersActionResOffersPublicationMarketplacesAdditional] `json:"additional" yaml:"additional"`
}

// The base class definition for base
type GetSellersOffersActionResOffersPublicationMarketplacesBase struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for additional
type GetSellersOffersActionResOffersPublicationMarketplacesAdditional struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for afterSalesServices
type GetSellersOffersActionResOffersAfterSalesServices struct {
	ImpliedWarranty GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        GetSellersOffersActionResOffersAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

// The base class definition for impliedWarranty
type GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for returnPolicy
type GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for warranty
type GetSellersOffersActionResOffersAfterSalesServicesWarranty struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for additionalServices
type GetSellersOffersActionResOffersAdditionalServices struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for external
type GetSellersOffersActionResOffersExternal struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for delivery
type GetSellersOffersActionResOffersDelivery struct {
	ShippingRates GetSellersOffersActionResOffersDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

// The base class definition for shippingRates
type GetSellersOffersActionResOffersDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for b2b
type GetSellersOffersActionResOffersB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

// The base class definition for fundraisingCampaign
type GetSellersOffersActionResOffersFundraisingCampaign struct {
	Id string `json:"id" yaml:"id"`
}

func (x *GetSellersOffersActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetSellersOffersActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetSellersOffersActionResponse) SetContentType(contentType string) *GetSellersOffersActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetSellersOffersActionResponse) AsStream(r io.Reader, contentType string) *GetSellersOffersActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetSellersOffersActionResponse) AsJSON(payload any) *GetSellersOffersActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetSellersOffersActionResponse) WithIdeal(payload GetSellersOffersActionRes) *GetSellersOffersActionResponse {
	x.Payload = payload
	return x
}
func (x *GetSellersOffersActionResponse) AsHTML(payload string) *GetSellersOffersActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetSellersOffersActionResponse) AsBytes(payload []byte) *GetSellersOffersActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetSellersOffersActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetSellersOffersActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetSellersOffersActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type GetSellersOffersActionRequestSig = func(c GetSellersOffersActionRequest) (*GetSellersOffersActionResponse, error)

/**
 * Query parameters for Get sellers offersAction
 */
// Query wrapper with private fields
type GetSellersOffersActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetSellersOffersActionQueryFromString(rawQuery string) GetSellersOffersActionQuery {
	v := GetSellersOffersActionQuery{}
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
func GetSellersOffersActionQueryFromHttp(r *http.Request) GetSellersOffersActionQuery {
	return GetSellersOffersActionQueryFromString(r.URL.RawQuery)
}
func (q GetSellersOffersActionQuery) Values() url.Values {
	return q.values
}
func (q GetSellersOffersActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetSellersOffersActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetSellersOffersActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetSellersOffersActionRequest struct {
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

func GetSellersOffersActionClientCreateUrl(
	req GetSellersOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetSellersOffersActionMeta()
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
func GetSellersOffersActionClientExecuteTyped(httpReq *http.Request) (*GetSellersOffersActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetSellersOffersActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetSellersOffersActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetSellersOffersActionResponse{Payload: result}, err
	}
	return &GetSellersOffersActionResponse{Payload: result}, nil
}
func GetSellersOffersActionClientBuildRequest(req GetSellersOffersActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetSellersOffersActionMeta()
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
func GetSellersOffersActionCall(
	req GetSellersOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetSellersOffersActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetSellersOffersActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetSellersOffersActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetSellersOffersActionClientExecuteTyped(r)
}

// GetSellersOffersActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the GetSellersOffersAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *GetSellersOffersActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func GetSellersOffersActionHttpHandler(
	handler func(c GetSellersOffersActionRequest) (*GetSellersOffersActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := GetSellersOffersActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := GetSellersOffersActionRequest{
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

// GetSellersOffersActionHttp is a high-level convenience wrapper around
// GetSellersOffersActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func GetSellersOffersActionHttp(
	mux *http.ServeMux,
	handler func(c GetSellersOffersActionRequest) (*GetSellersOffersActionResponse, error),
) {
	method, pattern, h := GetSellersOffersActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
