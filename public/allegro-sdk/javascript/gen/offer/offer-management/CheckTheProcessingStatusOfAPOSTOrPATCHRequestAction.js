import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Check the processing status of a POST or PATCH request
*/
	/**
 * CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction
 */
export class CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction { //
  static URL = 'https://api.{environment}/sale/product-offers/{offerId}/operations/{operationId}';
  static NewUrl = (
	qs
  ) => buildUrl(
		CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction.URL,
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
			overrideUrl ?? CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction.NewUrl(
				qs
			),
			{
				method: CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction.Method,
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
				creatorFn: (item) => new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(item))
		const res = await CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction.Fetch$(
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
  "name": "Check the processing status of a POST or PATCH request",
  "url": "https://api.{environment}/sale/product-offers/{offerId}/operations/{operationId}",
  "method": "get",
  "description": "The URI for the resource given by Location header of POST /sale/product-offers and PATCH /sale/product-offers/{offerId}. Use this resource to check processing status of a POST or PATCH request. Read more: PL / EN.",
  "out": {
    "fields": [
      {
        "name": "offer",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
          }
        ]
      },
      {
        "name": "operation",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "status",
            "type": "string"
          },
          {
            "name": "startedAt",
            "type": "string"
          }
        ]
      }
    ]
  }
}
}
/**
  * The base class definition for checkTheProcessingStatusOfAPOSTOrPATCHRequestActionRes
  **/
export class CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes {
		/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer}
  **/
 #offer
		/**
  * 
  * @returns {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer}
  **/
get offer () { return this.#offer }
/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer}
  **/
set offer (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer) {
			this.#offer = value
		} else {
			this.#offer = new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(value)
		}
}
setOffer (value) {
	this.offer = value
	return this
}
		/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation}
  **/
 #operation
		/**
  * 
  * @returns {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation}
  **/
get operation () { return this.#operation }
/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation}
  **/
set operation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation) {
			this.#operation = value
		} else {
			this.#operation = new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(value)
		}
}
setOperation (value) {
	this.operation = value
	return this
}
/**
  * The base class definition for offer
  **/
static Offer = class Offer {
		/**
  * 
  * @type {string}
  **/
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
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
			if (d.id !== undefined) { this.id = d.id }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
	  }
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(possibleDtoObject);
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(partialDtoObject);
	}
	copyWith(partial) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(this.toJSON());
	}
}
/**
  * The base class definition for operation
  **/
static Operation = class Operation {
		/**
  * 
  * @type {string}
  **/
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #status  =  ""
		/**
  * 
  * @returns {string}
  **/
get status () { return this.#status }
/**
  * 
  * @type {string}
  **/
set status (value) {
		this.#status = String(value);
}
setStatus (value) {
	this.status = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #startedAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get startedAt () { return this.#startedAt }
/**
  * 
  * @type {string}
  **/
set startedAt (value) {
		this.#startedAt = String(value);
}
setStartedAt (value) {
	this.startedAt = value
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
			if (d.id !== undefined) { this.id = d.id }
			if (d.status !== undefined) { this.status = d.status }
			if (d.startedAt !== undefined) { this.startedAt = d.startedAt }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				status: this.#status,
				startedAt: this.#startedAt,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			status: 'status',
			startedAt: 'startedAt',
	  }
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(possibleDtoObject);
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(partialDtoObject);
	}
	copyWith(partial) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(this.toJSON());
	}
}
	constructor(data) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
			if (d.offer !== undefined) { this.offer = d.offer }
			if (d.operation !== undefined) { this.operation = d.operation }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.offer instanceof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer)) { this.offer = new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(d.offer || {}) }	
			if (!(d.operation instanceof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation)) { this.operation = new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(d.operation || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				offer: this.#offer,
				operation: this.#operation,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			offer$: 'offer',
get offer() {
					return withPrefix(
						"offer",
						CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer.Fields
						);
						},
			operation$: 'operation',
get operation() {
					return withPrefix(
						"operation",
						CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(this.toJSON());
	}
}