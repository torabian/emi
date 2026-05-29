export class MOne<T, S = unknown> {
  private operation: string | null = null;
  private selector: S | undefined;

  private content: T | undefined = undefined;

  isNull(): boolean {
    return this.content === null;
  }

  isSelector(): boolean {
    return this.selector !== undefined && this.operation !== null;
  }

  get(): T {
    return this.content!;
  }

  static of<T>(value: T) {
    const one = new MOne<T>();
    one.content = value;

    return one;
  }

  static select<T = unknown>(selector: any) {
    const one = new MOne<T>();
    one.selector = selector;
    one.operation = "replace";

    return one;
  }

  toJSON() {
    // When its explicit replace, it means that we need to pass a selector, so reader will
    // be able to replace it via internal mechanism
    if (this.operation === "replace") {
      return {
        __operation: this.operation,
        selector: this.selector,
      };
    }

    return this.content;
  }
}

// In javascript, nullability and undefined is already working perfectly fine.
// hence, maybe just giving back one is enough.
export class MOneNullable<T, S = unknown> extends MOne<T, S> {}

// Array describes how an incoming list should be applied to an existing one,
// mirroring emigo.Array on the Go side. It lets a PATCH-style payload say
// "replace the whole set" versus "append to the existing set".
//
// On the wire a "replace" is implicit (a bare array), while an "append" is
// tagged with __operation so the reader keeps the existing rows. This keeps
// the payload identical to the Go client/backend generators.
export class MArray<T> {
  private operation: "replace" | "append" = "replace";
  private items: T[] = [];

  isAppend(): boolean {
    return this.operation === "append";
  }

  isReplace(): boolean {
    return this.operation === "replace";
  }

  len(): number {
    return this.items.length;
  }

  get(): T[] {
    return this.items;
  }

  // Full replacement — existing rows are cleared before these are applied.
  static of<T>(items: T[]) {
    const arr = new MArray<T>();
    arr.items = items;
    arr.operation = "replace";

    return arr;
  }

  // Append — existing rows are preserved and these are added alongside them.
  static append<T>(items: T[]) {
    const arr = new MArray<T>();
    arr.items = items;
    arr.operation = "append";

    return arr;
  }

  toJSON() {
    // "replace" is implicit on the wire, so we emit a bare array. Only the
    // "append" operation needs the explicit tagged-object form.
    if (this.operation === "append") {
      return {
        __operation: this.operation,
        items: this.items,
      };
    }

    return this.items;
  }
}

// In javascript, nullability and undefined is already working perfectly fine.
// hence, just extending Array is enough.
export class MArrayNullable<T> extends Array<T> {}

// Collection mirrors emigo.Collection on the Go side. Structurally it is the
// same as Array — a list carrying a "replace"/"append" operation — but it is a
// distinct field type: a collection holds a list of a target entity, whereas an
// array holds a list of an inline DTO.
export class MCollection<T> {
  private operation: "replace" | "append" = "replace";
  private items: T[] = [];

  isAppend(): boolean {
    return this.operation === "append";
  }

  isReplace(): boolean {
    return this.operation === "replace";
  }

  len(): number {
    return this.items.length;
  }

  get(): T[] {
    return this.items;
  }

  // Full replacement — existing rows are cleared before these are applied.
  static of<T>(items: T[]) {
    const collection = new MCollection<T>();
    collection.items = items;
    collection.operation = "replace";

    return collection;
  }

  // Append — existing rows are preserved and these are added alongside them.
  static append<T>(items: T[]) {
    const collection = new MCollection<T>();
    collection.items = items;
    collection.operation = "append";

    return collection;
  }

  toJSON() {
    // "replace" is implicit on the wire, so we emit a bare array. Only the
    // "append" operation needs the explicit tagged-object form.
    if (this.operation === "append") {
      return {
        __operation: this.operation,
        items: this.items,
      };
    }

    return this.items;
  }
}

// In javascript, nullability and undefined is already working perfectly fine.
// hence, just extending Collection is enough.
export class MCollectionNullable<T> extends MCollection<T> {}
