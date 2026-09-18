import { MArray } from "./sdk/common/operators";
import { type PartialDeep } from "./sdk/common/fetchx";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for insertUsersVsqlParams
 **/
export class InsertUsersVsqlParams {
  /**
   *
   * @type {InsertUsersVsqlParams.Users}
   **/
  #users: MArray<InstanceType<typeof InsertUsersVsqlParams.Users>> = MArray.of(
    [],
  );
  /**
   *
   * @returns {InsertUsersVsqlParams.Users}
   **/
  get users() {
    return this.#users;
  }
  /**
   *
   * @type {InsertUsersVsqlParams.Users}
   **/
  set users(
    value:
      | MArray<InstanceType<typeof InsertUsersVsqlParams.Users>>
      | InstanceType<typeof InsertUsersVsqlParams.Users>[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof InsertUsersVsqlParams.Users) {
        this.#users = MArray.of(value);
      } else {
        this.#users = MArray.of(
          value.map((item) => new InsertUsersVsqlParams.Users(item)),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#users = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast<unknown>(value);
    if (ok) {
      this.#users = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to users, because it needs MArray instance or an Array.",
    );
  }
  setUsers(
    value:
      | MArray<InstanceType<typeof InsertUsersVsqlParams.Users>>
      | InstanceType<typeof InsertUsersVsqlParams.Users>[],
  ) {
    this.users = value;
    return this;
  }
  /**
   * The base class definition for users
   **/
  static Users = class Users {
    /**
     *
     * @type {string}
     **/
    #email: string = "";
    /**
     *
     * @returns {string}
     **/
    get email() {
      return this.#email;
    }
    /**
     *
     * @type {string}
     **/
    set email(value: string) {
      this.#email = String(value);
    }
    setEmail(value: string) {
      this.email = value;
      return this;
    }
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
      const d = data as Partial<Users>;
      if (d.email !== undefined) {
        this.email = d.email;
      }
      if (d.firstName !== undefined) {
        this.firstName = d.firstName;
      }
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
    toJSON(): InsertUsersVsqlParamsType.UsersType {
      return {
        email: this.#toPlainJSON(this.#email),
        firstName: this.#toPlainJSON(this.#firstName),
        lastName: this.#toPlainJSON(this.#lastName),
      } as InsertUsersVsqlParamsType.UsersType;
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
        email: "email",
        firstName: "firstName",
        lastName: "lastName",
      };
    }
    /**
     * Creates an instance of InsertUsersVsqlParams.Users, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: InsertUsersVsqlParamsType.UsersType) {
      return new InsertUsersVsqlParams.Users(possibleDtoObject);
    }
    /**
     * Creates an instance of InsertUsersVsqlParams.Users, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<InsertUsersVsqlParamsType.UsersType>,
    ) {
      return new InsertUsersVsqlParams.Users(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<InsertUsersVsqlParamsType.UsersType>,
    ): InstanceType<typeof InsertUsersVsqlParams.Users> {
      return new InsertUsersVsqlParams.Users({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof InsertUsersVsqlParams.Users> {
      return new InsertUsersVsqlParams.Users(this.toJSON());
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
    const d = data as Partial<InsertUsersVsqlParams>;
    if (d.users !== undefined) {
      this.users = d.users;
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
  toJSON(): InsertUsersVsqlParamsType {
    return {
      users: this.#toPlainJSON(this.#users),
    } as InsertUsersVsqlParamsType;
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
      users$: "users",
      get users() {
        return withPrefix("users[:i]", InsertUsersVsqlParams.Users.Fields);
      },
    };
  }
  /**
   * Creates an instance of InsertUsersVsqlParams, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: InsertUsersVsqlParamsType) {
    return new InsertUsersVsqlParams(possibleDtoObject);
  }
  /**
   * Creates an instance of InsertUsersVsqlParams, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<InsertUsersVsqlParamsType>) {
    return new InsertUsersVsqlParams(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<InsertUsersVsqlParamsType>,
  ): InstanceType<typeof InsertUsersVsqlParams> {
    return new InsertUsersVsqlParams({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof InsertUsersVsqlParams> {
    return new InsertUsersVsqlParams(this.toJSON());
  }
}
export abstract class InsertUsersVsqlParamsFactory {
  abstract create(data: unknown): InsertUsersVsqlParams;
}
/**
 * The base type definition for insertUsersVsqlParams
 **/
export type InsertUsersVsqlParamsType = {
  /**
   *
   * @type {InsertUsersVsqlParamsType.UsersType[]}
   **/
  users: InsertUsersVsqlParamsType.UsersType[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace InsertUsersVsqlParamsType {
  /**
   * The base type definition for usersType
   **/
  export type UsersType = {
    /**
     *
     * @type {string}
     **/
    email: string;
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
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace UsersType {}
}
