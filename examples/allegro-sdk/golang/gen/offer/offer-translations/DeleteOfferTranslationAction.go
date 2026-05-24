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
* Action to communicate with the action DeleteOfferTranslationAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of DeleteOfferTranslationAction
func DeleteOfferTranslationAction(c DeleteOfferTranslationActionRequest) (*DeleteOfferTranslationActionResponse, error) {
	return &DeleteOfferTranslationActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func DeleteOfferTranslationActionMeta() struct {
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
		Name:        "DeleteOfferTranslationAction",
		CliName:     "delete offer translation-action",
		URL:         "https://api.{environment}/sale/offers/{offerId}/translations/{language}",
		Method:      "DELETE",
		Description: `Delete single element or entire manual translation. Read more: PL / EN.`,
	}
}

type DeleteOfferTranslationActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *DeleteOfferTranslationActionResponse) SetContentType(contentType string) *DeleteOfferTranslationActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *DeleteOfferTranslationActionResponse) AsStream(r io.Reader, contentType string) *DeleteOfferTranslationActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *DeleteOfferTranslationActionResponse) AsJSON(payload any) *DeleteOfferTranslationActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *DeleteOfferTranslationActionResponse) AsHTML(payload string) *DeleteOfferTranslationActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *DeleteOfferTranslationActionResponse) AsBytes(payload []byte) *DeleteOfferTranslationActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x DeleteOfferTranslationActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x DeleteOfferTranslationActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x DeleteOfferTranslationActionResponse) GetPayload() interface{} {
	return x.Payload
}

// DeleteOfferTranslationActionRaw registers a raw Gin route for the DeleteOfferTranslationAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func DeleteOfferTranslationActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := DeleteOfferTranslationActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type DeleteOfferTranslationActionRequestSig = func(c DeleteOfferTranslationActionRequest) (*DeleteOfferTranslationActionResponse, error)

// DeleteOfferTranslationActionHandler returns the HTTP method, route URL, and a typed Gin handler for the DeleteOfferTranslationAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func DeleteOfferTranslationActionHandler(
	handler DeleteOfferTranslationActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := DeleteOfferTranslationActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := DeleteOfferTranslationActionRequest{
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

// DeleteOfferTranslationAction is a high-level convenience wrapper around DeleteOfferTranslationActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func DeleteOfferTranslationActionGin(r gin.IRoutes, handler DeleteOfferTranslationActionRequestSig) {
	method, url, h := DeleteOfferTranslationActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Delete offer translationAction
 */
// Query wrapper with private fields
type DeleteOfferTranslationActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func DeleteOfferTranslationActionQueryFromString(rawQuery string) DeleteOfferTranslationActionQuery {
	v := DeleteOfferTranslationActionQuery{}
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
func DeleteOfferTranslationActionQueryFromGin(c *gin.Context) DeleteOfferTranslationActionQuery {
	return DeleteOfferTranslationActionQueryFromString(c.Request.URL.RawQuery)
}
func DeleteOfferTranslationActionQueryFromHttp(r *http.Request) DeleteOfferTranslationActionQuery {
	return DeleteOfferTranslationActionQueryFromString(r.URL.RawQuery)
}
func (q DeleteOfferTranslationActionQuery) Values() url.Values {
	return q.values
}
func (q DeleteOfferTranslationActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *DeleteOfferTranslationActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *DeleteOfferTranslationActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type DeleteOfferTranslationActionRequest struct {
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

func (x DeleteOfferTranslationActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x DeleteOfferTranslationActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type DeleteOfferTranslationActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func DeleteOfferTranslationActionClientCreateUrl(
	req DeleteOfferTranslationActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := DeleteOfferTranslationActionMeta()
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
func DeleteOfferTranslationActionClientExecuteTyped(httpReq *http.Request) (*DeleteOfferTranslationActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result DeleteOfferTranslationActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &DeleteOfferTranslationActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &DeleteOfferTranslationActionResponse{Payload: result}, err
	}
	return &DeleteOfferTranslationActionResponse{Payload: result}, nil
}
func DeleteOfferTranslationActionClientBuildRequest(req DeleteOfferTranslationActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := DeleteOfferTranslationActionMeta()
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
func DeleteOfferTranslationActionCall(
	req DeleteOfferTranslationActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*DeleteOfferTranslationActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := DeleteOfferTranslationActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := DeleteOfferTranslationActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return DeleteOfferTranslationActionClientExecuteTyped(r)
}
