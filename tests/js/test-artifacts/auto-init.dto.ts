import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for autoInitClassDto
 **/
export class AutoInitClassDto {
  /**
   * This field will be always initialised
   * @type {AutoInitClassDto.Object1}
   **/
  #object1!: InstanceType<typeof AutoInitClassDto.Object1>;
  /**
   * This field will be always initialised
   * @returns {AutoInitClassDto.Object1}
   **/
  get object1() {
    return this.#object1;
  }
  /**
   * This field will be always initialised
   * @type {AutoInitClassDto.Object1}
   **/
  set object1(value: InstanceType<typeof AutoInitClassDto.Object1>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof AutoInitClassDto.Object1) {
      this.#object1 = value;
    } else {
      this.#object1 = new AutoInitClassDto.Object1(value);
    }
  }
  setObject1(value: InstanceType<typeof AutoInitClassDto.Object1>) {
    this.object1 = value;
    return this;
  }
  /**
   * The base class definition for object1
   **/
  static Object1 = class Object1 {
    /**
     * This field also will be initialised, always.
     * @type {AutoInitClassDto.Object1.Object2}
     **/
    #object2!: InstanceType<typeof AutoInitClassDto.Object1.Object2>;
    /**
     * This field also will be initialised, always.
     * @returns {AutoInitClassDto.Object1.Object2}
     **/
    get object2() {
      return this.#object2;
    }
    /**
     * This field also will be initialised, always.
     * @type {AutoInitClassDto.Object1.Object2}
     **/
    set object2(value: InstanceType<typeof AutoInitClassDto.Object1.Object2>) {
      // For objects, the sub type needs to always be instance of the sub class.
      if (value instanceof AutoInitClassDto.Object1.Object2) {
        this.#object2 = value;
      } else {
        this.#object2 = new AutoInitClassDto.Object1.Object2(value);
      }
    }
    setObject2(value: InstanceType<typeof AutoInitClassDto.Object1.Object2>) {
      this.object2 = value;
      return this;
    }
    /**
     * The base class definition for object2
     **/
    static Object2 = class Object2 {
      /**
       * This field will be always an array
       * @type {AutoInitClassDto.Object1.Object2.Contacts}
       **/
      #contacts: InstanceType<
        typeof AutoInitClassDto.Object1.Object2.Contacts
      >[] = [];
      /**
       * This field will be always an array
       * @returns {AutoInitClassDto.Object1.Object2.Contacts}
       **/
      get contacts() {
        return this.#contacts;
      }
      /**
       * This field will be always an array
       * @type {AutoInitClassDto.Object1.Object2.Contacts}
       **/
      set contacts(
        value: InstanceType<typeof AutoInitClassDto.Object1.Object2.Contacts>[],
      ) {
        // For arrays, you only can pass arrays to the object
        if (!Array.isArray(value)) {
          return;
        }
        if (
          value.length > 0 &&
          value[0] instanceof AutoInitClassDto.Object1.Object2.Contacts
        ) {
          this.#contacts = value;
        } else {
          this.#contacts = value.map(
            (item) => new AutoInitClassDto.Object1.Object2.Contacts(item),
          );
        }
      }
      setContacts(
        value: InstanceType<typeof AutoInitClassDto.Object1.Object2.Contacts>[],
      ) {
        this.contacts = value;
        return this;
      }
      /**
       * The base class definition for contacts
       **/
      static Contacts = class Contacts {
        /**
         *
         * @type {string}
         **/
        #email: string = "emi-compiler@emi-compiler.com";
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
        set email(value: string) {
          const correctType = typeof value === "string";
          this.#email = correctType ? value : "" + value;
        }
        setEmail(value: string) {
          this.email = value;
          return this;
        }
        /**
         *
         * @type {string}
         **/
        #phone?: string | null = undefined;
        /**
         *
         * @returns {string}
         **/
        get phone() {
          return this.#phone;
        }
        /**
         *
         * @type {string}
         **/
        set phone(value: string | null | undefined) {
          const correctType =
            typeof value === "string" || value === undefined || value === null;
          this.#phone = correctType ? value : "" + value;
        }
        setPhone(value: string | null | undefined) {
          this.phone = value;
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
          const d = data as Partial<Contacts>;
          if (d.email !== undefined) {
            this.email = d.email;
          }
          if (d.phone !== undefined) {
            this.phone = d.phone;
          }
        }
        /**
         *	Special toJSON override, since the field are private,
         *	Json stringify won't see them unless we mention it explicitly.
         **/
        toJSON() {
          return {
            email: this.#email,
            phone: this.#phone,
          };
        }
        toString() {
          return JSON.stringify(this);
        }
        static get Fields() {
          return {
            email: "email",
            phone: "phone",
          };
        }
        /**
         * Creates an instance of AutoInitClassDto.Object1.Object2.Contacts, and possibleDtoObject
         * needs to satisfy the type requirement fully, otherwise typescript compile would
         * be complaining.
         **/
        static from(
          possibleDtoObject: AutoInitClassDtoType.Object1Type.Object2Type.ContactsType,
        ) {
          return new AutoInitClassDto.Object1.Object2.Contacts(
            possibleDtoObject,
          );
        }
        /**
         * Creates an instance of AutoInitClassDto.Object1.Object2.Contacts, and partialDtoObject
         * needs to satisfy the type, but partially, and rest of the content would
         * be constructed according to data types and nullability.
         **/
        static with(
          partialDtoObject: Partial<AutoInitClassDtoType.Object1Type.Object2Type.ContactsType>,
        ) {
          return new AutoInitClassDto.Object1.Object2.Contacts(
            partialDtoObject,
          );
        }
        copyWith(
          partial: Partial<AutoInitClassDtoType.Object1Type.Object2Type.ContactsType>,
        ): InstanceType<typeof AutoInitClassDto.Object1.Object2.Contacts> {
          return new AutoInitClassDto.Object1.Object2.Contacts({
            ...this.toJSON(),
            ...partial,
          });
        }
        clone(): InstanceType<
          typeof AutoInitClassDto.Object1.Object2.Contacts
        > {
          return new AutoInitClassDto.Object1.Object2.Contacts(this.toJSON());
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
        const d = data as Partial<Object2>;
        if (d.contacts !== undefined) {
          this.contacts = d.contacts;
        }
      }
      /**
       *	Special toJSON override, since the field are private,
       *	Json stringify won't see them unless we mention it explicitly.
       **/
      toJSON() {
        return {
          contacts: this.#contacts,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          contacts$: "contacts",
          get contacts() {
            return withPrefix(
              "object1.object2.contacts[:i]",
              AutoInitClassDto.Object1.Object2.Contacts.Fields,
            );
          },
        };
      }
      /**
       * Creates an instance of AutoInitClassDto.Object1.Object2, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(
        possibleDtoObject: AutoInitClassDtoType.Object1Type.Object2Type,
      ) {
        return new AutoInitClassDto.Object1.Object2(possibleDtoObject);
      }
      /**
       * Creates an instance of AutoInitClassDto.Object1.Object2, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(
        partialDtoObject: Partial<AutoInitClassDtoType.Object1Type.Object2Type>,
      ) {
        return new AutoInitClassDto.Object1.Object2(partialDtoObject);
      }
      copyWith(
        partial: Partial<AutoInitClassDtoType.Object1Type.Object2Type>,
      ): InstanceType<typeof AutoInitClassDto.Object1.Object2> {
        return new AutoInitClassDto.Object1.Object2({
          ...this.toJSON(),
          ...partial,
        });
      }
      clone(): InstanceType<typeof AutoInitClassDto.Object1.Object2> {
        return new AutoInitClassDto.Object1.Object2(this.toJSON());
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
      const d = data as Partial<Object1>;
      if (d.object2 !== undefined) {
        this.object2 = d.object2;
      }
      this.#lateInitFields(data);
    }
    /**
     * These are the class instances, which need to be initialised, regardless of the constructor incoming data
     **/
    #lateInitFields(data = {}) {
      const d = data as Partial<Object1>;
      if (!(d.object2 instanceof AutoInitClassDto.Object1.Object2)) {
        this.object2 = new AutoInitClassDto.Object1.Object2(d.object2 || {});
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        object2: this.#object2,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        object2$: "object2",
        get object2() {
          return withPrefix(
            "object1.object2",
            AutoInitClassDto.Object1.Object2.Fields,
          );
        },
      };
    }
    /**
     * Creates an instance of AutoInitClassDto.Object1, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: AutoInitClassDtoType.Object1Type) {
      return new AutoInitClassDto.Object1(possibleDtoObject);
    }
    /**
     * Creates an instance of AutoInitClassDto.Object1, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: Partial<AutoInitClassDtoType.Object1Type>) {
      return new AutoInitClassDto.Object1(partialDtoObject);
    }
    copyWith(
      partial: Partial<AutoInitClassDtoType.Object1Type>,
    ): InstanceType<typeof AutoInitClassDto.Object1> {
      return new AutoInitClassDto.Object1({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof AutoInitClassDto.Object1> {
      return new AutoInitClassDto.Object1(this.toJSON());
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
    const d = data as Partial<AutoInitClassDto>;
    if (d.object1 !== undefined) {
      this.object1 = d.object1;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<AutoInitClassDto>;
    if (!(d.object1 instanceof AutoInitClassDto.Object1)) {
      this.object1 = new AutoInitClassDto.Object1(d.object1 || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      object1: this.#object1,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      object1$: "object1",
      get object1() {
        return withPrefix("object1", AutoInitClassDto.Object1.Fields);
      },
    };
  }
  /**
   * Creates an instance of AutoInitClassDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AutoInitClassDtoType) {
    return new AutoInitClassDto(possibleDtoObject);
  }
  /**
   * Creates an instance of AutoInitClassDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: Partial<AutoInitClassDtoType>) {
    return new AutoInitClassDto(partialDtoObject);
  }
  copyWith(
    partial: Partial<AutoInitClassDtoType>,
  ): InstanceType<typeof AutoInitClassDto> {
    return new AutoInitClassDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AutoInitClassDto> {
    return new AutoInitClassDto(this.toJSON());
  }
}
export abstract class AutoInitClassDtoFactory {
  abstract create(data: unknown): AutoInitClassDto;
}
/**
 * The base type definition for autoInitClassDto
 **/
export type AutoInitClassDtoType = {
  /**
   * This field will be always initialised
   * @type {AutoInitClassDtoType.Object1Type}
   **/
  object1: AutoInitClassDtoType.Object1Type;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AutoInitClassDtoType {
  /**
   * The base type definition for object1Type
   **/
  export type Object1Type = {
    /**
     * This field also will be initialised, always.
     * @type {AutoInitClassDtoType.Object1Type.Object2Type}
     **/
    object2: AutoInitClassDtoType.Object1Type.Object2Type;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace Object1Type {
    /**
     * The base type definition for object2Type
     **/
    export type Object2Type = {
      /**
       * This field will be always an array
       * @type {AutoInitClassDtoType.Object1Type.Object2Type.ContactsType[]}
       **/
      contacts: AutoInitClassDtoType.Object1Type.Object2Type.ContactsType[];
    };
    // eslint-disable-next-line @typescript-eslint/no-namespace
    export namespace Object2Type {
      /**
       * The base type definition for contactsType
       **/
      export type ContactsType = {
        /**
         *
         * @type {string}
         **/
        email: string;
        /**
         *
         * @type {string}
         **/
        phone?: string;
      };
      // eslint-disable-next-line @typescript-eslint/no-namespace
      export namespace ContactsType {}
    }
  }
}
