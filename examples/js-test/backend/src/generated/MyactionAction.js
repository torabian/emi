import { TypedRequestInit, URLSearchParamsX, fetchx } from './sdk';
import { createParamDecorator } from '@nestjs/common';
/**
 * Action to communicate with the action myaction
 */
/**
 * FetchMyactionAction
 */
export class FetchMyactionAction {
  static URL = '/';
  static Method = 'post';
  static Axios = clientInstance.request(config).then((res) => {
    return {
      ...res,
      // if there is a output class, create instance out of it.
      data: new MyactionRes(res.data),
    };
  });
  static Fetch = (init, qs, overrideUrl) =>
    fetchx(
      new URL((overrideUrl ?? FetchMyactionAction.URL) + '?' + qs?.toString()),
      init,
    )
      .then((res) => res.json())
      .then((data) => new MyactionRes(data));
}
/**
 * @decription The base class definition for myactionReq
 **/
export class MyactionReq {
  constructor(data) {
    // This probably doesn't cover the nested objects
    Object.assign(this, data);
  }
  /**
   * @type {string}
   * @description
   **/
  toNumber;
  /**
   * @returns {string}
   * @description
   **/
  getToNumber() {
    return this[`toNumber`];
  }
  /**
   * @param {string}
   * @description
   **/
  setToNumber(value) {
    this[`toNumber`] = value;
    return this;
  }
  /**
   * @type {string}
   * @description
   **/
  body;
  /**
   * @returns {string}
   * @description
   **/
  getBody() {
    return this[`body`];
  }
  /**
   * @param {string}
   * @description
   **/
  setBody(value) {
    this[`body`] = value;
    return this;
  }
  /**
   * @type {MyactionReq.User}
   * @description
   **/
  user;
  /**
   * @returns {MyactionReq.User}
   * @description
   **/
  getUser() {
    return this[`user`];
  }
  /**
   * @param {MyactionReq.User}
   * @description
   **/
  setUser(value) {
    this[`user`] = value;
    return this;
  }
  /**
   * @decription The base class definition for user
   **/
  static User = class User {
    constructor(data) {
      // This probably doesn't cover the nested objects
      Object.assign(this, data);
    }
    /**
     * @type {string}
     * @description
     **/
    firstName;
    /**
     * @returns {string}
     * @description
     **/
    getFirstName() {
      return this[`firstName`];
    }
    /**
     * @param {string}
     * @description
     **/
    setFirstName(value) {
      this[`firstName`] = value;
      return this;
    }
    /**
     * @type {string}
     * @description
     **/
    lastName;
    /**
     * @returns {string}
     * @description
     **/
    getLastName() {
      return this[`lastName`];
    }
    /**
     * @param {string}
     * @description
     **/
    setLastName(value) {
      this[`lastName`] = value;
      return this;
    }
    /**
     * @type {MyactionReq.User.Visits}
     * @description
     **/
    visits;
    /**
     * @returns {MyactionReq.User.Visits}
     * @description
     **/
    getVisits() {
      return this[`visits`];
    }
    /**
     * @param {MyactionReq.User.Visits}
     * @description
     **/
    setVisits(value) {
      this[`visits`] = value;
      return this;
    }
    /**
     * @decription The base class definition for visits
     **/
    static Visits = class Visits {
      constructor(data) {
        // This probably doesn't cover the nested objects
        Object.assign(this, data);
      }
      /**
       * @type {number}
       * @description
       **/
      field1;
      /**
       * @returns {number}
       * @description
       **/
      getField1() {
        return this[`field1`];
      }
      /**
       * @param {number}
       * @description
       **/
      setField1(value) {
        this[`field1`] = value;
        return this;
      }
      /** a placeholder for WebRequestX auto patching the json content to the object **/
      static __jsonParsable;
    };
    /** a placeholder for WebRequestX auto patching the json content to the object **/
    static __jsonParsable;
  };
  /** a placeholder for WebRequestX auto patching the json content to the object **/
  static __jsonParsable;
  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@MyactionReq.Nest() headers: MyactionReq): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator((_data, ctx) => {
    const request = ctx.switchToHttp().getRequest();
    return new MyactionReq(request.body);
  });
}
/**
 * @decription The base class definition for myactionRes
 **/
export class MyactionRes {
  constructor(data) {
    // This probably doesn't cover the nested objects
    Object.assign(this, data);
  }
  /**
   * @type {string}
   * @description
   **/
  queueId;
  /**
   * @returns {string}
   * @description
   **/
  getQueueId() {
    return this[`queueId`];
  }
  /**
   * @param {string}
   * @description
   **/
  setQueueId(value) {
    this[`queueId`] = value;
    return this;
  }
  /** a placeholder for WebRequestX auto patching the json content to the object **/
  static __jsonParsable;
  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@MyactionRes.Nest() headers: MyactionRes): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator((_data, ctx) => {
    const request = ctx.switchToHttp().getRequest();
    return new MyactionRes(request.body);
  });
}
/**
 * MyactionHeaders class
 * Auto-generated from Module3Action
 */
export class MyactionHeaders extends Headers {
  // the getters generated by us would be casting types before returning.
  // you still can use .get function to get the string value.
  // eslint-disable-next-line no-unused-private-class-members
  #getTyped(key, type) {
    const val = this.get(key);
    if (val == null) return null;
    const t = type.toLowerCase();
    if (t.includes('number')) return Number(val);
    if (t.includes('bool')) return val === 'true';
    return val; // string or any other fallback
  }
  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }
  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@MyactionHeaders.Nest() headers: MyactionHeaders): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator((_data, ctx) => {
    const request = ctx.switchToHttp().getRequest();
    return new MyactionHeaders(Object.entries(request.headers));
  });
}
/**
 * MyactionQueryParams class
 * Auto-generated from Module3Action
 */
export class MyactionQueryParams extends URLSearchParamsX {
  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@MyactionQueryParams.Nest() headers: MyactionQueryParams): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator((_data, ctx) => {
    const request = ctx.switchToHttp().getRequest();
    return new MyactionQueryParams(request.query);
  });
}
