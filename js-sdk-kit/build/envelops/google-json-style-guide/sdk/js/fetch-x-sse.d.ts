import { TypedResponse } from "./fetch-x";
export declare const SSEFetch: <T = string>(res: TypedResponse<T>, onMessage?: (ev: MessageEvent) => void, signal?: AbortSignal | null) => {
    response: TypedResponse<T>;
    done: Promise<void>;
};
export declare function handleFetchResponse<T>(res: TypedResponse<T>, dto?: new (data: any) => T, onMessage?: (msg: any) => void, signal?: AbortSignal | null): Promise<{
    done: Promise<void>;
    response: Response & {
        result?: any;
    };
}>;
