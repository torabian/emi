import { type PartialDeep } from './sdk/common/fetchx';
/**
  * The base class definition for user
  **/
export class User {
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
		const g = globalThis as unknown as { Buffer: any; Blob: any };
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
		const d = data as Partial<User>;
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
	  }
	}
	/**
	* Creates an instance of User, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: UserType) {
		return new User(possibleDtoObject);
	}
	/**
	* Creates an instance of User, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<UserType>) {
		return new User(partialDtoObject);
	}
	copyWith(partial: PartialDeep<UserType>): InstanceType<typeof User> {
		return new User ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof User> {
		return new User(this.toJSON());
	}
}
export abstract class UserFactory {
	abstract create(data: unknown): User;
}
	/**
  * The base type definition for user
  **/
	export type UserType =  {
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UserType {
}
  
  function isPlausibleObject(v: any) { return false }
  