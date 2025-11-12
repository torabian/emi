/**
 * The base class definition for yokeMessageDto
 **/
export class YokeMessageDto {
  /**
   *
   * @type {number}
   **/
  #pitch = 0.0;
  /**
   *
   * @returns {number}
   **/
  get pitch() {
    return this.#pitch;
  }
  /**
   *
   * @type {number}
   **/
  set pitch(value) {}
  setPitch(value) {
    this.pitch = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #role = 0.0;
  /**
   *
   * @returns {number}
   **/
  get role() {
    return this.#role;
  }
  /**
   *
   * @type {number}
   **/
  set role(value) {}
  setRole(value) {
    this.role = value;
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
          typeof data
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
    if (d.pitch !== undefined) {
      this.pitch = d.pitch;
    }
    if (d.role !== undefined) {
      this.role = d.role;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      pitch: this.#pitch,
      role: this.#role,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      pitch: "pitch",
      role: "role",
    };
  }
  /**
   * Creates an instance of YokeMessageDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new YokeMessageDto(possibleDtoObject);
  }
  /**
   * Creates an instance of YokeMessageDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new YokeMessageDto(partialDtoObject);
  }
  copyWith(partial) {
    return new YokeMessageDto({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new YokeMessageDto(this.toJSON());
  }
}
