package external

import (
	"fmt"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

/**
* Action to communicate with the action ReactiveEchoAction
 */
func ReactiveEchoActionMeta() struct {
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
		Name:        "ReactiveEchoAction",
		URL:         "/reactive/echo/:channel",
		Method:      "REACTIVE",
		CliName:     "reactive-echo-action",
		Description: "Echoes messages back over a websocket, scoped to a channel.",
	}
}

/**
 * Path parameters for ReactiveEchoAction
 */
type ReactiveEchoActionPathParameter struct {
	Channel string
}

// Converts a placeholder url, and applies the parameters to it.
func ReactiveEchoActionPathParameterApply(params ReactiveEchoActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":channel", fmt.Sprintf("%v", params.Channel))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func ReactiveEchoActionPathParameterFromFn(fn func(key string) string) ReactiveEchoActionPathParameter {
	res := ReactiveEchoActionPathParameter{}
	res.Channel = fn("channel")
	return res
}

/**
 * Query parameters for ReactiveEchoAction
 */
// Query wrapper with private fields
type ReactiveEchoActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ReactiveEchoActionQueryFromString(rawQuery string) ReactiveEchoActionQuery {
	v := ReactiveEchoActionQuery{}
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
func ReactiveEchoActionQueryFromHttp(r *http.Request) ReactiveEchoActionQuery {
	return ReactiveEchoActionQueryFromString(r.URL.RawQuery)
}
func (q ReactiveEchoActionQuery) Values() url.Values {
	return q.values
}
func (q ReactiveEchoActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ReactiveEchoActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ReactiveEchoActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ReactiveEchoActionMessage struct {
	Raw []byte
	// Conn *websocket.Conn
	Conn        interface{}
	MessageType int
	Error       error
	PathParams  ReactiveEchoActionPathParameter
}

// Developer handler type
type ReactiveEchoActionHandler func(msg ReactiveEchoActionMessage) error
type ReactiveEchoActionSession struct {
	// Ctx    *gin.Context
	// Socket *websocket.Conn
	Ctx         interface{}
	Socket      interface{}
	Done        chan bool
	Read        chan ReactiveEchoActionReadChan
	QueryParams ReactiveEchoActionQuery
}
type ReactiveEchoActionHandlerDuplex func(*ReactiveEchoActionSession)
type ReactiveEchoActionReadChan struct {
	Data        []byte
	Error       error
	MessageType int
}

// ReactiveEchoActionClientSession is the client-side mirror of
// ReactiveEchoActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type ReactiveEchoActionClientSession struct {
	// Socket *websocket.Conn
	Socket interface{}
	Done   chan bool
	Read   chan ReactiveEchoActionReadChan
	Write  chan []byte
}
