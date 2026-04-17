import { CommonVectorResponseDto } from "./CommonVectorResponseDto";
import { buildUrl } from "./sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "./sdk/common/fetchx";
import { type UseMutationOptions, useMutation } from "react-query";
import { useFetchxContext } from "./sdk/react/useFetchx";
import { useState } from "react";
import { withPrefix } from "./sdk/common/withPrefix";
/**
 * Action to communicate with the action allData
 */
export type AllDataActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type AllDataActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  AllDataActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => AllDataActionRes;
  }>;
export const useAllDataAction = (options?: AllDataActionMutationOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return AllDataAction.Fetch(
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
 * AllDataAction
 */
export class AllDataAction {
  //
  static URL = "/res/22";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(AllDataAction.URL, undefined, qs);
  static Method = "";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<AllDataActionRes, unknown, unknown>(
      overrideUrl ?? AllDataAction.NewUrl(qs),
      {
        method: AllDataAction.Method,
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
      creatorFn?: ((item: unknown) => AllDataActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new AllDataActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new AllDataActionRes(item));
    const res = await AllDataAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "allData",
    url: "/res/22",
    out: {
      fields: [
        {
          name: "stringType",
          type: "string",
        },
        {
          name: "stringTypeNull",
          type: "string?",
        },
        {
          name: "collection",
          type: "collection",
          target: "CommonVectorResponseDto",
        },
      ],
    },
  };
}
/**
 * The base class definition for allDataActionRes
 **/
export class AllDataActionRes {
  /**
   *
   * @type {string}
   **/
  #stringType: string = "";
  /**
   *
   * @returns {string}
   **/
  get stringType() {
    return this.#stringType;
  }
  /**
   *
   * @type {string}
   **/
  set stringType(value: string) {
    this.#stringType = String(value);
  }
  setStringType(value: string) {
    this.stringType = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #stringTypeNull?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get stringTypeNull() {
    return this.#stringTypeNull;
  }
  /**
   *
   * @type {string}
   **/
  set stringTypeNull(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#stringTypeNull = correctType ? value : String(value);
  }
  setStringTypeNull(value: string | null | undefined) {
    this.stringTypeNull = value;
    return this;
  }
  /**
   *
   * @type {CommonVectorResponseDto[]}
   **/
  #collection: CommonVectorResponseDto[] = [];
  /**
   *
   * @returns {CommonVectorResponseDto[]}
   **/
  get collection() {
    return this.#collection;
  }
  /**
   *
   * @type {CommonVectorResponseDto[]}
   **/
  set collection(value: CommonVectorResponseDto[]) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof CommonVectorResponseDto) {
      this.#collection = value;
    } else {
      this.#collection = value.map((item) => new CommonVectorResponseDto(item));
    }
  }
  setCollection(value: CommonVectorResponseDto[]) {
    this.collection = value;
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
    const d = data as Partial<AllDataActionRes>;
    if (d.stringType !== undefined) {
      this.stringType = d.stringType;
    }
    if (d.stringTypeNull !== undefined) {
      this.stringTypeNull = d.stringTypeNull;
    }
    if (d.collection !== undefined) {
      this.collection = d.collection;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      stringType: this.#stringType,
      stringTypeNull: this.#stringTypeNull,
      collection: this.#collection,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      stringType: "stringType",
      stringTypeNull: "stringTypeNull",
      collection$: "collection",
      get collection() {
        return withPrefix("collection[:i]", CommonVectorResponseDto.Fields);
      },
    };
  }
  /**
   * Creates an instance of AllDataActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AllDataActionResType) {
    return new AllDataActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of AllDataActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AllDataActionResType>) {
    return new AllDataActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AllDataActionResType>,
  ): InstanceType<typeof AllDataActionRes> {
    return new AllDataActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AllDataActionRes> {
    return new AllDataActionRes(this.toJSON());
  }
}
export abstract class AllDataActionResFactory {
  abstract create(data: unknown): AllDataActionRes;
}
/**
 * The base type definition for allDataActionRes
 **/
export type AllDataActionResType = {
  /**
   *
   * @type {string}
   **/
  stringType: string;
  /**
   *
   * @type {string}
   **/
  stringTypeNull?: string;
  /**
   *
   * @type {CommonVectorResponseDto[]}
   **/
  collection: CommonVectorResponseDto[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AllDataActionResType {}
