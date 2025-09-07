export type TypedRequestInit<TBody = unknown, THeaders = unknown> = Omit<
  RequestInit,
  "body" | "headers"
> & {
  body?: TBody;
  headers?: THeaders;
};

export class TypedResponse<T> extends Response {
  override json(): Promise<T> {
    return super.json();
  }

  result:
    | T
    | undefined
    | ReadableStream<Uint8Array<ArrayBuffer>>
    | null
    | string;
}

export function fetchx<
  TResponse = unknown,
  TBody = unknown,
  THeaders = unknown
>(
  input: RequestInfo | URL,
  init?: TypedRequestInit<TBody, THeaders>
): Promise<TypedResponse<TResponse>> {
  return fetch(input, init as RequestInit) as Promise<TypedResponse<TResponse>>;
}

export async function handleFetchResponse<T>(
  res: TypedResponse<T>,
  dto?: new (data: any) => T,
  onMessage?: (msg: any) => void,
  signal?: AbortSignal | null
): Promise<{ done: Promise<void>; response: Response & { result?: any } }> {
  const ct = res.headers.get("content-type") || "";
  const cd = res.headers.get("content-disposition") || "";

  if (ct.includes("text/event-stream")) {
    return SSEFetch(res, onMessage, signal);
  }

  if (
    cd.includes("attachment") ||
    (!ct.includes("json") && !ct.startsWith("text/"))
  ) {
    (res as any).result = res.body;
  } else if (ct.includes("application/json")) {
    const json = await res.json();
    (res as any).result = dto ? new dto(json) : json;
  } else {
    (res as any).result = await res.text();
  }

  return { done: Promise.resolve(), response: res as any };
}

export const SSEFetch = <T = string>(
  res: TypedResponse<T>,
  onMessage?: (ev: MessageEvent) => void,
  signal?: AbortSignal | null
): { response: TypedResponse<T>; done: Promise<void> } => {
  if (!res.body) throw new Error("SSE requires readable body");

  const reader = res.body.getReader();
  const decoder = new TextDecoder();
  let buffer = "";

  const done = new Promise<void>((resolve, reject) => {
    function readChunk() {
      reader
        .read()
        .then(({ done: finished, value }) => {
          if (signal?.aborted) {
            reader.cancel();
            return resolve(); // resolve on abort
          }

          if (finished) return resolve(); // normal end

          buffer += decoder.decode(value, { stream: true });
          const parts = buffer.split("\n\n");
          buffer = parts.pop() || "";

          for (const part of parts) {
            let data = "";
            let event = "message";

            part.split("\n").forEach((line) => {
              if (line.startsWith("data:")) data += line.slice(5).trim();
              else if (line.startsWith("event:")) event = line.slice(6).trim();
            });

            if (data) {
              if (data === "[DONE]") return resolve();
              onMessage?.(new MessageEvent(event, { data }));
            }
          }

          readChunk();
        })
        .catch((err) => {
          if (err.name === "AbortError") resolve();
          else reject(err);
        });
    }

    readChunk();
  });

  return { response: res, done };
};
