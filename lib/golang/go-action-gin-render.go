package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// GoActionGinRender renders the gin-specific portion of an action: the
// raw/handler/convenience route wrappers plus the IsGin() helper method
// on the action's Request type.
//
// The caller decides whether to append this to the main action file or
// emit it as its own file (controlled by the "split-gin" tag).
func GoActionGinRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	realms goActionRealms,
) (*core.CodeChunkCompiled, error) {

	const tmpl = `
{{ if .realms.SkipGinWasm }}
//go:build !wasm
{{ end }}
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
		{{ if .realms.RequestClassName }}
		var body {{ .realms.RequestClassName }}
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		{{ end }}

		// Build typed request wrapper
		req := {{ .realms.ActionName }}Request{
			{{ if .realms.RequestClassName }}
				Body:        body,
			{{ else }}
				Body: nil,
			{{ end }}
			{{ if .realms.PathParameter }}
			Params: {{ .realms.ActionName }}PathParameterFromGin(m),
			{{ end }}
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx: m,
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

// {{ .realms.ActionName }}Gin is a high-level convenience wrapper around {{ .realms.ActionName }}Handler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func {{ .realms.ActionName }}Gin(r gin.IRoutes, handler func(c {{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error),) {
	method, url, h := {{ .realms.ActionName }}Handler(handler)
	r.Handle(method, url, h)
}

func (x {{ .realms.ActionName }}Request) IsGin() bool {
	if x.GinCtx == nil {
		return false
	}

	v := reflect.ValueOf(x.GinCtx)

	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}

	return true
}

func {{ .realms.ActionName }}QueryFromGin(c *gin.Context) {{ .realms.ActionName }}Query {
	return {{ .realms.ActionName }}QueryFromString(c.Request.URL.RawQuery)
}
`

	t := template.Must(template.New("action_gin").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action": action,
		"realms": realms,
	}); err != nil {
		return nil, err
	}

	deps := []core.CodeChunkDependency{
		{Location: "net/http"},
		{Location: "reflect"},
		{Location: "github.com/gin-gonic/gin"},
	}

	return &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		CodeChunkDependensies: deps,
		SuggestedFileName:     realms.ActionName + "Gin",
		SuggestedExtension:    ".go",
	}, nil
}
