import { URLSearchParamsX, WebSocketX, buildUrl } from './sdk/js';
/**
* Action to communicate with the action webSocketOrgEcho
*/
export type WebSocketOrgEchoActionOptions = {
	queryKey?: unknown[];
	qs?: WebSocketOrgEchoQueryParams;
	headers?: WebSocketOrgEchoReqHeaders;
};
	/**
 * WebSocketOrgEchoAction
 */
export class WebSocketOrgEchoAction {
  static URL = 'wss://echo.websocket.org/.ws';
  static NewUrl = (
	qs?: WebSocketOrgEchoQueryParams
  ) => buildUrl(
		WebSocketOrgEchoAction.URL,
		 undefined,
		qs
	);
  static Method = 'reactive';
	static Create = (
		overrideUrl?: string,
		qs?: WebSocketOrgEchoQueryParams,
	) => {
		const url = overrideUrl ?? WebSocketOrgEchoAction.NewUrl(
			qs
		)
		return new WebSocketX<WebSocketOrgEchoReq, WebSocketOrgEchoRes>(
			url
		);
	}
}
	/**
  * @description The base type definition for webSocketOrgEchoReq
  **/
	export type WebSocketOrgEchoReqType =  {
			/**
  * @type {string}
  * @description 
  **/
 firstName?: string;
			/**
  * @type {string}
  * @description 
  **/
 lastName?: string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WebSocketOrgEchoReqType {
}
/**
  * @decription The base class definition for webSocketOrgEchoReq
  **/
export class WebSocketOrgEchoReq {
	constructor(data: unknown) {
		// This probably doesn't cover the nested objects
		const d = data as Partial<WebSocketOrgEchoReq>;
			if (d[`firstName`] !== undefined) { 
 this.setFirstName (d[`firstName`]) 
}
			if (d[`lastName`] !== undefined) { 
 this.setLastName (d[`lastName`]) 
}
	}
		/**
  * 
  * @type {string}
  **/
 firstName: string = ""
		/**
  * @returns {string}
  * @description 
  **/
getFirstName () { return this[`firstName`] }
		/**
  * 
  * @param {string}
  **/
setFirstName (value: string ) { this[`firstName`] = value; return this; } 
		/**
  * 
  * @type {string}
  **/
 lastName: string = ""
		/**
  * @returns {string}
  * @description 
  **/
getLastName () { return this[`lastName`] }
		/**
  * 
  * @param {string}
  **/
setLastName (value: string ) { this[`lastName`] = value; return this; } 
}
export abstract class WebSocketOrgEchoReqFactory {
	abstract create(data: unknown): WebSocketOrgEchoReq;
}
	/**
  * @description The base type definition for webSocketOrgEchoRes
  **/
	export type WebSocketOrgEchoResType =  {
			/**
  * @type {string}
  * @description 
  **/
 lastName?: string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WebSocketOrgEchoResType {
}
/**
  * @decription The base class definition for webSocketOrgEchoRes
  **/
export class WebSocketOrgEchoRes {
	constructor(data: unknown) {
		// This probably doesn't cover the nested objects
		const d = data as Partial<WebSocketOrgEchoRes>;
			if (d[`lastName`] !== undefined) { 
 this.setLastName (d[`lastName`]) 
}
	}
		/**
  * 
  * @type {string}
  **/
 lastName: string = ""
		/**
  * @returns {string}
  * @description 
  **/
getLastName () { return this[`lastName`] }
		/**
  * 
  * @param {string}
  **/
setLastName (value: string ) { this[`lastName`] = value; return this; } 
}
export abstract class WebSocketOrgEchoResFactory {
	abstract create(data: unknown): WebSocketOrgEchoRes;
}
/**
 * WebSocketOrgEchoReqHeaders class
 * Auto-generated from Module3Action
 */
export class WebSocketOrgEchoReqHeaders extends Headers {
  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }
}
/**
 * WebSocketOrgEchoResHeaders class
 * Auto-generated from Module3Action
 */
export class WebSocketOrgEchoResHeaders extends Headers {
  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }
}
/**
 * WebSocketOrgEchoQueryParams class
 * Auto-generated from Module3Action
 */
export class WebSocketOrgEchoQueryParams extends URLSearchParamsX {
}