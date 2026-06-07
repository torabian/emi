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
* Action to communicate with the action CreateOfferBasedOnProductAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CreateOfferBasedOnProductAction
func CreateOfferBasedOnProductAction(c CreateOfferBasedOnProductActionRequest) (*CreateOfferBasedOnProductActionResponse, error) {
	return &CreateOfferBasedOnProductActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CreateOfferBasedOnProductActionMeta() struct {
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
		Name:        "CreateOfferBasedOnProductAction",
		CliName:     "create offer based on product-action",
		URL:         "https://api.{environment}/sale/product-offers",
		Method:      "POST",
		Description: ``,
	}
}

// The base class definition for createOfferBasedOnProductActionReq
type CreateOfferBasedOnProductActionReq struct {
	// Offer title
	Name string `json:"name" yaml:"name"`
	// Offer language code (e.g., pl-PL)
	Language string                                     `json:"language" yaml:"language"`
	Category CreateOfferBasedOnProductActionReqCategory `json:"category" yaml:"category"`
	// Product details and associated quantities
	ProductSet              emigo.Array[CreateOfferBasedOnProductActionReqProductSet]  `json:"productSet" yaml:"productSet"`
	Stock                   CreateOfferBasedOnProductActionReqStock                    `json:"stock" yaml:"stock"`
	SellingMode             CreateOfferBasedOnProductActionReqSellingMode              `json:"sellingMode" yaml:"sellingMode"`
	Payments                CreateOfferBasedOnProductActionReqPayments                 `json:"payments" yaml:"payments"`
	Delivery                CreateOfferBasedOnProductActionReqDelivery                 `json:"delivery" yaml:"delivery"`
	Publication             CreateOfferBasedOnProductActionReqPublication              `json:"publication" yaml:"publication"`
	AdditionalMarketplaces  emigo.Nullable[map[any]any]                                `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
	CompatibilityList       CreateOfferBasedOnProductActionReqCompatibilityList        `json:"compatibilityList" yaml:"compatibilityList"`
	Images                  []string                                                   `json:"images" yaml:"images"`
	Description             CreateOfferBasedOnProductActionReqDescription              `json:"description" yaml:"description"`
	B2b                     CreateOfferBasedOnProductActionReqB2b                      `json:"b2b" yaml:"b2b"`
	Attachments             emigo.Array[CreateOfferBasedOnProductActionReqAttachments] `json:"attachments" yaml:"attachments"`
	FundraisingCampaign     CreateOfferBasedOnProductActionReqFundraisingCampaign      `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	AdditionalServices      CreateOfferBasedOnProductActionReqAdditionalServices       `json:"additionalServices" yaml:"additionalServices"`
	AfterSalesServices      CreateOfferBasedOnProductActionReqAfterSalesServices       `json:"afterSalesServices" yaml:"afterSalesServices"`
	SizeTable               CreateOfferBasedOnProductActionReqSizeTable                `json:"sizeTable" yaml:"sizeTable"`
	Contact                 CreateOfferBasedOnProductActionReqContact                  `json:"contact" yaml:"contact"`
	Discounts               CreateOfferBasedOnProductActionReqDiscounts                `json:"discounts" yaml:"discounts"`
	Location                CreateOfferBasedOnProductActionReqLocation                 `json:"location" yaml:"location"`
	External                CreateOfferBasedOnProductActionReqExternal                 `json:"external" yaml:"external"`
	TaxSettings             CreateOfferBasedOnProductActionReqTaxSettings              `json:"taxSettings" yaml:"taxSettings"`
	MessageToSellerSettings CreateOfferBasedOnProductActionReqMessageToSellerSettings  `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
}

// The base class definition for category
type CreateOfferBasedOnProductActionReqCategory struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for productSet
type CreateOfferBasedOnProductActionReqProductSet struct {
	Product                      CreateOfferBasedOnProductActionReqProductSetProduct               `json:"product" yaml:"product"`
	Quantity                     CreateOfferBasedOnProductActionReqProductSetQuantity              `json:"quantity" yaml:"quantity"`
	ResponsiblePerson            CreateOfferBasedOnProductActionReqProductSetResponsiblePerson     `json:"responsiblePerson" yaml:"responsiblePerson"`
	ResponsibleProducer          CreateOfferBasedOnProductActionReqProductSetResponsibleProducer   `json:"responsibleProducer" yaml:"responsibleProducer"`
	SafetyInformation            CreateOfferBasedOnProductActionReqProductSetSafetyInformation     `json:"safetyInformation" yaml:"safetyInformation"`
	MarketedBeforeGPSRObligation bool                                                              `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
	Deposits                     emigo.Array[CreateOfferBasedOnProductActionReqProductSetDeposits] `json:"deposits" yaml:"deposits"`
}

// The base class definition for product
type CreateOfferBasedOnProductActionReqProductSetProduct struct {
	Id         string                                                                     `json:"id" yaml:"id"`
	IdType     string                                                                     `json:"idType" yaml:"idType"`
	Name       string                                                                     `json:"name" yaml:"name"`
	Category   CreateOfferBasedOnProductActionReqProductSetProductCategory                `json:"category" yaml:"category"`
	Parameters emigo.Array[CreateOfferBasedOnProductActionReqProductSetProductParameters] `json:"parameters" yaml:"parameters"`
	Images     []string                                                                   `json:"images" yaml:"images"`
}

