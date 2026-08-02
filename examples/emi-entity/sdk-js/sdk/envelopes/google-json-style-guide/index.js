import { ResponseDto } from "./generated/ResponseDto";
// Use this class to generate a GResponse.
export class GResponse extends ResponseDto {
    constructor(data) {
        super(data);
    }
    setCreator(fn) {
        this.creator = fn;
        return this;
    }
    /**
     * GResponse can have data.item or data.items
     * We create that based on incoming data tpye, so there is no need for 2 different
     * classes, one for array and other for singular
     * @param data
     * @returns
     */
    inject(body) {
        var _a, _b, _c, _d, _e, _f, _g, _h;
        this.applyFromObject(body);
        if (body === null || body === void 0 ? void 0 : body.data) {
            if (!this.data) {
                this.setData({});
            }
            if (Array.isArray(body === null || body === void 0 ? void 0 : body.data.items) &&
                typeof this.creator !== "undefined" &&
                this.creator !== null) {
                (_a = this.data) === null || _a === void 0 ? void 0 : _a.setItems((_c = (_b = body === null || body === void 0 ? void 0 : body.data) === null || _b === void 0 ? void 0 : _b.items) === null || _c === void 0 ? void 0 : _c.map((item) => { var _a; return (_a = this.creator) === null || _a === void 0 ? void 0 : _a.call(this, item); }));
            }
            else if (typeof ((_d = body === null || body === void 0 ? void 0 : body.data) === null || _d === void 0 ? void 0 : _d.item) === "object" &&
                typeof this.creator !== "undefined" &&
                this.creator !== null) {
                (_e = this.data) === null || _e === void 0 ? void 0 : _e.setItem(this.creator((_f = body === null || body === void 0 ? void 0 : body.data) === null || _f === void 0 ? void 0 : _f.item));
            }
            else {
                (_g = this.data) === null || _g === void 0 ? void 0 : _g.setItem((_h = body === null || body === void 0 ? void 0 : body.data) === null || _h === void 0 ? void 0 : _h.item);
            }
        }
        return this;
    }
}
