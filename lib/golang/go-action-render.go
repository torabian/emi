package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	realms, deps, err := GoActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  core.TOKEN_ORIGINAL_NAME,
				Value: realms.ActionName,
			},
		},
	}

	const tmpl = `/**
* Action to communicate with the action {{ .realms.ActionName }}
*/


func {{ .realms.ActionName }}Meta() struct {
    Name   string
    URL    string
    Method string
} {
    return struct {
        Name   string
        URL    string
        Method string
    }{
        Name:   "{{ .realms.ActionName }}",
        URL:    "{{ .action.Url }}",
        Method: "{{ .action.Method }}",
    }
}


{{ if .realms.RequestClass }}
	{{ b2s .realms.RequestClass.ActualScript }}
{{ end }}

{{ if .realms.ResponseClass }}
	{{ b2s .realms.ResponseClass.ActualScript }}
{{ end }}


// {{ .realms.ActionName }}Request wraps the current Gin context.
// It can be extended later to include typed request data (headers, body, params, etc.)
// so developers can access both the raw context and strongly typed fields.
type {{ .realms.ActionName }}Request struct {
	C *gin.Context
}

type {{ .realms.ActionName }}Response struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// {{ .realms.ActionName }}Raw registers a raw Gin route for the {{ .realms.ActionName }} action.
// This gives the developer full control over middleware, handlers, and response handling.
func {{ .realms.ActionName }}Raw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := {{ .realms.ActionName }}Meta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// {{ .realms.ActionName }}Handler returns the HTTP method, route URL, and a typed Gin handler for the {{ .realms.ActionName }} action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func {{ .realms.ActionName }}Handler(
	handler func(c {{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error),
) (method, url string, h gin.HandlerFunc) {
	meta := {{ .realms.ActionName }}Meta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		req := {{ .realms.ActionName }}Request{C: m}

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

// {{ .realms.ActionName }} is a high-level convenience wrapper around {{ .realms.ActionName }}Handler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func {{ .realms.ActionName }}(r gin.IRoutes, handler func({{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error)) {
	method, url, h := {{ .realms.ActionName }}Handler(handler)
	r.Handle(method, url, h)
}


// Using in client code.

type {{ .realms.ActionName }}Query struct {
	url.Values
}

type {{ .realms.ActionName }}Context struct {
	Body {{ .realms.ActionName }}Req
	QueryParams interface{}
	Headers http.Header
	UrlValues   {{ .realms.ActionName }}Query
}

type {{ .realms.ActionName }}Result struct {
	*http.Response                      // embed original response
	Payload interface{}
}


func {{ .realms.ActionName }}Call(
	req {{ .realms.ActionName }}Context,
) (*{{ .realms.ActionName }}Result, error) {
	meta := {{ .realms.ActionName }}Meta()

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

	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(meta.Method, u.String(), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	httpReq.Header = req.Headers

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed: %s", respBody)
	}

	var result {{ .realms.ActionName }}Result
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}



`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":       action,
		"realms":       realms,
		"shouldExport": true,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".go"
	res.CodeChunkDependensies = deps

	return res, nil
}
