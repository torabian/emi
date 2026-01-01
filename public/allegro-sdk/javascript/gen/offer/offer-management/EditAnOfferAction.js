import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Edit an offer
*/
	/**
 * EditAnOfferAction
 */
export class EditAnOfferAction { //
  static URL = 'https://api.{environment}/sale/product-offers/{offerId}';
  static NewUrl = (
	qs
  ) => buildUrl(
		EditAnOfferAction.URL,
		 undefined,
		qs
	);
  static Method = 'patch';
	static Fetch$ = async (
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
			overrideUrl ?? EditAnOfferAction.NewUrl(
				qs
			),
			{
				method: EditAnOfferAction.Method,
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
				creatorFn: (item) => new EditAnOfferActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new EditAnOfferActionRes(item))
		const res = await EditAnOfferAction.Fetch$(
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
  "name": "Edit an offer",
  "url": "https://api.{environment}/sale/product-offers/{offerId}",
  "method": "patch",
  "description": "Use this resource to edit offer. This resource allows you to edit each field independently, so use it if you want to change only, for example, the price or the quantity in an offer. Read more: PL / EN. Note that requests may be limited.",
  "in": {
    "fields": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "language",
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
        "type": "map",
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
  },
  "out": {
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
        "name": "language",
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
                    "type": "slice",
                    "primitive": "string"
                  },
                  {
                    "name": "valuesIds",
                    "type": "slice",
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
        "name": "additionalMarketplaces",
        "type": "object",
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
                        "type": "slice",
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
        "name": "warnings",
        "type": "slice",
        "primitive": "string"
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
      },
      {
        "name": "createdAt",
        "type": "string"
      },
      {
        "name": "updatedAt",
        "type": "string"
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
      }
    ]
  }
}
}
/**
  * The base class definition for editAnOfferActionReq
  **/
export class EditAnOfferActionReq {
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
  * @type {string}
  **/
 #language  =  ""
		/**
  * 
  * @returns {string}
  **/
get language () { return this.#language }
/**
  * 
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
  * 
  * @type {EditAnOfferActionReq.Category}
  **/
 #category
		/**
  * 
  * @returns {EditAnOfferActionReq.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {EditAnOfferActionReq.Category}
  **/
set category (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Category) {
			this.#category = value
		} else {
			this.#category = new EditAnOfferActionReq.Category(value)
		}
}
setCategory (value) {
	this.category = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet}
  **/
 #productSet  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet}
  **/
get productSet () { return this.#productSet }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet}
  **/
set productSet (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionReq.ProductSet) {
			this.#productSet = value
		} else {
			this.#productSet = value.map(item => new EditAnOfferActionReq.ProductSet(item))
		}
}
setProductSet (value) {
	this.productSet = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Stock}
  **/
 #stock
		/**
  * 
  * @returns {EditAnOfferActionReq.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {EditAnOfferActionReq.Stock}
  **/
set stock (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Stock) {
			this.#stock = value
		} else {
			this.#stock = new EditAnOfferActionReq.Stock(value)
		}
}
setStock (value) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.SellingMode}
  **/
 #sellingMode
		/**
  * 
  * @returns {EditAnOfferActionReq.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {EditAnOfferActionReq.SellingMode}
  **/
set sellingMode (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new EditAnOfferActionReq.SellingMode(value)
		}
}
setSellingMode (value) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Payments}
  **/
 #payments
		/**
  * 
  * @returns {EditAnOfferActionReq.Payments}
  **/
get payments () { return this.#payments }
/**
  * 
  * @type {EditAnOfferActionReq.Payments}
  **/
set payments (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Payments) {
			this.#payments = value
		} else {
			this.#payments = new EditAnOfferActionReq.Payments(value)
		}
}
setPayments (value) {
	this.payments = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Delivery}
  **/
 #delivery
		/**
  * 
  * @returns {EditAnOfferActionReq.Delivery}
  **/
get delivery () { return this.#delivery }
/**
  * 
  * @type {EditAnOfferActionReq.Delivery}
  **/
set delivery (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Delivery) {
			this.#delivery = value
		} else {
			this.#delivery = new EditAnOfferActionReq.Delivery(value)
		}
}
setDelivery (value) {
	this.delivery = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Publication}
  **/
 #publication
		/**
  * 
  * @returns {EditAnOfferActionReq.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {EditAnOfferActionReq.Publication}
  **/
set publication (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Publication) {
			this.#publication = value
		} else {
			this.#publication = new EditAnOfferActionReq.Publication(value)
		}
}
setPublication (value) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {{[key: string]: any}}
  **/
 #additionalMarketplaces
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
  * @type {EditAnOfferActionReq.CompatibilityList}
  **/
 #compatibilityList
		/**
  * 
  * @returns {EditAnOfferActionReq.CompatibilityList}
  **/
get compatibilityList () { return this.#compatibilityList }
/**
  * 
  * @type {EditAnOfferActionReq.CompatibilityList}
  **/
set compatibilityList (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.CompatibilityList) {
			this.#compatibilityList = value
		} else {
			this.#compatibilityList = new EditAnOfferActionReq.CompatibilityList(value)
		}
}
setCompatibilityList (value) {
	this.compatibilityList = value
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
  * @type {EditAnOfferActionReq.Description}
  **/
 #description
		/**
  * 
  * @returns {EditAnOfferActionReq.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {EditAnOfferActionReq.Description}
  **/
set description (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Description) {
			this.#description = value
		} else {
			this.#description = new EditAnOfferActionReq.Description(value)
		}
}
setDescription (value) {
	this.description = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.B2b}
  **/
 #b2b
		/**
  * 
  * @returns {EditAnOfferActionReq.B2b}
  **/
get b2b () { return this.#b2b }
/**
  * 
  * @type {EditAnOfferActionReq.B2b}
  **/
set b2b (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.B2b) {
			this.#b2b = value
		} else {
			this.#b2b = new EditAnOfferActionReq.B2b(value)
		}
}
setB2b (value) {
	this.b2b = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Attachments}
  **/
 #attachments  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.Attachments}
  **/
get attachments () { return this.#attachments }
/**
  * 
  * @type {EditAnOfferActionReq.Attachments}
  **/
set attachments (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionReq.Attachments) {
			this.#attachments = value
		} else {
			this.#attachments = value.map(item => new EditAnOfferActionReq.Attachments(item))
		}
}
setAttachments (value) {
	this.attachments = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.FundraisingCampaign}
  **/
 #fundraisingCampaign
		/**
  * 
  * @returns {EditAnOfferActionReq.FundraisingCampaign}
  **/
get fundraisingCampaign () { return this.#fundraisingCampaign }
/**
  * 
  * @type {EditAnOfferActionReq.FundraisingCampaign}
  **/
set fundraisingCampaign (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.FundraisingCampaign) {
			this.#fundraisingCampaign = value
		} else {
			this.#fundraisingCampaign = new EditAnOfferActionReq.FundraisingCampaign(value)
		}
}
setFundraisingCampaign (value) {
	this.fundraisingCampaign = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.AdditionalServices}
  **/
 #additionalServices
		/**
  * 
  * @returns {EditAnOfferActionReq.AdditionalServices}
  **/
get additionalServices () { return this.#additionalServices }
/**
  * 
  * @type {EditAnOfferActionReq.AdditionalServices}
  **/
set additionalServices (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AdditionalServices) {
			this.#additionalServices = value
		} else {
			this.#additionalServices = new EditAnOfferActionReq.AdditionalServices(value)
		}
}
setAdditionalServices (value) {
	this.additionalServices = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices}
  **/
 #afterSalesServices
		/**
  * 
  * @returns {EditAnOfferActionReq.AfterSalesServices}
  **/
get afterSalesServices () { return this.#afterSalesServices }
/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices}
  **/
set afterSalesServices (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AfterSalesServices) {
			this.#afterSalesServices = value
		} else {
			this.#afterSalesServices = new EditAnOfferActionReq.AfterSalesServices(value)
		}
}
setAfterSalesServices (value) {
	this.afterSalesServices = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.SizeTable}
  **/
 #sizeTable
		/**
  * 
  * @returns {EditAnOfferActionReq.SizeTable}
  **/
get sizeTable () { return this.#sizeTable }
/**
  * 
  * @type {EditAnOfferActionReq.SizeTable}
  **/
set sizeTable (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SizeTable) {
			this.#sizeTable = value
		} else {
			this.#sizeTable = new EditAnOfferActionReq.SizeTable(value)
		}
}
setSizeTable (value) {
	this.sizeTable = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Contact}
  **/
 #contact
		/**
  * 
  * @returns {EditAnOfferActionReq.Contact}
  **/
get contact () { return this.#contact }
/**
  * 
  * @type {EditAnOfferActionReq.Contact}
  **/
set contact (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Contact) {
			this.#contact = value
		} else {
			this.#contact = new EditAnOfferActionReq.Contact(value)
		}
}
setContact (value) {
	this.contact = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Discounts}
  **/
 #discounts
		/**
  * 
  * @returns {EditAnOfferActionReq.Discounts}
  **/
get discounts () { return this.#discounts }
/**
  * 
  * @type {EditAnOfferActionReq.Discounts}
  **/
set discounts (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Discounts) {
			this.#discounts = value
		} else {
			this.#discounts = new EditAnOfferActionReq.Discounts(value)
		}
}
setDiscounts (value) {
	this.discounts = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Location}
  **/
 #location
		/**
  * 
  * @returns {EditAnOfferActionReq.Location}
  **/
get location () { return this.#location }
/**
  * 
  * @type {EditAnOfferActionReq.Location}
  **/
set location (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Location) {
			this.#location = value
		} else {
			this.#location = new EditAnOfferActionReq.Location(value)
		}
}
setLocation (value) {
	this.location = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.External}
  **/
 #external
		/**
  * 
  * @returns {EditAnOfferActionReq.External}
  **/
get external () { return this.#external }
/**
  * 
  * @type {EditAnOfferActionReq.External}
  **/
set external (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.External) {
			this.#external = value
		} else {
			this.#external = new EditAnOfferActionReq.External(value)
		}
}
setExternal (value) {
	this.external = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.TaxSettings}
  **/
 #taxSettings
		/**
  * 
  * @returns {EditAnOfferActionReq.TaxSettings}
  **/
get taxSettings () { return this.#taxSettings }
/**
  * 
  * @type {EditAnOfferActionReq.TaxSettings}
  **/
set taxSettings (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.TaxSettings) {
			this.#taxSettings = value
		} else {
			this.#taxSettings = new EditAnOfferActionReq.TaxSettings(value)
		}
}
setTaxSettings (value) {
	this.taxSettings = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.MessageToSellerSettings}
  **/
 #messageToSellerSettings
		/**
  * 
  * @returns {EditAnOfferActionReq.MessageToSellerSettings}
  **/
get messageToSellerSettings () { return this.#messageToSellerSettings }
/**
  * 
  * @type {EditAnOfferActionReq.MessageToSellerSettings}
  **/
set messageToSellerSettings (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.MessageToSellerSettings) {
			this.#messageToSellerSettings = value
		} else {
			this.#messageToSellerSettings = new EditAnOfferActionReq.MessageToSellerSettings(value)
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
	* Creates an instance of EditAnOfferActionReq.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Category(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Category ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Category(this.toJSON());
	}
}
/**
  * The base class definition for productSet
  **/
static ProductSet = class ProductSet {
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product}
  **/
 #product
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Product}
  **/
get product () { return this.#product }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product}
  **/
set product (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.Product) {
			this.#product = value
		} else {
			this.#product = new EditAnOfferActionReq.ProductSet.Product(value)
		}
}
setProduct (value) {
	this.product = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Quantity}
  **/
 #quantity
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Quantity}
  **/
