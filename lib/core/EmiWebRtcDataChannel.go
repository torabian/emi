package core

type EmiWebRtcDataChannel struct {
	// Name of the data channel in the webrtc.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the data channel in the webrtc"`

	// Channel data which will be sent to.
	In *EmiActionBody `yaml:"in,omitempty" json:"in,omitempty" jsonschema:"description=Channel data which will be sent to"`
}
