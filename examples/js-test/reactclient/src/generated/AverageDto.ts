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
	constructor(data: unknown = undefined) {
		if (data === null || data === undefined) {
			return;
		}
		if (typeof data === "string") {
			this.applyFromObject(JSON.parse(data));
		} else if (this.#isJsonAppliable(data)) {
			this.applyFromObject(data);
		} else {
			throw new Error("Instance cannot be created on an unknown value, check the content being passed. got: "  + typeof data);
		}
	}
	#isJsonAppliable(obj: unknown) {
		const g = globalThis as any
		const isBuffer =
			typeof g.Buffer !== "undefined" &&
			typeof g.Buffer.isBuffer === "function" &&
			g.Buffer.isBuffer(obj);
		const isBlob =
			typeof g.Blob !== "undefined" && obj instanceof g.Blob;
		return (
			obj &&
			typeof obj === "object" &&
			!Array.isArray(obj) &&
			!isBuffer &&
			!(obj instanceof ArrayBuffer) &&
			!isBlob
		);
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
	/**
	* Creates an instance of AverageDto, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: AverageDtoType) {
		return new AverageDto(possibleDtoObject);
	}
	/**
	* Creates an instance of AverageDto, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<AverageDtoType>) {
		return new AverageDto(partialDtoObject);
	}
	copyWith(partial: PartialDeep<AverageDtoType>): InstanceType<typeof AverageDto> {
		return new AverageDto ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof AverageDto> {
		return new AverageDto(this.toJSON());
	}
}
export abstract class AverageDtoFactory {
	abstract create(data: unknown): AverageDto;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for averageDto
  **/
	export type AverageDtoType =  {
			/**
  * 
  * @type {number}
  **/
 number : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AverageDtoType {
}