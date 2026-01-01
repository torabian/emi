import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get promo options for seller's offers
*/
	/**
 * GetPromoOptionsForSellerSOffersAction
 */
export class GetPromoOptionsForSellerSOffersAction { //
  static URL = 'https://api.{environment}/sale/offers/promo-options';
  static NewUrl = (
	qs
  ) => buildUrl(
		GetPromoOptionsForSellerSOffersAction.URL,
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
		init,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
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
 #promoOptions  =  []
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions}
  **/
get promoOptions () { return this.#promoOptions }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions}
  **/
set promoOptions (value) {
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
setPromoOptions (value) {
	this.promoOptions = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #count  =  0
		/**
  * 
  * @returns {number}
  **/
get count () { return this.#count }
/**
  * 
  * @type {number}
  **/
set count (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#count = parsedValue;
		}
}
setCount (value) {
	this.count = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #totalCount  =  0
		/**
  * 
  * @returns {number}
  **/
get totalCount () { return this.#totalCount }
/**
  * 
  * @type {number}
  **/
set totalCount (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#totalCount = parsedValue;
		}
}
setTotalCount (value) {
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
 #offerId  =  ""
		/**
  * 
  * @returns {string}
  **/
get offerId () { return this.#offerId }
/**
  * 
  * @type {string}
  **/
set offerId (value) {
		this.#offerId = String(value);
}
setOfferId (value) {
	this.offerId = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #marketplaceId  =  ""
		/**
  * 
  * @returns {string}
  **/
get marketplaceId () { return this.#marketplaceId }
/**
  * 
  * @type {string}
  **/
set marketplaceId (value) {
		this.#marketplaceId = String(value);
}
setMarketplaceId (value) {
	this.marketplaceId = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage}
  **/
 #basePackage
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage}
  **/
set basePackage (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage(value)
		}
}
setBasePackage (value) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages}
  **/
 #extraPackages  =  []
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages}
  **/
set extraPackages (value) {
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
setExtraPackages (value) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges}
  **/
 #pendingChanges
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges}
  **/
get pendingChanges () { return this.#pendingChanges }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges}
  **/
set pendingChanges (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges) {
			this.#pendingChanges = value
		} else {
			this.#pendingChanges = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges(value)
		}
}
setPendingChanges (value) {
	this.pendingChanges = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces}
  **/
 #additionalMarketplaces  =  []
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value) {
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
setAdditionalMarketplaces (value) {
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
 #validFrom  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value) {
		this.#validFrom = String(value);
}
setValidFrom (value) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value) {
		this.#validTo = String(value);
}
setValidTo (value) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value) {
	this.nextCycleDate = value
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #validFrom  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value) {
		this.#validFrom = String(value);
}
setValidFrom (value) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value) {
		this.#validTo = String(value);
}
setValidTo (value) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value) {
	this.nextCycleDate = value
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #basePackage
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage}
  **/
set basePackage (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage(value)
		}
}
setBasePackage (value) {
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
 #validFrom  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value) {
		this.#validFrom = String(value);
}
setValidFrom (value) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value) {
		this.#validTo = String(value);
}
setValidTo (value) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value) {
	this.nextCycleDate = value
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges.BasePackage(this.toJSON());
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
			if (d.basePackage !== undefined) { this.basePackage = d.basePackage }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.PendingChanges ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #marketplaceId  =  ""
		/**
  * 
  * @returns {string}
  **/
get marketplaceId () { return this.#marketplaceId }
/**
  * 
  * @type {string}
  **/
set marketplaceId (value) {
		this.#marketplaceId = String(value);
}
setMarketplaceId (value) {
	this.marketplaceId = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage}
  **/
 #basePackage
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage}
  **/
set basePackage (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage(value)
		}
}
setBasePackage (value) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages}
  **/
 #extraPackages  =  []
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages}
  **/
set extraPackages (value) {
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
setExtraPackages (value) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges}
  **/
 #pendingChanges
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges}
  **/
get pendingChanges () { return this.#pendingChanges }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges}
  **/
set pendingChanges (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges) {
			this.#pendingChanges = value
		} else {
			this.#pendingChanges = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges(value)
		}
}
setPendingChanges (value) {
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
 #validFrom  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value) {
		this.#validFrom = String(value);
}
setValidFrom (value) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value) {
		this.#validTo = String(value);
}
setValidTo (value) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value) {
	this.nextCycleDate = value
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #validFrom  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value) {
		this.#validFrom = String(value);
}
setValidFrom (value) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value) {
		this.#validTo = String(value);
}
setValidTo (value) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value) {
	this.nextCycleDate = value
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #basePackage
		/**
  * 
  * @returns {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
set basePackage (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage(value)
		}
}
setBasePackage (value) {
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
 #validFrom  =  ""
		/**
  * 
  * @returns {string}
  **/
get validFrom () { return this.#validFrom }
/**
  * 
  * @type {string}
  **/
set validFrom (value) {
		this.#validFrom = String(value);
}
setValidFrom (value) {
	this.validFrom = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #validTo  =  ""
		/**
  * 
  * @returns {string}
  **/
get validTo () { return this.#validTo }
/**
  * 
  * @type {string}
  **/
set validTo (value) {
		this.#validTo = String(value);
}
setValidTo (value) {
	this.validTo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #nextCycleDate  =  ""
		/**
  * 
  * @returns {string}
  **/
get nextCycleDate () { return this.#nextCycleDate }
/**
  * 
  * @type {string}
  **/
set nextCycleDate (value) {
		this.#nextCycleDate = String(value);
}
setNextCycleDate (value) {
	this.nextCycleDate = value
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges.BasePackage(this.toJSON());
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
			if (d.basePackage !== undefined) { this.basePackage = d.basePackage }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces.PendingChanges(this.toJSON());
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions.AdditionalMarketplaces(this.toJSON());
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes.PromoOptions, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetPromoOptionsForSellerSOffersActionRes.PromoOptions(this.toJSON());
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
	static from(possibleDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetPromoOptionsForSellerSOffersActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetPromoOptionsForSellerSOffersActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new GetPromoOptionsForSellerSOffersActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetPromoOptionsForSellerSOffersActionRes(this.toJSON());
	}
}