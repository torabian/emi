import { Money } from "../complexes/Money";
import { type PartialDeep, type PlainOf } from "./sdk/common/fetchx";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for insertUsersVsqlFilters
 **/
export class InsertUsersVsqlFilters {
  /**
   * Primary key. Selected by default - stays a plain int64.
   * @type {number}
   **/
  #id: number = 0;
  /**
   * Primary key. Selected by default - stays a plain int64.
   * @returns {number}
   **/
  get id() {
    return this.#id;
  }
  /**
   * Primary key. Selected by default - stays a plain int64.
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
   * Off by default - becomes emigo.Nullable[int64] in the row DTO.
   * @type {number}
   **/
  #balanceCents?: number | null | undefined = undefined;
  /**
   * Off by default - becomes emigo.Nullable[int64] in the row DTO.
   * @returns {number}
   **/
  get balanceCents() {
    return this.#balanceCents;
  }
  /**
   * Off by default - becomes emigo.Nullable[int64] in the row DTO.
   * @type {number}
   **/
  set balanceCents(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#balanceCents = parsedValue;
    }
  }
  setBalanceCents(value: number | null | undefined) {
    this.balanceCents = value;
    return this;
  }
  /**
   * Selected by default - stays a plain string.
   * @type {string}
   **/
  #email: string = "";
  /**
   * Selected by default - stays a plain string.
   * @returns {string}
   **/
  get email() {
    return this.#email;
  }
  /**
   * Selected by default - stays a plain string.
   * @type {string}
   **/
  set email(value: string) {
    this.#email = String(value);
  }
  setEmail(value: string) {
    this.email = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #firstName?: string | null | undefined = undefined;
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
  set firstName(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#firstName = correctType ? value : String(value);
  }
  setFirstName(value: string | null | undefined) {
    this.firstName = value;
    return this;
  }
  /**
   * Selected by default - stays a plain bool.
   * @type {boolean}
   **/
  #isVerified!: boolean;
  /**
   * Selected by default - stays a plain bool.
   * @returns {boolean}
   **/
  get isVerified() {
    return this.#isVerified;
  }
  /**
   * Selected by default - stays a plain bool.
   * @type {boolean}
   **/
  set isVerified(value: boolean) {
    this.#isVerified = Boolean(value);
  }
  setIsVerified(value: boolean) {
    this.isVerified = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #isSuspended?: boolean | null | undefined = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get isSuspended() {
    return this.#isSuspended;
  }
  /**
   *
   * @type {boolean}
   **/
  set isSuspended(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#isSuspended = correctType ? value : Boolean(value);
  }
  setIsSuspended(value: boolean | null | undefined) {
    this.isSuspended = value;
    return this;
  }
  /**
   * Off by default - becomes emigo.Nullable[float64].
   * @type {number}
   **/
  #rating?: number | null | undefined = undefined;
  /**
   * Off by default - becomes emigo.Nullable[float64].
   * @returns {number}
   **/
  get rating() {
    return this.#rating;
  }
  /**
   * Off by default - becomes emigo.Nullable[float64].
   * @type {number}
   **/
  set rating(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#rating = parsedValue;
    }
  }
  setRating(value: number | null | undefined) {
    this.rating = value;
    return this;
  }
  /**
   * Selected by default - enums render as plain string. The real column is named account_status, not status - `column:` overrides the default snake_cased-Name projection for exactly this kind of mismatch.
   * @type {"Active" | "Suspended"}
   **/
  #status!: "Active" | "Suspended";
  /**
   * Selected by default - enums render as plain string. The real column is named account_status, not status - `column:` overrides the default snake_cased-Name projection for exactly this kind of mismatch.
   * @returns {"Active" | "Suspended"}
   **/
  get status() {
    return this.#status;
  }
  /**
   * Selected by default - enums render as plain string. The real column is named account_status, not status - `column:` overrides the default snake_cased-Name projection for exactly this kind of mismatch.
   * @type {"Active" | "Suspended"}
   **/
  set status(value: "Active" | "Suspended") {
    this.#status = value;
  }
  setStatus(value: "Active" | "Suspended") {
    this.status = value;
    return this;
  }
  /**
   * Off by default - becomes emigo.Nullable[string], same as any other nullable enum.
   * @type {any}
   **/
  #role?: any | null | undefined = undefined;
  /**
   * Off by default - becomes emigo.Nullable[string], same as any other nullable enum.
   * @returns {any}
   **/
  get role() {
    return this.#role;
  }
  /**
   * Off by default - becomes emigo.Nullable[string], same as any other nullable enum.
   * @type {any}
   **/
  set role(value: any | null | undefined) {
    this.#role = value;
  }
  setRole(value: any | null | undefined) {
    this.role = value;
    return this;
  }
  /**
   * Selected by default - stays the plain nested struct.
   * @type {InsertUsersVsqlFilters.Profile}
   **/
  #profile!: InstanceType<typeof InsertUsersVsqlFilters.Profile>;
  /**
   * Selected by default - stays the plain nested struct.
   * @returns {InsertUsersVsqlFilters.Profile}
   **/
  get profile() {
    return this.#profile;
  }
  /**
   * Selected by default - stays the plain nested struct.
   * @type {InsertUsersVsqlFilters.Profile}
   **/
  set profile(value: InstanceType<typeof InsertUsersVsqlFilters.Profile>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof InsertUsersVsqlFilters.Profile) {
      this.#profile = value;
    } else {
      this.#profile = new InsertUsersVsqlFilters.Profile(value);
    }
  }
  setProfile(value: InstanceType<typeof InsertUsersVsqlFilters.Profile>) {
    this.profile = value;
    return this;
  }
  /**
   * Off by default - becomes emigo.Nullable[InsertUsersVsqlRowPreferences], not just a plain nested struct with zero values, so "not fetched" is distinguishable from "fetched, all fields empty".
   * @type {InsertUsersVsqlFilters.Preferences}
   **/
  #preferences?:
    | InstanceType<typeof InsertUsersVsqlFilters.Preferences>
    | null
    | undefined = undefined;
  /**
   * Off by default - becomes emigo.Nullable[InsertUsersVsqlRowPreferences], not just a plain nested struct with zero values, so "not fetched" is distinguishable from "fetched, all fields empty".
   * @returns {InsertUsersVsqlFilters.Preferences}
   **/
  get preferences() {
    return this.#preferences;
  }
  /**
   * Off by default - becomes emigo.Nullable[InsertUsersVsqlRowPreferences], not just a plain nested struct with zero values, so "not fetched" is distinguishable from "fetched, all fields empty".
   * @type {InsertUsersVsqlFilters.Preferences}
   **/
  set preferences(
    value:
      | InstanceType<typeof InsertUsersVsqlFilters.Preferences>
      | null
      | undefined
      | null
      | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof InsertUsersVsqlFilters.Preferences) {
      this.#preferences = value;
    } else {
      this.#preferences = new InsertUsersVsqlFilters.Preferences(value);
    }
  }
  setPreferences(
    value:
      | InstanceType<typeof InsertUsersVsqlFilters.Preferences>
      | null
      | undefined
      | null
      | undefined,
  ) {
    this.preferences = value;
    return this;
  }
  /**
   * Demonstrates a complex column spanning two physical SQL columns under one projection/selection toggle - see sdkgen/money.go. `complex` has no wire-level nullability of its own (unlike a primitive or object), so being off by default only flips its type to `complex?`, a no-op for the Go backend - Money is responsible for its own optionality (see Money.SQLValue).
   * @type {Money}
   **/
  #money?: Money | null | undefined = undefined;
  /**
   * Demonstrates a complex column spanning two physical SQL columns under one projection/selection toggle - see sdkgen/money.go. `complex` has no wire-level nullability of its own (unlike a primitive or object), so being off by default only flips its type to `complex?`, a no-op for the Go backend - Money is responsible for its own optionality (see Money.SQLValue).
   * @returns {Money}
   **/
  get money() {
    return this.#money;
  }
  /**
   * Demonstrates a complex column spanning two physical SQL columns under one projection/selection toggle - see sdkgen/money.go. `complex` has no wire-level nullability of its own (unlike a primitive or object), so being off by default only flips its type to `complex?`, a no-op for the Go backend - Money is responsible for its own optionality (see Money.SQLValue).
   * @type {Money}
   **/
  set money(value: Money | null | undefined) {
    // For a nullable complex field, an explicit undefined/null is a
    // deliberate value and has to pass through untouched - same as every
    // other nullable field's setter (array?, one?, object?, ...) above.
    // Anything else always becomes a real instance, exactly like a
    // non-nullable "complex" field does.
    if (value === null || value === undefined) {
      this.#money = value === null ? null : undefined;
      return;
    }
    if (value instanceof Money) {
      this.#money = value;
    } else {
      this.#money = new Money(value);
    }
  }
  setMoney(value: Money | null | undefined) {
    this.money = value;
    return this;
  }
  /**
   * The base class definition for profile
   **/
  static Profile = class Profile {
    /**
     *
     * @type {string}
     **/
    #bio: string = "";
    /**
     *
     * @returns {string}
     **/
    get bio() {
      return this.#bio;
    }
    /**
     *
     * @type {string}
     **/
    set bio(value: string) {
      this.#bio = String(value);
    }
    setBio(value: string) {
      this.bio = value;
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
      const d = data as Partial<Profile>;
      if (d.bio !== undefined) {
        this.bio = d.bio;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly. Each
     *	field goes through #toPlainJSON rather than a bare this.#field: a
     *	nested dto instance, or an emigo Array/One/Collection wrapper, has its
     *	own toJSON that JSON.stringify would only reach on a *second* pass -
     *	calling it here means toJSON()'s own return value is already the same
     *	plain shape a consumer choosing the exported type instead of this
     *	class gets, rather than a shallow object still holding class instances.
     **/
    toJSON(): InsertUsersVsqlFiltersType.ProfileType {
      return {
        bio: this.#toPlainJSON(this.#bio),
      } as InsertUsersVsqlFiltersType.ProfileType;
    }
    /**
     * Resolves value into a plain, JSON-serializable value: recurses into
     * arrays, and - the case JSON.stringify's own recursion would only get to
     * after this method already returned - calls a nested value's own
     * toJSON() (a dto instance, or an emigo Array/One/Collection wrapper)
     * rather than leaving it as-is. A plain value (string, number, a map's
     * own plain object, ...) is returned unchanged.
     **/
    #toPlainJSON(value: unknown): unknown {
      if (value === null || value === undefined) {
        return value;
      }
      if (Array.isArray(value)) {
        return value.map((item) => this.#toPlainJSON(item));
      }
      const asAny = value as any;
      if (typeof asAny.toJSON === "function") {
        return this.#toPlainJSON(asAny.toJSON());
      }
      return value;
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        bio: "bio",
      };
    }
    /**
     * Creates an instance of InsertUsersVsqlFilters.Profile, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: InsertUsersVsqlFiltersType.ProfileType) {
      return new InsertUsersVsqlFilters.Profile(possibleDtoObject);
    }
    /**
     * Creates an instance of InsertUsersVsqlFilters.Profile, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<InsertUsersVsqlFiltersType.ProfileType>,
    ) {
      return new InsertUsersVsqlFilters.Profile(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<InsertUsersVsqlFiltersType.ProfileType>,
    ): InstanceType<typeof InsertUsersVsqlFilters.Profile> {
      return new InsertUsersVsqlFilters.Profile({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof InsertUsersVsqlFilters.Profile> {
      return new InsertUsersVsqlFilters.Profile(this.toJSON());
    }
  };
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
     *	Json stringify won't see them unless we mention it explicitly. Each
     *	field goes through #toPlainJSON rather than a bare this.#field: a
     *	nested dto instance, or an emigo Array/One/Collection wrapper, has its
     *	own toJSON that JSON.stringify would only reach on a *second* pass -
     *	calling it here means toJSON()'s own return value is already the same
     *	plain shape a consumer choosing the exported type instead of this
     *	class gets, rather than a shallow object still holding class instances.
     **/
    toJSON(): InsertUsersVsqlFiltersType.PreferencesType {
      return {
        theme: this.#toPlainJSON(this.#theme),
        locale: this.#toPlainJSON(this.#locale),
      } as InsertUsersVsqlFiltersType.PreferencesType;
    }
    /**
     * Resolves value into a plain, JSON-serializable value: recurses into
     * arrays, and - the case JSON.stringify's own recursion would only get to
     * after this method already returned - calls a nested value's own
     * toJSON() (a dto instance, or an emigo Array/One/Collection wrapper)
     * rather than leaving it as-is. A plain value (string, number, a map's
     * own plain object, ...) is returned unchanged.
     **/
    #toPlainJSON(value: unknown): unknown {
      if (value === null || value === undefined) {
        return value;
      }
      if (Array.isArray(value)) {
        return value.map((item) => this.#toPlainJSON(item));
      }
      const asAny = value as any;
      if (typeof asAny.toJSON === "function") {
        return this.#toPlainJSON(asAny.toJSON());
      }
      return value;
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
     * Creates an instance of InsertUsersVsqlFilters.Preferences, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: InsertUsersVsqlFiltersType.PreferencesType) {
      return new InsertUsersVsqlFilters.Preferences(possibleDtoObject);
    }
    /**
     * Creates an instance of InsertUsersVsqlFilters.Preferences, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<InsertUsersVsqlFiltersType.PreferencesType>,
    ) {
      return new InsertUsersVsqlFilters.Preferences(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<InsertUsersVsqlFiltersType.PreferencesType>,
    ): InstanceType<typeof InsertUsersVsqlFilters.Preferences> {
      return new InsertUsersVsqlFilters.Preferences({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof InsertUsersVsqlFilters.Preferences> {
      return new InsertUsersVsqlFilters.Preferences(this.toJSON());
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
    const d = data as Partial<InsertUsersVsqlFilters>;
    if (d.id !== undefined) {
      this.id = d.id;
    }
    if (d.balanceCents !== undefined) {
      this.balanceCents = d.balanceCents;
    }
    if (d.email !== undefined) {
      this.email = d.email;
    }
    if (d.firstName !== undefined) {
      this.firstName = d.firstName;
    }
    if (d.isVerified !== undefined) {
      this.isVerified = d.isVerified;
    }
    if (d.isSuspended !== undefined) {
      this.isSuspended = d.isSuspended;
    }
    if (d.rating !== undefined) {
      this.rating = d.rating;
    }
    if (d.status !== undefined) {
      this.status = d.status;
    }
    if (d.role !== undefined) {
      this.role = d.role;
    }
    if (d.profile !== undefined) {
      this.profile = d.profile;
    }
    if (d.preferences !== undefined) {
      this.preferences = d.preferences;
    }
    if (d.money !== undefined) {
      this.money = d.money;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<InsertUsersVsqlFilters>;
    if (!(d.profile instanceof InsertUsersVsqlFilters.Profile)) {
      this.profile = new InsertUsersVsqlFilters.Profile(d.profile || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly. Each
   *	field goes through #toPlainJSON rather than a bare this.#field: a
   *	nested dto instance, or an emigo Array/One/Collection wrapper, has its
   *	own toJSON that JSON.stringify would only reach on a *second* pass -
   *	calling it here means toJSON()'s own return value is already the same
   *	plain shape a consumer choosing the exported type instead of this
   *	class gets, rather than a shallow object still holding class instances.
   **/
  toJSON(): InsertUsersVsqlFiltersType {
    return {
      id: this.#toPlainJSON(this.#id),
      balanceCents: this.#toPlainJSON(this.#balanceCents),
      email: this.#toPlainJSON(this.#email),
      firstName: this.#toPlainJSON(this.#firstName),
      isVerified: this.#toPlainJSON(this.#isVerified),
      isSuspended: this.#toPlainJSON(this.#isSuspended),
      rating: this.#toPlainJSON(this.#rating),
      status: this.#toPlainJSON(this.#status),
      role: this.#toPlainJSON(this.#role),
      profile: this.#toPlainJSON(this.#profile),
      preferences: this.#toPlainJSON(this.#preferences),
      money: this.#toPlainJSON(this.#money),
    } as InsertUsersVsqlFiltersType;
  }
  /**
   * Resolves value into a plain, JSON-serializable value: recurses into
   * arrays, and - the case JSON.stringify's own recursion would only get to
   * after this method already returned - calls a nested value's own
   * toJSON() (a dto instance, or an emigo Array/One/Collection wrapper)
   * rather than leaving it as-is. A plain value (string, number, a map's
   * own plain object, ...) is returned unchanged.
   **/
  #toPlainJSON(value: unknown): unknown {
    if (value === null || value === undefined) {
      return value;
    }
    if (Array.isArray(value)) {
      return value.map((item) => this.#toPlainJSON(item));
    }
    const asAny = value as any;
    if (typeof asAny.toJSON === "function") {
      return this.#toPlainJSON(asAny.toJSON());
    }
    return value;
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      id: "id",
      balanceCents: "balanceCents",
      email: "email",
      firstName: "firstName",
      isVerified: "isVerified",
      isSuspended: "isSuspended",
      rating: "rating",
      status: "status",
      role: "role",
      profile$: "profile",
      get profile() {
        return withPrefix("profile", InsertUsersVsqlFilters.Profile.Fields);
      },
      preferences$: "preferences",
      get preferences() {
        return withPrefix(
          "preferences",
          InsertUsersVsqlFilters.Preferences.Fields,
        );
      },
      money: "money",
    };
  }
  /**
   * Creates an instance of InsertUsersVsqlFilters, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: InsertUsersVsqlFiltersType) {
    return new InsertUsersVsqlFilters(possibleDtoObject);
  }
  /**
   * Creates an instance of InsertUsersVsqlFilters, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<InsertUsersVsqlFiltersType>) {
    return new InsertUsersVsqlFilters(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<InsertUsersVsqlFiltersType>,
  ): InstanceType<typeof InsertUsersVsqlFilters> {
    return new InsertUsersVsqlFilters({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof InsertUsersVsqlFilters> {
    return new InsertUsersVsqlFilters(this.toJSON());
  }
}
export abstract class InsertUsersVsqlFiltersFactory {
  abstract create(data: unknown): InsertUsersVsqlFilters;
}
/**
 * The base type definition for insertUsersVsqlFilters
 **/
export type InsertUsersVsqlFiltersType = {
  /**
   * Primary key. Selected by default - stays a plain int64.
   * @type {number}
   **/
  id: number;
  /**
   * Off by default - becomes emigo.Nullable[int64] in the row DTO.
   * @type {number}
   **/
  balanceCents?: number;
  /**
   * Selected by default - stays a plain string.
   * @type {string}
   **/
  email: string;
  /**
   *
   * @type {string}
   **/
  firstName?: string;
  /**
   * Selected by default - stays a plain bool.
   * @type {boolean}
   **/
  isVerified: boolean;
  /**
   *
   * @type {boolean}
   **/
  isSuspended?: boolean;
  /**
   * Off by default - becomes emigo.Nullable[float64].
   * @type {number}
   **/
  rating?: number;
  /**
   * Selected by default - enums render as plain string. The real column is named account_status, not status - `column:` overrides the default snake_cased-Name projection for exactly this kind of mismatch.
   * @type {"Active" | "Suspended"}
   **/
  status: "Active" | "Suspended";
  /**
   * Off by default - becomes emigo.Nullable[string], same as any other nullable enum.
   * @type {any}
   **/
  role?: any;
  /**
   * Selected by default - stays the plain nested struct.
   * @type {InsertUsersVsqlFiltersType.ProfileType}
   **/
  profile: InsertUsersVsqlFiltersType.ProfileType;
  /**
   * Off by default - becomes emigo.Nullable[InsertUsersVsqlRowPreferences], not just a plain nested struct with zero values, so "not fetched" is distinguishable from "fetched, all fields empty".
   * @type {InsertUsersVsqlFiltersType.PreferencesType}
   **/
  preferences?: InsertUsersVsqlFiltersType.PreferencesType;
  /**
   * Demonstrates a complex column spanning two physical SQL columns under one projection/selection toggle - see sdkgen/money.go. `complex` has no wire-level nullability of its own (unlike a primitive or object), so being off by default only flips its type to `complex?`, a no-op for the Go backend - Money is responsible for its own optionality (see Money.SQLValue).
   * @type {PlainOf<Money>}
   **/
  money?: PlainOf<Money>;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace InsertUsersVsqlFiltersType {
  /**
   * The base type definition for profileType
   **/
  export type ProfileType = {
    /**
     *
     * @type {string}
     **/
    bio: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace ProfileType {}
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
}
/** Field names a JSON-Logic filter against the insertUsers vsql query may legally reference. */
export type InsertUsersVsqlFilterField =
  | "id"
  | "balanceCents"
  | "email"
  | "firstName"
  | "isVerified"
  | "isSuspended"
  | "rating"
  | "status"
  | "role"
  | "profile"
  | "preferences"
  | "money";
/**
 * Minimal typed JSON-Logic node (https://jsonlogic.com): only `{ var: ... }` is
 * constrained to a real field - every other operator stays loosely typed.
 */
export type InsertUsersVsqlFilterVarNode = { var: InsertUsersVsqlFilterField };
export type InsertUsersVsqlFilterOperatorNode = {
  [operator: string]: InsertUsersVsqlFilterValue;
} & { var?: never };
export type InsertUsersVsqlFilterValue =
  | InsertUsersVsqlFilter
  | string
  | number
  | boolean
  | null
  | InsertUsersVsqlFilterValue[];
export type InsertUsersVsqlFilter =
  | InsertUsersVsqlFilterVarNode
  | InsertUsersVsqlFilterOperatorNode;