get quantity () { return this.#quantity }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Quantity}
  **/
set quantity (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.Quantity) {
			this.#quantity = value
		} else {
			this.#quantity = new EditAnOfferActionReq.ProductSet.Quantity(value)
		}
}
setQuantity (value) {
	this.quantity = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.ResponsiblePerson}
  **/
 #responsiblePerson
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.ResponsiblePerson}
  **/
get responsiblePerson () { return this.#responsiblePerson }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.ResponsiblePerson}
  **/
set responsiblePerson (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.ResponsiblePerson) {
			this.#responsiblePerson = value
		} else {
			this.#responsiblePerson = new EditAnOfferActionReq.ProductSet.ResponsiblePerson(value)
		}
}
setResponsiblePerson (value) {
	this.responsiblePerson = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.ResponsibleProducer}
  **/
 #responsibleProducer
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.ResponsibleProducer}
  **/
get responsibleProducer () { return this.#responsibleProducer }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.ResponsibleProducer}
  **/
set responsibleProducer (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.ResponsibleProducer) {
			this.#responsibleProducer = value
		} else {
			this.#responsibleProducer = new EditAnOfferActionReq.ProductSet.ResponsibleProducer(value)
		}
}
setResponsibleProducer (value) {
	this.responsibleProducer = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.SafetyInformation}
  **/
 #safetyInformation
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.SafetyInformation}
  **/
set safetyInformation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new EditAnOfferActionReq.ProductSet.SafetyInformation(value)
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
  * @type {EditAnOfferActionReq.ProductSet.Deposits}
  **/
 #deposits  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Deposits}
  **/
