package emigo

import (
	"encoding/json"
	"fmt"
)

/**
* Typed errors returned by the request-body binders (BindGinRequestBody,
* BindHttpRequestBody, and their format-specific helpers below). They carry
* structured fields (offset/line/col, field name, expected/actual type, ...)
* instead of a single flattened message string, so a caller that wants
* localized/catalog-driven messages (see fireback's gintools package) can
* switch on the concrete type and produce its own translated error, while a
* caller with no catalog of its own still gets a sane English message and a
* renderable JSON body for free via ToPublicJSON, which
* RenderGinError/RenderHttpError already know how to render.
*
* The body shape - {"message": <code>, "messageTranslated": <text>} -
* matches fireback's ferror.PublicError (the actual public-facing error
* projection apps built on this already expect), not ferror.Error's
* internal {"$": <code>, "en": <text>, "fa": <text>, ...} multi-language
* map. This package has no dependency on ferror, but messageTranslated is
* resolved for real against BindErrorCatalog.go's own small catalog (en,
* fa, pl, ru, es - see that file), keyed off the same lang RenderGinError/
* RenderHttpError already resolve from "?acceptLanguage" or
* Accept-Language before calling ToPublicJSON.
**/

func bindErrorPublicJSON(code string, lang string, args ...any) ([]byte, int32) {
	b, _ := json.Marshal(map[string]string{"message": code, "messageTranslated": bindErrorMessage(code, lang, args...)})
	return b, 400
}

// bindErrorPublicJSONAt is bindErrorPublicJSON plus the offset/line/col a
// syntax error was found at, so a client doesn't just hear "the body is
// malformed" but where - both in the human-readable messageTranslated text
// and, nested under "messageParams" (matching fireback's
// ferror.PublicError.MessageParams field), as separate numeric values a
// caller can highlight in an editor/UI without re-parsing the sentence.
func bindErrorPublicJSONAt(code string, lang string, offset int64, line int, col int, args ...any) ([]byte, int32) {
	b, _ := json.Marshal(map[string]any{
		"message":           code,
		"messageTranslated": bindErrorMessage(code, lang, args...),
		"messageParams": map[string]any{
			"offset": offset, "line": line, "col": col,
		},
	})
	return b, 400
}

// BindBodyReadError reports a failure reading the raw request body, before
// any format-specific unmarshalling was attempted.
type BindBodyReadError struct {
	// Kind is one of "empty", "unexpectedEof", "readAfterClose", "unknown".
	Kind string
	Err  error
}

func (e *BindBodyReadError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("cannot read request body (%s): %s", e.Kind, e.Err.Error())
	}
	return fmt.Sprintf("cannot read request body (%s)", e.Kind)
}

func (e *BindBodyReadError) Unwrap() error { return e.Err }

func (e *BindBodyReadError) ToPublicJSON(lang string) ([]byte, int32) {
	switch e.Kind {
	case "empty":
		return bindErrorPublicJSON("BodyIsEmptyEof", lang)
	case "unexpectedEof":
		return bindErrorPublicJSON("BodyUnexpectedEof", lang)
	case "readAfterClose":
		return bindErrorPublicJSON("BodyReadAfterClose", lang)
	default:
		return bindErrorPublicJSON("UnknownErrorReadingBody", lang)
	}
}

// BindFieldTypeError reports a JSON/XML/etc value that doesn't match the
// target struct field's type (e.g. a string sent for an int field).
type BindFieldTypeError struct {
	Field    string
	Expected string
	Actual   interface{}
	Offset   int64
	Line     int
	Col      int
}

func (e *BindFieldTypeError) Error() string {
	return fmt.Sprintf("field %q: expected type %s but got %v (offset %d, line %d, col %d)",
		e.Field, e.Expected, e.Actual, e.Offset, e.Line, e.Col)
}

