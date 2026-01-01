package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
)

/**
* Action to communicate with the action ModificationCommandDetailedResultAction
 */
func ModificationCommandDetailedResultActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "ModificationCommandDetailedResultAction",
		URL:    "https://api.{environment}/sale/offers/promo-options-commands/{commandId}/tasks",
		Method: "GET",
	}
}

// The base class definition for modificationCommandDetailedResultActionRes
type ModificationCommandDetailedResultActionRes struct {
	Tasks                  []ModificationCommandDetailedResultActionResTasks                  `json:"tasks" yaml:"tasks"`
	Modification           ModificationCommandDetailedResultActionResModification             `json:"modification" yaml:"modification"`
	AdditionalMarketplaces []ModificationCommandDetailedResultActionResAdditionalMarketplaces `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

// The base class definition for tasks
type ModificationCommandDetailedResultActionResTasks struct {
	Offer         ModificationCommandDetailedResultActionResTasksOffer    `json:"offer" yaml:"offer"`
	MarketplaceId string                                                  `json:"marketplaceId" yaml:"marketplaceId"`
	ScheduledAt   string                                                  `json:"scheduledAt" yaml:"scheduledAt"`
	FinishedAt    string                                                  `json:"finishedAt" yaml:"finishedAt"`
	Status        string                                                  `json:"status" yaml:"status"`
	Errors        []ModificationCommandDetailedResultActionResTasksErrors `json:"errors" yaml:"errors"`
}

// The base class definition for offer
type ModificationCommandDetailedResultActionResTasksOffer struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for errors
type ModificationCommandDetailedResultActionResTasksErrors struct {
	Code        string                                                        `json:"code" yaml:"code"`
	Details     string                                                        `json:"details" yaml:"details"`
	Message     string                                                        `json:"message" yaml:"message"`
	Path        string                                                        `json:"path" yaml:"path"`
	UserMessage string                                                        `json:"userMessage" yaml:"userMessage"`
	Metadata    ModificationCommandDetailedResultActionResTasksErrorsMetadata `json:"metadata" yaml:"metadata"`
}

// The base class definition for metadata
type ModificationCommandDetailedResultActionResTasksErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

// The base class definition for modification
type ModificationCommandDetailedResultActionResModification struct {
	BasePackage      ModificationCommandDetailedResultActionResModificationBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages    []ModificationCommandDetailedResultActionResModificationExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	ModificationTime string                                                                `json:"modificationTime" yaml:"modificationTime"`
}

// The base class definition for basePackage
type ModificationCommandDetailedResultActionResModificationBasePackage struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for extraPackages
type ModificationCommandDetailedResultActionResModificationExtraPackages struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for additionalMarketplaces
type ModificationCommandDetailedResultActionResAdditionalMarketplaces struct {
	MarketplaceId string                                                                       `json:"marketplaceId" yaml:"marketplaceId"`
	Modification  ModificationCommandDetailedResultActionResAdditionalMarketplacesModification `json:"modification" yaml:"modification"`
}

// The base class definition for modification
type ModificationCommandDetailedResultActionResAdditionalMarketplacesModification struct {
	BasePackage      ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage     `json:"basePackage" yaml:"basePackage"`
	ExtraPackages    []ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages `json:"extraPackages" yaml:"extraPackages"`
	ModificationTime string                                                                                      `json:"modificationTime" yaml:"modificationTime"`
}

// The base class definition for basePackage
type ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for extraPackages
type ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages struct {
	Id string `json:"id" yaml:"id"`
}
type ModificationCommandDetailedResultActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// ModificationCommandDetailedResultActionRaw registers a raw Gin route for the ModificationCommandDetailedResultAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ModificationCommandDetailedResultActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ModificationCommandDetailedResultActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // ModificationCommandDetailedResultActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ModificationCommandDetailedResultAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ModificationCommandDetailedResultActionHandler(
	handler func(c ModificationCommandDetailedResultActionRequest, gin *gin.Context) (*ModificationCommandDetailedResultActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ModificationCommandDetailedResultActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := ModificationCommandDetailedResultActionRequest{
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

// ModificationCommandDetailedResultAction is a high-level convenience wrapper around ModificationCommandDetailedResultActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ModificationCommandDetailedResultAction(r gin.IRoutes, handler func(c ModificationCommandDetailedResultActionRequest, gin *gin.Context) (*ModificationCommandDetailedResultActionResponse, error)) {
	method, url, h := ModificationCommandDetailedResultActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Modification command detailed resultAction
 */
// Query wrapper with private fields
type ModificationCommandDetailedResultActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ModificationCommandDetailedResultActionQueryFromString(rawQuery string) ModificationCommandDetailedResultActionQuery {
	v := ModificationCommandDetailedResultActionQuery{}
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
func ModificationCommandDetailedResultActionQueryFromGin(c *gin.Context) ModificationCommandDetailedResultActionQuery {
	return ModificationCommandDetailedResultActionQueryFromString(c.Request.URL.RawQuery)
}
func ModificationCommandDetailedResultActionQueryFromHttp(r *http.Request) ModificationCommandDetailedResultActionQuery {
	return ModificationCommandDetailedResultActionQueryFromString(r.URL.RawQuery)
}
func (q ModificationCommandDetailedResultActionQuery) Values() url.Values {
	return q.values
}
func (q ModificationCommandDetailedResultActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ModificationCommandDetailedResultActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ModificationCommandDetailedResultActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ModificationCommandDetailedResultActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type ModificationCommandDetailedResultActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ModificationCommandDetailedResultActionCall(
	req ModificationCommandDetailedResultActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ModificationCommandDetailedResultActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ModificationCommandDetailedResultActionMeta()
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
	var result ModificationCommandDetailedResultActionResult
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
