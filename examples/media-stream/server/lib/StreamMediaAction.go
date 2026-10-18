package mediagen

import (
	"fmt"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

/**
* Action to communicate with the action StreamMediaAction
 */
func StreamMediaActionMeta() struct {
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
		Name:        "StreamMediaAction",
		URL:         "/media/stream/:deviceId",
		Method:      "REACTIVE",
		CliName:     "stream-media-action",
		Description: "Device -> server raw audio (PCM/Opus) and video (MJPEG/H264) frames.",
	}
}

/**
 * Path parameters for StreamMediaAction
 */
type StreamMediaActionPathParameter struct {
	DeviceId string
}

// Converts a placeholder url, and applies the parameters to it.
func StreamMediaActionPathParameterApply(params StreamMediaActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":deviceId", fmt.Sprintf("%v", params.DeviceId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func StreamMediaActionPathParameterFromFn(fn func(key string) string) StreamMediaActionPathParameter {
	res := StreamMediaActionPathParameter{}
	res.DeviceId = fn("deviceId")
	return res
}

/**
 * Query parameters for StreamMediaAction
 */
// Query wrapper with private fields
type StreamMediaActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	SampleRate int    `json:"sampleRate"`
	Channels   int    `json:"channels"`
	AudioCodec string `json:"audioCodec"`
	VideoCodec string `json:"videoCodec"`
}

func (q StreamMediaActionQuery) GetSampleRate() int    { return q.SampleRate }
func (q StreamMediaActionQuery) GetChannels() int      { return q.Channels }
func (q StreamMediaActionQuery) GetAudioCodec() string { return q.AudioCodec }
func (q StreamMediaActionQuery) GetVideoCodec() string { return q.VideoCodec }
func StreamMediaActionQueryFromString(rawQuery string) StreamMediaActionQuery {
	v := StreamMediaActionQuery{}
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
func StreamMediaActionQueryFromHttp(r *http.Request) StreamMediaActionQuery {
	return StreamMediaActionQueryFromString(r.URL.RawQuery)
}
func (q StreamMediaActionQuery) Values() url.Values {
	return q.values
}
func (q StreamMediaActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *StreamMediaActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *StreamMediaActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type StreamMediaActionMessage struct {
	Raw []byte
	// Conn *websocket.Conn
	Conn        interface{}
	MessageType int
	Error       error
	PathParams  StreamMediaActionPathParameter
}

// Developer handler type
type StreamMediaActionHandler func(msg StreamMediaActionMessage) error
type StreamMediaActionSession struct {
	// Ctx    *gin.Context
	// Socket *websocket.Conn
	Ctx         interface{}
	Socket      interface{}
	Done        chan bool
	Read        chan StreamMediaActionReadChan
	QueryParams StreamMediaActionQuery
}
type StreamMediaActionHandlerDuplex func(*StreamMediaActionSession)
type StreamMediaActionReadChan struct {
	Data        []byte
	Error       error
	MessageType int
}

// StreamMediaActionClientSession is the client-side mirror of
// StreamMediaActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type StreamMediaActionClientSession struct {
	// Socket *websocket.Conn
	Socket interface{}
	Done   chan bool
	Read   chan StreamMediaActionReadChan
	Write  chan []byte
}
