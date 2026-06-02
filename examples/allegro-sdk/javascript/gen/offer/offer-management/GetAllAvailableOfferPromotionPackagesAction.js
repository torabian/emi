import { MArray } from "./sdk/common/operators";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * Action to communicate with the action Get all available offer promotion packages
 */
/**
 * GetAllAvailableOfferPromotionPackagesAction
 */
export class GetAllAvailableOfferPromotionPackagesAction {
  //
  static URL = "https://api.{environment}/sale/offer-promotion-packages";
  static NewUrl = (qs) =>
    buildUrl(GetAllAvailableOfferPromotionPackagesAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? GetAllAvailableOfferPromotionPackagesAction.NewUrl(qs),
      {
        method: GetAllAvailableOfferPromotionPackagesAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) =>
        new GetAllAvailableOfferPromotionPackagesActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn ||
      ((item) => new GetAllAvailableOfferPromotionPackagesActionRes(item));
    const res = await GetAllAvailableOfferPromotionPackagesAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "Get all available offer promotion packages",
    url: "https://api.{environment}/sale/offer-promotion-packages",
    method: "get",
    description:
      "Use this resource to retrieve all available offer promotion packages. Read more: PL / EN.",
    out: {
      fields: [
        {
          name: "marketplaceId",
          type: "string",
        },
        {
          name: "basePackages",
          type: "array",
          fields: [
            {
              name: "id",
              type: "string",
            },
            {
              name: "name",
              type: "string",
            },
            {
              name: "cycleDuration",
              type: "string",
            },
          ],
        },
        {
          name: "extraPackages",
          type: "array",
          fields: [
            {
              name: "id",
              type: "string",
            },
            {
              name: "name",
              type: "string",
            },
            {
              name: "cycleDuration",
              type: "string",
            },
          ],
        },
        {
          name: "additionalMarketplaces",
          type: "array",
          fields: [
            {
              name: "marketplaceId",
              type: "string",
            },
            {
              name: "basePackages",
              type: "array",
              fields: [
                {
                  name: "id",
                  type: "string",
                },
                {
                  name: "name",
                  type: "string",
                },
                {
                  name: "cycleDuration",
                  type: "string",
                },
              ],
            },
            {
              name: "extraPackages",
              type: "array",
              fields: [
                {
                  name: "id",
                  type: "string",
                },
                {
                  name: "name",
                  type: "string",
                },
                {
                  name: "cycleDuration",
                  type: "string",
                },
              ],
            },
          ],
        },
      ],
    },
  };
}
/**
 * The base class definition for getAllAvailableOfferPromotionPackagesActionRes
 **/
