import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get promo options for seller's offers
*/
export type GetPromoOptionsForSellerSOffersActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * GetPromoOptionsForSellerSOffersAction
 */
export class GetPromoOptionsForSellerSOffersAction { //
  static URL = 'https://api.{environment}/sale/offers/promo-options';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		GetPromoOptionsForSellerSOffersAction.URL,
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
		return fetchx<GetPromoOptionsForSellerSOffersActionRes, unknown, unknown>(
			overrideUrl ?? GetPromoOptionsForSellerSOffersAction.NewUrl(
				qs
			),
			{
				method: GetPromoOptionsForSellerSOffersAction.Method,
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
				creatorFn?: ((item: unknown) => GetPromoOptionsForSellerSOffersActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new GetPromoOptionsForSellerSOffersActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetPromoOptionsForSellerSOffersActionRes(item))
		const res = await GetPromoOptionsForSellerSOffersAction.Fetch$(
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
  "name": "Get promo options for seller's offers",
  "url": "https://api.{environment}/sale/offers/promo-options",
  "method": "get",
  "description": "Use this resource to retrieve promo options for seller offers. Read more: PL / EN.",
  "out": {
    "fields": [
      {
        "name": "promoOptions",
        "type": "array",
        "fields": [
          {
            "name": "offerId",
            "type": "string"
          },
          {
            "name": "marketplaceId",
            "type": "string"
          },
          {
            "name": "basePackage",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "validFrom",
                "type": "string"
              },
              {
                "name": "validTo",
                "type": "string"
              },
              {
                "name": "nextCycleDate",
                "type": "string"
              }
            ]
          },
          {
            "name": "extraPackages",
            "type": "array",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "validFrom",
                "type": "string"
              },
              {
                "name": "validTo",
                "type": "string"
              },
              {
                "name": "nextCycleDate",
                "type": "string"
              }
            ]
          },
          {
            "name": "pendingChanges",
            "type": "object",
            "fields": [
              {
                "name": "basePackage",
                "type": "object",
                "fields": [
                  {
                    "name": "id",
                    "type": "string"
                  },
                  {
                    "name": "validFrom",
                    "type": "string"
                  },
                  {
                    "name": "validTo",
                    "type": "string"
                  },
                  {
                    "name": "nextCycleDate",
                    "type": "string"
                  }
                ]
              }
            ]
          },
          {
            "name": "additionalMarketplaces",
            "type": "array",
            "fields": [
              {
                "name": "marketplaceId",
                "type": "string"
              },
              {
                "name": "basePackage",
                "type": "object",
                "fields": [
                  {
                    "name": "id",
                    "type": "string"
                  },
                  {
                    "name": "validFrom",
                    "type": "string"
                  },
                  {
                    "name": "validTo",
                    "type": "string"
                  },
                  {
                    "name": "nextCycleDate",
                    "type": "string"
                  }
                ]
              },
              {
                "name": "extraPackages",
                "type": "array",
                "fields": [
                  {
                    "name": "id",
                    "type": "string"
                  },
                  {
                    "name": "validFrom",
                    "type": "string"
                  },
                  {
                    "name": "validTo",
                    "type": "string"
                  },
                  {
                    "name": "nextCycleDate",
                    "type": "string"
                  }
                ]
              },
              {
                "name": "pendingChanges",
                "type": "object",
                "fields": [
                  {
                    "name": "basePackage",
                    "type": "object",
                    "fields": [
                      {
                        "name": "id",
                        "type": "string"
                      },
                      {
                        "name": "validFrom",
                        "type": "string"
                      },
                      {
                        "name": "validTo",
                        "type": "string"
                      },
                      {
                        "name": "nextCycleDate",
                        "type": "string"
                      }
                    ]
                  }
                ]
              }
            ]
          }
        ]
      },
      {
        "name": "count",
        "type": "int"
      },
      {
        "name": "totalCount",
        "type": "int"
      }
    ]
  }
}
}
/**
  * The base class definition for getPromoOptionsForSellerSOffersActionRes
  **/
export class GetPromoOptionsForSellerSOffersActionRes {
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions}
  **/
 #promoOptions : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions>[]  =  []
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions}
  **/
