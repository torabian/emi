/**
 * Used in fetch context, to detect if the response is not a buffer
 * or arraybufer, then create class instance of it.
 * In such cases when response is blob, its better to pass the original classes to caller.
 * @param obj
 * @returns
 */
export declare const isPlausibleObject: (obj: any) => boolean;
