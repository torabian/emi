package external

import (
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action SubscribeYokeAction
 */
func SubscribeYokeActionMeta() struct {
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
		Name:        "SubscribeYokeAction",
		URL:         "/subscribe-yoke",
		Method:      "REACTIVE",
		CliName:     "",
		Description: "Used by the unreal engine, to fetch the yoke information and moving the elements.",
	}
}

/**
 * Query parameters for SubscribeYokeAction
 */
// Query wrapper with private fields
type SubscribeYokeActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func SubscribeYokeActionQueryFromString(rawQuery string) SubscribeYokeActionQuery {
	v := SubscribeYokeActionQuery{}
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
func SubscribeYokeActionQueryFromHttp(r *http.Request) SubscribeYokeActionQuery {
	return SubscribeYokeActionQueryFromString(r.URL.RawQuery)
}
func (q SubscribeYokeActionQuery) Values() url.Values {
	return q.values
}
func (q SubscribeYokeActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *SubscribeYokeActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *SubscribeYokeActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type SubscribeYokeActionMessage struct {
	Raw []byte
	// Conn *websocket.Conn
	Conn        interface{}
	MessageType int
	Error       error
}

// Developer handler type
type SubscribeYokeActionHandler func(msg SubscribeYokeActionMessage) error
type SubscribeYokeActionSession struct {
	// Ctx    *gin.Context
	// Socket *websocket.Conn
	Ctx         interface{}
	Socket      interface{}
	Done        chan bool
	Read        chan SubscribeYokeActionReadChan
	QueryParams SubscribeYokeActionQuery
}
type SubscribeYokeActionHandlerDuplex func(*SubscribeYokeActionSession)
type SubscribeYokeActionReadChan struct {
	Data        []byte
	Error       error
	MessageType int
}

// SubscribeYokeActionClientSession is the client-side mirror of
// SubscribeYokeActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type SubscribeYokeActionClientSession struct {
	// Socket *websocket.Conn
	Socket interface{}
	Done   chan bool
	Read   chan SubscribeYokeActionReadChan
	Write  chan []byte
}
