package emigo

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

/**
* Response-body content negotiation shared by GinBinding.go and
* HttpBinding.go: picking a format off the request's Accept header, and
* marshaling a response payload (or the {"error": ...} envelope) into it.
* Lives in its own file because both binding files need the exact same
* logic and, being in the same package, can't each declare it.
**/

// responseFormat is the wire format chosen for a response body, resolved
// from the request's Accept header by detectAcceptFormat.
type responseFormat string

const (
	responseFormatJSON responseFormat = "json"
	responseFormatYAML responseFormat = "yaml"
	responseFormatTOML responseFormat = "toml"
	responseFormatCSV  responseFormat = "csv"
	responseFormatXML  responseFormat = "xml"
)

// contentType is the Content-Type header value that pairs with f.
func (f responseFormat) contentType() string {
	switch f {
	case responseFormatYAML:
		return "application/yaml"
	case responseFormatTOML:
		return "application/toml"
	case responseFormatCSV:
		return "text/csv"
	case responseFormatXML:
		return "application/xml"
	default:
		return "application/json"
	}
}

// detectAcceptFormat picks a responseFormat off an Accept header (as sent by
// the client): its comma-separated media types are checked in the order the
// client listed them (ignoring "q=..." and any other parameter), and the
// first one this package recognizes wins. Alongside the real MIME types,
// each format also accepts its bare name ("yaml"/"yml", "toml", "csv",
// "xml") for clients that don't bother with a proper media type. JSON is
// the default when nothing matches (including an empty/missing header) -
// this is what every caller got before other formats existed, so JSON
// stays the fallback rather than becoming one more type to ask for
// explicitly.
func detectAcceptFormat(accept string) responseFormat {
	for _, part := range strings.Split(accept, ",") {
		mt := strings.TrimSpace(part)
		if i := strings.IndexByte(mt, ';'); i >= 0 {
			mt = strings.TrimSpace(mt[:i])
		}
		switch mt {
		case "application/yaml", "application/x-yaml", "text/yaml", "yaml", "yml":
			return responseFormatYAML
		case "application/toml", "text/toml", "application/x-toml", "toml":
			return responseFormatTOML
		case "text/csv", "application/csv", "csv":
			return responseFormatCSV
		case "application/xml", "text/xml", "xml":
			return responseFormatXML
		case "application/json":
			return responseFormatJSON
		}
	}
	return responseFormatJSON
}

// marshalResponse renders payload as f, returning the response body. JSON
// and YAML accept any Go value payload already satisfies (struct, map,
// slice, scalar, ...) unchanged. TOML and CSV can't represent a bare
// top-level scalar/array the way JSON/YAML can - see marshalToml and
// marshalCSV for how each copes with that.
func marshalResponse(f responseFormat, payload any) ([]byte, error) {
	switch f {
	case responseFormatYAML:
		return yaml.Marshal(payload)
	case responseFormatTOML:
		return marshalToml(payload)
	case responseFormatCSV:
		return marshalCSV(payload)
	case responseFormatXML:
		return marshalXML(payload)
	default:
		return json.Marshal(payload)
	}
}

// normalizeForTable round-trips payload through JSON into generic
// map[string]any / []any / scalar values, regardless of what concrete Go
// type payload actually is (a struct, a typed map, ...). marshalToml and
// marshalCSV both need this: they inspect payload's shape (is it a map? a
// slice of records?) generically rather than via reflection on arbitrary
// struct types.
func normalizeForTable(payload any) (any, error) {
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	var normalized any
	if err := json.Unmarshal(raw, &normalized); err != nil {
		return nil, err
	}
	return normalized, nil
}

