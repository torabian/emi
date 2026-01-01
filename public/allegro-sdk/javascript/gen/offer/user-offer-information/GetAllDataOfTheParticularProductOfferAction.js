import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get all data of the particular product-offer
*/
	/**
 * GetAllDataOfTheParticularProductOfferAction
 */
export class GetAllDataOfTheParticularProductOfferAction { //
  static URL = 'https://api.{environment}/sale/product-offers/{offerId}';
  static NewUrl = (
	qs
  ) => buildUrl(
		GetAllDataOfTheParticularProductOfferAction.URL,
		 undefined,
		qs
	);
  static Method = 'get';
	static Fetch$ = async (
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
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
		init,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
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
 #id  =  ""
		/**
  * Unique offer identifier
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * Unique offer identifier
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
		/**
  * Offer title
  * @type {string}
  **/
 #name  =  ""
		/**
  * Offer title
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * Offer title
  * @type {string}
  **/
set name (value) {
		this.#name = String(value);
}
setName (value) {
	this.name = value
	return this
}
		/**
  * Offer language code (e.g. pl-PL)
  * @type {string}
  **/
 #language  =  ""
		/**
  * Offer language code (e.g. pl-PL)
  * @returns {string}
  **/
get language () { return this.#language }
/**
  * Offer language code (e.g. pl-PL)
  * @type {string}
  **/
set language (value) {
		this.#language = String(value);
}
setLanguage (value) {
	this.language = value
	return this
}
		/**
  * Offer creation timestamp (ISO8601)
  * @type {string}
  **/
 #createdAt  =  ""
		/**
  * Offer creation timestamp (ISO8601)
  * @returns {string}
  **/
get createdAt () { return this.#createdAt }
/**
  * Offer creation timestamp (ISO8601)
  * @type {string}
  **/
set createdAt (value) {
		this.#createdAt = String(value);
}
setCreatedAt (value) {
	this.createdAt = value
	return this
}
		/**
  * Offer last update timestamp (ISO8601)
  * @type {string}
  **/
 #updatedAt  =  ""
		/**
  * Offer last update timestamp (ISO8601)
  * @returns {string}
  **/
get updatedAt () { return this.#updatedAt }
/**
  * Offer last update timestamp (ISO8601)
  * @type {string}
  **/
set updatedAt (value) {
		this.#updatedAt = String(value);
}
setUpdatedAt (value) {
	this.updatedAt = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Category}
  **/
 #category
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Category}
  **/
set category (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Category) {
			this.#category = value
		} else {
			this.#category = new GetAllDataOfTheParticularProductOfferActionRes.Category(value)
		}
}
setCategory (value) {
	this.category = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Stock}
  **/
 #stock
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Stock}
  **/
set stock (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Stock) {
			this.#stock = value
		} else {
			this.#stock = new GetAllDataOfTheParticularProductOfferActionRes.Stock(value)
		}
}
setStock (value) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Contact}
  **/
 #contact
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Contact}
  **/
get contact () { return this.#contact }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Contact}
  **/
set contact (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Contact) {
			this.#contact = value
		} else {
			this.#contact = new GetAllDataOfTheParticularProductOfferActionRes.Contact(value)
		}
}
setContact (value) {
	this.contact = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication}
  **/
 #publication
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication}
  **/
set publication (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication) {
			this.#publication = value
		} else {
			this.#publication = new GetAllDataOfTheParticularProductOfferActionRes.Publication(value)
		}
}
setPublication (value) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
 #sellingMode
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode}
  **/
set sellingMode (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode(value)
		}
}
setSellingMode (value) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Payments}
  **/
 #payments
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Payments}
  **/
get payments () { return this.#payments }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Payments}
  **/
set payments (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Payments) {
			this.#payments = value
		} else {
			this.#payments = new GetAllDataOfTheParticularProductOfferActionRes.Payments(value)
		}
}
setPayments (value) {
	this.payments = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Delivery}
  **/
 #delivery
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Delivery}
  **/
get delivery () { return this.#delivery }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Delivery}
  **/
set delivery (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Delivery) {
			this.#delivery = value
		} else {
			this.#delivery = new GetAllDataOfTheParticularProductOfferActionRes.Delivery(value)
		}
}
setDelivery (value) {
	this.delivery = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices}
  **/
 #afterSalesServices
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices}
  **/
