import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Batch offer publish / unpublish
*/
	/**
 * BatchOfferPublishUnpublishAction
 */
export class BatchOfferPublishUnpublishAction { //
  static URL = 'https://api.{environment}/sale/offer-publication-commands/{commandId}';
  static NewUrl = (
	qs
  ) => buildUrl(
		BatchOfferPublishUnpublishAction.URL,
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
		init,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
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
 #offerCriteria  =  []
		/**
  * 
  * @returns {BatchOfferPublishUnpublishActionReq.OfferCriteria}
  **/
get offerCriteria () { return this.#offerCriteria }
/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.OfferCriteria}
  **/
set offerCriteria (value) {
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
setOfferCriteria (value) {
	this.offerCriteria = value
	return this
}
		/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.Publication}
  **/
 #publication
		/**
  * 
  * @returns {BatchOfferPublishUnpublishActionReq.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.Publication}
  **/
set publication (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof BatchOfferPublishUnpublishActionReq.Publication) {
			this.#publication = value
		} else {
			this.#publication = new BatchOfferPublishUnpublishActionReq.Publication(value)
		}
}
setPublication (value) {
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
 #offers  =  []
		/**
  * 
  * @returns {BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers}
  **/
get offers () { return this.#offers }
/**
  * 
  * @type {BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers}
  **/
set offers (value) {
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
	* Creates an instance of BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria.Offers(this.toJSON());
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
	static from(possibleDtoObject) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq.OfferCriteria, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPublishUnpublishActionReq.OfferCriteria ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #action  =  ""
		/**
  * 
  * @returns {string}
  **/
get action () { return this.#action }
/**
  * 
  * @type {string}
  **/
set action (value) {
		this.#action = String(value);
}
setAction (value) {
	this.action = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #scheduledFor  =  ""
		/**
  * 
  * @returns {string}
  **/
get scheduledFor () { return this.#scheduledFor }
/**
  * 
  * @type {string}
  **/
set scheduledFor (value) {
		this.#scheduledFor = String(value);
}
setScheduledFor (value) {
	this.scheduledFor = value
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
	static from(possibleDtoObject) {
		return new BatchOfferPublishUnpublishActionReq.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPublishUnpublishActionReq.Publication(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPublishUnpublishActionReq.Publication ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPublishUnpublishActionReq.Publication(this.toJSON());
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
			if (d.publication !== undefined) { this.publication = d.publication }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
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
	static from(possibleDtoObject) {
		return new BatchOfferPublishUnpublishActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPublishUnpublishActionReq(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPublishUnpublishActionReq ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPublishUnpublishActionReq(this.toJSON());
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
 #createdAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get createdAt () { return this.#createdAt }
/**
  * 
  * @type {string}
  **/
set createdAt (value) {
		this.#createdAt = String(value);
}
setCreatedAt (value) {
	this.createdAt = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #completedAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get completedAt () { return this.#completedAt }
/**
  * 
  * @type {string}
  **/
set completedAt (value) {
		this.#completedAt = String(value);
}
setCompletedAt (value) {
	this.completedAt = value
	return this
}
		/**
  * 
  * @type {BatchOfferPublishUnpublishActionRes.TaskCount}
  **/
 #taskCount
		/**
  * 
  * @returns {BatchOfferPublishUnpublishActionRes.TaskCount}
  **/
get taskCount () { return this.#taskCount }
/**
  * 
  * @type {BatchOfferPublishUnpublishActionRes.TaskCount}
  **/
set taskCount (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof BatchOfferPublishUnpublishActionRes.TaskCount) {
			this.#taskCount = value
		} else {
			this.#taskCount = new BatchOfferPublishUnpublishActionRes.TaskCount(value)
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
	* Creates an instance of BatchOfferPublishUnpublishActionRes.TaskCount, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new BatchOfferPublishUnpublishActionRes.TaskCount(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionRes.TaskCount, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPublishUnpublishActionRes.TaskCount(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPublishUnpublishActionRes.TaskCount ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPublishUnpublishActionRes.TaskCount(this.toJSON());
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
			if (d.createdAt !== undefined) { this.createdAt = d.createdAt }
			if (d.completedAt !== undefined) { this.completedAt = d.completedAt }
			if (d.taskCount !== undefined) { this.taskCount = d.taskCount }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
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
	static from(possibleDtoObject) {
		return new BatchOfferPublishUnpublishActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of BatchOfferPublishUnpublishActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new BatchOfferPublishUnpublishActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new BatchOfferPublishUnpublishActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new BatchOfferPublishUnpublishActionRes(this.toJSON());
	}
}