import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get selected data of the particular product-offer
*/
export type GetSelectedDataOfTheParticularProductOfferActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * GetSelectedDataOfTheParticularProductOfferAction
 */
export class GetSelectedDataOfTheParticularProductOfferAction { //
  static URL = 'https://api.{environment}/sale/product-offers/{offerId}/parts';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		GetSelectedDataOfTheParticularProductOfferAction.URL,
		 undefined,
		qs
	);
  static Method = 'get';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<unknown, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<GetSelectedDataOfTheParticularProductOfferActionRes, unknown, unknown>(
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
		init?: TypedRequestInit<unknown, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => GetSelectedDataOfTheParticularProductOfferActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
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
 #id : string  =  ""
		/**
  * Unique offer identifier
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * Unique offer identifier
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
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.Stock}
  **/
 #stock ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.Stock>
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.Stock}
  **/
set stock (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.Stock>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.Stock) {
			this.#stock = value
		} else {
			this.#stock = new GetSelectedDataOfTheParticularProductOfferActionRes.Stock(value)
		}
}
setStock (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.Stock>) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
 #sellingMode ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode>
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
set sellingMode (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode(value)
		}
}
setSellingMode (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode>) {
	this.sellingMode = value
	return this
}
		/**
  * Marketplace-specific price information
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces}
  **/
 #additionalMarketplaces ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces>
		/**
  * Marketplace-specific price information
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * Marketplace-specific price information
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces(value)
		}
}
setAdditionalMarketplaces (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces>) {
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
 #available : number  =  0
		/**
  * Number of available items in stock
  * @returns {number}
  **/
get available () { return this.#available }
/**
  * Number of available items in stock
  * @type {number}
  **/
set available (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#available = parsedValue;
		}
}
setAvailable (value: number) {
	this.available = value
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
		const d = data as Partial<Stock>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.StockType) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.StockType>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.Stock(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.StockType>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.Stock> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.Stock ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.Stock> {
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
 #price ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price>
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
set price (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price(value)
		}
}
setPrice (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price>) {
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
		const d = data as Partial<Price>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode.Price(this.toJSON());
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
		const d = data as Partial<SellingMode>;
			if (d.price !== undefined) { this.price = d.price }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<SellingMode>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.SellingMode> {
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
 #marketplaceId1 ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1>
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1}
  **/
get marketplaceId1 () { return this.#marketplaceId1 }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1}
  **/
set marketplaceId1 (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1) {
			this.#marketplaceId1 = value
		} else {
			this.#marketplaceId1 = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1(value)
		}
}
setMarketplaceId1 (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1>) {
	this.marketplaceId1 = value
	return this
}
		/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2}
  **/
 #marketplaceId2 ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2>
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2}
  **/
get marketplaceId2 () { return this.#marketplaceId2 }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2}
  **/
set marketplaceId2 (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2) {
			this.#marketplaceId2 = value
		} else {
			this.#marketplaceId2 = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2(value)
		}
}
setMarketplaceId2 (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2>) {
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
 #sellingMode ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode>
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode}
  **/
set sellingMode (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode(value)
		}
}
setSellingMode (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode>) {
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
 #price ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price>
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price}
  **/
set price (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price(value)
		}
}
setPrice (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price>) {
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
		const d = data as Partial<Price>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType.PriceType) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType.PriceType>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType.PriceType>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode.Price(this.toJSON());
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
		const d = data as Partial<SellingMode>;
			if (d.price !== undefined) { this.price = d.price }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<SellingMode>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1.SellingMode(this.toJSON());
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
		const d = data as Partial<MarketplaceId1>;
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<MarketplaceId1>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1 ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId1> {
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
 #sellingMode ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode>
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode}
  **/
