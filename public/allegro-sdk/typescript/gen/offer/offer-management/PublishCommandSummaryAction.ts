import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Publish command summary
*/
export type PublishCommandSummaryActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * PublishCommandSummaryAction
 */
export class PublishCommandSummaryAction { //
  static URL = 'https://api.{environment}/sale/offer-publication-commands/{commandId}';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		PublishCommandSummaryAction.URL,
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
		return fetchx<PublishCommandSummaryActionRes, unknown, unknown>(
			overrideUrl ?? PublishCommandSummaryAction.NewUrl(
				qs
			),
			{
				method: PublishCommandSummaryAction.Method,
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
				creatorFn?: ((item: unknown) => PublishCommandSummaryActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new PublishCommandSummaryActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new PublishCommandSummaryActionRes(item))
		const res = await PublishCommandSummaryAction.Fetch$(
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
  "name": "Publish command summary",
  "url": "https://api.{environment}/sale/offer-publication-commands/{commandId}",
  "method": "get",
  "description": "Use this resource to retrieve information about the offer listing statuses.  You will receive a summary with a number of correctly listed offers and errors.  Read more: PL / EN. This resource is rate limited to retrieving information about 270 000 offer changes per minute.",
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
  * The base class definition for publishCommandSummaryActionRes
  **/
export class PublishCommandSummaryActionRes {
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
  * @type {PublishCommandSummaryActionRes.TaskCount}
  **/
 #taskCount ! : InstanceType<typeof PublishCommandSummaryActionRes.TaskCount>
		/**
  * 
  * @returns {PublishCommandSummaryActionRes.TaskCount}
  **/
get taskCount () { return this.#taskCount }
/**
  * 
  * @type {PublishCommandSummaryActionRes.TaskCount}
  **/
set taskCount (value: InstanceType<typeof PublishCommandSummaryActionRes.TaskCount>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof PublishCommandSummaryActionRes.TaskCount) {
			this.#taskCount = value
		} else {
			this.#taskCount = new PublishCommandSummaryActionRes.TaskCount(value)
		}
}
setTaskCount (value: InstanceType<typeof PublishCommandSummaryActionRes.TaskCount>) {
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
	* Creates an instance of PublishCommandSummaryActionRes.TaskCount, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: PublishCommandSummaryActionResType.TaskCountType) {
		return new PublishCommandSummaryActionRes.TaskCount(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandSummaryActionRes.TaskCount, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<PublishCommandSummaryActionResType.TaskCountType>) {
		return new PublishCommandSummaryActionRes.TaskCount(partialDtoObject);
	}
	copyWith(partial: PartialDeep<PublishCommandSummaryActionResType.TaskCountType>): InstanceType<typeof PublishCommandSummaryActionRes.TaskCount> {
		return new PublishCommandSummaryActionRes.TaskCount ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof PublishCommandSummaryActionRes.TaskCount> {
		return new PublishCommandSummaryActionRes.TaskCount(this.toJSON());
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
		const d = data as Partial<PublishCommandSummaryActionRes>;
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
		const d = data as Partial<PublishCommandSummaryActionRes>;
			if (!(d.taskCount instanceof PublishCommandSummaryActionRes.TaskCount)) { this.taskCount = new PublishCommandSummaryActionRes.TaskCount(d.taskCount || {}) }	
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
						PublishCommandSummaryActionRes.TaskCount.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of PublishCommandSummaryActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: PublishCommandSummaryActionResType) {
		return new PublishCommandSummaryActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandSummaryActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<PublishCommandSummaryActionResType>) {
		return new PublishCommandSummaryActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<PublishCommandSummaryActionResType>): InstanceType<typeof PublishCommandSummaryActionRes> {
		return new PublishCommandSummaryActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof PublishCommandSummaryActionRes> {
		return new PublishCommandSummaryActionRes(this.toJSON());
	}
}
export abstract class PublishCommandSummaryActionResFactory {
	abstract create(data: unknown): PublishCommandSummaryActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for publishCommandSummaryActionRes
  **/
	export type PublishCommandSummaryActionResType =  {
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
  * @type {PublishCommandSummaryActionResType.TaskCountType}
  **/
 taskCount : PublishCommandSummaryActionResType.TaskCountType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PublishCommandSummaryActionResType {
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