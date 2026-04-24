package external

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction
 */
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta() struct {
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
		Name:        "CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction",
		CliName:     "check the processing status of a post or patch request-action",
		URL:         "https://api.{environment}/sale/product-offers/{offerId}/operations/{operationId}",
		Method:      "GET",
		Description: `The URI for the resource given by Location header of POST /sale/product-offers and PATCH /sale/product-offers/{offerId}. Use this resource to check processing status of a POST or PATCH request. Read more: PL / EN.`,
	}
}
func GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "offer",
			Type:     "object",
			Children: GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOfferCliFlags("offer-"),
		},
		{
			Name:     prefix + "operation",
			Type:     "object",
			Children: GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperationCliFlags("operation-"),
		},
	}
}
func CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResFromCli(c emigo.CliCastable) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes {
	data := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes{}
	if c.IsSet("offer") {
		data.Offer = CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOfferFromCli(c)
	}
	if c.IsSet("operation") {
		data.Operation = CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperationFromCli(c)
	}
	return data
}

// The base class definition for checkTheProcessingStatusOfAPOSTOrPATCHRequestActionRes
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes struct {
	Offer     CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer     `json:"offer" yaml:"offer"`
	Operation CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation `json:"operation" yaml:"operation"`
}

func GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOfferCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOfferFromCli(c emigo.CliCastable) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer {
	data := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for offer
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer struct {
	Id string `json:"id" yaml:"id"`
}

func GetCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "status",
			Type: "string",
		},
		{
			Name: prefix + "started-at",
			Type: "string",
		},
	}
}
func CastCheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperationFromCli(c emigo.CliCastable) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation {
	data := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("started-at") {
		data.StartedAt = c.String("started-at")
	}
	return data
}

// The base class definition for operation
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation struct {
	Id        string `json:"id" yaml:"id"`
	Status    string `json:"status" yaml:"status"`
	StartedAt string `json:"startedAt" yaml:"startedAt"`
}

func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) SetContentType(contentType string) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) AsStream(r io.Reader, contentType string) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) AsJSON(payload any) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) AsHTML(payload string) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) AsBytes(payload []byte) *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse) GetPayload() interface{} {
	return x.Payload
}

// CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRaw registers a raw Gin route for the CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequestSig = func(c CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse, error)

// CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHandler(
	handler CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest{
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

// CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction is a high-level convenience wrapper around CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionGin(r gin.IRoutes, handler CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequestSig) {
	method, url, h := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Check the processing status of a POST or PATCH requestAction
 */
// Query wrapper with private fields
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromString(rawQuery string) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery {
	v := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery{}
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
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromGin(c *gin.Context) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery {
	return CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromString(c.Request.URL.RawQuery)
}
func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromHttp(r *http.Request) CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery {
	return CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQueryFromString(r.URL.RawQuery)
}
func (q CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) Values() url.Values {
	return q.values
}
func (q CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionCall(
	req CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
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
	var result CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResult
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
