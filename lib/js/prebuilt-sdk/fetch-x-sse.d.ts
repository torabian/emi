import { TypedResponse } from "./fetch-x";
export declare const SSEFetch: <T = string>(res: TypedResponse<T>, onMessage?: (ev: MessageEvent) => void, signal?: AbortSignal) => {
    response: TypedResponse<T>;
    done: Promise<void>;
};
