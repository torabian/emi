import { CommonVectorResponseDto } from "./CommonVectorResponseDto";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * Action to communicate with the action allData
 */
/**
 * AllDataAction
 */
export class AllDataAction {
  //
  static URL = "/res/22";
  static NewUrl = (qs) => buildUrl(AllDataAction.URL, undefined, qs);
  static Method = "";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? AllDataAction.NewUrl(qs),
      {
        method: AllDataAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new AllDataActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new AllDataActionRes(item));
    const res = await AllDataAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "allData",
    url: "/res/22",
    out: {
      fields: [
        {
          name: "stringType",
          type: "string",
        },
        {
          name: "stringTypeNull",
          type: "string?",
        },
        {
          name: "collection",
          type: "collection",
          target: "CommonVectorResponseDto",
        },
      ],
    },
  };
}
/**
 * The base class definition for allDataActionRes
 **/
export class AllDataActionRes {
  /**
   *
   * @type {string}
   **/
  #stringType = "";
  /**
   *
   * @returns {string}
   **/
  get stringType() {
    return this.#stringType;
  }
  /**
   *
   * @type {string}
   **/
  set stringType(value) {
    this.#stringType = String(value);
  }
  setStringType(value) {
    this.stringType = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #stringTypeNull = undefined;
  /**
   *
   * @returns {string}
   **/
  get stringTypeNull() {
    return this.#stringTypeNull;
  }
  /**
   *
   * @type {string}
   **/
  set stringTypeNull(value) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#stringTypeNull = correctType ? value : String(value);
  }
  setStringTypeNull(value) {
    this.stringTypeNull = value;
    return this;
  }
  /**
   *
   * @type {CommonVectorResponseDto[]}
   **/
  #collection = [];
  /**
   *
   * @returns {CommonVectorResponseDto[]}
   **/
  get collection() {
    return this.#collection;
  }
  /**
   *
   * @type {CommonVectorResponseDto[]}
   **/
  set collection(value) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof CommonVectorResponseDto) {
      this.#collection = value;
    } else {
      this.#collection = value.map((item) => new CommonVectorResponseDto(item));
    }
  }
  setCollection(value) {
    this.collection = value;
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
    if (d.stringType !== undefined) {
      this.stringType = d.stringType;
    }
    if (d.stringTypeNull !== undefined) {
      this.stringTypeNull = d.stringTypeNull;
    }
    if (d.collection !== undefined) {
      this.collection = d.collection;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      stringType: this.#stringType,
      stringTypeNull: this.#stringTypeNull,
      collection: this.#collection,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      stringType: "stringType",
      stringTypeNull: "stringTypeNull",
      collection$: "collection",
      get collection() {
        return withPrefix("collection[:i]", CommonVectorResponseDto.Fields);
      },
    };
  }
  /**
   * Creates an instance of AllDataActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new AllDataActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of AllDataActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new AllDataActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new AllDataActionRes({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new AllDataActionRes(this.toJSON());
  }
}
