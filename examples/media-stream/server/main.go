// Receives raw microphone + camera frames from a device (ESP32, desktop, ...)
// over the generated StreamMediaAction websocket, and writes them to disk:
//
//	out/<deviceId>.wav      PCM16 audio, header patched when the device disconnects
//	GET /latest/<deviceId>.jpg   newest camera frame
//	GET /app/                    browser demo that streams the visitor's mic + camera
package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/torabian/emi/emigo"

	mediagen "github.com/torabian/emi/examples/media-stream/server/lib"
)

var (
	latestMu sync.RWMutex
	latest   = map[string][]byte{}
)

func main() {
	if err := os.MkdirAll("out", 0o755); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// The whole server side of the media action. The factory runs once per
	// connected device; it returns the channel for server -> device messages.
	mediagen.StreamMediaActionGin(r, func(s mediagen.StreamMediaActionSession) (chan []byte, error) {
		deviceId := mediagen.StreamMediaActionPathParameterFromGin(s.GinCtx()).DeviceId
		q := s.QueryParams // typed: SampleRate, Channels, AudioCodec, VideoCodec

		if q.AudioCodec != "" && q.AudioCodec != "pcm16" {
			return nil, fmt.Errorf("audioCodec %q not supported by this demo (use pcm16)", q.AudioCodec)
		}
		if q.SampleRate == 0 {
			q.SampleRate = 16000
		}
		if q.Channels == 0 {
			q.Channels = 1
		}

		wav, err := newWavWriter(filepath.Join("out", deviceId+".wav"), q.SampleRate, q.Channels)
		if err != nil {
			return nil, err
		}
		log.Printf("[%s] connected: %d Hz, %d ch, video=%q", deviceId, q.SampleRate, q.Channels, q.VideoCodec)

		go func() {
			defer wav.Close()
			var audioFrames, videoFrames, audioLost, videoLost uint32
			var lastAudioSeq, lastVideoSeq uint32
			for msg := range s.Read {
				if msg.Error != nil {
					log.Printf("[%s] disconnected: audio=%d (lost %d) video=%d (lost %d)",
						deviceId, audioFrames, audioLost, videoFrames, videoLost)
					return
				}
				if msg.MessageType != websocket.BinaryMessage {
					continue // text frames are not part of this protocol
				}
				f, err := emigo.ParseMediaFrame(msg.Data)
				if err != nil {
					log.Printf("[%s] bad frame: %v", deviceId, err)
					continue
				}
				switch f.Kind {
				case emigo.MediaKindAudio:
					if audioFrames > 0 {
						audioLost += f.Seq - lastAudioSeq - 1
					}
					lastAudioSeq, audioFrames = f.Seq, audioFrames+1
					wav.Write(f.Payload)
				case emigo.MediaKindVideo:
					if videoFrames > 0 {
						videoLost += f.Seq - lastVideoSeq - 1
					}
					lastVideoSeq, videoFrames = f.Seq, videoFrames+1
					jpg := append([]byte(nil), f.Payload...) // Payload aliases the read buffer
					latestMu.Lock()
					latest[deviceId] = jpg
					latestMu.Unlock()
				}
			}
		}()

		// Nothing to send back in this demo. Keep the channel open; closing
		// it would end the connection.
		return make(chan []byte), nil
	})

	r.GET("/latest/:file", func(c *gin.Context) {
		id := c.Param("file")
		if len(id) > 4 && id[len(id)-4:] == ".jpg" {
			id = id[:len(id)-4]
		}
		latestMu.RLock()
		jpg := latest[id]
		latestMu.RUnlock()
		if jpg == nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "image/jpeg", jpg)
	})

	// The browser demo (web/): built bundle + index.html.
	webDir := os.Getenv("WEB_DIR")
	if webDir == "" {
		webDir = "../web/public"
	}
	r.StaticFS("/app", http.Dir(webDir))

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(r.Run(addr))
}

// wavWriter streams PCM16LE into a .wav and fixes the sizes on Close.
type wavWriter struct {
	f    *os.File
	size uint32
}

func newWavWriter(path string, sampleRate, channels int) (*wavWriter, error) {
	if channels <= 0 || sampleRate <= 0 {
		return nil, errors.New("sampleRate and channels must be positive")
	}
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	h := make([]byte, 44)
	copy(h[0:], "RIFF")
	copy(h[8:], "WAVEfmt ")
	binary.LittleEndian.PutUint32(h[16:], 16)
	binary.LittleEndian.PutUint16(h[20:], 1) // PCM
	binary.LittleEndian.PutUint16(h[22:], uint16(channels))
	binary.LittleEndian.PutUint32(h[24:], uint32(sampleRate))
	binary.LittleEndian.PutUint32(h[28:], uint32(sampleRate*channels*2))
	binary.LittleEndian.PutUint16(h[32:], uint16(channels*2))
	binary.LittleEndian.PutUint16(h[34:], 16)
	copy(h[36:], "data")
	if _, err := f.Write(h); err != nil {
		f.Close()
		return nil, err
	}
	return &wavWriter{f: f}, nil
}

func (w *wavWriter) Write(p []byte) {
	n, _ := w.f.Write(p)
	w.size += uint32(n)
}

func (w *wavWriter) Close() {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], 36+w.size)
	w.f.WriteAt(b[:], 4)
	binary.LittleEndian.PutUint32(b[:], w.size)
	w.f.WriteAt(b[:], 40)
	w.f.Close()
}
