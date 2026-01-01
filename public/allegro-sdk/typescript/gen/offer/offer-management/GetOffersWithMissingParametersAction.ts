import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get offers with missing parameters
*/
export type GetOffersWithMissingParametersActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * GetOffersWithMissingParametersAction
 */
export class GetOffersWithMissingParametersAction { //
  static URL = 'https://api.{environment}/sale/offers/unfilled-parameters';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		GetOffersWithMissingParametersAction.URL,
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
		return fetchx<GetOffersWithMissingParametersActionRes, unknown, unknown>(
			overrideUrl ?? GetOffersWithMissingParametersAction.NewUrl(
				qs
			),
			{
				method: GetOffersWithMissingParametersAction.Method,
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
				creatorFn?: ((item: unknown) => GetOffersWithMissingParametersActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new GetOffersWithMissingParametersActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetOffersWithMissingParametersActionRes(item))
		const res = await GetOffersWithMissingParametersAction.Fetch$(
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
  "name": "Get offers with missing parameters",
  "url": "https://api.{environment}/sale/offers/unfilled-parameters",
  "method": "get",
  "description": "Use this resource to get information about required parameters or parameters scheduled to become required that are not filled in offers. Read more: PL / EN.",
  "out": {
    "fields": [
      {
        "name": "offers",
        "type": "array",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "parameters",
            "type": "array",
            "fields": [
              {
                "name": "id",
                "type": "string"
              }
            ]
          },
          {
            "name": "category",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              }
            ]
          }
        ]
      },
      {
        "name": "count",
        "type": "int"
      },
      {
        "name": "totalCount",
        "type": "int"
      }
    ]
  }
}
}
/**
  * The base class definition for getOffersWithMissingParametersActionRes
  **/
