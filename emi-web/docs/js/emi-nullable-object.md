---
sidebar_position: 4
---

# Nullable object vs non-nullable

Emi definition allows for nullable object vs non-nullable objects. In case an object is not nullable, should always be present and initialised
by the parent class upon instantiation. For types, simply it would indicate that it needs to be present.

 
```yaml
name: nullableResponseAction
fields:
  - name: mother
    type: object
    description: Mother info should be present
    fields:
      - name: firstName
        type: string
  - name: father
    type: object?
    description: Father name is not essential for some goverment papers.
    fields:
      - name: firstName
        type: string
  - name: firstUncle
    type: one
    target: UncleDto
    description: Uncle is a separate dto, therefor we use that entity
  - name: secondUncle
    type: one?
    target: UncleDto
    description: Second uncle is optional

```


```ts
import { UncleDto } from "./UncleDto";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for nullableResponseActionDto
 **/
export class NullableResponseActionDto {
  /**
   * Mother info should be present
   * @type {NullableResponseActionDto.Mother}
   **/
  #mother!: InstanceType<typeof NullableResponseActionDto.Mother>;
  /**
   * Mother info should be present
   * @returns {NullableResponseActionDto.Mother}
   **/
  get mother() {
    return this.#mother;
  }
  /**
   * Mother info should be present
   * @type {NullableResponseActionDto.Mother}
   **/
  set mother(value: InstanceType<typeof NullableResponseActionDto.Mother>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof NullableResponseActionDto.Mother) {
      this.#mother = value;
    } else {
      this.#mother = new NullableResponseActionDto.Mother(value);
    }
  }
  setMother(value: InstanceType<typeof NullableResponseActionDto.Mother>) {
    this.mother = value;
    return this;
  }
  /**
   * Father name is not essential for some goverment papers.
   * @type {NullableResponseActionDto.Father}
   **/
  #father?:
    | InstanceType<typeof NullableResponseActionDto.Father>
    | null
    | undefined
    | null = undefined;
  /**
   * Father name is not essential for some goverment papers.
   * @returns {NullableResponseActionDto.Father}
   **/
  get father() {
    return this.#father;
  }
  /**
   * Father name is not essential for some goverment papers.
   * @type {NullableResponseActionDto.Father}
   **/
  set father(
    value:
      | InstanceType<typeof NullableResponseActionDto.Father>
      | null
      | undefined
      | null
      | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof NullableResponseActionDto.Father) {
      this.#father = value;
    } else {
      this.#father = new NullableResponseActionDto.Father(value);
    }
  }
  setFather(
    value:
      | InstanceType<typeof NullableResponseActionDto.Father>
      | null
      | undefined
      | null
      | undefined,
  ) {
    this.father = value;
    return this;
  }
  /**
   * Uncle is a separate dto, therefor we use that entity
   * @type {UncleDto}
   **/
  #firstUncle!: UncleDto;
  /**
   * Uncle is a separate dto, therefor we use that entity
   * @returns {UncleDto}
   **/
  get firstUncle() {
    return this.#firstUncle;
  }
  /**
   * Uncle is a separate dto, therefor we use that entity
   * @type {UncleDto}
   **/
  set firstUncle(value: UncleDto) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UncleDto) {
      this.#firstUncle = value;
    } else {
      this.#firstUncle = new UncleDto(value);
    }
  }
  setFirstUncle(value: UncleDto) {
    this.firstUncle = value;
    return this;
  }
  /**
   * Second uncle is optional
   * @type {UncleDto}
   **/
  #secondUncle?: UncleDto | null = undefined;
  /**
   * Second uncle is optional
   * @returns {UncleDto}
   **/
  get secondUncle() {
    return this.#secondUncle;
  }
  /**
   * Second uncle is optional
   * @type {UncleDto}
   **/
  set secondUncle(value: UncleDto | null | undefined) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UncleDto) {
      this.#secondUncle = value;
    } else {
      this.#secondUncle = new UncleDto(value);
    }
  }
  setSecondUncle(value: UncleDto | null | undefined) {
    this.secondUncle = value;
    return this;
  }
  /**
   * The base class definition for mother
   **/
  static Mother = class Mother {
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
      const g = globalThis as any;
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
      const d = data as Partial<Mother>;
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
    /**
     * Creates an instance of NullableResponseActionDto.Mother, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: NullableResponseActionDtoType.MotherType) {
      return new NullableResponseActionDto.Mother(possibleDtoObject);
    }
    /**
     * Creates an instance of NullableResponseActionDto.Mother, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: Partial<NullableResponseActionDtoType.MotherType>,
    ) {
      return new NullableResponseActionDto.Mother(partialDtoObject);
    }
    copyWith(
      partial: Partial<NullableResponseActionDtoType.MotherType>,
    ): InstanceType<typeof NullableResponseActionDto.Mother> {
      return new NullableResponseActionDto.Mother({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof NullableResponseActionDto.Mother> {
      return new NullableResponseActionDto.Mother(this.toJSON());
    }
  };
  /**
   * The base class definition for father
   **/
  static Father = class Father {
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
      const g = globalThis as any;
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
      const d = data as Partial<Father>;
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
    /**
     * Creates an instance of NullableResponseActionDto.Father, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: NullableResponseActionDtoType.FatherType) {
      return new NullableResponseActionDto.Father(possibleDtoObject);
    }
    /**
     * Creates an instance of NullableResponseActionDto.Father, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: Partial<NullableResponseActionDtoType.FatherType>,
    ) {
      return new NullableResponseActionDto.Father(partialDtoObject);
    }
    copyWith(
      partial: Partial<NullableResponseActionDtoType.FatherType>,
    ): InstanceType<typeof NullableResponseActionDto.Father> {
      return new NullableResponseActionDto.Father({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof NullableResponseActionDto.Father> {
      return new NullableResponseActionDto.Father(this.toJSON());
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
    const g = globalThis as any;
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
    const d = data as Partial<NullableResponseActionDto>;
    if (d.mother !== undefined) {
      this.mother = d.mother;
    }
    if (d.father !== undefined) {
      this.father = d.father;
    }
    if (d.firstUncle !== undefined) {
      this.firstUncle = d.firstUncle;
    }
    if (d.secondUncle !== undefined) {
      this.secondUncle = d.secondUncle;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<NullableResponseActionDto>;
    if (!(d.mother instanceof NullableResponseActionDto.Mother)) {
      this.mother = new NullableResponseActionDto.Mother(d.mother || {});
    }
    if (!(d.firstUncle instanceof UncleDto)) {
      this.firstUncle = new UncleDto(d.firstUncle || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      mother: this.#mother,
      father: this.#father,
      firstUncle: this.#firstUncle,
      secondUncle: this.#secondUncle,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      mother$: "mother",
      get mother() {
        return withPrefix("mother", NullableResponseActionDto.Mother.Fields);
      },
      father$: "father",
      get father() {
        return withPrefix("father", NullableResponseActionDto.Father.Fields);
      },
      firstUncle$: "firstUncle",
      get firstUncle() {
        return withPrefix("firstUncle", UncleDto.Fields);
      },
      secondUncle: "secondUncle",
    };
  }
  /**
   * Creates an instance of NullableResponseActionDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: NullableResponseActionDtoType) {
    return new NullableResponseActionDto(possibleDtoObject);
  }
  /**
   * Creates an instance of NullableResponseActionDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: Partial<NullableResponseActionDtoType>) {
    return new NullableResponseActionDto(partialDtoObject);
  }
  copyWith(
    partial: Partial<NullableResponseActionDtoType>,
  ): InstanceType<typeof NullableResponseActionDto> {
    return new NullableResponseActionDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof NullableResponseActionDto> {
    return new NullableResponseActionDto(this.toJSON());
  }
}
export abstract class NullableResponseActionDtoFactory {
  abstract create(data: unknown): NullableResponseActionDto;
}
/**
 * The base type definition for nullableResponseActionDto
 **/
export type NullableResponseActionDtoType = {
  /**
   * Mother info should be present
   * @type {NullableResponseActionDtoType.MotherType}
   **/
  mother: NullableResponseActionDtoType.MotherType;
  /**
   * Father name is not essential for some goverment papers.
   * @type {NullableResponseActionDtoType.FatherType}
   **/
  father?: NullableResponseActionDtoType.FatherType;
  /**
   * Uncle is a separate dto, therefor we use that entity
   * @type {UncleDto}
   **/
  firstUncle: UncleDto;
  /**
   * Second uncle is optional
   * @type {UncleDto}
   **/
  secondUncle?: UncleDto;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace NullableResponseActionDtoType {
  /**
   * The base type definition for motherType
   **/
  export type MotherType = {
    /**
     *
     * @type {string}
     **/
    firstName: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace MotherType {}
  /**
   * The base type definition for fatherType
   **/
  export type FatherType = {
    /**
     *
     * @type {string}
     **/
    firstName: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace FatherType {}
}

```