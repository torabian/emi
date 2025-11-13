package unknownpackage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"test.com/emi-go-header-demo/generalgen/emigo"
)

/**
* Action to communicate with the action ComputeApiAction
 */
func ComputeApiActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "ComputeApiAction",
		URL:    "",
		Method: "POST",
	}
}

// The base class definition for computeApiActionReq
type ComputeApiActionReq struct {
	InitialVector1 []int                  `json:"initialVector1" yaml:"initialVector1"`
	Value          emigo.Nullable[string] `json:"value" yaml:"value"`
	InitialVector2 []int                  `json:"initialVector2" yaml:"initialVector2"`
}

// The base class definition for computeApiActionRes
type ComputeApiActionRes struct {
	OutputVector []int `json:"outputVector" yaml:"outputVector"`
}
type ComputeApiActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// ComputeApiActionRaw registers a raw Gin route for the ComputeApiAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ComputeApiActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ComputeApiActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // ComputeApiActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ComputeApiAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ComputeApiActionHandler(
	handler func(c ComputeApiActionRequest, gin *gin.Context) (*ComputeApiActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ComputeApiActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ComputeApiActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ComputeApiActionRequest{
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

// ComputeApiAction is a high-level convenience wrapper around ComputeApiActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ComputeApiAction(r gin.IRoutes, handler func(c ComputeApiActionRequest, gin *gin.Context) (*ComputeApiActionResponse, error)) {
	method, url, h := ComputeApiActionHandler(handler)
	r.Handle(method, url, h)
}

// Using in client code.
type ComputeApiActionQuery struct {
	url.Values
}
type ComputeApiActionRequest struct {
	Body        ComputeApiActionReq
	QueryParams url.Values
	Headers     http.Header
	UrlValues   ComputeApiActionQuery
}
type ComputeApiActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ComputeApiActionCall(
	req ComputeApiActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ComputeApiActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ComputeApiActionMeta()
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
	var result ComputeApiActionResult
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