get promoOptions () { return this.#promoOptions }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions}
  **/
set promoOptions (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions) {
			this.#promoOptions = value
		} else {
			this.#promoOptions = value.map(item => new GetPromoOptionsForSellerSOffersActionRes.PromoOptions(item))
		}
}
setPromoOptions (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions>[]) {
	this.promoOptions = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #count : number  =  0
		/**
  * 
  * @returns {number}
  **/
get count () { return this.#count }
/**
  * 
  * @type {number}
  **/
set count (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#count = parsedValue;
		}
}
setCount (value: number) {
	this.count = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #totalCount : number  =  0
		/**
  * 
  * @returns {number}
  **/
get totalCount () { return this.#totalCount }
/**
  * 
  * @type {number}
  **/
set totalCount (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#totalCount = parsedValue;
		}
}
setTotalCount (value: number) {
	this.totalCount = value
	return this
}
/**
  * The base class definition for promoOptions
  **/
static PromoOptions = class PromoOptions {
		/**
  * 
  * @type {string}
  **/
 #offerId : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get offerId () { return this.#offerId }
/**
  * 
  * @type {string}
  **/
set offerId (value: string) {
		this.#offerId = String(value);
}
setOfferId (value: string) {
	this.offerId = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #marketplaceId : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get marketplaceId () { return this.#marketplaceId }
/**
  * 
  * @type {string}
  **/
set marketplaceId (value: string) {
		this.#marketplaceId = String(value);
}
setMarketplaceId (value: string) {
	this.marketplaceId = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage>
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage}
  **/
set basePackage (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage>) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages(item))
		}
}
setExtraPackages (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges}
  **/
 #pendingChanges ! : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges>
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges}
  **/
get pendingChanges () { return this.#pendingChanges }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges}
  **/
set pendingChanges (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges) {
			this.#pendingChanges = value
		} else {
			this.#pendingChanges = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges(value)
		}
}
setPendingChanges (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges>) {
	this.pendingChanges = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces}
  **/
 #additionalMarketplaces : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces>[]  =  []
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = value.map(item => new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces(item))
		}
}
setAdditionalMarketplaces (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces>[]) {
	this.additionalMarketplaces = value
	return this
}
/**
  * The base class definition for basePackage
  **/
static BasePackage = class BasePackage {
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
  * @type {string}
  **/
 #validFrom : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value: string) {
		this.#validFrom = String(value);
}
setValidFrom (value: string) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value: string) {
		this.#validTo = String(value);
}
setValidTo (value: string) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value: string) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value: string) {
	this.nextCycleDate = value
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
		const d = data as Partial<BasePackage>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.validFrom !== undefined) { this.validFrom = d.validFrom }
			if (d.validTo !== undefined) { this.validTo = d.validTo }
			if (d.nextCycleDate !== undefined) { this.nextCycleDate = d.nextCycleDate }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				validFrom: this.#validFrom,
				validTo: this.#validTo,
				nextCycleDate: this.#nextCycleDate,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			validFrom: 'validFrom',
			validTo: 'validTo',
			nextCycleDate: 'nextCycleDate',
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.BasePackageType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.BasePackageType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.BasePackageType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage(this.toJSON());
	}
}
/**
  * The base class definition for extraPackages
  **/
static ExtraPackages = class ExtraPackages {
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
  * @type {string}
  **/
 #validFrom : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value: string) {
		this.#validFrom = String(value);
}
setValidFrom (value: string) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value: string) {
		this.#validTo = String(value);
}
setValidTo (value: string) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value: string) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value: string) {
	this.nextCycleDate = value
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
		const d = data as Partial<ExtraPackages>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.validFrom !== undefined) { this.validFrom = d.validFrom }
			if (d.validTo !== undefined) { this.validTo = d.validTo }
			if (d.nextCycleDate !== undefined) { this.nextCycleDate = d.nextCycleDate }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				validFrom: this.#validFrom,
				validTo: this.#validTo,
				nextCycleDate: this.#nextCycleDate,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			validFrom: 'validFrom',
			validTo: 'validTo',
			nextCycleDate: 'nextCycleDate',
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.ExtraPackagesType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.ExtraPackagesType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.ExtraPackagesType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages(this.toJSON());
	}
}
/**
  * The base class definition for pendingChanges
  **/
static PendingChanges = class PendingChanges {
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage>
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage}
  **/
set basePackage (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage>) {
	this.basePackage = value
	return this
}
/**
  * The base class definition for basePackage
  **/
