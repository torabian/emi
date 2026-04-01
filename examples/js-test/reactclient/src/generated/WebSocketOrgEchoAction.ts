import { WebSocketX } from './sdk/common/WebSocketX';
import { buildUrl } from './sdk/common/buildUrl';
import { useWebSocketX } from './sdk/react/useWebSocketX';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action webSocketOrgEcho
*/
export type WebSocketOrgEchoActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
export const useWebSocketOrgEchoAction = (options?: {
	qs?: URLSearchParams,
	overrideUrl?: string
}) => {
	return useWebSocketX(
		() => WebSocketOrgEchoAction.Create(options?.overrideUrl, options?.qs)
	);
};
	/**
 * WebSocketOrgEchoAction
 */
export class WebSocketOrgEchoAction { //
  static URL = 'wss://echo.websocket.org/.ws';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		WebSocketOrgEchoAction.URL,
		 undefined,
		qs
	);
  static Method = 'reactive';
	static Create = (
		overrideUrl?: string,
		qs?: URLSearchParams,
	) => {
		const url = overrideUrl ?? WebSocketOrgEchoAction.NewUrl(
			qs
		)
		return new WebSocketX<WebSocketOrgEchoActionReq, WebSocketOrgEchoActionRes>(
			url,
			undefined,
			{
				MessageFactoryClass: WebSocketOrgEchoActionRes,
			}
		);
	}
  static Definition = {
  "name": "webSocketOrgEcho",
  "url": "wss://echo.websocket.org/.ws",
  "method": "reactive",
  "description": "Websocket.org eco server, to send a json and recieve back",
  "in": {
    "fields": [
      {
        "name": "firstName",
        "type": "string"
      },
      {
        "name": "lastName",
        "type": "string"
      },
      {
        "name": "user",
        "type": "object",
        "fields": [
          {
            "name": "item1",
            "type": "string"
          },
          {
            "name": "item2array",
            "type": "array",
            "fields": [
              {
                "name": "subItem1",
                "type": "int64"
              },
              {
                "name": "subItem2",
                "type": "int64"
              }
            ]
          }
        ]
      }
    ]
  },
  "out": {
    "fields": [
      {
        "name": "lastName",
        "type": "string"
      }
    ]
  }
}
}
/**
  * The base class definition for webSocketOrgEchoActionReq
  **/