export class GetAllAvailableOfferPromotionPackagesActionRes {
  /**
   *
   * @type {string}
   **/
  #marketplaceId = "";
  /**
   *
   * @returns {string}
   **/
  get marketplaceId() {
    return this.#marketplaceId;
  }
  /**
   *
   * @type {string}
   **/
  set marketplaceId(value) {
    this.#marketplaceId = String(value);
  }
  setMarketplaceId(value) {
    this.marketplaceId = value;
    return this;
  }
  /**
   *
   * @type {GetAllAvailableOfferPromotionPackagesActionRes.BasePackages}
   **/
  #basePackages = MArray.of([]);
  /**
   *
   * @returns {GetAllAvailableOfferPromotionPackagesActionRes.BasePackages}
   **/
  get basePackages() {
    return this.#basePackages;
  }
  /**
   *
   * @type {GetAllAvailableOfferPromotionPackagesActionRes.BasePackages}
   **/
  set basePackages(value) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (
        value.length > 0 &&
        value[0] instanceof
          GetAllAvailableOfferPromotionPackagesActionRes.BasePackages
      ) {
        this.#basePackages = MArray.of(value);
      } else {
        this.#basePackages = MArray.of(
          value.map(
            (item) =>
              new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages(
                item,
              ),
          ),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#basePackages = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast(value);
    if (ok) {
      this.#basePackages = mcastValue;
      return;
    }
    console.warn(
      "Cannot assing value to basePackages, because it needs MArray instance or an Array.",
    );
  }
  setBasePackages(value) {
    this.basePackages = value;
    return this;
  }
  /**
   *
   * @type {GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages}
   **/
  #extraPackages = MArray.of([]);
  /**
   *
   * @returns {GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages}
   **/
  get extraPackages() {
    return this.#extraPackages;
  }
  /**
   *
   * @type {GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages}
   **/
  set extraPackages(value) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (
        value.length > 0 &&
        value[0] instanceof
          GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages
      ) {
        this.#extraPackages = MArray.of(value);
      } else {
        this.#extraPackages = MArray.of(
          value.map(
            (item) =>
              new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages(
                item,
              ),
          ),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#extraPackages = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast(value);
    if (ok) {
      this.#extraPackages = mcastValue;
      return;
    }
    console.warn(
      "Cannot assing value to extraPackages, because it needs MArray instance or an Array.",
    );
  }
  setExtraPackages(value) {
    this.extraPackages = value;
    return this;
  }
  /**
   *
   * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces}
   **/
  #additionalMarketplaces = MArray.of([]);
  /**
   *
   * @returns {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces}
   **/
  get additionalMarketplaces() {
    return this.#additionalMarketplaces;
  }
  /**
   *
   * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces}
   **/
  set additionalMarketplaces(value) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (
        value.length > 0 &&
        value[0] instanceof
          GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces
      ) {
        this.#additionalMarketplaces = MArray.of(value);
      } else {
        this.#additionalMarketplaces = MArray.of(
          value.map(
            (item) =>
              new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces(
                item,
              ),
          ),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#additionalMarketplaces = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast(value);
    if (ok) {
      this.#additionalMarketplaces = mcastValue;
      return;
    }
    console.warn(
      "Cannot assing value to additionalMarketplaces, because it needs MArray instance or an Array.",
    );
  }
  setAdditionalMarketplaces(value) {
    this.additionalMarketplaces = value;
    return this;
  }
  /**
   * The base class definition for basePackages
   **/
  static BasePackages = class BasePackages {
    /**
     *
     * @type {string}
     **/
    #id = "";
    /**
     *
     * @returns {string}
     **/
    get id() {
      return this.#id;
    }
    /**
     *
     * @type {string}
     **/
    set id(value) {
      this.#id = String(value);
    }
    setId(value) {
      this.id = value;
      return this;
    }
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
    #cycleDuration = "";
    /**
     *
     * @returns {string}
     **/
    get cycleDuration() {
      return this.#cycleDuration;
    }
    /**
     *
     * @type {string}
     **/
    set cycleDuration(value) {
      this.#cycleDuration = String(value);
    }
    setCycleDuration(value) {
      this.cycleDuration = value;
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
      if (d.name !== undefined) {
        this.name = d.name;
      }
      if (d.cycleDuration !== undefined) {
        this.cycleDuration = d.cycleDuration;
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
        cycleDuration: this.#cycleDuration,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        id: "id",
        name: "name",
        cycleDuration: "cycleDuration",
      };
    }
    /**
     * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.BasePackages, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages(
        possibleDtoObject,
      );
    }
    /**
     * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.BasePackages, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages(
        partialDtoObject,
      );
    }
    copyWith(partial) {
      return new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone() {
      return new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages(
        this.toJSON(),
      );
    }
  };
  /**
   * The base class definition for extraPackages
   **/
  static ExtraPackages = class ExtraPackages {
    /**
     *
     * @type {string}
     **/
    #id = "";
    /**
     *
     * @returns {string}
     **/
    get id() {
      return this.#id;
    }
    /**
     *
     * @type {string}
     **/
    set id(value) {
      this.#id = String(value);
    }
    setId(value) {
      this.id = value;
      return this;
    }
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
    #cycleDuration = "";
    /**
     *
     * @returns {string}
     **/
    get cycleDuration() {
      return this.#cycleDuration;
    }
    /**
     *
     * @type {string}
     **/
    set cycleDuration(value) {
      this.#cycleDuration = String(value);
    }
    setCycleDuration(value) {
      this.cycleDuration = value;
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
      if (d.name !== undefined) {
        this.name = d.name;
      }
      if (d.cycleDuration !== undefined) {
        this.cycleDuration = d.cycleDuration;
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
        cycleDuration: this.#cycleDuration,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        id: "id",
        name: "name",
        cycleDuration: "cycleDuration",
      };
    }
    /**
     * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages(
        possibleDtoObject,
      );
    }
    /**
     * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages(
        partialDtoObject,
      );
    }
    copyWith(partial) {
      return new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone() {
      return new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages(
        this.toJSON(),
      );
    }
  };
  /**
   * The base class definition for additionalMarketplaces
   **/
  static AdditionalMarketplaces = class AdditionalMarketplaces {
    /**
     *
     * @type {string}
     **/
    #marketplaceId = "";
    /**
     *
     * @returns {string}
     **/
    get marketplaceId() {
      return this.#marketplaceId;
    }
    /**
     *
     * @type {string}
     **/
    set marketplaceId(value) {
      this.#marketplaceId = String(value);
    }
    setMarketplaceId(value) {
      this.marketplaceId = value;
      return this;
    }
    /**
     *
     * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages}
     **/
    #basePackages = MArray.of([]);
    /**
     *
     * @returns {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages}
     **/
    get basePackages() {
      return this.#basePackages;
    }
    /**
     *
     * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages}
     **/
    set basePackages(value) {
      // When the passed value is already an array, we check if we need to
      // cast the inner items into class instance.
      if (Array.isArray(value)) {
        if (
          value.length > 0 &&
          value[0] instanceof
            GetAllAvailableOfferPromotionPackagesActionRes
              .AdditionalMarketplaces.BasePackages
        ) {
          this.#basePackages = MArray.of(value);
        } else {
          this.#basePackages = MArray.of(
            value.map(
              (item) =>
                new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages(
                  item,
                ),
            ),
          );
        }
        return;
      }
      // If the instance is already an MArray, we assume it's all good.
      if (value instanceof MArray) {
        this.#basePackages = value;
        return;
      }
      // If the value is not array, and is not a MArray, we need to be consider,
      // it might be eligible to be casted into MArray.
      const { ok, value: mcastValue } = MArray.cast(value);
      if (ok) {
        this.#basePackages = mcastValue;
        return;
      }
      console.warn(
        "Cannot assing value to basePackages, because it needs MArray instance or an Array.",
      );
    }
    setBasePackages(value) {
      this.basePackages = value;
      return this;
    }
    /**
     *
     * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
     **/
    #extraPackages = MArray.of([]);
    /**
     *
     * @returns {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
     **/
    get extraPackages() {
      return this.#extraPackages;
    }
    /**
     *
     * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
     **/
    set extraPackages(value) {
      // When the passed value is already an array, we check if we need to
      // cast the inner items into class instance.
      if (Array.isArray(value)) {
        if (
          value.length > 0 &&
          value[0] instanceof
            GetAllAvailableOfferPromotionPackagesActionRes
              .AdditionalMarketplaces.ExtraPackages
        ) {
          this.#extraPackages = MArray.of(value);
        } else {
          this.#extraPackages = MArray.of(
            value.map(
              (item) =>
                new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(
                  item,
                ),
            ),
          );
        }
        return;
      }
      // If the instance is already an MArray, we assume it's all good.
      if (value instanceof MArray) {
        this.#extraPackages = value;
        return;
      }
      // If the value is not array, and is not a MArray, we need to be consider,
      // it might be eligible to be casted into MArray.
      const { ok, value: mcastValue } = MArray.cast(value);
      if (ok) {
        this.#extraPackages = mcastValue;
        return;
      }
      console.warn(
        "Cannot assing value to extraPackages, because it needs MArray instance or an Array.",
      );
    }
    setExtraPackages(value) {
      this.extraPackages = value;
      return this;
    }
    /**
     * The base class definition for basePackages
     **/
    static BasePackages = class BasePackages {
      /**
       *
       * @type {string}
       **/
      #id = "";
      /**
       *
       * @returns {string}
       **/
      get id() {
        return this.#id;
      }
      /**
       *
       * @type {string}
       **/
      set id(value) {
        this.#id = String(value);
      }
      setId(value) {
        this.id = value;
        return this;
      }
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
      #cycleDuration = "";
      /**
       *
       * @returns {string}
       **/
      get cycleDuration() {
        return this.#cycleDuration;
      }
      /**
       *
       * @type {string}
       **/
      set cycleDuration(value) {
        this.#cycleDuration = String(value);
      }
      setCycleDuration(value) {
        this.cycleDuration = value;
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
        if (d.name !== undefined) {
          this.name = d.name;
        }
        if (d.cycleDuration !== undefined) {
          this.cycleDuration = d.cycleDuration;
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
          cycleDuration: this.#cycleDuration,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          id: "id",
          name: "name",
          cycleDuration: "cycleDuration",
        };
      }
      /**
       * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(possibleDtoObject) {
        return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages(
          possibleDtoObject,
        );
      }
      /**
       * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(partialDtoObject) {
        return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages(
          partialDtoObject,
        );
      }
      copyWith(partial) {
        return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages(
          { ...this.toJSON(), ...partial },
        );
      }
      clone() {
        return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages(
          this.toJSON(),
        );
      }
    };
    /**
     * The base class definition for extraPackages
     **/
    static ExtraPackages = class ExtraPackages {
      /**
       *
       * @type {string}
       **/
      #id = "";
      /**
       *
       * @returns {string}
       **/
      get id() {
        return this.#id;
      }
      /**
       *
       * @type {string}
       **/
      set id(value) {
        this.#id = String(value);
      }
      setId(value) {
        this.id = value;
        return this;
      }
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
      #cycleDuration = "";
      /**
       *
       * @returns {string}
       **/
      get cycleDuration() {
        return this.#cycleDuration;
      }
      /**
       *
       * @type {string}
       **/
      set cycleDuration(value) {
        this.#cycleDuration = String(value);
      }
      setCycleDuration(value) {
        this.cycleDuration = value;
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
        if (d.name !== undefined) {
          this.name = d.name;
        }
        if (d.cycleDuration !== undefined) {
          this.cycleDuration = d.cycleDuration;
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
          cycleDuration: this.#cycleDuration,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          id: "id",
          name: "name",
          cycleDuration: "cycleDuration",
        };
      }
      /**
       * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(possibleDtoObject) {
        return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(
          possibleDtoObject,
        );
      }
      /**
       * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(partialDtoObject) {
        return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(
          partialDtoObject,
        );
      }
      copyWith(partial) {
        return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(
          { ...this.toJSON(), ...partial },
        );
      }
      clone() {
        return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(
          this.toJSON(),
        );
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
      if (d.marketplaceId !== undefined) {
        this.marketplaceId = d.marketplaceId;
      }
      if (d.basePackages !== undefined) {
        this.basePackages = d.basePackages;
      }
      if (d.extraPackages !== undefined) {
        this.extraPackages = d.extraPackages;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        marketplaceId: this.#marketplaceId,
        basePackages: this.#basePackages,
        extraPackages: this.#extraPackages,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        marketplaceId: "marketplaceId",
        basePackages$: "basePackages",
        get basePackages() {
          return withPrefix(
            "additionalMarketplaces.basePackages[:i]",
            GetAllAvailableOfferPromotionPackagesActionRes
              .AdditionalMarketplaces.BasePackages.Fields,
          );
        },
        extraPackages$: "extraPackages",
        get extraPackages() {
          return withPrefix(
            "additionalMarketplaces.extraPackages[:i]",
            GetAllAvailableOfferPromotionPackagesActionRes
              .AdditionalMarketplaces.ExtraPackages.Fields,
          );
        },
      };
    }
    /**
     * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces(
        possibleDtoObject,
      );
    }
    /**
     * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces(
        partialDtoObject,
      );
    }
    copyWith(partial) {
      return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces(
        { ...this.toJSON(), ...partial },
      );
    }
    clone() {
      return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces(
        this.toJSON(),
      );
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
    if (d.marketplaceId !== undefined) {
      this.marketplaceId = d.marketplaceId;
    }
    if (d.basePackages !== undefined) {
      this.basePackages = d.basePackages;
    }
    if (d.extraPackages !== undefined) {
      this.extraPackages = d.extraPackages;
    }
    if (d.additionalMarketplaces !== undefined) {
      this.additionalMarketplaces = d.additionalMarketplaces;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      marketplaceId: this.#marketplaceId,
      basePackages: this.#basePackages,
      extraPackages: this.#extraPackages,
      additionalMarketplaces: this.#additionalMarketplaces,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      marketplaceId: "marketplaceId",
      basePackages$: "basePackages",
      get basePackages() {
        return withPrefix(
          "basePackages[:i]",
          GetAllAvailableOfferPromotionPackagesActionRes.BasePackages.Fields,
        );
      },
      extraPackages$: "extraPackages",
      get extraPackages() {
        return withPrefix(
          "extraPackages[:i]",
          GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages.Fields,
        );
      },
      additionalMarketplaces$: "additionalMarketplaces",
      get additionalMarketplaces() {
        return withPrefix(
          "additionalMarketplaces[:i]",
          GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces
            .Fields,
        );
      },
    };
  }
  /**
   * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new GetAllAvailableOfferPromotionPackagesActionRes(
      possibleDtoObject,
    );
  }
  /**
   * Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new GetAllAvailableOfferPromotionPackagesActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new GetAllAvailableOfferPromotionPackagesActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone() {
    return new GetAllAvailableOfferPromotionPackagesActionRes(this.toJSON());
  }
}
