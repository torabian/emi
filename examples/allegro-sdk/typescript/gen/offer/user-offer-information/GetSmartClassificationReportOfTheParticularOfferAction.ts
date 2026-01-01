import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get Smart! classification report of the particular offer
*/
export type GetSmartClassificationReportOfTheParticularOfferActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * GetSmartClassificationReportOfTheParticularOfferAction
 */
export class GetSmartClassificationReportOfTheParticularOfferAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}/smart';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		GetSmartClassificationReportOfTheParticularOfferAction.URL,
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
		return fetchx<GetSmartClassificationReportOfTheParticularOfferActionRes, unknown, unknown>(
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
		init?: TypedRequestInit<unknown, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => GetSmartClassificationReportOfTheParticularOfferActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
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
 #scheduledForReclassification ! : boolean
		/**
  * Indicates if offer is queued for reclassification
  * @returns {boolean}
  **/
get scheduledForReclassification () { return this.#scheduledForReclassification }
/**
  * Indicates if offer is queued for reclassification
  * @type {boolean}
  **/
set scheduledForReclassification (value: boolean) {
		this.#scheduledForReclassification = Boolean(value);
}
setScheduledForReclassification (value: boolean) {
	this.scheduledForReclassification = value
	return this
}
		/**
  * Offer classification status and last change date
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Classification}
  **/
 #classification ! : InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Classification>
		/**
  * Offer classification status and last change date
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.Classification}
  **/
get classification () { return this.#classification }
/**
  * Offer classification status and last change date
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Classification}
  **/
set classification (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Classification>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSmartClassificationReportOfTheParticularOfferActionRes.Classification) {
			this.#classification = value
		} else {
			this.#classification = new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification(value)
		}
}
setClassification (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Classification>) {
	this.classification = value
	return this
}
		/**
  * List of smart delivery method identifiers
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods}
  **/
 #smartDeliveryMethods : InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods>[]  =  []
		/**
  * List of smart delivery method identifiers
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods}
  **/
get smartDeliveryMethods () { return this.#smartDeliveryMethods }
/**
  * List of smart delivery method identifiers
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods}
  **/
set smartDeliveryMethods (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods>[]) {
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
setSmartDeliveryMethods (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods>[]) {
	this.smartDeliveryMethods = value
	return this
}
		/**
  * List of classification conditions with delivery method checks
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions}
  **/
 #conditions : InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions>[]  =  []
		/**
  * List of classification conditions with delivery method checks
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions}
  **/
get conditions () { return this.#conditions }
/**
  * List of classification conditions with delivery method checks
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions}
  **/
set conditions (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions>[]) {
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
setConditions (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions>[]) {
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
 #fulfilled ! : boolean
		/**
  * Whether the classification conditions are fulfilled
  * @returns {boolean}
  **/
get fulfilled () { return this.#fulfilled }
/**
  * Whether the classification conditions are fulfilled
  * @type {boolean}
  **/
set fulfilled (value: boolean) {
		this.#fulfilled = Boolean(value);
}
setFulfilled (value: boolean) {
	this.fulfilled = value
	return this
}
		/**
  * ISO8601 timestamp of last classification change
  * @type {string}
  **/
 #lastChanged : string  =  ""
		/**
  * ISO8601 timestamp of last classification change
  * @returns {string}
  **/
get lastChanged () { return this.#lastChanged }
/**
  * ISO8601 timestamp of last classification change
  * @type {string}
  **/
set lastChanged (value: string) {
		this.#lastChanged = String(value);
}
setLastChanged (value: string) {
	this.lastChanged = value
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
		const d = data as Partial<Classification>;
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
	static from(possibleDtoObject: GetSmartClassificationReportOfTheParticularOfferActionResType.ClassificationType) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Classification, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.ClassificationType>) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.ClassificationType>): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Classification> {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Classification ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Classification> {
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
		const d = data as Partial<SmartDeliveryMethods>;
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
	static from(possibleDtoObject: GetSmartClassificationReportOfTheParticularOfferActionResType.SmartDeliveryMethodsType) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.SmartDeliveryMethodsType>) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.SmartDeliveryMethodsType>): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods> {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.SmartDeliveryMethods> {
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
 #code : string  =  ""
		/**
  * Condition code identifier
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * Condition code identifier
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
  * Human-readable condition name
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * Human-readable condition name
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * Human-readable condition name
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
		/**
  * Detailed condition description
  * @type {string}
  **/
 #description : string  =  ""
		/**
  * Detailed condition description
  * @returns {string}
  **/
get description () { return this.#description }
/**
  * Detailed condition description
  * @type {string}
  **/
set description (value: string) {
		this.#description = String(value);
}
setDescription (value: string) {
	this.description = value
	return this
}
		/**
  * Indicates if this condition is fulfilled
  * @type {boolean}
  **/
 #fulfilled ! : boolean
		/**
  * Indicates if this condition is fulfilled
  * @returns {boolean}
  **/
get fulfilled () { return this.#fulfilled }
/**
  * Indicates if this condition is fulfilled
  * @type {boolean}
  **/
set fulfilled (value: boolean) {
		this.#fulfilled = Boolean(value);
}
setFulfilled (value: boolean) {
	this.fulfilled = value
	return this
}
		/**
  * Delivery methods that passed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods}
  **/
 #passedDeliveryMethods : InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods>[]  =  []
		/**
  * Delivery methods that passed validation for this condition
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods}
  **/
get passedDeliveryMethods () { return this.#passedDeliveryMethods }
/**
  * Delivery methods that passed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods}
  **/
set passedDeliveryMethods (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods>[]) {
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
setPassedDeliveryMethods (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods>[]) {
	this.passedDeliveryMethods = value
	return this
}
		/**
  * Delivery methods that failed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods}
  **/
 #failedDeliveryMethods : InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods>[]  =  []
		/**
  * Delivery methods that failed validation for this condition
  * @returns {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods}
  **/
get failedDeliveryMethods () { return this.#failedDeliveryMethods }
/**
  * Delivery methods that failed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods}
  **/
set failedDeliveryMethods (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods>[]) {
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
setFailedDeliveryMethods (value: InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods>[]) {
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
		const d = data as Partial<PassedDeliveryMethods>;
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
	static from(possibleDtoObject: GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.PassedDeliveryMethodsType) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.PassedDeliveryMethodsType>) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.PassedDeliveryMethodsType>): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods> {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.PassedDeliveryMethods> {
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
		const d = data as Partial<FailedDeliveryMethods>;
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
	static from(possibleDtoObject: GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.FailedDeliveryMethodsType) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.FailedDeliveryMethodsType>) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.FailedDeliveryMethodsType>): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods> {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods> {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions.FailedDeliveryMethods(this.toJSON());
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
		const d = data as Partial<Conditions>;
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
	static from(possibleDtoObject: GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType>) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType>): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions> {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions> {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes.Conditions(this.toJSON());
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
		const d = data as Partial<GetSmartClassificationReportOfTheParticularOfferActionRes>;
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
		const d = data as Partial<GetSmartClassificationReportOfTheParticularOfferActionRes>;
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
	static from(possibleDtoObject: GetSmartClassificationReportOfTheParticularOfferActionResType) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSmartClassificationReportOfTheParticularOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType>) {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSmartClassificationReportOfTheParticularOfferActionResType>): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes> {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSmartClassificationReportOfTheParticularOfferActionRes> {
		return new GetSmartClassificationReportOfTheParticularOfferActionRes(this.toJSON());
	}
}
export abstract class GetSmartClassificationReportOfTheParticularOfferActionResFactory {
	abstract create(data: unknown): GetSmartClassificationReportOfTheParticularOfferActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for getSmartClassificationReportOfTheParticularOfferActionRes
  **/
	export type GetSmartClassificationReportOfTheParticularOfferActionResType =  {
			/**
  * Indicates if offer is queued for reclassification
  * @type {boolean}
  **/
 scheduledForReclassification : boolean;
			/**
  * Offer classification status and last change date
  * @type {GetSmartClassificationReportOfTheParticularOfferActionResType.ClassificationType}
  **/
 classification : GetSmartClassificationReportOfTheParticularOfferActionResType.ClassificationType;
			/**
  * List of smart delivery method identifiers
  * @type {GetSmartClassificationReportOfTheParticularOfferActionResType.SmartDeliveryMethodsType[]}
  **/
 smartDeliveryMethods : GetSmartClassificationReportOfTheParticularOfferActionResType.SmartDeliveryMethodsType[];
			/**
  * List of classification conditions with delivery method checks
  * @type {GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType[]}
  **/
 conditions : GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetSmartClassificationReportOfTheParticularOfferActionResType {
	/**
  * The base type definition for classificationType
  **/
	export type ClassificationType =  {
			/**
  * Whether the classification conditions are fulfilled
  * @type {boolean}
  **/
 fulfilled : boolean;
			/**
  * ISO8601 timestamp of last classification change
  * @type {string}
  **/
 lastChanged : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassificationType {
}
	/**
  * The base type definition for smartDeliveryMethodsType
  **/
	export type SmartDeliveryMethodsType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SmartDeliveryMethodsType {
}
	/**
  * The base type definition for conditionsType
  **/
	export type ConditionsType =  {
			/**
  * Condition code identifier
  * @type {string}
  **/
 code : string;
			/**
  * Human-readable condition name
  * @type {string}
  **/
 name : string;
			/**
  * Detailed condition description
  * @type {string}
  **/
 description : string;
			/**
  * Indicates if this condition is fulfilled
  * @type {boolean}
  **/
 fulfilled : boolean;
			/**
  * Delivery methods that passed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.PassedDeliveryMethodsType[]}
  **/
 passedDeliveryMethods : GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.PassedDeliveryMethodsType[];
			/**
  * Delivery methods that failed validation for this condition
  * @type {GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.FailedDeliveryMethodsType[]}
  **/
 failedDeliveryMethods : GetSmartClassificationReportOfTheParticularOfferActionResType.ConditionsType.FailedDeliveryMethodsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ConditionsType {
	/**
  * The base type definition for passedDeliveryMethodsType
  **/
	export type PassedDeliveryMethodsType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PassedDeliveryMethodsType {
}
	/**
  * The base type definition for failedDeliveryMethodsType
  **/
	export type FailedDeliveryMethodsType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace FailedDeliveryMethodsType {
}
}
}