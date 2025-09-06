
/**
  * The base class definition for anonymouse
  **/
export class Anonymouse {
		/**
  * string field, non-nullable
  * @type {string}
  **/
 #stringField : string  =  ""
		/**
  * string field, non-nullable
  * @returns {string}
  **/
get stringField () { return this.#stringField }
/**
  * string field, non-nullable
  * @type {string}
  **/
set stringField (value: string) {
	 	const correctType = typeof value === 'string';
		this.#stringField = correctType ? value : ('' + value);
}
setStringField (value: string) {
	this.stringField = value
	return this
}
		/**
  * string field with default
  * @type {string}
  **/
 #stringFieldWithValue : string  =  "defaultstring"
		/**
  * string field with default
  * @returns {string}
  **/
get stringFieldWithValue () { return this.#stringFieldWithValue }
/**
  * string field with default
  * @type {string}
  **/
set stringFieldWithValue (value: string) {
	 	const correctType = typeof value === 'string';
		this.#stringFieldWithValue = correctType ? value : ('' + value);
}
setStringFieldWithValue (value: string) {
	this.stringFieldWithValue = value
	return this
}
		/**
  * nullable string
  * @type {string}
  **/
 #nullablestringField ? : string  | null  =  undefined
		/**
  * nullable string
  * @returns {string}
  **/
get nullablestringField () { return this.#nullablestringField }
/**
  * nullable string
  * @type {string}
  **/
set nullablestringField (value: string) {
	 	const correctType = typeof value === 'string' || value === undefined || value === null
		this.#nullablestringField = correctType ? value : ('' + value);
}
setNullablestringField (value: string) {
	this.nullablestringField = value
	return this
}
		/**
  * nullable string with default
  * @type {string}
  **/
 #nullablestringFieldWithValue ? : string  | null  =  "defaultstring"
		/**
  * nullable string with default
  * @returns {string}
  **/
get nullablestringFieldWithValue () { return this.#nullablestringFieldWithValue }
/**
  * nullable string with default
  * @type {string}
  **/
set nullablestringFieldWithValue (value: string) {
	 	const correctType = typeof value === 'string' || value === undefined || value === null
		this.#nullablestringFieldWithValue = correctType ? value : ('' + value);
}
setNullablestringFieldWithValue (value: string) {
	this.nullablestringFieldWithValue = value
	return this
}
		/**
  * bool field, non-nullable
  * @type {boolean}
  **/
 #boolField : boolean  =  null
		/**
  * bool field, non-nullable
  * @returns {boolean}
  **/
get boolField () { return this.#boolField }
/**
  * bool field, non-nullable
  * @type {boolean}
  **/
set boolField (value: boolean) {
	 	const correctType = value === true || value === false
		this.#boolField = correctType ? value : Boolean(value);
}
setBoolField (value: boolean) {
	this.boolField = value
	return this
}
		/**
  * bool field with default
  * @type {boolean}
  **/
 #boolFieldWithValue : boolean  =  true
		/**
  * bool field with default
  * @returns {boolean}
  **/
get boolFieldWithValue () { return this.#boolFieldWithValue }
/**
  * bool field with default
  * @type {boolean}
  **/
set boolFieldWithValue (value: boolean) {
	 	const correctType = value === true || value === false
		this.#boolFieldWithValue = correctType ? value : Boolean(value);
}
setBoolFieldWithValue (value: boolean) {
	this.boolFieldWithValue = value
	return this
}
		/**
  * nullable bool
  * @type {boolean}
  **/
 #nullableboolField ? : boolean  | null  =  undefined
		/**
  * nullable bool
  * @returns {boolean}
  **/
get nullableboolField () { return this.#nullableboolField }
/**
  * nullable bool
  * @type {boolean}
  **/
set nullableboolField (value: boolean) {
	 	const correctType = value === true || value === false || value === undefined || value === null
		this.#nullableboolField = correctType ? value : Boolean(value);
}
setNullableboolField (value: boolean) {
	this.nullableboolField = value
	return this
}
		/**
  * nullable bool with default
  * @type {boolean}
  **/
 #nullableboolFieldWithValue ? : boolean  | null  =  true
		/**
  * nullable bool with default
  * @returns {boolean}
  **/
get nullableboolFieldWithValue () { return this.#nullableboolFieldWithValue }
/**
  * nullable bool with default
  * @type {boolean}
  **/
set nullableboolFieldWithValue (value: boolean) {
	 	const correctType = value === true || value === false || value === undefined || value === null
		this.#nullableboolFieldWithValue = correctType ? value : Boolean(value);
}
setNullableboolFieldWithValue (value: boolean) {
	this.nullableboolFieldWithValue = value
	return this
}
		/**
  * int field, non-nullable
  * @type {number}
  **/
 #intField : number  =  0
		/**
  * int field, non-nullable
  * @returns {number}
  **/
get intField () { return this.#intField }
/**
  * int field, non-nullable
  * @type {number}
  **/
set intField (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#intField = parsedValue;
		}
}
setIntField (value: number) {
	this.intField = value
	return this
}
		/**
  * int field with default
  * @type {number}
  **/
 #intFieldWithValue : number  =  42
		/**
  * int field with default
  * @returns {number}
  **/
get intFieldWithValue () { return this.#intFieldWithValue }
/**
  * int field with default
  * @type {number}
  **/
set intFieldWithValue (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#intFieldWithValue = parsedValue;
		}
}
setIntFieldWithValue (value: number) {
	this.intFieldWithValue = value
	return this
}
		/**
  * nullable int
  * @type {number}
  **/
 #nullableIntField ? : number  | null  =  undefined
		/**
  * nullable int
  * @returns {number}
  **/
get nullableIntField () { return this.#nullableIntField }
/**
  * nullable int
  * @type {number}
  **/
set nullableIntField (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableIntField = parsedValue;
		}
}
setNullableIntField (value: number) {
	this.nullableIntField = value
	return this
}
		/**
  * nullable int with default
  * @type {number}
  **/
 #nullableIntFieldWithValue ? : number  | null  =  7
		/**
  * nullable int with default
  * @returns {number}
  **/
get nullableIntFieldWithValue () { return this.#nullableIntFieldWithValue }
/**
  * nullable int with default
  * @type {number}
  **/
set nullableIntFieldWithValue (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableIntFieldWithValue = parsedValue;
		}
}
setNullableIntFieldWithValue (value: number) {
	this.nullableIntFieldWithValue = value
	return this
}
		/**
  * int32 field, non-nullable
  * @type {number}
  **/
 #int32Field : number  =  0
		/**
  * int32 field, non-nullable
  * @returns {number}
  **/
get int32Field () { return this.#int32Field }
/**
  * int32 field, non-nullable
  * @type {number}
  **/
set int32Field (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#int32Field = parsedValue;
		}
}
setInt32Field (value: number) {
	this.int32Field = value
	return this
}
		/**
  * int32 with default
  * @type {number}
  **/
 #int32FieldWithValue : number  =  100
		/**
  * int32 with default
  * @returns {number}
  **/
get int32FieldWithValue () { return this.#int32FieldWithValue }
/**
  * int32 with default
  * @type {number}
  **/
set int32FieldWithValue (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#int32FieldWithValue = parsedValue;
		}
}
setInt32FieldWithValue (value: number) {
	this.int32FieldWithValue = value
	return this
}
		/**
  * nullable int32
  * @type {number}
  **/
 #nullableInt32Field ? : number  | null  =  undefined
		/**
  * nullable int32
  * @returns {number}
  **/
get nullableInt32Field () { return this.#nullableInt32Field }
/**
  * nullable int32
  * @type {number}
  **/
set nullableInt32Field (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableInt32Field = parsedValue;
		}
}
setNullableInt32Field (value: number) {
	this.nullableInt32Field = value
	return this
}
		/**
  * nullable int32 with default
  * @type {number}
  **/
 #nullableInt32FieldWithValue ? : number  | null  =  200
		/**
  * nullable int32 with default
  * @returns {number}
  **/
get nullableInt32FieldWithValue () { return this.#nullableInt32FieldWithValue }
/**
  * nullable int32 with default
  * @type {number}
  **/
set nullableInt32FieldWithValue (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableInt32FieldWithValue = parsedValue;
		}
}
setNullableInt32FieldWithValue (value: number) {
	this.nullableInt32FieldWithValue = value
	return this
}
		/**
  * int64 field
  * @type {number}
  **/
 #int64Field : number  =  0
		/**
  * int64 field
  * @returns {number}
  **/
get int64Field () { return this.#int64Field }
/**
  * int64 field
  * @type {number}
  **/
set int64Field (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#int64Field = parsedValue;
		}
}
setInt64Field (value: number) {
	this.int64Field = value
	return this
}
		/**
  * int64 with default
  * @type {number}
  **/
 #int64FieldWithValue ? : number  | null  =  123
		/**
  * int64 with default
  * @returns {number}
  **/
get int64FieldWithValue () { return this.#int64FieldWithValue }
/**
  * int64 with default
  * @type {number}
  **/
set int64FieldWithValue (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#int64FieldWithValue = parsedValue;
		}
}
setInt64FieldWithValue (value: number) {
	this.int64FieldWithValue = value
	return this
}
		/**
  * nullable int64
  * @type {number}
  **/
 #nullableInt64Field ? : number  | null  =  undefined
		/**
  * nullable int64
  * @returns {number}
  **/
get nullableInt64Field () { return this.#nullableInt64Field }
/**
  * nullable int64
  * @type {number}
  **/
set nullableInt64Field (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableInt64Field = parsedValue;
		}
}
setNullableInt64Field (value: number) {
	this.nullableInt64Field = value
	return this
}
		/**
  * nullable int64 with default
  * @type {number}
  **/
 #nullableInt64FieldWithValue ? : number  | null  =  456
		/**
  * nullable int64 with default
  * @returns {number}
  **/
get nullableInt64FieldWithValue () { return this.#nullableInt64FieldWithValue }
/**
  * nullable int64 with default
  * @type {number}
  **/
set nullableInt64FieldWithValue (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableInt64FieldWithValue = parsedValue;
		}
}
setNullableInt64FieldWithValue (value: number) {
	this.nullableInt64FieldWithValue = value
	return this
}
		/**
  * float32 field
  * @type {number}
  **/
 #float32Field : number  =  0.0
		/**
  * float32 field
  * @returns {number}
  **/
get float32Field () { return this.#float32Field }
/**
  * float32 field
  * @type {number}
  **/
set float32Field (value: number) {
}
setFloat32Field (value: number) {
	this.float32Field = value
	return this
}
		/**
  * float32 with default
  * @type {number}
  **/
 #float32FieldWithValue : number  =  1.23
		/**
  * float32 with default
  * @returns {number}
  **/
get float32FieldWithValue () { return this.#float32FieldWithValue }
/**
  * float32 with default
  * @type {number}
  **/
set float32FieldWithValue (value: number) {
}
setFloat32FieldWithValue (value: number) {
	this.float32FieldWithValue = value
	return this
}
		/**
  * nullable float32
  * @type {number}
  **/
 #nullableFloat32Field ? : number  | null  =  undefined
		/**
  * nullable float32
  * @returns {number}
  **/
get nullableFloat32Field () { return this.#nullableFloat32Field }
/**
  * nullable float32
  * @type {number}
  **/
set nullableFloat32Field (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableFloat32Field = parsedValue;
		}
}
setNullableFloat32Field (value: number) {
	this.nullableFloat32Field = value
	return this
}
		/**
  * nullable float32 with default
  * @type {number}
  **/
 #nullableFloat32FieldWithValue ? : number  | null  =  4.56
		/**
  * nullable float32 with default
  * @returns {number}
  **/
get nullableFloat32FieldWithValue () { return this.#nullableFloat32FieldWithValue }
/**
  * nullable float32 with default
  * @type {number}
  **/
set nullableFloat32FieldWithValue (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableFloat32FieldWithValue = parsedValue;
		}
}
setNullableFloat32FieldWithValue (value: number) {
	this.nullableFloat32FieldWithValue = value
	return this
}
		/**
  * float64 field
  * @type {number}
  **/
 #float64Field : number  =  0.0
		/**
  * float64 field
  * @returns {number}
  **/
get float64Field () { return this.#float64Field }
/**
  * float64 field
  * @type {number}
  **/
set float64Field (value: number) {
}
setFloat64Field (value: number) {
	this.float64Field = value
	return this
}
		/**
  * float64 with default
  * @type {number}
  **/
 #float64FieldWithValue : number  =  7.89
		/**
  * float64 with default
  * @returns {number}
  **/
get float64FieldWithValue () { return this.#float64FieldWithValue }
/**
  * float64 with default
  * @type {number}
  **/
set float64FieldWithValue (value: number) {
}
setFloat64FieldWithValue (value: number) {
	this.float64FieldWithValue = value
	return this
}
		/**
  * nullable float64
  * @type {number}
  **/
 #nullableFloat64Field ? : number  | null  =  undefined
		/**
  * nullable float64
  * @returns {number}
  **/
get nullableFloat64Field () { return this.#nullableFloat64Field }
/**
  * nullable float64
  * @type {number}
  **/
set nullableFloat64Field (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableFloat64Field = parsedValue;
		}
}
setNullableFloat64Field (value: number) {
	this.nullableFloat64Field = value
	return this
}
		/**
  * nullable float64 with default
  * @type {number}
  **/
 #nullableFloat64FieldWithValue ? : number  | null  =  0.12
		/**
  * nullable float64 with default
  * @returns {number}
  **/
get nullableFloat64FieldWithValue () { return this.#nullableFloat64FieldWithValue }
/**
  * nullable float64 with default
  * @type {number}
  **/
set nullableFloat64FieldWithValue (value: number) {
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#nullableFloat64FieldWithValue = parsedValue;
		}
}
setNullableFloat64FieldWithValue (value: number) {
	this.nullableFloat64FieldWithValue = value
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
		const d = data as Partial<Anonymouse>;
			if (d.stringField !== undefined) { this.stringField = d.stringField }
			if (d.stringFieldWithValue !== undefined) { this.stringFieldWithValue = d.stringFieldWithValue }
			if (d.nullablestringField !== undefined) { this.nullablestringField = d.nullablestringField }
			if (d.nullablestringFieldWithValue !== undefined) { this.nullablestringFieldWithValue = d.nullablestringFieldWithValue }
			if (d.boolField !== undefined) { this.boolField = d.boolField }
			if (d.boolFieldWithValue !== undefined) { this.boolFieldWithValue = d.boolFieldWithValue }
			if (d.nullableboolField !== undefined) { this.nullableboolField = d.nullableboolField }
			if (d.nullableboolFieldWithValue !== undefined) { this.nullableboolFieldWithValue = d.nullableboolFieldWithValue }
			if (d.intField !== undefined) { this.intField = d.intField }
			if (d.intFieldWithValue !== undefined) { this.intFieldWithValue = d.intFieldWithValue }
			if (d.nullableIntField !== undefined) { this.nullableIntField = d.nullableIntField }
			if (d.nullableIntFieldWithValue !== undefined) { this.nullableIntFieldWithValue = d.nullableIntFieldWithValue }
			if (d.int32Field !== undefined) { this.int32Field = d.int32Field }
			if (d.int32FieldWithValue !== undefined) { this.int32FieldWithValue = d.int32FieldWithValue }
			if (d.nullableInt32Field !== undefined) { this.nullableInt32Field = d.nullableInt32Field }
			if (d.nullableInt32FieldWithValue !== undefined) { this.nullableInt32FieldWithValue = d.nullableInt32FieldWithValue }
			if (d.int64Field !== undefined) { this.int64Field = d.int64Field }
			if (d.int64FieldWithValue !== undefined) { this.int64FieldWithValue = d.int64FieldWithValue }
			if (d.nullableInt64Field !== undefined) { this.nullableInt64Field = d.nullableInt64Field }
			if (d.nullableInt64FieldWithValue !== undefined) { this.nullableInt64FieldWithValue = d.nullableInt64FieldWithValue }
			if (d.float32Field !== undefined) { this.float32Field = d.float32Field }
			if (d.float32FieldWithValue !== undefined) { this.float32FieldWithValue = d.float32FieldWithValue }
			if (d.nullableFloat32Field !== undefined) { this.nullableFloat32Field = d.nullableFloat32Field }
			if (d.nullableFloat32FieldWithValue !== undefined) { this.nullableFloat32FieldWithValue = d.nullableFloat32FieldWithValue }
			if (d.float64Field !== undefined) { this.float64Field = d.float64Field }
			if (d.float64FieldWithValue !== undefined) { this.float64FieldWithValue = d.float64FieldWithValue }
			if (d.nullableFloat64Field !== undefined) { this.nullableFloat64Field = d.nullableFloat64Field }
			if (d.nullableFloat64FieldWithValue !== undefined) { this.nullableFloat64FieldWithValue = d.nullableFloat64FieldWithValue }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				stringField: this.#stringField,
				stringFieldWithValue: this.#stringFieldWithValue,
				nullablestringField: this.#nullablestringField,
				nullablestringFieldWithValue: this.#nullablestringFieldWithValue,
				boolField: this.#boolField,
				boolFieldWithValue: this.#boolFieldWithValue,
				nullableboolField: this.#nullableboolField,
				nullableboolFieldWithValue: this.#nullableboolFieldWithValue,
				intField: this.#intField,
				intFieldWithValue: this.#intFieldWithValue,
				nullableIntField: this.#nullableIntField,
				nullableIntFieldWithValue: this.#nullableIntFieldWithValue,
				int32Field: this.#int32Field,
				int32FieldWithValue: this.#int32FieldWithValue,
				nullableInt32Field: this.#nullableInt32Field,
				nullableInt32FieldWithValue: this.#nullableInt32FieldWithValue,
				int64Field: this.#int64Field,
				int64FieldWithValue: this.#int64FieldWithValue,
				nullableInt64Field: this.#nullableInt64Field,
				nullableInt64FieldWithValue: this.#nullableInt64FieldWithValue,
				float32Field: this.#float32Field,
				float32FieldWithValue: this.#float32FieldWithValue,
				nullableFloat32Field: this.#nullableFloat32Field,
				nullableFloat32FieldWithValue: this.#nullableFloat32FieldWithValue,
				float64Field: this.#float64Field,
				float64FieldWithValue: this.#float64FieldWithValue,
				nullableFloat64Field: this.#nullableFloat64Field,
				nullableFloat64FieldWithValue: this.#nullableFloat64FieldWithValue,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			stringField: 'stringField',
			stringFieldWithValue: 'stringFieldWithValue',
			nullablestringField: 'nullablestringField',
			nullablestringFieldWithValue: 'nullablestringFieldWithValue',
			boolField: 'boolField',
			boolFieldWithValue: 'boolFieldWithValue',
			nullableboolField: 'nullableboolField',
			nullableboolFieldWithValue: 'nullableboolFieldWithValue',
			intField: 'intField',
			intFieldWithValue: 'intFieldWithValue',
			nullableIntField: 'nullableIntField',
			nullableIntFieldWithValue: 'nullableIntFieldWithValue',
			int32Field: 'int32Field',
			int32FieldWithValue: 'int32FieldWithValue',
			nullableInt32Field: 'nullableInt32Field',
			nullableInt32FieldWithValue: 'nullableInt32FieldWithValue',
			int64Field: 'int64Field',
			int64FieldWithValue: 'int64FieldWithValue',
			nullableInt64Field: 'nullableInt64Field',
			nullableInt64FieldWithValue: 'nullableInt64FieldWithValue',
			float32Field: 'float32Field',
			float32FieldWithValue: 'float32FieldWithValue',
			nullableFloat32Field: 'nullableFloat32Field',
			nullableFloat32FieldWithValue: 'nullableFloat32FieldWithValue',
			float64Field: 'float64Field',
			float64FieldWithValue: 'float64FieldWithValue',
			nullableFloat64Field: 'nullableFloat64Field',
			nullableFloat64FieldWithValue: 'nullableFloat64FieldWithValue',
	  }
	}
}
export abstract class AnonymouseFactory {
	abstract create(data: unknown): Anonymouse;
}
	/**
  * The base type definition for anonymouse
  **/
	export type AnonymouseType =  {
			/**
  * string field, non-nullable
  * @type {string}
  **/
 stringField?: string;
			/**
  * string field with default
  * @type {string}
  **/
 stringFieldWithValue?: string;
			/**
  * nullable string
  * @type {string}
  **/
 nullablestringField?: string;
			/**
  * nullable string with default
  * @type {string}
  **/
 nullablestringFieldWithValue?: string;
			/**
  * bool field, non-nullable
  * @type {boolean}
  **/
 boolField?: boolean;
			/**
  * bool field with default
  * @type {boolean}
  **/
 boolFieldWithValue?: boolean;
			/**
  * nullable bool
  * @type {boolean}
  **/
 nullableboolField?: boolean;
			/**
  * nullable bool with default
  * @type {boolean}
  **/
 nullableboolFieldWithValue?: boolean;
			/**
  * int field, non-nullable
  * @type {number}
  **/
 intField?: number;
			/**
  * int field with default
  * @type {number}
  **/
 intFieldWithValue?: number;
			/**
  * nullable int
  * @type {number}
  **/
 nullableIntField?: number;
			/**
  * nullable int with default
  * @type {number}
  **/
 nullableIntFieldWithValue?: number;
			/**
  * int32 field, non-nullable
  * @type {number}
  **/
 int32Field?: number;
			/**
  * int32 with default
  * @type {number}
  **/
 int32FieldWithValue?: number;
			/**
  * nullable int32
  * @type {number}
  **/
 nullableInt32Field?: number;
			/**
  * nullable int32 with default
  * @type {number}
  **/
 nullableInt32FieldWithValue?: number;
			/**
  * int64 field
  * @type {number}
  **/
 int64Field?: number;
			/**
  * int64 with default
  * @type {number}
  **/
 int64FieldWithValue?: number;
			/**
  * nullable int64
  * @type {number}
  **/
 nullableInt64Field?: number;
			/**
  * nullable int64 with default
  * @type {number}
  **/
 nullableInt64FieldWithValue?: number;
			/**
  * float32 field
  * @type {number}
  **/
 float32Field?: number;
			/**
  * float32 with default
  * @type {number}
  **/
 float32FieldWithValue?: number;
			/**
  * nullable float32
  * @type {number}
  **/
 nullableFloat32Field?: number;
			/**
  * nullable float32 with default
  * @type {number}
  **/
 nullableFloat32FieldWithValue?: number;
			/**
  * float64 field
  * @type {number}
  **/
 float64Field?: number;
			/**
  * float64 with default
  * @type {number}
  **/
 float64FieldWithValue?: number;
			/**
  * nullable float64
  * @type {number}
  **/
 nullableFloat64Field?: number;
			/**
  * nullable float64 with default
  * @type {number}
  **/
 nullableFloat64FieldWithValue?: number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AnonymouseType {
}
  
  function isPlausibleObject(v: any) { return false }
  