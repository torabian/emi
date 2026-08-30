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

// RenderHttpError writes err to w as JSON. See RenderGinError for the shared
// rendering rules (ToPublicJSON, JSON-string forwarding, plain fallback);
// this is the same logic against http.ResponseWriter/*http.Request instead
// of *gin.Context.
func RenderHttpError(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusInternalServerError
	w.Header().Set("Content-Type", "application/json")

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
		wrapped, wErr := json.Marshal(map[string]json.RawMessage{"error": json.RawMessage(body)})
		w.WriteHeader(status)
		if wErr == nil {
			w.Write(wrapped)
		} else {
			w.Write(body)
		}
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
		wrapped, wErr := json.Marshal(map[string]json.RawMessage{"error": json.RawMessage(trimmed)})
		w.WriteHeader(status)
		if wErr == nil {
			w.Write(wrapped)
		} else {
			w.Write([]byte(trimmed))
		}
		return
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// RenderHttpResult writes a successful action response: headers, then status
// (defaulting to 200) and JSON payload, or a bare status when there's no
// payload. Every emi-generated {{Action}}Response already satisfies
// EmiActionResult.
func RenderHttpResult(w http.ResponseWriter, resp EmiActionResult) {
	for k, v := range resp.GetRespHeaders() {
		w.Header().Set(k, v)
	}

	status := resp.GetStatusCode()
	if status == 0 {
		status = http.StatusOK
	}

	if payload := resp.GetPayload(); payload != nil {
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", "application/json")
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(payload)
	} else {
		w.WriteHeader(status)
	}
}
