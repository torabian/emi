import { MArray } from "./sdk/common/operators";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * Action to communicate with the action entity2Query
 */
/**
 * Entity2QueryAction
 */
export class Entity2QueryAction {
  //
  static URL = "/entity2/query";
  static NewUrl = (qs) => buildUrl(Entity2QueryAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity2QueryAction.NewUrl(qs),
      {
        method: Entity2QueryAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity2QueryActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity2QueryActionRes(item));
    const res = await Entity2QueryAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "entity2Query",
    url: "/entity2/query",
    method: "post",
    description:
      'Returns "entity2" rows matching a JSON-logic filter, sorted/paged (see emigorm.QueryDSL).',
    in: {
      fields: [
        {
          name: "filter",
          type: "string?",
        },
        {
          name: "sort",
          type: "string?",
        },
        {
          name: "startIndex",
          type: "int?",
        },
        {
          name: "itemsPerPage",
          type: "int?",
        },
      ],
    },
    out: {
      fields: [
        {
          name: "items",
          type: "array",
          fields: [
            {
              name: "uniqueId",
              type: "string",
            },
            {
              name: "label2",
              type: "string",
            },
          ],
        },
        {
          name: "total",
          type: "int64",
        },
      ],
    },
  };
}
/**
 * The base class definition for entity2QueryActionReq
 **/
export class Entity2QueryActionReq {
  /**
   *
   * @type {string}
   **/
  #filter = undefined;
  /**
   *
   * @returns {string}
   **/
  get filter() {
    return this.#filter;
  }
  /**
   *
   * @type {string}
   **/
  set filter(value) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#filter = correctType ? value : String(value);
  }
  setFilter(value) {
    this.filter = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #sort = undefined;
  /**
   *
   * @returns {string}
   **/
  get sort() {
    return this.#sort;
  }
  /**
   *
   * @type {string}
   **/
  set sort(value) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#sort = correctType ? value : String(value);
  }
  setSort(value) {
    this.sort = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #startIndex = undefined;
  /**
   *
   * @returns {number}
   **/
  get startIndex() {
    return this.#startIndex;
  }
  /**
   *
   * @type {number}
   **/
  set startIndex(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#startIndex = parsedValue;
    }
  }
  setStartIndex(value) {
    this.startIndex = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #itemsPerPage = undefined;
  /**
   *
   * @returns {number}
   **/
  get itemsPerPage() {
    return this.#itemsPerPage;
  }
  /**
   *
   * @type {number}
   **/
  set itemsPerPage(value) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#itemsPerPage = parsedValue;
    }
  }
  setItemsPerPage(value) {
    this.itemsPerPage = value;
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
    if (d.filter !== undefined) {
      this.filter = d.filter;
    }
    if (d.sort !== undefined) {
      this.sort = d.sort;
    }
    if (d.startIndex !== undefined) {
      this.startIndex = d.startIndex;
    }
    if (d.itemsPerPage !== undefined) {
      this.itemsPerPage = d.itemsPerPage;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      filter: this.#filter,
      sort: this.#sort,
      startIndex: this.#startIndex,
      itemsPerPage: this.#itemsPerPage,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      filter: "filter",
      sort: "sort",
      startIndex: "startIndex",
      itemsPerPage: "itemsPerPage",
    };
  }
  /**
   * Creates an instance of Entity2QueryActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new Entity2QueryActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of Entity2QueryActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new Entity2QueryActionReq(partialDtoObject);
  }
  copyWith(partial) {
    return new Entity2QueryActionReq({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new Entity2QueryActionReq(this.toJSON());
  }
}
/**
 * The base class definition for entity2QueryActionRes
 **/
export class Entity2QueryActionRes {
  /**
   *
   * @type {Entity2QueryActionRes.Items}
   **/
  #items = MArray.of([]);
  /**
   *
   * @returns {Entity2QueryActionRes.Items}
   **/
  get items() {
    return this.#items;
  }
  /**
   *
   * @type {Entity2QueryActionRes.Items}
   **/
  set items(value) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof Entity2QueryActionRes.Items) {
        this.#items = MArray.of(value);
      } else {
        this.#items = MArray.of(
          value.map((item) => new Entity2QueryActionRes.Items(item)),
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
   * @type {number}
   **/
  #total = 0;
  /**
   *
   * @returns {number}
   **/
  get total() {
    return this.#total;
  }
  /**
   *
   * @type {number}
   **/
  set total(value) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#total = parsedValue;
    }
  }
  setTotal(value) {
    this.total = value;
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
    #uniqueId = "";
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
      this.#uniqueId = String(value);
    }
    setUniqueId(value) {
      this.uniqueId = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #label2 = "";
    /**
     *
     * @returns {string}
     **/
    get label2() {
      return this.#label2;
    }
    /**
     *
     * @type {string}
     **/
    set label2(value) {
      this.#label2 = String(value);
    }
    setLabel2(value) {
      this.label2 = value;
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
      if (d.uniqueId !== undefined) {
        this.uniqueId = d.uniqueId;
      }
      if (d.label2 !== undefined) {
        this.label2 = d.label2;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        uniqueId: this.#uniqueId,
        label2: this.#label2,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        uniqueId: "uniqueId",
        label2: "label2",
      };
    }
    /**
     * Creates an instance of Entity2QueryActionRes.Items, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject) {
      return new Entity2QueryActionRes.Items(possibleDtoObject);
    }
    /**
     * Creates an instance of Entity2QueryActionRes.Items, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject) {
      return new Entity2QueryActionRes.Items(partialDtoObject);
    }
    copyWith(partial) {
      return new Entity2QueryActionRes.Items({ ...this.toJSON(), ...partial });
    }
    clone() {
      return new Entity2QueryActionRes.Items(this.toJSON());
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
    if (d.items !== undefined) {
      this.items = d.items;
    }
    if (d.total !== undefined) {
      this.total = d.total;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      items: this.#items,
      total: this.#total,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      items$: "items",
      get items() {
        return withPrefix("items[:i]", Entity2QueryActionRes.Items.Fields);
      },
      total: "total",
    };
  }
  /**
   * Creates an instance of Entity2QueryActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new Entity2QueryActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of Entity2QueryActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new Entity2QueryActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new Entity2QueryActionRes({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new Entity2QueryActionRes(this.toJSON());
  }
}
