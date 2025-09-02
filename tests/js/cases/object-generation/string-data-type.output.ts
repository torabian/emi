	/**
  * The base type definition for anonymouse
  **/
	export type AnonymouseType =  {
			/**
  * This is a pure string field, there for never can be null, and by default needs to be empty string
  * @type {string}
  **/
 stringField?: string;
			/**
  * Pure string field, but with an intial string value, and never can be undefined or null
  * @type {string}
  **/
 stringFieldWithValue?: string;
			/**
  * Nullable string field. Can be undefined, or set to null to indicate intentional emptiness
  * @type {string}
  **/
 nullableStringField?: string;
			/**
  * Nullable string field, can be undefined or null, but with an initial value
  * @type {string}
  **/
 nullableStringFieldWithValue?: string;
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
			if (d[`stringField`] !== undefined) { 
 this.setStringField (d[`stringField`]) 
}
			if (d[`stringFieldWithValue`] !== undefined) { 
 this.setStringFieldWithValue (d[`stringFieldWithValue`]) 
}
			if (d[`nullableStringField`] !== undefined) { 
 this.setNullableStringField (d[`nullableStringField`]) 
}
			if (d[`nullableStringFieldWithValue`] !== undefined) { 
 this.setNullableStringFieldWithValue (d[`nullableStringFieldWithValue`]) 
}
	}
		/**
  * This is a pure string field, there for never can be null, and by default needs to be empty string
  * @type {string}
  **/
 stringField: string = ""
		/**
  * This is a pure string field, there for never can be null, and by default needs to be empty string
  * @returns {string}
  **/
getStringField () { return this[`stringField`] }
		/**
  * This is a pure string field, there for never can be null, and by default needs to be empty string
  * @param {string}
  **/
/**
  * This is a pure string field, there for never can be null, and by default needs to be empty string
  * @param {string}
  **/
/// XXX
setStringField(value: string) {
	// Only accept array types
	this["stringField"] = value
	return this
}
		/**
  * Pure string field, but with an intial string value, and never can be undefined or null
  * @type {string}
  **/
 stringFieldWithValue: string = "testvalue"
		/**
  * Pure string field, but with an intial string value, and never can be undefined or null
  * @returns {string}
  **/
getStringFieldWithValue () { return this[`stringFieldWithValue`] }
		/**
  * Pure string field, but with an intial string value, and never can be undefined or null
  * @param {string}
  **/
/**
  * Pure string field, but with an intial string value, and never can be undefined or null
  * @param {string}
  **/
/// XXX
setStringFieldWithValue(value: string) {
	// Only accept array types
	this["stringFieldWithValue"] = value
	return this
}
		/**
  * Nullable string field. Can be undefined, or set to null to indicate intentional emptiness
  * @type {string}
  **/
 nullableStringField?: string | null = undefined
		/**
  * Nullable string field. Can be undefined, or set to null to indicate intentional emptiness
  * @returns {string}
  **/
getNullableStringField () { return this[`nullableStringField`] }
		/**
  * Nullable string field. Can be undefined, or set to null to indicate intentional emptiness
  * @param {string}
  **/
/**
  * Nullable string field. Can be undefined, or set to null to indicate intentional emptiness
  * @param {string}
  **/
/// XXX
setNullableStringField(value: string) {
	// Only accept array types
	this["nullableStringField"] = value
	return this
}
		/**
  * Nullable string field, can be undefined or null, but with an initial value
  * @type {string}
  **/
 nullableStringFieldWithValue?: string | null  = "stringvalue"
		/**
  * Nullable string field, can be undefined or null, but with an initial value
  * @returns {string}
  **/
getNullableStringFieldWithValue () { return this[`nullableStringFieldWithValue`] }
		/**
  * Nullable string field, can be undefined or null, but with an initial value
  * @param {string}
  **/
/**
  * Nullable string field, can be undefined or null, but with an initial value
  * @param {string}
  **/
/// XXX
setNullableStringFieldWithValue(value: string) {
	// Only accept array types
	this["nullableStringFieldWithValue"] = value
	return this
}
}
export abstract class AnonymouseFactory {
	abstract create(data: unknown): Anonymouse;
}