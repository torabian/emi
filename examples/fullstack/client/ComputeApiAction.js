import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { URLSearchParamsX } from './sdk/common/URLSearchParamsX';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action computeApi
*/
	/**
 * ComputeApiAction
 */
export class ComputeApiAction { //
  static URL = '/compute/api';
  static NewUrl = (
	qs
  ) => buildUrl(
		ComputeApiAction.URL,
		 undefined,
		qs
	);
  static Method = 'post';
	static Fetch$ = async (
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
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
		init,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
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
	* Creates an instance of ComputeApiActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ComputeApiActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeApiActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ComputeApiActionReq(partialDtoObject);
	}
	copyWith(partial) {
		return new ComputeApiActionReq ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ComputeApiActionReq(this.toJSON());
	}
}
/**
  * The base class definition for computeApiActionRes
  **/
export class ComputeApiActionRes {
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
	* Creates an instance of ComputeApiActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ComputeApiActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ComputeApiActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ComputeApiActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new ComputeApiActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ComputeApiActionRes(this.toJSON());
	}
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