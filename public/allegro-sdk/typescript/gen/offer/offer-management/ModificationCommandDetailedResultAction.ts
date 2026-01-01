import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Modification command detailed result
*/
export type ModificationCommandDetailedResultActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * ModificationCommandDetailedResultAction
 */
export class ModificationCommandDetailedResultAction { //
  static URL = 'https://api.{environment}/sale/offers/promo-options-commands/{commandId}/tasks';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		ModificationCommandDetailedResultAction.URL,
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
		return fetchx<ModificationCommandDetailedResultActionRes, unknown, unknown>(
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
		init?: TypedRequestInit<unknown, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => ModificationCommandDetailedResultActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
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
 #tasks : InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks>[]  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Tasks}
  **/
get tasks () { return this.#tasks }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks}
  **/
set tasks (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks>[]) {
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
setTasks (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks>[]) {
	this.tasks = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification}
  **/
 #modification ! : InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification>
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Modification}
  **/
get modification () { return this.#modification }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification}
  **/
set modification (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.Modification) {
			this.#modification = value
		} else {
			this.#modification = new ModificationCommandDetailedResultActionRes.Modification(value)
		}
}
setModification (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification>) {
	this.modification = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces}
  **/
 #additionalMarketplaces : InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces>[]  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces>[]) {
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
setAdditionalMarketplaces (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces>[]) {
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
 #offer ! : InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Offer>
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Tasks.Offer}
  **/
get offer () { return this.#offer }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Offer}
  **/
set offer (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Offer>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.Tasks.Offer) {
			this.#offer = value
		} else {
			this.#offer = new ModificationCommandDetailedResultActionRes.Tasks.Offer(value)
		}
}
setOffer (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Offer>) {
	this.offer = value
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
  * @type {string}
  **/
 #scheduledAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get scheduledAt () { return this.#scheduledAt }
/**
  * 
  * @type {string}
  **/
set scheduledAt (value: string) {
		this.#scheduledAt = String(value);
}
setScheduledAt (value: string) {
	this.scheduledAt = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #finishedAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get finishedAt () { return this.#finishedAt }
/**
  * 
  * @type {string}
  **/
set finishedAt (value: string) {
		this.#finishedAt = String(value);
}
setFinishedAt (value: string) {
	this.finishedAt = value
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
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Errors}
  **/
 #errors : InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors>[]  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Tasks.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Errors}
  **/
set errors (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors>[]) {
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
setErrors (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors>[]) {
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
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Offer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.TasksType.OfferType) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Offer(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Offer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.TasksType.OfferType>) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Offer(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.TasksType.OfferType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Offer> {
		return new ModificationCommandDetailedResultActionRes.Tasks.Offer ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Offer> {
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
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata}
  **/
 #metadata ! : InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata>
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata}
  **/
set metadata (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata(value)
		}
}
setMetadata (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata>) {
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
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.TasksType.ErrorsType.MetadataType) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.TasksType.ErrorsType.MetadataType>) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.TasksType.ErrorsType.MetadataType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata> {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata> {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors.Metadata(this.toJSON());
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.TasksType.ErrorsType) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.TasksType.ErrorsType>) {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.TasksType.ErrorsType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors> {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks.Errors> {
		return new ModificationCommandDetailedResultActionRes.Tasks.Errors(this.toJSON());
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
		const d = data as Partial<Tasks>;
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.TasksType) {
		return new ModificationCommandDetailedResultActionRes.Tasks(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Tasks, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.TasksType>) {
		return new ModificationCommandDetailedResultActionRes.Tasks(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.TasksType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks> {
		return new ModificationCommandDetailedResultActionRes.Tasks ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.Tasks> {
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
 #basePackage ! : InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.BasePackage>
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Modification.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification.BasePackage}
  **/
set basePackage (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.Modification.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new ModificationCommandDetailedResultActionRes.Modification.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.BasePackage>) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.Modification.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.Modification.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.ExtraPackages>[]) {
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
setExtraPackages (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #modificationTime : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get modificationTime () { return this.#modificationTime }
/**
  * 
  * @type {string}
  **/
set modificationTime (value: string) {
		this.#modificationTime = String(value);
}
setModificationTime (value: string) {
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
		const d = data as Partial<BasePackage>;
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.ModificationType.BasePackageType) {
		return new ModificationCommandDetailedResultActionRes.Modification.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Modification.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.ModificationType.BasePackageType>) {
		return new ModificationCommandDetailedResultActionRes.Modification.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.ModificationType.BasePackageType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.BasePackage> {
		return new ModificationCommandDetailedResultActionRes.Modification.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.BasePackage> {
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
		const d = data as Partial<ExtraPackages>;
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.ModificationType.ExtraPackagesType) {
		return new ModificationCommandDetailedResultActionRes.Modification.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Modification.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.ModificationType.ExtraPackagesType>) {
		return new ModificationCommandDetailedResultActionRes.Modification.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.ModificationType.ExtraPackagesType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.ExtraPackages> {
		return new ModificationCommandDetailedResultActionRes.Modification.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification.ExtraPackages> {
		return new ModificationCommandDetailedResultActionRes.Modification.ExtraPackages(this.toJSON());
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
		const d = data as Partial<Modification>;
			if (d.basePackage !== undefined) { this.basePackage = d.basePackage }
			if (d.extraPackages !== undefined) { this.extraPackages = d.extraPackages }
			if (d.modificationTime !== undefined) { this.modificationTime = d.modificationTime }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Modification>;
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.ModificationType) {
		return new ModificationCommandDetailedResultActionRes.Modification(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.Modification, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.ModificationType>) {
		return new ModificationCommandDetailedResultActionRes.Modification(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.ModificationType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification> {
		return new ModificationCommandDetailedResultActionRes.Modification ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.Modification> {
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
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification}
  **/
 #modification ! : InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification>
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification}
  **/
get modification () { return this.#modification }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification}
  **/
set modification (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification) {
			this.#modification = value
		} else {
			this.#modification = new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification(value)
		}
}
setModification (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification>) {
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
 #basePackage ! : InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage>
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage}
  **/
get basePackage () { return this.#basePackage }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage}
  **/
set basePackage (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage) {
			this.#basePackage = value
		} else {
			this.#basePackage = new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage(value)
		}
}
setBasePackage (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage>) {
	this.basePackage = value
	return this
}
		/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages}
  **/
 #extraPackages : InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages>[]  =  []
		/**
  * 
  * @returns {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages}
  **/
get extraPackages () { return this.#extraPackages }
/**
  * 
  * @type {ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages}
  **/
set extraPackages (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages>[]) {
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
setExtraPackages (value: InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages>[]) {
	this.extraPackages = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #modificationTime : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get modificationTime () { return this.#modificationTime }
/**
  * 
  * @type {string}
  **/
set modificationTime (value: string) {
		this.#modificationTime = String(value);
}
setModificationTime (value: string) {
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
		const d = data as Partial<BasePackage>;
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.BasePackageType) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.BasePackageType>) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.BasePackageType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage> {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.BasePackage> {
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
		const d = data as Partial<ExtraPackages>;
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.ExtraPackagesType) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.ExtraPackagesType>) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.ExtraPackagesType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages> {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages> {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification.ExtraPackages(this.toJSON());
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
		const d = data as Partial<Modification>;
			if (d.basePackage !== undefined) { this.basePackage = d.basePackage }
			if (d.extraPackages !== undefined) { this.extraPackages = d.extraPackages }
			if (d.modificationTime !== undefined) { this.modificationTime = d.modificationTime }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Modification>;
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType>) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification> {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification> {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces.Modification(this.toJSON());
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
			if (d.modification !== undefined) { this.modification = d.modification }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<AdditionalMarketplaces>;
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType>) {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType>): InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces> {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes.AdditionalMarketplaces> {
		return new ModificationCommandDetailedResultActionRes.AdditionalMarketplaces(this.toJSON());
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
		const d = data as Partial<ModificationCommandDetailedResultActionRes>;
			if (d.tasks !== undefined) { this.tasks = d.tasks }
			if (d.modification !== undefined) { this.modification = d.modification }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<ModificationCommandDetailedResultActionRes>;
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
	static from(possibleDtoObject: ModificationCommandDetailedResultActionResType) {
		return new ModificationCommandDetailedResultActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandDetailedResultActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandDetailedResultActionResType>) {
		return new ModificationCommandDetailedResultActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandDetailedResultActionResType>): InstanceType<typeof ModificationCommandDetailedResultActionRes> {
		return new ModificationCommandDetailedResultActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandDetailedResultActionRes> {
		return new ModificationCommandDetailedResultActionRes(this.toJSON());
	}
}
export abstract class ModificationCommandDetailedResultActionResFactory {
	abstract create(data: unknown): ModificationCommandDetailedResultActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for modificationCommandDetailedResultActionRes
  **/
	export type ModificationCommandDetailedResultActionResType =  {
			/**
  * 
  * @type {ModificationCommandDetailedResultActionResType.TasksType[]}
  **/
 tasks : ModificationCommandDetailedResultActionResType.TasksType[];
			/**
  * 
  * @type {ModificationCommandDetailedResultActionResType.ModificationType}
  **/
 modification : ModificationCommandDetailedResultActionResType.ModificationType;
			/**
  * 
  * @type {ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType[]}
  **/
 additionalMarketplaces : ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModificationCommandDetailedResultActionResType {
	/**
  * The base type definition for tasksType
  **/
	export type TasksType =  {
			/**
  * 
  * @type {ModificationCommandDetailedResultActionResType.TasksType.OfferType}
  **/
 offer : ModificationCommandDetailedResultActionResType.TasksType.OfferType;
			/**
  * 
  * @type {string}
  **/
 marketplaceId : string;
			/**
  * 
  * @type {string}
  **/
 scheduledAt : string;
			/**
  * 
  * @type {string}
  **/
 finishedAt : string;
			/**
  * 
  * @type {string}
  **/
 status : string;
			/**
  * 
  * @type {ModificationCommandDetailedResultActionResType.TasksType.ErrorsType[]}
  **/
 errors : ModificationCommandDetailedResultActionResType.TasksType.ErrorsType[];
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
  * @type {ModificationCommandDetailedResultActionResType.TasksType.ErrorsType.MetadataType}
  **/
 metadata : ModificationCommandDetailedResultActionResType.TasksType.ErrorsType.MetadataType;
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
	/**
  * The base type definition for modificationType
  **/
	export type ModificationType =  {
			/**
  * 
  * @type {ModificationCommandDetailedResultActionResType.ModificationType.BasePackageType}
  **/
 basePackage : ModificationCommandDetailedResultActionResType.ModificationType.BasePackageType;
			/**
  * 
  * @type {ModificationCommandDetailedResultActionResType.ModificationType.ExtraPackagesType[]}
  **/
 extraPackages : ModificationCommandDetailedResultActionResType.ModificationType.ExtraPackagesType[];
			/**
  * 
  * @type {string}
  **/
 modificationTime : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModificationType {
	/**
  * The base type definition for basePackageType
  **/
	export type BasePackageType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
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
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ExtraPackagesType {
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
  * @type {ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType}
  **/
 modification : ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdditionalMarketplacesType {
	/**
  * The base type definition for modificationType
  **/
	export type ModificationType =  {
			/**
  * 
  * @type {ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.BasePackageType}
  **/
 basePackage : ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.BasePackageType;
			/**
  * 
  * @type {ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.ExtraPackagesType[]}
  **/
 extraPackages : ModificationCommandDetailedResultActionResType.AdditionalMarketplacesType.ModificationType.ExtraPackagesType[];
			/**
  * 
  * @type {string}
  **/
 modificationTime : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModificationType {
	/**
  * The base type definition for basePackageType
  **/
	export type BasePackageType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
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
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ExtraPackagesType {
}
}
}
}