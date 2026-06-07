package external

import (
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
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

// The base class definition for getPromoOptionsForSellerSOffersActionRes
type GetPromoOptionsForSellerSOffersActionRes struct {
	PromoOptions emigo.Array[GetPromoOptionsForSellerSOffersActionResPromoOptions] `json:"promoOptions" yaml:"promoOptions"`
	Count        int                                                               `json:"count" yaml:"count"`
	TotalCount   int                                                               `json:"totalCount" yaml:"totalCount"`
}

// The base class definition for promoOptions
type GetPromoOptionsForSellerSOffersActionResPromoOptions struct {
	OfferId                string                                                                                  `json:"offerId" yaml:"offerId"`
	MarketplaceId          string                                                                                  `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage            GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage                         `json:"basePackage" yaml:"basePackage"`
	ExtraPackages          emigo.Array[GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages]          `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges         GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges                      `json:"pendingChanges" yaml:"pendingChanges"`
	AdditionalMarketplaces emigo.Array[GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for extraPackages
type GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for pendingChanges
type GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges struct {
	BasePackage GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

// The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for additionalMarketplaces
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces struct {
	MarketplaceId  string                                                                                               `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage    GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage                `json:"basePackage" yaml:"basePackage"`
	ExtraPackages  emigo.Array[GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages] `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges             `json:"pendingChanges" yaml:"pendingChanges"`
}

// The base class definition for basePackage
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for extraPackages
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for pendingChanges
type GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges struct {
	BasePackage GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
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

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type GetPromoOptionsForSellerSOffersActionRequestSig = func(c GetPromoOptionsForSellerSOffersActionRequest) (*GetPromoOptionsForSellerSOffersActionResponse, error)

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

// GetPromoOptionsForSellerSOffersActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the GetPromoOptionsForSellerSOffersAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *GetPromoOptionsForSellerSOffersActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func GetPromoOptionsForSellerSOffersActionHttpHandler(
	handler func(c GetPromoOptionsForSellerSOffersActionRequest) (*GetPromoOptionsForSellerSOffersActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := GetPromoOptionsForSellerSOffersActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := GetPromoOptionsForSellerSOffersActionRequest{
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

// GetPromoOptionsForSellerSOffersActionHttp is a high-level convenience wrapper around
// GetPromoOptionsForSellerSOffersActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func GetPromoOptionsForSellerSOffersActionHttp(
	mux *http.ServeMux,
	handler func(c GetPromoOptionsForSellerSOffersActionRequest) (*GetPromoOptionsForSellerSOffersActionResponse, error),
) {
	method, pattern, h := GetPromoOptionsForSellerSOffersActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
