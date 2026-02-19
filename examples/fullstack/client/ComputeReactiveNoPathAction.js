import { URLSearchParamsX } from './sdk/common/URLSearchParamsX';
import { WebSocketX } from './sdk/common/WebSocketX';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action computeReactiveNoPath
*/
	/**
 * ComputeReactiveNoPathAction
 */
export class ComputeReactiveNoPathAction { //
  static URL = '/compute/reactive';
  static NewUrl = (
	qs
  ) => buildUrl(
		ComputeReactiveNoPathAction.URL,
		 undefined,
		qs
	);
  static Method = 'reactive';
	static Create = (
		overrideUrl,
		qs,
	) => {
		const url = overrideUrl ?? ComputeReactiveNoPathAction.NewUrl(
			qs
		)
		return new WebSocketX(
			url,
			undefined,
			{
				MessageFactoryClass: undefined,
			}
		);
	}
  static Definition = {
  "name": "computeReactiveNoPath",
  "url": "/compute/reactive",
  "method": "reactive",
  "qs": [
    {
      "name": "queryParam1",
      "type": "string"
    },
    {
      "name": "securityToken",
      "type": "string"
    }
  ],
  "description": "Reactive compute elsasements."
}
}
/**
 * ComputeReactiveNoPathActionQueryParams class
 * Auto-generated from EmiAction
 */
export class ComputeReactiveNoPathActionQueryParams extends URLSearchParamsX {
  /**
   * 
   * @returns { string | null }
   */
  getQueryParam1 () {
    return this.getTyped('queryParam1' , 'string | null');
  }
  /**
   * 
   * @param { string | null } value
   */
  setQueryParam1 (value: string | null) {
    this.set('queryParam1', value);
    return this;
  }
  /**
   * 
   * @returns { string | null }
   */
  getSecurityToken () {
    return this.getTyped('securityToken' , 'string | null');
  }
  /**
   * 
   * @param { string | null } value
   */
  setSecurityToken (value: string | null) {
    this.set('securityToken', value);
    return this;
  }
}