static BasePackage = class BasePackage {
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
  * @type {string}
  **/
 #validFrom : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value: string) {
		this.#validFrom = String(value);
}
setValidFrom (value: string) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value: string) {
		this.#validTo = String(value);
}
setValidTo (value: string) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value: string) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value: string) {
	this.nextCycleDate = value
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
		const d = data as Partial<BasePackage>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.validFrom !== undefined) { this.validFrom = d.validFrom }
			if (d.validTo !== undefined) { this.validTo = d.validTo }
			if (d.nextCycleDate !== undefined) { this.nextCycleDate = d.nextCycleDate }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				validFrom: this.#validFrom,
				validTo: this.#validTo,
				nextCycleDate: this.#nextCycleDate,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			validFrom: 'validFrom',
			validTo: 'validTo',
			nextCycleDate: 'nextCycleDate',
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType.BasePackageType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType.BasePackageType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType.BasePackageType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage(this.toJSON());
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
		const d = data as Partial<PendingChanges>;
			if (d.basePackage !== undefined) { this.basePackage = d.basePackage }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<PendingChanges>;
			if (!(d.basePackage instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage)) { this.basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage(d.basePackage || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				basePackage: this.#basePackage,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			basePackage$: 'basePackage',
get basePackage() {
					return withPrefix(
						"promoOptions.pendingChanges.basePackage",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges(this.toJSON());
	}
}
/**
  * The base class definition for additionalMarketplaces
  **/
static AdditionalMarketplaces = class AdditionalMarketplaces {
		/**
  * 
  * @type {string}
  **/
 #marketplaceId : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get marketplaceId () { return this.#marketplaceId }
/**
  * 
  * @type {string}
  **/
set marketplaceId (value: string) {
		this.#marketplaceId = String(value);
}
setMarketplaceId (value: string) {
	this.marketplaceId = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage>
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage}
  **/
set basePackage (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage>) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages(item))
		}
}
setExtraPackages (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges}
  **/
 #pendingChanges ! : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges>
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges}
  **/
get pendingChanges () { return this.#pendingChanges }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges}
  **/
set pendingChanges (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges) {
			this.#pendingChanges = value
		} else {
			this.#pendingChanges = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges(value)
		}
}
setPendingChanges (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges>) {
	this.pendingChanges = value
	return this
}
/**
  * The base class definition for basePackage
  **/
static BasePackage = class BasePackage {
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
  * @type {string}
  **/
 #validFrom : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value: string) {
		this.#validFrom = String(value);
}
setValidFrom (value: string) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value: string) {
		this.#validTo = String(value);
}
setValidTo (value: string) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value: string) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value: string) {
	this.nextCycleDate = value
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
		const d = data as Partial<BasePackage>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.validFrom !== undefined) { this.validFrom = d.validFrom }
			if (d.validTo !== undefined) { this.validTo = d.validTo }
			if (d.nextCycleDate !== undefined) { this.nextCycleDate = d.nextCycleDate }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				validFrom: this.#validFrom,
				validTo: this.#validTo,
				nextCycleDate: this.#nextCycleDate,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			validFrom: 'validFrom',
			validTo: 'validTo',
			nextCycleDate: 'nextCycleDate',
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.BasePackageType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.BasePackageType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.BasePackageType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage(this.toJSON());
	}
}
/**
  * The base class definition for extraPackages
  **/
static ExtraPackages = class ExtraPackages {
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
  * @type {string}
  **/
 #validFrom : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value: string) {
		this.#validFrom = String(value);
}
setValidFrom (value: string) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value: string) {
		this.#validTo = String(value);
}
setValidTo (value: string) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value: string) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value: string) {
	this.nextCycleDate = value
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
		const d = data as Partial<ExtraPackages>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.validFrom !== undefined) { this.validFrom = d.validFrom }
			if (d.validTo !== undefined) { this.validTo = d.validTo }
			if (d.nextCycleDate !== undefined) { this.nextCycleDate = d.nextCycleDate }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				validFrom: this.#validFrom,
				validTo: this.#validTo,
				nextCycleDate: this.#nextCycleDate,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			validFrom: 'validFrom',
			validTo: 'validTo',
			nextCycleDate: 'nextCycleDate',
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.ExtraPackagesType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.ExtraPackagesType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.ExtraPackagesType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages(this.toJSON());
	}
}
/**
  * The base class definition for pendingChanges
  **/
