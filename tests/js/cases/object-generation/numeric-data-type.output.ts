	/**
  * @description The base type definition for anonymouse
  **/
	export type AnonymouseType =  {
			/**
  * @type {boolean}
  * @description bool field, non-nullable
  **/
 boolField?: boolean;
			/**
  * @type {boolean}
  * @description bool field with default
  **/
 boolFieldWithValue?: boolean;
			/**
  * @type {boolean}
  * @description nullable bool
  **/
 nullableboolField?: boolean;
			/**
  * @type {boolean}
  * @description nullable bool with default
  **/
 nullableboolFieldWithValue?: boolean;
			/**
  * @type {number}
  * @description int field, non-nullable
  **/
 intField?: number;
			/**
  * @type {number}
  * @description int field with default
  **/
 intFieldWithValue?: number;
			/**
  * @type {number}
  * @description nullable int
  **/
 nullableIntField?: number;
			/**
  * @type {number}
  * @description nullable int with default
  **/
 nullableIntFieldWithValue?: number;
			/**
  * @type {number}
  * @description int32 field, non-nullable
  **/
 int32Field?: number;
			/**
  * @type {number}
  * @description int32 with default
  **/
 int32FieldWithValue?: number;
			/**
  * @type {number}
  * @description nullable int32
  **/
 nullableInt32Field?: number;
			/**
  * @type {number}
  * @description nullable int32 with default
  **/
 nullableInt32FieldWithValue?: number;
			/**
  * @type {number}
  * @description int64 field
  **/
 int64Field?: number;
			/**
  * @type {number}
  * @description int64 with default
  **/
 int64FieldWithValue?: number;
			/**
  * @type {number}
  * @description nullable int64
  **/
 nullableInt64Field?: number;
			/**
  * @type {number}
  * @description nullable int64 with default
  **/
 nullableInt64FieldWithValue?: number;
			/**
  * @type {number}
  * @description float32 field
  **/
 float32Field?: number;
			/**
  * @type {number}
  * @description float32 with default
  **/
 float32FieldWithValue?: number;
			/**
  * @type {number}
  * @description nullable float32
  **/
 nullableFloat32Field?: number;
			/**
  * @type {number}
  * @description nullable float32 with default
  **/
 nullableFloat32FieldWithValue?: number;
			/**
  * @type {number}
  * @description float64 field
  **/
 float64Field?: number;
			/**
  * @type {number}
  * @description float64 with default
  **/
 float64FieldWithValue?: number;
			/**
  * @type {number}
  * @description nullable float64
  **/
 nullableFloat64Field?: number;
			/**
  * @type {number}
  * @description nullable float64 with default
  **/
 nullableFloat64FieldWithValue?: number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AnonymouseType {
}
/**
  * @decription The base class definition for anonymouse
  **/
export class Anonymouse {
	constructor(data: unknown) {
		// This probably doesn't cover the nested objects
		const d = data as Partial<Anonymouse>;
			if (d[`boolField`] !== undefined) { 
 this.setBoolField (d[`boolField`]) 
}
			if (d[`boolFieldWithValue`] !== undefined) { 
 this.setBoolFieldWithValue (d[`boolFieldWithValue`]) 
}
			if (d[`nullableboolField`] !== undefined) { 
 this.setNullableboolField (d[`nullableboolField`]) 
}
			if (d[`nullableboolFieldWithValue`] !== undefined) { 
 this.setNullableboolFieldWithValue (d[`nullableboolFieldWithValue`]) 
}
			if (d[`intField`] !== undefined) { 
 this.setIntField (d[`intField`]) 
}
			if (d[`intFieldWithValue`] !== undefined) { 
 this.setIntFieldWithValue (d[`intFieldWithValue`]) 
}
			if (d[`nullableIntField`] !== undefined) { 
 this.setNullableIntField (d[`nullableIntField`]) 
}
			if (d[`nullableIntFieldWithValue`] !== undefined) { 
 this.setNullableIntFieldWithValue (d[`nullableIntFieldWithValue`]) 
}
			if (d[`int32Field`] !== undefined) { 
 this.setInt32Field (d[`int32Field`]) 
}
			if (d[`int32FieldWithValue`] !== undefined) { 
 this.setInt32FieldWithValue (d[`int32FieldWithValue`]) 
}
			if (d[`nullableInt32Field`] !== undefined) { 
 this.setNullableInt32Field (d[`nullableInt32Field`]) 
}
			if (d[`nullableInt32FieldWithValue`] !== undefined) { 
 this.setNullableInt32FieldWithValue (d[`nullableInt32FieldWithValue`]) 
}
			if (d[`int64Field`] !== undefined) { 
 this.setInt64Field (d[`int64Field`]) 
}
			if (d[`int64FieldWithValue`] !== undefined) { 
 this.setInt64FieldWithValue (d[`int64FieldWithValue`]) 
}
			if (d[`nullableInt64Field`] !== undefined) { 
 this.setNullableInt64Field (d[`nullableInt64Field`]) 
}
			if (d[`nullableInt64FieldWithValue`] !== undefined) { 
 this.setNullableInt64FieldWithValue (d[`nullableInt64FieldWithValue`]) 
}
			if (d[`float32Field`] !== undefined) { 
 this.setFloat32Field (d[`float32Field`]) 
}
			if (d[`float32FieldWithValue`] !== undefined) { 
 this.setFloat32FieldWithValue (d[`float32FieldWithValue`]) 
}
			if (d[`nullableFloat32Field`] !== undefined) { 
 this.setNullableFloat32Field (d[`nullableFloat32Field`]) 
}
			if (d[`nullableFloat32FieldWithValue`] !== undefined) { 
 this.setNullableFloat32FieldWithValue (d[`nullableFloat32FieldWithValue`]) 
}
			if (d[`float64Field`] !== undefined) { 
 this.setFloat64Field (d[`float64Field`]) 
}
			if (d[`float64FieldWithValue`] !== undefined) { 
 this.setFloat64FieldWithValue (d[`float64FieldWithValue`]) 
}
			if (d[`nullableFloat64Field`] !== undefined) { 
 this.setNullableFloat64Field (d[`nullableFloat64Field`]) 
}
			if (d[`nullableFloat64FieldWithValue`] !== undefined) { 
 this.setNullableFloat64FieldWithValue (d[`nullableFloat64FieldWithValue`]) 
}
	}
		/**
  * bool field, non-nullable
  * @type {boolean}
  **/
 boolField: boolean = null
		/**
  * @returns {boolean}
  * @description bool field, non-nullable
  **/
getBoolField () { return this[`boolField`] }
		/**
  * bool field, non-nullable
  * @param {boolean}
  **/
setBoolField (value: boolean ) { this[`boolField`] = value; return this; } 
		/**
  * bool field with default
  * @type {boolean}
  **/
 boolFieldWithValue: boolean = true
		/**
  * @returns {boolean}
  * @description bool field with default
  **/
getBoolFieldWithValue () { return this[`boolFieldWithValue`] }
		/**
  * bool field with default
  * @param {boolean}
  **/
setBoolFieldWithValue (value: boolean ) { this[`boolFieldWithValue`] = value; return this; } 
		/**
  * nullable bool
  * @type {boolean}
  **/
 nullableboolField?: boolean | null = undefined
		/**
  * @returns {boolean}
  * @description nullable bool
  **/
getNullableboolField () { return this[`nullableboolField`] }
		/**
  * nullable bool
  * @param {boolean}
  **/
setNullableboolField (value: boolean  | null) { this[`nullableboolField`] = value; return this; } 
		/**
  * nullable bool with default
  * @type {boolean}
  **/
 nullableboolFieldWithValue?: boolean | null  = true
		/**
  * @returns {boolean}
  * @description nullable bool with default
  **/
getNullableboolFieldWithValue () { return this[`nullableboolFieldWithValue`] }
		/**
  * nullable bool with default
  * @param {boolean}
  **/
setNullableboolFieldWithValue (value: boolean  | null) { this[`nullableboolFieldWithValue`] = value; return this; } 
		/**
  * int field, non-nullable
  * @type {number}
  **/
 intField: number = 0
		/**
  * @returns {number}
  * @description int field, non-nullable
  **/
getIntField () { return this[`intField`] }
		/**
  * int field, non-nullable
  * @param {number}
  **/
setIntField (value: number ) { this[`intField`] = value; return this; } 
		/**
  * int field with default
  * @type {number}
  **/
 intFieldWithValue: number = 42
		/**
  * @returns {number}
  * @description int field with default
  **/
getIntFieldWithValue () { return this[`intFieldWithValue`] }
		/**
  * int field with default
  * @param {number}
  **/
setIntFieldWithValue (value: number ) { this[`intFieldWithValue`] = value; return this; } 
		/**
  * nullable int
  * @type {number}
  **/
 nullableIntField?: number | null = undefined
		/**
  * @returns {number}
  * @description nullable int
  **/
getNullableIntField () { return this[`nullableIntField`] }
		/**
  * nullable int
  * @param {number}
  **/
setNullableIntField (value: number  | null) { this[`nullableIntField`] = value; return this; } 
		/**
  * nullable int with default
  * @type {number}
  **/
 nullableIntFieldWithValue?: number | null  = 7
		/**
  * @returns {number}
  * @description nullable int with default
  **/
getNullableIntFieldWithValue () { return this[`nullableIntFieldWithValue`] }
		/**
  * nullable int with default
  * @param {number}
  **/
setNullableIntFieldWithValue (value: number  | null) { this[`nullableIntFieldWithValue`] = value; return this; } 
		/**
  * int32 field, non-nullable
  * @type {number}
  **/
 int32Field: number = 0
		/**
  * @returns {number}
  * @description int32 field, non-nullable
  **/
getInt32Field () { return this[`int32Field`] }
		/**
  * int32 field, non-nullable
  * @param {number}
  **/
setInt32Field (value: number ) { this[`int32Field`] = value; return this; } 
		/**
  * int32 with default
  * @type {number}
  **/
 int32FieldWithValue: number = 100
		/**
  * @returns {number}
  * @description int32 with default
  **/
getInt32FieldWithValue () { return this[`int32FieldWithValue`] }
		/**
  * int32 with default
  * @param {number}
  **/
setInt32FieldWithValue (value: number ) { this[`int32FieldWithValue`] = value; return this; } 
		/**
  * nullable int32
  * @type {number}
  **/
 nullableInt32Field?: number | null = undefined
		/**
  * @returns {number}
  * @description nullable int32
  **/
getNullableInt32Field () { return this[`nullableInt32Field`] }
		/**
  * nullable int32
  * @param {number}
  **/
setNullableInt32Field (value: number  | null) { this[`nullableInt32Field`] = value; return this; } 
		/**
  * nullable int32 with default
  * @type {number}
  **/
 nullableInt32FieldWithValue?: number | null  = 200
		/**
  * @returns {number}
  * @description nullable int32 with default
  **/
getNullableInt32FieldWithValue () { return this[`nullableInt32FieldWithValue`] }
		/**
  * nullable int32 with default
  * @param {number}
  **/
setNullableInt32FieldWithValue (value: number  | null) { this[`nullableInt32FieldWithValue`] = value; return this; } 
		/**
  * int64 field
  * @type {number}
  **/
 int64Field: number = 0
		/**
  * @returns {number}
  * @description int64 field
  **/
getInt64Field () { return this[`int64Field`] }
		/**
  * int64 field
  * @param {number}
  **/
setInt64Field (value: number ) { this[`int64Field`] = value; return this; } 
		/**
  * int64 with default
  * @type {number}
  **/
 int64FieldWithValue?: number | null  = 123
		/**
  * @returns {number}
  * @description int64 with default
  **/
getInt64FieldWithValue () { return this[`int64FieldWithValue`] }
		/**
  * int64 with default
  * @param {number}
  **/
setInt64FieldWithValue (value: number  | null) { this[`int64FieldWithValue`] = value; return this; } 
		/**
  * nullable int64
  * @type {number}
  **/
 nullableInt64Field?: number | null = undefined
		/**
  * @returns {number}
  * @description nullable int64
  **/
getNullableInt64Field () { return this[`nullableInt64Field`] }
		/**
  * nullable int64
  * @param {number}
  **/
setNullableInt64Field (value: number  | null) { this[`nullableInt64Field`] = value; return this; } 
		/**
  * nullable int64 with default
  * @type {number}
  **/
 nullableInt64FieldWithValue?: number | null  = 456
		/**
  * @returns {number}
  * @description nullable int64 with default
  **/
getNullableInt64FieldWithValue () { return this[`nullableInt64FieldWithValue`] }
		/**
  * nullable int64 with default
  * @param {number}
  **/
setNullableInt64FieldWithValue (value: number  | null) { this[`nullableInt64FieldWithValue`] = value; return this; } 
		/**
  * float32 field
  * @type {number}
  **/
 float32Field: number = 0.0
		/**
  * @returns {number}
  * @description float32 field
  **/
getFloat32Field () { return this[`float32Field`] }
		/**
  * float32 field
  * @param {number}
  **/
setFloat32Field (value: number ) { this[`float32Field`] = value; return this; } 
		/**
  * float32 with default
  * @type {number}
  **/
 float32FieldWithValue: number = 1.23
		/**
  * @returns {number}
  * @description float32 with default
  **/
getFloat32FieldWithValue () { return this[`float32FieldWithValue`] }
		/**
  * float32 with default
  * @param {number}
  **/
setFloat32FieldWithValue (value: number ) { this[`float32FieldWithValue`] = value; return this; } 
		/**
  * nullable float32
  * @type {number}
  **/
 nullableFloat32Field?: number | null = undefined
		/**
  * @returns {number}
  * @description nullable float32
  **/
getNullableFloat32Field () { return this[`nullableFloat32Field`] }
		/**
  * nullable float32
  * @param {number}
  **/
setNullableFloat32Field (value: number  | null) { this[`nullableFloat32Field`] = value; return this; } 
		/**
  * nullable float32 with default
  * @type {number}
  **/
 nullableFloat32FieldWithValue?: number | null  = 4.56
		/**
  * @returns {number}
  * @description nullable float32 with default
  **/
getNullableFloat32FieldWithValue () { return this[`nullableFloat32FieldWithValue`] }
		/**
  * nullable float32 with default
  * @param {number}
  **/
setNullableFloat32FieldWithValue (value: number  | null) { this[`nullableFloat32FieldWithValue`] = value; return this; } 
		/**
  * float64 field
  * @type {number}
  **/
 float64Field: number = 0.0
		/**
  * @returns {number}
  * @description float64 field
  **/
getFloat64Field () { return this[`float64Field`] }
		/**
  * float64 field
  * @param {number}
  **/
setFloat64Field (value: number ) { this[`float64Field`] = value; return this; } 
		/**
  * float64 with default
  * @type {number}
  **/
 float64FieldWithValue: number = 7.89
		/**
  * @returns {number}
  * @description float64 with default
  **/
getFloat64FieldWithValue () { return this[`float64FieldWithValue`] }
		/**
  * float64 with default
  * @param {number}
  **/
setFloat64FieldWithValue (value: number ) { this[`float64FieldWithValue`] = value; return this; } 
		/**
  * nullable float64
  * @type {number}
  **/
 nullableFloat64Field?: number | null = undefined
		/**
  * @returns {number}
  * @description nullable float64
  **/
getNullableFloat64Field () { return this[`nullableFloat64Field`] }
		/**
  * nullable float64
  * @param {number}
  **/
setNullableFloat64Field (value: number  | null) { this[`nullableFloat64Field`] = value; return this; } 
		/**
  * nullable float64 with default
  * @type {number}
  **/
 nullableFloat64FieldWithValue?: number | null  = 0.12
		/**
  * @returns {number}
  * @description nullable float64 with default
  **/
getNullableFloat64FieldWithValue () { return this[`nullableFloat64FieldWithValue`] }
		/**
  * nullable float64 with default
  * @param {number}
  **/
setNullableFloat64FieldWithValue (value: number  | null) { this[`nullableFloat64FieldWithValue`] = value; return this; } 
}
export abstract class AnonymouseFactory {
	abstract create(data: unknown): Anonymouse;
}