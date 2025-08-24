import {
  URLSearchParamsX,
  buildUrl,
  fetchx,
  type TypedRequestInit,
} from "./sdk";
import { type UseQueryOptions, useQuery } from "@tanstack/react-query";
/**
 * Action to communicate with the action getSinglePost
 */
export type GetSinglePostActionOptions = {
  queryKey?: unknown[];
  params: FetchGetSinglePostActionPathParameter;
  qs?: GetSinglePostQueryParams;
  headers?: GetSinglePostHeaders;
};
export type GetSinglePostActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, unknown, unknown[]>,
  "queryKey"
> &
  GetSinglePostActionOptions;
// 2️⃣ Hook: uses the query options with useQuery
export const useGetSinglePost = (options: GetSinglePostActionQueryOptions) => {
  return useQuery({
    ...options,
    queryKey: [FetchGetSinglePostAction.NewUrl(options.params, options.qs)],
    queryFn: () =>
      FetchGetSinglePostAction.Fetch(options.params, options.qs, {
        headers: options.headers,
      }),
  });
};
/**
 * Path parameters for FetchGetSinglePostAction
 */
export type FetchGetSinglePostActionPathParameter = {
  id: string | number | boolean;
};
/**
 * FetchGetSinglePostAction
 */
export class FetchGetSinglePostAction {
  static URL = "https://jsonplaceholder.typicode.com/posts/:id";
  static NewUrl = (
    params: FetchGetSinglePostActionPathParameter,
    qs?: GetSinglePostQueryParams
  ) => buildUrl(FetchGetSinglePostAction.URL, params, qs);
  static Method = "get";
  static Fetch = async (
    params: FetchGetSinglePostActionPathParameter,
    qs?: GetSinglePostQueryParams,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string
  ) => {
    const res = await fetchx<unknown, unknown, unknown>(
      overrideUrl ?? FetchGetSinglePostAction.NewUrl(params, qs),
      init
    );
    const result = await res.json();
    res.result = new GetSinglePostRes(result);
    return res;
  };
}
/**
 * @decription The base class definition for getSinglePostRes
 **/
export class GetSinglePostRes {
  constructor(data: unknown) {
    // This probably doesn't cover the nested objects
    Object.assign(this, data);
  }
  /**
   * @type {number}
   * @description
   **/
  userId?: number;
  /**
   * @returns {number}
   * @description
   **/
  getUserId() {
    return this[`userId`];
  }
  /**
   * @param {number}
   * @description
   **/
  setUserId(value: number) {
    this[`userId`] = value;
    return this;
  }
  /**
   * @type {number}
   * @description
   **/
  id?: number;
  /**
   * @returns {number}
   * @description
   **/
  getId() {
    return this[`id`];
  }
  /**
   * @param {number}
   * @description
   **/
  setId(value: number) {
    this[`id`] = value;
    return this;
  }
  /**
   * @type {string}
   * @description
   **/
  title?: string;
  /**
   * @returns {string}
   * @description
   **/
  getTitle() {
    return this[`title`];
  }
  /**
   * @param {string}
   * @description
   **/
  setTitle(value: string) {
    this[`title`] = value;
    return this;
  }
  /**
   * @type {string}
   * @description
   **/
  body?: string;
  /**
   * @returns {string}
   * @description
   **/
  getBody() {
    return this[`body`];
  }
  /**
   * @param {string}
   * @description
   **/
  setBody(value: string) {
    this[`body`] = value;
    return this;
  }
}
/**
 * GetSinglePostHeaders class
 * Auto-generated from Module3Action
 */
export class GetSinglePostHeaders extends Headers {
  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }
}
/**
 * GetSinglePostQueryParams class
 * Auto-generated from Module3Action
 */
export class GetSinglePostQueryParams extends URLSearchParamsX {}
