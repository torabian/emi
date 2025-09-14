import { Money } from '../Money';
import { buildUrl } from './sdk/common/buildUrl';
import { fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { type AxiosInstance, type AxiosRequestConfig, type AxiosResponse } from 'axios';
import { type UseMutationOptions, type UseQueryOptions, useMutation, useQuery } from '@tanstack/react-query';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action getSinglePost
*/
export type GetSinglePostActionOptions = {
	queryKey?: unknown[];
	params: GetSinglePostActionPathParameter;
	qs?: URLSearchParams;
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
export type GetSinglePostActionMutationOptions = Omit<
	UseMutationOptions<unknown, unknown, unknown, unknown>,
	"mutationFn"
> &
	GetSinglePostActionOptions;
export const useGetSinglePostActionMutation = (
	options: GetSinglePostActionMutationOptions
) => {
	return useMutation({
		mutationFn: (vars: unknown) =>
			GetSinglePostAction.Fetch(
				options.params,
				options.qs,
				{
					body: vars,
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
	qs?: URLSearchParams
  ) => buildUrl(
		GetSinglePostAction.URL,
		params,
		qs
	);
  static Method = 'get';
	static Fetch$ = async (
			params: GetSinglePostActionPathParameter,
		qs?: URLSearchParams,
		init?: TypedRequestInit<unknown, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<GetSinglePostActionRes, unknown, unknown>(
			overrideUrl ?? GetSinglePostAction.NewUrl(
				params,
				qs
			),
			{
				method: GetSinglePostAction.Method,
				...(init || {})
			}
		)
	}
	static Fetch = async (
			params: GetSinglePostActionPathParameter,
			creatorFn: (item: unknown) => GetSinglePostActionRes = (item) => new GetSinglePostActionRes(item),
		qs?: URLSearchParams,
		init?: TypedRequestInit<unknown, unknown>,
		onMessage?: (ev: MessageEvent) => void,
		overrideUrl?: string,
	) => {
		const res = await GetSinglePostAction.Fetch$(
			params,
			qs,
			init,
			overrideUrl,
		);
			return handleFetchResponse(
				res, 
				GetSinglePostActionRes,
				onMessage,
				init?.signal,
			);
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
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#userId = parsedValue;
		}
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
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#id = parsedValue;
		}
}
setId (value: number) {
	this.id = value
	return this
}
		/**
  * 
  * @type {Money}
  **/
 #title22 : Money
		/**
  * 
  * @returns {Money}
  **/
get title22 () { return this.#title22 }
/**
  * 
  * @type {Money}
  **/
set title22 (value: Money) {
	 	if (value instanceof Money) {
			this.#title22 = value
		} else {
		 	this.#title22 = new Money(value)
		}
}
setTitle22 (value: Money) {
	this.title22 = value
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
		} else if (this.#isJsonAppliable(data)) {
			this.applyFromObject(data);
		} else {
			throw new Error("Instance cannot be created on an unknown value, check the content being passed. got: "  + typeof data);
		}
	}
	#isJsonAppliable(obj) {
		const isBuffer =
			typeof globalThis.Buffer !== "undefined" &&
			typeof globalThis.Buffer.isBuffer === "function" &&
			globalThis.Buffer.isBuffer(obj);
		const isBlob =
			typeof globalThis.Blob !== "undefined" && obj instanceof globalThis.Blob;
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
		const d = data as Partial<GetSinglePostActionRes>;
			if (d.userId !== undefined) { this.userId = d.userId }
			if (d.id !== undefined) { this.id = d.id }
			if (d.title22 !== undefined) { this.title22 = d.title22 }
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
				title22: this.#title22,
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
			title22: 'title22',
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
 userId : number;
			/**
  * 
  * @type {number}
  **/
 id : number;
			/**
  * 
  * @type {Money}
  **/
 title22 : Money;
			/**
  * 
  * @type {string}
  **/
 body : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetSinglePostActionResType {
}