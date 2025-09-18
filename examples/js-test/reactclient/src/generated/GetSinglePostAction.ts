import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit, type TypedResponse } from './sdk/common/fetchx';
import { Money } from '../Money';
import { buildUrl } from './sdk/common/buildUrl';
import { type UseMutationOptions, type UseQueryOptions, useMutation, useQuery } from '@tanstack/react-query';
import { useFetchxContext } from './sdk/react/useFetchx';
import { useState } from 'react';
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
	GetSinglePostActionOptions
& Partial<{
	creatorFn: (item: unknown) => GetSinglePostActionRes
}>
& {
	onMessage?: (ev: MessageEvent) => void;
	overrideUrl?: string;
	headers?: Headers;
	ctx?: FetchxContext<unknown, unknown>;
}
export const useGetSinglePostActionQuery = (
	options: GetSinglePostActionQueryOptions
) => {
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;
	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState<TypedResponse<unknown>>();
	const fn = (
	) =>
		{
			setCompleteState(false);
			GetSinglePostAction.Fetch(
					options.params,
				options?.creatorFn,
				options?.qs,
				{
					headers: options?.headers,
				},
				options?.onMessage,
				options?.overrideUrl,
				ctx,
			).then((x) => {
				x.done.then(() => {
					setCompleteState(true);
				});
				setResponse(x.response)
				return x.response.result;
			})
		}
	const result = useQuery({
		queryKey: [
			GetSinglePostAction.NewUrl (
				options.params,
				options?.qs
			)
		],
		queryFn: fn,
		...(options || {}),
	});
	return {
		...result,
		isCompleted,
		response
	}
};
export type GetSinglePostActionMutationOptions = Omit<
	UseMutationOptions<unknown, unknown, unknown, unknown>,
	"mutationFn"
> &
	GetSinglePostActionOptions
& {
	ctx?: FetchxContext<unknown, unknown>;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  }
& Partial<{
	creatorFn: (item: unknown) => GetSinglePostActionRes
}>
export const useGetSinglePostAction = (
	options: GetSinglePostActionMutationOptions
) => {
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;
	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState<TypedResponse<unknown>>();
	const fn = (
			body: unknown,
	) =>
		{
			setCompleteState(false);
			GetSinglePostAction.Fetch(
					options.params,
				options?.creatorFn,
				options?.qs,
				{
						body,
					headers: options?.headers,
				},
				options?.onMessage,
				options?.overrideUrl,
				ctx,
			).then((x) => {
				x.done.then(() => {
					setCompleteState(true);
				});
				setResponse(x.response)
				return x.response.result;
			})
		}
	const result =  useMutation({
		mutationFn: fn,
		...(options || {}),
	});
	return {
		...result,
		isCompleted,
		response
	}
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
		ctx?: FetchxContext<unknown, unknown>,
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
			},
			ctx
		)
	}
	static Fetch = async (
			params: GetSinglePostActionPathParameter,
			creatorFn: (item: unknown) => GetSinglePostActionRes = (item) => new GetSinglePostActionRes(item),
		qs?: URLSearchParams,
		ctx?: FetchxContext<unknown, unknown>,
		init?: TypedRequestInit<unknown, unknown>,
		onMessage?: (ev: MessageEvent) => void,
		overrideUrl?: string,
	) => {
		const res = await GetSinglePostAction.Fetch$(
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
	}
  static Definition = {
  "name": "getSinglePost",
  "cliName": "get-single-post",
  "url": "https://jsonplaceholder.typicode.com/posts/:id",
  "method": "get",
  "description": "Get's an specific post from the endpoint",
  "out": {
    "fields": [
      {
        "name": "userId",
        "type": "int64"
      },
      {
        "name": "id",
        "type": "int64"
      },
      {
        "name": "title",
        "complex": "+Money"
      },
      {
        "name": "body",
        "type": "string"
      }
    ]
  }
}
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
 #title : Money
		/**
  * 
  * @returns {Money}
  **/
get title () { return this.#title }
/**
  * 
  * @type {Money}
  **/
set title (value: Money) {
	 	if (value instanceof Money) {
			this.#title = value
		} else {
		 	this.#title = new Money(value)
		}
}
setTitle (value: Money) {
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
	constructor(data: unknown) {
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
	#isJsonAppliable(obj: unknown) {
		const g = globalThis as any
		const isBuffer =
			typeof g.Buffer !== "undefined" &&
			typeof g.Buffer.isBuffer === "function" &&
			g.Buffer.isBuffer(obj);
		const isBlob =
			typeof g.Blob !== "undefined" && obj instanceof g.Blob;
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
 title : Money;
			/**
  * 
  * @type {string}
  **/
 body : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetSinglePostActionResType {
}