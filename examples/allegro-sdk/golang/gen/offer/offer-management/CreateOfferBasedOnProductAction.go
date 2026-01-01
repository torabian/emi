package external
import (
"bytes"
"encoding/json"
"fmt"
"github.com/gin-gonic/gin"
"io"
"net/http"
"net/url"
"test.com/emi/golang/emigo"
)
/**
* Action to communicate with the action CreateOfferBasedOnProductAction
*/
func CreateOfferBasedOnProductActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "CreateOfferBasedOnProductAction",
        URL:    "https://api.{environment}/sale/product-offers",
        Method: "POST",
    }
}
  // The base class definition for createOfferBasedOnProductActionReq
type CreateOfferBasedOnProductActionReq struct {
		  // Offer title
 Name string `json:"name" yaml:"name"`
		  // Offer language code (e.g., pl-PL)
 Language string `json:"language" yaml:"language"`
		Category  CreateOfferBasedOnProductActionReqCategory `json:"category" yaml:"category"`
		  // Product details and associated quantities
 ProductSet []CreateOfferBasedOnProductActionReqProductSet `json:"productSet" yaml:"productSet"`
		Stock  CreateOfferBasedOnProductActionReqStock `json:"stock" yaml:"stock"`
		SellingMode  CreateOfferBasedOnProductActionReqSellingMode `json:"sellingMode" yaml:"sellingMode"`
		Payments  CreateOfferBasedOnProductActionReqPayments `json:"payments" yaml:"payments"`
		Delivery  CreateOfferBasedOnProductActionReqDelivery `json:"delivery" yaml:"delivery"`
		Publication  CreateOfferBasedOnProductActionReqPublication `json:"publication" yaml:"publication"`
		AdditionalMarketplaces emigo.Nullable[interface{}] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
		CompatibilityList  CreateOfferBasedOnProductActionReqCompatibilityList `json:"compatibilityList" yaml:"compatibilityList"`
		Images []string `json:"images" yaml:"images"`
		Description  CreateOfferBasedOnProductActionReqDescription `json:"description" yaml:"description"`
		B2b  CreateOfferBasedOnProductActionReqB2b `json:"b2b" yaml:"b2b"`
		Attachments []CreateOfferBasedOnProductActionReqAttachments `json:"attachments" yaml:"attachments"`
		FundraisingCampaign  CreateOfferBasedOnProductActionReqFundraisingCampaign `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
		AdditionalServices  CreateOfferBasedOnProductActionReqAdditionalServices `json:"additionalServices" yaml:"additionalServices"`
		AfterSalesServices  CreateOfferBasedOnProductActionReqAfterSalesServices `json:"afterSalesServices" yaml:"afterSalesServices"`
		SizeTable  CreateOfferBasedOnProductActionReqSizeTable `json:"sizeTable" yaml:"sizeTable"`
		Contact  CreateOfferBasedOnProductActionReqContact `json:"contact" yaml:"contact"`
		Discounts  CreateOfferBasedOnProductActionReqDiscounts `json:"discounts" yaml:"discounts"`
		Location  CreateOfferBasedOnProductActionReqLocation `json:"location" yaml:"location"`
		External  CreateOfferBasedOnProductActionReqExternal `json:"external" yaml:"external"`
		TaxSettings  CreateOfferBasedOnProductActionReqTaxSettings `json:"taxSettings" yaml:"taxSettings"`
		MessageToSellerSettings  CreateOfferBasedOnProductActionReqMessageToSellerSettings `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
}
  // The base class definition for category