// The base class definition for category
type CreateOfferBasedOnProductActionReqProductSetProductCategory struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for parameters
type CreateOfferBasedOnProductActionReqProductSetProductParameters struct {
	Id         string                                                                  `json:"id" yaml:"id"`
	Name       string                                                                  `json:"name" yaml:"name"`
	RangeValue CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue `json:"rangeValue" yaml:"rangeValue"`
	Values     []string                                                                `json:"values" yaml:"values"`
	ValuesIds  []string                                                                `json:"valuesIds" yaml:"valuesIds"`
}

// The base class definition for rangeValue
type CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

// The base class definition for quantity
type CreateOfferBasedOnProductActionReqProductSetQuantity struct {
	Value int `json:"value" yaml:"value"`
}

// The base class definition for responsiblePerson
type CreateOfferBasedOnProductActionReqProductSetResponsiblePerson struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for responsibleProducer
type CreateOfferBasedOnProductActionReqProductSetResponsibleProducer struct {
	Id   string `json:"id" yaml:"id"`
	Type string `json:"type" yaml:"type"`
}

// The base class definition for safetyInformation
type CreateOfferBasedOnProductActionReqProductSetSafetyInformation struct {
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description" yaml:"description"`
}

// The base class definition for deposits
type CreateOfferBasedOnProductActionReqProductSetDeposits struct {
	Id       string `json:"id" yaml:"id"`
	Quantity int    `json:"quantity" yaml:"quantity"`
}

// The base class definition for stock
type CreateOfferBasedOnProductActionReqStock struct {
	Available int    `json:"available" yaml:"available"`
	Unit      string `json:"unit" yaml:"unit"`
}

// The base class definition for sellingMode
type CreateOfferBasedOnProductActionReqSellingMode struct {
	Format        string                                                     `json:"format" yaml:"format"`
	Price         CreateOfferBasedOnProductActionReqSellingModePrice         `json:"price" yaml:"price"`
	MinimalPrice  CreateOfferBasedOnProductActionReqSellingModeMinimalPrice  `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice CreateOfferBasedOnProductActionReqSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}

// The base class definition for price
type CreateOfferBasedOnProductActionReqSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for minimalPrice
type CreateOfferBasedOnProductActionReqSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for startingPrice
type CreateOfferBasedOnProductActionReqSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for payments
type CreateOfferBasedOnProductActionReqPayments struct {
	Invoice string `json:"invoice" yaml:"invoice"`
}

// The base class definition for delivery
type CreateOfferBasedOnProductActionReqDelivery struct {
	HandlingTime   string `json:"handlingTime" yaml:"handlingTime"`
	AdditionalInfo string `json:"additionalInfo" yaml:"additionalInfo"`
	ShipmentDate   string `json:"shipmentDate" yaml:"shipmentDate"`
	// Optional; may be null
	ShippingRates CreateOfferBasedOnProductActionReqDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

// The base class definition for shippingRates
type CreateOfferBasedOnProductActionReqDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for publication
type CreateOfferBasedOnProductActionReqPublication struct {
	Duration   string `json:"duration" yaml:"duration"`
	StartingAt string `json:"startingAt" yaml:"startingAt"`
	EndingAt   string `json:"endingAt" yaml:"endingAt"`
	Status     string `json:"status" yaml:"status"`
	Republish  bool   `json:"republish" yaml:"republish"`
}

// The base class definition for compatibilityList
type CreateOfferBasedOnProductActionReqCompatibilityList struct {
	Items emigo.Array[CreateOfferBasedOnProductActionReqCompatibilityListItems] `json:"items" yaml:"items"`
}

// The base class definition for items
type CreateOfferBasedOnProductActionReqCompatibilityListItems struct {
	Type string `json:"type" yaml:"type"`
	Text string `json:"text" yaml:"text"`
}

// The base class definition for description
type CreateOfferBasedOnProductActionReqDescription struct {
	Sections emigo.Array[CreateOfferBasedOnProductActionReqDescriptionSections] `json:"sections" yaml:"sections"`
}

// The base class definition for sections
type CreateOfferBasedOnProductActionReqDescriptionSections struct {
	Items emigo.Array[CreateOfferBasedOnProductActionReqDescriptionSectionsItems] `json:"items" yaml:"items"`
}

// The base class definition for items
type CreateOfferBasedOnProductActionReqDescriptionSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

// The base class definition for b2b
type CreateOfferBasedOnProductActionReqB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

// The base class definition for attachments
type CreateOfferBasedOnProductActionReqAttachments struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for fundraisingCampaign
type CreateOfferBasedOnProductActionReqFundraisingCampaign struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for additionalServices
type CreateOfferBasedOnProductActionReqAdditionalServices struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for afterSalesServices
type CreateOfferBasedOnProductActionReqAfterSalesServices struct {
	ImpliedWarranty CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

// The base class definition for impliedWarranty
type CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for returnPolicy
type CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for warranty
type CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for sizeTable
type CreateOfferBasedOnProductActionReqSizeTable struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for contact
type CreateOfferBasedOnProductActionReqContact struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for discounts
type CreateOfferBasedOnProductActionReqDiscounts struct {
	WholesalePriceList CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}

// The base class definition for wholesalePriceList
type CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for location
type CreateOfferBasedOnProductActionReqLocation struct {
	City        string `json:"city" yaml:"city"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
	PostCode    string `json:"postCode" yaml:"postCode"`
	Province    string `json:"province" yaml:"province"`
}

