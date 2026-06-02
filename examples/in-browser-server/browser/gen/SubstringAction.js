import { buildUrl } from "./sdk/common/buildUrl.js";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx.js";
/**
 * Action to communicate with the action substring
 */
/**
 * SubstringAction
 */
export class SubstringAction {
  //
  static URL = "/";
  static NewUrl = (qs) => buildUrl(SubstringAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? SubstringAction.NewUrl(qs),
      {
        method: SubstringAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new SubstringActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new SubstringActionRes(item));
    const res = await SubstringAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "substring",
    url: "/",
    method: "post",
    in: {
      fields: [
        {
          name: "input",
          description: "The string you want to do substring",
          type: "string",
        },
        {
          name: "start",
          description: "Start position",
          type: "int",
        },
        {
          name: "end",
          description: "End position",
          type: "int",
        },
      ],
    },
    out: {
      fields: [
        {
          name: "result",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for substringActionReq
 **/
export class SubstringActionReq {
  /**
   * The string you want to do substring
   * @type {string}
   **/
  #input = "";
  /**
   * The string you want to do substring
   * @returns {string}
   **/
  get input() {
    return this.#input;
  }
  /**
   * The string you want to do substring
   * @type {string}
   **/
  set input(value) {
    this.#input = String(value);
  }
  setInput(value) {
    this.input = value;
    return this;
  }
  /**
   * Start position
   * @type {number}
   **/
  #start = 0;
  /**
   * Start position
   * @returns {number}
   **/
  get start() {
    return this.#start;
  }
  /**
   * Start position
   * @type {number}
   **/
  set start(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#start = parsedValue;
    }
  }
  setStart(value) {
    this.start = value;
    return this;
  }
  /**
   * End position
   * @type {number}
   **/
  #end = 0;
  /**
   * End position
   * @returns {number}
   **/
  get end() {
    return this.#end;
  }
  /**
   * End position
   * @type {number}
   **/
  set end(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#end = parsedValue;
    }
  }
  setEnd(value) {
    this.end = value;
    return this;
  }
  constructor(data) {
    if (data === null || data === undefined) {
      return;
    }
    if (typeof data === "string") {
      this.applyFromObject(JSON.parse(data));
    } else if (this.#isJsonAppliable(data)) {
      this.applyFromObject(data);
    } else {
      throw new Error(
        "Instance cannot be created on an unknown value, check the content being passed. got: " +
          typeof data,
      );
    }
  }
  #isJsonAppliable(obj) {
    const g = globalThis;
    const isBuffer =
      typeof g.Buffer !== "undefined" &&
      typeof g.Buffer.isBuffer === "function" &&
      g.Buffer.isBuffer(obj);
    const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
    return (
      obj &&
      typeof obj === "object" &&
      !Array.isArray(obj) &&
      !isBuffer &&
      !(obj instanceof ArrayBuffer) &&
      !isBlob
    );
  }
  /**
   * casts the fields of a javascript object into the class properties one by one
   **/
  applyFromObject(data = {}) {
    const d = data;
    if (d.input !== undefined) {
      this.input = d.input;
    }
    if (d.start !== undefined) {
      this.start = d.start;
    }
    if (d.end !== undefined) {
      this.end = d.end;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      input: this.#input,
      start: this.#start,
      end: this.#end,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      input: "input",
      start: "start",
      end: "end",
    };
  }
  /**
   * Creates an instance of SubstringActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new SubstringActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of SubstringActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new SubstringActionReq(partialDtoObject);
  }
  copyWith(partial) {
    return new SubstringActionReq({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new SubstringActionReq(this.toJSON());
  }
}
/**
 * The base class definition for substringActionRes
 **/
export class SubstringActionRes {
  /**
   *
   * @type {string}
   **/
  #result = "";
  /**
   *
   * @returns {string}
   **/
  get result() {
    return this.#result;
  }
  /**
   *
   * @type {string}
   **/
  set result(value) {
    this.#result = String(value);
  }
  setResult(value) {
    this.result = value;
    return this;
  }
  constructor(data) {
    if (data === null || data === undefined) {
      return;
    }
    if (typeof data === "string") {
      this.applyFromObject(JSON.parse(data));
    } else if (this.#isJsonAppliable(data)) {
      this.applyFromObject(data);
    } else {
      throw new Error(
        "Instance cannot be created on an unknown value, check the content being passed. got: " +
          typeof data,
      );
    }
  }
  #isJsonAppliable(obj) {
    const g = globalThis;
    const isBuffer =
      typeof g.Buffer !== "undefined" &&
      typeof g.Buffer.isBuffer === "function" &&
      g.Buffer.isBuffer(obj);
    const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
    return (
      obj &&
      typeof obj === "object" &&
      !Array.isArray(obj) &&
      !isBuffer &&
      !(obj instanceof ArrayBuffer) &&
      !isBlob
    );
  }
  /**
   * casts the fields of a javascript object into the class properties one by one
   **/
  applyFromObject(data = {}) {
    const d = data;
    if (d.result !== undefined) {
      this.result = d.result;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      result: this.#result,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      result: "result",
    };
  }
  /**
   * Creates an instance of SubstringActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new SubstringActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of SubstringActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new SubstringActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new SubstringActionRes({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new SubstringActionRes(this.toJSON());
  }
}
