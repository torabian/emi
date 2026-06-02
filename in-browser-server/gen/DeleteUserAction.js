import { buildUrl } from "./sdk/common/buildUrl.js";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx.js";
/**
 * Action to communicate with the action deleteUser
 */
/**
 * DeleteUserAction
 */
export class DeleteUserAction {
  //
  static URL = "/users";
  static NewUrl = (qs) => buildUrl(DeleteUserAction.URL, undefined, qs);
  static Method = "delete";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? DeleteUserAction.NewUrl(qs),
      {
        method: DeleteUserAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new DeleteUserActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new DeleteUserActionRes(item));
    const res = await DeleteUserAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "deleteUser",
    url: "/users",
    method: "delete",
    description: "Delete a single user row by id.",
    in: {
      fields: [
        {
          name: "id",
          description: "Id of the user to delete",
          type: "int",
        },
      ],
    },
    out: {
      fields: [
        {
          name: "deleted",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for deleteUserActionReq
 **/
export class DeleteUserActionReq {
  /**
   * Id of the user to delete
   * @type {number}
   **/
  #id = 0;
  /**
   * Id of the user to delete
   * @returns {number}
   **/
  get id() {
    return this.#id;
  }
  /**
   * Id of the user to delete
   * @type {number}
   **/
  set id(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#id = parsedValue;
    }
  }
  setId(value) {
    this.id = value;
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
    if (d.id !== undefined) {
      this.id = d.id;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      id: this.#id,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      id: "id",
    };
  }
  /**
   * Creates an instance of DeleteUserActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new DeleteUserActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of DeleteUserActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new DeleteUserActionReq(partialDtoObject);
  }
  copyWith(partial) {
    return new DeleteUserActionReq({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new DeleteUserActionReq(this.toJSON());
  }
}
/**
 * The base class definition for deleteUserActionRes
 **/
export class DeleteUserActionRes {
  /**
   *
   * @type {boolean}
   **/
  #deleted;
  /**
   *
   * @returns {boolean}
   **/
  get deleted() {
    return this.#deleted;
  }
  /**
   *
   * @type {boolean}
   **/
  set deleted(value) {
    this.#deleted = Boolean(value);
  }
  setDeleted(value) {
    this.deleted = value;
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
    if (d.deleted !== undefined) {
      this.deleted = d.deleted;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      deleted: this.#deleted,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      deleted: "deleted",
    };
  }
  /**
   * Creates an instance of DeleteUserActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new DeleteUserActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of DeleteUserActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new DeleteUserActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new DeleteUserActionRes({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new DeleteUserActionRes(this.toJSON());
  }
}