export class GetOffersWithMissingParametersActionRes {
		/**
  * 
  * @type {GetOffersWithMissingParametersActionRes.Offers}
  **/
 #offers : InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers>[]  =  []
		/**
  * 
  * @returns {GetOffersWithMissingParametersActionRes.Offers}
  **/
get offers () { return this.#offers }
/**
  * 
  * @type {GetOffersWithMissingParametersActionRes.Offers}
  **/
set offers (value: InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetOffersWithMissingParametersActionRes.Offers) {
			this.#offers = value
		} else {
			this.#offers = value.map(item => new GetOffersWithMissingParametersActionRes.Offers(item))
		}
}
setOffers (value: InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers>[]) {
	this.offers = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #count : number  =  0
		/**
  * 
  * @returns {number}
  **/
get count () { return this.#count }
/**
  * 
  * @type {number}
  **/
set count (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#count = parsedValue;
		}
}
setCount (value: number) {
	this.count = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #totalCount : number  =  0
		/**
  * 
  * @returns {number}
  **/
get totalCount () { return this.#totalCount }
/**
  * 
  * @type {number}
  **/
set totalCount (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#totalCount = parsedValue;
		}
}
setTotalCount (value: number) {
	this.totalCount = value
	return this
}
/**
  * The base class definition for offers
  **/
static Offers = class Offers {
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
  * @type {GetOffersWithMissingParametersActionRes.Offers.Parameters}
  **/
 #parameters : InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Parameters>[]  =  []
		/**
  * 
  * @returns {GetOffersWithMissingParametersActionRes.Offers.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {GetOffersWithMissingParametersActionRes.Offers.Parameters}
  **/
set parameters (value: InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Parameters>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetOffersWithMissingParametersActionRes.Offers.Parameters) {
			this.#parameters = value
		} else {
			this.#parameters = value.map(item => new GetOffersWithMissingParametersActionRes.Offers.Parameters(item))
		}
}
setParameters (value: InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Parameters>[]) {
	this.parameters = value
	return this
}
		/**
  * 
  * @type {GetOffersWithMissingParametersActionRes.Offers.Category}
  **/
 #category ! : InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Category>
		/**
  * 
  * @returns {GetOffersWithMissingParametersActionRes.Offers.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {GetOffersWithMissingParametersActionRes.Offers.Category}
  **/
set category (value: InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Category>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOffersWithMissingParametersActionRes.Offers.Category) {
			this.#category = value
		} else {
			this.#category = new GetOffersWithMissingParametersActionRes.Offers.Category(value)
		}
}
setCategory (value: InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Category>) {
	this.category = value
	return this
}
/**
  * The base class definition for parameters
  **/
static Parameters = class Parameters {
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
		const d = data as Partial<Parameters>;
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
	* Creates an instance of GetOffersWithMissingParametersActionRes.Offers.Parameters, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOffersWithMissingParametersActionResType.OffersType.ParametersType) {
		return new GetOffersWithMissingParametersActionRes.Offers.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOffersWithMissingParametersActionRes.Offers.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOffersWithMissingParametersActionResType.OffersType.ParametersType>) {
		return new GetOffersWithMissingParametersActionRes.Offers.Parameters(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOffersWithMissingParametersActionResType.OffersType.ParametersType>): InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Parameters> {
		return new GetOffersWithMissingParametersActionRes.Offers.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Parameters> {
		return new GetOffersWithMissingParametersActionRes.Offers.Parameters(this.toJSON());
	}
}
/**
  * The base class definition for category
  **/
static Category = class Category {
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
		const d = data as Partial<Category>;
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
	* Creates an instance of GetOffersWithMissingParametersActionRes.Offers.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOffersWithMissingParametersActionResType.OffersType.CategoryType) {
		return new GetOffersWithMissingParametersActionRes.Offers.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOffersWithMissingParametersActionRes.Offers.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOffersWithMissingParametersActionResType.OffersType.CategoryType>) {
		return new GetOffersWithMissingParametersActionRes.Offers.Category(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOffersWithMissingParametersActionResType.OffersType.CategoryType>): InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Category> {
		return new GetOffersWithMissingParametersActionRes.Offers.Category ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers.Category> {
		return new GetOffersWithMissingParametersActionRes.Offers.Category(this.toJSON());
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
		const d = data as Partial<Offers>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.parameters !== undefined) { this.parameters = d.parameters }
			if (d.category !== undefined) { this.category = d.category }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Offers>;
			if (!(d.category instanceof GetOffersWithMissingParametersActionRes.Offers.Category)) { this.category = new GetOffersWithMissingParametersActionRes.Offers.Category(d.category || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				parameters: this.#parameters,
				category: this.#category,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			parameters$: 'parameters',
get parameters() {
					return withPrefix(
						"offers.parameters[:i]",
						GetOffersWithMissingParametersActionRes.Offers.Parameters.Fields
						);
						},
			category$: 'category',
get category() {
					return withPrefix(
						"offers.category",
						GetOffersWithMissingParametersActionRes.Offers.Category.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOffersWithMissingParametersActionRes.Offers, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOffersWithMissingParametersActionResType.OffersType) {
		return new GetOffersWithMissingParametersActionRes.Offers(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOffersWithMissingParametersActionRes.Offers, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOffersWithMissingParametersActionResType.OffersType>) {
		return new GetOffersWithMissingParametersActionRes.Offers(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOffersWithMissingParametersActionResType.OffersType>): InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers> {
		return new GetOffersWithMissingParametersActionRes.Offers ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOffersWithMissingParametersActionRes.Offers> {
		return new GetOffersWithMissingParametersActionRes.Offers(this.toJSON());
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
		const d = data as Partial<GetOffersWithMissingParametersActionRes>;
			if (d.offers !== undefined) { this.offers = d.offers }
			if (d.count !== undefined) { this.count = d.count }
			if (d.totalCount !== undefined) { this.totalCount = d.totalCount }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				offers: this.#offers,
				count: this.#count,
				totalCount: this.#totalCount,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			offers$: 'offers',
get offers() {
					return withPrefix(
						"offers[:i]",
						GetOffersWithMissingParametersActionRes.Offers.Fields
						);
						},
			count: 'count',
			totalCount: 'totalCount',
	  }
	}
	/**
	* Creates an instance of GetOffersWithMissingParametersActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetOffersWithMissingParametersActionResType) {
		return new GetOffersWithMissingParametersActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOffersWithMissingParametersActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOffersWithMissingParametersActionResType>) {
		return new GetOffersWithMissingParametersActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOffersWithMissingParametersActionResType>): InstanceType<typeof GetOffersWithMissingParametersActionRes> {
		return new GetOffersWithMissingParametersActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOffersWithMissingParametersActionRes> {
		return new GetOffersWithMissingParametersActionRes(this.toJSON());
	}
}
export abstract class GetOffersWithMissingParametersActionResFactory {
	abstract create(data: unknown): GetOffersWithMissingParametersActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for getOffersWithMissingParametersActionRes
  **/
	export type GetOffersWithMissingParametersActionResType =  {
			/**
  * 
  * @type {GetOffersWithMissingParametersActionResType.OffersType[]}
  **/
 offers : GetOffersWithMissingParametersActionResType.OffersType[];
			/**
  * 
  * @type {number}
  **/
 count : number;
			/**
  * 
  * @type {number}
  **/
 totalCount : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetOffersWithMissingParametersActionResType {
	/**
  * The base type definition for offersType
  **/
	export type OffersType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {GetOffersWithMissingParametersActionResType.OffersType.ParametersType[]}
  **/
 parameters : GetOffersWithMissingParametersActionResType.OffersType.ParametersType[];
			/**
  * 
  * @type {GetOffersWithMissingParametersActionResType.OffersType.CategoryType}
  **/
 category : GetOffersWithMissingParametersActionResType.OffersType.CategoryType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OffersType {
	/**
  * The base type definition for parametersType
  **/
	export type ParametersType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ParametersType {
}
	/**
  * The base type definition for categoryType
  **/
	export type CategoryType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CategoryType {
}
}
}