import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Modify the Buy Now price in an offer
*/
export type ModifyTheBuyNowPriceInAnOfferActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * ModifyTheBuyNowPriceInAnOfferAction
 */
export class ModifyTheBuyNowPriceInAnOfferAction { //
  static URL = 'https://api.{environment}/offers/{offerId}/change-price-commands/{commandId}';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		ModifyTheBuyNowPriceInAnOfferAction.URL,
		 undefined,
		qs
	);
  static Method = 'put';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<ModifyTheBuyNowPriceInAnOfferActionReq, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<ModifyTheBuyNowPriceInAnOfferActionRes, ModifyTheBuyNowPriceInAnOfferActionReq, unknown>(
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
		init?: TypedRequestInit<ModifyTheBuyNowPriceInAnOfferActionReq, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => ModifyTheBuyNowPriceInAnOfferActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
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
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionReq.Input}
  **/
 #input ! : InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input>
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionReq.Input}
  **/
get input () { return this.#input }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionReq.Input}
  **/
set input (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionReq.Input) {
			this.#input = value
		} else {
			this.#input = new ModifyTheBuyNowPriceInAnOfferActionReq.Input(value)
		}
}
setInput (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input>) {
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
 #buyNowPrice ! : InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice>
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice}
  **/
get buyNowPrice () { return this.#buyNowPrice }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice}
  **/
set buyNowPrice (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice) {
			this.#buyNowPrice = value
		} else {
			this.#buyNowPrice = new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice(value)
		}
}
setBuyNowPrice (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice>) {
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
 #amount : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value: string) {
		this.#amount = String(value);
}
setAmount (value: string) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value: string) {
		this.#currency = String(value);
}
setCurrency (value: string) {
	this.currency = value
	return this
}
	constructor(data: unknown = undefined) {
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
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<BuyNowPrice>;
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
	static from(possibleDtoObject: ModifyTheBuyNowPriceInAnOfferActionReqType.InputType.BuyNowPriceType) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionReqType.InputType.BuyNowPriceType>) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionReqType.InputType.BuyNowPriceType>): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice> {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice> {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input.BuyNowPrice(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
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
	#isJsonAppliable(obj: unknown) {
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<Input>;
			if (d.buyNowPrice !== undefined) { this.buyNowPrice = d.buyNowPrice }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Input>;
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
	static from(possibleDtoObject: ModifyTheBuyNowPriceInAnOfferActionReqType.InputType) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionReq.Input, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionReqType.InputType>) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionReqType.InputType>): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input> {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq.Input> {
		return new ModifyTheBuyNowPriceInAnOfferActionReq.Input(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
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
	#isJsonAppliable(obj: unknown) {
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<ModifyTheBuyNowPriceInAnOfferActionReq>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.input !== undefined) { this.input = d.input }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<ModifyTheBuyNowPriceInAnOfferActionReq>;
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
	static from(possibleDtoObject: ModifyTheBuyNowPriceInAnOfferActionReqType) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionReqType>) {
		return new ModifyTheBuyNowPriceInAnOfferActionReq(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionReqType>): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq> {
		return new ModifyTheBuyNowPriceInAnOfferActionReq ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionReq> {
		return new ModifyTheBuyNowPriceInAnOfferActionReq(this.toJSON());
	}
}
export abstract class ModifyTheBuyNowPriceInAnOfferActionReqFactory {
	abstract create(data: unknown): ModifyTheBuyNowPriceInAnOfferActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for modifyTheBuyNowPriceInAnOfferActionReq
  **/
	export type ModifyTheBuyNowPriceInAnOfferActionReqType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionReqType.InputType}
  **/
 input : ModifyTheBuyNowPriceInAnOfferActionReqType.InputType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModifyTheBuyNowPriceInAnOfferActionReqType {
	/**
  * The base type definition for inputType
  **/
	export type InputType =  {
			/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionReqType.InputType.BuyNowPriceType}
  **/
 buyNowPrice : ModifyTheBuyNowPriceInAnOfferActionReqType.InputType.BuyNowPriceType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace InputType {
	/**
  * The base type definition for buyNowPriceType
  **/
	export type BuyNowPriceType =  {
			/**
  * 
  * @type {string}
  **/
 amount : string;
			/**
  * 
  * @type {string}
  **/
 currency : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BuyNowPriceType {
}
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
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Input}
  **/
 #input ! : InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input>
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Input}
  **/
get input () { return this.#input }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Input}
  **/
set input (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Input) {
			this.#input = value
		} else {
			this.#input = new ModifyTheBuyNowPriceInAnOfferActionRes.Input(value)
		}
}
setInput (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input>) {
	this.input = value
	return this
}
		/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output}
  **/
 #output ! : InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output>
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Output}
  **/
get output () { return this.#output }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output}
  **/
set output (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Output) {
			this.#output = value
		} else {
			this.#output = new ModifyTheBuyNowPriceInAnOfferActionRes.Output(value)
		}
}
setOutput (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output>) {
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
 #buyNowPrice ! : InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice>
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice}
  **/
get buyNowPrice () { return this.#buyNowPrice }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice}
  **/
set buyNowPrice (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice) {
			this.#buyNowPrice = value
		} else {
			this.#buyNowPrice = new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice(value)
		}
}
setBuyNowPrice (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice>) {
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
 #amount : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value: string) {
		this.#amount = String(value);
}
setAmount (value: string) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value: string) {
		this.#currency = String(value);
}
setCurrency (value: string) {
	this.currency = value
	return this
}
	constructor(data: unknown = undefined) {
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
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<BuyNowPrice>;
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
	static from(possibleDtoObject: ModifyTheBuyNowPriceInAnOfferActionResType.InputType.BuyNowPriceType) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.InputType.BuyNowPriceType>) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.InputType.BuyNowPriceType>): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input.BuyNowPrice(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
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
	#isJsonAppliable(obj: unknown) {
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<Input>;
			if (d.buyNowPrice !== undefined) { this.buyNowPrice = d.buyNowPrice }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Input>;
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
	static from(possibleDtoObject: ModifyTheBuyNowPriceInAnOfferActionResType.InputType) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Input, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.InputType>) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.InputType>): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Input ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Input> {
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
 #status : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get status () { return this.#status }
/**
  * 
  * @type {string}
  **/
set status (value: string) {
		this.#status = String(value);
}
setStatus (value: string) {
	this.status = value
	return this
}
		/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors}
  **/
 #errors : InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors>[]  =  []
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors}
  **/