export class WebSocketOrgEchoActionReq {
		/**
  * 
  * @type {string}
  **/
 #firstName : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get firstName () { return this.#firstName }
/**
  * 
  * @type {string}
  **/
set firstName (value: string) {
		this.#firstName = String(value);
}
setFirstName (value: string) {
	this.firstName = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #lastName : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get lastName () { return this.#lastName }
/**
  * 
  * @type {string}
  **/
set lastName (value: string) {
		this.#lastName = String(value);
}
setLastName (value: string) {
	this.lastName = value
	return this
}
		/**
  * 
  * @type {WebSocketOrgEchoActionReq.User}
  **/
 #user ! : InstanceType<typeof WebSocketOrgEchoActionReq.User>
		/**
  * 
  * @returns {WebSocketOrgEchoActionReq.User}
  **/
get user () { return this.#user }
/**
  * 
  * @type {WebSocketOrgEchoActionReq.User}
  **/
set user (value: InstanceType<typeof WebSocketOrgEchoActionReq.User>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof WebSocketOrgEchoActionReq.User) {
			this.#user = value
		} else {
			this.#user = new WebSocketOrgEchoActionReq.User(value)
		}
}
setUser (value: InstanceType<typeof WebSocketOrgEchoActionReq.User>) {
	this.user = value
	return this
}
/**
  * The base class definition for user
  **/
static User = class User {
		/**
  * 
  * @type {string}
  **/
 #item1 : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get item1 () { return this.#item1 }
/**
  * 
  * @type {string}
  **/
set item1 (value: string) {
		this.#item1 = String(value);
}
setItem1 (value: string) {
	this.item1 = value
	return this
}
		/**
  * 
  * @type {WebSocketOrgEchoActionReq.User.Item2array}
  **/
 #item2array : InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array>[]  =  []
		/**
  * 
  * @returns {WebSocketOrgEchoActionReq.User.Item2array}
  **/
get item2array () { return this.#item2array }
/**
  * 
  * @type {WebSocketOrgEchoActionReq.User.Item2array}
  **/
set item2array (value: InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof WebSocketOrgEchoActionReq.User.Item2array) {
			this.#item2array = value
		} else {
			this.#item2array = value.map(item => new WebSocketOrgEchoActionReq.User.Item2array(item))
		}
}
setItem2array (value: InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array>[]) {
	this.item2array = value
	return this
}
/**
  * The base class definition for item2array
  **/
static Item2array = class Item2array {
		/**
  * 
  * @type {number}
  **/
 #subItem1 : number  =  0
		/**
  * 
  * @returns {number}
  **/
get subItem1 () { return this.#subItem1 }
/**
  * 
  * @type {number}
  **/
set subItem1 (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#subItem1 = parsedValue;
		}
}
setSubItem1 (value: number) {
	this.subItem1 = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #subItem2 : number  =  0
		/**
  * 
  * @returns {number}
  **/
get subItem2 () { return this.#subItem2 }
/**
  * 
  * @type {number}
  **/
set subItem2 (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#subItem2 = parsedValue;
		}
}
setSubItem2 (value: number) {
	this.subItem2 = value
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
		const d = data as Partial<Item2array>;
			if (d.subItem1 !== undefined) { this.subItem1 = d.subItem1 }
			if (d.subItem2 !== undefined) { this.subItem2 = d.subItem2 }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				subItem1: this.#subItem1,
				subItem2: this.#subItem2,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			subItem1: 'subItem1',
			subItem2: 'subItem2',
	  }
	}
	/**
	* Creates an instance of WebSocketOrgEchoActionReq.User.Item2array, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: WebSocketOrgEchoActionReqType.UserType.Item2arrayType) {
		return new WebSocketOrgEchoActionReq.User.Item2array(possibleDtoObject);
	}
	/**
	* Creates an instance of WebSocketOrgEchoActionReq.User.Item2array, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<WebSocketOrgEchoActionReqType.UserType.Item2arrayType>) {
		return new WebSocketOrgEchoActionReq.User.Item2array(partialDtoObject);
	}
	copyWith(partial: PartialDeep<WebSocketOrgEchoActionReqType.UserType.Item2arrayType>): InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array> {
		return new WebSocketOrgEchoActionReq.User.Item2array ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof WebSocketOrgEchoActionReq.User.Item2array> {
		return new WebSocketOrgEchoActionReq.User.Item2array(this.toJSON());
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
		const d = data as Partial<User>;
			if (d.item1 !== undefined) { this.item1 = d.item1 }
			if (d.item2array !== undefined) { this.item2array = d.item2array }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				item1: this.#item1,
				item2array: this.#item2array,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			item1: 'item1',
			item2array$: 'item2array',
get item2array() {
					return withPrefix(
						"user.item2array[:i]",
						WebSocketOrgEchoActionReq.User.Item2array.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of WebSocketOrgEchoActionReq.User, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: WebSocketOrgEchoActionReqType.UserType) {
		return new WebSocketOrgEchoActionReq.User(possibleDtoObject);
	}
	/**
	* Creates an instance of WebSocketOrgEchoActionReq.User, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<WebSocketOrgEchoActionReqType.UserType>) {
		return new WebSocketOrgEchoActionReq.User(partialDtoObject);
	}
	copyWith(partial: PartialDeep<WebSocketOrgEchoActionReqType.UserType>): InstanceType<typeof WebSocketOrgEchoActionReq.User> {
		return new WebSocketOrgEchoActionReq.User ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof WebSocketOrgEchoActionReq.User> {
		return new WebSocketOrgEchoActionReq.User(this.toJSON());
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
		const d = data as Partial<WebSocketOrgEchoActionReq>;
			if (d.firstName !== undefined) { this.firstName = d.firstName }
			if (d.lastName !== undefined) { this.lastName = d.lastName }
			if (d.user !== undefined) { this.user = d.user }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<WebSocketOrgEchoActionReq>;
			if (!(d.user instanceof WebSocketOrgEchoActionReq.User)) { this.user = new WebSocketOrgEchoActionReq.User(d.user || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				firstName: this.#firstName,
				lastName: this.#lastName,
				user: this.#user,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			firstName: 'firstName',
			lastName: 'lastName',
			user$: 'user',
get user() {
					return withPrefix(
						"user",
						WebSocketOrgEchoActionReq.User.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of WebSocketOrgEchoActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: WebSocketOrgEchoActionReqType) {
		return new WebSocketOrgEchoActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of WebSocketOrgEchoActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<WebSocketOrgEchoActionReqType>) {
		return new WebSocketOrgEchoActionReq(partialDtoObject);
	}
	copyWith(partial: PartialDeep<WebSocketOrgEchoActionReqType>): InstanceType<typeof WebSocketOrgEchoActionReq> {
		return new WebSocketOrgEchoActionReq ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof WebSocketOrgEchoActionReq> {
		return new WebSocketOrgEchoActionReq(this.toJSON());
	}
}
export abstract class WebSocketOrgEchoActionReqFactory {
	abstract create(data: unknown): WebSocketOrgEchoActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for webSocketOrgEchoActionReq
  **/
	export type WebSocketOrgEchoActionReqType =  {
			/**
  * 
  * @type {string}
  **/
 firstName : string;
			/**
  * 
  * @type {string}
  **/
 lastName : string;
			/**
  * 
  * @type {WebSocketOrgEchoActionReqType.UserType}
  **/
 user : WebSocketOrgEchoActionReqType.UserType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WebSocketOrgEchoActionReqType {
	/**
  * The base type definition for userType
  **/
	export type UserType =  {
			/**
  * 
  * @type {string}
  **/
 item1 : string;
			/**
  * 
  * @type {WebSocketOrgEchoActionReqType.UserType.Item2arrayType[]}
  **/
 item2array : WebSocketOrgEchoActionReqType.UserType.Item2arrayType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UserType {
	/**
  * The base type definition for item2arrayType
  **/
	export type Item2arrayType =  {
			/**
  * 
  * @type {number}
  **/
 subItem1 : number;
			/**
  * 
  * @type {number}
  **/
 subItem2 : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Item2arrayType {
}
}
}
/**
  * The base class definition for webSocketOrgEchoActionRes
  **/
export class WebSocketOrgEchoActionRes {
		/**
  * 
  * @type {string}
  **/
 #lastName : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get lastName () { return this.#lastName }
/**
  * 
  * @type {string}
  **/
set lastName (value: string) {
		this.#lastName = String(value);
}
setLastName (value: string) {
	this.lastName = value
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
		const d = data as Partial<WebSocketOrgEchoActionRes>;
			if (d.lastName !== undefined) { this.lastName = d.lastName }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				lastName: this.#lastName,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			lastName: 'lastName',
	  }
	}
	/**
	* Creates an instance of WebSocketOrgEchoActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: WebSocketOrgEchoActionResType) {
		return new WebSocketOrgEchoActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of WebSocketOrgEchoActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<WebSocketOrgEchoActionResType>) {
		return new WebSocketOrgEchoActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<WebSocketOrgEchoActionResType>): InstanceType<typeof WebSocketOrgEchoActionRes> {
		return new WebSocketOrgEchoActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof WebSocketOrgEchoActionRes> {
		return new WebSocketOrgEchoActionRes(this.toJSON());
	}
}
export abstract class WebSocketOrgEchoActionResFactory {
	abstract create(data: unknown): WebSocketOrgEchoActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for webSocketOrgEchoActionRes
  **/
	export type WebSocketOrgEchoActionResType =  {
			/**
  * 
  * @type {string}
  **/
 lastName : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WebSocketOrgEchoActionResType {
}