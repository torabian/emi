import { withPrefix } from './sdk/common/withPrefix';
/**
  * The base class definition for giantDto
  **/
export class GiantDto {
		/**
  * 
  * @type {string}
  **/
 #firstName  =  ""
		/**
  * 
  * @returns {string}
  **/
get firstName () { return this.#firstName }
/**
  * 
  * @type {string}
  **/
set firstName (value) {
		this.#firstName = String(value);
}
setFirstName (value) {
	this.firstName = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #firstNameNullable  =  undefined
		/**
  * 
  * @returns {string}
  **/
get firstNameNullable () { return this.#firstNameNullable }
/**
  * 
  * @type {string}
  **/
set firstNameNullable (value) {
	 	const correctType = typeof value === 'string' || value === undefined || value === null
		this.#firstNameNullable = correctType ? value : String(value);
}
setFirstNameNullable (value) {
	this.firstNameNullable = value
	return this
}
		/**
  * 
  * @type {GiantDto.Array}
  **/
 #array  =  []
		/**
  * 
  * @returns {GiantDto.Array}
  **/
get array () { return this.#array }
/**
  * 
  * @type {GiantDto.Array}
  **/
set array (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GiantDto.Array) {
			this.#array = value
		} else {
			this.#array = value.map(item => new GiantDto.Array(item))
		}
}
setArray (value) {
	this.array = value
	return this
}
		/**
  * 
  * @type {GiantDto.ArrayNullable}
  **/
 #arrayNullable  =  undefined
		/**
  * 
  * @returns {GiantDto.ArrayNullable}
  **/
get arrayNullable () { return this.#arrayNullable }
/**
  * 
  * @type {GiantDto.ArrayNullable}
  **/
set arrayNullable (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GiantDto.ArrayNullable) {
			this.#arrayNullable = value
		} else {
			this.#arrayNullable = value.map(item => new GiantDto.ArrayNullable(item))
		}
}
setArrayNullable (value) {
	this.arrayNullable = value
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
			if (d.firstName !== undefined) { this.firstName = d.firstName }
			if (d.firstNameNullable !== undefined) { this.firstNameNullable = d.firstNameNullable }
			if (d.array !== undefined) { this.array = d.array }
			if (d.arrayNullable !== undefined) { this.arrayNullable = d.arrayNullable }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				firstName: this.#firstName,
				firstNameNullable: this.#firstNameNullable,
				array: this.#array,
				arrayNullable: this.#arrayNullable,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			firstName: 'firstName',
			firstNameNullable: 'firstNameNullable',
			array$: 'array',
get array() {
					return withPrefix(
						"array[:i]",
						GiantDto.Array.Fields
						);
						},
			arrayNullable$: 'arrayNullable',
get arrayNullable() {
					return withPrefix(
						"arrayNullable[:i]",
						GiantDto.ArrayNullable.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GiantDto, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GiantDto(possibleDtoObject);
	}
	/**
	* Creates an instance of GiantDto, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GiantDto(partialDtoObject);
	}
	copyWith(partial) {
		return new GiantDto ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GiantDto(this.toJSON());
	}
}