import { WebSocketX } from "./sdk/common/WebSocketX";
import { buildUrl } from "./sdk/common/buildUrl";
/**
 * Action to communicate with the action reactiveEcho
 */
/**
 * ReactiveEchoAction
 */
export class ReactiveEchoAction {
  //
  static URL = "/reactive/echo/:channel";
  static NewUrl = (params, qs) => buildUrl(ReactiveEchoAction.URL, params, qs);
  static Method = "REACTIVE";
  static Create = (overrideUrl, qs, params, options) => {
    const url = overrideUrl ?? ReactiveEchoAction.NewUrl(params, qs);
    const Cls = options?.SocketClass ? options.SocketClass : WebSocketX;
    return new Cls(url, undefined, {
      MessageFactoryClass: ReactiveEchoActionRes,
    });
  };
  static Definition = {
    name: "reactiveEcho",
    url: "/reactive/echo/:channel string",
    method: "reactive",
    description: "Echoes messages back over a websocket, scoped to a channel.",
    in: {
      fields: [
        {
          name: "message",
          type: "string",
        },
      ],
    },
    out: {
      fields: [
        {
          name: "message",
          type: "string",
        },
        {
          name: "echoedAt",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for reactiveEchoActionReq
 **/
export class ReactiveEchoActionReq {
  /**
   *
   * @type {string}
   **/
  #message = "";
  /**
   *
   * @returns {string}
   **/
  get message() {
    return this.#message;
  }
  /**
   *
   * @type {string}
   **/
  set message(value) {
    this.#message = String(value);
  }
  setMessage(value) {
    this.message = value;
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
    if (d.message !== undefined) {
      this.message = d.message;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      message: this.#message,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      message: "message",
    };
  }
  /**
   * Creates an instance of ReactiveEchoActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new ReactiveEchoActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of ReactiveEchoActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new ReactiveEchoActionReq(partialDtoObject);
  }
  copyWith(partial) {
    return new ReactiveEchoActionReq({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new ReactiveEchoActionReq(this.toJSON());
  }
}
/**
 * The base class definition for reactiveEchoActionRes
 **/
export class ReactiveEchoActionRes {
  /**
   *
   * @type {string}
   **/
  #message = "";
  /**
   *
   * @returns {string}
   **/
  get message() {
    return this.#message;
  }
  /**
   *
   * @type {string}
   **/
  set message(value) {
    this.#message = String(value);
  }
  setMessage(value) {
    this.message = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #echoedAt = "";
  /**
   *
   * @returns {string}
   **/
  get echoedAt() {
    return this.#echoedAt;
  }
  /**
   *
   * @type {string}
   **/
  set echoedAt(value) {
    this.#echoedAt = String(value);
  }
  setEchoedAt(value) {
    this.echoedAt = value;
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
    if (d.message !== undefined) {
      this.message = d.message;
    }
    if (d.echoedAt !== undefined) {
      this.echoedAt = d.echoedAt;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      message: this.#message,
      echoedAt: this.#echoedAt,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      message: "message",
      echoedAt: "echoedAt",
    };
  }
  /**
   * Creates an instance of ReactiveEchoActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new ReactiveEchoActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ReactiveEchoActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new ReactiveEchoActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new ReactiveEchoActionRes({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new ReactiveEchoActionRes(this.toJSON());
  }
}
