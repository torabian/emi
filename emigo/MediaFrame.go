package emigo

import (
	"encoding/binary"
	"errors"
)

// MediaFrame is one microphone / camera frame received over a binary reactive
// action (`method: reactive` + `binaryType: arraybuffer`). It is the Go twin of
// the C++ runtime header EmiMediaFrame.hpp - both sides must agree on this layout:
//
//	offset size field
//	  0     1   Kind        1 audio, 2 video, 3 other
//	  1     1   Flags       bit0 = keyframe
//	  2     2   reserved    always 0
//	  4     4   Seq         per-kind counter, big endian (a gap means a dropped frame)
//	  8     4   TimestampMs milliseconds since the stream started, big endian
//	 12     n   Payload     codec bytes as produced (PCM16LE, Opus, JPEG, H.264 ...)
//
// The codec is not repeated per frame; announce it through the action's typed
// query parameters instead.
type MediaFrame struct {
	Kind        uint8
	Flags       uint8
	Seq         uint32
	TimestampMs uint32
	Payload     []byte
}

const (
	MediaKindAudio uint8 = 1
	MediaKindVideo uint8 = 2
	MediaKindOther uint8 = 3

	MediaFlagKeyframe uint8 = 0x01

	MediaFrameHeaderSize = 12
)

var ErrMediaFrameTooShort = errors.New("media frame shorter than its 12 byte header")

// ParseMediaFrame decodes one websocket binary message. Payload aliases data;
// copy it if you keep it after the read loop reuses the buffer.
func ParseMediaFrame(data []byte) (MediaFrame, error) {
	if len(data) < MediaFrameHeaderSize {
		return MediaFrame{}, ErrMediaFrameTooShort
	}
	return MediaFrame{
		Kind:        data[0],
		Flags:       data[1],
		Seq:         binary.BigEndian.Uint32(data[4:8]),
		TimestampMs: binary.BigEndian.Uint32(data[8:12]),
		Payload:     data[MediaFrameHeaderSize:],
	}, nil
}

func (f MediaFrame) IsKeyframe() bool { return f.Flags&MediaFlagKeyframe != 0 }

// Marshal encodes the frame, ready to be sent on the client session's Write channel.
func (f MediaFrame) Marshal() []byte {
	out := make([]byte, MediaFrameHeaderSize+len(f.Payload))
	out[0] = f.Kind
	out[1] = f.Flags
	binary.BigEndian.PutUint32(out[4:8], f.Seq)
	binary.BigEndian.PutUint32(out[8:12], f.TimestampMs)
	copy(out[MediaFrameHeaderSize:], f.Payload)
	return out
}
