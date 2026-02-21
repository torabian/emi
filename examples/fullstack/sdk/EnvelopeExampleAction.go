package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/fullstack/emigo"
)

/**
* Action to communicate with the action EnvelopeExampleAction
 */
func EnvelopeExampleActionMeta() struct {
	Name    string
	CliName string
	URL     string
	Method  string
} {
	return struct {
		Name    string
		CliName string
		URL     string
		Method  string
	}{
		Name:    "EnvelopeExampleAction",
		CliName: "envelope-example-action",
		URL:     "/response/with/envelop",
		Method:  "GET",
	}
}
func GetEnvelopeExampleActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "content",
			Type: "string",
		},
	}
}
func CastEnvelopeExampleActionResFromCli(c emigo.CliCastable) EnvelopeExampleActionRes {
	data := EnvelopeExampleActionRes{}
	if c.IsSet("content") {
		data.Content = c.String("content")
	}
	return data
}

// The base class definition for envelopeExampleActionRes
type EnvelopeExampleActionRes struct {
	Content string `json:"content" yaml:"content"`
}

func (x *EnvelopeExampleActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type EnvelopeExampleActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// EnvelopeExampleActionRaw registers a raw Gin route for the EnvelopeExampleAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func EnvelopeExampleActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := EnvelopeExampleActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // EnvelopeExampleActionHandler returns the HTTP method, route URL, and a typed Gin handler for the EnvelopeExampleAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func EnvelopeExampleActionHandler(
	handler func(c EnvelopeExampleActionRequest, gin *gin.Context) (*EnvelopeExampleActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := EnvelopeExampleActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := EnvelopeExampleActionRequest{
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

// EnvelopeExampleAction is a high-level convenience wrapper around EnvelopeExampleActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func EnvelopeExampleAction(r gin.IRoutes, handler func(c EnvelopeExampleActionRequest, gin *gin.Context) (*EnvelopeExampleActionResponse, error)) {
	method, url, h := EnvelopeExampleActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for EnvelopeExampleAction
 */
// Query wrapper with private fields
type EnvelopeExampleActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func EnvelopeExampleActionQueryFromString(rawQuery string) EnvelopeExampleActionQuery {
	v := EnvelopeExampleActionQuery{}
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
func EnvelopeExampleActionQueryFromGin(c *gin.Context) EnvelopeExampleActionQuery {
	return EnvelopeExampleActionQueryFromString(c.Request.URL.RawQuery)
}
func EnvelopeExampleActionQueryFromHttp(r *http.Request) EnvelopeExampleActionQuery {
	return EnvelopeExampleActionQueryFromString(r.URL.RawQuery)
}
func (q EnvelopeExampleActionQuery) Values() url.Values {
	return q.values
}
func (q EnvelopeExampleActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *EnvelopeExampleActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *EnvelopeExampleActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type EnvelopeExampleActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type EnvelopeExampleActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func EnvelopeExampleActionCall(
	req EnvelopeExampleActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*EnvelopeExampleActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := EnvelopeExampleActionMeta()
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
	var result EnvelopeExampleActionResult
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
