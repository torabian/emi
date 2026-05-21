import { type PartialDeep } from "./sdk/common/fetchx";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for userentityDto
 **/
export class UserentityDto {
  /**
   *
   * @type {string}
   **/
  #name: string = "";
  /**
   *
   * @returns {string}
   **/
  get name() {
    return this.#name;
  }
  /**
   *
   * @type {string}
   **/
  set name(value: string) {
    this.#name = String(value);
  }
  setName(value: string) {
    this.name = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #email?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get email() {
    return this.#email;
  }
  /**
   *
   * @type {string}
   **/
  set email(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#email = correctType ? value : String(value);
  }
  setEmail(value: string | null | undefined) {
    this.email = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #age: number = 0;
  /**
   *
   * @returns {number}
   **/
  get age() {
    return this.#age;
  }
  /**
   *
   * @type {number}
   **/
  set age(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#age = parsedValue;
    }
  }
  setAge(value: number) {
    this.age = value;
    return this;
  }
  /**
   *
   * @type {UserentityDto.Preferences}
   **/
  #preferences!: InstanceType<typeof UserentityDto.Preferences>;
  /**
   *
   * @returns {UserentityDto.Preferences}
   **/
  get preferences() {
    return this.#preferences;
  }
  /**
   *
   * @type {UserentityDto.Preferences}
   **/
  set preferences(value: InstanceType<typeof UserentityDto.Preferences>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UserentityDto.Preferences) {
      this.#preferences = value;
    } else {
      this.#preferences = new UserentityDto.Preferences(value);
    }
  }
  setPreferences(value: InstanceType<typeof UserentityDto.Preferences>) {
    this.preferences = value;
    return this;
  }
  /**
   *
   * @type {UserentityDto.Tags}
   **/
  #tags: InstanceType<typeof UserentityDto.Tags>[] = [];
  /**
   *
   * @returns {UserentityDto.Tags}
   **/
  get tags() {
    return this.#tags;
  }
  /**
   *
   * @type {UserentityDto.Tags}
   **/
  set tags(value: InstanceType<typeof UserentityDto.Tags>[]) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof UserentityDto.Tags) {
      this.#tags = value;
    } else {
      this.#tags = value.map((item) => new UserentityDto.Tags(item));
    }
  }
  setTags(value: InstanceType<typeof UserentityDto.Tags>[]) {
    this.tags = value;
    return this;
  }
  /**
   * Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.
   * @type {UserentityDto.Address}
   **/
  #address!: InstanceType<typeof UserentityDto.Address>;
  /**
   * Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.
   * @returns {UserentityDto.Address}
   **/
  get address() {
    return this.#address;
  }
  /**
   * Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.
   * @type {UserentityDto.Address}
   **/
  set address(value: InstanceType<typeof UserentityDto.Address>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UserentityDto.Address) {
      this.#address = value;
    } else {
      this.#address = new UserentityDto.Address(value);
    }
  }
  setAddress(value: InstanceType<typeof UserentityDto.Address>) {
    this.address = value;
    return this;
  }
  /**
   * The base class definition for preferences
   **/
  static Preferences = class Preferences {
    /**
     *
     * @type {string}
     **/
    #theme: string = "";
    /**
     *
     * @returns {string}
     **/
    get theme() {
      return this.#theme;
    }
    /**
     *
     * @type {string}
     **/
    set theme(value: string) {
      this.#theme = String(value);
    }
    setTheme(value: string) {
      this.theme = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #locale: string = "";
    /**
     *
     * @returns {string}
     **/
    get locale() {
      return this.#locale;
    }
    /**
     *
     * @type {string}
     **/
    set locale(value: string) {
      this.#locale = String(value);
    }
    setLocale(value: string) {
      this.locale = value;
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
      const d = data as Partial<Preferences>;
      if (d.theme !== undefined) {
        this.theme = d.theme;
      }
      if (d.locale !== undefined) {
        this.locale = d.locale;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        theme: this.#theme,
        locale: this.#locale,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        theme: "theme",
        locale: "locale",
      };
    }
    /**
     * Creates an instance of UserentityDto.Preferences, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: UserentityDtoType.PreferencesType) {
      return new UserentityDto.Preferences(possibleDtoObject);
    }
    /**
     * Creates an instance of UserentityDto.Preferences, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<UserentityDtoType.PreferencesType>,
    ) {
      return new UserentityDto.Preferences(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<UserentityDtoType.PreferencesType>,
    ): InstanceType<typeof UserentityDto.Preferences> {
      return new UserentityDto.Preferences({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof UserentityDto.Preferences> {
      return new UserentityDto.Preferences(this.toJSON());
    }
  };
  /**
   * The base class definition for tags
   **/
  static Tags = class Tags {
    /**
     *
     * @type {string}
     **/
    #key: string = "";
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
    set key(value: string) {
      this.#key = String(value);
    }
    setKey(value: string) {
      this.key = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #value: string = "";
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
    set value(value: string) {
      this.#value = String(value);
    }
    setValue(value: string) {
      this.value = value;
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
      const d = data as Partial<Tags>;
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
     * Creates an instance of UserentityDto.Tags, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: UserentityDtoType.TagsType) {
      return new UserentityDto.Tags(possibleDtoObject);
    }
    /**
     * Creates an instance of UserentityDto.Tags, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<UserentityDtoType.TagsType>) {
      return new UserentityDto.Tags(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<UserentityDtoType.TagsType>,
    ): InstanceType<typeof UserentityDto.Tags> {
      return new UserentityDto.Tags({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof UserentityDto.Tags> {
      return new UserentityDto.Tags(this.toJSON());
    }
  };
  /**
   * The base class definition for address
   **/
  static Address = class Address {
    /**
     *
     * @type {number}
     **/
    #id: number = 0;
    /**
     *
     * @returns {number}
     **/
    get id() {
      return this.#id;
    }
    /**
     *
     * @type {number}
     **/
    set id(value: number) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#id = parsedValue;
      }
    }
    setId(value: number) {
      this.id = value;
      return this;
    }
    /**
     *
     * @type {number}
     **/
    #userId: number = 0;
    /**
     *
     * @returns {number}
     **/
    get userId() {
      return this.#userId;
    }
    /**
     *
     * @type {number}
     **/
    set userId(value: number) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#userId = parsedValue;
      }
    }
    setUserId(value: number) {
      this.userId = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #street: string = "";
    /**
     *
     * @returns {string}
     **/
    get street() {
      return this.#street;
    }
    /**
     *
     * @type {string}
     **/
    set street(value: string) {
      this.#street = String(value);
    }
    setStreet(value: string) {
      this.street = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #city: string = "";
    /**
     *
     * @returns {string}
     **/
    get city() {
      return this.#city;
    }
    /**
     *
     * @type {string}
     **/
    set city(value: string) {
      this.#city = String(value);
    }
    setCity(value: string) {
      this.city = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #postcode?: string | null = undefined;
    /**
     *
     * @returns {string}
     **/
    get postcode() {
      return this.#postcode;
    }
    /**
     *
     * @type {string}
     **/
    set postcode(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#postcode = correctType ? value : String(value);
    }
    setPostcode(value: string | null | undefined) {
      this.postcode = value;
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
      const d = data as Partial<Address>;
      if (d.id !== undefined) {
        this.id = d.id;
      }
      if (d.userId !== undefined) {
        this.userId = d.userId;
      }
      if (d.street !== undefined) {
        this.street = d.street;
      }
      if (d.city !== undefined) {
        this.city = d.city;
      }
      if (d.postcode !== undefined) {
        this.postcode = d.postcode;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        id: this.#id,
        userId: this.#userId,
        street: this.#street,
        city: this.#city,
        postcode: this.#postcode,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        id: "id",
        userId: "userId",
        street: "street",
        city: "city",
        postcode: "postcode",
      };
    }
    /**
     * Creates an instance of UserentityDto.Address, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: UserentityDtoType.AddressType) {
      return new UserentityDto.Address(possibleDtoObject);
    }
    /**
     * Creates an instance of UserentityDto.Address, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<UserentityDtoType.AddressType>) {
      return new UserentityDto.Address(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<UserentityDtoType.AddressType>,
    ): InstanceType<typeof UserentityDto.Address> {
      return new UserentityDto.Address({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof UserentityDto.Address> {
      return new UserentityDto.Address(this.toJSON());
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
    const d = data as Partial<UserentityDto>;
    if (d.name !== undefined) {
      this.name = d.name;
    }
    if (d.email !== undefined) {
      this.email = d.email;
    }
    if (d.age !== undefined) {
      this.age = d.age;
    }
    if (d.preferences !== undefined) {
      this.preferences = d.preferences;
    }
    if (d.tags !== undefined) {
      this.tags = d.tags;
    }
    if (d.address !== undefined) {
      this.address = d.address;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<UserentityDto>;
    if (!(d.preferences instanceof UserentityDto.Preferences)) {
      this.preferences = new UserentityDto.Preferences(d.preferences || {});
    }
    if (!(d.address instanceof UserentityDto.Address)) {
      this.address = new UserentityDto.Address(d.address || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      name: this.#name,
      email: this.#email,
      age: this.#age,
      preferences: this.#preferences,
      tags: this.#tags,
      address: this.#address,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      name: "name",
      email: "email",
      age: "age",
      preferences$: "preferences",
      get preferences() {
        return withPrefix("preferences", UserentityDto.Preferences.Fields);
      },
      tags$: "tags",
      get tags() {
        return withPrefix("tags[:i]", UserentityDto.Tags.Fields);
      },
      address$: "address",
      get address() {
        return withPrefix("address", UserentityDto.Address.Fields);
      },
    };
  }
  /**
   * Creates an instance of UserentityDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UserentityDtoType) {
    return new UserentityDto(possibleDtoObject);
  }
  /**
   * Creates an instance of UserentityDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UserentityDtoType>) {
    return new UserentityDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UserentityDtoType>,
  ): InstanceType<typeof UserentityDto> {
    return new UserentityDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UserentityDto> {
    return new UserentityDto(this.toJSON());
  }
}
export abstract class UserentityDtoFactory {
  abstract create(data: unknown): UserentityDto;
}
/**
 * The base type definition for userentityDto
 **/
export type UserentityDtoType = {
  /**
   *
   * @type {string}
   **/
  name: string;
  /**
   *
   * @type {string}
   **/
  email?: string;
  /**
   *
   * @type {number}
   **/
  age: number;
  /**
   *
   * @type {UserentityDtoType.PreferencesType}
   **/
  preferences: UserentityDtoType.PreferencesType;
  /**
   *
   * @type {UserentityDtoType.TagsType[]}
   **/
  tags: UserentityDtoType.TagsType[];
  /**
   * Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.
   * @type {UserentityDtoType.AddressType}
   **/
  address: UserentityDtoType.AddressType;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UserentityDtoType {
  /**
   * The base type definition for preferencesType
   **/
  export type PreferencesType = {
    /**
     *
     * @type {string}
     **/
    theme: string;
    /**
     *
     * @type {string}
     **/
    locale: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace PreferencesType {}
  /**
   * The base type definition for tagsType
   **/
  export type TagsType = {
    /**
     *
     * @type {string}
     **/
    key: string;
    /**
     *
     * @type {string}
     **/
    value: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace TagsType {}
  /**
   * The base type definition for addressType
   **/
  export type AddressType = {
    /**
     *
     * @type {number}
     **/
    id: number;
    /**
     *
     * @type {number}
     **/
    userId: number;
    /**
     *
     * @type {string}
     **/
    street: string;
    /**
     *
     * @type {string}
     **/
    city: string;
    /**
     *
     * @type {string}
     **/
    postcode?: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace AddressType {}
}
