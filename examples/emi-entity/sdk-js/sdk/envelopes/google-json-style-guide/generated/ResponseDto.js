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
var _ResponseDto_instances, _a, _ResponseDto_apiVersion, _ResponseDto_context, _ResponseDto_id, _ResponseDto_method, _ResponseDto_params, _ResponseDto_data, _ResponseDto_error, _ResponseDto_isJsonAppliable, _ResponseDto_lateInitFields, _Data_instances, _Data_item, _Data_items, _Data_editLink, _Data_selfLink, _Data_kind, _Data_fields, _Data_etag, _Data_cursor, _Data_id, _Data_lang, _Data_updated, _Data_currentItemCount, _Data_itemsPerPage, _Data_startIndex, _Data_totalItems, _Data_totalAvailableItems, _Data_pageIndex, _Data_totalPages, _Data_isJsonAppliable, _b, _Error_instances, _Error_code, _Error_message, _Error_messageTranslated, _Error_errors, _Error_isJsonAppliable, _c, _Errors_instances, _Errors_domain, _Errors_reason, _Errors_message, _Errors_messageTranslated, _Errors_location, _Errors_locationType, _Errors_extendedHelp, _Errors_sendReport, _Errors_isJsonAppliable, _d;
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for responseDto
 **/
export class ResponseDto {
    /**
     * Version of the API used for this response.
     * @returns {string}
     **/
    get apiVersion() {
        return __classPrivateFieldGet(this, _ResponseDto_apiVersion, "f");
    }
    /**
     * Version of the API used for this response.
     * @type {string}
     **/
    set apiVersion(value) {
        const correctType = typeof value === "string" || value === undefined || value === null;
        __classPrivateFieldSet(this, _ResponseDto_apiVersion, correctType ? value : String(value), "f");
    }
    setApiVersion(value) {
        this.apiVersion = value;
        return this;
    }
    /**
     * Context string provided by the client or system for request tracking.
     * @returns {string}
     **/
    get context() {
        return __classPrivateFieldGet(this, _ResponseDto_context, "f");
    }
    /**
     * Context string provided by the client or system for request tracking.
     * @type {string}
     **/
    set context(value) {
        const correctType = typeof value === "string" || value === undefined || value === null;
        __classPrivateFieldSet(this, _ResponseDto_context, correctType ? value : String(value), "f");
    }
    setContext(value) {
        this.context = value;
        return this;
    }
    /**
     * Unique identifier assigned to the request/response.
     * @returns {string}
     **/
    get id() {
        return __classPrivateFieldGet(this, _ResponseDto_id, "f");
    }
    /**
     * Unique identifier assigned to the request/response.
     * @type {string}
     **/
    set id(value) {
        const correctType = typeof value === "string" || value === undefined || value === null;
        __classPrivateFieldSet(this, _ResponseDto_id, correctType ? value : String(value), "f");
    }
    setId(value) {
        this.id = value;
        return this;
    }
    /**
     * Name of the API method invoked.
     * @returns {string}
     **/
    get method() {
        return __classPrivateFieldGet(this, _ResponseDto_method, "f");
    }
    /**
     * Name of the API method invoked.
     * @type {string}
     **/
    set method(value) {
        const correctType = typeof value === "string" || value === undefined || value === null;
        __classPrivateFieldSet(this, _ResponseDto_method, correctType ? value : String(value), "f");
    }
    setMethod(value) {
        this.method = value;
        return this;
    }
    /**
     * Parameters sent with the request.
     * @returns {any}
     **/
    get params() {
        return __classPrivateFieldGet(this, _ResponseDto_params, "f");
    }
    /**
     * Parameters sent with the request.
     * @type {any}
     **/
    set params(value) {
        __classPrivateFieldSet(this, _ResponseDto_params, value, "f");
    }
    setParams(value) {
        this.params = value;
        return this;
    }
    /**
     * Main data payload of the response.
     * @returns {ResponseDto.Data<T>}
     **/
    get data() {
        return __classPrivateFieldGet(this, _ResponseDto_data, "f");
    }
    /**
     * Main data payload of the response.
     * @type {ResponseDto.Data}
     **/
    set data(value) {
        // For objects, the sub type needs to always be instance of the sub class.
        if (value instanceof _a.Data) {
            __classPrivateFieldSet(this, _ResponseDto_data, value, "f");
        }
        else {
            __classPrivateFieldSet(this, _ResponseDto_data, new _a.Data(value), "f");
        }
    }
    setData(value) {
        this.data = value;
        return this;
    }
    /**
     * Error details, if the request failed.
     * @returns {ResponseDto.Error}
     **/
    get error() {
        return __classPrivateFieldGet(this, _ResponseDto_error, "f");
    }
    /**
     * Error details, if the request failed.
     * @type {ResponseDto.Error}
     **/
    set error(value) {
        // For objects, the sub type needs to always be instance of the sub class.
        if (value instanceof _a.Error) {
            __classPrivateFieldSet(this, _ResponseDto_error, value, "f");
        }
        else {
            __classPrivateFieldSet(this, _ResponseDto_error, new _a.Error(value), "f");
        }
    }
    setError(value) {
        this.error = value;
        return this;
    }
    constructor(data = undefined) {
        _ResponseDto_instances.add(this);
        /**
         * Version of the API used for this response.
         * @type {string}
         **/
        _ResponseDto_apiVersion.set(this, undefined);
        /**
         * Context string provided by the client or system for request tracking.
         * @type {string}
         **/
        _ResponseDto_context.set(this, undefined);
        /**
         * Unique identifier assigned to the request/response.
         * @type {string}
         **/
        _ResponseDto_id.set(this, undefined);
        /**
         * Name of the API method invoked.
         * @type {string}
         **/
        _ResponseDto_method.set(this, undefined);
        /**
         * Parameters sent with the request.
         * @type {any}
         **/
        _ResponseDto_params.set(this, null);
        /**
         * Main data payload of the response.
         * @type {ResponseDto.Data<T>}
         **/
        _ResponseDto_data.set(this, void 0);
        /**
         * Error details, if the request failed.
         * @type {ResponseDto.Error}
         **/
        _ResponseDto_error.set(this, void 0);
        if (data === null || data === undefined) {
            __classPrivateFieldGet(this, _ResponseDto_instances, "m", _ResponseDto_lateInitFields).call(this);
            return;
        }
        if (typeof data === "string") {
            this.applyFromObject(JSON.parse(data));
        }
        else if (__classPrivateFieldGet(this, _ResponseDto_instances, "m", _ResponseDto_isJsonAppliable).call(this, data)) {
            this.applyFromObject(data);
        }
        else {
            throw new Error("Instance cannot be created on an unknown value, check the content being passed. got: " +
                typeof data);
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
        if (d.id !== undefined) {
            this.id = d.id;
        }
        if (d.method !== undefined) {
            this.method = d.method;
        }
        if (d.params !== undefined) {
            this.params = d.params;
        }
        if (d.data !== undefined) {
            this.data = d.data;
        }
        if (d.error !== undefined) {
            this.error = d.error;
        }
        __classPrivateFieldGet(this, _ResponseDto_instances, "m", _ResponseDto_lateInitFields).call(this, data);
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
        return {
            apiVersion: __classPrivateFieldGet(this, _ResponseDto_apiVersion, "f"),
            context: __classPrivateFieldGet(this, _ResponseDto_context, "f"),
            id: __classPrivateFieldGet(this, _ResponseDto_id, "f"),
            method: __classPrivateFieldGet(this, _ResponseDto_method, "f"),
            params: __classPrivateFieldGet(this, _ResponseDto_params, "f"),
            data: __classPrivateFieldGet(this, _ResponseDto_data, "f"),
            error: __classPrivateFieldGet(this, _ResponseDto_error, "f"),
        };
    }
    toString() {
        return JSON.stringify(this);
    }
    static get Fields() {
        return {
            apiVersion: "apiVersion",
            context: "context",
            id: "id",
            method: "method",
            params: "params",
            data$: "data",
            get data() {
                return withPrefix("data", _a.Data.Fields);
            },
            error$: "error",
            get error() {
                return withPrefix("error", _a.Error.Fields);
            },
        };
    }
    /**
     * Creates an instance of ResponseDto, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
        return new _a(possibleDtoObject);
    }
    /**
     * Creates an instance of ResponseDto, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
        return new _a(partialDtoObject);
    }
    copyWith(partial) {
        return new _a({ ...this.toJSON(), ...partial });
    }
    clone() {
        return new _a(this.toJSON());
    }
}
_a = ResponseDto, _ResponseDto_apiVersion = new WeakMap(), _ResponseDto_context = new WeakMap(), _ResponseDto_id = new WeakMap(), _ResponseDto_method = new WeakMap(), _ResponseDto_params = new WeakMap(), _ResponseDto_data = new WeakMap(), _ResponseDto_error = new WeakMap(), _ResponseDto_instances = new WeakSet(), _ResponseDto_isJsonAppliable = function _ResponseDto_isJsonAppliable(obj) {
    const g = globalThis;
    const isBuffer = typeof g.Buffer !== "undefined" &&
        typeof g.Buffer.isBuffer === "function" &&
        g.Buffer.isBuffer(obj);
    const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
    return (obj &&
        typeof obj === "object" &&
        !Array.isArray(obj) &&
        !isBuffer &&
        !(obj instanceof ArrayBuffer) &&
        !isBlob);
}, _ResponseDto_lateInitFields = function _ResponseDto_lateInitFields(data = {}) {
    const d = data;
    if (!(d.data instanceof _a.Data)) {
        this.data = new _a.Data(d.data || {});
    }
    if (!(d.error instanceof _a.Error)) {
        this.error = new _a.Error(d.error || {});
    }
};
/**
 * The base class definition for data
 **/
ResponseDto.Data = (_b = class Data {
        /**
         * Single item returned by the API.
         * @returns {T}
         **/
        get item() {
            return __classPrivateFieldGet(this, _Data_item, "f");
        }
        /**
         * Single item returned by the API.
         * @type {any}
         **/
        set item(value) {
            __classPrivateFieldSet(this, _Data_item, value, "f");
        }
        setItem(value) {
            this.item = value;
            return this;
        }
        /**
         * List of items returned by the API.
         * @returns {T[]}
         **/
        get items() {
            return __classPrivateFieldGet(this, _Data_items, "f");
        }
        /**
         * List of items returned by the API.
         * @type {T[]}
         **/
        set items(value) {
            __classPrivateFieldSet(this, _Data_items, value, "f");
        }
        setItems(value) {
            this.items = value;
            return this;
        }
        /**
         * Link to edit this resource.
         * @returns {string}
         **/
        get editLink() {
            return __classPrivateFieldGet(this, _Data_editLink, "f");
        }
        /**
         * Link to edit this resource.
         * @type {string}
         **/
        set editLink(value) {
            const correctType = typeof value === "string" || value === undefined || value === null;
            __classPrivateFieldSet(this, _Data_editLink, correctType ? value : String(value), "f");
        }
        setEditLink(value) {
            this.editLink = value;
            return this;
        }
        /**
         * Link to retrieve this resource.
         * @returns {string}
         **/
        get selfLink() {
            return __classPrivateFieldGet(this, _Data_selfLink, "f");
        }
        /**
         * Link to retrieve this resource.
         * @type {string}
         **/
        set selfLink(value) {
            const correctType = typeof value === "string" || value === undefined || value === null;
            __classPrivateFieldSet(this, _Data_selfLink, correctType ? value : String(value), "f");
        }
        setSelfLink(value) {
            this.selfLink = value;
            return this;
        }
        /**
         * Resource type (kind) identifier.
         * @returns {string}
         **/
        get kind() {
            return __classPrivateFieldGet(this, _Data_kind, "f");
        }
        /**
         * Resource type (kind) identifier.
         * @type {string}
         **/
        set kind(value) {
            const correctType = typeof value === "string" || value === undefined || value === null;
            __classPrivateFieldSet(this, _Data_kind, correctType ? value : String(value), "f");
        }
        setKind(value) {
            this.kind = value;
            return this;
        }
        /**
         * Selector specifying which fields are included in a partial response.
         * @returns {string}
         **/
        get fields() {
            return __classPrivateFieldGet(this, _Data_fields, "f");
        }
        /**
         * Selector specifying which fields are included in a partial response.
         * @type {string}
         **/
        set fields(value) {
            const correctType = typeof value === "string" || value === undefined || value === null;
            __classPrivateFieldSet(this, _Data_fields, correctType ? value : String(value), "f");
        }
        setFields(value) {
            this.fields = value;
            return this;
        }
        /**
         * ETag of the resource, used for caching/version control.
         * @returns {string}
         **/
        get etag() {
            return __classPrivateFieldGet(this, _Data_etag, "f");
        }
        /**
         * ETag of the resource, used for caching/version control.
         * @type {string}
         **/
        set etag(value) {
            const correctType = typeof value === "string" || value === undefined || value === null;
            __classPrivateFieldSet(this, _Data_etag, correctType ? value : String(value), "f");
        }
        setEtag(value) {
            this.etag = value;
            return this;
        }
        /**
         * Cursor for paginated data fetching.
         * @returns {string}
         **/
        get cursor() {
            return __classPrivateFieldGet(this, _Data_cursor, "f");
        }
        /**
         * Cursor for paginated data fetching.
         * @type {string}
         **/
        set cursor(value) {
            const correctType = typeof value === "string" || value === undefined || value === null;
            __classPrivateFieldSet(this, _Data_cursor, correctType ? value : String(value), "f");
        }
        setCursor(value) {
            this.cursor = value;
            return this;
        }
        /**
         * Unique identifier of the resource.
         * @returns {string}
         **/
        get id() {
            return __classPrivateFieldGet(this, _Data_id, "f");
        }
        /**
         * Unique identifier of the resource.
         * @type {string}
         **/
        set id(value) {
            const correctType = typeof value === "string" || value === undefined || value === null;
            __classPrivateFieldSet(this, _Data_id, correctType ? value : String(value), "f");
        }
        setId(value) {
            this.id = value;
            return this;
        }
        /**
         * Language code of the response data.
         * @returns {string}
         **/
        get lang() {
            return __classPrivateFieldGet(this, _Data_lang, "f");
        }
        /**
         * Language code of the response data.
         * @type {string}
         **/
        set lang(value) {
            const correctType = typeof value === "string" || value === undefined || value === null;
            __classPrivateFieldSet(this, _Data_lang, correctType ? value : String(value), "f");
        }
        setLang(value) {
            this.lang = value;
            return this;
        }
        /**
         * Last modification time of the resource.
         * @returns {string}
         **/
        get updated() {
            return __classPrivateFieldGet(this, _Data_updated, "f");
        }
        /**
         * Last modification time of the resource.
         * @type {string}
         **/
        set updated(value) {
            const correctType = typeof value === "string" || value === undefined || value === null;
            __classPrivateFieldSet(this, _Data_updated, correctType ? value : String(value), "f");
        }
        setUpdated(value) {
            this.updated = value;
            return this;
        }
        /**
         * Number of items in the current response page.
         * @returns {number}
         **/
        get currentItemCount() {
            return __classPrivateFieldGet(this, _Data_currentItemCount, "f");
        }
        /**
         * Number of items in the current response page.
         * @type {number}
         **/
        set currentItemCount(value) {
            const correctType = typeof value === "number" || value === undefined || value === null;
            const parsedValue = correctType ? value : Number(value);
            if (!Number.isNaN(parsedValue)) {
                __classPrivateFieldSet(this, _Data_currentItemCount, parsedValue, "f");
            }
        }
        setCurrentItemCount(value) {
            this.currentItemCount = value;
            return this;
        }
        /**
         * Maximum number of items per page.
         * @returns {number}
         **/
        get itemsPerPage() {
            return __classPrivateFieldGet(this, _Data_itemsPerPage, "f");
        }
        /**
         * Maximum number of items per page.
         * @type {number}
         **/
        set itemsPerPage(value) {
            const correctType = typeof value === "number" || value === undefined || value === null;
            const parsedValue = correctType ? value : Number(value);
            if (!Number.isNaN(parsedValue)) {
                __classPrivateFieldSet(this, _Data_itemsPerPage, parsedValue, "f");
            }
        }
        setItemsPerPage(value) {
            this.itemsPerPage = value;
            return this;
        }
        /**
         * Index of the first item in the current page.
         * @returns {number}
         **/
        get startIndex() {
            return __classPrivateFieldGet(this, _Data_startIndex, "f");
        }
        /**
         * Index of the first item in the current page.
         * @type {number}
         **/
        set startIndex(value) {
            const correctType = typeof value === "number" || value === undefined || value === null;
            const parsedValue = correctType ? value : Number(value);
            if (!Number.isNaN(parsedValue)) {
                __classPrivateFieldSet(this, _Data_startIndex, parsedValue, "f");
            }
        }
        setStartIndex(value) {
            this.startIndex = value;
            return this;
        }
        /**
         * Total number of items available.
         * @returns {number}
         **/
        get totalItems() {
            return __classPrivateFieldGet(this, _Data_totalItems, "f");
        }
        /**
         * Total number of items available.
         * @type {number}
         **/
        set totalItems(value) {
            const correctType = typeof value === "number" || value === undefined || value === null;
            const parsedValue = correctType ? value : Number(value);
            if (!Number.isNaN(parsedValue)) {
                __classPrivateFieldSet(this, _Data_totalItems, parsedValue, "f");
            }
        }
        setTotalItems(value) {
            this.totalItems = value;
            return this;
        }
        /**
         * Number of items available for this user/query.
         * @returns {number}
         **/
        get totalAvailableItems() {
            return __classPrivateFieldGet(this, _Data_totalAvailableItems, "f");
        }
        /**
         * Number of items available for this user/query.
         * @type {number}
         **/
        set totalAvailableItems(value) {
            const correctType = typeof value === "number" || value === undefined || value === null;
            const parsedValue = correctType ? value : Number(value);
            if (!Number.isNaN(parsedValue)) {
                __classPrivateFieldSet(this, _Data_totalAvailableItems, parsedValue, "f");
            }
        }
        setTotalAvailableItems(value) {
            this.totalAvailableItems = value;
            return this;
        }
        /**
         * Current page index in the pagination.
         * @returns {number}
         **/
        get pageIndex() {
            return __classPrivateFieldGet(this, _Data_pageIndex, "f");
        }
        /**
         * Current page index in the pagination.
         * @type {number}
         **/
        set pageIndex(value) {
            const correctType = typeof value === "number" || value === undefined || value === null;
            const parsedValue = correctType ? value : Number(value);
            if (!Number.isNaN(parsedValue)) {
                __classPrivateFieldSet(this, _Data_pageIndex, parsedValue, "f");
            }
        }
        setPageIndex(value) {
            this.pageIndex = value;
            return this;
        }
        /**
         * Total number of pages in the pagination.
         * @returns {number}
         **/
        get totalPages() {
            return __classPrivateFieldGet(this, _Data_totalPages, "f");
        }
        /**
         * Total number of pages in the pagination.
         * @type {number}
         **/
        set totalPages(value) {
            const correctType = typeof value === "number" || value === undefined || value === null;
            const parsedValue = correctType ? value : Number(value);
            if (!Number.isNaN(parsedValue)) {
                __classPrivateFieldSet(this, _Data_totalPages, parsedValue, "f");
            }
        }
        setTotalPages(value) {
            this.totalPages = value;
            return this;
        }
        constructor(data = undefined) {
            _Data_instances.add(this);
            /**
             * Single item returned by the API.
             * @type {any}
             **/
            _Data_item.set(this, null);
            /**
             * List of items returned by the API.
             * @type {any}
             **/
            _Data_items.set(this, []);
            /**
             * Link to edit this resource.
             * @type {string}
             **/
            _Data_editLink.set(this, undefined);
            /**
             * Link to retrieve this resource.
             * @type {string}
             **/
            _Data_selfLink.set(this, undefined);
            /**
             * Resource type (kind) identifier.
             * @type {string}
             **/
            _Data_kind.set(this, undefined);
            /**
             * Selector specifying which fields are included in a partial response.
             * @type {string}
             **/
            _Data_fields.set(this, undefined);
            /**
             * ETag of the resource, used for caching/version control.
             * @type {string}
             **/
            _Data_etag.set(this, undefined);
            /**
             * Cursor for paginated data fetching.
             * @type {string}
             **/
            _Data_cursor.set(this, undefined);
            /**
             * Unique identifier of the resource.
             * @type {string}
             **/
            _Data_id.set(this, undefined);
            /**
             * Language code of the response data.
             * @type {string}
             **/
            _Data_lang.set(this, undefined);
            /**
             * Last modification time of the resource.
             * @type {string}
             **/
            _Data_updated.set(this, undefined);
            /**
             * Number of items in the current response page.
             * @type {number}
             **/
            _Data_currentItemCount.set(this, undefined);
            /**
             * Maximum number of items per page.
             * @type {number}
             **/
            _Data_itemsPerPage.set(this, undefined);
            /**
             * Index of the first item in the current page.
             * @type {number}
             **/
            _Data_startIndex.set(this, undefined);
            /**
             * Total number of items available.
             * @type {number}
             **/
            _Data_totalItems.set(this, undefined);
            /**
             * Number of items available for this user/query.
             * @type {number}
             **/
            _Data_totalAvailableItems.set(this, undefined);
            /**
             * Current page index in the pagination.
             * @type {number}
             **/
            _Data_pageIndex.set(this, undefined);
            /**
             * Total number of pages in the pagination.
             * @type {number}
             **/
            _Data_totalPages.set(this, undefined);
            if (data === null || data === undefined) {
                return;
            }
            if (typeof data === "string") {
                this.applyFromObject(JSON.parse(data));
            }
            else if (__classPrivateFieldGet(this, _Data_instances, "m", _Data_isJsonAppliable).call(this, data)) {
                this.applyFromObject(data);
            }
            else {
                throw new Error("Instance cannot be created on an unknown value, check the content being passed. got: " +
                    typeof data);
            }
        }
        /**
         * casts the fields of a javascript object into the class properties one by one
         **/
        applyFromObject(data = {}) {
            const d = data;
            if (d.item !== undefined) {
                this.item = d.item;
            }
            if (d.items !== undefined) {
                this.items = d.items;
            }
            if (d.editLink !== undefined) {
                this.editLink = d.editLink;
            }
            if (d.selfLink !== undefined) {
                this.selfLink = d.selfLink;
            }
            if (d.kind !== undefined) {
                this.kind = d.kind;
            }
            if (d.fields !== undefined) {
                this.fields = d.fields;
            }
            if (d.etag !== undefined) {
                this.etag = d.etag;
            }
            if (d.cursor !== undefined) {
                this.cursor = d.cursor;
            }
            if (d.id !== undefined) {
                this.id = d.id;
            }
            if (d.lang !== undefined) {
                this.lang = d.lang;
            }
            if (d.updated !== undefined) {
                this.updated = d.updated;
            }
            if (d.currentItemCount !== undefined) {
                this.currentItemCount = d.currentItemCount;
            }
            if (d.itemsPerPage !== undefined) {
                this.itemsPerPage = d.itemsPerPage;
            }
            if (d.startIndex !== undefined) {
                this.startIndex = d.startIndex;
            }
            if (d.totalItems !== undefined) {
                this.totalItems = d.totalItems;
            }
            if (d.totalAvailableItems !== undefined) {
                this.totalAvailableItems = d.totalAvailableItems;
            }
            if (d.pageIndex !== undefined) {
                this.pageIndex = d.pageIndex;
            }
            if (d.totalPages !== undefined) {
                this.totalPages = d.totalPages;
            }
        }
        /**
         *	Special toJSON override, since the field are private,
         *	Json stringify won't see them unless we mention it explicitly.
         **/
        toJSON() {
            return {
                item: __classPrivateFieldGet(this, _Data_item, "f"),
                items: __classPrivateFieldGet(this, _Data_items, "f"),
                editLink: __classPrivateFieldGet(this, _Data_editLink, "f"),
                selfLink: __classPrivateFieldGet(this, _Data_selfLink, "f"),
                kind: __classPrivateFieldGet(this, _Data_kind, "f"),
                fields: __classPrivateFieldGet(this, _Data_fields, "f"),
                etag: __classPrivateFieldGet(this, _Data_etag, "f"),
                cursor: __classPrivateFieldGet(this, _Data_cursor, "f"),
                id: __classPrivateFieldGet(this, _Data_id, "f"),
                lang: __classPrivateFieldGet(this, _Data_lang, "f"),
                updated: __classPrivateFieldGet(this, _Data_updated, "f"),
                currentItemCount: __classPrivateFieldGet(this, _Data_currentItemCount, "f"),
                itemsPerPage: __classPrivateFieldGet(this, _Data_itemsPerPage, "f"),
                startIndex: __classPrivateFieldGet(this, _Data_startIndex, "f"),
                totalItems: __classPrivateFieldGet(this, _Data_totalItems, "f"),
                totalAvailableItems: __classPrivateFieldGet(this, _Data_totalAvailableItems, "f"),
                pageIndex: __classPrivateFieldGet(this, _Data_pageIndex, "f"),
                totalPages: __classPrivateFieldGet(this, _Data_totalPages, "f"),
            };
        }
        toString() {
            return JSON.stringify(this);
        }
        static get Fields() {
            return {
                item: "item",
                items: "items",
                editLink: "editLink",
                selfLink: "selfLink",
                kind: "kind",
                fields: "fields",
                etag: "etag",
                cursor: "cursor",
                id: "id",
                lang: "lang",
                updated: "updated",
                currentItemCount: "currentItemCount",
                itemsPerPage: "itemsPerPage",
                startIndex: "startIndex",
                totalItems: "totalItems",
                totalAvailableItems: "totalAvailableItems",
                pageIndex: "pageIndex",
                totalPages: "totalPages",
            };
        }
        /**
         * Creates an instance of ResponseDto.Data, and possibleDtoObject
         * needs to satisfy the type requirement fully, otherwise typescript compile would
         * be complaining.
         **/
        static from(possibleDtoObject) {
            return new _a.Data(possibleDtoObject);
        }
        /**
         * Creates an instance of ResponseDto.Data, and partialDtoObject
         * needs to satisfy the type, but partially, and rest of the content would
         * be constructed according to data types and nullability.
         **/
        static with(partialDtoObject) {
            return new _a.Data(partialDtoObject);
        }
        copyWith(partial) {
            return new _a.Data({ ...this.toJSON(), ...partial });
        }
        clone() {
            return new _a.Data(this.toJSON());
        }
    },
    _Data_item = new WeakMap(),
    _Data_items = new WeakMap(),
    _Data_editLink = new WeakMap(),
    _Data_selfLink = new WeakMap(),
    _Data_kind = new WeakMap(),
    _Data_fields = new WeakMap(),
    _Data_etag = new WeakMap(),
    _Data_cursor = new WeakMap(),
    _Data_id = new WeakMap(),
    _Data_lang = new WeakMap(),
    _Data_updated = new WeakMap(),
    _Data_currentItemCount = new WeakMap(),
    _Data_itemsPerPage = new WeakMap(),
    _Data_startIndex = new WeakMap(),
    _Data_totalItems = new WeakMap(),
    _Data_totalAvailableItems = new WeakMap(),
    _Data_pageIndex = new WeakMap(),
    _Data_totalPages = new WeakMap(),
    _Data_instances = new WeakSet(),
    _Data_isJsonAppliable = function _Data_isJsonAppliable(obj) {
        const g = globalThis;
        const isBuffer = typeof g.Buffer !== "undefined" &&
            typeof g.Buffer.isBuffer === "function" &&
            g.Buffer.isBuffer(obj);
        const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
        return (obj &&
            typeof obj === "object" &&
            !Array.isArray(obj) &&
            !isBuffer &&
            !(obj instanceof ArrayBuffer) &&
            !isBlob);
    },
    _b);
/**
 * The base class definition for error
 **/
ResponseDto.Error = (_c = class Error {
        /**
         * Numeric error code representing the failure.
         * @returns {number}
         **/
        get code() {
            return __classPrivateFieldGet(this, _Error_code, "f");
        }
        /**
         * Numeric error code representing the failure.
         * @type {number}
         **/
        set code(value) {
            const correctType = typeof value === "number";
            const parsedValue = correctType ? value : Number(value);
            if (!Number.isNaN(parsedValue)) {
                __classPrivateFieldSet(this, _Error_code, parsedValue, "f");
            }
        }
        setCode(value) {
            this.code = value;
            return this;
        }
        /**
         * Human-readable explanation of the error.
         * @returns {string}
         **/
        get message() {
            return __classPrivateFieldGet(this, _Error_message, "f");
        }
        /**
         * Human-readable explanation of the error.
         * @type {string}
         **/
        set message(value) {
            __classPrivateFieldSet(this, _Error_message, String(value), "f");
        }
        setMessage(value) {
            this.message = value;
            return this;
        }
        /**
         * Localized/translated version of the error message.
         * @returns {string}
         **/
        get messageTranslated() {
            return __classPrivateFieldGet(this, _Error_messageTranslated, "f");
        }
        /**
         * Localized/translated version of the error message.
         * @type {string}
         **/
        set messageTranslated(value) {
            __classPrivateFieldSet(this, _Error_messageTranslated, String(value), "f");
        }
        setMessageTranslated(value) {
            this.messageTranslated = value;
            return this;
        }
        /**
         * Detailed list of error objects.
         * @returns {ResponseDto.Error.Errors}
         **/
        get errors() {
            return __classPrivateFieldGet(this, _Error_errors, "f");
        }
        /**
         * Detailed list of error objects.
         * @type {ResponseDto.Error.Errors}
         **/
        set errors(value) {
            // For arrays, you only can pass arrays to the object
            if (!Array.isArray(value)) {
                return;
            }
            if (value.length > 0 && value[0] instanceof _a.Error.Errors) {
                __classPrivateFieldSet(this, _Error_errors, value, "f");
            }
            else {
                __classPrivateFieldSet(this, _Error_errors, value.map((item) => new _a.Error.Errors(item)), "f");
            }
        }
        setErrors(value) {
            this.errors = value;
            return this;
        }
        constructor(data = undefined) {
            _Error_instances.add(this);
            /**
             * Numeric error code representing the failure.
             * @type {number}
             **/
            _Error_code.set(this, 0);
            /**
             * Human-readable explanation of the error.
             * @type {string}
             **/
            _Error_message.set(this, "");
            /**
             * Localized/translated version of the error message.
             * @type {string}
             **/
            _Error_messageTranslated.set(this, "");
            /**
             * Detailed list of error objects.
             * @type {ResponseDto.Error.Errors}
             **/
            _Error_errors.set(this, []);
            if (data === null || data === undefined) {
                return;
            }
            if (typeof data === "string") {
                this.applyFromObject(JSON.parse(data));
            }
            else if (__classPrivateFieldGet(this, _Error_instances, "m", _Error_isJsonAppliable).call(this, data)) {
                this.applyFromObject(data);
            }
            else {
                throw new _c("Instance cannot be created on an unknown value, check the content being passed. got: " +
                    typeof data);
            }
        }
        /**
         * casts the fields of a javascript object into the class properties one by one
         **/
        applyFromObject(data = {}) {
            const d = data;
            if (d.code !== undefined) {
                this.code = d.code;
            }
            if (d.message !== undefined) {
                this.message = d.message;
            }
            if (d.messageTranslated !== undefined) {
                this.messageTranslated = d.messageTranslated;
            }
            if (d.errors !== undefined) {
                this.errors = d.errors;
            }
        }
        /**
         *	Special toJSON override, since the field are private,
         *	Json stringify won't see them unless we mention it explicitly.
         **/
        toJSON() {
            return {
                code: __classPrivateFieldGet(this, _Error_code, "f"),
                message: __classPrivateFieldGet(this, _Error_message, "f"),
                messageTranslated: __classPrivateFieldGet(this, _Error_messageTranslated, "f"),
                errors: __classPrivateFieldGet(this, _Error_errors, "f"),
            };
        }
        toString() {
            return JSON.stringify(this);
        }
        static get Fields() {
            return {
                code: "code",
                message: "message",
                messageTranslated: "messageTranslated",
                errors$: "errors",
                get errors() {
                    return withPrefix("error.errors[:i]", _a.Error.Errors.Fields);
                },
            };
        }
        /**
         * Creates an instance of ResponseDto.Error, and possibleDtoObject
         * needs to satisfy the type requirement fully, otherwise typescript compile would
         * be complaining.
         **/
        static from(possibleDtoObject) {
            return new _a.Error(possibleDtoObject);
        }
        /**
         * Creates an instance of ResponseDto.Error, and partialDtoObject
         * needs to satisfy the type, but partially, and rest of the content would
         * be constructed according to data types and nullability.
         **/
        static with(partialDtoObject) {
            return new _a.Error(partialDtoObject);
        }
        copyWith(partial) {
            return new _a.Error({ ...this.toJSON(), ...partial });
        }
        clone() {
            return new _a.Error(this.toJSON());
        }
    },
    _Error_code = new WeakMap(),
    _Error_message = new WeakMap(),
    _Error_messageTranslated = new WeakMap(),
    _Error_errors = new WeakMap(),
    _Error_instances = new WeakSet(),
    _Error_isJsonAppliable = function _Error_isJsonAppliable(obj) {
        const g = globalThis;
        const isBuffer = typeof g.Buffer !== "undefined" &&
            typeof g.Buffer.isBuffer === "function" &&
            g.Buffer.isBuffer(obj);
        const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
        return (obj &&
            typeof obj === "object" &&
            !Array.isArray(obj) &&
            !isBuffer &&
            !(obj instanceof ArrayBuffer) &&
            !isBlob);
    },
    /**
     * The base class definition for errors
     **/
    _c.Errors = (_d = class Errors {
            /**
             * Logical grouping of the error (e.g., global, usageLimits).
             * @returns {string}
             **/
            get domain() {
                return __classPrivateFieldGet(this, _Errors_domain, "f");
            }
            /**
             * Logical grouping of the error (e.g., global, usageLimits).
             * @type {string}
             **/
            set domain(value) {
                const correctType = typeof value === "string" || value === undefined || value === null;
                __classPrivateFieldSet(this, _Errors_domain, correctType ? value : String(value), "f");
            }
            setDomain(value) {
                this.domain = value;
                return this;
            }
            /**
             * Reason identifier for the error.
             * @returns {string}
             **/
            get reason() {
                return __classPrivateFieldGet(this, _Errors_reason, "f");
            }
            /**
             * Reason identifier for the error.
             * @type {string}
             **/
            set reason(value) {
                const correctType = typeof value === "string" || value === undefined || value === null;
                __classPrivateFieldSet(this, _Errors_reason, correctType ? value : String(value), "f");
            }
            setReason(value) {
                this.reason = value;
                return this;
            }
            /**
             * Human-readable explanation of the sub-error.
             * @returns {string}
             **/
            get message() {
                return __classPrivateFieldGet(this, _Errors_message, "f");
            }
            /**
             * Human-readable explanation of the sub-error.
             * @type {string}
             **/
            set message(value) {
                const correctType = typeof value === "string" || value === undefined || value === null;
                __classPrivateFieldSet(this, _Errors_message, correctType ? value : String(value), "f");
            }
            setMessage(value) {
                this.message = value;
                return this;
            }
            /**
             * Localized/translated version of the sub-error message.
             * @returns {string}
             **/
            get messageTranslated() {
                return __classPrivateFieldGet(this, _Errors_messageTranslated, "f");
            }
            /**
             * Localized/translated version of the sub-error message.
             * @type {string}
             **/
            set messageTranslated(value) {
                const correctType = typeof value === "string" || value === undefined || value === null;
                __classPrivateFieldSet(this, _Errors_messageTranslated, correctType ? value : String(value), "f");
            }
            setMessageTranslated(value) {
                this.messageTranslated = value;
                return this;
            }
            /**
             * Field or parameter in which the error occurred.
             * @returns {string}
             **/
            get location() {
                return __classPrivateFieldGet(this, _Errors_location, "f");
            }
            /**
             * Field or parameter in which the error occurred.
             * @type {string}
             **/
            set location(value) {
                const correctType = typeof value === "string" || value === undefined || value === null;
                __classPrivateFieldSet(this, _Errors_location, correctType ? value : String(value), "f");
            }
            setLocation(value) {
                this.location = value;
                return this;
            }
            /**
             * Type of location (e.g., parameter, header).
             * @returns {string}
             **/
            get locationType() {
                return __classPrivateFieldGet(this, _Errors_locationType, "f");
            }
            /**
             * Type of location (e.g., parameter, header).
             * @type {string}
             **/
            set locationType(value) {
                const correctType = typeof value === "string" || value === undefined || value === null;
                __classPrivateFieldSet(this, _Errors_locationType, correctType ? value : String(value), "f");
            }
            setLocationType(value) {
                this.locationType = value;
                return this;
            }
            /**
             * URL linking to additional documentation about the error.
             * @returns {string}
             **/
            get extendedHelp() {
                return __classPrivateFieldGet(this, _Errors_extendedHelp, "f");
            }
            /**
             * URL linking to additional documentation about the error.
             * @type {string}
             **/
            set extendedHelp(value) {
                const correctType = typeof value === "string" || value === undefined || value === null;
                __classPrivateFieldSet(this, _Errors_extendedHelp, correctType ? value : String(value), "f");
            }
            setExtendedHelp(value) {
                this.extendedHelp = value;
                return this;
            }
            /**
             * URL to submit a report for this error.
             * @returns {string}
             **/
            get sendReport() {
                return __classPrivateFieldGet(this, _Errors_sendReport, "f");
            }
            /**
             * URL to submit a report for this error.
             * @type {string}
             **/
            set sendReport(value) {
                const correctType = typeof value === "string" || value === undefined || value === null;
                __classPrivateFieldSet(this, _Errors_sendReport, correctType ? value : String(value), "f");
            }
            setSendReport(value) {
                this.sendReport = value;
                return this;
            }
            constructor(data = undefined) {
                _Errors_instances.add(this);
                /**
                 * Logical grouping of the error (e.g., global, usageLimits).
                 * @type {string}
                 **/
                _Errors_domain.set(this, undefined);
                /**
                 * Reason identifier for the error.
                 * @type {string}
                 **/
                _Errors_reason.set(this, undefined);
                /**
                 * Human-readable explanation of the sub-error.
                 * @type {string}
                 **/
                _Errors_message.set(this, undefined);
                /**
                 * Localized/translated version of the sub-error message.
                 * @type {string}
                 **/
                _Errors_messageTranslated.set(this, undefined);
                /**
                 * Field or parameter in which the error occurred.
                 * @type {string}
                 **/
                _Errors_location.set(this, undefined);
                /**
                 * Type of location (e.g., parameter, header).
                 * @type {string}
                 **/
                _Errors_locationType.set(this, undefined);
                /**
                 * URL linking to additional documentation about the error.
                 * @type {string}
                 **/
                _Errors_extendedHelp.set(this, undefined);
                /**
                 * URL to submit a report for this error.
                 * @type {string}
                 **/
                _Errors_sendReport.set(this, undefined);
                if (data === null || data === undefined) {
                    return;
                }
                if (typeof data === "string") {
                    this.applyFromObject(JSON.parse(data));
                }
                else if (__classPrivateFieldGet(this, _Errors_instances, "m", _Errors_isJsonAppliable).call(this, data)) {
                    this.applyFromObject(data);
                }
                else {
                    throw new _c("Instance cannot be created on an unknown value, check the content being passed. got: " +
                        typeof data);
                }
            }
            /**
             * casts the fields of a javascript object into the class properties one by one
             **/
            applyFromObject(data = {}) {
                const d = data;
                if (d.domain !== undefined) {
                    this.domain = d.domain;
                }
                if (d.reason !== undefined) {
                    this.reason = d.reason;
                }
                if (d.message !== undefined) {
                    this.message = d.message;
                }
                if (d.messageTranslated !== undefined) {
                    this.messageTranslated = d.messageTranslated;
                }
                if (d.location !== undefined) {
                    this.location = d.location;
                }
                if (d.locationType !== undefined) {
                    this.locationType = d.locationType;
                }
                if (d.extendedHelp !== undefined) {
                    this.extendedHelp = d.extendedHelp;
                }
                if (d.sendReport !== undefined) {
                    this.sendReport = d.sendReport;
                }
            }
            /**
             *	Special toJSON override, since the field are private,
             *	Json stringify won't see them unless we mention it explicitly.
             **/
            toJSON() {
                return {
                    domain: __classPrivateFieldGet(this, _Errors_domain, "f"),
                    reason: __classPrivateFieldGet(this, _Errors_reason, "f"),
                    message: __classPrivateFieldGet(this, _Errors_message, "f"),
                    messageTranslated: __classPrivateFieldGet(this, _Errors_messageTranslated, "f"),
                    location: __classPrivateFieldGet(this, _Errors_location, "f"),
                    locationType: __classPrivateFieldGet(this, _Errors_locationType, "f"),
                    extendedHelp: __classPrivateFieldGet(this, _Errors_extendedHelp, "f"),
                    sendReport: __classPrivateFieldGet(this, _Errors_sendReport, "f"),
                };
            }
            toString() {
                return JSON.stringify(this);
            }
            static get Fields() {
                return {
                    domain: "domain",
                    reason: "reason",
                    message: "message",
                    messageTranslated: "messageTranslated",
                    location: "location",
                    locationType: "locationType",
                    extendedHelp: "extendedHelp",
                    sendReport: "sendReport",
                };
            }
            /**
             * Creates an instance of ResponseDto.Error.Errors, and possibleDtoObject
             * needs to satisfy the type requirement fully, otherwise typescript compile would
             * be complaining.
             **/
            static from(possibleDtoObject) {
                return new _a.Error.Errors(possibleDtoObject);
            }
            /**
             * Creates an instance of ResponseDto.Error.Errors, and partialDtoObject
             * needs to satisfy the type, but partially, and rest of the content would
             * be constructed according to data types and nullability.
             **/
            static with(partialDtoObject) {
                return new _a.Error.Errors(partialDtoObject);
            }
            copyWith(partial) {
                return new _a.Error.Errors({ ...this.toJSON(), ...partial });
            }
            clone() {
                return new _a.Error.Errors(this.toJSON());
            }
        },
        _Errors_domain = new WeakMap(),
        _Errors_reason = new WeakMap(),
        _Errors_message = new WeakMap(),
        _Errors_messageTranslated = new WeakMap(),
        _Errors_location = new WeakMap(),
        _Errors_locationType = new WeakMap(),
        _Errors_extendedHelp = new WeakMap(),
        _Errors_sendReport = new WeakMap(),
        _Errors_instances = new WeakSet(),
        _Errors_isJsonAppliable = function _Errors_isJsonAppliable(obj) {
            const g = globalThis;
            const isBuffer = typeof g.Buffer !== "undefined" &&
                typeof g.Buffer.isBuffer === "function" &&
                g.Buffer.isBuffer(obj);
            const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
            return (obj &&
                typeof obj === "object" &&
                !Array.isArray(obj) &&
                !isBuffer &&
                !(obj instanceof ArrayBuffer) &&
                !isBlob);
        },
        _d),
    _c);
export class ResponseDtoFactory {
}
