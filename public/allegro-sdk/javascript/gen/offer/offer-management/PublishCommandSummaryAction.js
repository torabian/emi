import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Publish command summary
*/
	/**
 * PublishCommandSummaryAction
 */
export class PublishCommandSummaryAction { //
  static URL = 'https://api.{environment}/sale/offer-publication-commands/{commandId}';
  static NewUrl = (
	qs
  ) => buildUrl(
		PublishCommandSummaryAction.URL,
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
		init,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
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
  * @type {PublishCommandSummaryActionRes.TaskCount}
  **/
 #taskCount
		/**
  * 
  * @returns {PublishCommandSummaryActionRes.TaskCount}
  **/
get taskCount () { return this.#taskCount }
/**
  * 
  * @type {PublishCommandSummaryActionRes.TaskCount}
  **/
set taskCount (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof PublishCommandSummaryActionRes.TaskCount) {
			this.#taskCount = value
		} else {
			this.#taskCount = new PublishCommandSummaryActionRes.TaskCount(value)
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
	* Creates an instance of PublishCommandSummaryActionRes.TaskCount, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new PublishCommandSummaryActionRes.TaskCount(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandSummaryActionRes.TaskCount, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new PublishCommandSummaryActionRes.TaskCount(partialDtoObject);
	}
	copyWith(partial) {
		return new PublishCommandSummaryActionRes.TaskCount ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new PublishCommandSummaryActionRes.TaskCount(this.toJSON());
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
	static from(possibleDtoObject) {
		return new PublishCommandSummaryActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of PublishCommandSummaryActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new PublishCommandSummaryActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new PublishCommandSummaryActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new PublishCommandSummaryActionRes(this.toJSON());
	}
}