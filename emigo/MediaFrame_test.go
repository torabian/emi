package emigo

import (
	"bytes"
	"testing"
)

func TestMediaFrameRoundTrip(t *testing.T) {
	in := MediaFrame{Kind: MediaKindVideo, Flags: MediaFlagKeyframe, Seq: 0x01020304, TimestampMs: 0xAABBCCDD, Payload: []byte{0xFF, 0xD8, 0xFF}}
	raw := in.Marshal()
	if len(raw) != MediaFrameHeaderSize+3 {
		t.Fatalf("unexpected length %d", len(raw))
	}
	out, err := ParseMediaFrame(raw)
	if err != nil {
		t.Fatal(err)
	}
	if out.Kind != in.Kind || out.Seq != in.Seq || out.TimestampMs != in.TimestampMs || !out.IsKeyframe() || !bytes.Equal(out.Payload, in.Payload) {
		t.Fatalf("round trip mismatch: %+v", out)
	}
}

func TestMediaFrameTooShort(t *testing.T) {
	if _, err := ParseMediaFrame(make([]byte, 11)); err != ErrMediaFrameTooShort {
		t.Fatalf("want ErrMediaFrameTooShort, got %v", err)
	}
	if f, err := ParseMediaFrame(make([]byte, 12)); err != nil || len(f.Payload) != 0 {
		t.Fatalf("header-only frame should parse with empty payload: %v", err)
	}
}
