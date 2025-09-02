	/**
  * The base type definition for anonymouse
  **/
	export type AnonymouseType =  {
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
/**
  * The base class definition for anonymouse
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
  * bool field, non-nullable
  * @returns {boolean}
  **/
getBoolField () { return this[`boolField`] }
		/**
  * bool field, non-nullable
  * @param {boolean}
  **/
/**
  * bool field, non-nullable
  * @param {boolean}
  **/
/// XXX
setBoolField(value: boolean) {
	// Only accept array types
	this["boolField"] = value
	return this
}
		/**
  * bool field with default
  * @type {boolean}
  **/
 boolFieldWithValue: boolean = true
		/**
  * bool field with default
  * @returns {boolean}
  **/
getBoolFieldWithValue () { return this[`boolFieldWithValue`] }
		/**
  * bool field with default
  * @param {boolean}
  **/
/**
  * bool field with default
  * @param {boolean}
  **/
/// XXX
setBoolFieldWithValue(value: boolean) {
	// Only accept array types
	this["boolFieldWithValue"] = value
	return this
}
		/**
  * nullable bool
  * @type {boolean}
  **/
 nullableboolField?: boolean | null = undefined
		/**
  * nullable bool
  * @returns {boolean}
  **/
getNullableboolField () { return this[`nullableboolField`] }
		/**
  * nullable bool
  * @param {boolean}
  **/
/**
  * nullable bool
  * @param {boolean}
  **/
/// XXX
setNullableboolField(value: boolean) {
	// Only accept array types
	this["nullableboolField"] = value
	return this
}
		/**
  * nullable bool with default
  * @type {boolean}
  **/
 nullableboolFieldWithValue?: boolean | null  = true
		/**
  * nullable bool with default
  * @returns {boolean}
  **/
getNullableboolFieldWithValue () { return this[`nullableboolFieldWithValue`] }
		/**
  * nullable bool with default
  * @param {boolean}
  **/
/**
  * nullable bool with default
  * @param {boolean}
  **/
/// XXX
setNullableboolFieldWithValue(value: boolean) {
	// Only accept array types
	this["nullableboolFieldWithValue"] = value
	return this
}
		/**
  * int field, non-nullable
  * @type {number}
  **/
 intField: number = 0
		/**
  * int field, non-nullable
  * @returns {number}
  **/
getIntField () { return this[`intField`] }
		/**
  * int field, non-nullable
  * @param {number}
  **/
/**
  * int field, non-nullable
  * @param {number}
  **/
/// XXX
setIntField(value: number) {
	// Only accept array types
	this["intField"] = value
	return this
}
		/**
  * int field with default
  * @type {number}
  **/
 intFieldWithValue: number = 42
		/**
  * int field with default
  * @returns {number}
  **/
getIntFieldWithValue () { return this[`intFieldWithValue`] }
		/**
  * int field with default
  * @param {number}
  **/
/**
  * int field with default
  * @param {number}
  **/
/// XXX
setIntFieldWithValue(value: number) {
	// Only accept array types
	this["intFieldWithValue"] = value
	return this
}
		/**
  * nullable int
  * @type {number}
  **/
 nullableIntField?: number | null = undefined
		/**
  * nullable int
  * @returns {number}
  **/
getNullableIntField () { return this[`nullableIntField`] }
		/**
  * nullable int
  * @param {number}
  **/
/**
  * nullable int
  * @param {number}
  **/
/// XXX
setNullableIntField(value: number) {
	// Only accept array types
	this["nullableIntField"] = value
	return this
}
		/**
  * nullable int with default
  * @type {number}
  **/
 nullableIntFieldWithValue?: number | null  = 7
		/**
  * nullable int with default
  * @returns {number}
  **/
getNullableIntFieldWithValue () { return this[`nullableIntFieldWithValue`] }
		/**
  * nullable int with default
  * @param {number}
  **/
/**
  * nullable int with default
  * @param {number}
  **/
/// XXX
setNullableIntFieldWithValue(value: number) {
	// Only accept array types
	this["nullableIntFieldWithValue"] = value
	return this
}
		/**
  * int32 field, non-nullable
  * @type {number}
  **/
 int32Field: number = 0
		/**
  * int32 field, non-nullable
  * @returns {number}
  **/
getInt32Field () { return this[`int32Field`] }
		/**
  * int32 field, non-nullable
  * @param {number}
  **/
/**
  * int32 field, non-nullable
  * @param {number}
  **/
/// XXX
setInt32Field(value: number) {
	// Only accept array types
	this["int32Field"] = value
	return this
}
		/**
  * int32 with default
  * @type {number}
  **/
 int32FieldWithValue: number = 100
		/**
  * int32 with default
  * @returns {number}
  **/
getInt32FieldWithValue () { return this[`int32FieldWithValue`] }
		/**
  * int32 with default
  * @param {number}
  **/
/**
  * int32 with default
  * @param {number}
  **/
/// XXX
setInt32FieldWithValue(value: number) {
	// Only accept array types
	this["int32FieldWithValue"] = value
	return this
}
		/**
  * nullable int32
  * @type {number}
  **/
 nullableInt32Field?: number | null = undefined
		/**
  * nullable int32
  * @returns {number}
  **/
getNullableInt32Field () { return this[`nullableInt32Field`] }
		/**
  * nullable int32
  * @param {number}
  **/
/**
  * nullable int32
  * @param {number}
  **/
/// XXX
setNullableInt32Field(value: number) {
	// Only accept array types
	this["nullableInt32Field"] = value
	return this
}
		/**
  * nullable int32 with default
  * @type {number}
  **/
 nullableInt32FieldWithValue?: number | null  = 200
		/**
  * nullable int32 with default
  * @returns {number}
  **/
getNullableInt32FieldWithValue () { return this[`nullableInt32FieldWithValue`] }
		/**
  * nullable int32 with default
  * @param {number}
  **/
/**
  * nullable int32 with default
  * @param {number}
  **/
/// XXX
setNullableInt32FieldWithValue(value: number) {
	// Only accept array types
	this["nullableInt32FieldWithValue"] = value
	return this
}
		/**
  * int64 field
  * @type {number}
  **/
 int64Field: number = 0
		/**
  * int64 field
  * @returns {number}
  **/
getInt64Field () { return this[`int64Field`] }
		/**
  * int64 field
  * @param {number}
  **/
/**
  * int64 field
  * @param {number}
  **/
/// XXX
setInt64Field(value: number) {
	// Only accept array types
	this["int64Field"] = value
	return this
}
		/**
  * int64 with default
  * @type {number}
  **/
 int64FieldWithValue?: number | null  = 123
		/**
  * int64 with default
  * @returns {number}
  **/
getInt64FieldWithValue () { return this[`int64FieldWithValue`] }
		/**
  * int64 with default
  * @param {number}
  **/
/**
  * int64 with default
  * @param {number}
  **/
/// XXX
setInt64FieldWithValue(value: number) {
	// Only accept array types
	this["int64FieldWithValue"] = value
	return this
}
		/**
  * nullable int64
  * @type {number}
  **/
 nullableInt64Field?: number | null = undefined
		/**
  * nullable int64
  * @returns {number}
  **/
getNullableInt64Field () { return this[`nullableInt64Field`] }
		/**
  * nullable int64
  * @param {number}
  **/
/**
  * nullable int64
  * @param {number}
  **/
/// XXX
setNullableInt64Field(value: number) {
	// Only accept array types
	this["nullableInt64Field"] = value
	return this
}
		/**
  * nullable int64 with default
  * @type {number}
  **/
 nullableInt64FieldWithValue?: number | null  = 456
		/**
  * nullable int64 with default
  * @returns {number}
  **/
getNullableInt64FieldWithValue () { return this[`nullableInt64FieldWithValue`] }
		/**
  * nullable int64 with default
  * @param {number}
  **/
/**
  * nullable int64 with default
  * @param {number}
  **/
/// XXX
setNullableInt64FieldWithValue(value: number) {
	// Only accept array types
	this["nullableInt64FieldWithValue"] = value
	return this
}
		/**
  * float32 field
  * @type {number}
  **/
 float32Field: number = 0.0
		/**
  * float32 field
  * @returns {number}
  **/
getFloat32Field () { return this[`float32Field`] }
		/**
  * float32 field
  * @param {number}
  **/
/**
  * float32 field
  * @param {number}
  **/
/// XXX
setFloat32Field(value: number) {
	// Only accept array types
	this["float32Field"] = value
	return this
}
		/**
  * float32 with default
  * @type {number}
  **/
 float32FieldWithValue: number = 1.23
		/**
  * float32 with default
  * @returns {number}
  **/
getFloat32FieldWithValue () { return this[`float32FieldWithValue`] }
		/**
  * float32 with default
  * @param {number}
  **/
/**
  * float32 with default
  * @param {number}
  **/
/// XXX
setFloat32FieldWithValue(value: number) {
	// Only accept array types
	this["float32FieldWithValue"] = value
	return this
}
		/**
  * nullable float32
  * @type {number}
  **/
 nullableFloat32Field?: number | null = undefined
		/**
  * nullable float32
  * @returns {number}
  **/
getNullableFloat32Field () { return this[`nullableFloat32Field`] }
		/**
  * nullable float32
  * @param {number}
  **/
/**
  * nullable float32
  * @param {number}
  **/
/// XXX
setNullableFloat32Field(value: number) {
	// Only accept array types
	this["nullableFloat32Field"] = value
	return this
}
		/**
  * nullable float32 with default
  * @type {number}
  **/
 nullableFloat32FieldWithValue?: number | null  = 4.56
		/**
  * nullable float32 with default
  * @returns {number}
  **/
getNullableFloat32FieldWithValue () { return this[`nullableFloat32FieldWithValue`] }
		/**
  * nullable float32 with default
  * @param {number}
  **/
/**
  * nullable float32 with default
  * @param {number}
  **/
/// XXX
setNullableFloat32FieldWithValue(value: number) {
	// Only accept array types
	this["nullableFloat32FieldWithValue"] = value
	return this
}
		/**
  * float64 field
  * @type {number}
  **/
 float64Field: number = 0.0
		/**
  * float64 field
  * @returns {number}
  **/
getFloat64Field () { return this[`float64Field`] }
		/**
  * float64 field
  * @param {number}
  **/
/**
  * float64 field
  * @param {number}
  **/
/// XXX
setFloat64Field(value: number) {
	// Only accept array types
	this["float64Field"] = value
	return this
}
		/**
  * float64 with default
  * @type {number}
  **/
 float64FieldWithValue: number = 7.89
		/**
  * float64 with default
  * @returns {number}
  **/
getFloat64FieldWithValue () { return this[`float64FieldWithValue`] }
		/**
  * float64 with default
  * @param {number}
  **/
/**
  * float64 with default
  * @param {number}
  **/
/// XXX
setFloat64FieldWithValue(value: number) {
	// Only accept array types
	this["float64FieldWithValue"] = value
	return this
}
		/**
  * nullable float64
  * @type {number}
  **/
 nullableFloat64Field?: number | null = undefined
		/**
  * nullable float64
  * @returns {number}
  **/
getNullableFloat64Field () { return this[`nullableFloat64Field`] }
		/**
  * nullable float64
  * @param {number}
  **/
/**
  * nullable float64
  * @param {number}
  **/
/// XXX
setNullableFloat64Field(value: number) {
	// Only accept array types
	this["nullableFloat64Field"] = value
	return this
}
		/**
  * nullable float64 with default
  * @type {number}
  **/
 nullableFloat64FieldWithValue?: number | null  = 0.12
		/**
  * nullable float64 with default
  * @returns {number}
  **/
getNullableFloat64FieldWithValue () { return this[`nullableFloat64FieldWithValue`] }
		/**
  * nullable float64 with default
  * @param {number}
  **/
/**
  * nullable float64 with default
  * @param {number}
  **/
/// XXX
setNullableFloat64FieldWithValue(value: number) {
	// Only accept array types
	this["nullableFloat64FieldWithValue"] = value
	return this
}
}
export abstract class AnonymouseFactory {
	abstract create(data: unknown): Anonymouse;
}