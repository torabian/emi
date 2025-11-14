package unk

import (
	"bytes"
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
* Action to communicate with the action PatchGiantAction
 */
func PatchGiantActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "PatchGiantAction",
		URL:    "/get/giant/:id",
		Method: "PATCH",
	}
}

type PatchGiantActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// PatchGiantActionRaw registers a raw Gin route for the PatchGiantAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PatchGiantActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PatchGiantActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // PatchGiantActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PatchGiantAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PatchGiantActionHandler(
	handler func(c PatchGiantActionRequest, gin *gin.Context) (*PatchGiantActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PatchGiantActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body GiantDto
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := PatchGiantActionRequest{
			Body:        body,
			Params:      PatchGiantActionPathParameterFromGin(m),
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

// PatchGiantAction is a high-level convenience wrapper around PatchGiantActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PatchGiantAction(r gin.IRoutes, handler func(c PatchGiantActionRequest, gin *gin.Context) (*PatchGiantActionResponse, error)) {
	method, url, h := PatchGiantActionHandler(handler)
	r.Handle(method, url, h)
}

// Using in client code.
type PatchGiantActionQuery struct {
	url.Values
}

/**
 * Path parameters for PatchGiantAction
 */
type PatchGiantActionPathParameter struct {
	Id string
}

// Converts a placeholder url, and applies the parameters to it.
func PatchGiantActionPathParameterApply(params PatchGiantActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, "id", params.Id)
	return templateUrl
}

// Creates the parameters from the gin
func PatchGiantActionPathParameterFromGin(g *gin.Context) PatchGiantActionPathParameter {
	res := PatchGiantActionPathParameter{}
	res.Id = g.Param("id")
	return res
}

type PatchGiantActionRequest struct {
	Body        GiantDto
	Params      PatchGiantActionPathParameter
	QueryParams url.Values
	Headers     http.Header
	UrlValues   PatchGiantActionQuery
}
type PatchGiantActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func PatchGiantActionCall(
	req PatchGiantActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PatchGiantActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := PatchGiantActionMeta()
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
	var result PatchGiantActionResult
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
