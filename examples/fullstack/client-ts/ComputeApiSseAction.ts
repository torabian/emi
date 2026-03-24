import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit, type TypedResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { type UseMutationOptions, type UseQueryOptions, useMutation, useQuery } from 'react-query';
import { useFetchxContext } from './sdk/react/useFetchx';
import { useState } from 'react';
/**
* Action to communicate with the action computeApiSse
*/
export type ComputeApiSseActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
export type ComputeApiSseActionQueryOptions = Omit<
	UseQueryOptions<
		unknown,
		unknown,
			ComputeApiSseActionRes,
		unknown[]
	>,
	"queryKey"
> &
	ComputeApiSseActionOptions
& Partial<{
	creatorFn: (item: unknown) => ComputeApiSseActionRes
}>
& {
	onMessage?: (ev: MessageEvent) => void;
	overrideUrl?: string;
	headers?: Headers;
	ctx?: FetchxContext;
}
export const useComputeApiSseActionQuery = (
	options: ComputeApiSseActionQueryOptions
) => {
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;
	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState<TypedResponse<unknown>>();
	const fn = (
	) =>
		{
			setCompleteState(false);
			return ComputeApiSseAction.Fetch(
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
			ComputeApiSseAction.NewUrl (
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
export type ComputeApiSseActionMutationOptions = Omit<
	UseMutationOptions<unknown, unknown, unknown, unknown>,
	"mutationFn"
> &
	ComputeApiSseActionOptions
& {
	ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  }
& Partial<{
	creatorFn: (item: unknown) => ComputeApiSseActionRes
}>
export const useComputeApiSseAction = (
	options?: ComputeApiSseActionMutationOptions
) => {
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;
	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState<TypedResponse<unknown>>();
	const fn = (
			body: ComputeApiSseActionReq
	) =>
		{
			setCompleteState(false);
			return ComputeApiSseAction.Fetch(
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
 * ComputeApiSseAction
 */
export class ComputeApiSseAction { //
  static URL = '/compute/sse';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		ComputeApiSseAction.URL,
		 undefined,
		qs
	);
  static Method = 'get';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<ComputeApiSseActionReq, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<ComputeApiSseActionRes, ComputeApiSseActionReq, unknown>(
			overrideUrl ?? ComputeApiSseAction.NewUrl(
				qs
			),
			{
				method: ComputeApiSseAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init?: TypedRequestInit<ComputeApiSseActionReq, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => ComputeApiSseActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new ComputeApiSseActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new ComputeApiSseActionRes(item))
		const res = await ComputeApiSseAction.Fetch$(
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
  "name": "computeApiSse",
  "url": "/compute/sse",
  "method": "get",
  "description": "The same compute api, but it would return the response as SSE.",
  "in": {
    "fields": [
      {
        "name": "initialVector1",
        "type": "slice",
        "primitive": "int"
      },
      {
        "name": "value",
        "type": "string?"
      },
      {
        "name": "initialVector2",
        "type": "slice",
        "primitive": "int"
      }
    ]
  },
  "out": {
    "fields": [
      {
        "name": "outputVector",
        "type": "slice",
        "primitive": "int"
      }
    ]
  }
}
}
/**
  * The base class definition for computeApiSseActionReq
  **/
export class ComputeApiSseActionReq {
		/**
  * 
  * @type {number[]}
  **/
 #initialVector1 : number[]  =  []
		/**
  * 
  * @returns {number[]}
  **/
get initialVector1 () { return this.#initialVector1 }
/**
  * 
  * @type {number[]}
  **/
set initialVector1 (value: number[]) {
}
setInitialVector1 (value: number[]) {
	this.initialVector1 = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #value ? : string  | null  =  undefined
		/**
  * 
  * @returns {string}
  **/
get value () { return this.#value }
/**
  * 
  * @type {string}
  **/
set value (value: string | null | undefined) {
	 	const correctType = typeof value === 'string' || value === undefined || value === null
		this.#value = correctType ? value : String(value);
}
setValue (value: string | null | undefined) {
	this.value = value
	return this
}
		/**
  * 
  * @type {number[]}
  **/
 #initialVector2 : number[]  =  []
		/**
  * 
  * @returns {number[]}
  **/
get initialVector2 () { return this.#initialVector2 }
/**
  * 
  * @type {number[]}
  **/
set initialVector2 (value: number[]) {
}
setInitialVector2 (value: number[]) {
	this.initialVector2 = value
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
		const d = data as Partial<ComputeApiSseActionReq>;
			if (d.initialVector1 !== undefined) { this.initialVector1 = d.initialVector1 }
			if (d.value !== undefined) { this.value = d.value }
			if (d.initialVector2 !== undefined) { this.initialVector2 = d.initialVector2 }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				initialVector1: this.#initialVector1,
				value: this.#value,
				initialVector2: this.#initialVector2,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			initialVector1$: 'initialVector1',
get initialVector1() {
					return "initialVector1[:i]";
						},
			value: 'value',
			initialVector2$: 'initialVector2',
get initialVector2() {
					return "initialVector2[:i]";
						},
	  }
	}
	/**
	* Creates an instance of ComputeApiSseActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ComputeApiSseActionReqType) {
		return new ComputeApiSseActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeApiSseActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ComputeApiSseActionReqType>) {
		return new ComputeApiSseActionReq(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ComputeApiSseActionReqType>): InstanceType<typeof ComputeApiSseActionReq> {
		return new ComputeApiSseActionReq ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ComputeApiSseActionReq> {
		return new ComputeApiSseActionReq(this.toJSON());
	}
}
export abstract class ComputeApiSseActionReqFactory {
	abstract create(data: unknown): ComputeApiSseActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for computeApiSseActionReq
  **/
	export type ComputeApiSseActionReqType =  {
			/**
  * 
  * @type {number[]}
  **/
 initialVector1 : number[];
			/**
  * 
  * @type {string}
  **/
 value ?: string;
			/**
  * 
  * @type {number[]}
  **/
 initialVector2 : number[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ComputeApiSseActionReqType {
}
/**
  * The base class definition for computeApiSseActionRes
  **/
export class ComputeApiSseActionRes {
		/**
  * 
  * @type {number[]}
  **/
 #outputVector : number[]  =  []
		/**
  * 
  * @returns {number[]}
  **/
get outputVector () { return this.#outputVector }
/**
  * 
  * @type {number[]}
  **/
set outputVector (value: number[]) {
}
setOutputVector (value: number[]) {
	this.outputVector = value
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
		const d = data as Partial<ComputeApiSseActionRes>;
			if (d.outputVector !== undefined) { this.outputVector = d.outputVector }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				outputVector: this.#outputVector,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			outputVector$: 'outputVector',
get outputVector() {
					return "outputVector[:i]";
						},
	  }
	}
	/**
	* Creates an instance of ComputeApiSseActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ComputeApiSseActionResType) {
		return new ComputeApiSseActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeApiSseActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ComputeApiSseActionResType>) {
		return new ComputeApiSseActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ComputeApiSseActionResType>): InstanceType<typeof ComputeApiSseActionRes> {
		return new ComputeApiSseActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ComputeApiSseActionRes> {
		return new ComputeApiSseActionRes(this.toJSON());
	}
}
export abstract class ComputeApiSseActionResFactory {
	abstract create(data: unknown): ComputeApiSseActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for computeApiSseActionRes
  **/
	export type ComputeApiSseActionResType =  {
			/**
  * 
  * @type {number[]}
  **/
 outputVector : number[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ComputeApiSseActionResType {
}