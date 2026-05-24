import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for userentityDto
 **/
export class UserentityDto {
  /**
   *
   * @type {string}
   **/
  #name = "";
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
  set name(value) {
    this.#name = String(value);
  }
  setName(value) {
    this.name = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #email = undefined;
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
  set email(value) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#email = correctType ? value : String(value);
  }
  setEmail(value) {
    this.email = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #age = 0;
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
  set age(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#age = parsedValue;
    }
  }
  setAge(value) {
    this.age = value;
    return this;
  }
  /**
   *
   * @type {UserentityDto.Preferences}
   **/
  #preferences;
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
  set preferences(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UserentityDto.Preferences) {
      this.#preferences = value;
    } else {
      this.#preferences = new UserentityDto.Preferences(value);
    }
  }
  setPreferences(value) {
    this.preferences = value;
    return this;
  }
  /**
   *
   * @type {UserentityDto.Tags}
   **/
  #tags = [];
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
  set tags(value) {
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
  setTags(value) {
    this.tags = value;
    return this;
  }
  /**
   * Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.
   * @type {UserentityDto.Address}
   **/
  #address;
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
  set address(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UserentityDto.Address) {
      this.#address = value;
    } else {
      this.#address = new UserentityDto.Address(value);
    }
  }
  setAddress(value) {
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
    #theme = "";
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
    set theme(value) {
      this.#theme = String(value);
    }
    setTheme(value) {
      this.theme = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #locale = "";
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
    set locale(value) {
      this.#locale = String(value);
    }
    setLocale(value) {
      this.locale = value;
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
    static from(possibleDtoObject) {
      return new UserentityDto.Preferences(possibleDtoObject);
    }
    /**
     * Creates an instance of UserentityDto.Preferences, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new UserentityDto.Preferences(partialDtoObject);
    }
    copyWith(partial) {
      return new UserentityDto.Preferences({ ...this.toJSON(), ...partial });
    }
    clone() {
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
     * @type {string}
     **/
    #value = "";
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
    set value(value) {
      this.#value = String(value);
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
     * Creates an instance of UserentityDto.Tags, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new UserentityDto.Tags(possibleDtoObject);
    }
    /**
     * Creates an instance of UserentityDto.Tags, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new UserentityDto.Tags(partialDtoObject);
    }
    copyWith(partial) {
      return new UserentityDto.Tags({ ...this.toJSON(), ...partial });
    }
    clone() {
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
    #id = 0;
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
    set id(value) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#id = parsedValue;
      }
    }
    setId(value) {
      this.id = value;
      return this;
    }
    /**
     *
     * @type {number}
     **/
    #userId = 0;
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
    set userId(value) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#userId = parsedValue;
      }
    }
    setUserId(value) {
      this.userId = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #street = "";
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
    set street(value) {
      this.#street = String(value);
    }
    setStreet(value) {
      this.street = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #city = "";
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
    set city(value) {
      this.#city = String(value);
    }
    setCity(value) {
      this.city = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #postcode = undefined;
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
    set postcode(value) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#postcode = correctType ? value : String(value);
    }
    setPostcode(value) {
      this.postcode = value;
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
    static from(possibleDtoObject) {
      return new UserentityDto.Address(possibleDtoObject);
    }
    /**
     * Creates an instance of UserentityDto.Address, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new UserentityDto.Address(partialDtoObject);
    }
    copyWith(partial) {
      return new UserentityDto.Address({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new UserentityDto.Address(this.toJSON());
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
    const d = data;
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
  static from(possibleDtoObject) {
    return new UserentityDto(possibleDtoObject);
  }
  /**
   * Creates an instance of UserentityDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new UserentityDto(partialDtoObject);
  }
  copyWith(partial) {
    return new UserentityDto({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new UserentityDto(this.toJSON());
  }
}
