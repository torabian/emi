import {
  FetchxContext,
  fetchx,
  handleFetchResponse,
  type TypedRequestInit,
  type TypedResponse,
} from "./sdk/common/fetchx";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "./sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action httpAction
 */
export type HttpActionActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type HttpActionActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, HttpActionActionRes, unknown[]>,
  "queryKey"
> &
  HttpActionActionOptions &
  Partial<{
    creatorFn: (item: unknown) => HttpActionActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useHttpActionActionQuery = (
  options: HttpActionActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return HttpActionAction.Fetch(
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
    queryKey: [HttpActionAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type HttpActionActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  HttpActionActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => HttpActionActionRes;
  }>;
export const useHttpActionAction = (
  options?: HttpActionActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return HttpActionAction.Fetch(
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
 * HttpActionAction
 */
export class HttpActionAction {
  static URL = "http://localhost:8081 (for test we use override)";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(HttpActionAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
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
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => HttpActionActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new HttpActionActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new HttpActionActionRes(item));
    const res = await HttpActionAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<HttpActionActionRes>();
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
    name: "httpAction",
    url: "http://localhost:8081 (for test we use override)",
    method: "get",
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
  /**
   * Creates an instance of HttpActionActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: HttpActionActionResType) {
    return new HttpActionActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of HttpActionActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<HttpActionActionResType>) {
    return new HttpActionActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<HttpActionActionResType>,
  ): InstanceType<typeof HttpActionActionRes> {
    return new HttpActionActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof HttpActionActionRes> {
    return new HttpActionActionRes(this.toJSON());
  }
}
export abstract class HttpActionActionResFactory {
  abstract create(data: unknown): HttpActionActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
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
