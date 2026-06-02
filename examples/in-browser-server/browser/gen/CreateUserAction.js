import { buildUrl } from "./sdk/common/buildUrl.js";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx.js";
/**
 * Action to communicate with the action createUser
 */
/**
 * CreateUserAction
 */
export class CreateUserAction {
  //
  static URL = "/users";
  static NewUrl = (qs) => buildUrl(CreateUserAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? CreateUserAction.NewUrl(qs),
      {
        method: CreateUserAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new CreateUserActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new CreateUserActionRes(item));
    const res = await CreateUserAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "createUser",
    url: "/users",
    method: "post",
    description: "Insert a new user row and return the created record.",
    in: {
      fields: [
        {
          name: "firstName",
          description: "User's first name",
          type: "string",
        },
        {
          name: "lastName",
          description: "User's last name",
          type: "string",
        },
        {
          name: "birthDate",
          description: "User's birth date, ISO format YYYY-MM-DD",
          type: "string",
        },
      ],
    },
    out: {
      fields: [
        {
          name: "id",
          type: "int",
        },
        {
          name: "firstName",
          type: "string",
        },
        {
          name: "lastName",
          type: "string",
        },
        {
          name: "birthDate",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for createUserActionReq
 **/
export class CreateUserActionReq {
  /**
   * User's first name
   * @type {string}
   **/
  #firstName = "";
  /**
   * User's first name
   * @returns {string}
   **/
  get firstName() {
    return this.#firstName;
  }
  /**
   * User's first name
   * @type {string}
   **/
  set firstName(value) {
    this.#firstName = String(value);
  }
  setFirstName(value) {
    this.firstName = value;
    return this;
  }
  /**
   * User's last name
   * @type {string}
   **/
  #lastName = "";
  /**
   * User's last name
   * @returns {string}
   **/
  get lastName() {
    return this.#lastName;
  }
  /**
   * User's last name
   * @type {string}
   **/
  set lastName(value) {
    this.#lastName = String(value);
  }
  setLastName(value) {
    this.lastName = value;
    return this;
  }
  /**
   * User's birth date, ISO format YYYY-MM-DD
   * @type {string}
   **/
  #birthDate = "";
  /**
   * User's birth date, ISO format YYYY-MM-DD
   * @returns {string}
   **/
  get birthDate() {
    return this.#birthDate;
  }
  /**
   * User's birth date, ISO format YYYY-MM-DD
   * @type {string}
   **/
  set birthDate(value) {
    this.#birthDate = String(value);
  }
  setBirthDate(value) {
    this.birthDate = value;
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
    if (d.firstName !== undefined) {
      this.firstName = d.firstName;
    }
    if (d.lastName !== undefined) {
      this.lastName = d.lastName;
    }
    if (d.birthDate !== undefined) {
      this.birthDate = d.birthDate;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      firstName: this.#firstName,
      lastName: this.#lastName,
      birthDate: this.#birthDate,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      firstName: "firstName",
      lastName: "lastName",
      birthDate: "birthDate",
    };
  }
  /**
   * Creates an instance of CreateUserActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new CreateUserActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of CreateUserActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new CreateUserActionReq(partialDtoObject);
  }
  copyWith(partial) {
    return new CreateUserActionReq({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new CreateUserActionReq(this.toJSON());
  }
}
/**
 * The base class definition for createUserActionRes
 **/
export class CreateUserActionRes {
  /**
   *
   * @type {number}
   **/
  #id = 0;
  /**
   *
   * @returns {number}
   **/
  get id() {
    return this.#id;
  }
  /**
   *
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
  /**
   *
   * @type {string}
   **/
  #firstName = "";
  /**
   *
   * @returns {string}
   **/
  get firstName() {
    return this.#firstName;
  }
  /**
   *
   * @type {string}
   **/
  set firstName(value) {
    this.#firstName = String(value);
  }
  setFirstName(value) {
    this.firstName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #lastName = "";
  /**
   *
   * @returns {string}
   **/
  get lastName() {
    return this.#lastName;
  }
  /**
   *
   * @type {string}
   **/
  set lastName(value) {
    this.#lastName = String(value);
  }
  setLastName(value) {
    this.lastName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #birthDate = "";
  /**
   *
   * @returns {string}
   **/
  get birthDate() {
    return this.#birthDate;
  }
  /**
   *
   * @type {string}
   **/
  set birthDate(value) {
    this.#birthDate = String(value);
  }
  setBirthDate(value) {
    this.birthDate = value;
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
    if (d.firstName !== undefined) {
      this.firstName = d.firstName;
    }
    if (d.lastName !== undefined) {
      this.lastName = d.lastName;
    }
    if (d.birthDate !== undefined) {
      this.birthDate = d.birthDate;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      id: this.#id,
      firstName: this.#firstName,
      lastName: this.#lastName,
      birthDate: this.#birthDate,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      id: "id",
      firstName: "firstName",
      lastName: "lastName",
      birthDate: "birthDate",
    };
  }
  /**
   * Creates an instance of CreateUserActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new CreateUserActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of CreateUserActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new CreateUserActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new CreateUserActionRes({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new CreateUserActionRes(this.toJSON());
  }
}
