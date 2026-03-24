import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action computeExp
*/
	/**
 * ComputeExpAction
 */
export class ComputeExpAction { //
  static URL = '/big/exp/:first/:second';
  static NewUrl = (
	params,
	qs
  ) => buildUrl(
		ComputeExpAction.URL,
		params,
		qs
	);
  static Method = '';
	static Fetch$ = async (
			params,
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
			overrideUrl ?? ComputeExpAction.NewUrl(
				params,
				qs
			),
			{
				method: ComputeExpAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
			params,
		init,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
				creatorFn: (item) => new ComputeExpActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new ComputeExpActionRes(item))
		const res = await ComputeExpAction.Fetch$(
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
  "name": "computeExp",
  "url": "/big/exp/:first/:second",
  "description": "Computes the exp value using big integer",
  "in": {
    "fields": [
      {
        "name": "base",
        "type": "complex",
        "complex": "Int"
      },
      {
        "name": "exponent",
        "type": "complex",
        "complex": "Int"
      }
    ]
  },
  "out": {
    "fields": [
      {
        "name": "result",
        "type": "complex",
        "complex": "big.Int"
      }
    ]
  }
}
}
/**
  * The base class definition for computeExpActionReq
  **/
export class ComputeExpActionReq {
		/**
  * 
  * @type {Int}
  **/
 #base
		/**
  * 
  * @returns {Int}
  **/
get base () { return this.#base }
/**
  * 
  * @type {Int}
  **/
set base (value) {
}
setBase (value) {
	this.base = value
	return this
}
		/**
  * 
  * @type {Int}
  **/
 #exponent
		/**
  * 
  * @returns {Int}
  **/
get exponent () { return this.#exponent }
/**
  * 
  * @type {Int}
  **/
set exponent (value) {
}
setExponent (value) {
	this.exponent = value
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
		const g = globalThis
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
		const d = data;
			if (d.base !== undefined) { this.base = d.base }
			if (d.exponent !== undefined) { this.exponent = d.exponent }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				base: this.#base,
				exponent: this.#exponent,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			base: 'base',
			exponent: 'exponent',
	  }
	}
	/**
	* Creates an instance of ComputeExpActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ComputeExpActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeExpActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ComputeExpActionReq(partialDtoObject);
	}
	copyWith(partial) {
		return new ComputeExpActionReq ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ComputeExpActionReq(this.toJSON());
	}
}
/**
  * The base class definition for computeExpActionRes
  **/
export class ComputeExpActionRes {
		/**
  * 
  * @type {big.Int}
  **/
 #result
		/**
  * 
  * @returns {big.Int}
  **/
get result () { return this.#result }
/**
  * 
  * @type {big.Int}
  **/
set result (value) {
}
setResult (value) {
	this.result = value
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
		const g = globalThis
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
		const d = data;
			if (d.result !== undefined) { this.result = d.result }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				result: this.#result,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			result: 'result',
	  }
	}
	/**
	* Creates an instance of ComputeExpActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ComputeExpActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeExpActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ComputeExpActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new ComputeExpActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ComputeExpActionRes(this.toJSON());
	}
}