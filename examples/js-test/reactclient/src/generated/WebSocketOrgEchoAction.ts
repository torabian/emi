import { WebSocketX } from './sdk/common/WebSocketX';
import { buildUrl } from './sdk/common/buildUrl';
import { isPlausibleObject } from './sdk/common/isPlausibleObject';
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
export class WebSocketOrgEchoAction {
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
	 	const correctType = typeof value === 'string';
		this.#firstName = correctType ? value : ('' + value);
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
	 	const correctType = typeof value === 'string';
		this.#lastName = correctType ? value : ('' + value);
}
setLastName (value: string) {
	this.lastName = value
	return this
}
		/**
  * 
  * @type {WebSocketOrgEchoActionReq.User}
  **/
 #user : InstanceType<typeof WebSocketOrgEchoActionReq.User> | null  =  null
		/**
  * 
  * @returns {WebSocketOrgEchoActionReq.User}
  **/
get user () { return this.#user }
/**
  * 
  * @type {WebSocketOrgEchoActionReq.User}
  **/
set user (value: InstanceType<typeof WebSocketOrgEchoActionReq.User> | null) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof WebSocketOrgEchoActionReq.User) {
			this.#user = value
		} else {
			this.#user = new WebSocketOrgEchoActionReq.User(value)
		}
}
setUser (value: InstanceType<typeof WebSocketOrgEchoActionReq.User> | null) {
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
	 	const correctType = typeof value === 'string';
		this.#item1 = correctType ? value : ('' + value);
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
	constructor(data) {
		if (data === null || data === undefined) {
			return;
		}
		if (typeof data === "string") {
			this.applyFromObject(JSON.parse(data));
		} else if (isPlausibleObject(data)) {
			this.applyFromObject(data);
		} else {
			throw new Error("Instance is not implemented.");
		}
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
}
	constructor(data) {
		if (data === null || data === undefined) {
			return;
		}
		if (typeof data === "string") {
			this.applyFromObject(JSON.parse(data));
		} else if (isPlausibleObject(data)) {
			this.applyFromObject(data);
		} else {
			throw new Error("Instance is not implemented.");
		}
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
}
	constructor(data) {
		if (data === null || data === undefined) {
			return;
		}
		if (typeof data === "string") {
			this.applyFromObject(JSON.parse(data));
		} else if (isPlausibleObject(data)) {
			this.applyFromObject(data);
		} else {
			throw new Error("Instance is not implemented.");
		}
	}
	/**
	* casts the fields of a javascript object into the class properties one by one
	**/
	applyFromObject(data = {}) {
		const d = data as Partial<WebSocketOrgEchoActionReq>;
			if (d.firstName !== undefined) { this.firstName = d.firstName }
			if (d.lastName !== undefined) { this.lastName = d.lastName }
			if (d.user !== undefined) { this.user = d.user }
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
}
export abstract class WebSocketOrgEchoActionReqFactory {
	abstract create(data: unknown): WebSocketOrgEchoActionReq;
}
	/**
  * The base type definition for webSocketOrgEchoActionReq
  **/
	export type WebSocketOrgEchoActionReqType =  {
			/**
  * 
  * @type {string}
  **/
 firstName?: string;
			/**
  * 
  * @type {string}
  **/
 lastName?: string;
			/**
  * 
  * @type {WebSocketOrgEchoActionReqType.UserType}
  **/
 user?: WebSocketOrgEchoActionReqType.UserType;
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
 item1?: string;
			/**
  * 
  * @type {WebSocketOrgEchoActionReqType.UserType.Item2arrayType[]}
  **/
 item2array?: WebSocketOrgEchoActionReqType.UserType.Item2arrayType[];
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
 subItem1?: number;
			/**
  * 
  * @type {number}
  **/
 subItem2?: number;
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
	 	const correctType = typeof value === 'string';
		this.#lastName = correctType ? value : ('' + value);
}
setLastName (value: string) {
	this.lastName = value
	return this
}
	constructor(data) {
		if (data === null || data === undefined) {
			return;
		}
		if (typeof data === "string") {
			this.applyFromObject(JSON.parse(data));
		} else if (isPlausibleObject(data)) {
			this.applyFromObject(data);
		} else {
			throw new Error("Instance is not implemented.");
		}
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
}
export abstract class WebSocketOrgEchoActionResFactory {
	abstract create(data: unknown): WebSocketOrgEchoActionRes;
}
	/**
  * The base type definition for webSocketOrgEchoActionRes
  **/
	export type WebSocketOrgEchoActionResType =  {
			/**
  * 
  * @type {string}
  **/
 lastName?: string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WebSocketOrgEchoActionResType {
}