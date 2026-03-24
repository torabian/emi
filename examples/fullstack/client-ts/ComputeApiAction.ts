import { FetchxContext, TypedRequestInit, TypedResponse, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { URLSearchParamsX } from './sdk/common/URLSearchParamsX';
import { UseMutationOptions, useMutation } from 'react-query';
import { buildUrl } from './sdk/common/buildUrl';
import { useFetchxContext } from './sdk/react/useFetchx';
import { useState } from 'react';
/**
* Action to communicate with the action computeApi
*/
export type ComputeApiActionOptions = {
	queryKey?: unknown[];
	qs?: ComputeApiActionQueryParams;
};
export type ComputeApiActionMutationOptions = Omit<
	UseMutationOptions<unknown, unknown, unknown, unknown>,
	"mutationFn"
> &
	ComputeApiActionOptions
& {
	ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  }
& Partial<{
	creatorFn: (item: unknown) => ComputeApiActionRes
}>
export const useComputeApiAction = (
	options?: ComputeApiActionMutationOptions
) => {
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;
	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState<TypedResponse<unknown>>();
	const fn = (
			body: ComputeApiActionReq
	) =>
		{
			setCompleteState(false);
			return ComputeApiAction.Fetch(
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
 * ComputeApiAction
 */
export class ComputeApiAction { //
  static URL = '/compute/api';
  static NewUrl = (
	qs?: ComputeApiActionQueryParams
  ) => buildUrl(
		ComputeApiAction.URL,
		 undefined,
		qs
	);
  static Method = 'post';
	static Fetch$ = async (
		qs?: ComputeApiActionQueryParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<ComputeApiActionReq, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<ComputeApiActionRes, ComputeApiActionReq, unknown>(
			overrideUrl ?? ComputeApiAction.NewUrl(
				qs
			),
			{
				method: ComputeApiAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init?: TypedRequestInit<ComputeApiActionReq, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => ComputeApiActionRes) | undefined,
			qs?: ComputeApiActionQueryParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new ComputeApiActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new ComputeApiActionRes(item))
		const res = await ComputeApiAction.Fetch$(
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
  "name": "computeApi",
  "url": "/compute/api",
  "method": "post",
  "qs": [
    {
      "name": "action",
      "type": "string"
    },
    {
      "name": "snapPoints",
      "type": "slice",
      "primitive": "int"
    }
  ],
  "description": "An API, which computes two vectors",
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
  * The base class definition for computeApiActionReq
  **/
export class ComputeApiActionReq {
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
		const d = data as Partial<ComputeApiActionReq>;
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
	* Creates an instance of ComputeApiActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ComputeApiActionReqType) {
		return new ComputeApiActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeApiActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ComputeApiActionReqType>) {
		return new ComputeApiActionReq(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ComputeApiActionReqType>): InstanceType<typeof ComputeApiActionReq> {
		return new ComputeApiActionReq ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ComputeApiActionReq> {
		return new ComputeApiActionReq(this.toJSON());
	}
}
export abstract class ComputeApiActionReqFactory {
	abstract create(data: unknown): ComputeApiActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for computeApiActionReq
  **/
	export type ComputeApiActionReqType =  {
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
export namespace ComputeApiActionReqType {
}
/**
  * The base class definition for computeApiActionRes
  **/
export class ComputeApiActionRes {
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
		const d = data as Partial<ComputeApiActionRes>;
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
	* Creates an instance of ComputeApiActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ComputeApiActionResType) {
		return new ComputeApiActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeApiActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ComputeApiActionResType>) {
		return new ComputeApiActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ComputeApiActionResType>): InstanceType<typeof ComputeApiActionRes> {
		return new ComputeApiActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ComputeApiActionRes> {
		return new ComputeApiActionRes(this.toJSON());
	}
}
export abstract class ComputeApiActionResFactory {
	abstract create(data: unknown): ComputeApiActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for computeApiActionRes
  **/
	export type ComputeApiActionResType =  {
			/**
  * 
  * @type {number[]}
  **/
 outputVector : number[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ComputeApiActionResType {
}
/**
 * ComputeApiActionQueryParams class
 * Auto-generated from EmiAction
 */
export class ComputeApiActionQueryParams extends URLSearchParamsX {
  /**
   * 
   * @returns { string | null }
   */
  getAction () {
    return this.getTyped('action' , 'string | null');
  }
  /**
   * 
   * @param { string | null } value
   */
  setAction (value: string | null) {
    this.set('action', value);
    return this;
  }
  /**
   * 
   * @returns { any }
   */
  getSnapPoints () {
    return this.getTyped('snapPoints' , 'any');
  }
  /**
   * 
   * @param { any } value
   */
  setSnapPoints (value: any) {
    this.set('snapPoints', value);
    return this;
  }
}