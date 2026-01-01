import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Create offer based on product
*/
export type CreateOfferBasedOnProductActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * CreateOfferBasedOnProductAction
 */
export class CreateOfferBasedOnProductAction { //
  static URL = 'https://api.{environment}/sale/product-offers';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		CreateOfferBasedOnProductAction.URL,
		 undefined,
		qs
	);
  static Method = 'post';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<CreateOfferBasedOnProductActionReq, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<unknown, CreateOfferBasedOnProductActionReq, unknown>(
			overrideUrl ?? CreateOfferBasedOnProductAction.NewUrl(
				qs
			),
			{
				method: CreateOfferBasedOnProductAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init?: TypedRequestInit<CreateOfferBasedOnProductActionReq, unknown>,
		{
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
		}
	) => {
		const res = await CreateOfferBasedOnProductAction.Fetch$(
			qs,
			ctx,
			init,
			overrideUrl,
			);
			return handleFetchResponse(
				res, 
				undefined,
				onMessage,
				init?.signal,
			);
	}
  static Definition = {
  "name": "Create offer based on product",
  "url": "https://api.{environment}/sale/product-offers",
  "method": "post",
  "in": {
    "fields": [
      {
        "name": "name",
        "description": "Offer title",
        "type": "string"
      },
      {
        "name": "language",
        "description": "Offer language code (e.g., pl-PL)",
        "type": "string"
      },
      {
        "name": "category",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
          }
        ]
      },
      {
        "name": "productSet",
        "description": "Product details and associated quantities",
        "type": "array",
        "fields": [
          {
            "name": "product",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "idType",
                "type": "string"
              },
              {
                "name": "name",
                "type": "string"
              },
              {
                "name": "category",
                "type": "object",
                "fields": [
                  {
                    "name": "id",
                    "type": "string"
                  }
                ]
              },
              {
                "name": "parameters",
                "type": "array",
                "fields": [
                  {
                    "name": "id",
                    "type": "string"
                  },
                  {
                    "name": "name",
                    "type": "string"
                  },
                  {
                    "name": "rangeValue",
                    "type": "object",
                    "fields": [
                      {
                        "name": "from",
                        "type": "string"
                      },
                      {
                        "name": "to",
                        "type": "string"
                      }
                    ]
                  },
                  {
                    "name": "values",
                    "type": "slice",
                    "primitive": "string"
                  },
                  {
                    "name": "valuesIds",
                    "type": "slice",
                    "primitive": "string"
                  }
                ]
              },
              {
                "name": "images",
                "type": "slice",
                "primitive": "string"
              }
            ]
          },
          {
            "name": "quantity",
            "type": "object",
            "fields": [
              {
                "name": "value",
                "type": "int"
              }
            ]
          },
          {
            "name": "responsiblePerson",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "name",
                "type": "string"
              }
            ]
          },
          {
            "name": "responsibleProducer",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "type",
                "type": "string"
              }
            ]
          },
          {
            "name": "safetyInformation",
            "type": "object",
            "fields": [
              {
                "name": "type",
                "type": "string"
              },
              {
                "name": "description",
                "type": "string"
              }
            ]
          },
          {
            "name": "marketedBeforeGPSRObligation",
            "type": "bool"
          },
          {
            "name": "deposits",
            "type": "array",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "quantity",
                "type": "int"
              }
            ]
          }
        ]
      },
      {
        "name": "stock",
        "type": "object",
        "fields": [
          {
            "name": "available",
            "type": "int"
          },
          {
            "name": "unit",
            "type": "string"
          }
        ]
      },
      {
        "name": "sellingMode",
        "type": "object",
        "fields": [
          {
            "name": "format",
            "type": "string"
          },
          {
            "name": "price",
            "type": "object",
            "fields": [
              {
                "name": "amount",
                "type": "string"
              },
              {
                "name": "currency",
                "type": "string"
              }
            ]
          },
          {
            "name": "minimalPrice",
            "type": "object",
            "fields": [
              {
                "name": "amount",
                "type": "string"
              },
              {
                "name": "currency",
                "type": "string"
              }
            ]
          },
          {
            "name": "startingPrice",
            "type": "object",
            "fields": [
              {
                "name": "amount",
                "type": "string"
              },
              {
                "name": "currency",
                "type": "string"
              }
            ]
          }
        ]
      },
      {
        "name": "payments",
        "type": "object",
        "fields": [
          {
            "name": "invoice",
            "type": "string"
          }
        ]
      },
      {
        "name": "delivery",
        "type": "object",
        "fields": [
          {
            "name": "handlingTime",
            "type": "string"
          },
          {
            "name": "additionalInfo",
            "type": "string"
          },
          {
            "name": "shipmentDate",
            "type": "string"
          },
          {
            "name": "shippingRates",
            "description": "Optional; may be null",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              }
            ]
          }
        ]
      },
      {
        "name": "publication",
        "type": "object",
        "fields": [
          {
            "name": "duration",
            "type": "string"
          },
          {
            "name": "startingAt",
            "type": "string"
          },
          {
            "name": "endingAt",
            "type": "string"
          },
          {
            "name": "status",
            "type": "string"
          },
          {
            "name": "republish",
            "type": "bool"
          }
        ]
      },
      {
        "name": "additionalMarketplaces",
        "type": "map?",
        "mapKey": "string",
        "fields": [
          {
            "name": "sellingMode",
            "type": "object",
            "fields": [
              {
                "name": "price",
                "type": "object",
                "fields": [
                  {
                    "name": "amount",
                    "type": "string"
                  },
                  {
                    "name": "currency",
                    "type": "string"
                  }
                ]
              }
            ]
          }
        ]
      },
      {
        "name": "compatibilityList",
        "type": "object",
        "fields": [
          {
            "name": "items",
            "type": "array",
            "fields": [
              {
                "name": "type",
                "type": "string"
              },
              {
                "name": "text",
                "type": "string"
              }
            ]
          }
        ]
      },
      {
        "name": "images",
        "type": "slice",
        "primitive": "string"
      },
      {
        "name": "description",
        "type": "object",
        "fields": [
          {
            "name": "sections",
            "type": "array",
            "fields": [
              {
                "name": "items",
                "type": "array",
                "fields": [
                  {
                    "name": "type",
                    "type": "string"
                  }
                ]
              }
            ]
          }
        ]
      },
      {
        "name": "b2b",
        "type": "object",
        "fields": [
          {
            "name": "buyableOnlyByBusiness",
            "type": "bool"
          }
        ]
      },
      {
        "name": "attachments",
        "type": "array",
        "fields": [
          {
            "name": "id",
            "type": "string"
          }
        ]
      },
      {
        "name": "fundraisingCampaign",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "name",
            "type": "string"
          }
        ]
      },
      {
        "name": "additionalServices",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "name",
            "type": "string"
          }
        ]
      },
      {
        "name": "afterSalesServices",
        "type": "object",
        "fields": [
          {
            "name": "impliedWarranty",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "name",
                "type": "string"
              }
            ]
          },
          {
            "name": "returnPolicy",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "name",
                "type": "string"
              }
            ]
          },
          {
            "name": "warranty",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "name",
                "type": "string"
              }
            ]
          }
        ]
      },
      {
        "name": "sizeTable",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "name",
            "type": "string"
          }
        ]
      },
      {
        "name": "contact",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "name",
            "type": "string"
          }
        ]
      },
      {
        "name": "discounts",
        "type": "object",
        "fields": [
          {
            "name": "wholesalePriceList",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "name",
                "type": "string"
              }
            ]
          }
        ]
      },
      {
        "name": "location",
        "type": "object",
        "fields": [
          {
            "name": "city",
            "type": "string"
          },
          {
            "name": "countryCode",
            "type": "string"
          },
          {
            "name": "postCode",
            "type": "string"
          },
          {
            "name": "province",
            "type": "string"
          }
        ]
      },
      {
        "name": "external",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
          }
        ]
      },
      {
        "name": "taxSettings",
        "type": "object",
        "fields": [
          {
            "name": "subject",
            "type": "string"
          },
          {
            "name": "exemption",
            "type": "string"
          },
          {
            "name": "rates",
            "type": "array",
            "fields": [
              {
                "name": "rate",
                "type": "string"
              },
              {
                "name": "countryCode",
                "type": "string"
              }
            ]
          }
        ]
      },
      {
        "name": "messageToSellerSettings",
        "type": "object",
        "fields": [
          {
            "name": "mode",
            "type": "string"
          },
          {
            "name": "hint",
            "type": "string"
          }
        ]
      }
    ]
  }
}
}
/**
  * The base class definition for createOfferBasedOnProductActionReq
  **/
export class CreateOfferBasedOnProductActionReq {
		/**
  * Offer title
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * Offer title
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * Offer title
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
		/**
  * Offer language code (e.g., pl-PL)
  * @type {string}
  **/
 #language : string  =  ""
		/**
  * Offer language code (e.g., pl-PL)
  * @returns {string}
  **/
get language () { return this.#language }
/**
  * Offer language code (e.g., pl-PL)
  * @type {string}
  **/
set language (value: string) {
		this.#language = String(value);
}
setLanguage (value: string) {
	this.language = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Category}
  **/
 #category ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Category>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Category}
  **/
set category (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Category>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Category) {
			this.#category = value
		} else {
			this.#category = new CreateOfferBasedOnProductActionReq.Category(value)
		}
}
setCategory (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Category>) {
	this.category = value
	return this
}
		/**
  * Product details and associated quantities
  * @type {CreateOfferBasedOnProductActionReq.ProductSet}
  **/
 #productSet : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet>[]  =  []
		/**
  * Product details and associated quantities
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet}
  **/
get productSet () { return this.#productSet }
/**
  * Product details and associated quantities
  * @type {CreateOfferBasedOnProductActionReq.ProductSet}
  **/
set productSet (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CreateOfferBasedOnProductActionReq.ProductSet) {
			this.#productSet = value
		} else {
			this.#productSet = value.map(item => new CreateOfferBasedOnProductActionReq.ProductSet(item))
		}
}
setProductSet (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet>[]) {
	this.productSet = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Stock}
  **/
 #stock ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Stock>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Stock}
  **/
set stock (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Stock>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Stock) {
			this.#stock = value
		} else {
			this.#stock = new CreateOfferBasedOnProductActionReq.Stock(value)
		}
}
setStock (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Stock>) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SellingMode}
  **/
 #sellingMode ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SellingMode}
  **/
set sellingMode (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new CreateOfferBasedOnProductActionReq.SellingMode(value)
		}
}
setSellingMode (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode>) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Payments}
  **/
 #payments ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Payments>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Payments}
  **/
get payments () { return this.#payments }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Payments}
  **/
set payments (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Payments>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Payments) {
			this.#payments = value
		} else {
			this.#payments = new CreateOfferBasedOnProductActionReq.Payments(value)
		}
}
setPayments (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Payments>) {
	this.payments = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Delivery}
  **/
 #delivery ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Delivery}
  **/
get delivery () { return this.#delivery }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Delivery}
  **/
set delivery (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Delivery) {
			this.#delivery = value
		} else {
			this.#delivery = new CreateOfferBasedOnProductActionReq.Delivery(value)
		}
}
setDelivery (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery>) {
	this.delivery = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Publication}
  **/
 #publication ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Publication>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Publication}
  **/
set publication (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Publication>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Publication) {
			this.#publication = value
		} else {
			this.#publication = new CreateOfferBasedOnProductActionReq.Publication(value)
		}
}
setPublication (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Publication>) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {{[key: string]: any}}
  **/
 #additionalMarketplaces ? : {[key: string]: any}  | null  =  undefined
		/**
  * 
  * @returns {{[key: string]: any}}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {{[key: string]: any}}
  **/
set additionalMarketplaces (value: {[key: string]: any} | null | undefined) {
}
setAdditionalMarketplaces (value: {[key: string]: any} | null | undefined) {
	this.additionalMarketplaces = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.CompatibilityList}
  **/
 #compatibilityList ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.CompatibilityList}
  **/
