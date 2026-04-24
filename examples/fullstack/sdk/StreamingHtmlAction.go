package external

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action StreamingHtmlAction
 */
func StreamingHtmlActionMeta() struct {
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
		Name:        "StreamingHtmlAction",
		CliName:     "streaming-html-action",
		URL:         "/stream/html",
		Method:      "GET",
		Description: ``,
	}
}

type StreamingHtmlActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *StreamingHtmlActionResponse) SetContentType(contentType string) *StreamingHtmlActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *StreamingHtmlActionResponse) AsStream(r io.Reader, contentType string) *StreamingHtmlActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *StreamingHtmlActionResponse) AsJSON(payload any) *StreamingHtmlActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *StreamingHtmlActionResponse) AsHTML(payload string) *StreamingHtmlActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *StreamingHtmlActionResponse) AsBytes(payload []byte) *StreamingHtmlActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x StreamingHtmlActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x StreamingHtmlActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x StreamingHtmlActionResponse) GetPayload() interface{} {
	return x.Payload
}

// StreamingHtmlActionRaw registers a raw Gin route for the StreamingHtmlAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func StreamingHtmlActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := StreamingHtmlActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type StreamingHtmlActionRequestSig = func(c StreamingHtmlActionRequest) (*StreamingHtmlActionResponse, error)

// StreamingHtmlActionHandler returns the HTTP method, route URL, and a typed Gin handler for the StreamingHtmlAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func StreamingHtmlActionHandler(
	handler StreamingHtmlActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := StreamingHtmlActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := StreamingHtmlActionRequest{
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

// StreamingHtmlAction is a high-level convenience wrapper around StreamingHtmlActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func StreamingHtmlActionGin(r gin.IRoutes, handler StreamingHtmlActionRequestSig) {
	method, url, h := StreamingHtmlActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for StreamingHtmlAction
 */
// Query wrapper with private fields
type StreamingHtmlActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func StreamingHtmlActionQueryFromString(rawQuery string) StreamingHtmlActionQuery {
	v := StreamingHtmlActionQuery{}
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
func StreamingHtmlActionQueryFromGin(c *gin.Context) StreamingHtmlActionQuery {
	return StreamingHtmlActionQueryFromString(c.Request.URL.RawQuery)
}
func StreamingHtmlActionQueryFromHttp(r *http.Request) StreamingHtmlActionQuery {
	return StreamingHtmlActionQueryFromString(r.URL.RawQuery)
}
func (q StreamingHtmlActionQuery) Values() url.Values {
	return q.values
}
func (q StreamingHtmlActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *StreamingHtmlActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *StreamingHtmlActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type StreamingHtmlActionRequest struct {
	Body        interface{}
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type StreamingHtmlActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func StreamingHtmlActionCall(
	req StreamingHtmlActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*StreamingHtmlActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := StreamingHtmlActionMeta()
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
	var result StreamingHtmlActionResult
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