get afterSalesServices () { return this.#afterSalesServices }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices}
  **/
set afterSalesServices (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices) {
			this.#afterSalesServices = value
		} else {
			this.#afterSalesServices = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices(value)
		}
}
setAfterSalesServices (value) {
	this.afterSalesServices = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Discounts}
  **/
 #discounts
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Discounts}
  **/
get discounts () { return this.#discounts }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Discounts}
  **/
set discounts (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Discounts) {
			this.#discounts = value
		} else {
			this.#discounts = new GetAllDataOfTheParticularProductOfferActionRes.Discounts(value)
		}
}
setDiscounts (value) {
	this.discounts = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description}
  **/
 #description
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description}
  **/
set description (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Description) {
			this.#description = value
		} else {
			this.#description = new GetAllDataOfTheParticularProductOfferActionRes.Description(value)
		}
}
setDescription (value) {
	this.description = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #images  =  []
		/**
  * 
  * @returns {string[]}
  **/
get images () { return this.#images }
/**
  * 
  * @type {string[]}
  **/
set images (value) {
}
setImages (value) {
	this.images = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet}
  **/
 #productSet  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet}
  **/
get productSet () { return this.#productSet }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet}
  **/
set productSet (value) {
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
setProductSet (value) {
	this.productSet = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Attachments}
  **/
 #attachments  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Attachments}
  **/
get attachments () { return this.#attachments }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Attachments}
  **/
set attachments (value) {
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
setAttachments (value) {
	this.attachments = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign}
  **/
 #fundraisingCampaign
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign}
  **/
get fundraisingCampaign () { return this.#fundraisingCampaign }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign}
  **/
set fundraisingCampaign (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign) {
			this.#fundraisingCampaign = value
		} else {
			this.#fundraisingCampaign = new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign(value)
		}
}
setFundraisingCampaign (value) {
	this.fundraisingCampaign = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices}
  **/
 #additionalServices
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices}
  **/
get additionalServices () { return this.#additionalServices }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices}
  **/
set additionalServices (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices) {
			this.#additionalServices = value
		} else {
			this.#additionalServices = new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices(value)
		}
}
setAdditionalServices (value) {
	this.additionalServices = value
	return this
}
		/**
  * 
  * @type {{[key: string]: any}}
  **/
 #additionalMarketplaces  =  undefined
		/**
  * 
  * @returns {{[key: string]: any}}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {{[key: string]: any}}
  **/
set additionalMarketplaces (value) {
}
setAdditionalMarketplaces (value) {
	this.additionalMarketplaces = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.B2b}
  **/
 #b2b
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.B2b}
  **/
get b2b () { return this.#b2b }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.B2b}
  **/
set b2b (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.B2b) {
			this.#b2b = value
		} else {
			this.#b2b = new GetAllDataOfTheParticularProductOfferActionRes.B2b(value)
		}
}
setB2b (value) {
	this.b2b = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList}
  **/
 #compatibilityList
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList}
  **/
get compatibilityList () { return this.#compatibilityList }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList}
  **/
set compatibilityList (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList) {
			this.#compatibilityList = value
		} else {
			this.#compatibilityList = new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList(value)
		}
}
setCompatibilityList (value) {
	this.compatibilityList = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation}
  **/
 #validation
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation}
  **/
get validation () { return this.#validation }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation}
  **/
set validation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation) {
			this.#validation = value
		} else {
			this.#validation = new GetAllDataOfTheParticularProductOfferActionRes.Validation(value)
		}
}
setValidation (value) {
	this.validation = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.External}
  **/
 #external
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.External}
  **/
get external () { return this.#external }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.External}
  **/
set external (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.External) {
			this.#external = value
		} else {
			this.#external = new GetAllDataOfTheParticularProductOfferActionRes.External(value)
		}
}
setExternal (value) {
	this.external = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SizeTable}
  **/
 #sizeTable
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SizeTable}
  **/
get sizeTable () { return this.#sizeTable }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SizeTable}
  **/
set sizeTable (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SizeTable) {
			this.#sizeTable = value
		} else {
			this.#sizeTable = new GetAllDataOfTheParticularProductOfferActionRes.SizeTable(value)
		}
}
setSizeTable (value) {
	this.sizeTable = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings}
  **/
 #taxSettings
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings}
  **/
