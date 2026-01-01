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
* Action to communicate with the action GetAllDataOfTheParticularProductOfferAction
*/
func GetAllDataOfTheParticularProductOfferActionMeta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "GetAllDataOfTheParticularProductOfferAction",
        URL:    "https://api.{environment}/sale/product-offers/{offerId}",
        Method: "GET",
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
 UpdatedAt string `json:"updatedAt" yaml:"updatedAt"`
		Category  GetAllDataOfTheParticularProductOfferActionResCategory `json:"category" yaml:"category"`
		Stock  GetAllDataOfTheParticularProductOfferActionResStock `json:"stock" yaml:"stock"`
		Contact  GetAllDataOfTheParticularProductOfferActionResContact `json:"contact" yaml:"contact"`
		Publication  GetAllDataOfTheParticularProductOfferActionResPublication `json:"publication" yaml:"publication"`
		SellingMode  GetAllDataOfTheParticularProductOfferActionResSellingMode `json:"sellingMode" yaml:"sellingMode"`
		Payments  GetAllDataOfTheParticularProductOfferActionResPayments `json:"payments" yaml:"payments"`
		Delivery  GetAllDataOfTheParticularProductOfferActionResDelivery `json:"delivery" yaml:"delivery"`
		AfterSalesServices  GetAllDataOfTheParticularProductOfferActionResAfterSalesServices `json:"afterSalesServices" yaml:"afterSalesServices"`
		Discounts  GetAllDataOfTheParticularProductOfferActionResDiscounts `json:"discounts" yaml:"discounts"`
		Description  GetAllDataOfTheParticularProductOfferActionResDescription `json:"description" yaml:"description"`
		Images []string `json:"images" yaml:"images"`
		ProductSet []GetAllDataOfTheParticularProductOfferActionResProductSet `json:"productSet" yaml:"productSet"`
		Attachments []GetAllDataOfTheParticularProductOfferActionResAttachments `json:"attachments" yaml:"attachments"`
		FundraisingCampaign  GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
		AdditionalServices  GetAllDataOfTheParticularProductOfferActionResAdditionalServices `json:"additionalServices" yaml:"additionalServices"`
		AdditionalMarketplaces emigo.Nullable[interface{}] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
		B2b  GetAllDataOfTheParticularProductOfferActionResB2b `json:"b2b" yaml:"b2b"`
		CompatibilityList  GetAllDataOfTheParticularProductOfferActionResCompatibilityList `json:"compatibilityList" yaml:"compatibilityList"`
		Validation  GetAllDataOfTheParticularProductOfferActionResValidation `json:"validation" yaml:"validation"`
		External  GetAllDataOfTheParticularProductOfferActionResExternal `json:"external" yaml:"external"`
		SizeTable  GetAllDataOfTheParticularProductOfferActionResSizeTable `json:"sizeTable" yaml:"sizeTable"`
		TaxSettings  GetAllDataOfTheParticularProductOfferActionResTaxSettings `json:"taxSettings" yaml:"taxSettings"`
		MessageToSellerSettings  GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
}
  // The base class definition for category
type GetAllDataOfTheParticularProductOfferActionResCategory struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for stock
type GetAllDataOfTheParticularProductOfferActionResStock struct {
		Available int `json:"available" yaml:"available"`
		Unit string `json:"unit" yaml:"unit"`
}
  // The base class definition for contact
type GetAllDataOfTheParticularProductOfferActionResContact struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for publication
type GetAllDataOfTheParticularProductOfferActionResPublication struct {
		Duration string `json:"duration" yaml:"duration"`
		StartingAt string `json:"startingAt" yaml:"startingAt"`
		EndingAt string `json:"endingAt" yaml:"endingAt"`
		EndedBy string `json:"endedBy" yaml:"endedBy"`
		Status string `json:"status" yaml:"status"`
		Republish bool `json:"republish" yaml:"republish"`
		Marketplaces  GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces `json:"marketplaces" yaml:"marketplaces"`
}
  // The base class definition for marketplaces
type GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces struct {
		Base  GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase `json:"base" yaml:"base"`
		Additional []GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional `json:"additional" yaml:"additional"`
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
		Format string `json:"format" yaml:"format"`
		Price  GetAllDataOfTheParticularProductOfferActionResSellingModePrice `json:"price" yaml:"price"`
		MinimalPrice  GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice `json:"minimalPrice" yaml:"minimalPrice"`
		StartingPrice  GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}
  // The base class definition for price
type GetAllDataOfTheParticularProductOfferActionResSellingModePrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for minimalPrice
type GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for startingPrice
type GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice struct {
		Amount string `json:"amount" yaml:"amount"`
		Currency string `json:"currency" yaml:"currency"`
}
  // The base class definition for payments
type GetAllDataOfTheParticularProductOfferActionResPayments struct {
		Invoice string `json:"invoice" yaml:"invoice"`
}
  // The base class definition for delivery
type GetAllDataOfTheParticularProductOfferActionResDelivery struct {
		HandlingTime string `json:"handlingTime" yaml:"handlingTime"`
		AdditionalInfo string `json:"additionalInfo" yaml:"additionalInfo"`
		ShipmentDate string `json:"shipmentDate" yaml:"shipmentDate"`
		ShippingRates  GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}
  // The base class definition for shippingRates
type GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for afterSalesServices
type GetAllDataOfTheParticularProductOfferActionResAfterSalesServices struct {
		ImpliedWarranty  GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
		ReturnPolicy  GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy `json:"returnPolicy" yaml:"returnPolicy"`
		Warranty  GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty `json:"warranty" yaml:"warranty"`
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
		WholesalePriceList  GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}
  // The base class definition for wholesalePriceList
type GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList struct {
		Id string `json:"id" yaml:"id"`
}
  // The base class definition for description
type GetAllDataOfTheParticularProductOfferActionResDescription struct {
		Sections []GetAllDataOfTheParticularProductOfferActionResDescriptionSections `json:"sections" yaml:"sections"`
}
  // The base class definition for sections
type GetAllDataOfTheParticularProductOfferActionResDescriptionSections struct {
		Items []GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems `json:"items" yaml:"items"`
}
  // The base class definition for items
type GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems struct {
		Type string `json:"type" yaml:"type"`
}
  // The base class definition for productSet
