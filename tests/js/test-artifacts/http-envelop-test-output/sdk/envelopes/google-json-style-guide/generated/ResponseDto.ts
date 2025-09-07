import { isPlausibleObject } from './sdk/common/isPlausibleObject';
import { withPrefix } from './sdk/common/withPrefix';
/**
  * The base class definition for responseDto
  **/
export class ResponseDto {
		/**
  * 
  * @type {string}
  **/
 #apiVersion : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get apiVersion () { return this.#apiVersion }
/**
  * 
  * @type {string}
  **/
set apiVersion (value: string) {
	 	const correctType = typeof value === 'string';
		this.#apiVersion = correctType ? value : ('' + value);
}
setApiVersion (value: string) {
	this.apiVersion = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #context : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get context () { return this.#context }
/**
  * 
  * @type {string}
  **/
set context (value: string) {
	 	const correctType = typeof value === 'string';
		this.#context = correctType ? value : ('' + value);
}
setContext (value: string) {
	this.context = value
	return this
}
		/**
  * 
  * @type {ResponseDto.Data}
  **/
 #data : InstanceType<typeof ResponseDto.Data> | null  =  null
		/**
  * 
  * @returns {ResponseDto.Data}
  **/
get data () { return this.#data }
/**
  * 
  * @type {ResponseDto.Data}
  **/
set data (value: InstanceType<typeof ResponseDto.Data> | null) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof ResponseDto.Data) {
			this.#data = value
		} else {
			this.#data = new ResponseDto.Data(value)
		}
}
setData (value: InstanceType<typeof ResponseDto.Data> | null) {
	this.data = value
	return this
}
/**
  * The base class definition for data
  **/
static Data = class Data {
		/**
  * 
  * @type {any}
  **/
 #item : any  =  null
		/**
  * 
  * @returns {any}
  **/
get item () { return this.#item }
/**
  * 
  * @type {any}
  **/
set item (value: any) {
		this.#item = value;
}
setItem (value: any) {
	this.item = value
	return this
}
		/**
  * 
  * @type {any}
  **/
 #items : any  =  null
		/**
  * 
  * @returns {any}
  **/
get items () { return this.#items }
/**
  * 
  * @type {any}
  **/
set items (value: any) {
		this.#items = value;
}
setItems (value: any) {
	this.items = value
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
		const d = data as Partial<Data>;
			if (d.item !== undefined) { this.item = d.item }
			if (d.items !== undefined) { this.items = d.items }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				item: this.#item,
				items: this.#items,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			item: 'item',
			items: 'items',
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
		const d = data as Partial<ResponseDto>;
			if (d.apiVersion !== undefined) { this.apiVersion = d.apiVersion }
			if (d.context !== undefined) { this.context = d.context }
			if (d.data !== undefined) { this.data = d.data }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				apiVersion: this.#apiVersion,
				context: this.#context,
				data: this.#data,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			apiVersion: 'apiVersion',
			context: 'context',
			data$: 'data',
get data() {
				return withPrefix(
					"data",
					ResponseDto.Data.Fields
				);
			},
	  }
	}
}
export abstract class ResponseDtoFactory {
	abstract create(data: unknown): ResponseDto;
}
	/**
  * The base type definition for responseDto
  **/
	export type ResponseDtoType =  {
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
  * @type {ResponseDtoType.DataType}
  **/
 data?: ResponseDtoType.DataType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ResponseDtoType {
	/**
  * The base type definition for dataType
  **/
	export type DataType =  {
			/**
  * 
  * @type {any}
  **/
 item?: any;
			/**
  * 
  * @type {any}
  **/
 items?: any;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DataType {
}
}