var __classPrivateFieldSet =
  (this && this.__classPrivateFieldSet) ||
  function (receiver, state, value, kind, f) {
    if (kind === "m") throw new TypeError("Private method is not writable");
    if (kind === "a" && !f)
      throw new TypeError("Private accessor was defined without a setter");
    if (
      typeof state === "function"
        ? receiver !== state || !f
        : !state.has(receiver)
    )
      throw new TypeError(
        "Cannot write private member to an object whose class did not declare it"
      );
    return (
      kind === "a"
        ? f.call(receiver, value)
        : f
        ? (f.value = value)
        : state.set(receiver, value),
      value
    );
  };
var __classPrivateFieldGet =
  (this && this.__classPrivateFieldGet) ||
  function (receiver, state, kind, f) {
    if (kind === "a" && !f)
      throw new TypeError("Private accessor was defined without a getter");
    if (
      typeof state === "function"
        ? receiver !== state || !f
        : !state.has(receiver)
    )
      throw new TypeError(
        "Cannot read private member from an object whose class did not declare it"
      );
    return kind === "m"
      ? f
      : kind === "a"
      ? f.call(receiver)
      : f
      ? f.value
      : state.get(receiver);
  };
var _WebSocketX_factoryCls;
export class WebSocketX extends WebSocket {
  constructor(url, protocols, options) {
    super(url, protocols);
    _WebSocketX_factoryCls.set(this, void 0);
    this.sendRaw = super.send.bind(this);
    this.addEventListenerRaw = super.addEventListener.bind(this);
    if (
      options === null || options === void 0
        ? void 0
        : options.MessageFactoryClass
    ) {
      __classPrivateFieldSet(
        this,
        _WebSocketX_factoryCls,
        options.MessageFactoryClass,
        "f"
      );
    }
  }
  set onmessage(fn) {
    if (fn) {
      this.addEventListener("message", fn);
    } else {
      super.onmessage = null;
    }
  }
  get onmessage() {
    return super.onmessage;
  }
  // @ts-expect-error override to customize send
  send(data) {
    if (
      typeof data === "string" ||
      data instanceof Blob ||
      data instanceof ArrayBuffer ||
      ArrayBuffer.isView(data)
    ) {
      super.send(data);
    } else if (data !== undefined && data !== null) {
      super.send(data.toString());
    }
  }
  // implementation
  addEventListener(type, listener, options) {
    if (type === "message") {
      const wrapped = (ev) => {
        let parsed = ev.data;
        if (__classPrivateFieldGet(this, _WebSocketX_factoryCls, "f")) {
          try {
            parsed = new (__classPrivateFieldGet(
              this,
              _WebSocketX_factoryCls,
              "f"
            ))(ev.data);
          } catch (_a) {
            // if constructor rejects (e.g. ArrayBuffer not supported), keep raw
          }
        }
        listener.call(this, new MessageEvent("message", { data: parsed }));
      };
      super.addEventListener(type, wrapped, options);
    } else {
      super.addEventListener(type, listener, options);
    }
  }
}
_WebSocketX_factoryCls = new WeakMap();
