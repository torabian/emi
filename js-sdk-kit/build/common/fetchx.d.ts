export type TypedRequestInit<TBody = unknown, THeaders = unknown> = Omit<RequestInit, "body" | "headers"> & {
    body?: TBody;
    headers?: THeaders;
};
export declare class TypedResponse<T> extends Response {
    json(): Promise<T>;
    result: T | undefined | ReadableStream<Uint8Array<ArrayBuffer>> | null | string;
}
export declare function fetchx<TResponse = unknown, TBody = unknown, THeaders = unknown>(input: RequestInfo | URL, init?: TypedRequestInit<TBody, THeaders>, ctx?: FetchxContext<TBody, THeaders>): Promise<TypedResponse<TResponse>>;
type DtoFactory<T> = {
    new (data: any): T;
} | ((data: any) => T);
export declare function handleFetchResponse<T>(res: TypedResponse<T>, dto?: DtoFactory<T>, onMessage?: (msg: any) => void, signal?: AbortSignal | null): Promise<{
    done: Promise<void>;
    response: TypedResponse<T>;
}>;
export declare const SSEFetch: <T = string>(res: TypedResponse<T>, onMessage?: (ev: MessageEvent) => void, signal?: AbortSignal | null) => {
    response: TypedResponse<T>;
    done: Promise<void>;
};
export declare class FetchxContext<TBody, THeaders> {
    baseUrl: string;
    defaultHeaders: Record<string, string>;
    requestInterceptor?: (url: string, init: TypedRequestInit<TBody, THeaders>) => Promise<[string, TypedRequestInit<TBody, THeaders>]> | [string, TypedRequestInit<TBody, THeaders>];
    responseInterceptor?: <T>(res: TypedResponse<T>) => Promise<TypedResponse<T>>;
    constructor(baseUrl?: string, defaultHeaders?: Record<string, string>, requestInterceptor?: (url: string, init: TypedRequestInit<TBody, THeaders>) => Promise<[string, TypedRequestInit<TBody, THeaders>]> | [string, TypedRequestInit<TBody, THeaders>], responseInterceptor?: <T>(res: TypedResponse<T>) => Promise<TypedResponse<T>>);
    apply<T>(url: string, init: TypedRequestInit<TBody, THeaders>): Promise<[string, TypedRequestInit<TBody, THeaders>]>;
    handle<T>(res: TypedResponse<T>): Promise<TypedResponse<T>>;
    clone(overrides?: Partial<FetchxContext<TBody, THeaders>>): FetchxContext<TBody, THeaders>;
}
export {};
