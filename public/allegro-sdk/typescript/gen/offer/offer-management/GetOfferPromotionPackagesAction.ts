import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get offer promotion packages
*/
export type GetOfferPromotionPackagesActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * GetOfferPromotionPackagesAction
 */
export class GetOfferPromotionPackagesAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}/promo-options';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		GetOfferPromotionPackagesAction.URL,
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
		return fetchx<GetOfferPromotionPackagesActionRes, unknown, unknown>(
			overrideUrl ?? GetOfferPromotionPackagesAction.NewUrl(
				qs
			),
			{
				method: GetOfferPromotionPackagesAction.Method,
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
				creatorFn?: ((item: unknown) => GetOfferPromotionPackagesActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new GetOfferPromotionPackagesActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetOfferPromotionPackagesActionRes(item))
		const res = await GetOfferPromotionPackagesAction.Fetch$(
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
  "name": "Get offer promotion packages",
  "url": "https://api.{environment}/sale/offers/{offerId}/promo-options",
  "method": "get",
  "out": {
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
  }
}
}
/**
  * The base class definition for getOfferPromotionPackagesActionRes
  **/
export class GetOfferPromotionPackagesActionRes {
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
  * @type {GetOfferPromotionPackagesActionRes.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof GetOfferPromotionPackagesActionRes.BasePackage>
		/**
  * 
  * @returns {GetOfferPromotionPackagesActionRes.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.BasePackage}
  **/
set basePackage (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferPromotionPackagesActionRes.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetOfferPromotionPackagesActionRes.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.BasePackage>) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof GetOfferPromotionPackagesActionRes.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {GetOfferPromotionPackagesActionRes.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.ExtraPackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetOfferPromotionPackagesActionRes.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new GetOfferPromotionPackagesActionRes.ExtraPackages(item))
		}
}
setExtraPackages (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.PendingChanges}
  **/
 #pendingChanges ! : InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges>
		/**
  * 
  * @returns {GetOfferPromotionPackagesActionRes.PendingChanges}
  **/
get pendingChanges () { return this.#pendingChanges }
/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.PendingChanges}
  **/
set pendingChanges (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferPromotionPackagesActionRes.PendingChanges) {
			this.#pendingChanges = value
		} else {
			this.#pendingChanges = new GetOfferPromotionPackagesActionRes.PendingChanges(value)
		}
}
setPendingChanges (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges>) {
	this.pendingChanges = value
	return this
}
		/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces}
  **/
 #additionalMarketplaces : InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces>[]  =  []
		/**
  * 
  * @returns {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = value.map(item => new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces(item))
		}
}
setAdditionalMarketplaces (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces>[]) {
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
	* Creates an instance of GetOfferPromotionPackagesActionRes.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType.BasePackageType) {
		return new GetOfferPromotionPackagesActionRes.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType.BasePackageType>) {
		return new GetOfferPromotionPackagesActionRes.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType.BasePackageType>): InstanceType<typeof GetOfferPromotionPackagesActionRes.BasePackage> {
		return new GetOfferPromotionPackagesActionRes.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes.BasePackage> {
		return new GetOfferPromotionPackagesActionRes.BasePackage(this.toJSON());
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
	* Creates an instance of GetOfferPromotionPackagesActionRes.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType.ExtraPackagesType) {
		return new GetOfferPromotionPackagesActionRes.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType.ExtraPackagesType>) {
		return new GetOfferPromotionPackagesActionRes.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType.ExtraPackagesType>): InstanceType<typeof GetOfferPromotionPackagesActionRes.ExtraPackages> {
		return new GetOfferPromotionPackagesActionRes.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes.ExtraPackages> {
		return new GetOfferPromotionPackagesActionRes.ExtraPackages(this.toJSON());
	}
}
/**
  * The base class definition for pendingChanges
  **/
static PendingChanges = class PendingChanges {
		/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage>
		/**
  * 
  * @returns {GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage}
  **/
set basePackage (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage>) {
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
	* Creates an instance of GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType) {
		return new GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType>) {
		return new GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType>): InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage> {
		return new GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage> {
		return new GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage(this.toJSON());
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
			if (!(d.basePackage instanceof GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage)) { this.basePackage = new GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage(d.basePackage || {}) }	
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
						"pendingChanges.basePackage",
						GetOfferPromotionPackagesActionRes.PendingChanges.BasePackage.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.PendingChanges, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType.PendingChangesType) {
		return new GetOfferPromotionPackagesActionRes.PendingChanges(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.PendingChanges, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType.PendingChangesType>) {
		return new GetOfferPromotionPackagesActionRes.PendingChanges(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType.PendingChangesType>): InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges> {
		return new GetOfferPromotionPackagesActionRes.PendingChanges ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes.PendingChanges> {
		return new GetOfferPromotionPackagesActionRes.PendingChanges(this.toJSON());
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
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage>
		/**
  * 
  * @returns {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage}
  **/
set basePackage (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage>) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(item))
		}
}
setExtraPackages (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges}
  **/
 #pendingChanges ! : InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges>
		/**
  * 
  * @returns {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges}
  **/
get pendingChanges () { return this.#pendingChanges }
/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges}
  **/
set pendingChanges (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges) {
			this.#pendingChanges = value
		} else {
			this.#pendingChanges = new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(value)
		}
}
setPendingChanges (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges>) {
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
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType>) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType>): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(this.toJSON());
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
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType>) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType>): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(this.toJSON());
	}
}
/**
  * The base class definition for pendingChanges
  **/
static PendingChanges = class PendingChanges {
		/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage>
		/**
  * 
  * @returns {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
set basePackage (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage>) {
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
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType>) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType>): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(this.toJSON());
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
			if (!(d.basePackage instanceof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage)) { this.basePackage = new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(d.basePackage || {}) }	
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
						"additionalMarketplaces.pendingChanges.basePackage",
						GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType>) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType>): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(this.toJSON());
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
			if (!(d.basePackage instanceof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage)) { this.basePackage = new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(d.basePackage || {}) }	
			if (!(d.pendingChanges instanceof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges)) { this.pendingChanges = new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(d.pendingChanges || {}) }	
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
						"additionalMarketplaces.basePackage",
						GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"additionalMarketplaces.extraPackages[:i]",
						GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages.Fields
						);
						},
			pendingChanges$: 'pendingChanges',
get pendingChanges() {
					return withPrefix(
						"additionalMarketplaces.pendingChanges",
						GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType>) {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType>): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes.AdditionalMarketplaces> {
		return new GetOfferPromotionPackagesActionRes.AdditionalMarketplaces(this.toJSON());
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
		const d = data as Partial<GetOfferPromotionPackagesActionRes>;
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
		const d = data as Partial<GetOfferPromotionPackagesActionRes>;
			if (!(d.basePackage instanceof GetOfferPromotionPackagesActionRes.BasePackage)) { this.basePackage = new GetOfferPromotionPackagesActionRes.BasePackage(d.basePackage || {}) }	
			if (!(d.pendingChanges instanceof GetOfferPromotionPackagesActionRes.PendingChanges)) { this.pendingChanges = new GetOfferPromotionPackagesActionRes.PendingChanges(d.pendingChanges || {}) }	
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
						"basePackage",
						GetOfferPromotionPackagesActionRes.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"extraPackages[:i]",
						GetOfferPromotionPackagesActionRes.ExtraPackages.Fields
						);
						},
			pendingChanges$: 'pendingChanges',
get pendingChanges() {
					return withPrefix(
						"pendingChanges",
						GetOfferPromotionPackagesActionRes.PendingChanges.Fields
						);
						},
			additionalMarketplaces$: 'additionalMarketplaces',
get additionalMarketplaces() {
					return withPrefix(
						"additionalMarketplaces[:i]",
						GetOfferPromotionPackagesActionRes.AdditionalMarketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOfferPromotionPackagesActionResType) {
		return new GetOfferPromotionPackagesActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferPromotionPackagesActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferPromotionPackagesActionResType>) {
		return new GetOfferPromotionPackagesActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferPromotionPackagesActionResType>): InstanceType<typeof GetOfferPromotionPackagesActionRes> {
		return new GetOfferPromotionPackagesActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferPromotionPackagesActionRes> {
		return new GetOfferPromotionPackagesActionRes(this.toJSON());
	}
}
export abstract class GetOfferPromotionPackagesActionResFactory {
	abstract create(data: unknown): GetOfferPromotionPackagesActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for getOfferPromotionPackagesActionRes
  **/
	export type GetOfferPromotionPackagesActionResType =  {
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
  * @type {GetOfferPromotionPackagesActionResType.BasePackageType}
  **/
 basePackage : GetOfferPromotionPackagesActionResType.BasePackageType;
			/**
  * 
  * @type {GetOfferPromotionPackagesActionResType.ExtraPackagesType[]}
  **/
 extraPackages : GetOfferPromotionPackagesActionResType.ExtraPackagesType[];
			/**
  * 
  * @type {GetOfferPromotionPackagesActionResType.PendingChangesType}
  **/
 pendingChanges : GetOfferPromotionPackagesActionResType.PendingChangesType;
			/**
  * 
  * @type {GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType[]}
  **/
 additionalMarketplaces : GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetOfferPromotionPackagesActionResType {
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
  * @type {GetOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType}
  **/
 basePackage : GetOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType;
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
  * @type {GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType}
  **/
 basePackage : GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType;
			/**
  * 
  * @type {GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType[]}
  **/
 extraPackages : GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType[];
			/**
  * 
  * @type {GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType}
  **/
 pendingChanges : GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType;
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
  * @type {GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType}
  **/
 basePackage : GetOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType;
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