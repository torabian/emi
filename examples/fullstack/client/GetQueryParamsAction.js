import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action getQueryParams
 */
/**
 * GetQueryParamsAction
 */
export class GetQueryParamsAction {
  //
  static URL = "/stream/:addres1/:addressName2 string/:count int";
  static NewUrl = (params, qs) =>
    buildUrl(GetQueryParamsAction.URL, params, qs);
  static Method = "get";
  static Fetch$ = async (params, qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? GetQueryParamsAction.NewUrl(params, qs),
      {
        method: GetQueryParamsAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params,
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new GetQueryParamsActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new GetQueryParamsActionRes(item));
    const res = await GetQueryParamsAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "getQueryParams",
    url: "/stream/:addres1/:addressName2 string/:count int",
    method: "get",
    description:
      "The goal is to check that if query params also become available in the cli params",
    out: {
      fields: [
        {
          name: "name",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for getQueryParamsActionRes
 **/
export class GetQueryParamsActionRes {
  /**
   *
   * @type {string}
   **/
  #name = "";
  /**
   *
   * @returns {string}
   **/
  get name() {
    return this.#name;
  }
  /**
   *
   * @type {string}
   **/
  set name(value) {
    this.#name = String(value);
  }
  setName(value) {
    this.name = value;
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
    if (d.name !== undefined) {
      this.name = d.name;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      name: this.#name,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      name: "name",
    };
  }
  /**
   * Creates an instance of GetQueryParamsActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new GetQueryParamsActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of GetQueryParamsActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new GetQueryParamsActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new GetQueryParamsActionRes({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new GetQueryParamsActionRes(this.toJSON());
  }
}
