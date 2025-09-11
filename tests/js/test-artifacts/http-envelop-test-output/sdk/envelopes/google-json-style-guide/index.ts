import { CreatorSignature, EnvelopeClass } from "../common/EnvelopeClass";
import { ResponseDto } from "./generated/ResponseDto";

// Use this class to generate a GResponse.
export class GResponse<T> extends ResponseDto implements EnvelopeClass<T> {
  creator: CreatorSignature<T>;

  constructor(data?: unknown) {
    super(data);
  }

  setCreator(fn: CreatorSignature<T>) {
    this.creator = fn;

    return this;
  }

  /**
   * GResponse can have data.item or data.items
   * We create that based on incoming data tpye, so there is no need for 2 different
   * classes, one for array and other for singular
   * @param data
   * @returns
   */
  inject(body: any): this {
    this.applyFromObject(body);

    if ((body as any)?.data) {
      if (!this.data) {
        this.setData({} as any);
      }
      if (Array.isArray(body?.data.items)) {
        this.data?.setItems(
          body?.data?.items?.map((item) => this.creator(item)),
        );
      } else if (typeof body?.data?.item === "object") {
        this.data?.setItem(this.creator(body?.data?.item));
      }
    }

    return this;
  }
}
