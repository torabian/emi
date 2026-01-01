import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Modification command detailed result
*/
	/**
 * ModificationCommandDetailedResultAction
 */
export class ModificationCommandDetailedResultAction { //
  static URL = 'https://api.{environment}/sale/offers/promo-options-commands/{commandId}/tasks';
  static NewUrl = (
	qs
  ) => buildUrl(
		ModificationCommandDetailedResultAction.URL,
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
			overrideUrl ?? ModificationCommandDetailedResultAction.NewUrl(
				qs
			),
			{
				method: ModificationCommandDetailedResultAction.Method,
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
				creatorFn: (item) => new ModificationCommandDetailedResultActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new ModificationCommandDetailedResultActionRes(item))
		const res = await ModificationCommandDetailedResultAction.Fetch$(
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
  "name": "Modification command detailed result",
  "url": "https://api.{environment}/sale/offers/promo-options-commands/{commandId}/tasks",
  "method": "get",
  "out": {
    "fields": [
      {
        "name": "tasks",
        "type": "array",
        "fields": [
          {
            "name": "offer",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              }
            ]
          },
          {
            "name": "marketplaceId",
            "type": "string"
          },
          {
            "name": "scheduledAt",
            "type": "string"
          },
          {
            "name": "finishedAt",
            "type": "string"
          },
          {
            "name": "status",
            "type": "string"
          },
          {
            "name": "errors",
            "type": "array",
            "fields": [
              {
                "name": "code",
                "type": "string"
              },
              {
                "name": "details",
                "type": "string"
              },
              {
                "name": "message",
                "type": "string"
              },
              {
                "name": "path",
                "type": "string"
              },
              {
                "name": "userMessage",
                "type": "string"
              },
              {
                "name": "metadata",
                "type": "object",
                "fields": [
                  {
                    "name": "productId",
                    "type": "string"
                  }
                ]
              }
            ]
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
  }
}
}
/**
  * The base class definition for modificationCommandDetailedResultActionRes
  **/
export class ModificationCommandDetailedResultActionRes {
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks}
  **/
 #tasks  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Tasks}
  **/
get tasks () { return this.#tasks }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks}
  **/
set tasks (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModificationCommandDetailedResultActionRes.Tasks) {
			this.#tasks = value
		} else {
			this.#tasks = value.map(item => new ModificationCommandDetailedResultActionRes.Tasks(item))
		}
}
setTasks (value) {
	this.tasks = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification}
  **/
 #modification
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Modification}
  **/
get modification () { return this.#modification }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification}
  **/
set modification (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.Modification) {
			this.#modification = value
		} else {
			this.#modification = new ModificationCommandDetailedResultActionRes.Modification(value)
		}
}
setModification (value) {
	this.modification = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces}
  **/
 #additionalMarketplaces  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = value.map(item => new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces(item))
		}
}
setAdditionalMarketplaces (value) {
	this.additionalMarketplaces = value
	return this
}
/**
  * The base class definition for tasks
  **/
static Tasks = class Tasks {
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Offer}
  **/
 #offer
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Tasks.Offer}
  **/
