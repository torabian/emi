import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get events about the seller's offers
*/
	/**
 * GetEventsAboutTheSellerSOffersAction
 */
export class GetEventsAboutTheSellerSOffersAction { //
  static URL = 'https://api.{environment}/sale/offer-events';
  static NewUrl = (
	qs
  ) => buildUrl(
		GetEventsAboutTheSellerSOffersAction.URL,
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
			overrideUrl ?? GetEventsAboutTheSellerSOffersAction.NewUrl(
				qs
			),
			{
				method: GetEventsAboutTheSellerSOffersAction.Method,
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
				creatorFn: (item) => new GetEventsAboutTheSellerSOffersActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetEventsAboutTheSellerSOffersActionRes(item))
		const res = await GetEventsAboutTheSellerSOffersAction.Fetch$(
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
  "name": "Get events about the seller's offers",
  "url": "https://api.{environment}/sale/offer-events",
  "method": "get",
  "description": "Use this endpoint to get events from the last 24 hours concerning changes in the authorized seller's offers. At present we support the following events:\nOFFER_ACTIVATED - offer is visible on site and available for purchase, occurs when offer status changes from ACTIVATING to ACTIVE. OFFER_CHANGED - occurs when offer's fields has been changed e.g. description or photos. OFFER_ENDED - offer is no longer available for purchase, occurs when offer status changes from ACTIVE to ENDED. OFFER_STOCK_CHANGED - stock in an offer was changed either via purchase or by seller. OFFER_PRICE_CHANGED - occurs when price in an offer was changed. OFFER_ARCHIVED - offer is no longer available on listing and has been archived. OFFER_BID_PLACED - bid was placed on the offer. OFFER_BID_CANCELED - bid for offer was canceled. OFFER_TRANSLATION_UPDATED - translation of offer was updated. OFFER_VISIBILITY_CHANGED - visibility of offer was changed on marketplaces. OFFER_DELIVERY_COUNTRIES_BLOCKED - the offer has been blocked in selected countries. Returned events may occur by actions made via browser or API. The resource allows you to get events concerning active offers and offers scheduled for activation (status ACTIVE and ACTIVATING). Returned events do not concern offers in INACTIVE and ENDED status (the exception is OFFER_ARCHIVED event). External id is returned for all event types except OFFER_BID_PLACED and OFFER_BID_CANCELED. Please note that one change may result in more than one event. Read more: PL / EN.",
  "out": {
    "fields": [
      {
        "name": "offerEvents",
        "description": "List of events related to offer state changes",
        "type": "array",
        "fields": [
          {
            "name": "id",
            "description": "Unique event identifier (base64 encoded)",
            "type": "string"
          },
          {
            "name": "occurredAt",
            "description": "ISO8601 timestamp when the event occurred",
            "type": "string"
          },
          {
            "name": "type",
            "description": "Event type (e.g., OFFER_ACTIVATED, OFFER_ENDED, etc.)",
            "type": "string"
          },
          {
            "name": "offer",
            "description": "Basic offer information for which event occurred",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "external",
                "type": "object",
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
    ]
  }
}
}
/**
  * The base class definition for getEventsAboutTheSellerSOffersActionRes
  **/
export class GetEventsAboutTheSellerSOffersActionRes {
		/**
  * List of events related to offer state changes
  * @type {GetEventsAboutTheSellerSOffersActionRes.OfferEvents}
  **/
 #offerEvents  =  []
		/**
  * List of events related to offer state changes
  * @returns {GetEventsAboutTheSellerSOffersActionRes.OfferEvents}
  **/
get offerEvents () { return this.#offerEvents }
/**
  * List of events related to offer state changes
  * @type {GetEventsAboutTheSellerSOffersActionRes.OfferEvents}
  **/
set offerEvents (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetEventsAboutTheSellerSOffersActionRes.OfferEvents) {
			this.#offerEvents = value
		} else {
			this.#offerEvents = value.map(item => new GetEventsAboutTheSellerSOffersActionRes.OfferEvents(item))
		}
}
setOfferEvents (value) {
	this.offerEvents = value
	return this
}
/**
  * The base class definition for offerEvents
  **/
static OfferEvents = class OfferEvents {
		/**
  * Unique event identifier (base64 encoded)
  * @type {string}
  **/
 #id  =  ""
		/**
  * Unique event identifier (base64 encoded)
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * Unique event identifier (base64 encoded)
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
  * ISO8601 timestamp when the event occurred
  * @type {string}
  **/
 #occurredAt  =  ""
		/**
  * ISO8601 timestamp when the event occurred
  * @returns {string}
  **/
get occurredAt () { return this.#occurredAt }
/**
  * ISO8601 timestamp when the event occurred
  * @type {string}
  **/
set occurredAt (value) {
		this.#occurredAt = String(value);
}
setOccurredAt (value) {
	this.occurredAt = value
	return this
}
		/**
  * Event type (e.g., OFFER_ACTIVATED, OFFER_ENDED, etc.)
  * @type {string}
  **/
 #type  =  ""
		/**
  * Event type (e.g., OFFER_ACTIVATED, OFFER_ENDED, etc.)
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * Event type (e.g., OFFER_ACTIVATED, OFFER_ENDED, etc.)
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
  * Basic offer information for which event occurred
  * @type {GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer}
  **/
 #offer
		/**
  * Basic offer information for which event occurred
  * @returns {GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer}
  **/
get offer () { return this.#offer }
/**
  * Basic offer information for which event occurred
  * @type {GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer}
  **/
set offer (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer) {
			this.#offer = value
		} else {
			this.#offer = new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer(value)
		}
}
setOffer (value) {
	this.offer = value
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
  * @type {GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External}
  **/
 #external
		/**
  * 
  * @returns {GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External}
  **/
get external () { return this.#external }
/**
  * 
  * @type {GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External}
  **/
set external (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External) {
			this.#external = value
		} else {
			this.#external = new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External(value)
		}
}
setExternal (value) {
	this.external = value
	return this
}
/**
  * The base class definition for external
  **/
static External = class External {
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
	* Creates an instance of GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External(possibleDtoObject);
	}
	/**
	* Creates an instance of GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External(partialDtoObject);
	}
	copyWith(partial) {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External(this.toJSON());
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
			if (d.external !== undefined) { this.external = d.external }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.external instanceof GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External)) { this.external = new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External(d.external || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				external: this.#external,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			external$: 'external',
get external() {
					return withPrefix(
						"offerEvents.offer.external",
						GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.External.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer(possibleDtoObject);
	}
	/**
	* Creates an instance of GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer(partialDtoObject);
	}
	copyWith(partial) {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer(this.toJSON());
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
			if (d.occurredAt !== undefined) { this.occurredAt = d.occurredAt }
			if (d.type !== undefined) { this.type = d.type }
			if (d.offer !== undefined) { this.offer = d.offer }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.offer instanceof GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer)) { this.offer = new GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer(d.offer || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				occurredAt: this.#occurredAt,
				type: this.#type,
				offer: this.#offer,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			occurredAt: 'occurredAt',
			type: 'type',
			offer$: 'offer',
get offer() {
					return withPrefix(
						"offerEvents.offer",
						GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Offer.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetEventsAboutTheSellerSOffersActionRes.OfferEvents, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents(possibleDtoObject);
	}
	/**
	* Creates an instance of GetEventsAboutTheSellerSOffersActionRes.OfferEvents, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents(partialDtoObject);
	}
	copyWith(partial) {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetEventsAboutTheSellerSOffersActionRes.OfferEvents(this.toJSON());
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
			if (d.offerEvents !== undefined) { this.offerEvents = d.offerEvents }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				offerEvents: this.#offerEvents,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			offerEvents$: 'offerEvents',
get offerEvents() {
					return withPrefix(
						"offerEvents[:i]",
						GetEventsAboutTheSellerSOffersActionRes.OfferEvents.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetEventsAboutTheSellerSOffersActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GetEventsAboutTheSellerSOffersActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetEventsAboutTheSellerSOffersActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetEventsAboutTheSellerSOffersActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new GetEventsAboutTheSellerSOffersActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetEventsAboutTheSellerSOffersActionRes(this.toJSON());
	}
}