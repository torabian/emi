import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Publish command detailed report
*/
	/**
 * PublishCommandDetailedReportAction
 */
export class PublishCommandDetailedReportAction { //
  static URL = 'https://api.{environment}/sale/offer-publication-commands/{commandId}/tasks';
  static NewUrl = (
	qs
  ) => buildUrl(
		PublishCommandDetailedReportAction.URL,
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
			overrideUrl ?? PublishCommandDetailedReportAction.NewUrl(
				qs
			),
			{
				method: PublishCommandDetailedReportAction.Method,
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
				creatorFn: (item) => new PublishCommandDetailedReportActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new PublishCommandDetailedReportActionRes(item))
		const res = await PublishCommandDetailedReportAction.Fetch$(
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
  "name": "Publish command detailed report",
  "url": "https://api.{environment}/sale/offer-publication-commands/{commandId}/tasks",
  "method": "get",
  "description": "Use this resource to retrieve information about the offer statuses on the site (Defaults: limit = 100, offset = 0). Read more: PL / EN. This resource is rate limited to retrieving information about 270 000 offer changes per minute.",
  "out": {
    "fields": [
      {
        "name": "tasks",
        "type": "array",
        "fields": [
          {
            "name": "field",
            "type": "string"
          },
          {
            "name": "message",
            "type": "string"
          },
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
      }
    ]
  }
}
}
/**
  * The base class definition for publishCommandDetailedReportActionRes
  **/
export class PublishCommandDetailedReportActionRes {
		/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks}
  **/
 #tasks  =  []
		/**
  * 
  * @returns {PublishCommandDetailedReportActionRes.Tasks}
  **/
get tasks () { return this.#tasks }
/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks}
  **/
set tasks (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof PublishCommandDetailedReportActionRes.Tasks) {
			this.#tasks = value
		} else {
			this.#tasks = value.map(item => new PublishCommandDetailedReportActionRes.Tasks(item))
		}
}
setTasks (value) {
	this.tasks = value
	return this
}
/**
  * The base class definition for tasks
  **/
static Tasks = class Tasks {
		/**
  * 
  * @type {string}
  **/
 #field  =  ""
		/**
  * 
  * @returns {string}
  **/
get field () { return this.#field }
/**
  * 
  * @type {string}
  **/
set field (value) {
		this.#field = String(value);
}
setField (value) {
	this.field = value
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
  * @type {PublishCommandDetailedReportActionRes.Tasks.Offer}
  **/
 #offer
		/**
  * 
  * @returns {PublishCommandDetailedReportActionRes.Tasks.Offer}
  **/
get offer () { return this.#offer }
/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks.Offer}
  **/
set offer (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof PublishCommandDetailedReportActionRes.Tasks.Offer) {
			this.#offer = value
		} else {
			this.#offer = new PublishCommandDetailedReportActionRes.Tasks.Offer(value)
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
  * @type {PublishCommandDetailedReportActionRes.Tasks.Errors}
  **/
 #errors  =  []
		/**
  * 
  * @returns {PublishCommandDetailedReportActionRes.Tasks.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks.Errors}
  **/
set errors (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof PublishCommandDetailedReportActionRes.Tasks.Errors) {
			this.#errors = value
		} else {
			this.#errors = value.map(item => new PublishCommandDetailedReportActionRes.Tasks.Errors(item))
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
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks.Offer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new PublishCommandDetailedReportActionRes.Tasks.Offer(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks.Offer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new PublishCommandDetailedReportActionRes.Tasks.Offer(partialDtoObject);
	}
	copyWith(partial) {
		return new PublishCommandDetailedReportActionRes.Tasks.Offer ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new PublishCommandDetailedReportActionRes.Tasks.Offer(this.toJSON());
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
  * @type {PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata}
  **/
 #metadata
		/**
  * 
  * @returns {PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata}
  **/
set metadata (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata(value)
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
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata(this.toJSON());
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
			if (!(d.metadata instanceof PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata)) { this.metadata = new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata(d.metadata || {}) }	
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
						PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks.Errors, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors(partialDtoObject);
	}
	copyWith(partial) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors(this.toJSON());
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
			if (d.field !== undefined) { this.field = d.field }
			if (d.message !== undefined) { this.message = d.message }
			if (d.offer !== undefined) { this.offer = d.offer }
			if (d.status !== undefined) { this.status = d.status }
			if (d.errors !== undefined) { this.errors = d.errors }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.offer instanceof PublishCommandDetailedReportActionRes.Tasks.Offer)) { this.offer = new PublishCommandDetailedReportActionRes.Tasks.Offer(d.offer || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				field: this.#field,
				message: this.#message,
				offer: this.#offer,
				status: this.#status,
				errors: this.#errors,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			field: 'field',
			message: 'message',
			offer$: 'offer',
get offer() {
					return withPrefix(
						"tasks.offer",
						PublishCommandDetailedReportActionRes.Tasks.Offer.Fields
						);
						},
			status: 'status',
			errors$: 'errors',
get errors() {
					return withPrefix(
						"tasks.errors[:i]",
						PublishCommandDetailedReportActionRes.Tasks.Errors.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new PublishCommandDetailedReportActionRes.Tasks(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new PublishCommandDetailedReportActionRes.Tasks(partialDtoObject);
	}
	copyWith(partial) {
		return new PublishCommandDetailedReportActionRes.Tasks ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new PublishCommandDetailedReportActionRes.Tasks(this.toJSON());
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
			if (d.tasks !== undefined) { this.tasks = d.tasks }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				tasks: this.#tasks,
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
						PublishCommandDetailedReportActionRes.Tasks.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new PublishCommandDetailedReportActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new PublishCommandDetailedReportActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new PublishCommandDetailedReportActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new PublishCommandDetailedReportActionRes(this.toJSON());
	}
}