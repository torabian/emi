package external
import (
"bytes"
"encoding/json"
"fmt"
"github.com/gin-gonic/gin"
"github.com/torabian/emi/examples/fullstack/emigo"
"io"
"net/http"
"net/url"
)
/**
* Action to communicate with the action ComputeApiSseChannelAction
*/
func ComputeApiSseChannelActionMeta() struct {
    Name   string
	CliName   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        CliName   string
        URL    string
        Method string
    }{
        Name:   "ComputeApiSseChannelAction",
        CliName:   "compute-api-sse-channel-action",
        URL:    "/compute/sse/ch",
        Method: "GET",
    }
}
func GetComputeApiSseChannelActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "initial-vector1",
			Type: "slice",
		},
		{
			Name: prefix + "value",
			Type: "string?",
		},
		{
			Name: prefix + "initial-vector2",
			Type: "slice",
		},
	}
}
func CastComputeApiSseChannelActionReqFromCli(c emigo.CliCastable) ComputeApiSseChannelActionReq {
	data := ComputeApiSseChannelActionReq{}
			if c.IsSet("initial-vector1") { 
 emigo.InflatePossibleSlice(c.String("initial-vector1"), &data.InitialVector1) 
}
			if c.IsSet("value") { 
 emigo.ParseNullable(c.String("value"), &data.Value) 
}
			if c.IsSet("initial-vector2") { 
 emigo.InflatePossibleSlice(c.String("initial-vector2"), &data.InitialVector2) 
}
	return data
}
  // The base class definition for computeApiSseChannelActionReq
type ComputeApiSseChannelActionReq struct {
		InitialVector1 []int `json:"initialVector1" yaml:"initialVector1"`
		Value emigo.Nullable[string] `json:"value" yaml:"value"`
		InitialVector2 []int `json:"initialVector2" yaml:"initialVector2"`
}
func (x *ComputeApiSseChannelActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetComputeApiSseChannelActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "output-vector",
			Type: "slice",
		},
	}
}
func CastComputeApiSseChannelActionResFromCli(c emigo.CliCastable) ComputeApiSseChannelActionRes {
	data := ComputeApiSseChannelActionRes{}
			if c.IsSet("output-vector") { 
 emigo.InflatePossibleSlice(c.String("output-vector"), &data.OutputVector) 
}
	return data
}
  // The base class definition for computeApiSseChannelActionRes
type ComputeApiSseChannelActionRes struct {
		OutputVector []int `json:"outputVector" yaml:"outputVector"`
}
func (x *ComputeApiSseChannelActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
type ComputeApiSseChannelActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// ComputeApiSseChannelActionRaw registers a raw Gin route for the ComputeApiSseChannelAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ComputeApiSseChannelActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ComputeApiSseChannelActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// ComputeApiSseChannelActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ComputeApiSseChannelAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ComputeApiSseChannelActionHandler(
	handler func(c ComputeApiSseChannelActionRequest, gin *gin.Context) (*ComputeApiSseChannelActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ComputeApiSseChannelActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ComputeApiSseChannelActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ComputeApiSseChannelActionRequest{
			Body:        body,
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
// ComputeApiSseChannelAction is a high-level convenience wrapper around ComputeApiSseChannelActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ComputeApiSseChannelAction(r gin.IRoutes, handler func(c ComputeApiSseChannelActionRequest, gin *gin.Context) (*ComputeApiSseChannelActionResponse, error),) {
	method, url, h := ComputeApiSseChannelActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for ComputeApiSseChannelAction
 */
// Query wrapper with private fields
type ComputeApiSseChannelActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func ComputeApiSseChannelActionQueryFromString(rawQuery string) ComputeApiSseChannelActionQuery {
	v := ComputeApiSseChannelActionQuery{}
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
func ComputeApiSseChannelActionQueryFromGin(c *gin.Context) ComputeApiSseChannelActionQuery {
	return ComputeApiSseChannelActionQueryFromString(c.Request.URL.RawQuery)
}
func ComputeApiSseChannelActionQueryFromHttp(r *http.Request) ComputeApiSseChannelActionQuery {
	return ComputeApiSseChannelActionQueryFromString(r.URL.RawQuery)
}
func (q ComputeApiSseChannelActionQuery) Values() url.Values {
	return q.values
}
func (q ComputeApiSseChannelActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ComputeApiSseChannelActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ComputeApiSseChannelActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type ComputeApiSseChannelActionRequest struct {
	Body ComputeApiSseChannelActionReq
	QueryParams url.Values
	Headers http.Header
}
type ComputeApiSseChannelActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func ComputeApiSseChannelActionCall(
	req ComputeApiSseChannelActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ComputeApiSseChannelActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ComputeApiSseChannelActionMeta()
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
			bodyBytes, err := json.Marshal(req.Body)
			if err != nil {
				return nil, err
			}
			req0, err := http.NewRequest(meta.Method, u.String(), bytes.NewReader(bodyBytes))
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
	var result ComputeApiSseChannelActionResult
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