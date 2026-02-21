package external
import (
"bytes"
"encoding"
"encoding/json"
"fmt"
"github.com/gin-gonic/gin"
"github.com/torabian/emi/examples/fullstack/emigo"
"io"
"math/big"
"net/http"
"net/url"
"strings"
)
/**
* Action to communicate with the action ComputeExpAction
*/
func ComputeExpActionMeta() struct {
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
        Name:   "ComputeExpAction",
        CliName:   "compute-exp-action",
        URL:    "/big/exp/:first/:second",
        Method: "",
    }
}
func GetComputeExpActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "base",
			Type: "complex",
		},
		{
			Name: prefix + "exponent",
			Type: "complex",
		},
	}
}
func CastComputeExpActionReqFromCli(c emigo.CliCastable) ComputeExpActionReq {
	data := ComputeExpActionReq{}
			if c.IsSet("base") { 
 if u, ok := any(&data.Base).(encoding.TextUnmarshaler); ok { u.UnmarshalText([]byte(c.String("base"))) } 
}
			if c.IsSet("exponent") { 
 if u, ok := any(&data.Exponent).(encoding.TextUnmarshaler); ok { u.UnmarshalText([]byte(c.String("exponent"))) } 
}
	return data
}
  // The base class definition for computeExpActionReq
type ComputeExpActionReq struct {
		Base big.Int `json:"base" yaml:"base"`
		Exponent big.Int `json:"exponent" yaml:"exponent"`
}
func (x *ComputeExpActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetComputeExpActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "result",
			Type: "complex",
		},
	}
}
func CastComputeExpActionResFromCli(c emigo.CliCastable) ComputeExpActionRes {
	data := ComputeExpActionRes{}
			if c.IsSet("result") { 
 if u, ok := any(&data.Result).(encoding.TextUnmarshaler); ok { u.UnmarshalText([]byte(c.String("result"))) } 
}
	return data
}
  // The base class definition for computeExpActionRes
type ComputeExpActionRes struct {
		Result  big.Int `json:"result" yaml:"result"`
}
func (x *ComputeExpActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
type ComputeExpActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}
// ComputeExpActionRaw registers a raw Gin route for the ComputeExpAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ComputeExpActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ComputeExpActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}// ComputeExpActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ComputeExpAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ComputeExpActionHandler(
	handler func(c ComputeExpActionRequest, gin *gin.Context) (*ComputeExpActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ComputeExpActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ComputeExpActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ComputeExpActionRequest{
			Body:        body,
			Params: ComputeExpActionPathParameterFromGin(m),
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
// ComputeExpAction is a high-level convenience wrapper around ComputeExpActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ComputeExpAction(r gin.IRoutes, handler func(c ComputeExpActionRequest, gin *gin.Context) (*ComputeExpActionResponse, error),) {
	method, url, h := ComputeExpActionHandler(handler)
	r.Handle(method, url, h)
}
	/**
 * Path parameters for ComputeExpAction
 */
type ComputeExpActionPathParameter struct {
	First string
	Second string
}
// Converts a placeholder url, and applies the parameters to it.
func ComputeExpActionPathParameterApply(params ComputeExpActionPathParameter, templateUrl string) string {
		templateUrl = strings.ReplaceAll(templateUrl, "first", fmt.Sprintf("%v", params.First))
		templateUrl = strings.ReplaceAll(templateUrl, "second", fmt.Sprintf("%v", params.Second))
	return templateUrl
}
// Creates the parameters from the gin
// Creates the parameters from the gin
func ComputeExpActionPathParameterFromGin(g *gin.Context) ComputeExpActionPathParameter {
	res := ComputeExpActionPathParameter{}
			res.First = g.Param("first")
			res.Second = g.Param("second")
	return res
}
	/**
 * Query parameters for ComputeExpAction
 */
// Query wrapper with private fields
type ComputeExpActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}
func ComputeExpActionQueryFromString(rawQuery string) ComputeExpActionQuery {
	v := ComputeExpActionQuery{}
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
func ComputeExpActionQueryFromGin(c *gin.Context) ComputeExpActionQuery {
	return ComputeExpActionQueryFromString(c.Request.URL.RawQuery)
}
func ComputeExpActionQueryFromHttp(r *http.Request) ComputeExpActionQuery {
	return ComputeExpActionQueryFromString(r.URL.RawQuery)
}
func (q ComputeExpActionQuery) Values() url.Values {
	return q.values
}
func (q ComputeExpActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ComputeExpActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ComputeExpActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}
type ComputeExpActionRequest struct {
	Body ComputeExpActionReq
	Params ComputeExpActionPathParameter
	QueryParams url.Values
	Headers http.Header
}
type ComputeExpActionResult struct {
	resp *http.Response                      // embed original response
	Payload interface{}
}
func ComputeExpActionCall(
	req ComputeExpActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ComputeExpActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ComputeExpActionMeta()
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
	var result ComputeExpActionResult
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