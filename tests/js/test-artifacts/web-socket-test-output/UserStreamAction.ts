import { WebSocketX } from "./sdk/common/WebSocketX";
import { buildUrl } from "./sdk/common/buildUrl";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * Action to communicate with the action userStream
 */
export type UserStreamActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
/**
 * UserStreamAction
 */
export class UserStreamAction {
  static URL = "ws://localhost:8081";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(UserStreamAction.URL, undefined, qs);
  static Method = "reactive";
  static Create = (overrideUrl?: string, qs?: URLSearchParams) => {
    const url = overrideUrl ?? UserStreamAction.NewUrl(qs);
    return new WebSocketX<UserStreamActionReq, UserStreamActionRes>(
      url,
      undefined,
      {
        MessageFactoryClass: UserStreamActionRes,
      },
    );
  };
}
/**
 * The base class definition for userStreamActionReq
 **/
export class UserStreamActionReq {
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
    const isBuffer =
      typeof globalThis.Buffer !== "undefined" &&
      typeof globalThis.Buffer.isBuffer === "function" &&
      globalThis.Buffer.isBuffer(obj);
    const isBlob =
      typeof globalThis.Blob !== "undefined" && obj instanceof globalThis.Blob;
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
    const d = data as Partial<UserStreamActionReq>;
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
}
export abstract class UserStreamActionReqFactory {
  abstract create(data: unknown): UserStreamActionReq;
}
/**
 * The base type definition for userStreamActionReq
 **/
export type UserStreamActionReqType = {
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
export namespace UserStreamActionReqType {}
/**
 * The base class definition for userStreamActionRes
 **/
export class UserStreamActionRes {
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
    const isBuffer =
      typeof globalThis.Buffer !== "undefined" &&
      typeof globalThis.Buffer.isBuffer === "function" &&
      globalThis.Buffer.isBuffer(obj);
    const isBlob =
      typeof globalThis.Blob !== "undefined" && obj instanceof globalThis.Blob;
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
    const d = data as Partial<UserStreamActionRes>;
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
}
export abstract class UserStreamActionResFactory {
  abstract create(data: unknown): UserStreamActionRes;
}
/**
 * The base type definition for userStreamActionRes
 **/
export type UserStreamActionResType = {
  /**
   *
   * @type {number}
   **/
  number: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UserStreamActionResType {}
