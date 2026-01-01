import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Modify the Buy Now price in an offer
*/
	/**
 * ModifyTheBuyNowPriceInAnOfferAction
 */
export class ModifyTheBuyNowPriceInAnOfferAction { //
  static URL = 'https://api.{environment}/offers/{offerId}/change-price-commands/{commandId}';
  static NewUrl = (
	qs
  ) => buildUrl(
		ModifyTheBuyNowPriceInAnOfferAction.URL,
		 undefined,
		qs
	);
  static Method = 'put';
	static Fetch$ = async (
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
			overrideUrl ?? ModifyTheBuyNowPriceInAnOfferAction.NewUrl(
				qs
			),
			{
				method: ModifyTheBuyNowPriceInAnOfferAction.Method,
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
				creatorFn: (item) => new ModifyTheBuyNowPriceInAnOfferActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new ModifyTheBuyNowPriceInAnOfferActionRes(item))
		const res = await ModifyTheBuyNowPriceInAnOfferAction.Fetch$(
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
  "name": "Modify the Buy Now price in an offer",
  "url": "https://api.{environment}/offers/{offerId}/change-price-commands/{commandId}",
  "method": "put",
  "description": "Use this resource to change the Buy Now price in a single offer. Read more: PL / EN.",
  "in": {
    "fields": [
      {
        "name": "id",
        "type": "string"
      },
      {
        "name": "input",
        "type": "object",
        "fields": [
          {
            "name": "buyNowPrice",
            "type": "object",
            "fields": [
              {
                "name": "amount",
                "type": "string"
              },
              {
                "name": "currency",
                "type": "string"
              }
            ]
          }
        ]
      }
    ]
  },
  "out": {
    "fields": [
      {
        "name": "id",
        "type": "string"
      },
      {
        "name": "input",
        "type": "object",
        "fields": [
          {
            "name": "buyNowPrice",
            "type": "object",
            "fields": [
              {
                "name": "amount",
                "type": "string"
              },
              {
                "name": "currency",
                "type": "string"
              }
            ]
          }
        ]
      },
      {
        "name": "output",
        "type": "object",
        "fields": [
          {
            "name": "status",
            "type": "string"
          },
          {
            "name": "errors",
            "type": "array",
            "fields": [
              {
                "name": "code",
                "type": "string"
              },
              {
                "name": "details",
                "type": "string"
              },
              {
                "name": "message",
                "type": "string"
              },
              {
                "name": "path",
                "type": "string"
              },
              {
                "name": "userMessage",
                "type": "string"
              },
              {
                "name": "metadata",
                "type": "object",
                "fields": [
                  {
                    "name": "productId",
                    "type": "string"
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  }
}
}
/**
  * The base class definition for modifyTheBuyNowPriceInAnOfferActionReq
  **/
export class ModifyTheBuyNowPriceInAnOfferActionReq {
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
  * @type {ModifyTheBuyNowPriceInAnOfferActionReq.Input}
  **/
 #input
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionReq.Input}
  **/
get input () { return this.#input }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionReq.Input}
  **/
set input (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionReq.Input) {
			this.#input = value
		} else {
			this.#input = new ModifyTheBuyNowPriceInAnOfferActionReq.Input(value)
		}
}
setInput (value) {
	this.input = value
	return this
}
/**
  * The base class definition for input
  **/
static Input = class Input {
		/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice}
  **/
 #buyNowPrice
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice}
  **/
get buyNowPrice () { return this.#buyNowPrice }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice}
  **/
set buyNowPrice (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice) {
			this.#buyNowPrice = value
		} else {
			this.#buyNowPrice = new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice(value)
		}
}
setBuyNowPrice (value) {
	this.buyNowPrice = value
	return this
}
/**
  * The base class definition for buyNowPrice
  **/
static BuyNowPrice = class BuyNowPrice {
		/**
  * 
  * @type {string}
  **/
 #amount  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value) {
		this.#amount = String(value);
}
setAmount (value) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value) {
		this.#currency = String(value);
}
setCurrency (value) {
	this.currency = value
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
			if (d.amount !== undefined) { this.amount = d.amount }
			if (d.currency !== undefined) { this.currency = d.currency }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				amount: this.#amount,
				currency: this.#currency,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			amount: 'amount',
			currency: 'currency',
	  }
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice(partialDtoObject);
	}
	copyWith(partial) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice(this.toJSON());
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
			if (d.buyNowPrice !== undefined) { this.buyNowPrice = d.buyNowPrice }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.buyNowPrice instanceof ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice)) { this.buyNowPrice = new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice(d.buyNowPrice || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				buyNowPrice: this.#buyNowPrice,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			buyNowPrice$: 'buyNowPrice',
get buyNowPrice() {
					return withPrefix(
						"input.buyNowPrice",
						ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionReq.Input, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionReq.Input, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input(partialDtoObject);
	}
	copyWith(partial) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input(this.toJSON());
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
			if (d.id !== undefined) { this.id = d.id }
			if (d.input !== undefined) { this.input = d.input }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.input instanceof ModifyTheBuyNowPriceInAnOfferActionReq.Input)) { this.input = new ModifyTheBuyNowPriceInAnOfferActionReq.Input(d.input || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				input: this.#input,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			input$: 'input',
get input() {
					return withPrefix(
						"input",
						ModifyTheBuyNowPriceInAnOfferActionReq.Input.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq(partialDtoObject);
	}
	copyWith(partial) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModifyTheBuyNowPriceInAnOfferActionReq(this.toJSON());
	}
}
/**
  * The base class definition for modifyTheBuyNowPriceInAnOfferActionRes
  **/
export class ModifyTheBuyNowPriceInAnOfferActionRes {
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
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Input}
  **/
 #input
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Input}
  **/
get input () { return this.#input }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Input}
  **/
set input (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Input) {
			this.#input = value
		} else {
			this.#input = new ModifyTheBuyNowPriceInAnOfferActionRes.Input(value)
		}
}
setInput (value) {
	this.input = value
	return this
}
		/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output}
  **/
 #output
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Output}
  **/
get output () { return this.#output }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output}
  **/
set output (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Output) {
			this.#output = value
		} else {
			this.#output = new ModifyTheBuyNowPriceInAnOfferActionRes.Output(value)
		}
}
setOutput (value) {
	this.output = value
	return this
}
/**
  * The base class definition for input
  **/
static Input = class Input {
		/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice}
  **/
 #buyNowPrice
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice}
  **/
get buyNowPrice () { return this.#buyNowPrice }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice}
  **/
set buyNowPrice (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice) {
			this.#buyNowPrice = value
		} else {
			this.#buyNowPrice = new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice(value)
		}
}
setBuyNowPrice (value) {
	this.buyNowPrice = value
	return this
}
/**
  * The base class definition for buyNowPrice
  **/
static BuyNowPrice = class BuyNowPrice {
		/**
  * 
  * @type {string}
  **/
 #amount  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value) {
		this.#amount = String(value);
}
setAmount (value) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value) {
		this.#currency = String(value);
}
setCurrency (value) {
	this.currency = value
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
			if (d.amount !== undefined) { this.amount = d.amount }
			if (d.currency !== undefined) { this.currency = d.currency }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				amount: this.#amount,
				currency: this.#currency,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			amount: 'amount',
			currency: 'currency',
	  }
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice(partialDtoObject);
	}
	copyWith(partial) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice(this.toJSON());
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
			if (d.buyNowPrice !== undefined) { this.buyNowPrice = d.buyNowPrice }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.buyNowPrice instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice)) { this.buyNowPrice = new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice(d.buyNowPrice || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				buyNowPrice: this.#buyNowPrice,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			buyNowPrice$: 'buyNowPrice',
get buyNowPrice() {
					return withPrefix(
						"input.buyNowPrice",
						ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Input, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Input, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input(partialDtoObject);
	}
	copyWith(partial) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input(this.toJSON());
	}
}
/**
  * The base class definition for output
  **/
static Output = class Output {
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
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors}
  **/
 #errors  =  []
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors}
  **/
