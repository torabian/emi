package external

import (
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action GetOfferPromotionPackagesAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetOfferPromotionPackagesAction
func GetOfferPromotionPackagesAction(c GetOfferPromotionPackagesActionRequest) (*GetOfferPromotionPackagesActionResponse, error) {
	return &GetOfferPromotionPackagesActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetOfferPromotionPackagesActionMeta() struct {
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
		Name:        "GetOfferPromotionPackagesAction",
		CliName:     "get offer promotion packages-action",
		URL:         "https://api.{environment}/sale/offers/{offerId}/promo-options",
		Method:      "GET",
		Description: ``,
	}
}

// The base class definition for getOfferPromotionPackagesActionRes
type GetOfferPromotionPackagesActionRes struct {
	OfferId                string                                                                `json:"offerId" yaml:"offerId"`
	MarketplaceId          string                                                                `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage            GetOfferPromotionPackagesActionResBasePackage                         `json:"basePackage" yaml:"basePackage"`
	ExtraPackages          emigo.Array[GetOfferPromotionPackagesActionResExtraPackages]          `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges         GetOfferPromotionPackagesActionResPendingChanges                      `json:"pendingChanges" yaml:"pendingChanges"`
	AdditionalMarketplaces emigo.Array[GetOfferPromotionPackagesActionResAdditionalMarketplaces] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for basePackage
type GetOfferPromotionPackagesActionResBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for extraPackages
type GetOfferPromotionPackagesActionResExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for pendingChanges
type GetOfferPromotionPackagesActionResPendingChanges struct {
	BasePackage GetOfferPromotionPackagesActionResPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

// The base class definition for basePackage
type GetOfferPromotionPackagesActionResPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for additionalMarketplaces
type GetOfferPromotionPackagesActionResAdditionalMarketplaces struct {
	MarketplaceId  string                                                                             `json:"marketplaceId" yaml:"marketplaceId"`
	BasePackage    GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage                `json:"basePackage" yaml:"basePackage"`
	ExtraPackages  emigo.Array[GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages] `json:"extraPackages" yaml:"extraPackages"`
	PendingChanges GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges             `json:"pendingChanges" yaml:"pendingChanges"`
}

// The base class definition for basePackage
type GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for extraPackages
type GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

// The base class definition for pendingChanges
type GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges struct {
	BasePackage GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage `json:"basePackage" yaml:"basePackage"`
}

// The base class definition for basePackage
type GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage struct {
	Id            string `json:"id" yaml:"id"`
	ValidFrom     string `json:"validFrom" yaml:"validFrom"`
	ValidTo       string `json:"validTo" yaml:"validTo"`
	NextCycleDate string `json:"nextCycleDate" yaml:"nextCycleDate"`
}

func (x *GetOfferPromotionPackagesActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetOfferPromotionPackagesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetOfferPromotionPackagesActionResponse) SetContentType(contentType string) *GetOfferPromotionPackagesActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetOfferPromotionPackagesActionResponse) AsStream(r io.Reader, contentType string) *GetOfferPromotionPackagesActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetOfferPromotionPackagesActionResponse) AsJSON(payload any) *GetOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetOfferPromotionPackagesActionResponse) WithIdeal(payload GetOfferPromotionPackagesActionRes) *GetOfferPromotionPackagesActionResponse {
	x.Payload = payload
	return x
}
func (x *GetOfferPromotionPackagesActionResponse) AsHTML(payload string) *GetOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetOfferPromotionPackagesActionResponse) AsBytes(payload []byte) *GetOfferPromotionPackagesActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetOfferPromotionPackagesActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetOfferPromotionPackagesActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetOfferPromotionPackagesActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type GetOfferPromotionPackagesActionRequestSig = func(c GetOfferPromotionPackagesActionRequest) (*GetOfferPromotionPackagesActionResponse, error)

/**
 * Query parameters for Get offer promotion packagesAction
 */
// Query wrapper with private fields
type GetOfferPromotionPackagesActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetOfferPromotionPackagesActionQueryFromString(rawQuery string) GetOfferPromotionPackagesActionQuery {
	v := GetOfferPromotionPackagesActionQuery{}
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
func GetOfferPromotionPackagesActionQueryFromHttp(r *http.Request) GetOfferPromotionPackagesActionQuery {
	return GetOfferPromotionPackagesActionQueryFromString(r.URL.RawQuery)
}
func (q GetOfferPromotionPackagesActionQuery) Values() url.Values {
	return q.values
}
func (q GetOfferPromotionPackagesActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetOfferPromotionPackagesActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetOfferPromotionPackagesActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetOfferPromotionPackagesActionRequest struct {
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

func GetOfferPromotionPackagesActionClientCreateUrl(
	req GetOfferPromotionPackagesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetOfferPromotionPackagesActionMeta()
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
func GetOfferPromotionPackagesActionClientExecuteTyped(httpReq *http.Request) (*GetOfferPromotionPackagesActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetOfferPromotionPackagesActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetOfferPromotionPackagesActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetOfferPromotionPackagesActionResponse{Payload: result}, err
	}
	return &GetOfferPromotionPackagesActionResponse{Payload: result}, nil
}
func GetOfferPromotionPackagesActionClientBuildRequest(req GetOfferPromotionPackagesActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetOfferPromotionPackagesActionMeta()
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
func GetOfferPromotionPackagesActionCall(
	req GetOfferPromotionPackagesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetOfferPromotionPackagesActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetOfferPromotionPackagesActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetOfferPromotionPackagesActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetOfferPromotionPackagesActionClientExecuteTyped(r)
}

// GetOfferPromotionPackagesActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the GetOfferPromotionPackagesAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *GetOfferPromotionPackagesActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func GetOfferPromotionPackagesActionHttpHandler(
	handler func(c GetOfferPromotionPackagesActionRequest) (*GetOfferPromotionPackagesActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := GetOfferPromotionPackagesActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := GetOfferPromotionPackagesActionRequest{
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

// GetOfferPromotionPackagesActionHttp is a high-level convenience wrapper around
// GetOfferPromotionPackagesActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func GetOfferPromotionPackagesActionHttp(
	mux *http.ServeMux,
	handler func(c GetOfferPromotionPackagesActionRequest) (*GetOfferPromotionPackagesActionResponse, error),
) {
	method, pattern, h := GetOfferPromotionPackagesActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
