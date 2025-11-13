package unk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"test/emi/emigo"

	"github.com/gin-gonic/gin"
)

/**
* Action to communicate with the action GetAsGiantsAction
 */
func GetAsGiantsActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "GetAsGiantsAction",
		URL:    "/get/giant/:id",
		Method: "GET",
	}
}

type GetAsGiantsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// GetAsGiantsActionRaw registers a raw Gin route for the GetAsGiantsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetAsGiantsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetAsGiantsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // GetAsGiantsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetAsGiantsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetAsGiantsActionHandler(
	handler func(c GetAsGiantsActionRequest, gin *gin.Context) (*GetAsGiantsActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetAsGiantsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetAsGiantsActionRequest{
			Params:      GetAsGiantsActionPathParameterFromGin(m),
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

// GetAsGiantsAction is a high-level convenience wrapper around GetAsGiantsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetAsGiantsAction(r gin.IRoutes, handler func(c GetAsGiantsActionRequest, gin *gin.Context) (*GetAsGiantsActionResponse, error)) {
	method, url, h := GetAsGiantsActionHandler(handler)
	r.Handle(method, url, h)
}

// Using in client code.
type GetAsGiantsActionQuery struct {
	url.Values
}

/**
 * Path parameters for GetAsGiantsAction
 */
type GetAsGiantsActionPathParameter struct {
	Id string
}

// Converts a placeholder url, and applies the parameters to it.
func GetAsGiantsActionPathParameterApply(params GetAsGiantsActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, "id", params.Id)
	return templateUrl
}

// Creates the parameters from the gin
func GetAsGiantsActionPathParameterFromGin(g *gin.Context) GetAsGiantsActionPathParameter {
	res := GetAsGiantsActionPathParameter{}
	res.Id = g.Param("id")
	return res
}

type GetAsGiantsActionRequest struct {
	Params      GetAsGiantsActionPathParameter
	QueryParams url.Values
	Headers     http.Header
	UrlValues   GetAsGiantsActionQuery
}
type GetAsGiantsActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func GetAsGiantsActionCall(
	req GetAsGiantsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetAsGiantsActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetAsGiantsActionMeta()
		baseURL := meta.URL
		// Build final URL with query string
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}
		// if UrlValues present, encode and append
		if len(req.UrlValues.Values) > 0 {
			u.RawQuery = req.UrlValues.Encode()
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
	var result GetAsGiantsActionResult
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
