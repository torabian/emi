import {
  FetchxContext,
  fetchx,
  handleFetchResponse,
  type TypedRequestInit,
} from "./sdk/common/fetchx";
import { buildUrl } from "./sdk/common/buildUrl";
/**
 * Action to communicate with the action computeApiSseChannel
 */
export type ComputeApiSseChannelActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
/**
 * ComputeApiSseChannelAction
 */
export class ComputeApiSseChannelAction {
  //
  static URL = "/compute/sse/ch";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(ComputeApiSseChannelAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<ComputeApiSseChannelActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      ComputeApiSseChannelActionRes,
      ComputeApiSseChannelActionReq,
      unknown
    >(
      overrideUrl ?? ComputeApiSseChannelAction.NewUrl(qs),
      {
        method: ComputeApiSseChannelAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<ComputeApiSseChannelActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => ComputeApiSseChannelActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new ComputeApiSseChannelActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new ComputeApiSseChannelActionRes(item));
    const res = await ComputeApiSseChannelAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (item) => creatorFn(item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "computeApiSseChannel",
    url: "/compute/sse/ch",
    method: "get",
    description:
      "We use a helper in order to send SSE, instead of pure code in gin.",
    in: {
      fields: [
        {
          name: "initialVector1",
          type: "slice",
          primitive: "int",
        },
        {
          name: "value",
          type: "string?",
        },
        {
          name: "initialVector2",
          type: "slice",
          primitive: "int",
        },
      ],
    },
    out: {
      fields: [
        {
          name: "outputVector",
          type: "slice",
          primitive: "int",
        },
      ],
    },
  };
}
/**
 * The base class definition for computeApiSseChannelActionReq
 **/
export class ComputeApiSseChannelActionReq {
  /**
   *
   * @type {number[]}
   **/
  #initialVector1: number[] = [];
  /**
   *
   * @returns {number[]}
   **/
  get initialVector1() {
    return this.#initialVector1;
  }
  /**
   *
   * @type {number[]}
   **/
  set initialVector1(value: number[]) {}
  setInitialVector1(value: number[]) {
    this.initialVector1 = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #value?: string | null = undefined;
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
  set value(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#value = correctType ? value : String(value);
  }
  setValue(value: string | null | undefined) {
    this.value = value;
    return this;
  }
  /**
   *
   * @type {number[]}
   **/
  #initialVector2: number[] = [];
  /**
   *
   * @returns {number[]}
   **/
  get initialVector2() {
    return this.#initialVector2;
  }
  /**
   *
   * @type {number[]}
   **/
  set initialVector2(value: number[]) {}
  setInitialVector2(value: number[]) {
    this.initialVector2 = value;
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
    const d = data as Partial<ComputeApiSseChannelActionReq>;
    if (d.initialVector1 !== undefined) {
      this.initialVector1 = d.initialVector1;
    }
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.initialVector2 !== undefined) {
      this.initialVector2 = d.initialVector2;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      initialVector1: this.#initialVector1,
      value: this.#value,
      initialVector2: this.#initialVector2,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      initialVector1$: "initialVector1",
      get initialVector1() {
        return "initialVector1[:i]";
      },
      value: "value",
      initialVector2$: "initialVector2",
      get initialVector2() {
        return "initialVector2[:i]";
      },
    };
  }
  /**
   * Creates an instance of ComputeApiSseChannelActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ComputeApiSseChannelActionReqType) {
    return new ComputeApiSseChannelActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of ComputeApiSseChannelActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<ComputeApiSseChannelActionReqType>,
  ) {
    return new ComputeApiSseChannelActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ComputeApiSseChannelActionReqType>,
  ): InstanceType<typeof ComputeApiSseChannelActionReq> {
    return new ComputeApiSseChannelActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ComputeApiSseChannelActionReq> {
    return new ComputeApiSseChannelActionReq(this.toJSON());
  }
}
export abstract class ComputeApiSseChannelActionReqFactory {
  abstract create(data: unknown): ComputeApiSseChannelActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for computeApiSseChannelActionReq
 **/
export type ComputeApiSseChannelActionReqType = {
  /**
   *
   * @type {number[]}
   **/
  initialVector1: number[];
  /**
   *
   * @type {string}
   **/
  value?: string;
  /**
   *
   * @type {number[]}
   **/
  initialVector2: number[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ComputeApiSseChannelActionReqType {}
/**
 * The base class definition for computeApiSseChannelActionRes
 **/
export class ComputeApiSseChannelActionRes {
  /**
   *
   * @type {number[]}
   **/
  #outputVector: number[] = [];
  /**
   *
   * @returns {number[]}
   **/
  get outputVector() {
    return this.#outputVector;
  }
  /**
   *
   * @type {number[]}
   **/
  set outputVector(value: number[]) {}
  setOutputVector(value: number[]) {
    this.outputVector = value;
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
    const d = data as Partial<ComputeApiSseChannelActionRes>;
    if (d.outputVector !== undefined) {
      this.outputVector = d.outputVector;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      outputVector: this.#outputVector,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      outputVector$: "outputVector",
      get outputVector() {
        return "outputVector[:i]";
      },
    };
  }
  /**
   * Creates an instance of ComputeApiSseChannelActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ComputeApiSseChannelActionResType) {
    return new ComputeApiSseChannelActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ComputeApiSseChannelActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<ComputeApiSseChannelActionResType>,
  ) {
    return new ComputeApiSseChannelActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ComputeApiSseChannelActionResType>,
  ): InstanceType<typeof ComputeApiSseChannelActionRes> {
    return new ComputeApiSseChannelActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ComputeApiSseChannelActionRes> {
    return new ComputeApiSseChannelActionRes(this.toJSON());
  }
}
export abstract class ComputeApiSseChannelActionResFactory {
  abstract create(data: unknown): ComputeApiSseChannelActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for computeApiSseChannelActionRes
 **/
export type ComputeApiSseChannelActionResType = {
  /**
   *
   * @type {number[]}
   **/
  outputVector: number[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ComputeApiSseChannelActionResType {}
