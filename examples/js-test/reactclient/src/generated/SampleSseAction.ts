import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit, type TypedResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { type UseMutationOptions, type UseQueryOptions, useMutation, useQuery } from '@tanstack/react-query';
import { useFetchxContext } from './sdk/react/useFetchx';
import { useState } from 'react';
/**
* Action to communicate with the action sampleSse
*/
export type SampleSseActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
export type SampleSseActionQueryOptions = Omit<
	UseQueryOptions<
		unknown,
		unknown,
			SampleSseActionRes,
		unknown[]
	>,
	"queryKey"
> &
	SampleSseActionOptions
& Partial<{
	creatorFn: (item: unknown) => SampleSseActionRes
}>
& {
	onMessage?: (ev: MessageEvent) => void;
	overrideUrl?: string;
	headers?: Headers;
	ctx?: FetchxContext;
}
export const useSampleSseActionQuery = (
	options: SampleSseActionQueryOptions
) => {
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;
	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState<TypedResponse<unknown>>();
	const fn = (
	) =>
		{
			setCompleteState(false);
			return SampleSseAction.Fetch(
				{
					headers: options?.headers,
				},
				{
					creatorFn: options?.creatorFn,
					qs: options?.qs,
					ctx,
					onMessage: options?.onMessage,
					overrideUrl: options?.overrideUrl,
				}
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
			SampleSseAction.NewUrl (
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
export type SampleSseActionMutationOptions = Omit<
	UseMutationOptions<unknown, unknown, unknown, unknown>,
	"mutationFn"
> &
	SampleSseActionOptions
& {
	ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  }
& Partial<{
	creatorFn: (item: unknown) => SampleSseActionRes
}>
export const useSampleSseAction = (
	options?: SampleSseActionMutationOptions
) => {
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;
	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState<TypedResponse<unknown>>();
	const fn = (
			body: unknown
	) =>
		{
			setCompleteState(false);
			return SampleSseAction.Fetch(
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
				}
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
 * SampleSseAction
 */
export class SampleSseAction { //
  static URL = 'http://localhost:3000/stream';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		SampleSseAction.URL,
		 undefined,
		qs
	);
  static Method = 'get';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<unknown, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<SampleSseActionRes, unknown, unknown>(
			overrideUrl ?? SampleSseAction.NewUrl(
				qs
			),
			{
				method: SampleSseAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init?: TypedRequestInit<unknown, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => SampleSseActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new SampleSseActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new SampleSseActionRes(item))
		const res = await SampleSseAction.Fetch$(
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
  "name": "sampleSse",
  "url": "http://localhost:3000/stream",
  "method": "get",
  "description": "SSE Sample",
  "out": {
    "fields": [
      {
        "name": "message",
        "type": "string"
      }
    ]
  }
}
}
/**
  * The base class definition for sampleSseActionRes
  **/
export class SampleSseActionRes {
		/**
  * 
  * @type {string}
  **/
 #message : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get message () { return this.#message }
/**
  * 
  * @type {string}
  **/
set message (value: string) {
		this.#message = String(value);
}
setMessage (value: string) {
	this.message = value
	return this
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
			throw new Error("Instance cannot be created on an unknown value, check the content being passed. got: "  + typeof data);
		}
	}
	#isJsonAppliable(obj: unknown) {
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<SampleSseActionRes>;
			if (d.message !== undefined) { this.message = d.message }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				message: this.#message,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			message: 'message',
	  }
	}
	/**
	* Creates an instance of SampleSseActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: SampleSseActionResType) {
		return new SampleSseActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of SampleSseActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<SampleSseActionResType>) {
		return new SampleSseActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<SampleSseActionResType>): InstanceType<typeof SampleSseActionRes> {
		return new SampleSseActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof SampleSseActionRes> {
		return new SampleSseActionRes(this.toJSON());
	}
}
export abstract class SampleSseActionResFactory {
	abstract create(data: unknown): SampleSseActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for sampleSseActionRes
  **/
	export type SampleSseActionResType =  {
			/**
  * 
  * @type {string}
  **/
 message : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SampleSseActionResType {
}