type CreateOfferBasedOnProductActionReqCategory struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for productSet
type CreateOfferBasedOnProductActionReqProductSet struct {
		Product  CreateOfferBasedOnProductActionReqProductSetProduct `json:"product" yaml:"product"`
		Quantity  CreateOfferBasedOnProductActionReqProductSetQuantity `json:"quantity" yaml:"quantity"`
		ResponsiblePerson  CreateOfferBasedOnProductActionReqProductSetResponsiblePerson `json:"responsiblePerson" yaml:"responsiblePerson"`
		ResponsibleProducer  CreateOfferBasedOnProductActionReqProductSetResponsibleProducer `json:"responsibleProducer" yaml:"responsibleProducer"`
		SafetyInformation  CreateOfferBasedOnProductActionReqProductSetSafetyInformation `json:"safetyInformation" yaml:"safetyInformation"`
		MarketedBeforeGPSRObligation bool `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
		Deposits []CreateOfferBasedOnProductActionReqProductSetDeposits `json:"deposits" yaml:"deposits"`
}
  // The base class definition for product
type CreateOfferBasedOnProductActionReqProductSetProduct struct {
		Id string `json:"id" yaml:"id"`
		IdType string `json:"idType" yaml:"idType"`
		Name string `json:"name" yaml:"name"`
		Category  CreateOfferBasedOnProductActionReqProductSetProductCategory `json:"category" yaml:"category"`
		Parameters []CreateOfferBasedOnProductActionReqProductSetProductParameters `json:"parameters" yaml:"parameters"`
		Images []string `json:"images" yaml:"images"`
}
  // The base class definition for category
type CreateOfferBasedOnProductActionReqProductSetProductCategory struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for parameters
type CreateOfferBasedOnProductActionReqProductSetProductParameters struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
		RangeValue  CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue `json:"rangeValue" yaml:"rangeValue"`
		Values []string `json:"values" yaml:"values"`
		ValuesIds []string `json:"valuesIds" yaml:"valuesIds"`
}
  // The base class definition for rangeValue
type CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue struct {
		From string `json:"from" yaml:"from"`
		To string `json:"to" yaml:"to"`
}
  // The base class definition for quantity
type CreateOfferBasedOnProductActionReqProductSetQuantity struct {
		Value int `json:"value" yaml:"value"`
}
  // The base class definition for responsiblePerson
type CreateOfferBasedOnProductActionReqProductSetResponsiblePerson struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
}
  // The base class definition for responsibleProducer
type CreateOfferBasedOnProductActionReqProductSetResponsibleProducer struct {
		Id string `json:"id" yaml:"id"`
		Type string `json:"type" yaml:"type"`
}
  // The base class definition for safetyInformation
type CreateOfferBasedOnProductActionReqProductSetSafetyInformation struct {
		Type string `json:"type" yaml:"type"`
		Description string `json:"description" yaml:"description"`
}
  // The base class definition for deposits
type CreateOfferBasedOnProductActionReqProductSetDeposits struct {
		Id string `json:"id" yaml:"id"`
		Quantity int `json:"quantity" yaml:"quantity"`
}
  // The base class definition for stock
type CreateOfferBasedOnProductActionReqStock struct {
		Available int `json:"available" yaml:"available"`
		Unit string `json:"unit" yaml:"unit"`
}
  // The base class definition for sellingMode
type CreateOfferBasedOnProductActionReqSellingMode struct {
		Format string `json:"format" yaml:"format"`
		Price  CreateOfferBasedOnProductActionReqSellingModePrice `json:"price" yaml:"price"`
		MinimalPrice  CreateOfferBasedOnProductActionReqSellingModeMinimalPrice `json:"minimalPrice" yaml:"minimalPrice"`
		StartingPrice  CreateOfferBasedOnProductActionReqSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}
  // The base class definition for price
type CreateOfferBasedOnProductActionReqSellingModePrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for minimalPrice
type CreateOfferBasedOnProductActionReqSellingModeMinimalPrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for startingPrice
type CreateOfferBasedOnProductActionReqSellingModeStartingPrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for payments
type CreateOfferBasedOnProductActionReqPayments struct {
		Invoice string `json:"invoice" yaml:"invoice"`
}
  // The base class definition for delivery
type CreateOfferBasedOnProductActionReqDelivery struct {
		HandlingTime string `json:"handlingTime" yaml:"handlingTime"`
		AdditionalInfo string `json:"additionalInfo" yaml:"additionalInfo"`
		ShipmentDate string `json:"shipmentDate" yaml:"shipmentDate"`
		  // Optional; may be null
 ShippingRates  CreateOfferBasedOnProductActionReqDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}
  // The base class definition for shippingRates
type CreateOfferBasedOnProductActionReqDeliveryShippingRates struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for publication
type CreateOfferBasedOnProductActionReqPublication struct {
		Duration string `json:"duration" yaml:"duration"`
		StartingAt string `json:"startingAt" yaml:"startingAt"`
		EndingAt string `json:"endingAt" yaml:"endingAt"`
		Status string `json:"status" yaml:"status"`
		Republish bool `json:"republish" yaml:"republish"`
}
  // The base class definition for compatibilityList
type CreateOfferBasedOnProductActionReqCompatibilityList struct {
		Items []CreateOfferBasedOnProductActionReqCompatibilityListItems `json:"items" yaml:"items"`
}
  // The base class definition for items
type CreateOfferBasedOnProductActionReqCompatibilityListItems struct {
		Type string `json:"type" yaml:"type"`
		Text string `json:"text" yaml:"text"`
}
  // The base class definition for description
type CreateOfferBasedOnProductActionReqDescription struct {
		Sections []CreateOfferBasedOnProductActionReqDescriptionSections `json:"sections" yaml:"sections"`
}
  // The base class definition for sections
type CreateOfferBasedOnProductActionReqDescriptionSections struct {
		Items []CreateOfferBasedOnProductActionReqDescriptionSectionsItems `json:"items" yaml:"items"`
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
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
}
  // The base class definition for additionalServices
type CreateOfferBasedOnProductActionReqAdditionalServices struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
}
  // The base class definition for afterSalesServices
type CreateOfferBasedOnProductActionReqAfterSalesServices struct {
		ImpliedWarranty  CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
		ReturnPolicy  CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy `json:"returnPolicy" yaml:"returnPolicy"`
		Warranty  CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty `json:"warranty" yaml:"warranty"`
}
  // The base class definition for impliedWarranty
type CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
}
  // The base class definition for returnPolicy
type CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
}
  // The base class definition for warranty
type CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
}
  // The base class definition for sizeTable
type CreateOfferBasedOnProductActionReqSizeTable struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
}
  // The base class definition for contact
type CreateOfferBasedOnProductActionReqContact struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
}
  // The base class definition for discounts
type CreateOfferBasedOnProductActionReqDiscounts struct {
		WholesalePriceList  CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}
  // The base class definition for wholesalePriceList
type CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
}
  // The base class definition for location
type CreateOfferBasedOnProductActionReqLocation struct {
		City string `json:"city" yaml:"city"`
		CountryCode string `json:"countryCode" yaml:"countryCode"`
		PostCode string `json:"postCode" yaml:"postCode"`
		Province string `json:"province" yaml:"province"`
}
  // The base class definition for external
type CreateOfferBasedOnProductActionReqExternal struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for taxSettings
type CreateOfferBasedOnProductActionReqTaxSettings struct {
		Subject string `json:"subject" yaml:"subject"`
		Exemption string `json:"exemption" yaml:"exemption"`
		Rates []CreateOfferBasedOnProductActionReqTaxSettingsRates `json:"rates" yaml:"rates"`
}
  // The base class definition for rates
type CreateOfferBasedOnProductActionReqTaxSettingsRates struct {
		Rate string `json:"rate" yaml:"rate"`
		CountryCode string `json:"countryCode" yaml:"countryCode"`
}
  // The base class definition for messageToSellerSettings
type CreateOfferBasedOnProductActionReqMessageToSellerSettings struct {
		Mode string `json:"mode" yaml:"mode"`
		Hint string `json:"hint" yaml:"hint"`
}
type CreateOfferBasedOnProductActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// CreateOfferBasedOnProductActionRaw registers a raw Gin route for the CreateOfferBasedOnProductAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CreateOfferBasedOnProductActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CreateOfferBasedOnProductActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// CreateOfferBasedOnProductActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CreateOfferBasedOnProductAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CreateOfferBasedOnProductActionHandler(
	handler func(c CreateOfferBasedOnProductActionRequest, gin *gin.Context) (*CreateOfferBasedOnProductActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := CreateOfferBasedOnProductActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body CreateOfferBasedOnProductActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := CreateOfferBasedOnProductActionRequest{
			Body:        body,
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
		}
		resp, err := handler(req, m)
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
// CreateOfferBasedOnProductAction is a high-level convenience wrapper around CreateOfferBasedOnProductActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CreateOfferBasedOnProductAction(r gin.IRoutes, handler func(c CreateOfferBasedOnProductActionRequest, gin *gin.Context) (*CreateOfferBasedOnProductActionResponse, error),) {
	method, url, h := CreateOfferBasedOnProductActionHandler(handler)
	r.Handle(method, url, h)
}
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
func CreateOfferBasedOnProductActionQueryFromGin(c *gin.Context) CreateOfferBasedOnProductActionQuery {
	return CreateOfferBasedOnProductActionQueryFromString(c.Request.URL.RawQuery)
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
	Body CreateOfferBasedOnProductActionReq
	QueryParams url.Values
	Headers http.Header
}
type CreateOfferBasedOnProductActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func CreateOfferBasedOnProductActionCall(
	req CreateOfferBasedOnProductActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CreateOfferBasedOnProductActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := CreateOfferBasedOnProductActionMeta()
		baseURL := meta.URL
		// Build final URL with query string
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}
		// if UrlValues present, encode and append
		if len(req.QueryParams) > 0 {
			u.RawQuery = req.QueryParams.Encode()
		}
			bodyBytes, err := json.Marshal(req.Body)
			if err != nil {
				return nil, err
			}
			req0, err := http.NewRequest(meta.Method, u.String(), bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, err
		}
		httpReq = req0
	} else {
		httpReq = config.Httpr
	}
	httpReq.Header = req.Headers
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var result CreateOfferBasedOnProductActionResult
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	if resp.StatusCode >= 400 {
		return &result, fmt.Errorf("request failed: %s", respBody)
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &result, err
	}
	return &result, nil
}