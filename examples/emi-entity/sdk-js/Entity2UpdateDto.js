/**
 * The base class definition for entity2UpdateDto
 **/
export class Entity2UpdateDto {
  /**
   *
   * @type {string}
   **/
  #label2 = undefined;
  /**
   *
   * @returns {string}
   **/
  get label2() {
    return this.#label2;
  }
  /**
   *
   * @type {string}
   **/
  set label2(value) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#label2 = correctType ? value : String(value);
  }
  setLabel2(value) {
    this.label2 = value;
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
    if (d.label2 !== undefined) {
      this.label2 = d.label2;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      label2: this.#label2,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      label2: "label2",
    };
  }
  /**
   * Creates an instance of Entity2UpdateDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new Entity2UpdateDto(possibleDtoObject);
  }
  /**
   * Creates an instance of Entity2UpdateDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new Entity2UpdateDto(partialDtoObject);
  }
  copyWith(partial) {
    return new Entity2UpdateDto({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new Entity2UpdateDto(this.toJSON());
  }
}
