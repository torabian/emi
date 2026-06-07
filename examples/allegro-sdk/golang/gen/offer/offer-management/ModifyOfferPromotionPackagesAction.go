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

// The base class definition for modifyOfferPromotionPackagesActionReq
type ModifyOfferPromotionPackagesActionReq struct {
	Modifications          emigo.Array[ModifyOfferPromotionPackagesActionReqModifications]          `json:"modifications" yaml:"modifications"`
	AdditionalMarketplaces emigo.Array[ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for modifications
type ModifyOfferPromotionPackagesActionReqModifications struct {
	ModificationType string `json:"modificationType" yaml:"modificationType"`
	PackageType      string `json:"packageType" yaml:"packageType"`
	PackageId        string `json:"packageId" yaml:"packageId"`
}

// The base class definition for additionalMarketplaces
type ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces struct {
	MarketplaceId string                                                                                `json:"marketplaceId" yaml:"marketplaceId"`
	Modifications emigo.Array[ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications] `json:"modifications" yaml:"modifications"`
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

// The base class definition for modifyOfferPromotionPackagesActionRes
type ModifyOfferPromotionPackagesActionRes struct {
	OfferId                string                                                                   `json:"offerId" yaml:"offerId"`
	MarketplaceId          string                                                                   `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage            ModifyOfferPromotionPackagesActionResBasePackage                         `json:"basePackage" yaml:"basePackage"`
	ExtraPackages          emigo.Array[ModifyOfferPromotionPackagesActionResExtraPackages]          `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges         ModifyOfferPromotionPackagesActionResPendingChanges                      `json:"pendingChanges" yaml:"pendingChanges"`
	AdditionalMarketplaces emigo.Array[ModifyOfferPromotionPackagesActionResAdditionalMarketplaces] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for basePackage
type ModifyOfferPromotionPackagesActionResBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for extraPackages
type ModifyOfferPromotionPackagesActionResExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for pendingChanges
type ModifyOfferPromotionPackagesActionResPendingChanges struct {
	BasePackage ModifyOfferPromotionPackagesActionResPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

// The base class definition for basePackage
type ModifyOfferPromotionPackagesActionResPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for additionalMarketplaces
type ModifyOfferPromotionPackagesActionResAdditionalMarketplaces struct {
	MarketplaceId  string                                                                                `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage    ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage                `json:"basePackage" yaml:"basePackage"`
	ExtraPackages  emigo.Array[ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages] `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges             `json:"pendingChanges" yaml:"pendingChanges"`
}

// The base class definition for basePackage
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for extraPackages
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for pendingChanges
type ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges struct {
	BasePackage ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
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

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type ModifyOfferPromotionPackagesActionRequestSig = func(c ModifyOfferPromotionPackagesActionRequest) (*ModifyOfferPromotionPackagesActionResponse, error)

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

// ModifyOfferPromotionPackagesActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the ModifyOfferPromotionPackagesAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *ModifyOfferPromotionPackagesActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func ModifyOfferPromotionPackagesActionHttpHandler(
	handler func(c ModifyOfferPromotionPackagesActionRequest) (*ModifyOfferPromotionPackagesActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := ModifyOfferPromotionPackagesActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body ModifyOfferPromotionPackagesActionReq
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
		req := ModifyOfferPromotionPackagesActionRequest{
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

// ModifyOfferPromotionPackagesActionHttp is a high-level convenience wrapper around
// ModifyOfferPromotionPackagesActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func ModifyOfferPromotionPackagesActionHttp(
	mux *http.ServeMux,
	handler func(c ModifyOfferPromotionPackagesActionRequest) (*ModifyOfferPromotionPackagesActionResponse, error),
) {
	method, pattern, h := ModifyOfferPromotionPackagesActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
