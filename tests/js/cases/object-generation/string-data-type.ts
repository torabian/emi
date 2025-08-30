	/**
  * @description The base type definition for anonymouse
  **/
	export type AnonymouseType =  {
			/**
  * @type {string}
  * @description This is a pure string field, there for never can be null, and by default needs to be empty string
  **/
 stringField?: string;
			/**
  * @type {string}
  * @description Pure string field, but with an intial string value, and never can be undefined or null
  **/
 stringFieldWithValue?: string;
			/**
  * @type {string}
  * @description Nullable string field. Can be undefined, or set to null to indicate intentional emptiness
  **/
 nullableStringField?: string;
			/**
  * @type {string}
  * @description Nullable string field, can be undefined or null, but with an initial value
  **/
 nullableStringFieldWithValue?: string;
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
  * @returns {string}
  * @description This is a pure string field, there for never can be null, and by default needs to be empty string
  **/
getStringField () { return this[`stringField`] }
		/**
  * This is a pure string field, there for never can be null, and by default needs to be empty string
  * @param {string}
  **/
setStringField (value: string ) { this[`stringField`] = value; return this; } 
		/**
  * Pure string field, but with an intial string value, and never can be undefined or null
  * @type {string}
  **/
 stringFieldWithValue: string = "testvalue"
		/**
  * @returns {string}
  * @description Pure string field, but with an intial string value, and never can be undefined or null
  **/
getStringFieldWithValue () { return this[`stringFieldWithValue`] }
		/**
  * Pure string field, but with an intial string value, and never can be undefined or null
  * @param {string}
  **/
setStringFieldWithValue (value: string ) { this[`stringFieldWithValue`] = value; return this; } 
		/**
  * Nullable string field. Can be undefined, or set to null to indicate intentional emptiness
  * @type {string}
  **/
 nullableStringField?: string | null = undefined
		/**
  * @returns {string}
  * @description Nullable string field. Can be undefined, or set to null to indicate intentional emptiness
  **/
getNullableStringField () { return this[`nullableStringField`] }
		/**
  * Nullable string field. Can be undefined, or set to null to indicate intentional emptiness
  * @param {string}
  **/
setNullableStringField (value: string  | null) { this[`nullableStringField`] = value; return this; } 
		/**
  * Nullable string field, can be undefined or null, but with an initial value
  * @type {string}
  **/
 nullableStringFieldWithValue?: string | null  = "stringvalue"
		/**
  * @returns {string}
  * @description Nullable string field, can be undefined or null, but with an initial value
  **/
getNullableStringFieldWithValue () { return this[`nullableStringFieldWithValue`] }
		/**
  * Nullable string field, can be undefined or null, but with an initial value
  * @param {string}
  **/
setNullableStringFieldWithValue (value: string  | null) { this[`nullableStringFieldWithValue`] = value; return this; } 
}
export abstract class AnonymouseFactory {
	abstract create(data: unknown): Anonymouse;
}