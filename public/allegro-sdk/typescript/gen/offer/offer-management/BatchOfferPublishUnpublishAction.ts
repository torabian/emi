import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Batch offer publish / unpublish
*/
export type BatchOfferPublishUnpublishActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * BatchOfferPublishUnpublishAction
 */
export class BatchOfferPublishUnpublishAction { //
  static URL = 'https://api.{environment}/sale/offer-publication-commands/{commandId}';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		BatchOfferPublishUnpublishAction.URL,
		 undefined,
		qs
	);
  static Method = 'put';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<BatchOfferPublishUnpublishActionReq, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<BatchOfferPublishUnpublishActionRes, BatchOfferPublishUnpublishActionReq, unknown>(
			overrideUrl ?? BatchOfferPublishUnpublishAction.NewUrl(
				qs
			),
			{
				method: BatchOfferPublishUnpublishAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init?: TypedRequestInit<BatchOfferPublishUnpublishActionReq, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => BatchOfferPublishUnpublishActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new BatchOfferPublishUnpublishActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new BatchOfferPublishUnpublishActionRes(item))
		const res = await BatchOfferPublishUnpublishAction.Fetch$(
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
  "name": "Batch offer publish / unpublish",
  "url": "https://api.{environment}/sale/offer-publication-commands/{commandId}",
  "method": "put",
  "description": "Use this resource to modify multiple offers publication at once. Read more: PL / EN. This resource is rate limited to 250 000 offer changes per hour or 9000 offer changes per minute.",
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
        "name": "publication",
        "type": "object",
        "fields": [
          {
            "name": "action",
            "type": "string"
          },
          {
            "name": "scheduledFor",
            "type": "string"
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
        "name": "createdAt",
        "type": "string"
      },
      {
        "name": "completedAt",
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
  * The base class definition for batchOfferPublishUnpublishActionReq
  **/
export class BatchOfferPublishUnpublishActionReq {
		/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.OfferCriteria}
  **/
 #offerCriteria : InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria>[]  =  []
		/**
  * 
  * @returns {BatchOfferPublishUnpublishActionReq.OfferCriteria}
  **/
get offerCriteria () { return this.#offerCriteria }
/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.OfferCriteria}
  **/
set offerCriteria (value: InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof BatchOfferPublishUnpublishActionReq.OfferCriteria) {
			this.#offerCriteria = value
		} else {
			this.#offerCriteria = value.map(item => new BatchOfferPublishUnpublishActionReq.OfferCriteria(item))
		}
}
setOfferCriteria (value: InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria>[]) {
	this.offerCriteria = value
	return this
}
		/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.Publication}
  **/
 #publication ! : InstanceType<typeof BatchOfferPublishUnpublishActionReq.Publication>
		/**
  * 
  * @returns {BatchOfferPublishUnpublishActionReq.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.Publication}
  **/
set publication (value: InstanceType<typeof BatchOfferPublishUnpublishActionReq.Publication>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof BatchOfferPublishUnpublishActionReq.Publication) {
			this.#publication = value
		} else {
			this.#publication = new BatchOfferPublishUnpublishActionReq.Publication(value)
		}
}
setPublication (value: InstanceType<typeof BatchOfferPublishUnpublishActionReq.Publication>) {
	this.publication = value
	return this
}
/**
  * The base class definition for offerCriteria
  **/
static OfferCriteria = class OfferCriteria {
		/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers}
  **/
 #offers : InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers>[]  =  []
		/**
  * 
  * @returns {BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers}
  **/
get offers () { return this.#offers }
/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers}
  **/
set offers (value: InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers) {
			this.#offers = value
		} else {
			this.#offers = value.map(item => new BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers(item))
		}
}
setOffers (value: InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers>[]) {
	this.offers = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #type : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value: string) {
		this.#type = String(value);
}
setType (value: string) {
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
		const d = data as Partial<Offers>;
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
	* Creates an instance of BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: BatchOfferPublishUnpublishActionReqType.OfferCriteriaType.OffersType) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<BatchOfferPublishUnpublishActionReqType.OfferCriteriaType.OffersType>) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers(partialDtoObject);
	}
	copyWith(partial: PartialDeep<BatchOfferPublishUnpublishActionReqType.OfferCriteriaType.OffersType>): InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers> {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers> {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers(this.toJSON());
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
		const d = data as Partial<OfferCriteria>;
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
						BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers.Fields
						);
						},
			type: 'type',
	  }
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq.OfferCriteria, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: BatchOfferPublishUnpublishActionReqType.OfferCriteriaType) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq.OfferCriteria, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<BatchOfferPublishUnpublishActionReqType.OfferCriteriaType>) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria(partialDtoObject);
	}
	copyWith(partial: PartialDeep<BatchOfferPublishUnpublishActionReqType.OfferCriteriaType>): InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria> {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof BatchOfferPublishUnpublishActionReq.OfferCriteria> {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria(this.toJSON());
	}
}
/**
  * The base class definition for publication
  **/
