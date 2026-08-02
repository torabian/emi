import { Entity2Dto } from "./Entity2Dto";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity2Create
 */
/**
 * Entity2CreateAction
 */
export class Entity2CreateAction {
  //
  static URL = "/entity2";
  static NewUrl = (qs) => buildUrl(Entity2CreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity2CreateAction.NewUrl(qs),
      {
        method: Entity2CreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity2CreateActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity2CreateActionRes(item));
    const res = await Entity2CreateAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "entity2Create",
    url: "/entity2",
    method: "post",
    description: 'Creates a new "entity2" row.',
    in: {
      dto: "Entity2Dto",
    },
    out: {
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
  };
}
/**
 * The base class definition for entity2CreateActionRes
 **/
export class Entity2CreateActionRes {
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
   * Creates an instance of Entity2CreateActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new Entity2CreateActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of Entity2CreateActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new Entity2CreateActionRes(partialDtoObject);
  }
  copyWith(partial) {
    return new Entity2CreateActionRes({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new Entity2CreateActionRes(this.toJSON());
  }
}
