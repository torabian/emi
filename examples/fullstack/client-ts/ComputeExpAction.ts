import {
  FetchxContext,
  fetchx,
  handleFetchResponse,
  type TypedRequestInit,
  type TypedResponse,
} from "./sdk/common/fetchx";
import { buildUrl } from "./sdk/common/buildUrl";
import { type UseMutationOptions, useMutation } from "react-query";
import { useFetchxContext } from "./sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action computeExp
 */
export type ComputeExpActionOptions = {
  queryKey?: unknown[];
  params: ComputeExpActionPathParameter;
  qs?: URLSearchParams;
};
export type ComputeExpActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  ComputeExpActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => ComputeExpActionRes;
  }>;
export const useComputeExpAction = (
  options: ComputeExpActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: ComputeExpActionReq) => {
    setCompleteState(false);
    return ComputeExpAction.Fetch(
      options.params,
      {
        body,
        headers: options?.headers,
      },
      {
        creatorFn: options?.creatorFn,
        qs: options?.qs,
        ctx,
        onMessage: options?.onMessage,
        overrideUrl: options?.overrideUrl,
      },
    ).then((x) => {
      x.done.then(() => {
        setCompleteState(true);
      });
      setResponse(x.response);
      return x.response.result;
    });
  };
  const result = useMutation({
    mutationFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
/**
 * Path parameters for ComputeExpAction
 */
export type ComputeExpActionPathParameter = {
  first: string;
  second: string;
};
/**
 * ComputeExpAction
 */
export class ComputeExpAction {
  //
  static URL = "/big/exp/:first/:second";
  static NewUrl = (
    params: ComputeExpActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(ComputeExpAction.URL, params, qs);
  static Method = "";
  static Fetch$ = async (
    params: ComputeExpActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<ComputeExpActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<ComputeExpActionRes, ComputeExpActionReq, unknown>(
      overrideUrl ?? ComputeExpAction.NewUrl(params, qs),
      {
        method: ComputeExpAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: ComputeExpActionPathParameter,
    init?: TypedRequestInit<ComputeExpActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => ComputeExpActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new ComputeExpActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new ComputeExpActionRes(item));
    const res = await ComputeExpAction.Fetch$(
      params,
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
    name: "computeExp",
    url: "/big/exp/:first/:second",
    description: "Computes the exp value using big integer",
    in: {
      fields: [
        {
          name: "base",
          type: "complex",
          complex: "Int",
        },
        {
          name: "exponent",
          type: "complex",
          complex: "Int",
        },
      ],
    },
    out: {
      fields: [
        {
          name: "result",
          type: "complex",
          complex: "big.Int",
        },
      ],
    },
  };
}
/**
 * The base class definition for computeExpActionReq
 **/
export class ComputeExpActionReq {
  /**
   *
   * @type {Int}
   **/
  #base!: Int;
  /**
   *
   * @returns {Int}
   **/
  get base() {
    return this.#base;
  }
  /**
   *
   * @type {Int}
   **/
  set base(value: Int) {}
  setBase(value: Int) {
    this.base = value;
    return this;
  }
  /**
   *
   * @type {Int}
   **/
  #exponent!: Int;
  /**
   *
   * @returns {Int}
   **/
  get exponent() {
    return this.#exponent;
  }
  /**
   *
   * @type {Int}
   **/
  set exponent(value: Int) {}
  setExponent(value: Int) {
    this.exponent = value;
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
    const d = data as Partial<ComputeExpActionReq>;
    if (d.base !== undefined) {
      this.base = d.base;
    }
    if (d.exponent !== undefined) {
      this.exponent = d.exponent;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      base: this.#base,
      exponent: this.#exponent,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      base: "base",
      exponent: "exponent",
    };
  }
  /**
   * Creates an instance of ComputeExpActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ComputeExpActionReqType) {
    return new ComputeExpActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of ComputeExpActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ComputeExpActionReqType>) {
    return new ComputeExpActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ComputeExpActionReqType>,
  ): InstanceType<typeof ComputeExpActionReq> {
    return new ComputeExpActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ComputeExpActionReq> {
    return new ComputeExpActionReq(this.toJSON());
  }
}
export abstract class ComputeExpActionReqFactory {
  abstract create(data: unknown): ComputeExpActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for computeExpActionReq
 **/
export type ComputeExpActionReqType = {
  /**
   *
   * @type {Int}
   **/
  base: Int;
  /**
   *
   * @type {Int}
   **/
  exponent: Int;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ComputeExpActionReqType {}
/**
 * The base class definition for computeExpActionRes
 **/
export class ComputeExpActionRes {
  /**
   *
   * @type {big.Int}
   **/
  #result!: big.Int;
  /**
   *
   * @returns {big.Int}
   **/
  get result() {
    return this.#result;
  }
  /**
   *
   * @type {big.Int}
   **/
  set result(value: big.Int) {}
  setResult(value: big.Int) {
    this.result = value;
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
    const d = data as Partial<ComputeExpActionRes>;
    if (d.result !== undefined) {
      this.result = d.result;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      result: this.#result,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      result: "result",
    };
  }
  /**
   * Creates an instance of ComputeExpActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ComputeExpActionResType) {
    return new ComputeExpActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ComputeExpActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ComputeExpActionResType>) {
    return new ComputeExpActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ComputeExpActionResType>,
  ): InstanceType<typeof ComputeExpActionRes> {
    return new ComputeExpActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ComputeExpActionRes> {
    return new ComputeExpActionRes(this.toJSON());
  }
}
export abstract class ComputeExpActionResFactory {
  abstract create(data: unknown): ComputeExpActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for computeExpActionRes
 **/
export type ComputeExpActionResType = {
  /**
   *
   * @type {big.Int}
   **/
  result: big.Int;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ComputeExpActionResType {}
