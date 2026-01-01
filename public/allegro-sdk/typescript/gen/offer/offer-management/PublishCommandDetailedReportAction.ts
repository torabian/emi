import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Publish command detailed report
*/
export type PublishCommandDetailedReportActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * PublishCommandDetailedReportAction
 */
export class PublishCommandDetailedReportAction { //
  static URL = 'https://api.{environment}/sale/offer-publication-commands/{commandId}/tasks';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		PublishCommandDetailedReportAction.URL,
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
		return fetchx<PublishCommandDetailedReportActionRes, unknown, unknown>(
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
		init?: TypedRequestInit<unknown, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => PublishCommandDetailedReportActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
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
 #tasks : InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks>[]  =  []
		/**
  * 
  * @returns {PublishCommandDetailedReportActionRes.Tasks}
  **/
get tasks () { return this.#tasks }
/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks}
  **/
set tasks (value: InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks>[]) {
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
setTasks (value: InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks>[]) {
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
 #field : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get field () { return this.#field }
/**
  * 
  * @type {string}
  **/
set field (value: string) {
		this.#field = String(value);
}
setField (value: string) {
	this.field = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #message : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get message () { return this.#message }
/**
  * 
  * @type {string}
  **/
set message (value: string) {
		this.#message = String(value);
}
setMessage (value: string) {
	this.message = value
	return this
}
		/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks.Offer}
  **/
 #offer ! : InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Offer>
		/**
  * 
  * @returns {PublishCommandDetailedReportActionRes.Tasks.Offer}
  **/
get offer () { return this.#offer }
/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks.Offer}
  **/
set offer (value: InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Offer>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof PublishCommandDetailedReportActionRes.Tasks.Offer) {
			this.#offer = value
		} else {
			this.#offer = new PublishCommandDetailedReportActionRes.Tasks.Offer(value)
		}
}
setOffer (value: InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Offer>) {
	this.offer = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #status : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get status () { return this.#status }
/**
  * 
  * @type {string}
  **/
set status (value: string) {
		this.#status = String(value);
}
setStatus (value: string) {
	this.status = value
	return this
}
		/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks.Errors}
  **/
 #errors : InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors>[]  =  []
		/**
  * 
  * @returns {PublishCommandDetailedReportActionRes.Tasks.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks.Errors}
  **/
set errors (value: InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors>[]) {
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
setErrors (value: InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors>[]) {
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
		const d = data as Partial<Offer>;
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
	static from(possibleDtoObject: PublishCommandDetailedReportActionResType.TasksType.OfferType) {
		return new PublishCommandDetailedReportActionRes.Tasks.Offer(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks.Offer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<PublishCommandDetailedReportActionResType.TasksType.OfferType>) {
		return new PublishCommandDetailedReportActionRes.Tasks.Offer(partialDtoObject);
	}
	copyWith(partial: PartialDeep<PublishCommandDetailedReportActionResType.TasksType.OfferType>): InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Offer> {
		return new PublishCommandDetailedReportActionRes.Tasks.Offer ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Offer> {
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
 #code : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * 
  * @type {string}
  **/
set code (value: string) {
		this.#code = String(value);
}
setCode (value: string) {
	this.code = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #details : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get details () { return this.#details }
/**
  * 
  * @type {string}
  **/
set details (value: string) {
		this.#details = String(value);
}
setDetails (value: string) {
	this.details = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #message : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get message () { return this.#message }
/**
  * 
  * @type {string}
  **/
set message (value: string) {
		this.#message = String(value);
}
setMessage (value: string) {
	this.message = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #path : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get path () { return this.#path }
/**
  * 
  * @type {string}
  **/
set path (value: string) {
		this.#path = String(value);
}
setPath (value: string) {
	this.path = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #userMessage : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get userMessage () { return this.#userMessage }
/**
  * 
  * @type {string}
  **/
set userMessage (value: string) {
		this.#userMessage = String(value);
}
setUserMessage (value: string) {
	this.userMessage = value
	return this
}
		/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata}
  **/
 #metadata ! : InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata>
		/**
  * 
  * @returns {PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata}
  **/
set metadata (value: InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata(value)
		}
}
setMetadata (value: InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata>) {
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
 #productId : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get productId () { return this.#productId }
/**
  * 
  * @type {string}
  **/
set productId (value: string) {
		this.#productId = String(value);
}
setProductId (value: string) {
	this.productId = value
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
		const d = data as Partial<Metadata>;
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
	static from(possibleDtoObject: PublishCommandDetailedReportActionResType.TasksType.ErrorsType.MetadataType) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<PublishCommandDetailedReportActionResType.TasksType.ErrorsType.MetadataType>) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial: PartialDeep<PublishCommandDetailedReportActionResType.TasksType.ErrorsType.MetadataType>): InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata> {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata> {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors.Metadata(this.toJSON());
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
		const d = data as Partial<Errors>;
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
		const d = data as Partial<Errors>;
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
	static from(possibleDtoObject: PublishCommandDetailedReportActionResType.TasksType.ErrorsType) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<PublishCommandDetailedReportActionResType.TasksType.ErrorsType>) {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors(partialDtoObject);
	}
	copyWith(partial: PartialDeep<PublishCommandDetailedReportActionResType.TasksType.ErrorsType>): InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors> {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks.Errors> {
		return new PublishCommandDetailedReportActionRes.Tasks.Errors(this.toJSON());
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
		const d = data as Partial<Tasks>;
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
		const d = data as Partial<Tasks>;
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
	static from(possibleDtoObject: PublishCommandDetailedReportActionResType.TasksType) {
		return new PublishCommandDetailedReportActionRes.Tasks(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes.Tasks, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<PublishCommandDetailedReportActionResType.TasksType>) {
		return new PublishCommandDetailedReportActionRes.Tasks(partialDtoObject);
	}
	copyWith(partial: PartialDeep<PublishCommandDetailedReportActionResType.TasksType>): InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks> {
		return new PublishCommandDetailedReportActionRes.Tasks ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof PublishCommandDetailedReportActionRes.Tasks> {
		return new PublishCommandDetailedReportActionRes.Tasks(this.toJSON());
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
		const d = data as Partial<PublishCommandDetailedReportActionRes>;
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
	static from(possibleDtoObject: PublishCommandDetailedReportActionResType) {
		return new PublishCommandDetailedReportActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandDetailedReportActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<PublishCommandDetailedReportActionResType>) {
		return new PublishCommandDetailedReportActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<PublishCommandDetailedReportActionResType>): InstanceType<typeof PublishCommandDetailedReportActionRes> {
		return new PublishCommandDetailedReportActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof PublishCommandDetailedReportActionRes> {
		return new PublishCommandDetailedReportActionRes(this.toJSON());
	}
}
export abstract class PublishCommandDetailedReportActionResFactory {
	abstract create(data: unknown): PublishCommandDetailedReportActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for publishCommandDetailedReportActionRes
  **/
	export type PublishCommandDetailedReportActionResType =  {
			/**
  * 
  * @type {PublishCommandDetailedReportActionResType.TasksType[]}
  **/
 tasks : PublishCommandDetailedReportActionResType.TasksType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PublishCommandDetailedReportActionResType {
	/**
  * The base type definition for tasksType
  **/
	export type TasksType =  {
			/**
  * 
  * @type {string}
  **/
 field : string;
			/**
  * 
  * @type {string}
  **/
 message : string;
			/**
  * 
  * @type {PublishCommandDetailedReportActionResType.TasksType.OfferType}
  **/
 offer : PublishCommandDetailedReportActionResType.TasksType.OfferType;
			/**
  * 
  * @type {string}
  **/
 status : string;
			/**
  * 
  * @type {PublishCommandDetailedReportActionResType.TasksType.ErrorsType[]}
  **/
 errors : PublishCommandDetailedReportActionResType.TasksType.ErrorsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace TasksType {
	/**
  * The base type definition for offerType
  **/
	export type OfferType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OfferType {
}
	/**
  * The base type definition for errorsType
  **/
	export type ErrorsType =  {
			/**
  * 
  * @type {string}
  **/
 code : string;
			/**
  * 
  * @type {string}
  **/
 details : string;
			/**
  * 
  * @type {string}
  **/
 message : string;
			/**
  * 
  * @type {string}
  **/
 path : string;
			/**
  * 
  * @type {string}
  **/
 userMessage : string;
			/**
  * 
  * @type {PublishCommandDetailedReportActionResType.TasksType.ErrorsType.MetadataType}
  **/
 metadata : PublishCommandDetailedReportActionResType.TasksType.ErrorsType.MetadataType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ErrorsType {
	/**
  * The base type definition for metadataType
  **/
	export type MetadataType =  {
			/**
  * 
  * @type {string}
  **/
 productId : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MetadataType {
}
}
}
}