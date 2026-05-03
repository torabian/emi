package golang

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {

	if action.GetMethod() == "reactive" {
		return GoActionRenderReactive(action, ctx, complexes)
	}
	skipGoClient := strings.Contains(ctx.Tags, GEN_GO_SKIP_CLIENT)

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


/*
Here is a quick function implementation to make your life easier:

// Actual implementation of {{ .realms.ActionName }}
func {{ .realms.ActionName }}(c {{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error) {

	return &{{ .realms.ActionName }}Response{
		// Payload is an interface. Use it at carefully.
	}, nil
}

*/


func {{ .realms.ActionName }}Meta() struct {
    Name   string
	CliName   string
	{{ if .realms.CliShort }}
		CliShort   string
	{{ end }}
    URL    string
    Method string
	Description string
} {
    return struct {
        Name   string
        CliName   string
        {{ if .realms.CliShort }}
			CliShort   string
		{{ end }}
        URL    string
        Method string
		Description string
    }{
        Name:   "{{ .realms.ActionName }}",
        CliName:   "{{ .realms.CliName }}",
		{{ if .realms.CliShort }}
			CliShort:   "{{ .realms.CliShort }}",
		{{ end }}
        URL:    "{{ .realms.SafeUrl }}",
        Method: "{{ UPPER .action.Method }}",
		Description: ` + "`" + `{{ .action.Description }}` + "`" + `,
    }
}


{{ if .realms.RequestClass }}
	{{ b2s .realms.RequestClass.ActualScript }}
{{ end }}

{{ if .realms.ResponseClass }}
	{{ b2s .realms.ResponseClass.ActualScript }}
{{ end }}




type {{ .realms.ActionName }}Response struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}

	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}


func (x *{{ .realms.ActionName }}Response) SetContentType(contentType string) *{{ .realms.ActionName }}Response {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}

	x.Headers["Content-Type"] = contentType
	return x
}

func (x *{{ .realms.ActionName }}Response) AsStream(r io.Reader, contentType string) *{{ .realms.ActionName }}Response {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}

func (x *{{ .realms.ActionName }}Response) AsJSON(payload any) *{{ .realms.ActionName }}Response {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}


{{ if .realms.IdealResponseType }}
// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *{{ .realms.ActionName }}Response) WithIdeal(payload {{ .realms.IdealResponseType }}) *{{ .realms.ActionName }}Response {
	x.Payload = payload
	return x
}
{{ end }}

func (x *{{ .realms.ActionName }}Response) AsHTML(payload string) *{{ .realms.ActionName }}Response {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}

func (x *{{ .realms.ActionName }}Response) AsBytes(payload []byte) *{{ .realms.ActionName }}Response {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}

func (x {{ .realms.ActionName }}Response) GetStatusCode() int {
	return x.StatusCode
}

func (x {{ .realms.ActionName }}Response) GetRespHeaders() map[string]string {
	return x.Headers
}

func (x {{ .realms.ActionName }}Response) GetPayload() interface{} {
	return x.Payload
}


// {{ .realms.ActionName }}Raw registers a raw Gin route for the {{ .realms.ActionName }} action.
// This gives the developer full control over middleware, handlers, and response handling.
func {{ .realms.ActionName }}Raw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := {{ .realms.ActionName }}Meta()
	r.Handle(meta.Method, meta.URL, handlers...)
}


