import { GiantDto } from "./GiantDto";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for giantDto
 **/
export class GiantDto {
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
  #firstNameNullable?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get firstNameNullable() {
    return this.#firstNameNullable;
  }
  /**
   *
   * @type {string}
   **/
  set firstNameNullable(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#firstNameNullable = correctType ? value : String(value);
  }
  setFirstNameNullable(value: string | null | undefined) {
    this.firstNameNullable = value;
    return this;
  }
  /**
   *
   * @type {GiantDto.Array}
   **/
  #array: InstanceType<typeof GiantDto.Array>[] = [];
  /**
   *
   * @returns {GiantDto.Array}
   **/
  get array() {
    return this.#array;
  }
  /**
   *
   * @type {GiantDto.Array}
   **/
  set array(value: InstanceType<typeof GiantDto.Array>[]) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof GiantDto.Array) {
      this.#array = value;
    } else {
      this.#array = value.map((item) => new GiantDto.Array(item));
    }
  }
  setArray(value: InstanceType<typeof GiantDto.Array>[]) {
    this.array = value;
    return this;
  }
  /**
   *
   * @type {GiantDto.ArrayNullable}
   **/
  #arrayNullable?:
    | InstanceType<typeof GiantDto.ArrayNullable>[]
    | null
    | undefined
    | null = undefined;
  /**
   *
   * @returns {GiantDto.ArrayNullable}
   **/
  get arrayNullable() {
    return this.#arrayNullable;
  }
  /**
   *
   * @type {GiantDto.ArrayNullable}
   **/
  set arrayNullable(
    value:
      | InstanceType<typeof GiantDto.ArrayNullable>[]
      | null
      | undefined
      | null
      | undefined
  ) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof GiantDto.ArrayNullable) {
      this.#arrayNullable = value;
    } else {
      this.#arrayNullable = value.map(
        (item) => new GiantDto.ArrayNullable(item)
      );
    }
  }
  setArrayNullable(
    value:
      | InstanceType<typeof GiantDto.ArrayNullable>[]
      | null
      | undefined
      | null
      | undefined
  ) {
    this.arrayNullable = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #booleanField!: boolean;
  /**
   *
   * @returns {boolean}
   **/
  get booleanField() {
    return this.#booleanField;
  }
  /**
   *
   * @type {boolean}
   **/
  set booleanField(value: boolean) {
    this.#booleanField = Boolean(value);
  }
  setBooleanField(value: boolean) {
    this.booleanField = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #booleanFieldNullable?: boolean | null = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get booleanFieldNullable() {
    return this.#booleanFieldNullable;
  }
  /**
   *
   * @type {boolean}
   **/
  set booleanFieldNullable(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#booleanFieldNullable = correctType ? value : Boolean(value);
  }
  setBooleanFieldNullable(value: boolean | null | undefined) {
    this.booleanFieldNullable = value;
    return this;
  }
  /**
   *
   * @type {GiantDto[]}
   **/
  #collectionItems: GiantDto[] = [];
  /**
   *
   * @returns {GiantDto[]}
   **/
  get collectionItems() {
    return this.#collectionItems;
  }
  /**
   *
   * @type {GiantDto[]}
   **/
  set collectionItems(value: GiantDto[]) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof GiantDto) {
      this.#collectionItems = value;
    } else {
      this.#collectionItems = value.map((item) => new GiantDto(item));
    }
  }
  setCollectionItems(value: GiantDto[]) {
    this.collectionItems = value;
    return this;
  }
  /**
   *
   * @type {GiantDto[]}
   **/
  #collectionItemsNullable?: GiantDto[] | null = undefined;
  /**
   *
   * @returns {GiantDto[]}
   **/
  get collectionItemsNullable() {
    return this.#collectionItemsNullable;
  }
  /**
   *
   * @type {GiantDto[]}
   **/
  set collectionItemsNullable(value: GiantDto[] | null | undefined) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof GiantDto) {
      this.#collectionItemsNullable = value;
    } else {
      this.#collectionItemsNullable = value.map((item) => new GiantDto(item));
    }
  }
  setCollectionItemsNullable(value: GiantDto[] | null | undefined) {
    this.collectionItemsNullable = value;
    return this;
  }
  /**
   *
   * @type {Date}
   **/
  #dateObject!: Date;
  /**
   *
   * @returns {Date}
   **/
  get dateObject() {
    return this.#dateObject;
  }
  /**
   *
   * @type {Date}
   **/
  set dateObject(value: Date) {}
  setDateObject(value: Date) {
    this.dateObject = value;
    return this;
  }
  /**
   *
   * @type {GiantDto}
   **/
  #singleRef!: GiantDto;
  /**
   *
   * @returns {GiantDto}
   **/
  get singleRef() {
    return this.#singleRef;
  }
  /**
   *
   * @type {GiantDto}
   **/
  set singleRef(value: GiantDto) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof GiantDto) {
      this.#singleRef = value;
    } else {
      this.#singleRef = new GiantDto(value);
    }
  }
  setSingleRef(value: GiantDto) {
    this.singleRef = value;
    return this;
  }
  /**
   *
   * @type {GiantDto}
   **/
  #singleRefNullable?: GiantDto | null = undefined;
  /**
   *
   * @returns {GiantDto}
   **/
  get singleRefNullable() {
    return this.#singleRefNullable;
  }
  /**
   *
   * @type {GiantDto}
   **/
  set singleRefNullable(value: GiantDto | null | undefined) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof GiantDto) {
      this.#singleRefNullable = value;
    } else {
      this.#singleRefNullable = new GiantDto(value);
    }
  }
  setSingleRefNullable(value: GiantDto | null | undefined) {
    this.singleRefNullable = value;
    return this;
  }
  /**
   *
   * @type {"Key3" | "Key4"}
   **/
  #enumeration!: "Key3" | "Key4";
  /**
   *
   * @returns {"Key3" | "Key4"}
   **/
  get enumeration() {
    return this.#enumeration;
  }
  /**
   *
   * @type {"Key3" | "Key4"}
   **/
  set enumeration(value: "Key3" | "Key4") {}
  setEnumeration(value: "Key3" | "Key4") {
    this.enumeration = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #enumerationNullable?: any | null = undefined;
  /**
   *
   * @returns {any}
   **/
  get enumerationNullable() {
    return this.#enumerationNullable;
  }
  /**
   *
   * @type {any}
   **/
  set enumerationNullable(value: any | null | undefined) {}
  setEnumerationNullable(value: any | null | undefined) {
    this.enumerationNullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #floatingPoint32: number = 0.0;
  /**
   *
   * @returns {number}
   **/
  get floatingPoint32() {
    return this.#floatingPoint32;
  }
  /**
   *
   * @type {number}
   **/
  set floatingPoint32(value: number) {}
  setFloatingPoint32(value: number) {
    this.floatingPoint32 = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #floatingPoint32Nullable?: number | null = undefined;
  /**
   *
   * @returns {number}
   **/
  get floatingPoint32Nullable() {
    return this.#floatingPoint32Nullable;
  }
  /**
   *
   * @type {number}
   **/
  set floatingPoint32Nullable(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#floatingPoint32Nullable = parsedValue;
    }
  }
  setFloatingPoint32Nullable(value: number | null | undefined) {
    this.floatingPoint32Nullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #floatingPoint64: number = 0.0;
  /**
   *
   * @returns {number}
   **/
  get floatingPoint64() {
    return this.#floatingPoint64;
  }
  /**
   *
   * @type {number}
   **/
  set floatingPoint64(value: number) {}
  setFloatingPoint64(value: number) {
    this.floatingPoint64 = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #floatingPoint64Nullable?: number | null = undefined;
  /**
   *
   * @returns {number}
   **/
  get floatingPoint64Nullable() {
    return this.#floatingPoint64Nullable;
  }
  /**
   *
   * @type {number}
   **/
  set floatingPoint64Nullable(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#floatingPoint64Nullable = parsedValue;
    }
  }
  setFloatingPoint64Nullable(value: number | null | undefined) {
    this.floatingPoint64Nullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #integerValue: number = 0;
  /**
   *
   * @returns {number}
   **/
  get integerValue() {
    return this.#integerValue;
  }
  /**
   *
   * @type {number}
   **/
  set integerValue(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#integerValue = parsedValue;
    }
  }
  setIntegerValue(value: number) {
    this.integerValue = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #integer32ValueNullable?: number | null = undefined;
  /**
   *
   * @returns {number}
   **/
  get integer32ValueNullable() {
    return this.#integer32ValueNullable;
  }
  /**
   *
   * @type {number}
   **/
  set integer32ValueNullable(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#integer32ValueNullable = parsedValue;
    }
  }
  setInteger32ValueNullable(value: number | null | undefined) {
    this.integer32ValueNullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #integer32Value: number = 0;
  /**
   *
   * @returns {number}
   **/
  get integer32Value() {
    return this.#integer32Value;
  }
  /**
   *
   * @type {number}
   **/
  set integer32Value(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#integer32Value = parsedValue;
    }
  }
  setInteger32Value(value: number) {
    this.integer32Value = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #integer64ValueNullable?: number | null = undefined;
  /**
   *
   * @returns {number}
   **/
  get integer64ValueNullable() {
    return this.#integer64ValueNullable;
  }
  /**
   *
   * @type {number}
   **/
  set integer64ValueNullable(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#integer64ValueNullable = parsedValue;
    }
  }
  setInteger64ValueNullable(value: number | null | undefined) {
    this.integer64ValueNullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #integer64Value: number = 0;
  /**
   *
   * @returns {number}
   **/
  get integer64Value() {
    return this.#integer64Value;
  }
  /**
   *
   * @type {number}
   **/
  set integer64Value(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#integer64Value = parsedValue;
    }
  }
  setInteger64Value(value: number) {
    this.integer64Value = value;
    return this;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  #mapValue!: { [key: string]: any };
  /**
   *
   * @returns {{[key: string]: any}}
   **/
  get mapValue() {
    return this.#mapValue;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  set mapValue(value: { [key: string]: any }) {}
  setMapValue(value: { [key: string]: any }) {
    this.mapValue = value;
    return this;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  #mapValueNullable?: { [key: string]: any } | null = undefined;
  /**
   *
   * @returns {{[key: string]: any}}
   **/
  get mapValueNullable() {
    return this.#mapValueNullable;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  set mapValueNullable(value: { [key: string]: any } | null | undefined) {}
  setMapValueNullable(value: { [key: string]: any } | null | undefined) {
    this.mapValueNullable = value;
    return this;
  }
  /**
   *
   * @type {unknown[]}
   **/
  #sliceValue: unknown[] = [];
  /**
   *
   * @returns {unknown[]}
   **/
  get sliceValue() {
    return this.#sliceValue;
  }
  /**
   *
   * @type {unknown[]}
   **/
  set sliceValue(value: unknown[]) {}
  setSliceValue(value: unknown[]) {
    this.sliceValue = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #sliceValueNullable?: any | null = undefined;
  /**
   *
   * @returns {any}
   **/
  get sliceValueNullable() {
    return this.#sliceValueNullable;
  }
  /**
   *
   * @type {any}
   **/
  set sliceValueNullable(value: any | null | undefined) {}
  setSliceValueNullable(value: any | null | undefined) {
    this.sliceValueNullable = value;
    return this;
  }
  /**
   *
   * @type {GiantDto.ObjectValue}
   **/
  #objectValue!: InstanceType<typeof GiantDto.ObjectValue>;
  /**
   *
   * @returns {GiantDto.ObjectValue}
   **/
  get objectValue() {
    return this.#objectValue;
  }
  /**
   *
   * @type {GiantDto.ObjectValue}
   **/
  set objectValue(value: InstanceType<typeof GiantDto.ObjectValue>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof GiantDto.ObjectValue) {
      this.#objectValue = value;
    } else {
      this.#objectValue = new GiantDto.ObjectValue(value);
    }
  }
  setObjectValue(value: InstanceType<typeof GiantDto.ObjectValue>) {
    this.objectValue = value;
    return this;
  }
  /**
   * The base class definition for array
   **/
  static Array = class Array {
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
            typeof data
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
      const d = data as Partial<Array>;
      if (d.subItem1 !== undefined) {
        this.subItem1 = d.subItem1;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        subItem1: this.#subItem1,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        subItem1: "subItem1",
      };
    }
    /**
     * Creates an instance of GiantDto.Array, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: GiantDtoType.ArrayType) {
      return new GiantDto.Array(possibleDtoObject);
    }
    /**
     * Creates an instance of GiantDto.Array, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<GiantDtoType.ArrayType>) {
      return new GiantDto.Array(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<GiantDtoType.ArrayType>
    ): InstanceType<typeof GiantDto.Array> {
      return new GiantDto.Array({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof GiantDto.Array> {
      return new GiantDto.Array(this.toJSON());
    }
  };
  /**
   * The base class definition for arrayNullable
   **/
  static ArrayNullable = class ArrayNullable {
    /**
     *
     * @type {number}
     **/
    #subItemNullable1?: number | null = undefined;
    /**
     *
     * @returns {number}
     **/
    get subItemNullable1() {
      return this.#subItemNullable1;
    }
    /**
     *
     * @type {number}
     **/
    set subItemNullable1(value: number | null | undefined) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#subItemNullable1 = parsedValue;
      }
    }
    setSubItemNullable1(value: number | null | undefined) {
      this.subItemNullable1 = value;
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
            typeof data
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
      const d = data as Partial<ArrayNullable>;
      if (d.subItemNullable1 !== undefined) {
        this.subItemNullable1 = d.subItemNullable1;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        subItemNullable1: this.#subItemNullable1,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        subItemNullable1: "subItemNullable1",
      };
    }
    /**
     * Creates an instance of GiantDto.ArrayNullable, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: GiantDtoType.ArrayNullableType) {
      return new GiantDto.ArrayNullable(possibleDtoObject);
    }
    /**
     * Creates an instance of GiantDto.ArrayNullable, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<GiantDtoType.ArrayNullableType>) {
      return new GiantDto.ArrayNullable(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<GiantDtoType.ArrayNullableType>
    ): InstanceType<typeof GiantDto.ArrayNullable> {
      return new GiantDto.ArrayNullable({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof GiantDto.ArrayNullable> {
      return new GiantDto.ArrayNullable(this.toJSON());
    }
  };
  /**
   * The base class definition for objectValue
   **/
  static ObjectValue = class ObjectValue {
    /**
     *
     * @type {GiantDto.ObjectValue.InnerObject}
     **/
    #innerObject?:
      | InstanceType<typeof GiantDto.ObjectValue.InnerObject>
      | null
      | undefined
      | null = undefined;
    /**
     *
     * @returns {GiantDto.ObjectValue.InnerObject}
     **/
    get innerObject() {
      return this.#innerObject;
    }
    /**
     *
     * @type {GiantDto.ObjectValue.InnerObject}
     **/
    set innerObject(
      value:
        | InstanceType<typeof GiantDto.ObjectValue.InnerObject>
        | null
        | undefined
        | null
        | undefined
    ) {
      // For objects, the sub type needs to always be instance of the sub class.
      if (value instanceof GiantDto.ObjectValue.InnerObject) {
        this.#innerObject = value;
      } else {
        this.#innerObject = new GiantDto.ObjectValue.InnerObject(value);
      }
    }
    setInnerObject(
      value:
        | InstanceType<typeof GiantDto.ObjectValue.InnerObject>
        | null
        | undefined
        | null
        | undefined
    ) {
      this.innerObject = value;
      return this;
    }
    /**
     * The base class definition for innerObject
     **/
    static InnerObject = class InnerObject {
      /**
       *
       * @type {string}
       **/
      #innerObjText: string = "";
      /**
       *
       * @returns {string}
       **/
      get innerObjText() {
        return this.#innerObjText;
      }
      /**
       *
       * @type {string}
       **/
      set innerObjText(value: string) {
        this.#innerObjText = String(value);
      }
      setInnerObjText(value: string) {
        this.innerObjText = value;
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
              typeof data
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
        const d = data as Partial<InnerObject>;
        if (d.innerObjText !== undefined) {
          this.innerObjText = d.innerObjText;
        }
      }
      /**
       *	Special toJSON override, since the field are private,
       *	Json stringify won't see them unless we mention it explicitly.
       **/
      toJSON() {
        return {
          innerObjText: this.#innerObjText,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          innerObjText: "innerObjText",
        };
      }
      /**
       * Creates an instance of GiantDto.ObjectValue.InnerObject, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(
        possibleDtoObject: GiantDtoType.ObjectValueType.InnerObjectType
      ) {
        return new GiantDto.ObjectValue.InnerObject(possibleDtoObject);
      }
      /**
       * Creates an instance of GiantDto.ObjectValue.InnerObject, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(
        partialDtoObject: PartialDeep<GiantDtoType.ObjectValueType.InnerObjectType>
      ) {
        return new GiantDto.ObjectValue.InnerObject(partialDtoObject);
      }
      copyWith(
        partial: PartialDeep<GiantDtoType.ObjectValueType.InnerObjectType>
      ): InstanceType<typeof GiantDto.ObjectValue.InnerObject> {
        return new GiantDto.ObjectValue.InnerObject({
          ...this.toJSON(),
          ...partial,
        });
      }
      clone(): InstanceType<typeof GiantDto.ObjectValue.InnerObject> {
        return new GiantDto.ObjectValue.InnerObject(this.toJSON());
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
            typeof data
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
      const d = data as Partial<ObjectValue>;
      if (d.innerObject !== undefined) {
        this.innerObject = d.innerObject;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        innerObject: this.#innerObject,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        innerObject$: "innerObject",
        get innerObject() {
          return withPrefix(
            "objectValue.innerObject",
            GiantDto.ObjectValue.InnerObject.Fields
          );
        },
      };
    }
    /**
     * Creates an instance of GiantDto.ObjectValue, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: GiantDtoType.ObjectValueType) {
      return new GiantDto.ObjectValue(possibleDtoObject);
    }
    /**
     * Creates an instance of GiantDto.ObjectValue, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<GiantDtoType.ObjectValueType>) {
      return new GiantDto.ObjectValue(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<GiantDtoType.ObjectValueType>
    ): InstanceType<typeof GiantDto.ObjectValue> {
      return new GiantDto.ObjectValue({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof GiantDto.ObjectValue> {
      return new GiantDto.ObjectValue(this.toJSON());
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
          typeof data
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
    const d = data as Partial<GiantDto>;
    if (d.firstName !== undefined) {
      this.firstName = d.firstName;
    }
    if (d.firstNameNullable !== undefined) {
      this.firstNameNullable = d.firstNameNullable;
    }
    if (d.array !== undefined) {
      this.array = d.array;
    }
    if (d.arrayNullable !== undefined) {
      this.arrayNullable = d.arrayNullable;
    }
    if (d.booleanField !== undefined) {
      this.booleanField = d.booleanField;
    }
    if (d.booleanFieldNullable !== undefined) {
      this.booleanFieldNullable = d.booleanFieldNullable;
    }
    if (d.collectionItems !== undefined) {
      this.collectionItems = d.collectionItems;
    }
    if (d.collectionItemsNullable !== undefined) {
      this.collectionItemsNullable = d.collectionItemsNullable;
    }
    if (d.dateObject !== undefined) {
      this.dateObject = d.dateObject;
    }
    if (d.singleRef !== undefined) {
      this.singleRef = d.singleRef;
    }
    if (d.singleRefNullable !== undefined) {
      this.singleRefNullable = d.singleRefNullable;
    }
    if (d.enumeration !== undefined) {
      this.enumeration = d.enumeration;
    }
    if (d.enumerationNullable !== undefined) {
      this.enumerationNullable = d.enumerationNullable;
    }
    if (d.floatingPoint32 !== undefined) {
      this.floatingPoint32 = d.floatingPoint32;
    }
    if (d.floatingPoint32Nullable !== undefined) {
      this.floatingPoint32Nullable = d.floatingPoint32Nullable;
    }
    if (d.floatingPoint64 !== undefined) {
      this.floatingPoint64 = d.floatingPoint64;
    }
    if (d.floatingPoint64Nullable !== undefined) {
      this.floatingPoint64Nullable = d.floatingPoint64Nullable;
    }
    if (d.integerValue !== undefined) {
      this.integerValue = d.integerValue;
    }
    if (d.integer32ValueNullable !== undefined) {
      this.integer32ValueNullable = d.integer32ValueNullable;
    }
    if (d.integer32Value !== undefined) {
      this.integer32Value = d.integer32Value;
    }
    if (d.integer64ValueNullable !== undefined) {
      this.integer64ValueNullable = d.integer64ValueNullable;
    }
    if (d.integer64Value !== undefined) {
      this.integer64Value = d.integer64Value;
    }
    if (d.mapValue !== undefined) {
      this.mapValue = d.mapValue;
    }
    if (d.mapValueNullable !== undefined) {
      this.mapValueNullable = d.mapValueNullable;
    }
    if (d.sliceValue !== undefined) {
      this.sliceValue = d.sliceValue;
    }
    if (d.sliceValueNullable !== undefined) {
      this.sliceValueNullable = d.sliceValueNullable;
    }
    if (d.objectValue !== undefined) {
      this.objectValue = d.objectValue;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<GiantDto>;
    if (!(d.singleRef instanceof GiantDto)) {
      this.singleRef = new GiantDto(d.singleRef || {});
    }
    if (!(d.objectValue instanceof GiantDto.ObjectValue)) {
      this.objectValue = new GiantDto.ObjectValue(d.objectValue || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      firstName: this.#firstName,
      firstNameNullable: this.#firstNameNullable,
      array: this.#array,
      arrayNullable: this.#arrayNullable,
      booleanField: this.#booleanField,
      booleanFieldNullable: this.#booleanFieldNullable,
      collectionItems: this.#collectionItems,
      collectionItemsNullable: this.#collectionItemsNullable,
      dateObject: this.#dateObject,
      singleRef: this.#singleRef,
      singleRefNullable: this.#singleRefNullable,
      enumeration: this.#enumeration,
      enumerationNullable: this.#enumerationNullable,
      floatingPoint32: this.#floatingPoint32,
      floatingPoint32Nullable: this.#floatingPoint32Nullable,
      floatingPoint64: this.#floatingPoint64,
      floatingPoint64Nullable: this.#floatingPoint64Nullable,
      integerValue: this.#integerValue,
      integer32ValueNullable: this.#integer32ValueNullable,
      integer32Value: this.#integer32Value,
      integer64ValueNullable: this.#integer64ValueNullable,
      integer64Value: this.#integer64Value,
      mapValue: this.#mapValue,
      mapValueNullable: this.#mapValueNullable,
      sliceValue: this.#sliceValue,
      sliceValueNullable: this.#sliceValueNullable,
      objectValue: this.#objectValue,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      firstName: "firstName",
      firstNameNullable: "firstNameNullable",
      array$: "array",
      get array() {
        return withPrefix("array[:i]", GiantDto.Array.Fields);
      },
      arrayNullable$: "arrayNullable",
      get arrayNullable() {
        return withPrefix("arrayNullable[:i]", GiantDto.ArrayNullable.Fields);
      },
      booleanField: "booleanField",
      booleanFieldNullable: "booleanFieldNullable",
      collectionItems$: "collectionItems",
      get collectionItems() {
        return withPrefix("collectionItems[:i]", GiantDto.Fields);
      },
      collectionItemsNullable$: "collectionItemsNullable",
      get collectionItemsNullable() {
        return withPrefix("collectionItemsNullable", GiantDto.Fields);
      },
      dateObject: "dateObject",
      singleRef$: "singleRef",
      get singleRef() {
        return withPrefix("singleRef", GiantDto.Fields);
      },
      singleRefNullable: "singleRefNullable",
      enumeration: "enumeration",
      enumerationNullable: "enumerationNullable",
      floatingPoint32: "floatingPoint32",
      floatingPoint32Nullable: "floatingPoint32Nullable",
      floatingPoint64: "floatingPoint64",
      floatingPoint64Nullable: "floatingPoint64Nullable",
      integerValue: "integerValue",
      integer32ValueNullable: "integer32ValueNullable",
      integer32Value: "integer32Value",
      integer64ValueNullable: "integer64ValueNullable",
      integer64Value: "integer64Value",
      mapValue: "mapValue",
      mapValueNullable: "mapValueNullable",
      sliceValue$: "sliceValue",
      get sliceValue() {
        return "sliceValue[:i]";
      },
      sliceValueNullable: "sliceValueNullable",
      objectValue$: "objectValue",
      get objectValue() {
        return withPrefix("objectValue", GiantDto.ObjectValue.Fields);
      },
    };
  }
  /**
   * Creates an instance of GiantDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: GiantDtoType) {
    return new GiantDto(possibleDtoObject);
  }
  /**
   * Creates an instance of GiantDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<GiantDtoType>) {
    return new GiantDto(partialDtoObject);
  }
  copyWith(partial: PartialDeep<GiantDtoType>): InstanceType<typeof GiantDto> {
    return new GiantDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof GiantDto> {
    return new GiantDto(this.toJSON());
  }
}
export abstract class GiantDtoFactory {
  abstract create(data: unknown): GiantDto;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
    ? PartialDeep<T[P]>
    : T[P];
};
/**
 * The base type definition for giantDto
 **/
export type GiantDtoType = {
  /**
   *
   * @type {string}
   **/
  firstName: string;
  /**
   *
   * @type {string}
   **/
  firstNameNullable?: string;
  /**
   *
   * @type {GiantDtoType.ArrayType[]}
   **/
  array: GiantDtoType.ArrayType[];
  /**
   *
   * @type {GiantDtoType.ArrayNullableType[]}
   **/
  arrayNullable?: GiantDtoType.ArrayNullableType[];
  /**
   *
   * @type {boolean}
   **/
  booleanField: boolean;
  /**
   *
   * @type {boolean}
   **/
  booleanFieldNullable?: boolean;
  /**
   *
   * @type {GiantDto[]}
   **/
  collectionItems: GiantDto[];
  /**
   *
   * @type {GiantDto[]}
   **/
  collectionItemsNullable?: GiantDto[];
  /**
   *
   * @type {Date}
   **/
  dateObject: Date;
  /**
   *
   * @type {GiantDto}
   **/
  singleRef: GiantDto;
  /**
   *
   * @type {GiantDto}
   **/
  singleRefNullable?: GiantDto;
  /**
   *
   * @type {"Key3" | "Key4"}
   **/
  enumeration: "Key3" | "Key4";
  /**
   *
   * @type {any}
   **/
  enumerationNullable?: any;
  /**
   *
   * @type {number}
   **/
  floatingPoint32: number;
  /**
   *
   * @type {number}
   **/
  floatingPoint32Nullable?: number;
  /**
   *
   * @type {number}
   **/
  floatingPoint64: number;
  /**
   *
   * @type {number}
   **/
  floatingPoint64Nullable?: number;
  /**
   *
   * @type {number}
   **/
  integerValue: number;
  /**
   *
   * @type {number}
   **/
  integer32ValueNullable?: number;
  /**
   *
   * @type {number}
   **/
  integer32Value: number;
  /**
   *
   * @type {number}
   **/
  integer64ValueNullable?: number;
  /**
   *
   * @type {number}
   **/
  integer64Value: number;
  /**
   *
   * @type {{[key: string]: any}}
   **/
  mapValue: { [key: string]: any };
  /**
   *
   * @type {{[key: string]: any}}
   **/
  mapValueNullable?: { [key: string]: any };
  /**
   *
   * @type {unknown[]}
   **/
  sliceValue: unknown[];
  /**
   *
   * @type {any}
   **/
  sliceValueNullable?: any;
  /**
   *
   * @type {GiantDtoType.ObjectValueType}
   **/
  objectValue: GiantDtoType.ObjectValueType;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GiantDtoType {
  /**
   * The base type definition for arrayType
   **/
  export type ArrayType = {
    /**
     *
     * @type {number}
     **/
    subItem1: number;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace ArrayType {}
  /**
   * The base type definition for arrayNullableType
   **/
  export type ArrayNullableType = {
    /**
     *
     * @type {number}
     **/
    subItemNullable1?: number;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace ArrayNullableType {}
  /**
   * The base type definition for objectValueType
   **/
  export type ObjectValueType = {
    /**
     *
     * @type {GiantDtoType.ObjectValueType.InnerObjectType}
     **/
    innerObject?: GiantDtoType.ObjectValueType.InnerObjectType;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace ObjectValueType {
    /**
     * The base type definition for innerObjectType
     **/
    export type InnerObjectType = {
      /**
       *
       * @type {string}
       **/
      innerObjText: string;
    };
    // eslint-disable-next-line @typescript-eslint/no-namespace
    export namespace InnerObjectType {}
  }
}
