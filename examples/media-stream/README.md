# media-stream: microphone + camera -> Go server

Devices stream **raw audio and camera frames** to one Go server through a single emi
action. The contract is `media.emi.yml`; every client below is generated from it.

```yaml
- name: streamMedia
  url: /media/stream/:deviceId
  method: reactive          # = a real WebSocket
  binaryType: arraybuffer   # frames are binary, not JSON
  qs: [sampleRate, channels, audioCodec, videoCodec]   # stream parameters, typed on every side
```

| Client | Folder | Runtime emi ships | Verified |
|---|---|---|---|
| Go server (gin) | `server/` | generated handler + `emigo.MediaFrame` | yes |
| **Browser** (mic + camera) | `web/` | `MediaStreamer` + `MediaFrame` in the JS/TS sdk | yes, headless Chrome with fake devices |
| **ESP32, ESP-IDF, C** (I2S mic) | `esp-idf-mic/` | `emiws.h` + `emimedia.h` (no libcurl) | **builds** (ESP-IDF v5.4.4, esp32); not flashed, no hardware |
| Plain C on a host | `device-c/` | same as the ESP project | yes, against the server |
| C++17 (desktop / ESP-IDF / Arduino) | `device/` | `EmiWebSocketX` + `EmiMediaFrame.hpp` | yes, against the server |

```
make gen          # regenerate server/lib, device/emi, device-c/emi, esp-idf-mic/main/emi, web/sdk
make run-server   # Go on :8080 (ADDR=host:port to change)
make web          # bundle the browser demo (npm), then open http://localhost:8080/app/
make device && ./device/media-device 127.0.0.1 8080 3      # C++ fake mic + camera
make device-c && ./device-c/media-device-c 127.0.0.1 8080 3 # C fake mic
```

The server writes `server/out/<deviceId>.wav` (PCM16) and serves the newest camera frame at
`GET /latest/<deviceId>.jpg`; it logs frames received and lost per device.

## The ESP32 microphone project

`esp-idf-mic/` is a complete ESP-IDF project: Wi-Fi station, I2S microphone (INMP441 / SPH0645
style), and a reconnect loop around the generated `stream_media_action()`.

```
cd esp-idf-mic
idf.py set-target esp32      # or esp32s3 ...
idf.py menuconfig            # "Emi mic stream": Wi-Fi, server url, device id, I2S pins, gain
idf.py build flash monitor
```

What was and was not checked: it compiles and links with `-Wall -Werror=all` on ESP-IDF v5.4.4 for
`esp32` (the first build found a real generator bug, `int32_t` printed with `%d`, fixed). It has not
run on a board, so the I2S slot configuration and the gain shift (`EMI_MIC_SHIFT`) are the parts to
tune on your hardware. Not tested on `esp32s3`/`c3` either; the pins in `Kconfig` are only defaults.

## Frame layout (identical in Go, C, C++ and JS)

| bytes | field | |
|---|---|---|
| 0 | kind | 1 audio, 2 video, 3 other |
| 1 | flags | bit0 keyframe |
| 2-3 | reserved | 0 |
| 4-7 | seq | per kind, big endian; a gap = dropped frame |
| 8-11 | timestampMs | since stream start, big endian |
| 12.. | payload | PCM16LE audio, JPEG video |

The codec is announced once through the typed query (`sampleRate`, `audioCodec`, ...), not per frame.

## Browser notes

`MediaStreamer` sends mono PCM16 (it asks the browser for an `AudioContext` at your sample rate and
taps it with an `AudioWorklet`) and JPEG frames from a canvas. It skips a camera frame when the
socket's `bufferedAmount` is high, so a slow link drops frames instead of building latency.
`getUserMedia` needs `https://` or `localhost`. Safari: older versions ignore the requested
`AudioContext` sample rate; check `ctx.sampleRate` if the audio sounds fast or slow.

## Which transport?

| Option | In emi? | Verdict for mic/camera |
|---|---|---|
| **WebSocket, binary frames** (`method: reactive`) | **Yes** (Go, C, C++, JS, Swift) | Best fit. Reliable, ordered, TLS and proxies work. Under bad Wi-Fi it can stall: drop frames on the device. |
| HTTP POST | JSON bodies only; raw `bytes` bodies are **not** supported | OK for occasional snapshots. |
| SSE | Yes, server to client only | Wrong direction for upload. |
| UDP / RTP | No | Lowest latency; you write the socket code. |
| WebRTC | No (emi can model the *signaling* as normal actions) | Right for low-latency video to browsers. |
| RTSP / RTMP / HLS | No | Put a media server (MediaMTX, go2rtc) behind the Go server; emi defines its control API. |

## Limits

- No fragmentation or permessage-deflate; an incoming message is buffered whole (C: capped at 1 MiB).
- No `wss://` out of the box: supply a TLS `EmiStream` (C) / `IEmiByteStream` (C++), e.g. over esp-tls.
- `emi::EmiMediaSender` / `EmiMediaSender` are not thread safe; the C websocket has `lock`/`unlock`
  hooks (the ESP project uses a FreeRTOS mutex) for sending from several tasks.
- The C target's path parameters are not percent-encoded (pre-existing); keep device ids URL-safe.
- Chrome fake devices were used for the browser test; a real camera/microphone and Safari/Firefox were not.
