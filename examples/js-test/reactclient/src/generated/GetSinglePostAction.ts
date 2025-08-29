import { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { URLSearchParamsX, buildUrl, fetchx, type TypedRequestInit } from './sdk';
import { type UseQueryOptions, useQuery } from '@tanstack/react-query';
/**
* Action to communicate with the action getSinglePost
*/
export type GetSinglePostActionOptions = {
	queryKey?: unknown[];
	params: FetchGetSinglePostActionPathParameter;
	qs?: GetSinglePostQueryParams;
	headers?: GetSinglePostReqHeaders;
};
export type GetSinglePostActionQueryOptions = Omit<
	UseQueryOptions<
		unknown,
		unknown,
			GetSinglePostRes,
		unknown[]
	>,
	"queryKey"
> &
	GetSinglePostActionOptions;
export const useGetSinglePost = (
	options: GetSinglePostActionQueryOptions
) => {
	return useQuery({
		queryKey: [
			FetchGetSinglePostAction.NewUrl (
				options.params,
				options.qs
			)
		],
		queryFn: () =>
		FetchGetSinglePostAction.Fetch(
				options.params,
			options.qs,
			{
				headers: options.headers,
			}
		),
		...(options || {}),
	});
};
	/**
 * Path parameters for FetchGetSinglePostAction
 */
export type FetchGetSinglePostActionPathParameter = {
	id: string | number | boolean;
}
	/**
 * FetchGetSinglePostAction
 */
export class FetchGetSinglePostAction {
  static URL = 'https://jsonplaceholder.typicode.com/posts/:id';
  static NewUrl = (
	params: FetchGetSinglePostActionPathParameter,
	qs?: GetSinglePostQueryParams
  ) => buildUrl(
		FetchGetSinglePostAction.URL,
		params,
		qs
	);
  static Method = 'get';
	static Axios : (
		clientInstance: AxiosInstance,
		config: AxiosRequestConfig<unknown>,
	)  => Promise<AxiosResponse<unknown>> = (
		clientInstance,
		config,
	) => 
		clientInstance
		.request<unknown, AxiosResponse<unknown>, unknown>(
			{
				method: FetchGetSinglePostAction.Method,
				...(config || {})
			}
		)
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
		init?: TypedRequestInit<GetSinglePostRes, GetSinglePostReqHeaders>,
		overrideUrl?: string
	) => {
		const res = await fetchx<GetSinglePostRes, unknown, GetSinglePostReqHeaders>(
			overrideUrl ?? FetchGetSinglePostAction.NewUrl(
				params,
				qs
			),
			{
				method: FetchGetSinglePostAction.Method,
				...(init || {})
			}
		)
		const result = await res.json();
			res.result = new GetSinglePostRes (result);
		return res;
	}
}
	/**
  * @description The base type definition for getSinglePostRes
  **/
	export type GetSinglePostResType =  {
			/**
  * @type {number}
  * @description 
  **/
 userId?: number;
			/**
  * @type {number}
  * @description 
  **/
 id?: number;
			/**
  * @type {string}
  * @description 
  **/
 title?: string;
			/**
  * @type {string}
  * @description 
  **/
 body?: string;
	}
export namespace GetSinglePostResType {
}
/**
  * @decription The base class definition for getSinglePostRes
  **/
export class GetSinglePostRes implements GetSinglePostResType {
	constructor(data: unknown) {
		// This probably doesn't cover the nested objects
		const d = data as Partial<GetSinglePostRes>;
			if (d[`userId`] !== undefined) { 
 this.setUserId (d[`userId`]) 
}
			if (d[`id`] !== undefined) { 
 this.setId (d[`id`]) 
}
			if (d[`title`] !== undefined) { 
 this.setTitle (d[`title`]) 
}
			if (d[`body`] !== undefined) { 
 this.setBody (d[`body`]) 
}
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
getUserId () { return this[`userId`] }
		/**
  * @param {number}
  * @description 
  **/
setUserId (value: number) { this[`userId`] = value; return this; } 
		/**
  * @type {number}
  * @description 
  **/
 id?: number;
		/**
  * @returns {number}
  * @description 
  **/
getId () { return this[`id`] }
		/**
  * @param {number}
  * @description 
  **/
setId (value: number) { this[`id`] = value; return this; } 
		/**
  * @type {string}
  * @description 
  **/
 title?: string;
		/**
  * @returns {string}
  * @description 
  **/
getTitle () { return this[`title`] }
		/**
  * @param {string}
  * @description 
  **/
setTitle (value: string) { this[`title`] = value; return this; } 
		/**
  * @type {string}
  * @description 
  **/
 body?: string;
		/**
  * @returns {string}
  * @description 
  **/
getBody () { return this[`body`] }
		/**
  * @param {string}
  * @description 
  **/
setBody (value: string) { this[`body`] = value; return this; } 
}
export abstract class GetSinglePostResFactory {
	abstract create(data: unknown): GetSinglePostRes;
}
/**
 * GetSinglePostReqHeaders class
 * Auto-generated from Module3Action
 */
export class GetSinglePostReqHeaders extends Headers {
  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }
}
/**
 * GetSinglePostResHeaders class
 * Auto-generated from Module3Action
 */
export class GetSinglePostResHeaders extends Headers {
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
export class GetSinglePostQueryParams extends URLSearchParamsX {
}