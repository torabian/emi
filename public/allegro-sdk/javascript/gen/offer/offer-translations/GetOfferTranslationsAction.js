import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get offer translations
*/
	/**
 * GetOfferTranslationsAction
 */
export class GetOfferTranslationsAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}/translations';
  static NewUrl = (
	qs
  ) => buildUrl(
		GetOfferTranslationsAction.URL,
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
			overrideUrl ?? GetOfferTranslationsAction.NewUrl(
				qs
			),
			{
				method: GetOfferTranslationsAction.Method,
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
				creatorFn: (item) => new GetOfferTranslationsActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetOfferTranslationsActionRes(item))
		const res = await GetOfferTranslationsAction.Fetch$(
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
  "name": "Get offer translations",
  "url": "https://api.{environment}/sale/offers/{offerId}/translations",
  "method": "get",
  "description": "Get offer translation for given language or all present. Read more: PL / EN.",
  "out": {
    "fields": [
      {
        "name": "translations",
        "type": "array",
        "fields": [
          {
            "name": "language",
            "type": "string"
          },
          {
            "name": "title",
            "type": "object",
            "fields": [
              {
                "name": "translation",
                "type": "string"
              },
              {
                "name": "type",
                "type": "string"
              }
            ]
          },
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
              },
              {
                "name": "type",
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
                  },
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
  }
}
}
/**
  * The base class definition for getOfferTranslationsActionRes
  **/
export class GetOfferTranslationsActionRes {
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations}
  **/
 #translations  =  []
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations}
  **/
get translations () { return this.#translations }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations}
  **/
set translations (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetOfferTranslationsActionRes.Translations) {
			this.#translations = value
		} else {
			this.#translations = value.map(item => new GetOfferTranslationsActionRes.Translations(item))
		}
}
setTranslations (value) {
	this.translations = value
	return this
}
/**
  * The base class definition for translations
  **/
static Translations = class Translations {
		/**
  * 
  * @type {string}
  **/
 #language  =  ""
		/**
  * 
  * @returns {string}
  **/
get language () { return this.#language }
/**
  * 
  * @type {string}
  **/
set language (value) {
		this.#language = String(value);
}
setLanguage (value) {
	this.language = value
	return this
}
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Title}
  **/
 #title
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Title}
  **/
get title () { return this.#title }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Title}
  **/
set title (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferTranslationsActionRes.Translations.Title) {
			this.#title = value
		} else {
			this.#title = new GetOfferTranslationsActionRes.Translations.Title(value)
		}
}
setTitle (value) {
	this.title = value
	return this
}
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description}
  **/
 #description
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description}
  **/
set description (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferTranslationsActionRes.Translations.Description) {
			this.#description = value
		} else {
			this.#description = new GetOfferTranslationsActionRes.Translations.Description(value)
		}
}
setDescription (value) {
	this.description = value
	return this
}
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.SafetyInformation}
  **/
 #safetyInformation
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.SafetyInformation}
  **/
set safetyInformation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferTranslationsActionRes.Translations.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new GetOfferTranslationsActionRes.Translations.SafetyInformation(value)
		}
}
setSafetyInformation (value) {
	this.safetyInformation = value
	return this
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
			if (d.translation !== undefined) { this.translation = d.translation }
			if (d.type !== undefined) { this.type = d.type }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				translation: this.#translation,
				type: this.#type,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			translation: 'translation',
			type: 'type',
	  }
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Title, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Title(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Title, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Title(partialDtoObject);
	}
	copyWith(partial) {
		return new GetOfferTranslationsActionRes.Translations.Title ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetOfferTranslationsActionRes.Translations.Title(this.toJSON());
	}
}
/**
  * The base class definition for description
  **/
static Description = class Description {
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description.Translation}
  **/
 #translation
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Description.Translation}
  **/
get translation () { return this.#translation }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description.Translation}
  **/
set translation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferTranslationsActionRes.Translations.Description.Translation) {
			this.#translation = value
		} else {
			this.#translation = new GetOfferTranslationsActionRes.Translations.Description.Translation(value)
		}
}
setTranslation (value) {
	this.translation = value
	return this
}
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
/**
  * The base class definition for translation
  **/
static Translation = class Translation {
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections}
  **/
 #sections  =  []
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections}
  **/
set sections (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections) {
			this.#sections = value
		} else {
			this.#sections = value.map(item => new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections(item))
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
  * @type {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items}
  **/
 #items  =  []
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items}
  **/
