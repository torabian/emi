import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { GResponse } from './sdk/envelopes/index';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action envelopeExample
*/
	/**
 * EnvelopeExampleAction
 */
export class EnvelopeExampleAction { //
  static URL = '/response/with/envelop';
  static NewUrl = (
	qs
  ) => buildUrl(
		EnvelopeExampleAction.URL,
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
			overrideUrl ?? EnvelopeExampleAction.NewUrl(
				qs
			),
			{
				method: EnvelopeExampleAction.Method,
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
				creatorFn: (item) => new EnvelopeExampleActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new EnvelopeExampleActionRes(item))
		const res = await EnvelopeExampleAction.Fetch$(
			qs,
			ctx,
			init,
			overrideUrl,
			);
			return handleFetchResponse(
				res, 
				(data) => { 
					const resp = new GResponse<EnvelopeExampleActionRes>();
					if (creatorFn) {
						resp.setCreator(creatorFn);
					}
					resp.inject(data);
					return resp;
			},
				onMessage,
				init?.signal,
			);
	}
  static Definition = {
  "name": "envelopeExample",
  "url": "/response/with/envelop",
  "method": "get",
  "out": {
    "envelope": "GResponse",
    "fields": [
      {
        "name": "content",
        "type": "string"
      }
    ]
  }
}
}
/**
  * The base class definition for envelopeExampleActionRes
  **/
export class EnvelopeExampleActionRes {
		/**
  * 
  * @type {string}
  **/
 #content  =  ""
		/**
  * 
  * @returns {string}
  **/
get content () { return this.#content }
/**
  * 
  * @type {string}
  **/
set content (value) {
		this.#content = String(value);
}
setContent (value) {
	this.content = value
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
			if (d.content !== undefined) { this.content = d.content }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				content: this.#content,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			content: 'content',
	  }
	}
	/**
	* Creates an instance of EnvelopeExampleActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EnvelopeExampleActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of EnvelopeExampleActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EnvelopeExampleActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new EnvelopeExampleActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EnvelopeExampleActionRes(this.toJSON());
	}
}