type GetAllDataOfTheParticularProductOfferActionResProductSet struct {
		Quantity  GetAllDataOfTheParticularProductOfferActionResProductSetQuantity `json:"quantity" yaml:"quantity"`
		Product  GetAllDataOfTheParticularProductOfferActionResProductSetProduct `json:"product" yaml:"product"`
		ResponsiblePerson  GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson `json:"responsiblePerson" yaml:"responsiblePerson"`
		ResponsibleProducer  GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer `json:"responsibleProducer" yaml:"responsibleProducer"`
		SafetyInformation  GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation `json:"safetyInformation" yaml:"safetyInformation"`
		MarketedBeforeGPSRObligation bool `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
		Deposits []GetAllDataOfTheParticularProductOfferActionResProductSetDeposits `json:"deposits" yaml:"deposits"`
}
  // The base class definition for quantity
type GetAllDataOfTheParticularProductOfferActionResProductSetQuantity struct {
		Value int `json:"value" yaml:"value"`
}
  // The base class definition for product
type GetAllDataOfTheParticularProductOfferActionResProductSetProduct struct {
		Id string `json:"id" yaml:"id"`
		IsAiCoCreated bool `json:"isAiCoCreated" yaml:"isAiCoCreated"`
		Publication  GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication `json:"publication" yaml:"publication"`
		Parameters []GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters `json:"parameters" yaml:"parameters"`
}
  // The base class definition for publication
type GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication struct {
		Status string `json:"status" yaml:"status"`
}
  // The base class definition for parameters
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters struct {
		Id string `json:"id" yaml:"id"`
		Name string `json:"name" yaml:"name"`
		RangeValue  GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue `json:"rangeValue" yaml:"rangeValue"`
		Values []GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues `json:"values" yaml:"values"`
		ValuesIds []GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds `json:"valuesIds" yaml:"valuesIds"`
}
  // The base class definition for rangeValue
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue struct {
		From string `json:"from" yaml:"from"`
		To string `json:"to" yaml:"to"`
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
		Type string `json:"type" yaml:"type"`
		Description string `json:"description" yaml:"description"`
}
  // The base class definition for deposits
type GetAllDataOfTheParticularProductOfferActionResProductSetDeposits struct {
		Id string `json:"id" yaml:"id"`
		Quantity int `json:"quantity" yaml:"quantity"`
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
		ValidatedAt string `json:"validatedAt" yaml:"validatedAt"`
		Errors []GetAllDataOfTheParticularProductOfferActionResValidationErrors `json:"errors" yaml:"errors"`
		Warnings []GetAllDataOfTheParticularProductOfferActionResValidationWarnings `json:"warnings" yaml:"warnings"`
}
  // The base class definition for errors
type GetAllDataOfTheParticularProductOfferActionResValidationErrors struct {
		Code string `json:"code" yaml:"code"`
		Details string `json:"details" yaml:"details"`
		Message string `json:"message" yaml:"message"`
		Path string `json:"path" yaml:"path"`
		UserMessage string `json:"userMessage" yaml:"userMessage"`
		Metadata  GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata `json:"metadata" yaml:"metadata"`
}
  // The base class definition for metadata
type GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata struct {
		ProductId string `json:"productId" yaml:"productId"`
}
  // The base class definition for warnings
type GetAllDataOfTheParticularProductOfferActionResValidationWarnings struct {
		Code string `json:"code" yaml:"code"`
		Details string `json:"details" yaml:"details"`
		Message string `json:"message" yaml:"message"`
		Path string `json:"path" yaml:"path"`
		UserMessage string `json:"userMessage" yaml:"userMessage"`
		Metadata  GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata `json:"metadata" yaml:"metadata"`
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
		Subject string `json:"subject" yaml:"subject"`
		Exemption string `json:"exemption" yaml:"exemption"`
		Rates []GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates `json:"rates" yaml:"rates"`
}
  // The base class definition for rates
type GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates struct {
		Rate string `json:"rate" yaml:"rate"`
		CountryCode string `json:"countryCode" yaml:"countryCode"`
}
  // The base class definition for messageToSellerSettings
type GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings struct {
		Mode string `json:"mode" yaml:"mode"`
		Hint string `json:"hint" yaml:"hint"`
}
type GetAllDataOfTheParticularProductOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// GetAllDataOfTheParticularProductOfferActionRaw registers a raw Gin route for the GetAllDataOfTheParticularProductOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetAllDataOfTheParticularProductOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// GetAllDataOfTheParticularProductOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetAllDataOfTheParticularProductOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetAllDataOfTheParticularProductOfferActionHandler(
	handler func(c GetAllDataOfTheParticularProductOfferActionRequest, gin *gin.Context) (*GetAllDataOfTheParticularProductOfferActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetAllDataOfTheParticularProductOfferActionRequest{
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
// GetAllDataOfTheParticularProductOfferAction is a high-level convenience wrapper around GetAllDataOfTheParticularProductOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetAllDataOfTheParticularProductOfferAction(r gin.IRoutes, handler func(c GetAllDataOfTheParticularProductOfferActionRequest, gin *gin.Context) (*GetAllDataOfTheParticularProductOfferActionResponse, error),) {
	method, url, h := GetAllDataOfTheParticularProductOfferActionHandler(handler)
	r.Handle(method, url, h)
}
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
func GetAllDataOfTheParticularProductOfferActionQueryFromGin(c *gin.Context) GetAllDataOfTheParticularProductOfferActionQuery {
	return GetAllDataOfTheParticularProductOfferActionQueryFromString(c.Request.URL.RawQuery)
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
	QueryParams url.Values
	Headers http.Header
}
type GetAllDataOfTheParticularProductOfferActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func GetAllDataOfTheParticularProductOfferActionCall(
	req GetAllDataOfTheParticularProductOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetAllDataOfTheParticularProductOfferActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetAllDataOfTheParticularProductOfferActionMeta()
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
			req0, err := http.NewRequest(meta.Method, u.String(), nil)
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
	var result GetAllDataOfTheParticularProductOfferActionResult
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