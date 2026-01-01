import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get Smart! classification report of the particular offer
*/
	/**
 * GetSmartClassificationReportOfTheParticularOfferAction
 */
export class GetSmartClassificationReportOfTheParticularOfferAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}/smart';
  static NewUrl = (
	qs
  ) => buildUrl(
		GetSmartClassificationReportOfTheParticularOfferAction.URL,
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
			overrideUrl ?? GetSmartClassificationReportOfTheParticularOfferAction.NewUrl(
				qs
			),
			{
				method: GetSmartClassificationReportOfTheParticularOfferAction.Method,
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
				creatorFn: (item) => new GetSmartClassificationReportOfTheParticularOfferActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetSmartClassificationReportOfTheParticularOfferActionRes(item))
		const res = await GetSmartClassificationReportOfTheParticularOfferAction.Fetch$(
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
  "name": "Get Smart! classification report of the particular offer",
  "url": "https://api.{environment}/sale/offers/{offerId}/smart",
  "method": "get",
  "description": "Use this resource to get a full Smart! offer classification report of one of your offers. Please keep in mind you have to meet Smart! seller conditions first - for more details, use GET /sale/smart. To learn more about Smart! offer requirements, see our knowledge base article: PL / EN. Read more: PL / EN.",
  "out": {
    "fields": [
      {
        "name": "scheduledForReclassification",
        "description": "Indicates if offer is queued for reclassification",
        "type": "bool"
      },
      {
        "name": "classification",
        "description": "Offer classification status and last change date",
        "type": "object",
        "fields": [
          {
            "name": "fulfilled",
            "description": "Whether the classification conditions are fulfilled",
            "type": "bool"
          },
          {
            "name": "lastChanged",
            "description": "ISO8601 timestamp of last classification change",
            "type": "string"
          }
        ]
      },
      {
        "name": "smartDeliveryMethods",
        "description": "List of smart delivery method identifiers",
        "type": "array",
        "fields": [
          {
            "name": "id",
            "type": "string"
          }
        ]
      },
      {
        "name": "conditions",
        "description": "List of classification conditions with delivery method checks",
        "type": "array",
        "fields": [
          {
            "name": "code",
            "description": "Condition code identifier",
            "type": "string"
          },
          {
            "name": "name",
            "description": "Human-readable condition name",
            "type": "string"
          },
          {
            "name": "description",
            "description": "Detailed condition description",
            "type": "string"
          },
          {
            "name": "fulfilled",
            "description": "Indicates if this condition is fulfilled",
            "type": "bool"
          },
          {
            "name": "passedDeliveryMethods",
            "description": "Delivery methods that passed validation for this condition",
            "type": "array",
            "fields": [
              {
                "name": "id",
                "type": "string"
              }
            ]
          },
          {
            "name": "failedDeliveryMethods",
            "description": "Delivery methods that failed validation for this condition",
            "type": "array",
            "fields": [
              {
                "name": "id",
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
  * The base class definition for getSmartClassificationReportOfTheParticularOfferActionRes
  **/
export class GetSmartClassificationReportOfTheParticularOfferActionRes {
		/**
  * Indicates if offer is queued for reclassification
  * @type {boolean}
  **/
 #scheduledForReclassification
		/**
  * Indicates if offer is queued for reclassification
  * @returns {boolean}
  **/
get scheduledForReclassification () { return this.#scheduledForReclassification }
/**
  * Indicates if offer is queued for reclassification
  * @type {boolean}
  **/
set scheduledForReclassification (value) {
		this.#scheduledForReclassification = Boolean(value);
}
setScheduledForReclassification (value) {
	this.scheduledForReclassification = value
	return this
}
		/**
  * Offer classification status and last change date
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Classification}
  **/
 #classification
		/**
  * Offer classification status and last change date
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.Classification}
  **/
get classification () { return this.#classification }
/**
  * Offer classification status and last change date
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Classification}
  **/
set classification (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSmartClassificationReportOfTheParticularOfferActionRes.Classification) {
			this.#classification = value
		} else {
			this.#classification = new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification(value)
		}
}
setClassification (value) {
	this.classification = value
	return this
}
		/**
  * List of smart delivery method identifiers
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods}
  **/
 #smartDeliveryMethods  =  []
		/**
  * List of smart delivery method identifiers
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods}
  **/
get smartDeliveryMethods () { return this.#smartDeliveryMethods }
/**
  * List of smart delivery method identifiers
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods}
  **/
set smartDeliveryMethods (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods) {
			this.#smartDeliveryMethods = value
		} else {
			this.#smartDeliveryMethods = value.map(item => new GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods(item))
		}
}
setSmartDeliveryMethods (value) {
	this.smartDeliveryMethods = value
	return this
}
		/**
  * List of classification conditions with delivery method checks
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions}
  **/
 #conditions  =  []
		/**
  * List of classification conditions with delivery method checks
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions}
  **/
get conditions () { return this.#conditions }
/**
  * List of classification conditions with delivery method checks
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions}
  **/
set conditions (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions) {
			this.#conditions = value
		} else {
			this.#conditions = value.map(item => new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions(item))
		}
}
setConditions (value) {
	this.conditions = value
	return this
}
/**
  * The base class definition for classification
  **/
static Classification = class Classification {
		/**
  * Whether the classification conditions are fulfilled
  * @type {boolean}
  **/
 #fulfilled
		/**
  * Whether the classification conditions are fulfilled
  * @returns {boolean}
  **/
get fulfilled () { return this.#fulfilled }
/**
  * Whether the classification conditions are fulfilled
  * @type {boolean}
  **/
set fulfilled (value) {
		this.#fulfilled = Boolean(value);
}
setFulfilled (value) {
	this.fulfilled = value
	return this
}
		/**
  * ISO8601 timestamp of last classification change
  * @type {string}
  **/
 #lastChanged  =  ""
		/**
  * ISO8601 timestamp of last classification change
  * @returns {string}
  **/
get lastChanged () { return this.#lastChanged }
/**
  * ISO8601 timestamp of last classification change
  * @type {string}
  **/
set lastChanged (value) {
		this.#lastChanged = String(value);
}
setLastChanged (value) {
	this.lastChanged = value
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
			if (d.fulfilled !== undefined) { this.fulfilled = d.fulfilled }
			if (d.lastChanged !== undefined) { this.lastChanged = d.lastChanged }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				fulfilled: this.#fulfilled,
				lastChanged: this.#lastChanged,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			fulfilled: 'fulfilled',
			lastChanged: 'lastChanged',
	  }
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Classification, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Classification, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification(this.toJSON());
	}
}
/**
  * The base class definition for smartDeliveryMethods
  **/
static SmartDeliveryMethods = class SmartDeliveryMethods {
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
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods(this.toJSON());
	}
}
/**
  * The base class definition for conditions
  **/
static Conditions = class Conditions {
		/**
  * Condition code identifier
  * @type {string}
  **/
 #code  =  ""
		/**
  * Condition code identifier
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * Condition code identifier
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
  * Human-readable condition name
  * @type {string}
  **/
 #name  =  ""
		/**
  * Human-readable condition name
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * Human-readable condition name
  * @type {string}
  **/
set name (value) {
		this.#name = String(value);
}
setName (value) {
	this.name = value
	return this
}
		/**
  * Detailed condition description
  * @type {string}
  **/
 #description  =  ""
		/**
  * Detailed condition description
  * @returns {string}
  **/
get description () { return this.#description }
/**
  * Detailed condition description
  * @type {string}
  **/
set description (value) {
		this.#description = String(value);
}
setDescription (value) {
	this.description = value
	return this
}
		/**
  * Indicates if this condition is fulfilled
  * @type {boolean}
  **/
 #fulfilled
		/**
  * Indicates if this condition is fulfilled
  * @returns {boolean}
  **/
get fulfilled () { return this.#fulfilled }
/**
  * Indicates if this condition is fulfilled
  * @type {boolean}
  **/
set fulfilled (value) {
		this.#fulfilled = Boolean(value);
}
setFulfilled (value) {
	this.fulfilled = value
	return this
}
		/**
  * Delivery methods that passed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods}
  **/
 #passedDeliveryMethods  =  []
		/**
  * Delivery methods that passed validation for this condition
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods}
  **/
get passedDeliveryMethods () { return this.#passedDeliveryMethods }
/**
  * Delivery methods that passed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods}
  **/
set passedDeliveryMethods (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods) {
			this.#passedDeliveryMethods = value
		} else {
			this.#passedDeliveryMethods = value.map(item => new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods(item))
		}
}
setPassedDeliveryMethods (value) {
	this.passedDeliveryMethods = value
	return this
}
		/**
  * Delivery methods that failed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods}
  **/
 #failedDeliveryMethods  =  []
		/**
  * Delivery methods that failed validation for this condition
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods}
  **/
get failedDeliveryMethods () { return this.#failedDeliveryMethods }
/**
  * Delivery methods that failed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods}
  **/
set failedDeliveryMethods (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods) {
			this.#failedDeliveryMethods = value
		} else {
			this.#failedDeliveryMethods = value.map(item => new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods(item))
		}
}
setFailedDeliveryMethods (value) {
	this.failedDeliveryMethods = value
	return this
}
/**
  * The base class definition for passedDeliveryMethods
  **/
static PassedDeliveryMethods = class PassedDeliveryMethods {
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
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods(this.toJSON());
	}
}
/**
  * The base class definition for failedDeliveryMethods
  **/
static FailedDeliveryMethods = class FailedDeliveryMethods {
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
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods(this.toJSON());
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
			if (d.code !== undefined) { this.code = d.code }
			if (d.name !== undefined) { this.name = d.name }
			if (d.description !== undefined) { this.description = d.description }
			if (d.fulfilled !== undefined) { this.fulfilled = d.fulfilled }
			if (d.passedDeliveryMethods !== undefined) { this.passedDeliveryMethods = d.passedDeliveryMethods }
			if (d.failedDeliveryMethods !== undefined) { this.failedDeliveryMethods = d.failedDeliveryMethods }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				code: this.#code,
				name: this.#name,
				description: this.#description,
				fulfilled: this.#fulfilled,
				passedDeliveryMethods: this.#passedDeliveryMethods,
				failedDeliveryMethods: this.#failedDeliveryMethods,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			code: 'code',
			name: 'name',
			description: 'description',
			fulfilled: 'fulfilled',
			passedDeliveryMethods$: 'passedDeliveryMethods',
get passedDeliveryMethods() {
					return withPrefix(
						"conditions.passedDeliveryMethods[:i]",
						GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods.Fields
						);
						},
			failedDeliveryMethods$: 'failedDeliveryMethods',
get failedDeliveryMethods() {
					return withPrefix(
						"conditions.failedDeliveryMethods[:i]",
						GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions(this.toJSON());
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
			if (d.scheduledForReclassification !== undefined) { this.scheduledForReclassification = d.scheduledForReclassification }
			if (d.classification !== undefined) { this.classification = d.classification }
			if (d.smartDeliveryMethods !== undefined) { this.smartDeliveryMethods = d.smartDeliveryMethods }
			if (d.conditions !== undefined) { this.conditions = d.conditions }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.classification instanceof GetSmartClassificationReportOfTheParticularOfferActionRes.Classification)) { this.classification = new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification(d.classification || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				scheduledForReclassification: this.#scheduledForReclassification,
				classification: this.#classification,
				smartDeliveryMethods: this.#smartDeliveryMethods,
				conditions: this.#conditions,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			scheduledForReclassification: 'scheduledForReclassification',
			classification$: 'classification',
get classification() {
					return withPrefix(
						"classification",
						GetSmartClassificationReportOfTheParticularOfferActionRes.Classification.Fields
						);
						},
			smartDeliveryMethods$: 'smartDeliveryMethods',
get smartDeliveryMethods() {
					return withPrefix(
						"smartDeliveryMethods[:i]",
						GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods.Fields
						);
						},
			conditions$: 'conditions',
get conditions() {
					return withPrefix(
						"conditions[:i]",
						GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes(this.toJSON());
	}
}