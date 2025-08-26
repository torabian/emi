import { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { ExecutionContext, createParamDecorator } from '@nestjs/common';
import {
  URLSearchParamsX,
  buildUrl,
  fetchx,
  type TypedRequestInit,
} from './sdk';
/**
 * Action to communicate with the action getSinglePost
 */
export type GetSinglePostActionOptions = {
  queryKey?: unknown[];
  params: FetchGetSinglePostActionPathParameter;
  qs?: GetSinglePostQueryParams;
  headers?: GetSinglePostHeaders;
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
  static URL = 'https://jsonplaceholder.typicode.com/posts/:id';
  static NewUrl = (
    params: FetchGetSinglePostActionPathParameter,
    qs?: GetSinglePostQueryParams,
  ) => buildUrl(FetchGetSinglePostAction.URL, params, qs);
  static Method = 'get';
  static Axios: (
    clientInstance: AxiosInstance,
    config: AxiosRequestConfig<unknown>,
  ) => Promise<AxiosResponse<unknown>> = (clientInstance, config) =>
    clientInstance
      .request<unknown, AxiosResponse<unknown>, unknown>(config)
      .then((res) => {
        return {
          ...res,
          // if there is a output class, create instance out of it.
          data: new GetSinglePostRes(res.data),
        };
      });
  static Fetch = async (
    params: FetchGetSinglePostActionPathParameter,
    qs?: GetSinglePostQueryParams,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    const res = await fetchx<unknown, unknown, unknown>(
      overrideUrl ?? FetchGetSinglePostAction.NewUrl(params, qs),
      init,
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
  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@GetSinglePostRes.Nest() headers: GetSinglePostRes): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
    (_data: unknown, ctx: ExecutionContext): GetSinglePostRes => {
      const request = ctx.switchToHttp().getRequest<{
        body: unknown;
        headers: { [s: string]: string };
        query: Record<string, string>;
      }>();
      return new GetSinglePostRes(request.body as Partial<GetSinglePostRes>);
    },
  );
}
export abstract class GetSinglePostResFactory {
  abstract create(data: unknown): GetSinglePostRes;
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
  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@GetSinglePostHeaders.Nest() headers: GetSinglePostHeaders): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
    (_data: unknown, ctx: ExecutionContext): GetSinglePostHeaders => {
      const request = ctx.switchToHttp().getRequest<{
        body: unknown;
        headers: { [s: string]: string };
        query: Record<string, string>;
      }>();
      return new GetSinglePostHeaders(Object.entries(request.headers));
    },
  );
}
/**
 * GetSinglePostQueryParams class
 * Auto-generated from Module3Action
 */
export class GetSinglePostQueryParams extends URLSearchParamsX {
  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@GetSinglePostQueryParams.Nest() headers: GetSinglePostQueryParams): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
    (_data: unknown, ctx: ExecutionContext): GetSinglePostQueryParams => {
      const request = ctx.switchToHttp().getRequest<{
        body: unknown;
        headers: { [s: string]: string };
        query: Record<string, string>;
      }>();
      return new GetSinglePostQueryParams(request.query);
    },
  );
}