get offer () { return this.#offer }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Offer}
  **/
set offer (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.Tasks.Offer) {
			this.#offer = value
		} else {
			this.#offer = new ModificationCommandDetailedResultActionRes.Tasks.Offer(value)
		}
}
setOffer (value) {
	this.offer = value
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
  * @type {string}
  **/
 #scheduledAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get scheduledAt () { return this.#scheduledAt }
/**
  * 
  * @type {string}
  **/
set scheduledAt (value) {
		this.#scheduledAt = String(value);
}
setScheduledAt (value) {
	this.scheduledAt = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #finishedAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get finishedAt () { return this.#finishedAt }
/**
  * 
  * @type {string}
  **/
set finishedAt (value) {
		this.#finishedAt = String(value);
}
setFinishedAt (value) {
	this.finishedAt = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #status  =  ""
		/**
  * 
  * @returns {string}
  **/
get status () { return this.#status }
/**
  * 
  * @type {string}
  **/
set status (value) {
		this.#status = String(value);
}
setStatus (value) {
	this.status = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Errors}
  **/
 #errors  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Tasks.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Errors}
  **/
set errors (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModificationCommandDetailedResultActionRes.Tasks.Errors) {
			this.#errors = value
		} else {
			this.#errors = value.map(item => new ModificationCommandDetailedResultActionRes.Tasks.Errors(item))
		}
}
setErrors (value) {
	this.errors = value
	return this
}
/**
  * The base class definition for offer
  **/
static Offer = class Offer {
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
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Offer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Offer(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Offer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Offer(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Offer ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.Tasks.Offer(this.toJSON());
	}
}
/**
  * The base class definition for errors
  **/
static Errors = class Errors {
		/**
  * 
  * @type {string}
  **/
 #code  =  ""
		/**
  * 
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * 
  * @type {string}
  **/
set code (value) {
		this.#code = String(value);
}
setCode (value) {
	this.code = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #details  =  ""
		/**
  * 
  * @returns {string}
  **/
get details () { return this.#details }
/**
  * 
  * @type {string}
  **/
set details (value) {
		this.#details = String(value);
}
setDetails (value) {
	this.details = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #message  =  ""
		/**
  * 
  * @returns {string}
  **/
get message () { return this.#message }
/**
  * 
  * @type {string}
  **/
set message (value) {
		this.#message = String(value);
}
setMessage (value) {
	this.message = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #path  =  ""
		/**
  * 
  * @returns {string}
  **/
get path () { return this.#path }
/**
  * 
  * @type {string}
  **/
set path (value) {
		this.#path = String(value);
}
setPath (value) {
	this.path = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #userMessage  =  ""
		/**
  * 
  * @returns {string}
  **/
get userMessage () { return this.#userMessage }
/**
  * 
  * @type {string}
  **/
set userMessage (value) {
		this.#userMessage = String(value);
}
setUserMessage (value) {
	this.userMessage = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata}
  **/
 #metadata
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata}
  **/
set metadata (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata(value)
		}
}
setMetadata (value) {
	this.metadata = value
	return this
}
/**
  * The base class definition for metadata
  **/
static Metadata = class Metadata {
		/**
  * 
  * @type {string}
  **/
 #productId  =  ""
		/**
  * 
  * @returns {string}
  **/
get productId () { return this.#productId }
/**
  * 
  * @type {string}
  **/
set productId (value) {
		this.#productId = String(value);
}
setProductId (value) {
	this.productId = value
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
			if (d.productId !== undefined) { this.productId = d.productId }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				productId: this.#productId,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			productId: 'productId',
	  }
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata(this.toJSON());
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
			if (d.code !== undefined) { this.code = d.code }
			if (d.details !== undefined) { this.details = d.details }
			if (d.message !== undefined) { this.message = d.message }
			if (d.path !== undefined) { this.path = d.path }
			if (d.userMessage !== undefined) { this.userMessage = d.userMessage }
			if (d.metadata !== undefined) { this.metadata = d.metadata }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.metadata instanceof ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata)) { this.metadata = new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata(d.metadata || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				code: this.#code,
				details: this.#details,
				message: this.#message,
				path: this.#path,
				userMessage: this.#userMessage,
				metadata: this.#metadata,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			code: 'code',
			details: 'details',
			message: 'message',
			path: 'path',
			userMessage: 'userMessage',
			metadata$: 'metadata',
get metadata() {
					return withPrefix(
						"tasks.errors.metadata",
						ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Errors, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors(this.toJSON());
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
			if (d.offer !== undefined) { this.offer = d.offer }
			if (d.marketplaceId !== undefined) { this.marketplaceId = d.marketplaceId }
			if (d.scheduledAt !== undefined) { this.scheduledAt = d.scheduledAt }
			if (d.finishedAt !== undefined) { this.finishedAt = d.finishedAt }
			if (d.status !== undefined) { this.status = d.status }
			if (d.errors !== undefined) { this.errors = d.errors }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.offer instanceof ModificationCommandDetailedResultActionRes.Tasks.Offer)) { this.offer = new ModificationCommandDetailedResultActionRes.Tasks.Offer(d.offer || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				offer: this.#offer,
				marketplaceId: this.#marketplaceId,
				scheduledAt: this.#scheduledAt,
				finishedAt: this.#finishedAt,
				status: this.#status,
				errors: this.#errors,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			offer$: 'offer',
get offer() {
					return withPrefix(
						"tasks.offer",
						ModificationCommandDetailedResultActionRes.Tasks.Offer.Fields
						);
						},
			marketplaceId: 'marketplaceId',
			scheduledAt: 'scheduledAt',
			finishedAt: 'finishedAt',
			status: 'status',
			errors$: 'errors',
get errors() {
					return withPrefix(
						"tasks.errors[:i]",
						ModificationCommandDetailedResultActionRes.Tasks.Errors.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Tasks(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Tasks(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.Tasks ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.Tasks(this.toJSON());
	}
}
/**
  * The base class definition for modification
  **/
static Modification = class Modification {
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification.BasePackage}
  **/
 #basePackage
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Modification.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification.BasePackage}
  **/
set basePackage (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.Modification.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new ModificationCommandDetailedResultActionRes.Modification.BasePackage(value)
		}
}
setBasePackage (value) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification.ExtraPackages}
  **/
 #extraPackages  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Modification.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification.ExtraPackages}
  **/
set extraPackages (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModificationCommandDetailedResultActionRes.Modification.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new ModificationCommandDetailedResultActionRes.Modification.ExtraPackages(item))
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
	* Creates an instance of ModificationCommandDetailedResultActionRes.Modification.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Modification.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Modification.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Modification.BasePackage(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.Modification.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.Modification.BasePackage(this.toJSON());
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
	* Creates an instance of ModificationCommandDetailedResultActionRes.Modification.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Modification.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Modification.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Modification.ExtraPackages(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.Modification.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.Modification.ExtraPackages(this.toJSON());
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
			if (!(d.basePackage instanceof ModificationCommandDetailedResultActionRes.Modification.BasePackage)) { this.basePackage = new ModificationCommandDetailedResultActionRes.Modification.BasePackage(d.basePackage || {}) }	
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
						ModificationCommandDetailedResultActionRes.Modification.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"modification.extraPackages[:i]",
						ModificationCommandDetailedResultActionRes.Modification.ExtraPackages.Fields
						);
						},
			modificationTime: 'modificationTime',
	  }
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Modification, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Modification(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Modification, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.Modification(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.Modification ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.Modification(this.toJSON());
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
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification}
  **/
 #modification
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification}
  **/
get modification () { return this.#modification }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification}
  **/
set modification (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification) {
			this.#modification = value
		} else {
			this.#modification = new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification(value)
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
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage}
  **/
 #basePackage
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage}
  **/
set basePackage (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage(value)
		}
}
setBasePackage (value) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages}
  **/
 #extraPackages  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages}
  **/
set extraPackages (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages) {
			this.#extraPackages = value
		} else {
			this.#extraPackages = value.map(item => new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages(item))
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
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage(this.toJSON());
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
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages(this.toJSON());
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
			if (!(d.basePackage instanceof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage)) { this.basePackage = new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage(d.basePackage || {}) }	
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
						ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage.Fields
						);
						},
			extraPackages$: 'extraPackages',
get extraPackages() {
					return withPrefix(
						"additionalMarketplaces.modification.extraPackages[:i]",
						ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages.Fields
						);
						},
			modificationTime: 'modificationTime',
	  }
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification(this.toJSON());
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
			if (!(d.modification instanceof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification)) { this.modification = new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification(d.modification || {}) }	
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
						ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces(this.toJSON());
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
			if (d.tasks !== undefined) { this.tasks = d.tasks }
			if (d.modification !== undefined) { this.modification = d.modification }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.modification instanceof ModificationCommandDetailedResultActionRes.Modification)) { this.modification = new ModificationCommandDetailedResultActionRes.Modification(d.modification || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				tasks: this.#tasks,
				modification: this.#modification,
				additionalMarketplaces: this.#additionalMarketplaces,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			tasks$: 'tasks',
get tasks() {
					return withPrefix(
						"tasks[:i]",
						ModificationCommandDetailedResultActionRes.Tasks.Fields
						);
						},
			modification$: 'modification',
get modification() {
					return withPrefix(
						"modification",
						ModificationCommandDetailedResultActionRes.Modification.Fields
						);
						},
			additionalMarketplaces$: 'additionalMarketplaces',
get additionalMarketplaces() {
					return withPrefix(
						"additionalMarketplaces[:i]",
						ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new ModificationCommandDetailedResultActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new ModificationCommandDetailedResultActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new ModificationCommandDetailedResultActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new ModificationCommandDetailedResultActionRes(this.toJSON());
	}
}