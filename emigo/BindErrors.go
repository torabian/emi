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
* renderable {"$": ..., "en": ...} JSON body for free via ToPublicJSON,
* which RenderGinError/RenderHttpError already know how to render.
**/

func bindErrorPublicJSON(code string, en string) ([]byte, int32) {
	b, _ := json.Marshal(map[string]string{"$": code, "en": en})
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
		return bindErrorPublicJSON("BodyIsEmptyEof", "Body is empty. Please provide the necessary data and try again.")
	case "unexpectedEof":
		return bindErrorPublicJSON("BodyUnexpectedEof", "Body unexpected EOF. The data you sent appears incomplete. Please check your request and try again.")
	case "readAfterClose":
		return bindErrorPublicJSON("BodyReadAfterClose", "Body is read after closed. The request might have been processed incorrectly.")
	default:
		return bindErrorPublicJSON("UnknownErrorReadingBody", "We cannot read the body of your request.")
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
	return bindErrorPublicJSON("JsonInvalidFieldType", e.Error())
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
		return bindErrorPublicJSON("XmlMalformed", "The XML format is broken or incomplete. Please make sure all tags are properly opened and closed.")
	}
	return bindErrorPublicJSON("JsonMalformed", "Body is malformed. Check your commas, braces, tags, etc.")
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
	return bindErrorPublicJSON("XmlUnmarshalError", "The XML structure doesn't match the expected format. Some elements may be missing or in the wrong place.")
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
	return bindErrorPublicJSON("JsonUnmarshalUnsupportedType", "Unsupported type when unmarshalling body.")
}

// BindYamlTypeError reports one or more per-value YAML type mismatches.
type BindYamlTypeError struct {
	Errors []string
}

func (e *BindYamlTypeError) Error() string {
	return fmt.Sprintf("%d yaml type error(s): %v", len(e.Errors), e.Errors)
}

func (e *BindYamlTypeError) ToPublicJSON(lang string) ([]byte, int32) {
	return bindErrorPublicJSON("YamlTypeError", "One of the values is in the wrong format. For example, you might have entered text instead of a number or used quotes incorrectly.")
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
		return bindErrorPublicJSON("YamlDecodingError", "There's something wrong with the format of your YAML. Please check indentation, colons, and line breaks to fix the formatting.")
	case "xml":
		return bindErrorPublicJSON("XmlDecodingError", "Something went wrong while processing the XML. Please check the content or try again later.")
	case "form":
		return bindErrorPublicJSON("FormDataMalformed", "The form data submitted is malformed or contains invalid fields. Please check the form and ensure all required fields are properly filled out.")
	default:
		return bindErrorPublicJSON("JsonDecodingError", "Unknown error happened upon decoding.")
	}
}
