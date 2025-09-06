import { isPlausibleObject } from './sdk/js';
/**
  * The base class definition for averageDto
  **/
export class AverageDto {
		/**
  * 
  * @type {number}
  **/
 #number : number  =  0
		/**
  * 
  * @returns {number}
  **/
get number () { return this.#number }
/**
  * 
  * @type {number}
  **/
set number (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#number = parsedValue;
		}
}
setNumber (value: number) {
	this.number = value
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
		const d = data as Partial<AverageDto>;
			if (d.number !== undefined) { this.number = d.number }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				number: this.#number,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			number: 'number',
	  }
	}
}
export abstract class AverageDtoFactory {
	abstract create(data: unknown): AverageDto;
}
	/**
  * The base type definition for averageDto
  **/
	export type AverageDtoType =  {
			/**
  * 
  * @type {number}
  **/
 number?: number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AverageDtoType {
}