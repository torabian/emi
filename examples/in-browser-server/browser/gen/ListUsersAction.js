import {
  MArray,
  MArrayNullable,
  MCollection,
  MCollectionNullable,
  MOne,
  MOneNullable,
} from "./sdk/common/operators.js";
import { buildUrl } from "./sdk/common/buildUrl.js";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx.js";
import { withPrefix } from "./sdk/common/withPrefix.js";
/**
 * Action to communicate with the action listUsers
 */
/**
 * ListUsersAction
 */
export class ListUsersAction {
  //
  static URL = "/users";
  static NewUrl = (qs) => buildUrl(ListUsersAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? ListUsersAction.NewUrl(qs),
      {
        method: ListUsersAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new ListUsersActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new ListUsersActionRes(item));
    const res = await ListUsersAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "listUsers",
    url: "/users",
    method: "get",
    description: "Return every user row, ordered by id.",
    out: {
      fields: [
        {
          name: "users",
          type: "array",
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
      ],
    },
  };
}
/**
 * The base class definition for listUsersActionRes
 **/
export class ListUsersActionRes {
  /**
   *
   * @type {ListUsersActionRes.Users}
   **/
  #users = [];
  /**
   *
   * @returns {ListUsersActionRes.Users}
   **/
  get users() {
    return this.#users;
  }
  /**
   *
   * @type {ListUsersActionRes.Users}
   **/
  set users(value) {
    if (!Array.isArray(value) && !(value instanceof MCollection)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof ListUsersActionRes.Users) {
      this.#users = value;
    } else {
      this.#users = value.map((item) => new ListUsersActionRes.Users(item));
    }
  }
  setUsers(value) {
    this.users = value;
    return this;
  }
  /**
   * The base class definition for users
   **/
  static Users = class Users {
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
     * Creates an instance of ListUsersActionRes.Users, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new ListUsersActionRes.Users(possibleDtoObject);
    }
    /**
     * Creates an instance of ListUsersActionRes.Users, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new ListUsersActionRes.Users(partialDtoObject);
    }
    copyWith(partial) {
      return new ListUsersActionRes.Users({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new ListUsersActionRes.Users(this.toJSON());
    }
  };
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
    if (d.users !== undefined) {
      this.users = d.users;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      users: this.#users,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      users$: "users",
      get users() {
        return withPrefix("users[:i]", ListUsersActionRes.Users.Fields);
      },
    };
  }
  /**
   * Creates an instance of ListUsersActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new ListUsersActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ListUsersActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new ListUsersActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new ListUsersActionRes({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new ListUsersActionRes(this.toJSON());
  }
}
