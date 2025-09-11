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
  min?: number;
  /**
   * Maximum number which can be generated
   * @type {number}
   **/
  max?: number;
  /**
   * How many numbers you want to be generated based on maximum and minimum
   * @type {number}
   **/
  count?: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SingleDtoDtoType {}
