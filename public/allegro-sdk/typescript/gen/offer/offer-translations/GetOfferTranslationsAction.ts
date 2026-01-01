import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get offer translations
*/
export type GetOfferTranslationsActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * GetOfferTranslationsAction
 */
export class GetOfferTranslationsAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}/translations';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		GetOfferTranslationsAction.URL,
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
		return fetchx<GetOfferTranslationsActionRes, unknown, unknown>(
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
		init?: TypedRequestInit<unknown, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => GetOfferTranslationsActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
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
 #translations : InstanceType<typeof GetOfferTranslationsActionRes.Translations>[]  =  []
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations}
  **/
get translations () { return this.#translations }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations}
  **/
set translations (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations>[]) {
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
setTranslations (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations>[]) {
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
 #language : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get language () { return this.#language }
/**
  * 
  * @type {string}
  **/
set language (value: string) {
		this.#language = String(value);
}
setLanguage (value: string) {
	this.language = value
	return this
}
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Title}
  **/
 #title ! : InstanceType<typeof GetOfferTranslationsActionRes.Translations.Title>
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Title}
  **/
get title () { return this.#title }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Title}
  **/
set title (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Title>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferTranslationsActionRes.Translations.Title) {
			this.#title = value
		} else {
			this.#title = new GetOfferTranslationsActionRes.Translations.Title(value)
		}
}
setTitle (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Title>) {
	this.title = value
	return this
}
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description}
  **/
 #description ! : InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description>
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description}
  **/
set description (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferTranslationsActionRes.Translations.Description) {
			this.#description = value
		} else {
			this.#description = new GetOfferTranslationsActionRes.Translations.Description(value)
		}
}
setDescription (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description>) {
	this.description = value
	return this
}
		/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.SafetyInformation}
  **/
 #safetyInformation ! : InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation>
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.SafetyInformation}
  **/