// marshalToml renders payload as a TOML document. TOML has no bare
// top-level scalar or array the way JSON/YAML do - a document is always a
// table - so a payload that isn't itself an object (a slice, e.g. a list
// response, or a scalar) is wrapped as {"data": payload} first.
func marshalToml(payload any) ([]byte, error) {
	normalized, err := normalizeForTable(payload)
	if err != nil {
		return nil, err
	}
	// Every JSON number round-trips as float64, and unlike YAML's encoder,
	// go-toml prints a whole-number float64 with its ".0" suffix (1 becomes
	// "1.0") since TOML's int and float are genuinely distinct types. That
	// would misrepresent, say, an original int id as a float, so whole
	// numbers are converted back to int64 first.
	normalized = denumberWholeFloats(normalized)
	if _, ok := normalized.(map[string]any); !ok {
		normalized = map[string]any{"data": normalized}
	}
	return toml.Marshal(normalized)
}

// denumberWholeFloats recursively converts every whole-number float64 in v
// (a value shaped by normalizeForTable: nested map[string]any/[]any/scalar)
// to an int64, leaving genuine fractional floats, strings, bools, and nil
// untouched.
func denumberWholeFloats(v any) any {
	switch t := v.(type) {
	case map[string]any:
		for k, vv := range t {
			t[k] = denumberWholeFloats(vv)
		}
		return t
	case []any:
		for i, vv := range t {
			t[i] = denumberWholeFloats(vv)
		}
		return t
	case float64:
		if !math.IsInf(t, 0) && !math.IsNaN(t) && t == math.Trunc(t) {
			return int64(t)
		}
		return t
	default:
		return t
	}
}

