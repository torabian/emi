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
* Action to communicate with the action GetPromoOptionsForSellerSOffersAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetPromoOptionsForSellerSOffersAction
func GetPromoOptionsForSellerSOffersAction(c GetPromoOptionsForSellerSOffersActionRequest) (*GetPromoOptionsForSellerSOffersActionResponse, error) {
	return &GetPromoOptionsForSellerSOffersActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetPromoOptionsForSellerSOffersActionMeta() struct {
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
		Name:        "GetPromoOptionsForSellerSOffersAction",
		CliName:     "get promo options for seller's offers-action",
		URL:         "https://api.{environment}/sale/offers/promo-options",
		Method:      "GET",
		Description: `Use this resource to retrieve promo options for seller offers. Read more: PL / EN.`,
	}
}
func GetGetPromoOptionsForSellerSOffersActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "promo-options",
			Type: "array",
		},
		{
			Name: prefix + "count",
			Type: "int",
		},
		{
			Name: prefix + "total-count",
			Type: "int",
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionRes {
	data := GetPromoOptionsForSellerSOffersActionRes{}
	if c.IsSet("promo-options") {
		data.PromoOptions = emigo.CapturePossibleArray(CastGetPromoOptionsForSellerSOffersActionResPromoOptionsFromCli, "promo-options", c)
	}
	if c.IsSet("count") {
		data.Count = int(c.Int64("count"))
	}
	if c.IsSet("total-count") {
		data.TotalCount = int(c.Int64("total-count"))
	}
	return data
}

// The base class definition for getPromoOptionsForSellerSOffersActionRes
type GetPromoOptionsForSellerSOffersActionRes struct {
	PromoOptions []GetPromoOptionsForSellerSOffersActionResPromoOptions `json:"promoOptions" yaml:"promoOptions"`
	Count        int                                                    `json:"count" yaml:"count"`
	TotalCount   int                                                    `json:"totalCount" yaml:"totalCount"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "offer-id",
			Type: "string",
		},
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetGetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name:     prefix + "pending-changes",
			Type:     "object",
			Children: GetGetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesCliFlags("pending-changes-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "array",
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptions {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptions{}
	if c.IsSet("offer-id") {
		data.OfferId = c.String("offer-id")
	}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("base-package") {
		data.BasePackage = CastGetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastGetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("pending-changes") {
		data.PendingChanges = CastGetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = emigo.CapturePossibleArray(CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesFromCli, "additional-marketplaces", c)
	}
	return data
}

// The base class definition for promoOptions
type GetPromoOptionsForSellerSOffersActionResPromoOptions struct {
	OfferId                string                                                                       `json:"offerId" yaml:"offerId"`
	MarketplaceId          string                                                                       `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage            GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage              `json:"basePackage" yaml:"basePackage"`
	ExtraPackages          []GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages          `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges         GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges           `json:"pendingChanges" yaml:"pendingChanges"`
	AdditionalMarketplaces []GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackageFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackagesFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for extraPackages
type GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetGetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackageCliFlags("base-package-"),
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges{}
	if c.IsSet("base-package") {
		data.BasePackage = CastGetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackageFromCli(c)
	}
	return data
}

// The base class definition for pendingChanges
type GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges struct {
	BasePackage GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackageFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name:     prefix + "pending-changes",
			Type:     "object",
			Children: GetGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesCliFlags("pending-changes-"),
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("base-package") {
		data.BasePackage = CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("pending-changes") {
		data.PendingChanges = CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesFromCli(c)
	}
	return data
}

// The base class definition for additionalMarketplaces
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces struct {
	MarketplaceId  string                                                                                    `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage    GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages  []GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges  `json:"pendingChanges" yaml:"pendingChanges"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackageFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackagesFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for extraPackages
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackageCliFlags("base-package-"),
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges{}
	if c.IsSet("base-package") {
		data.BasePackage = CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackageFromCli(c)
	}
	return data
}

// The base class definition for pendingChanges
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges struct {
	BasePackage GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

func GetGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "valid-from",
			Type: "string",
		},
		{
			Name: prefix + "valid-to",
			Type: "string",
		},
		{
			Name: prefix + "next-cycle-date",
			Type: "string",
		},
	}
}
func CastGetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackageFromCli(c emigo.CliCastable) GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage {
	data := GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("valid-from") {
		data.ValidFrom = c.String("valid-from")
	}
	if c.IsSet("valid-to") {
		data.ValidTo = c.String("valid-to")
	}
	if c.IsSet("next-cycle-date") {
		data.NextCycleDate = c.String("next-cycle-date")
	}
	return data
}

// The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func (x *GetPromoOptionsForSellerSOffersActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetPromoOptionsForSellerSOffersActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetPromoOptionsForSellerSOffersActionResponse) SetContentType(contentType string) *GetPromoOptionsForSellerSOffersActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetPromoOptionsForSellerSOffersActionResponse) AsStream(r io.Reader, contentType string) *GetPromoOptionsForSellerSOffersActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetPromoOptionsForSellerSOffersActionResponse) AsJSON(payload any) *GetPromoOptionsForSellerSOffersActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetPromoOptionsForSellerSOffersActionResponse) WithIdeal(payload GetPromoOptionsForSellerSOffersActionRes) *GetPromoOptionsForSellerSOffersActionResponse {
	x.Payload = payload
	return x
}
func (x *GetPromoOptionsForSellerSOffersActionResponse) AsHTML(payload string) *GetPromoOptionsForSellerSOffersActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetPromoOptionsForSellerSOffersActionResponse) AsBytes(payload []byte) *GetPromoOptionsForSellerSOffersActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetPromoOptionsForSellerSOffersActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetPromoOptionsForSellerSOffersActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetPromoOptionsForSellerSOffersActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetPromoOptionsForSellerSOffersActionRaw registers a raw Gin route for the GetPromoOptionsForSellerSOffersAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetPromoOptionsForSellerSOffersActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetPromoOptionsForSellerSOffersActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetPromoOptionsForSellerSOffersActionRequestSig = func(c GetPromoOptionsForSellerSOffersActionRequest) (*GetPromoOptionsForSellerSOffersActionResponse, error)

// GetPromoOptionsForSellerSOffersActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetPromoOptionsForSellerSOffersAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetPromoOptionsForSellerSOffersActionHandler(
	handler GetPromoOptionsForSellerSOffersActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetPromoOptionsForSellerSOffersActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetPromoOptionsForSellerSOffersActionRequest{
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

// GetPromoOptionsForSellerSOffersAction is a high-level convenience wrapper around GetPromoOptionsForSellerSOffersActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetPromoOptionsForSellerSOffersActionGin(r gin.IRoutes, handler GetPromoOptionsForSellerSOffersActionRequestSig) {
	method, url, h := GetPromoOptionsForSellerSOffersActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get promo options for seller's offersAction
 */
// Query wrapper with private fields
type GetPromoOptionsForSellerSOffersActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetPromoOptionsForSellerSOffersActionQueryFromString(rawQuery string) GetPromoOptionsForSellerSOffersActionQuery {
	v := GetPromoOptionsForSellerSOffersActionQuery{}
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
func GetPromoOptionsForSellerSOffersActionQueryFromGin(c *gin.Context) GetPromoOptionsForSellerSOffersActionQuery {
	return GetPromoOptionsForSellerSOffersActionQueryFromString(c.Request.URL.RawQuery)
}
func GetPromoOptionsForSellerSOffersActionQueryFromHttp(r *http.Request) GetPromoOptionsForSellerSOffersActionQuery {
	return GetPromoOptionsForSellerSOffersActionQueryFromString(r.URL.RawQuery)
}
func (q GetPromoOptionsForSellerSOffersActionQuery) Values() url.Values {
	return q.values
}
func (q GetPromoOptionsForSellerSOffersActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetPromoOptionsForSellerSOffersActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetPromoOptionsForSellerSOffersActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetPromoOptionsForSellerSOffersActionRequest struct {
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

func (x GetPromoOptionsForSellerSOffersActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GetPromoOptionsForSellerSOffersActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GetPromoOptionsForSellerSOffersActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GetPromoOptionsForSellerSOffersActionClientCreateUrl(
	req GetPromoOptionsForSellerSOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetPromoOptionsForSellerSOffersActionMeta()
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
func GetPromoOptionsForSellerSOffersActionClientExecuteTyped(httpReq *http.Request) (*GetPromoOptionsForSellerSOffersActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetPromoOptionsForSellerSOffersActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetPromoOptionsForSellerSOffersActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetPromoOptionsForSellerSOffersActionResponse{Payload: result}, err
	}
	return &GetPromoOptionsForSellerSOffersActionResponse{Payload: result}, nil
}
func GetPromoOptionsForSellerSOffersActionClientBuildRequest(req GetPromoOptionsForSellerSOffersActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetPromoOptionsForSellerSOffersActionMeta()
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
func GetPromoOptionsForSellerSOffersActionCall(
	req GetPromoOptionsForSellerSOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetPromoOptionsForSellerSOffersActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetPromoOptionsForSellerSOffersActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetPromoOptionsForSellerSOffersActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetPromoOptionsForSellerSOffersActionClientExecuteTyped(r)
}
