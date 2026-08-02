import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity1AwareDelete
 */
/**
 * Entity1AwareDeleteAction
 */
export class Entity1AwareDeleteAction {
  //
  static URL = "/entity1/delete";
  static NewUrl = (qs) => buildUrl(Entity1AwareDeleteAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity1AwareDeleteAction.NewUrl(qs),
      {
        method: Entity1AwareDeleteAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (init, { qs, ctx, onMessage, overrideUrl } = {}) => {
    const res = await Entity1AwareDeleteAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(res, undefined, onMessage, init?.signal);
  };
  static Definition = {
    name: "entity1AwareDelete",
    url: "/entity1/delete",
    method: "post",
    description:
      'Deletes the given "entity1" uniqueIds, along with everything entity1AwareDeletePreview reports.',
    in: {
      fields: [
        {
          name: "uniqueIds",
          type: "slice",
          primitive: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for entity1AwareDeleteActionReq
 **/
export class Entity1AwareDeleteActionReq {
  /**
   *
   * @type {string[]}
   **/
  #uniqueIds = [];
  /**
   *
   * @returns {string[]}
   **/
  get uniqueIds() {
    return this.#uniqueIds;
  }
  /**
   *
   * @type {string[]}
   **/
  set uniqueIds(value) {
    this.#uniqueIds = value;
  }
  setUniqueIds(value) {
    this.uniqueIds = value;
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
    if (d.uniqueIds !== undefined) {
      this.uniqueIds = d.uniqueIds;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueIds: this.#uniqueIds,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueIds$: "uniqueIds",
      get uniqueIds() {
        return "uniqueIds[:i]";
      },
    };
  }
  /**
   * Creates an instance of Entity1AwareDeleteActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject) {
    return new Entity1AwareDeleteActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of Entity1AwareDeleteActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject) {
    return new Entity1AwareDeleteActionReq(partialDtoObject);
  }
  copyWith(partial) {
    return new Entity1AwareDeleteActionReq({ ...this.toJSON(), ...partial });
  }
  clone() {
    return new Entity1AwareDeleteActionReq(this.toJSON());
  }
}
