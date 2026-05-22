import { GiantDto } from './GiantDto';
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
		/**
  * 
  * @type {boolean}
  **/
 #booleanField
		/**
  * 
  * @returns {boolean}
  **/
get booleanField () { return this.#booleanField }
/**
  * 
  * @type {boolean}
  **/
set booleanField (value) {
		this.#booleanField = Boolean(value);
}
setBooleanField (value) {
	this.booleanField = value
	return this
}
		/**
  * 
  * @type {boolean}
  **/
 #booleanFieldNullable  =  undefined
		/**
  * 
  * @returns {boolean}
  **/
get booleanFieldNullable () { return this.#booleanFieldNullable }
/**
  * 
  * @type {boolean}
  **/
set booleanFieldNullable (value) {
	 	const correctType = value === true || value === false || value === undefined || value === null
		this.#booleanFieldNullable = correctType ? value : Boolean(value);
}
setBooleanFieldNullable (value) {
	this.booleanFieldNullable = value
	return this
}
		/**
  * 
  * @type {GiantDto[]}
  **/
 #collectionItems  =  []
		/**
  * 
  * @returns {GiantDto[]}
  **/
get collectionItems () { return this.#collectionItems }
/**
  * 
  * @type {GiantDto[]}
  **/
set collectionItems (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GiantDto) {
			this.#collectionItems = value
		} else {
			this.#collectionItems = value.map(item => new GiantDto(item))
		}
}
setCollectionItems (value) {
	this.collectionItems = value
	return this
}
		/**
  * 
  * @type {GiantDto[]}
  **/
 #collectionItemsNullable  =  undefined
		/**
  * 
  * @returns {GiantDto[]}
  **/
get collectionItemsNullable () { return this.#collectionItemsNullable }
/**
  * 
  * @type {GiantDto[]}
  **/
set collectionItemsNullable (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GiantDto) {
			this.#collectionItemsNullable = value
		} else {
			this.#collectionItemsNullable = value.map(item => new GiantDto(item))
		}
}
setCollectionItemsNullable (value) {
	this.collectionItemsNullable = value
	return this
}
		/**
  * 
  * @type {Date}
  **/
 #dateObject
		/**
  * 
  * @returns {Date}
  **/
get dateObject () { return this.#dateObject }
/**
  * 
  * @type {Date}
  **/
set dateObject (value) {
}
setDateObject (value) {
	this.dateObject = value
	return this
}
		/**
  * 
  * @type {GiantDto}
  **/
 #singleRefNullable  =  undefined
		/**
  * 
  * @returns {GiantDto}
  **/
get singleRefNullable () { return this.#singleRefNullable }
/**
  * 
  * @type {GiantDto}
  **/
set singleRefNullable (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GiantDto) {
			this.#singleRefNullable = value
		} else {
			this.#singleRefNullable = new GiantDto(value)
		}
}
setSingleRefNullable (value) {
	this.singleRefNullable = value
	return this
}
		/**
  * 
  * @type {"Key3" | "Key4"}
  **/
 #enumeration
		/**
  * 
  * @returns {"Key3" | "Key4"}
  **/
get enumeration () { return this.#enumeration }
/**
  * 
  * @type {"Key3" | "Key4"}
  **/
set enumeration (value) {
}
setEnumeration (value) {
	this.enumeration = value
	return this
}
		/**
  * 
  * @type {any}
  **/
 #enumerationNullable  =  undefined
		/**
  * 
  * @returns {any}
  **/
get enumerationNullable () { return this.#enumerationNullable }
/**
  * 
  * @type {any}
  **/
set enumerationNullable (value) {
}
setEnumerationNullable (value) {
	this.enumerationNullable = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #floatingPoint32  =  0.0
		/**
  * 
  * @returns {number}
  **/
get floatingPoint32 () { return this.#floatingPoint32 }
/**
  * 
  * @type {number}
  **/
set floatingPoint32 (value) {
}
setFloatingPoint32 (value) {
	this.floatingPoint32 = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #floatingPoint32Nullable  =  undefined
		/**
  * 
  * @returns {number}
  **/
get floatingPoint32Nullable () { return this.#floatingPoint32Nullable }
/**
  * 
  * @type {number}
  **/
set floatingPoint32Nullable (value) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#floatingPoint32Nullable = parsedValue;
		}
}
setFloatingPoint32Nullable (value) {
	this.floatingPoint32Nullable = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #floatingPoint64  =  0.0
		/**
  * 
  * @returns {number}
  **/
get floatingPoint64 () { return this.#floatingPoint64 }
/**
  * 
  * @type {number}
  **/
set floatingPoint64 (value) {
}
setFloatingPoint64 (value) {
	this.floatingPoint64 = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #floatingPoint64Nullable  =  undefined
		/**
  * 
  * @returns {number}
  **/
get floatingPoint64Nullable () { return this.#floatingPoint64Nullable }
/**
  * 
  * @type {number}
  **/
set floatingPoint64Nullable (value) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#floatingPoint64Nullable = parsedValue;
		}
}
setFloatingPoint64Nullable (value) {
	this.floatingPoint64Nullable = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #integerValue  =  0
		/**
  * 
  * @returns {number}
  **/
get integerValue () { return this.#integerValue }
/**
  * 
  * @type {number}
  **/
set integerValue (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#integerValue = parsedValue;
		}
}
setIntegerValue (value) {
	this.integerValue = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #integer32ValueNullable  =  undefined
		/**
  * 
  * @returns {number}
  **/
get integer32ValueNullable () { return this.#integer32ValueNullable }
/**
  * 
  * @type {number}
  **/
set integer32ValueNullable (value) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#integer32ValueNullable = parsedValue;
		}
}
setInteger32ValueNullable (value) {
	this.integer32ValueNullable = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #integer32Value  =  0
		/**
  * 
  * @returns {number}
  **/
get integer32Value () { return this.#integer32Value }
/**
  * 
  * @type {number}
  **/
set integer32Value (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#integer32Value = parsedValue;
		}
}
setInteger32Value (value) {
	this.integer32Value = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #integer64ValueNullable  =  undefined
		/**
  * 
  * @returns {number}
  **/
get integer64ValueNullable () { return this.#integer64ValueNullable }
/**
  * 
  * @type {number}
  **/
set integer64ValueNullable (value) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#integer64ValueNullable = parsedValue;
		}
}
setInteger64ValueNullable (value) {
	this.integer64ValueNullable = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #integer64Value  =  0
		/**
  * 
  * @returns {number}
  **/
get integer64Value () { return this.#integer64Value }
/**
  * 
  * @type {number}
  **/
set integer64Value (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#integer64Value = parsedValue;
		}
}
setInteger64Value (value) {
	this.integer64Value = value
	return this
}
		/**
  * 
  * @type {{[key: string]: any}}
  **/
 #mapValue
		/**
  * 
  * @returns {{[key: string]: any}}
  **/
get mapValue () { return this.#mapValue }
/**
  * 
  * @type {{[key: string]: any}}
  **/
set mapValue (value) {
}
setMapValue (value) {
	this.mapValue = value
	return this
}
		/**
  * 
  * @type {{[key: string]: any}}
  **/
 #mapValueNullable  =  undefined
		/**
  * 
  * @returns {{[key: string]: any}}
  **/
get mapValueNullable () { return this.#mapValueNullable }
/**
  * 
  * @type {{[key: string]: any}}
  **/
set mapValueNullable (value) {
}
setMapValueNullable (value) {
	this.mapValueNullable = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #sliceValue  =  []
		/**
  * 
  * @returns {string[]}
  **/
get sliceValue () { return this.#sliceValue }
/**
  * 
  * @type {string[]}
  **/
set sliceValue (value) {
}
setSliceValue (value) {
	this.sliceValue = value
	return this
}
		/**
  * 
  * @type {any}
  **/
 #sliceValueNullable  =  undefined
		/**
  * 
  * @returns {any}
  **/
get sliceValueNullable () { return this.#sliceValueNullable }
/**
  * 
  * @type {any}
  **/
set sliceValueNullable (value) {
}
setSliceValueNullable (value) {
	this.sliceValueNullable = value
	return this
}
		/**
  * 
  * @type {GiantDto.ObjectValue}
  **/
 #objectValue
		/**
  * 
  * @returns {GiantDto.ObjectValue}
  **/
get objectValue () { return this.#objectValue }
/**
  * 
  * @type {GiantDto.ObjectValue}
  **/
set objectValue (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GiantDto.ObjectValue) {
			this.#objectValue = value
		} else {
			this.#objectValue = new GiantDto.ObjectValue(value)
		}
}
setObjectValue (value) {
	this.objectValue = value
	return this
}
/**
  * The base class definition for array
  **/
static Array = class Array {
		/**
  * 
  * @type {number}
  **/
 #subItem1  =  0
		/**
  * 
  * @returns {number}
  **/
get subItem1 () { return this.#subItem1 }
/**
  * 
  * @type {number}
  **/
set subItem1 (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#subItem1 = parsedValue;
		}
}
setSubItem1 (value) {
	this.subItem1 = value
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
			if (d.subItem1 !== undefined) { this.subItem1 = d.subItem1 }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				subItem1: this.#subItem1,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			subItem1: 'subItem1',
	  }
	}
	/**
	* Creates an instance of GiantDto.Array, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GiantDto.Array(possibleDtoObject);
	}
	/**
	* Creates an instance of GiantDto.Array, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GiantDto.Array(partialDtoObject);
	}
	copyWith(partial) {
		return new GiantDto.Array ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GiantDto.Array(this.toJSON());
	}
}
/**
  * The base class definition for arrayNullable
  **/
static ArrayNullable = class ArrayNullable {
		/**
  * 
  * @type {number}
  **/
 #subItemNullable1  =  undefined
		/**
  * 
  * @returns {number}
  **/
get subItemNullable1 () { return this.#subItemNullable1 }
/**
  * 
  * @type {number}
  **/
set subItemNullable1 (value) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#subItemNullable1 = parsedValue;
		}
}
setSubItemNullable1 (value) {
	this.subItemNullable1 = value
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
			if (d.subItemNullable1 !== undefined) { this.subItemNullable1 = d.subItemNullable1 }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				subItemNullable1: this.#subItemNullable1,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			subItemNullable1: 'subItemNullable1',
	  }
	}
	/**
	* Creates an instance of GiantDto.ArrayNullable, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GiantDto.ArrayNullable(possibleDtoObject);
	}
	/**
	* Creates an instance of GiantDto.ArrayNullable, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GiantDto.ArrayNullable(partialDtoObject);
	}
	copyWith(partial) {
		return new GiantDto.ArrayNullable ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GiantDto.ArrayNullable(this.toJSON());
	}
}
/**
  * The base class definition for objectValue
  **/
static ObjectValue = class ObjectValue {
		/**
  * 
  * @type {GiantDto.ObjectValue.InnerObject}
  **/
 #innerObject  =  undefined
		/**
  * 
  * @returns {GiantDto.ObjectValue.InnerObject}
  **/
get innerObject () { return this.#innerObject }
/**
  * 
  * @type {GiantDto.ObjectValue.InnerObject}
  **/
set innerObject (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GiantDto.ObjectValue.InnerObject) {
			this.#innerObject = value
		} else {
			this.#innerObject = new GiantDto.ObjectValue.InnerObject(value)
		}
}
setInnerObject (value) {
	this.innerObject = value
	return this
}
/**
  * The base class definition for innerObject
  **/
static InnerObject = class InnerObject {
		/**
  * 
  * @type {string}
  **/
 #innerObjText  =  ""
		/**
  * 
  * @returns {string}
  **/
get innerObjText () { return this.#innerObjText }
/**
  * 
  * @type {string}
  **/
set innerObjText (value) {
		this.#innerObjText = String(value);
}
setInnerObjText (value) {
	this.innerObjText = value
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
			if (d.innerObjText !== undefined) { this.innerObjText = d.innerObjText }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				innerObjText: this.#innerObjText,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			innerObjText: 'innerObjText',
	  }
	}
	/**
	* Creates an instance of GiantDto.ObjectValue.InnerObject, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GiantDto.ObjectValue.InnerObject(possibleDtoObject);
	}
	/**
	* Creates an instance of GiantDto.ObjectValue.InnerObject, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GiantDto.ObjectValue.InnerObject(partialDtoObject);
	}
	copyWith(partial) {
		return new GiantDto.ObjectValue.InnerObject ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GiantDto.ObjectValue.InnerObject(this.toJSON());
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
			if (d.innerObject !== undefined) { this.innerObject = d.innerObject }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				innerObject: this.#innerObject,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			innerObject$: 'innerObject',
get innerObject() {
					return withPrefix(
						"objectValue.innerObject",
						GiantDto.ObjectValue.InnerObject.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GiantDto.ObjectValue, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new GiantDto.ObjectValue(possibleDtoObject);
	}
	/**
	* Creates an instance of GiantDto.ObjectValue, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GiantDto.ObjectValue(partialDtoObject);
	}
	copyWith(partial) {
		return new GiantDto.ObjectValue ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GiantDto.ObjectValue(this.toJSON());
	}
}
	constructor(data) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
			if (d.booleanField !== undefined) { this.booleanField = d.booleanField }
			if (d.booleanFieldNullable !== undefined) { this.booleanFieldNullable = d.booleanFieldNullable }
			if (d.collectionItems !== undefined) { this.collectionItems = d.collectionItems }
			if (d.collectionItemsNullable !== undefined) { this.collectionItemsNullable = d.collectionItemsNullable }
			if (d.dateObject !== undefined) { this.dateObject = d.dateObject }
			if (d.singleRefNullable !== undefined) { this.singleRefNullable = d.singleRefNullable }
			if (d.enumeration !== undefined) { this.enumeration = d.enumeration }
			if (d.enumerationNullable !== undefined) { this.enumerationNullable = d.enumerationNullable }
			if (d.floatingPoint32 !== undefined) { this.floatingPoint32 = d.floatingPoint32 }
			if (d.floatingPoint32Nullable !== undefined) { this.floatingPoint32Nullable = d.floatingPoint32Nullable }
			if (d.floatingPoint64 !== undefined) { this.floatingPoint64 = d.floatingPoint64 }
			if (d.floatingPoint64Nullable !== undefined) { this.floatingPoint64Nullable = d.floatingPoint64Nullable }
			if (d.integerValue !== undefined) { this.integerValue = d.integerValue }
			if (d.integer32ValueNullable !== undefined) { this.integer32ValueNullable = d.integer32ValueNullable }
			if (d.integer32Value !== undefined) { this.integer32Value = d.integer32Value }
			if (d.integer64ValueNullable !== undefined) { this.integer64ValueNullable = d.integer64ValueNullable }
			if (d.integer64Value !== undefined) { this.integer64Value = d.integer64Value }
			if (d.mapValue !== undefined) { this.mapValue = d.mapValue }
			if (d.mapValueNullable !== undefined) { this.mapValueNullable = d.mapValueNullable }
			if (d.sliceValue !== undefined) { this.sliceValue = d.sliceValue }
			if (d.sliceValueNullable !== undefined) { this.sliceValueNullable = d.sliceValueNullable }
			if (d.objectValue !== undefined) { this.objectValue = d.objectValue }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.objectValue instanceof GiantDto.ObjectValue)) { this.objectValue = new GiantDto.ObjectValue(d.objectValue || {}) }	
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
				booleanField: this.#booleanField,
				booleanFieldNullable: this.#booleanFieldNullable,
				collectionItems: this.#collectionItems,
				collectionItemsNullable: this.#collectionItemsNullable,
				dateObject: this.#dateObject,
				singleRefNullable: this.#singleRefNullable,
				enumeration: this.#enumeration,
				enumerationNullable: this.#enumerationNullable,
				floatingPoint32: this.#floatingPoint32,
				floatingPoint32Nullable: this.#floatingPoint32Nullable,
				floatingPoint64: this.#floatingPoint64,
				floatingPoint64Nullable: this.#floatingPoint64Nullable,
				integerValue: this.#integerValue,
				integer32ValueNullable: this.#integer32ValueNullable,
				integer32Value: this.#integer32Value,
				integer64ValueNullable: this.#integer64ValueNullable,
				integer64Value: this.#integer64Value,
				mapValue: this.#mapValue,
				mapValueNullable: this.#mapValueNullable,
				sliceValue: this.#sliceValue,
				sliceValueNullable: this.#sliceValueNullable,
				objectValue: this.#objectValue,
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
			booleanField: 'booleanField',
			booleanFieldNullable: 'booleanFieldNullable',
			collectionItems$: 'collectionItems',
get collectionItems() {
					return withPrefix(
						"collectionItems[:i]",
						GiantDto.Fields
						);
						},
			collectionItemsNullable$: 'collectionItemsNullable',
get collectionItemsNullable() {
					return withPrefix(
						"collectionItemsNullable",
						GiantDto.Fields
						);
						},
			dateObject: 'dateObject',
			singleRefNullable: 'singleRefNullable',
			enumeration: 'enumeration',
			enumerationNullable: 'enumerationNullable',
			floatingPoint32: 'floatingPoint32',
			floatingPoint32Nullable: 'floatingPoint32Nullable',
			floatingPoint64: 'floatingPoint64',
			floatingPoint64Nullable: 'floatingPoint64Nullable',
			integerValue: 'integerValue',
			integer32ValueNullable: 'integer32ValueNullable',
			integer32Value: 'integer32Value',
			integer64ValueNullable: 'integer64ValueNullable',
			integer64Value: 'integer64Value',
			mapValue: 'mapValue',
			mapValueNullable: 'mapValueNullable',
			sliceValue$: 'sliceValue',
get sliceValue() {
					return "sliceValue[:i]";
						},
			sliceValueNullable: 'sliceValueNullable',
			objectValue$: 'objectValue',
get objectValue() {
					return withPrefix(
						"objectValue",
						GiantDto.ObjectValue.Fields
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