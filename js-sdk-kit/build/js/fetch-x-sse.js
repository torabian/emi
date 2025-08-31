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