// marshalCSV renders payload as a CSV table, following a "wrap a single
// object as a one-row table" policy:
//   - a JSON array of objects becomes one row per element; the header is
//     the union of every element's keys (a key missing from a given
//     element leaves that row's cell empty);
//   - anything else - a single object (including the {"error": ...}
//     envelope), a scalar, or an array of non-objects - becomes a table
//     with exactly one row.
//
// Columns are sorted alphabetically for a stable, deterministic header:
// payload round-trips through JSON first (see normalizeForTable), which
// already discards Go struct field order, so there's no "original order"
// left to preserve.
func marshalCSV(payload any) ([]byte, error) {
	normalized, err := normalizeForTable(payload)
	if err != nil {
		return nil, err
	}

	var rows []map[string]any
	switch v := normalized.(type) {
	case []any:
		for _, item := range v {
			if m, ok := item.(map[string]any); ok {
				rows = append(rows, m)
			} else {
				rows = append(rows, map[string]any{"value": item})
			}
		}
	case map[string]any:
		rows = append(rows, v)
	default:
		rows = append(rows, map[string]any{"value": v})
	}

	columns := map[string]bool{}
	for _, row := range rows {
		for k := range row {
			columns[k] = true
		}
	}
	header := make([]string, 0, len(columns))
	for k := range columns {
		header = append(header, k)
	}
	sort.Strings(header)

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(header); err != nil {
		return nil, err
	}
	for _, row := range rows {
		record := make([]string, len(header))
		for i, col := range header {
			record[i] = formatScalarText(row[col])
		}
		if err := w.Write(record); err != nil {
			return nil, err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// formatScalarText renders one JSON-decoded value as plain text - a CSV
// cell, or an XML element's character data: scalars print as their plain
// textual form (a whole-number float64 - which is how every JSON number
// round-trips - prints without a trailing ".0" or exponent), and anything
// structured (a nested object/array) falls back to its compact JSON form
// so no data is silently dropped.
func formatScalarText(v any) string {
	switch t := v.(type) {
	case nil:
		return ""
	case string:
		return t
	case bool:
		return strconv.FormatBool(t)
	case float64:
		if !math.IsInf(t, 0) && !math.IsNaN(t) && t == math.Trunc(t) {
			return strconv.FormatInt(int64(t), 10)
		}
		return strconv.FormatFloat(t, 'f', -1, 64)
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return ""
		}
		return string(b)
	}
}

// xmlRootElement is the fixed root tag marshalXML wraps every response in.
// encoding/xml requires exactly one root element and payload/the
// {"error": ...} envelope aren't always themselves a single-key object (a
// list response is a bare array; a plain object may have several
// top-level keys), so there's no payload-derived name to use instead - see
// marshalXML.
const xmlRootElement = "response"

// xmlArrayItemElement names each element of a top-level array payload, the
// one case marshalXML hits an array with no parent key to repeat (see its
// map[string]any case, which repeats the key itself for a nested array).
const xmlArrayItemElement = "item"

// marshalXML renders payload as an XML document. encoding/xml can't
// marshal a bare map[string]any at all (unlike json/yaml, which accept any
// shape), and XML documents need exactly one root element, so payload is
// first normalized through JSON (see normalizeForTable) and then written
// out by hand, following the common JSON<->XML convention: a map's keys
// become child elements (sorted, same as marshalCSV's column order), and
// an array nested under a key repeats that key as the element name once
// per item (e.g. {"tags":["a","b"]} -> <tags>a</tags><tags>b</tags>)
// rather than introducing a wrapper. A top-level array (no key to repeat)
// falls back to xmlArrayItemElement per element. Everything lands under a
// fixed xmlRootElement root - including the {"error": ...} envelope, which
// comes out as <response><error>...</error></response>.
func marshalXML(payload any) ([]byte, error) {
	normalized, err := normalizeForTable(payload)
	if err != nil {
		return nil, err
	}
	// Same rationale as marshalToml: a JSON number is always float64 after
	// normalizeForTable, and printing e.g. "1" as "1.0" would misrepresent
	// an original int as a float.
	normalized = denumberWholeFloats(normalized)

	var buf bytes.Buffer
	buf.WriteString(xml.Header)
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")
	if err := writeXMLElement(enc, xmlRootElement, normalized); err != nil {
		return nil, err
	}
	if err := enc.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// writeXMLElement writes v as the element named name (children recursively,
// per marshalXML's doc comment).
func writeXMLElement(enc *xml.Encoder, name string, v any) error {
	start := xml.StartElement{Name: xml.Name{Local: sanitizeXMLName(name)}}

	switch t := v.(type) {
	case map[string]any:
		if err := enc.EncodeToken(start); err != nil {
			return err
		}
		keys := make([]string, 0, len(t))
		for k := range t {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			child := t[k]
			if arr, ok := child.([]any); ok {
				// Repeat the key itself per element instead of wrapping the
				// array in its own element - see marshalXML's doc comment.
				for _, item := range arr {
					if err := writeXMLElement(enc, k, item); err != nil {
						return err
					}
				}
				continue
			}
			if err := writeXMLElement(enc, k, child); err != nil {
				return err
			}
		}
		return enc.EncodeToken(start.End())
	case []any:
		// Only reached for a top-level array payload - a nested array is
		// handled inline by the map case above, since there its parent key
		// is available to repeat instead of using xmlArrayItemElement.
		if err := enc.EncodeToken(start); err != nil {
			return err
		}
		for _, item := range t {
			if err := writeXMLElement(enc, xmlArrayItemElement, item); err != nil {
				return err
			}
		}
		return enc.EncodeToken(start.End())
	default:
		if err := enc.EncodeToken(start); err != nil {
			return err
		}
		if t != nil {
			if err := enc.EncodeToken(xml.CharData([]byte(formatScalarText(t)))); err != nil {
				return err
			}
		}
		return enc.EncodeToken(start.End())
	}
}

// sanitizeXMLName makes name safe to use as an XML element's local name:
// leading/trailing whitespace is trimmed, and any character XML forbids in
// a Name (anything but a letter, digit, '_', '-', or '.', and never a
// leading digit or '-'/'.') becomes '_'. Real API field names (JSON keys,
// Go struct field names) already satisfy this, so this only guards against
// the unusual key that wouldn't.
func sanitizeXMLName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "_"
	}
	var b strings.Builder
	for i, r := range name {
		valid := unicode.IsLetter(r) || r == '_' ||
			(i > 0 && (unicode.IsDigit(r) || r == '-' || r == '.'))
		if valid {
			b.WriteRune(r)
		} else {
			b.WriteRune('_')
		}
	}
	return b.String()
}
