import { MArray } from "./sdk/common/operators";
import { WebSocketX } from "./sdk/common/WebSocketX";
import { buildUrl } from "./sdk/common/buildUrl";
import { type PartialDeep } from "./sdk/common/fetchx";
import { useWebSocketX } from "./sdk/react/useWebSocketX";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * Action to communicate with the action webSocketOrgEcho
 */
export type WebSocketOrgEchoActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export const useWebSocketOrgEchoAction = (options?: {
  qs?: URLSearchParams;
  overrideUrl?: string;
}) => {
  return useWebSocketX(() =>
    WebSocketOrgEchoAction.Create(options?.overrideUrl, options?.qs),
  );
};
/**
 * WebSocketOrgEchoAction
 */
export class WebSocketOrgEchoAction {
  //
  static URL = "wss://echo.websocket.org/.ws";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WebSocketOrgEchoAction.URL, undefined, qs);
  static Method = "REACTIVE";
  static Create = (overrideUrl?: string, qs?: URLSearchParams, options) => {
    const url = overrideUrl ?? WebSocketOrgEchoAction.NewUrl(qs);
    const Cls = options?.SocketClass
      ? options.SocketClass
      : WebSocketX<WebSocketOrgEchoActionReq, WebSocketOrgEchoActionRes>;
    return new Cls(url, undefined, {
      MessageFactoryClass: WebSocketOrgEchoActionRes,
    });
  };
  static Definition = {
    name: "webSocketOrgEcho",
    url: "wss://echo.websocket.org/.ws",
    method: "reactive",
    description: "Websocket.org eco server, to send a json and recieve back",
    in: {
      fields: [
        {
          name: "firstName",
          type: "string",
        },
        {
          name: "lastName",
          type: "string",
        },
        {
          name: "user",
          type: "object",
          fields: [
            {
              name: "item1",
              type: "string",
            },
            {
              name: "item2array",
              type: "array",
              fields: [
                {
                  name: "subItem1",
                  type: "int64",
                },
                {
                  name: "subItem2",
                  type: "int64",
                },
              ],
            },
          ],
        },
      ],
    },
    out: {
      fields: [
        {
          name: "lastName",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for webSocketOrgEchoActionReq
 **/
export class WebSocketOrgEchoActionReq {
  /**
   *
   * @type {string}
   **/
  #firstName: string = "";
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
  set firstName(value: string) {
    this.#firstName = String(value);
  }
  setFirstName(value: string) {
    this.firstName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #lastName: string = "";
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
  set lastName(value: string) {
    this.#lastName = String(value);
  }
  setLastName(value: string) {
    this.lastName = value;
    return this;
  }
  /**
   *
   * @type {WebSocketOrgEchoActionReq.User}
   **/
  #user!: InstanceType<typeof WebSocketOrgEchoActionReq.User>;
  /**
   *
   * @returns {WebSocketOrgEchoActionReq.User}
   **/
  get user() {
    return this.#user;
  }
  /**
   *
   * @type {WebSocketOrgEchoActionReq.User}
   **/
  set user(value: InstanceType<typeof WebSocketOrgEchoActionReq.User>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof WebSocketOrgEchoActionReq.User) {
      this.#user = value;
    } else {
      this.#user = new WebSocketOrgEchoActionReq.User(value);
    }
  }
  setUser(value: InstanceType<typeof WebSocketOrgEchoActionReq.User>) {
    this.user = value;
    return this;
  }
  /**
   * The base class definition for user
   **/
  static User = class User {
    /**
     *
     * @type {string}
     **/
    #item1: string = "";
    /**
     *
     * @returns {string}
     **/
    get item1() {
      return this.#item1;
    }
    /**
     *
     * @type {string}
     **/
    set item1(value: string) {
      this.#item1 = String(value);
    }
    setItem1(value: string) {
      this.item1 = value;
      return this;
    }
    /**
     *
     * @type {WebSocketOrgEchoActionReq.User.Item2array}
     **/
    #item2array: MArray<
      InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array>
    > = MArray.of([]);
    /**
     *
     * @returns {WebSocketOrgEchoActionReq.User.Item2array}
     **/
    get item2array() {
      return this.#item2array;
    }
    /**
     *
     * @type {WebSocketOrgEchoActionReq.User.Item2array}
     **/
    set item2array(
      value:
        | MArray<InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array>>
        | InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array>[],
    ) {
      // When the passed value is already an array, we check if we need to
      // cast the inner items into class instance.
      if (Array.isArray(value)) {
        if (
          value.length > 0 &&
          value[0] instanceof WebSocketOrgEchoActionReq.User.Item2array
        ) {
          this.#item2array = MArray.of(value);
        } else {
          this.#item2array = MArray.of(
            value.map(
              (item) => new WebSocketOrgEchoActionReq.User.Item2array(item),
            ),
          );
        }
        return;
      }
      // If the instance is already an MArray, we assume it's all good.
      if (value instanceof MArray) {
        this.#item2array = value;
        return;
      }
      // If the value is not array, and is not a MArray, we need to be consider,
      // it might be eligible to be casted into MArray.
      const { ok, value: mcastValue } = MArray.cast<unknown>(value);
      if (ok) {
        this.#item2array = mcastValue as any;
        return;
      }
      console.warn(
        "Cannot assing value to item2array, because it needs MArray instance or an Array.",
      );
    }
    setItem2array(
      value:
        | MArray<InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array>>
        | InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array>[],
    ) {
      this.item2array = value;
      return this;
    }
    /**
     * The base class definition for item2array
     **/
    static Item2array = class Item2array {
      /**
       *
       * @type {number}
       **/
      #subItem1: number = 0;
      /**
       *
       * @returns {number}
       **/
      get subItem1() {
        return this.#subItem1;
      }
      /**
       *
       * @type {number}
       **/
      set subItem1(value: number) {
        const correctType = typeof value === "number";
        const parsedValue = correctType ? value : Number(value);
        if (!Number.isNaN(parsedValue)) {
          this.#subItem1 = parsedValue;
        }
      }
      setSubItem1(value: number) {
        this.subItem1 = value;
        return this;
      }
      /**
       *
       * @type {number}
       **/
      #subItem2: number = 0;
      /**
       *
       * @returns {number}
       **/
      get subItem2() {
        return this.#subItem2;
      }
      /**
       *
       * @type {number}
       **/
      set subItem2(value: number) {
        const correctType = typeof value === "number";
        const parsedValue = correctType ? value : Number(value);
        if (!Number.isNaN(parsedValue)) {
          this.#subItem2 = parsedValue;
        }
      }
      setSubItem2(value: number) {
        this.subItem2 = value;
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
        const d = data as Partial<Item2array>;
        if (d.subItem1 !== undefined) {
          this.subItem1 = d.subItem1;
        }
        if (d.subItem2 !== undefined) {
          this.subItem2 = d.subItem2;
        }
      }
      /**
       *	Special toJSON override, since the field are private,
       *	Json stringify won't see them unless we mention it explicitly. Each
       *	field goes through #toPlainJSON rather than a bare this.#field: a
       *	nested dto instance, or an emigo Array/One/Collection wrapper, has its
       *	own toJSON that JSON.stringify would only reach on a *second* pass -
       *	calling it here means toJSON()'s own return value is already the same
       *	plain shape a consumer choosing the exported type instead of this
       *	class gets, rather than a shallow object still holding class instances.
       **/
      toJSON(): WebSocketOrgEchoActionReqType.UserType.Item2arrayType {
        return {
          subItem1: this.#toPlainJSON(this.#subItem1),
          subItem2: this.#toPlainJSON(this.#subItem2),
        } as WebSocketOrgEchoActionReqType.UserType.Item2arrayType;
      }
      /**
       * Resolves value into a plain, JSON-serializable value: recurses into
       * arrays, and - the case JSON.stringify's own recursion would only get to
       * after this method already returned - calls a nested value's own
       * toJSON() (a dto instance, or an emigo Array/One/Collection wrapper)
       * rather than leaving it as-is. A plain value (string, number, a map's
       * own plain object, ...) is returned unchanged.
       **/
      #toPlainJSON(value: unknown): unknown {
        if (value === null || value === undefined) {
          return value;
        }
        if (Array.isArray(value)) {
          return value.map((item) => this.#toPlainJSON(item));
        }
        const asAny = value as any;
        if (typeof asAny.toJSON === "function") {
          return this.#toPlainJSON(asAny.toJSON());
        }
        return value;
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          subItem1: "subItem1",
          subItem2: "subItem2",
        };
      }
      /**
       * Creates an instance of WebSocketOrgEchoActionReq.User.Item2array, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(
        possibleDtoObject: WebSocketOrgEchoActionReqType.UserType.Item2arrayType,
      ) {
        return new WebSocketOrgEchoActionReq.User.Item2array(possibleDtoObject);
      }
      /**
       * Creates an instance of WebSocketOrgEchoActionReq.User.Item2array, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(
        partialDtoObject: PartialDeep<WebSocketOrgEchoActionReqType.UserType.Item2arrayType>,
      ) {
        return new WebSocketOrgEchoActionReq.User.Item2array(partialDtoObject);
      }
      copyWith(
        partial: PartialDeep<WebSocketOrgEchoActionReqType.UserType.Item2arrayType>,
      ): InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array> {
        return new WebSocketOrgEchoActionReq.User.Item2array({
          ...this.toJSON(),
          ...partial,
        });
      }
      clone(): InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array> {
        return new WebSocketOrgEchoActionReq.User.Item2array(this.toJSON());
      }
    };
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
      const d = data as Partial<User>;
      if (d.item1 !== undefined) {
        this.item1 = d.item1;
      }
      if (d.item2array !== undefined) {
        this.item2array = d.item2array;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly. Each
     *	field goes through #toPlainJSON rather than a bare this.#field: a
     *	nested dto instance, or an emigo Array/One/Collection wrapper, has its
     *	own toJSON that JSON.stringify would only reach on a *second* pass -
     *	calling it here means toJSON()'s own return value is already the same
     *	plain shape a consumer choosing the exported type instead of this
     *	class gets, rather than a shallow object still holding class instances.
     **/
    toJSON(): WebSocketOrgEchoActionReqType.UserType {
      return {
        item1: this.#toPlainJSON(this.#item1),
        item2array: this.#toPlainJSON(this.#item2array),
      } as WebSocketOrgEchoActionReqType.UserType;
    }
    /**
     * Resolves value into a plain, JSON-serializable value: recurses into
     * arrays, and - the case JSON.stringify's own recursion would only get to
     * after this method already returned - calls a nested value's own
     * toJSON() (a dto instance, or an emigo Array/One/Collection wrapper)
     * rather than leaving it as-is. A plain value (string, number, a map's
     * own plain object, ...) is returned unchanged.
     **/
    #toPlainJSON(value: unknown): unknown {
      if (value === null || value === undefined) {
        return value;
      }
      if (Array.isArray(value)) {
        return value.map((item) => this.#toPlainJSON(item));
      }
      const asAny = value as any;
      if (typeof asAny.toJSON === "function") {
        return this.#toPlainJSON(asAny.toJSON());
      }
      return value;
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        item1: "item1",
        item2array$: "item2array",
        get item2array() {
          return withPrefix(
            "user.item2array[:i]",
            WebSocketOrgEchoActionReq.User.Item2array.Fields,
          );
        },
      };
    }
    /**
     * Creates an instance of WebSocketOrgEchoActionReq.User, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: WebSocketOrgEchoActionReqType.UserType) {
      return new WebSocketOrgEchoActionReq.User(possibleDtoObject);
    }
    /**
     * Creates an instance of WebSocketOrgEchoActionReq.User, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<WebSocketOrgEchoActionReqType.UserType>,
    ) {
      return new WebSocketOrgEchoActionReq.User(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<WebSocketOrgEchoActionReqType.UserType>,
    ): InstanceType<typeof WebSocketOrgEchoActionReq.User> {
      return new WebSocketOrgEchoActionReq.User({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof WebSocketOrgEchoActionReq.User> {
      return new WebSocketOrgEchoActionReq.User(this.toJSON());
    }
  };
  constructor(data: unknown = undefined) {
    if (data === null || data === undefined) {
      this.#lateInitFields();
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
    const d = data as Partial<WebSocketOrgEchoActionReq>;
    if (d.firstName !== undefined) {
      this.firstName = d.firstName;
    }
    if (d.lastName !== undefined) {
      this.lastName = d.lastName;
    }
    if (d.user !== undefined) {
      this.user = d.user;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<WebSocketOrgEchoActionReq>;
    if (!(d.user instanceof WebSocketOrgEchoActionReq.User)) {
      this.user = new WebSocketOrgEchoActionReq.User(d.user || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly. Each
   *	field goes through #toPlainJSON rather than a bare this.#field: a
   *	nested dto instance, or an emigo Array/One/Collection wrapper, has its
   *	own toJSON that JSON.stringify would only reach on a *second* pass -
   *	calling it here means toJSON()'s own return value is already the same
   *	plain shape a consumer choosing the exported type instead of this
   *	class gets, rather than a shallow object still holding class instances.
   **/
  toJSON(): WebSocketOrgEchoActionReqType {
    return {
      firstName: this.#toPlainJSON(this.#firstName),
      lastName: this.#toPlainJSON(this.#lastName),
      user: this.#toPlainJSON(this.#user),
    } as WebSocketOrgEchoActionReqType;
  }
  /**
   * Resolves value into a plain, JSON-serializable value: recurses into
   * arrays, and - the case JSON.stringify's own recursion would only get to
   * after this method already returned - calls a nested value's own
   * toJSON() (a dto instance, or an emigo Array/One/Collection wrapper)
   * rather than leaving it as-is. A plain value (string, number, a map's
   * own plain object, ...) is returned unchanged.
   **/
  #toPlainJSON(value: unknown): unknown {
    if (value === null || value === undefined) {
      return value;
    }
    if (Array.isArray(value)) {
      return value.map((item) => this.#toPlainJSON(item));
    }
    const asAny = value as any;
    if (typeof asAny.toJSON === "function") {
      return this.#toPlainJSON(asAny.toJSON());
    }
    return value;
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      firstName: "firstName",
      lastName: "lastName",
      user$: "user",
      get user() {
        return withPrefix("user", WebSocketOrgEchoActionReq.User.Fields);
      },
    };
  }
  /**
   * Creates an instance of WebSocketOrgEchoActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WebSocketOrgEchoActionReqType) {
    return new WebSocketOrgEchoActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of WebSocketOrgEchoActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WebSocketOrgEchoActionReqType>) {
    return new WebSocketOrgEchoActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WebSocketOrgEchoActionReqType>,
  ): InstanceType<typeof WebSocketOrgEchoActionReq> {
    return new WebSocketOrgEchoActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WebSocketOrgEchoActionReq> {
    return new WebSocketOrgEchoActionReq(this.toJSON());
  }
}
export abstract class WebSocketOrgEchoActionReqFactory {
  abstract create(data: unknown): WebSocketOrgEchoActionReq;
}
/**
 * The base type definition for webSocketOrgEchoActionReq
 **/
export type WebSocketOrgEchoActionReqType = {
  /**
   *
   * @type {string}
   **/
  firstName: string;
  /**
   *
   * @type {string}
   **/
  lastName: string;
  /**
   *
   * @type {WebSocketOrgEchoActionReqType.UserType}
   **/
  user: WebSocketOrgEchoActionReqType.UserType;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WebSocketOrgEchoActionReqType {
  /**
   * The base type definition for userType
   **/
  export type UserType = {
    /**
     *
     * @type {string}
     **/
    item1: string;
    /**
     *
     * @type {WebSocketOrgEchoActionReqType.UserType.Item2arrayType[]}
     **/
    item2array: WebSocketOrgEchoActionReqType.UserType.Item2arrayType[];
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace UserType {
    /**
     * The base type definition for item2arrayType
     **/
    export type Item2arrayType = {
      /**
       *
       * @type {number}
       **/
      subItem1: number;
      /**
       *
       * @type {number}
       **/
      subItem2: number;
    };
    // eslint-disable-next-line @typescript-eslint/no-namespace
    export namespace Item2arrayType {}
  }
}
/**
 * The base class definition for webSocketOrgEchoActionRes
 **/
export class WebSocketOrgEchoActionRes {
  /**
   *
   * @type {string}
   **/
  #lastName: string = "";
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
  set lastName(value: string) {
    this.#lastName = String(value);
  }
  setLastName(value: string) {
    this.lastName = value;
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
    const d = data as Partial<WebSocketOrgEchoActionRes>;
    if (d.lastName !== undefined) {
      this.lastName = d.lastName;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly. Each
   *	field goes through #toPlainJSON rather than a bare this.#field: a
   *	nested dto instance, or an emigo Array/One/Collection wrapper, has its
   *	own toJSON that JSON.stringify would only reach on a *second* pass -
   *	calling it here means toJSON()'s own return value is already the same
   *	plain shape a consumer choosing the exported type instead of this
   *	class gets, rather than a shallow object still holding class instances.
   **/
  toJSON(): WebSocketOrgEchoActionResType {
    return {
      lastName: this.#toPlainJSON(this.#lastName),
    } as WebSocketOrgEchoActionResType;
  }
  /**
   * Resolves value into a plain, JSON-serializable value: recurses into
   * arrays, and - the case JSON.stringify's own recursion would only get to
   * after this method already returned - calls a nested value's own
   * toJSON() (a dto instance, or an emigo Array/One/Collection wrapper)
   * rather than leaving it as-is. A plain value (string, number, a map's
   * own plain object, ...) is returned unchanged.
   **/
  #toPlainJSON(value: unknown): unknown {
    if (value === null || value === undefined) {
      return value;
    }
    if (Array.isArray(value)) {
      return value.map((item) => this.#toPlainJSON(item));
    }
    const asAny = value as any;
    if (typeof asAny.toJSON === "function") {
      return this.#toPlainJSON(asAny.toJSON());
    }
    return value;
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      lastName: "lastName",
    };
  }
  /**
   * Creates an instance of WebSocketOrgEchoActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WebSocketOrgEchoActionResType) {
    return new WebSocketOrgEchoActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of WebSocketOrgEchoActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WebSocketOrgEchoActionResType>) {
    return new WebSocketOrgEchoActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WebSocketOrgEchoActionResType>,
  ): InstanceType<typeof WebSocketOrgEchoActionRes> {
    return new WebSocketOrgEchoActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WebSocketOrgEchoActionRes> {
    return new WebSocketOrgEchoActionRes(this.toJSON());
  }
}
export abstract class WebSocketOrgEchoActionResFactory {
  abstract create(data: unknown): WebSocketOrgEchoActionRes;
}
/**
 * The base type definition for webSocketOrgEchoActionRes
 **/
export type WebSocketOrgEchoActionResType = {
  /**
   *
   * @type {string}
   **/
  lastName: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WebSocketOrgEchoActionResType {}
