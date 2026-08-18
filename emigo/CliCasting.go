package emigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/urfave/cli/v3"
)

// Often happens that a dto is being created from cli
// arguments. Here we put some tools to cast them into
// the emi types.

// This is based on urfave, but you can implement
// this for other common frameworks. it should work
// right away with urfave v3
type CliCastable interface {
	String(field string) string
	Int(field string) int
	Int32(field string) int32
	Int64(field string) int64
	Float64(field string) float64
	Bool(field string) bool
	IsSet(field string) bool
}

// Flags are being defined using this, which is again similar to urfave2/3
// Proper casting functions might be required, to convert them into string flag.
// This is just a representation
type CliFlag struct {
	Name        string
	Required    bool
	Usage       string
	Description string
	Type        string

	Children []CliFlag
}

// When on cli it's passed as array, then we need to get it this way.
func CapturePossibleArray[T any](generator func(c CliCastable) T, fieldName string, c CliCastable) Array[T] {
	var result Array[T]
	json.Unmarshal([]byte(c.String(fieldName)), &result.Items)

	return result
}

func CapturePossibleArrayNullable[T any](generator func(c CliCastable) T, fieldName string, c CliCastable) ArrayNullable[T] {
	var result ArrayNullable[T]
	json.Unmarshal([]byte(c.String(fieldName)), &result.Items)

	return result
}

func CapturePossibleCollection[T any](generator func(c CliCastable) T, fieldName string, c CliCastable) Collection[T] {
	var result Collection[T]
	json.Unmarshal([]byte(c.String(fieldName)), &result.Items)

	return result
}

func CapturePossibleCollectionNullable[T any](generator func(c CliCastable) T, fieldName string, c CliCastable) CollectionNullable[T] {
	var result CollectionNullable[T]
	json.Unmarshal([]byte(c.String(fieldName)), &result.Items)

	return result
}

func CapturePossibleOne[T any](generator func(c CliCastable) T, fieldName string, c CliCastable) One[T] {
	var result One[T]
	json.Unmarshal([]byte(c.String(fieldName)), &result.Item)

	return result
}

func CapturePossibleOneNullable[T any](generator func(c CliCastable) T, fieldName string, c CliCastable) OneNullable[T] {
	var result OneNullable[T]
	json.Unmarshal([]byte(c.String(fieldName)), &result.Item)

	return result
}

// CastEmiFlagToUrfave converts a framework-agnostic CliFlag slice (as produced by the
// generated Get*CliFlags functions) into concrete urfave/cli v3 flags. Flags carrying
// Children (nested "object" fields) are flattened - only their leaves ever become an
// actual flag, since urfave has no notion of a nested flag.
func CastEmiFlagToUrfave(flags []CliFlag) []cli.Flag {
	out := make([]cli.Flag, 0, len(flags))

	for _, f := range flags {
		if len(f.Children) > 0 {
			out = append(out, CastEmiFlagToUrfave(f.Children)...)
			continue
		}

		usage := f.Description
		if usage == "" {
			usage = f.Usage
		}

		fieldType := strings.TrimSuffix(f.Type, "?")

		switch fieldType {
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			out = append(out, &cli.Int64Flag{Name: f.Name, Usage: usage, Required: f.Required})
		case "float32", "float64":
			out = append(out, &cli.Float64Flag{Name: f.Name, Usage: usage, Required: f.Required})
		case "bool":
			out = append(out, &cli.BoolFlag{Name: f.Name, Usage: usage, Required: f.Required})
		default:
			// Slice/array/collection/object fields are all captured as a single string
			// (comma-separated, or JSON) by the generated Cast*FromCli functions, via
			// c.String(...) - so they must be registered as a plain StringFlag too.
			out = append(out, &cli.StringFlag{Name: f.Name, Usage: usage, Required: f.Required})
		}
	}

	return out
}

// ParseCliHeaders turns "Key: Value" entries (as collected by a repeatable --header/-H
// flag) into an http.Header, the same shape the Gin/http transports populate
// {Action}Request.Headers with from a live request.
func ParseCliHeaders(raw []string) http.Header {
	headers := http.Header{}
	for _, entry := range raw {
		parts := strings.SplitN(entry, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if key == "" {
			continue
		}
		headers.Add(key, value)
	}
	return headers
}

// CliActionResponse is satisfied by every generated {Action}Response: it exposes
// GetStatusCode/GetRespHeaders/GetPayload with value receivers, so *{Action}Response
// implements this automatically without any extra wiring.
type CliActionResponse interface {
	GetStatusCode() int
	GetRespHeaders() map[string]string
	GetPayload() interface{}
}

// HandleActionInCli is the CLI-side counterpart of what the Gin/http handlers do with an
// action's response: it prints the payload as JSON to stdout and turns a returned error
// (or a >=400 status code) into a non-nil error so urfave reports a non-zero exit code.
func HandleActionInCli(resp CliActionResponse, err error) error {
	if err != nil {
		return err
	}

	if resp == nil {
		return nil
	}

	// resp is a typed nil pointer wrapped in a non-nil interface value; guard against
	// dereferencing it via the interface methods below.
	v := reflect.ValueOf(resp)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil
	}

	if payload := resp.GetPayload(); payload != nil {
		out, encErr := json.MarshalIndent(payload, "", "  ")
		if encErr != nil {
			return encErr
		}
		fmt.Println(string(out))
	}

	if status := resp.GetStatusCode(); status >= 400 {
		return fmt.Errorf("action failed with status %d", status)
	}

	return nil
}

func InflatePossibleSlice[T any](raw string, target *[]T) error {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}

	// JSON array support (still useful)
	if strings.HasPrefix(raw, "[") {
		return json.Unmarshal([]byte(raw), target)
	}

	parts := strings.Split(raw, ",")
	result := make([]T, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)

		v, err := CastPrimitive[T](p)
		if err != nil {
			return err
		}

		result = append(result, v)
	}

	*target = result
	return nil
}
