package external

import (
	"encoding"
	"testing"
)

// Covers the "complex" scenario: a hand-written type (Money, in support.go) plugged in
// via complexes: with no location/namespace, so it's expected to already live in this
// same generated package.

func TestComplexField_TypeAndTextRoundtrip(t *testing.T) {
	entity := Entity1Entity{Complex1: Money{Cents: 1250}}

	text, err := entity.Complex1.MarshalText()
	if err != nil {
		t.Fatalf("MarshalText error: %v", err)
	}
	if string(text) != "1250" {
		t.Fatalf("MarshalText() = %q, want %q", text, "1250")
	}

	var roundtripped Money
	if err := roundtripped.UnmarshalText(text); err != nil {
		t.Fatalf("UnmarshalText error: %v", err)
	}
	if roundtripped.Cents != 1250 {
		t.Fatalf("roundtripped Cents = %d, want 1250", roundtripped.Cents)
	}
}

func TestComplexField_CliCastUsesTextUnmarshaler(t *testing.T) {
	data := Entity1Entity{}

	if _, ok := any(&data.Complex1).(encoding.TextUnmarshaler); !ok {
		t.Fatal("expected *Money to implement encoding.TextUnmarshaler for CLI casting")
	}

	u := any(&data.Complex1).(encoding.TextUnmarshaler)
	if err := u.UnmarshalText([]byte("999")); err != nil {
		t.Fatalf("UnmarshalText error: %v", err)
	}
	if data.Complex1.Cents != 999 {
		t.Fatalf("Complex1.Cents = %d, want 999", data.Complex1.Cents)
	}
}
