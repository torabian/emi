import { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { ExecutionContext, createParamDecorator } from '@nestjs/common';
import { TypedRequestInit, URLSearchParamsX, fetchx } from './sdk';
/**
 * Action to communicate with the action myaction
 */
/**
 * FetchMyactionAction
 */
export class FetchMyactionAction {
  static URL = '/';
  static Method = 'post';
  static Axios: (
    clientInstance: AxiosInstance,
    config: AxiosRequestConfig<unknown>,
  ) => Promise<AxiosResponse<unknown>> = (clientInstance, config) =>
    clientInstance
      .request<unknown, AxiosResponse<unknown>, unknown>(config)
      .then((res) => {
        return {
          ...res,
          // if there is a output class, create instance out of it.
          data: new MyactionRes(res.data),
        };
      });
  static Fetch = (
    init?: TypedRequestInit<unknown, unknown>,
    qs?: MyactionQueryParams,
    overrideUrl?: string,
  ) =>
    fetchx<unknown, unknown, unknown>(
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
  toNumber: string;
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
  setToNumber(value: string) {
    this[`toNumber`] = value;
    return this;
  }
  /**
   * @type {string}
   * @description
   **/
  body: string;
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
  setBody(value: string) {
    this[`body`] = value;
    return this;
  }
  /**
   * @type {MyactionReq.User}
   * @description
   **/
  user: InstanceType<typeof MyactionReq.User>;
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
  setUser(value: InstanceType<typeof MyactionReq.User>) {
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
    firstName: string;
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
    setFirstName(value: string) {
      this[`firstName`] = value;
      return this;
    }
    /**
     * @type {string}
     * @description
     **/
    lastName: string;
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
    setLastName(value: string) {
      this[`lastName`] = value;
      return this;
    }
    /**
     * @type {MyactionReq.User.Visits}
     * @description
     **/
    visits: InstanceType<typeof MyactionReq.User.Visits>;
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
    setVisits(value: InstanceType<typeof MyactionReq.User.Visits>) {
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
      field1: number;
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
      setField1(value: number) {
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
  static Nest = createParamDecorator(
    (_data: unknown, ctx: ExecutionContext): MyactionReq => {
      const request = ctx.switchToHttp().getRequest<{
        body: unknown;
        headers: { [s: string]: string };
        query: Record<string, string>;
      }>();
      return new MyactionReq(request.body as Partial<MyactionReq>);
    },
  );
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
  queueId: string;
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
  setQueueId(value: string) {
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
  static Nest = createParamDecorator(
    (_data: unknown, ctx: ExecutionContext): MyactionRes => {
      const request = ctx.switchToHttp().getRequest<{
        body: unknown;
        headers: { [s: string]: string };
        query: Record<string, string>;
      }>();
      return new MyactionRes(request.body as Partial<MyactionRes>);
    },
  );
}
/**
 * MyactionHeaders class
 * Auto-generated from Module3Action
 */
export class MyactionHeaders extends Headers {
  // the getters generated by us would be casting types before returning.
  // you still can use .get function to get the string value.
  // eslint-disable-next-line no-unused-private-class-members
  #getTyped(key, type: string) {
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
  static Nest = createParamDecorator(
    (_data: unknown, ctx: ExecutionContext): MyactionHeaders => {
      const request = ctx.switchToHttp().getRequest<{
        body: unknown;
        headers: { [s: string]: string };
        query: Record<string, string>;
      }>();
      return new MyactionHeaders(Object.entries(request.headers));
    },
  );
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
  static Nest = createParamDecorator(
    (_data: unknown, ctx: ExecutionContext): MyactionQueryParams => {
      const request = ctx.switchToHttp().getRequest<{
        body: unknown;
        headers: { [s: string]: string };
        query: Record<string, string>;
      }>();
      return new MyactionQueryParams(request.query);
    },
  );
}