get compatibilityList () { return this.#compatibilityList }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.CompatibilityList}
  **/
set compatibilityList (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.CompatibilityList) {
			this.#compatibilityList = value
		} else {
			this.#compatibilityList = new CreateOfferBasedOnProductActionReq.CompatibilityList(value)
		}
}
setCompatibilityList (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList>) {
	this.compatibilityList = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #images : string[]  =  []
		/**
  * 
  * @returns {string[]}
  **/
get images () { return this.#images }
/**
  * 
  * @type {string[]}
  **/
set images (value: string[]) {
}
setImages (value: string[]) {
	this.images = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Description}
  **/
 #description ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Description>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Description}
  **/
set description (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Description>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Description) {
			this.#description = value
		} else {
			this.#description = new CreateOfferBasedOnProductActionReq.Description(value)
		}
}
setDescription (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Description>) {
	this.description = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.B2b}
  **/
 #b2b ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.B2b>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.B2b}
  **/
get b2b () { return this.#b2b }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.B2b}
  **/
set b2b (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.B2b>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.B2b) {
			this.#b2b = value
		} else {
			this.#b2b = new CreateOfferBasedOnProductActionReq.B2b(value)
		}
}
setB2b (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.B2b>) {
	this.b2b = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Attachments}
  **/
 #attachments : InstanceType<typeof CreateOfferBasedOnProductActionReq.Attachments>[]  =  []
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Attachments}
  **/
get attachments () { return this.#attachments }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Attachments}
  **/
set attachments (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Attachments>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CreateOfferBasedOnProductActionReq.Attachments) {
			this.#attachments = value
		} else {
			this.#attachments = value.map(item => new CreateOfferBasedOnProductActionReq.Attachments(item))
		}
}
setAttachments (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Attachments>[]) {
	this.attachments = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.FundraisingCampaign}
  **/
 #fundraisingCampaign ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.FundraisingCampaign>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.FundraisingCampaign}
  **/
get fundraisingCampaign () { return this.#fundraisingCampaign }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.FundraisingCampaign}
  **/
set fundraisingCampaign (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.FundraisingCampaign>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.FundraisingCampaign) {
			this.#fundraisingCampaign = value
		} else {
			this.#fundraisingCampaign = new CreateOfferBasedOnProductActionReq.FundraisingCampaign(value)
		}
}
setFundraisingCampaign (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.FundraisingCampaign>) {
	this.fundraisingCampaign = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AdditionalServices}
  **/
 #additionalServices ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.AdditionalServices>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.AdditionalServices}
  **/
get additionalServices () { return this.#additionalServices }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AdditionalServices}
  **/
set additionalServices (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AdditionalServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.AdditionalServices) {
			this.#additionalServices = value
		} else {
			this.#additionalServices = new CreateOfferBasedOnProductActionReq.AdditionalServices(value)
		}
}
setAdditionalServices (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AdditionalServices>) {
	this.additionalServices = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AfterSalesServices}
  **/
 #afterSalesServices ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.AfterSalesServices}
  **/
get afterSalesServices () { return this.#afterSalesServices }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AfterSalesServices}
  **/
set afterSalesServices (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.AfterSalesServices) {
			this.#afterSalesServices = value
		} else {
			this.#afterSalesServices = new CreateOfferBasedOnProductActionReq.AfterSalesServices(value)
		}
}
setAfterSalesServices (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices>) {
	this.afterSalesServices = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SizeTable}
  **/
 #sizeTable ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.SizeTable>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.SizeTable}
  **/
get sizeTable () { return this.#sizeTable }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SizeTable}
  **/
set sizeTable (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SizeTable>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.SizeTable) {
			this.#sizeTable = value
		} else {
			this.#sizeTable = new CreateOfferBasedOnProductActionReq.SizeTable(value)
		}
}
setSizeTable (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SizeTable>) {
	this.sizeTable = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Contact}
  **/
 #contact ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Contact>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Contact}
  **/
get contact () { return this.#contact }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Contact}
  **/
set contact (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Contact>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Contact) {
			this.#contact = value
		} else {
			this.#contact = new CreateOfferBasedOnProductActionReq.Contact(value)
		}
}
setContact (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Contact>) {
	this.contact = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Discounts}
  **/
 #discounts ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Discounts}
  **/
get discounts () { return this.#discounts }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Discounts}
  **/
set discounts (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Discounts) {
			this.#discounts = value
		} else {
			this.#discounts = new CreateOfferBasedOnProductActionReq.Discounts(value)
		}
}
setDiscounts (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts>) {
	this.discounts = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Location}
  **/
 #location ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Location>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Location}
  **/
get location () { return this.#location }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Location}
  **/
set location (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Location>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Location) {
			this.#location = value
		} else {
			this.#location = new CreateOfferBasedOnProductActionReq.Location(value)
		}
}
setLocation (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Location>) {
	this.location = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.External}
  **/
 #external ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.External>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.External}
  **/
get external () { return this.#external }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.External}
  **/
set external (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.External>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.External) {
			this.#external = value
		} else {
			this.#external = new CreateOfferBasedOnProductActionReq.External(value)
		}
}
setExternal (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.External>) {
	this.external = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.TaxSettings}
  **/
 #taxSettings ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.TaxSettings}
  **/
get taxSettings () { return this.#taxSettings }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.TaxSettings}
  **/
set taxSettings (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.TaxSettings) {
			this.#taxSettings = value
		} else {
			this.#taxSettings = new CreateOfferBasedOnProductActionReq.TaxSettings(value)
		}
}
setTaxSettings (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings>) {
	this.taxSettings = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.MessageToSellerSettings}
  **/
 #messageToSellerSettings ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.MessageToSellerSettings>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.MessageToSellerSettings}
  **/
get messageToSellerSettings () { return this.#messageToSellerSettings }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.MessageToSellerSettings}
  **/
set messageToSellerSettings (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.MessageToSellerSettings>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.MessageToSellerSettings) {
			this.#messageToSellerSettings = value
		} else {
			this.#messageToSellerSettings = new CreateOfferBasedOnProductActionReq.MessageToSellerSettings(value)
		}
}
setMessageToSellerSettings (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.MessageToSellerSettings>) {
	this.messageToSellerSettings = value
	return this
}
/**
  * The base class definition for category
  **/
static Category = class Category {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
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
		const d = data as Partial<Category>;
			if (d.id !== undefined) { this.id = d.id }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.CategoryType) {
		return new CreateOfferBasedOnProductActionReq.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.CategoryType>) {
		return new CreateOfferBasedOnProductActionReq.Category(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.CategoryType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Category> {
		return new CreateOfferBasedOnProductActionReq.Category ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Category> {
		return new CreateOfferBasedOnProductActionReq.Category(this.toJSON());
	}
}
/**
  * The base class definition for productSet
  **/
static ProductSet = class ProductSet {
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Product}
  **/
 #product ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet.Product}
  **/
get product () { return this.#product }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Product}
  **/
set product (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.ProductSet.Product) {
			this.#product = value
		} else {
			this.#product = new CreateOfferBasedOnProductActionReq.ProductSet.Product(value)
		}
}
setProduct (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product>) {
	this.product = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Quantity}
  **/
 #quantity ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Quantity>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet.Quantity}
  **/
get quantity () { return this.#quantity }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Quantity}
  **/
set quantity (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Quantity>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.ProductSet.Quantity) {
			this.#quantity = value
		} else {
			this.#quantity = new CreateOfferBasedOnProductActionReq.ProductSet.Quantity(value)
		}
}
setQuantity (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Quantity>) {
	this.quantity = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson}
  **/
 #responsiblePerson ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson}
  **/
get responsiblePerson () { return this.#responsiblePerson }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson}
  **/
set responsiblePerson (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson) {
			this.#responsiblePerson = value
		} else {
			this.#responsiblePerson = new CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson(value)
		}
}
setResponsiblePerson (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson>) {
	this.responsiblePerson = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer}
  **/
 #responsibleProducer ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer}
  **/
get responsibleProducer () { return this.#responsibleProducer }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer}
  **/
set responsibleProducer (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer) {
			this.#responsibleProducer = value
		} else {
			this.#responsibleProducer = new CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer(value)
		}
}
setResponsibleProducer (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer>) {
	this.responsibleProducer = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation}
  **/
 #safetyInformation ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation}
  **/
set safetyInformation (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation(value)
		}
}
setSafetyInformation (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation>) {
	this.safetyInformation = value
	return this
}
		/**
  * 
  * @type {boolean}
  **/
 #marketedBeforeGPSRObligation ! : boolean
		/**
  * 
  * @returns {boolean}
  **/
get marketedBeforeGPSRObligation () { return this.#marketedBeforeGPSRObligation }
/**
  * 
  * @type {boolean}
  **/
set marketedBeforeGPSRObligation (value: boolean) {
		this.#marketedBeforeGPSRObligation = Boolean(value);
}
setMarketedBeforeGPSRObligation (value: boolean) {
	this.marketedBeforeGPSRObligation = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Deposits}
  **/
 #deposits : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Deposits>[]  =  []
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet.Deposits}
  **/
get deposits () { return this.#deposits }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Deposits}
  **/
set deposits (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Deposits>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CreateOfferBasedOnProductActionReq.ProductSet.Deposits) {
			this.#deposits = value
		} else {
			this.#deposits = value.map(item => new CreateOfferBasedOnProductActionReq.ProductSet.Deposits(item))
		}
}
setDeposits (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Deposits>[]) {
	this.deposits = value
	return this
}
/**
  * The base class definition for product
  **/
static Product = class Product {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #idType : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get idType () { return this.#idType }
/**
  * 
  * @type {string}
  **/
set idType (value: string) {
		this.#idType = String(value);
}
setIdType (value: string) {
	this.idType = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Product.Category}
  **/
 #category ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Category>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet.Product.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Product.Category}
  **/
set category (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Category>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.ProductSet.Product.Category) {
			this.#category = value
		} else {
			this.#category = new CreateOfferBasedOnProductActionReq.ProductSet.Product.Category(value)
		}
}
setCategory (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Category>) {
	this.category = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters}
  **/
 #parameters : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters>[]  =  []
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters}
  **/
