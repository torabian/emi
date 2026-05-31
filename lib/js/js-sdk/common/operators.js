export class MOne {
    constructor() {
        this.operation = null;
        this.content = undefined;
    }
    isNull() {
        return this.content === null;
    }
    static cast(value) {
        if (typeof value === "object" && (value === null || value === void 0 ? void 0 : value.__operation)) {
            const res = MOne.of(value === null || value === void 0 ? void 0 : value.content);
            res.operation = value === null || value === void 0 ? void 0 : value.__operation;
            return {
                ok: true,
                value: res,
            };
        }
        return {
            value: null,
            ok: false,
        };
    }
    isSelector() {
        return this.selector !== undefined && this.operation !== null;
    }
    get() {
        return this.content;
    }
    static of(value) {
        const one = new MOne();
        one.content = value;
        return one;
    }
    static select(selector) {
        const one = new MOne();
        one.selector = selector;
        one.operation = "replace";
        return one;
    }
    toJSON() {
        // When its explicit replace, it means that we need to pass a selector, so reader will
        // be able to replace it via internal mechanism
        if (this.operation === "replace") {
            return {
                __operation: this.operation,
                selector: this.selector,
            };
        }
        return this.content;
    }
}
// Array describes how an incoming list should be applied to an existing one,
// mirroring emigo.Array on the Go side. It lets a PATCH-style payload say
// "replace the whole set" versus "append to the existing set".
//
// On the wire a "replace" is implicit (a bare array), while an "append" is
// tagged with __operation so the reader keeps the existing rows. This keeps
// the payload identical to the Go client/backend generators.
export class MArray {
    constructor() {
        this.operation = "replace";
        this.items = [];
    }
    static cast(value) {
        if (typeof value === "object" && (value === null || value === void 0 ? void 0 : value.__operation)) {
            const res = MArray.of(value === null || value === void 0 ? void 0 : value.items);
            res.operation = value === null || value === void 0 ? void 0 : value.__operation;
            return {
                ok: true,
                value: res,
            };
        }
        return {
            value: null,
            ok: false,
        };
    }
    isAppend() {
        return this.operation === "append";
    }
    isReplace() {
        return this.operation === "replace";
    }
    len() {
        return this.items.length;
    }
    get() {
        return this.items;
    }
    // Full replacement — existing rows are cleared before these are applied.
    static of(items) {
        const arr = new MArray();
        arr.items = items;
        arr.operation = "replace";
        return arr;
    }
    // Append — existing rows are preserved and these are added alongside them.
    static append(items) {
        const arr = new MArray();
        arr.items = items;
        arr.operation = "append";
        return arr;
    }
    toJSON() {
        // "replace" is implicit on the wire, so we emit a bare array. Only the
        // "append" operation needs the explicit tagged-object form.
        if (this.operation === "append") {
            return {
                __operation: this.operation,
                items: this.items,
            };
        }
        return this.items;
    }
}
// Collection mirrors emigo.Collection on the Go side. Structurally it is the
// same as Array — a list carrying a "replace"/"append" operation — but it is a
// distinct field type: a collection holds a list of a target entity, whereas an
// array holds a list of an inline DTO.
export class MCollection {
    constructor() {
        this.operation = "replace";
        this.items = [];
    }
    isAppend() {
        return this.operation === "append";
    }
    static cast(value) {
        if (typeof value === "object" && (value === null || value === void 0 ? void 0 : value.__operation)) {
            const res = MCollection.of(value === null || value === void 0 ? void 0 : value.items);
            res.operation = value === null || value === void 0 ? void 0 : value.__operation;
            return {
                ok: true,
                value: res,
            };
        }
        return {
            value: null,
            ok: false,
        };
    }
    isReplace() {
        return this.operation === "replace";
    }
    len() {
        return this.items.length;
    }
    get() {
        return this.items;
    }
    // Full replacement — existing rows are cleared before these are applied.
    static of(items) {
        const collection = new MCollection();
        collection.items = items;
        collection.operation = "replace";
        return collection;
    }
    // Append — existing rows are preserved and these are added alongside them.
    static append(items) {
        const collection = new MCollection();
        collection.items = items;
        collection.operation = "append";
        return collection;
    }
    toJSON() {
        // "replace" is implicit on the wire, so we emit a bare array. Only the
        // "append" operation needs the explicit tagged-object form.
        if (this.operation === "append") {
            return {
                __operation: this.operation,
                items: this.items,
            };
        }
        return this.items;
    }
}
