type ConstructorWithArg<T = any, R = any> = new (arg: T, ...rest: any[]) => R;
export declare class WebSocketX<SendType = string | ArrayBufferLike | Blob | ArrayBufferView, RecieveData = string> extends WebSocket {
    #private;
    readonly addEventListenerRaw: WebSocket["addEventListener"];
    readonly sendRaw: WebSocket["send"];
    constructor(url: string | URL, protocols?: string | string[], options?: {
        MessageFactoryClass: ConstructorWithArg<any>;
    });
    set onmessage(fn: ((this: WebSocket, ev: MessageEvent<RecieveData>) => any) | null);
    get onmessage(): ((this: WebSocket, ev: MessageEvent<RecieveData>) => any) | null;
    send(data: SendType): void;
    addEventListener(type: "message", listener: (this: WebSocket, ev: MessageEvent<RecieveData>) => unknown, options?: boolean | AddEventListenerOptions): void;
    addEventListener<K extends Exclude<keyof WebSocketEventMap, "message">>(type: K, listener: (this: WebSocket, ev: WebSocketEventMap[K]) => unknown, options?: boolean | AddEventListenerOptions): void;
}
export {};