static PendingChanges = class PendingChanges {
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage>
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
set basePackage (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage>) {
	this.basePackage = value
	return this
}
/**
  * The base class definition for basePackage
  **/
static BasePackage = class BasePackage {
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
  * @type {string}
  **/
 #validFrom : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value: string) {
		this.#validFrom = String(value);
}
setValidFrom (value: string) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value: string) {
		this.#validTo = String(value);
}
setValidTo (value: string) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value: string) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value: string) {
	this.nextCycleDate = value
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
		const d = data as Partial<BasePackage>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.validFrom !== undefined) { this.validFrom = d.validFrom }
			if (d.validTo !== undefined) { this.validTo = d.validTo }
			if (d.nextCycleDate !== undefined) { this.nextCycleDate = d.nextCycleDate }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				validFrom: this.#validFrom,
				validTo: this.#validTo,
				nextCycleDate: this.#nextCycleDate,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			validFrom: 'validFrom',
			validTo: 'validTo',
			nextCycleDate: 'nextCycleDate',
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType.BasePackageType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType.BasePackageType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType.BasePackageType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage(this.toJSON());
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
		const d = data as Partial<PendingChanges>;
			if (d.basePackage !== undefined) { this.basePackage = d.basePackage }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<PendingChanges>;
			if (!(d.basePackage instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage)) { this.basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage(d.basePackage || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				basePackage: this.#basePackage,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			basePackage$: 'basePackage',
get basePackage() {
					return withPrefix(
						"promoOptions.additionalMarketplaces.pendingChanges.basePackage",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges(this.toJSON());
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
			if (d.marketplaceId !== undefined) { this.marketplaceId = d.marketplaceId }
			if (d.basePackage !== undefined) { this.basePackage = d.basePackage }
			if (d.extraPackages !== undefined) { this.extraPackages = d.extraPackages }
			if (d.pendingChanges !== undefined) { this.pendingChanges = d.pendingChanges }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<AdditionalMarketplaces>;
			if (!(d.basePackage instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage)) { this.basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage(d.basePackage || {}) }	
			if (!(d.pendingChanges instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges)) { this.pendingChanges = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges(d.pendingChanges || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				marketplaceId: this.#marketplaceId,
				basePackage: this.#basePackage,
				extraPackages: this.#extraPackages,
				pendingChanges: this.#pendingChanges,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			marketplaceId: 'marketplaceId',
			basePackage$: 'basePackage',
get basePackage() {
					return withPrefix(
						"promoOptions.additionalMarketplaces.basePackage",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"promoOptions.additionalMarketplaces.extraPackages[:i]",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages.Fields
						);
						},
			pendingChanges$: 'pendingChanges',
get pendingChanges() {
					return withPrefix(
						"promoOptions.additionalMarketplaces.pendingChanges",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces(this.toJSON());
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
		const d = data as Partial<PromoOptions>;
			if (d.offerId !== undefined) { this.offerId = d.offerId }
			if (d.marketplaceId !== undefined) { this.marketplaceId = d.marketplaceId }
			if (d.basePackage !== undefined) { this.basePackage = d.basePackage }
			if (d.extraPackages !== undefined) { this.extraPackages = d.extraPackages }
			if (d.pendingChanges !== undefined) { this.pendingChanges = d.pendingChanges }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<PromoOptions>;
			if (!(d.basePackage instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage)) { this.basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage(d.basePackage || {}) }	
			if (!(d.pendingChanges instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges)) { this.pendingChanges = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges(d.pendingChanges || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				offerId: this.#offerId,
				marketplaceId: this.#marketplaceId,
				basePackage: this.#basePackage,
				extraPackages: this.#extraPackages,
				pendingChanges: this.#pendingChanges,
				additionalMarketplaces: this.#additionalMarketplaces,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			offerId: 'offerId',
			marketplaceId: 'marketplaceId',
			basePackage$: 'basePackage',
get basePackage() {
					return withPrefix(
						"promoOptions.basePackage",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"promoOptions.extraPackages[:i]",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages.Fields
						);
						},
			pendingChanges$: 'pendingChanges',
get pendingChanges() {
					return withPrefix(
						"promoOptions.pendingChanges",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.Fields
						);
						},
			additionalMarketplaces$: 'additionalMarketplaces',
get additionalMarketplaces() {
					return withPrefix(
						"promoOptions.additionalMarketplaces[:i]",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType>) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes.PromoOptions> {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions(this.toJSON());
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
		const d = data as Partial<GetPromoOptionsForSellerSOffersActionRes>;
			if (d.promoOptions !== undefined) { this.promoOptions = d.promoOptions }
			if (d.count !== undefined) { this.count = d.count }
			if (d.totalCount !== undefined) { this.totalCount = d.totalCount }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				promoOptions: this.#promoOptions,
				count: this.#count,
				totalCount: this.#totalCount,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			promoOptions$: 'promoOptions',
get promoOptions() {
					return withPrefix(
						"promoOptions[:i]",
						GetPromoOptionsForSellerSOffersActionRes.PromoOptions.Fields
						);
						},
			count: 'count',
			totalCount: 'totalCount',
	  }
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetPromoOptionsForSellerSOffersActionResType) {
		return new GetPromoOptionsForSellerSOffersActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetPromoOptionsForSellerSOffersActionResType>) {
		return new GetPromoOptionsForSellerSOffersActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetPromoOptionsForSellerSOffersActionResType>): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes> {
		return new GetPromoOptionsForSellerSOffersActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetPromoOptionsForSellerSOffersActionRes> {
		return new GetPromoOptionsForSellerSOffersActionRes(this.toJSON());
	}
}
export abstract class GetPromoOptionsForSellerSOffersActionResFactory {
	abstract create(data: unknown): GetPromoOptionsForSellerSOffersActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for getPromoOptionsForSellerSOffersActionRes
  **/
	export type GetPromoOptionsForSellerSOffersActionResType =  {
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType[]}
  **/
 promoOptions : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType[];
			/**
  * 
  * @type {number}
  **/
 count : number;
			/**
  * 
  * @type {number}
  **/
 totalCount : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetPromoOptionsForSellerSOffersActionResType {
	/**
  * The base type definition for promoOptionsType
  **/
	export type PromoOptionsType =  {
			/**
  * 
  * @type {string}
  **/
 offerId : string;
			/**
  * 
  * @type {string}
  **/
 marketplaceId : string;
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.BasePackageType}
  **/
 basePackage : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.BasePackageType;
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.ExtraPackagesType[]}
  **/
 extraPackages : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.ExtraPackagesType[];
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType}
  **/
 pendingChanges : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType;
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType[]}
  **/
 additionalMarketplaces : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PromoOptionsType {
	/**
  * The base type definition for basePackageType
  **/
	export type BasePackageType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 validFrom : string;
			/**
  * 
  * @type {string}
  **/
 validTo : string;
			/**
  * 
  * @type {string}
  **/
 nextCycleDate : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BasePackageType {
}
	/**
  * The base type definition for extraPackagesType
  **/
	export type ExtraPackagesType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 validFrom : string;
			/**
  * 
  * @type {string}
  **/
 validTo : string;
			/**
  * 
  * @type {string}
  **/
 nextCycleDate : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ExtraPackagesType {
}
	/**
  * The base type definition for pendingChangesType
  **/
	export type PendingChangesType =  {
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType.BasePackageType}
  **/
 basePackage : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.PendingChangesType.BasePackageType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PendingChangesType {
	/**
  * The base type definition for basePackageType
  **/
	export type BasePackageType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 validFrom : string;
			/**
  * 
  * @type {string}
  **/
 validTo : string;
			/**
  * 
  * @type {string}
  **/
 nextCycleDate : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BasePackageType {
}
}
	/**
  * The base type definition for additionalMarketplacesType
  **/
	export type AdditionalMarketplacesType =  {
			/**
  * 
  * @type {string}
  **/
 marketplaceId : string;
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.BasePackageType}
  **/
 basePackage : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.BasePackageType;
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.ExtraPackagesType[]}
  **/
 extraPackages : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.ExtraPackagesType[];
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType}
  **/
 pendingChanges : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdditionalMarketplacesType {
	/**
  * The base type definition for basePackageType
  **/
	export type BasePackageType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 validFrom : string;
			/**
  * 
  * @type {string}
  **/
 validTo : string;
			/**
  * 
  * @type {string}
  **/
 nextCycleDate : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BasePackageType {
}
	/**
  * The base type definition for extraPackagesType
  **/
	export type ExtraPackagesType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 validFrom : string;
			/**
  * 
  * @type {string}
  **/
 validTo : string;
			/**
  * 
  * @type {string}
  **/
 nextCycleDate : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ExtraPackagesType {
}
	/**
  * The base type definition for pendingChangesType
  **/
	export type PendingChangesType =  {
			/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType.BasePackageType}
  **/
 basePackage : GetPromoOptionsForSellerSOffersActionResType.PromoOptionsType.AdditionalMarketplacesType.PendingChangesType.BasePackageType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PendingChangesType {
	/**
  * The base type definition for basePackageType
  **/
	export type BasePackageType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 validFrom : string;
			/**
  * 
  * @type {string}
  **/
 validTo : string;
			/**
  * 
  * @type {string}
  **/
 nextCycleDate : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BasePackageType {
}
}
}
}
}