//go:build wasm

package emigo

import (
	"net/url"
	"sync"
	"sync/atomic"
	"syscall/js"
)

// WasmReactiveConn is the server-side view of one logical client connection.
// It is the websocket-conn analogue for the browser: there is no socket, just
// two channels bridged to JS callbacks.
//
//	Read  carries client -> server frames (fed by wasmWsSend)
//	Write carries server -> client frames (drained into the onMessage callback)
//	Done  fires once when either side tears the connection down
//
// This mirrors the channel contract of the gorilla-backed reactive session, so
// the same developer factory can drive either transport.
type WasmReactiveConn struct {
	Read  chan []byte
	Write chan []byte
	Done  chan bool
	Query url.Values
}

// WasmReactiveHandler sets up one connection. It runs once per WebSocketWasm the
// browser opens. Wire your read/write loops against conn here; return an error
// to reject the connection (the error text is delivered as one final message).
type WasmReactiveHandler func(conn *WasmReactiveConn) error

// WasmReactor is the reactive counterpart of LiftWasmServer. Where LiftWasmServer
// turns a single JS call into one HTTP request/response, the reactor manages
// long-lived connections and exposes three JS entry points the WebSocketWasm
// class drives:
//
//	wasmWsOpen(path, query, onMessage, onClose) -> connId | -1
//	wasmWsSend(connId, data)                    -> bool
//	wasmWsClose(connId)                         -> bool
//
// The reactive path deliberately bypasses net/http: gorilla's Upgrade needs to
// hijack a real net.Conn, which the in-browser httptest server has no notion of.
// Talking channels straight to JS sidesteps the handshake and framing entirely.
type WasmReactor struct {
	mu       sync.Mutex
	handlers map[string]WasmReactiveHandler
	conns    map[int]*WasmReactiveConn
	nextID   int64
}

func NewWasmReactor() *WasmReactor {
	return &WasmReactor{
		handlers: map[string]WasmReactiveHandler{},
		conns:    map[int]*WasmReactiveConn{},
	}
}

// Handle registers a reactive handler under a path. The browser selects it via
// the pathname of the URL passed to new WebSocketWasm(url).
func (r *WasmReactor) Handle(path string, h WasmReactiveHandler) {
	r.mu.Lock()
	r.handlers[path] = h
	r.mu.Unlock()
}

// Lift exposes the JS entry points. Call it once, after registering handlers.
func (r *WasmReactor) Lift() {
	js.Global().Set("wasmWsOpen", js.FuncOf(r.jsOpen))
	js.Global().Set("wasmWsSend", js.FuncOf(r.jsSend))
	js.Global().Set("wasmWsClose", js.FuncOf(r.jsClose))
}

func (r *WasmReactor) drop(id int) {
	r.mu.Lock()
	delete(r.conns, id)
	r.mu.Unlock()
}

// jsOpen: wasmWsOpen(path, query, onMessage, onClose) -> connId | -1
func (r *WasmReactor) jsOpen(_ js.Value, a []js.Value) any {
	path := a[0].String()
	rawQuery := a[1].String()
	onMessage := a[2] // function(dataString)
	onClose := a[3]   // function(reasonString)

	r.mu.Lock()
	h, ok := r.handlers[path]
	r.mu.Unlock()
	if !ok {
		return -1
	}

	q, _ := url.ParseQuery(rawQuery)
	conn := &WasmReactiveConn{
		Read:  make(chan []byte, 16),
		Write: make(chan []byte, 16),
		Done:  make(chan bool, 1),
		Query: q,
	}
	id := int(atomic.AddInt64(&r.nextID, 1))
	r.mu.Lock()
	r.conns[id] = conn
	r.mu.Unlock()

	// server -> client pump. Exactly one onClose fires: whichever of Write-close
	// or Done lands first wins, then the goroutine returns.
	go func() {
		for {
			select {
			case msg, ok := <-conn.Write:
				if !ok {
					onClose.Invoke("server closed")
					r.drop(id)
					return
				}
				onMessage.Invoke(string(msg))
			case <-conn.Done:
				onClose.Invoke("closed")
				r.drop(id)
				return
			}
		}
	}()

	// Run the developer handler. On setup error, deliver the text and tear down.
	go func() {
		if err := h(conn); err != nil {
			onMessage.Invoke("error: " + err.Error())
			select {
			case conn.Done <- true:
			default:
			}
		}
	}()

	return id
}

// jsSend: wasmWsSend(connId, data) -> bool
func (r *WasmReactor) jsSend(_ js.Value, a []js.Value) any {
	id := a[0].Int()
	data := a[1].String()

	r.mu.Lock()
	conn := r.conns[id]
	r.mu.Unlock()
	if conn == nil {
		return false
	}
	// Feed the frame on a goroutine: a synchronous JS->Go call must not block
	// the event loop waiting on a busy handler.
	go func() {
		select {
		case conn.Read <- []byte(data):
		case <-conn.Done:
		}
	}()
	return true
}

// jsClose: wasmWsClose(connId) -> bool
func (r *WasmReactor) jsClose(_ js.Value, a []js.Value) any {
	id := a[0].Int()

	r.mu.Lock()
	conn := r.conns[id]
	r.mu.Unlock()
	if conn == nil {
		return false
	}
	select {
	case conn.Done <- true:
	default:
	}
	return true
}
