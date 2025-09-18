import {
  FetchxContext,
  fetchx,
  handleFetchResponse,
  type TypedRequestInit,
} from "./sdk/common/fetchx";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * Action to communicate with the action httpAction
 */
export type HttpActionActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
/**
 * HttpActionAction
 */
export class HttpActionAction {
  static URL = "http://localhost:8081 (for test we use override)";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(HttpActionAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext<unknown, unknown>,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<HttpActionActionRes>, unknown, unknown>(
      overrideUrl ?? HttpActionAction.NewUrl(qs),
      {
        method: HttpActionAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    creatorFn: (item: unknown) => HttpActionActionRes = (item) =>
      new HttpActionActionRes(item),
    qs?: URLSearchParams,
    ctx?: FetchxContext<unknown, unknown>,
    init?: TypedRequestInit<unknown, unknown>,
    onMessage?: (ev: MessageEvent) => void,
    overrideUrl?: string,
  ) => {
    const res = await HttpActionAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        return new GResponse<HttpActionActionRes>()
          .setCreator(creatorFn)
          .inject(data);
      },
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "httpAction",
    url: "http://localhost:8081 (for test we use override)",
    method: "post",
    description:
      "A post request which would return an array, but enveloped in google json styleguide.",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "recordNumber",
          description:
            "Fake record number to simulate a id from database table.",
          type: "int",
        },
      ],
    },
  };
}
/**
 * The base class definition for httpActionActionRes
 **/
export class HttpActionActionRes {
  /**
   * Fake record number to simulate a id from database table.
   * @type {number}
   **/
  #recordNumber: number = 0;
  /**
   * Fake record number to simulate a id from database table.
   * @returns {number}
   **/
  get recordNumber() {
    return this.#recordNumber;
  }
  /**
   * Fake record number to simulate a id from database table.
   * @type {number}
   **/
  set recordNumber(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#recordNumber = parsedValue;
    }
  }
  setRecordNumber(value: number) {
    this.recordNumber = value;
    return this;
  }
  constructor(data: unknown) {
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
    const d = data as Partial<HttpActionActionRes>;
    if (d.recordNumber !== undefined) {
      this.recordNumber = d.recordNumber;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      recordNumber: this.#recordNumber,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      recordNumber: "recordNumber",
    };
  }
}
export abstract class HttpActionActionResFactory {
  abstract create(data: unknown): HttpActionActionRes;
}
/**
 * The base type definition for httpActionActionRes
 **/
export type HttpActionActionResType = {
  /**
   * Fake record number to simulate a id from database table.
   * @type {number}
   **/
  recordNumber: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace HttpActionActionResType {}
