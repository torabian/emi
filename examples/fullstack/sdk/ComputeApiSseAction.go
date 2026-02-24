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
* Action to communicate with the action ComputeApiSseAction
*/
func ComputeApiSseActionMeta() struct {
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
        Name:   "ComputeApiSseAction",
        CliName:   "compute-api-sse-action",
        URL:    "/compute/sse",
        Method: "GET",
    }
}
func GetComputeApiSseActionReqCliFlags(prefix string) []emigo.CliFlag {
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
func CastComputeApiSseActionReqFromCli(c emigo.CliCastable) ComputeApiSseActionReq {
	data := ComputeApiSseActionReq{}
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
  // The base class definition for computeApiSseActionReq
type ComputeApiSseActionReq struct {
		InitialVector1 []int `json:"initialVector1" yaml:"initialVector1"`
		Value emigo.Nullable[string] `yaml:"value" json:"value"`
		InitialVector2 []int `json:"initialVector2" yaml:"initialVector2"`
}
func (x *ComputeApiSseActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetComputeApiSseActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "output-vector",
			Type: "slice",
		},
	}
}
func CastComputeApiSseActionResFromCli(c emigo.CliCastable) ComputeApiSseActionRes {
	data := ComputeApiSseActionRes{}
			if c.IsSet("output-vector") { 
 emigo.InflatePossibleSlice(c.String("output-vector"), &data.OutputVector) 
}
	return data
}
  // The base class definition for computeApiSseActionRes
type ComputeApiSseActionRes struct {
		OutputVector []int `json:"outputVector" yaml:"outputVector"`
}
func (x *ComputeApiSseActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
type ComputeApiSseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// ComputeApiSseActionRaw registers a raw Gin route for the ComputeApiSseAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ComputeApiSseActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ComputeApiSseActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// ComputeApiSseActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ComputeApiSseAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ComputeApiSseActionHandler(
	handler func(c ComputeApiSseActionRequest, gin *gin.Context) (*ComputeApiSseActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ComputeApiSseActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ComputeApiSseActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ComputeApiSseActionRequest{
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
// ComputeApiSseAction is a high-level convenience wrapper around ComputeApiSseActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ComputeApiSseAction(r gin.IRoutes, handler func(c ComputeApiSseActionRequest, gin *gin.Context) (*ComputeApiSseActionResponse, error),) {
	method, url, h := ComputeApiSseActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Query parameters for ComputeApiSseAction
 */
// Query wrapper with private fields
type ComputeApiSseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func ComputeApiSseActionQueryFromString(rawQuery string) ComputeApiSseActionQuery {
	v := ComputeApiSseActionQuery{}
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
func ComputeApiSseActionQueryFromGin(c *gin.Context) ComputeApiSseActionQuery {
	return ComputeApiSseActionQueryFromString(c.Request.URL.RawQuery)
}
func ComputeApiSseActionQueryFromHttp(r *http.Request) ComputeApiSseActionQuery {
	return ComputeApiSseActionQueryFromString(r.URL.RawQuery)
}
func (q ComputeApiSseActionQuery) Values() url.Values {
	return q.values
}
func (q ComputeApiSseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ComputeApiSseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ComputeApiSseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type ComputeApiSseActionRequest struct {
	Body ComputeApiSseActionReq
	QueryParams url.Values
	Headers http.Header
}
type ComputeApiSseActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func ComputeApiSseActionCall(
	req ComputeApiSseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ComputeApiSseActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ComputeApiSseActionMeta()
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
	var result ComputeApiSseActionResult
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