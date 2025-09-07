export function withPrefix(prefix, fields) {
    const out = {};
    for (const [k, v] of Object.entries(fields)) {
        if (typeof v === "string") {
            out[k] = `${prefix}.${v}`;
        }
        else if (typeof v === "object" && v !== null) {
            out[k] = v;
        }
    }
    return out;
}
export function at(source, ...args) {
    args.forEach((item) => {
        source = source.replace("[:i]", `[${item}]`);
    });
    return source;
}
