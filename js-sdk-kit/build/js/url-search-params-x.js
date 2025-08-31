import { stringify, parse } from "qs";
/**
 * Extended URLSearchParams that stores data in a nested object
 * and keeps compatibility with URLSearchParams methods.
 */
export class URLSearchParamsX extends URLSearchParams {
    constructor(init) {
        super(init);
        /** Internal data store */
        this.data = {};
        if (init) {
            if (typeof init === "string") {
                Object.assign(this.data, parse(init));
            }
            else if (init instanceof URLSearchParams) {
                Object.assign(this.data, parse(init.toString()));
            }
            else if (Array.isArray(init)) {
                init.forEach(([k, v]) => (this.data[k] = v));
            }
            else {
                Object.assign(this.data, init);
            }
        }
    }
    /** Remove a key from the store */
    delete(name) {
        delete this.data[name];
    }
    /** Append a value to an array or create a new array */
    append(name, value) {
        if (this.data[name] === undefined)
            this.data[name] = value;
        else if (Array.isArray(this.data[name]))
            this.data[name].push(value);
        else
            this.data[name] = [this.data[name], value];
    }
    /** Get an iterator of top-level keys */
    keys() {
        const obj = this.data;
        return (function* () {
            for (const key of Object.keys(obj)) {
                yield key;
            }
            return undefined;
        })();
    }
    /** Number of top-level keys */
    get size() {
        return Object.keys(this.data).length;
    }
    /** Sort top-level keys */
    sort() {
        const sorted = {};
        Object.keys(this.data)
            .sort()
            .forEach((key) => {
            // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
            sorted[key] = this.data[key];
        });
        this.data = sorted;
    }
    /** Get an iterator of top-level values */
    values() {
        const obj = this.data;
        return (function* () {
            for (const key of Object.keys(obj)) {
                const val = obj[key];
                // Make sure val is string
                yield String(val);
            }
            return undefined;
        })();
    }
    /** Get a single value by key */
    get(name) {
        // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
        const val = this.data[name];
        if (val == null)
            return null;
        return Array.isArray(val) ? String(val[0]) : String(val);
    }
    /** Check if key exists */
    has(name) {
        return this.data[name] !== undefined;
    }
    /** Iterate over top-level keys and values */
    forEach(
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    callbackfn) {
        for (const key of Object.keys(this.data)) {
            callbackfn(this.data[key], key, this);
        }
    }
    /** Get all values for a key as array */
    getAll(name) {
        // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
        const val = this.data[name];
        if (val === undefined)
            return [];
        return Array.isArray(val) ? val.map(String) : [String(val)];
    }
    /** Get an iterator of key/value pairs (flattened) */
    entries() {
        const params = new URLSearchParams(stringify(this.data));
        return params.entries();
    }
    /** Convert to query string */
    toString() {
        return stringify(this.data);
    }
    /** Set a key to a value */
    set(name, value) {
        // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
        this.data[name] = value;
        return this;
    }
    /** Convert entries to plain object */
    toObject() {
        return Object.fromEntries(this.entries());
    }
    // eslint-disable-next-line no-unused-private-class-members
    getTyped(key, type) {
        const val = this.get(key);
        if (val == null)
            return null;
        const t = type.toLowerCase();
        if (t.includes("number"))
            return Number(val);
        if (t.includes("bool"))
            return val === "true";
        return val;
    }
}
/**
 * Handy tool to create a final callable url, from query string, query params,
 * and the actual url.
 * @param template
 * @param params
 * @param qs
 * @returns
 */
export function buildUrl(url, params, qs) {
    // Replace :placeholders
    if (params) {
        Object.entries(params).forEach(([key, value]) => {
            url = url.replace(new RegExp(`:${key}`, "g"), encodeURIComponent(String(value)));
        });
    }
    if (qs && qs instanceof URLSearchParamsX) {
        url += `?${qs.toString()}`;
    }
    else if (qs && Object.keys(qs).length) {
        const query = new URLSearchParams(Object.entries(qs).map(([k, v]) => [k, String(v)])).toString();
        url += `?${query}`;
    }
    return url;
}
