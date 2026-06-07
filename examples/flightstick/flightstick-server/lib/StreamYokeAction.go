package external

import (
	"fmt"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

/**
* Action to communicate with the action StreamYokeAction
 */
func StreamYokeActionMeta() struct {
	Name        string
	URL         string
	Method      string
	CliName     string
	Description string
} {
	return struct {
		Name        string
		URL         string
		Method      string
		CliName     string
		Description string
	}{
		Name:        "StreamYokeAction",
		URL:         "/stream-yoke/:path222",
		Method:      "REACTIVE",
		CliName:     "",
		Description: "Used by the kotlin to golang, to send the yoke information",
	}
}

/**
 * Path parameters for StreamYokeAction
 */
type StreamYokeActionPathParameter struct {
	Path222 string
}

// Converts a placeholder url, and applies the parameters to it.
func StreamYokeActionPathParameterApply(params StreamYokeActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":path222", fmt.Sprintf("%v", params.Path222))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func StreamYokeActionPathParameterFromFn(fn func(key string) string) StreamYokeActionPathParameter {
	res := StreamYokeActionPathParameter{}
	res.Path222 = fn("path222")
	return res
}

/**
 * Query parameters for StreamYokeAction
 */
// Query wrapper with private fields
type StreamYokeActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func StreamYokeActionQueryFromString(rawQuery string) StreamYokeActionQuery {
	v := StreamYokeActionQuery{}
	values, _ := url.ParseQuery(rawQuery)
	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}
	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json", // reuse json tags
		WeaklyTypedInput: true,   // "1" -> int, "true" -> bool
		Result:           &v,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}
	v.values = values
	v.mapped = mapped
	return v
}
func StreamYokeActionQueryFromHttp(r *http.Request) StreamYokeActionQuery {
	return StreamYokeActionQueryFromString(r.URL.RawQuery)
}
func (q StreamYokeActionQuery) Values() url.Values {
	return q.values
}
func (q StreamYokeActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *StreamYokeActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *StreamYokeActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type StreamYokeActionMessage struct {
	Raw []byte
	// Conn *websocket.Conn
	Conn        interface{}
	MessageType int
	Error       error
	PathParams  StreamYokeActionPathParameter
}

// Developer handler type
type StreamYokeActionHandler func(msg StreamYokeActionMessage) error
type StreamYokeActionSession struct {
	// Ctx    *gin.Context
	// Socket *websocket.Conn
	Ctx         interface{}
	Socket      interface{}
	Done        chan bool
	Read        chan StreamYokeActionReadChan
	QueryParams StreamYokeActionQuery
}
type StreamYokeActionHandlerDuplex func(*StreamYokeActionSession)
type StreamYokeActionReadChan struct {
	Data        []byte
	Error       error
	MessageType int
}

// StreamYokeActionClientSession is the client-side mirror of
// StreamYokeActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type StreamYokeActionClientSession struct {
	// Socket *websocket.Conn
	Socket interface{}
	Done   chan bool
	Read   chan StreamYokeActionReadChan
	Write  chan []byte
}
