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
	*	Json stringify won't see them unless we mention it explicitly. Each
	*	field goes through #toPlainJSON rather than a bare this.#field: a
	*	nested dto instance, or an emigo Array/One/Collection wrapper, has its
	*	own toJSON that JSON.stringify would only reach on a *second* pass -
	*	calling it here means toJSON()'s own return value is already the same
	*	plain shape a consumer choosing the exported type instead of this
	*	class gets, rather than a shallow object still holding class instances.
	**/
	toJSON(): UserType {
    	return {
		} as UserType;
  	}
	/**
	* Resolves value into a plain, JSON-serializable value: recurses into
	* arrays, and - the case JSON.stringify's own recursion would only get to
	* after this method already returned - calls a nested value's own
	* toJSON() (a dto instance, or an emigo Array/One/Collection wrapper)
	* rather than leaving it as-is. A plain value (string, number, a map's
	* own plain object, ...) is returned unchanged.
	**/
	#toPlainJSON(value: unknown): unknown {
		if (value === null || value === undefined) {
			return value;
		}
		if (Array.isArray(value)) {
			return value.map((item) => this.#toPlainJSON(item));
		}
		const asAny = value as any;
		if (typeof asAny.toJSON === "function") {
			return this.#toPlainJSON(asAny.toJSON());
		}
		return value;
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
  