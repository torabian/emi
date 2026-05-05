package external

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ModifyOfferPromotionPackagesAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ModifyOfferPromotionPackagesAction
func ModifyOfferPromotionPackagesAction(c ModifyOfferPromotionPackagesActionRequest) (*ModifyOfferPromotionPackagesActionResponse, error) {
	return &ModifyOfferPromotionPackagesActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ModifyOfferPromotionPackagesActionMeta() struct {
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
		Name:        "ModifyOfferPromotionPackagesAction",
		CliName:     "modify offer promotion packages-action",
		URL:         "https://api.{environment}/sale/offers/{offerId}/promo-options-modification",
		Method:      "POST",
		Description: `Use this resource to modify offer promotion packages. Read more: PL / EN.`,
	}
}
func GetModifyOfferPromotionPackagesActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "modifications",
			Type: "array",
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "array",
		},
	}
}
func CastModifyOfferPromotionPackagesActionReqFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionReq {
	data := ModifyOfferPromotionPackagesActionReq{}
	if c.IsSet("modifications") {
		data.Modifications = emigo.CapturePossibleArray(CastModifyOfferPromotionPackagesActionReqModificationsFromCli, "modifications", c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = emigo.CapturePossibleArray(CastModifyOfferPromotionPackagesActionReqAdditionalMarketplacesFromCli, "additional-marketplaces", c)
	}
	return data
}

// The base class definition for modifyOfferPromotionPackagesActionReq
type ModifyOfferPromotionPackagesActionReq struct {
	Modifications          []ModifyOfferPromotionPackagesActionReqModifications          `json:"modifications" yaml:"modifications"`
	AdditionalMarketplaces []ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

func GetModifyOfferPromotionPackagesActionReqModificationsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "modification-type",
			Type: "string",
		},
		{
			Name: prefix + "package-type",
			Type: "string",
		},
		{
			Name: prefix + "package-id",
			Type: "string",
		},
	}
}
func CastModifyOfferPromotionPackagesActionReqModificationsFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionReqModifications {
	data := ModifyOfferPromotionPackagesActionReqModifications{}
	if c.IsSet("modification-type") {
		data.ModificationType = c.String("modification-type")
	}
	if c.IsSet("package-type") {
		data.PackageType = c.String("package-type")
	}
	if c.IsSet("package-id") {
		data.PackageId = c.String("package-id")
	}
	return data
}

// The base class definition for modifications
type ModifyOfferPromotionPackagesActionReqModifications struct {
	ModificationType string `json:"modificationType" yaml:"modificationType"`
	PackageType      string `json:"packageType" yaml:"packageType"`
	PackageId        string `json:"packageId" yaml:"packageId"`
}

func GetModifyOfferPromotionPackagesActionReqAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name: prefix + "modifications",
			Type: "array",
		},
	}
}
func CastModifyOfferPromotionPackagesActionReqAdditionalMarketplacesFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces {
	data := ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("modifications") {
		data.Modifications = emigo.CapturePossibleArray(CastModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModificationsFromCli, "modifications", c)
	}
	return data
}

// The base class definition for additionalMarketplaces
type ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces struct {
	MarketplaceId string                                                                     `json:"marketplaceId" yaml:"marketplaceId"`
	Modifications []ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications `json:"modifications" yaml:"modifications"`
}

func GetModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModificationsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "modification-type",
			Type: "string",
		},
		{
			Name: prefix + "package-type",
			Type: "string",
		},
		{
			Name: prefix + "package-id",
			Type: "string",
		},
	}
}
func CastModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModificationsFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications {
	data := ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications{}
	if c.IsSet("modification-type") {
		data.ModificationType = c.String("modification-type")
	}
	if c.IsSet("package-type") {
		data.PackageType = c.String("package-type")
	}
	if c.IsSet("package-id") {
		data.PackageId = c.String("package-id")
	}
	return data
}

// The base class definition for modifications
type ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications struct {
	ModificationType string `json:"modificationType" yaml:"modificationType"`
	PackageType      string `json:"packageType" yaml:"packageType"`
	PackageId        string `json:"packageId" yaml:"packageId"`
}

func (x *ModifyOfferPromotionPackagesActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetModifyOfferPromotionPackagesActionResCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetModifyOfferPromotionPackagesActionResBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name:     prefix + "pending-changes",
			Type:     "object",
			Children: GetModifyOfferPromotionPackagesActionResPendingChangesCliFlags("pending-changes-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "array",
		},
	}
}
func CastModifyOfferPromotionPackagesActionResFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionRes {
	data := ModifyOfferPromotionPackagesActionRes{}
	if c.IsSet("offer-id") {
		data.OfferId = c.String("offer-id")
	}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("base-package") {
		data.BasePackage = CastModifyOfferPromotionPackagesActionResBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastModifyOfferPromotionPackagesActionResExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("pending-changes") {
		data.PendingChanges = CastModifyOfferPromotionPackagesActionResPendingChangesFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = emigo.CapturePossibleArray(CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesFromCli, "additional-marketplaces", c)
	}
	return data
}

// The base class definition for modifyOfferPromotionPackagesActionRes
type ModifyOfferPromotionPackagesActionRes struct {
	OfferId                string                                                        `json:"offerId" yaml:"offerId"`
	MarketplaceId          string                                                        `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage            ModifyOfferPromotionPackagesActionResBasePackage              `json:"basePackage" yaml:"basePackage"`
	ExtraPackages          []ModifyOfferPromotionPackagesActionResExtraPackages          `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges         ModifyOfferPromotionPackagesActionResPendingChanges           `json:"pendingChanges" yaml:"pendingChanges"`
	AdditionalMarketplaces []ModifyOfferPromotionPackagesActionResAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

func GetModifyOfferPromotionPackagesActionResBasePackageCliFlags(prefix string) []emigo.CliFlag {
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
func CastModifyOfferPromotionPackagesActionResBasePackageFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionResBasePackage {
	data := ModifyOfferPromotionPackagesActionResBasePackage{}
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
type ModifyOfferPromotionPackagesActionResBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetModifyOfferPromotionPackagesActionResExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
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
func CastModifyOfferPromotionPackagesActionResExtraPackagesFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionResExtraPackages {
	data := ModifyOfferPromotionPackagesActionResExtraPackages{}
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
type ModifyOfferPromotionPackagesActionResExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetModifyOfferPromotionPackagesActionResPendingChangesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetModifyOfferPromotionPackagesActionResPendingChangesBasePackageCliFlags("base-package-"),
		},
	}
}
func CastModifyOfferPromotionPackagesActionResPendingChangesFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionResPendingChanges {
	data := ModifyOfferPromotionPackagesActionResPendingChanges{}
	if c.IsSet("base-package") {
		data.BasePackage = CastModifyOfferPromotionPackagesActionResPendingChangesBasePackageFromCli(c)
	}
	return data
}

// The base class definition for pendingChanges
type ModifyOfferPromotionPackagesActionResPendingChanges struct {
	BasePackage ModifyOfferPromotionPackagesActionResPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

func GetModifyOfferPromotionPackagesActionResPendingChangesBasePackageCliFlags(prefix string) []emigo.CliFlag {
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
func CastModifyOfferPromotionPackagesActionResPendingChangesBasePackageFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionResPendingChangesBasePackage {
	data := ModifyOfferPromotionPackagesActionResPendingChangesBasePackage{}
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
type ModifyOfferPromotionPackagesActionResPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetModifyOfferPromotionPackagesActionResAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name:     prefix + "pending-changes",
			Type:     "object",
			Children: GetModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesCliFlags("pending-changes-"),
		},
	}
}
func CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionResAdditionalMarketplaces {
	data := ModifyOfferPromotionPackagesActionResAdditionalMarketplaces{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("base-package") {
		data.BasePackage = CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("pending-changes") {
		data.PendingChanges = CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesFromCli(c)
	}
	return data
}

// The base class definition for additionalMarketplaces
type ModifyOfferPromotionPackagesActionResAdditionalMarketplaces struct {
	MarketplaceId  string                                                                     `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage    ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages  []ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges  `json:"pendingChanges" yaml:"pendingChanges"`
}

func GetModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackageCliFlags(prefix string) []emigo.CliFlag {
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
func CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackageFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage {
	data := ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage{}
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
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
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
func CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackagesFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages {
	data := ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages{}
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
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func GetModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackageCliFlags("base-package-"),
		},
	}
}
func CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges {
	data := ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges{}
	if c.IsSet("base-package") {
		data.BasePackage = CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackageFromCli(c)
	}
	return data
}

// The base class definition for pendingChanges
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges struct {
	BasePackage ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

func GetModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackageCliFlags(prefix string) []emigo.CliFlag {
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
func CastModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackageFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage {
	data := ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage{}
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
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func (x *ModifyOfferPromotionPackagesActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ModifyOfferPromotionPackagesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ModifyOfferPromotionPackagesActionResponse) SetContentType(contentType string) *ModifyOfferPromotionPackagesActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ModifyOfferPromotionPackagesActionResponse) AsStream(r io.Reader, contentType string) *ModifyOfferPromotionPackagesActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ModifyOfferPromotionPackagesActionResponse) AsJSON(payload any) *ModifyOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ModifyOfferPromotionPackagesActionResponse) WithIdeal(payload ModifyOfferPromotionPackagesActionRes) *ModifyOfferPromotionPackagesActionResponse {
	x.Payload = payload
	return x
}
func (x *ModifyOfferPromotionPackagesActionResponse) AsHTML(payload string) *ModifyOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ModifyOfferPromotionPackagesActionResponse) AsBytes(payload []byte) *ModifyOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ModifyOfferPromotionPackagesActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ModifyOfferPromotionPackagesActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ModifyOfferPromotionPackagesActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ModifyOfferPromotionPackagesActionRaw registers a raw Gin route for the ModifyOfferPromotionPackagesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModifyOfferPromotionPackagesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModifyOfferPromotionPackagesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ModifyOfferPromotionPackagesActionRequestSig = func(c ModifyOfferPromotionPackagesActionRequest) (*ModifyOfferPromotionPackagesActionResponse, error)

// ModifyOfferPromotionPackagesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModifyOfferPromotionPackagesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModifyOfferPromotionPackagesActionHandler(
	handler ModifyOfferPromotionPackagesActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ModifyOfferPromotionPackagesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ModifyOfferPromotionPackagesActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ModifyOfferPromotionPackagesActionRequest{
			Body:        body,
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

// ModifyOfferPromotionPackagesAction is a high-level convenience wrapper around ModifyOfferPromotionPackagesActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModifyOfferPromotionPackagesActionGin(r gin.IRoutes, handler ModifyOfferPromotionPackagesActionRequestSig) {
	method, url, h := ModifyOfferPromotionPackagesActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Modify offer promotion packagesAction
 */
// Query wrapper with private fields
type ModifyOfferPromotionPackagesActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModifyOfferPromotionPackagesActionQueryFromString(rawQuery string) ModifyOfferPromotionPackagesActionQuery {
	v := ModifyOfferPromotionPackagesActionQuery{}
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
func ModifyOfferPromotionPackagesActionQueryFromGin(c *gin.Context) ModifyOfferPromotionPackagesActionQuery {
	return ModifyOfferPromotionPackagesActionQueryFromString(c.Request.URL.RawQuery)
}
func ModifyOfferPromotionPackagesActionQueryFromHttp(r *http.Request) ModifyOfferPromotionPackagesActionQuery {
	return ModifyOfferPromotionPackagesActionQueryFromString(r.URL.RawQuery)
}
func (q ModifyOfferPromotionPackagesActionQuery) Values() url.Values {
	return q.values
}
func (q ModifyOfferPromotionPackagesActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModifyOfferPromotionPackagesActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModifyOfferPromotionPackagesActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModifyOfferPromotionPackagesActionRequest struct {
	Body        ModifyOfferPromotionPackagesActionReq
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

func (x ModifyOfferPromotionPackagesActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x ModifyOfferPromotionPackagesActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type ModifyOfferPromotionPackagesActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func ModifyOfferPromotionPackagesActionClientCreateUrl(
	req ModifyOfferPromotionPackagesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ModifyOfferPromotionPackagesActionMeta()
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
func ModifyOfferPromotionPackagesActionClientExecuteTyped(httpReq *http.Request) (*ModifyOfferPromotionPackagesActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ModifyOfferPromotionPackagesActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ModifyOfferPromotionPackagesActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ModifyOfferPromotionPackagesActionResponse{Payload: result}, err
	}
	return &ModifyOfferPromotionPackagesActionResponse{Payload: result}, nil
}
func ModifyOfferPromotionPackagesActionClientBuildRequest(req ModifyOfferPromotionPackagesActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ModifyOfferPromotionPackagesActionMeta()
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
func ModifyOfferPromotionPackagesActionCall(
	req ModifyOfferPromotionPackagesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModifyOfferPromotionPackagesActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ModifyOfferPromotionPackagesActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ModifyOfferPromotionPackagesActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ModifyOfferPromotionPackagesActionClientExecuteTyped(r)
}
