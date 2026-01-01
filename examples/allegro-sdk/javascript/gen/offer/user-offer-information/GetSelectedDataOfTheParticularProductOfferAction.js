import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get selected data of the particular product-offer
*/
	/**
 * GetSelectedDataOfTheParticularProductOfferAction
 */
export class GetSelectedDataOfTheParticularProductOfferAction { //
  static URL = 'https://api.{environment}/sale/product-offers/{offerId}/parts';
  static NewUrl = (
	qs
  ) => buildUrl(
		GetSelectedDataOfTheParticularProductOfferAction.URL,
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
			overrideUrl ?? GetSelectedDataOfTheParticularProductOfferAction.NewUrl(
				qs
			),
			{
				method: GetSelectedDataOfTheParticularProductOfferAction.Method,
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
				creatorFn: (item) => new GetSelectedDataOfTheParticularProductOfferActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetSelectedDataOfTheParticularProductOfferActionRes(item))
		const res = await GetSelectedDataOfTheParticularProductOfferAction.Fetch$(
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
  "name": "Get selected data of the particular product-offer",
  "url": "https://api.{environment}/sale/product-offers/{offerId}/parts",
  "method": "get",
  "description": "Use this resource to retrieve selected data of the particular product-offer. The model and functionality is a subset of the full product offer get endpoint (GET /sale/product-offers/{offerId}), but it is faster and more reliable.",
  "out": {
    "fields": [
      {
        "name": "id",
        "description": "Unique offer identifier",
        "type": "string"
      },
      {
        "name": "stock",
        "type": "object",
        "fields": [
          {
            "name": "available",
            "description": "Number of available items in stock",
            "type": "int"
          }
        ]
      },
      {
        "name": "sellingMode",
        "type": "object",
        "fields": [
          {
            "name": "price",
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
        "name": "additionalMarketplaces",
        "description": "Marketplace-specific price information",
        "type": "object",
        "fields": [
          {
            "name": "marketplaceId1",
            "type": "object",
            "fields": [
              {
                "name": "sellingMode",
                "type": "object",
                "fields": [
                  {
                    "name": "price",
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
          {
            "name": "marketplaceId2",
            "type": "object",
            "fields": [
              {
                "name": "sellingMode",
                "type": "object",
                "fields": [
                  {
                    "name": "price",
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
          }
        ]
      }
    ]
  }
}
}
/**
  * The base class definition for getSelectedDataOfTheParticularProductOfferActionRes
  **/
export class GetSelectedDataOfTheParticularProductOfferActionRes {
		/**
  * Unique offer identifier
  * @type {string}
  **/
 #id  =  ""
		/**
  * Unique offer identifier
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * Unique offer identifier
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
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.Stock}
  **/
 #stock
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.Stock}
  **/
set stock (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.Stock) {
			this.#stock = value
		} else {
			this.#stock = new GetSelectedDataOfTheParticularProductOfferActionRes.Stock(value)
		}
}
setStock (value) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
 #sellingMode
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
set sellingMode (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode(value)
		}
}
setSellingMode (value) {
	this.sellingMode = value
	return this
}
		/**
  * Marketplace-specific price information
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces}
  **/
 #additionalMarketplaces
		/**
  * Marketplace-specific price information
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * Marketplace-specific price information
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces(value)
		}
}
setAdditionalMarketplaces (value) {
	this.additionalMarketplaces = value
	return this
}
/**
  * The base class definition for stock
  **/
static Stock = class Stock {
		/**
  * Number of available items in stock
  * @type {number}
  **/
 #available  =  0
		/**
  * Number of available items in stock
  * @returns {number}
  **/
get available () { return this.#available }
/**
  * Number of available items in stock
  * @type {number}
  **/
set available (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#available = parsedValue;
		}
}
setAvailable (value) {
	this.available = value
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
			if (d.available !== undefined) { this.available = d.available }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				available: this.#available,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			available: 'available',
	  }
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.Stock, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.Stock(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.Stock ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.Stock(this.toJSON());
	}
}
/**
  * The base class definition for sellingMode
  **/
static SellingMode = class SellingMode {
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
 #price
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
set price (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price(value)
		}
}
setPrice (value) {
	this.price = value
	return this
}
/**
  * The base class definition for price
  **/
static Price = class Price {
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
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price(this.toJSON());
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
			if (d.price !== undefined) { this.price = d.price }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.price instanceof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price)) { this.price = new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price(d.price || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				price: this.#price,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			price$: 'price',
get price() {
					return withPrefix(
						"sellingMode.price",
						GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode(this.toJSON());
	}
}
/**
  * The base class definition for additionalMarketplaces
  **/
static AdditionalMarketplaces = class AdditionalMarketplaces {
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1}
  **/
 #marketplaceId1
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1}
  **/
get marketplaceId1 () { return this.#marketplaceId1 }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1}
  **/
set marketplaceId1 (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1) {
			this.#marketplaceId1 = value
		} else {
			this.#marketplaceId1 = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1(value)
		}
}
setMarketplaceId1 (value) {
	this.marketplaceId1 = value
	return this
}
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2}
  **/
 #marketplaceId2
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2}
  **/
get marketplaceId2 () { return this.#marketplaceId2 }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2}
  **/
set marketplaceId2 (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2) {
			this.#marketplaceId2 = value
		} else {
			this.#marketplaceId2 = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2(value)
		}
}
setMarketplaceId2 (value) {
	this.marketplaceId2 = value
	return this
}
/**
  * The base class definition for marketplaceId1
  **/
static MarketplaceId1 = class MarketplaceId1 {
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode}
  **/
 #sellingMode
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode}
  **/
set sellingMode (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode(value)
		}
}
setSellingMode (value) {
	this.sellingMode = value
	return this
}
/**
  * The base class definition for sellingMode
  **/
static SellingMode = class SellingMode {
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price}
  **/
 #price
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price}
  **/
set price (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price(value)
		}
}
setPrice (value) {
	this.price = value
	return this
}
/**
  * The base class definition for price
  **/
static Price = class Price {
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
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price(this.toJSON());
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
			if (d.price !== undefined) { this.price = d.price }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.price instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price)) { this.price = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price(d.price || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				price: this.#price,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			price$: 'price',
get price() {
					return withPrefix(
						"additionalMarketplaces.marketplaceId1.sellingMode.price",
						GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode(this.toJSON());
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
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.sellingMode instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode)) { this.sellingMode = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode(d.sellingMode || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				sellingMode: this.#sellingMode,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			sellingMode$: 'sellingMode',
get sellingMode() {
					return withPrefix(
						"additionalMarketplaces.marketplaceId1.sellingMode",
						GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1 ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1(this.toJSON());
	}
}
/**
  * The base class definition for marketplaceId2
  **/
static MarketplaceId2 = class MarketplaceId2 {
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode}
  **/
 #sellingMode
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode}
  **/
set sellingMode (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode(value)
		}
}
setSellingMode (value) {
	this.sellingMode = value
	return this
}
/**
  * The base class definition for sellingMode
  **/
static SellingMode = class SellingMode {
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price}
  **/
 #price
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price}
  **/
set price (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price(value)
		}
}
setPrice (value) {
	this.price = value
	return this
}
/**
  * The base class definition for price
  **/
static Price = class Price {
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
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price(this.toJSON());
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
			if (d.price !== undefined) { this.price = d.price }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.price instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price)) { this.price = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price(d.price || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				price: this.#price,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			price$: 'price',
get price() {
					return withPrefix(
						"additionalMarketplaces.marketplaceId2.sellingMode.price",
						GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode(this.toJSON());
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
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.sellingMode instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode)) { this.sellingMode = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode(d.sellingMode || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				sellingMode: this.#sellingMode,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			sellingMode$: 'sellingMode',
get sellingMode() {
					return withPrefix(
						"additionalMarketplaces.marketplaceId2.sellingMode",
						GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2 ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2(this.toJSON());
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
			if (d.marketplaceId1 !== undefined) { this.marketplaceId1 = d.marketplaceId1 }
			if (d.marketplaceId2 !== undefined) { this.marketplaceId2 = d.marketplaceId2 }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.marketplaceId1 instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1)) { this.marketplaceId1 = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1(d.marketplaceId1 || {}) }	
			if (!(d.marketplaceId2 instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2)) { this.marketplaceId2 = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2(d.marketplaceId2 || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				marketplaceId1: this.#marketplaceId1,
				marketplaceId2: this.#marketplaceId2,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			marketplaceId1$: 'marketplaceId1',
get marketplaceId1() {
					return withPrefix(
						"additionalMarketplaces.marketplaceId1",
						GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.Fields
						);
						},
			marketplaceId2$: 'marketplaceId2',
get marketplaceId2() {
					return withPrefix(
						"additionalMarketplaces.marketplaceId2",
						GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces(this.toJSON());
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
			if (d.stock !== undefined) { this.stock = d.stock }
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.stock instanceof GetSelectedDataOfTheParticularProductOfferActionRes.Stock)) { this.stock = new GetSelectedDataOfTheParticularProductOfferActionRes.Stock(d.stock || {}) }	
			if (!(d.sellingMode instanceof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode)) { this.sellingMode = new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode(d.sellingMode || {}) }	
			if (!(d.additionalMarketplaces instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces)) { this.additionalMarketplaces = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces(d.additionalMarketplaces || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				stock: this.#stock,
				sellingMode: this.#sellingMode,
				additionalMarketplaces: this.#additionalMarketplaces,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			stock$: 'stock',
get stock() {
					return withPrefix(
						"stock",
						GetSelectedDataOfTheParticularProductOfferActionRes.Stock.Fields
						);
						},
			sellingMode$: 'sellingMode',
get sellingMode() {
					return withPrefix(
						"sellingMode",
						GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Fields
						);
						},
			additionalMarketplaces$: 'additionalMarketplaces',
get additionalMarketplaces() {
					return withPrefix(
						"additionalMarketplaces",
						GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSelectedDataOfTheParticularProductOfferActionRes(this.toJSON());
	}
}