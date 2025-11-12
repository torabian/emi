import {
  FetchxContext,
  fetchx,
  handleFetchResponse,
  type TypedRequestInit,
} from "./sdk/common/fetchx";
import { buildUrl } from "./sdk/common/buildUrl";
/**
 * Action to communicate with the action httpAction
 */
export type HttpActionActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
/**
 * HttpActionAction
 */
export class HttpActionAction {
  //
  static URL = "http://localhost:8081 (for test we use override)";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(HttpActionAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<HttpActionActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<HttpActionActionRes, HttpActionActionReq, unknown>(
      overrideUrl ?? HttpActionAction.NewUrl(qs),
      {
        method: HttpActionAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<HttpActionActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => HttpActionActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new HttpActionActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new HttpActionActionRes(item));
    const res = await HttpActionAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => creatorFn(item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "httpAction",
    url: "http://localhost:8081 (for test we use override)",
    method: "post",
    description:
      "A post connection which would generate random numbers, based on min, max, and count. For normal response, it would return an object, for SSE would be streaming.",
    in: {
      fields: [
        {
          name: "min",
          description: "Minimum number which can be generated",
          type: "int",
        },
        {
          name: "max",
          description: "Maximum number which can be generated",
          type: "int",
        },
        {
          name: "count",
          description:
            "How many numbers you want to be generated based on maximum and minimum",
          type: "int",
        },
      ],
    },
    out: {
      fields: [
        {
          name: "number",
          type: "int",
        },
      ],
    },
  };
}
/**
 * The base class definition for httpActionActionReq
 **/
export class HttpActionActionReq {
  /**
   * Minimum number which can be generated
   * @type {number}
   **/
  #min: number = 0;
  /**
   * Minimum number which can be generated
   * @returns {number}
   **/
  get min() {
    return this.#min;
  }
  /**
   * Minimum number which can be generated
   * @type {number}
   **/
  set min(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#min = parsedValue;
    }
  }
  setMin(value: number) {
    this.min = value;
    return this;
  }
  /**
   * Maximum number which can be generated
   * @type {number}
   **/
  #max: number = 0;
  /**
   * Maximum number which can be generated
   * @returns {number}
   **/
  get max() {
    return this.#max;
  }
  /**
   * Maximum number which can be generated
   * @type {number}
   **/
  set max(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#max = parsedValue;
    }
  }
  setMax(value: number) {
    this.max = value;
    return this;
  }
  /**
   * How many numbers you want to be generated based on maximum and minimum
   * @type {number}
   **/
  #count: number = 0;
  /**
   * How many numbers you want to be generated based on maximum and minimum
   * @returns {number}
   **/
  get count() {
    return this.#count;
  }
  /**
   * How many numbers you want to be generated based on maximum and minimum
   * @type {number}
   **/
  set count(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#count = parsedValue;
    }
  }
  setCount(value: number) {
    this.count = value;
    return this;
  }
  constructor(data: unknown = undefined) {
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
  #isJsonAppliable(obj: unknown) {
    const g = globalThis as unknown as { Buffer: any; Blob: any };
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
    const d = data as Partial<HttpActionActionReq>;
    if (d.min !== undefined) {
      this.min = d.min;
    }
    if (d.max !== undefined) {
      this.max = d.max;
    }
    if (d.count !== undefined) {
      this.count = d.count;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      min: this.#min,
      max: this.#max,
      count: this.#count,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      min: "min",
      max: "max",
      count: "count",
    };
  }
  /**
   * Creates an instance of HttpActionActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: HttpActionActionReqType) {
    return new HttpActionActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of HttpActionActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<HttpActionActionReqType>) {
    return new HttpActionActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<HttpActionActionReqType>,
  ): InstanceType<typeof HttpActionActionReq> {
    return new HttpActionActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof HttpActionActionReq> {
    return new HttpActionActionReq(this.toJSON());
  }
}
export abstract class HttpActionActionReqFactory {
  abstract create(data: unknown): HttpActionActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for httpActionActionReq
 **/
export type HttpActionActionReqType = {
  /**
   * Minimum number which can be generated
   * @type {number}
   **/
  min: number;
  /**
   * Maximum number which can be generated
   * @type {number}
   **/
  max: number;
  /**
   * How many numbers you want to be generated based on maximum and minimum
   * @type {number}
   **/
  count: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace HttpActionActionReqType {}
/**
 * The base class definition for httpActionActionRes
 **/
export class HttpActionActionRes {
  /**
   *
   * @type {number}
   **/
  #number: number = 0;
  /**
   *
   * @returns {number}
   **/
  get number() {
    return this.#number;
  }
  /**
   *
   * @type {number}
   **/
  set number(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#number = parsedValue;
    }
  }
  setNumber(value: number) {
    this.number = value;
    return this;
  }
  constructor(data: unknown = undefined) {
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
  #isJsonAppliable(obj: unknown) {
    const g = globalThis as unknown as { Buffer: any; Blob: any };
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
    const d = data as Partial<HttpActionActionRes>;
    if (d.number !== undefined) {
      this.number = d.number;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      number: this.#number,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      number: "number",
    };
  }
  /**
   * Creates an instance of HttpActionActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: HttpActionActionResType) {
    return new HttpActionActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of HttpActionActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<HttpActionActionResType>) {
    return new HttpActionActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<HttpActionActionResType>,
  ): InstanceType<typeof HttpActionActionRes> {
    return new HttpActionActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof HttpActionActionRes> {
    return new HttpActionActionRes(this.toJSON());
  }
}
export abstract class HttpActionActionResFactory {
  abstract create(data: unknown): HttpActionActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for httpActionActionRes
 **/
export type HttpActionActionResType = {
  /**
   *
   * @type {number}
   **/
  number: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace HttpActionActionResType {}
