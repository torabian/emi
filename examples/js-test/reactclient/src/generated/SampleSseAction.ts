import { buildUrl } from './sdk/common/buildUrl';
import { fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { type AxiosInstance, type AxiosRequestConfig, type AxiosResponse } from 'axios';
import { type UseMutationOptions, useMutation } from '@tanstack/react-query';
import { useSse } from './sdk/react/useSse';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action sampleSse
*/
export type SampleSseActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
export type SampleSseActionMutationOptions = Omit<
	UseMutationOptions<unknown, unknown, unknown, unknown>,
	"mutationFn"
> &
	SampleSseActionOptions;
export const useSampleSseActionMutation = (
	options: SampleSseActionMutationOptions
) => {
	return useMutation({
		mutationFn: (vars: unknown) =>
			SampleSseAction.Fetch(
				options.qs,
				{
					body: vars,
					headers: options.headers,
				}
			),
		...(options || {}),
	});
};
export const useSampleSseAction = (options: {
	qs?: URLSearchParams,
	init?: TypedRequestInit<unknown, unknown>,
	overrideUrl?: string
}) => {
	return useSse(SampleSseAction.Fetch, options);
};
	/**
 * SampleSseAction
 */
export class SampleSseAction {
  static URL = 'http://localhost:3000/stream';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		SampleSseAction.URL,
		 undefined,
		qs
	);
  static Method = 'sse';
	static Fetch$ = async (
		qs?: URLSearchParams,
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
			}
		)
	}
	static Fetch = async (
			creatorFn: (item: unknown) => SampleSseActionRes = (item) => new SampleSseActionRes(item),
		qs?: URLSearchParams,
		init?: TypedRequestInit<unknown, unknown>,
		onMessage?: (ev: MessageEvent) => void,
		overrideUrl?: string,
	) => {
		const res = await SampleSseAction.Fetch$(
			qs,
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
				method: SampleSseAction.Method,
				...(config || {})
			}
		)
		.then((res) => {
			return {
			...res,
			// if there is a output class, create instance out of it.
			data: new SampleSseActionRes(res.data),
			};
		});
  static Definition = {
  "name": "sampleSse",
  "url": "http://localhost:3000/stream",
  "method": "sse",
  "description": "SSE Sample",
  "out": {
    "fields": [
      {
        "name": "message",
        "type": "string",
        "gormMap": {}
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
	 	const correctType = typeof value === 'string';
		this.#message = correctType ? value : ('' + value);
}
setMessage (value: string) {
	this.message = value
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
}
export abstract class SampleSseActionResFactory {
	abstract create(data: unknown): SampleSseActionRes;
}
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