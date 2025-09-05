import { URLSearchParamsX, buildUrl, fetchx, isPlausibleObject, type TypedRequestInit, withPrefix } from './sdk/js';
import { type AxiosInstance, type AxiosRequestConfig, type AxiosResponse } from 'axios';
import { type UseQueryOptions, useQuery } from '@tanstack/react-query';
/**
* Action to communicate with the action getSinglePost
*/
export type GetSinglePostActionOptions = {
	queryKey?: unknown[];
	params: GetSinglePostActionPathParameter;
	qs?: GetSinglePostActionQueryParams;
};
export type GetSinglePostActionQueryOptions = Omit<
	UseQueryOptions<
		unknown,
		unknown,
			GetSinglePostActionRes,
		unknown[]
	>,
	"queryKey"
> &
	GetSinglePostActionOptions;
export const useGetSinglePostAction = (
	options: GetSinglePostActionQueryOptions
) => {
	return useQuery({
		queryKey: [
			GetSinglePostAction.NewUrl (
				options.params,
				options.qs
			)
		],
		queryFn: () =>
		GetSinglePostAction.Fetch(
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
 * Path parameters for GetSinglePostAction
 */
export type GetSinglePostActionPathParameter = {
	id: string | number | boolean;
}
	/**
 * GetSinglePostAction
 */
export class GetSinglePostAction {
  static URL = 'https://jsonplaceholder.typicode.com/posts/:id';
  static NewUrl = (
	params: GetSinglePostActionPathParameter,
	qs?: GetSinglePostActionQueryParams
  ) => buildUrl(
		GetSinglePostAction.URL,
		params,
		qs
	);
  static Method = 'get';
	static Fetch = async (
			params: GetSinglePostActionPathParameter,
		qs?: GetSinglePostActionQueryParams,
		init?: TypedRequestInit<GetSinglePostActionRes, unknown>,
		overrideUrl?: string
	) => {
		const res = await fetchx<GetSinglePostActionRes, unknown, unknown>(
			overrideUrl ?? GetSinglePostAction.NewUrl(
				params,
				qs
			),
			{
				method: GetSinglePostAction.Method,
				...(init || {})
			}
		)
			const result = await res.json();
				res.result = new GetSinglePostActionRes (result);
			return res;
	}
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
				method: GetSinglePostAction.Method,
				...(config || {})
			}
		)
		.then((res) => {
			return {
			...res,
			// if there is a output class, create instance out of it.
			data: new GetSinglePostActionRes(res.data),
			};
		});
}
/**
  * The base class definition for getSinglePostActionRes
  **/
export class GetSinglePostActionRes {
		/**
  * 
  * @type {number}
  **/
 #userId : number  =  0
		/**
  * 
  * @returns {number}
  **/
get userId () { return this.#userId }
/**
  * 
  * @type {number}
  **/
set userId (value: number) {
	 	const correctType = typeof value === 'number'
		this.#userId = correctType ? value : Number(value);
}
setUserId (value: number) {
	this.userId = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #id : number  =  0
		/**
  * 
  * @returns {number}
  **/
get id () { return this.#id }
/**
  * 
  * @type {number}
  **/
set id (value: number) {
	 	const correctType = typeof value === 'number'
		this.#id = correctType ? value : Number(value);
}
setId (value: number) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #title : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get title () { return this.#title }
/**
  * 
  * @type {string}
  **/
set title (value: string) {
	 	const correctType = typeof value === 'string';
		this.#title = correctType ? value : ('' + value);
}
setTitle (value: string) {
	this.title = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #body : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get body () { return this.#body }
/**
  * 
  * @type {string}
  **/
set body (value: string) {
	 	const correctType = typeof value === 'string';
		this.#body = correctType ? value : ('' + value);
}
setBody (value: string) {
	this.body = value
	return this
}
	constructor(data) {
		if (data === null || data === undefined) {
			return;
		}
		if (typeof data === "string") {
			this.applyFromObject(JSON.parse(data));
		} else if (isPlausibleObject(data)) {
			this.applyFromObject(data);
		} else {
			throw new Error("Instance is not implemented.");
		}
	}
	/**
	* casts the fields of a javascript object into the class properties one by one
	**/
	applyFromObject(data = {}) {
		const d = data as Partial<GetSinglePostActionRes>;
			if (d.userId !== undefined) { this.userId = d.userId }
			if (d.id !== undefined) { this.id = d.id }
			if (d.title !== undefined) { this.title = d.title }
			if (d.body !== undefined) { this.body = d.body }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				userId: this.#userId,
				id: this.#id,
				title: this.#title,
				body: this.#body,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			userId: 'userId',
			id: 'id',
			title: 'title',
			body: 'body',
	  }
	}
}
export abstract class GetSinglePostActionResFactory {
	abstract create(data: unknown): GetSinglePostActionRes;
}
	/**
  * The base type definition for getSinglePostActionRes
  **/
	export type GetSinglePostActionResType =  {
			/**
  * 
  * @type {number}
  **/
 userId?: number;
			/**
  * 
  * @type {number}
  **/
 id?: number;
			/**
  * 
  * @type {string}
  **/
 title?: string;
			/**
  * 
  * @type {string}
  **/
 body?: string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetSinglePostActionResType {
}