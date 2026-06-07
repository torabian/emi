package external

import (
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action GetAllDataOfTheParticularProductOfferAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetAllDataOfTheParticularProductOfferAction
func GetAllDataOfTheParticularProductOfferAction(c GetAllDataOfTheParticularProductOfferActionRequest) (*GetAllDataOfTheParticularProductOfferActionResponse, error) {
	return &GetAllDataOfTheParticularProductOfferActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetAllDataOfTheParticularProductOfferActionMeta() struct {
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
		Name:        "GetAllDataOfTheParticularProductOfferAction",
		CliName:     "get all data of the particular product-offer-action",
		URL:         "https://api.{environment}/sale/product-offers/{offerId}",
		Method:      "GET",
		Description: `Full response model returned by GET /sale/product-offers/{offerId}`,
	}
}

// The base class definition for getAllDataOfTheParticularProductOfferActionRes
type GetAllDataOfTheParticularProductOfferActionRes struct {
	// Unique offer identifier
	Id string `json:"id" yaml:"id"`
	// Offer title
	Name string `json:"name" yaml:"name"`
	// Offer language code (e.g. pl-PL)
	Language string `json:"language" yaml:"language"`
	// Offer creation timestamp (ISO8601)
	CreatedAt string `json:"createdAt" yaml:"createdAt"`
	// Offer last update timestamp (ISO8601)
	UpdatedAt               string                                                                 `json:"updatedAt" yaml:"updatedAt"`
	Category                GetAllDataOfTheParticularProductOfferActionResCategory                 `json:"category" yaml:"category"`
	Stock                   GetAllDataOfTheParticularProductOfferActionResStock                    `json:"stock" yaml:"stock"`
	Contact                 GetAllDataOfTheParticularProductOfferActionResContact                  `json:"contact" yaml:"contact"`
	Publication             GetAllDataOfTheParticularProductOfferActionResPublication              `json:"publication" yaml:"publication"`
	SellingMode             GetAllDataOfTheParticularProductOfferActionResSellingMode              `json:"sellingMode" yaml:"sellingMode"`
	Payments                GetAllDataOfTheParticularProductOfferActionResPayments                 `json:"payments" yaml:"payments"`
	Delivery                GetAllDataOfTheParticularProductOfferActionResDelivery                 `json:"delivery" yaml:"delivery"`
	AfterSalesServices      GetAllDataOfTheParticularProductOfferActionResAfterSalesServices       `json:"afterSalesServices" yaml:"afterSalesServices"`
	Discounts               GetAllDataOfTheParticularProductOfferActionResDiscounts                `json:"discounts" yaml:"discounts"`
	Description             GetAllDataOfTheParticularProductOfferActionResDescription              `json:"description" yaml:"description"`
	Images                  []string                                                               `json:"images" yaml:"images"`
	ProductSet              emigo.Array[GetAllDataOfTheParticularProductOfferActionResProductSet]  `json:"productSet" yaml:"productSet"`
	Attachments             emigo.Array[GetAllDataOfTheParticularProductOfferActionResAttachments] `json:"attachments" yaml:"attachments"`
	FundraisingCampaign     GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign      `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	AdditionalServices      GetAllDataOfTheParticularProductOfferActionResAdditionalServices       `json:"additionalServices" yaml:"additionalServices"`
	AdditionalMarketplaces  emigo.Nullable[map[any]any]                                            `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
	B2b                     GetAllDataOfTheParticularProductOfferActionResB2b                      `json:"b2b" yaml:"b2b"`
	CompatibilityList       GetAllDataOfTheParticularProductOfferActionResCompatibilityList        `json:"compatibilityList" yaml:"compatibilityList"`
	Validation              GetAllDataOfTheParticularProductOfferActionResValidation               `json:"validation" yaml:"validation"`
	External                GetAllDataOfTheParticularProductOfferActionResExternal                 `json:"external" yaml:"external"`
	SizeTable               GetAllDataOfTheParticularProductOfferActionResSizeTable                `json:"sizeTable" yaml:"sizeTable"`
	TaxSettings             GetAllDataOfTheParticularProductOfferActionResTaxSettings              `json:"taxSettings" yaml:"taxSettings"`
	MessageToSellerSettings GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings  `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
}

// The base class definition for category
type GetAllDataOfTheParticularProductOfferActionResCategory struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for stock
type GetAllDataOfTheParticularProductOfferActionResStock struct {
	Available int    `json:"available" yaml:"available"`
	Unit      string `json:"unit" yaml:"unit"`
}

// The base class definition for contact
type GetAllDataOfTheParticularProductOfferActionResContact struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for publication
type GetAllDataOfTheParticularProductOfferActionResPublication struct {
	Duration     string                                                                `json:"duration" yaml:"duration"`
	StartingAt   string                                                                `json:"startingAt" yaml:"startingAt"`
	EndingAt     string                                                                `json:"endingAt" yaml:"endingAt"`
	EndedBy      string                                                                `json:"endedBy" yaml:"endedBy"`
	Status       string                                                                `json:"status" yaml:"status"`
	Republish    bool                                                                  `json:"republish" yaml:"republish"`
	Marketplaces GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces `json:"marketplaces" yaml:"marketplaces"`
}

// The base class definition for marketplaces
type GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces struct {
	Base       GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase                    `json:"base" yaml:"base"`
	Additional emigo.Array[GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional] `json:"additional" yaml:"additional"`
}

// The base class definition for base
type GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for additional
type GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for sellingMode
type GetAllDataOfTheParticularProductOfferActionResSellingMode struct {
	Format        string                                                                 `json:"format" yaml:"format"`
	Price         GetAllDataOfTheParticularProductOfferActionResSellingModePrice         `json:"price" yaml:"price"`
	MinimalPrice  GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice  `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}

// The base class definition for price
type GetAllDataOfTheParticularProductOfferActionResSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for minimalPrice
type GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for startingPrice
type GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for payments
type GetAllDataOfTheParticularProductOfferActionResPayments struct {
	Invoice string `json:"invoice" yaml:"invoice"`
}

// The base class definition for delivery
type GetAllDataOfTheParticularProductOfferActionResDelivery struct {
	HandlingTime   string                                                              `json:"handlingTime" yaml:"handlingTime"`
	AdditionalInfo string                                                              `json:"additionalInfo" yaml:"additionalInfo"`
	ShipmentDate   string                                                              `json:"shipmentDate" yaml:"shipmentDate"`
	ShippingRates  GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

// The base class definition for shippingRates
type GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for afterSalesServices
type GetAllDataOfTheParticularProductOfferActionResAfterSalesServices struct {
	ImpliedWarranty GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

// The base class definition for impliedWarranty
type GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for returnPolicy
type GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for warranty
type GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for discounts
type GetAllDataOfTheParticularProductOfferActionResDiscounts struct {
	WholesalePriceList GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}

// The base class definition for wholesalePriceList
type GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for description
type GetAllDataOfTheParticularProductOfferActionResDescription struct {
	Sections emigo.Array[GetAllDataOfTheParticularProductOfferActionResDescriptionSections] `json:"sections" yaml:"sections"`
}

// The base class definition for sections
type GetAllDataOfTheParticularProductOfferActionResDescriptionSections struct {
	Items emigo.Array[GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems] `json:"items" yaml:"items"`
}

// The base class definition for items
type GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

// The base class definition for productSet
type GetAllDataOfTheParticularProductOfferActionResProductSet struct {
	Quantity                     GetAllDataOfTheParticularProductOfferActionResProductSetQuantity              `json:"quantity" yaml:"quantity"`
	Product                      GetAllDataOfTheParticularProductOfferActionResProductSetProduct               `json:"product" yaml:"product"`
	ResponsiblePerson            GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson     `json:"responsiblePerson" yaml:"responsiblePerson"`
	ResponsibleProducer          GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer   `json:"responsibleProducer" yaml:"responsibleProducer"`
	SafetyInformation            GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation     `json:"safetyInformation" yaml:"safetyInformation"`
	MarketedBeforeGPSRObligation bool                                                                          `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
	Deposits                     emigo.Array[GetAllDataOfTheParticularProductOfferActionResProductSetDeposits] `json:"deposits" yaml:"deposits"`
}

// The base class definition for quantity
type GetAllDataOfTheParticularProductOfferActionResProductSetQuantity struct {
	Value int `json:"value" yaml:"value"`
}

// The base class definition for product
type GetAllDataOfTheParticularProductOfferActionResProductSetProduct struct {
	Id            string                                                                                 `json:"id" yaml:"id"`
	IsAiCoCreated bool                                                                                   `json:"isAiCoCreated" yaml:"isAiCoCreated"`
	Publication   GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication             `json:"publication" yaml:"publication"`
	Parameters    emigo.Array[GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters] `json:"parameters" yaml:"parameters"`
}

// The base class definition for publication
type GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication struct {
	Status string `json:"status" yaml:"status"`
}

// The base class definition for parameters
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters struct {
	Id         string                                                                                          `json:"id" yaml:"id"`
	Name       string                                                                                          `json:"name" yaml:"name"`
	RangeValue GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue             `json:"rangeValue" yaml:"rangeValue"`
	Values     emigo.Array[GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues]    `json:"values" yaml:"values"`
	ValuesIds  emigo.Array[GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds] `json:"valuesIds" yaml:"valuesIds"`
}

// The base class definition for rangeValue
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

// The base class definition for values
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues struct {
}

// The base class definition for valuesIds
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds struct {
}

// The base class definition for responsiblePerson
type GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for responsibleProducer
type GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for safetyInformation
type GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation struct {
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description" yaml:"description"`
}

// The base class definition for deposits
type GetAllDataOfTheParticularProductOfferActionResProductSetDeposits struct {
	Id       string `json:"id" yaml:"id"`
	Quantity int    `json:"quantity" yaml:"quantity"`
}

// The base class definition for attachments
type GetAllDataOfTheParticularProductOfferActionResAttachments struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for fundraisingCampaign
type GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for additionalServices
type GetAllDataOfTheParticularProductOfferActionResAdditionalServices struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for b2b
type GetAllDataOfTheParticularProductOfferActionResB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

// The base class definition for compatibilityList
type GetAllDataOfTheParticularProductOfferActionResCompatibilityList struct {
	Type string `json:"type" yaml:"type"`
}

// The base class definition for validation
type GetAllDataOfTheParticularProductOfferActionResValidation struct {
	ValidatedAt string                                                                        `json:"validatedAt" yaml:"validatedAt"`
	Errors      emigo.Array[GetAllDataOfTheParticularProductOfferActionResValidationErrors]   `json:"errors" yaml:"errors"`
	Warnings    emigo.Array[GetAllDataOfTheParticularProductOfferActionResValidationWarnings] `json:"warnings" yaml:"warnings"`
}

// The base class definition for errors
type GetAllDataOfTheParticularProductOfferActionResValidationErrors struct {
	Code        string                                                                 `json:"code" yaml:"code"`
	Details     string                                                                 `json:"details" yaml:"details"`
	Message     string                                                                 `json:"message" yaml:"message"`
	Path        string                                                                 `json:"path" yaml:"path"`
	UserMessage string                                                                 `json:"userMessage" yaml:"userMessage"`
	Metadata    GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata `json:"metadata" yaml:"metadata"`
}

// The base class definition for metadata
type GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

// The base class definition for warnings
type GetAllDataOfTheParticularProductOfferActionResValidationWarnings struct {
	Code        string                                                                   `json:"code" yaml:"code"`
	Details     string                                                                   `json:"details" yaml:"details"`
	Message     string                                                                   `json:"message" yaml:"message"`
	Path        string                                                                   `json:"path" yaml:"path"`
	UserMessage string                                                                   `json:"userMessage" yaml:"userMessage"`
	Metadata    GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata `json:"metadata" yaml:"metadata"`
}

// The base class definition for metadata
type GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

// The base class definition for external
type GetAllDataOfTheParticularProductOfferActionResExternal struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for sizeTable
type GetAllDataOfTheParticularProductOfferActionResSizeTable struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for taxSettings
type GetAllDataOfTheParticularProductOfferActionResTaxSettings struct {
	Subject   string                                                                      `json:"subject" yaml:"subject"`
	Exemption string                                                                      `json:"exemption" yaml:"exemption"`
	Rates     emigo.Array[GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates] `json:"rates" yaml:"rates"`
}

// The base class definition for rates
type GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates struct {
	Rate        string `json:"rate" yaml:"rate"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
}

// The base class definition for messageToSellerSettings
type GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings struct {
	Mode string `json:"mode" yaml:"mode"`
	Hint string `json:"hint" yaml:"hint"`
}

func (x *GetAllDataOfTheParticularProductOfferActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetAllDataOfTheParticularProductOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetAllDataOfTheParticularProductOfferActionResponse) SetContentType(contentType string) *GetAllDataOfTheParticularProductOfferActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetAllDataOfTheParticularProductOfferActionResponse) AsStream(r io.Reader, contentType string) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetAllDataOfTheParticularProductOfferActionResponse) AsJSON(payload any) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetAllDataOfTheParticularProductOfferActionResponse) WithIdeal(payload GetAllDataOfTheParticularProductOfferActionRes) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	return x
}
func (x *GetAllDataOfTheParticularProductOfferActionResponse) AsHTML(payload string) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetAllDataOfTheParticularProductOfferActionResponse) AsBytes(payload []byte) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetAllDataOfTheParticularProductOfferActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetAllDataOfTheParticularProductOfferActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetAllDataOfTheParticularProductOfferActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type GetAllDataOfTheParticularProductOfferActionRequestSig = func(c GetAllDataOfTheParticularProductOfferActionRequest) (*GetAllDataOfTheParticularProductOfferActionResponse, error)

/**
 * Query parameters for Get all data of the particular product-offerAction
 */
// Query wrapper with private fields
type GetAllDataOfTheParticularProductOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetAllDataOfTheParticularProductOfferActionQueryFromString(rawQuery string) GetAllDataOfTheParticularProductOfferActionQuery {
	v := GetAllDataOfTheParticularProductOfferActionQuery{}
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
func GetAllDataOfTheParticularProductOfferActionQueryFromHttp(r *http.Request) GetAllDataOfTheParticularProductOfferActionQuery {
	return GetAllDataOfTheParticularProductOfferActionQueryFromString(r.URL.RawQuery)
}
func (q GetAllDataOfTheParticularProductOfferActionQuery) Values() url.Values {
	return q.values
}
func (q GetAllDataOfTheParticularProductOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetAllDataOfTheParticularProductOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetAllDataOfTheParticularProductOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetAllDataOfTheParticularProductOfferActionRequest struct {
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

func GetAllDataOfTheParticularProductOfferActionClientCreateUrl(
	req GetAllDataOfTheParticularProductOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
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
func GetAllDataOfTheParticularProductOfferActionClientExecuteTyped(httpReq *http.Request) (*GetAllDataOfTheParticularProductOfferActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetAllDataOfTheParticularProductOfferActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetAllDataOfTheParticularProductOfferActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetAllDataOfTheParticularProductOfferActionResponse{Payload: result}, err
	}
	return &GetAllDataOfTheParticularProductOfferActionResponse{Payload: result}, nil
}
func GetAllDataOfTheParticularProductOfferActionClientBuildRequest(req GetAllDataOfTheParticularProductOfferActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
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
func GetAllDataOfTheParticularProductOfferActionCall(
	req GetAllDataOfTheParticularProductOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetAllDataOfTheParticularProductOfferActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetAllDataOfTheParticularProductOfferActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetAllDataOfTheParticularProductOfferActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetAllDataOfTheParticularProductOfferActionClientExecuteTyped(r)
}

// GetAllDataOfTheParticularProductOfferActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the GetAllDataOfTheParticularProductOfferAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *GetAllDataOfTheParticularProductOfferActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func GetAllDataOfTheParticularProductOfferActionHttpHandler(
	handler func(c GetAllDataOfTheParticularProductOfferActionRequest) (*GetAllDataOfTheParticularProductOfferActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := GetAllDataOfTheParticularProductOfferActionRequest{
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

// GetAllDataOfTheParticularProductOfferActionHttp is a high-level convenience wrapper around
// GetAllDataOfTheParticularProductOfferActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func GetAllDataOfTheParticularProductOfferActionHttp(
	mux *http.ServeMux,
	handler func(c GetAllDataOfTheParticularProductOfferActionRequest) (*GetAllDataOfTheParticularProductOfferActionResponse, error),
) {
	method, pattern, h := GetAllDataOfTheParticularProductOfferActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
