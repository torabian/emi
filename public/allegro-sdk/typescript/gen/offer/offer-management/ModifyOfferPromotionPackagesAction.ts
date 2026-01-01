import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Modify offer promotion packages
*/
export type ModifyOfferPromotionPackagesActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * ModifyOfferPromotionPackagesAction
 */
export class ModifyOfferPromotionPackagesAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}/promo-options-modification';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		ModifyOfferPromotionPackagesAction.URL,
		 undefined,
		qs
	);
  static Method = 'post';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<ModifyOfferPromotionPackagesActionReq, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<ModifyOfferPromotionPackagesActionRes, ModifyOfferPromotionPackagesActionReq, unknown>(
			overrideUrl ?? ModifyOfferPromotionPackagesAction.NewUrl(
				qs
			),
			{
				method: ModifyOfferPromotionPackagesAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init?: TypedRequestInit<ModifyOfferPromotionPackagesActionReq, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => ModifyOfferPromotionPackagesActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new ModifyOfferPromotionPackagesActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new ModifyOfferPromotionPackagesActionRes(item))
		const res = await ModifyOfferPromotionPackagesAction.Fetch$(
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
  "name": "Modify offer promotion packages",
  "url": "https://api.{environment}/sale/offers/{offerId}/promo-options-modification",
  "method": "post",
  "description": "Use this resource to modify offer promotion packages. Read more: PL / EN.",
  "in": {
    "fields": [
      {
        "name": "modifications",
        "type": "array",
        "fields": [
          {
            "name": "modificationType",
            "type": "string"
          },
          {
            "name": "packageType",
            "type": "string"
          },
          {
            "name": "packageId",
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
            "name": "modifications",
            "type": "array",
            "fields": [
              {
                "name": "modificationType",
                "type": "string"
              },
              {
                "name": "packageType",
                "type": "string"
              },
              {
                "name": "packageId",
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
  * The base class definition for modifyOfferPromotionPackagesActionReq
  **/
export class ModifyOfferPromotionPackagesActionReq {
		/**
  * 
  * @type {ModifyOfferPromotionPackagesActionReq.Modifications}
  **/
 #modifications : InstanceType<typeof ModifyOfferPromotionPackagesActionReq.Modifications>[]  =  []
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionReq.Modifications}
  **/
get modifications () { return this.#modifications }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionReq.Modifications}
  **/
set modifications (value: InstanceType<typeof ModifyOfferPromotionPackagesActionReq.Modifications>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModifyOfferPromotionPackagesActionReq.Modifications) {
			this.#modifications = value
		} else {
			this.#modifications = value.map(item => new ModifyOfferPromotionPackagesActionReq.Modifications(item))
		}
}
setModifications (value: InstanceType<typeof ModifyOfferPromotionPackagesActionReq.Modifications>[]) {
	this.modifications = value
	return this
}
		/**
  * 
  * @type {ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces}
  **/
 #additionalMarketplaces : InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces>[]  =  []
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value: InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = value.map(item => new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces(item))
		}
}
setAdditionalMarketplaces (value: InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces>[]) {
	this.additionalMarketplaces = value
	return this
}
/**
  * The base class definition for modifications
  **/
static Modifications = class Modifications {
		/**
  * 
  * @type {string}
  **/
 #modificationType : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get modificationType () { return this.#modificationType }
/**
  * 
  * @type {string}
  **/
set modificationType (value: string) {
		this.#modificationType = String(value);
}
setModificationType (value: string) {
	this.modificationType = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #packageType : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get packageType () { return this.#packageType }
/**
  * 
  * @type {string}
  **/
set packageType (value: string) {
		this.#packageType = String(value);
}
setPackageType (value: string) {
	this.packageType = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #packageId : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get packageId () { return this.#packageId }
/**
  * 
  * @type {string}
  **/
set packageId (value: string) {
		this.#packageId = String(value);
}
setPackageId (value: string) {
	this.packageId = value
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
		const d = data as Partial<Modifications>;
			if (d.modificationType !== undefined) { this.modificationType = d.modificationType }
			if (d.packageType !== undefined) { this.packageType = d.packageType }
			if (d.packageId !== undefined) { this.packageId = d.packageId }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				modificationType: this.#modificationType,
				packageType: this.#packageType,
				packageId: this.#packageId,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			modificationType: 'modificationType',
			packageType: 'packageType',
			packageId: 'packageId',
	  }
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionReq.Modifications, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionReqType.ModificationsType) {
		return new ModifyOfferPromotionPackagesActionReq.Modifications(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionReq.Modifications, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionReqType.ModificationsType>) {
		return new ModifyOfferPromotionPackagesActionReq.Modifications(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionReqType.ModificationsType>): InstanceType<typeof ModifyOfferPromotionPackagesActionReq.Modifications> {
		return new ModifyOfferPromotionPackagesActionReq.Modifications ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionReq.Modifications> {
		return new ModifyOfferPromotionPackagesActionReq.Modifications(this.toJSON());
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
  * @type {ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications}
  **/
 #modifications : InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications>[]  =  []
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications}
  **/
get modifications () { return this.#modifications }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications}
  **/
set modifications (value: InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications) {
			this.#modifications = value
		} else {
			this.#modifications = value.map(item => new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications(item))
		}
}
setModifications (value: InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications>[]) {
	this.modifications = value
	return this
}
/**
  * The base class definition for modifications
  **/
static Modifications = class Modifications {
		/**
  * 
  * @type {string}
  **/
 #modificationType : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get modificationType () { return this.#modificationType }
/**
  * 
  * @type {string}
  **/
set modificationType (value: string) {
		this.#modificationType = String(value);
}
setModificationType (value: string) {
	this.modificationType = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #packageType : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get packageType () { return this.#packageType }
/**
  * 
  * @type {string}
  **/
set packageType (value: string) {
		this.#packageType = String(value);
}
setPackageType (value: string) {
	this.packageType = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #packageId : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get packageId () { return this.#packageId }
/**
  * 
  * @type {string}
  **/
set packageId (value: string) {
		this.#packageId = String(value);
}
setPackageId (value: string) {
	this.packageId = value
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
		const d = data as Partial<Modifications>;
			if (d.modificationType !== undefined) { this.modificationType = d.modificationType }
			if (d.packageType !== undefined) { this.packageType = d.packageType }
			if (d.packageId !== undefined) { this.packageId = d.packageId }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				modificationType: this.#modificationType,
				packageType: this.#packageType,
				packageId: this.#packageId,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			modificationType: 'modificationType',
			packageType: 'packageType',
			packageId: 'packageId',
	  }
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType.ModificationsType) {
		return new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType.ModificationsType>) {
		return new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType.ModificationsType>): InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications> {
		return new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications> {
		return new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications(this.toJSON());
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
			if (d.modifications !== undefined) { this.modifications = d.modifications }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				marketplaceId: this.#marketplaceId,
				modifications: this.#modifications,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			marketplaceId: 'marketplaceId',
			modifications$: 'modifications',
get modifications() {
					return withPrefix(
						"additionalMarketplaces.modifications[:i]",
						ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Modifications.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType) {
		return new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType>) {
		return new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType>): InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces> {
		return new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces> {
		return new ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces(this.toJSON());
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
		const d = data as Partial<ModifyOfferPromotionPackagesActionReq>;
			if (d.modifications !== undefined) { this.modifications = d.modifications }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				modifications: this.#modifications,
				additionalMarketplaces: this.#additionalMarketplaces,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			modifications$: 'modifications',
get modifications() {
					return withPrefix(
						"modifications[:i]",
						ModifyOfferPromotionPackagesActionReq.Modifications.Fields
						);
						},
			additionalMarketplaces$: 'additionalMarketplaces',
get additionalMarketplaces() {
					return withPrefix(
						"additionalMarketplaces[:i]",
						ModifyOfferPromotionPackagesActionReq.AdditionalMarketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionReqType) {
		return new ModifyOfferPromotionPackagesActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionReqType>) {
		return new ModifyOfferPromotionPackagesActionReq(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionReqType>): InstanceType<typeof ModifyOfferPromotionPackagesActionReq> {
		return new ModifyOfferPromotionPackagesActionReq ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionReq> {
		return new ModifyOfferPromotionPackagesActionReq(this.toJSON());
	}
}
export abstract class ModifyOfferPromotionPackagesActionReqFactory {
	abstract create(data: unknown): ModifyOfferPromotionPackagesActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for modifyOfferPromotionPackagesActionReq
  **/
	export type ModifyOfferPromotionPackagesActionReqType =  {
			/**
  * 
  * @type {ModifyOfferPromotionPackagesActionReqType.ModificationsType[]}
  **/
 modifications : ModifyOfferPromotionPackagesActionReqType.ModificationsType[];
			/**
  * 
  * @type {ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType[]}
  **/
 additionalMarketplaces : ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModifyOfferPromotionPackagesActionReqType {
	/**
  * The base type definition for modificationsType
  **/
	export type ModificationsType =  {
			/**
  * 
  * @type {string}
  **/
 modificationType : string;
			/**
  * 
  * @type {string}
  **/
 packageType : string;
			/**
  * 
  * @type {string}
  **/
 packageId : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModificationsType {
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
  * @type {ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType.ModificationsType[]}
  **/
 modifications : ModifyOfferPromotionPackagesActionReqType.AdditionalMarketplacesType.ModificationsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdditionalMarketplacesType {
	/**
  * The base type definition for modificationsType
  **/
	export type ModificationsType =  {
			/**
  * 
  * @type {string}
  **/
 modificationType : string;
			/**
  * 
  * @type {string}
  **/
 packageType : string;
			/**
  * 
  * @type {string}
  **/
 packageId : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModificationsType {
}
}
}
/**
  * The base class definition for modifyOfferPromotionPackagesActionRes
  **/
export class ModifyOfferPromotionPackagesActionRes {
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
  * @type {ModifyOfferPromotionPackagesActionRes.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof ModifyOfferPromotionPackagesActionRes.BasePackage>
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionRes.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.BasePackage}
  **/
set basePackage (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyOfferPromotionPackagesActionRes.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new ModifyOfferPromotionPackagesActionRes.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.BasePackage>) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof ModifyOfferPromotionPackagesActionRes.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionRes.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.ExtraPackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModifyOfferPromotionPackagesActionRes.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new ModifyOfferPromotionPackagesActionRes.ExtraPackages(item))
		}
}
setExtraPackages (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.PendingChanges}
  **/
 #pendingChanges ! : InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges>
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionRes.PendingChanges}
  **/
get pendingChanges () { return this.#pendingChanges }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.PendingChanges}
  **/
set pendingChanges (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyOfferPromotionPackagesActionRes.PendingChanges) {
			this.#pendingChanges = value
		} else {
			this.#pendingChanges = new ModifyOfferPromotionPackagesActionRes.PendingChanges(value)
		}
}
setPendingChanges (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges>) {
	this.pendingChanges = value
	return this
}
		/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces}
  **/
 #additionalMarketplaces : InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces>[]  =  []
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = value.map(item => new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces(item))
		}
}
setAdditionalMarketplaces (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces>[]) {
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
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType.BasePackageType) {
		return new ModifyOfferPromotionPackagesActionRes.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType.BasePackageType>) {
		return new ModifyOfferPromotionPackagesActionRes.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType.BasePackageType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.BasePackage> {
		return new ModifyOfferPromotionPackagesActionRes.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.BasePackage> {
		return new ModifyOfferPromotionPackagesActionRes.BasePackage(this.toJSON());
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
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType.ExtraPackagesType) {
		return new ModifyOfferPromotionPackagesActionRes.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType.ExtraPackagesType>) {
		return new ModifyOfferPromotionPackagesActionRes.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType.ExtraPackagesType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.ExtraPackages> {
		return new ModifyOfferPromotionPackagesActionRes.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.ExtraPackages> {
		return new ModifyOfferPromotionPackagesActionRes.ExtraPackages(this.toJSON());
	}
}
/**
  * The base class definition for pendingChanges
  **/
static PendingChanges = class PendingChanges {
		/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage>
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage}
  **/
set basePackage (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage>) {
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
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType) {
		return new ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType>) {
		return new ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage> {
		return new ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage> {
		return new ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage(this.toJSON());
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
			if (!(d.basePackage instanceof ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage)) { this.basePackage = new ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage(d.basePackage || {}) }	
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
						ModifyOfferPromotionPackagesActionRes.PendingChanges.BasePackage.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.PendingChanges, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType.PendingChangesType) {
		return new ModifyOfferPromotionPackagesActionRes.PendingChanges(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.PendingChanges, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType.PendingChangesType>) {
		return new ModifyOfferPromotionPackagesActionRes.PendingChanges(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType.PendingChangesType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges> {
		return new ModifyOfferPromotionPackagesActionRes.PendingChanges ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.PendingChanges> {
		return new ModifyOfferPromotionPackagesActionRes.PendingChanges(this.toJSON());
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
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage>
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage}
  **/
set basePackage (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage>) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(item))
		}
}
setExtraPackages (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges}
  **/
 #pendingChanges ! : InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges>
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges}
  **/
get pendingChanges () { return this.#pendingChanges }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges}
  **/
set pendingChanges (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges) {
			this.#pendingChanges = value
		} else {
			this.#pendingChanges = new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(value)
		}
}
setPendingChanges (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges>) {
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
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType>) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(this.toJSON());
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
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType>) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages(this.toJSON());
	}
}
/**
  * The base class definition for pendingChanges
  **/
static PendingChanges = class PendingChanges {
		/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
 #basePackage ! : InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage>
		/**
  * 
  * @returns {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage}
  **/
set basePackage (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage>) {
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
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType>) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(this.toJSON());
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
			if (!(d.basePackage instanceof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage)) { this.basePackage = new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage(d.basePackage || {}) }	
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
						ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.BasePackage.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType>) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(this.toJSON());
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
			if (!(d.basePackage instanceof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage)) { this.basePackage = new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage(d.basePackage || {}) }	
			if (!(d.pendingChanges instanceof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges)) { this.pendingChanges = new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges(d.pendingChanges || {}) }	
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
						ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"additionalMarketplaces.extraPackages[:i]",
						ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.ExtraPackages.Fields
						);
						},
			pendingChanges$: 'pendingChanges',
get pendingChanges() {
					return withPrefix(
						"additionalMarketplaces.pendingChanges",
						ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.PendingChanges.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType>) {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces> {
		return new ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces(this.toJSON());
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
		const d = data as Partial<ModifyOfferPromotionPackagesActionRes>;
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
		const d = data as Partial<ModifyOfferPromotionPackagesActionRes>;
			if (!(d.basePackage instanceof ModifyOfferPromotionPackagesActionRes.BasePackage)) { this.basePackage = new ModifyOfferPromotionPackagesActionRes.BasePackage(d.basePackage || {}) }	
			if (!(d.pendingChanges instanceof ModifyOfferPromotionPackagesActionRes.PendingChanges)) { this.pendingChanges = new ModifyOfferPromotionPackagesActionRes.PendingChanges(d.pendingChanges || {}) }	
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
						ModifyOfferPromotionPackagesActionRes.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"extraPackages[:i]",
						ModifyOfferPromotionPackagesActionRes.ExtraPackages.Fields
						);
						},
			pendingChanges$: 'pendingChanges',
get pendingChanges() {
					return withPrefix(
						"pendingChanges",
						ModifyOfferPromotionPackagesActionRes.PendingChanges.Fields
						);
						},
			additionalMarketplaces$: 'additionalMarketplaces',
get additionalMarketplaces() {
					return withPrefix(
						"additionalMarketplaces[:i]",
						ModifyOfferPromotionPackagesActionRes.AdditionalMarketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModifyOfferPromotionPackagesActionResType) {
		return new ModifyOfferPromotionPackagesActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ModifyOfferPromotionPackagesActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModifyOfferPromotionPackagesActionResType>) {
		return new ModifyOfferPromotionPackagesActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModifyOfferPromotionPackagesActionResType>): InstanceType<typeof ModifyOfferPromotionPackagesActionRes> {
		return new ModifyOfferPromotionPackagesActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModifyOfferPromotionPackagesActionRes> {
		return new ModifyOfferPromotionPackagesActionRes(this.toJSON());
	}
}
export abstract class ModifyOfferPromotionPackagesActionResFactory {
	abstract create(data: unknown): ModifyOfferPromotionPackagesActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for modifyOfferPromotionPackagesActionRes
  **/
	export type ModifyOfferPromotionPackagesActionResType =  {
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
  * @type {ModifyOfferPromotionPackagesActionResType.BasePackageType}
  **/
 basePackage : ModifyOfferPromotionPackagesActionResType.BasePackageType;
			/**
  * 
  * @type {ModifyOfferPromotionPackagesActionResType.ExtraPackagesType[]}
  **/
 extraPackages : ModifyOfferPromotionPackagesActionResType.ExtraPackagesType[];
			/**
  * 
  * @type {ModifyOfferPromotionPackagesActionResType.PendingChangesType}
  **/
 pendingChanges : ModifyOfferPromotionPackagesActionResType.PendingChangesType;
			/**
  * 
  * @type {ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType[]}
  **/
 additionalMarketplaces : ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModifyOfferPromotionPackagesActionResType {
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
  * @type {ModifyOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType}
  **/
 basePackage : ModifyOfferPromotionPackagesActionResType.PendingChangesType.BasePackageType;
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
  * @type {ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType}
  **/
 basePackage : ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.BasePackageType;
			/**
  * 
  * @type {ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType[]}
  **/
 extraPackages : ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.ExtraPackagesType[];
			/**
  * 
  * @type {ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType}
  **/
 pendingChanges : ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType;
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
  * @type {ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType}
  **/
 basePackage : ModifyOfferPromotionPackagesActionResType.AdditionalMarketplacesType.PendingChangesType.BasePackageType;
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