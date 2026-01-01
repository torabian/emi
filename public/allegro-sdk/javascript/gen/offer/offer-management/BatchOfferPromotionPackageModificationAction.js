import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Batch offer promotion package modification
*/
	/**
 * BatchOfferPromotionPackageModificationAction
 */
export class BatchOfferPromotionPackageModificationAction { //
  static URL = 'https://api.{environment}/sale/offers/promo-options-commands/{commandId}';
  static NewUrl = (
	qs
  ) => buildUrl(
		BatchOfferPromotionPackageModificationAction.URL,
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
			overrideUrl ?? BatchOfferPromotionPackageModificationAction.NewUrl(
				qs
			),
			{
				method: BatchOfferPromotionPackageModificationAction.Method,
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
				creatorFn: (item) => new BatchOfferPromotionPackageModificationActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new BatchOfferPromotionPackageModificationActionRes(item))
		const res = await BatchOfferPromotionPackageModificationAction.Fetch$(
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
  "name": "Batch offer promotion package modification",
  "url": "https://api.{environment}/sale/offers/promo-options-commands/{commandId}",
  "method": "put",
  "description": "Use this resource to modify promotion packages on multiple offers at once. Read more: PL / EN.",
  "in": {
    "fields": [
      {
        "name": "offerCriteria",
        "type": "array",
        "fields": [
          {
            "name": "offers",
            "type": "array",
            "fields": [
              {
                "name": "id",
                "type": "string"
              }
            ]
          },
          {
            "name": "type",
            "type": "string"
          }
        ]
      },
      {
        "name": "modification",
        "type": "object",
        "fields": [
          {
            "name": "basePackage",
            "type": "object",
            "fields": [
              {
                "name": "id",
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
              }
            ]
          },
          {
            "name": "modificationTime",
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
            "name": "modification",
            "type": "object",
            "fields": [
              {
                "name": "basePackage",
                "type": "object",
                "fields": [
                  {
                    "name": "id",
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
                  }
                ]
              },
              {
                "name": "modificationTime",
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
        "name": "taskCount",
        "type": "object",
        "fields": [
          {
            "name": "failed",
            "type": "int"
          },
          {
            "name": "success",
            "type": "int"
          },
          {
            "name": "total",
            "type": "int"
          }
        ]
      }
    ]
  }
}
}
/**
  * The base class definition for batchOfferPromotionPackageModificationActionReq
  **/
export class BatchOfferPromotionPackageModificationActionReq {
		/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.OfferCriteria}
  **/
 #offerCriteria  =  []
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionReq.OfferCriteria}
  **/
get offerCriteria () { return this.#offerCriteria }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.OfferCriteria}
  **/
set offerCriteria (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof BatchOfferPromotionPackageModificationActionReq.OfferCriteria) {
			this.#offerCriteria = value
		} else {
			this.#offerCriteria = value.map(item => new BatchOfferPromotionPackageModificationActionReq.OfferCriteria(item))
		}
}
setOfferCriteria (value) {
	this.offerCriteria = value
	return this
}
		/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.Modification}
  **/
 #modification
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionReq.Modification}
  **/
get modification () { return this.#modification }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.Modification}
  **/
set modification (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof BatchOfferPromotionPackageModificationActionReq.Modification) {
			this.#modification = value
		} else {
			this.#modification = new BatchOfferPromotionPackageModificationActionReq.Modification(value)
		}
}
setModification (value) {
	this.modification = value
	return this
}
		/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces}
  **/
 #additionalMarketplaces  =  []
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = value.map(item => new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces(item))
		}
}
setAdditionalMarketplaces (value) {
	this.additionalMarketplaces = value
	return this
}
/**
  * The base class definition for offerCriteria
  **/
static OfferCriteria = class OfferCriteria {
		/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers}
  **/
 #offers  =  []
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers}
  **/
get offers () { return this.#offers }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers}
  **/
set offers (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers) {
			this.#offers = value
		} else {
			this.#offers = value.map(item => new BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers(item))
		}
}
setOffers (value) {
	this.offers = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #type  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value) {
		this.#type = String(value);
}
setType (value) {
	this.type = value
	return this
}
/**
  * The base class definition for offers
  **/
