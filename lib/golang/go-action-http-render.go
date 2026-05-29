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
// Unlike the gin transport this carries no external dependency, so it also
// compiles under GOOS=js/wasm. The same typed handler runs on a real net/http
// server AND inside wasm via httptest.ResponseRecorder.
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
// returns either an *{{ .realms.ActionName }}Response or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func {{ .realms.ActionName }}HttpHandler(
	handler func(c {{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := {{ .realms.ActionName }}Meta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		{{ if .realms.RequestClassName }}
		var body {{ .realms.RequestClassName }}
		if r.Body != nil {
			defer r.Body.Close()
			if data, _ := io.ReadAll(r.Body); len(data) > 0 {
				if err := json.Unmarshal(data, &body); err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON: " + err.Error()})
					return
				}
			}
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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// If the handler returned nil (and no error), the response was handled
		// manually.
		if resp == nil {
			return
		}

		// Apply headers
		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}

		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}

		if resp.Payload != nil {
			if w.Header().Get("Content-Type") == "" {
				w.Header().Set("Content-Type", "application/json")
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(resp.Payload)
		} else {
			w.WriteHeader(status)
		}
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
		{Location: "encoding/json"},
		{Location: "io"},
	}

	return &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		CodeChunkDependensies: deps,
		SuggestedFileName:     realms.ActionName + "Http",
		SuggestedExtension:    ".go",
	}, nil
}
