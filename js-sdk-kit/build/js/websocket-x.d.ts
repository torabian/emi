export declare class WebSocketX<SendType = string | ArrayBufferLike | Blob | ArrayBufferView, RecieveData = string> extends WebSocket {
    readonly addEventListenerRaw: WebSocket["addEventListener"];
    readonly sendRaw: WebSocket["send"];
    constructor(url: string | URL, protocols?: string | string[]);
    send(data: SendType): void;
    addEventListener(type: "message", listener: (this: WebSocket, ev: MessageEvent<RecieveData>) => unknown, options?: boolean | AddEventListenerOptions): void;
    addEventListener<K extends Exclude<keyof WebSocketEventMap, "message">>(type: K, listener: (this: WebSocket, ev: WebSocketEventMap[K]) => unknown, options?: boolean | AddEventListenerOptions): void;
}
