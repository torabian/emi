import { withPrefix } from './sdk/common/withPrefix';
/**
  * The base class definition for commonVectorComputeDto
  **/
export class CommonVectorComputeDto {
		/**
  * 
  * @type {number[]}
  **/
 #initialVector1  =  []
		/**
  * 
  * @returns {number[]}
  **/
get initialVector1 () { return this.#initialVector1 }
/**
  * 
  * @type {number[]}
  **/
set initialVector1 (value) {
}
setInitialVector1 (value) {
	this.initialVector1 = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #value  =  undefined
		/**
  * 
  * @returns {string}
  **/
get value () { return this.#value }
/**
  * 
  * @type {string}
  **/
set value (value) {
	 	const correctType = typeof value === 'string' || value === undefined || value === null
		this.#value = correctType ? value : String(value);
}
setValue (value) {
	this.value = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #valuex  =  ""
		/**
  * 
  * @returns {string}
  **/
get valuex () { return this.#valuex }
/**
  * 
  * @type {string}
  **/
set valuex (value) {
		this.#valuex = String(value);
}
setValuex (value) {
	this.valuex = value
	return this
}
		/**
  * 
  * @type {number[]}
  **/
 #initialVector2  =  []
		/**
  * 
  * @returns {number[]}
  **/
get initialVector2 () { return this.#initialVector2 }
/**
  * 
  * @type {number[]}
  **/
set initialVector2 (value) {
}
setInitialVector2 (value) {
	this.initialVector2 = value
	return this
}
		/**
  * 
  * @type {CommonVectorComputeDto.FieldTypeArray}
  **/
 #fieldTypeArray  =  []
		/**
  * 
  * @returns {CommonVectorComputeDto.FieldTypeArray}
  **/
get fieldTypeArray () { return this.#fieldTypeArray }
/**
  * 
  * @type {CommonVectorComputeDto.FieldTypeArray}
  **/
set fieldTypeArray (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CommonVectorComputeDto.FieldTypeArray) {
			this.#fieldTypeArray = value
		} else {
			this.#fieldTypeArray = value.map(item => new CommonVectorComputeDto.FieldTypeArray(item))
		}
}
setFieldTypeArray (value) {
	this.fieldTypeArray = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #fieldTypeSlice  =  []
		/**
  * 
  * @returns {string[]}
  **/
get fieldTypeSlice () { return this.#fieldTypeSlice }
/**
  * 
  * @type {string[]}
  **/
set fieldTypeSlice (value) {
}
setFieldTypeSlice (value) {
	this.fieldTypeSlice = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #fieldInt  =  0
		/**
  * 
  * @returns {number}
  **/
get fieldInt () { return this.#fieldInt }
/**
  * 
  * @type {number}
  **/
set fieldInt (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#fieldInt = parsedValue;
		}
}
setFieldInt (value) {
	this.fieldInt = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #fieldIntNullable  =  undefined
		/**
  * 
  * @returns {number}
  **/
get fieldIntNullable () { return this.#fieldIntNullable }
/**
  * 
  * @type {number}
  **/
set fieldIntNullable (value) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#fieldIntNullable = parsedValue;
		}
}
setFieldIntNullable (value) {
	this.fieldIntNullable = value
	return this
}
		/**
  * 
  * @type {Money}
  **/
 #complexMoney
		/**
  * 
  * @returns {Money}
  **/
get complexMoney () { return this.#complexMoney }
/**
  * 
  * @type {Money}
  **/
set complexMoney (value) {
}
setComplexMoney (value) {
	this.complexMoney = value
	return this
}
/**
  * The base class definition for fieldTypeArray
  **/
static FieldTypeArray = class FieldTypeArray {
		/**
  * 
  * @type {string}
  **/
 #arrayField1  =  ""
		/**
  * 
  * @returns {string}
  **/
get arrayField1 () { return this.#arrayField1 }
/**
  * 
  * @type {string}
  **/
set arrayField1 (value) {
		this.#arrayField1 = String(value);
}
setArrayField1 (value) {
	this.arrayField1 = value
	return this
}
		/**
  * 
  * @type {CommonVectorComputeDto.FieldTypeArray.ArrayField2}
  **/
 #arrayField2  =  []
		/**
  * 
  * @returns {CommonVectorComputeDto.FieldTypeArray.ArrayField2}
  **/
get arrayField2 () { return this.#arrayField2 }
/**
  * 
  * @type {CommonVectorComputeDto.FieldTypeArray.ArrayField2}
  **/
set arrayField2 (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CommonVectorComputeDto.FieldTypeArray.ArrayField2) {
			this.#arrayField2 = value
		} else {
			this.#arrayField2 = value.map(item => new CommonVectorComputeDto.FieldTypeArray.ArrayField2(item))
		}
}
setArrayField2 (value) {
	this.arrayField2 = value
	return this
}
/**
  * The base class definition for arrayField2
  **/
static ArrayField2 = class ArrayField2 {
		/**
  * 
  * @type {string}
  **/
 #lastItem  =  ""
		/**
  * 
  * @returns {string}
  **/
get lastItem () { return this.#lastItem }
/**
  * 
  * @type {string}
  **/
set lastItem (value) {
		this.#lastItem = String(value);
}
setLastItem (value) {
	this.lastItem = value
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
			if (d.lastItem !== undefined) { this.lastItem = d.lastItem }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				lastItem: this.#lastItem,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			lastItem: 'lastItem',
	  }
	}
	/**
	* Creates an instance of CommonVectorComputeDto.FieldTypeArray.ArrayField2, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new CommonVectorComputeDto.FieldTypeArray.ArrayField2(possibleDtoObject);
	}
	/**
	* Creates an instance of CommonVectorComputeDto.FieldTypeArray.ArrayField2, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new CommonVectorComputeDto.FieldTypeArray.ArrayField2(partialDtoObject);
	}
	copyWith(partial) {
		return new CommonVectorComputeDto.FieldTypeArray.ArrayField2 ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new CommonVectorComputeDto.FieldTypeArray.ArrayField2(this.toJSON());
	}
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
			if (d.arrayField1 !== undefined) { this.arrayField1 = d.arrayField1 }
			if (d.arrayField2 !== undefined) { this.arrayField2 = d.arrayField2 }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				arrayField1: this.#arrayField1,
				arrayField2: this.#arrayField2,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			arrayField1: 'arrayField1',
			arrayField2$: 'arrayField2',
get arrayField2() {
					return withPrefix(
						"fieldTypeArray.arrayField2[:i]",
						CommonVectorComputeDto.FieldTypeArray.ArrayField2.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CommonVectorComputeDto.FieldTypeArray, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new CommonVectorComputeDto.FieldTypeArray(possibleDtoObject);
	}
	/**
	* Creates an instance of CommonVectorComputeDto.FieldTypeArray, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new CommonVectorComputeDto.FieldTypeArray(partialDtoObject);
	}
	copyWith(partial) {
		return new CommonVectorComputeDto.FieldTypeArray ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new CommonVectorComputeDto.FieldTypeArray(this.toJSON());
	}
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
			if (d.initialVector1 !== undefined) { this.initialVector1 = d.initialVector1 }
			if (d.value !== undefined) { this.value = d.value }
			if (d.valuex !== undefined) { this.valuex = d.valuex }
			if (d.initialVector2 !== undefined) { this.initialVector2 = d.initialVector2 }
			if (d.fieldTypeArray !== undefined) { this.fieldTypeArray = d.fieldTypeArray }
			if (d.fieldTypeSlice !== undefined) { this.fieldTypeSlice = d.fieldTypeSlice }
			if (d.fieldInt !== undefined) { this.fieldInt = d.fieldInt }
			if (d.fieldIntNullable !== undefined) { this.fieldIntNullable = d.fieldIntNullable }
			if (d.complexMoney !== undefined) { this.complexMoney = d.complexMoney }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				initialVector1: this.#initialVector1,
				value: this.#value,
				valuex: this.#valuex,
				initialVector2: this.#initialVector2,
				fieldTypeArray: this.#fieldTypeArray,
				fieldTypeSlice: this.#fieldTypeSlice,
				fieldInt: this.#fieldInt,
				fieldIntNullable: this.#fieldIntNullable,
				complexMoney: this.#complexMoney,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			initialVector1$: 'initialVector1',
get initialVector1() {
					return "initialVector1[:i]";
						},
			value: 'value',
			valuex: 'valuex',
			initialVector2$: 'initialVector2',
get initialVector2() {
					return "initialVector2[:i]";
						},
			fieldTypeArray$: 'fieldTypeArray',
get fieldTypeArray() {
					return withPrefix(
						"fieldTypeArray[:i]",
						CommonVectorComputeDto.FieldTypeArray.Fields
						);
						},
			fieldTypeSlice$: 'fieldTypeSlice',
get fieldTypeSlice() {
					return "fieldTypeSlice[:i]";
						},
			fieldInt: 'fieldInt',
			fieldIntNullable: 'fieldIntNullable',
			complexMoney: 'complexMoney',
	  }
	}
	/**
	* Creates an instance of CommonVectorComputeDto, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new CommonVectorComputeDto(possibleDtoObject);
	}
	/**
	* Creates an instance of CommonVectorComputeDto, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new CommonVectorComputeDto(partialDtoObject);
	}
	copyWith(partial) {
		return new CommonVectorComputeDto ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new CommonVectorComputeDto(this.toJSON());
	}
}