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
* Action to communicate with the action DeleteADraftOfferAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of DeleteADraftOfferAction
func DeleteADraftOfferAction(c DeleteADraftOfferActionRequest) (*DeleteADraftOfferActionResponse, error) {
	return &DeleteADraftOfferActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func DeleteADraftOfferActionMeta() struct {
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
		Name:        "DeleteADraftOfferAction",
		CliName:     "delete a draft offer-action",
		URL:         "https://api.{environment}/sale/offers/{offerId}",
		Method:      "DELETE",
		Description: `Use this resource to delete a draft offer. Read more: PL / EN.`,
	}
}

type DeleteADraftOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *DeleteADraftOfferActionResponse) SetContentType(contentType string) *DeleteADraftOfferActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *DeleteADraftOfferActionResponse) AsStream(r io.Reader, contentType string) *DeleteADraftOfferActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *DeleteADraftOfferActionResponse) AsJSON(payload any) *DeleteADraftOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *DeleteADraftOfferActionResponse) AsHTML(payload string) *DeleteADraftOfferActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *DeleteADraftOfferActionResponse) AsBytes(payload []byte) *DeleteADraftOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x DeleteADraftOfferActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x DeleteADraftOfferActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x DeleteADraftOfferActionResponse) GetPayload() interface{} {
	return x.Payload
}

// DeleteADraftOfferActionRaw registers a raw Gin route for the DeleteADraftOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func DeleteADraftOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := DeleteADraftOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type DeleteADraftOfferActionRequestSig = func(c DeleteADraftOfferActionRequest) (*DeleteADraftOfferActionResponse, error)

// DeleteADraftOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the DeleteADraftOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func DeleteADraftOfferActionHandler(
	handler DeleteADraftOfferActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := DeleteADraftOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := DeleteADraftOfferActionRequest{
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

// DeleteADraftOfferAction is a high-level convenience wrapper around DeleteADraftOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func DeleteADraftOfferActionGin(r gin.IRoutes, handler DeleteADraftOfferActionRequestSig) {
	method, url, h := DeleteADraftOfferActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Delete a draft offerAction
 */
// Query wrapper with private fields
type DeleteADraftOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func DeleteADraftOfferActionQueryFromString(rawQuery string) DeleteADraftOfferActionQuery {
	v := DeleteADraftOfferActionQuery{}
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
func DeleteADraftOfferActionQueryFromGin(c *gin.Context) DeleteADraftOfferActionQuery {
	return DeleteADraftOfferActionQueryFromString(c.Request.URL.RawQuery)
}
func DeleteADraftOfferActionQueryFromHttp(r *http.Request) DeleteADraftOfferActionQuery {
	return DeleteADraftOfferActionQueryFromString(r.URL.RawQuery)
}
func (q DeleteADraftOfferActionQuery) Values() url.Values {
	return q.values
}
func (q DeleteADraftOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *DeleteADraftOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *DeleteADraftOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type DeleteADraftOfferActionRequest struct {
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

func (x DeleteADraftOfferActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x DeleteADraftOfferActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type DeleteADraftOfferActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func DeleteADraftOfferActionClientCreateUrl(
	req DeleteADraftOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := DeleteADraftOfferActionMeta()
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
func DeleteADraftOfferActionClientExecuteTyped(httpReq *http.Request) (*DeleteADraftOfferActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result DeleteADraftOfferActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &DeleteADraftOfferActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &DeleteADraftOfferActionResponse{Payload: result}, err
	}
	return &DeleteADraftOfferActionResponse{Payload: result}, nil
}
func DeleteADraftOfferActionClientBuildRequest(req DeleteADraftOfferActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := DeleteADraftOfferActionMeta()
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
func DeleteADraftOfferActionCall(
	req DeleteADraftOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*DeleteADraftOfferActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := DeleteADraftOfferActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := DeleteADraftOfferActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return DeleteADraftOfferActionClientExecuteTyped(r)
}
