export class WebSocketX<
  SendType = string | ArrayBufferLike | Blob | ArrayBufferView,
  RecieveData = string
> extends WebSocket {
  public readonly addEventListenerRaw: WebSocket["addEventListener"];
  public readonly sendRaw: WebSocket["send"];

  constructor(url: string | URL, protocols?: string | string[]) {
    super(url, protocols);

    this.sendRaw = super.send.bind(this);
    this.addEventListenerRaw = super.addEventListener.bind(this);
  }

  // @ts-expect-error override to customize send
  send(data: SendType): void {
    if (
      typeof data === "string" ||
      data instanceof Blob ||
      data instanceof ArrayBuffer ||
      ArrayBuffer.isView(data)
    ) {
      super.send(data);
    } else {
      super.send(JSON.stringify(data));
    }
  }

  addEventListener(
    type: "message",
    listener: (this: WebSocket, ev: MessageEvent<RecieveData>) => unknown,
    options?: boolean | AddEventListenerOptions
  ): void;

  // fallback overloads (other event types)
  addEventListener<K extends Exclude<keyof WebSocketEventMap, "message">>(
    type: K,
    listener: (this: WebSocket, ev: WebSocketEventMap[K]) => unknown,
    options?: boolean | AddEventListenerOptions
  ): void;

  // implementation
  addEventListener(
    type: string,
    listener: EventListenerOrEventListenerObject,
    options?: boolean | AddEventListenerOptions
  ): void {
    if (type === "message") {
      const wrapped = ((ev: MessageEvent) => {
        let parsed: unknown;
        try {
          parsed = JSON.parse(ev.data);
        } catch {
          parsed = ev.data;
        }
        (listener as EventListener).call(
          this,
          new MessageEvent("message", { data: parsed })
        );
      }) as EventListener;

      super.addEventListener(type, wrapped, options);
    } else {
      super.addEventListener(type, listener, options);
    }
  }
}
