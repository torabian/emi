import { SSEFetch, URLSearchParamsX, buildUrl, fetchx, isPlausibleObject, type TypedRequestInit, withPrefix } from './sdk/js';
import { useSse } from './sdk/react';
/**
* Action to communicate with the action sampleSse
*/
export type SampleSseActionOptions = {
	queryKey?: unknown[];
	qs?: SampleSseActionQueryParams;
	headers?: SampleSseActionReqHeaders;
};
export const useSampleSseAction = (options: {
	qs?: SampleSseActionQueryParams,
	init?: TypedRequestInit<SampleSseActionRes, SampleSseActionReqHeaders>,
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
	qs?: SampleSseActionQueryParams
  ) => buildUrl(
		SampleSseAction.URL,
		 undefined,
		qs
	);
  static Method = 'get';
	static Fetch = async (
			onMessage?: (ev: MessageEvent) => void,
		qs?: SampleSseActionQueryParams,
		init?: TypedRequestInit<SampleSseActionRes, SampleSseActionReqHeaders>,
		overrideUrl?: string
	) => {
		const res = await fetchx<SampleSseActionRes, unknown, SampleSseActionReqHeaders>(
			overrideUrl ?? SampleSseAction.NewUrl(
				qs
			),
			{
				method: SampleSseAction.Method,
				...(init || {})
			}
		)
			return SSEFetch(res, onMessage, init?.signal || undefined);
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
 message?: string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SampleSseActionResType {
}
/**
 * SampleSseActionReqHeaders class
 * Auto-generated from EmiAction
 */
export class SampleSseActionReqHeaders extends Headers {
  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }
}
/**
 * SampleSseActionResHeaders class
 * Auto-generated from EmiAction
 */
export class SampleSseActionResHeaders extends Headers {
  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }
}
/**
 * SampleSseActionQueryParams class
 * Auto-generated from EmiAction
 */
export class SampleSseActionQueryParams extends URLSearchParamsX {
}