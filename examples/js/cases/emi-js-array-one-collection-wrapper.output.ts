import { MArray, MCollection, MOne } from "./sdk/common/operators";
import { User } from "./User";
import { type PartialDeep } from "./sdk/common/fetchx";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for anonymouse
 **/
export class Anonymouse {
  /**
   * array field, non-nullable
   * @type {Anonymouse.ArrayField}
   **/
  #arrayField: MArray<InstanceType<typeof Anonymouse.ArrayField>> = MArray.of(
    [],
  );
  /**
   * array field, non-nullable
   * @returns {Anonymouse.ArrayField}
   **/
  get arrayField() {
    return this.#arrayField;
  }
  /**
   * array field, non-nullable
   * @type {Anonymouse.ArrayField}
   **/
  set arrayField(
    value:
      | MArray<InstanceType<typeof Anonymouse.ArrayField>>
      | InstanceType<typeof Anonymouse.ArrayField>[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof Anonymouse.ArrayField) {
        this.#arrayField = MArray.of(value);
      } else {
        this.#arrayField = MArray.of(
          value.map((item) => new Anonymouse.ArrayField(item)),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#arrayField = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast<unknown>(value);
    if (ok) {
      this.#arrayField = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to arrayField, because it needs MArray instance or an Array.",
    );
  }
  setArrayField(
    value:
      | MArray<InstanceType<typeof Anonymouse.ArrayField>>
      | InstanceType<typeof Anonymouse.ArrayField>[],
  ) {
    this.arrayField = value;
    return this;
  }
  /**
   * arrayNullable field, non-nullable
   * @type {any}
   **/
  #arrayNullableField!: any;
  /**
   * arrayNullable field, non-nullable
   * @returns {any}
   **/
  get arrayNullableField() {
    return this.#arrayNullableField;
  }
  /**
   * arrayNullable field, non-nullable
   * @type {any}
   **/
  set arrayNullableField(value: any) {
    this.#arrayNullableField = value;
  }
  setArrayNullableField(value: any) {
    this.arrayNullableField = value;
    return this;
  }
  /**
   * collection field, non-nullable
   * @type {User[]}
   **/
  #collectionField: MCollection<User> = MCollection.of([]);
  /**
   * collection field, non-nullable
   * @returns {User[]}
   **/
  get collectionField() {
    return this.#collectionField;
  }
  /**
   * collection field, non-nullable
   * @type {User[]}
   **/
  set collectionField(value: MCollection<User> | InstanceType<typeof User>[]) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof User) {
        this.#collectionField = MCollection.of(value);
      } else {
        this.#collectionField = MCollection.of(
          value.map((item) => new User(item)),
        );
      }
      return;
    }
    // If the instance is already an MCollection, we assume it's all good.
    if (value instanceof MCollection) {
      this.#collectionField = value;
      return;
    }
    // If the value is not array, and is not a MCollection, we need to be consider,
    // it might be eligible to be casted into MCollection.
    const { ok, value: mcastValue } = MCollection.cast<unknown>(value);
    if (ok) {
      this.#collectionField = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to collectionField, because it needs MCollection instance or an Array.",
    );
  }
  setCollectionField(value: MCollection<User> | InstanceType<typeof User>[]) {
    this.collectionField = value;
    return this;
  }
  /**
   * collectionNullable field, non-nullable
   * @type {User[]}
   **/
  #collectionNullableField?: MCollection<User> | null = undefined;
  /**
   * collectionNullable field, non-nullable
   * @returns {User[]}
   **/
  get collectionNullableField() {
    return this.#collectionNullableField;
  }
  /**
   * collectionNullable field, non-nullable
   * @type {User[]}
   **/
  set collectionNullableField(
    value: MCollection<User> | InstanceType<typeof User>[] | null | undefined,
  ) {
    // For nullable collection, we allow explicit undefined or null values
    if (value === null || value === undefined) {
      this.#collectionNullableField = value;
      return;
    }
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof User) {
        this.#collectionNullableField = MCollection.of(value);
      } else {
        this.#collectionNullableField = MCollection.of(
          value.map((item) => new User(item)),
        );
      }
      return;
    }
    // If the instance is already an MCollection, we assume it's all good.
    if (value instanceof MCollection) {
      this.#collectionNullableField = value;
      return;
    }
    // If the value is not array, and is not a MCollection, we need to be consider,
    // it might be eligible to be casted into MCollection.
    const { ok, value: mcastValue } = MCollection.cast<unknown>(value);
    if (ok) {
      this.#collectionNullableField = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to collectionNullableField, because it needs MCollection instance or an Array.",
    );
  }
  setCollectionNullableField(
    value: MCollection<User> | InstanceType<typeof User>[] | null | undefined,
  ) {
    this.collectionNullableField = value;
    return this;
  }
  /**
   * one field, non-nullable
   * @type {User}
   **/
  #oneField!: MOne<User>;
  /**
   * one field, non-nullable
   * @returns {User}
   **/
  get oneField() {
    return this.#oneField;
  }
  /**
   * one field, non-nullable
   * @type {User}
   **/
  set oneField(value: MOne<User> | InstanceType<typeof User>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#oneField = value;
    } else if (value instanceof User) {
      this.#oneField = MOne.of(value);
    } else {
      this.#oneField = MOne.of(new User(value));
    }
  }
  setOneField(value: MOne<User> | InstanceType<typeof User>) {
    this.oneField = value;
    return this;
  }
  /**
   * oneNullable field, non-nullable
   * @type {User}
   **/
  #oneNullableField?: MOne<User> | null = undefined;
  /**
   * oneNullable field, non-nullable
   * @returns {User}
   **/
  get oneNullableField() {
    return this.#oneNullableField;
  }
  /**
   * oneNullable field, non-nullable
   * @type {User}
   **/
  set oneNullableField(
    value: MOne<User> | InstanceType<typeof User> | null | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#oneNullableField = value;
    } else if (value instanceof User) {
      this.#oneNullableField = MOne.of(value);
    } else {
      this.#oneNullableField = MOne.of(new User(value));
    }
  }
  setOneNullableField(
    value: MOne<User> | InstanceType<typeof User> | null | undefined,
  ) {
    this.oneNullableField = value;
    return this;
  }
  /**
   * The base class definition for arrayField
   **/
  static ArrayField = class ArrayField {
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
      const d = data as Partial<ArrayField>;
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {};
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {};
    }
    /**
     * Creates an instance of Anonymouse.ArrayField, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: AnonymouseType.ArrayFieldType) {
      return new Anonymouse.ArrayField(possibleDtoObject);
    }
    /**
     * Creates an instance of Anonymouse.ArrayField, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<AnonymouseType.ArrayFieldType>) {
      return new Anonymouse.ArrayField(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<AnonymouseType.ArrayFieldType>,
    ): InstanceType<typeof Anonymouse.ArrayField> {
      return new Anonymouse.ArrayField({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof Anonymouse.ArrayField> {
      return new Anonymouse.ArrayField(this.toJSON());
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
    const d = data as Partial<Anonymouse>;
    if (d.arrayField !== undefined) {
      this.arrayField = d.arrayField;
    }
    if (d.arrayNullableField !== undefined) {
      this.arrayNullableField = d.arrayNullableField;
    }
    if (d.collectionField !== undefined) {
      this.collectionField = d.collectionField;
    }
    if (d.collectionNullableField !== undefined) {
      this.collectionNullableField = d.collectionNullableField;
    }
    if (d.oneField !== undefined) {
      this.oneField = d.oneField;
    }
    if (d.oneNullableField !== undefined) {
      this.oneNullableField = d.oneNullableField;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<Anonymouse>;
    if (!(d.oneField instanceof User)) {
      this.oneField = MOne.of(new User(d.oneField || {}));
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      arrayField: this.#arrayField,
      arrayNullableField: this.#arrayNullableField,
      collectionField: this.#collectionField,
      collectionNullableField: this.#collectionNullableField,
      oneField: this.#oneField,
      oneNullableField: this.#oneNullableField,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      arrayField$: "arrayField",
      get arrayField() {
        return withPrefix("arrayField[:i]", Anonymouse.ArrayField.Fields);
      },
      arrayNullableField: "arrayNullableField",
      collectionField$: "collectionField",
      get collectionField() {
        return withPrefix("collectionField[:i]", User.Fields);
      },
      collectionNullableField$: "collectionNullableField",
      get collectionNullableField() {
        return withPrefix("collectionNullableField", User.Fields);
      },
      oneField$: "oneField",
      get oneField() {
        return withPrefix("oneField", User.Fields);
      },
      oneNullableField: "oneNullableField",
    };
  }
  /**
   * Creates an instance of Anonymouse, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AnonymouseType) {
    return new Anonymouse(possibleDtoObject);
  }
  /**
   * Creates an instance of Anonymouse, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AnonymouseType>) {
    return new Anonymouse(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AnonymouseType>,
  ): InstanceType<typeof Anonymouse> {
    return new Anonymouse({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof Anonymouse> {
    return new Anonymouse(this.toJSON());
  }
}
export abstract class AnonymouseFactory {
  abstract create(data: unknown): Anonymouse;
}
/**
 * The base type definition for anonymouse
 **/
export type AnonymouseType = {
  /**
   * array field, non-nullable
   * @type {any[]}
   **/
  arrayField: any[];
  /**
   * arrayNullable field, non-nullable
   * @type {any}
   **/
  arrayNullableField: any;
  /**
   * collection field, non-nullable
   * @type {User[]}
   **/
  collectionField: User[];
  /**
   * collectionNullable field, non-nullable
   * @type {User[]}
   **/
  collectionNullableField?: User[];
  /**
   * one field, non-nullable
   * @type {User}
   **/
  oneField: User;
  /**
   * oneNullable field, non-nullable
   * @type {User}
   **/
  oneNullableField?: User;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AnonymouseType {
  /**
   * The base type definition for arrayFieldType
   **/
  export type ArrayFieldType = {};
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace ArrayFieldType {}
}
