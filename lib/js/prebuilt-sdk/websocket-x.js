export class WebSocketX extends WebSocket {
    constructor(url, protocols) {
        super(url, protocols);
        this.sendRaw = super.send.bind(this);
        this.addEventListenerRaw = super.addEventListener.bind(this);
    }
    // @ts-expect-error override to customize send
    send(data) {
        if (typeof data === "string" ||
            data instanceof Blob ||
            data instanceof ArrayBuffer ||
            ArrayBuffer.isView(data)) {
            super.send(data);
        }
        else {
            super.send(JSON.stringify(data));
        }
    }
    // implementation
    addEventListener(type, listener, options) {
        if (type === "message") {
            const wrapped = ((ev) => {
                let parsed;
                try {
                    parsed = JSON.parse(ev.data);
                }
                catch {
                    parsed = ev.data;
                }
                listener.call(this, new MessageEvent("message", { data: parsed }));
            });
            super.addEventListener(type, wrapped, options);
        }
        else {
            super.addEventListener(type, listener, options);
        }
    }
}