func (e *BindFieldTypeError) ToPublicJSON(lang string) ([]byte, int32) {
	return bindErrorPublicJSONAt("JsonInvalidFieldType", lang, e.Offset, e.Line, e.Col,
		e.Field, e.Expected, e.Actual, e.Line, e.Col)
}

// BindSyntaxError reports malformed JSON/XML (broken braces, commas, tags...).
// Format is one of "json", "xml"; it only selects the ToPublicJSON message.
type BindSyntaxError struct {
	Format string
	Offset int64
	Line   int
	Col    int
}

func (e *BindSyntaxError) Error() string {
	return fmt.Sprintf("malformed body at offset %d (line %d, col %d)", e.Offset, e.Line, e.Col)
}

func (e *BindSyntaxError) ToPublicJSON(lang string) ([]byte, int32) {
	if e.Format == "xml" {
		return bindErrorPublicJSONAt("XmlMalformed", lang, e.Offset, e.Line, e.Col, e.Line, e.Col)
	}
	return bindErrorPublicJSONAt("JsonMalformed", lang, e.Offset, e.Line, e.Col, e.Line, e.Col)
}

// BindXmlUnmarshalError reports an XML document that parses fine but whose
// structure doesn't match the target Go struct (missing/misplaced elements).
type BindXmlUnmarshalError struct {
	Err error
}

func (e *BindXmlUnmarshalError) Error() string {
	if e.Err != nil {
		return "xml unmarshal error: " + e.Err.Error()
	}
	return "xml unmarshal error"
}

func (e *BindXmlUnmarshalError) Unwrap() error { return e.Err }

func (e *BindXmlUnmarshalError) ToPublicJSON(lang string) ([]byte, int32) {
	return bindErrorPublicJSON("XmlUnmarshalError", lang)
}

// getLineAndCharFromOffset walks body up to offset, converting a byte offset
// (as reported by json/xml syntax errors) into a 1-indexed line and column.
func getLineAndCharFromOffset(body []byte, offset int64) (line int, col int) {
	line = 1
	col = 1
	for i := int64(0); i < offset && i < int64(len(body)); i++ {
		if body[i] == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return
}

// BindUnsupportedTypeError reports a target field type json/xml cannot decode into.
type BindUnsupportedTypeError struct {
	Type string
}

func (e *BindUnsupportedTypeError) Error() string {
	return fmt.Sprintf("unsupported type %s", e.Type)
}

func (e *BindUnsupportedTypeError) ToPublicJSON(lang string) ([]byte, int32) {
	return bindErrorPublicJSON("JsonUnmarshalUnsupportedType", lang)
}

// BindYamlTypeError reports one or more per-value YAML type mismatches.
type BindYamlTypeError struct {
	Errors []string
}

func (e *BindYamlTypeError) Error() string {
	return fmt.Sprintf("%d yaml type error(s): %v", len(e.Errors), e.Errors)
}

func (e *BindYamlTypeError) ToPublicJSON(lang string) ([]byte, int32) {
	return bindErrorPublicJSON("YamlTypeError", lang)
}

// BindDecodingError is the catch-all for a format-specific unmarshal failure
// that doesn't fit one of the more specific types above. Format is one of
// "json", "yaml", "xml", "form".
type BindDecodingError struct {
	Format string
	Err    error
}

func (e *BindDecodingError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s decoding error: %s", e.Format, e.Err.Error())
	}
	return fmt.Sprintf("%s decoding error", e.Format)
}

func (e *BindDecodingError) Unwrap() error { return e.Err }

func (e *BindDecodingError) ToPublicJSON(lang string) ([]byte, int32) {
	switch e.Format {
	case "yaml":
		return bindErrorPublicJSON("YamlDecodingError", lang)
	case "xml":
		return bindErrorPublicJSON("XmlDecodingError", lang)
	case "form":
		return bindErrorPublicJSON("FormDataMalformed", lang)
	default:
		return bindErrorPublicJSON("JsonDecodingError", lang)
	}
}