set safetyInformation (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferTranslationsActionRes.Translations.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new GetOfferTranslationsActionRes.Translations.SafetyInformation(value)
		}
}
setSafetyInformation (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation>) {
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
 #translation : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get translation () { return this.#translation }
/**
  * 
  * @type {string}
  **/
set translation (value: string) {
		this.#translation = String(value);
}
setTranslation (value: string) {
	this.translation = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #type : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value: string) {
		this.#type = String(value);
}
setType (value: string) {
	this.type = value
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
		const d = data as Partial<Title>;
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
	static from(possibleDtoObject: GetOfferTranslationsActionResType.TranslationsType.TitleType) {
		return new GetOfferTranslationsActionRes.Translations.Title(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Title, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.TitleType>) {
		return new GetOfferTranslationsActionRes.Translations.Title(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.TitleType>): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Title> {
		return new GetOfferTranslationsActionRes.Translations.Title ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Title> {
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
 #translation ! : InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation>
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Description.Translation}
  **/
get translation () { return this.#translation }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description.Translation}
  **/
set translation (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetOfferTranslationsActionRes.Translations.Description.Translation) {
			this.#translation = value
		} else {
			this.#translation = new GetOfferTranslationsActionRes.Translations.Description.Translation(value)
		}
}
setTranslation (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation>) {
	this.translation = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #type : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value: string) {
		this.#type = String(value);
}
setType (value: string) {
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
 #sections : InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections>[]  =  []
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections}
  **/
set sections (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections>[]) {
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
setSections (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections>[]) {
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
 #items : InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items>[]  =  []
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items}
  **/
set items (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items>[]) {
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
setItems (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items>[]) {
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
 #type : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value: string) {
		this.#type = String(value);
}
setType (value: string) {
	this.type = value
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
		const d = data as Partial<Items>;
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
	static from(possibleDtoObject: GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType.ItemsType) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType.ItemsType>) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType.ItemsType>): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items> {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items> {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections.Items(this.toJSON());
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
		const d = data as Partial<Sections>;
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
	static from(possibleDtoObject: GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description.Translation.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType>) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType>): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections> {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation.Sections> {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation.Sections(this.toJSON());
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
		const d = data as Partial<Translation>;
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
	static from(possibleDtoObject: GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description.Translation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType>) {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType>): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation> {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description.Translation> {
		return new GetOfferTranslationsActionRes.Translations.Description.Translation(this.toJSON());
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
		const d = data as Partial<Description>;
			if (d.translation !== undefined) { this.translation = d.translation }
			if (d.type !== undefined) { this.type = d.type }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Description>;
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
	static from(possibleDtoObject: GetOfferTranslationsActionResType.TranslationsType.DescriptionType) {
		return new GetOfferTranslationsActionRes.Translations.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.DescriptionType>) {
		return new GetOfferTranslationsActionRes.Translations.Description(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.DescriptionType>): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description> {
		return new GetOfferTranslationsActionRes.Translations.Description ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferTranslationsActionRes.Translations.Description> {
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
 #products : InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation.Products>[]  =  []
		/**
  * 
  * @returns {GetOfferTranslationsActionRes.Translations.SafetyInformation.Products}
  **/
get products () { return this.#products }
/**
  * 
  * @type {GetOfferTranslationsActionRes.Translations.SafetyInformation.Products}
  **/
set products (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation.Products>[]) {
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
setProducts (value: InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation.Products>[]) {
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
 #translation : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get translation () { return this.#translation }
/**
  * 
  * @type {string}
  **/
set translation (value: string) {
		this.#translation = String(value);
}
setTranslation (value: string) {
	this.translation = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #type : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value: string) {
		this.#type = String(value);
}
setType (value: string) {
	this.type = value
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
		const d = data as Partial<Products>;
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
	static from(possibleDtoObject: GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType.ProductsType) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation.Products(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.SafetyInformation.Products, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType.ProductsType>) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation.Products(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType.ProductsType>): InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation.Products> {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation.Products ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation.Products> {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation.Products(this.toJSON());
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
		const d = data as Partial<SafetyInformation>;
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
	static from(possibleDtoObject: GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType>) {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType>): InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation> {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferTranslationsActionRes.Translations.SafetyInformation> {
		return new GetOfferTranslationsActionRes.Translations.SafetyInformation(this.toJSON());
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
		const d = data as Partial<Translations>;
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
		const d = data as Partial<Translations>;
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
	static from(possibleDtoObject: GetOfferTranslationsActionResType.TranslationsType) {
		return new GetOfferTranslationsActionRes.Translations(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes.Translations, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferTranslationsActionResType.TranslationsType>) {
		return new GetOfferTranslationsActionRes.Translations(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferTranslationsActionResType.TranslationsType>): InstanceType<typeof GetOfferTranslationsActionRes.Translations> {
		return new GetOfferTranslationsActionRes.Translations ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferTranslationsActionRes.Translations> {
		return new GetOfferTranslationsActionRes.Translations(this.toJSON());
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
		const d = data as Partial<GetOfferTranslationsActionRes>;
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
	static from(possibleDtoObject: GetOfferTranslationsActionResType) {
		return new GetOfferTranslationsActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetOfferTranslationsActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetOfferTranslationsActionResType>) {
		return new GetOfferTranslationsActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetOfferTranslationsActionResType>): InstanceType<typeof GetOfferTranslationsActionRes> {
		return new GetOfferTranslationsActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetOfferTranslationsActionRes> {
		return new GetOfferTranslationsActionRes(this.toJSON());
	}
}
export abstract class GetOfferTranslationsActionResFactory {
	abstract create(data: unknown): GetOfferTranslationsActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for getOfferTranslationsActionRes
  **/
	export type GetOfferTranslationsActionResType =  {
			/**
  * 
  * @type {GetOfferTranslationsActionResType.TranslationsType[]}
  **/
 translations : GetOfferTranslationsActionResType.TranslationsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetOfferTranslationsActionResType {
	/**
  * The base type definition for translationsType
  **/
	export type TranslationsType =  {
			/**
  * 
  * @type {string}
  **/
 language : string;
			/**
  * 
  * @type {GetOfferTranslationsActionResType.TranslationsType.TitleType}
  **/
 title : GetOfferTranslationsActionResType.TranslationsType.TitleType;
			/**
  * 
  * @type {GetOfferTranslationsActionResType.TranslationsType.DescriptionType}
  **/
 description : GetOfferTranslationsActionResType.TranslationsType.DescriptionType;
			/**
  * 
  * @type {GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType}
  **/
 safetyInformation : GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace TranslationsType {
	/**
  * The base type definition for titleType
  **/
	export type TitleType =  {
			/**
  * 
  * @type {string}
  **/
 translation : string;
			/**
  * 
  * @type {string}
  **/
 type : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace TitleType {
}
	/**
  * The base type definition for descriptionType
  **/
	export type DescriptionType =  {
			/**
  * 
  * @type {GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType}
  **/
 translation : GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType;
			/**
  * 
  * @type {string}
  **/
 type : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DescriptionType {
	/**
  * The base type definition for translationType
  **/
	export type TranslationType =  {
			/**
  * 
  * @type {GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType[]}
  **/
 sections : GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace TranslationType {
	/**
  * The base type definition for sectionsType
  **/
	export type SectionsType =  {
			/**
  * 
  * @type {GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType.ItemsType[]}
  **/
 items : GetOfferTranslationsActionResType.TranslationsType.DescriptionType.TranslationType.SectionsType.ItemsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SectionsType {
	/**
  * The base type definition for itemsType
  **/
	export type ItemsType =  {
			/**
  * 
  * @type {string}
  **/
 type : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ItemsType {
}
}
}
}
	/**
  * The base type definition for safetyInformationType
  **/
	export type SafetyInformationType =  {
			/**
  * 
  * @type {GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType.ProductsType[]}
  **/
 products : GetOfferTranslationsActionResType.TranslationsType.SafetyInformationType.ProductsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SafetyInformationType {
	/**
  * The base type definition for productsType
  **/
	export type ProductsType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 translation : string;
			/**
  * 
  * @type {string}
  **/
 type : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ProductsType {
}
}
}
}