set items (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items) {
			this.#items = value
		} else {
			this.#items = value.map(item => new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items(item))
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
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items(partialDtoObject);
	}
	copyWith(partial) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items(this.toJSON());
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
						"translations.description.translation.sections.items[:i]",
						GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description.Translation.Sections, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description.Translation.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections(partialDtoObject);
	}
	copyWith(partial) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections(this.toJSON());
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
						"translations.description.translation.sections[:i]",
						GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description.Translation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description.Translation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation(partialDtoObject);
	}
	copyWith(partial) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation(this.toJSON());
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
			if (d.type !== undefined) { this.type = d.type }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.translation instanceof GetOfferTranslationsActionRes.Translations.Description.Translation)) { this.translation = new GetOfferTranslationsActionRes.Translations.Description.Translation(d.translation || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				translation: this.#translation,
				type: this.#type,
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
						"translations.description.translation",
						GetOfferTranslationsActionRes.Translations.Description.Translation.Fields
						);
						},
			type: 'type',
	  }
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.Description(partialDtoObject);
	}
	copyWith(partial) {
		return new GetOfferTranslationsActionRes.Translations.Description ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetOfferTranslationsActionRes.Translations.Description(this.toJSON());
	}
}
/**
  * The base class definition for safetyInformation
  **/
static SafetyInformation = class SafetyInformation {
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.SafetyInformation.Products}
  **/
 #products  =  []
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.SafetyInformation.Products}
  **/
get products () { return this.#products }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.SafetyInformation.Products}
  **/
set products (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetOfferTranslationsActionRes.Translations.SafetyInformation.Products) {
			this.#products = value
		} else {
			this.#products = value.map(item => new GetOfferTranslationsActionRes.Translations.SafetyInformation.Products(item))
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
			if (d.id !== undefined) { this.id = d.id }
			if (d.translation !== undefined) { this.translation = d.translation }
			if (d.type !== undefined) { this.type = d.type }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				translation: this.#translation,
				type: this.#type,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			translation: 'translation',
			type: 'type',
	  }
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.SafetyInformation.Products, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation.Products(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.SafetyInformation.Products, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation.Products(partialDtoObject);
	}
	copyWith(partial) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation.Products ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation.Products(this.toJSON());
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
						"translations.safetyInformation.products[:i]",
						GetOfferTranslationsActionRes.Translations.SafetyInformation.Products.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.SafetyInformation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation(partialDtoObject);
	}
	copyWith(partial) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation(this.toJSON());
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
			if (d.language !== undefined) { this.language = d.language }
			if (d.title !== undefined) { this.title = d.title }
			if (d.description !== undefined) { this.description = d.description }
			if (d.safetyInformation !== undefined) { this.safetyInformation = d.safetyInformation }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.title instanceof GetOfferTranslationsActionRes.Translations.Title)) { this.title = new GetOfferTranslationsActionRes.Translations.Title(d.title || {}) }	
			if (!(d.description instanceof GetOfferTranslationsActionRes.Translations.Description)) { this.description = new GetOfferTranslationsActionRes.Translations.Description(d.description || {}) }	
			if (!(d.safetyInformation instanceof GetOfferTranslationsActionRes.Translations.SafetyInformation)) { this.safetyInformation = new GetOfferTranslationsActionRes.Translations.SafetyInformation(d.safetyInformation || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				language: this.#language,
				title: this.#title,
				description: this.#description,
				safetyInformation: this.#safetyInformation,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			language: 'language',
			title$: 'title',
get title() {
					return withPrefix(
						"translations.title",
						GetOfferTranslationsActionRes.Translations.Title.Fields
						);
						},
			description$: 'description',
get description() {
					return withPrefix(
						"translations.description",
						GetOfferTranslationsActionRes.Translations.Description.Fields
						);
						},
			safetyInformation$: 'safetyInformation',
get safetyInformation() {
					return withPrefix(
						"translations.safetyInformation",
						GetOfferTranslationsActionRes.Translations.SafetyInformation.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetOfferTranslationsActionRes.Translations(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetOfferTranslationsActionRes.Translations(partialDtoObject);
	}
	copyWith(partial) {
		return new GetOfferTranslationsActionRes.Translations ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetOfferTranslationsActionRes.Translations(this.toJSON());
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
			if (d.translations !== undefined) { this.translations = d.translations }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				translations: this.#translations,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			translations$: 'translations',
get translations() {
					return withPrefix(
						"translations[:i]",
						GetOfferTranslationsActionRes.Translations.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetOfferTranslationsActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetOfferTranslationsActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new GetOfferTranslationsActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetOfferTranslationsActionRes(this.toJSON());
	}
}