var __classPrivateFieldGet = (this && this.__classPrivateFieldGet) || function (receiver, state, kind, f) {
    if (kind === "a" && !f) throw new TypeError("Private accessor was defined without a getter");
    if (typeof state === "function" ? receiver !== state || !f : !state.has(receiver)) throw new TypeError("Cannot read private member from an object whose class did not declare it");
    return kind === "m" ? f : kind === "a" ? f.call(receiver) : f ? f.value : state.get(receiver);
};
var __classPrivateFieldSet = (this && this.__classPrivateFieldSet) || function (receiver, state, value, kind, f) {
    if (kind === "m") throw new TypeError("Private method is not writable");
    if (kind === "a" && !f) throw new TypeError("Private accessor was defined without a setter");
    if (typeof state === "function" ? receiver !== state || !f : !state.has(receiver)) throw new TypeError("Cannot write private member to an object whose class did not declare it");
    return (kind === "a" ? f.call(receiver, value) : f ? f.value = value : state.set(receiver, value)), value;
};
var _ResponseDto_apiVersion, _ResponseDto_context, _ResponseDto_data;
import { isPlausibleObject } from "./sdk/js/index";
/**
 * The base class definition for responseDto
 **/
export class ResponseDto {
    /**
     *
     * @returns {string}
     **/
    get apiVersion() {
        return __classPrivateFieldGet(this, _ResponseDto_apiVersion, "f");
    }
    /**
     *
     * @type {string}
     **/
    set apiVersion(value) {
        const correctType = typeof value === "string";
        __classPrivateFieldSet(this, _ResponseDto_apiVersion, correctType ? value : "" + value, "f");
    }
    setApiVersion(value) {
        this.apiVersion = value;
        return this;
    }
    /**
     *
     * @returns {string}
     **/
    get context() {
        return __classPrivateFieldGet(this, _ResponseDto_context, "f");
    }
    /**
     *
     * @type {string}
     **/
    set context(value) {
        const correctType = typeof value === "string";
        __classPrivateFieldSet(this, _ResponseDto_context, correctType ? value : "" + value, "f");
    }
    setContext(value) {
        this.context = value;
        return this;
    }
    /**
     *
     * @returns {string}
     **/
    get data() {
        return __classPrivateFieldGet(this, _ResponseDto_data, "f");
    }
    /**
     *
     * @type {string}
     **/
    set data(value) { }
    setData(value) {
        this.data = value;
        return this;
    }
    constructor(data) {
        /**
         *
         * @type {string}
         **/
        _ResponseDto_apiVersion.set(this, "");
        /**
         *
         * @type {string}
         **/
        _ResponseDto_context.set(this, "");
        /**
         *
         * @type {string}
         **/
        _ResponseDto_data.set(this, null);
        if (data === null || data === undefined) {
            return;
        }
        if (typeof data === "string") {
            this.applyFromObject(JSON.parse(data));
        }
        else if (isPlausibleObject(data)) {
            this.applyFromObject(data);
        }
        else {
            throw new Error("Instance is not implemented.");
        }
    }
    /**
     * casts the fields of a javascript object into the class properties one by one
     **/
    applyFromObject(data = {}) {
        const d = data;
        if (d.apiVersion !== undefined) {
            this.apiVersion = d.apiVersion;
        }
        if (d.context !== undefined) {
            this.context = d.context;
        }
        if (d.data !== undefined) {
            this.data = d.data;
        }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
        return {
            apiVersion: __classPrivateFieldGet(this, _ResponseDto_apiVersion, "f"),
            context: __classPrivateFieldGet(this, _ResponseDto_context, "f"),
            data: __classPrivateFieldGet(this, _ResponseDto_data, "f"),
        };
    }
    toString() {
        return JSON.stringify(this);
    }
    static get Fields() {
        return {
            apiVersion: "apiVersion",
            context: "context",
            data: "data",
        };
    }
}
_ResponseDto_apiVersion = new WeakMap(), _ResponseDto_context = new WeakMap(), _ResponseDto_data = new WeakMap();
export class ResponseDtoFactory {
}