static Offers = class Offers {
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
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers(this.toJSON());
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
			if (d.offers !== undefined) { this.offers = d.offers }
			if (d.type !== undefined) { this.type = d.type }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				offers: this.#offers,
				type: this.#type,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			offers$: 'offers',
get offers() {
					return withPrefix(
						"offerCriteria.offers[:i]",
						BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Offers.Fields
						);
						},
			type: 'type',
	  }
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.OfferCriteria, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.OfferCriteria(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.OfferCriteria, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.OfferCriteria(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq.OfferCriteria ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq.OfferCriteria(this.toJSON());
	}
}
/**
  * The base class definition for modification
  **/
static Modification = class Modification {
		/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage}
  **/
 #basePackage
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage}
  **/
set basePackage (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage(value)
		}
}
setBasePackage (value) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages}
  **/
 #extraPackages  =  []
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages}
  **/
set extraPackages (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages(item))
		}
}
setExtraPackages (value) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #modificationTime  =  ""
		/**
  * 
  * @returns {string}
  **/
get modificationTime () { return this.#modificationTime }
/**
  * 
  * @type {string}
  **/
set modificationTime (value) {
		this.#modificationTime = String(value);
}
setModificationTime (value) {
	this.modificationTime = value
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
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage(this.toJSON());
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
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages(this.toJSON());
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
			if (d.extraPackages !== undefined) { this.extraPackages = d.extraPackages }
			if (d.modificationTime !== undefined) { this.modificationTime = d.modificationTime }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.basePackage instanceof BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage)) { this.basePackage = new BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage(d.basePackage || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				basePackage: this.#basePackage,
				extraPackages: this.#extraPackages,
				modificationTime: this.#modificationTime,
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
						"modification.basePackage",
						BatchOfferPromotionPackageModificationActionReq.Modification.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"modification.extraPackages[:i]",
						BatchOfferPromotionPackageModificationActionReq.Modification.ExtraPackages.Fields
						);
						},
			modificationTime: 'modificationTime',
	  }
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.Modification, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.Modification(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.Modification, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.Modification(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq.Modification ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq.Modification(this.toJSON());
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
  * @type {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification}
  **/
 #modification
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification}
  **/
get modification () { return this.#modification }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification}
  **/
set modification (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification) {
			this.#modification = value
		} else {
			this.#modification = new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification(value)
		}
}
setModification (value) {
	this.modification = value
	return this
}
/**
  * The base class definition for modification
  **/
static Modification = class Modification {
		/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage}
  **/
 #basePackage
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage}
  **/
set basePackage (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage(value)
		}
}
setBasePackage (value) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages}
  **/
 #extraPackages  =  []
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages}
  **/
set extraPackages (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages(item))
		}
}
setExtraPackages (value) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #modificationTime  =  ""
		/**
  * 
  * @returns {string}
  **/
