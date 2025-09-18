export class TypedResponse extends Response {
    json() {
        return super.json();
    }
}
export async function fetchx(input, init, ctx) {
    let url = input.toString();
    let reqInit = init || {};
    if (ctx) {
        [url, reqInit] = await ctx.apply(url, reqInit);
    }
    let res = (await fetch(url, reqInit));
    if (ctx) {
        res = await ctx.handle(res);
    }
    return res;
}
function isConstructor(fn) {
    return (typeof fn === "function" && fn.prototype && fn.prototype.constructor === fn);
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
        if (dto) {
            if (isConstructor(dto)) {
                res.result = new dto(json); // ✅ class constructor
            }
            else {
                res.result = dto(json); // ✅ factory function
            }
        }
        else {
            res.result = json;
        }
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
export class FetchxContext {
    constructor(baseUrl = "", defaultHeaders = {}, requestInterceptor, responseInterceptor) {
        this.baseUrl = baseUrl;
        this.defaultHeaders = defaultHeaders;
        this.requestInterceptor = requestInterceptor;
        this.responseInterceptor = responseInterceptor;
    }
    async apply(url, init) {
        // prefix baseUrl
        if (!/^https?:\/\//.test(url)) {
            url = this.baseUrl + url;
        }
        // merge default headers
        init.headers = {
            ...this.defaultHeaders,
            ...(init.headers || {}),
        };
        // call request interceptor if present
        if (this.requestInterceptor) {
            return this.requestInterceptor(url, init);
        }
        return [url, init];
    }
    async handle(res) {
        if (this.responseInterceptor) {
            return this.responseInterceptor(res);
        }
        return res;
    }
    clone(overrides) {
        var _a, _b, _c;
        return new FetchxContext((_a = overrides === null || overrides === void 0 ? void 0 : overrides.baseUrl) !== null && _a !== void 0 ? _a : this.baseUrl, { ...this.defaultHeaders, ...((overrides === null || overrides === void 0 ? void 0 : overrides.defaultHeaders) || {}) }, (_b = overrides === null || overrides === void 0 ? void 0 : overrides.requestInterceptor) !== null && _b !== void 0 ? _b : this.requestInterceptor, (_c = overrides === null || overrides === void 0 ? void 0 : overrides.responseInterceptor) !== null && _c !== void 0 ? _c : this.responseInterceptor);
    }
}