set errors (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors>[]) {
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
setErrors (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors>[]) {
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
 #code : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * 
  * @type {string}
  **/
set code (value: string) {
		this.#code = String(value);
}
setCode (value: string) {
	this.code = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #details : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get details () { return this.#details }
/**
  * 
  * @type {string}
  **/
set details (value: string) {
		this.#details = String(value);
}
setDetails (value: string) {
	this.details = value
	return this
}
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
		this.#message = String(value);
}
setMessage (value: string) {
	this.message = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #path : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get path () { return this.#path }
/**
  * 
  * @type {string}
  **/
set path (value: string) {
		this.#path = String(value);
}
setPath (value: string) {
	this.path = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #userMessage : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get userMessage () { return this.#userMessage }
/**
  * 
  * @type {string}
  **/
set userMessage (value: string) {
		this.#userMessage = String(value);
}
setUserMessage (value: string) {
	this.userMessage = value
	return this
}
		/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata}
  **/
 #metadata ! : InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata>
		/**
  * 
  * @returns {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata}
  **/
set metadata (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata(value)
		}
}
setMetadata (value: InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata>) {
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
 #productId : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get productId () { return this.#productId }
/**
  * 
  * @type {string}
  **/
set productId (value: string) {
		this.#productId = String(value);
}
setProductId (value: string) {
	this.productId = value
	return this
}
	constructor(data: unknown = undefined) {
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
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<Metadata>;
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
	static from(possibleDtoObject: ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType.MetadataType) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType.MetadataType>) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType.MetadataType>): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors.Metadata(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
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
	#isJsonAppliable(obj: unknown) {
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<Errors>;
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
		const d = data as Partial<Errors>;
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
	static from(possibleDtoObject: ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType>) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType>): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output.Errors(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
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
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<Output>;
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
	static from(possibleDtoObject: ModifyTheBuyNowPriceInAnOfferActionResType.OutputType) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes.Output, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.OutputType>) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType.OutputType>): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes.Output> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes.Output(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
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
	#isJsonAppliable(obj: unknown) {
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<ModifyTheBuyNowPriceInAnOfferActionRes>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.input !== undefined) { this.input = d.input }
			if (d.output !== undefined) { this.output = d.output }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<ModifyTheBuyNowPriceInAnOfferActionRes>;
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
	static from(possibleDtoObject: ModifyTheBuyNowPriceInAnOfferActionResType) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyTheBuyNowPriceInAnOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType>) {
		return new ModifyTheBuyNowPriceInAnOfferActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyTheBuyNowPriceInAnOfferActionResType>): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyTheBuyNowPriceInAnOfferActionRes> {
		return new ModifyTheBuyNowPriceInAnOfferActionRes(this.toJSON());
	}
}
export abstract class ModifyTheBuyNowPriceInAnOfferActionResFactory {
	abstract create(data: unknown): ModifyTheBuyNowPriceInAnOfferActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for modifyTheBuyNowPriceInAnOfferActionRes
  **/
	export type ModifyTheBuyNowPriceInAnOfferActionResType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionResType.InputType}
  **/
 input : ModifyTheBuyNowPriceInAnOfferActionResType.InputType;
			/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionResType.OutputType}
  **/
 output : ModifyTheBuyNowPriceInAnOfferActionResType.OutputType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModifyTheBuyNowPriceInAnOfferActionResType {
	/**
  * The base type definition for inputType
  **/
	export type InputType =  {
			/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionResType.InputType.BuyNowPriceType}
  **/
 buyNowPrice : ModifyTheBuyNowPriceInAnOfferActionResType.InputType.BuyNowPriceType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace InputType {
	/**
  * The base type definition for buyNowPriceType
  **/
	export type BuyNowPriceType =  {
			/**
  * 
  * @type {string}
  **/
 amount : string;
			/**
  * 
  * @type {string}
  **/
 currency : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BuyNowPriceType {
}
}
	/**
  * The base type definition for outputType
  **/
	export type OutputType =  {
			/**
  * 
  * @type {string}
  **/
 status : string;
			/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType[]}
  **/
 errors : ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OutputType {
	/**
  * The base type definition for errorsType
  **/
	export type ErrorsType =  {
			/**
  * 
  * @type {string}
  **/
 code : string;
			/**
  * 
  * @type {string}
  **/
 details : string;
			/**
  * 
  * @type {string}
  **/
 message : string;
			/**
  * 
  * @type {string}
  **/
 path : string;
			/**
  * 
  * @type {string}
  **/
 userMessage : string;
			/**
  * 
  * @type {ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType.MetadataType}
  **/
 metadata : ModifyTheBuyNowPriceInAnOfferActionResType.OutputType.ErrorsType.MetadataType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ErrorsType {
	/**
  * The base type definition for metadataType
  **/
	export type MetadataType =  {
			/**
  * 
  * @type {string}
  **/
 productId : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MetadataType {
}
}
}
}