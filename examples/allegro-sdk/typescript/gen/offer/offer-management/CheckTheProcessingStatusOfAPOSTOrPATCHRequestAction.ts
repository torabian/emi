import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Check the processing status of a POST or PATCH request
*/
export type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction
 */
export class CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction { //
  static URL = 'https://api.{environment}/sale/product-offers/{offerId}/operations/{operationId}';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction.URL,
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
		return fetchx<CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes, unknown, unknown>(
			overrideUrl ?? CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction.NewUrl(
				qs
			),
			{
				method: CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction.Method,
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
				creatorFn?: ((item: unknown) => CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(item))
		const res = await CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction.Fetch$(
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
  "name": "Check the processing status of a POST or PATCH request",
  "url": "https://api.{environment}/sale/product-offers/{offerId}/operations/{operationId}",
  "method": "get",
  "description": "The URI for the resource given by Location header of POST /sale/product-offers and PATCH /sale/product-offers/{offerId}. Use this resource to check processing status of a POST or PATCH request. Read more: PL / EN.",
  "out": {
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
        "name": "operation",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "status",
            "type": "string"
          },
          {
            "name": "startedAt",
            "type": "string"
          }
        ]
      }
    ]
  }
}
}
/**
  * The base class definition for checkTheProcessingStatusOfAPOSTOrPATCHRequestActionRes
  **/
export class CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes {
		/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer}
  **/
 #offer ! : InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer>
		/**
  * 
  * @returns {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer}
  **/
get offer () { return this.#offer }
/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer}
  **/
set offer (value: InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer) {
			this.#offer = value
		} else {
			this.#offer = new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(value)
		}
}
setOffer (value: InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer>) {
	this.offer = value
	return this
}
		/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation}
  **/
 #operation ! : InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation>
		/**
  * 
  * @returns {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation}
  **/
get operation () { return this.#operation }
/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation}
  **/
set operation (value: InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation) {
			this.#operation = value
		} else {
			this.#operation = new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(value)
		}
}
setOperation (value: InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation>) {
	this.operation = value
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
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OfferType) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(possibleDtoObject);
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OfferType>) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OfferType>): InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer> {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer> {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(this.toJSON());
	}
}
/**
  * The base class definition for operation
  **/
static Operation = class Operation {
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
  * @type {string}
  **/
 #startedAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get startedAt () { return this.#startedAt }
/**
  * 
  * @type {string}
  **/
set startedAt (value: string) {
		this.#startedAt = String(value);
}
setStartedAt (value: string) {
	this.startedAt = value
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
		const d = data as Partial<Operation>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.status !== undefined) { this.status = d.status }
			if (d.startedAt !== undefined) { this.startedAt = d.startedAt }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				status: this.#status,
				startedAt: this.#startedAt,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			status: 'status',
			startedAt: 'startedAt',
	  }
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OperationType) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(possibleDtoObject);
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OperationType>) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OperationType>): InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation> {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation> {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(this.toJSON());
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
		const d = data as Partial<CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes>;
			if (d.offer !== undefined) { this.offer = d.offer }
			if (d.operation !== undefined) { this.operation = d.operation }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes>;
			if (!(d.offer instanceof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer)) { this.offer = new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer(d.offer || {}) }	
			if (!(d.operation instanceof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation)) { this.operation = new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation(d.operation || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				offer: this.#offer,
				operation: this.#operation,
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
						"offer",
						CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Offer.Fields
						);
						},
			operation$: 'operation',
get operation() {
					return withPrefix(
						"operation",
						CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes.Operation.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType>) {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType>): InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes> {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes> {
		return new CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes(this.toJSON());
	}
}
export abstract class CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResFactory {
	abstract create(data: unknown): CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for checkTheProcessingStatusOfAPOSTOrPATCHRequestActionRes
  **/
	export type CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType =  {
			/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OfferType}
  **/
 offer : CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OfferType;
			/**
  * 
  * @type {CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OperationType}
  **/
 operation : CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType.OperationType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResType {
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
  * The base type definition for operationType
  **/
	export type OperationType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 status : string;
			/**
  * 
  * @type {string}
  **/
 startedAt : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OperationType {
}
}