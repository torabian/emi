import { buildUrl } from './sdk/common/buildUrl';
import { fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { type UseMutationOptions, useMutation } from '@tanstack/react-query';
import { useState } from 'react';
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
	SampleSseActionOptions
& {
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
 	const [isCompleted, setCompleteState] = useState(false);
	const mutationResult =  useMutation({
		mutationFn: (body: unknown) =>
			SampleSseAction.Fetch(
				options?.creatorFn,
				options?.qs,
				{
					body,
					headers: options?.headers,
				},
				options?.onMessage,
				options?.overrideUrl,
			).then((x) => {
				x.done.then(() => {
					setCompleteState(true);
				});
				return x.response;
			}),
		...(options || {}),
	});
	return {
		...mutationResult,
		isCompleted
	}
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
  static Definition = {
  "name": "sampleSse",
  "url": "http://localhost:3000/stream",
  "method": "sse",
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