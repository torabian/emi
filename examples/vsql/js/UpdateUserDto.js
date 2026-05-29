import { operators } from "./sdk/common/operators";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for updateUserDto
 **/
export class UpdateUserDto {
  /**
   *
   * @type {number}
   **/
  #id = undefined;
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
    const correctType =
      typeof value === "number" || value === undefined || value === null;
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
   * @type {string}
   **/
  #name = undefined;
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
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#name = correctType ? value : String(value);
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
  #age = undefined;
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
    const correctType =
      typeof value === "number" || value === undefined || value === null;
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
   * @type {UpdateUserDto.Preferences}
   **/
  #preferences = undefined;
  /**
   *
   * @returns {UpdateUserDto.Preferences}
   **/
  get preferences() {
    return this.#preferences;
  }
  /**
   *
   * @type {UpdateUserDto.Preferences}
   **/
  set preferences(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UpdateUserDto.Preferences) {
      this.#preferences = value;
    } else {
      this.#preferences = new UpdateUserDto.Preferences(value);
    }
  }
  setPreferences(value) {
    this.preferences = value;
    return this;
  }
  /**
   *
   * @type {UpdateUserDto.Tags}
   **/
  #tags = undefined;
  /**
   *
   * @returns {UpdateUserDto.Tags}
   **/
  get tags() {
    return this.#tags;
  }
  /**
   *
   * @type {UpdateUserDto.Tags}
   **/
  set tags(value) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof UpdateUserDto.Tags) {
      this.#tags = value;
    } else {
      this.#tags = value.map((item) => new UpdateUserDto.Tags(item));
    }
  }
  setTags(value) {
    this.tags = value;
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
     * Creates an instance of UpdateUserDto.Preferences, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new UpdateUserDto.Preferences(possibleDtoObject);
    }
    /**
     * Creates an instance of UpdateUserDto.Preferences, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new UpdateUserDto.Preferences(partialDtoObject);
    }
    copyWith(partial) {
      return new UpdateUserDto.Preferences({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new UpdateUserDto.Preferences(this.toJSON());
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
     * Creates an instance of UpdateUserDto.Tags, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new UpdateUserDto.Tags(possibleDtoObject);
    }
    /**
     * Creates an instance of UpdateUserDto.Tags, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new UpdateUserDto.Tags(partialDtoObject);
    }
    copyWith(partial) {
      return new UpdateUserDto.Tags({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new UpdateUserDto.Tags(this.toJSON());
    }
  };
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
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      id: this.#id,
      name: this.#name,
      email: this.#email,
      age: this.#age,
      preferences: this.#preferences,
      tags: this.#tags,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      id: "id",
      name: "name",
      email: "email",
      age: "age",
      preferences$: "preferences",
      get preferences() {
        return withPrefix("preferences", UpdateUserDto.Preferences.Fields);
      },
      tags$: "tags",
      get tags() {
        return withPrefix("tags[:i]", UpdateUserDto.Tags.Fields);
      },
    };
  }
  /**
   * Creates an instance of UpdateUserDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new UpdateUserDto(possibleDtoObject);
  }
  /**
   * Creates an instance of UpdateUserDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new UpdateUserDto(partialDtoObject);
  }
  copyWith(partial) {
    return new UpdateUserDto({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new UpdateUserDto(this.toJSON());
  }
}