get modificationTime () { return this.#modificationTime }
/**
  * 
  * @type {string}
  **/
set modificationTime (value) {
		this.#modificationTime = String(value);
}
setModificationTime (value) {
	this.modificationTime = value
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
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage(this.toJSON());
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
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages(this.toJSON());
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
			if (d.extraPackages !== undefined) { this.extraPackages = d.extraPackages }
			if (d.modificationTime !== undefined) { this.modificationTime = d.modificationTime }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.basePackage instanceof BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage)) { this.basePackage = new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage(d.basePackage || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				basePackage: this.#basePackage,
				extraPackages: this.#extraPackages,
				modificationTime: this.#modificationTime,
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
						"additionalMarketplaces.modification.basePackage",
						BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"additionalMarketplaces.modification.extraPackages[:i]",
						BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.ExtraPackages.Fields
						);
						},
			modificationTime: 'modificationTime',
	  }
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification(this.toJSON());
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
			if (d.modification !== undefined) { this.modification = d.modification }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.modification instanceof BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification)) { this.modification = new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification(d.modification || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				marketplaceId: this.#marketplaceId,
				modification: this.#modification,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			marketplaceId: 'marketplaceId',
			modification$: 'modification',
get modification() {
					return withPrefix(
						"additionalMarketplaces.modification",
						BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Modification.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces(this.toJSON());
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
			if (d.offerCriteria !== undefined) { this.offerCriteria = d.offerCriteria }
			if (d.modification !== undefined) { this.modification = d.modification }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.modification instanceof BatchOfferPromotionPackageModificationActionReq.Modification)) { this.modification = new BatchOfferPromotionPackageModificationActionReq.Modification(d.modification || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				offerCriteria: this.#offerCriteria,
				modification: this.#modification,
				additionalMarketplaces: this.#additionalMarketplaces,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			offerCriteria$: 'offerCriteria',
get offerCriteria() {
					return withPrefix(
						"offerCriteria[:i]",
						BatchOfferPromotionPackageModificationActionReq.OfferCriteria.Fields
						);
						},
			modification$: 'modification',
get modification() {
					return withPrefix(
						"modification",
						BatchOfferPromotionPackageModificationActionReq.Modification.Fields
						);
						},
			additionalMarketplaces$: 'additionalMarketplaces',
get additionalMarketplaces() {
					return withPrefix(
						"additionalMarketplaces[:i]",
						BatchOfferPromotionPackageModificationActionReq.AdditionalMarketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionReq(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionReq ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionReq(this.toJSON());
	}
}
/**
  * The base class definition for batchOfferPromotionPackageModificationActionRes
  **/
export class BatchOfferPromotionPackageModificationActionRes {
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
  * @type {BatchOfferPromotionPackageModificationActionRes.TaskCount}
  **/
 #taskCount
		/**
  * 
  * @returns {BatchOfferPromotionPackageModificationActionRes.TaskCount}
  **/
get taskCount () { return this.#taskCount }
/**
  * 
  * @type {BatchOfferPromotionPackageModificationActionRes.TaskCount}
  **/
set taskCount (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof BatchOfferPromotionPackageModificationActionRes.TaskCount) {
			this.#taskCount = value
		} else {
			this.#taskCount = new BatchOfferPromotionPackageModificationActionRes.TaskCount(value)
		}
}
setTaskCount (value) {
	this.taskCount = value
	return this
}
/**
  * The base class definition for taskCount
  **/
static TaskCount = class TaskCount {
		/**
  * 
  * @type {number}
  **/
 #failed  =  0
		/**
  * 
  * @returns {number}
  **/
get failed () { return this.#failed }
/**
  * 
  * @type {number}
  **/
set failed (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#failed = parsedValue;
		}
}
setFailed (value) {
	this.failed = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #success  =  0
		/**
  * 
  * @returns {number}
  **/
get success () { return this.#success }
/**
  * 
  * @type {number}
  **/
set success (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#success = parsedValue;
		}
}
setSuccess (value) {
	this.success = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #total  =  0
		/**
  * 
  * @returns {number}
  **/
get total () { return this.#total }
/**
  * 
  * @type {number}
  **/
set total (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#total = parsedValue;
		}
}
setTotal (value) {
	this.total = value
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
			if (d.failed !== undefined) { this.failed = d.failed }
			if (d.success !== undefined) { this.success = d.success }
			if (d.total !== undefined) { this.total = d.total }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				failed: this.#failed,
				success: this.#success,
				total: this.#total,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			failed: 'failed',
			success: 'success',
			total: 'total',
	  }
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionRes.TaskCount, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionRes.TaskCount(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionRes.TaskCount, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionRes.TaskCount(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionRes.TaskCount ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionRes.TaskCount(this.toJSON());
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
			if (d.taskCount !== undefined) { this.taskCount = d.taskCount }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.taskCount instanceof BatchOfferPromotionPackageModificationActionRes.TaskCount)) { this.taskCount = new BatchOfferPromotionPackageModificationActionRes.TaskCount(d.taskCount || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				taskCount: this.#taskCount,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			taskCount$: 'taskCount',
get taskCount() {
					return withPrefix(
						"taskCount",
						BatchOfferPromotionPackageModificationActionRes.TaskCount.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPromotionPackageModificationActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPromotionPackageModificationActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPromotionPackageModificationActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPromotionPackageModificationActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPromotionPackageModificationActionRes(this.toJSON());
	}
}