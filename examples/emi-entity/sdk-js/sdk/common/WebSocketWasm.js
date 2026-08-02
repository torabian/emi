// WebSocketWasm — a drop-in replacement for the native WebSocket, backed by the
// Go-WASM reactor (emigo.WasmReactor). Swap `new WebSocket(url)` for
// `new WebSocketWasm(url)` and the rest of your code is identical: readyState,
// onopen/onmessage/onclose/onerror, addEventListener, send(), close().
//
// It is NOT a network client — there is no handshake and no socket. The
// constructor's URL is only used for its pathname (to pick the reactor handler)
// and query string; the host is ignored.
export class WebSocketWasm extends EventTarget {
    constructor(url, _protocols) {
        super();
        this.url = url;
        this.readyState = WebSocketWasm.CONNECTING;
        this.onopen = null;
        this.onmessage = null;
        this.onerror = null;
        this.onclose = null;
        this._id = -1;
        const u = new URL(url, "ws://wasm.local"); // base lets bare paths parse
        const path = u.pathname;
        const query = u.search.replace(/^\?/, "");
        // Defer so listeners attached right after `new` still catch 'open', exactly
        // like the native async connect.
        queueMicrotask(() => this._open(path, query));
    }
    _open(path, query) {
        var _a;
        const onMessage = (data) => {
            var _a;
            const ev = new MessageEvent("message", { data });
            (_a = this.onmessage) === null || _a === void 0 ? void 0 : _a.call(this, ev);
            this.dispatchEvent(ev);
        };
        const onClose = (reason) => {
            var _a;
            if (this.readyState === WebSocketWasm.CLOSED)
                return;
            this.readyState = WebSocketWasm.CLOSED;
            const ev = new CloseEvent("close", {
                reason,
                code: 1000,
                wasClean: true,
            });
            (_a = this.onclose) === null || _a === void 0 ? void 0 : _a.call(this, ev);
            this.dispatchEvent(ev);
        };
        if (typeof window.wasmWsOpen !== "function") {
            return this._fail("reactor not ready (window.wasmWsOpen missing)");
        }
        this._id = window.wasmWsOpen(path, query, onMessage, onClose);
        if (this._id < 0) {
            return this._fail(`no reactive handler registered for "${path}"`);
        }
        this.readyState = WebSocketWasm.OPEN;
        const ev = new Event("open");
        (_a = this.onopen) === null || _a === void 0 ? void 0 : _a.call(this, ev);
        this.dispatchEvent(ev);
    }
    _fail(message) {
        var _a, _b;
        this.readyState = WebSocketWasm.CLOSED;
        const err = new Event("error");
        err.message = message;
        (_a = this.onerror) === null || _a === void 0 ? void 0 : _a.call(this, err);
        this.dispatchEvent(err);
        const close = new CloseEvent("close", {
            reason: message,
            code: 1006,
            wasClean: false,
        });
        (_b = this.onclose) === null || _b === void 0 ? void 0 : _b.call(this, close);
        this.dispatchEvent(close);
    }
    send(data) {
        if (this.readyState !== WebSocketWasm.OPEN) {
            throw new DOMException("WebSocketWasm is not open", "InvalidStateError");
        }
        // Strings pass through; everything else is stringified for now. Real binary
        // (ArrayBuffer/Blob) would carry a type flag across the bridge — future work.
        window.wasmWsSend(this._id, typeof data === "string" ? data : String(data));
    }
    close(_code, _reason) {
        if (this.readyState === WebSocketWasm.CLOSING ||
            this.readyState === WebSocketWasm.CLOSED) {
            return;
        }
        this.readyState = WebSocketWasm.CLOSING;
        if (this._id >= 0)
            window.wasmWsClose(this._id);
    }
}
WebSocketWasm.CONNECTING = 0;
WebSocketWasm.OPEN = 1;
WebSocketWasm.CLOSING = 2;
WebSocketWasm.CLOSED = 3;