get taxSettings () { return this.#taxSettings }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings}
  **/
set taxSettings (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.TaxSettings) {
			this.#taxSettings = value
		} else {
			this.#taxSettings = new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings(value)
		}
}
setTaxSettings (value) {
	this.taxSettings = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings}
  **/
 #messageToSellerSettings
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings}
  **/
get messageToSellerSettings () { return this.#messageToSellerSettings }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings}
  **/
set messageToSellerSettings (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings) {
			this.#messageToSellerSettings = value
		} else {
			this.#messageToSellerSettings = new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings(value)
		}
}
setMessageToSellerSettings (value) {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Category(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Category ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #available  =  0
		/**
  * 
  * @returns {number}
  **/
get available () { return this.#available }
/**
  * 
  * @type {number}
  **/
set available (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#available = parsedValue;
		}
}
setAvailable (value) {
	this.available = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #unit  =  ""
		/**
  * 
  * @returns {string}
  **/
get unit () { return this.#unit }
/**
  * 
  * @type {string}
  **/
set unit (value) {
		this.#unit = String(value);
}
setUnit (value) {
	this.unit = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Stock(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Stock ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Contact(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Contact, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Contact(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Contact ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #duration  =  ""
		/**
  * 
  * @returns {string}
  **/
get duration () { return this.#duration }
/**
  * 
  * @type {string}
  **/
set duration (value) {
		this.#duration = String(value);
}
setDuration (value) {
	this.duration = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #startingAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get startingAt () { return this.#startingAt }
/**
  * 
  * @type {string}
  **/
set startingAt (value) {
		this.#startingAt = String(value);
}
setStartingAt (value) {
	this.startingAt = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #endingAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get endingAt () { return this.#endingAt }
/**
  * 
  * @type {string}
  **/
set endingAt (value) {
		this.#endingAt = String(value);
}
setEndingAt (value) {
	this.endingAt = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #endedBy  =  ""
		/**
  * 
  * @returns {string}
  **/
get endedBy () { return this.#endedBy }
/**
  * 
  * @type {string}
  **/
set endedBy (value) {
		this.#endedBy = String(value);
}
setEndedBy (value) {
	this.endedBy = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #status  =  ""
		/**
  * 
  * @returns {string}
  **/
get status () { return this.#status }
/**
  * 
  * @type {string}
  **/
set status (value) {
		this.#status = String(value);
}
setStatus (value) {
	this.status = value
	return this
}
		/**
  * 
  * @type {boolean}
  **/
 #republish
		/**
  * 
  * @returns {boolean}
  **/
get republish () { return this.#republish }
/**
  * 
  * @type {boolean}
  **/
set republish (value) {
		this.#republish = Boolean(value);
}
setRepublish (value) {
	this.republish = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces}
  **/
 #marketplaces
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces}
  **/
get marketplaces () { return this.#marketplaces }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces}
  **/
set marketplaces (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces) {
			this.#marketplaces = value
		} else {
			this.#marketplaces = new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces(value)
		}
}
setMarketplaces (value) {
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
 #base
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base}
  **/
get base () { return this.#base }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base}
  **/
set base (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base) {
			this.#base = value
		} else {
			this.#base = new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base(value)
		}
}
setBase (value) {
	this.base = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional}
  **/
 #additional  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional}
  **/
get additional () { return this.#additional }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional}
  **/
set additional (value) {
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
setAdditional (value) {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Base ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces.Additional(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
			if (d.base !== undefined) { this.base = d.base }
			if (d.additional !== undefined) { this.additional = d.additional }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication.Marketplaces(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Publication ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #format  =  ""
		/**
  * 
  * @returns {string}
  **/
get format () { return this.#format }
/**
  * 
  * @type {string}
  **/
set format (value) {
		this.#format = String(value);
}
setFormat (value) {
	this.format = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
 #price
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price}
  **/
set price (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price(value)
		}
}
setPrice (value) {
	this.price = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice}
  **/
 #minimalPrice
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice}
  **/
get minimalPrice () { return this.#minimalPrice }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice}
  **/
set minimalPrice (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice) {
			this.#minimalPrice = value
		} else {
			this.#minimalPrice = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice(value)
		}
}
setMinimalPrice (value) {
	this.minimalPrice = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice}
  **/
 #startingPrice
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice}
  **/
get startingPrice () { return this.#startingPrice }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice}
  **/
set startingPrice (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice) {
			this.#startingPrice = value
		} else {
			this.#startingPrice = new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice(value)
		}
}
setStartingPrice (value) {
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
 #amount  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value) {
		this.#amount = String(value);
}
setAmount (value) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value) {
		this.#currency = String(value);
}
setCurrency (value) {
	this.currency = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #amount  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value) {
		this.#amount = String(value);
}
setAmount (value) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value) {
		this.#currency = String(value);
}
setCurrency (value) {
	this.currency = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.MinimalPrice ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #amount  =  ""
		/**
  * 
  * @returns {string}
  **/
get amount () { return this.#amount }
/**
  * 
  * @type {string}
  **/
set amount (value) {
		this.#amount = String(value);
}
setAmount (value) {
	this.amount = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #currency  =  ""
		/**
  * 
  * @returns {string}
  **/
get currency () { return this.#currency }
/**
  * 
  * @type {string}
  **/
set currency (value) {
		this.#currency = String(value);
}
setCurrency (value) {
	this.currency = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode.StartingPrice(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #invoice  =  ""
		/**
  * 
  * @returns {string}
  **/
get invoice () { return this.#invoice }
/**
  * 
  * @type {string}
  **/
set invoice (value) {
		this.#invoice = String(value);
}
setInvoice (value) {
	this.invoice = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Payments(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Payments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Payments(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Payments ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #handlingTime  =  ""
		/**
  * 
  * @returns {string}
  **/
get handlingTime () { return this.#handlingTime }
/**
  * 
  * @type {string}
  **/
set handlingTime (value) {
		this.#handlingTime = String(value);
}
setHandlingTime (value) {
	this.handlingTime = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #additionalInfo  =  ""
		/**
  * 
  * @returns {string}
  **/
get additionalInfo () { return this.#additionalInfo }
/**
  * 
  * @type {string}
  **/
set additionalInfo (value) {
		this.#additionalInfo = String(value);
}
setAdditionalInfo (value) {
	this.additionalInfo = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #shipmentDate  =  ""
		/**
  * 
  * @returns {string}
  **/
get shipmentDate () { return this.#shipmentDate }
/**
  * 
  * @type {string}
  **/
set shipmentDate (value) {
		this.#shipmentDate = String(value);
}
setShipmentDate (value) {
	this.shipmentDate = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates}
  **/
 #shippingRates
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates}
  **/
get shippingRates () { return this.#shippingRates }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates}
  **/
set shippingRates (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates) {
			this.#shippingRates = value
		} else {
			this.#shippingRates = new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates(value)
		}
}
setShippingRates (value) {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery.ShippingRates(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Delivery, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Delivery ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #impliedWarranty
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
get impliedWarranty () { return this.#impliedWarranty }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
set impliedWarranty (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty) {
			this.#impliedWarranty = value
		} else {
			this.#impliedWarranty = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty(value)
		}
}
setImpliedWarranty (value) {
	this.impliedWarranty = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
 #returnPolicy
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
get returnPolicy () { return this.#returnPolicy }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
set returnPolicy (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy) {
			this.#returnPolicy = value
		} else {
			this.#returnPolicy = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy(value)
		}
}
setReturnPolicy (value) {
	this.returnPolicy = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty}
  **/
 #warranty
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty}
  **/
get warranty () { return this.#warranty }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty}
  **/
set warranty (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty) {
			this.#warranty = value
		} else {
			this.#warranty = new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty(value)
		}
}
setWarranty (value) {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ImpliedWarranty ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.ReturnPolicy ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices.Warranty(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
			if (d.impliedWarranty !== undefined) { this.impliedWarranty = d.impliedWarranty }
			if (d.returnPolicy !== undefined) { this.returnPolicy = d.returnPolicy }
			if (d.warranty !== undefined) { this.warranty = d.warranty }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AfterSalesServices ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #wholesalePriceList
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList}
  **/
get wholesalePriceList () { return this.#wholesalePriceList }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList}
  **/
set wholesalePriceList (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList) {
			this.#wholesalePriceList = value
		} else {
			this.#wholesalePriceList = new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList(value)
		}
}
setWholesalePriceList (value) {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts.WholesalePriceList(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
			if (d.wholesalePriceList !== undefined) { this.wholesalePriceList = d.wholesalePriceList }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Discounts, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Discounts ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #sections  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections}
  **/
set sections (value) {
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
setSections (value) {
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
 #items  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items}
  **/
set items (value) {
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
setItems (value) {
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
 #type  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value) {
		this.#type = String(value);
}
setType (value) {
	this.type = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections.Items(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Description.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description.Sections(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Description ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #quantity
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity}
  **/
get quantity () { return this.#quantity }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity}
  **/
set quantity (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity) {
			this.#quantity = value
		} else {
			this.#quantity = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity(value)
		}
}
setQuantity (value) {
	this.quantity = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product}
  **/
 #product
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product}
  **/
get product () { return this.#product }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product}
  **/
set product (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product) {
			this.#product = value
		} else {
			this.#product = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product(value)
		}
}
setProduct (value) {
	this.product = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson}
  **/
 #responsiblePerson
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson}
  **/
get responsiblePerson () { return this.#responsiblePerson }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson}
  **/
set responsiblePerson (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson) {
			this.#responsiblePerson = value
		} else {
			this.#responsiblePerson = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson(value)
		}
}
setResponsiblePerson (value) {
	this.responsiblePerson = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer}
  **/
 #responsibleProducer
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer}
  **/
get responsibleProducer () { return this.#responsibleProducer }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer}
  **/
set responsibleProducer (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer) {
			this.#responsibleProducer = value
		} else {
			this.#responsibleProducer = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer(value)
		}
}
setResponsibleProducer (value) {
	this.responsibleProducer = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation}
  **/
 #safetyInformation
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation}
  **/
set safetyInformation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation(value)
		}
}
setSafetyInformation (value) {
	this.safetyInformation = value
	return this
}
		/**
  * 
  * @type {boolean}
  **/
 #marketedBeforeGPSRObligation
		/**
  * 
  * @returns {boolean}
  **/
get marketedBeforeGPSRObligation () { return this.#marketedBeforeGPSRObligation }
/**
  * 
  * @type {boolean}
  **/
set marketedBeforeGPSRObligation (value) {
		this.#marketedBeforeGPSRObligation = Boolean(value);
}
setMarketedBeforeGPSRObligation (value) {
	this.marketedBeforeGPSRObligation = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits}
  **/
 #deposits  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits}
  **/
get deposits () { return this.#deposits }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits}
  **/
set deposits (value) {
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
setDeposits (value) {
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
 #value  =  0
		/**
  * 
  * @returns {number}
  **/
get value () { return this.#value }
/**
  * 
  * @type {number}
  **/
set value (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#value = parsedValue;
		}
}
setValue (value) {
	this.value = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Quantity ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
		/**
  * 
  * @type {boolean}
  **/
 #isAiCoCreated
		/**
  * 
  * @returns {boolean}
  **/
get isAiCoCreated () { return this.#isAiCoCreated }
/**
  * 
  * @type {boolean}
  **/
set isAiCoCreated (value) {
		this.#isAiCoCreated = Boolean(value);
}
setIsAiCoCreated (value) {
	this.isAiCoCreated = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication}
  **/
 #publication
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication}
  **/
set publication (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication) {
			this.#publication = value
		} else {
			this.#publication = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication(value)
		}
}
setPublication (value) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters}
  **/
 #parameters  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters}
  **/
set parameters (value) {
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
setParameters (value) {
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
 #status  =  ""
		/**
  * 
  * @returns {string}
  **/
get status () { return this.#status }
/**
  * 
  * @type {string}
  **/
set status (value) {
		this.#status = String(value);
}
setStatus (value) {
	this.status = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Publication ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #name  =  ""
		/**
  * 
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * 
  * @type {string}
  **/
set name (value) {
		this.#name = String(value);
}
setName (value) {
	this.name = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
 #rangeValue
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
get rangeValue () { return this.#rangeValue }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
set rangeValue (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue) {
			this.#rangeValue = value
		} else {
			this.#rangeValue = new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue(value)
		}
}
setRangeValue (value) {
	this.rangeValue = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values}
  **/
 #values  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values}
  **/
get values () { return this.#values }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.Values}
  **/
set values (value) {
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
setValues (value) {
	this.values = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds}
  **/
 #valuesIds  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds}
  **/
get valuesIds () { return this.#valuesIds }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.ValuesIds}
  **/
set valuesIds (value) {
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
setValuesIds (value) {
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
 #from  =  ""
		/**
  * 
  * @returns {string}
  **/
get from () { return this.#from }
/**
  * 
  * @type {string}
  **/
set from (value) {
		this.#from = String(value);
}
setFrom (value) {
	this.from = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #to  =  ""
		/**
  * 
  * @returns {string}
  **/
get to () { return this.#to }
/**
  * 
  * @type {string}
  **/
set to (value) {
		this.#to = String(value);
}
setTo (value) {
	this.to = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters.RangeValue(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product.Parameters(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Product ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsiblePerson ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.ResponsibleProducer ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #type  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value) {
		this.#type = String(value);
}
setType (value) {
	this.type = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #description  =  ""
		/**
  * 
  * @returns {string}
  **/
get description () { return this.#description }
/**
  * 
  * @type {string}
  **/
set description (value) {
		this.#description = String(value);
}
setDescription (value) {
	this.description = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #quantity  =  0
		/**
  * 
  * @returns {number}
  **/
get quantity () { return this.#quantity }
/**
  * 
  * @type {number}
  **/
set quantity (value) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#quantity = parsedValue;
		}
}
setQuantity (value) {
	this.quantity = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet.Deposits(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.ProductSet, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.ProductSet ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Attachments(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Attachments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Attachments(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Attachments ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.FundraisingCampaign ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.AdditionalServices ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #buyableOnlyByBusiness
		/**
  * 
  * @returns {boolean}
  **/
get buyableOnlyByBusiness () { return this.#buyableOnlyByBusiness }
/**
  * 
  * @type {boolean}
  **/
set buyableOnlyByBusiness (value) {
		this.#buyableOnlyByBusiness = Boolean(value);
}
setBuyableOnlyByBusiness (value) {
	this.buyableOnlyByBusiness = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.B2b(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.B2b, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.B2b(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.B2b ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #type  =  ""
		/**
  * 
  * @returns {string}
  **/
get type () { return this.#type }
/**
  * 
  * @type {string}
  **/
set type (value) {
		this.#type = String(value);
}
setType (value) {
	this.type = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.CompatibilityList ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #validatedAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get validatedAt () { return this.#validatedAt }
/**
  * 
  * @type {string}
  **/
set validatedAt (value) {
		this.#validatedAt = String(value);
}
setValidatedAt (value) {
	this.validatedAt = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors}
  **/
 #errors  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors}
  **/
set errors (value) {
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
setErrors (value) {
	this.errors = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings}
  **/
 #warnings  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings}
  **/
get warnings () { return this.#warnings }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings}
  **/
set warnings (value) {
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
setWarnings (value) {
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
 #code  =  ""
		/**
  * 
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * 
  * @type {string}
  **/
set code (value) {
		this.#code = String(value);
}
setCode (value) {
	this.code = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #details  =  ""
		/**
  * 
  * @returns {string}
  **/
get details () { return this.#details }
/**
  * 
  * @type {string}
  **/
set details (value) {
		this.#details = String(value);
}
setDetails (value) {
	this.details = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #message  =  ""
		/**
  * 
  * @returns {string}
  **/
get message () { return this.#message }
/**
  * 
  * @type {string}
  **/
set message (value) {
		this.#message = String(value);
}
setMessage (value) {
	this.message = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #path  =  ""
		/**
  * 
  * @returns {string}
  **/
get path () { return this.#path }
/**
  * 
  * @type {string}
  **/
set path (value) {
		this.#path = String(value);
}
setPath (value) {
	this.path = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #userMessage  =  ""
		/**
  * 
  * @returns {string}
  **/
get userMessage () { return this.#userMessage }
/**
  * 
  * @type {string}
  **/
set userMessage (value) {
		this.#userMessage = String(value);
}
setUserMessage (value) {
	this.userMessage = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata}
  **/
 #metadata
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata}
  **/
set metadata (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata(value)
		}
}
setMetadata (value) {
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
 #productId  =  ""
		/**
  * 
  * @returns {string}
  **/
get productId () { return this.#productId }
/**
  * 
  * @type {string}
  **/
set productId (value) {
		this.#productId = String(value);
}
setProductId (value) {
	this.productId = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors.Metadata(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Errors ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #code  =  ""
		/**
  * 
  * @returns {string}
  **/
get code () { return this.#code }
/**
  * 
  * @type {string}
  **/
set code (value) {
		this.#code = String(value);
}
setCode (value) {
	this.code = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #details  =  ""
		/**
  * 
  * @returns {string}
  **/
get details () { return this.#details }
/**
  * 
  * @type {string}
  **/
set details (value) {
		this.#details = String(value);
}
setDetails (value) {
	this.details = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #message  =  ""
		/**
  * 
  * @returns {string}
  **/
get message () { return this.#message }
/**
  * 
  * @type {string}
  **/
set message (value) {
		this.#message = String(value);
}
setMessage (value) {
	this.message = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #path  =  ""
		/**
  * 
  * @returns {string}
  **/
get path () { return this.#path }
/**
  * 
  * @type {string}
  **/
set path (value) {
		this.#path = String(value);
}
setPath (value) {
	this.path = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #userMessage  =  ""
		/**
  * 
  * @returns {string}
  **/
get userMessage () { return this.#userMessage }
/**
  * 
  * @type {string}
  **/
set userMessage (value) {
		this.#userMessage = String(value);
}
setUserMessage (value) {
	this.userMessage = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata}
  **/
 #metadata
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata}
  **/
set metadata (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata(value)
		}
}
setMetadata (value) {
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
 #productId  =  ""
		/**
  * 
  * @returns {string}
  **/
get productId () { return this.#productId }
/**
  * 
  * @type {string}
  **/
set productId (value) {
		this.#productId = String(value);
}
setProductId (value) {
	this.productId = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings.Metadata(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation.Warnings(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.Validation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.Validation ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.External(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.External, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.External(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.External ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #id  =  ""
		/**
  * 
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * 
  * @type {string}
  **/
set id (value) {
		this.#id = String(value);
}
setId (value) {
	this.id = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SizeTable(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.SizeTable, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SizeTable(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.SizeTable ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #subject  =  ""
		/**
  * 
  * @returns {string}
  **/
get subject () { return this.#subject }
/**
  * 
  * @type {string}
  **/
set subject (value) {
		this.#subject = String(value);
}
setSubject (value) {
	this.subject = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #exemption  =  ""
		/**
  * 
  * @returns {string}
  **/
get exemption () { return this.#exemption }
/**
  * 
  * @type {string}
  **/
set exemption (value) {
		this.#exemption = String(value);
}
setExemption (value) {
	this.exemption = value
	return this
}
		/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates}
  **/
 #rates  =  []
		/**
  * 
  * @returns {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates}
  **/
get rates () { return this.#rates }
/**
  * 
  * @type {GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates}
  **/
set rates (value) {
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
setRates (value) {
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
 #rate  =  ""
		/**
  * 
  * @returns {string}
  **/
get rate () { return this.#rate }
/**
  * 
  * @type {string}
  **/
set rate (value) {
		this.#rate = String(value);
}
setRate (value) {
	this.rate = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #countryCode  =  ""
		/**
  * 
  * @returns {string}
  **/
get countryCode () { return this.#countryCode }
/**
  * 
  * @type {string}
  **/
set countryCode (value) {
		this.#countryCode = String(value);
}
setCountryCode (value) {
	this.countryCode = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings.Rates(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.TaxSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.TaxSettings ({ ...this.toJSON(), ...partial });
	}
	clone() {
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
 #mode  =  ""
		/**
  * 
  * @returns {string}
  **/
get mode () { return this.#mode }
/**
  * 
  * @type {string}
  **/
set mode (value) {
		this.#mode = String(value);
}
setMode (value) {
	this.mode = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #hint  =  ""
		/**
  * 
  * @returns {string}
  **/
get hint () { return this.#hint }
/**
  * 
  * @type {string}
  **/
set hint (value) {
		this.#hint = String(value);
}
setHint (value) {
	this.hint = value
	return this
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes.MessageToSellerSettings(this.toJSON());
	}
}
	constructor(data) {
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
	#isJsonAppliable(obj) {
		const g = globalThis
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
		const d = data;
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
		const d = data;
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
	static from(possibleDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetAllDataOfTheParticularProductOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new GetAllDataOfTheParticularProductOfferActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new GetAllDataOfTheParticularProductOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new GetAllDataOfTheParticularProductOfferActionRes(this.toJSON());
	}
}