package emigo

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

/**
* The net/http counterpart of GinBinding.go. Deliberately gin-free so it
* compiles under GOOS=js/wasm as well as a real net/http server - the same
* typed handler emi generates runs on both. Only JSON is supported here
* (matching what the generated {{Action}}HttpHandler needs); an application
* that wants YAML/XML/form support over plain net/http can call
* BindJsonBytes/BindYamlBytes/BindXmlBytes from BindFormats.go directly.
**/

// BindHttpRequestBody reads and JSON-decodes r's body into body (a pointer).
// A missing/empty body is not an error - the target is simply left
// untouched, since net/http (unlike gin's ShouldBindJSON) has no separate
// "was there a body at all" signal to key off. On a decoding failure it
// returns one of the typed errors in BindErrors.go, ready to hand to
// RenderHttpError.
func BindHttpRequestBody(r *http.Request, body any) error {
	if r.Body == nil {
		return nil
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return &BindBodyReadError{Kind: "unknown", Err: err}
	}
	if len(data) == 0 {
		return nil
	}

	return BindJsonBytes(data, body)
}

// writeHttpEnvelope writes status/envelope to w in the format the request's
// Accept header names (see detectAcceptFormat), falling back to JSON if
// that format's marshaler rejects envelope (all defined formats accept a
// map[string]any today, but this keeps a response going out either way).
func writeHttpEnvelope(w http.ResponseWriter, accept string, status int, envelope map[string]any) {
	format := detectAcceptFormat(accept)
	body, err := marshalResponse(format, envelope)
	if err != nil {
		format = responseFormatJSON
		body, err = marshalResponse(format, envelope)
		if err != nil {
			// json.Marshal on a map[string]any is not expected to fail;
			// this is the last-resort path if it somehow does.
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", format.contentType())
	w.WriteHeader(status)
	w.Write(body)
}

// RenderHttpError writes err to w in the format the request's Accept header
// names (see detectAcceptFormat). See RenderGinError for the shared
// rendering rules (ToPublicJSON, JSON-string forwarding, plain fallback);
// this is the same logic against http.ResponseWriter/*http.Request instead
// of *gin.Context.
func RenderHttpError(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusInternalServerError
	accept := r.Header.Get("Accept")

	if converter, ok := err.(interface {
		ToPublicJSON(lang string) ([]byte, int32)
	}); ok {
		lang := r.URL.Query().Get("acceptLanguage")
		if lang == "" {
			lang = r.Header.Get("Accept-Language")
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
		// Decode body back into a generic value (rather than keeping it as
		// json.RawMessage) so writeHttpEnvelope can re-encode it as YAML
		// when that's what the client asked for, instead of dumping raw
		// JSON bytes as a base64-ish YAML string.
		var decoded any
		if err := json.Unmarshal(body, &decoded); err != nil {
			decoded = json.RawMessage(body)
		}
		writeHttpEnvelope(w, accept, status, map[string]any{"error": decoded})
		return
	}

	msg := err.Error()
	trimmed := strings.TrimSpace(msg)
	if strings.HasPrefix(trimmed, "{") && json.Valid([]byte(trimmed)) {
		var probe struct {
			HttpCode int32 `json:"httpCode"`
		}
		if uErr := json.Unmarshal([]byte(trimmed), &probe); uErr == nil && probe.HttpCode != 0 {
			status = int(probe.HttpCode)
		}
		var decoded any
		if err := json.Unmarshal([]byte(trimmed), &decoded); err != nil {
			decoded = trimmed
		}
		writeHttpEnvelope(w, accept, status, map[string]any{"error": decoded})
		return
	}
	writeHttpEnvelope(w, accept, status, map[string]any{"error": msg})
}

// RenderHttpResult writes a successful action response: headers, then status
// (defaulting to 200) and payload, or a bare status when there's no payload.
// The payload is rendered in whatever format the request's Accept header
// names (see detectAcceptFormat: JSON, YAML, TOML, or CSV), JSON by
// default. Every emi-generated {{Action}}Response already satisfies
// EmiActionResult.
func RenderHttpResult(w http.ResponseWriter, r *http.Request, resp EmiActionResult) {
	for k, v := range resp.GetRespHeaders() {
		w.Header().Set(k, v)
	}

	status := resp.GetStatusCode()
	if status == 0 {
		status = http.StatusOK
	}

	payload := resp.GetPayload()
	if payload == nil {
		w.WriteHeader(status)
		return
	}

	format := detectAcceptFormat(r.Header.Get("Accept"))
	body, err := marshalResponse(format, payload)
	if err != nil {
		// The requested format's marshaler rejected this payload (e.g. a
		// TOML/CSV shape it can't represent) - fall back to JSON rather
		// than dropping the response.
		format = responseFormatJSON
		body, err = json.Marshal(payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", format.contentType())
	}
	w.WriteHeader(status)
	w.Write(body)
}
