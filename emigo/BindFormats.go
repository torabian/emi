package emigo

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	reflect "reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

/**
* Pure format decoders shared by BindGinRequestBody (GinBinding.go, excluded
* from wasm) and BindHttpRequestBody (HttpBinding.go, wasm-safe). Kept in
* their own file, with no gin/http import, so both transports - and wasm,
* which can't link gin - can use them.
**/

// BindJsonBytes decodes jsonInput into target, returning one of the typed
// errors in BindErrors.go (with offset/line/col detail) on failure.
func BindJsonBytes(jsonInput []byte, target any) error {
	var syntaxErr *json.SyntaxError
	var unmarshalTypeErr *json.UnmarshalTypeError
	var unsupportedTypeErr *json.UnsupportedTypeError

	err := json.Unmarshal(jsonInput, target)
	if err == nil {
		return nil
	}

	switch {
	case errors.As(err, &unmarshalTypeErr):
		field := unmarshalTypeErr.Field
		if field == "" {
			// json only fills Field when the mismatch is nested; for a
			// top-level field it just gives the struct name, so fall back to
			// scanning target's own fields for one whose type matches.
			field = findFieldNameByStructName(target, unmarshalTypeErr.Struct)
		}
		line, col := getLineAndCharFromOffset(jsonInput, unmarshalTypeErr.Offset)
		return &BindFieldTypeError{
			Field:    field,
			Expected: unmarshalTypeErr.Type.String(),
			Actual:   unmarshalTypeErr.Value,
			Offset:   unmarshalTypeErr.Offset,
			Line:     line,
			Col:      col,
		}
	case errors.As(err, &syntaxErr):
		line, col := getLineAndCharFromOffset(jsonInput, syntaxErr.Offset)
		return &BindSyntaxError{Format: "json", Offset: syntaxErr.Offset, Line: line, Col: col}
	case errors.As(err, &unsupportedTypeErr):
		return &BindUnsupportedTypeError{Type: unsupportedTypeErr.Type.String()}
	default:
		return &BindDecodingError{Format: "json", Err: err}
	}
}

// BindYamlBytes decodes yamlInput into target, returning one of the typed
// errors in BindErrors.go on failure.
func BindYamlBytes(yamlInput []byte, target any) error {
	var node yaml.Node
	if err := yaml.Unmarshal(yamlInput, &node); err != nil {
		if typeErr, ok := err.(*yaml.TypeError); ok && len(typeErr.Errors) > 0 {
			return &BindYamlTypeError{Errors: typeErr.Errors}
		}
		return &BindDecodingError{Format: "yaml", Err: err}
	}

	if err := node.Decode(target); err != nil {
		var typeErr *yaml.TypeError
		if errors.As(err, &typeErr) && len(typeErr.Errors) > 0 {
			return &BindYamlTypeError{Errors: typeErr.Errors}
		}
		return &BindDecodingError{Format: "yaml", Err: err}
	}

	return nil
}

// BindXmlBytes decodes xmlInput into target, returning one of the typed
// errors in BindErrors.go on failure.
func BindXmlBytes(xmlInput []byte, target any) error {
	var syntaxErr *xml.SyntaxError
	var unmarshalErr *xml.UnmarshalError

	err := xml.Unmarshal(xmlInput, target)
	if err == nil {
		return nil
	}

	switch {
	case errors.As(err, &syntaxErr):
		return &BindSyntaxError{Format: "xml", Line: syntaxErr.Line}
	case errors.As(err, &unmarshalErr):
		return &BindXmlUnmarshalError{Err: unmarshalErr}
	default:
		return &BindDecodingError{Format: "xml", Err: err}
	}
}

// findFieldNameByStructName scans target's own fields for one whose Go type
// name contains structName - json.UnmarshalTypeError only gives the struct
// name (not the field) when the offending value is a top-level field rather
// than a nested one.
func findFieldNameByStructName(target any, structName string) string {
	t := reflect.TypeOf(target)
	if t == nil {
		return ""
	}
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ""
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if strings.Contains(f.Type.String(), structName) {
			return f.Name
		}
	}
	return ""
}

// bindFormMap round-trips a flattened form value map through JSON into
// target. Inefficient, but it lets a plain form submission populate the same
// struct tags a JSON body would.
func bindFormMap(formMap map[string]any, target any) error {
	formJSON, err := json.Marshal(formMap)
	if err != nil {
		return &BindDecodingError{Format: "form", Err: err}
	}
	if err := json.Unmarshal(formJSON, target); err != nil {
		return &BindDecodingError{Format: "form", Err: err}
	}
	return nil
}
