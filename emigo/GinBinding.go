package emigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
* Request-body binding and response rendering for gin, generic enough that
* any emi-generated action handler can share it instead of reimplementing
* JSON/YAML/XML/form dispatch and error rendering per action.
*
* Bug fix: this file used to carry a `//go:build !wasm` tag, on the
* assumption that gin doesn't compile under GOOS=js/GOARCH=wasm. It does
* (confirmed directly: `GOOS=js GOARCH=wasm go build github.com/gin-gonic/gin`
* succeeds) - what a wasm build actually can't do is *run* a gin.Engine
* (no real net.Listen inside a browser sandbox), which is a separate,
* runtime-only concern nothing here triggers. The tag's real effect wasn't
* excluding gin from wasm builds at all: every emi-generated `*ActionGin.go`
* file (one per action, in the same package as its wasm-safe
* `*ActionHttpHandler` net/http counterpart - see HttpBinding.go) has no
* build tag of its own, so removing this file's `!wasm` tag was the only
* way those packages - and the real net/http handlers living alongside the
* Gin ones - could compile for wasm at all. Before this fix, `GOOS=js
* GOARCH=wasm go build ./cmd/fireback-wasm` failed outright with `undefined:
* emigo.RenderGinError` (and friends) the moment it touched any package
* with generated actions in it (modules/abac/defs, .../interfacetools/defs,
* .../messaging/defs, ...) - which is every module cmd/fireback-wasm/main.go
* imports, so nothing past `go build` itself ever ran; the ui/wasm-demo app
* had no fireback.wasm to fetch at all. The Gin-only symbols (RenderGinError,
* RenderGinResult, BindGinRequestBody, ...) below just need to exist so
* those packages build; nothing under wasm ever actually constructs a
* *gin.Context or calls them - cmd/fireback-wasm/main.go's own hand-written
* JSON stand-ins (mock /whoami, mock /passports/available-methods, ...) are
* a separate, still-open piece of work (those real actions still need a
* wasm-safe session/auth story before they can replace the stand-ins), not
* something this fix does on its own.
*
* The wasm/http1.1-free transport a wasm build's handlers actually run
* through is emi's net/http renderer (see HttpBinding.go), which
* BindGinRequestBody's format helpers are shared with.
*
* This package cannot depend on any specific application's error catalog
* (fireback's ferror, or anyone else's), so binding failures come back as the
* typed errors in BindErrors.go, each with its own generic English
* ToPublicJSON. An application that wants localized/catalog-driven messages
* instead switches on these concrete types and wraps them in its own error
* before it reaches RenderGinError - see fireback's gintools package for an
* example of that translation layer.
**/

// BindGinFormData, when non-nil, is consulted by BindGinRequestBody for a
// multipart/form-data upload's file fields. Representing an uploaded file
// (storage, mime handling, size limits, ...) is usually application-specific
// - fireback, for instance, wants a complexes.XFile - so this is left as a
// hook rather than a hard dependency. When nil, BindGinRequestBody falls back
// to a generic BindMultipartFile value per file field.
var ConvertMultipartFile func(fh *multipart.FileHeader) (any, error)

// BindMultipartFile is the generic per-file value used by BindGinRequestBody
// when ConvertMultipartFile is not set.
type BindMultipartFile struct {
	Filename string `json:"filename"`
	Mime     string `json:"mime"`
	Data     []byte `json:"data"`
}

