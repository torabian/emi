import { ResponseDto } from "./generated/ResponseDto";

/**
 * To declare a new envelope class, you must follow these rules:
 * 1. The constructor must accept any object, which will contain the parsed JSON message
 *    from a response.
 * 2. Envelope classes must provide a function to update the payload. Since payloads are
 *    type-safe, they must be instantiated and passed to the envelope; the common constructor
 *    alone is not enough.
 * 3. Enveope must have a way to provide the content back actually, in order to create a class out of them.
 */

interface EnvelopeClass<T> {
  updatePayload(payload: T): void;

  getPayload(): T;
}

// Use this class to generate a GResponse.
export class GResponse<T> extends ResponseDto implements EnvelopeClass<T> {
  constructor(data: unknown) {
    super(data);
  }

  updatePayload(payload: T) {
    if (!this.data) {
      this.setData({} as any);
    }

    this.data?.setItem(payload as any);

    return this;
  }

  getPayload(): T {
    return this.data?.item as T;
  }
}