get deposits () { return this.#deposits }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Deposits}
  **/
set deposits (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionReq.ProductSet.Deposits) {
			this.#deposits = value
		} else {
			this.#deposits = value.map(item => new EditAnOfferActionReq.ProductSet.Deposits(item))
		}
}
setDeposits (value) {
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
 #idType  =  ""
		/**
  * 
  * @returns {string}
  **/
get idType () { return this.#idType }
/**
  * 
  * @type {string}
  **/
set idType (value) {
		this.#idType = String(value);
}
setIdType (value) {
	this.idType = value
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
  * @type {EditAnOfferActionReq.ProductSet.Product.Category}
  **/
 #category
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Product.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product.Category}
  **/
set category (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.Product.Category) {
			this.#category = value
		} else {
			this.#category = new EditAnOfferActionReq.ProductSet.Product.Category(value)
		}
}
setCategory (value) {
	this.category = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product.Parameters}
  **/
 #parameters  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Product.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product.Parameters}
  **/
set parameters (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionReq.ProductSet.Product.Parameters) {
			this.#parameters = value
		} else {
			this.#parameters = value.map(item => new EditAnOfferActionReq.ProductSet.Product.Parameters(item))
		}
}
setParameters (value) {
	this.parameters = value
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Product.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Product.Category(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet.Product.Category ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet.Product.Category(this.toJSON());
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
  * @type {EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue}
  **/
 #rangeValue
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue}
  **/
get rangeValue () { return this.#rangeValue }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue}
  **/
set rangeValue (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue) {
			this.#rangeValue = value
		} else {
			this.#rangeValue = new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue(value)
		}
}
setRangeValue (value) {
	this.rangeValue = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #values  =  []
		/**
  * 
  * @returns {string[]}
  **/
get values () { return this.#values }
/**
  * 
  * @type {string[]}
  **/
set values (value) {
}
setValues (value) {
	this.values = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #valuesIds  =  []
		/**
  * 
  * @returns {string[]}
  **/
get valuesIds () { return this.#valuesIds }
/**
  * 
  * @type {string[]}
  **/
set valuesIds (value) {
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue(this.toJSON());
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
			if (!(d.rangeValue instanceof EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue)) { this.rangeValue = new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue(d.rangeValue || {}) }	
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
						EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue.Fields
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Parameters, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters(this.toJSON());
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
		const d = data;
			if (!(d.category instanceof EditAnOfferActionReq.ProductSet.Product.Category)) { this.category = new EditAnOfferActionReq.ProductSet.Product.Category(d.category || {}) }	
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
						EditAnOfferActionReq.ProductSet.Product.Category.Fields
						);
						},
			parameters$: 'parameters',
get parameters() {
					return withPrefix(
						"productSet.product.parameters[:i]",
						EditAnOfferActionReq.ProductSet.Product.Parameters.Fields
						);
						},
			images$: 'images',
get images() {
					return "productSet.product.images[:i]";
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Product(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Product(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet.Product ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet.Product(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.Quantity, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Quantity(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Quantity, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Quantity(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet.Quantity ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet.Quantity(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.ResponsiblePerson, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet.ResponsiblePerson(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.ResponsiblePerson, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet.ResponsiblePerson(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet.ResponsiblePerson ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet.ResponsiblePerson(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.ResponsibleProducer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet.ResponsibleProducer(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.ResponsibleProducer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet.ResponsibleProducer(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet.ResponsibleProducer ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet.ResponsibleProducer(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.SafetyInformation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet.SafetyInformation(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet.SafetyInformation(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.Deposits, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Deposits(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Deposits, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet.Deposits(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet.Deposits ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet.Deposits(this.toJSON());
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
		const d = data;
			if (!(d.product instanceof EditAnOfferActionReq.ProductSet.Product)) { this.product = new EditAnOfferActionReq.ProductSet.Product(d.product || {}) }	
			if (!(d.quantity instanceof EditAnOfferActionReq.ProductSet.Quantity)) { this.quantity = new EditAnOfferActionReq.ProductSet.Quantity(d.quantity || {}) }	
			if (!(d.responsiblePerson instanceof EditAnOfferActionReq.ProductSet.ResponsiblePerson)) { this.responsiblePerson = new EditAnOfferActionReq.ProductSet.ResponsiblePerson(d.responsiblePerson || {}) }	
			if (!(d.responsibleProducer instanceof EditAnOfferActionReq.ProductSet.ResponsibleProducer)) { this.responsibleProducer = new EditAnOfferActionReq.ProductSet.ResponsibleProducer(d.responsibleProducer || {}) }	
			if (!(d.safetyInformation instanceof EditAnOfferActionReq.ProductSet.SafetyInformation)) { this.safetyInformation = new EditAnOfferActionReq.ProductSet.SafetyInformation(d.safetyInformation || {}) }	
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
						EditAnOfferActionReq.ProductSet.Product.Fields
						);
						},
			quantity$: 'quantity',
get quantity() {
					return withPrefix(
						"productSet.quantity",
						EditAnOfferActionReq.ProductSet.Quantity.Fields
						);
						},
			responsiblePerson$: 'responsiblePerson',
get responsiblePerson() {
					return withPrefix(
						"productSet.responsiblePerson",
						EditAnOfferActionReq.ProductSet.ResponsiblePerson.Fields
						);
						},
			responsibleProducer$: 'responsibleProducer',
get responsibleProducer() {
					return withPrefix(
						"productSet.responsibleProducer",
						EditAnOfferActionReq.ProductSet.ResponsibleProducer.Fields
						);
						},
			safetyInformation$: 'safetyInformation',
get safetyInformation() {
					return withPrefix(
						"productSet.safetyInformation",
						EditAnOfferActionReq.ProductSet.SafetyInformation.Fields
						);
						},
			marketedBeforeGPSRObligation: 'marketedBeforeGPSRObligation',
			deposits$: 'deposits',
get deposits() {
					return withPrefix(
						"productSet.deposits[:i]",
						EditAnOfferActionReq.ProductSet.Deposits.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.ProductSet(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.ProductSet(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.ProductSet ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.ProductSet(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.Stock, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Stock(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Stock ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Stock(this.toJSON());
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
  * @type {EditAnOfferActionReq.SellingMode.Price}
  **/
 #price
		/**
  * 
  * @returns {EditAnOfferActionReq.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.Price}
  **/
set price (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new EditAnOfferActionReq.SellingMode.Price(value)
		}
}
setPrice (value) {
	this.price = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.MinimalPrice}
  **/
 #minimalPrice
		/**
  * 
  * @returns {EditAnOfferActionReq.SellingMode.MinimalPrice}
  **/
get minimalPrice () { return this.#minimalPrice }
/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.MinimalPrice}
  **/
set minimalPrice (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SellingMode.MinimalPrice) {
			this.#minimalPrice = value
		} else {
			this.#minimalPrice = new EditAnOfferActionReq.SellingMode.MinimalPrice(value)
		}
}
setMinimalPrice (value) {
	this.minimalPrice = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.StartingPrice}
  **/
 #startingPrice
		/**
  * 
  * @returns {EditAnOfferActionReq.SellingMode.StartingPrice}
  **/
get startingPrice () { return this.#startingPrice }
/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.StartingPrice}
  **/
set startingPrice (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SellingMode.StartingPrice) {
			this.#startingPrice = value
		} else {
			this.#startingPrice = new EditAnOfferActionReq.SellingMode.StartingPrice(value)
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
	* Creates an instance of EditAnOfferActionReq.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.SellingMode.Price(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.SellingMode.MinimalPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.SellingMode.MinimalPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SellingMode.MinimalPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.SellingMode.MinimalPrice(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.SellingMode.MinimalPrice ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.SellingMode.MinimalPrice(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.SellingMode.StartingPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.SellingMode.StartingPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SellingMode.StartingPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.SellingMode.StartingPrice(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.SellingMode.StartingPrice ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.SellingMode.StartingPrice(this.toJSON());
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
			if (!(d.price instanceof EditAnOfferActionReq.SellingMode.Price)) { this.price = new EditAnOfferActionReq.SellingMode.Price(d.price || {}) }	
			if (!(d.minimalPrice instanceof EditAnOfferActionReq.SellingMode.MinimalPrice)) { this.minimalPrice = new EditAnOfferActionReq.SellingMode.MinimalPrice(d.minimalPrice || {}) }	
			if (!(d.startingPrice instanceof EditAnOfferActionReq.SellingMode.StartingPrice)) { this.startingPrice = new EditAnOfferActionReq.SellingMode.StartingPrice(d.startingPrice || {}) }	
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
						EditAnOfferActionReq.SellingMode.Price.Fields
						);
						},
			minimalPrice$: 'minimalPrice',
get minimalPrice() {
					return withPrefix(
						"sellingMode.minimalPrice",
						EditAnOfferActionReq.SellingMode.MinimalPrice.Fields
						);
						},
			startingPrice$: 'startingPrice',
get startingPrice() {
					return withPrefix(
						"sellingMode.startingPrice",
						EditAnOfferActionReq.SellingMode.StartingPrice.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SellingMode, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.SellingMode(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.SellingMode(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.Payments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Payments(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Payments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Payments(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Payments ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Payments(this.toJSON());
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
  * @type {EditAnOfferActionReq.Delivery.ShippingRates}
  **/
 #shippingRates
		/**
  * 
  * @returns {EditAnOfferActionReq.Delivery.ShippingRates}
  **/
get shippingRates () { return this.#shippingRates }
/**
  * 
  * @type {EditAnOfferActionReq.Delivery.ShippingRates}
  **/
set shippingRates (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Delivery.ShippingRates) {
			this.#shippingRates = value
		} else {
			this.#shippingRates = new EditAnOfferActionReq.Delivery.ShippingRates(value)
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
	* Creates an instance of EditAnOfferActionReq.Delivery.ShippingRates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Delivery.ShippingRates(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Delivery.ShippingRates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Delivery.ShippingRates(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Delivery.ShippingRates ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Delivery.ShippingRates(this.toJSON());
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
			if (!(d.shippingRates instanceof EditAnOfferActionReq.Delivery.ShippingRates)) { this.shippingRates = new EditAnOfferActionReq.Delivery.ShippingRates(d.shippingRates || {}) }	
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
						EditAnOfferActionReq.Delivery.ShippingRates.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Delivery, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Delivery(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Delivery, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Delivery(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Delivery ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Delivery(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Publication(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Publication ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Publication(this.toJSON());
	}
}
/**
  * The base class definition for compatibilityList
  **/
static CompatibilityList = class CompatibilityList {
		/**
  * 
  * @type {EditAnOfferActionReq.CompatibilityList.Items}
  **/
 #items  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.CompatibilityList.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {EditAnOfferActionReq.CompatibilityList.Items}
  **/
set items (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionReq.CompatibilityList.Items) {
			this.#items = value
		} else {
			this.#items = value.map(item => new EditAnOfferActionReq.CompatibilityList.Items(item))
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
		/**
  * 
  * @type {string}
  **/
 #text  =  ""
		/**
  * 
  * @returns {string}
  **/
get text () { return this.#text }
/**
  * 
  * @type {string}
  **/
set text (value) {
		this.#text = String(value);
}
setText (value) {
	this.text = value
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
	* Creates an instance of EditAnOfferActionReq.CompatibilityList.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.CompatibilityList.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.CompatibilityList.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.CompatibilityList.Items(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.CompatibilityList.Items ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.CompatibilityList.Items(this.toJSON());
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
						"compatibilityList.items[:i]",
						EditAnOfferActionReq.CompatibilityList.Items.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.CompatibilityList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.CompatibilityList(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.CompatibilityList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.CompatibilityList(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.CompatibilityList ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.CompatibilityList(this.toJSON());
	}
}
/**
  * The base class definition for description
  **/
static Description = class Description {
		/**
  * 
  * @type {EditAnOfferActionReq.Description.Sections}
  **/
 #sections  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.Description.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {EditAnOfferActionReq.Description.Sections}
  **/
set sections (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionReq.Description.Sections) {
			this.#sections = value
		} else {
			this.#sections = value.map(item => new EditAnOfferActionReq.Description.Sections(item))
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
  * @type {EditAnOfferActionReq.Description.Sections.Items}
  **/
 #items  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.Description.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {EditAnOfferActionReq.Description.Sections.Items}
  **/
set items (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionReq.Description.Sections.Items) {
			this.#items = value
		} else {
			this.#items = value.map(item => new EditAnOfferActionReq.Description.Sections.Items(item))
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
	* Creates an instance of EditAnOfferActionReq.Description.Sections.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Description.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Description.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Description.Sections.Items(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Description.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Description.Sections.Items(this.toJSON());
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
						EditAnOfferActionReq.Description.Sections.Items.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Description.Sections, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Description.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Description.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Description.Sections(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Description.Sections ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Description.Sections(this.toJSON());
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
						EditAnOfferActionReq.Description.Sections.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Description, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Description(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Description ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Description(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.B2b, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.B2b(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.B2b, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.B2b(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.B2b ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.B2b(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.Attachments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Attachments(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Attachments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Attachments(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Attachments ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Attachments(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.FundraisingCampaign, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.FundraisingCampaign(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.FundraisingCampaign, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.FundraisingCampaign(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.FundraisingCampaign ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.FundraisingCampaign(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.AdditionalServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.AdditionalServices(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AdditionalServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.AdditionalServices(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.AdditionalServices ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.AdditionalServices(this.toJSON());
	}
}
/**
  * The base class definition for afterSalesServices
  **/
static AfterSalesServices = class AfterSalesServices {
		/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty}
  **/
 #impliedWarranty
		/**
  * 
  * @returns {EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty}
  **/
get impliedWarranty () { return this.#impliedWarranty }
/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty}
  **/
set impliedWarranty (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty) {
			this.#impliedWarranty = value
		} else {
			this.#impliedWarranty = new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty(value)
		}
}
setImpliedWarranty (value) {
	this.impliedWarranty = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.ReturnPolicy}
  **/
 #returnPolicy
		/**
  * 
  * @returns {EditAnOfferActionReq.AfterSalesServices.ReturnPolicy}
  **/
get returnPolicy () { return this.#returnPolicy }
/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.ReturnPolicy}
  **/
set returnPolicy (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AfterSalesServices.ReturnPolicy) {
			this.#returnPolicy = value
		} else {
			this.#returnPolicy = new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy(value)
		}
}
setReturnPolicy (value) {
	this.returnPolicy = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.Warranty}
  **/
 #warranty
		/**
  * 
  * @returns {EditAnOfferActionReq.AfterSalesServices.Warranty}
  **/
get warranty () { return this.#warranty }
/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.Warranty}
  **/
set warranty (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AfterSalesServices.Warranty) {
			this.#warranty = value
		} else {
			this.#warranty = new EditAnOfferActionReq.AfterSalesServices.Warranty(value)
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
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.ReturnPolicy, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.ReturnPolicy, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.Warranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.AfterSalesServices.Warranty(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.Warranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.AfterSalesServices.Warranty(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.AfterSalesServices.Warranty ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.AfterSalesServices.Warranty(this.toJSON());
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
			if (!(d.impliedWarranty instanceof EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty)) { this.impliedWarranty = new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty(d.impliedWarranty || {}) }	
			if (!(d.returnPolicy instanceof EditAnOfferActionReq.AfterSalesServices.ReturnPolicy)) { this.returnPolicy = new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy(d.returnPolicy || {}) }	
			if (!(d.warranty instanceof EditAnOfferActionReq.AfterSalesServices.Warranty)) { this.warranty = new EditAnOfferActionReq.AfterSalesServices.Warranty(d.warranty || {}) }	
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
						EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty.Fields
						);
						},
			returnPolicy$: 'returnPolicy',
get returnPolicy() {
					return withPrefix(
						"afterSalesServices.returnPolicy",
						EditAnOfferActionReq.AfterSalesServices.ReturnPolicy.Fields
						);
						},
			warranty$: 'warranty',
get warranty() {
					return withPrefix(
						"afterSalesServices.warranty",
						EditAnOfferActionReq.AfterSalesServices.Warranty.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.AfterSalesServices(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.AfterSalesServices(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.AfterSalesServices ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.AfterSalesServices(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.SizeTable, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.SizeTable(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SizeTable, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.SizeTable(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.SizeTable ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.SizeTable(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.Contact, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Contact(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Contact, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Contact(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Contact ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Contact(this.toJSON());
	}
}
/**
  * The base class definition for discounts
  **/
static Discounts = class Discounts {
		/**
  * 
  * @type {EditAnOfferActionReq.Discounts.WholesalePriceList}
  **/
 #wholesalePriceList
		/**
  * 
  * @returns {EditAnOfferActionReq.Discounts.WholesalePriceList}
  **/
get wholesalePriceList () { return this.#wholesalePriceList }
/**
  * 
  * @type {EditAnOfferActionReq.Discounts.WholesalePriceList}
  **/
set wholesalePriceList (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Discounts.WholesalePriceList) {
			this.#wholesalePriceList = value
		} else {
			this.#wholesalePriceList = new EditAnOfferActionReq.Discounts.WholesalePriceList(value)
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
	* Creates an instance of EditAnOfferActionReq.Discounts.WholesalePriceList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Discounts.WholesalePriceList(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Discounts.WholesalePriceList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Discounts.WholesalePriceList(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Discounts.WholesalePriceList ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Discounts.WholesalePriceList(this.toJSON());
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
			if (!(d.wholesalePriceList instanceof EditAnOfferActionReq.Discounts.WholesalePriceList)) { this.wholesalePriceList = new EditAnOfferActionReq.Discounts.WholesalePriceList(d.wholesalePriceList || {}) }	
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
						EditAnOfferActionReq.Discounts.WholesalePriceList.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Discounts, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Discounts(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Discounts, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Discounts(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Discounts ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Discounts(this.toJSON());
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
 #city  =  ""
		/**
  * 
  * @returns {string}
  **/
get city () { return this.#city }
/**
  * 
  * @type {string}
  **/
set city (value) {
		this.#city = String(value);
}
setCity (value) {
	this.city = value
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
		/**
  * 
  * @type {string}
  **/
 #postCode  =  ""
		/**
  * 
  * @returns {string}
  **/
get postCode () { return this.#postCode }
/**
  * 
  * @type {string}
  **/
set postCode (value) {
		this.#postCode = String(value);
}
setPostCode (value) {
	this.postCode = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #province  =  ""
		/**
  * 
  * @returns {string}
  **/
get province () { return this.#province }
/**
  * 
  * @type {string}
  **/
set province (value) {
		this.#province = String(value);
}
setProvince (value) {
	this.province = value
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
	* Creates an instance of EditAnOfferActionReq.Location, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.Location(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Location, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.Location(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.Location ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.Location(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.External, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.External(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.External, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.External(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.External ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.External(this.toJSON());
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
  * @type {EditAnOfferActionReq.TaxSettings.Rates}
  **/
 #rates  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.TaxSettings.Rates}
  **/
get rates () { return this.#rates }
/**
  * 
  * @type {EditAnOfferActionReq.TaxSettings.Rates}
  **/
set rates (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionReq.TaxSettings.Rates) {
			this.#rates = value
		} else {
			this.#rates = value.map(item => new EditAnOfferActionReq.TaxSettings.Rates(item))
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
	* Creates an instance of EditAnOfferActionReq.TaxSettings.Rates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.TaxSettings.Rates(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.TaxSettings.Rates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.TaxSettings.Rates(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.TaxSettings.Rates ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.TaxSettings.Rates(this.toJSON());
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
						EditAnOfferActionReq.TaxSettings.Rates.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq.TaxSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.TaxSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.TaxSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.TaxSettings(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.TaxSettings ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.TaxSettings(this.toJSON());
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
	* Creates an instance of EditAnOfferActionReq.MessageToSellerSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq.MessageToSellerSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.MessageToSellerSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq.MessageToSellerSettings(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq.MessageToSellerSettings ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq.MessageToSellerSettings(this.toJSON());
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
		const d = data;
			if (!(d.category instanceof EditAnOfferActionReq.Category)) { this.category = new EditAnOfferActionReq.Category(d.category || {}) }	
			if (!(d.stock instanceof EditAnOfferActionReq.Stock)) { this.stock = new EditAnOfferActionReq.Stock(d.stock || {}) }	
			if (!(d.sellingMode instanceof EditAnOfferActionReq.SellingMode)) { this.sellingMode = new EditAnOfferActionReq.SellingMode(d.sellingMode || {}) }	
			if (!(d.payments instanceof EditAnOfferActionReq.Payments)) { this.payments = new EditAnOfferActionReq.Payments(d.payments || {}) }	
			if (!(d.delivery instanceof EditAnOfferActionReq.Delivery)) { this.delivery = new EditAnOfferActionReq.Delivery(d.delivery || {}) }	
			if (!(d.publication instanceof EditAnOfferActionReq.Publication)) { this.publication = new EditAnOfferActionReq.Publication(d.publication || {}) }	
			if (!(d.compatibilityList instanceof EditAnOfferActionReq.CompatibilityList)) { this.compatibilityList = new EditAnOfferActionReq.CompatibilityList(d.compatibilityList || {}) }	
			if (!(d.description instanceof EditAnOfferActionReq.Description)) { this.description = new EditAnOfferActionReq.Description(d.description || {}) }	
			if (!(d.b2b instanceof EditAnOfferActionReq.B2b)) { this.b2b = new EditAnOfferActionReq.B2b(d.b2b || {}) }	
			if (!(d.fundraisingCampaign instanceof EditAnOfferActionReq.FundraisingCampaign)) { this.fundraisingCampaign = new EditAnOfferActionReq.FundraisingCampaign(d.fundraisingCampaign || {}) }	
			if (!(d.additionalServices instanceof EditAnOfferActionReq.AdditionalServices)) { this.additionalServices = new EditAnOfferActionReq.AdditionalServices(d.additionalServices || {}) }	
			if (!(d.afterSalesServices instanceof EditAnOfferActionReq.AfterSalesServices)) { this.afterSalesServices = new EditAnOfferActionReq.AfterSalesServices(d.afterSalesServices || {}) }	
			if (!(d.sizeTable instanceof EditAnOfferActionReq.SizeTable)) { this.sizeTable = new EditAnOfferActionReq.SizeTable(d.sizeTable || {}) }	
			if (!(d.contact instanceof EditAnOfferActionReq.Contact)) { this.contact = new EditAnOfferActionReq.Contact(d.contact || {}) }	
			if (!(d.discounts instanceof EditAnOfferActionReq.Discounts)) { this.discounts = new EditAnOfferActionReq.Discounts(d.discounts || {}) }	
			if (!(d.location instanceof EditAnOfferActionReq.Location)) { this.location = new EditAnOfferActionReq.Location(d.location || {}) }	
			if (!(d.external instanceof EditAnOfferActionReq.External)) { this.external = new EditAnOfferActionReq.External(d.external || {}) }	
			if (!(d.taxSettings instanceof EditAnOfferActionReq.TaxSettings)) { this.taxSettings = new EditAnOfferActionReq.TaxSettings(d.taxSettings || {}) }	
			if (!(d.messageToSellerSettings instanceof EditAnOfferActionReq.MessageToSellerSettings)) { this.messageToSellerSettings = new EditAnOfferActionReq.MessageToSellerSettings(d.messageToSellerSettings || {}) }	
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
						EditAnOfferActionReq.Category.Fields
						);
						},
			productSet$: 'productSet',
get productSet() {
					return withPrefix(
						"productSet[:i]",
						EditAnOfferActionReq.ProductSet.Fields
						);
						},
			stock$: 'stock',
get stock() {
					return withPrefix(
						"stock",
						EditAnOfferActionReq.Stock.Fields
						);
						},
			sellingMode$: 'sellingMode',
get sellingMode() {
					return withPrefix(
						"sellingMode",
						EditAnOfferActionReq.SellingMode.Fields
						);
						},
			payments$: 'payments',
get payments() {
					return withPrefix(
						"payments",
						EditAnOfferActionReq.Payments.Fields
						);
						},
			delivery$: 'delivery',
get delivery() {
					return withPrefix(
						"delivery",
						EditAnOfferActionReq.Delivery.Fields
						);
						},
			publication$: 'publication',
get publication() {
					return withPrefix(
						"publication",
						EditAnOfferActionReq.Publication.Fields
						);
						},
			additionalMarketplaces: 'additionalMarketplaces',
			compatibilityList$: 'compatibilityList',
get compatibilityList() {
					return withPrefix(
						"compatibilityList",
						EditAnOfferActionReq.CompatibilityList.Fields
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
						EditAnOfferActionReq.Description.Fields
						);
						},
			b2b$: 'b2b',
get b2b() {
					return withPrefix(
						"b2b",
						EditAnOfferActionReq.B2b.Fields
						);
						},
			attachments$: 'attachments',
get attachments() {
					return withPrefix(
						"attachments[:i]",
						EditAnOfferActionReq.Attachments.Fields
						);
						},
			fundraisingCampaign$: 'fundraisingCampaign',
get fundraisingCampaign() {
					return withPrefix(
						"fundraisingCampaign",
						EditAnOfferActionReq.FundraisingCampaign.Fields
						);
						},
			additionalServices$: 'additionalServices',
get additionalServices() {
					return withPrefix(
						"additionalServices",
						EditAnOfferActionReq.AdditionalServices.Fields
						);
						},
			afterSalesServices$: 'afterSalesServices',
get afterSalesServices() {
					return withPrefix(
						"afterSalesServices",
						EditAnOfferActionReq.AfterSalesServices.Fields
						);
						},
			sizeTable$: 'sizeTable',
get sizeTable() {
					return withPrefix(
						"sizeTable",
						EditAnOfferActionReq.SizeTable.Fields
						);
						},
			contact$: 'contact',
get contact() {
					return withPrefix(
						"contact",
						EditAnOfferActionReq.Contact.Fields
						);
						},
			discounts$: 'discounts',
get discounts() {
					return withPrefix(
						"discounts",
						EditAnOfferActionReq.Discounts.Fields
						);
						},
			location$: 'location',
get location() {
					return withPrefix(
						"location",
						EditAnOfferActionReq.Location.Fields
						);
						},
			external$: 'external',
get external() {
					return withPrefix(
						"external",
						EditAnOfferActionReq.External.Fields
						);
						},
			taxSettings$: 'taxSettings',
get taxSettings() {
					return withPrefix(
						"taxSettings",
						EditAnOfferActionReq.TaxSettings.Fields
						);
						},
			messageToSellerSettings$: 'messageToSellerSettings',
get messageToSellerSettings() {
					return withPrefix(
						"messageToSellerSettings",
						EditAnOfferActionReq.MessageToSellerSettings.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionReq, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionReq(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionReq ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionReq(this.toJSON());
	}
}
/**
  * The base class definition for editAnOfferActionRes
  **/
export class EditAnOfferActionRes {
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
  * @type {string}
  **/
 #language  =  ""
		/**
  * 
  * @returns {string}
  **/
get language () { return this.#language }
/**
  * 
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
  * 
  * @type {EditAnOfferActionRes.Category}
  **/
 #category
		/**
  * 
  * @returns {EditAnOfferActionRes.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {EditAnOfferActionRes.Category}
  **/
set category (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Category) {
			this.#category = value
		} else {
			this.#category = new EditAnOfferActionRes.Category(value)
		}
}
setCategory (value) {
	this.category = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet}
  **/
 #productSet  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet}
  **/
get productSet () { return this.#productSet }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet}
  **/
set productSet (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.ProductSet) {
			this.#productSet = value
		} else {
			this.#productSet = value.map(item => new EditAnOfferActionRes.ProductSet(item))
		}
}
setProductSet (value) {
	this.productSet = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Stock}
  **/
 #stock
		/**
  * 
  * @returns {EditAnOfferActionRes.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {EditAnOfferActionRes.Stock}
  **/
set stock (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Stock) {
			this.#stock = value
		} else {
			this.#stock = new EditAnOfferActionRes.Stock(value)
		}
}
setStock (value) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Payments}
  **/
 #payments
		/**
  * 
  * @returns {EditAnOfferActionRes.Payments}
  **/
get payments () { return this.#payments }
/**
  * 
  * @type {EditAnOfferActionRes.Payments}
  **/
set payments (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Payments) {
			this.#payments = value
		} else {
			this.#payments = new EditAnOfferActionRes.Payments(value)
		}
}
setPayments (value) {
	this.payments = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.SellingMode}
  **/
 #sellingMode
		/**
  * 
  * @returns {EditAnOfferActionRes.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {EditAnOfferActionRes.SellingMode}
  **/
set sellingMode (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new EditAnOfferActionRes.SellingMode(value)
		}
}
setSellingMode (value) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Delivery}
  **/
 #delivery
		/**
  * 
  * @returns {EditAnOfferActionRes.Delivery}
  **/
get delivery () { return this.#delivery }
/**
  * 
  * @type {EditAnOfferActionRes.Delivery}
  **/
set delivery (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Delivery) {
			this.#delivery = value
		} else {
			this.#delivery = new EditAnOfferActionRes.Delivery(value)
		}
}
setDelivery (value) {
	this.delivery = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Publication}
  **/
 #publication
		/**
  * 
  * @returns {EditAnOfferActionRes.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {EditAnOfferActionRes.Publication}
  **/
set publication (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Publication) {
			this.#publication = value
		} else {
			this.#publication = new EditAnOfferActionRes.Publication(value)
		}
}
setPublication (value) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces}
  **/
 #additionalMarketplaces
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = new EditAnOfferActionRes.AdditionalMarketplaces(value)
		}
}
setAdditionalMarketplaces (value) {
	this.additionalMarketplaces = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.B2b}
  **/
 #b2b
		/**
  * 
  * @returns {EditAnOfferActionRes.B2b}
  **/
get b2b () { return this.#b2b }
/**
  * 
  * @type {EditAnOfferActionRes.B2b}
  **/
set b2b (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.B2b) {
			this.#b2b = value
		} else {
			this.#b2b = new EditAnOfferActionRes.B2b(value)
		}
}
setB2b (value) {
	this.b2b = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.CompatibilityList}
  **/
 #compatibilityList
		/**
  * 
  * @returns {EditAnOfferActionRes.CompatibilityList}
  **/
get compatibilityList () { return this.#compatibilityList }
/**
  * 
  * @type {EditAnOfferActionRes.CompatibilityList}
  **/
set compatibilityList (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.CompatibilityList) {
			this.#compatibilityList = value
		} else {
			this.#compatibilityList = new EditAnOfferActionRes.CompatibilityList(value)
		}
}
setCompatibilityList (value) {
	this.compatibilityList = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Validation}
  **/
 #validation
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation}
  **/
get validation () { return this.#validation }
/**
  * 
  * @type {EditAnOfferActionRes.Validation}
  **/
set validation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Validation) {
			this.#validation = value
		} else {
			this.#validation = new EditAnOfferActionRes.Validation(value)
		}
}
setValidation (value) {
	this.validation = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #warnings  =  []
		/**
  * 
  * @returns {string[]}
  **/
get warnings () { return this.#warnings }
/**
  * 
  * @type {string[]}
  **/
set warnings (value) {
}
setWarnings (value) {
	this.warnings = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices}
  **/
 #afterSalesServices
		/**
  * 
  * @returns {EditAnOfferActionRes.AfterSalesServices}
  **/
get afterSalesServices () { return this.#afterSalesServices }
/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices}
  **/
set afterSalesServices (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AfterSalesServices) {
			this.#afterSalesServices = value
		} else {
			this.#afterSalesServices = new EditAnOfferActionRes.AfterSalesServices(value)
		}
}
setAfterSalesServices (value) {
	this.afterSalesServices = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Discounts}
  **/
 #discounts
		/**
  * 
  * @returns {EditAnOfferActionRes.Discounts}
  **/
get discounts () { return this.#discounts }
/**
  * 
  * @type {EditAnOfferActionRes.Discounts}
  **/
set discounts (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Discounts) {
			this.#discounts = value
		} else {
			this.#discounts = new EditAnOfferActionRes.Discounts(value)
		}
}
setDiscounts (value) {
	this.discounts = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Contact}
  **/
 #contact
		/**
  * 
  * @returns {EditAnOfferActionRes.Contact}
  **/
get contact () { return this.#contact }
/**
  * 
  * @type {EditAnOfferActionRes.Contact}
  **/
set contact (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Contact) {
			this.#contact = value
		} else {
			this.#contact = new EditAnOfferActionRes.Contact(value)
		}
}
setContact (value) {
	this.contact = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Attachments}
  **/
 #attachments  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Attachments}
  **/
get attachments () { return this.#attachments }
/**
  * 
  * @type {EditAnOfferActionRes.Attachments}
  **/
set attachments (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.Attachments) {
			this.#attachments = value
		} else {
			this.#attachments = value.map(item => new EditAnOfferActionRes.Attachments(item))
		}
}
setAttachments (value) {
	this.attachments = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.FundraisingCampaign}
  **/
 #fundraisingCampaign
		/**
  * 
  * @returns {EditAnOfferActionRes.FundraisingCampaign}
  **/
get fundraisingCampaign () { return this.#fundraisingCampaign }
/**
  * 
  * @type {EditAnOfferActionRes.FundraisingCampaign}
  **/
set fundraisingCampaign (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.FundraisingCampaign) {
			this.#fundraisingCampaign = value
		} else {
			this.#fundraisingCampaign = new EditAnOfferActionRes.FundraisingCampaign(value)
		}
}
setFundraisingCampaign (value) {
	this.fundraisingCampaign = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalServices}
  **/
 #additionalServices
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalServices}
  **/
get additionalServices () { return this.#additionalServices }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalServices}
  **/
set additionalServices (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalServices) {
			this.#additionalServices = value
		} else {
			this.#additionalServices = new EditAnOfferActionRes.AdditionalServices(value)
		}
}
setAdditionalServices (value) {
	this.additionalServices = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.SizeTable}
  **/
 #sizeTable
		/**
  * 
  * @returns {EditAnOfferActionRes.SizeTable}
  **/
get sizeTable () { return this.#sizeTable }
/**
  * 
  * @type {EditAnOfferActionRes.SizeTable}
  **/
set sizeTable (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SizeTable) {
			this.#sizeTable = value
		} else {
			this.#sizeTable = new EditAnOfferActionRes.SizeTable(value)
		}
}
setSizeTable (value) {
	this.sizeTable = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Location}
  **/
 #location
		/**
  * 
  * @returns {EditAnOfferActionRes.Location}
  **/
get location () { return this.#location }
/**
  * 
  * @type {EditAnOfferActionRes.Location}
  **/
set location (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Location) {
			this.#location = value
		} else {
			this.#location = new EditAnOfferActionRes.Location(value)
		}
}
setLocation (value) {
	this.location = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.External}
  **/
 #external
		/**
  * 
  * @returns {EditAnOfferActionRes.External}
  **/
get external () { return this.#external }
/**
  * 
  * @type {EditAnOfferActionRes.External}
  **/
set external (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.External) {
			this.#external = value
		} else {
			this.#external = new EditAnOfferActionRes.External(value)
		}
}
setExternal (value) {
	this.external = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.TaxSettings}
  **/
 #taxSettings
		/**
  * 
  * @returns {EditAnOfferActionRes.TaxSettings}
  **/
get taxSettings () { return this.#taxSettings }
/**
  * 
  * @type {EditAnOfferActionRes.TaxSettings}
  **/
set taxSettings (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.TaxSettings) {
			this.#taxSettings = value
		} else {
			this.#taxSettings = new EditAnOfferActionRes.TaxSettings(value)
		}
}
setTaxSettings (value) {
	this.taxSettings = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.MessageToSellerSettings}
  **/
 #messageToSellerSettings
		/**
  * 
  * @returns {EditAnOfferActionRes.MessageToSellerSettings}
  **/
get messageToSellerSettings () { return this.#messageToSellerSettings }
/**
  * 
  * @type {EditAnOfferActionRes.MessageToSellerSettings}
  **/
set messageToSellerSettings (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.MessageToSellerSettings) {
			this.#messageToSellerSettings = value
		} else {
			this.#messageToSellerSettings = new EditAnOfferActionRes.MessageToSellerSettings(value)
		}
}
setMessageToSellerSettings (value) {
	this.messageToSellerSettings = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #createdAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get createdAt () { return this.#createdAt }
/**
  * 
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
  * 
  * @type {string}
  **/
 #updatedAt  =  ""
		/**
  * 
  * @returns {string}
  **/
get updatedAt () { return this.#updatedAt }
/**
  * 
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
  * @type {EditAnOfferActionRes.Description}
  **/
 #description
		/**
  * 
  * @returns {EditAnOfferActionRes.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {EditAnOfferActionRes.Description}
  **/
set description (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Description) {
			this.#description = value
		} else {
			this.#description = new EditAnOfferActionRes.Description(value)
		}
}
setDescription (value) {
	this.description = value
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
	* Creates an instance of EditAnOfferActionRes.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Category(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Category ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Category(this.toJSON());
	}
}
/**
  * The base class definition for productSet
  **/
static ProductSet = class ProductSet {
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Quantity}
  **/
 #quantity
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Quantity}
  **/
get quantity () { return this.#quantity }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Quantity}
  **/
set quantity (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.Quantity) {
			this.#quantity = value
		} else {
			this.#quantity = new EditAnOfferActionRes.ProductSet.Quantity(value)
		}
}
setQuantity (value) {
	this.quantity = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product}
  **/
 #product
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Product}
  **/
get product () { return this.#product }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product}
  **/
set product (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.Product) {
			this.#product = value
		} else {
			this.#product = new EditAnOfferActionRes.ProductSet.Product(value)
		}
}
setProduct (value) {
	this.product = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.ResponsiblePerson}
  **/
 #responsiblePerson
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.ResponsiblePerson}
  **/
get responsiblePerson () { return this.#responsiblePerson }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.ResponsiblePerson}
  **/
set responsiblePerson (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.ResponsiblePerson) {
			this.#responsiblePerson = value
		} else {
			this.#responsiblePerson = new EditAnOfferActionRes.ProductSet.ResponsiblePerson(value)
		}
}
setResponsiblePerson (value) {
	this.responsiblePerson = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.ResponsibleProducer}
  **/
 #responsibleProducer
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.ResponsibleProducer}
  **/
get responsibleProducer () { return this.#responsibleProducer }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.ResponsibleProducer}
  **/
set responsibleProducer (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.ResponsibleProducer) {
			this.#responsibleProducer = value
		} else {
			this.#responsibleProducer = new EditAnOfferActionRes.ProductSet.ResponsibleProducer(value)
		}
}
setResponsibleProducer (value) {
	this.responsibleProducer = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.SafetyInformation}
  **/
 #safetyInformation
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.SafetyInformation}
  **/
set safetyInformation (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new EditAnOfferActionRes.ProductSet.SafetyInformation(value)
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
  * @type {EditAnOfferActionRes.ProductSet.Deposits}
  **/
 #deposits  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Deposits}
  **/
get deposits () { return this.#deposits }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Deposits}
  **/
set deposits (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.ProductSet.Deposits) {
			this.#deposits = value
		} else {
			this.#deposits = value.map(item => new EditAnOfferActionRes.ProductSet.Deposits(item))
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.Quantity, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Quantity(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Quantity, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Quantity(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet.Quantity ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet.Quantity(this.toJSON());
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
  * @type {EditAnOfferActionRes.ProductSet.Product.Publication}
  **/
 #publication
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Product.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product.Publication}
  **/
set publication (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.Product.Publication) {
			this.#publication = value
		} else {
			this.#publication = new EditAnOfferActionRes.ProductSet.Product.Publication(value)
		}
}
setPublication (value) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product.Parameters}
  **/
 #parameters  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Product.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product.Parameters}
  **/
set parameters (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.ProductSet.Product.Parameters) {
			this.#parameters = value
		} else {
			this.#parameters = value.map(item => new EditAnOfferActionRes.ProductSet.Product.Parameters(item))
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Product.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Product.Publication(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet.Product.Publication ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet.Product.Publication(this.toJSON());
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
  * @type {EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
 #rangeValue
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
get rangeValue () { return this.#rangeValue }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
set rangeValue (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue) {
			this.#rangeValue = value
		} else {
			this.#rangeValue = new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue(value)
		}
}
setRangeValue (value) {
	this.rangeValue = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #values  =  []
		/**
  * 
  * @returns {string[]}
  **/
get values () { return this.#values }
/**
  * 
  * @type {string[]}
  **/
set values (value) {
}
setValues (value) {
	this.values = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #valuesIds  =  []
		/**
  * 
  * @returns {string[]}
  **/
get valuesIds () { return this.#valuesIds }
/**
  * 
  * @type {string[]}
  **/
set valuesIds (value) {
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue(this.toJSON());
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
			if (!(d.rangeValue instanceof EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue)) { this.rangeValue = new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue(d.rangeValue || {}) }	
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
						EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue.Fields
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Parameters, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters(this.toJSON());
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
			if (!(d.publication instanceof EditAnOfferActionRes.ProductSet.Product.Publication)) { this.publication = new EditAnOfferActionRes.ProductSet.Product.Publication(d.publication || {}) }	
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
						EditAnOfferActionRes.ProductSet.Product.Publication.Fields
						);
						},
			parameters$: 'parameters',
get parameters() {
					return withPrefix(
						"productSet.product.parameters[:i]",
						EditAnOfferActionRes.ProductSet.Product.Parameters.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Product(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Product(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet.Product ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet.Product(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.ResponsiblePerson, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet.ResponsiblePerson(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.ResponsiblePerson, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet.ResponsiblePerson(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet.ResponsiblePerson ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet.ResponsiblePerson(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.ResponsibleProducer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet.ResponsibleProducer(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.ResponsibleProducer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet.ResponsibleProducer(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet.ResponsibleProducer ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet.ResponsibleProducer(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.SafetyInformation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet.SafetyInformation(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet.SafetyInformation(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.Deposits, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Deposits(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Deposits, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet.Deposits(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet.Deposits ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet.Deposits(this.toJSON());
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
			if (!(d.quantity instanceof EditAnOfferActionRes.ProductSet.Quantity)) { this.quantity = new EditAnOfferActionRes.ProductSet.Quantity(d.quantity || {}) }	
			if (!(d.product instanceof EditAnOfferActionRes.ProductSet.Product)) { this.product = new EditAnOfferActionRes.ProductSet.Product(d.product || {}) }	
			if (!(d.responsiblePerson instanceof EditAnOfferActionRes.ProductSet.ResponsiblePerson)) { this.responsiblePerson = new EditAnOfferActionRes.ProductSet.ResponsiblePerson(d.responsiblePerson || {}) }	
			if (!(d.responsibleProducer instanceof EditAnOfferActionRes.ProductSet.ResponsibleProducer)) { this.responsibleProducer = new EditAnOfferActionRes.ProductSet.ResponsibleProducer(d.responsibleProducer || {}) }	
			if (!(d.safetyInformation instanceof EditAnOfferActionRes.ProductSet.SafetyInformation)) { this.safetyInformation = new EditAnOfferActionRes.ProductSet.SafetyInformation(d.safetyInformation || {}) }	
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
						EditAnOfferActionRes.ProductSet.Quantity.Fields
						);
						},
			product$: 'product',
get product() {
					return withPrefix(
						"productSet.product",
						EditAnOfferActionRes.ProductSet.Product.Fields
						);
						},
			responsiblePerson$: 'responsiblePerson',
get responsiblePerson() {
					return withPrefix(
						"productSet.responsiblePerson",
						EditAnOfferActionRes.ProductSet.ResponsiblePerson.Fields
						);
						},
			responsibleProducer$: 'responsibleProducer',
get responsibleProducer() {
					return withPrefix(
						"productSet.responsibleProducer",
						EditAnOfferActionRes.ProductSet.ResponsibleProducer.Fields
						);
						},
			safetyInformation$: 'safetyInformation',
get safetyInformation() {
					return withPrefix(
						"productSet.safetyInformation",
						EditAnOfferActionRes.ProductSet.SafetyInformation.Fields
						);
						},
			marketedBeforeGPSRObligation: 'marketedBeforeGPSRObligation',
			deposits$: 'deposits',
get deposits() {
					return withPrefix(
						"productSet.deposits[:i]",
						EditAnOfferActionRes.ProductSet.Deposits.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.ProductSet(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.ProductSet(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.ProductSet ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.ProductSet(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.Stock, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Stock(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Stock ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Stock(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.Payments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Payments(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Payments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Payments(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Payments ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Payments(this.toJSON());
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
  * @type {EditAnOfferActionRes.SellingMode.Price}
  **/
 #price
		/**
  * 
  * @returns {EditAnOfferActionRes.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.Price}
  **/
set price (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new EditAnOfferActionRes.SellingMode.Price(value)
		}
}
setPrice (value) {
	this.price = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.MinimalPrice}
  **/
 #minimalPrice
		/**
  * 
  * @returns {EditAnOfferActionRes.SellingMode.MinimalPrice}
  **/
get minimalPrice () { return this.#minimalPrice }
/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.MinimalPrice}
  **/
set minimalPrice (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SellingMode.MinimalPrice) {
			this.#minimalPrice = value
		} else {
			this.#minimalPrice = new EditAnOfferActionRes.SellingMode.MinimalPrice(value)
		}
}
setMinimalPrice (value) {
	this.minimalPrice = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.StartingPrice}
  **/
 #startingPrice
		/**
  * 
  * @returns {EditAnOfferActionRes.SellingMode.StartingPrice}
  **/
get startingPrice () { return this.#startingPrice }
/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.StartingPrice}
  **/
set startingPrice (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SellingMode.StartingPrice) {
			this.#startingPrice = value
		} else {
			this.#startingPrice = new EditAnOfferActionRes.SellingMode.StartingPrice(value)
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
	* Creates an instance of EditAnOfferActionRes.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.SellingMode.Price(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.SellingMode.MinimalPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.SellingMode.MinimalPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SellingMode.MinimalPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.SellingMode.MinimalPrice(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.SellingMode.MinimalPrice ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.SellingMode.MinimalPrice(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.SellingMode.StartingPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.SellingMode.StartingPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SellingMode.StartingPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.SellingMode.StartingPrice(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.SellingMode.StartingPrice ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.SellingMode.StartingPrice(this.toJSON());
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
			if (!(d.price instanceof EditAnOfferActionRes.SellingMode.Price)) { this.price = new EditAnOfferActionRes.SellingMode.Price(d.price || {}) }	
			if (!(d.minimalPrice instanceof EditAnOfferActionRes.SellingMode.MinimalPrice)) { this.minimalPrice = new EditAnOfferActionRes.SellingMode.MinimalPrice(d.minimalPrice || {}) }	
			if (!(d.startingPrice instanceof EditAnOfferActionRes.SellingMode.StartingPrice)) { this.startingPrice = new EditAnOfferActionRes.SellingMode.StartingPrice(d.startingPrice || {}) }	
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
						EditAnOfferActionRes.SellingMode.Price.Fields
						);
						},
			minimalPrice$: 'minimalPrice',
get minimalPrice() {
					return withPrefix(
						"sellingMode.minimalPrice",
						EditAnOfferActionRes.SellingMode.MinimalPrice.Fields
						);
						},
			startingPrice$: 'startingPrice',
get startingPrice() {
					return withPrefix(
						"sellingMode.startingPrice",
						EditAnOfferActionRes.SellingMode.StartingPrice.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SellingMode, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.SellingMode(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.SellingMode(this.toJSON());
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
  * @type {EditAnOfferActionRes.Delivery.ShippingRates}
  **/
 #shippingRates
		/**
  * 
  * @returns {EditAnOfferActionRes.Delivery.ShippingRates}
  **/
get shippingRates () { return this.#shippingRates }
/**
  * 
  * @type {EditAnOfferActionRes.Delivery.ShippingRates}
  **/
set shippingRates (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Delivery.ShippingRates) {
			this.#shippingRates = value
		} else {
			this.#shippingRates = new EditAnOfferActionRes.Delivery.ShippingRates(value)
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
	* Creates an instance of EditAnOfferActionRes.Delivery.ShippingRates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Delivery.ShippingRates(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Delivery.ShippingRates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Delivery.ShippingRates(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Delivery.ShippingRates ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Delivery.ShippingRates(this.toJSON());
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
			if (!(d.shippingRates instanceof EditAnOfferActionRes.Delivery.ShippingRates)) { this.shippingRates = new EditAnOfferActionRes.Delivery.ShippingRates(d.shippingRates || {}) }	
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
						EditAnOfferActionRes.Delivery.ShippingRates.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Delivery, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Delivery(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Delivery, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Delivery(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Delivery ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Delivery(this.toJSON());
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
  * @type {EditAnOfferActionRes.Publication.Marketplaces}
  **/
 #marketplaces
		/**
  * 
  * @returns {EditAnOfferActionRes.Publication.Marketplaces}
  **/
get marketplaces () { return this.#marketplaces }
/**
  * 
  * @type {EditAnOfferActionRes.Publication.Marketplaces}
  **/
set marketplaces (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Publication.Marketplaces) {
			this.#marketplaces = value
		} else {
			this.#marketplaces = new EditAnOfferActionRes.Publication.Marketplaces(value)
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
  * @type {EditAnOfferActionRes.Publication.Marketplaces.Base}
  **/
 #base
		/**
  * 
  * @returns {EditAnOfferActionRes.Publication.Marketplaces.Base}
  **/
get base () { return this.#base }
/**
  * 
  * @type {EditAnOfferActionRes.Publication.Marketplaces.Base}
  **/
set base (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Publication.Marketplaces.Base) {
			this.#base = value
		} else {
			this.#base = new EditAnOfferActionRes.Publication.Marketplaces.Base(value)
		}
}
setBase (value) {
	this.base = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Publication.Marketplaces.Additional}
  **/
 #additional  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Publication.Marketplaces.Additional}
  **/
get additional () { return this.#additional }
/**
  * 
  * @type {EditAnOfferActionRes.Publication.Marketplaces.Additional}
  **/
set additional (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.Publication.Marketplaces.Additional) {
			this.#additional = value
		} else {
			this.#additional = value.map(item => new EditAnOfferActionRes.Publication.Marketplaces.Additional(item))
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
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces.Base, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Base(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces.Base, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Base(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Base ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Publication.Marketplaces.Base(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces.Additional, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Additional(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces.Additional, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Additional(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Additional ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Publication.Marketplaces.Additional(this.toJSON());
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
			if (!(d.base instanceof EditAnOfferActionRes.Publication.Marketplaces.Base)) { this.base = new EditAnOfferActionRes.Publication.Marketplaces.Base(d.base || {}) }	
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
						EditAnOfferActionRes.Publication.Marketplaces.Base.Fields
						);
						},
			additional$: 'additional',
get additional() {
					return withPrefix(
						"publication.marketplaces.additional[:i]",
						EditAnOfferActionRes.Publication.Marketplaces.Additional.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Publication.Marketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Publication.Marketplaces(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Publication.Marketplaces ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Publication.Marketplaces(this.toJSON());
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
			if (!(d.marketplaces instanceof EditAnOfferActionRes.Publication.Marketplaces)) { this.marketplaces = new EditAnOfferActionRes.Publication.Marketplaces(d.marketplaces || {}) }	
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
						EditAnOfferActionRes.Publication.Marketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Publication(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Publication ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Publication(this.toJSON());
	}
}
/**
  * The base class definition for additionalMarketplaces
  **/
static AdditionalMarketplaces = class AdditionalMarketplaces {
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode}
  **/
 #sellingMode
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode}
  **/
set sellingMode (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode(value)
		}
}
setSellingMode (value) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication}
  **/
 #publication
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication}
  **/
set publication (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces.Publication) {
			this.#publication = value
		} else {
			this.#publication = new EditAnOfferActionRes.AdditionalMarketplaces.Publication(value)
		}
}
setPublication (value) {
	this.publication = value
	return this
}
/**
  * The base class definition for sellingMode
  **/
static SellingMode = class SellingMode {
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price}
  **/
 #price
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price}
  **/
set price (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price(value)
		}
}
setPrice (value) {
	this.price = value
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
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price(this.toJSON());
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
			if (d.price !== undefined) { this.price = d.price }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.price instanceof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price)) { this.price = new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price(d.price || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				price: this.#price,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			price$: 'price',
get price() {
					return withPrefix(
						"additionalMarketplaces.sellingMode.price",
						EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.SellingMode, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode(this.toJSON());
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
 #state  =  ""
		/**
  * 
  * @returns {string}
  **/
get state () { return this.#state }
/**
  * 
  * @type {string}
  **/
set state (value) {
		this.#state = String(value);
}
setState (value) {
	this.state = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons}
  **/
 #refusalReasons  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons}
  **/
get refusalReasons () { return this.#refusalReasons }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons}
  **/
set refusalReasons (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons) {
			this.#refusalReasons = value
		} else {
			this.#refusalReasons = value.map(item => new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons(item))
		}
}
setRefusalReasons (value) {
	this.refusalReasons = value
	return this
}
/**
  * The base class definition for refusalReasons
  **/
static RefusalReasons = class RefusalReasons {
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
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters}
  **/
 #parameters
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters}
  **/
set parameters (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters) {
			this.#parameters = value
		} else {
			this.#parameters = new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters(value)
		}
}
setParameters (value) {
	this.parameters = value
	return this
}
/**
  * The base class definition for parameters
  **/
static Parameters = class Parameters {
		/**
  * 
  * @type {string[]}
  **/
 #maxAllowedPriceDecreasePercent  =  []
		/**
  * 
  * @returns {string[]}
  **/
get maxAllowedPriceDecreasePercent () { return this.#maxAllowedPriceDecreasePercent }
/**
  * 
  * @type {string[]}
  **/
set maxAllowedPriceDecreasePercent (value) {
}
setMaxAllowedPriceDecreasePercent (value) {
	this.maxAllowedPriceDecreasePercent = value
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
			if (d.maxAllowedPriceDecreasePercent !== undefined) { this.maxAllowedPriceDecreasePercent = d.maxAllowedPriceDecreasePercent }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				maxAllowedPriceDecreasePercent: this.#maxAllowedPriceDecreasePercent,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			maxAllowedPriceDecreasePercent$: 'maxAllowedPriceDecreasePercent',
get maxAllowedPriceDecreasePercent() {
					return "additionalMarketplaces.publication.refusalReasons.parameters.maxAllowedPriceDecreasePercent[:i]";
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters(this.toJSON());
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
			if (d.userMessage !== undefined) { this.userMessage = d.userMessage }
			if (d.parameters !== undefined) { this.parameters = d.parameters }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.parameters instanceof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters)) { this.parameters = new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters(d.parameters || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				code: this.#code,
				userMessage: this.#userMessage,
				parameters: this.#parameters,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			code: 'code',
			userMessage: 'userMessage',
			parameters$: 'parameters',
get parameters() {
					return withPrefix(
						"additionalMarketplaces.publication.refusalReasons.parameters",
						EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons(this.toJSON());
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
			if (d.state !== undefined) { this.state = d.state }
			if (d.refusalReasons !== undefined) { this.refusalReasons = d.refusalReasons }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				state: this.#state,
				refusalReasons: this.#refusalReasons,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			state: 'state',
			refusalReasons$: 'refusalReasons',
get refusalReasons() {
					return withPrefix(
						"additionalMarketplaces.publication.refusalReasons[:i]",
						EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication(this.toJSON());
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
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
			if (d.publication !== undefined) { this.publication = d.publication }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.sellingMode instanceof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode)) { this.sellingMode = new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode(d.sellingMode || {}) }	
			if (!(d.publication instanceof EditAnOfferActionRes.AdditionalMarketplaces.Publication)) { this.publication = new EditAnOfferActionRes.AdditionalMarketplaces.Publication(d.publication || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				sellingMode: this.#sellingMode,
				publication: this.#publication,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			sellingMode$: 'sellingMode',
get sellingMode() {
					return withPrefix(
						"additionalMarketplaces.sellingMode",
						EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Fields
						);
						},
			publication$: 'publication',
get publication() {
					return withPrefix(
						"additionalMarketplaces.publication",
						EditAnOfferActionRes.AdditionalMarketplaces.Publication.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AdditionalMarketplaces(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.B2b, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.B2b(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.B2b, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.B2b(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.B2b ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.B2b(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.CompatibilityList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.CompatibilityList(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.CompatibilityList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.CompatibilityList(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.CompatibilityList ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.CompatibilityList(this.toJSON());
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
  * @type {EditAnOfferActionRes.Validation.Errors}
  **/
 #errors  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {EditAnOfferActionRes.Validation.Errors}
  **/
set errors (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.Validation.Errors) {
			this.#errors = value
		} else {
			this.#errors = value.map(item => new EditAnOfferActionRes.Validation.Errors(item))
		}
}
setErrors (value) {
	this.errors = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Validation.Warnings}
  **/
 #warnings  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation.Warnings}
  **/
get warnings () { return this.#warnings }
/**
  * 
  * @type {EditAnOfferActionRes.Validation.Warnings}
  **/
set warnings (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.Validation.Warnings) {
			this.#warnings = value
		} else {
			this.#warnings = value.map(item => new EditAnOfferActionRes.Validation.Warnings(item))
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
  * @type {EditAnOfferActionRes.Validation.Errors.Metadata}
  **/
 #metadata
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {EditAnOfferActionRes.Validation.Errors.Metadata}
  **/
set metadata (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Validation.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new EditAnOfferActionRes.Validation.Errors.Metadata(value)
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
	* Creates an instance of EditAnOfferActionRes.Validation.Errors.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Validation.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Validation.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Validation.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Validation.Errors.Metadata(this.toJSON());
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
			if (!(d.metadata instanceof EditAnOfferActionRes.Validation.Errors.Metadata)) { this.metadata = new EditAnOfferActionRes.Validation.Errors.Metadata(d.metadata || {}) }	
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
						EditAnOfferActionRes.Validation.Errors.Metadata.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Errors, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Validation.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Validation.Errors(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Validation.Errors ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Validation.Errors(this.toJSON());
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
  * @type {EditAnOfferActionRes.Validation.Warnings.Metadata}
  **/
 #metadata
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation.Warnings.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {EditAnOfferActionRes.Validation.Warnings.Metadata}
  **/
set metadata (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Validation.Warnings.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new EditAnOfferActionRes.Validation.Warnings.Metadata(value)
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
	* Creates an instance of EditAnOfferActionRes.Validation.Warnings.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Validation.Warnings.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Warnings.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Validation.Warnings.Metadata(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Validation.Warnings.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Validation.Warnings.Metadata(this.toJSON());
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
			if (!(d.metadata instanceof EditAnOfferActionRes.Validation.Warnings.Metadata)) { this.metadata = new EditAnOfferActionRes.Validation.Warnings.Metadata(d.metadata || {}) }	
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
						EditAnOfferActionRes.Validation.Warnings.Metadata.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Warnings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Validation.Warnings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Warnings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Validation.Warnings(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Validation.Warnings ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Validation.Warnings(this.toJSON());
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
						EditAnOfferActionRes.Validation.Errors.Fields
						);
						},
			warnings$: 'warnings',
get warnings() {
					return withPrefix(
						"validation.warnings[:i]",
						EditAnOfferActionRes.Validation.Warnings.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Validation(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Validation(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Validation ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Validation(this.toJSON());
	}
}
/**
  * The base class definition for afterSalesServices
  **/
static AfterSalesServices = class AfterSalesServices {
		/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
 #impliedWarranty
		/**
  * 
  * @returns {EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
get impliedWarranty () { return this.#impliedWarranty }
/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
set impliedWarranty (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty) {
			this.#impliedWarranty = value
		} else {
			this.#impliedWarranty = new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty(value)
		}
}
setImpliedWarranty (value) {
	this.impliedWarranty = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
 #returnPolicy
		/**
  * 
  * @returns {EditAnOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
get returnPolicy () { return this.#returnPolicy }
/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
set returnPolicy (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AfterSalesServices.ReturnPolicy) {
			this.#returnPolicy = value
		} else {
			this.#returnPolicy = new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy(value)
		}
}
setReturnPolicy (value) {
	this.returnPolicy = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.Warranty}
  **/
 #warranty
		/**
  * 
  * @returns {EditAnOfferActionRes.AfterSalesServices.Warranty}
  **/
get warranty () { return this.#warranty }
/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.Warranty}
  **/
set warranty (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AfterSalesServices.Warranty) {
			this.#warranty = value
		} else {
			this.#warranty = new EditAnOfferActionRes.AfterSalesServices.Warranty(value)
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
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.ReturnPolicy, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.ReturnPolicy, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.Warranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AfterSalesServices.Warranty(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.Warranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AfterSalesServices.Warranty(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AfterSalesServices.Warranty ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AfterSalesServices.Warranty(this.toJSON());
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
			if (!(d.impliedWarranty instanceof EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty)) { this.impliedWarranty = new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty(d.impliedWarranty || {}) }	
			if (!(d.returnPolicy instanceof EditAnOfferActionRes.AfterSalesServices.ReturnPolicy)) { this.returnPolicy = new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy(d.returnPolicy || {}) }	
			if (!(d.warranty instanceof EditAnOfferActionRes.AfterSalesServices.Warranty)) { this.warranty = new EditAnOfferActionRes.AfterSalesServices.Warranty(d.warranty || {}) }	
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
						EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty.Fields
						);
						},
			returnPolicy$: 'returnPolicy',
get returnPolicy() {
					return withPrefix(
						"afterSalesServices.returnPolicy",
						EditAnOfferActionRes.AfterSalesServices.ReturnPolicy.Fields
						);
						},
			warranty$: 'warranty',
get warranty() {
					return withPrefix(
						"afterSalesServices.warranty",
						EditAnOfferActionRes.AfterSalesServices.Warranty.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AfterSalesServices(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AfterSalesServices(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AfterSalesServices ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AfterSalesServices(this.toJSON());
	}
}
/**
  * The base class definition for discounts
  **/
static Discounts = class Discounts {
		/**
  * 
  * @type {EditAnOfferActionRes.Discounts.WholesalePriceList}
  **/
 #wholesalePriceList
		/**
  * 
  * @returns {EditAnOfferActionRes.Discounts.WholesalePriceList}
  **/
get wholesalePriceList () { return this.#wholesalePriceList }
/**
  * 
  * @type {EditAnOfferActionRes.Discounts.WholesalePriceList}
  **/
set wholesalePriceList (value) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Discounts.WholesalePriceList) {
			this.#wholesalePriceList = value
		} else {
			this.#wholesalePriceList = new EditAnOfferActionRes.Discounts.WholesalePriceList(value)
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
	* Creates an instance of EditAnOfferActionRes.Discounts.WholesalePriceList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Discounts.WholesalePriceList(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Discounts.WholesalePriceList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Discounts.WholesalePriceList(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Discounts.WholesalePriceList ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Discounts.WholesalePriceList(this.toJSON());
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
			if (!(d.wholesalePriceList instanceof EditAnOfferActionRes.Discounts.WholesalePriceList)) { this.wholesalePriceList = new EditAnOfferActionRes.Discounts.WholesalePriceList(d.wholesalePriceList || {}) }	
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
						EditAnOfferActionRes.Discounts.WholesalePriceList.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Discounts, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Discounts(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Discounts, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Discounts(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Discounts ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Discounts(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.Contact, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Contact(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Contact, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Contact(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Contact ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Contact(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.Attachments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Attachments(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Attachments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Attachments(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Attachments ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Attachments(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.FundraisingCampaign, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.FundraisingCampaign(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.FundraisingCampaign, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.FundraisingCampaign(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.FundraisingCampaign ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.FundraisingCampaign(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.AdditionalServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.AdditionalServices(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.AdditionalServices(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.AdditionalServices ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.AdditionalServices(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.SizeTable, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.SizeTable(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SizeTable, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.SizeTable(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.SizeTable ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.SizeTable(this.toJSON());
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
 #city  =  ""
		/**
  * 
  * @returns {string}
  **/
get city () { return this.#city }
/**
  * 
  * @type {string}
  **/
set city (value) {
		this.#city = String(value);
}
setCity (value) {
	this.city = value
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
		/**
  * 
  * @type {string}
  **/
 #postCode  =  ""
		/**
  * 
  * @returns {string}
  **/
get postCode () { return this.#postCode }
/**
  * 
  * @type {string}
  **/
set postCode (value) {
		this.#postCode = String(value);
}
setPostCode (value) {
	this.postCode = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #province  =  ""
		/**
  * 
  * @returns {string}
  **/
get province () { return this.#province }
/**
  * 
  * @type {string}
  **/
set province (value) {
		this.#province = String(value);
}
setProvince (value) {
	this.province = value
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
	* Creates an instance of EditAnOfferActionRes.Location, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Location(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Location, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Location(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Location ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Location(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.External, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.External(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.External, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.External(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.External ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.External(this.toJSON());
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
  * @type {EditAnOfferActionRes.TaxSettings.Rates}
  **/
 #rates  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.TaxSettings.Rates}
  **/
get rates () { return this.#rates }
/**
  * 
  * @type {EditAnOfferActionRes.TaxSettings.Rates}
  **/
set rates (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.TaxSettings.Rates) {
			this.#rates = value
		} else {
			this.#rates = value.map(item => new EditAnOfferActionRes.TaxSettings.Rates(item))
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
	* Creates an instance of EditAnOfferActionRes.TaxSettings.Rates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.TaxSettings.Rates(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.TaxSettings.Rates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.TaxSettings.Rates(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.TaxSettings.Rates ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.TaxSettings.Rates(this.toJSON());
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
						EditAnOfferActionRes.TaxSettings.Rates.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.TaxSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.TaxSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.TaxSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.TaxSettings(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.TaxSettings ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.TaxSettings(this.toJSON());
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
	* Creates an instance of EditAnOfferActionRes.MessageToSellerSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.MessageToSellerSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.MessageToSellerSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.MessageToSellerSettings(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.MessageToSellerSettings ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.MessageToSellerSettings(this.toJSON());
	}
}
/**
  * The base class definition for description
  **/
static Description = class Description {
		/**
  * 
  * @type {EditAnOfferActionRes.Description.Sections}
  **/
 #sections  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Description.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {EditAnOfferActionRes.Description.Sections}
  **/
set sections (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.Description.Sections) {
			this.#sections = value
		} else {
			this.#sections = value.map(item => new EditAnOfferActionRes.Description.Sections(item))
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
  * @type {EditAnOfferActionRes.Description.Sections.Items}
  **/
 #items  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Description.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {EditAnOfferActionRes.Description.Sections.Items}
  **/
set items (value) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof EditAnOfferActionRes.Description.Sections.Items) {
			this.#items = value
		} else {
			this.#items = value.map(item => new EditAnOfferActionRes.Description.Sections.Items(item))
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
	* Creates an instance of EditAnOfferActionRes.Description.Sections.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Description.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Description.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Description.Sections.Items(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Description.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Description.Sections.Items(this.toJSON());
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
						EditAnOfferActionRes.Description.Sections.Items.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Description.Sections, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Description.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Description.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Description.Sections(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Description.Sections ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Description.Sections(this.toJSON());
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
						EditAnOfferActionRes.Description.Sections.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Description, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes.Description(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes.Description ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes.Description(this.toJSON());
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
			if (d.category !== undefined) { this.category = d.category }
			if (d.productSet !== undefined) { this.productSet = d.productSet }
			if (d.stock !== undefined) { this.stock = d.stock }
			if (d.payments !== undefined) { this.payments = d.payments }
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
			if (d.delivery !== undefined) { this.delivery = d.delivery }
			if (d.publication !== undefined) { this.publication = d.publication }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
			if (d.b2b !== undefined) { this.b2b = d.b2b }
			if (d.compatibilityList !== undefined) { this.compatibilityList = d.compatibilityList }
			if (d.validation !== undefined) { this.validation = d.validation }
			if (d.warnings !== undefined) { this.warnings = d.warnings }
			if (d.afterSalesServices !== undefined) { this.afterSalesServices = d.afterSalesServices }
			if (d.discounts !== undefined) { this.discounts = d.discounts }
			if (d.contact !== undefined) { this.contact = d.contact }
			if (d.attachments !== undefined) { this.attachments = d.attachments }
			if (d.fundraisingCampaign !== undefined) { this.fundraisingCampaign = d.fundraisingCampaign }
			if (d.additionalServices !== undefined) { this.additionalServices = d.additionalServices }
			if (d.sizeTable !== undefined) { this.sizeTable = d.sizeTable }
			if (d.location !== undefined) { this.location = d.location }
			if (d.external !== undefined) { this.external = d.external }
			if (d.taxSettings !== undefined) { this.taxSettings = d.taxSettings }
			if (d.messageToSellerSettings !== undefined) { this.messageToSellerSettings = d.messageToSellerSettings }
			if (d.createdAt !== undefined) { this.createdAt = d.createdAt }
			if (d.updatedAt !== undefined) { this.updatedAt = d.updatedAt }
			if (d.images !== undefined) { this.images = d.images }
			if (d.description !== undefined) { this.description = d.description }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data;
			if (!(d.category instanceof EditAnOfferActionRes.Category)) { this.category = new EditAnOfferActionRes.Category(d.category || {}) }	
			if (!(d.stock instanceof EditAnOfferActionRes.Stock)) { this.stock = new EditAnOfferActionRes.Stock(d.stock || {}) }	
			if (!(d.payments instanceof EditAnOfferActionRes.Payments)) { this.payments = new EditAnOfferActionRes.Payments(d.payments || {}) }	
			if (!(d.sellingMode instanceof EditAnOfferActionRes.SellingMode)) { this.sellingMode = new EditAnOfferActionRes.SellingMode(d.sellingMode || {}) }	
			if (!(d.delivery instanceof EditAnOfferActionRes.Delivery)) { this.delivery = new EditAnOfferActionRes.Delivery(d.delivery || {}) }	
			if (!(d.publication instanceof EditAnOfferActionRes.Publication)) { this.publication = new EditAnOfferActionRes.Publication(d.publication || {}) }	
			if (!(d.additionalMarketplaces instanceof EditAnOfferActionRes.AdditionalMarketplaces)) { this.additionalMarketplaces = new EditAnOfferActionRes.AdditionalMarketplaces(d.additionalMarketplaces || {}) }	
			if (!(d.b2b instanceof EditAnOfferActionRes.B2b)) { this.b2b = new EditAnOfferActionRes.B2b(d.b2b || {}) }	
			if (!(d.compatibilityList instanceof EditAnOfferActionRes.CompatibilityList)) { this.compatibilityList = new EditAnOfferActionRes.CompatibilityList(d.compatibilityList || {}) }	
			if (!(d.validation instanceof EditAnOfferActionRes.Validation)) { this.validation = new EditAnOfferActionRes.Validation(d.validation || {}) }	
			if (!(d.afterSalesServices instanceof EditAnOfferActionRes.AfterSalesServices)) { this.afterSalesServices = new EditAnOfferActionRes.AfterSalesServices(d.afterSalesServices || {}) }	
			if (!(d.discounts instanceof EditAnOfferActionRes.Discounts)) { this.discounts = new EditAnOfferActionRes.Discounts(d.discounts || {}) }	
			if (!(d.contact instanceof EditAnOfferActionRes.Contact)) { this.contact = new EditAnOfferActionRes.Contact(d.contact || {}) }	
			if (!(d.fundraisingCampaign instanceof EditAnOfferActionRes.FundraisingCampaign)) { this.fundraisingCampaign = new EditAnOfferActionRes.FundraisingCampaign(d.fundraisingCampaign || {}) }	
			if (!(d.additionalServices instanceof EditAnOfferActionRes.AdditionalServices)) { this.additionalServices = new EditAnOfferActionRes.AdditionalServices(d.additionalServices || {}) }	
			if (!(d.sizeTable instanceof EditAnOfferActionRes.SizeTable)) { this.sizeTable = new EditAnOfferActionRes.SizeTable(d.sizeTable || {}) }	
			if (!(d.location instanceof EditAnOfferActionRes.Location)) { this.location = new EditAnOfferActionRes.Location(d.location || {}) }	
			if (!(d.external instanceof EditAnOfferActionRes.External)) { this.external = new EditAnOfferActionRes.External(d.external || {}) }	
			if (!(d.taxSettings instanceof EditAnOfferActionRes.TaxSettings)) { this.taxSettings = new EditAnOfferActionRes.TaxSettings(d.taxSettings || {}) }	
			if (!(d.messageToSellerSettings instanceof EditAnOfferActionRes.MessageToSellerSettings)) { this.messageToSellerSettings = new EditAnOfferActionRes.MessageToSellerSettings(d.messageToSellerSettings || {}) }	
			if (!(d.description instanceof EditAnOfferActionRes.Description)) { this.description = new EditAnOfferActionRes.Description(d.description || {}) }	
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
				category: this.#category,
				productSet: this.#productSet,
				stock: this.#stock,
				payments: this.#payments,
				sellingMode: this.#sellingMode,
				delivery: this.#delivery,
				publication: this.#publication,
				additionalMarketplaces: this.#additionalMarketplaces,
				b2b: this.#b2b,
				compatibilityList: this.#compatibilityList,
				validation: this.#validation,
				warnings: this.#warnings,
				afterSalesServices: this.#afterSalesServices,
				discounts: this.#discounts,
				contact: this.#contact,
				attachments: this.#attachments,
				fundraisingCampaign: this.#fundraisingCampaign,
				additionalServices: this.#additionalServices,
				sizeTable: this.#sizeTable,
				location: this.#location,
				external: this.#external,
				taxSettings: this.#taxSettings,
				messageToSellerSettings: this.#messageToSellerSettings,
				createdAt: this.#createdAt,
				updatedAt: this.#updatedAt,
				images: this.#images,
				description: this.#description,
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
			category$: 'category',
get category() {
					return withPrefix(
						"category",
						EditAnOfferActionRes.Category.Fields
						);
						},
			productSet$: 'productSet',
get productSet() {
					return withPrefix(
						"productSet[:i]",
						EditAnOfferActionRes.ProductSet.Fields
						);
						},
			stock$: 'stock',
get stock() {
					return withPrefix(
						"stock",
						EditAnOfferActionRes.Stock.Fields
						);
						},
			payments$: 'payments',
get payments() {
					return withPrefix(
						"payments",
						EditAnOfferActionRes.Payments.Fields
						);
						},
			sellingMode$: 'sellingMode',
get sellingMode() {
					return withPrefix(
						"sellingMode",
						EditAnOfferActionRes.SellingMode.Fields
						);
						},
			delivery$: 'delivery',
get delivery() {
					return withPrefix(
						"delivery",
						EditAnOfferActionRes.Delivery.Fields
						);
						},
			publication$: 'publication',
get publication() {
					return withPrefix(
						"publication",
						EditAnOfferActionRes.Publication.Fields
						);
						},
			additionalMarketplaces$: 'additionalMarketplaces',
get additionalMarketplaces() {
					return withPrefix(
						"additionalMarketplaces",
						EditAnOfferActionRes.AdditionalMarketplaces.Fields
						);
						},
			b2b$: 'b2b',
get b2b() {
					return withPrefix(
						"b2b",
						EditAnOfferActionRes.B2b.Fields
						);
						},
			compatibilityList$: 'compatibilityList',
get compatibilityList() {
					return withPrefix(
						"compatibilityList",
						EditAnOfferActionRes.CompatibilityList.Fields
						);
						},
			validation$: 'validation',
get validation() {
					return withPrefix(
						"validation",
						EditAnOfferActionRes.Validation.Fields
						);
						},
			warnings$: 'warnings',
get warnings() {
					return "warnings[:i]";
						},
			afterSalesServices$: 'afterSalesServices',
get afterSalesServices() {
					return withPrefix(
						"afterSalesServices",
						EditAnOfferActionRes.AfterSalesServices.Fields
						);
						},
			discounts$: 'discounts',
get discounts() {
					return withPrefix(
						"discounts",
						EditAnOfferActionRes.Discounts.Fields
						);
						},
			contact$: 'contact',
get contact() {
					return withPrefix(
						"contact",
						EditAnOfferActionRes.Contact.Fields
						);
						},
			attachments$: 'attachments',
get attachments() {
					return withPrefix(
						"attachments[:i]",
						EditAnOfferActionRes.Attachments.Fields
						);
						},
			fundraisingCampaign$: 'fundraisingCampaign',
get fundraisingCampaign() {
					return withPrefix(
						"fundraisingCampaign",
						EditAnOfferActionRes.FundraisingCampaign.Fields
						);
						},
			additionalServices$: 'additionalServices',
get additionalServices() {
					return withPrefix(
						"additionalServices",
						EditAnOfferActionRes.AdditionalServices.Fields
						);
						},
			sizeTable$: 'sizeTable',
get sizeTable() {
					return withPrefix(
						"sizeTable",
						EditAnOfferActionRes.SizeTable.Fields
						);
						},
			location$: 'location',
get location() {
					return withPrefix(
						"location",
						EditAnOfferActionRes.Location.Fields
						);
						},
			external$: 'external',
get external() {
					return withPrefix(
						"external",
						EditAnOfferActionRes.External.Fields
						);
						},
			taxSettings$: 'taxSettings',
get taxSettings() {
					return withPrefix(
						"taxSettings",
						EditAnOfferActionRes.TaxSettings.Fields
						);
						},
			messageToSellerSettings$: 'messageToSellerSettings',
get messageToSellerSettings() {
					return withPrefix(
						"messageToSellerSettings",
						EditAnOfferActionRes.MessageToSellerSettings.Fields
						);
						},
			createdAt: 'createdAt',
			updatedAt: 'updatedAt',
			images$: 'images',
get images() {
					return "images[:i]";
						},
			description$: 'description',
get description() {
					return withPrefix(
						"description",
						EditAnOfferActionRes.Description.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of EditAnOfferActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject) {
		return new EditAnOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject) {
		return new EditAnOfferActionRes(partialDtoObject);
	}
	copyWith(partial) {
		return new EditAnOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone() {
		return new EditAnOfferActionRes(this.toJSON());
	}
}