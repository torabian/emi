import { SSEFetch, buildUrl, fetchx, isPlausibleObject, type TypedRequestInit, withPrefix } from './sdk/js';
/**
* Action to communicate with the action httpAction
*/
export type HttpActionActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * HttpActionAction
 */
export class HttpActionAction {
  static URL = 'http://localhost:8081 (for test we use override)';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		HttpActionAction.URL,
		 undefined,
		qs
	);
  static Method = 'post';
	static Fetch = async (
		qs?: URLSearchParams,
		init?: TypedRequestInit<HttpActionActionRes, unknown>,
		onMessage?: (ev: MessageEvent) => void,
		overrideUrl?: string,
	) => {
		const res = await fetchx<HttpActionActionRes, unknown, unknown>(
			overrideUrl ?? HttpActionAction.NewUrl(
				qs
			),
			{
				method: HttpActionAction.Method,
				...(init || {})
			}
		)
			const ct = res.headers.get("content-type") || "";
			const cd = res.headers.get("content-disposition") || "";
			if (ct.includes("text/event-stream")) {
				// delegate to SSEFetch
				return SSEFetch(res, onMessage, init?.signal);
			}
			if (cd.includes("attachment") || (!ct.includes("json") && !ct.startsWith("text/"))) {
				res.result = res.body;
			} else if (ct.includes("application/json")) {
				const json = await res.json();
				res.result = new HttpActionActionRes (json);
			} else {
				// plain text or fallback
				res.result = await res.text();
			}
			return { done: Promise.resolve(), response: res };
	}
}
/**
  * The base class definition for httpActionActionReq
  **/
export class HttpActionActionReq {
		/**
  * Minimum number which can be generated
  * @type {number}
  **/
 #min : number  =  0
		/**
  * Minimum number which can be generated
  * @returns {number}
  **/
get min () { return this.#min }
/**
  * Minimum number which can be generated
  * @type {number}
  **/
set min (value: number) {
	 	const correctType = typeof value === 'number'
		this.#min = correctType ? value : Number(value);
}
setMin (value: number) {
	this.min = value
	return this
}
		/**
  * Maximum number which can be generated
  * @type {number}
  **/
 #max : number  =  0
		/**
  * Maximum number which can be generated
  * @returns {number}
  **/
get max () { return this.#max }
/**
  * Maximum number which can be generated
  * @type {number}
  **/
set max (value: number) {
	 	const correctType = typeof value === 'number'
		this.#max = correctType ? value : Number(value);
}
setMax (value: number) {
	this.max = value
	return this
}
		/**
  * How many numbers you want to be generated based on maximum and minimum
  * @type {number}
  **/
 #count : number  =  0
		/**
  * How many numbers you want to be generated based on maximum and minimum
  * @returns {number}
  **/
get count () { return this.#count }
/**
  * How many numbers you want to be generated based on maximum and minimum
  * @type {number}
  **/
set count (value: number) {
	 	const correctType = typeof value === 'number'
		this.#count = correctType ? value : Number(value);
}
setCount (value: number) {
	this.count = value
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
		const d = data as Partial<HttpActionActionReq>;
			if (d.min !== undefined) { this.min = d.min }
			if (d.max !== undefined) { this.max = d.max }
			if (d.count !== undefined) { this.count = d.count }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				min: this.#min,
				max: this.#max,
				count: this.#count,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			min: 'min',
			max: 'max',
			count: 'count',
	  }
	}
}
export abstract class HttpActionActionReqFactory {
	abstract create(data: unknown): HttpActionActionReq;
}
	/**
  * The base type definition for httpActionActionReq
  **/
	export type HttpActionActionReqType =  {
			/**
  * Minimum number which can be generated
  * @type {number}
  **/
 min?: number;
			/**
  * Maximum number which can be generated
  * @type {number}
  **/
 max?: number;
			/**
  * How many numbers you want to be generated based on maximum and minimum
  * @type {number}
  **/
 count?: number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace HttpActionActionReqType {
}
/**
  * The base class definition for httpActionActionRes
  **/
export class HttpActionActionRes {
		/**
  * 
  * @type {number}
  **/
 #number : number  =  0
		/**
  * 
  * @returns {number}
  **/
get number () { return this.#number }
/**
  * 
  * @type {number}
  **/
set number (value: number) {
	 	const correctType = typeof value === 'number'
		this.#number = correctType ? value : Number(value);
}
setNumber (value: number) {
	this.number = value
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
		const d = data as Partial<HttpActionActionRes>;
			if (d.number !== undefined) { this.number = d.number }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				number: this.#number,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			number: 'number',
	  }
	}
}
export abstract class HttpActionActionResFactory {
	abstract create(data: unknown): HttpActionActionRes;
}
	/**
  * The base type definition for httpActionActionRes
  **/
	export type HttpActionActionResType =  {
			/**
  * 
  * @type {number}
  **/
 number?: number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace HttpActionActionResType {
}