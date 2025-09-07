export type TypedRequestInit<TBody = unknown, THeaders = unknown> = Omit<RequestInit, "body" | "headers"> & {
    body?: TBody;
    headers?: THeaders;
};
export declare class TypedResponse<T> extends Response {
    json(): Promise<T>;
    result: T | undefined | ReadableStream<Uint8Array<ArrayBuffer>> | null | string;
}
export declare function fetchx<TResponse = unknown, TBody = unknown, THeaders = unknown>(input: RequestInfo | URL, init?: TypedRequestInit<TBody, THeaders>): Promise<TypedResponse<TResponse>>;
