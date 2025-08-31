import { SSEFetch, URLSearchParamsX, buildUrl, fetchx, type TypedRequestInit } from './sdk/js';
/**
* Action to communicate with the action sampleSse
*/
export type SampleSseActionOptions = {
	queryKey?: unknown[];
	qs?: SampleSseQueryParams;
	headers?: SampleSseReqHeaders;
};
	/**
 * SampleSseAction
 */
export class SampleSseAction {
  static URL = 'http://localhost:3000/stream';
  static NewUrl = (
	qs?: SampleSseQueryParams
  ) => buildUrl(
		SampleSseAction.URL,
		 undefined,
		qs
	);
  static Method = 'get';
	static Fetch = async (
			onMessage?: (ev: MessageEvent) => void,
		qs?: SampleSseQueryParams,
		init?: TypedRequestInit<SampleSseRes, SampleSseReqHeaders>,
		overrideUrl?: string
	) => {
		const res = await fetchx<SampleSseRes, unknown, SampleSseReqHeaders>(
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
  * @description The base type definition for sampleSseRes
  **/
	export type SampleSseResType =  {
			/**
  * @type {string}
  * @description 
  **/
 message?: string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SampleSseResType {
}
/**
  * @decription The base class definition for sampleSseRes
  **/
export class SampleSseRes {
	constructor(data: unknown) {
		// This probably doesn't cover the nested objects
		const d = data as Partial<SampleSseRes>;
			if (d[`message`] !== undefined) { 
 this.setMessage (d[`message`]) 
}
	}
		/**
  * 
  * @type {string}
  **/
 message: string = ""
		/**
  * @returns {string}
  * @description 
  **/
getMessage () { return this[`message`] }
		/**
  * 
  * @param {string}
  **/
setMessage (value: string ) { this[`message`] = value; return this; } 
}
export abstract class SampleSseResFactory {
	abstract create(data: unknown): SampleSseRes;
}
/**
 * SampleSseReqHeaders class
 * Auto-generated from Module3Action
 */
export class SampleSseReqHeaders extends Headers {
  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }
}
/**
 * SampleSseResHeaders class
 * Auto-generated from Module3Action
 */
export class SampleSseResHeaders extends Headers {
  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }
}
/**
 * SampleSseQueryParams class
 * Auto-generated from Module3Action
 */
export class SampleSseQueryParams extends URLSearchParamsX {
}