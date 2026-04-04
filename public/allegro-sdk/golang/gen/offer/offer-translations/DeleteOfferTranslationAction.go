package external

import (
	"bytes"
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
* Action to communicate with the action DeleteOfferTranslationAction
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
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type DeleteOfferTranslationActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func DeleteOfferTranslationActionCall(
	req DeleteOfferTranslationActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*DeleteOfferTranslationActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := DeleteOfferTranslationActionMeta()
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
	var result DeleteOfferTranslationActionResult
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
