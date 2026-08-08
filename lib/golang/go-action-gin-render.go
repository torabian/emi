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

{{ if .realms.PathParameterGin }}
	{{ b2s .realms.PathParameterGin.ActualScript }}
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
			// Some deeper call inside handler (e.g. a security/authorization check
			// that rejects the request before the handler's own business logic ever
			// runs) may have already written and aborted the response itself - gin
			// tracks that on the ResponseWriter regardless of who did the writing.
			// Rendering the bubbled-up error on top of that would append a second,
			// invalid JSON body after the first.
			if m.Writer.Written() {
				return
			}

			status := http.StatusInternalServerError

			// If the error knows how to render itself for a given language (e.g.
			// fireback.IError, whose ferror.Error.ToPublicJSON resolves its
			// {"$": ..., "en": ..., "fa": ...} message map down to one string), let it -
			// picking the language the same way the rest of the app resolves it: the
			// "acceptLanguage" query param first, else the Accept-Language header, else
			// "en".
			if converter, ok := err.(interface {
				ToPublicJSON(lang string) ([]byte, int32)
			}); ok {
				lang := m.Query("acceptLanguage")
				if lang == "" {
					lang = m.GetHeader("Accept-Language")
					if i := strings.IndexAny(lang, ",;-"); i >= 0 {
						lang = lang[:i]
					}
					lang = strings.ToLower(strings.TrimSpace(lang))
				}
				if lang == "" {
					lang = "en"
				}
				body, code := converter.ToPublicJSON(lang)
				if code != 0 {
					status = int(code)
				}
				// Nest the resolved object under "error" (rather than writing it as the
				// bare response body) so every error shape - this one, the generic
				// forwarded-JSON one below, and the plain-string one - answers with the
				// same {"error": ...} envelope. json.RawMessage keeps body embedded as
				// real JSON instead of being re-escaped into a string.
				m.JSON(status, gin.H{"error": json.RawMessage(body)})
				return
			}

			// Otherwise, other action errors may still stringify themselves as an
			// indented JSON object via their Error() method. If that's what we got,
			// forward it nested under "error" as real JSON (optionally honoring its own
			// "httpCode" field for the response status) instead of re-escaping it into a
			// string, which is what plain errors still get.
			msg := err.Error()
			trimmed := strings.TrimSpace(msg)
			if strings.HasPrefix(trimmed, "{") && json.Valid([]byte(trimmed)) {
				var probe struct {
					HttpCode int32 ` + "`json:\"httpCode\"`" + `
				}
				if uErr := json.Unmarshal([]byte(trimmed), &probe); uErr == nil && probe.HttpCode != 0 {
					status = int(probe.HttpCode)
				}
				m.JSON(status, gin.H{"error": json.RawMessage(trimmed)})
				return
			}
			m.JSON(status, gin.H{"error": msg})
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
		{Location: "strings"},
		{Location: "encoding/json"},
		{Location: "github.com/gin-gonic/gin"},
	}

	return &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		CodeChunkDependensies: deps,
		SuggestedFileName:     realms.ActionName + "Gin",
		SuggestedExtension:    ".go",
	}, nil
}
