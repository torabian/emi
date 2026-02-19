/**
  * The base class definition for commonVectorResponseDto
  **/
export class CommonVectorResponseDto {
		/**
  * 
  * @type {number[]}
  **/
 #outputVector  =  []
		/**
  * 
  * @returns {number[]}
  **/
get outputVector () { return this.#outputVector }
/**
  * 
  * @type {number[]}
  **/
set outputVector (value) {
}
setOutputVector (value) {
	this.outputVector = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
			if (d.outputVector !== undefined) { this.outputVector = d.outputVector }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				outputVector: this.#outputVector,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			outputVector$: 'outputVector',
get outputVector() {
					return "outputVector[:i]";
						},
	  }
	}
	/**
	* Creates an instance of CommonVectorResponseDto, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new CommonVectorResponseDto(possibleDtoObject);
	}
	/**
	* Creates an instance of CommonVectorResponseDto, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new CommonVectorResponseDto(partialDtoObject);
	}
	copyWith(partial) {
		return new CommonVectorResponseDto ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new CommonVectorResponseDto(this.toJSON());
	}
}