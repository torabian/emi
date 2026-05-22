/**
 * The base class definition for responseDto
 **/
export declare class ResponseDto {
    #private;
    /**
     *
     * @returns {string}
     **/
    get apiVersion(): string;
    /**
     *
     * @type {string}
     **/
    set apiVersion(value: string);
    setApiVersion(value: string): this;
    /**
     *
     * @returns {string}
     **/
    get context(): string;
    /**
     *
     * @type {string}
     **/
    set context(value: string);
    setContext(value: string): this;
    /**
     *
     * @returns {string}
     **/
    get data(): string;
    /**
     *
     * @type {string}
     **/
    set data(value: string);
    setData(value: string): this;
    constructor(data: any);
    /**
     * casts the fields of a javascript object into the class properties one by one
     **/
    applyFromObject(data?: {}): void;
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON(): {
        apiVersion: string;
        context: string;
        data: string;
    };
    toString(): string;
    static get Fields(): {
        apiVersion: string;
        context: string;
        data: string;
    };
}
export declare abstract class ResponseDtoFactory {
    abstract create(data: unknown): ResponseDto;
}
/**
 * The base type definition for responseDto
 **/
export type ResponseDtoType = {
    /**
     *
     * @type {string}
     **/
    apiVersion?: string;
    /**
     *
     * @type {string}
     **/
    context?: string;
    /**
     *
     * @type {string}
     **/
    data?: string;
};
export declare namespace ResponseDtoType { }
