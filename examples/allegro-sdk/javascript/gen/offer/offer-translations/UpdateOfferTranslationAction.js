import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Update offer translation
*/
	/**
 * UpdateOfferTranslationAction
 */
export class UpdateOfferTranslationAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}/translations/{language}';
  static NewUrl = (
	qs
  ) => buildUrl(
		UpdateOfferTranslationAction.URL,
		 undefined,
		qs
	);
  static Method = 'patch';
	static Fetch$ = async (
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
			overrideUrl ?? UpdateOfferTranslationAction.NewUrl(
				qs
			),
			{
				method: UpdateOfferTranslationAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init,
		{
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
		}
	) => {
		const res = await UpdateOfferTranslationAction.Fetch$(
			qs,
			ctx,
			init,
			overrideUrl,
			);
			return handleFetchResponse(
				res, 
				undefined,
				onMessage,
				init?.signal,
			);
	}
  static Definition = {
  "name": "Update offer translation",
  "url": "https://api.{environment}/sale/offers/{offerId}/translations/{language}",
  "method": "patch",
  "description": "Update manual translation for offer. Read more: PL / EN.",
  "in": {
    "fields": [
      {
        "name": "description",
        "type": "object",
        "fields": [
          {
            "name": "translation",
            "type": "object",
            "fields": [
              {
                "name": "sections",
                "type": "array",
                "fields": [
                  {
                    "name": "items",
                    "type": "array",
                    "fields": [
                      {
                        "name": "type",
                        "type": "string"
                      }
                    ]
                  }
                ]
              }
            ]
          }
        ]
      },
      {
        "name": "title",
        "type": "object",
        "fields": [
          {
            "name": "translation",
            "type": "string"
          }
        ]
      },
      {
        "name": "safetyInformation",
        "type": "object",
        "fields": [
          {
            "name": "products",
            "type": "array",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "translation",
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
  * The base class definition for updateOfferTranslationActionReq
  **/
export class UpdateOfferTranslationActionReq {
		/**
  * 
  * @type {UpdateOfferTranslationActionReq.Description}
  **/
 #description
		/**
  * 
  * @returns {UpdateOfferTranslationActionReq.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {UpdateOfferTranslationActionReq.Description}
  **/
set description (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof UpdateOfferTranslationActionReq.Description) {
			this.#description = value
		} else {
			this.#description = new UpdateOfferTranslationActionReq.Description(value)
		}
}
setDescription (value) {
	this.description = value
	return this
}
		/**
  * 
  * @type {UpdateOfferTranslationActionReq.Title}
  **/
 #title
		/**
  * 
  * @returns {UpdateOfferTranslationActionReq.Title}
  **/
get title () { return this.#title }
/**
  * 
  * @type {UpdateOfferTranslationActionReq.Title}
  **/
set title (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof UpdateOfferTranslationActionReq.Title) {
			this.#title = value
		} else {
			this.#title = new UpdateOfferTranslationActionReq.Title(value)
		}
}
setTitle (value) {
	this.title = value
	return this
}
		/**
  * 
  * @type {UpdateOfferTranslationActionReq.SafetyInformation}
  **/
 #safetyInformation
		/**
  * 
  * @returns {UpdateOfferTranslationActionReq.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {UpdateOfferTranslationActionReq.SafetyInformation}
  **/
set safetyInformation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof UpdateOfferTranslationActionReq.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new UpdateOfferTranslationActionReq.SafetyInformation(value)
		}
}
setSafetyInformation (value) {
	this.safetyInformation = value
	return this
}
/**
  * The base class definition for description
  **/
static Description = class Description {
		/**
  * 
  * @type {UpdateOfferTranslationActionReq.Description.Translation}
  **/
 #translation
		/**
  * 
  * @returns {UpdateOfferTranslationActionReq.Description.Translation}
  **/
get translation () { return this.#translation }
/**
  * 
  * @type {UpdateOfferTranslationActionReq.Description.Translation}
  **/
set translation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof UpdateOfferTranslationActionReq.Description.Translation) {
			this.#translation = value
		} else {
			this.#translation = new UpdateOfferTranslationActionReq.Description.Translation(value)
		}
}
setTranslation (value) {
	this.translation = value
	return this
}
/**
  * The base class definition for translation
  **/
static Translation = class Translation {
		/**
  * 
  * @type {UpdateOfferTranslationActionReq.Description.Translation.Sections}
  **/
 #sections  =  []
		/**
  * 
  * @returns {UpdateOfferTranslationActionReq.Description.Translation.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {UpdateOfferTranslationActionReq.Description.Translation.Sections}
  **/
set sections (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof UpdateOfferTranslationActionReq.Description.Translation.Sections) {
			this.#sections = value
		} else {
			this.#sections = value.map(item => new UpdateOfferTranslationActionReq.Description.Translation.Sections(item))
		}
}
setSections (value) {
	this.sections = value
	return this
}
/**
  * The base class definition for sections
  **/
static Sections = class Sections {
		/**
  * 
  * @type {UpdateOfferTranslationActionReq.Description.Translation.Sections.Items}
  **/
 #items  =  []
		/**
  * 
  * @returns {UpdateOfferTranslationActionReq.Description.Translation.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {UpdateOfferTranslationActionReq.Description.Translation.Sections.Items}
  **/
set items (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof UpdateOfferTranslationActionReq.Description.Translation.Sections.Items) {
			this.#items = value
		} else {
			this.#items = value.map(item => new UpdateOfferTranslationActionReq.Description.Translation.Sections.Items(item))
		}
}
setItems (value) {
	this.items = value
	return this
}
/**
  * The base class definition for items
  **/
static Items = class Items {
		/**
  * 
  * @type {string}
  **/
 #type  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value) {
		this.#type = String(value);
}
setType (value) {
	this.type = value
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
			if (d.type !== undefined) { this.type = d.type }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				type: this.#type,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			type: 'type',
	  }
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Description.Translation.Sections.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new UpdateOfferTranslationActionReq.Description.Translation.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Description.Translation.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new UpdateOfferTranslationActionReq.Description.Translation.Sections.Items(partialDtoObject);
	}
	copyWith(partial) {
		return new UpdateOfferTranslationActionReq.Description.Translation.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new UpdateOfferTranslationActionReq.Description.Translation.Sections.Items(this.toJSON());
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
			if (d.items !== undefined) { this.items = d.items }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				items: this.#items,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			items$: 'items',
get items() {
					return withPrefix(
						"description.translation.sections.items[:i]",
						UpdateOfferTranslationActionReq.Description.Translation.Sections.Items.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Description.Translation.Sections, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new UpdateOfferTranslationActionReq.Description.Translation.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Description.Translation.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new UpdateOfferTranslationActionReq.Description.Translation.Sections(partialDtoObject);
	}
	copyWith(partial) {
		return new UpdateOfferTranslationActionReq.Description.Translation.Sections ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new UpdateOfferTranslationActionReq.Description.Translation.Sections(this.toJSON());
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
			if (d.sections !== undefined) { this.sections = d.sections }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				sections: this.#sections,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			sections$: 'sections',
get sections() {
					return withPrefix(
						"description.translation.sections[:i]",
						UpdateOfferTranslationActionReq.Description.Translation.Sections.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Description.Translation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new UpdateOfferTranslationActionReq.Description.Translation(possibleDtoObject);
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Description.Translation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new UpdateOfferTranslationActionReq.Description.Translation(partialDtoObject);
	}
	copyWith(partial) {
		return new UpdateOfferTranslationActionReq.Description.Translation ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new UpdateOfferTranslationActionReq.Description.Translation(this.toJSON());
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
			if (d.translation !== undefined) { this.translation = d.translation }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.translation instanceof UpdateOfferTranslationActionReq.Description.Translation)) { this.translation = new UpdateOfferTranslationActionReq.Description.Translation(d.translation || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				translation: this.#translation,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			translation$: 'translation',
get translation() {
					return withPrefix(
						"description.translation",
						UpdateOfferTranslationActionReq.Description.Translation.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Description, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new UpdateOfferTranslationActionReq.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new UpdateOfferTranslationActionReq.Description(partialDtoObject);
	}
	copyWith(partial) {
		return new UpdateOfferTranslationActionReq.Description ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new UpdateOfferTranslationActionReq.Description(this.toJSON());
	}
}
/**
  * The base class definition for title
  **/
static Title = class Title {
		/**
  * 
  * @type {string}
  **/
 #translation  =  ""
		/**
  * 
  * @returns {string}
  **/
get translation () { return this.#translation }
/**
  * 
  * @type {string}
  **/
set translation (value) {
		this.#translation = String(value);
}
setTranslation (value) {
	this.translation = value
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
			if (d.translation !== undefined) { this.translation = d.translation }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				translation: this.#translation,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			translation: 'translation',
	  }
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Title, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new UpdateOfferTranslationActionReq.Title(possibleDtoObject);
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.Title, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new UpdateOfferTranslationActionReq.Title(partialDtoObject);
	}
	copyWith(partial) {
		return new UpdateOfferTranslationActionReq.Title ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new UpdateOfferTranslationActionReq.Title(this.toJSON());
	}
}
/**
  * The base class definition for safetyInformation
  **/
static SafetyInformation = class SafetyInformation {
		/**
  * 
  * @type {UpdateOfferTranslationActionReq.SafetyInformation.Products}
  **/
 #products  =  []
		/**
  * 
  * @returns {UpdateOfferTranslationActionReq.SafetyInformation.Products}
  **/
get products () { return this.#products }
/**
  * 
  * @type {UpdateOfferTranslationActionReq.SafetyInformation.Products}
  **/
set products (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof UpdateOfferTranslationActionReq.SafetyInformation.Products) {
			this.#products = value
		} else {
			this.#products = value.map(item => new UpdateOfferTranslationActionReq.SafetyInformation.Products(item))
		}
}
setProducts (value) {
	this.products = value
	return this
}
/**
  * The base class definition for products
  **/
static Products = class Products {
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
 #translation  =  ""
		/**
  * 
  * @returns {string}
  **/
get translation () { return this.#translation }
/**
  * 
  * @type {string}
  **/
set translation (value) {
		this.#translation = String(value);
}
setTranslation (value) {
	this.translation = value
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
			if (d.translation !== undefined) { this.translation = d.translation }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				translation: this.#translation,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			translation: 'translation',
	  }
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.SafetyInformation.Products, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new UpdateOfferTranslationActionReq.SafetyInformation.Products(possibleDtoObject);
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.SafetyInformation.Products, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new UpdateOfferTranslationActionReq.SafetyInformation.Products(partialDtoObject);
	}
	copyWith(partial) {
		return new UpdateOfferTranslationActionReq.SafetyInformation.Products ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new UpdateOfferTranslationActionReq.SafetyInformation.Products(this.toJSON());
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
			if (d.products !== undefined) { this.products = d.products }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				products: this.#products,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			products$: 'products',
get products() {
					return withPrefix(
						"safetyInformation.products[:i]",
						UpdateOfferTranslationActionReq.SafetyInformation.Products.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.SafetyInformation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new UpdateOfferTranslationActionReq.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new UpdateOfferTranslationActionReq.SafetyInformation(partialDtoObject);
	}
	copyWith(partial) {
		return new UpdateOfferTranslationActionReq.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new UpdateOfferTranslationActionReq.SafetyInformation(this.toJSON());
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
			if (d.description !== undefined) { this.description = d.description }
			if (d.title !== undefined) { this.title = d.title }
			if (d.safetyInformation !== undefined) { this.safetyInformation = d.safetyInformation }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.description instanceof UpdateOfferTranslationActionReq.Description)) { this.description = new UpdateOfferTranslationActionReq.Description(d.description || {}) }	
			if (!(d.title instanceof UpdateOfferTranslationActionReq.Title)) { this.title = new UpdateOfferTranslationActionReq.Title(d.title || {}) }	
			if (!(d.safetyInformation instanceof UpdateOfferTranslationActionReq.SafetyInformation)) { this.safetyInformation = new UpdateOfferTranslationActionReq.SafetyInformation(d.safetyInformation || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				description: this.#description,
				title: this.#title,
				safetyInformation: this.#safetyInformation,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			description$: 'description',
get description() {
					return withPrefix(
						"description",
						UpdateOfferTranslationActionReq.Description.Fields
						);
						},
			title$: 'title',
get title() {
					return withPrefix(
						"title",
						UpdateOfferTranslationActionReq.Title.Fields
						);
						},
			safetyInformation$: 'safetyInformation',
get safetyInformation() {
					return withPrefix(
						"safetyInformation",
						UpdateOfferTranslationActionReq.SafetyInformation.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new UpdateOfferTranslationActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of UpdateOfferTranslationActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new UpdateOfferTranslationActionReq(partialDtoObject);
	}
	copyWith(partial) {
		return new UpdateOfferTranslationActionReq ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new UpdateOfferTranslationActionReq(this.toJSON());
	}
}