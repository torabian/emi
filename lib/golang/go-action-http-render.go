package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// GoActionHttpRender renders the std net/http portion of an action: a typed
// handler plus a convenience wrapper that registers the route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
//
// Unlike the gin transport this carries no external (non-stdlib) dependency
// beyond emigo itself - whose HttpBinding.go is wasm-safe by the same rule -
// so it also compiles under GOOS=js/wasm. The same typed handler runs on a
// real net/http server AND inside wasm via httptest.ResponseRecorder.
//
// The caller decides whether to append this to the main action file or emit it
// as its own file (controlled by the "split-http" tag).
func GoActionHttpRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	realms goActionRealms,
) (*core.CodeChunkCompiled, error) {

	const tmpl = `
// {{ .realms.ActionName }}HttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the {{ .realms.ActionName }} action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *{{ .realms.ActionName }}Response or nil. Body binding, headers, status
// codes, and errors are all handled by emigo - see BindHttpRequestBody, RenderHttpError
// and RenderHttpResult in github.com/torabian/emi/emigo.
func {{ .realms.ActionName }}HttpHandler(
	handler func(c {{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := {{ .realms.ActionName }}Meta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		{{ if .realms.RequestClassName }}
		var body {{ .realms.RequestClassName }}
		if err := emigo.BindHttpRequestBody(r, &body); err != nil {
			emigo.RenderHttpError(w, r, err)
			return
		}
		{{ end }}

		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := {{ .realms.ActionName }}Request{
			{{ if .realms.RequestClassName }}
			Body:        body,
			{{ else }}
			Body: nil,
			{{ end }}
			{{ if .realms.PathParameter }}
			Params: {{ .realms.ActionName }}PathParameterFromFn(func(key string) string {
				return r.PathValue(key)
			}),
			{{ end }}
			QueryParams: r.URL.Query(),
			Headers:     r.Header,
		}

		resp, err := handler(req)
		if err != nil {
			emigo.RenderHttpError(w, r, err)
			return
		}

		// If the handler returned nil (and no error), the response was handled
		// manually.
		if resp == nil {
			return
		}

		emigo.RenderHttpResult(w, r, resp)
	}
}

// {{ .realms.ActionName }}Http is a high-level convenience wrapper around
// {{ .realms.ActionName }}HttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func {{ .realms.ActionName }}Http(
	mux *http.ServeMux,
	handler func(c {{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error),
) {
	method, pattern, h := {{ .realms.ActionName }}HttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
`

	t := template.Must(template.New("action_http").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action": action,
		"realms": realms,
	}); err != nil {
		return nil, err
	}

	deps := []core.CodeChunkDependency{
		{Location: "net/http"},
		{Location: "github.com/torabian/emi/emigo"},
	}

	return &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		CodeChunkDependensies: deps,
		SuggestedFileName:     realms.ActionName + "Http",
		SuggestedExtension:    ".go",
	}, nil
}