// BindGinRequestBody reads and decodes the request body of c into body
// (a pointer), picking the decoder from the Content-Type header (JSON, YAML,
// XML, multipart/form-data, or application/x-www-form-urlencoded; JSON is the
// default when the header is missing/unrecognized). Only POST/PATCH/PUT
// requests carry a body worth reading - for any other method this is a no-op
// that returns nil, matching how the rest of HTTP treats GET/DELETE/etc.
//
// On failure it returns one of the typed errors in BindErrors.go, ready to
// hand to RenderGinError.
func BindGinRequestBody(c *gin.Context, body any) error {
	if c.Request.Method != http.MethodPost && c.Request.Method != http.MethodPatch && c.Request.Method != http.MethodPut {
		return nil
	}

	switch detectGinContentType(c) {
	case bindContentTypeYAML:
		bodyBytes, err := ginBodyToBytes(c)
		if err != nil {
			return err
		}
		return BindYamlBytes(bodyBytes, body)
	case bindContentTypeFormData:
		return BindGinMultipartForm(c, body)
	case bindContentTypeURLEncoded:
		return BindGinUrlEncoded(c, body)
	case bindContentTypeXML:
		bodyBytes, err := ginBodyToBytes(c)
		if err != nil {
			return err
		}
		return BindXmlBytes(bodyBytes, body)
	default:
		bodyBytes, err := ginBodyToBytes(c)
		if err != nil {
			return err
		}
		return BindJsonBytes(bodyBytes, body)
	}
}

type bindContentType string

const (
	bindContentTypeJSON       bindContentType = "json"
	bindContentTypeURLEncoded bindContentType = "urlencoded"
	bindContentTypeFormData   bindContentType = "form-data"
	bindContentTypeYAML       bindContentType = "yaml"
	bindContentTypeXML        bindContentType = "xml"
)

func detectGinContentType(c *gin.Context) bindContentType {
	contentType := c.GetHeader("Content-Type")

	switch {
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		return bindContentTypeURLEncoded
	case strings.HasPrefix(contentType, "multipart/form-data"):
		return bindContentTypeFormData
	case strings.HasPrefix(contentType, "application/yaml"), strings.HasPrefix(contentType, "application/x-yaml"), strings.HasPrefix(contentType, "text/yaml"):
		return bindContentTypeYAML
	case strings.HasPrefix(contentType, "application/xml"), strings.HasPrefix(contentType, "text/xml"):
		return bindContentTypeXML
	default:
		return bindContentTypeJSON
	}
}

func ginBodyToBytes(c *gin.Context) ([]byte, error) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		switch {
		case errors.Is(err, io.EOF):
			return nil, &BindBodyReadError{Kind: "empty", Err: err}
		case errors.Is(err, io.ErrUnexpectedEOF):
			return nil, &BindBodyReadError{Kind: "unexpectedEof", Err: err}
		case errors.Is(err, http.ErrBodyReadAfterClose):
			return nil, &BindBodyReadError{Kind: "readAfterClose", Err: err}
		default:
			return nil, &BindBodyReadError{Kind: "unknown", Err: err}
		}
	}

	// Reset the body so it can be read again later (e.g. by logging middleware).
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}

// BindGinMultipartForm parses c's multipart/form-data body into target: each
// non-file field round-trips through JSON (so it lands in whatever struct
// tags a JSON body would use), and each uploaded file becomes a
// BindMultipartFile unless ConvertMultipartFile is set. Exported (rather than
// folded privately into BindGinRequestBody) so a caller that already knows
// its content-type can invoke it directly.
func BindGinMultipartForm(c *gin.Context, target any) error {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB limit
		return &BindDecodingError{Format: "form", Err: err}
	}

	formData := c.Request.MultipartForm
	formMap := make(map[string]any)

	for fieldName, files := range formData.File {
		for _, fileHeader := range files {
			if ConvertMultipartFile != nil {
				converted, err := ConvertMultipartFile(fileHeader)
				if err != nil {
					return &BindDecodingError{Format: "form", Err: err}
				}
				formMap[fieldName] = converted
				continue
			}

			file, err := fileHeader.Open()
			if err != nil {
				return &BindDecodingError{Format: "form", Err: err}
			}
			data, err := io.ReadAll(file)
			file.Close()
			if err != nil {
				return &BindDecodingError{Format: "form", Err: err}
			}

			formMap[fieldName] = BindMultipartFile{
				Filename: fileHeader.Filename,
				Mime:     fileHeader.Header.Get("Content-Type"),
				Data:     data,
			}
		}
	}

	for key, values := range formData.Value {
		if len(values) > 1 {
			formMap[key] = values
		} else {
			formMap[key] = values[0]
		}
	}

	return bindFormMap(formMap, target)
}

