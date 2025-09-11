import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for singleDtoDto
 **/
export class SingleDtoDto {
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
  /**
   * This object is optional. Can be passed or not. On typ has ? operator
   * @type {SingleDtoDto.NullableObject}
   **/
  #nullableObject?:
    | InstanceType<typeof SingleDtoDto.NullableObject>
    | null
    | undefined
    | null = undefined;
  /**
   * This object is optional. Can be passed or not. On typ has ? operator
   * @returns {SingleDtoDto.NullableObject}
   **/
  get nullableObject() {
    return this.#nullableObject;
  }
  /**
   * This object is optional. Can be passed or not. On typ has ? operator
   * @type {SingleDtoDto.NullableObject}
   **/
  set nullableObject(
    value: InstanceType<typeof SingleDtoDto.NullableObject> | null | undefined
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof SingleDtoDto.NullableObject) {
      this.#nullableObject = value;
    } else {
      this.#nullableObject = new SingleDtoDto.NullableObject(value);
    }
  }
  setNullableObject(
    value: InstanceType<typeof SingleDtoDto.NullableObject> | null | undefined
  ) {
    this.nullableObject = value;
    return this;
  }
  /**
   * This object is not nullable. On classes, will be always instantiated, also on types, is not having ? operator
   * @type {SingleDtoDto.StaticObject}
   **/
  #staticObject: InstanceType<typeof SingleDtoDto.StaticObject>;
  /**
   * This object is not nullable. On classes, will be always instantiated, also on types, is not having ? operator
   * @returns {SingleDtoDto.StaticObject}
   **/
  get staticObject() {
    return this.#staticObject;
  }
  /**
   * This object is not nullable. On classes, will be always instantiated, also on types, is not having ? operator
   * @type {SingleDtoDto.StaticObject}
   **/
  set staticObject(value: InstanceType<typeof SingleDtoDto.StaticObject>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof SingleDtoDto.StaticObject) {
      this.#staticObject = value;
    } else {
      this.#staticObject = new SingleDtoDto.StaticObject(value);
    }
  }
  setStaticObject(value: InstanceType<typeof SingleDtoDto.StaticObject>) {
    this.staticObject = value;
    return this;
  }
  /**
   * The base class definition for nullableObject
   **/
  static NullableObject = class NullableObject {
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
      const correctType = typeof value === "string";
      this.#firstName = correctType ? value : "" + value;
    }
    setFirstName(value: string) {
      this.firstName = value;
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
        throw new Error("Instance is not implemented.");
      }
    }
    #isJsonAppliable(obj) {
      const isBuffer =
        typeof globalThis.Buffer !== "undefined" &&
        typeof globalThis.Buffer.isBuffer === "function" &&
        globalThis.Buffer.isBuffer(obj);
      const isBlob =
        typeof globalThis.Blob !== "undefined" &&
        obj instanceof globalThis.Blob;
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
      const d = data as Partial<NullableObject>;
      if (d.firstName !== undefined) {
        this.firstName = d.firstName;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        firstName: this.#firstName,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        firstName: "firstName",
      };
    }
  };
  /**
   * The base class definition for staticObject
   **/
  static StaticObject = class StaticObject {
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
      const correctType = typeof value === "string";
      this.#firstName = correctType ? value : "" + value;
    }
    setFirstName(value: string) {
      this.firstName = value;
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
        throw new Error("Instance is not implemented.");
      }
    }
    #isJsonAppliable(obj) {
      const isBuffer =
        typeof globalThis.Buffer !== "undefined" &&
        typeof globalThis.Buffer.isBuffer === "function" &&
        globalThis.Buffer.isBuffer(obj);
      const isBlob =
        typeof globalThis.Blob !== "undefined" &&
        obj instanceof globalThis.Blob;
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
      const d = data as Partial<StaticObject>;
      if (d.firstName !== undefined) {
        this.firstName = d.firstName;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        firstName: this.#firstName,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        firstName: "firstName",
      };
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
      throw new Error("Instance is not implemented.");
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
    const d = data as Partial<SingleDtoDto>;
    if (d.min !== undefined) {
      this.min = d.min;
    }
    if (d.max !== undefined) {
      this.max = d.max;
    }
    if (d.count !== undefined) {
      this.count = d.count;
    }
    if (d.nullableObject !== undefined) {
      this.nullableObject = d.nullableObject;
    }
    if (d.staticObject !== undefined) {
      this.staticObject = d.staticObject;
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
      nullableObject: this.#nullableObject,
      staticObject: this.#staticObject,
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
      nullableObject$: "nullableObject",
      get nullableObject() {
        return withPrefix("nullableObject", SingleDtoDto.NullableObject.Fields);
      },
      staticObject$: "staticObject",
      get staticObject() {
        return withPrefix("staticObject", SingleDtoDto.StaticObject.Fields);
      },
    };
  }
}
export abstract class SingleDtoDtoFactory {
  abstract create(data: unknown): SingleDtoDto;
}
/**
 * The base type definition for singleDtoDto
 **/
export type SingleDtoDtoType = {
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
  /**
   * This object is optional. Can be passed or not. On typ has ? operator
   * @type {SingleDtoDtoType.NullableObjectType}
   **/
  nullableObject?: SingleDtoDtoType.NullableObjectType;
  /**
   * This object is not nullable. On classes, will be always instantiated, also on types, is not having ? operator
   * @type {SingleDtoDtoType.StaticObjectType}
   **/
  staticObject: SingleDtoDtoType.StaticObjectType;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SingleDtoDtoType {
  /**
   * The base type definition for nullableObjectType
   **/
  export type NullableObjectType = {
    /**
     *
     * @type {string}
     **/
    firstName: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace NullableObjectType {}
  /**
   * The base type definition for staticObjectType
   **/
  export type StaticObjectType = {
    /**
     *
     * @type {string}
     **/
    firstName: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace StaticObjectType {}
}
