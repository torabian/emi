package defs

import (
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ChatAction
 */
func ChatActionMeta() struct {
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
		Name:        "ChatAction",
		URL:         "/chat",
		Method:      "REACTIVE",
		CliName:     "",
		Description: "A Reactive chating application, that returns length of strings you've typed.",
	}
}

/**
 * Query parameters for ChatAction
 */
// Query wrapper with private fields
type ChatActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ChatActionQueryFromString(rawQuery string) ChatActionQuery {
	v := ChatActionQuery{}
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
func ChatActionQueryFromHttp(r *http.Request) ChatActionQuery {
	return ChatActionQueryFromString(r.URL.RawQuery)
}
func (q ChatActionQuery) Values() url.Values {
	return q.values
}
func (q ChatActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ChatActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ChatActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ChatActionMessage struct {
	Raw []byte
	// Conn *websocket.Conn
	Conn        interface{}
	MessageType int
	Error       error
}

// Developer handler type
type ChatActionHandler func(msg ChatActionMessage) error
type ChatActionSession struct {
	// Ctx    *gin.Context
	// Socket *websocket.Conn
	Ctx         interface{}
	Socket      interface{}
	Done        chan bool
	Read        chan ChatActionReadChan
	QueryParams ChatActionQuery
}
type ChatActionHandlerDuplex func(*ChatActionSession)
type ChatActionReadChan struct {
	Data        []byte
	Error       error
	MessageType int
}

// ChatActionClientSession is the client-side mirror of
// ChatActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type ChatActionClientSession struct {
	// Socket *websocket.Conn
	Socket interface{}
	Done   chan bool
	Read   chan ChatActionReadChan
	Write  chan []byte
}
