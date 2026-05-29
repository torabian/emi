import { CreateUserDtoTags } from "./CreateUserDtoTags";
import { operators } from "./sdk/common/operators";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for dataTypesDto
 **/
export class DataTypesDto {
  /**
   *
   * @type {string}
   **/
  #stringField = "";
  /**
   *
   * @returns {string}
   **/
  get stringField() {
    return this.#stringField;
  }
  /**
   *
   * @type {string}
   **/
  set stringField(value) {
    this.#stringField = String(value);
  }
  setStringField(value) {
    this.stringField = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #stringFieldNullable = undefined;
  /**
   *
   * @returns {string}
   **/
  get stringFieldNullable() {
    return this.#stringFieldNullable;
  }
  /**
   *
   * @type {string}
   **/
  set stringFieldNullable(value) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#stringFieldNullable = correctType ? value : String(value);
  }
  setStringFieldNullable(value) {
    this.stringFieldNullable = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #boolField;
  /**
   *
   * @returns {boolean}
   **/
  get boolField() {
    return this.#boolField;
  }
  /**
   *
   * @type {boolean}
   **/
  set boolField(value) {
    this.#boolField = Boolean(value);
  }
  setBoolField(value) {
    this.boolField = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #boolFieldNullable = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get boolFieldNullable() {
    return this.#boolFieldNullable;
  }
  /**
   *
   * @type {boolean}
   **/
  set boolFieldNullable(value) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#boolFieldNullable = correctType ? value : Boolean(value);
  }
  setBoolFieldNullable(value) {
    this.boolFieldNullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #intField = 0;
  /**
   *
   * @returns {number}
   **/
  get intField() {
    return this.#intField;
  }
  /**
   *
   * @type {number}
   **/
  set intField(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#intField = parsedValue;
    }
  }
  setIntField(value) {
    this.intField = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #intFieldNullable = undefined;
  /**
   *
   * @returns {number}
   **/
  get intFieldNullable() {
    return this.#intFieldNullable;
  }
  /**
   *
   * @type {number}
   **/
  set intFieldNullable(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#intFieldNullable = parsedValue;
    }
  }
  setIntFieldNullable(value) {
    this.intFieldNullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #int32Field = 0;
  /**
   *
   * @returns {number}
   **/
  get int32Field() {
    return this.#int32Field;
  }
  /**
   *
   * @type {number}
   **/
  set int32Field(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#int32Field = parsedValue;
    }
  }
  setInt32Field(value) {
    this.int32Field = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #int32FieldNullable = undefined;
  /**
   *
   * @returns {number}
   **/
  get int32FieldNullable() {
    return this.#int32FieldNullable;
  }
  /**
   *
   * @type {number}
   **/
  set int32FieldNullable(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#int32FieldNullable = parsedValue;
    }
  }
  setInt32FieldNullable(value) {
    this.int32FieldNullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #int64Field = 0;
  /**
   *
   * @returns {number}
   **/
  get int64Field() {
    return this.#int64Field;
  }
  /**
   *
   * @type {number}
   **/
  set int64Field(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#int64Field = parsedValue;
    }
  }
  setInt64Field(value) {
    this.int64Field = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #int64FieldNullable = undefined;
  /**
   *
   * @returns {number}
   **/
  get int64FieldNullable() {
    return this.#int64FieldNullable;
  }
  /**
   *
   * @type {number}
   **/
  set int64FieldNullable(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#int64FieldNullable = parsedValue;
    }
  }
  setInt64FieldNullable(value) {
    this.int64FieldNullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #float32Field = 0.0;
  /**
   *
   * @returns {number}
   **/
  get float32Field() {
    return this.#float32Field;
  }
  /**
   *
   * @type {number}
   **/
  set float32Field(value) {
    this.#float32Field = value;
  }
  setFloat32Field(value) {
    this.float32Field = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #float32FieldNullable = undefined;
  /**
   *
   * @returns {number}
   **/
  get float32FieldNullable() {
    return this.#float32FieldNullable;
  }
  /**
   *
   * @type {number}
   **/
  set float32FieldNullable(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#float32FieldNullable = parsedValue;
    }
  }
  setFloat32FieldNullable(value) {
    this.float32FieldNullable = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #float64Field = 0.0;
  /**
   *
   * @returns {number}
   **/
  get float64Field() {
    return this.#float64Field;
  }
  /**
   *
   * @type {number}
   **/
  set float64Field(value) {
    this.#float64Field = value;
  }
  setFloat64Field(value) {
    this.float64Field = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #float64FieldNullable = undefined;
  /**
   *
   * @returns {number}
   **/
  get float64FieldNullable() {
    return this.#float64FieldNullable;
  }
  /**
   *
   * @type {number}
   **/
  set float64FieldNullable(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#float64FieldNullable = parsedValue;
    }
  }
  setFloat64FieldNullable(value) {
    this.float64FieldNullable = value;
    return this;
  }
  /**
   *
   * @type {"Alpha" | "Beta"}
   **/
  #enumField;
  /**
   *
   * @returns {"Alpha" | "Beta"}
   **/
  get enumField() {
    return this.#enumField;
  }
  /**
   *
   * @type {"Alpha" | "Beta"}
   **/
  set enumField(value) {
    this.#enumField = value;
  }
  setEnumField(value) {
    this.enumField = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #enumFieldNullable = undefined;
  /**
   *
   * @returns {any}
   **/
  get enumFieldNullable() {
    return this.#enumFieldNullable;
  }
  /**
   *
   * @type {any}
   **/
  set enumFieldNullable(value) {
    this.#enumFieldNullable = value;
  }
  setEnumFieldNullable(value) {
    this.enumFieldNullable = value;
    return this;
  }
  /**
   *
   * @type {DataTypesDto.ObjectField}
   **/
  #objectField;
  /**
   *
   * @returns {DataTypesDto.ObjectField}
   **/
  get objectField() {
    return this.#objectField;
  }
  /**
   *
   * @type {DataTypesDto.ObjectField}
   **/
  set objectField(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof DataTypesDto.ObjectField) {
      this.#objectField = value;
    } else {
      this.#objectField = new DataTypesDto.ObjectField(value);
    }
  }
  setObjectField(value) {
    this.objectField = value;
    return this;
  }
  /**
   *
   * @type {DataTypesDto.ObjectFieldNullable}
   **/
  #objectFieldNullable = undefined;
  /**
   *
   * @returns {DataTypesDto.ObjectFieldNullable}
   **/
  get objectFieldNullable() {
    return this.#objectFieldNullable;
  }
  /**
   *
   * @type {DataTypesDto.ObjectFieldNullable}
   **/
  set objectFieldNullable(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof DataTypesDto.ObjectFieldNullable) {
      this.#objectFieldNullable = value;
    } else {
      this.#objectFieldNullable = new DataTypesDto.ObjectFieldNullable(value);
    }
  }
  setObjectFieldNullable(value) {
    this.objectFieldNullable = value;
    return this;
  }
  /**
   *
   * @type {DataTypesDto.ArrayField}
   **/
  #arrayField = [];
  /**
   *
   * @returns {DataTypesDto.ArrayField}
   **/
  get arrayField() {
    return this.#arrayField;
  }
  /**
   *
   * @type {DataTypesDto.ArrayField}
   **/
  set arrayField(value) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof DataTypesDto.ArrayField) {
      this.#arrayField = value;
    } else {
      this.#arrayField = value.map((item) => new DataTypesDto.ArrayField(item));
    }
  }
  setArrayField(value) {
    this.arrayField = value;
    return this;
  }
  /**
   *
   * @type {DataTypesDto.ArrayFieldNullable}
   **/
  #arrayFieldNullable = undefined;
  /**
   *
   * @returns {DataTypesDto.ArrayFieldNullable}
   **/
  get arrayFieldNullable() {
    return this.#arrayFieldNullable;
  }
  /**
   *
   * @type {DataTypesDto.ArrayFieldNullable}
   **/
  set arrayFieldNullable(value) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (
      value.length > 0 &&
      value[0] instanceof DataTypesDto.ArrayFieldNullable
    ) {
      this.#arrayFieldNullable = value;
    } else {
      this.#arrayFieldNullable = value.map(
        (item) => new DataTypesDto.ArrayFieldNullable(item),
      );
    }
  }
  setArrayFieldNullable(value) {
    this.arrayFieldNullable = value;
    return this;
  }
  /**
   *
   * @type {string[]}
   **/
  #sliceField = [];
  /**
   *
   * @returns {string[]}
   **/
  get sliceField() {
    return this.#sliceField;
  }
  /**
   *
   * @type {string[]}
   **/
  set sliceField(value) {
    this.#sliceField = value;
  }
  setSliceField(value) {
    this.sliceField = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #sliceFieldNullable = undefined;
  /**
   *
   * @returns {any}
   **/
  get sliceFieldNullable() {
    return this.#sliceFieldNullable;
  }
  /**
   *
   * @type {any}
   **/
  set sliceFieldNullable(value) {
    this.#sliceFieldNullable = value;
  }
  setSliceFieldNullable(value) {
    this.sliceFieldNullable = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #anyField = null;
  /**
   *
   * @returns {any}
   **/
  get anyField() {
    return this.#anyField;
  }
  /**
   *
   * @type {any}
   **/
  set anyField(value) {
    this.#anyField = value;
  }
  setAnyField(value) {
    this.anyField = value;
    return this;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  #mapField;
  /**
   *
   * @returns {{[key: string]: any}}
   **/
  get mapField() {
    return this.#mapField;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  set mapField(value) {
    this.#mapField = value;
  }
  setMapField(value) {
    this.mapField = value;
    return this;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  #mapFieldNullable = undefined;
  /**
   *
   * @returns {{[key: string]: any}}
   **/
  get mapFieldNullable() {
    return this.#mapFieldNullable;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  set mapFieldNullable(value) {
    this.#mapFieldNullable = value;
  }
  setMapFieldNullable(value) {
    this.mapFieldNullable = value;
    return this;
  }
  /**
   *
   * @type {GeoPoint}
   **/
  #complexField;
  /**
   *
   * @returns {GeoPoint}
   **/
  get complexField() {
    return this.#complexField;
  }
  /**
   *
   * @type {GeoPoint}
   **/
  set complexField(value) {
    this.#complexField = value;
  }
  setComplexField(value) {
    this.complexField = value;
    return this;
  }
  /**
   *
   * @type {CreateUserDtoTags}
   **/
  #oneRef;
  /**
   *
   * @returns {CreateUserDtoTags}
   **/
  get oneRef() {
    return this.#oneRef;
  }
  /**
   *
   * @type {CreateUserDtoTags}
   **/
  set oneRef(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof CreateUserDtoTags) {
      this.#oneRef = value;
    } else {
      this.#oneRef = new CreateUserDtoTags(value);
    }
  }
  setOneRef(value) {
    this.oneRef = value;
    return this;
  }
  /**
   *
   * @type {CreateUserDtoTags}
   **/
  #oneRefNullable = undefined;
  /**
   *
   * @returns {CreateUserDtoTags}
   **/
  get oneRefNullable() {
    return this.#oneRefNullable;
  }
  /**
   *
   * @type {CreateUserDtoTags}
   **/
  set oneRefNullable(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof CreateUserDtoTags) {
      this.#oneRefNullable = value;
    } else {
      this.#oneRefNullable = new CreateUserDtoTags(value);
    }
  }
  setOneRefNullable(value) {
    this.oneRefNullable = value;
    return this;
  }
  /**
   *
   * @type {CreateUserDtoTags[]}
   **/
  #collectionRef = [];
  /**
   *
   * @returns {CreateUserDtoTags[]}
   **/
  get collectionRef() {
    return this.#collectionRef;
  }
  /**
   *
   * @type {CreateUserDtoTags[]}
   **/
  set collectionRef(value) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof CreateUserDtoTags) {
      this.#collectionRef = value;
    } else {
      this.#collectionRef = value.map((item) => new CreateUserDtoTags(item));
    }
  }
  setCollectionRef(value) {
    this.collectionRef = value;
    return this;
  }
  /**
   *
   * @type {CreateUserDtoTags[]}
   **/
  #collectionRefNullable = undefined;
  /**
   *
   * @returns {CreateUserDtoTags[]}
   **/
  get collectionRefNullable() {
    return this.#collectionRefNullable;
  }
  /**
   *
   * @type {CreateUserDtoTags[]}
   **/
  set collectionRefNullable(value) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof CreateUserDtoTags) {
      this.#collectionRefNullable = value;
    } else {
      this.#collectionRefNullable = value.map(
        (item) => new CreateUserDtoTags(item),
      );
    }
  }
  setCollectionRefNullable(value) {
    this.collectionRefNullable = value;
    return this;
  }
  /**
   * The base class definition for objectField
   **/
  static ObjectField = class ObjectField {
    /**
     *
     * @type {string}
     **/
    #nestedString = "";
    /**
     *
     * @returns {string}
     **/
    get nestedString() {
      return this.#nestedString;
    }
    /**
     *
     * @type {string}
     **/
    set nestedString(value) {
      this.#nestedString = String(value);
    }
    setNestedString(value) {
      this.nestedString = value;
      return this;
    }
    /**
     *
     * @type {number}
     **/
    #nestedInt = 0;
    /**
     *
     * @returns {number}
     **/
    get nestedInt() {
      return this.#nestedInt;
    }
    /**
     *
     * @type {number}
     **/
    set nestedInt(value) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#nestedInt = parsedValue;
      }
    }
    setNestedInt(value) {
      this.nestedInt = value;
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
      if (d.nestedString !== undefined) {
        this.nestedString = d.nestedString;
      }
      if (d.nestedInt !== undefined) {
        this.nestedInt = d.nestedInt;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        nestedString: this.#nestedString,
        nestedInt: this.#nestedInt,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        nestedString: "nestedString",
        nestedInt: "nestedInt",
      };
    }
    /**
     * Creates an instance of DataTypesDto.ObjectField, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new DataTypesDto.ObjectField(possibleDtoObject);
    }
    /**
     * Creates an instance of DataTypesDto.ObjectField, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new DataTypesDto.ObjectField(partialDtoObject);
    }
    copyWith(partial) {
      return new DataTypesDto.ObjectField({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new DataTypesDto.ObjectField(this.toJSON());
    }
  };
  /**
   * The base class definition for objectFieldNullable
   **/
  static ObjectFieldNullable = class ObjectFieldNullable {
    /**
     *
     * @type {string}
     **/
    #optionalString = "";
    /**
     *
     * @returns {string}
     **/
    get optionalString() {
      return this.#optionalString;
    }
    /**
     *
     * @type {string}
     **/
    set optionalString(value) {
      this.#optionalString = String(value);
    }
    setOptionalString(value) {
      this.optionalString = value;
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
      if (d.optionalString !== undefined) {
        this.optionalString = d.optionalString;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        optionalString: this.#optionalString,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        optionalString: "optionalString",
      };
    }
    /**
     * Creates an instance of DataTypesDto.ObjectFieldNullable, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new DataTypesDto.ObjectFieldNullable(possibleDtoObject);
    }
    /**
     * Creates an instance of DataTypesDto.ObjectFieldNullable, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new DataTypesDto.ObjectFieldNullable(partialDtoObject);
    }
    copyWith(partial) {
      return new DataTypesDto.ObjectFieldNullable({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone() {
      return new DataTypesDto.ObjectFieldNullable(this.toJSON());
    }
  };
  /**
   * The base class definition for arrayField
   **/
  static ArrayField = class ArrayField {
    /**
     *
     * @type {string}
     **/
    #key = "";
    /**
     *
     * @returns {string}
     **/
    get key() {
      return this.#key;
    }
    /**
     *
     * @type {string}
     **/
    set key(value) {
      this.#key = String(value);
    }
    setKey(value) {
      this.key = value;
      return this;
    }
    /**
     *
     * @type {number}
     **/
    #value = 0;
    /**
     *
     * @returns {number}
     **/
    get value() {
      return this.#value;
    }
    /**
     *
     * @type {number}
     **/
    set value(value) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#value = parsedValue;
      }
    }
    setValue(value) {
      this.value = value;
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
      if (d.key !== undefined) {
        this.key = d.key;
      }
      if (d.value !== undefined) {
        this.value = d.value;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        key: this.#key,
        value: this.#value,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        key: "key",
        value: "value",
      };
    }
    /**
     * Creates an instance of DataTypesDto.ArrayField, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new DataTypesDto.ArrayField(possibleDtoObject);
    }
    /**
     * Creates an instance of DataTypesDto.ArrayField, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new DataTypesDto.ArrayField(partialDtoObject);
    }
    copyWith(partial) {
      return new DataTypesDto.ArrayField({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new DataTypesDto.ArrayField(this.toJSON());
    }
  };
  /**
   * The base class definition for arrayFieldNullable
   **/
  static ArrayFieldNullable = class ArrayFieldNullable {
    /**
     *
     * @type {string}
     **/
    #key = "";
    /**
     *
     * @returns {string}
     **/
    get key() {
      return this.#key;
    }
    /**
     *
     * @type {string}
     **/
    set key(value) {
      this.#key = String(value);
    }
    setKey(value) {
      this.key = value;
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
      if (d.key !== undefined) {
        this.key = d.key;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        key: this.#key,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        key: "key",
      };
    }
    /**
     * Creates an instance of DataTypesDto.ArrayFieldNullable, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new DataTypesDto.ArrayFieldNullable(possibleDtoObject);
    }
    /**
     * Creates an instance of DataTypesDto.ArrayFieldNullable, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new DataTypesDto.ArrayFieldNullable(partialDtoObject);
    }
    copyWith(partial) {
      return new DataTypesDto.ArrayFieldNullable({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone() {
      return new DataTypesDto.ArrayFieldNullable(this.toJSON());
    }
  };
  constructor(data) {
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
    if (d.stringField !== undefined) {
      this.stringField = d.stringField;
    }
    if (d.stringFieldNullable !== undefined) {
      this.stringFieldNullable = d.stringFieldNullable;
    }
    if (d.boolField !== undefined) {
      this.boolField = d.boolField;
    }
    if (d.boolFieldNullable !== undefined) {
      this.boolFieldNullable = d.boolFieldNullable;
    }
    if (d.intField !== undefined) {
      this.intField = d.intField;
    }
    if (d.intFieldNullable !== undefined) {
      this.intFieldNullable = d.intFieldNullable;
    }
    if (d.int32Field !== undefined) {
      this.int32Field = d.int32Field;
    }
    if (d.int32FieldNullable !== undefined) {
      this.int32FieldNullable = d.int32FieldNullable;
    }
    if (d.int64Field !== undefined) {
      this.int64Field = d.int64Field;
    }
    if (d.int64FieldNullable !== undefined) {
      this.int64FieldNullable = d.int64FieldNullable;
    }
    if (d.float32Field !== undefined) {
      this.float32Field = d.float32Field;
    }
    if (d.float32FieldNullable !== undefined) {
      this.float32FieldNullable = d.float32FieldNullable;
    }
    if (d.float64Field !== undefined) {
      this.float64Field = d.float64Field;
    }
    if (d.float64FieldNullable !== undefined) {
      this.float64FieldNullable = d.float64FieldNullable;
    }
    if (d.enumField !== undefined) {
      this.enumField = d.enumField;
    }
    if (d.enumFieldNullable !== undefined) {
      this.enumFieldNullable = d.enumFieldNullable;
    }
    if (d.objectField !== undefined) {
      this.objectField = d.objectField;
    }
    if (d.objectFieldNullable !== undefined) {
      this.objectFieldNullable = d.objectFieldNullable;
    }
    if (d.arrayField !== undefined) {
      this.arrayField = d.arrayField;
    }
    if (d.arrayFieldNullable !== undefined) {
      this.arrayFieldNullable = d.arrayFieldNullable;
    }
    if (d.sliceField !== undefined) {
      this.sliceField = d.sliceField;
    }
    if (d.sliceFieldNullable !== undefined) {
      this.sliceFieldNullable = d.sliceFieldNullable;
    }
    if (d.anyField !== undefined) {
      this.anyField = d.anyField;
    }
    if (d.mapField !== undefined) {
      this.mapField = d.mapField;
    }
    if (d.mapFieldNullable !== undefined) {
      this.mapFieldNullable = d.mapFieldNullable;
    }
    if (d.complexField !== undefined) {
      this.complexField = d.complexField;
    }
    if (d.oneRef !== undefined) {
      this.oneRef = d.oneRef;
    }
    if (d.oneRefNullable !== undefined) {
      this.oneRefNullable = d.oneRefNullable;
    }
    if (d.collectionRef !== undefined) {
      this.collectionRef = d.collectionRef;
    }
    if (d.collectionRefNullable !== undefined) {
      this.collectionRefNullable = d.collectionRefNullable;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data;
    if (!(d.objectField instanceof DataTypesDto.ObjectField)) {
      this.objectField = new DataTypesDto.ObjectField(d.objectField || {});
    }
    if (!(d.oneRef instanceof CreateUserDtoTags)) {
      this.oneRef = new CreateUserDtoTags(d.oneRef || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      stringField: this.#stringField,
      stringFieldNullable: this.#stringFieldNullable,
      boolField: this.#boolField,
      boolFieldNullable: this.#boolFieldNullable,
      intField: this.#intField,
      intFieldNullable: this.#intFieldNullable,
      int32Field: this.#int32Field,
      int32FieldNullable: this.#int32FieldNullable,
      int64Field: this.#int64Field,
      int64FieldNullable: this.#int64FieldNullable,
      float32Field: this.#float32Field,
      float32FieldNullable: this.#float32FieldNullable,
      float64Field: this.#float64Field,
      float64FieldNullable: this.#float64FieldNullable,
      enumField: this.#enumField,
      enumFieldNullable: this.#enumFieldNullable,
      objectField: this.#objectField,
      objectFieldNullable: this.#objectFieldNullable,
      arrayField: this.#arrayField,
      arrayFieldNullable: this.#arrayFieldNullable,
      sliceField: this.#sliceField,
      sliceFieldNullable: this.#sliceFieldNullable,
      anyField: this.#anyField,
      mapField: this.#mapField,
      mapFieldNullable: this.#mapFieldNullable,
      complexField: this.#complexField,
      oneRef: this.#oneRef,
      oneRefNullable: this.#oneRefNullable,
      collectionRef: this.#collectionRef,
      collectionRefNullable: this.#collectionRefNullable,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      stringField: "stringField",
      stringFieldNullable: "stringFieldNullable",
      boolField: "boolField",
      boolFieldNullable: "boolFieldNullable",
      intField: "intField",
      intFieldNullable: "intFieldNullable",
      int32Field: "int32Field",
      int32FieldNullable: "int32FieldNullable",
      int64Field: "int64Field",
      int64FieldNullable: "int64FieldNullable",
      float32Field: "float32Field",
      float32FieldNullable: "float32FieldNullable",
      float64Field: "float64Field",
      float64FieldNullable: "float64FieldNullable",
      enumField: "enumField",
      enumFieldNullable: "enumFieldNullable",
      objectField$: "objectField",
      get objectField() {
        return withPrefix("objectField", DataTypesDto.ObjectField.Fields);
      },
      objectFieldNullable$: "objectFieldNullable",
      get objectFieldNullable() {
        return withPrefix(
          "objectFieldNullable",
          DataTypesDto.ObjectFieldNullable.Fields,
        );
      },
      arrayField$: "arrayField",
      get arrayField() {
        return withPrefix("arrayField[:i]", DataTypesDto.ArrayField.Fields);
      },
      arrayFieldNullable$: "arrayFieldNullable",
      get arrayFieldNullable() {
        return withPrefix(
          "arrayFieldNullable[:i]",
          DataTypesDto.ArrayFieldNullable.Fields,
        );
      },
      sliceField$: "sliceField",
      get sliceField() {
        return "sliceField[:i]";
      },
      sliceFieldNullable: "sliceFieldNullable",
      anyField: "anyField",
      mapField: "mapField",
      mapFieldNullable: "mapFieldNullable",
      complexField: "complexField",
      oneRef$: "oneRef",
      get oneRef() {
        return withPrefix("oneRef", CreateUserDtoTags.Fields);
      },
      oneRefNullable: "oneRefNullable",
      collectionRef$: "collectionRef",
      get collectionRef() {
        return withPrefix("collectionRef[:i]", CreateUserDtoTags.Fields);
      },
      collectionRefNullable$: "collectionRefNullable",
      get collectionRefNullable() {
        return withPrefix("collectionRefNullable", CreateUserDtoTags.Fields);
      },
    };
  }
  /**
   * Creates an instance of DataTypesDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new DataTypesDto(possibleDtoObject);
  }
  /**
   * Creates an instance of DataTypesDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new DataTypesDto(partialDtoObject);
  }
  copyWith(partial) {
    return new DataTypesDto({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new DataTypesDto(this.toJSON());
  }
}