static Publication = class Publication {
		/**
  * 
  * @type {string}
  **/
 #action : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get action () { return this.#action }
/**
  * 
  * @type {string}
  **/
set action (value: string) {
		this.#action = String(value);
}
setAction (value: string) {
	this.action = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #scheduledFor : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get scheduledFor () { return this.#scheduledFor }
/**
  * 
  * @type {string}
  **/
set scheduledFor (value: string) {
		this.#scheduledFor = String(value);
}
setScheduledFor (value: string) {
	this.scheduledFor = value
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
		const d = data as Partial<Publication>;
			if (d.action !== undefined) { this.action = d.action }
			if (d.scheduledFor !== undefined) { this.scheduledFor = d.scheduledFor }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				action: this.#action,
				scheduledFor: this.#scheduledFor,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			action: 'action',
			scheduledFor: 'scheduledFor',
	  }
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: BatchOfferPublishUnpublishActionReqType.PublicationType) {
		return new BatchOfferPublishUnpublishActionReq.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<BatchOfferPublishUnpublishActionReqType.PublicationType>) {
		return new BatchOfferPublishUnpublishActionReq.Publication(partialDtoObject);
	}
	copyWith(partial: PartialDeep<BatchOfferPublishUnpublishActionReqType.PublicationType>): InstanceType<typeof BatchOfferPublishUnpublishActionReq.Publication> {
		return new BatchOfferPublishUnpublishActionReq.Publication ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof BatchOfferPublishUnpublishActionReq.Publication> {
		return new BatchOfferPublishUnpublishActionReq.Publication(this.toJSON());
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
		const d = data as Partial<BatchOfferPublishUnpublishActionReq>;
			if (d.offerCriteria !== undefined) { this.offerCriteria = d.offerCriteria }
			if (d.publication !== undefined) { this.publication = d.publication }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<BatchOfferPublishUnpublishActionReq>;
			if (!(d.publication instanceof BatchOfferPublishUnpublishActionReq.Publication)) { this.publication = new BatchOfferPublishUnpublishActionReq.Publication(d.publication || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				offerCriteria: this.#offerCriteria,
				publication: this.#publication,
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
						BatchOfferPublishUnpublishActionReq.OfferCriteria.Fields
						);
						},
			publication$: 'publication',
get publication() {
					return withPrefix(
						"publication",
						BatchOfferPublishUnpublishActionReq.Publication.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: BatchOfferPublishUnpublishActionReqType) {
		return new BatchOfferPublishUnpublishActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<BatchOfferPublishUnpublishActionReqType>) {
		return new BatchOfferPublishUnpublishActionReq(partialDtoObject);
	}
	copyWith(partial: PartialDeep<BatchOfferPublishUnpublishActionReqType>): InstanceType<typeof BatchOfferPublishUnpublishActionReq> {
		return new BatchOfferPublishUnpublishActionReq ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof BatchOfferPublishUnpublishActionReq> {
		return new BatchOfferPublishUnpublishActionReq(this.toJSON());
	}
}
export abstract class BatchOfferPublishUnpublishActionReqFactory {
	abstract create(data: unknown): BatchOfferPublishUnpublishActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for batchOfferPublishUnpublishActionReq
  **/
	export type BatchOfferPublishUnpublishActionReqType =  {
			/**
  * 
  * @type {BatchOfferPublishUnpublishActionReqType.OfferCriteriaType[]}
  **/
 offerCriteria : BatchOfferPublishUnpublishActionReqType.OfferCriteriaType[];
			/**
  * 
  * @type {BatchOfferPublishUnpublishActionReqType.PublicationType}
  **/
 publication : BatchOfferPublishUnpublishActionReqType.PublicationType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BatchOfferPublishUnpublishActionReqType {
	/**
  * The base type definition for offerCriteriaType
  **/
	export type OfferCriteriaType =  {
			/**
  * 
  * @type {BatchOfferPublishUnpublishActionReqType.OfferCriteriaType.OffersType[]}
  **/
 offers : BatchOfferPublishUnpublishActionReqType.OfferCriteriaType.OffersType[];
			/**
  * 
  * @type {string}
  **/
 type : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OfferCriteriaType {
	/**
  * The base type definition for offersType
  **/
	export type OffersType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OffersType {
}
}
	/**
  * The base type definition for publicationType
  **/
	export type PublicationType =  {
			/**
  * 
  * @type {string}
  **/
 action : string;
			/**
  * 
  * @type {string}
  **/
 scheduledFor : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PublicationType {
}
}
/**
  * The base class definition for batchOfferPublishUnpublishActionRes
  **/
export class BatchOfferPublishUnpublishActionRes {
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
 #createdAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get createdAt () { return this.#createdAt }
/**
  * 
  * @type {string}
  **/
set createdAt (value: string) {
		this.#createdAt = String(value);
}
setCreatedAt (value: string) {
	this.createdAt = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #completedAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get completedAt () { return this.#completedAt }
/**
  * 
  * @type {string}
  **/
set completedAt (value: string) {
		this.#completedAt = String(value);
}
setCompletedAt (value: string) {
	this.completedAt = value
	return this
}
		/**
  * 
  * @type {BatchOfferPublishUnpublishActionRes.TaskCount}
  **/
 #taskCount ! : InstanceType<typeof BatchOfferPublishUnpublishActionRes.TaskCount>
		/**
  * 
  * @returns {BatchOfferPublishUnpublishActionRes.TaskCount}
  **/
get taskCount () { return this.#taskCount }
/**
  * 
  * @type {BatchOfferPublishUnpublishActionRes.TaskCount}
  **/
set taskCount (value: InstanceType<typeof BatchOfferPublishUnpublishActionRes.TaskCount>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof BatchOfferPublishUnpublishActionRes.TaskCount) {
			this.#taskCount = value
		} else {
			this.#taskCount = new BatchOfferPublishUnpublishActionRes.TaskCount(value)
		}
}
setTaskCount (value: InstanceType<typeof BatchOfferPublishUnpublishActionRes.TaskCount>) {
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
 #failed : number  =  0
		/**
  * 
  * @returns {number}
  **/
get failed () { return this.#failed }
/**
  * 
  * @type {number}
  **/
set failed (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#failed = parsedValue;
		}
}
setFailed (value: number) {
	this.failed = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #success : number  =  0
		/**
  * 
  * @returns {number}
  **/
get success () { return this.#success }
/**
  * 
  * @type {number}
  **/
set success (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#success = parsedValue;
		}
}
setSuccess (value: number) {
	this.success = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #total : number  =  0
		/**
  * 
  * @returns {number}
  **/
get total () { return this.#total }
/**
  * 
  * @type {number}
  **/
set total (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#total = parsedValue;
		}
}
setTotal (value: number) {
	this.total = value
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
		const d = data as Partial<TaskCount>;
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
	* Creates an instance of BatchOfferPublishUnpublishActionRes.TaskCount, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: BatchOfferPublishUnpublishActionResType.TaskCountType) {
		return new BatchOfferPublishUnpublishActionRes.TaskCount(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionRes.TaskCount, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<BatchOfferPublishUnpublishActionResType.TaskCountType>) {
		return new BatchOfferPublishUnpublishActionRes.TaskCount(partialDtoObject);
	}
	copyWith(partial: PartialDeep<BatchOfferPublishUnpublishActionResType.TaskCountType>): InstanceType<typeof BatchOfferPublishUnpublishActionRes.TaskCount> {
		return new BatchOfferPublishUnpublishActionRes.TaskCount ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof BatchOfferPublishUnpublishActionRes.TaskCount> {
		return new BatchOfferPublishUnpublishActionRes.TaskCount(this.toJSON());
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
		const d = data as Partial<BatchOfferPublishUnpublishActionRes>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.createdAt !== undefined) { this.createdAt = d.createdAt }
			if (d.completedAt !== undefined) { this.completedAt = d.completedAt }
			if (d.taskCount !== undefined) { this.taskCount = d.taskCount }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<BatchOfferPublishUnpublishActionRes>;
			if (!(d.taskCount instanceof BatchOfferPublishUnpublishActionRes.TaskCount)) { this.taskCount = new BatchOfferPublishUnpublishActionRes.TaskCount(d.taskCount || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				createdAt: this.#createdAt,
				completedAt: this.#completedAt,
				taskCount: this.#taskCount,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			createdAt: 'createdAt',
			completedAt: 'completedAt',
			taskCount$: 'taskCount',
get taskCount() {
					return withPrefix(
						"taskCount",
						BatchOfferPublishUnpublishActionRes.TaskCount.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: BatchOfferPublishUnpublishActionResType) {
		return new BatchOfferPublishUnpublishActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<BatchOfferPublishUnpublishActionResType>) {
		return new BatchOfferPublishUnpublishActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<BatchOfferPublishUnpublishActionResType>): InstanceType<typeof BatchOfferPublishUnpublishActionRes> {
		return new BatchOfferPublishUnpublishActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof BatchOfferPublishUnpublishActionRes> {
		return new BatchOfferPublishUnpublishActionRes(this.toJSON());
	}
}
export abstract class BatchOfferPublishUnpublishActionResFactory {
	abstract create(data: unknown): BatchOfferPublishUnpublishActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for batchOfferPublishUnpublishActionRes
  **/
	export type BatchOfferPublishUnpublishActionResType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 createdAt : string;
			/**
  * 
  * @type {string}
  **/
 completedAt : string;
			/**
  * 
  * @type {BatchOfferPublishUnpublishActionResType.TaskCountType}
  **/
 taskCount : BatchOfferPublishUnpublishActionResType.TaskCountType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BatchOfferPublishUnpublishActionResType {
	/**
  * The base type definition for taskCountType
  **/
	export type TaskCountType =  {
			/**
  * 
  * @type {number}
  **/
 failed : number;
			/**
  * 
  * @type {number}
  **/
 success : number;
			/**
  * 
  * @type {number}
  **/
 total : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace TaskCountType {
}
}