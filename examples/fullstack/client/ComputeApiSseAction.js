import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action computeApiSse
*/
	/**
 * ComputeApiSseAction
 */
export class ComputeApiSseAction { //
  static URL = '/compute/sse';
  static NewUrl = (
	qs
  ) => buildUrl(
		ComputeApiSseAction.URL,
		 undefined,
		qs
	);
  static Method = 'get';
	static Fetch$ = async (
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
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
		init,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
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
 #initialVector1  =  []
		/**
  * 
  * @returns {number[]}
  **/
get initialVector1 () { return this.#initialVector1 }
/**
  * 
  * @type {number[]}
  **/
set initialVector1 (value) {
}
setInitialVector1 (value) {
	this.initialVector1 = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #value  =  undefined
		/**
  * 
  * @returns {string}
  **/
get value () { return this.#value }
/**
  * 
  * @type {string}
  **/
set value (value) {
	 	const correctType = typeof value === 'string' || value === undefined || value === null
		this.#value = correctType ? value : String(value);
}
setValue (value) {
	this.value = value
	return this
}
		/**
  * 
  * @type {number[]}
  **/
 #initialVector2  =  []
		/**
  * 
  * @returns {number[]}
  **/
get initialVector2 () { return this.#initialVector2 }
/**
  * 
  * @type {number[]}
  **/
set initialVector2 (value) {
}
setInitialVector2 (value) {
	this.initialVector2 = value
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
	static from(possibleDtoObject) {
		return new ComputeApiSseActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeApiSseActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ComputeApiSseActionReq(partialDtoObject);
	}
	copyWith(partial) {
		return new ComputeApiSseActionReq ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ComputeApiSseActionReq(this.toJSON());
	}
}
/**
  * The base class definition for computeApiSseActionRes
  **/
export class ComputeApiSseActionRes {
		/**
  * 
  * @type {number[]}
  **/
 #outputVector  =  []
		/**
  * 
  * @returns {number[]}
  **/
get outputVector () { return this.#outputVector }
/**
  * 
  * @type {number[]}
  **/
set outputVector (value) {
}
setOutputVector (value) {
	this.outputVector = value
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
	static from(possibleDtoObject) {
		return new ComputeApiSseActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeApiSseActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ComputeApiSseActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new ComputeApiSseActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ComputeApiSseActionRes(this.toJSON());
	}
}