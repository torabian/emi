import { buildUrl } from "./sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "./sdk/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "react-query";
import { useFetchxContext } from "./sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action getQueryParams
 */
export type GetQueryParamsActionOptions = {
  queryKey?: unknown[];
  params: GetQueryParamsActionPathParameter;
  qs?: URLSearchParams;
};
export type GetQueryParamsActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GetQueryParamsActionRes, unknown[]>,
  "queryKey"
> &
  GetQueryParamsActionOptions &
  Partial<{
    creatorFn: (item: unknown) => GetQueryParamsActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useGetQueryParamsActionQuery = (
  options: GetQueryParamsActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return GetQueryParamsAction.Fetch(
      options.params,
      {
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
  const result = useQuery({
    queryKey: [GetQueryParamsAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type GetQueryParamsActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  GetQueryParamsActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => GetQueryParamsActionRes;
  }>;
export const useGetQueryParamsAction = (
  options: GetQueryParamsActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return GetQueryParamsAction.Fetch(
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
 * Path parameters for GetQueryParamsAction
 */
export type GetQueryParamsActionPathParameter = {
  addres1: string;
  addressName2: string;
  count: number;
};
/**
 * GetQueryParamsAction
 */
export class GetQueryParamsAction {
  //
  static URL = "/stream/:addres1/:addressName2 string/:count int";
  static NewUrl = (
    params: GetQueryParamsActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(GetQueryParamsAction.URL, params, qs);
  static Method = "get";
  static Fetch$ = async (
    params: GetQueryParamsActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GetQueryParamsActionRes, unknown, unknown>(
      overrideUrl ?? GetQueryParamsAction.NewUrl(params, qs),
      {
        method: GetQueryParamsAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: GetQueryParamsActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => GetQueryParamsActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new GetQueryParamsActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new GetQueryParamsActionRes(item));
    const res = await GetQueryParamsAction.Fetch$(
      params,
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
    name: "getQueryParams",
    url: "/stream/:addres1/:addressName2 string/:count int",
    method: "get",
    description:
      "The goal is to check that if query params also become available in the cli params",
    out: {
      fields: [
        {
          name: "name",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for getQueryParamsActionRes
 **/
export class GetQueryParamsActionRes {
  /**
   *
   * @type {string}
   **/
  #name: string = "";
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
  set name(value: string) {
    this.#name = String(value);
  }
  setName(value: string) {
    this.name = value;
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
    const d = data as Partial<GetQueryParamsActionRes>;
    if (d.name !== undefined) {
      this.name = d.name;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      name: this.#name,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      name: "name",
    };
  }
  /**
   * Creates an instance of GetQueryParamsActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: GetQueryParamsActionResType) {
    return new GetQueryParamsActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of GetQueryParamsActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<GetQueryParamsActionResType>) {
    return new GetQueryParamsActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<GetQueryParamsActionResType>,
  ): InstanceType<typeof GetQueryParamsActionRes> {
    return new GetQueryParamsActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof GetQueryParamsActionRes> {
    return new GetQueryParamsActionRes(this.toJSON());
  }
}
export abstract class GetQueryParamsActionResFactory {
  abstract create(data: unknown): GetQueryParamsActionRes;
}
/**
 * The base type definition for getQueryParamsActionRes
 **/
export type GetQueryParamsActionResType = {
  /**
   *
   * @type {string}
   **/
  name: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetQueryParamsActionResType {}
