import type { TypedRequestInit } from "../js/fetch-x";
export interface UseSSEResult<T> {
    messages: Array<T>;
    error?: Error | null;
    cancel: () => void;
    restart: () => void;
}
type SseFetchFn<Q, H, R> = (onMessage?: (ev: MessageEvent) => void, qs?: Q, init?: TypedRequestInit<R, H>, overrideUrl?: string) => Promise<any>;
export declare function useSse<T = string, Q = any, H = any, R = any>(fetchFn: SseFetchFn<Q, H, R>, props?: {
    qs?: Q;
    init?: TypedRequestInit<R, H>;
    overrideUrl?: string;
}): {
    cancel: () => void;
    restart: () => void;
    messages: T[];
    error?: Error | null;
};
export {};
