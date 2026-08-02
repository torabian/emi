/**
 * The base class definition for entity4OptionalDto
 **/
export class Entity4OptionalDto {
  /**
   *
   * @type {string}
   **/
  #note = undefined;
  /**
   *
   * @returns {string}
   **/
  get note() {
    return this.#note;
  }
  /**
   *
   * @type {string}
   **/
  set note(value) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#note = correctType ? value : String(value);
  }
  setNote(value) {
    this.note = value;
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
    if (d.note !== undefined) {
      this.note = d.note;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      note: this.#note,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      note: "note",
    };
  }
  /**
   * Creates an instance of Entity4OptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new Entity4OptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of Entity4OptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new Entity4OptionalDto(partialDtoObject);
  }
  copyWith(partial) {
    return new Entity4OptionalDto({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new Entity4OptionalDto(this.toJSON());
  }
}