set sellingMode (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode(value)
		}
}
setSellingMode (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode>) {
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
 #price ! : InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price>
		/**
  * 
  * @returns {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price}
  **/
set price (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price(value)
		}
}
setPrice (value: InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price>) {
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
		const d = data as Partial<Price>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType.PriceType) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType.PriceType>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType.PriceType>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode.Price(this.toJSON());
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
		const d = data as Partial<SellingMode>;
			if (d.price !== undefined) { this.price = d.price }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<SellingMode>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2.SellingMode(this.toJSON());
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
		const d = data as Partial<MarketplaceId2>;
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<MarketplaceId2>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2 ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces.MarketplaceId2(this.toJSON());
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
		const d = data as Partial<AdditionalMarketplaces>;
			if (d.marketplaceId1 !== undefined) { this.marketplaceId1 = d.marketplaceId1 }
			if (d.marketplaceId2 !== undefined) { this.marketplaceId2 = d.marketplaceId2 }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<AdditionalMarketplaces>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes.AdditionalMarketplaces(this.toJSON());
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
		const d = data as Partial<GetSelectedDataOfTheParticularProductOfferActionRes>;
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
		const d = data as Partial<GetSelectedDataOfTheParticularProductOfferActionRes>;
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
	static from(possibleDtoObject: GetSelectedDataOfTheParticularProductOfferActionResType) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSelectedDataOfTheParticularProductOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType>) {
		return new GetSelectedDataOfTheParticularProductOfferActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSelectedDataOfTheParticularProductOfferActionResType>): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSelectedDataOfTheParticularProductOfferActionRes> {
		return new GetSelectedDataOfTheParticularProductOfferActionRes(this.toJSON());
	}
}
export abstract class GetSelectedDataOfTheParticularProductOfferActionResFactory {
	abstract create(data: unknown): GetSelectedDataOfTheParticularProductOfferActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for getSelectedDataOfTheParticularProductOfferActionRes
  **/
	export type GetSelectedDataOfTheParticularProductOfferActionResType =  {
			/**
  * Unique offer identifier
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.StockType}
  **/
 stock : GetSelectedDataOfTheParticularProductOfferActionResType.StockType;
			/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType}
  **/
 sellingMode : GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType;
			/**
  * Marketplace-specific price information
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType}
  **/
 additionalMarketplaces : GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetSelectedDataOfTheParticularProductOfferActionResType {
	/**
  * The base type definition for stockType
  **/
	export type StockType =  {
			/**
  * Number of available items in stock
  * @type {number}
  **/
 available : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace StockType {
}
	/**
  * The base type definition for sellingModeType
  **/
	export type SellingModeType =  {
			/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType}
  **/
 price : GetSelectedDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SellingModeType {
	/**
  * The base type definition for priceType
  **/
	export type PriceType =  {
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
export namespace PriceType {
}
}
	/**
  * The base type definition for additionalMarketplacesType
  **/
	export type AdditionalMarketplacesType =  {
			/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type}
  **/
 marketplaceId1 : GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type;
			/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type}
  **/
 marketplaceId2 : GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdditionalMarketplacesType {
	/**
  * The base type definition for marketplaceId1Type
  **/
	export type MarketplaceId1Type =  {
			/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType}
  **/
 sellingMode : GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MarketplaceId1Type {
	/**
  * The base type definition for sellingModeType
  **/
	export type SellingModeType =  {
			/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType.PriceType}
  **/
 price : GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId1Type.SellingModeType.PriceType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SellingModeType {
	/**
  * The base type definition for priceType
  **/
	export type PriceType =  {
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
export namespace PriceType {
}
}
}
	/**
  * The base type definition for marketplaceId2Type
  **/
	export type MarketplaceId2Type =  {
			/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType}
  **/
 sellingMode : GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MarketplaceId2Type {
	/**
  * The base type definition for sellingModeType
  **/
	export type SellingModeType =  {
			/**
  * 
  * @type {GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType.PriceType}
  **/
 price : GetSelectedDataOfTheParticularProductOfferActionResType.AdditionalMarketplacesType.MarketplaceId2Type.SellingModeType.PriceType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SellingModeType {
	/**
  * The base type definition for priceType
  **/
	export type PriceType =  {
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
export namespace PriceType {
}
}
}
}
}