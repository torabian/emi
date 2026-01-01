import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Modification command summary
*/
export type ModificationCommandSummaryActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * ModificationCommandSummaryAction
 */
export class ModificationCommandSummaryAction { //
  static URL = 'https://api.{environment}/sale/offers/promo-options-commands/{commandId}';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		ModificationCommandSummaryAction.URL,
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
		return fetchx<ModificationCommandSummaryActionRes, unknown, unknown>(
			overrideUrl ?? ModificationCommandSummaryAction.NewUrl(
				qs
			),
			{
				method: ModificationCommandSummaryAction.Method,
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
				creatorFn?: ((item: unknown) => ModificationCommandSummaryActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new ModificationCommandSummaryActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new ModificationCommandSummaryActionRes(item))
		const res = await ModificationCommandSummaryAction.Fetch$(
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
  "name": "Modification command summary",
  "url": "https://api.{environment}/sale/offers/promo-options-commands/{commandId}",
  "method": "get",
  "description": "Use this resource to find out how many offers were edited within one {commandId}. You will receive a summary with a number of successfully edited offers and errors. Read more: PL / EN.",
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
  * The base class definition for modificationCommandSummaryActionRes
  **/
export class ModificationCommandSummaryActionRes {
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
  * @type {ModificationCommandSummaryActionRes.TaskCount}
  **/
 #taskCount ! : InstanceType<typeof ModificationCommandSummaryActionRes.TaskCount>
		/**
  * 
  * @returns {ModificationCommandSummaryActionRes.TaskCount}
  **/
get taskCount () { return this.#taskCount }
/**
  * 
  * @type {ModificationCommandSummaryActionRes.TaskCount}
  **/
set taskCount (value: InstanceType<typeof ModificationCommandSummaryActionRes.TaskCount>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ModificationCommandSummaryActionRes.TaskCount) {
			this.#taskCount = value
		} else {
			this.#taskCount = new ModificationCommandSummaryActionRes.TaskCount(value)
		}
}
setTaskCount (value: InstanceType<typeof ModificationCommandSummaryActionRes.TaskCount>) {
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
	* Creates an instance of ModificationCommandSummaryActionRes.TaskCount, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModificationCommandSummaryActionResType.TaskCountType) {
		return new ModificationCommandSummaryActionRes.TaskCount(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandSummaryActionRes.TaskCount, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandSummaryActionResType.TaskCountType>) {
		return new ModificationCommandSummaryActionRes.TaskCount(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandSummaryActionResType.TaskCountType>): InstanceType<typeof ModificationCommandSummaryActionRes.TaskCount> {
		return new ModificationCommandSummaryActionRes.TaskCount ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandSummaryActionRes.TaskCount> {
		return new ModificationCommandSummaryActionRes.TaskCount(this.toJSON());
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
		const d = data as Partial<ModificationCommandSummaryActionRes>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.taskCount !== undefined) { this.taskCount = d.taskCount }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<ModificationCommandSummaryActionRes>;
			if (!(d.taskCount instanceof ModificationCommandSummaryActionRes.TaskCount)) { this.taskCount = new ModificationCommandSummaryActionRes.TaskCount(d.taskCount || {}) }	
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
						ModificationCommandSummaryActionRes.TaskCount.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of ModificationCommandSummaryActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: ModificationCommandSummaryActionResType) {
		return new ModificationCommandSummaryActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of ModificationCommandSummaryActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<ModificationCommandSummaryActionResType>) {
		return new ModificationCommandSummaryActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<ModificationCommandSummaryActionResType>): InstanceType<typeof ModificationCommandSummaryActionRes> {
		return new ModificationCommandSummaryActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof ModificationCommandSummaryActionRes> {
		return new ModificationCommandSummaryActionRes(this.toJSON());
	}
}
export abstract class ModificationCommandSummaryActionResFactory {
	abstract create(data: unknown): ModificationCommandSummaryActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for modificationCommandSummaryActionRes
  **/
	export type ModificationCommandSummaryActionResType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {ModificationCommandSummaryActionResType.TaskCountType}
  **/
 taskCount : ModificationCommandSummaryActionResType.TaskCountType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ModificationCommandSummaryActionResType {
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