// The base class definition for external
type CreateOfferBasedOnProductActionReqExternal struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for taxSettings
type CreateOfferBasedOnProductActionReqTaxSettings struct {
	Subject   string                                                          `json:"subject" yaml:"subject"`
	Exemption string                                                          `json:"exemption" yaml:"exemption"`
	Rates     emigo.Array[CreateOfferBasedOnProductActionReqTaxSettingsRates] `json:"rates" yaml:"rates"`
}

// The base class definition for rates
type CreateOfferBasedOnProductActionReqTaxSettingsRates struct {
	Rate        string `json:"rate" yaml:"rate"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
}

// The base class definition for messageToSellerSettings
type CreateOfferBasedOnProductActionReqMessageToSellerSettings struct {
	Mode string `json:"mode" yaml:"mode"`
	Hint string `json:"hint" yaml:"hint"`
}

func (x *CreateOfferBasedOnProductActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CreateOfferBasedOnProductActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CreateOfferBasedOnProductActionResponse) SetContentType(contentType string) *CreateOfferBasedOnProductActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CreateOfferBasedOnProductActionResponse) AsStream(r io.Reader, contentType string) *CreateOfferBasedOnProductActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CreateOfferBasedOnProductActionResponse) AsJSON(payload any) *CreateOfferBasedOnProductActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *CreateOfferBasedOnProductActionResponse) AsHTML(payload string) *CreateOfferBasedOnProductActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CreateOfferBasedOnProductActionResponse) AsBytes(payload []byte) *CreateOfferBasedOnProductActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CreateOfferBasedOnProductActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CreateOfferBasedOnProductActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CreateOfferBasedOnProductActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type CreateOfferBasedOnProductActionRequestSig = func(c CreateOfferBasedOnProductActionRequest) (*CreateOfferBasedOnProductActionResponse, error)

/**
 * Query parameters for Create offer based on productAction
 */
// Query wrapper with private fields
type CreateOfferBasedOnProductActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CreateOfferBasedOnProductActionQueryFromString(rawQuery string) CreateOfferBasedOnProductActionQuery {
	v := CreateOfferBasedOnProductActionQuery{}
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
func CreateOfferBasedOnProductActionQueryFromHttp(r *http.Request) CreateOfferBasedOnProductActionQuery {
	return CreateOfferBasedOnProductActionQueryFromString(r.URL.RawQuery)
}
func (q CreateOfferBasedOnProductActionQuery) Values() url.Values {
	return q.values
}
func (q CreateOfferBasedOnProductActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CreateOfferBasedOnProductActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CreateOfferBasedOnProductActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CreateOfferBasedOnProductActionRequest struct {
	Body        CreateOfferBasedOnProductActionReq
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

func CreateOfferBasedOnProductActionClientCreateUrl(
	req CreateOfferBasedOnProductActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CreateOfferBasedOnProductActionMeta()
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
func CreateOfferBasedOnProductActionClientExecuteTyped(httpReq *http.Request) (*CreateOfferBasedOnProductActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CreateOfferBasedOnProductActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &CreateOfferBasedOnProductActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &CreateOfferBasedOnProductActionResponse{Payload: result}, err
	}
	return &CreateOfferBasedOnProductActionResponse{Payload: result}, nil
}
func CreateOfferBasedOnProductActionClientBuildRequest(req CreateOfferBasedOnProductActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CreateOfferBasedOnProductActionMeta()
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
func CreateOfferBasedOnProductActionCall(
	req CreateOfferBasedOnProductActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CreateOfferBasedOnProductActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CreateOfferBasedOnProductActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CreateOfferBasedOnProductActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CreateOfferBasedOnProductActionClientExecuteTyped(r)
}

// CreateOfferBasedOnProductActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the CreateOfferBasedOnProductAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *CreateOfferBasedOnProductActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func CreateOfferBasedOnProductActionHttpHandler(
	handler func(c CreateOfferBasedOnProductActionRequest) (*CreateOfferBasedOnProductActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := CreateOfferBasedOnProductActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body CreateOfferBasedOnProductActionReq
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
		req := CreateOfferBasedOnProductActionRequest{
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

// CreateOfferBasedOnProductActionHttp is a high-level convenience wrapper around
// CreateOfferBasedOnProductActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func CreateOfferBasedOnProductActionHttp(
	mux *http.ServeMux,
	handler func(c CreateOfferBasedOnProductActionRequest) (*CreateOfferBasedOnProductActionResponse, error),
) {
	method, pattern, h := CreateOfferBasedOnProductActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