set errors (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors) {
			this.#errors = value
		} else {
			this.#errors = value.map(item => new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors(item))
		}
}
setErrors (value) {
	this.errors = value
	return this
}
/**
  * The base class definition for errors
  **/
static Errors = class Errors {
		/**
  * 
  * @type {string}
  **/
 #code  =  ""
		/**
  * 
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * 
  * @type {string}
  **/
set code (value) {
		this.#code = String(value);
}
setCode (value) {
	this.code = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #details  =  ""
		/**
  * 
  * @returns {string}
  **/
get details () { return this.#details }
/**
  * 
  * @type {string}
  **/
set details (value) {
		this.#details = String(value);
}
setDetails (value) {
	this.details = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #message  =  ""
		/**
  * 
  * @returns {string}
  **/
get message () { return this.#message }
/**
  * 
  * @type {string}
  **/
set message (value) {
		this.#message = String(value);
}
setMessage (value) {
	this.message = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #path  =  ""
		/**
  * 
  * @returns {string}
  **/
get path () { return this.#path }
/**
  * 
  * @type {string}
  **/
set path (value) {
		this.#path = String(value);
}
setPath (value) {
	this.path = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #userMessage  =  ""
		/**
  * 
  * @returns {string}
  **/
get userMessage () { return this.#userMessage }
/**
  * 
  * @type {string}
  **/
set userMessage (value) {
		this.#userMessage = String(value);
}
setUserMessage (value) {
	this.userMessage = value
	return this
}
		/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata}
  **/
 #metadata
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata}
  **/
set metadata (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata(value)
		}
}
setMetadata (value) {
	this.metadata = value
	return this
}
/**
  * The base class definition for metadata
  **/
static Metadata = class Metadata {
		/**
  * 
  * @type {string}
  **/
 #productId  =  ""
		/**
  * 
  * @returns {string}
  **/
get productId () { return this.#productId }
/**
  * 
  * @type {string}
  **/
set productId (value) {
		this.#productId = String(value);
}
setProductId (value) {
	this.productId = value
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
			if (d.productId !== undefined) { this.productId = d.productId }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				productId: this.#productId,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			productId: 'productId',
	  }
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata(this.toJSON());
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
			if (d.code !== undefined) { this.code = d.code }
			if (d.details !== undefined) { this.details = d.details }
			if (d.message !== undefined) { this.message = d.message }
			if (d.path !== undefined) { this.path = d.path }
			if (d.userMessage !== undefined) { this.userMessage = d.userMessage }
			if (d.metadata !== undefined) { this.metadata = d.metadata }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.metadata instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata)) { this.metadata = new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata(d.metadata || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				code: this.#code,
				details: this.#details,
				message: this.#message,
				path: this.#path,
				userMessage: this.#userMessage,
				metadata: this.#metadata,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			code: 'code',
			details: 'details',
			message: 'message',
			path: 'path',
			userMessage: 'userMessage',
			metadata$: 'metadata',
get metadata() {
					return withPrefix(
						"output.errors.metadata",
						ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors(partialDtoObject);
	}
	copyWith(partial) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors(this.toJSON());
	}
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
			if (d.status !== undefined) { this.status = d.status }
			if (d.errors !== undefined) { this.errors = d.errors }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				status: this.#status,
				errors: this.#errors,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			status: 'status',
			errors$: 'errors',
get errors() {
					return withPrefix(
						"output.errors[:i]",
						ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Output, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Output, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output(partialDtoObject);
	}
	copyWith(partial) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output(this.toJSON());
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
			if (d.id !== undefined) { this.id = d.id }
			if (d.input !== undefined) { this.input = d.input }
			if (d.output !== undefined) { this.output = d.output }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.input instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Input)) { this.input = new ModifyTheBuyNowPriceInAnOfferActionRes.Input(d.input || {}) }	
			if (!(d.output instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Output)) { this.output = new ModifyTheBuyNowPriceInAnOfferActionRes.Output(d.output || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				input: this.#input,
				output: this.#output,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			input$: 'input',
get input() {
					return withPrefix(
						"input",
						ModifyTheBuyNowPriceInAnOfferActionRes.Input.Fields
						);
						},
			output$: 'output',
get output() {
					return withPrefix(
						"output",
						ModifyTheBuyNowPriceInAnOfferActionRes.Output.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModifyTheBuyNowPriceInAnOfferActionRes(this.toJSON());
	}
}