import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get all data of the particular product-offer
*/
export type GetAllDataOfTheParticularProductOfferActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * GetAllDataOfTheParticularProductOfferAction
 */
export class GetAllDataOfTheParticularProductOfferAction { //
  static URL = 'https://api.{environment}/sale/product-offers/{offerId}';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		GetAllDataOfTheParticularProductOfferAction.URL,
		 undefined,
		qs
	);
  static Method = 'get';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<unknown, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<GetAllDataOfTheParticularProductOfferActionRes, unknown, unknown>(
			overrideUrl ?? GetAllDataOfTheParticularProductOfferAction.NewUrl(
				qs
			),
			{
				method: GetAllDataOfTheParticularProductOfferAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init?: TypedRequestInit<unknown, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => GetAllDataOfTheParticularProductOfferActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new GetAllDataOfTheParticularProductOfferActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetAllDataOfTheParticularProductOfferActionRes(item))
		const res = await GetAllDataOfTheParticularProductOfferAction.Fetch$(
			qs,
			ctx,
			init,
			overrideUrl,
			);
			return handleFetchResponse(
				res, 
				(item) => creatorFn(item),
				onMessage,
				init?.signal,
			);
	}
  static Definition = {
  "name": "Get all data of the particular product-offer",
  "url": "https://api.{environment}/sale/product-offers/{offerId}",
  "method": "get",
  "description": "Full response model returned by GET /sale/product-offers/{offerId}",
  "out": {
    "fields": [
      {
        "name": "id",
        "description": "Unique offer identifier",
        "type": "string"
      },
      {
        "name": "name",
        "description": "Offer title",
        "type": "string"
      },
      {
        "name": "language",
        "description": "Offer language code (e.g. pl-PL)",
        "type": "string"
      },
      {
        "name": "createdAt",
        "description": "Offer creation timestamp (ISO8601)",
        "type": "string"
      },
      {
        "name": "updatedAt",
        "description": "Offer last update timestamp (ISO8601)",
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
        "name": "contact",
        "type": "object",
        "fields": [
          {
            "name": "id",
            "type": "string"
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
            "name": "endedBy",
            "type": "string"
          },
          {
            "name": "status",
            "type": "string"
          },
          {
            "name": "republish",
            "type": "bool"
          },
          {
            "name": "marketplaces",
            "type": "object",
            "fields": [
              {
                "name": "base",
                "type": "object",
                "fields": [
                  {
                    "name": "id",
                    "type": "string"
                  }
                ]
              },
              {
                "name": "additional",
                "type": "array",
                "fields": [
                  {
                    "name": "id",
                    "type": "string"
                  }
                ]
              }
            ]
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
              }
            ]
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
              }
            ]
          }
        ]
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
        "name": "images",
        "type": "slice",
        "primitive": "string"
      },
      {
        "name": "productSet",
        "type": "array",
        "fields": [
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
            "name": "product",
            "type": "object",
            "fields": [
              {
                "name": "id",
                "type": "string"
              },
              {
                "name": "isAiCoCreated",
                "type": "bool"
              },
              {
                "name": "publication",
                "type": "object",
                "fields": [
                  {
                    "name": "status",
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
                    "type": "array",
                    "primitive": "string"
                  },
                  {
                    "name": "valuesIds",
                    "type": "array",
                    "primitive": "string"
                  }
                ]
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
          }
        ]
      },
      {
        "name": "additionalMarketplaces",
        "type": "map?",
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
          },
          {
            "name": "publication",
            "type": "object",
            "fields": [
              {
                "name": "state",
                "type": "string"
              },
              {
                "name": "refusalReasons",
                "type": "array",
                "fields": [
                  {
                    "name": "code",
                    "type": "string"
                  },
                  {
                    "name": "userMessage",
                    "type": "string"
                  },
                  {
                    "name": "parameters",
                    "type": "object",
                    "fields": [
                      {
                        "name": "maxAllowedPriceDecreasePercent",
                        "type": "array",
                        "primitive": "string"
                      }
                    ]
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
        "name": "compatibilityList",
        "type": "object",
        "fields": [
          {
            "name": "type",
            "type": "string"
          }
        ]
      },
      {
        "name": "validation",
        "type": "object",
        "fields": [
          {
            "name": "validatedAt",
            "type": "string"
          },
          {
            "name": "errors",
            "type": "array",
            "fields": [
              {
                "name": "code",
                "type": "string"
              },
              {
                "name": "details",
                "type": "string"
              },
              {
                "name": "message",
                "type": "string"
              },
              {
                "name": "path",
                "type": "string"
              },
              {
                "name": "userMessage",
                "type": "string"
              },
              {
                "name": "metadata",
                "type": "object",
                "fields": [
                  {
                    "name": "productId",
                    "type": "string"
                  }
                ]
              }
            ]
          },
          {
            "name": "warnings",
            "type": "array",
            "fields": [
              {
                "name": "code",
                "type": "string"
              },
              {
                "name": "details",
                "type": "string"
              },
              {
                "name": "message",
                "type": "string"
              },
              {
                "name": "path",
                "type": "string"
              },
              {
                "name": "userMessage",
                "type": "string"
              },
              {
                "name": "metadata",
                "type": "object",
                "fields": [
                  {
                    "name": "productId",
                    "type": "string"
                  }
                ]
              }
            ]
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
        "name": "sizeTable",
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
  * The base class definition for getAllDataOfTheParticularProductOfferActionRes
  **/
export class GetAllDataOfTheParticularProductOfferActionRes {
		/**
  * Unique offer identifier
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * Unique offer identifier
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * Unique offer identifier
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
  * Offer language code (e.g. pl-PL)
  * @type {string}
  **/
 #language : string  =  ""
		/**
  * Offer language code (e.g. pl-PL)
  * @returns {string}
  **/
get language () { return this.#language }
/**
  * Offer language code (e.g. pl-PL)
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
  * Offer creation timestamp (ISO8601)
  * @type {string}
  **/
 #createdAt : string  =  ""
		/**
  * Offer creation timestamp (ISO8601)
  * @returns {string}
  **/
get createdAt () { return this.#createdAt }
/**
  * Offer creation timestamp (ISO8601)
  * @type {string}
  **/
set createdAt (value: string) {
		this.#createdAt = String(value);
}
setCreatedAt (value: string) {
	this.createdAt = value
	return this
}
		/**
  * Offer last update timestamp (ISO8601)
  * @type {string}
  **/
 #updatedAt : string  =  ""
		/**
  * Offer last update timestamp (ISO8601)
  * @returns {string}
  **/
get updatedAt () { return this.#updatedAt }
/**
  * Offer last update timestamp (ISO8601)
  * @type {string}
  **/
set updatedAt (value: string) {
		this.#updatedAt = String(value);
}
setUpdatedAt (value: string) {
	this.updatedAt = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Category}
  **/
 #category ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Category>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Category}
  **/
set category (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Category>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Category) {
			this.#category = value
		} else {
			this.#category = new GetAllDataOfTheParticularProductOfferActionRes.Category(value)
		}
}
setCategory (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Category>) {
	this.category = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Stock}
  **/
 #stock ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Stock>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Stock}
  **/
set stock (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Stock>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Stock) {
			this.#stock = value
		} else {
			this.#stock = new GetAllDataOfTheParticularProductOfferActionRes.Stock(value)
		}
}
setStock (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Stock>) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Contact}
  **/
 #contact ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Contact>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Contact}
  **/
get contact () { return this.#contact }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Contact}
  **/
set contact (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Contact>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Contact) {
			this.#contact = value
		} else {
			this.#contact = new GetAllDataOfTheParticularProductOfferActionRes.Contact(value)
		}
}
setContact (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Contact>) {
	this.contact = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication}
  **/
 #publication ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication}
  **/
set publication (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication) {
			this.#publication = value
		} else {
			this.#publication = new GetAllDataOfTheParticularProductOfferActionRes.Publication(value)
		}
}
setPublication (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication>) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
 #sellingMode ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
set sellingMode (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode(value)
		}
}
setSellingMode (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode>) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Payments}
  **/
 #payments ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Payments>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Payments}
  **/
get payments () { return this.#payments }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Payments}
  **/
set payments (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Payments>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Payments) {
			this.#payments = value
		} else {
			this.#payments = new GetAllDataOfTheParticularProductOfferActionRes.Payments(value)
		}
}
setPayments (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Payments>) {
	this.payments = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Delivery}
  **/
 #delivery ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Delivery}
  **/
get delivery () { return this.#delivery }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Delivery}
  **/
set delivery (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Delivery) {
			this.#delivery = value
		} else {
			this.#delivery = new GetAllDataOfTheParticularProductOfferActionRes.Delivery(value)
		}
}
setDelivery (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery>) {
	this.delivery = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices}
  **/
 #afterSalesServices ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices}
  **/
get afterSalesServices () { return this.#afterSalesServices }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices}
  **/
set afterSalesServices (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices) {
			this.#afterSalesServices = value
		} else {
			this.#afterSalesServices = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices(value)
		}
}
setAfterSalesServices (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices>) {
	this.afterSalesServices = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Discounts}
  **/
 #discounts ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Discounts}
  **/
get discounts () { return this.#discounts }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Discounts}
  **/
set discounts (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Discounts) {
			this.#discounts = value
		} else {
			this.#discounts = new GetAllDataOfTheParticularProductOfferActionRes.Discounts(value)
		}
}
setDiscounts (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts>) {
	this.discounts = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description}
  **/
 #description ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description}
  **/
set description (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Description) {
			this.#description = value
		} else {
			this.#description = new GetAllDataOfTheParticularProductOfferActionRes.Description(value)
		}
}
setDescription (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description>) {
	this.description = value
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
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet}
  **/
 #productSet : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet}
  **/
get productSet () { return this.#productSet }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet}
  **/
set productSet (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet) {
			this.#productSet = value
		} else {
			this.#productSet = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.ProductSet(item))
		}
}
setProductSet (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet>[]) {
	this.productSet = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Attachments}
  **/
 #attachments : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Attachments>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Attachments}
  **/
get attachments () { return this.#attachments }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Attachments}
  **/
set attachments (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Attachments>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.Attachments) {
			this.#attachments = value
		} else {
			this.#attachments = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.Attachments(item))
		}
}
setAttachments (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Attachments>[]) {
	this.attachments = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign}
  **/
 #fundraisingCampaign ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign}
  **/
get fundraisingCampaign () { return this.#fundraisingCampaign }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign}
  **/
set fundraisingCampaign (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign) {
			this.#fundraisingCampaign = value
		} else {
			this.#fundraisingCampaign = new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign(value)
		}
}
setFundraisingCampaign (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign>) {
	this.fundraisingCampaign = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices}
  **/
 #additionalServices ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices}
  **/
get additionalServices () { return this.#additionalServices }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices}
  **/
set additionalServices (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices) {
			this.#additionalServices = value
		} else {
			this.#additionalServices = new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices(value)
		}
}
setAdditionalServices (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices>) {
	this.additionalServices = value
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
  * @type {GetAllDataOfTheParticularProductOfferActionRes.B2b}
  **/
 #b2b ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.B2b>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.B2b}
  **/
get b2b () { return this.#b2b }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.B2b}
  **/
set b2b (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.B2b>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.B2b) {
			this.#b2b = value
		} else {
			this.#b2b = new GetAllDataOfTheParticularProductOfferActionRes.B2b(value)
		}
}
setB2b (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.B2b>) {
	this.b2b = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList}
  **/
 #compatibilityList ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList}
  **/
get compatibilityList () { return this.#compatibilityList }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList}
  **/
set compatibilityList (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList) {
			this.#compatibilityList = value
		} else {
			this.#compatibilityList = new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList(value)
		}
}
setCompatibilityList (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList>) {
	this.compatibilityList = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation}
  **/
 #validation ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation}
  **/
get validation () { return this.#validation }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation}
  **/
set validation (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation) {
			this.#validation = value
		} else {
			this.#validation = new GetAllDataOfTheParticularProductOfferActionRes.Validation(value)
		}
}
setValidation (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation>) {
	this.validation = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.External}
  **/
 #external ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.External>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.External}
  **/
get external () { return this.#external }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.External}
  **/
set external (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.External>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.External) {
			this.#external = value
		} else {
			this.#external = new GetAllDataOfTheParticularProductOfferActionRes.External(value)
		}
}
setExternal (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.External>) {
	this.external = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SizeTable}
  **/
 #sizeTable ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SizeTable>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SizeTable}
  **/
get sizeTable () { return this.#sizeTable }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SizeTable}
  **/
set sizeTable (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SizeTable>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SizeTable) {
			this.#sizeTable = value
		} else {
			this.#sizeTable = new GetAllDataOfTheParticularProductOfferActionRes.SizeTable(value)
		}
}
setSizeTable (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SizeTable>) {
	this.sizeTable = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings}
  **/
 #taxSettings ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings}
  **/
get taxSettings () { return this.#taxSettings }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings}
  **/
set taxSettings (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings) {
			this.#taxSettings = value
		} else {
			this.#taxSettings = new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings(value)
		}
}
setTaxSettings (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings>) {
	this.taxSettings = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings}
  **/
 #messageToSellerSettings ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings}
  **/
get messageToSellerSettings () { return this.#messageToSellerSettings }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings}
  **/
set messageToSellerSettings (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings) {
			this.#messageToSellerSettings = value
		} else {
			this.#messageToSellerSettings = new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings(value)
		}
}
setMessageToSellerSettings (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings>) {
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.CategoryType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.CategoryType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Category(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.CategoryType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Category> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Category ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Category> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Category(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Stock, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.StockType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.StockType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Stock(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.StockType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Stock> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Stock ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Stock> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Stock(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Contact, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ContactType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Contact(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Contact, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ContactType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Contact(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ContactType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Contact> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Contact ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Contact> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Contact(this.toJSON());
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
 #endedBy : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get endedBy () { return this.#endedBy }
/**
  * 
  * @type {string}
  **/
set endedBy (value: string) {
		this.#endedBy = String(value);
}
setEndedBy (value: string) {
	this.endedBy = value
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
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces}
  **/
 #marketplaces ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces}
  **/
get marketplaces () { return this.#marketplaces }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces}
  **/
set marketplaces (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces) {
			this.#marketplaces = value
		} else {
			this.#marketplaces = new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces(value)
		}
}
setMarketplaces (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces>) {
	this.marketplaces = value
	return this
}
/**
  * The base class definition for marketplaces
  **/
static Marketplaces = class Marketplaces {
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base}
  **/
 #base ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base}
  **/
get base () { return this.#base }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base}
  **/
set base (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base) {
			this.#base = value
		} else {
			this.#base = new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base(value)
		}
}
setBase (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base>) {
	this.base = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional}
  **/
 #additional : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional}
  **/
get additional () { return this.#additional }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional}
  **/
set additional (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional) {
			this.#additional = value
		} else {
			this.#additional = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional(item))
		}
}
setAdditional (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional>[]) {
	this.additional = value
	return this
}
/**
  * The base class definition for base
  **/
static Base = class Base {
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
		const d = data as Partial<Base>;
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.BaseType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.BaseType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.BaseType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base(this.toJSON());
	}
}
/**
  * The base class definition for additional
  **/
static Additional = class Additional {
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
		const d = data as Partial<Additional>;
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.AdditionalType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.AdditionalType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.AdditionalType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional(this.toJSON());
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
		const d = data as Partial<Marketplaces>;
			if (d.base !== undefined) { this.base = d.base }
			if (d.additional !== undefined) { this.additional = d.additional }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Marketplaces>;
			if (!(d.base instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base)) { this.base = new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base(d.base || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				base: this.#base,
				additional: this.#additional,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			base$: 'base',
get base() {
					return withPrefix(
						"publication.marketplaces.base",
						GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base.Fields
						);
						},
			additional$: 'additional',
get additional() {
					return withPrefix(
						"publication.marketplaces.additional[:i]",
						GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces(this.toJSON());
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
		const d = data as Partial<Publication>;
			if (d.duration !== undefined) { this.duration = d.duration }
			if (d.startingAt !== undefined) { this.startingAt = d.startingAt }
			if (d.endingAt !== undefined) { this.endingAt = d.endingAt }
			if (d.endedBy !== undefined) { this.endedBy = d.endedBy }
			if (d.status !== undefined) { this.status = d.status }
			if (d.republish !== undefined) { this.republish = d.republish }
			if (d.marketplaces !== undefined) { this.marketplaces = d.marketplaces }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Publication>;
			if (!(d.marketplaces instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces)) { this.marketplaces = new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces(d.marketplaces || {}) }	
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
				endedBy: this.#endedBy,
				status: this.#status,
				republish: this.#republish,
				marketplaces: this.#marketplaces,
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
			endedBy: 'endedBy',
			status: 'status',
			republish: 'republish',
			marketplaces$: 'marketplaces',
get marketplaces() {
					return withPrefix(
						"publication.marketplaces",
						GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.PublicationType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PublicationType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PublicationType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Publication> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication(this.toJSON());
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
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
 #price ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
set price (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price(value)
		}
}
setPrice (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price>) {
	this.price = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice}
  **/
 #minimalPrice ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice}
  **/
get minimalPrice () { return this.#minimalPrice }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice}
  **/
set minimalPrice (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice) {
			this.#minimalPrice = value
		} else {
			this.#minimalPrice = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice(value)
		}
}
setMinimalPrice (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice>) {
	this.minimalPrice = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice}
  **/
 #startingPrice ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice}
  **/
get startingPrice () { return this.#startingPrice }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice}
  **/
set startingPrice (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice) {
			this.#startingPrice = value
		} else {
			this.#startingPrice = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice(value)
		}
}
setStartingPrice (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice>) {
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.MinimalPriceType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.MinimalPriceType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.MinimalPriceType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.StartingPriceType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.StartingPriceType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.StartingPriceType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice(this.toJSON());
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
			if (!(d.price instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price)) { this.price = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price(d.price || {}) }	
			if (!(d.minimalPrice instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice)) { this.minimalPrice = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice(d.minimalPrice || {}) }	
			if (!(d.startingPrice instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice)) { this.startingPrice = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice(d.startingPrice || {}) }	
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
						GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price.Fields
						);
						},
			minimalPrice$: 'minimalPrice',
get minimalPrice() {
					return withPrefix(
						"sellingMode.minimalPrice",
						GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice.Fields
						);
						},
			startingPrice$: 'startingPrice',
get startingPrice() {
					return withPrefix(
						"sellingMode.startingPrice",
						GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.SellingModeType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SellingModeType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SellingModeType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SellingMode> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Payments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.PaymentsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Payments(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Payments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PaymentsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Payments(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.PaymentsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Payments> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Payments ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Payments> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Payments(this.toJSON());
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
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates}
  **/
 #shippingRates ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates}
  **/
get shippingRates () { return this.#shippingRates }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates}
  **/
set shippingRates (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates) {
			this.#shippingRates = value
		} else {
			this.#shippingRates = new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates(value)
		}
}
setShippingRates (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates>) {
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.DeliveryType.ShippingRatesType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DeliveryType.ShippingRatesType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DeliveryType.ShippingRatesType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates(this.toJSON());
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
			if (!(d.shippingRates instanceof GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates)) { this.shippingRates = new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates(d.shippingRates || {}) }	
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
						GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Delivery, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.DeliveryType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Delivery, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DeliveryType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DeliveryType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Delivery> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery(this.toJSON());
	}
}
/**
  * The base class definition for afterSalesServices
  **/
static AfterSalesServices = class AfterSalesServices {
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
 #impliedWarranty ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
get impliedWarranty () { return this.#impliedWarranty }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
set impliedWarranty (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty) {
			this.#impliedWarranty = value
		} else {
			this.#impliedWarranty = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty(value)
		}
}
setImpliedWarranty (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty>) {
	this.impliedWarranty = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
 #returnPolicy ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
get returnPolicy () { return this.#returnPolicy }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
set returnPolicy (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy) {
			this.#returnPolicy = value
		} else {
			this.#returnPolicy = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy(value)
		}
}
setReturnPolicy (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy>) {
	this.returnPolicy = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty}
  **/
 #warranty ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty}
  **/
get warranty () { return this.#warranty }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty}
  **/
set warranty (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty) {
			this.#warranty = value
		} else {
			this.#warranty = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty(value)
		}
}
setWarranty (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty>) {
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ReturnPolicyType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ReturnPolicyType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ReturnPolicyType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.WarrantyType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.WarrantyType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.WarrantyType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty(this.toJSON());
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
			if (!(d.impliedWarranty instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty)) { this.impliedWarranty = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty(d.impliedWarranty || {}) }	
			if (!(d.returnPolicy instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy)) { this.returnPolicy = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy(d.returnPolicy || {}) }	
			if (!(d.warranty instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty)) { this.warranty = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty(d.warranty || {}) }	
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
						GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty.Fields
						);
						},
			returnPolicy$: 'returnPolicy',
get returnPolicy() {
					return withPrefix(
						"afterSalesServices.returnPolicy",
						GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy.Fields
						);
						},
			warranty$: 'warranty',
get warranty() {
					return withPrefix(
						"afterSalesServices.warranty",
						GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices(this.toJSON());
	}
}
/**
  * The base class definition for discounts
  **/
static Discounts = class Discounts {
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList}
  **/
 #wholesalePriceList ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList}
  **/
get wholesalePriceList () { return this.#wholesalePriceList }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList}
  **/
set wholesalePriceList (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList) {
			this.#wholesalePriceList = value
		} else {
			this.#wholesalePriceList = new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList(value)
		}
}
setWholesalePriceList (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList>) {
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.DiscountsType.WholesalePriceListType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DiscountsType.WholesalePriceListType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DiscountsType.WholesalePriceListType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList(this.toJSON());
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
			if (!(d.wholesalePriceList instanceof GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList)) { this.wholesalePriceList = new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList(d.wholesalePriceList || {}) }	
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
						GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Discounts, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.DiscountsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Discounts, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DiscountsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DiscountsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Discounts> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts(this.toJSON());
	}
}
/**
  * The base class definition for description
  **/
static Description = class Description {
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections}
  **/
 #sections : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections}
  **/
set sections (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections) {
			this.#sections = value
		} else {
			this.#sections = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections(item))
		}
}
setSections (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections>[]) {
	this.sections = value
	return this
}
/**
  * The base class definition for sections
  **/
static Sections = class Sections {
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items}
  **/
 #items : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items}
  **/
set items (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items) {
			this.#items = value
		} else {
			this.#items = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items(item))
		}
}
setItems (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items>[]) {
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType.ItemsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType.ItemsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType.ItemsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items(this.toJSON());
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
						GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Description.Sections, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Description.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description.Sections> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections(this.toJSON());
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
						GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Description, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.DescriptionType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DescriptionType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.DescriptionType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Description> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description(this.toJSON());
	}
}
/**
  * The base class definition for productSet
  **/
static ProductSet = class ProductSet {
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity}
  **/
 #quantity ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity}
  **/
get quantity () { return this.#quantity }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity}
  **/
set quantity (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity) {
			this.#quantity = value
		} else {
			this.#quantity = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity(value)
		}
}
setQuantity (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity>) {
	this.quantity = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product}
  **/
 #product ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product}
  **/
get product () { return this.#product }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product}
  **/
set product (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product) {
			this.#product = value
		} else {
			this.#product = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product(value)
		}
}
setProduct (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product>) {
	this.product = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson}
  **/
 #responsiblePerson ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson}
  **/
get responsiblePerson () { return this.#responsiblePerson }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson}
  **/
set responsiblePerson (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson) {
			this.#responsiblePerson = value
		} else {
			this.#responsiblePerson = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson(value)
		}
}
setResponsiblePerson (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson>) {
	this.responsiblePerson = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer}
  **/
 #responsibleProducer ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer}
  **/
get responsibleProducer () { return this.#responsibleProducer }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer}
  **/
set responsibleProducer (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer) {
			this.#responsibleProducer = value
		} else {
			this.#responsibleProducer = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer(value)
		}
}
setResponsibleProducer (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer>) {
	this.responsibleProducer = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation}
  **/
 #safetyInformation ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation}
  **/
set safetyInformation (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation(value)
		}
}
setSafetyInformation (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation>) {
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
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits}
  **/
 #deposits : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits}
  **/
get deposits () { return this.#deposits }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits}
  **/
set deposits (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits) {
			this.#deposits = value
		} else {
			this.#deposits = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits(item))
		}
}
setDeposits (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits>[]) {
	this.deposits = value
	return this
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.QuantityType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.QuantityType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.QuantityType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity(this.toJSON());
	}
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
  * @type {boolean}
  **/
 #isAiCoCreated ! : boolean
		/**
  * 
  * @returns {boolean}
  **/
get isAiCoCreated () { return this.#isAiCoCreated }
/**
  * 
  * @type {boolean}
  **/
set isAiCoCreated (value: boolean) {
		this.#isAiCoCreated = Boolean(value);
}
setIsAiCoCreated (value: boolean) {
	this.isAiCoCreated = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication}
  **/
 #publication ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication}
  **/
set publication (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication) {
			this.#publication = value
		} else {
			this.#publication = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication(value)
		}
}
setPublication (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication>) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters}
  **/
 #parameters : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters}
  **/
set parameters (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters) {
			this.#parameters = value
		} else {
			this.#parameters = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters(item))
		}
}
setParameters (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters>[]) {
	this.parameters = value
	return this
}
/**
  * The base class definition for publication
  **/
static Publication = class Publication {
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
			if (d.status !== undefined) { this.status = d.status }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				status: this.#status,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			status: 'status',
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.PublicationType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.PublicationType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.PublicationType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication(this.toJSON());
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
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
 #rangeValue ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
get rangeValue () { return this.#rangeValue }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
set rangeValue (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue) {
			this.#rangeValue = value
		} else {
			this.#rangeValue = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue(value)
		}
}
setRangeValue (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue>) {
	this.rangeValue = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values}
  **/
 #values : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values}
  **/
get values () { return this.#values }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values}
  **/
set values (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values) {
			this.#values = value
		} else {
			this.#values = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values(item))
		}
}
setValues (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values>[]) {
	this.values = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds}
  **/
 #valuesIds : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds}
  **/
get valuesIds () { return this.#valuesIds }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds}
  **/
set valuesIds (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds) {
			this.#valuesIds = value
		} else {
			this.#valuesIds = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds(item))
		}
}
setValuesIds (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds>[]) {
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue(this.toJSON());
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
			if (!(d.rangeValue instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue)) { this.rangeValue = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue(d.rangeValue || {}) }	
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
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue.Fields
						);
						},
			values$: 'values',
get values() {
					return withPrefix(
						"productSet.product.parameters.values[:i]",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values.Fields
						);
						},
			valuesIds$: 'valuesIds',
get valuesIds() {
					return withPrefix(
						"productSet.product.parameters.valuesIds[:i]",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters(this.toJSON());
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
			if (d.isAiCoCreated !== undefined) { this.isAiCoCreated = d.isAiCoCreated }
			if (d.publication !== undefined) { this.publication = d.publication }
			if (d.parameters !== undefined) { this.parameters = d.parameters }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Product>;
			if (!(d.publication instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication)) { this.publication = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication(d.publication || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				isAiCoCreated: this.#isAiCoCreated,
				publication: this.#publication,
				parameters: this.#parameters,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			isAiCoCreated: 'isAiCoCreated',
			publication$: 'publication',
get publication() {
					return withPrefix(
						"productSet.product.publication",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication.Fields
						);
						},
			parameters$: 'parameters',
get parameters() {
					return withPrefix(
						"productSet.product.parameters[:i]",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsiblePersonType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsiblePersonType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsiblePersonType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsibleProducerType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsibleProducerType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsibleProducerType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.SafetyInformationType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.SafetyInformationType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.SafetyInformationType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.DepositsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.DepositsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.DepositsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits(this.toJSON());
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
			if (d.quantity !== undefined) { this.quantity = d.quantity }
			if (d.product !== undefined) { this.product = d.product }
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
			if (!(d.quantity instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity)) { this.quantity = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity(d.quantity || {}) }	
			if (!(d.product instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product)) { this.product = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product(d.product || {}) }	
			if (!(d.responsiblePerson instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson)) { this.responsiblePerson = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson(d.responsiblePerson || {}) }	
			if (!(d.responsibleProducer instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer)) { this.responsibleProducer = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer(d.responsibleProducer || {}) }	
			if (!(d.safetyInformation instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation)) { this.safetyInformation = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation(d.safetyInformation || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				quantity: this.#quantity,
				product: this.#product,
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
			quantity$: 'quantity',
get quantity() {
					return withPrefix(
						"productSet.quantity",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity.Fields
						);
						},
			product$: 'product',
get product() {
					return withPrefix(
						"productSet.product",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Fields
						);
						},
			responsiblePerson$: 'responsiblePerson',
get responsiblePerson() {
					return withPrefix(
						"productSet.responsiblePerson",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson.Fields
						);
						},
			responsibleProducer$: 'responsibleProducer',
get responsibleProducer() {
					return withPrefix(
						"productSet.responsibleProducer",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer.Fields
						);
						},
			safetyInformation$: 'safetyInformation',
get safetyInformation() {
					return withPrefix(
						"productSet.safetyInformation",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation.Fields
						);
						},
			marketedBeforeGPSRObligation: 'marketedBeforeGPSRObligation',
			deposits$: 'deposits',
get deposits() {
					return withPrefix(
						"productSet.deposits[:i]",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ProductSetType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ProductSetType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.ProductSet> {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Attachments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.AttachmentsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Attachments(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Attachments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AttachmentsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Attachments(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AttachmentsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Attachments> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Attachments ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Attachments> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Attachments(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.FundraisingCampaignType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.FundraisingCampaignType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.FundraisingCampaignType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign> {
		return new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign> {
		return new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.AdditionalServicesType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AdditionalServicesType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.AdditionalServicesType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices> {
		return new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.B2b, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.B2bType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.B2b(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.B2b, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.B2bType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.B2b(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.B2bType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.B2b> {
		return new GetAllDataOfTheParticularProductOfferActionRes.B2b ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.B2b> {
		return new GetAllDataOfTheParticularProductOfferActionRes.B2b(this.toJSON());
	}
}
/**
  * The base class definition for compatibilityList
  **/
static CompatibilityList = class CompatibilityList {
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
		const d = data as Partial<CompatibilityList>;
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.CompatibilityListType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.CompatibilityListType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.CompatibilityListType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList> {
		return new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList> {
		return new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList(this.toJSON());
	}
}
/**
  * The base class definition for validation
  **/
static Validation = class Validation {
		/**
  * 
  * @type {string}
  **/
 #validatedAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get validatedAt () { return this.#validatedAt }
/**
  * 
  * @type {string}
  **/
set validatedAt (value: string) {
		this.#validatedAt = String(value);
}
setValidatedAt (value: string) {
	this.validatedAt = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors}
  **/
 #errors : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors}
  **/
set errors (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors) {
			this.#errors = value
		} else {
			this.#errors = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors(item))
		}
}
setErrors (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors>[]) {
	this.errors = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings}
  **/
 #warnings : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings}
  **/
get warnings () { return this.#warnings }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings}
  **/
set warnings (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings) {
			this.#warnings = value
		} else {
			this.#warnings = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings(item))
		}
}
setWarnings (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings>[]) {
	this.warnings = value
	return this
}
/**
  * The base class definition for errors
  **/
static Errors = class Errors {
		/**
  * 
  * @type {string}
  **/
 #code : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * 
  * @type {string}
  **/
set code (value: string) {
		this.#code = String(value);
}
setCode (value: string) {
	this.code = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #details : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get details () { return this.#details }
/**
  * 
  * @type {string}
  **/
set details (value: string) {
		this.#details = String(value);
}
setDetails (value: string) {
	this.details = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #message : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get message () { return this.#message }
/**
  * 
  * @type {string}
  **/
set message (value: string) {
		this.#message = String(value);
}
setMessage (value: string) {
	this.message = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #path : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get path () { return this.#path }
/**
  * 
  * @type {string}
  **/
set path (value: string) {
		this.#path = String(value);
}
setPath (value: string) {
	this.path = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #userMessage : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get userMessage () { return this.#userMessage }
/**
  * 
  * @type {string}
  **/
set userMessage (value: string) {
		this.#userMessage = String(value);
}
setUserMessage (value: string) {
	this.userMessage = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata}
  **/
 #metadata ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata}
  **/
set metadata (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata(value)
		}
}
setMetadata (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata>) {
	this.metadata = value
	return this
}
/**
  * The base class definition for metadata
  **/
static Metadata = class Metadata {
		/**
  * 
  * @type {string}
  **/
 #productId : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get productId () { return this.#productId }
/**
  * 
  * @type {string}
  **/
set productId (value: string) {
		this.#productId = String(value);
}
setProductId (value: string) {
	this.productId = value
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
		const d = data as Partial<Metadata>;
			if (d.productId !== undefined) { this.productId = d.productId }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				productId: this.#productId,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			productId: 'productId',
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType.MetadataType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType.MetadataType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType.MetadataType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata(this.toJSON());
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
		const d = data as Partial<Errors>;
			if (d.code !== undefined) { this.code = d.code }
			if (d.details !== undefined) { this.details = d.details }
			if (d.message !== undefined) { this.message = d.message }
			if (d.path !== undefined) { this.path = d.path }
			if (d.userMessage !== undefined) { this.userMessage = d.userMessage }
			if (d.metadata !== undefined) { this.metadata = d.metadata }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Errors>;
			if (!(d.metadata instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata)) { this.metadata = new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata(d.metadata || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				code: this.#code,
				details: this.#details,
				message: this.#message,
				path: this.#path,
				userMessage: this.#userMessage,
				metadata: this.#metadata,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			code: 'code',
			details: 'details',
			message: 'message',
			path: 'path',
			userMessage: 'userMessage',
			metadata$: 'metadata',
get metadata() {
					return withPrefix(
						"validation.errors.metadata",
						GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors(this.toJSON());
	}
}
/**
  * The base class definition for warnings
  **/
static Warnings = class Warnings {
		/**
  * 
  * @type {string}
  **/
 #code : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * 
  * @type {string}
  **/
set code (value: string) {
		this.#code = String(value);
}
setCode (value: string) {
	this.code = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #details : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get details () { return this.#details }
/**
  * 
  * @type {string}
  **/
set details (value: string) {
		this.#details = String(value);
}
setDetails (value: string) {
	this.details = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #message : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get message () { return this.#message }
/**
  * 
  * @type {string}
  **/
set message (value: string) {
		this.#message = String(value);
}
setMessage (value: string) {
	this.message = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #path : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get path () { return this.#path }
/**
  * 
  * @type {string}
  **/
set path (value: string) {
		this.#path = String(value);
}
setPath (value: string) {
	this.path = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #userMessage : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get userMessage () { return this.#userMessage }
/**
  * 
  * @type {string}
  **/
set userMessage (value: string) {
		this.#userMessage = String(value);
}
setUserMessage (value: string) {
	this.userMessage = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata}
  **/
 #metadata ! : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata>
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata}
  **/
set metadata (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata(value)
		}
}
setMetadata (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata>) {
	this.metadata = value
	return this
}
/**
  * The base class definition for metadata
  **/
static Metadata = class Metadata {
		/**
  * 
  * @type {string}
  **/
 #productId : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get productId () { return this.#productId }
/**
  * 
  * @type {string}
  **/
set productId (value: string) {
		this.#productId = String(value);
}
setProductId (value: string) {
	this.productId = value
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
		const d = data as Partial<Metadata>;
			if (d.productId !== undefined) { this.productId = d.productId }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				productId: this.#productId,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			productId: 'productId',
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType.MetadataType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType.MetadataType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType.MetadataType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata(this.toJSON());
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
		const d = data as Partial<Warnings>;
			if (d.code !== undefined) { this.code = d.code }
			if (d.details !== undefined) { this.details = d.details }
			if (d.message !== undefined) { this.message = d.message }
			if (d.path !== undefined) { this.path = d.path }
			if (d.userMessage !== undefined) { this.userMessage = d.userMessage }
			if (d.metadata !== undefined) { this.metadata = d.metadata }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Warnings>;
			if (!(d.metadata instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata)) { this.metadata = new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata(d.metadata || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				code: this.#code,
				details: this.#details,
				message: this.#message,
				path: this.#path,
				userMessage: this.#userMessage,
				metadata: this.#metadata,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			code: 'code',
			details: 'details',
			message: 'message',
			path: 'path',
			userMessage: 'userMessage',
			metadata$: 'metadata',
get metadata() {
					return withPrefix(
						"validation.warnings.metadata",
						GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings(this.toJSON());
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
		const d = data as Partial<Validation>;
			if (d.validatedAt !== undefined) { this.validatedAt = d.validatedAt }
			if (d.errors !== undefined) { this.errors = d.errors }
			if (d.warnings !== undefined) { this.warnings = d.warnings }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				validatedAt: this.#validatedAt,
				errors: this.#errors,
				warnings: this.#warnings,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			validatedAt: 'validatedAt',
			errors$: 'errors',
get errors() {
					return withPrefix(
						"validation.errors[:i]",
						GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Fields
						);
						},
			warnings$: 'warnings',
get warnings() {
					return withPrefix(
						"validation.warnings[:i]",
						GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ValidationType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ValidationType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.Validation> {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.External, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.ExternalType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.External(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.External, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ExternalType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.External(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.ExternalType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.External> {
		return new GetAllDataOfTheParticularProductOfferActionRes.External ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.External> {
		return new GetAllDataOfTheParticularProductOfferActionRes.External(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SizeTable, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.SizeTableType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SizeTable(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SizeTable, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SizeTableType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SizeTable(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.SizeTableType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SizeTable> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SizeTable ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.SizeTable> {
		return new GetAllDataOfTheParticularProductOfferActionRes.SizeTable(this.toJSON());
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
  * @type {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates}
  **/
 #rates : InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates>[]  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates}
  **/
get rates () { return this.#rates }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates}
  **/
set rates (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates) {
			this.#rates = value
		} else {
			this.#rates = value.map(item => new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates(item))
		}
}
setRates (value: InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates>[]) {
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType.RatesType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType.RatesType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType.RatesType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates> {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates> {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates(this.toJSON());
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
						GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.TaxSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.TaxSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings> {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings> {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings(this.toJSON());
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
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType.MessageToSellerSettingsType) {
		return new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.MessageToSellerSettingsType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType.MessageToSellerSettingsType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings> {
		return new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings> {
		return new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings(this.toJSON());
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
		const d = data as Partial<GetAllDataOfTheParticularProductOfferActionRes>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
			if (d.language !== undefined) { this.language = d.language }
			if (d.createdAt !== undefined) { this.createdAt = d.createdAt }
			if (d.updatedAt !== undefined) { this.updatedAt = d.updatedAt }
			if (d.category !== undefined) { this.category = d.category }
			if (d.stock !== undefined) { this.stock = d.stock }
			if (d.contact !== undefined) { this.contact = d.contact }
			if (d.publication !== undefined) { this.publication = d.publication }
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
			if (d.payments !== undefined) { this.payments = d.payments }
			if (d.delivery !== undefined) { this.delivery = d.delivery }
			if (d.afterSalesServices !== undefined) { this.afterSalesServices = d.afterSalesServices }
			if (d.discounts !== undefined) { this.discounts = d.discounts }
			if (d.description !== undefined) { this.description = d.description }
			if (d.images !== undefined) { this.images = d.images }
			if (d.productSet !== undefined) { this.productSet = d.productSet }
			if (d.attachments !== undefined) { this.attachments = d.attachments }
			if (d.fundraisingCampaign !== undefined) { this.fundraisingCampaign = d.fundraisingCampaign }
			if (d.additionalServices !== undefined) { this.additionalServices = d.additionalServices }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
			if (d.b2b !== undefined) { this.b2b = d.b2b }
			if (d.compatibilityList !== undefined) { this.compatibilityList = d.compatibilityList }
			if (d.validation !== undefined) { this.validation = d.validation }
			if (d.external !== undefined) { this.external = d.external }
			if (d.sizeTable !== undefined) { this.sizeTable = d.sizeTable }
			if (d.taxSettings !== undefined) { this.taxSettings = d.taxSettings }
			if (d.messageToSellerSettings !== undefined) { this.messageToSellerSettings = d.messageToSellerSettings }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<GetAllDataOfTheParticularProductOfferActionRes>;
			if (!(d.category instanceof GetAllDataOfTheParticularProductOfferActionRes.Category)) { this.category = new GetAllDataOfTheParticularProductOfferActionRes.Category(d.category || {}) }	
			if (!(d.stock instanceof GetAllDataOfTheParticularProductOfferActionRes.Stock)) { this.stock = new GetAllDataOfTheParticularProductOfferActionRes.Stock(d.stock || {}) }	
			if (!(d.contact instanceof GetAllDataOfTheParticularProductOfferActionRes.Contact)) { this.contact = new GetAllDataOfTheParticularProductOfferActionRes.Contact(d.contact || {}) }	
			if (!(d.publication instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication)) { this.publication = new GetAllDataOfTheParticularProductOfferActionRes.Publication(d.publication || {}) }	
			if (!(d.sellingMode instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode)) { this.sellingMode = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode(d.sellingMode || {}) }	
			if (!(d.payments instanceof GetAllDataOfTheParticularProductOfferActionRes.Payments)) { this.payments = new GetAllDataOfTheParticularProductOfferActionRes.Payments(d.payments || {}) }	
			if (!(d.delivery instanceof GetAllDataOfTheParticularProductOfferActionRes.Delivery)) { this.delivery = new GetAllDataOfTheParticularProductOfferActionRes.Delivery(d.delivery || {}) }	
			if (!(d.afterSalesServices instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices)) { this.afterSalesServices = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices(d.afterSalesServices || {}) }	
			if (!(d.discounts instanceof GetAllDataOfTheParticularProductOfferActionRes.Discounts)) { this.discounts = new GetAllDataOfTheParticularProductOfferActionRes.Discounts(d.discounts || {}) }	
			if (!(d.description instanceof GetAllDataOfTheParticularProductOfferActionRes.Description)) { this.description = new GetAllDataOfTheParticularProductOfferActionRes.Description(d.description || {}) }	
			if (!(d.fundraisingCampaign instanceof GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign)) { this.fundraisingCampaign = new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign(d.fundraisingCampaign || {}) }	
			if (!(d.additionalServices instanceof GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices)) { this.additionalServices = new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices(d.additionalServices || {}) }	
			if (!(d.b2b instanceof GetAllDataOfTheParticularProductOfferActionRes.B2b)) { this.b2b = new GetAllDataOfTheParticularProductOfferActionRes.B2b(d.b2b || {}) }	
			if (!(d.compatibilityList instanceof GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList)) { this.compatibilityList = new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList(d.compatibilityList || {}) }	
			if (!(d.validation instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation)) { this.validation = new GetAllDataOfTheParticularProductOfferActionRes.Validation(d.validation || {}) }	
			if (!(d.external instanceof GetAllDataOfTheParticularProductOfferActionRes.External)) { this.external = new GetAllDataOfTheParticularProductOfferActionRes.External(d.external || {}) }	
			if (!(d.sizeTable instanceof GetAllDataOfTheParticularProductOfferActionRes.SizeTable)) { this.sizeTable = new GetAllDataOfTheParticularProductOfferActionRes.SizeTable(d.sizeTable || {}) }	
			if (!(d.taxSettings instanceof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings)) { this.taxSettings = new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings(d.taxSettings || {}) }	
			if (!(d.messageToSellerSettings instanceof GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings)) { this.messageToSellerSettings = new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings(d.messageToSellerSettings || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
				language: this.#language,
				createdAt: this.#createdAt,
				updatedAt: this.#updatedAt,
				category: this.#category,
				stock: this.#stock,
				contact: this.#contact,
				publication: this.#publication,
				sellingMode: this.#sellingMode,
				payments: this.#payments,
				delivery: this.#delivery,
				afterSalesServices: this.#afterSalesServices,
				discounts: this.#discounts,
				description: this.#description,
				images: this.#images,
				productSet: this.#productSet,
				attachments: this.#attachments,
				fundraisingCampaign: this.#fundraisingCampaign,
				additionalServices: this.#additionalServices,
				additionalMarketplaces: this.#additionalMarketplaces,
				b2b: this.#b2b,
				compatibilityList: this.#compatibilityList,
				validation: this.#validation,
				external: this.#external,
				sizeTable: this.#sizeTable,
				taxSettings: this.#taxSettings,
				messageToSellerSettings: this.#messageToSellerSettings,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
			language: 'language',
			createdAt: 'createdAt',
			updatedAt: 'updatedAt',
			category$: 'category',
get category() {
					return withPrefix(
						"category",
						GetAllDataOfTheParticularProductOfferActionRes.Category.Fields
						);
						},
			stock$: 'stock',
get stock() {
					return withPrefix(
						"stock",
						GetAllDataOfTheParticularProductOfferActionRes.Stock.Fields
						);
						},
			contact$: 'contact',
get contact() {
					return withPrefix(
						"contact",
						GetAllDataOfTheParticularProductOfferActionRes.Contact.Fields
						);
						},
			publication$: 'publication',
get publication() {
					return withPrefix(
						"publication",
						GetAllDataOfTheParticularProductOfferActionRes.Publication.Fields
						);
						},
			sellingMode$: 'sellingMode',
get sellingMode() {
					return withPrefix(
						"sellingMode",
						GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Fields
						);
						},
			payments$: 'payments',
get payments() {
					return withPrefix(
						"payments",
						GetAllDataOfTheParticularProductOfferActionRes.Payments.Fields
						);
						},
			delivery$: 'delivery',
get delivery() {
					return withPrefix(
						"delivery",
						GetAllDataOfTheParticularProductOfferActionRes.Delivery.Fields
						);
						},
			afterSalesServices$: 'afterSalesServices',
get afterSalesServices() {
					return withPrefix(
						"afterSalesServices",
						GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Fields
						);
						},
			discounts$: 'discounts',
get discounts() {
					return withPrefix(
						"discounts",
						GetAllDataOfTheParticularProductOfferActionRes.Discounts.Fields
						);
						},
			description$: 'description',
get description() {
					return withPrefix(
						"description",
						GetAllDataOfTheParticularProductOfferActionRes.Description.Fields
						);
						},
			images$: 'images',
get images() {
					return "images[:i]";
						},
			productSet$: 'productSet',
get productSet() {
					return withPrefix(
						"productSet[:i]",
						GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Fields
						);
						},
			attachments$: 'attachments',
get attachments() {
					return withPrefix(
						"attachments[:i]",
						GetAllDataOfTheParticularProductOfferActionRes.Attachments.Fields
						);
						},
			fundraisingCampaign$: 'fundraisingCampaign',
get fundraisingCampaign() {
					return withPrefix(
						"fundraisingCampaign",
						GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign.Fields
						);
						},
			additionalServices$: 'additionalServices',
get additionalServices() {
					return withPrefix(
						"additionalServices",
						GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices.Fields
						);
						},
			additionalMarketplaces: 'additionalMarketplaces',
			b2b$: 'b2b',
get b2b() {
					return withPrefix(
						"b2b",
						GetAllDataOfTheParticularProductOfferActionRes.B2b.Fields
						);
						},
			compatibilityList$: 'compatibilityList',
get compatibilityList() {
					return withPrefix(
						"compatibilityList",
						GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList.Fields
						);
						},
			validation$: 'validation',
get validation() {
					return withPrefix(
						"validation",
						GetAllDataOfTheParticularProductOfferActionRes.Validation.Fields
						);
						},
			external$: 'external',
get external() {
					return withPrefix(
						"external",
						GetAllDataOfTheParticularProductOfferActionRes.External.Fields
						);
						},
			sizeTable$: 'sizeTable',
get sizeTable() {
					return withPrefix(
						"sizeTable",
						GetAllDataOfTheParticularProductOfferActionRes.SizeTable.Fields
						);
						},
			taxSettings$: 'taxSettings',
get taxSettings() {
					return withPrefix(
						"taxSettings",
						GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Fields
						);
						},
			messageToSellerSettings$: 'messageToSellerSettings',
get messageToSellerSettings() {
					return withPrefix(
						"messageToSellerSettings",
						GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetAllDataOfTheParticularProductOfferActionResType) {
		return new GetAllDataOfTheParticularProductOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType>) {
		return new GetAllDataOfTheParticularProductOfferActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetAllDataOfTheParticularProductOfferActionResType>): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes> {
		return new GetAllDataOfTheParticularProductOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetAllDataOfTheParticularProductOfferActionRes> {
		return new GetAllDataOfTheParticularProductOfferActionRes(this.toJSON());
	}
}
export abstract class GetAllDataOfTheParticularProductOfferActionResFactory {
	abstract create(data: unknown): GetAllDataOfTheParticularProductOfferActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for getAllDataOfTheParticularProductOfferActionRes
  **/
	export type GetAllDataOfTheParticularProductOfferActionResType =  {
			/**
  * Unique offer identifier
  * @type {string}
  **/
 id : string;
			/**
  * Offer title
  * @type {string}
  **/
 name : string;
			/**
  * Offer language code (e.g. pl-PL)
  * @type {string}
  **/
 language : string;
			/**
  * Offer creation timestamp (ISO8601)
  * @type {string}
  **/
 createdAt : string;
			/**
  * Offer last update timestamp (ISO8601)
  * @type {string}
  **/
 updatedAt : string;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.CategoryType}
  **/
 category : GetAllDataOfTheParticularProductOfferActionResType.CategoryType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.StockType}
  **/
 stock : GetAllDataOfTheParticularProductOfferActionResType.StockType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ContactType}
  **/
 contact : GetAllDataOfTheParticularProductOfferActionResType.ContactType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.PublicationType}
  **/
 publication : GetAllDataOfTheParticularProductOfferActionResType.PublicationType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.SellingModeType}
  **/
 sellingMode : GetAllDataOfTheParticularProductOfferActionResType.SellingModeType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.PaymentsType}
  **/
 payments : GetAllDataOfTheParticularProductOfferActionResType.PaymentsType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.DeliveryType}
  **/
 delivery : GetAllDataOfTheParticularProductOfferActionResType.DeliveryType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType}
  **/
 afterSalesServices : GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.DiscountsType}
  **/
 discounts : GetAllDataOfTheParticularProductOfferActionResType.DiscountsType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.DescriptionType}
  **/
 description : GetAllDataOfTheParticularProductOfferActionResType.DescriptionType;
			/**
  * 
  * @type {string[]}
  **/
 images : string[];
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType[]}
  **/
 productSet : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType[];
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.AttachmentsType[]}
  **/
 attachments : GetAllDataOfTheParticularProductOfferActionResType.AttachmentsType[];
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.FundraisingCampaignType}
  **/
 fundraisingCampaign : GetAllDataOfTheParticularProductOfferActionResType.FundraisingCampaignType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.AdditionalServicesType}
  **/
 additionalServices : GetAllDataOfTheParticularProductOfferActionResType.AdditionalServicesType;
			/**
  * 
  * @type {{[key: string]: any}}
  **/
 additionalMarketplaces ?: {[key: string]: any};
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.B2bType}
  **/
 b2b : GetAllDataOfTheParticularProductOfferActionResType.B2bType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.CompatibilityListType}
  **/
 compatibilityList : GetAllDataOfTheParticularProductOfferActionResType.CompatibilityListType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ValidationType}
  **/
 validation : GetAllDataOfTheParticularProductOfferActionResType.ValidationType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ExternalType}
  **/
 external : GetAllDataOfTheParticularProductOfferActionResType.ExternalType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.SizeTableType}
  **/
 sizeTable : GetAllDataOfTheParticularProductOfferActionResType.SizeTableType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType}
  **/
 taxSettings : GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.MessageToSellerSettingsType}
  **/
 messageToSellerSettings : GetAllDataOfTheParticularProductOfferActionResType.MessageToSellerSettingsType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetAllDataOfTheParticularProductOfferActionResType {
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
  * The base type definition for contactType
  **/
	export type ContactType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ContactType {
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
 endedBy : string;
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
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType}
  **/
 marketplaces : GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PublicationType {
	/**
  * The base type definition for marketplacesType
  **/
	export type MarketplacesType =  {
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.BaseType}
  **/
 base : GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.BaseType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.AdditionalType[]}
  **/
 additional : GetAllDataOfTheParticularProductOfferActionResType.PublicationType.MarketplacesType.AdditionalType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MarketplacesType {
	/**
  * The base type definition for baseType
  **/
	export type BaseType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace BaseType {
}
	/**
  * The base type definition for additionalType
  **/
	export type AdditionalType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdditionalType {
}
}
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
  * @type {GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType}
  **/
 price : GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.PriceType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.MinimalPriceType}
  **/
 minimalPrice : GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.MinimalPriceType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.StartingPriceType}
  **/
 startingPrice : GetAllDataOfTheParticularProductOfferActionResType.SellingModeType.StartingPriceType;
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
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.DeliveryType.ShippingRatesType}
  **/
 shippingRates : GetAllDataOfTheParticularProductOfferActionResType.DeliveryType.ShippingRatesType;
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
  * The base type definition for afterSalesServicesType
  **/
	export type AfterSalesServicesType =  {
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType}
  **/
 impliedWarranty : GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ReturnPolicyType}
  **/
 returnPolicy : GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.ReturnPolicyType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.WarrantyType}
  **/
 warranty : GetAllDataOfTheParticularProductOfferActionResType.AfterSalesServicesType.WarrantyType;
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
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WarrantyType {
}
}
	/**
  * The base type definition for discountsType
  **/
	export type DiscountsType =  {
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.DiscountsType.WholesalePriceListType}
  **/
 wholesalePriceList : GetAllDataOfTheParticularProductOfferActionResType.DiscountsType.WholesalePriceListType;
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
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WholesalePriceListType {
}
}
	/**
  * The base type definition for descriptionType
  **/
	export type DescriptionType =  {
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType[]}
  **/
 sections : GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DescriptionType {
	/**
  * The base type definition for sectionsType
  **/
	export type SectionsType =  {
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType.ItemsType[]}
  **/
 items : GetAllDataOfTheParticularProductOfferActionResType.DescriptionType.SectionsType.ItemsType[];
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
  * The base type definition for productSetType
  **/
	export type ProductSetType =  {
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.QuantityType}
  **/
 quantity : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.QuantityType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType}
  **/
 product : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsiblePersonType}
  **/
 responsiblePerson : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsiblePersonType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsibleProducerType}
  **/
 responsibleProducer : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ResponsibleProducerType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.SafetyInformationType}
  **/
 safetyInformation : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.SafetyInformationType;
			/**
  * 
  * @type {boolean}
  **/
 marketedBeforeGPSRObligation : boolean;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.DepositsType[]}
  **/
 deposits : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.DepositsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ProductSetType {
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
  * @type {boolean}
  **/
 isAiCoCreated : boolean;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.PublicationType}
  **/
 publication : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.PublicationType;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType[]}
  **/
 parameters : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ProductType {
	/**
  * The base type definition for publicationType
  **/
	export type PublicationType =  {
			/**
  * 
  * @type {string}
  **/
 status : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PublicationType {
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
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType}
  **/
 rangeValue : GetAllDataOfTheParticularProductOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType;
			/**
  * 
  * @type {any[]}
  **/
 values : any[];
			/**
  * 
  * @type {any[]}
  **/
 valuesIds : any[];
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
  * The base type definition for responsiblePersonType
  **/
	export type ResponsiblePersonType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
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
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdditionalServicesType {
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
  * The base type definition for compatibilityListType
  **/
	export type CompatibilityListType =  {
			/**
  * 
  * @type {string}
  **/
 type : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CompatibilityListType {
}
	/**
  * The base type definition for validationType
  **/
	export type ValidationType =  {
			/**
  * 
  * @type {string}
  **/
 validatedAt : string;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType[]}
  **/
 errors : GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType[];
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType[]}
  **/
 warnings : GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ValidationType {
	/**
  * The base type definition for errorsType
  **/
	export type ErrorsType =  {
			/**
  * 
  * @type {string}
  **/
 code : string;
			/**
  * 
  * @type {string}
  **/
 details : string;
			/**
  * 
  * @type {string}
  **/
 message : string;
			/**
  * 
  * @type {string}
  **/
 path : string;
			/**
  * 
  * @type {string}
  **/
 userMessage : string;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType.MetadataType}
  **/
 metadata : GetAllDataOfTheParticularProductOfferActionResType.ValidationType.ErrorsType.MetadataType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ErrorsType {
	/**
  * The base type definition for metadataType
  **/
	export type MetadataType =  {
			/**
  * 
  * @type {string}
  **/
 productId : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MetadataType {
}
}
	/**
  * The base type definition for warningsType
  **/
	export type WarningsType =  {
			/**
  * 
  * @type {string}
  **/
 code : string;
			/**
  * 
  * @type {string}
  **/
 details : string;
			/**
  * 
  * @type {string}
  **/
 message : string;
			/**
  * 
  * @type {string}
  **/
 path : string;
			/**
  * 
  * @type {string}
  **/
 userMessage : string;
			/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType.MetadataType}
  **/
 metadata : GetAllDataOfTheParticularProductOfferActionResType.ValidationType.WarningsType.MetadataType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WarningsType {
	/**
  * The base type definition for metadataType
  **/
	export type MetadataType =  {
			/**
  * 
  * @type {string}
  **/
 productId : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MetadataType {
}
}
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
  * The base type definition for sizeTableType
  **/
	export type SizeTableType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SizeTableType {
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
  * @type {GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType.RatesType[]}
  **/
 rates : GetAllDataOfTheParticularProductOfferActionResType.TaxSettingsType.RatesType[];
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