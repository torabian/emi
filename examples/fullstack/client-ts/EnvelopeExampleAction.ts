import {
  FetchxContext,
  fetchx,
  handleFetchResponse,
  type TypedRequestInit,
} from "./sdk/common/fetchx";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
/**
 * Action to communicate with the action envelopeExample
 */
export type EnvelopeExampleActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
/**
 * EnvelopeExampleAction
 */
export class EnvelopeExampleAction {
  //
  static URL = "/response/with/envelop";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(EnvelopeExampleAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<EnvelopeExampleActionRes>, unknown, unknown>(
      overrideUrl ?? EnvelopeExampleAction.NewUrl(qs),
      {
        method: EnvelopeExampleAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => EnvelopeExampleActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new EnvelopeExampleActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new EnvelopeExampleActionRes(item));
    const res = await EnvelopeExampleAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<EnvelopeExampleActionRes>();
        if (creatorFn) {
          resp.setCreator(creatorFn);
        }
        resp.inject(data);
        return resp;
      },
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "envelopeExample",
    url: "/response/with/envelop",
    method: "get",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "content",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for envelopeExampleActionRes
 **/
export class EnvelopeExampleActionRes {
  /**
   *
   * @type {string}
   **/
  #content: string = "";
  /**
   *
   * @returns {string}
   **/
  get content() {
    return this.#content;
  }
  /**
   *
   * @type {string}
   **/
  set content(value: string) {
    this.#content = String(value);
  }
  setContent(value: string) {
    this.content = value;
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
    const d = data as Partial<EnvelopeExampleActionRes>;
    if (d.content !== undefined) {
      this.content = d.content;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      content: this.#content,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      content: "content",
    };
  }
  /**
   * Creates an instance of EnvelopeExampleActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: EnvelopeExampleActionResType) {
    return new EnvelopeExampleActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of EnvelopeExampleActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<EnvelopeExampleActionResType>) {
    return new EnvelopeExampleActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<EnvelopeExampleActionResType>,
  ): InstanceType<typeof EnvelopeExampleActionRes> {
    return new EnvelopeExampleActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof EnvelopeExampleActionRes> {
    return new EnvelopeExampleActionRes(this.toJSON());
  }
}
export abstract class EnvelopeExampleActionResFactory {
  abstract create(data: unknown): EnvelopeExampleActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for envelopeExampleActionRes
 **/
export type EnvelopeExampleActionResType = {
  /**
   *
   * @type {string}
   **/
  content: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace EnvelopeExampleActionResType {}
