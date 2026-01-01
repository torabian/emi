import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get all available offer promotion packages
*/
export type GetAllAvailableOfferPromotionPackagesActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * GetAllAvailableOfferPromotionPackagesAction
 */
export class GetAllAvailableOfferPromotionPackagesAction { //
  static URL = 'https://api.{environment}/sale/offer-promotion-packages';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		GetAllAvailableOfferPromotionPackagesAction.URL,
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
		return fetchx<GetAllAvailableOfferPromotionPackagesActionRes, unknown, unknown>(
			overrideUrl ?? GetAllAvailableOfferPromotionPackagesAction.NewUrl(
				qs
			),
			{
				method: GetAllAvailableOfferPromotionPackagesAction.Method,
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
				creatorFn?: ((item: unknown) => GetAllAvailableOfferPromotionPackagesActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new GetAllAvailableOfferPromotionPackagesActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetAllAvailableOfferPromotionPackagesActionRes(item))
		const res = await GetAllAvailableOfferPromotionPackagesAction.Fetch$(
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
  "name": "Get all available offer promotion packages",
  "url": "https://api.{environment}/sale/offer-promotion-packages",
  "method": "get",
  "description": "Use this resource to retrieve all available offer promotion packages. Read more: PL / EN.",
  "out": {
    "fields": [
      {
        "name": "marketplaceId",
        "type": "string"
      },
      {
        "name": "basePackages",
        "type": "array",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "cycleDuration",
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
            "name": "name",
            "type": "string"
          },
          {
            "name": "cycleDuration",
            "type": "string"
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
            "name": "basePackages",
            "type": "array",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "name",
                "type": "string"
              },
              {
                "name": "cycleDuration",
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
                "name": "name",
                "type": "string"
              },
              {
                "name": "cycleDuration",
                "type": "string"
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
  * The base class definition for getAllAvailableOfferPromotionPackagesActionRes
  **/
export class GetAllAvailableOfferPromotionPackagesActionRes {
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
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.BasePackages}
  **/
 #basePackages : InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.BasePackages>[]  =  []
		/**
  * 
  * @returns {GetAllAvailableOfferPromotionPackagesActionRes.BasePackages}
  **/
get basePackages () { return this.#basePackages }
/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.BasePackages}
  **/
set basePackages (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.BasePackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllAvailableOfferPromotionPackagesActionRes.BasePackages) {
			this.#basePackages = value
		} else {
			this.#basePackages = value.map(item => new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages(item))
		}
}
setBasePackages (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.BasePackages>[]) {
	this.basePackages = value
	return this
}
		/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages(item))
		}
}
setExtraPackages (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces}
  **/
 #additionalMarketplaces : InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces>[]  =  []
		/**
  * 
  * @returns {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = value.map(item => new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces(item))
		}
}
setAdditionalMarketplaces (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces>[]) {
	this.additionalMarketplaces = value
	return this
}
/**
  * The base class definition for basePackages
  **/
static BasePackages = class BasePackages {
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
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #cycleDuration : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get cycleDuration () { return this.#cycleDuration }
/**
  * 
  * @type {string}
  **/
set cycleDuration (value: string) {
		this.#cycleDuration = String(value);
}
setCycleDuration (value: string) {
	this.cycleDuration = value
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
		const d = data as Partial<BasePackages>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
			if (d.cycleDuration !== undefined) { this.cycleDuration = d.cycleDuration }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
				cycleDuration: this.#cycleDuration,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
			cycleDuration: 'cycleDuration',
	  }
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.BasePackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllAvailableOfferPromotionPackagesActionResType.BasePackagesType) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.BasePackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.BasePackagesType>) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.BasePackagesType>): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.BasePackages> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.BasePackages> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.BasePackages(this.toJSON());
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
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #cycleDuration : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get cycleDuration () { return this.#cycleDuration }
/**
  * 
  * @type {string}
  **/
set cycleDuration (value: string) {
		this.#cycleDuration = String(value);
}
setCycleDuration (value: string) {
	this.cycleDuration = value
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
			if (d.name !== undefined) { this.name = d.name }
			if (d.cycleDuration !== undefined) { this.cycleDuration = d.cycleDuration }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
				cycleDuration: this.#cycleDuration,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
			cycleDuration: 'cycleDuration',
	  }
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllAvailableOfferPromotionPackagesActionResType.ExtraPackagesType) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.ExtraPackagesType>) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.ExtraPackagesType>): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages(this.toJSON());
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
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages}
  **/
 #basePackages : InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages>[]  =  []
		/**
  * 
  * @returns {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages}
  **/
get basePackages () { return this.#basePackages }
/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages}
  **/
set basePackages (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages) {
			this.#basePackages = value
		} else {
			this.#basePackages = value.map(item => new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages(item))
		}
}
setBasePackages (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages>[]) {
	this.basePackages = value
	return this
}
		/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(item))
		}
}
setExtraPackages (value: InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
/**
  * The base class definition for basePackages
  **/
static BasePackages = class BasePackages {
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
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #cycleDuration : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get cycleDuration () { return this.#cycleDuration }
/**
  * 
  * @type {string}
  **/
set cycleDuration (value: string) {
		this.#cycleDuration = String(value);
}
setCycleDuration (value: string) {
	this.cycleDuration = value
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
		const d = data as Partial<BasePackages>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
			if (d.cycleDuration !== undefined) { this.cycleDuration = d.cycleDuration }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
				cycleDuration: this.#cycleDuration,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
			cycleDuration: 'cycleDuration',
	  }
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackagesType) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackagesType>) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackagesType>): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages(this.toJSON());
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
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #cycleDuration : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get cycleDuration () { return this.#cycleDuration }
/**
  * 
  * @type {string}
  **/
set cycleDuration (value: string) {
		this.#cycleDuration = String(value);
}
setCycleDuration (value: string) {
	this.cycleDuration = value
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
			if (d.name !== undefined) { this.name = d.name }
			if (d.cycleDuration !== undefined) { this.cycleDuration = d.cycleDuration }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
				cycleDuration: this.#cycleDuration,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
			cycleDuration: 'cycleDuration',
	  }
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType>) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType>): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(this.toJSON());
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
		const d = data as Partial<AdditionalMarketplaces>;
			if (d.marketplaceId !== undefined) { this.marketplaceId = d.marketplaceId }
			if (d.basePackages !== undefined) { this.basePackages = d.basePackages }
			if (d.extraPackages !== undefined) { this.extraPackages = d.extraPackages }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				marketplaceId: this.#marketplaceId,
				basePackages: this.#basePackages,
				extraPackages: this.#extraPackages,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			marketplaceId: 'marketplaceId',
			basePackages$: 'basePackages',
get basePackages() {
					return withPrefix(
						"additionalMarketplaces.basePackages[:i]",
						GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackages.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"additionalMarketplaces.extraPackages[:i]",
						GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType>) {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType>): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces> {
		return new GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces(this.toJSON());
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
		const d = data as Partial<GetAllAvailableOfferPromotionPackagesActionRes>;
			if (d.marketplaceId !== undefined) { this.marketplaceId = d.marketplaceId }
			if (d.basePackages !== undefined) { this.basePackages = d.basePackages }
			if (d.extraPackages !== undefined) { this.extraPackages = d.extraPackages }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				marketplaceId: this.#marketplaceId,
				basePackages: this.#basePackages,
				extraPackages: this.#extraPackages,
				additionalMarketplaces: this.#additionalMarketplaces,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			marketplaceId: 'marketplaceId',
			basePackages$: 'basePackages',
get basePackages() {
					return withPrefix(
						"basePackages[:i]",
						GetAllAvailableOfferPromotionPackagesActionRes.BasePackages.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"extraPackages[:i]",
						GetAllAvailableOfferPromotionPackagesActionRes.ExtraPackages.Fields
						);
						},
			additionalMarketplaces$: 'additionalMarketplaces',
get additionalMarketplaces() {
					return withPrefix(
						"additionalMarketplaces[:i]",
						GetAllAvailableOfferPromotionPackagesActionRes.AdditionalMarketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllAvailableOfferPromotionPackagesActionResType) {
		return new GetAllAvailableOfferPromotionPackagesActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllAvailableOfferPromotionPackagesActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType>) {
		return new GetAllAvailableOfferPromotionPackagesActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllAvailableOfferPromotionPackagesActionResType>): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes> {
		return new GetAllAvailableOfferPromotionPackagesActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllAvailableOfferPromotionPackagesActionRes> {
		return new GetAllAvailableOfferPromotionPackagesActionRes(this.toJSON());
	}
}
export abstract class GetAllAvailableOfferPromotionPackagesActionResFactory {
	abstract create(data: unknown): GetAllAvailableOfferPromotionPackagesActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for getAllAvailableOfferPromotionPackagesActionRes
  **/
	export type GetAllAvailableOfferPromotionPackagesActionResType =  {
			/**
  * 
  * @type {string}
  **/
 marketplaceId : string;
			/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionResType.BasePackagesType[]}
  **/
 basePackages : GetAllAvailableOfferPromotionPackagesActionResType.BasePackagesType[];
			/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionResType.ExtraPackagesType[]}
  **/
 extraPackages : GetAllAvailableOfferPromotionPackagesActionResType.ExtraPackagesType[];
			/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType[]}
  **/
 additionalMarketplaces : GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetAllAvailableOfferPromotionPackagesActionResType {
	/**
  * The base type definition for basePackagesType
  **/
	export type BasePackagesType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
			/**
  * 
  * @type {string}
  **/
 cycleDuration : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BasePackagesType {
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
 name : string;
			/**
  * 
  * @type {string}
  **/
 cycleDuration : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ExtraPackagesType {
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
  * @type {GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackagesType[]}
  **/
 basePackages : GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackagesType[];
			/**
  * 
  * @type {GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType[]}
  **/
 extraPackages : GetAllAvailableOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdditionalMarketplacesType {
	/**
  * The base type definition for basePackagesType
  **/
	export type BasePackagesType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
			/**
  * 
  * @type {string}
  **/
 cycleDuration : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BasePackagesType {
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
 name : string;
			/**
  * 
  * @type {string}
  **/
 cycleDuration : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ExtraPackagesType {
}
}
}