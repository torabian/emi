export class TypedResponse extends Response {
    json() {
        return super.json();
    }
}
export function fetchx(input, init) {
    return fetch(input, init);
}
export async function handleFetchResponse(res, dto, onMessage, signal) {
    const ct = res.headers.get("content-type") || "";
    const cd = res.headers.get("content-disposition") || "";
    if (ct.includes("text/event-stream")) {
        return SSEFetch(res, onMessage, signal);
    }
    if (cd.includes("attachment") ||
        (!ct.includes("json") && !ct.startsWith("text/"))) {
        res.result = res.body;
    }
    else if (ct.includes("application/json")) {
        const json = await res.json();
        res.result = dto ? new dto(json) : json;
    }
    else {
        res.result = await res.text();
    }
    return { done: Promise.resolve(), response: res };
}
export const SSEFetch = (res, onMessage, signal) => {
    if (!res.body)
        throw new Error("SSE requires readable body");
    const reader = res.body.getReader();
    const decoder = new TextDecoder();
    let buffer = "";
    const done = new Promise((resolve, reject) => {
        function readChunk() {
            reader
                .read()
                .then(({ done: finished, value }) => {
                if (signal === null || signal === void 0 ? void 0 : signal.aborted) {
                    reader.cancel();
                    return resolve(); // resolve on abort
                }
                if (finished)
                    return resolve(); // normal end
                buffer += decoder.decode(value, { stream: true });
                const parts = buffer.split("\n\n");
                buffer = parts.pop() || "";
                for (const part of parts) {
                    let data = "";
                    let event = "message";
                    part.split("\n").forEach((line) => {
                        if (line.startsWith("data:"))
                            data += line.slice(5).trim();
                        else if (line.startsWith("event:"))
                            event = line.slice(6).trim();
                    });
                    if (data) {
                        if (data === "[DONE]")
                            return resolve();
                        onMessage === null || onMessage === void 0 ? void 0 : onMessage(new MessageEvent(event, { data }));
                    }
                }
                readChunk();
            })
                .catch((err) => {
                if (err.name === "AbortError")
                    resolve();
                else
                    reject(err);
            });
        }
        readChunk();
    });
    return { response: res, done };
};