// BindGinUrlEncoded parses c's application/x-www-form-urlencoded body into
// target, round-tripping the flattened form values through JSON. Exported
// for the same reason as BindGinMultipartForm.
func BindGinUrlEncoded(c *gin.Context, target any) error {
	if err := c.Request.ParseForm(); err != nil {
		return &BindDecodingError{Format: "form", Err: err}
	}

	formMap := make(map[string]any)
	for key, values := range c.Request.Form {
		if len(values) > 1 {
			formMap[key] = values
		} else {
			formMap[key] = values[0]
		}
	}

	return bindFormMap(formMap, target)
}

// RenderGinError writes err to the response as JSON and aborts the request.
// If err (or something it wraps) implements ToPublicJSON(lang string)
// ([]byte, int32) - as fireback's ferror.Error and every type in
// BindErrors.go do - that resolves the response body and status, picking the
// language the same way the rest of the app resolves it: the
// "acceptLanguage" query param first, else the Accept-Language header, else
// "en". Otherwise, an error that stringifies itself as an indented JSON
// object (optionally with its own "httpCode") is forwarded as-is; anything
// else falls back to a plain {"error": err.Error()} body.
func RenderGinError(c *gin.Context, err error) {
	// Some deeper call (e.g. a security/authorization check that rejects the
	// request before the handler's own business logic ever runs) may have
	// already written and aborted the response itself - gin tracks that on
	// the ResponseWriter regardless of who did the writing. Rendering the
	// bubbled-up error on top of that would append a second, invalid JSON
	// body after the first.
	if c.Writer.Written() {
		return
	}

	status := http.StatusInternalServerError

	if converter, ok := err.(interface {
		ToPublicJSON(lang string) ([]byte, int32)
	}); ok {
		lang := c.Query("acceptLanguage")
		if lang == "" {
			lang = c.GetHeader("Accept-Language")
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
		// Nest the resolved object under "error" (rather than writing it as
		// the bare response body) so every error shape - this one, the
		// generic forwarded-JSON one below, and the plain-string one -
		// answers with the same {"error": ...} envelope. json.RawMessage
		// keeps body embedded as real JSON instead of being re-escaped into
		// a string.
		c.JSON(status, gin.H{"error": json.RawMessage(body)})
		return
	}

	// Otherwise, other action errors may still stringify themselves as an
	// indented JSON object via their Error() method. If that's what we got,
	// forward it nested under "error" as real JSON (optionally honoring its
	// own "httpCode" field for the response status) instead of re-escaping
	// it into a string, which is what plain errors still get.
	msg := err.Error()
	trimmed := strings.TrimSpace(msg)
	if strings.HasPrefix(trimmed, "{") && json.Valid([]byte(trimmed)) {
		var probe struct {
			HttpCode int32 `json:"httpCode"`
		}
		if uErr := json.Unmarshal([]byte(trimmed), &probe); uErr == nil && probe.HttpCode != 0 {
			status = int(probe.HttpCode)
		}
		c.JSON(status, gin.H{"error": json.RawMessage(trimmed)})
		return
	}
	c.JSON(status, gin.H{"error": msg})
}

// RenderGinResult writes a successful action response: headers, then status
// (defaulting to 200) and JSON payload, or a bare status when there's no
// payload. Every emi-generated {{Action}}Response already satisfies
// EmiActionResult.
func RenderGinResult(c *gin.Context, resp EmiActionResult) {
	for k, v := range resp.GetRespHeaders() {
		c.Header(k, v)
	}

	status := resp.GetStatusCode()
	if status == 0 {
		status = http.StatusOK
	}

	if payload := resp.GetPayload(); payload != nil {
		c.JSON(status, payload)
	} else {
		c.Status(status)
	}
}
