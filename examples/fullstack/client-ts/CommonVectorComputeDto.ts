import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for commonVectorComputeDto
 **/
export class CommonVectorComputeDto {
  /**
   *
   * @type {number[]}
   **/
  #initialVector1: number[] = [];
  /**
   *
   * @returns {number[]}
   **/
  get initialVector1() {
    return this.#initialVector1;
  }
  /**
   *
   * @type {number[]}
   **/
  set initialVector1(value: number[]) {}
  setInitialVector1(value: number[]) {
    this.initialVector1 = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #value?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   *
   * @type {string}
   **/
  set value(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#value = correctType ? value : String(value);
  }
  setValue(value: string | null | undefined) {
    this.value = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #valuex: string = "";
  /**
   *
   * @returns {string}
   **/
  get valuex() {
    return this.#valuex;
  }
  /**
   *
   * @type {string}
   **/
  set valuex(value: string) {
    this.#valuex = String(value);
  }
  setValuex(value: string) {
    this.valuex = value;
    return this;
  }
  /**
   *
   * @type {number[]}
   **/
  #initialVector2: number[] = [];
  /**
   *
   * @returns {number[]}
   **/
  get initialVector2() {
    return this.#initialVector2;
  }
  /**
   *
   * @type {number[]}
   **/
  set initialVector2(value: number[]) {}
  setInitialVector2(value: number[]) {
    this.initialVector2 = value;
    return this;
  }
  /**
   *
   * @type {CommonVectorComputeDto.FieldTypeArray}
   **/
  #fieldTypeArray: InstanceType<
    typeof CommonVectorComputeDto.FieldTypeArray
  >[] = [];
  /**
   *
   * @returns {CommonVectorComputeDto.FieldTypeArray}
   **/
  get fieldTypeArray() {
    return this.#fieldTypeArray;
  }
  /**
   *
   * @type {CommonVectorComputeDto.FieldTypeArray}
   **/
  set fieldTypeArray(
    value: InstanceType<typeof CommonVectorComputeDto.FieldTypeArray>[],
  ) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (
      value.length > 0 &&
      value[0] instanceof CommonVectorComputeDto.FieldTypeArray
    ) {
      this.#fieldTypeArray = value;
    } else {
      this.#fieldTypeArray = value.map(
        (item) => new CommonVectorComputeDto.FieldTypeArray(item),
      );
    }
  }
  setFieldTypeArray(
    value: InstanceType<typeof CommonVectorComputeDto.FieldTypeArray>[],
  ) {
    this.fieldTypeArray = value;
    return this;
  }
  /**
   *
   * @type {string[]}
   **/
  #fieldTypeSlice: string[] = [];
  /**
   *
   * @returns {string[]}
   **/
  get fieldTypeSlice() {
    return this.#fieldTypeSlice;
  }
  /**
   *
   * @type {string[]}
   **/
  set fieldTypeSlice(value: string[]) {}
  setFieldTypeSlice(value: string[]) {
    this.fieldTypeSlice = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #fieldInt: number = 0;
  /**
   *
   * @returns {number}
   **/
  get fieldInt() {
    return this.#fieldInt;
  }
  /**
   *
   * @type {number}
   **/
  set fieldInt(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#fieldInt = parsedValue;
    }
  }
  setFieldInt(value: number) {
    this.fieldInt = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #fieldIntNullable?: number | null = undefined;
  /**
   *
   * @returns {number}
   **/
  get fieldIntNullable() {
    return this.#fieldIntNullable;
  }
  /**
   *
   * @type {number}
   **/
  set fieldIntNullable(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#fieldIntNullable = parsedValue;
    }
  }
  setFieldIntNullable(value: number | null | undefined) {
    this.fieldIntNullable = value;
    return this;
  }
  /**
   *
   * @type {Money}
   **/
  #complexMoney!: Money;
  /**
   *
   * @returns {Money}
   **/
  get complexMoney() {
    return this.#complexMoney;
  }
  /**
   *
   * @type {Money}
   **/
  set complexMoney(value: Money) {}
  setComplexMoney(value: Money) {
    this.complexMoney = value;
    return this;
  }
  /**
   * The base class definition for fieldTypeArray
   **/
  static FieldTypeArray = class FieldTypeArray {
    /**
     *
     * @type {string}
     **/
    #arrayField1: string = "";
    /**
     *
     * @returns {string}
     **/
    get arrayField1() {
      return this.#arrayField1;
    }
    /**
     *
     * @type {string}
     **/
    set arrayField1(value: string) {
      this.#arrayField1 = String(value);
    }
    setArrayField1(value: string) {
      this.arrayField1 = value;
      return this;
    }
    /**
     *
     * @type {CommonVectorComputeDto.FieldTypeArray.ArrayField2}
     **/
    #arrayField2: InstanceType<
      typeof CommonVectorComputeDto.FieldTypeArray.ArrayField2
    >[] = [];
    /**
     *
     * @returns {CommonVectorComputeDto.FieldTypeArray.ArrayField2}
     **/
    get arrayField2() {
      return this.#arrayField2;
    }
    /**
     *
     * @type {CommonVectorComputeDto.FieldTypeArray.ArrayField2}
     **/
    set arrayField2(
      value: InstanceType<
        typeof CommonVectorComputeDto.FieldTypeArray.ArrayField2
      >[],
    ) {
      // For arrays, you only can pass arrays to the object
      if (!Array.isArray(value)) {
        return;
      }
      if (
        value.length > 0 &&
        value[0] instanceof CommonVectorComputeDto.FieldTypeArray.ArrayField2
      ) {
        this.#arrayField2 = value;
      } else {
        this.#arrayField2 = value.map(
          (item) => new CommonVectorComputeDto.FieldTypeArray.ArrayField2(item),
        );
      }
    }
    setArrayField2(
      value: InstanceType<
        typeof CommonVectorComputeDto.FieldTypeArray.ArrayField2
      >[],
    ) {
      this.arrayField2 = value;
      return this;
    }
    /**
     * The base class definition for arrayField2
     **/
    static ArrayField2 = class ArrayField2 {
      /**
       *
       * @type {string}
       **/
      #lastItem: string = "";
      /**
       *
       * @returns {string}
       **/
      get lastItem() {
        return this.#lastItem;
      }
      /**
       *
       * @type {string}
       **/
      set lastItem(value: string) {
        this.#lastItem = String(value);
      }
      setLastItem(value: string) {
        this.lastItem = value;
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
        const d = data as Partial<ArrayField2>;
        if (d.lastItem !== undefined) {
          this.lastItem = d.lastItem;
        }
      }
      /**
       *	Special toJSON override, since the field are private,
       *	Json stringify won't see them unless we mention it explicitly.
       **/
      toJSON() {
        return {
          lastItem: this.#lastItem,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          lastItem: "lastItem",
        };
      }
      /**
       * Creates an instance of CommonVectorComputeDto.FieldTypeArray.ArrayField2, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(
        possibleDtoObject: CommonVectorComputeDtoType.FieldTypeArrayType.ArrayField2Type,
      ) {
        return new CommonVectorComputeDto.FieldTypeArray.ArrayField2(
          possibleDtoObject,
        );
      }
      /**
       * Creates an instance of CommonVectorComputeDto.FieldTypeArray.ArrayField2, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(
        partialDtoObject: PartialDeep<CommonVectorComputeDtoType.FieldTypeArrayType.ArrayField2Type>,
      ) {
        return new CommonVectorComputeDto.FieldTypeArray.ArrayField2(
          partialDtoObject,
        );
      }
      copyWith(
        partial: PartialDeep<CommonVectorComputeDtoType.FieldTypeArrayType.ArrayField2Type>,
      ): InstanceType<
        typeof CommonVectorComputeDto.FieldTypeArray.ArrayField2
      > {
        return new CommonVectorComputeDto.FieldTypeArray.ArrayField2({
          ...this.toJSON(),
          ...partial,
        });
      }
      clone(): InstanceType<
        typeof CommonVectorComputeDto.FieldTypeArray.ArrayField2
      > {
        return new CommonVectorComputeDto.FieldTypeArray.ArrayField2(
          this.toJSON(),
        );
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
      const d = data as Partial<FieldTypeArray>;
      if (d.arrayField1 !== undefined) {
        this.arrayField1 = d.arrayField1;
      }
      if (d.arrayField2 !== undefined) {
        this.arrayField2 = d.arrayField2;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        arrayField1: this.#arrayField1,
        arrayField2: this.#arrayField2,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        arrayField1: "arrayField1",
        arrayField2$: "arrayField2",
        get arrayField2() {
          return withPrefix(
            "fieldTypeArray.arrayField2[:i]",
            CommonVectorComputeDto.FieldTypeArray.ArrayField2.Fields,
          );
        },
      };
    }
    /**
     * Creates an instance of CommonVectorComputeDto.FieldTypeArray, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(
      possibleDtoObject: CommonVectorComputeDtoType.FieldTypeArrayType,
    ) {
      return new CommonVectorComputeDto.FieldTypeArray(possibleDtoObject);
    }
    /**
     * Creates an instance of CommonVectorComputeDto.FieldTypeArray, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<CommonVectorComputeDtoType.FieldTypeArrayType>,
    ) {
      return new CommonVectorComputeDto.FieldTypeArray(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<CommonVectorComputeDtoType.FieldTypeArrayType>,
    ): InstanceType<typeof CommonVectorComputeDto.FieldTypeArray> {
      return new CommonVectorComputeDto.FieldTypeArray({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof CommonVectorComputeDto.FieldTypeArray> {
      return new CommonVectorComputeDto.FieldTypeArray(this.toJSON());
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
    const d = data as Partial<CommonVectorComputeDto>;
    if (d.initialVector1 !== undefined) {
      this.initialVector1 = d.initialVector1;
    }
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.valuex !== undefined) {
      this.valuex = d.valuex;
    }
    if (d.initialVector2 !== undefined) {
      this.initialVector2 = d.initialVector2;
    }
    if (d.fieldTypeArray !== undefined) {
      this.fieldTypeArray = d.fieldTypeArray;
    }
    if (d.fieldTypeSlice !== undefined) {
      this.fieldTypeSlice = d.fieldTypeSlice;
    }
    if (d.fieldInt !== undefined) {
      this.fieldInt = d.fieldInt;
    }
    if (d.fieldIntNullable !== undefined) {
      this.fieldIntNullable = d.fieldIntNullable;
    }
    if (d.complexMoney !== undefined) {
      this.complexMoney = d.complexMoney;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      initialVector1: this.#initialVector1,
      value: this.#value,
      valuex: this.#valuex,
      initialVector2: this.#initialVector2,
      fieldTypeArray: this.#fieldTypeArray,
      fieldTypeSlice: this.#fieldTypeSlice,
      fieldInt: this.#fieldInt,
      fieldIntNullable: this.#fieldIntNullable,
      complexMoney: this.#complexMoney,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      initialVector1$: "initialVector1",
      get initialVector1() {
        return "initialVector1[:i]";
      },
      value: "value",
      valuex: "valuex",
      initialVector2$: "initialVector2",
      get initialVector2() {
        return "initialVector2[:i]";
      },
      fieldTypeArray$: "fieldTypeArray",
      get fieldTypeArray() {
        return withPrefix(
          "fieldTypeArray[:i]",
          CommonVectorComputeDto.FieldTypeArray.Fields,
        );
      },
      fieldTypeSlice$: "fieldTypeSlice",
      get fieldTypeSlice() {
        return "fieldTypeSlice[:i]";
      },
      fieldInt: "fieldInt",
      fieldIntNullable: "fieldIntNullable",
      complexMoney: "complexMoney",
    };
  }
  /**
   * Creates an instance of CommonVectorComputeDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CommonVectorComputeDtoType) {
    return new CommonVectorComputeDto(possibleDtoObject);
  }
  /**
   * Creates an instance of CommonVectorComputeDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<CommonVectorComputeDtoType>) {
    return new CommonVectorComputeDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CommonVectorComputeDtoType>,
  ): InstanceType<typeof CommonVectorComputeDto> {
    return new CommonVectorComputeDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CommonVectorComputeDto> {
    return new CommonVectorComputeDto(this.toJSON());
  }
}
export abstract class CommonVectorComputeDtoFactory {
  abstract create(data: unknown): CommonVectorComputeDto;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for commonVectorComputeDto
 **/
export type CommonVectorComputeDtoType = {
  /**
   *
   * @type {number[]}
   **/
  initialVector1: number[];
  /**
   *
   * @type {string}
   **/
  value?: string;
  /**
   *
   * @type {string}
   **/
  valuex: string;
  /**
   *
   * @type {number[]}
   **/
  initialVector2: number[];
  /**
   *
   * @type {CommonVectorComputeDtoType.FieldTypeArrayType[]}
   **/
  fieldTypeArray: CommonVectorComputeDtoType.FieldTypeArrayType[];
  /**
   *
   * @type {string[]}
   **/
  fieldTypeSlice: string[];
  /**
   *
   * @type {number}
   **/
  fieldInt: number;
  /**
   *
   * @type {number}
   **/
  fieldIntNullable?: number;
  /**
   *
   * @type {Money}
   **/
  complexMoney: Money;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CommonVectorComputeDtoType {
  /**
   * The base type definition for fieldTypeArrayType
   **/
  export type FieldTypeArrayType = {
    /**
     *
     * @type {string}
     **/
    arrayField1: string;
    /**
     *
     * @type {CommonVectorComputeDtoType.FieldTypeArrayType.ArrayField2Type[]}
     **/
    arrayField2: CommonVectorComputeDtoType.FieldTypeArrayType.ArrayField2Type[];
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace FieldTypeArrayType {
    /**
     * The base type definition for arrayField2Type
     **/
    export type ArrayField2Type = {
      /**
       *
       * @type {string}
       **/
      lastItem: string;
    };
    // eslint-disable-next-line @typescript-eslint/no-namespace
    export namespace ArrayField2Type {}
  }
}