set parameters (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters) {
			this.#parameters = value
		} else {
			this.#parameters = value.map(item => new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters(item))
		}
}
setParameters (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters>[]) {
	this.parameters = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #images : string[]  =  []
		/**
  * 
  * @returns {string[]}
  **/
get images () { return this.#images }
/**
  * 
  * @type {string[]}
  **/
set images (value: string[]) {
}
setImages (value: string[]) {
	this.images = value
	return this
}
/**
  * The base class definition for category
  **/
static Category = class Category {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
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
		const d = data as Partial<Category>;
			if (d.id !== undefined) { this.id = d.id }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Product.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.CategoryType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Product.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.CategoryType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Category(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.CategoryType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Category> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Category ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Category> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Category(this.toJSON());
	}
}
/**
  * The base class definition for parameters
  **/
static Parameters = class Parameters {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue}
  **/
 #rangeValue ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue}
  **/
get rangeValue () { return this.#rangeValue }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue}
  **/
set rangeValue (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue) {
			this.#rangeValue = value
		} else {
			this.#rangeValue = new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue(value)
		}
}
setRangeValue (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue>) {
	this.rangeValue = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #values : string[]  =  []
		/**
  * 
  * @returns {string[]}
  **/
get values () { return this.#values }
/**
  * 
  * @type {string[]}
  **/
set values (value: string[]) {
}
setValues (value: string[]) {
	this.values = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #valuesIds : string[]  =  []
		/**
  * 
  * @returns {string[]}
  **/
get valuesIds () { return this.#valuesIds }
/**
  * 
  * @type {string[]}
  **/
set valuesIds (value: string[]) {
}
setValuesIds (value: string[]) {
	this.valuesIds = value
	return this
}
/**
  * The base class definition for rangeValue
  **/
static RangeValue = class RangeValue {
		/**
  * 
  * @type {string}
  **/
 #from : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get from () { return this.#from }
/**
  * 
  * @type {string}
  **/
set from (value: string) {
		this.#from = String(value);
}
setFrom (value: string) {
	this.from = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #to : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get to () { return this.#to }
/**
  * 
  * @type {string}
  **/
set to (value: string) {
		this.#to = String(value);
}
setTo (value: string) {
	this.to = value
	return this
}
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
		const d = data as Partial<RangeValue>;
			if (d.from !== undefined) { this.from = d.from }
			if (d.to !== undefined) { this.to = d.to }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				from: this.#from,
				to: this.#to,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			from: 'from',
			to: 'to',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
		const d = data as Partial<Parameters>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
			if (d.rangeValue !== undefined) { this.rangeValue = d.rangeValue }
			if (d.values !== undefined) { this.values = d.values }
			if (d.valuesIds !== undefined) { this.valuesIds = d.valuesIds }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Parameters>;
			if (!(d.rangeValue instanceof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue)) { this.rangeValue = new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue(d.rangeValue || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
				rangeValue: this.#rangeValue,
				values: this.#values,
				valuesIds: this.#valuesIds,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
			rangeValue$: 'rangeValue',
get rangeValue() {
					return withPrefix(
						"productSet.product.parameters.rangeValue",
						CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.RangeValue.Fields
						);
						},
			values$: 'values',
get values() {
					return "productSet.product.parameters.values[:i]";
						},
			valuesIds$: 'valuesIds',
get valuesIds() {
					return "productSet.product.parameters.valuesIds[:i]";
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
		const d = data as Partial<Product>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.idType !== undefined) { this.idType = d.idType }
			if (d.name !== undefined) { this.name = d.name }
			if (d.category !== undefined) { this.category = d.category }
			if (d.parameters !== undefined) { this.parameters = d.parameters }
			if (d.images !== undefined) { this.images = d.images }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Product>;
			if (!(d.category instanceof CreateOfferBasedOnProductActionReq.ProductSet.Product.Category)) { this.category = new CreateOfferBasedOnProductActionReq.ProductSet.Product.Category(d.category || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				idType: this.#idType,
				name: this.#name,
				category: this.#category,
				parameters: this.#parameters,
				images: this.#images,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			idType: 'idType',
			name: 'name',
			category$: 'category',
get category() {
					return withPrefix(
						"productSet.product.category",
						CreateOfferBasedOnProductActionReq.ProductSet.Product.Category.Fields
						);
						},
			parameters$: 'parameters',
get parameters() {
					return withPrefix(
						"productSet.product.parameters[:i]",
						CreateOfferBasedOnProductActionReq.ProductSet.Product.Parameters.Fields
						);
						},
			images$: 'images',
get images() {
					return "productSet.product.images[:i]";
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Product, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Product, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Product> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Product(this.toJSON());
	}
}
/**
  * The base class definition for quantity
  **/
static Quantity = class Quantity {
		/**
  * 
  * @type {number}
  **/
 #value : number  =  0
		/**
  * 
  * @returns {number}
  **/
get value () { return this.#value }
/**
  * 
  * @type {number}
  **/
set value (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#value = parsedValue;
		}
}
setValue (value: number) {
	this.value = value
	return this
}
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
		const d = data as Partial<Quantity>;
			if (d.value !== undefined) { this.value = d.value }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				value: this.#value,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			value: 'value',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Quantity, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType.QuantityType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Quantity(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Quantity, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.QuantityType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Quantity(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.QuantityType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Quantity> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Quantity ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Quantity> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Quantity(this.toJSON());
	}
}
/**
  * The base class definition for responsiblePerson
  **/
static ResponsiblePerson = class ResponsiblePerson {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
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
		const d = data as Partial<ResponsiblePerson>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsiblePersonType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsiblePersonType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsiblePersonType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson(this.toJSON());
	}
}
/**
  * The base class definition for responsibleProducer
  **/
static ResponsibleProducer = class ResponsibleProducer {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #type : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value: string) {
		this.#type = String(value);
}
setType (value: string) {
	this.type = value
	return this
}
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
		const d = data as Partial<ResponsibleProducer>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.type !== undefined) { this.type = d.type }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				type: this.#type,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			type: 'type',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsibleProducerType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsibleProducerType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsibleProducerType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer(this.toJSON());
	}
}
/**
  * The base class definition for safetyInformation
  **/
static SafetyInformation = class SafetyInformation {
		/**
  * 
  * @type {string}
  **/
 #type : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value: string) {
		this.#type = String(value);
}
setType (value: string) {
	this.type = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #description : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get description () { return this.#description }
/**
  * 
  * @type {string}
  **/
set description (value: string) {
		this.#description = String(value);
}
setDescription (value: string) {
	this.description = value
	return this
}
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
		const d = data as Partial<SafetyInformation>;
			if (d.type !== undefined) { this.type = d.type }
			if (d.description !== undefined) { this.description = d.description }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				type: this.#type,
				description: this.#description,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			type: 'type',
			description: 'description',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType.SafetyInformationType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.SafetyInformationType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.SafetyInformationType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation(this.toJSON());
	}
}
/**
  * The base class definition for deposits
  **/
static Deposits = class Deposits {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #quantity : number  =  0
		/**
  * 
  * @returns {number}
  **/
get quantity () { return this.#quantity }
/**
  * 
  * @type {number}
  **/
set quantity (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#quantity = parsedValue;
		}
}
setQuantity (value: number) {
	this.quantity = value
	return this
}
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
		const d = data as Partial<Deposits>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.quantity !== undefined) { this.quantity = d.quantity }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				quantity: this.#quantity,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			quantity: 'quantity',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Deposits, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType.DepositsType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Deposits(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet.Deposits, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.DepositsType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Deposits(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType.DepositsType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Deposits> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Deposits ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet.Deposits> {
		return new CreateOfferBasedOnProductActionReq.ProductSet.Deposits(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
		const d = data as Partial<ProductSet>;
			if (d.product !== undefined) { this.product = d.product }
			if (d.quantity !== undefined) { this.quantity = d.quantity }
			if (d.responsiblePerson !== undefined) { this.responsiblePerson = d.responsiblePerson }
			if (d.responsibleProducer !== undefined) { this.responsibleProducer = d.responsibleProducer }
			if (d.safetyInformation !== undefined) { this.safetyInformation = d.safetyInformation }
			if (d.marketedBeforeGPSRObligation !== undefined) { this.marketedBeforeGPSRObligation = d.marketedBeforeGPSRObligation }
			if (d.deposits !== undefined) { this.deposits = d.deposits }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<ProductSet>;
			if (!(d.product instanceof CreateOfferBasedOnProductActionReq.ProductSet.Product)) { this.product = new CreateOfferBasedOnProductActionReq.ProductSet.Product(d.product || {}) }	
			if (!(d.quantity instanceof CreateOfferBasedOnProductActionReq.ProductSet.Quantity)) { this.quantity = new CreateOfferBasedOnProductActionReq.ProductSet.Quantity(d.quantity || {}) }	
			if (!(d.responsiblePerson instanceof CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson)) { this.responsiblePerson = new CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson(d.responsiblePerson || {}) }	
			if (!(d.responsibleProducer instanceof CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer)) { this.responsibleProducer = new CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer(d.responsibleProducer || {}) }	
			if (!(d.safetyInformation instanceof CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation)) { this.safetyInformation = new CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation(d.safetyInformation || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				product: this.#product,
				quantity: this.#quantity,
				responsiblePerson: this.#responsiblePerson,
				responsibleProducer: this.#responsibleProducer,
				safetyInformation: this.#safetyInformation,
				marketedBeforeGPSRObligation: this.#marketedBeforeGPSRObligation,
				deposits: this.#deposits,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			product$: 'product',
get product() {
					return withPrefix(
						"productSet.product",
						CreateOfferBasedOnProductActionReq.ProductSet.Product.Fields
						);
						},
			quantity$: 'quantity',
get quantity() {
					return withPrefix(
						"productSet.quantity",
						CreateOfferBasedOnProductActionReq.ProductSet.Quantity.Fields
						);
						},
			responsiblePerson$: 'responsiblePerson',
get responsiblePerson() {
					return withPrefix(
						"productSet.responsiblePerson",
						CreateOfferBasedOnProductActionReq.ProductSet.ResponsiblePerson.Fields
						);
						},
			responsibleProducer$: 'responsibleProducer',
get responsibleProducer() {
					return withPrefix(
						"productSet.responsibleProducer",
						CreateOfferBasedOnProductActionReq.ProductSet.ResponsibleProducer.Fields
						);
						},
			safetyInformation$: 'safetyInformation',
get safetyInformation() {
					return withPrefix(
						"productSet.safetyInformation",
						CreateOfferBasedOnProductActionReq.ProductSet.SafetyInformation.Fields
						);
						},
			marketedBeforeGPSRObligation: 'marketedBeforeGPSRObligation',
			deposits$: 'deposits',
get deposits() {
					return withPrefix(
						"productSet.deposits[:i]",
						CreateOfferBasedOnProductActionReq.ProductSet.Deposits.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ProductSetType) {
		return new CreateOfferBasedOnProductActionReq.ProductSet(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.ProductSet, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType>) {
		return new CreateOfferBasedOnProductActionReq.ProductSet(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ProductSetType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet> {
		return new CreateOfferBasedOnProductActionReq.ProductSet ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.ProductSet> {
		return new CreateOfferBasedOnProductActionReq.ProductSet(this.toJSON());
	}
}
/**
  * The base class definition for stock
  **/
static Stock = class Stock {
		/**
  * 
  * @type {number}
  **/
 #available : number  =  0
		/**
  * 
  * @returns {number}
  **/
get available () { return this.#available }
/**
  * 
  * @type {number}
  **/
set available (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#available = parsedValue;
		}
}
setAvailable (value: number) {
	this.available = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #unit : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get unit () { return this.#unit }
/**
  * 
  * @type {string}
  **/
set unit (value: string) {
		this.#unit = String(value);
}
setUnit (value: string) {
	this.unit = value
	return this
}
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
		const d = data as Partial<Stock>;
			if (d.available !== undefined) { this.available = d.available }
			if (d.unit !== undefined) { this.unit = d.unit }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				available: this.#available,
				unit: this.#unit,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			available: 'available',
			unit: 'unit',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Stock, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.StockType) {
		return new CreateOfferBasedOnProductActionReq.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.StockType>) {
		return new CreateOfferBasedOnProductActionReq.Stock(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.StockType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Stock> {
		return new CreateOfferBasedOnProductActionReq.Stock ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Stock> {
		return new CreateOfferBasedOnProductActionReq.Stock(this.toJSON());
	}
}
/**
  * The base class definition for sellingMode
  **/
static SellingMode = class SellingMode {
		/**
  * 
  * @type {string}
  **/
 #format : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get format () { return this.#format }
/**
  * 
  * @type {string}
  **/
set format (value: string) {
		this.#format = String(value);
}
setFormat (value: string) {
	this.format = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SellingMode.Price}
  **/
 #price ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.Price>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SellingMode.Price}
  **/
set price (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.Price>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new CreateOfferBasedOnProductActionReq.SellingMode.Price(value)
		}
}
setPrice (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.Price>) {
	this.price = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice}
  **/
 #minimalPrice ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice}
  **/
get minimalPrice () { return this.#minimalPrice }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice}
  **/
set minimalPrice (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice) {
			this.#minimalPrice = value
		} else {
			this.#minimalPrice = new CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice(value)
		}
}
setMinimalPrice (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice>) {
	this.minimalPrice = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice}
  **/
 #startingPrice ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice}
  **/
get startingPrice () { return this.#startingPrice }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice}
  **/
set startingPrice (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice) {
			this.#startingPrice = value
		} else {
			this.#startingPrice = new CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice(value)
		}
}
setStartingPrice (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice>) {
	this.startingPrice = value
	return this
}
/**
  * The base class definition for price
  **/
static Price = class Price {
		/**
  * 
  * @type {string}
  **/
 #amount : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value: string) {
		this.#amount = String(value);
}
setAmount (value: string) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value: string) {
		this.#currency = String(value);
}
setCurrency (value: string) {
	this.currency = value
	return this
}
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
		const d = data as Partial<Price>;
			if (d.amount !== undefined) { this.amount = d.amount }
			if (d.currency !== undefined) { this.currency = d.currency }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				amount: this.#amount,
				currency: this.#currency,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			amount: 'amount',
			currency: 'currency',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.SellingModeType.PriceType) {
		return new CreateOfferBasedOnProductActionReq.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.SellingModeType.PriceType>) {
		return new CreateOfferBasedOnProductActionReq.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.SellingModeType.PriceType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.Price> {
		return new CreateOfferBasedOnProductActionReq.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.Price> {
		return new CreateOfferBasedOnProductActionReq.SellingMode.Price(this.toJSON());
	}
}
/**
  * The base class definition for minimalPrice
  **/
static MinimalPrice = class MinimalPrice {
		/**
  * 
  * @type {string}
  **/
 #amount : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value: string) {
		this.#amount = String(value);
}
setAmount (value: string) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value: string) {
		this.#currency = String(value);
}
setCurrency (value: string) {
	this.currency = value
	return this
}
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
		const d = data as Partial<MinimalPrice>;
			if (d.amount !== undefined) { this.amount = d.amount }
			if (d.currency !== undefined) { this.currency = d.currency }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				amount: this.#amount,
				currency: this.#currency,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			amount: 'amount',
			currency: 'currency',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.SellingModeType.MinimalPriceType) {
		return new CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.SellingModeType.MinimalPriceType>) {
		return new CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.SellingModeType.MinimalPriceType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice> {
		return new CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice> {
		return new CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice(this.toJSON());
	}
}
/**
  * The base class definition for startingPrice
  **/
static StartingPrice = class StartingPrice {
		/**
  * 
  * @type {string}
  **/
 #amount : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value: string) {
		this.#amount = String(value);
}
setAmount (value: string) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value: string) {
		this.#currency = String(value);
}
setCurrency (value: string) {
	this.currency = value
	return this
}
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
		const d = data as Partial<StartingPrice>;
			if (d.amount !== undefined) { this.amount = d.amount }
			if (d.currency !== undefined) { this.currency = d.currency }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				amount: this.#amount,
				currency: this.#currency,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			amount: 'amount',
			currency: 'currency',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.SellingModeType.StartingPriceType) {
		return new CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.SellingModeType.StartingPriceType>) {
		return new CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.SellingModeType.StartingPriceType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice> {
		return new CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice> {
		return new CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
		const d = data as Partial<SellingMode>;
			if (d.format !== undefined) { this.format = d.format }
			if (d.price !== undefined) { this.price = d.price }
			if (d.minimalPrice !== undefined) { this.minimalPrice = d.minimalPrice }
			if (d.startingPrice !== undefined) { this.startingPrice = d.startingPrice }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<SellingMode>;
			if (!(d.price instanceof CreateOfferBasedOnProductActionReq.SellingMode.Price)) { this.price = new CreateOfferBasedOnProductActionReq.SellingMode.Price(d.price || {}) }	
			if (!(d.minimalPrice instanceof CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice)) { this.minimalPrice = new CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice(d.minimalPrice || {}) }	
			if (!(d.startingPrice instanceof CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice)) { this.startingPrice = new CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice(d.startingPrice || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				format: this.#format,
				price: this.#price,
				minimalPrice: this.#minimalPrice,
				startingPrice: this.#startingPrice,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			format: 'format',
			price$: 'price',
get price() {
					return withPrefix(
						"sellingMode.price",
						CreateOfferBasedOnProductActionReq.SellingMode.Price.Fields
						);
						},
			minimalPrice$: 'minimalPrice',
get minimalPrice() {
					return withPrefix(
						"sellingMode.minimalPrice",
						CreateOfferBasedOnProductActionReq.SellingMode.MinimalPrice.Fields
						);
						},
			startingPrice$: 'startingPrice',
get startingPrice() {
					return withPrefix(
						"sellingMode.startingPrice",
						CreateOfferBasedOnProductActionReq.SellingMode.StartingPrice.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SellingMode, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.SellingModeType) {
		return new CreateOfferBasedOnProductActionReq.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.SellingModeType>) {
		return new CreateOfferBasedOnProductActionReq.SellingMode(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.SellingModeType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode> {
		return new CreateOfferBasedOnProductActionReq.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.SellingMode> {
		return new CreateOfferBasedOnProductActionReq.SellingMode(this.toJSON());
	}
}
/**
  * The base class definition for payments
  **/
static Payments = class Payments {
		/**
  * 
  * @type {string}
  **/
 #invoice : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get invoice () { return this.#invoice }
/**
  * 
  * @type {string}
  **/
set invoice (value: string) {
		this.#invoice = String(value);
}
setInvoice (value: string) {
	this.invoice = value
	return this
}
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
		const d = data as Partial<Payments>;
			if (d.invoice !== undefined) { this.invoice = d.invoice }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				invoice: this.#invoice,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			invoice: 'invoice',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Payments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.PaymentsType) {
		return new CreateOfferBasedOnProductActionReq.Payments(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Payments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.PaymentsType>) {
		return new CreateOfferBasedOnProductActionReq.Payments(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.PaymentsType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Payments> {
		return new CreateOfferBasedOnProductActionReq.Payments ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Payments> {
		return new CreateOfferBasedOnProductActionReq.Payments(this.toJSON());
	}
}
/**
  * The base class definition for delivery
  **/
static Delivery = class Delivery {
		/**
  * 
  * @type {string}
  **/
 #handlingTime : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get handlingTime () { return this.#handlingTime }
/**
  * 
  * @type {string}
  **/
set handlingTime (value: string) {
		this.#handlingTime = String(value);
}
setHandlingTime (value: string) {
	this.handlingTime = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #additionalInfo : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get additionalInfo () { return this.#additionalInfo }
/**
  * 
  * @type {string}
  **/
set additionalInfo (value: string) {
		this.#additionalInfo = String(value);
}
setAdditionalInfo (value: string) {
	this.additionalInfo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #shipmentDate : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get shipmentDate () { return this.#shipmentDate }
/**
  * 
  * @type {string}
  **/
set shipmentDate (value: string) {
		this.#shipmentDate = String(value);
}
setShipmentDate (value: string) {
	this.shipmentDate = value
	return this
}
		/**
  * Optional; may be null
  * @type {CreateOfferBasedOnProductActionReq.Delivery.ShippingRates}
  **/
 #shippingRates ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery.ShippingRates>
		/**
  * Optional; may be null
  * @returns {CreateOfferBasedOnProductActionReq.Delivery.ShippingRates}
  **/
get shippingRates () { return this.#shippingRates }
/**
  * Optional; may be null
  * @type {CreateOfferBasedOnProductActionReq.Delivery.ShippingRates}
  **/
set shippingRates (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery.ShippingRates>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Delivery.ShippingRates) {
			this.#shippingRates = value
		} else {
			this.#shippingRates = new CreateOfferBasedOnProductActionReq.Delivery.ShippingRates(value)
		}
}
setShippingRates (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery.ShippingRates>) {
	this.shippingRates = value
	return this
}
/**
  * The base class definition for shippingRates
  **/
static ShippingRates = class ShippingRates {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
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
		const d = data as Partial<ShippingRates>;
			if (d.id !== undefined) { this.id = d.id }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Delivery.ShippingRates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.DeliveryType.ShippingRatesType) {
		return new CreateOfferBasedOnProductActionReq.Delivery.ShippingRates(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Delivery.ShippingRates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.DeliveryType.ShippingRatesType>) {
		return new CreateOfferBasedOnProductActionReq.Delivery.ShippingRates(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.DeliveryType.ShippingRatesType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery.ShippingRates> {
		return new CreateOfferBasedOnProductActionReq.Delivery.ShippingRates ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery.ShippingRates> {
		return new CreateOfferBasedOnProductActionReq.Delivery.ShippingRates(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
		const d = data as Partial<Delivery>;
			if (d.handlingTime !== undefined) { this.handlingTime = d.handlingTime }
			if (d.additionalInfo !== undefined) { this.additionalInfo = d.additionalInfo }
			if (d.shipmentDate !== undefined) { this.shipmentDate = d.shipmentDate }
			if (d.shippingRates !== undefined) { this.shippingRates = d.shippingRates }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Delivery>;
			if (!(d.shippingRates instanceof CreateOfferBasedOnProductActionReq.Delivery.ShippingRates)) { this.shippingRates = new CreateOfferBasedOnProductActionReq.Delivery.ShippingRates(d.shippingRates || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				handlingTime: this.#handlingTime,
				additionalInfo: this.#additionalInfo,
				shipmentDate: this.#shipmentDate,
				shippingRates: this.#shippingRates,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			handlingTime: 'handlingTime',
			additionalInfo: 'additionalInfo',
			shipmentDate: 'shipmentDate',
			shippingRates$: 'shippingRates',
get shippingRates() {
					return withPrefix(
						"delivery.shippingRates",
						CreateOfferBasedOnProductActionReq.Delivery.ShippingRates.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Delivery, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.DeliveryType) {
		return new CreateOfferBasedOnProductActionReq.Delivery(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Delivery, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.DeliveryType>) {
		return new CreateOfferBasedOnProductActionReq.Delivery(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.DeliveryType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery> {
		return new CreateOfferBasedOnProductActionReq.Delivery ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Delivery> {
		return new CreateOfferBasedOnProductActionReq.Delivery(this.toJSON());
	}
}
/**
  * The base class definition for publication
  **/
static Publication = class Publication {
		/**
  * 
  * @type {string}
  **/
 #duration : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get duration () { return this.#duration }
/**
  * 
  * @type {string}
  **/
set duration (value: string) {
		this.#duration = String(value);
}
setDuration (value: string) {
	this.duration = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #startingAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get startingAt () { return this.#startingAt }
/**
  * 
  * @type {string}
  **/
set startingAt (value: string) {
		this.#startingAt = String(value);
}
setStartingAt (value: string) {
	this.startingAt = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #endingAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get endingAt () { return this.#endingAt }
/**
  * 
  * @type {string}
  **/
set endingAt (value: string) {
		this.#endingAt = String(value);
}
setEndingAt (value: string) {
	this.endingAt = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #status : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get status () { return this.#status }
/**
  * 
  * @type {string}
  **/
set status (value: string) {
		this.#status = String(value);
}
setStatus (value: string) {
	this.status = value
	return this
}
		/**
  * 
  * @type {boolean}
  **/
 #republish ! : boolean
		/**
  * 
  * @returns {boolean}
  **/
get republish () { return this.#republish }
/**
  * 
  * @type {boolean}
  **/
set republish (value: boolean) {
		this.#republish = Boolean(value);
}
setRepublish (value: boolean) {
	this.republish = value
	return this
}
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
		const d = data as Partial<Publication>;
			if (d.duration !== undefined) { this.duration = d.duration }
			if (d.startingAt !== undefined) { this.startingAt = d.startingAt }
			if (d.endingAt !== undefined) { this.endingAt = d.endingAt }
			if (d.status !== undefined) { this.status = d.status }
			if (d.republish !== undefined) { this.republish = d.republish }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				duration: this.#duration,
				startingAt: this.#startingAt,
				endingAt: this.#endingAt,
				status: this.#status,
				republish: this.#republish,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			duration: 'duration',
			startingAt: 'startingAt',
			endingAt: 'endingAt',
			status: 'status',
			republish: 'republish',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.PublicationType) {
		return new CreateOfferBasedOnProductActionReq.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.PublicationType>) {
		return new CreateOfferBasedOnProductActionReq.Publication(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.PublicationType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Publication> {
		return new CreateOfferBasedOnProductActionReq.Publication ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Publication> {
		return new CreateOfferBasedOnProductActionReq.Publication(this.toJSON());
	}
}
/**
  * The base class definition for compatibilityList
  **/
static CompatibilityList = class CompatibilityList {
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.CompatibilityList.Items}
  **/
 #items : InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList.Items>[]  =  []
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.CompatibilityList.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.CompatibilityList.Items}
  **/
set items (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList.Items>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CreateOfferBasedOnProductActionReq.CompatibilityList.Items) {
			this.#items = value
		} else {
			this.#items = value.map(item => new CreateOfferBasedOnProductActionReq.CompatibilityList.Items(item))
		}
}
setItems (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList.Items>[]) {
	this.items = value
	return this
}
/**
  * The base class definition for items
  **/
static Items = class Items {
		/**
  * 
  * @type {string}
  **/
 #type : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value: string) {
		this.#type = String(value);
}
setType (value: string) {
	this.type = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #text : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get text () { return this.#text }
/**
  * 
  * @type {string}
  **/
set text (value: string) {
		this.#text = String(value);
}
setText (value: string) {
	this.text = value
	return this
}
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
		const d = data as Partial<Items>;
			if (d.type !== undefined) { this.type = d.type }
			if (d.text !== undefined) { this.text = d.text }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				type: this.#type,
				text: this.#text,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			type: 'type',
			text: 'text',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.CompatibilityList.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.CompatibilityListType.ItemsType) {
		return new CreateOfferBasedOnProductActionReq.CompatibilityList.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.CompatibilityList.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.CompatibilityListType.ItemsType>) {
		return new CreateOfferBasedOnProductActionReq.CompatibilityList.Items(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.CompatibilityListType.ItemsType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList.Items> {
		return new CreateOfferBasedOnProductActionReq.CompatibilityList.Items ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList.Items> {
		return new CreateOfferBasedOnProductActionReq.CompatibilityList.Items(this.toJSON());
	}
}
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
		const d = data as Partial<CompatibilityList>;
			if (d.items !== undefined) { this.items = d.items }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				items: this.#items,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			items$: 'items',
get items() {
					return withPrefix(
						"compatibilityList.items[:i]",
						CreateOfferBasedOnProductActionReq.CompatibilityList.Items.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.CompatibilityList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.CompatibilityListType) {
		return new CreateOfferBasedOnProductActionReq.CompatibilityList(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.CompatibilityList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.CompatibilityListType>) {
		return new CreateOfferBasedOnProductActionReq.CompatibilityList(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.CompatibilityListType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList> {
		return new CreateOfferBasedOnProductActionReq.CompatibilityList ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.CompatibilityList> {
		return new CreateOfferBasedOnProductActionReq.CompatibilityList(this.toJSON());
	}
}
/**
  * The base class definition for description
  **/
static Description = class Description {
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Description.Sections}
  **/
 #sections : InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections>[]  =  []
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Description.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Description.Sections}
  **/
set sections (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CreateOfferBasedOnProductActionReq.Description.Sections) {
			this.#sections = value
		} else {
			this.#sections = value.map(item => new CreateOfferBasedOnProductActionReq.Description.Sections(item))
		}
}
setSections (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections>[]) {
	this.sections = value
	return this
}
/**
  * The base class definition for sections
  **/
static Sections = class Sections {
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Description.Sections.Items}
  **/
 #items : InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections.Items>[]  =  []
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Description.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Description.Sections.Items}
  **/
set items (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections.Items>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CreateOfferBasedOnProductActionReq.Description.Sections.Items) {
			this.#items = value
		} else {
			this.#items = value.map(item => new CreateOfferBasedOnProductActionReq.Description.Sections.Items(item))
		}
}
setItems (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections.Items>[]) {
	this.items = value
	return this
}
/**
  * The base class definition for items
  **/
static Items = class Items {
		/**
  * 
  * @type {string}
  **/
 #type : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value: string) {
		this.#type = String(value);
}
setType (value: string) {
	this.type = value
	return this
}
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
		const d = data as Partial<Items>;
			if (d.type !== undefined) { this.type = d.type }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				type: this.#type,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			type: 'type',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Description.Sections.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType.ItemsType) {
		return new CreateOfferBasedOnProductActionReq.Description.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Description.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType.ItemsType>) {
		return new CreateOfferBasedOnProductActionReq.Description.Sections.Items(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType.ItemsType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections.Items> {
		return new CreateOfferBasedOnProductActionReq.Description.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections.Items> {
		return new CreateOfferBasedOnProductActionReq.Description.Sections.Items(this.toJSON());
	}
}
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
		const d = data as Partial<Sections>;
			if (d.items !== undefined) { this.items = d.items }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				items: this.#items,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			items$: 'items',
get items() {
					return withPrefix(
						"description.sections.items[:i]",
						CreateOfferBasedOnProductActionReq.Description.Sections.Items.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Description.Sections, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType) {
		return new CreateOfferBasedOnProductActionReq.Description.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Description.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType>) {
		return new CreateOfferBasedOnProductActionReq.Description.Sections(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections> {
		return new CreateOfferBasedOnProductActionReq.Description.Sections ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Description.Sections> {
		return new CreateOfferBasedOnProductActionReq.Description.Sections(this.toJSON());
	}
}
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
		const d = data as Partial<Description>;
			if (d.sections !== undefined) { this.sections = d.sections }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				sections: this.#sections,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			sections$: 'sections',
get sections() {
					return withPrefix(
						"description.sections[:i]",
						CreateOfferBasedOnProductActionReq.Description.Sections.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Description, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.DescriptionType) {
		return new CreateOfferBasedOnProductActionReq.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.DescriptionType>) {
		return new CreateOfferBasedOnProductActionReq.Description(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.DescriptionType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Description> {
		return new CreateOfferBasedOnProductActionReq.Description ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Description> {
		return new CreateOfferBasedOnProductActionReq.Description(this.toJSON());
	}
}
/**
  * The base class definition for b2b
  **/
static B2b = class B2b {
		/**
  * 
  * @type {boolean}
  **/
 #buyableOnlyByBusiness ! : boolean
		/**
  * 
  * @returns {boolean}
  **/
get buyableOnlyByBusiness () { return this.#buyableOnlyByBusiness }
/**
  * 
  * @type {boolean}
  **/
set buyableOnlyByBusiness (value: boolean) {
		this.#buyableOnlyByBusiness = Boolean(value);
}
setBuyableOnlyByBusiness (value: boolean) {
	this.buyableOnlyByBusiness = value
	return this
}
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
		const d = data as Partial<B2b>;
			if (d.buyableOnlyByBusiness !== undefined) { this.buyableOnlyByBusiness = d.buyableOnlyByBusiness }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				buyableOnlyByBusiness: this.#buyableOnlyByBusiness,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			buyableOnlyByBusiness: 'buyableOnlyByBusiness',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.B2b, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.B2bType) {
		return new CreateOfferBasedOnProductActionReq.B2b(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.B2b, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.B2bType>) {
		return new CreateOfferBasedOnProductActionReq.B2b(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.B2bType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.B2b> {
		return new CreateOfferBasedOnProductActionReq.B2b ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.B2b> {
		return new CreateOfferBasedOnProductActionReq.B2b(this.toJSON());
	}
}
/**
  * The base class definition for attachments
  **/
static Attachments = class Attachments {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
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
		const d = data as Partial<Attachments>;
			if (d.id !== undefined) { this.id = d.id }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Attachments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.AttachmentsType) {
		return new CreateOfferBasedOnProductActionReq.Attachments(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Attachments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.AttachmentsType>) {
		return new CreateOfferBasedOnProductActionReq.Attachments(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.AttachmentsType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Attachments> {
		return new CreateOfferBasedOnProductActionReq.Attachments ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Attachments> {
		return new CreateOfferBasedOnProductActionReq.Attachments(this.toJSON());
	}
}
/**
  * The base class definition for fundraisingCampaign
  **/
static FundraisingCampaign = class FundraisingCampaign {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
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
		const d = data as Partial<FundraisingCampaign>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.FundraisingCampaign, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.FundraisingCampaignType) {
		return new CreateOfferBasedOnProductActionReq.FundraisingCampaign(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.FundraisingCampaign, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.FundraisingCampaignType>) {
		return new CreateOfferBasedOnProductActionReq.FundraisingCampaign(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.FundraisingCampaignType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.FundraisingCampaign> {
		return new CreateOfferBasedOnProductActionReq.FundraisingCampaign ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.FundraisingCampaign> {
		return new CreateOfferBasedOnProductActionReq.FundraisingCampaign(this.toJSON());
	}
}
/**
  * The base class definition for additionalServices
  **/
static AdditionalServices = class AdditionalServices {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
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
		const d = data as Partial<AdditionalServices>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AdditionalServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.AdditionalServicesType) {
		return new CreateOfferBasedOnProductActionReq.AdditionalServices(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AdditionalServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.AdditionalServicesType>) {
		return new CreateOfferBasedOnProductActionReq.AdditionalServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.AdditionalServicesType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.AdditionalServices> {
		return new CreateOfferBasedOnProductActionReq.AdditionalServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.AdditionalServices> {
		return new CreateOfferBasedOnProductActionReq.AdditionalServices(this.toJSON());
	}
}
/**
  * The base class definition for afterSalesServices
  **/
static AfterSalesServices = class AfterSalesServices {
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty}
  **/
 #impliedWarranty ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty}
  **/
get impliedWarranty () { return this.#impliedWarranty }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty}
  **/
set impliedWarranty (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty) {
			this.#impliedWarranty = value
		} else {
			this.#impliedWarranty = new CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty(value)
		}
}
setImpliedWarranty (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty>) {
	this.impliedWarranty = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy}
  **/
 #returnPolicy ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy}
  **/
get returnPolicy () { return this.#returnPolicy }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy}
  **/
set returnPolicy (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy) {
			this.#returnPolicy = value
		} else {
			this.#returnPolicy = new CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy(value)
		}
}
setReturnPolicy (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy>) {
	this.returnPolicy = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty}
  **/
 #warranty ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty}
  **/
get warranty () { return this.#warranty }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty}
  **/
set warranty (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty) {
			this.#warranty = value
		} else {
			this.#warranty = new CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty(value)
		}
}
setWarranty (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty>) {
	this.warranty = value
	return this
}
/**
  * The base class definition for impliedWarranty
  **/
static ImpliedWarranty = class ImpliedWarranty {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
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
		const d = data as Partial<ImpliedWarranty>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ImpliedWarrantyType) {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ImpliedWarrantyType>) {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ImpliedWarrantyType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty> {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty> {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty(this.toJSON());
	}
}
/**
  * The base class definition for returnPolicy
  **/
static ReturnPolicy = class ReturnPolicy {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
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
		const d = data as Partial<ReturnPolicy>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ReturnPolicyType) {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ReturnPolicyType>) {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ReturnPolicyType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy> {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy> {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy(this.toJSON());
	}
}
/**
  * The base class definition for warranty
  **/
static Warranty = class Warranty {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
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
		const d = data as Partial<Warranty>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.WarrantyType) {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.WarrantyType>) {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.WarrantyType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty> {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty> {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
		const d = data as Partial<AfterSalesServices>;
			if (d.impliedWarranty !== undefined) { this.impliedWarranty = d.impliedWarranty }
			if (d.returnPolicy !== undefined) { this.returnPolicy = d.returnPolicy }
			if (d.warranty !== undefined) { this.warranty = d.warranty }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<AfterSalesServices>;
			if (!(d.impliedWarranty instanceof CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty)) { this.impliedWarranty = new CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty(d.impliedWarranty || {}) }	
			if (!(d.returnPolicy instanceof CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy)) { this.returnPolicy = new CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy(d.returnPolicy || {}) }	
			if (!(d.warranty instanceof CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty)) { this.warranty = new CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty(d.warranty || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				impliedWarranty: this.#impliedWarranty,
				returnPolicy: this.#returnPolicy,
				warranty: this.#warranty,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			impliedWarranty$: 'impliedWarranty',
get impliedWarranty() {
					return withPrefix(
						"afterSalesServices.impliedWarranty",
						CreateOfferBasedOnProductActionReq.AfterSalesServices.ImpliedWarranty.Fields
						);
						},
			returnPolicy$: 'returnPolicy',
get returnPolicy() {
					return withPrefix(
						"afterSalesServices.returnPolicy",
						CreateOfferBasedOnProductActionReq.AfterSalesServices.ReturnPolicy.Fields
						);
						},
			warranty$: 'warranty',
get warranty() {
					return withPrefix(
						"afterSalesServices.warranty",
						CreateOfferBasedOnProductActionReq.AfterSalesServices.Warranty.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AfterSalesServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.AfterSalesServicesType) {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.AfterSalesServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.AfterSalesServicesType>) {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.AfterSalesServicesType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices> {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.AfterSalesServices> {
		return new CreateOfferBasedOnProductActionReq.AfterSalesServices(this.toJSON());
	}
}
/**
  * The base class definition for sizeTable
  **/
static SizeTable = class SizeTable {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
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
		const d = data as Partial<SizeTable>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SizeTable, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.SizeTableType) {
		return new CreateOfferBasedOnProductActionReq.SizeTable(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.SizeTable, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.SizeTableType>) {
		return new CreateOfferBasedOnProductActionReq.SizeTable(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.SizeTableType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.SizeTable> {
		return new CreateOfferBasedOnProductActionReq.SizeTable ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.SizeTable> {
		return new CreateOfferBasedOnProductActionReq.SizeTable(this.toJSON());
	}
}
/**
  * The base class definition for contact
  **/
static Contact = class Contact {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
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
		const d = data as Partial<Contact>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Contact, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ContactType) {
		return new CreateOfferBasedOnProductActionReq.Contact(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Contact, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ContactType>) {
		return new CreateOfferBasedOnProductActionReq.Contact(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ContactType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Contact> {
		return new CreateOfferBasedOnProductActionReq.Contact ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Contact> {
		return new CreateOfferBasedOnProductActionReq.Contact(this.toJSON());
	}
}
/**
  * The base class definition for discounts
  **/
static Discounts = class Discounts {
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList}
  **/
 #wholesalePriceList ! : InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList>
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList}
  **/
get wholesalePriceList () { return this.#wholesalePriceList }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList}
  **/
set wholesalePriceList (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList) {
			this.#wholesalePriceList = value
		} else {
			this.#wholesalePriceList = new CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList(value)
		}
}
setWholesalePriceList (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList>) {
	this.wholesalePriceList = value
	return this
}
/**
  * The base class definition for wholesalePriceList
  **/
static WholesalePriceList = class WholesalePriceList {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value: string) {
		this.#name = String(value);
}
setName (value: string) {
	this.name = value
	return this
}
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
		const d = data as Partial<WholesalePriceList>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.DiscountsType.WholesalePriceListType) {
		return new CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.DiscountsType.WholesalePriceListType>) {
		return new CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.DiscountsType.WholesalePriceListType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList> {
		return new CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList> {
		return new CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
		const d = data as Partial<Discounts>;
			if (d.wholesalePriceList !== undefined) { this.wholesalePriceList = d.wholesalePriceList }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Discounts>;
			if (!(d.wholesalePriceList instanceof CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList)) { this.wholesalePriceList = new CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList(d.wholesalePriceList || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				wholesalePriceList: this.#wholesalePriceList,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			wholesalePriceList$: 'wholesalePriceList',
get wholesalePriceList() {
					return withPrefix(
						"discounts.wholesalePriceList",
						CreateOfferBasedOnProductActionReq.Discounts.WholesalePriceList.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Discounts, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.DiscountsType) {
		return new CreateOfferBasedOnProductActionReq.Discounts(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Discounts, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.DiscountsType>) {
		return new CreateOfferBasedOnProductActionReq.Discounts(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.DiscountsType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts> {
		return new CreateOfferBasedOnProductActionReq.Discounts ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Discounts> {
		return new CreateOfferBasedOnProductActionReq.Discounts(this.toJSON());
	}
}
/**
  * The base class definition for location
  **/
static Location = class Location {
		/**
  * 
  * @type {string}
  **/
 #city : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get city () { return this.#city }
/**
  * 
  * @type {string}
  **/
set city (value: string) {
		this.#city = String(value);
}
setCity (value: string) {
	this.city = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #countryCode : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get countryCode () { return this.#countryCode }
/**
  * 
  * @type {string}
  **/
set countryCode (value: string) {
		this.#countryCode = String(value);
}
setCountryCode (value: string) {
	this.countryCode = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #postCode : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get postCode () { return this.#postCode }
/**
  * 
  * @type {string}
  **/
set postCode (value: string) {
		this.#postCode = String(value);
}
setPostCode (value: string) {
	this.postCode = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #province : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get province () { return this.#province }
/**
  * 
  * @type {string}
  **/
set province (value: string) {
		this.#province = String(value);
}
setProvince (value: string) {
	this.province = value
	return this
}
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
		const d = data as Partial<Location>;
			if (d.city !== undefined) { this.city = d.city }
			if (d.countryCode !== undefined) { this.countryCode = d.countryCode }
			if (d.postCode !== undefined) { this.postCode = d.postCode }
			if (d.province !== undefined) { this.province = d.province }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				city: this.#city,
				countryCode: this.#countryCode,
				postCode: this.#postCode,
				province: this.#province,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			city: 'city',
			countryCode: 'countryCode',
			postCode: 'postCode',
			province: 'province',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Location, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.LocationType) {
		return new CreateOfferBasedOnProductActionReq.Location(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.Location, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.LocationType>) {
		return new CreateOfferBasedOnProductActionReq.Location(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.LocationType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.Location> {
		return new CreateOfferBasedOnProductActionReq.Location ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.Location> {
		return new CreateOfferBasedOnProductActionReq.Location(this.toJSON());
	}
}
/**
  * The base class definition for external
  **/
static External = class External {
		/**
  * 
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value: string) {
		this.#id = String(value);
}
setId (value: string) {
	this.id = value
	return this
}
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
		const d = data as Partial<External>;
			if (d.id !== undefined) { this.id = d.id }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.External, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.ExternalType) {
		return new CreateOfferBasedOnProductActionReq.External(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.External, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.ExternalType>) {
		return new CreateOfferBasedOnProductActionReq.External(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.ExternalType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.External> {
		return new CreateOfferBasedOnProductActionReq.External ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.External> {
		return new CreateOfferBasedOnProductActionReq.External(this.toJSON());
	}
}
/**
  * The base class definition for taxSettings
  **/
static TaxSettings = class TaxSettings {
		/**
  * 
  * @type {string}
  **/
 #subject : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get subject () { return this.#subject }
/**
  * 
  * @type {string}
  **/
set subject (value: string) {
		this.#subject = String(value);
}
setSubject (value: string) {
	this.subject = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #exemption : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get exemption () { return this.#exemption }
/**
  * 
  * @type {string}
  **/
set exemption (value: string) {
		this.#exemption = String(value);
}
setExemption (value: string) {
	this.exemption = value
	return this
}
		/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.TaxSettings.Rates}
  **/
 #rates : InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings.Rates>[]  =  []
		/**
  * 
  * @returns {CreateOfferBasedOnProductActionReq.TaxSettings.Rates}
  **/
get rates () { return this.#rates }
/**
  * 
  * @type {CreateOfferBasedOnProductActionReq.TaxSettings.Rates}
  **/
set rates (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings.Rates>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof CreateOfferBasedOnProductActionReq.TaxSettings.Rates) {
			this.#rates = value
		} else {
			this.#rates = value.map(item => new CreateOfferBasedOnProductActionReq.TaxSettings.Rates(item))
		}
}
setRates (value: InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings.Rates>[]) {
	this.rates = value
	return this
}
/**
  * The base class definition for rates
  **/
static Rates = class Rates {
		/**
  * 
  * @type {string}
  **/
 #rate : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get rate () { return this.#rate }
/**
  * 
  * @type {string}
  **/
set rate (value: string) {
		this.#rate = String(value);
}
setRate (value: string) {
	this.rate = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #countryCode : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get countryCode () { return this.#countryCode }
/**
  * 
  * @type {string}
  **/
set countryCode (value: string) {
		this.#countryCode = String(value);
}
setCountryCode (value: string) {
	this.countryCode = value
	return this
}
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
		const d = data as Partial<Rates>;
			if (d.rate !== undefined) { this.rate = d.rate }
			if (d.countryCode !== undefined) { this.countryCode = d.countryCode }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				rate: this.#rate,
				countryCode: this.#countryCode,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			rate: 'rate',
			countryCode: 'countryCode',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.TaxSettings.Rates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.TaxSettingsType.RatesType) {
		return new CreateOfferBasedOnProductActionReq.TaxSettings.Rates(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.TaxSettings.Rates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.TaxSettingsType.RatesType>) {
		return new CreateOfferBasedOnProductActionReq.TaxSettings.Rates(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.TaxSettingsType.RatesType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings.Rates> {
		return new CreateOfferBasedOnProductActionReq.TaxSettings.Rates ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings.Rates> {
		return new CreateOfferBasedOnProductActionReq.TaxSettings.Rates(this.toJSON());
	}
}
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
		const d = data as Partial<TaxSettings>;
			if (d.subject !== undefined) { this.subject = d.subject }
			if (d.exemption !== undefined) { this.exemption = d.exemption }
			if (d.rates !== undefined) { this.rates = d.rates }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				subject: this.#subject,
				exemption: this.#exemption,
				rates: this.#rates,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			subject: 'subject',
			exemption: 'exemption',
			rates$: 'rates',
get rates() {
					return withPrefix(
						"taxSettings.rates[:i]",
						CreateOfferBasedOnProductActionReq.TaxSettings.Rates.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.TaxSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.TaxSettingsType) {
		return new CreateOfferBasedOnProductActionReq.TaxSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.TaxSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.TaxSettingsType>) {
		return new CreateOfferBasedOnProductActionReq.TaxSettings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.TaxSettingsType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings> {
		return new CreateOfferBasedOnProductActionReq.TaxSettings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.TaxSettings> {
		return new CreateOfferBasedOnProductActionReq.TaxSettings(this.toJSON());
	}
}
/**
  * The base class definition for messageToSellerSettings
  **/
static MessageToSellerSettings = class MessageToSellerSettings {
		/**
  * 
  * @type {string}
  **/
 #mode : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get mode () { return this.#mode }
/**
  * 
  * @type {string}
  **/
set mode (value: string) {
		this.#mode = String(value);
}
setMode (value: string) {
	this.mode = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #hint : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get hint () { return this.#hint }
/**
  * 
  * @type {string}
  **/
set hint (value: string) {
		this.#hint = String(value);
}
setHint (value: string) {
	this.hint = value
	return this
}
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
		const d = data as Partial<MessageToSellerSettings>;
			if (d.mode !== undefined) { this.mode = d.mode }
			if (d.hint !== undefined) { this.hint = d.hint }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				mode: this.#mode,
				hint: this.#hint,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			mode: 'mode',
			hint: 'hint',
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.MessageToSellerSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType.MessageToSellerSettingsType) {
		return new CreateOfferBasedOnProductActionReq.MessageToSellerSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq.MessageToSellerSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType.MessageToSellerSettingsType>) {
		return new CreateOfferBasedOnProductActionReq.MessageToSellerSettings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType.MessageToSellerSettingsType>): InstanceType<typeof CreateOfferBasedOnProductActionReq.MessageToSellerSettings> {
		return new CreateOfferBasedOnProductActionReq.MessageToSellerSettings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq.MessageToSellerSettings> {
		return new CreateOfferBasedOnProductActionReq.MessageToSellerSettings(this.toJSON());
	}
}
	constructor(data: unknown = undefined) {
		if (data === null || data === undefined) {
				this.#lateInitFields();
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
		const d = data as Partial<CreateOfferBasedOnProductActionReq>;
			if (d.name !== undefined) { this.name = d.name }
			if (d.language !== undefined) { this.language = d.language }
			if (d.category !== undefined) { this.category = d.category }
			if (d.productSet !== undefined) { this.productSet = d.productSet }
			if (d.stock !== undefined) { this.stock = d.stock }
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
			if (d.payments !== undefined) { this.payments = d.payments }
			if (d.delivery !== undefined) { this.delivery = d.delivery }
			if (d.publication !== undefined) { this.publication = d.publication }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
			if (d.compatibilityList !== undefined) { this.compatibilityList = d.compatibilityList }
			if (d.images !== undefined) { this.images = d.images }
			if (d.description !== undefined) { this.description = d.description }
			if (d.b2b !== undefined) { this.b2b = d.b2b }
			if (d.attachments !== undefined) { this.attachments = d.attachments }
			if (d.fundraisingCampaign !== undefined) { this.fundraisingCampaign = d.fundraisingCampaign }
			if (d.additionalServices !== undefined) { this.additionalServices = d.additionalServices }
			if (d.afterSalesServices !== undefined) { this.afterSalesServices = d.afterSalesServices }
			if (d.sizeTable !== undefined) { this.sizeTable = d.sizeTable }
			if (d.contact !== undefined) { this.contact = d.contact }
			if (d.discounts !== undefined) { this.discounts = d.discounts }
			if (d.location !== undefined) { this.location = d.location }
			if (d.external !== undefined) { this.external = d.external }
			if (d.taxSettings !== undefined) { this.taxSettings = d.taxSettings }
			if (d.messageToSellerSettings !== undefined) { this.messageToSellerSettings = d.messageToSellerSettings }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<CreateOfferBasedOnProductActionReq>;
			if (!(d.category instanceof CreateOfferBasedOnProductActionReq.Category)) { this.category = new CreateOfferBasedOnProductActionReq.Category(d.category || {}) }	
			if (!(d.stock instanceof CreateOfferBasedOnProductActionReq.Stock)) { this.stock = new CreateOfferBasedOnProductActionReq.Stock(d.stock || {}) }	
			if (!(d.sellingMode instanceof CreateOfferBasedOnProductActionReq.SellingMode)) { this.sellingMode = new CreateOfferBasedOnProductActionReq.SellingMode(d.sellingMode || {}) }	
			if (!(d.payments instanceof CreateOfferBasedOnProductActionReq.Payments)) { this.payments = new CreateOfferBasedOnProductActionReq.Payments(d.payments || {}) }	
			if (!(d.delivery instanceof CreateOfferBasedOnProductActionReq.Delivery)) { this.delivery = new CreateOfferBasedOnProductActionReq.Delivery(d.delivery || {}) }	
			if (!(d.publication instanceof CreateOfferBasedOnProductActionReq.Publication)) { this.publication = new CreateOfferBasedOnProductActionReq.Publication(d.publication || {}) }	
			if (!(d.compatibilityList instanceof CreateOfferBasedOnProductActionReq.CompatibilityList)) { this.compatibilityList = new CreateOfferBasedOnProductActionReq.CompatibilityList(d.compatibilityList || {}) }	
			if (!(d.description instanceof CreateOfferBasedOnProductActionReq.Description)) { this.description = new CreateOfferBasedOnProductActionReq.Description(d.description || {}) }	
			if (!(d.b2b instanceof CreateOfferBasedOnProductActionReq.B2b)) { this.b2b = new CreateOfferBasedOnProductActionReq.B2b(d.b2b || {}) }	
			if (!(d.fundraisingCampaign instanceof CreateOfferBasedOnProductActionReq.FundraisingCampaign)) { this.fundraisingCampaign = new CreateOfferBasedOnProductActionReq.FundraisingCampaign(d.fundraisingCampaign || {}) }	
			if (!(d.additionalServices instanceof CreateOfferBasedOnProductActionReq.AdditionalServices)) { this.additionalServices = new CreateOfferBasedOnProductActionReq.AdditionalServices(d.additionalServices || {}) }	
			if (!(d.afterSalesServices instanceof CreateOfferBasedOnProductActionReq.AfterSalesServices)) { this.afterSalesServices = new CreateOfferBasedOnProductActionReq.AfterSalesServices(d.afterSalesServices || {}) }	
			if (!(d.sizeTable instanceof CreateOfferBasedOnProductActionReq.SizeTable)) { this.sizeTable = new CreateOfferBasedOnProductActionReq.SizeTable(d.sizeTable || {}) }	
			if (!(d.contact instanceof CreateOfferBasedOnProductActionReq.Contact)) { this.contact = new CreateOfferBasedOnProductActionReq.Contact(d.contact || {}) }	
			if (!(d.discounts instanceof CreateOfferBasedOnProductActionReq.Discounts)) { this.discounts = new CreateOfferBasedOnProductActionReq.Discounts(d.discounts || {}) }	
			if (!(d.location instanceof CreateOfferBasedOnProductActionReq.Location)) { this.location = new CreateOfferBasedOnProductActionReq.Location(d.location || {}) }	
			if (!(d.external instanceof CreateOfferBasedOnProductActionReq.External)) { this.external = new CreateOfferBasedOnProductActionReq.External(d.external || {}) }	
			if (!(d.taxSettings instanceof CreateOfferBasedOnProductActionReq.TaxSettings)) { this.taxSettings = new CreateOfferBasedOnProductActionReq.TaxSettings(d.taxSettings || {}) }	
			if (!(d.messageToSellerSettings instanceof CreateOfferBasedOnProductActionReq.MessageToSellerSettings)) { this.messageToSellerSettings = new CreateOfferBasedOnProductActionReq.MessageToSellerSettings(d.messageToSellerSettings || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				name: this.#name,
				language: this.#language,
				category: this.#category,
				productSet: this.#productSet,
				stock: this.#stock,
				sellingMode: this.#sellingMode,
				payments: this.#payments,
				delivery: this.#delivery,
				publication: this.#publication,
				additionalMarketplaces: this.#additionalMarketplaces,
				compatibilityList: this.#compatibilityList,
				images: this.#images,
				description: this.#description,
				b2b: this.#b2b,
				attachments: this.#attachments,
				fundraisingCampaign: this.#fundraisingCampaign,
				additionalServices: this.#additionalServices,
				afterSalesServices: this.#afterSalesServices,
				sizeTable: this.#sizeTable,
				contact: this.#contact,
				discounts: this.#discounts,
				location: this.#location,
				external: this.#external,
				taxSettings: this.#taxSettings,
				messageToSellerSettings: this.#messageToSellerSettings,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			name: 'name',
			language: 'language',
			category$: 'category',
get category() {
					return withPrefix(
						"category",
						CreateOfferBasedOnProductActionReq.Category.Fields
						);
						},
			productSet$: 'productSet',
get productSet() {
					return withPrefix(
						"productSet[:i]",
						CreateOfferBasedOnProductActionReq.ProductSet.Fields
						);
						},
			stock$: 'stock',
get stock() {
					return withPrefix(
						"stock",
						CreateOfferBasedOnProductActionReq.Stock.Fields
						);
						},
			sellingMode$: 'sellingMode',
get sellingMode() {
					return withPrefix(
						"sellingMode",
						CreateOfferBasedOnProductActionReq.SellingMode.Fields
						);
						},
			payments$: 'payments',
get payments() {
					return withPrefix(
						"payments",
						CreateOfferBasedOnProductActionReq.Payments.Fields
						);
						},
			delivery$: 'delivery',
get delivery() {
					return withPrefix(
						"delivery",
						CreateOfferBasedOnProductActionReq.Delivery.Fields
						);
						},
			publication$: 'publication',
get publication() {
					return withPrefix(
						"publication",
						CreateOfferBasedOnProductActionReq.Publication.Fields
						);
						},
			additionalMarketplaces: 'additionalMarketplaces',
			compatibilityList$: 'compatibilityList',
get compatibilityList() {
					return withPrefix(
						"compatibilityList",
						CreateOfferBasedOnProductActionReq.CompatibilityList.Fields
						);
						},
			images$: 'images',
get images() {
					return "images[:i]";
						},
			description$: 'description',
get description() {
					return withPrefix(
						"description",
						CreateOfferBasedOnProductActionReq.Description.Fields
						);
						},
			b2b$: 'b2b',
get b2b() {
					return withPrefix(
						"b2b",
						CreateOfferBasedOnProductActionReq.B2b.Fields
						);
						},
			attachments$: 'attachments',
get attachments() {
					return withPrefix(
						"attachments[:i]",
						CreateOfferBasedOnProductActionReq.Attachments.Fields
						);
						},
			fundraisingCampaign$: 'fundraisingCampaign',
get fundraisingCampaign() {
					return withPrefix(
						"fundraisingCampaign",
						CreateOfferBasedOnProductActionReq.FundraisingCampaign.Fields
						);
						},
			additionalServices$: 'additionalServices',
get additionalServices() {
					return withPrefix(
						"additionalServices",
						CreateOfferBasedOnProductActionReq.AdditionalServices.Fields
						);
						},
			afterSalesServices$: 'afterSalesServices',
get afterSalesServices() {
					return withPrefix(
						"afterSalesServices",
						CreateOfferBasedOnProductActionReq.AfterSalesServices.Fields
						);
						},
			sizeTable$: 'sizeTable',
get sizeTable() {
					return withPrefix(
						"sizeTable",
						CreateOfferBasedOnProductActionReq.SizeTable.Fields
						);
						},
			contact$: 'contact',
get contact() {
					return withPrefix(
						"contact",
						CreateOfferBasedOnProductActionReq.Contact.Fields
						);
						},
			discounts$: 'discounts',
get discounts() {
					return withPrefix(
						"discounts",
						CreateOfferBasedOnProductActionReq.Discounts.Fields
						);
						},
			location$: 'location',
get location() {
					return withPrefix(
						"location",
						CreateOfferBasedOnProductActionReq.Location.Fields
						);
						},
			external$: 'external',
get external() {
					return withPrefix(
						"external",
						CreateOfferBasedOnProductActionReq.External.Fields
						);
						},
			taxSettings$: 'taxSettings',
get taxSettings() {
					return withPrefix(
						"taxSettings",
						CreateOfferBasedOnProductActionReq.TaxSettings.Fields
						);
						},
			messageToSellerSettings$: 'messageToSellerSettings',
get messageToSellerSettings() {
					return withPrefix(
						"messageToSellerSettings",
						CreateOfferBasedOnProductActionReq.MessageToSellerSettings.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: CreateOfferBasedOnProductActionReqType) {
		return new CreateOfferBasedOnProductActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of CreateOfferBasedOnProductActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<CreateOfferBasedOnProductActionReqType>) {
		return new CreateOfferBasedOnProductActionReq(partialDtoObject);
	}
	copyWith(partial: PartialDeep<CreateOfferBasedOnProductActionReqType>): InstanceType<typeof CreateOfferBasedOnProductActionReq> {
		return new CreateOfferBasedOnProductActionReq ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof CreateOfferBasedOnProductActionReq> {
		return new CreateOfferBasedOnProductActionReq(this.toJSON());
	}
}
export abstract class CreateOfferBasedOnProductActionReqFactory {
	abstract create(data: unknown): CreateOfferBasedOnProductActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for createOfferBasedOnProductActionReq
  **/
	export type CreateOfferBasedOnProductActionReqType =  {
			/**
  * Offer title
  * @type {string}
  **/
 name : string;
			/**
  * Offer language code (e.g., pl-PL)
  * @type {string}
  **/
 language : string;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.CategoryType}
  **/
 category : CreateOfferBasedOnProductActionReqType.CategoryType;
			/**
  * Product details and associated quantities
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType[]}
  **/
 productSet : CreateOfferBasedOnProductActionReqType.ProductSetType[];
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.StockType}
  **/
 stock : CreateOfferBasedOnProductActionReqType.StockType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.SellingModeType}
  **/
 sellingMode : CreateOfferBasedOnProductActionReqType.SellingModeType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.PaymentsType}
  **/
 payments : CreateOfferBasedOnProductActionReqType.PaymentsType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.DeliveryType}
  **/
 delivery : CreateOfferBasedOnProductActionReqType.DeliveryType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.PublicationType}
  **/
 publication : CreateOfferBasedOnProductActionReqType.PublicationType;
			/**
  * 
  * @type {{[key: string]: any}}
  **/
 additionalMarketplaces ?: {[key: string]: any};
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.CompatibilityListType}
  **/
 compatibilityList : CreateOfferBasedOnProductActionReqType.CompatibilityListType;
			/**
  * 
  * @type {string[]}
  **/
 images : string[];
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.DescriptionType}
  **/
 description : CreateOfferBasedOnProductActionReqType.DescriptionType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.B2bType}
  **/
 b2b : CreateOfferBasedOnProductActionReqType.B2bType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.AttachmentsType[]}
  **/
 attachments : CreateOfferBasedOnProductActionReqType.AttachmentsType[];
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.FundraisingCampaignType}
  **/
 fundraisingCampaign : CreateOfferBasedOnProductActionReqType.FundraisingCampaignType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.AdditionalServicesType}
  **/
 additionalServices : CreateOfferBasedOnProductActionReqType.AdditionalServicesType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.AfterSalesServicesType}
  **/
 afterSalesServices : CreateOfferBasedOnProductActionReqType.AfterSalesServicesType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.SizeTableType}
  **/
 sizeTable : CreateOfferBasedOnProductActionReqType.SizeTableType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ContactType}
  **/
 contact : CreateOfferBasedOnProductActionReqType.ContactType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.DiscountsType}
  **/
 discounts : CreateOfferBasedOnProductActionReqType.DiscountsType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.LocationType}
  **/
 location : CreateOfferBasedOnProductActionReqType.LocationType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ExternalType}
  **/
 external : CreateOfferBasedOnProductActionReqType.ExternalType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.TaxSettingsType}
  **/
 taxSettings : CreateOfferBasedOnProductActionReqType.TaxSettingsType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.MessageToSellerSettingsType}
  **/
 messageToSellerSettings : CreateOfferBasedOnProductActionReqType.MessageToSellerSettingsType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CreateOfferBasedOnProductActionReqType {
	/**
  * The base type definition for categoryType
  **/
	export type CategoryType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CategoryType {
}
	/**
  * The base type definition for productSetType
  **/
	export type ProductSetType =  {
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType}
  **/
 product : CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType.QuantityType}
  **/
 quantity : CreateOfferBasedOnProductActionReqType.ProductSetType.QuantityType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsiblePersonType}
  **/
 responsiblePerson : CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsiblePersonType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsibleProducerType}
  **/
 responsibleProducer : CreateOfferBasedOnProductActionReqType.ProductSetType.ResponsibleProducerType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType.SafetyInformationType}
  **/
 safetyInformation : CreateOfferBasedOnProductActionReqType.ProductSetType.SafetyInformationType;
			/**
  * 
  * @type {boolean}
  **/
 marketedBeforeGPSRObligation : boolean;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType.DepositsType[]}
  **/
 deposits : CreateOfferBasedOnProductActionReqType.ProductSetType.DepositsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ProductSetType {
	/**
  * The base type definition for productType
  **/
	export type ProductType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 idType : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.CategoryType}
  **/
 category : CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.CategoryType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType[]}
  **/
 parameters : CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType[];
			/**
  * 
  * @type {string[]}
  **/
 images : string[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ProductType {
	/**
  * The base type definition for categoryType
  **/
	export type CategoryType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CategoryType {
}
	/**
  * The base type definition for parametersType
  **/
	export type ParametersType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType}
  **/
 rangeValue : CreateOfferBasedOnProductActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType;
			/**
  * 
  * @type {string[]}
  **/
 values : string[];
			/**
  * 
  * @type {string[]}
  **/
 valuesIds : string[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ParametersType {
	/**
  * The base type definition for rangeValueType
  **/
	export type RangeValueType =  {
			/**
  * 
  * @type {string}
  **/
 from : string;
			/**
  * 
  * @type {string}
  **/
 to : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace RangeValueType {
}
}
}
	/**
  * The base type definition for quantityType
  **/
	export type QuantityType =  {
			/**
  * 
  * @type {number}
  **/
 value : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace QuantityType {
}
	/**
  * The base type definition for responsiblePersonType
  **/
	export type ResponsiblePersonType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ResponsiblePersonType {
}
	/**
  * The base type definition for responsibleProducerType
  **/
	export type ResponsibleProducerType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 type : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ResponsibleProducerType {
}
	/**
  * The base type definition for safetyInformationType
  **/
	export type SafetyInformationType =  {
			/**
  * 
  * @type {string}
  **/
 type : string;
			/**
  * 
  * @type {string}
  **/
 description : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SafetyInformationType {
}
	/**
  * The base type definition for depositsType
  **/
	export type DepositsType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {number}
  **/
 quantity : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DepositsType {
}
}
	/**
  * The base type definition for stockType
  **/
	export type StockType =  {
			/**
  * 
  * @type {number}
  **/
 available : number;
			/**
  * 
  * @type {string}
  **/
 unit : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace StockType {
}
	/**
  * The base type definition for sellingModeType
  **/
	export type SellingModeType =  {
			/**
  * 
  * @type {string}
  **/
 format : string;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.SellingModeType.PriceType}
  **/
 price : CreateOfferBasedOnProductActionReqType.SellingModeType.PriceType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.SellingModeType.MinimalPriceType}
  **/
 minimalPrice : CreateOfferBasedOnProductActionReqType.SellingModeType.MinimalPriceType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.SellingModeType.StartingPriceType}
  **/
 startingPrice : CreateOfferBasedOnProductActionReqType.SellingModeType.StartingPriceType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SellingModeType {
	/**
  * The base type definition for priceType
  **/
	export type PriceType =  {
			/**
  * 
  * @type {string}
  **/
 amount : string;
			/**
  * 
  * @type {string}
  **/
 currency : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PriceType {
}
	/**
  * The base type definition for minimalPriceType
  **/
	export type MinimalPriceType =  {
			/**
  * 
  * @type {string}
  **/
 amount : string;
			/**
  * 
  * @type {string}
  **/
 currency : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MinimalPriceType {
}
	/**
  * The base type definition for startingPriceType
  **/
	export type StartingPriceType =  {
			/**
  * 
  * @type {string}
  **/
 amount : string;
			/**
  * 
  * @type {string}
  **/
 currency : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace StartingPriceType {
}
}
	/**
  * The base type definition for paymentsType
  **/
	export type PaymentsType =  {
			/**
  * 
  * @type {string}
  **/
 invoice : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PaymentsType {
}
	/**
  * The base type definition for deliveryType
  **/
	export type DeliveryType =  {
			/**
  * 
  * @type {string}
  **/
 handlingTime : string;
			/**
  * 
  * @type {string}
  **/
 additionalInfo : string;
			/**
  * 
  * @type {string}
  **/
 shipmentDate : string;
			/**
  * Optional; may be null
  * @type {CreateOfferBasedOnProductActionReqType.DeliveryType.ShippingRatesType}
  **/
 shippingRates : CreateOfferBasedOnProductActionReqType.DeliveryType.ShippingRatesType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DeliveryType {
	/**
  * The base type definition for shippingRatesType
  **/
	export type ShippingRatesType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ShippingRatesType {
}
}
	/**
  * The base type definition for publicationType
  **/
	export type PublicationType =  {
			/**
  * 
  * @type {string}
  **/
 duration : string;
			/**
  * 
  * @type {string}
  **/
 startingAt : string;
			/**
  * 
  * @type {string}
  **/
 endingAt : string;
			/**
  * 
  * @type {string}
  **/
 status : string;
			/**
  * 
  * @type {boolean}
  **/
 republish : boolean;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PublicationType {
}
	/**
  * The base type definition for compatibilityListType
  **/
	export type CompatibilityListType =  {
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.CompatibilityListType.ItemsType[]}
  **/
 items : CreateOfferBasedOnProductActionReqType.CompatibilityListType.ItemsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CompatibilityListType {
	/**
  * The base type definition for itemsType
  **/
	export type ItemsType =  {
			/**
  * 
  * @type {string}
  **/
 type : string;
			/**
  * 
  * @type {string}
  **/
 text : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ItemsType {
}
}
	/**
  * The base type definition for descriptionType
  **/
	export type DescriptionType =  {
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType[]}
  **/
 sections : CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DescriptionType {
	/**
  * The base type definition for sectionsType
  **/
	export type SectionsType =  {
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType.ItemsType[]}
  **/
 items : CreateOfferBasedOnProductActionReqType.DescriptionType.SectionsType.ItemsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SectionsType {
	/**
  * The base type definition for itemsType
  **/
	export type ItemsType =  {
			/**
  * 
  * @type {string}
  **/
 type : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ItemsType {
}
}
}
	/**
  * The base type definition for b2bType
  **/
	export type B2bType =  {
			/**
  * 
  * @type {boolean}
  **/
 buyableOnlyByBusiness : boolean;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace B2bType {
}
	/**
  * The base type definition for attachmentsType
  **/
	export type AttachmentsType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AttachmentsType {
}
	/**
  * The base type definition for fundraisingCampaignType
  **/
	export type FundraisingCampaignType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace FundraisingCampaignType {
}
	/**
  * The base type definition for additionalServicesType
  **/
	export type AdditionalServicesType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdditionalServicesType {
}
	/**
  * The base type definition for afterSalesServicesType
  **/
	export type AfterSalesServicesType =  {
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ImpliedWarrantyType}
  **/
 impliedWarranty : CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ImpliedWarrantyType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ReturnPolicyType}
  **/
 returnPolicy : CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.ReturnPolicyType;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.WarrantyType}
  **/
 warranty : CreateOfferBasedOnProductActionReqType.AfterSalesServicesType.WarrantyType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AfterSalesServicesType {
	/**
  * The base type definition for impliedWarrantyType
  **/
	export type ImpliedWarrantyType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ImpliedWarrantyType {
}
	/**
  * The base type definition for returnPolicyType
  **/
	export type ReturnPolicyType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ReturnPolicyType {
}
	/**
  * The base type definition for warrantyType
  **/
	export type WarrantyType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WarrantyType {
}
}
	/**
  * The base type definition for sizeTableType
  **/
	export type SizeTableType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SizeTableType {
}
	/**
  * The base type definition for contactType
  **/
	export type ContactType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ContactType {
}
	/**
  * The base type definition for discountsType
  **/
	export type DiscountsType =  {
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.DiscountsType.WholesalePriceListType}
  **/
 wholesalePriceList : CreateOfferBasedOnProductActionReqType.DiscountsType.WholesalePriceListType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DiscountsType {
	/**
  * The base type definition for wholesalePriceListType
  **/
	export type WholesalePriceListType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
			/**
  * 
  * @type {string}
  **/
 name : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WholesalePriceListType {
}
}
	/**
  * The base type definition for locationType
  **/
	export type LocationType =  {
			/**
  * 
  * @type {string}
  **/
 city : string;
			/**
  * 
  * @type {string}
  **/
 countryCode : string;
			/**
  * 
  * @type {string}
  **/
 postCode : string;
			/**
  * 
  * @type {string}
  **/
 province : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace LocationType {
}
	/**
  * The base type definition for externalType
  **/
	export type ExternalType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ExternalType {
}
	/**
  * The base type definition for taxSettingsType
  **/
	export type TaxSettingsType =  {
			/**
  * 
  * @type {string}
  **/
 subject : string;
			/**
  * 
  * @type {string}
  **/
 exemption : string;
			/**
  * 
  * @type {CreateOfferBasedOnProductActionReqType.TaxSettingsType.RatesType[]}
  **/
 rates : CreateOfferBasedOnProductActionReqType.TaxSettingsType.RatesType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace TaxSettingsType {
	/**
  * The base type definition for ratesType
  **/
	export type RatesType =  {
			/**
  * 
  * @type {string}
  **/
 rate : string;
			/**
  * 
  * @type {string}
  **/
 countryCode : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace RatesType {
}
}
	/**
  * The base type definition for messageToSellerSettingsType
  **/
	export type MessageToSellerSettingsType =  {
			/**
  * 
  * @type {string}
  **/
 mode : string;
			/**
  * 
  * @type {string}
  **/
 hint : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MessageToSellerSettingsType {
}
}