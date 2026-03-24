import { URLSearchParamsX } from './sdk/common/URLSearchParamsX';
import { WebSocketX } from './sdk/common/WebSocketX';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action computeReactive
*/
	/**
 * ComputeReactiveAction
 */
export class ComputeReactiveAction { //
  static URL = '/compute/reactive/:id int32/:age int32';
  static NewUrl = (
	params,
	qs
  ) => buildUrl(
		ComputeReactiveAction.URL,
		params,
		qs
	);
  static Method = 'reactive';
	static Create = (
		overrideUrl,
		qs,
			params,
	) => {
		const url = overrideUrl ?? ComputeReactiveAction.NewUrl(
			params,
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
  "name": "computeReactive",
  "url": "/compute/reactive/:id int32/:age int32",
  "method": "reactive",
  "qs": [
    {
      "name": "queryParam1",
      "type": "string"
    },
    {
      "name": "securityToken",
      "type": "string"
    },
    {
      "name": "object1",
      "type": "object",
      "fields": [
        {
          "name": "field1",
          "type": "string"
        },
        {
          "name": "field2",
          "type": "string"
        }
      ]
    },
    {
      "name": "intSlice",
      "type": "slice",
      "primitive": "int"
    },
    {
      "name": "inlineArray",
      "type": "array",
      "fields": [
        {
          "name": "slice1num",
          "type": "int"
        },
        {
          "name": "innerSlice",
          "type": "slice",
          "primitive": "float64"
        }
      ]
    }
  ],
  "description": "Reactive compute elsasements."
}
}
/**
 * ComputeReactiveActionQueryParams class
 * Auto-generated from EmiAction
 */
export class ComputeReactiveActionQueryParams extends URLSearchParamsX {
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
  /**
   * 
   * @returns { any }
   */
  getObject1 () {
    return this.getTyped('object1' , 'any');
  }
  /**
   * 
   * @param { any } value
   */
  setObject1 (value: any) {
    this.set('object1', value);
    return this;
  }
  /**
   * 
   * @returns { any }
   */
  getIntSlice () {
    return this.getTyped('intSlice' , 'any');
  }
  /**
   * 
   * @param { any } value
   */
  setIntSlice (value: any) {
    this.set('intSlice', value);
    return this;
  }
  /**
   * 
   * @returns { any }
   */
  getInlineArray () {
    return this.getTyped('inlineArray' , 'any');
  }
  /**
   * 
   * @param { any } value
   */
  setInlineArray (value: any) {
    this.set('inlineArray', value);
    return this;
  }
}