type {{ .realms.ActionName }}RequestSig = func(c {{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error)



// {{ .realms.ActionName }}Handler returns the HTTP method, route URL, and a typed Gin handler for the {{ .realms.ActionName }} action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func {{ .realms.ActionName }}Handler(
	handler {{ .realms.ActionName }}RequestSig,
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

// {{ .realms.ActionName }} is a high-level convenience wrapper around {{ .realms.ActionName }}Handler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func {{ .realms.ActionName }}Gin(r gin.IRoutes, handler {{ .realms.ActionName }}RequestSig,) {
	method, url, h := {{ .realms.ActionName }}Handler(handler)
	r.Handle(method, url, h)
}


{{ if .realms.PathParameter }}
	{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}

{{ if .realms.QueryParams }}
	{{ b2s .realms.QueryParams.ActualScript }}
{{ end }}
 
type {{ .realms.ActionName }}Request struct {
	{{ if .realms.RequestClassName }}
	Body {{ .realms.RequestClassName }}
	{{ else }}
	Body interface{}
	{{ end }}
	{{ if .realms.PathParameter }}
	Params {{ .realms.ActionName }}PathParameter
	{{ end }}
	QueryParams url.Values

	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header

	// Gin context for each request in case of a direct access requirement
	GinCtx      *gin.Context

	// Urfave context, per each request
	CliCtx *cli.Command

	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func (x {{ .realms.ActionName }}Request) IsGin() bool {
	return x.GinCtx != nil
}

func (x {{ .realms.ActionName }}Request) IsCli() bool {
	return x.CliCtx != nil
}

//type {{ .realms.ActionName }}Result struct {
///resp *http.Response
///	Payload interface{}
///}

{{ if .EnableClientRequest }}

func {{ .realms.ActionName }}ClientCreateUrl(
	req {{ .realms.ActionName }}Request,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := {{ .realms.ActionName }}Meta()
	urlAddr := meta.URL

	urlAddr = config.BaseURL + urlAddr

	{{ if .realms.PathParameter }}
	// In case there is a path parameter, we need to apply that.
	urlAddr = {{ .realms.ActionName }}PathParameterApply(req.Params, urlAddr)
	{{ end }}

	// Build final URL with query string
	u, err := url.Parse(urlAddr)
	if err != nil {
		return nil, err
	}

	// if UrlValues present, encode and append
	if len(req.QueryParams) > 0 {
		u.RawQuery = req.QueryParams.Encode()
	}

	return u, nil
}

func {{ .realms.ActionName }}ClientExecuteTyped(httpReq *http.Request) (*{{ .realms.ActionName }}Response, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	// At this point, response is valid, and we need to return the results.
	var result {{ .realms.ActionName }}Response
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &{{ .realms.ActionName }}Response{Payload: result}, err
	}

	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &{{ .realms.ActionName }}Response{Payload: result}, err
	}

	return &{{ .realms.ActionName }}Response{Payload: result}, nil
}

func {{ .realms.ActionName }}ClientBuildRequest(req {{ .realms.ActionName }}Request, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := {{ .realms.ActionName }}Meta()

	{{ if .realms.RequestClass }}
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), bytes.NewReader(bodyBytes))
	{{ else }}
		httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), nil)
	{{ end }}
	if err != nil {
		return nil, err
	}

	httpReq.Header = make(http.Header)

	// copy defaults
	for k, v := range config.Headers {
		for _, vv := range v {
			httpReq.Header.Add(k, vv)
		}
	}

	// override with request-specific headers
	for k, v := range req.Headers {
		httpReq.Header.Del(k) // ensure override, not duplicate
		for _, vv := range v {
			httpReq.Header.Add(k, vv)
		}
	}

	return httpReq, nil
}

func {{ .realms.ActionName }}Call(
	req {{ .realms.ActionName }}Request,
	config *emigo.APIClient, // optional pre-built request
) (*{{ .realms.ActionName }}Response, error) {

	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.

	// first we create url, apply all path parameters, query params, etc
	u, err := {{ .realms.ActionName }}ClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}

	// We create the request from the body in second stage
	r, err := {{ .realms.ActionName }}ClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}

	// This one would execute the request and cast the result.
	return {{ .realms.ActionName }}ClientExecuteTyped(r)
}

{{ end }}
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":              action,
		"realms":              realms,
		"shouldExport":        true,
		"EnableClientRequest": !skipGoClient,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".go"
	res.CodeChunkDependensies = deps

	return res, nil
}
