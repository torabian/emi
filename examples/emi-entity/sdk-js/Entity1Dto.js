import { Entity2Entity } from "./Entity2Entity";
import { MArray, MCollection, MOne } from "./sdk/common/operators";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for entity1Dto
 **/
export class Entity1Dto {
  /**
   *
   * @type {string}
   **/
  #uniqueId = undefined;
  /**
   *
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   *
   * @type {string}
   **/
  set uniqueId(value) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#uniqueId = correctType ? value : String(value);
  }
  setUniqueId(value) {
    this.uniqueId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #title = "";
  /**
   *
   * @returns {string}
   **/
  get title() {
    return this.#title;
  }
  /**
   *
   * @type {string}
   **/
  set title(value) {
    this.#title = String(value);
  }
  setTitle(value) {
    this.title = value;
    return this;
  }
  /**
   *
   * @type {Entity1Dto.Items}
   **/
  #items = MArray.of([]);
  /**
   *
   * @returns {Entity1Dto.Items}
   **/
  get items() {
    return this.#items;
  }
  /**
   *
   * @type {Entity1Dto.Items}
   **/
  set items(value) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof Entity1Dto.Items) {
        this.#items = MArray.of(value);
      } else {
        this.#items = MArray.of(
          value.map((item) => new Entity1Dto.Items(item)),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#items = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast(value);
    if (ok) {
      this.#items = mcastValue;
      return;
    }
    console.warn(
      "Cannot assing value to items, because it needs MArray instance or an Array.",
    );
  }
  setItems(value) {
    this.items = value;
    return this;
  }
  /**
   *
   * @type {Entity1Dto.Items2}
   **/
  #items2 = undefined;
  /**
   *
   * @returns {Entity1Dto.Items2}
   **/
  get items2() {
    return this.#items2;
  }
  /**
   *
   * @type {Entity1Dto.Items2}
   **/
  set items2(value) {
    // For nullable array, we allow explicit undefined or null values
    if (value === null || value === undefined) {
      this.#items2 = value;
      return;
    }
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof Entity1Dto.Items2) {
        this.#items2 = MArray.of(value);
      } else {
        this.#items2 = MArray.of(
          value.map((item) => new Entity1Dto.Items2(item)),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#items2 = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast(value);
    if (ok) {
      this.#items2 = mcastValue;
      return;
    }
    console.warn(
      "Cannot assing value to items2, because it needs MArray instance or an Array.",
    );
  }
  setItems2(value) {
    this.items2 = value;
    return this;
  }
  /**
   *
   * @type {Entity2Entity[]}
   **/
  #items3 = MCollection.of([]);
  /**
   *
   * @returns {Entity2Entity[]}
   **/
  get items3() {
    return this.#items3;
  }
  /**
   *
   * @type {Entity2Entity[]}
   **/
  set items3(value) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof Entity2Entity) {
        this.#items3 = MCollection.of(value);
      } else {
        this.#items3 = MCollection.of(
          value.map((item) => new Entity2Entity(item)),
        );
      }
      return;
    }
    // If the instance is already an MCollection, we assume it's all good.
    if (value instanceof MCollection) {
      this.#items3 = value;
      return;
    }
    // If the value is not array, and is not a MCollection, we need to be consider,
    // it might be eligible to be casted into MCollection.
    const { ok, value: mcastValue } = MCollection.cast(value);
    if (ok) {
      this.#items3 = mcastValue;
      return;
    }
    console.warn(
      "Cannot assing value to items3, because it needs MCollection instance or an Array.",
    );
  }
  setItems3(value) {
    this.items3 = value;
    return this;
  }
  /**
   *
   * @type {Entity2Entity[]}
   **/
  #items4 = undefined;
  /**
   *
   * @returns {Entity2Entity[]}
   **/
  get items4() {
    return this.#items4;
  }
  /**
   *
   * @type {Entity2Entity[]}
   **/
  set items4(value) {
    // For nullable collection, we allow explicit undefined or null values
    if (value === null || value === undefined) {
      this.#items4 = value;
      return;
    }
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof Entity2Entity) {
        this.#items4 = MCollection.of(value);
      } else {
        this.#items4 = MCollection.of(
          value.map((item) => new Entity2Entity(item)),
        );
      }
      return;
    }
    // If the instance is already an MCollection, we assume it's all good.
    if (value instanceof MCollection) {
      this.#items4 = value;
      return;
    }
    // If the value is not array, and is not a MCollection, we need to be consider,
    // it might be eligible to be casted into MCollection.
    const { ok, value: mcastValue } = MCollection.cast(value);
    if (ok) {
      this.#items4 = mcastValue;
      return;
    }
    console.warn(
      "Cannot assing value to items4, because it needs MCollection instance or an Array.",
    );
  }
  setItems4(value) {
    this.items4 = value;
    return this;
  }
  /**
   *
   * @type {Entity2Entity}
   **/
  #owner;
  /**
   *
   * @returns {Entity2Entity}
   **/
  get owner() {
    return this.#owner;
  }
  /**
   *
   * @type {Entity2Entity}
   **/
  set owner(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#owner = value;
    } else if (value instanceof Entity2Entity) {
      this.#owner = MOne.of(value);
    } else {
      this.#owner = MOne.of(new Entity2Entity(value));
    }
  }
  setOwner(value) {
    this.owner = value;
    return this;
  }
  /**
   *
   * @type {Entity2Entity}
   **/
  #manager = undefined;
  /**
   *
   * @returns {Entity2Entity}
   **/
  get manager() {
    return this.#manager;
  }
  /**
   *
   * @type {Entity2Entity}
   **/
  set manager(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#manager = value;
    } else if (value instanceof Entity2Entity) {
      this.#manager = MOne.of(value);
    } else {
      this.#manager = MOne.of(new Entity2Entity(value));
    }
  }
  setManager(value) {
    this.manager = value;
    return this;
  }
  /**
   *
   * @type {Entity1Dto.Content1}
   **/
  #content1;
  /**
   *
   * @returns {Entity1Dto.Content1}
   **/
  get content1() {
    return this.#content1;
  }
  /**
   *
   * @type {Entity1Dto.Content1}
   **/
  set content1(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof Entity1Dto.Content1) {
      this.#content1 = value;
    } else {
      this.#content1 = new Entity1Dto.Content1(value);
    }
  }
  setContent1(value) {
    this.content1 = value;
    return this;
  }
  /**
   *
   * @type {Entity1Dto.Content2}
   **/
  #content2 = undefined;
  /**
   *
   * @returns {Entity1Dto.Content2}
   **/
  get content2() {
    return this.#content2;
  }
  /**
   *
   * @type {Entity1Dto.Content2}
   **/
  set content2(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof Entity1Dto.Content2) {
      this.#content2 = value;
    } else {
      this.#content2 = new Entity1Dto.Content2(value);
    }
  }
  setContent2(value) {
    this.content2 = value;
    return this;
  }
  /**
   *
   * @type {Money}
   **/
  #complex1;
  /**
   *
   * @returns {Money}
   **/
  get complex1() {
    return this.#complex1;
  }
  /**
   *
   * @type {Money}
   **/
  set complex1(value) {
    this.#complex1 = value;
  }
  setComplex1(value) {
    this.complex1 = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #subtitle = undefined;
  /**
   *
   * @returns {string}
   **/
  get subtitle() {
    return this.#subtitle;
  }
  /**
   *
   * @type {string}
   **/
  set subtitle(value) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#subtitle = correctType ? value : String(value);
  }
  setSubtitle(value) {
    this.subtitle = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #isActive;
  /**
   *
   * @returns {boolean}
   **/
  get isActive() {
    return this.#isActive;
  }
  /**
   *
   * @type {boolean}
   **/
  set isActive(value) {
    this.#isActive = Boolean(value);
  }
  setIsActive(value) {
    this.isActive = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #isFeatured = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get isFeatured() {
    return this.#isFeatured;
  }
  /**
   *
   * @type {boolean}
   **/
  set isFeatured(value) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#isFeatured = correctType ? value : Boolean(value);
  }
  setIsFeatured(value) {
    this.isFeatured = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #viewCount = 0;
  /**
   *
   * @returns {number}
   **/
  get viewCount() {
    return this.#viewCount;
  }
  /**
   *
   * @type {number}
   **/
  set viewCount(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#viewCount = parsedValue;
    }
  }
  setViewCount(value) {
    this.viewCount = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #viewCountOpt = undefined;
  /**
   *
   * @returns {number}
   **/
  get viewCountOpt() {
    return this.#viewCountOpt;
  }
  /**
   *
   * @type {number}
   **/
  set viewCountOpt(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#viewCountOpt = parsedValue;
    }
  }
  setViewCountOpt(value) {
    this.viewCountOpt = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #smallCount = 0;
  /**
   *
   * @returns {number}
   **/
  get smallCount() {
    return this.#smallCount;
  }
  /**
   *
   * @type {number}
   **/
  set smallCount(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#smallCount = parsedValue;
    }
  }
  setSmallCount(value) {
    this.smallCount = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #smallCountOpt = undefined;
  /**
   *
   * @returns {number}
   **/
  get smallCountOpt() {
    return this.#smallCountOpt;
  }
  /**
   *
   * @type {number}
   **/
  set smallCountOpt(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#smallCountOpt = parsedValue;
    }
  }
  setSmallCountOpt(value) {
    this.smallCountOpt = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #bigCount = 0;
  /**
   *
   * @returns {number}
   **/
  get bigCount() {
    return this.#bigCount;
  }
  /**
   *
   * @type {number}
   **/
  set bigCount(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#bigCount = parsedValue;
    }
  }
  setBigCount(value) {
    this.bigCount = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #bigCountOpt = undefined;
  /**
   *
   * @returns {number}
   **/
  get bigCountOpt() {
    return this.#bigCountOpt;
  }
  /**
   *
   * @type {number}
   **/
  set bigCountOpt(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#bigCountOpt = parsedValue;
    }
  }
  setBigCountOpt(value) {
    this.bigCountOpt = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #ratio32 = 0.0;
  /**
   *
   * @returns {number}
   **/
  get ratio32() {
    return this.#ratio32;
  }
  /**
   *
   * @type {number}
   **/
  set ratio32(value) {
    this.#ratio32 = value;
  }
  setRatio32(value) {
    this.ratio32 = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #ratio32Opt = undefined;
  /**
   *
   * @returns {number}
   **/
  get ratio32Opt() {
    return this.#ratio32Opt;
  }
  /**
   *
   * @type {number}
   **/
  set ratio32Opt(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#ratio32Opt = parsedValue;
    }
  }
  setRatio32Opt(value) {
    this.ratio32Opt = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #ratio64 = 0.0;
  /**
   *
   * @returns {number}
   **/
  get ratio64() {
    return this.#ratio64;
  }
  /**
   *
   * @type {number}
   **/
  set ratio64(value) {
    this.#ratio64 = value;
  }
  setRatio64(value) {
    this.ratio64 = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #ratio64Opt = undefined;
  /**
   *
   * @returns {number}
   **/
  get ratio64Opt() {
    return this.#ratio64Opt;
  }
  /**
   *
   * @type {number}
   **/
  set ratio64Opt(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#ratio64Opt = parsedValue;
    }
  }
  setRatio64Opt(value) {
    this.ratio64Opt = value;
    return this;
  }
  /**
   *
   * @type {"active" | "inactive"}
   **/
  #status;
  /**
   *
   * @returns {"active" | "inactive"}
   **/
  get status() {
    return this.#status;
  }
  /**
   *
   * @type {"active" | "inactive"}
   **/
  set status(value) {
    this.#status = value;
  }
  setStatus(value) {
    this.status = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #statusOpt = undefined;
  /**
   *
   * @returns {any}
   **/
  get statusOpt() {
    return this.#statusOpt;
  }
  /**
   *
   * @type {any}
   **/
  set statusOpt(value) {
    this.#statusOpt = value;
  }
  setStatusOpt(value) {
    this.statusOpt = value;
    return this;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  #metadata;
  /**
   *
   * @returns {{[key: string]: any}}
   **/
  get metadata() {
    return this.#metadata;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  set metadata(value) {
    this.#metadata = value;
  }
  setMetadata(value) {
    this.metadata = value;
    return this;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  #metadataOpt = undefined;
  /**
   *
   * @returns {{[key: string]: any}}
   **/
  get metadataOpt() {
    return this.#metadataOpt;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  set metadataOpt(value) {
    this.#metadataOpt = value;
  }
  setMetadataOpt(value) {
    this.metadataOpt = value;
    return this;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  #rawSettings;
  /**
   *
   * @returns {{[key: string]: any}}
   **/
  get rawSettings() {
    return this.#rawSettings;
  }
  /**
   *
   * @type {{[key: string]: any}}
   **/
  set rawSettings(value) {
    this.#rawSettings = value;
  }
  setRawSettings(value) {
    this.rawSettings = value;
    return this;
  }
  /**
   *
   * @type {string[]}
   **/
  #labels = [];
  /**
   *
   * @returns {string[]}
   **/
  get labels() {
    return this.#labels;
  }
  /**
   *
   * @type {string[]}
   **/
  set labels(value) {
    this.#labels = value;
  }
  setLabels(value) {
    this.labels = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #labelsOpt = undefined;
  /**
   *
   * @returns {any}
   **/
  get labelsOpt() {
    return this.#labelsOpt;
  }
  /**
   *
   * @type {any}
   **/
  set labelsOpt(value) {
    this.#labelsOpt = value;
  }
  setLabelsOpt(value) {
    this.labelsOpt = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #misc = null;
  /**
   *
   * @returns {any}
   **/
  get misc() {
    return this.#misc;
  }
  /**
   *
   * @type {any}
   **/
  set misc(value) {
    this.#misc = value;
  }
  setMisc(value) {
    this.misc = value;
    return this;
  }
  /**
   *
   * @type {Entity1Dto.NestedContainer}
   **/
  #nestedContainer;
  /**
   *
   * @returns {Entity1Dto.NestedContainer}
   **/
  get nestedContainer() {
    return this.#nestedContainer;
  }
  /**
   *
   * @type {Entity1Dto.NestedContainer}
   **/
  set nestedContainer(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof Entity1Dto.NestedContainer) {
      this.#nestedContainer = value;
    } else {
      this.#nestedContainer = new Entity1Dto.NestedContainer(value);
    }
  }
  setNestedContainer(value) {
    this.nestedContainer = value;
    return this;
  }
  /**
   *
   * @type {Entity1Dto.NestedContainerOpt}
   **/
  #nestedContainerOpt = undefined;
  /**
   *
   * @returns {Entity1Dto.NestedContainerOpt}
   **/
  get nestedContainerOpt() {
    return this.#nestedContainerOpt;
  }
  /**
   *
   * @type {Entity1Dto.NestedContainerOpt}
   **/
  set nestedContainerOpt(value) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof Entity1Dto.NestedContainerOpt) {
      this.#nestedContainerOpt = value;
    } else {
      this.#nestedContainerOpt = new Entity1Dto.NestedContainerOpt(value);
    }
  }
  setNestedContainerOpt(value) {
    this.nestedContainerOpt = value;
    return this;
  }
  /**
   * The base class definition for items
   **/
  static Items = class Items {
    /**
     *
     * @type {string}
     **/
    #item2 = "";
    /**
     *
     * @returns {string}
     **/
    get item2() {
      return this.#item2;
    }
    /**
     *
     * @type {string}
     **/
    set item2(value) {
      this.#item2 = String(value);
    }
    setItem2(value) {
      this.item2 = value;
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
      if (d.item2 !== undefined) {
        this.item2 = d.item2;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        item2: this.#item2,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        item2: "item2",
      };
    }
    /**
     * Creates an instance of Entity1Dto.Items, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new Entity1Dto.Items(possibleDtoObject);
    }
    /**
     * Creates an instance of Entity1Dto.Items, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new Entity1Dto.Items(partialDtoObject);
    }
    copyWith(partial) {
      return new Entity1Dto.Items({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new Entity1Dto.Items(this.toJSON());
    }
  };
  /**
   * The base class definition for items2
   **/
  static Items2 = class Items2 {
    /**
     *
     * @type {string}
     **/
    #item2 = "";
    /**
     *
     * @returns {string}
     **/
    get item2() {
      return this.#item2;
    }
    /**
     *
     * @type {string}
     **/
    set item2(value) {
      this.#item2 = String(value);
    }
    setItem2(value) {
      this.item2 = value;
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
      if (d.item2 !== undefined) {
        this.item2 = d.item2;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        item2: this.#item2,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        item2: "item2",
      };
    }
    /**
     * Creates an instance of Entity1Dto.Items2, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new Entity1Dto.Items2(possibleDtoObject);
    }
    /**
     * Creates an instance of Entity1Dto.Items2, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new Entity1Dto.Items2(partialDtoObject);
    }
    copyWith(partial) {
      return new Entity1Dto.Items2({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new Entity1Dto.Items2(this.toJSON());
    }
  };
  /**
   * The base class definition for content1
   **/
  static Content1 = class Content1 {
    /**
     *
     * @type {number}
     **/
    #item1 = undefined;
    /**
     *
     * @returns {number}
     **/
    get item1() {
      return this.#item1;
    }
    /**
     *
     * @type {number}
     **/
    set item1(value) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#item1 = parsedValue;
      }
    }
    setItem1(value) {
      this.item1 = value;
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
      if (d.item1 !== undefined) {
        this.item1 = d.item1;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        item1: this.#item1,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        item1: "item1",
      };
    }
    /**
     * Creates an instance of Entity1Dto.Content1, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new Entity1Dto.Content1(possibleDtoObject);
    }
    /**
     * Creates an instance of Entity1Dto.Content1, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new Entity1Dto.Content1(partialDtoObject);
    }
    copyWith(partial) {
      return new Entity1Dto.Content1({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new Entity1Dto.Content1(this.toJSON());
    }
  };
  /**
   * The base class definition for content2
   **/
  static Content2 = class Content2 {
    /**
     *
     * @type {number}
     **/
    #item2 = undefined;
    /**
     *
     * @returns {number}
     **/
    get item2() {
      return this.#item2;
    }
    /**
     *
     * @type {number}
     **/
    set item2(value) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#item2 = parsedValue;
      }
    }
    setItem2(value) {
      this.item2 = value;
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
      if (d.item2 !== undefined) {
        this.item2 = d.item2;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        item2: this.#item2,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        item2: "item2",
      };
    }
    /**
     * Creates an instance of Entity1Dto.Content2, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new Entity1Dto.Content2(possibleDtoObject);
    }
    /**
     * Creates an instance of Entity1Dto.Content2, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new Entity1Dto.Content2(partialDtoObject);
    }
    copyWith(partial) {
      return new Entity1Dto.Content2({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new Entity1Dto.Content2(this.toJSON());
    }
  };
  /**
   * The base class definition for nestedContainer
   **/
  static NestedContainer = class NestedContainer {
    /**
     *
     * @type {Entity1Dto.NestedContainer.NestedInner}
     **/
    #nestedInner;
    /**
     *
     * @returns {Entity1Dto.NestedContainer.NestedInner}
     **/
    get nestedInner() {
      return this.#nestedInner;
    }
    /**
     *
     * @type {Entity1Dto.NestedContainer.NestedInner}
     **/
    set nestedInner(value) {
      // For objects, the sub type needs to always be instance of the sub class.
      if (value instanceof Entity1Dto.NestedContainer.NestedInner) {
        this.#nestedInner = value;
      } else {
        this.#nestedInner = new Entity1Dto.NestedContainer.NestedInner(value);
      }
    }
    setNestedInner(value) {
      this.nestedInner = value;
      return this;
    }
    /**
     * The base class definition for nestedInner
     **/
    static NestedInner = class NestedInner {
      /**
       *
       * @type {Entity1Dto.NestedContainer.NestedInner.NestedItems}
       **/
      #nestedItems = MArray.of([]);
      /**
       *
       * @returns {Entity1Dto.NestedContainer.NestedInner.NestedItems}
       **/
      get nestedItems() {
        return this.#nestedItems;
      }
      /**
       *
       * @type {Entity1Dto.NestedContainer.NestedInner.NestedItems}
       **/
      set nestedItems(value) {
        // When the passed value is already an array, we check if we need to
        // cast the inner items into class instance.
        if (Array.isArray(value)) {
          if (
            value.length > 0 &&
            value[0] instanceof
              Entity1Dto.NestedContainer.NestedInner.NestedItems
          ) {
            this.#nestedItems = MArray.of(value);
          } else {
            this.#nestedItems = MArray.of(
              value.map(
                (item) =>
                  new Entity1Dto.NestedContainer.NestedInner.NestedItems(item),
              ),
            );
          }
          return;
        }
        // If the instance is already an MArray, we assume it's all good.
        if (value instanceof MArray) {
          this.#nestedItems = value;
          return;
        }
        // If the value is not array, and is not a MArray, we need to be consider,
        // it might be eligible to be casted into MArray.
        const { ok, value: mcastValue } = MArray.cast(value);
        if (ok) {
          this.#nestedItems = mcastValue;
          return;
        }
        console.warn(
          "Cannot assing value to nestedItems, because it needs MArray instance or an Array.",
        );
      }
      setNestedItems(value) {
        this.nestedItems = value;
        return this;
      }
      /**
       *
       * @type {Entity2Entity}
       **/
      #nestedOwner;
      /**
       *
       * @returns {Entity2Entity}
       **/
      get nestedOwner() {
        return this.#nestedOwner;
      }
      /**
       *
       * @type {Entity2Entity}
       **/
      set nestedOwner(value) {
        // For objects, the sub type needs to always be instance of the sub class.
        if (value instanceof MOne) {
          this.#nestedOwner = value;
        } else if (value instanceof Entity2Entity) {
          this.#nestedOwner = MOne.of(value);
        } else {
          this.#nestedOwner = MOne.of(new Entity2Entity(value));
        }
      }
      setNestedOwner(value) {
        this.nestedOwner = value;
        return this;
      }
      /**
       * The base class definition for nestedItems
       **/
      static NestedItems = class NestedItems {
        /**
         *
         * @type {string}
         **/
        #label = "";
        /**
         *
         * @returns {string}
         **/
        get label() {
          return this.#label;
        }
        /**
         *
         * @type {string}
         **/
        set label(value) {
          this.#label = String(value);
        }
        setLabel(value) {
          this.label = value;
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
          if (d.label !== undefined) {
            this.label = d.label;
          }
        }
        /**
         *	Special toJSON override, since the field are private,
         *	Json stringify won't see them unless we mention it explicitly.
         **/
        toJSON() {
          return {
            label: this.#label,
          };
        }
        toString() {
          return JSON.stringify(this);
        }
        static get Fields() {
          return {
            label: "label",
          };
        }
        /**
         * Creates an instance of Entity1Dto.NestedContainer.NestedInner.NestedItems, and possibleDtoObject
         * needs to satisfy the type requirement fully, otherwise typescript compile would
         * be complaining.
         **/
        static from(possibleDtoObject) {
          return new Entity1Dto.NestedContainer.NestedInner.NestedItems(
            possibleDtoObject,
          );
        }
        /**
         * Creates an instance of Entity1Dto.NestedContainer.NestedInner.NestedItems, and partialDtoObject
         * needs to satisfy the type, but partially, and rest of the content would
         * be constructed according to data types and nullability.
         **/
        static with(partialDtoObject) {
          return new Entity1Dto.NestedContainer.NestedInner.NestedItems(
            partialDtoObject,
          );
        }
        copyWith(partial) {
          return new Entity1Dto.NestedContainer.NestedInner.NestedItems({
            ...this.toJSON(),
            ...partial,
          });
        }
        clone() {
          return new Entity1Dto.NestedContainer.NestedInner.NestedItems(
            this.toJSON(),
          );
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
        if (d.nestedItems !== undefined) {
          this.nestedItems = d.nestedItems;
        }
        if (d.nestedOwner !== undefined) {
          this.nestedOwner = d.nestedOwner;
        }
        this.#lateInitFields(data);
      }
      /**
       * These are the class instances, which need to be initialised, regardless of the constructor incoming data
       **/
      #lateInitFields(data = {}) {
        const d = data;
        if (!(d.nestedOwner instanceof Entity2Entity)) {
          this.nestedOwner = MOne.of(new Entity2Entity(d.nestedOwner || {}));
        }
      }
      /**
       *	Special toJSON override, since the field are private,
       *	Json stringify won't see them unless we mention it explicitly.
       **/
      toJSON() {
        return {
          nestedItems: this.#nestedItems,
          nestedOwner: this.#nestedOwner,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          nestedItems$: "nestedItems",
          get nestedItems() {
            return withPrefix(
              "nestedContainer.nestedInner.nestedItems[:i]",
              Entity1Dto.NestedContainer.NestedInner.NestedItems.Fields,
            );
          },
          nestedOwner$: "nestedOwner",
          get nestedOwner() {
            return withPrefix(
              "nestedContainer.nestedInner.nestedOwner",
              Entity2Entity.Fields,
            );
          },
        };
      }
      /**
       * Creates an instance of Entity1Dto.NestedContainer.NestedInner, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(possibleDtoObject) {
        return new Entity1Dto.NestedContainer.NestedInner(possibleDtoObject);
      }
      /**
       * Creates an instance of Entity1Dto.NestedContainer.NestedInner, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(partialDtoObject) {
        return new Entity1Dto.NestedContainer.NestedInner(partialDtoObject);
      }
      copyWith(partial) {
        return new Entity1Dto.NestedContainer.NestedInner({
          ...this.toJSON(),
          ...partial,
        });
      }
      clone() {
        return new Entity1Dto.NestedContainer.NestedInner(this.toJSON());
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
      if (d.nestedInner !== undefined) {
        this.nestedInner = d.nestedInner;
      }
      this.#lateInitFields(data);
    }
    /**
     * These are the class instances, which need to be initialised, regardless of the constructor incoming data
     **/
    #lateInitFields(data = {}) {
      const d = data;
      if (!(d.nestedInner instanceof Entity1Dto.NestedContainer.NestedInner)) {
        this.nestedInner = new Entity1Dto.NestedContainer.NestedInner(
          d.nestedInner || {},
        );
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        nestedInner: this.#nestedInner,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        nestedInner$: "nestedInner",
        get nestedInner() {
          return withPrefix(
            "nestedContainer.nestedInner",
            Entity1Dto.NestedContainer.NestedInner.Fields,
          );
        },
      };
    }
    /**
     * Creates an instance of Entity1Dto.NestedContainer, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new Entity1Dto.NestedContainer(possibleDtoObject);
    }
    /**
     * Creates an instance of Entity1Dto.NestedContainer, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new Entity1Dto.NestedContainer(partialDtoObject);
    }
    copyWith(partial) {
      return new Entity1Dto.NestedContainer({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new Entity1Dto.NestedContainer(this.toJSON());
    }
  };
  /**
   * The base class definition for nestedContainerOpt
   **/
  static NestedContainerOpt = class NestedContainerOpt {
    /**
     *
     * @type {Entity1Dto.NestedContainerOpt.NestedInner}
     **/
    #nestedInner;
    /**
     *
     * @returns {Entity1Dto.NestedContainerOpt.NestedInner}
     **/
    get nestedInner() {
      return this.#nestedInner;
    }
    /**
     *
     * @type {Entity1Dto.NestedContainerOpt.NestedInner}
     **/
    set nestedInner(value) {
      // For objects, the sub type needs to always be instance of the sub class.
      if (value instanceof Entity1Dto.NestedContainerOpt.NestedInner) {
        this.#nestedInner = value;
      } else {
        this.#nestedInner = new Entity1Dto.NestedContainerOpt.NestedInner(
          value,
        );
      }
    }
    setNestedInner(value) {
      this.nestedInner = value;
      return this;
    }
    /**
     * The base class definition for nestedInner
     **/
    static NestedInner = class NestedInner {
      /**
       *
       * @type {Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt}
       **/
      #nestedItemsOpt = MArray.of([]);
      /**
       *
       * @returns {Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt}
       **/
      get nestedItemsOpt() {
        return this.#nestedItemsOpt;
      }
      /**
       *
       * @type {Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt}
       **/
      set nestedItemsOpt(value) {
        // When the passed value is already an array, we check if we need to
        // cast the inner items into class instance.
        if (Array.isArray(value)) {
          if (
            value.length > 0 &&
            value[0] instanceof
              Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt
          ) {
            this.#nestedItemsOpt = MArray.of(value);
          } else {
            this.#nestedItemsOpt = MArray.of(
              value.map(
                (item) =>
                  new Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt(
                    item,
                  ),
              ),
            );
          }
          return;
        }
        // If the instance is already an MArray, we assume it's all good.
        if (value instanceof MArray) {
          this.#nestedItemsOpt = value;
          return;
        }
        // If the value is not array, and is not a MArray, we need to be consider,
        // it might be eligible to be casted into MArray.
        const { ok, value: mcastValue } = MArray.cast(value);
        if (ok) {
          this.#nestedItemsOpt = mcastValue;
          return;
        }
        console.warn(
          "Cannot assing value to nestedItemsOpt, because it needs MArray instance or an Array.",
        );
      }
      setNestedItemsOpt(value) {
        this.nestedItemsOpt = value;
        return this;
      }
      /**
       * The base class definition for nestedItemsOpt
       **/
      static NestedItemsOpt = class NestedItemsOpt {
        /**
         *
         * @type {string}
         **/
        #label = "";
        /**
         *
         * @returns {string}
         **/
        get label() {
          return this.#label;
        }
        /**
         *
         * @type {string}
         **/
        set label(value) {
          this.#label = String(value);
        }
        setLabel(value) {
          this.label = value;
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
          if (d.label !== undefined) {
            this.label = d.label;
          }
        }
        /**
         *	Special toJSON override, since the field are private,
         *	Json stringify won't see them unless we mention it explicitly.
         **/
        toJSON() {
          return {
            label: this.#label,
          };
        }
        toString() {
          return JSON.stringify(this);
        }
        static get Fields() {
          return {
            label: "label",
          };
        }
        /**
         * Creates an instance of Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt, and possibleDtoObject
         * needs to satisfy the type requirement fully, otherwise typescript compile would
         * be complaining.
         **/
        static from(possibleDtoObject) {
          return new Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt(
            possibleDtoObject,
          );
        }
        /**
         * Creates an instance of Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt, and partialDtoObject
         * needs to satisfy the type, but partially, and rest of the content would
         * be constructed according to data types and nullability.
         **/
        static with(partialDtoObject) {
          return new Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt(
            partialDtoObject,
          );
        }
        copyWith(partial) {
          return new Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt({
            ...this.toJSON(),
            ...partial,
          });
        }
        clone() {
          return new Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt(
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
        if (d.nestedItemsOpt !== undefined) {
          this.nestedItemsOpt = d.nestedItemsOpt;
        }
      }
      /**
       *	Special toJSON override, since the field are private,
       *	Json stringify won't see them unless we mention it explicitly.
       **/
      toJSON() {
        return {
          nestedItemsOpt: this.#nestedItemsOpt,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          nestedItemsOpt$: "nestedItemsOpt",
          get nestedItemsOpt() {
            return withPrefix(
              "nestedContainerOpt.nestedInner.nestedItemsOpt[:i]",
              Entity1Dto.NestedContainerOpt.NestedInner.NestedItemsOpt.Fields,
            );
          },
        };
      }
      /**
       * Creates an instance of Entity1Dto.NestedContainerOpt.NestedInner, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(possibleDtoObject) {
        return new Entity1Dto.NestedContainerOpt.NestedInner(possibleDtoObject);
      }
      /**
       * Creates an instance of Entity1Dto.NestedContainerOpt.NestedInner, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(partialDtoObject) {
        return new Entity1Dto.NestedContainerOpt.NestedInner(partialDtoObject);
      }
      copyWith(partial) {
        return new Entity1Dto.NestedContainerOpt.NestedInner({
          ...this.toJSON(),
          ...partial,
        });
      }
      clone() {
        return new Entity1Dto.NestedContainerOpt.NestedInner(this.toJSON());
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
      if (d.nestedInner !== undefined) {
        this.nestedInner = d.nestedInner;
      }
      this.#lateInitFields(data);
    }
    /**
     * These are the class instances, which need to be initialised, regardless of the constructor incoming data
     **/
    #lateInitFields(data = {}) {
      const d = data;
      if (
        !(d.nestedInner instanceof Entity1Dto.NestedContainerOpt.NestedInner)
      ) {
        this.nestedInner = new Entity1Dto.NestedContainerOpt.NestedInner(
          d.nestedInner || {},
        );
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        nestedInner: this.#nestedInner,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        nestedInner$: "nestedInner",
        get nestedInner() {
          return withPrefix(
            "nestedContainerOpt.nestedInner",
            Entity1Dto.NestedContainerOpt.NestedInner.Fields,
          );
        },
      };
    }
    /**
     * Creates an instance of Entity1Dto.NestedContainerOpt, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new Entity1Dto.NestedContainerOpt(possibleDtoObject);
    }
    /**
     * Creates an instance of Entity1Dto.NestedContainerOpt, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new Entity1Dto.NestedContainerOpt(partialDtoObject);
    }
    copyWith(partial) {
      return new Entity1Dto.NestedContainerOpt({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone() {
      return new Entity1Dto.NestedContainerOpt(this.toJSON());
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
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.title !== undefined) {
      this.title = d.title;
    }
    if (d.items !== undefined) {
      this.items = d.items;
    }
    if (d.items2 !== undefined) {
      this.items2 = d.items2;
    }
    if (d.items3 !== undefined) {
      this.items3 = d.items3;
    }
    if (d.items4 !== undefined) {
      this.items4 = d.items4;
    }
    if (d.owner !== undefined) {
      this.owner = d.owner;
    }
    if (d.manager !== undefined) {
      this.manager = d.manager;
    }
    if (d.content1 !== undefined) {
      this.content1 = d.content1;
    }
    if (d.content2 !== undefined) {
      this.content2 = d.content2;
    }
    if (d.complex1 !== undefined) {
      this.complex1 = d.complex1;
    }
    if (d.subtitle !== undefined) {
      this.subtitle = d.subtitle;
    }
    if (d.isActive !== undefined) {
      this.isActive = d.isActive;
    }
    if (d.isFeatured !== undefined) {
      this.isFeatured = d.isFeatured;
    }
    if (d.viewCount !== undefined) {
      this.viewCount = d.viewCount;
    }
    if (d.viewCountOpt !== undefined) {
      this.viewCountOpt = d.viewCountOpt;
    }
    if (d.smallCount !== undefined) {
      this.smallCount = d.smallCount;
    }
    if (d.smallCountOpt !== undefined) {
      this.smallCountOpt = d.smallCountOpt;
    }
    if (d.bigCount !== undefined) {
      this.bigCount = d.bigCount;
    }
    if (d.bigCountOpt !== undefined) {
      this.bigCountOpt = d.bigCountOpt;
    }
    if (d.ratio32 !== undefined) {
      this.ratio32 = d.ratio32;
    }
    if (d.ratio32Opt !== undefined) {
      this.ratio32Opt = d.ratio32Opt;
    }
    if (d.ratio64 !== undefined) {
      this.ratio64 = d.ratio64;
    }
    if (d.ratio64Opt !== undefined) {
      this.ratio64Opt = d.ratio64Opt;
    }
    if (d.status !== undefined) {
      this.status = d.status;
    }
    if (d.statusOpt !== undefined) {
      this.statusOpt = d.statusOpt;
    }
    if (d.metadata !== undefined) {
      this.metadata = d.metadata;
    }
    if (d.metadataOpt !== undefined) {
      this.metadataOpt = d.metadataOpt;
    }
    if (d.rawSettings !== undefined) {
      this.rawSettings = d.rawSettings;
    }
    if (d.labels !== undefined) {
      this.labels = d.labels;
    }
    if (d.labelsOpt !== undefined) {
      this.labelsOpt = d.labelsOpt;
    }
    if (d.misc !== undefined) {
      this.misc = d.misc;
    }
    if (d.nestedContainer !== undefined) {
      this.nestedContainer = d.nestedContainer;
    }
    if (d.nestedContainerOpt !== undefined) {
      this.nestedContainerOpt = d.nestedContainerOpt;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data;
    if (!(d.owner instanceof Entity2Entity)) {
      this.owner = MOne.of(new Entity2Entity(d.owner || {}));
    }
    if (!(d.content1 instanceof Entity1Dto.Content1)) {
      this.content1 = new Entity1Dto.Content1(d.content1 || {});
    }
    if (!(d.nestedContainer instanceof Entity1Dto.NestedContainer)) {
      this.nestedContainer = new Entity1Dto.NestedContainer(
        d.nestedContainer || {},
      );
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      title: this.#title,
      items: this.#items,
      items2: this.#items2,
      items3: this.#items3,
      items4: this.#items4,
      owner: this.#owner,
      manager: this.#manager,
      content1: this.#content1,
      content2: this.#content2,
      complex1: this.#complex1,
      subtitle: this.#subtitle,
      isActive: this.#isActive,
      isFeatured: this.#isFeatured,
      viewCount: this.#viewCount,
      viewCountOpt: this.#viewCountOpt,
      smallCount: this.#smallCount,
      smallCountOpt: this.#smallCountOpt,
      bigCount: this.#bigCount,
      bigCountOpt: this.#bigCountOpt,
      ratio32: this.#ratio32,
      ratio32Opt: this.#ratio32Opt,
      ratio64: this.#ratio64,
      ratio64Opt: this.#ratio64Opt,
      status: this.#status,
      statusOpt: this.#statusOpt,
      metadata: this.#metadata,
      metadataOpt: this.#metadataOpt,
      rawSettings: this.#rawSettings,
      labels: this.#labels,
      labelsOpt: this.#labelsOpt,
      misc: this.#misc,
      nestedContainer: this.#nestedContainer,
      nestedContainerOpt: this.#nestedContainerOpt,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      title: "title",
      items$: "items",
      get items() {
        return withPrefix("items[:i]", Entity1Dto.Items.Fields);
      },
      items2$: "items2",
      get items2() {
        return withPrefix("items2[:i]", Entity1Dto.Items2.Fields);
      },
      items3$: "items3",
      get items3() {
        return withPrefix("items3[:i]", Entity2Entity.Fields);
      },
      items4$: "items4",
      get items4() {
        return withPrefix("items4", Entity2Entity.Fields);
      },
      owner$: "owner",
      get owner() {
        return withPrefix("owner", Entity2Entity.Fields);
      },
      manager: "manager",
      content1$: "content1",
      get content1() {
        return withPrefix("content1", Entity1Dto.Content1.Fields);
      },
      content2$: "content2",
      get content2() {
        return withPrefix("content2", Entity1Dto.Content2.Fields);
      },
      complex1: "complex1",
      subtitle: "subtitle",
      isActive: "isActive",
      isFeatured: "isFeatured",
      viewCount: "viewCount",
      viewCountOpt: "viewCountOpt",
      smallCount: "smallCount",
      smallCountOpt: "smallCountOpt",
      bigCount: "bigCount",
      bigCountOpt: "bigCountOpt",
      ratio32: "ratio32",
      ratio32Opt: "ratio32Opt",
      ratio64: "ratio64",
      ratio64Opt: "ratio64Opt",
      status: "status",
      statusOpt: "statusOpt",
      metadata: "metadata",
      metadataOpt: "metadataOpt",
      rawSettings: "rawSettings",
      labels$: "labels",
      get labels() {
        return "labels[:i]";
      },
      labelsOpt: "labelsOpt",
      misc: "misc",
      nestedContainer$: "nestedContainer",
      get nestedContainer() {
        return withPrefix("nestedContainer", Entity1Dto.NestedContainer.Fields);
      },
      nestedContainerOpt$: "nestedContainerOpt",
      get nestedContainerOpt() {
        return withPrefix(
          "nestedContainerOpt",
          Entity1Dto.NestedContainerOpt.Fields,
        );
      },
    };
  }
  /**
   * Creates an instance of Entity1Dto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new Entity1Dto(possibleDtoObject);
  }
  /**
   * Creates an instance of Entity1Dto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new Entity1Dto(partialDtoObject);
  }
  copyWith(partial) {
    return new Entity1Dto({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new Entity1Dto(this.toJSON());
  }
}
