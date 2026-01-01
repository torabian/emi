import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Edit an offer
*/
export type EditAnOfferActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * EditAnOfferAction
 */
export class EditAnOfferAction { //
  static URL = 'https://api.{environment}/sale/product-offers/{offerId}';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		EditAnOfferAction.URL,
		 undefined,
		qs
	);
  static Method = 'patch';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<EditAnOfferActionReq, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<EditAnOfferActionRes, EditAnOfferActionReq, unknown>(
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
		init?: TypedRequestInit<EditAnOfferActionReq, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => EditAnOfferActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
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
  * @type {string}
  **/
 #language : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get language () { return this.#language }
/**
  * 
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
  * @type {EditAnOfferActionReq.Category}
  **/
 #category ! : InstanceType<typeof EditAnOfferActionReq.Category>
		/**
  * 
  * @returns {EditAnOfferActionReq.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {EditAnOfferActionReq.Category}
  **/
set category (value: InstanceType<typeof EditAnOfferActionReq.Category>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Category) {
			this.#category = value
		} else {
			this.#category = new EditAnOfferActionReq.Category(value)
		}
}
setCategory (value: InstanceType<typeof EditAnOfferActionReq.Category>) {
	this.category = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet}
  **/
 #productSet : InstanceType<typeof EditAnOfferActionReq.ProductSet>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet}
  **/
get productSet () { return this.#productSet }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet}
  **/
set productSet (value: InstanceType<typeof EditAnOfferActionReq.ProductSet>[]) {
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
setProductSet (value: InstanceType<typeof EditAnOfferActionReq.ProductSet>[]) {
	this.productSet = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Stock}
  **/
 #stock ! : InstanceType<typeof EditAnOfferActionReq.Stock>
		/**
  * 
  * @returns {EditAnOfferActionReq.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {EditAnOfferActionReq.Stock}
  **/
set stock (value: InstanceType<typeof EditAnOfferActionReq.Stock>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Stock) {
			this.#stock = value
		} else {
			this.#stock = new EditAnOfferActionReq.Stock(value)
		}
}
setStock (value: InstanceType<typeof EditAnOfferActionReq.Stock>) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.SellingMode}
  **/
 #sellingMode ! : InstanceType<typeof EditAnOfferActionReq.SellingMode>
		/**
  * 
  * @returns {EditAnOfferActionReq.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {EditAnOfferActionReq.SellingMode}
  **/
set sellingMode (value: InstanceType<typeof EditAnOfferActionReq.SellingMode>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new EditAnOfferActionReq.SellingMode(value)
		}
}
setSellingMode (value: InstanceType<typeof EditAnOfferActionReq.SellingMode>) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Payments}
  **/
 #payments ! : InstanceType<typeof EditAnOfferActionReq.Payments>
		/**
  * 
  * @returns {EditAnOfferActionReq.Payments}
  **/
get payments () { return this.#payments }
/**
  * 
  * @type {EditAnOfferActionReq.Payments}
  **/
set payments (value: InstanceType<typeof EditAnOfferActionReq.Payments>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Payments) {
			this.#payments = value
		} else {
			this.#payments = new EditAnOfferActionReq.Payments(value)
		}
}
setPayments (value: InstanceType<typeof EditAnOfferActionReq.Payments>) {
	this.payments = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Delivery}
  **/
 #delivery ! : InstanceType<typeof EditAnOfferActionReq.Delivery>
		/**
  * 
  * @returns {EditAnOfferActionReq.Delivery}
  **/
get delivery () { return this.#delivery }
/**
  * 
  * @type {EditAnOfferActionReq.Delivery}
  **/
set delivery (value: InstanceType<typeof EditAnOfferActionReq.Delivery>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Delivery) {
			this.#delivery = value
		} else {
			this.#delivery = new EditAnOfferActionReq.Delivery(value)
		}
}
setDelivery (value: InstanceType<typeof EditAnOfferActionReq.Delivery>) {
	this.delivery = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Publication}
  **/
 #publication ! : InstanceType<typeof EditAnOfferActionReq.Publication>
		/**
  * 
  * @returns {EditAnOfferActionReq.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {EditAnOfferActionReq.Publication}
  **/
set publication (value: InstanceType<typeof EditAnOfferActionReq.Publication>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Publication) {
			this.#publication = value
		} else {
			this.#publication = new EditAnOfferActionReq.Publication(value)
		}
}
setPublication (value: InstanceType<typeof EditAnOfferActionReq.Publication>) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {{[key: string]: any}}
  **/
 #additionalMarketplaces ! : {[key: string]: any}
		/**
  * 
  * @returns {{[key: string]: any}}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {{[key: string]: any}}
  **/
set additionalMarketplaces (value: {[key: string]: any}) {
}
setAdditionalMarketplaces (value: {[key: string]: any}) {
	this.additionalMarketplaces = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.CompatibilityList}
  **/
 #compatibilityList ! : InstanceType<typeof EditAnOfferActionReq.CompatibilityList>
		/**
  * 
  * @returns {EditAnOfferActionReq.CompatibilityList}
  **/
get compatibilityList () { return this.#compatibilityList }
/**
  * 
  * @type {EditAnOfferActionReq.CompatibilityList}
  **/
set compatibilityList (value: InstanceType<typeof EditAnOfferActionReq.CompatibilityList>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.CompatibilityList) {
			this.#compatibilityList = value
		} else {
			this.#compatibilityList = new EditAnOfferActionReq.CompatibilityList(value)
		}
}
setCompatibilityList (value: InstanceType<typeof EditAnOfferActionReq.CompatibilityList>) {
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
  * @type {EditAnOfferActionReq.Description}
  **/
 #description ! : InstanceType<typeof EditAnOfferActionReq.Description>
		/**
  * 
  * @returns {EditAnOfferActionReq.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {EditAnOfferActionReq.Description}
  **/
set description (value: InstanceType<typeof EditAnOfferActionReq.Description>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Description) {
			this.#description = value
		} else {
			this.#description = new EditAnOfferActionReq.Description(value)
		}
}
setDescription (value: InstanceType<typeof EditAnOfferActionReq.Description>) {
	this.description = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.B2b}
  **/
 #b2b ! : InstanceType<typeof EditAnOfferActionReq.B2b>
		/**
  * 
  * @returns {EditAnOfferActionReq.B2b}
  **/
get b2b () { return this.#b2b }
/**
  * 
  * @type {EditAnOfferActionReq.B2b}
  **/
set b2b (value: InstanceType<typeof EditAnOfferActionReq.B2b>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.B2b) {
			this.#b2b = value
		} else {
			this.#b2b = new EditAnOfferActionReq.B2b(value)
		}
}
setB2b (value: InstanceType<typeof EditAnOfferActionReq.B2b>) {
	this.b2b = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Attachments}
  **/
 #attachments : InstanceType<typeof EditAnOfferActionReq.Attachments>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.Attachments}
  **/
get attachments () { return this.#attachments }
/**
  * 
  * @type {EditAnOfferActionReq.Attachments}
  **/
set attachments (value: InstanceType<typeof EditAnOfferActionReq.Attachments>[]) {
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
setAttachments (value: InstanceType<typeof EditAnOfferActionReq.Attachments>[]) {
	this.attachments = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.FundraisingCampaign}
  **/
 #fundraisingCampaign ! : InstanceType<typeof EditAnOfferActionReq.FundraisingCampaign>
		/**
  * 
  * @returns {EditAnOfferActionReq.FundraisingCampaign}
  **/
get fundraisingCampaign () { return this.#fundraisingCampaign }
/**
  * 
  * @type {EditAnOfferActionReq.FundraisingCampaign}
  **/
set fundraisingCampaign (value: InstanceType<typeof EditAnOfferActionReq.FundraisingCampaign>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.FundraisingCampaign) {
			this.#fundraisingCampaign = value
		} else {
			this.#fundraisingCampaign = new EditAnOfferActionReq.FundraisingCampaign(value)
		}
}
setFundraisingCampaign (value: InstanceType<typeof EditAnOfferActionReq.FundraisingCampaign>) {
	this.fundraisingCampaign = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.AdditionalServices}
  **/
 #additionalServices ! : InstanceType<typeof EditAnOfferActionReq.AdditionalServices>
		/**
  * 
  * @returns {EditAnOfferActionReq.AdditionalServices}
  **/
get additionalServices () { return this.#additionalServices }
/**
  * 
  * @type {EditAnOfferActionReq.AdditionalServices}
  **/
set additionalServices (value: InstanceType<typeof EditAnOfferActionReq.AdditionalServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AdditionalServices) {
			this.#additionalServices = value
		} else {
			this.#additionalServices = new EditAnOfferActionReq.AdditionalServices(value)
		}
}
setAdditionalServices (value: InstanceType<typeof EditAnOfferActionReq.AdditionalServices>) {
	this.additionalServices = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices}
  **/
 #afterSalesServices ! : InstanceType<typeof EditAnOfferActionReq.AfterSalesServices>
		/**
  * 
  * @returns {EditAnOfferActionReq.AfterSalesServices}
  **/
get afterSalesServices () { return this.#afterSalesServices }
/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices}
  **/
set afterSalesServices (value: InstanceType<typeof EditAnOfferActionReq.AfterSalesServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AfterSalesServices) {
			this.#afterSalesServices = value
		} else {
			this.#afterSalesServices = new EditAnOfferActionReq.AfterSalesServices(value)
		}
}
setAfterSalesServices (value: InstanceType<typeof EditAnOfferActionReq.AfterSalesServices>) {
	this.afterSalesServices = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.SizeTable}
  **/
 #sizeTable ! : InstanceType<typeof EditAnOfferActionReq.SizeTable>
		/**
  * 
  * @returns {EditAnOfferActionReq.SizeTable}
  **/
get sizeTable () { return this.#sizeTable }
/**
  * 
  * @type {EditAnOfferActionReq.SizeTable}
  **/
set sizeTable (value: InstanceType<typeof EditAnOfferActionReq.SizeTable>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SizeTable) {
			this.#sizeTable = value
		} else {
			this.#sizeTable = new EditAnOfferActionReq.SizeTable(value)
		}
}
setSizeTable (value: InstanceType<typeof EditAnOfferActionReq.SizeTable>) {
	this.sizeTable = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Contact}
  **/
 #contact ! : InstanceType<typeof EditAnOfferActionReq.Contact>
		/**
  * 
  * @returns {EditAnOfferActionReq.Contact}
  **/
get contact () { return this.#contact }
/**
  * 
  * @type {EditAnOfferActionReq.Contact}
  **/
set contact (value: InstanceType<typeof EditAnOfferActionReq.Contact>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Contact) {
			this.#contact = value
		} else {
			this.#contact = new EditAnOfferActionReq.Contact(value)
		}
}
setContact (value: InstanceType<typeof EditAnOfferActionReq.Contact>) {
	this.contact = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Discounts}
  **/
 #discounts ! : InstanceType<typeof EditAnOfferActionReq.Discounts>
		/**
  * 
  * @returns {EditAnOfferActionReq.Discounts}
  **/
get discounts () { return this.#discounts }
/**
  * 
  * @type {EditAnOfferActionReq.Discounts}
  **/
set discounts (value: InstanceType<typeof EditAnOfferActionReq.Discounts>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Discounts) {
			this.#discounts = value
		} else {
			this.#discounts = new EditAnOfferActionReq.Discounts(value)
		}
}
setDiscounts (value: InstanceType<typeof EditAnOfferActionReq.Discounts>) {
	this.discounts = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.Location}
  **/
 #location ! : InstanceType<typeof EditAnOfferActionReq.Location>
		/**
  * 
  * @returns {EditAnOfferActionReq.Location}
  **/
get location () { return this.#location }
/**
  * 
  * @type {EditAnOfferActionReq.Location}
  **/
set location (value: InstanceType<typeof EditAnOfferActionReq.Location>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Location) {
			this.#location = value
		} else {
			this.#location = new EditAnOfferActionReq.Location(value)
		}
}
setLocation (value: InstanceType<typeof EditAnOfferActionReq.Location>) {
	this.location = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.External}
  **/
 #external ! : InstanceType<typeof EditAnOfferActionReq.External>
		/**
  * 
  * @returns {EditAnOfferActionReq.External}
  **/
get external () { return this.#external }
/**
  * 
  * @type {EditAnOfferActionReq.External}
  **/
set external (value: InstanceType<typeof EditAnOfferActionReq.External>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.External) {
			this.#external = value
		} else {
			this.#external = new EditAnOfferActionReq.External(value)
		}
}
setExternal (value: InstanceType<typeof EditAnOfferActionReq.External>) {
	this.external = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.TaxSettings}
  **/
 #taxSettings ! : InstanceType<typeof EditAnOfferActionReq.TaxSettings>
		/**
  * 
  * @returns {EditAnOfferActionReq.TaxSettings}
  **/
get taxSettings () { return this.#taxSettings }
/**
  * 
  * @type {EditAnOfferActionReq.TaxSettings}
  **/
set taxSettings (value: InstanceType<typeof EditAnOfferActionReq.TaxSettings>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.TaxSettings) {
			this.#taxSettings = value
		} else {
			this.#taxSettings = new EditAnOfferActionReq.TaxSettings(value)
		}
}
setTaxSettings (value: InstanceType<typeof EditAnOfferActionReq.TaxSettings>) {
	this.taxSettings = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.MessageToSellerSettings}
  **/
 #messageToSellerSettings ! : InstanceType<typeof EditAnOfferActionReq.MessageToSellerSettings>
		/**
  * 
  * @returns {EditAnOfferActionReq.MessageToSellerSettings}
  **/
get messageToSellerSettings () { return this.#messageToSellerSettings }
/**
  * 
  * @type {EditAnOfferActionReq.MessageToSellerSettings}
  **/
set messageToSellerSettings (value: InstanceType<typeof EditAnOfferActionReq.MessageToSellerSettings>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.MessageToSellerSettings) {
			this.#messageToSellerSettings = value
		} else {
			this.#messageToSellerSettings = new EditAnOfferActionReq.MessageToSellerSettings(value)
		}
}
setMessageToSellerSettings (value: InstanceType<typeof EditAnOfferActionReq.MessageToSellerSettings>) {
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
	* Creates an instance of EditAnOfferActionReq.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.CategoryType) {
		return new EditAnOfferActionReq.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.CategoryType>) {
		return new EditAnOfferActionReq.Category(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.CategoryType>): InstanceType<typeof EditAnOfferActionReq.Category> {
		return new EditAnOfferActionReq.Category ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Category> {
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
 #product ! : InstanceType<typeof EditAnOfferActionReq.ProductSet.Product>
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Product}
  **/
get product () { return this.#product }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product}
  **/
set product (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Product>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.Product) {
			this.#product = value
		} else {
			this.#product = new EditAnOfferActionReq.ProductSet.Product(value)
		}
}
setProduct (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Product>) {
	this.product = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Quantity}
  **/
 #quantity ! : InstanceType<typeof EditAnOfferActionReq.ProductSet.Quantity>
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Quantity}
  **/
get quantity () { return this.#quantity }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Quantity}
  **/
set quantity (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Quantity>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.Quantity) {
			this.#quantity = value
		} else {
			this.#quantity = new EditAnOfferActionReq.ProductSet.Quantity(value)
		}
}
setQuantity (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Quantity>) {
	this.quantity = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.ResponsiblePerson}
  **/
 #responsiblePerson ! : InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsiblePerson>
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.ResponsiblePerson}
  **/
get responsiblePerson () { return this.#responsiblePerson }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.ResponsiblePerson}
  **/
set responsiblePerson (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsiblePerson>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.ResponsiblePerson) {
			this.#responsiblePerson = value
		} else {
			this.#responsiblePerson = new EditAnOfferActionReq.ProductSet.ResponsiblePerson(value)
		}
}
setResponsiblePerson (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsiblePerson>) {
	this.responsiblePerson = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.ResponsibleProducer}
  **/
 #responsibleProducer ! : InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsibleProducer>
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.ResponsibleProducer}
  **/
get responsibleProducer () { return this.#responsibleProducer }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.ResponsibleProducer}
  **/
set responsibleProducer (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsibleProducer>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.ResponsibleProducer) {
			this.#responsibleProducer = value
		} else {
			this.#responsibleProducer = new EditAnOfferActionReq.ProductSet.ResponsibleProducer(value)
		}
}
setResponsibleProducer (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsibleProducer>) {
	this.responsibleProducer = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.SafetyInformation}
  **/
 #safetyInformation ! : InstanceType<typeof EditAnOfferActionReq.ProductSet.SafetyInformation>
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.SafetyInformation}
  **/
set safetyInformation (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.SafetyInformation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new EditAnOfferActionReq.ProductSet.SafetyInformation(value)
		}
}
setSafetyInformation (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.SafetyInformation>) {
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
  * @type {EditAnOfferActionReq.ProductSet.Deposits}
  **/
 #deposits : InstanceType<typeof EditAnOfferActionReq.ProductSet.Deposits>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Deposits}
  **/
get deposits () { return this.#deposits }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Deposits}
  **/
set deposits (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Deposits>[]) {
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
setDeposits (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Deposits>[]) {
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
  * @type {EditAnOfferActionReq.ProductSet.Product.Category}
  **/
 #category ! : InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Category>
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Product.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product.Category}
  **/
set category (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Category>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.Product.Category) {
			this.#category = value
		} else {
			this.#category = new EditAnOfferActionReq.ProductSet.Product.Category(value)
		}
}
setCategory (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Category>) {
	this.category = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product.Parameters}
  **/
 #parameters : InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Product.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product.Parameters}
  **/
set parameters (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters>[]) {
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
setParameters (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters>[]) {
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType.ProductType.CategoryType) {
		return new EditAnOfferActionReq.ProductSet.Product.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType.ProductType.CategoryType>) {
		return new EditAnOfferActionReq.ProductSet.Product.Category(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType.ProductType.CategoryType>): InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Category> {
		return new EditAnOfferActionReq.ProductSet.Product.Category ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Category> {
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
  * @type {EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue}
  **/
 #rangeValue ! : InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue>
		/**
  * 
  * @returns {EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue}
  **/
get rangeValue () { return this.#rangeValue }
/**
  * 
  * @type {EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue}
  **/
set rangeValue (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue) {
			this.#rangeValue = value
		} else {
			this.#rangeValue = new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue(value)
		}
}
setRangeValue (value: InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue>) {
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType>) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType>): InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue> {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue> {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters.RangeValue(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType>) {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType>): InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters> {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet.Product.Parameters> {
		return new EditAnOfferActionReq.ProductSet.Product.Parameters(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType.ProductType) {
		return new EditAnOfferActionReq.ProductSet.Product(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Product, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType.ProductType>) {
		return new EditAnOfferActionReq.ProductSet.Product(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType.ProductType>): InstanceType<typeof EditAnOfferActionReq.ProductSet.Product> {
		return new EditAnOfferActionReq.ProductSet.Product ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet.Product> {
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.Quantity, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType.QuantityType) {
		return new EditAnOfferActionReq.ProductSet.Quantity(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Quantity, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType.QuantityType>) {
		return new EditAnOfferActionReq.ProductSet.Quantity(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType.QuantityType>): InstanceType<typeof EditAnOfferActionReq.ProductSet.Quantity> {
		return new EditAnOfferActionReq.ProductSet.Quantity ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet.Quantity> {
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.ResponsiblePerson, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType.ResponsiblePersonType) {
		return new EditAnOfferActionReq.ProductSet.ResponsiblePerson(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.ResponsiblePerson, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType.ResponsiblePersonType>) {
		return new EditAnOfferActionReq.ProductSet.ResponsiblePerson(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType.ResponsiblePersonType>): InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsiblePerson> {
		return new EditAnOfferActionReq.ProductSet.ResponsiblePerson ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsiblePerson> {
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.ResponsibleProducer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType.ResponsibleProducerType) {
		return new EditAnOfferActionReq.ProductSet.ResponsibleProducer(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.ResponsibleProducer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType.ResponsibleProducerType>) {
		return new EditAnOfferActionReq.ProductSet.ResponsibleProducer(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType.ResponsibleProducerType>): InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsibleProducer> {
		return new EditAnOfferActionReq.ProductSet.ResponsibleProducer ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet.ResponsibleProducer> {
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.SafetyInformation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType.SafetyInformationType) {
		return new EditAnOfferActionReq.ProductSet.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType.SafetyInformationType>) {
		return new EditAnOfferActionReq.ProductSet.SafetyInformation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType.SafetyInformationType>): InstanceType<typeof EditAnOfferActionReq.ProductSet.SafetyInformation> {
		return new EditAnOfferActionReq.ProductSet.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet.SafetyInformation> {
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
	* Creates an instance of EditAnOfferActionReq.ProductSet.Deposits, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType.DepositsType) {
		return new EditAnOfferActionReq.ProductSet.Deposits(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet.Deposits, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType.DepositsType>) {
		return new EditAnOfferActionReq.ProductSet.Deposits(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType.DepositsType>): InstanceType<typeof EditAnOfferActionReq.ProductSet.Deposits> {
		return new EditAnOfferActionReq.ProductSet.Deposits ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet.Deposits> {
		return new EditAnOfferActionReq.ProductSet.Deposits(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.ProductSetType) {
		return new EditAnOfferActionReq.ProductSet(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.ProductSet, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ProductSetType>) {
		return new EditAnOfferActionReq.ProductSet(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ProductSetType>): InstanceType<typeof EditAnOfferActionReq.ProductSet> {
		return new EditAnOfferActionReq.ProductSet ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.ProductSet> {
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
	* Creates an instance of EditAnOfferActionReq.Stock, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.StockType) {
		return new EditAnOfferActionReq.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.StockType>) {
		return new EditAnOfferActionReq.Stock(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.StockType>): InstanceType<typeof EditAnOfferActionReq.Stock> {
		return new EditAnOfferActionReq.Stock ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Stock> {
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
  * @type {EditAnOfferActionReq.SellingMode.Price}
  **/
 #price ! : InstanceType<typeof EditAnOfferActionReq.SellingMode.Price>
		/**
  * 
  * @returns {EditAnOfferActionReq.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.Price}
  **/
set price (value: InstanceType<typeof EditAnOfferActionReq.SellingMode.Price>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new EditAnOfferActionReq.SellingMode.Price(value)
		}
}
setPrice (value: InstanceType<typeof EditAnOfferActionReq.SellingMode.Price>) {
	this.price = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.MinimalPrice}
  **/
 #minimalPrice ! : InstanceType<typeof EditAnOfferActionReq.SellingMode.MinimalPrice>
		/**
  * 
  * @returns {EditAnOfferActionReq.SellingMode.MinimalPrice}
  **/
get minimalPrice () { return this.#minimalPrice }
/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.MinimalPrice}
  **/
set minimalPrice (value: InstanceType<typeof EditAnOfferActionReq.SellingMode.MinimalPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SellingMode.MinimalPrice) {
			this.#minimalPrice = value
		} else {
			this.#minimalPrice = new EditAnOfferActionReq.SellingMode.MinimalPrice(value)
		}
}
setMinimalPrice (value: InstanceType<typeof EditAnOfferActionReq.SellingMode.MinimalPrice>) {
	this.minimalPrice = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.StartingPrice}
  **/
 #startingPrice ! : InstanceType<typeof EditAnOfferActionReq.SellingMode.StartingPrice>
		/**
  * 
  * @returns {EditAnOfferActionReq.SellingMode.StartingPrice}
  **/
get startingPrice () { return this.#startingPrice }
/**
  * 
  * @type {EditAnOfferActionReq.SellingMode.StartingPrice}
  **/
set startingPrice (value: InstanceType<typeof EditAnOfferActionReq.SellingMode.StartingPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.SellingMode.StartingPrice) {
			this.#startingPrice = value
		} else {
			this.#startingPrice = new EditAnOfferActionReq.SellingMode.StartingPrice(value)
		}
}
setStartingPrice (value: InstanceType<typeof EditAnOfferActionReq.SellingMode.StartingPrice>) {
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
	* Creates an instance of EditAnOfferActionReq.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.SellingModeType.PriceType) {
		return new EditAnOfferActionReq.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.SellingModeType.PriceType>) {
		return new EditAnOfferActionReq.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.SellingModeType.PriceType>): InstanceType<typeof EditAnOfferActionReq.SellingMode.Price> {
		return new EditAnOfferActionReq.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.SellingMode.Price> {
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
	* Creates an instance of EditAnOfferActionReq.SellingMode.MinimalPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.SellingModeType.MinimalPriceType) {
		return new EditAnOfferActionReq.SellingMode.MinimalPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SellingMode.MinimalPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.SellingModeType.MinimalPriceType>) {
		return new EditAnOfferActionReq.SellingMode.MinimalPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.SellingModeType.MinimalPriceType>): InstanceType<typeof EditAnOfferActionReq.SellingMode.MinimalPrice> {
		return new EditAnOfferActionReq.SellingMode.MinimalPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.SellingMode.MinimalPrice> {
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
	* Creates an instance of EditAnOfferActionReq.SellingMode.StartingPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.SellingModeType.StartingPriceType) {
		return new EditAnOfferActionReq.SellingMode.StartingPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SellingMode.StartingPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.SellingModeType.StartingPriceType>) {
		return new EditAnOfferActionReq.SellingMode.StartingPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.SellingModeType.StartingPriceType>): InstanceType<typeof EditAnOfferActionReq.SellingMode.StartingPrice> {
		return new EditAnOfferActionReq.SellingMode.StartingPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.SellingMode.StartingPrice> {
		return new EditAnOfferActionReq.SellingMode.StartingPrice(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.SellingModeType) {
		return new EditAnOfferActionReq.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.SellingModeType>) {
		return new EditAnOfferActionReq.SellingMode(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.SellingModeType>): InstanceType<typeof EditAnOfferActionReq.SellingMode> {
		return new EditAnOfferActionReq.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.SellingMode> {
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
	* Creates an instance of EditAnOfferActionReq.Payments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.PaymentsType) {
		return new EditAnOfferActionReq.Payments(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Payments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.PaymentsType>) {
		return new EditAnOfferActionReq.Payments(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.PaymentsType>): InstanceType<typeof EditAnOfferActionReq.Payments> {
		return new EditAnOfferActionReq.Payments ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Payments> {
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
  * @type {EditAnOfferActionReq.Delivery.ShippingRates}
  **/
 #shippingRates ! : InstanceType<typeof EditAnOfferActionReq.Delivery.ShippingRates>
		/**
  * 
  * @returns {EditAnOfferActionReq.Delivery.ShippingRates}
  **/
get shippingRates () { return this.#shippingRates }
/**
  * 
  * @type {EditAnOfferActionReq.Delivery.ShippingRates}
  **/
set shippingRates (value: InstanceType<typeof EditAnOfferActionReq.Delivery.ShippingRates>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Delivery.ShippingRates) {
			this.#shippingRates = value
		} else {
			this.#shippingRates = new EditAnOfferActionReq.Delivery.ShippingRates(value)
		}
}
setShippingRates (value: InstanceType<typeof EditAnOfferActionReq.Delivery.ShippingRates>) {
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
	* Creates an instance of EditAnOfferActionReq.Delivery.ShippingRates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.DeliveryType.ShippingRatesType) {
		return new EditAnOfferActionReq.Delivery.ShippingRates(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Delivery.ShippingRates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.DeliveryType.ShippingRatesType>) {
		return new EditAnOfferActionReq.Delivery.ShippingRates(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.DeliveryType.ShippingRatesType>): InstanceType<typeof EditAnOfferActionReq.Delivery.ShippingRates> {
		return new EditAnOfferActionReq.Delivery.ShippingRates ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Delivery.ShippingRates> {
		return new EditAnOfferActionReq.Delivery.ShippingRates(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.DeliveryType) {
		return new EditAnOfferActionReq.Delivery(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Delivery, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.DeliveryType>) {
		return new EditAnOfferActionReq.Delivery(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.DeliveryType>): InstanceType<typeof EditAnOfferActionReq.Delivery> {
		return new EditAnOfferActionReq.Delivery ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Delivery> {
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
	* Creates an instance of EditAnOfferActionReq.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.PublicationType) {
		return new EditAnOfferActionReq.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.PublicationType>) {
		return new EditAnOfferActionReq.Publication(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.PublicationType>): InstanceType<typeof EditAnOfferActionReq.Publication> {
		return new EditAnOfferActionReq.Publication ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Publication> {
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
 #items : InstanceType<typeof EditAnOfferActionReq.CompatibilityList.Items>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.CompatibilityList.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {EditAnOfferActionReq.CompatibilityList.Items}
  **/
set items (value: InstanceType<typeof EditAnOfferActionReq.CompatibilityList.Items>[]) {
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
setItems (value: InstanceType<typeof EditAnOfferActionReq.CompatibilityList.Items>[]) {
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
	* Creates an instance of EditAnOfferActionReq.CompatibilityList.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.CompatibilityListType.ItemsType) {
		return new EditAnOfferActionReq.CompatibilityList.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.CompatibilityList.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.CompatibilityListType.ItemsType>) {
		return new EditAnOfferActionReq.CompatibilityList.Items(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.CompatibilityListType.ItemsType>): InstanceType<typeof EditAnOfferActionReq.CompatibilityList.Items> {
		return new EditAnOfferActionReq.CompatibilityList.Items ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.CompatibilityList.Items> {
		return new EditAnOfferActionReq.CompatibilityList.Items(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.CompatibilityListType) {
		return new EditAnOfferActionReq.CompatibilityList(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.CompatibilityList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.CompatibilityListType>) {
		return new EditAnOfferActionReq.CompatibilityList(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.CompatibilityListType>): InstanceType<typeof EditAnOfferActionReq.CompatibilityList> {
		return new EditAnOfferActionReq.CompatibilityList ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.CompatibilityList> {
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
 #sections : InstanceType<typeof EditAnOfferActionReq.Description.Sections>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.Description.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {EditAnOfferActionReq.Description.Sections}
  **/
set sections (value: InstanceType<typeof EditAnOfferActionReq.Description.Sections>[]) {
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
setSections (value: InstanceType<typeof EditAnOfferActionReq.Description.Sections>[]) {
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
 #items : InstanceType<typeof EditAnOfferActionReq.Description.Sections.Items>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.Description.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {EditAnOfferActionReq.Description.Sections.Items}
  **/
set items (value: InstanceType<typeof EditAnOfferActionReq.Description.Sections.Items>[]) {
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
setItems (value: InstanceType<typeof EditAnOfferActionReq.Description.Sections.Items>[]) {
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
	* Creates an instance of EditAnOfferActionReq.Description.Sections.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.DescriptionType.SectionsType.ItemsType) {
		return new EditAnOfferActionReq.Description.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Description.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.DescriptionType.SectionsType.ItemsType>) {
		return new EditAnOfferActionReq.Description.Sections.Items(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.DescriptionType.SectionsType.ItemsType>): InstanceType<typeof EditAnOfferActionReq.Description.Sections.Items> {
		return new EditAnOfferActionReq.Description.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Description.Sections.Items> {
		return new EditAnOfferActionReq.Description.Sections.Items(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.DescriptionType.SectionsType) {
		return new EditAnOfferActionReq.Description.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Description.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.DescriptionType.SectionsType>) {
		return new EditAnOfferActionReq.Description.Sections(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.DescriptionType.SectionsType>): InstanceType<typeof EditAnOfferActionReq.Description.Sections> {
		return new EditAnOfferActionReq.Description.Sections ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Description.Sections> {
		return new EditAnOfferActionReq.Description.Sections(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.DescriptionType) {
		return new EditAnOfferActionReq.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.DescriptionType>) {
		return new EditAnOfferActionReq.Description(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.DescriptionType>): InstanceType<typeof EditAnOfferActionReq.Description> {
		return new EditAnOfferActionReq.Description ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Description> {
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
	* Creates an instance of EditAnOfferActionReq.B2b, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.B2bType) {
		return new EditAnOfferActionReq.B2b(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.B2b, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.B2bType>) {
		return new EditAnOfferActionReq.B2b(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.B2bType>): InstanceType<typeof EditAnOfferActionReq.B2b> {
		return new EditAnOfferActionReq.B2b ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.B2b> {
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
	* Creates an instance of EditAnOfferActionReq.Attachments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.AttachmentsType) {
		return new EditAnOfferActionReq.Attachments(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Attachments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.AttachmentsType>) {
		return new EditAnOfferActionReq.Attachments(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.AttachmentsType>): InstanceType<typeof EditAnOfferActionReq.Attachments> {
		return new EditAnOfferActionReq.Attachments ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Attachments> {
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
	* Creates an instance of EditAnOfferActionReq.FundraisingCampaign, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.FundraisingCampaignType) {
		return new EditAnOfferActionReq.FundraisingCampaign(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.FundraisingCampaign, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.FundraisingCampaignType>) {
		return new EditAnOfferActionReq.FundraisingCampaign(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.FundraisingCampaignType>): InstanceType<typeof EditAnOfferActionReq.FundraisingCampaign> {
		return new EditAnOfferActionReq.FundraisingCampaign ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.FundraisingCampaign> {
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
	* Creates an instance of EditAnOfferActionReq.AdditionalServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.AdditionalServicesType) {
		return new EditAnOfferActionReq.AdditionalServices(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AdditionalServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.AdditionalServicesType>) {
		return new EditAnOfferActionReq.AdditionalServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.AdditionalServicesType>): InstanceType<typeof EditAnOfferActionReq.AdditionalServices> {
		return new EditAnOfferActionReq.AdditionalServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.AdditionalServices> {
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
 #impliedWarranty ! : InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty>
		/**
  * 
  * @returns {EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty}
  **/
get impliedWarranty () { return this.#impliedWarranty }
/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty}
  **/
set impliedWarranty (value: InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty) {
			this.#impliedWarranty = value
		} else {
			this.#impliedWarranty = new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty(value)
		}
}
setImpliedWarranty (value: InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty>) {
	this.impliedWarranty = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.ReturnPolicy}
  **/
 #returnPolicy ! : InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ReturnPolicy>
		/**
  * 
  * @returns {EditAnOfferActionReq.AfterSalesServices.ReturnPolicy}
  **/
get returnPolicy () { return this.#returnPolicy }
/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.ReturnPolicy}
  **/
set returnPolicy (value: InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ReturnPolicy>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AfterSalesServices.ReturnPolicy) {
			this.#returnPolicy = value
		} else {
			this.#returnPolicy = new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy(value)
		}
}
setReturnPolicy (value: InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ReturnPolicy>) {
	this.returnPolicy = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.Warranty}
  **/
 #warranty ! : InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.Warranty>
		/**
  * 
  * @returns {EditAnOfferActionReq.AfterSalesServices.Warranty}
  **/
get warranty () { return this.#warranty }
/**
  * 
  * @type {EditAnOfferActionReq.AfterSalesServices.Warranty}
  **/
set warranty (value: InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.Warranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.AfterSalesServices.Warranty) {
			this.#warranty = value
		} else {
			this.#warranty = new EditAnOfferActionReq.AfterSalesServices.Warranty(value)
		}
}
setWarranty (value: InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.Warranty>) {
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
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.AfterSalesServicesType.ImpliedWarrantyType) {
		return new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.AfterSalesServicesType.ImpliedWarrantyType>) {
		return new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.AfterSalesServicesType.ImpliedWarrantyType>): InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty> {
		return new EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ImpliedWarranty> {
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
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.ReturnPolicy, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.AfterSalesServicesType.ReturnPolicyType) {
		return new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.ReturnPolicy, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.AfterSalesServicesType.ReturnPolicyType>) {
		return new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.AfterSalesServicesType.ReturnPolicyType>): InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ReturnPolicy> {
		return new EditAnOfferActionReq.AfterSalesServices.ReturnPolicy ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.ReturnPolicy> {
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
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.Warranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.AfterSalesServicesType.WarrantyType) {
		return new EditAnOfferActionReq.AfterSalesServices.Warranty(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices.Warranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.AfterSalesServicesType.WarrantyType>) {
		return new EditAnOfferActionReq.AfterSalesServices.Warranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.AfterSalesServicesType.WarrantyType>): InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.Warranty> {
		return new EditAnOfferActionReq.AfterSalesServices.Warranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.AfterSalesServices.Warranty> {
		return new EditAnOfferActionReq.AfterSalesServices.Warranty(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.AfterSalesServicesType) {
		return new EditAnOfferActionReq.AfterSalesServices(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.AfterSalesServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.AfterSalesServicesType>) {
		return new EditAnOfferActionReq.AfterSalesServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.AfterSalesServicesType>): InstanceType<typeof EditAnOfferActionReq.AfterSalesServices> {
		return new EditAnOfferActionReq.AfterSalesServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.AfterSalesServices> {
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
	* Creates an instance of EditAnOfferActionReq.SizeTable, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.SizeTableType) {
		return new EditAnOfferActionReq.SizeTable(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.SizeTable, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.SizeTableType>) {
		return new EditAnOfferActionReq.SizeTable(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.SizeTableType>): InstanceType<typeof EditAnOfferActionReq.SizeTable> {
		return new EditAnOfferActionReq.SizeTable ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.SizeTable> {
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
	* Creates an instance of EditAnOfferActionReq.Contact, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.ContactType) {
		return new EditAnOfferActionReq.Contact(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Contact, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ContactType>) {
		return new EditAnOfferActionReq.Contact(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ContactType>): InstanceType<typeof EditAnOfferActionReq.Contact> {
		return new EditAnOfferActionReq.Contact ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Contact> {
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
 #wholesalePriceList ! : InstanceType<typeof EditAnOfferActionReq.Discounts.WholesalePriceList>
		/**
  * 
  * @returns {EditAnOfferActionReq.Discounts.WholesalePriceList}
  **/
get wholesalePriceList () { return this.#wholesalePriceList }
/**
  * 
  * @type {EditAnOfferActionReq.Discounts.WholesalePriceList}
  **/
set wholesalePriceList (value: InstanceType<typeof EditAnOfferActionReq.Discounts.WholesalePriceList>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionReq.Discounts.WholesalePriceList) {
			this.#wholesalePriceList = value
		} else {
			this.#wholesalePriceList = new EditAnOfferActionReq.Discounts.WholesalePriceList(value)
		}
}
setWholesalePriceList (value: InstanceType<typeof EditAnOfferActionReq.Discounts.WholesalePriceList>) {
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
	* Creates an instance of EditAnOfferActionReq.Discounts.WholesalePriceList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.DiscountsType.WholesalePriceListType) {
		return new EditAnOfferActionReq.Discounts.WholesalePriceList(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Discounts.WholesalePriceList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.DiscountsType.WholesalePriceListType>) {
		return new EditAnOfferActionReq.Discounts.WholesalePriceList(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.DiscountsType.WholesalePriceListType>): InstanceType<typeof EditAnOfferActionReq.Discounts.WholesalePriceList> {
		return new EditAnOfferActionReq.Discounts.WholesalePriceList ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Discounts.WholesalePriceList> {
		return new EditAnOfferActionReq.Discounts.WholesalePriceList(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.DiscountsType) {
		return new EditAnOfferActionReq.Discounts(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Discounts, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.DiscountsType>) {
		return new EditAnOfferActionReq.Discounts(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.DiscountsType>): InstanceType<typeof EditAnOfferActionReq.Discounts> {
		return new EditAnOfferActionReq.Discounts ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Discounts> {
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
	* Creates an instance of EditAnOfferActionReq.Location, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.LocationType) {
		return new EditAnOfferActionReq.Location(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.Location, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.LocationType>) {
		return new EditAnOfferActionReq.Location(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.LocationType>): InstanceType<typeof EditAnOfferActionReq.Location> {
		return new EditAnOfferActionReq.Location ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.Location> {
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
	* Creates an instance of EditAnOfferActionReq.External, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.ExternalType) {
		return new EditAnOfferActionReq.External(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.External, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.ExternalType>) {
		return new EditAnOfferActionReq.External(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.ExternalType>): InstanceType<typeof EditAnOfferActionReq.External> {
		return new EditAnOfferActionReq.External ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.External> {
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
  * @type {EditAnOfferActionReq.TaxSettings.Rates}
  **/
 #rates : InstanceType<typeof EditAnOfferActionReq.TaxSettings.Rates>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionReq.TaxSettings.Rates}
  **/
get rates () { return this.#rates }
/**
  * 
  * @type {EditAnOfferActionReq.TaxSettings.Rates}
  **/
set rates (value: InstanceType<typeof EditAnOfferActionReq.TaxSettings.Rates>[]) {
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
setRates (value: InstanceType<typeof EditAnOfferActionReq.TaxSettings.Rates>[]) {
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
	* Creates an instance of EditAnOfferActionReq.TaxSettings.Rates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.TaxSettingsType.RatesType) {
		return new EditAnOfferActionReq.TaxSettings.Rates(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.TaxSettings.Rates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.TaxSettingsType.RatesType>) {
		return new EditAnOfferActionReq.TaxSettings.Rates(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.TaxSettingsType.RatesType>): InstanceType<typeof EditAnOfferActionReq.TaxSettings.Rates> {
		return new EditAnOfferActionReq.TaxSettings.Rates ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.TaxSettings.Rates> {
		return new EditAnOfferActionReq.TaxSettings.Rates(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionReqType.TaxSettingsType) {
		return new EditAnOfferActionReq.TaxSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.TaxSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.TaxSettingsType>) {
		return new EditAnOfferActionReq.TaxSettings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.TaxSettingsType>): InstanceType<typeof EditAnOfferActionReq.TaxSettings> {
		return new EditAnOfferActionReq.TaxSettings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.TaxSettings> {
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
	* Creates an instance of EditAnOfferActionReq.MessageToSellerSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionReqType.MessageToSellerSettingsType) {
		return new EditAnOfferActionReq.MessageToSellerSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq.MessageToSellerSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType.MessageToSellerSettingsType>) {
		return new EditAnOfferActionReq.MessageToSellerSettings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType.MessageToSellerSettingsType>): InstanceType<typeof EditAnOfferActionReq.MessageToSellerSettings> {
		return new EditAnOfferActionReq.MessageToSellerSettings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq.MessageToSellerSettings> {
		return new EditAnOfferActionReq.MessageToSellerSettings(this.toJSON());
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
		const d = data as Partial<EditAnOfferActionReq>;
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
		const d = data as Partial<EditAnOfferActionReq>;
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
	static from(possibleDtoObject: EditAnOfferActionReqType) {
		return new EditAnOfferActionReq(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionReq, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionReqType>) {
		return new EditAnOfferActionReq(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionReqType>): InstanceType<typeof EditAnOfferActionReq> {
		return new EditAnOfferActionReq ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionReq> {
		return new EditAnOfferActionReq(this.toJSON());
	}
}
export abstract class EditAnOfferActionReqFactory {
	abstract create(data: unknown): EditAnOfferActionReq;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for editAnOfferActionReq
  **/
	export type EditAnOfferActionReqType =  {
			/**
  * 
  * @type {string}
  **/
 name : string;
			/**
  * 
  * @type {string}
  **/
 language : string;
			/**
  * 
  * @type {EditAnOfferActionReqType.CategoryType}
  **/
 category : EditAnOfferActionReqType.CategoryType;
			/**
  * 
  * @type {EditAnOfferActionReqType.ProductSetType[]}
  **/
 productSet : EditAnOfferActionReqType.ProductSetType[];
			/**
  * 
  * @type {EditAnOfferActionReqType.StockType}
  **/
 stock : EditAnOfferActionReqType.StockType;
			/**
  * 
  * @type {EditAnOfferActionReqType.SellingModeType}
  **/
 sellingMode : EditAnOfferActionReqType.SellingModeType;
			/**
  * 
  * @type {EditAnOfferActionReqType.PaymentsType}
  **/
 payments : EditAnOfferActionReqType.PaymentsType;
			/**
  * 
  * @type {EditAnOfferActionReqType.DeliveryType}
  **/
 delivery : EditAnOfferActionReqType.DeliveryType;
			/**
  * 
  * @type {EditAnOfferActionReqType.PublicationType}
  **/
 publication : EditAnOfferActionReqType.PublicationType;
			/**
  * 
  * @type {{[key: string]: any}}
  **/
 additionalMarketplaces : {[key: string]: any};
			/**
  * 
  * @type {EditAnOfferActionReqType.CompatibilityListType}
  **/
 compatibilityList : EditAnOfferActionReqType.CompatibilityListType;
			/**
  * 
  * @type {string[]}
  **/
 images : string[];
			/**
  * 
  * @type {EditAnOfferActionReqType.DescriptionType}
  **/
 description : EditAnOfferActionReqType.DescriptionType;
			/**
  * 
  * @type {EditAnOfferActionReqType.B2bType}
  **/
 b2b : EditAnOfferActionReqType.B2bType;
			/**
  * 
  * @type {EditAnOfferActionReqType.AttachmentsType[]}
  **/
 attachments : EditAnOfferActionReqType.AttachmentsType[];
			/**
  * 
  * @type {EditAnOfferActionReqType.FundraisingCampaignType}
  **/
 fundraisingCampaign : EditAnOfferActionReqType.FundraisingCampaignType;
			/**
  * 
  * @type {EditAnOfferActionReqType.AdditionalServicesType}
  **/
 additionalServices : EditAnOfferActionReqType.AdditionalServicesType;
			/**
  * 
  * @type {EditAnOfferActionReqType.AfterSalesServicesType}
  **/
 afterSalesServices : EditAnOfferActionReqType.AfterSalesServicesType;
			/**
  * 
  * @type {EditAnOfferActionReqType.SizeTableType}
  **/
 sizeTable : EditAnOfferActionReqType.SizeTableType;
			/**
  * 
  * @type {EditAnOfferActionReqType.ContactType}
  **/
 contact : EditAnOfferActionReqType.ContactType;
			/**
  * 
  * @type {EditAnOfferActionReqType.DiscountsType}
  **/
 discounts : EditAnOfferActionReqType.DiscountsType;
			/**
  * 
  * @type {EditAnOfferActionReqType.LocationType}
  **/
 location : EditAnOfferActionReqType.LocationType;
			/**
  * 
  * @type {EditAnOfferActionReqType.ExternalType}
  **/
 external : EditAnOfferActionReqType.ExternalType;
			/**
  * 
  * @type {EditAnOfferActionReqType.TaxSettingsType}
  **/
 taxSettings : EditAnOfferActionReqType.TaxSettingsType;
			/**
  * 
  * @type {EditAnOfferActionReqType.MessageToSellerSettingsType}
  **/
 messageToSellerSettings : EditAnOfferActionReqType.MessageToSellerSettingsType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace EditAnOfferActionReqType {
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
  * @type {EditAnOfferActionReqType.ProductSetType.ProductType}
  **/
 product : EditAnOfferActionReqType.ProductSetType.ProductType;
			/**
  * 
  * @type {EditAnOfferActionReqType.ProductSetType.QuantityType}
  **/
 quantity : EditAnOfferActionReqType.ProductSetType.QuantityType;
			/**
  * 
  * @type {EditAnOfferActionReqType.ProductSetType.ResponsiblePersonType}
  **/
 responsiblePerson : EditAnOfferActionReqType.ProductSetType.ResponsiblePersonType;
			/**
  * 
  * @type {EditAnOfferActionReqType.ProductSetType.ResponsibleProducerType}
  **/
 responsibleProducer : EditAnOfferActionReqType.ProductSetType.ResponsibleProducerType;
			/**
  * 
  * @type {EditAnOfferActionReqType.ProductSetType.SafetyInformationType}
  **/
 safetyInformation : EditAnOfferActionReqType.ProductSetType.SafetyInformationType;
			/**
  * 
  * @type {boolean}
  **/
 marketedBeforeGPSRObligation : boolean;
			/**
  * 
  * @type {EditAnOfferActionReqType.ProductSetType.DepositsType[]}
  **/
 deposits : EditAnOfferActionReqType.ProductSetType.DepositsType[];
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
  * @type {EditAnOfferActionReqType.ProductSetType.ProductType.CategoryType}
  **/
 category : EditAnOfferActionReqType.ProductSetType.ProductType.CategoryType;
			/**
  * 
  * @type {EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType[]}
  **/
 parameters : EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType[];
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
  * @type {EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType}
  **/
 rangeValue : EditAnOfferActionReqType.ProductSetType.ProductType.ParametersType.RangeValueType;
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
  * @type {EditAnOfferActionReqType.SellingModeType.PriceType}
  **/
 price : EditAnOfferActionReqType.SellingModeType.PriceType;
			/**
  * 
  * @type {EditAnOfferActionReqType.SellingModeType.MinimalPriceType}
  **/
 minimalPrice : EditAnOfferActionReqType.SellingModeType.MinimalPriceType;
			/**
  * 
  * @type {EditAnOfferActionReqType.SellingModeType.StartingPriceType}
  **/
 startingPrice : EditAnOfferActionReqType.SellingModeType.StartingPriceType;
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
  * @type {EditAnOfferActionReqType.DeliveryType.ShippingRatesType}
  **/
 shippingRates : EditAnOfferActionReqType.DeliveryType.ShippingRatesType;
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
  * @type {EditAnOfferActionReqType.CompatibilityListType.ItemsType[]}
  **/
 items : EditAnOfferActionReqType.CompatibilityListType.ItemsType[];
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
  * @type {EditAnOfferActionReqType.DescriptionType.SectionsType[]}
  **/
 sections : EditAnOfferActionReqType.DescriptionType.SectionsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DescriptionType {
	/**
  * The base type definition for sectionsType
  **/
	export type SectionsType =  {
			/**
  * 
  * @type {EditAnOfferActionReqType.DescriptionType.SectionsType.ItemsType[]}
  **/
 items : EditAnOfferActionReqType.DescriptionType.SectionsType.ItemsType[];
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
  * @type {EditAnOfferActionReqType.AfterSalesServicesType.ImpliedWarrantyType}
  **/
 impliedWarranty : EditAnOfferActionReqType.AfterSalesServicesType.ImpliedWarrantyType;
			/**
  * 
  * @type {EditAnOfferActionReqType.AfterSalesServicesType.ReturnPolicyType}
  **/
 returnPolicy : EditAnOfferActionReqType.AfterSalesServicesType.ReturnPolicyType;
			/**
  * 
  * @type {EditAnOfferActionReqType.AfterSalesServicesType.WarrantyType}
  **/
 warranty : EditAnOfferActionReqType.AfterSalesServicesType.WarrantyType;
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
  * @type {EditAnOfferActionReqType.DiscountsType.WholesalePriceListType}
  **/
 wholesalePriceList : EditAnOfferActionReqType.DiscountsType.WholesalePriceListType;
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
  * @type {EditAnOfferActionReqType.TaxSettingsType.RatesType[]}
  **/
 rates : EditAnOfferActionReqType.TaxSettingsType.RatesType[];
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
/**
  * The base class definition for editAnOfferActionRes
  **/
export class EditAnOfferActionRes {
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
  * @type {string}
  **/
 #language : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get language () { return this.#language }
/**
  * 
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
  * @type {EditAnOfferActionRes.Category}
  **/
 #category ! : InstanceType<typeof EditAnOfferActionRes.Category>
		/**
  * 
  * @returns {EditAnOfferActionRes.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {EditAnOfferActionRes.Category}
  **/
set category (value: InstanceType<typeof EditAnOfferActionRes.Category>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Category) {
			this.#category = value
		} else {
			this.#category = new EditAnOfferActionRes.Category(value)
		}
}
setCategory (value: InstanceType<typeof EditAnOfferActionRes.Category>) {
	this.category = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet}
  **/
 #productSet : InstanceType<typeof EditAnOfferActionRes.ProductSet>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet}
  **/
get productSet () { return this.#productSet }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet}
  **/
set productSet (value: InstanceType<typeof EditAnOfferActionRes.ProductSet>[]) {
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
setProductSet (value: InstanceType<typeof EditAnOfferActionRes.ProductSet>[]) {
	this.productSet = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Stock}
  **/
 #stock ! : InstanceType<typeof EditAnOfferActionRes.Stock>
		/**
  * 
  * @returns {EditAnOfferActionRes.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {EditAnOfferActionRes.Stock}
  **/
set stock (value: InstanceType<typeof EditAnOfferActionRes.Stock>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Stock) {
			this.#stock = value
		} else {
			this.#stock = new EditAnOfferActionRes.Stock(value)
		}
}
setStock (value: InstanceType<typeof EditAnOfferActionRes.Stock>) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Payments}
  **/
 #payments ! : InstanceType<typeof EditAnOfferActionRes.Payments>
		/**
  * 
  * @returns {EditAnOfferActionRes.Payments}
  **/
get payments () { return this.#payments }
/**
  * 
  * @type {EditAnOfferActionRes.Payments}
  **/
set payments (value: InstanceType<typeof EditAnOfferActionRes.Payments>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Payments) {
			this.#payments = value
		} else {
			this.#payments = new EditAnOfferActionRes.Payments(value)
		}
}
setPayments (value: InstanceType<typeof EditAnOfferActionRes.Payments>) {
	this.payments = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.SellingMode}
  **/
 #sellingMode ! : InstanceType<typeof EditAnOfferActionRes.SellingMode>
		/**
  * 
  * @returns {EditAnOfferActionRes.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {EditAnOfferActionRes.SellingMode}
  **/
set sellingMode (value: InstanceType<typeof EditAnOfferActionRes.SellingMode>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new EditAnOfferActionRes.SellingMode(value)
		}
}
setSellingMode (value: InstanceType<typeof EditAnOfferActionRes.SellingMode>) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Delivery}
  **/
 #delivery ! : InstanceType<typeof EditAnOfferActionRes.Delivery>
		/**
  * 
  * @returns {EditAnOfferActionRes.Delivery}
  **/
get delivery () { return this.#delivery }
/**
  * 
  * @type {EditAnOfferActionRes.Delivery}
  **/
set delivery (value: InstanceType<typeof EditAnOfferActionRes.Delivery>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Delivery) {
			this.#delivery = value
		} else {
			this.#delivery = new EditAnOfferActionRes.Delivery(value)
		}
}
setDelivery (value: InstanceType<typeof EditAnOfferActionRes.Delivery>) {
	this.delivery = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Publication}
  **/
 #publication ! : InstanceType<typeof EditAnOfferActionRes.Publication>
		/**
  * 
  * @returns {EditAnOfferActionRes.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {EditAnOfferActionRes.Publication}
  **/
set publication (value: InstanceType<typeof EditAnOfferActionRes.Publication>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Publication) {
			this.#publication = value
		} else {
			this.#publication = new EditAnOfferActionRes.Publication(value)
		}
}
setPublication (value: InstanceType<typeof EditAnOfferActionRes.Publication>) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces}
  **/
 #additionalMarketplaces ! : InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces>
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces}
  **/
set additionalMarketplaces (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces) {
			this.#additionalMarketplaces = value
		} else {
			this.#additionalMarketplaces = new EditAnOfferActionRes.AdditionalMarketplaces(value)
		}
}
setAdditionalMarketplaces (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces>) {
	this.additionalMarketplaces = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.B2b}
  **/
 #b2b ! : InstanceType<typeof EditAnOfferActionRes.B2b>
		/**
  * 
  * @returns {EditAnOfferActionRes.B2b}
  **/
get b2b () { return this.#b2b }
/**
  * 
  * @type {EditAnOfferActionRes.B2b}
  **/
set b2b (value: InstanceType<typeof EditAnOfferActionRes.B2b>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.B2b) {
			this.#b2b = value
		} else {
			this.#b2b = new EditAnOfferActionRes.B2b(value)
		}
}
setB2b (value: InstanceType<typeof EditAnOfferActionRes.B2b>) {
	this.b2b = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.CompatibilityList}
  **/
 #compatibilityList ! : InstanceType<typeof EditAnOfferActionRes.CompatibilityList>
		/**
  * 
  * @returns {EditAnOfferActionRes.CompatibilityList}
  **/
get compatibilityList () { return this.#compatibilityList }
/**
  * 
  * @type {EditAnOfferActionRes.CompatibilityList}
  **/
set compatibilityList (value: InstanceType<typeof EditAnOfferActionRes.CompatibilityList>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.CompatibilityList) {
			this.#compatibilityList = value
		} else {
			this.#compatibilityList = new EditAnOfferActionRes.CompatibilityList(value)
		}
}
setCompatibilityList (value: InstanceType<typeof EditAnOfferActionRes.CompatibilityList>) {
	this.compatibilityList = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Validation}
  **/
 #validation ! : InstanceType<typeof EditAnOfferActionRes.Validation>
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation}
  **/
get validation () { return this.#validation }
/**
  * 
  * @type {EditAnOfferActionRes.Validation}
  **/
set validation (value: InstanceType<typeof EditAnOfferActionRes.Validation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Validation) {
			this.#validation = value
		} else {
			this.#validation = new EditAnOfferActionRes.Validation(value)
		}
}
setValidation (value: InstanceType<typeof EditAnOfferActionRes.Validation>) {
	this.validation = value
	return this
}
		/**
  * 
  * @type {string[]}
  **/
 #warnings : string[]  =  []
		/**
  * 
  * @returns {string[]}
  **/
get warnings () { return this.#warnings }
/**
  * 
  * @type {string[]}
  **/
set warnings (value: string[]) {
}
setWarnings (value: string[]) {
	this.warnings = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices}
  **/
 #afterSalesServices ! : InstanceType<typeof EditAnOfferActionRes.AfterSalesServices>
		/**
  * 
  * @returns {EditAnOfferActionRes.AfterSalesServices}
  **/
get afterSalesServices () { return this.#afterSalesServices }
/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices}
  **/
set afterSalesServices (value: InstanceType<typeof EditAnOfferActionRes.AfterSalesServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AfterSalesServices) {
			this.#afterSalesServices = value
		} else {
			this.#afterSalesServices = new EditAnOfferActionRes.AfterSalesServices(value)
		}
}
setAfterSalesServices (value: InstanceType<typeof EditAnOfferActionRes.AfterSalesServices>) {
	this.afterSalesServices = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Discounts}
  **/
 #discounts ! : InstanceType<typeof EditAnOfferActionRes.Discounts>
		/**
  * 
  * @returns {EditAnOfferActionRes.Discounts}
  **/
get discounts () { return this.#discounts }
/**
  * 
  * @type {EditAnOfferActionRes.Discounts}
  **/
set discounts (value: InstanceType<typeof EditAnOfferActionRes.Discounts>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Discounts) {
			this.#discounts = value
		} else {
			this.#discounts = new EditAnOfferActionRes.Discounts(value)
		}
}
setDiscounts (value: InstanceType<typeof EditAnOfferActionRes.Discounts>) {
	this.discounts = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Contact}
  **/
 #contact ! : InstanceType<typeof EditAnOfferActionRes.Contact>
		/**
  * 
  * @returns {EditAnOfferActionRes.Contact}
  **/
get contact () { return this.#contact }
/**
  * 
  * @type {EditAnOfferActionRes.Contact}
  **/
set contact (value: InstanceType<typeof EditAnOfferActionRes.Contact>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Contact) {
			this.#contact = value
		} else {
			this.#contact = new EditAnOfferActionRes.Contact(value)
		}
}
setContact (value: InstanceType<typeof EditAnOfferActionRes.Contact>) {
	this.contact = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Attachments}
  **/
 #attachments : InstanceType<typeof EditAnOfferActionRes.Attachments>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Attachments}
  **/
get attachments () { return this.#attachments }
/**
  * 
  * @type {EditAnOfferActionRes.Attachments}
  **/
set attachments (value: InstanceType<typeof EditAnOfferActionRes.Attachments>[]) {
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
setAttachments (value: InstanceType<typeof EditAnOfferActionRes.Attachments>[]) {
	this.attachments = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.FundraisingCampaign}
  **/
 #fundraisingCampaign ! : InstanceType<typeof EditAnOfferActionRes.FundraisingCampaign>
		/**
  * 
  * @returns {EditAnOfferActionRes.FundraisingCampaign}
  **/
get fundraisingCampaign () { return this.#fundraisingCampaign }
/**
  * 
  * @type {EditAnOfferActionRes.FundraisingCampaign}
  **/
set fundraisingCampaign (value: InstanceType<typeof EditAnOfferActionRes.FundraisingCampaign>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.FundraisingCampaign) {
			this.#fundraisingCampaign = value
		} else {
			this.#fundraisingCampaign = new EditAnOfferActionRes.FundraisingCampaign(value)
		}
}
setFundraisingCampaign (value: InstanceType<typeof EditAnOfferActionRes.FundraisingCampaign>) {
	this.fundraisingCampaign = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalServices}
  **/
 #additionalServices ! : InstanceType<typeof EditAnOfferActionRes.AdditionalServices>
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalServices}
  **/
get additionalServices () { return this.#additionalServices }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalServices}
  **/
set additionalServices (value: InstanceType<typeof EditAnOfferActionRes.AdditionalServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalServices) {
			this.#additionalServices = value
		} else {
			this.#additionalServices = new EditAnOfferActionRes.AdditionalServices(value)
		}
}
setAdditionalServices (value: InstanceType<typeof EditAnOfferActionRes.AdditionalServices>) {
	this.additionalServices = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.SizeTable}
  **/
 #sizeTable ! : InstanceType<typeof EditAnOfferActionRes.SizeTable>
		/**
  * 
  * @returns {EditAnOfferActionRes.SizeTable}
  **/
get sizeTable () { return this.#sizeTable }
/**
  * 
  * @type {EditAnOfferActionRes.SizeTable}
  **/
set sizeTable (value: InstanceType<typeof EditAnOfferActionRes.SizeTable>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SizeTable) {
			this.#sizeTable = value
		} else {
			this.#sizeTable = new EditAnOfferActionRes.SizeTable(value)
		}
}
setSizeTable (value: InstanceType<typeof EditAnOfferActionRes.SizeTable>) {
	this.sizeTable = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Location}
  **/
 #location ! : InstanceType<typeof EditAnOfferActionRes.Location>
		/**
  * 
  * @returns {EditAnOfferActionRes.Location}
  **/
get location () { return this.#location }
/**
  * 
  * @type {EditAnOfferActionRes.Location}
  **/
set location (value: InstanceType<typeof EditAnOfferActionRes.Location>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Location) {
			this.#location = value
		} else {
			this.#location = new EditAnOfferActionRes.Location(value)
		}
}
setLocation (value: InstanceType<typeof EditAnOfferActionRes.Location>) {
	this.location = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.External}
  **/
 #external ! : InstanceType<typeof EditAnOfferActionRes.External>
		/**
  * 
  * @returns {EditAnOfferActionRes.External}
  **/
get external () { return this.#external }
/**
  * 
  * @type {EditAnOfferActionRes.External}
  **/
set external (value: InstanceType<typeof EditAnOfferActionRes.External>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.External) {
			this.#external = value
		} else {
			this.#external = new EditAnOfferActionRes.External(value)
		}
}
setExternal (value: InstanceType<typeof EditAnOfferActionRes.External>) {
	this.external = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.TaxSettings}
  **/
 #taxSettings ! : InstanceType<typeof EditAnOfferActionRes.TaxSettings>
		/**
  * 
  * @returns {EditAnOfferActionRes.TaxSettings}
  **/
get taxSettings () { return this.#taxSettings }
/**
  * 
  * @type {EditAnOfferActionRes.TaxSettings}
  **/
set taxSettings (value: InstanceType<typeof EditAnOfferActionRes.TaxSettings>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.TaxSettings) {
			this.#taxSettings = value
		} else {
			this.#taxSettings = new EditAnOfferActionRes.TaxSettings(value)
		}
}
setTaxSettings (value: InstanceType<typeof EditAnOfferActionRes.TaxSettings>) {
	this.taxSettings = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.MessageToSellerSettings}
  **/
 #messageToSellerSettings ! : InstanceType<typeof EditAnOfferActionRes.MessageToSellerSettings>
		/**
  * 
  * @returns {EditAnOfferActionRes.MessageToSellerSettings}
  **/
get messageToSellerSettings () { return this.#messageToSellerSettings }
/**
  * 
  * @type {EditAnOfferActionRes.MessageToSellerSettings}
  **/
set messageToSellerSettings (value: InstanceType<typeof EditAnOfferActionRes.MessageToSellerSettings>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.MessageToSellerSettings) {
			this.#messageToSellerSettings = value
		} else {
			this.#messageToSellerSettings = new EditAnOfferActionRes.MessageToSellerSettings(value)
		}
}
setMessageToSellerSettings (value: InstanceType<typeof EditAnOfferActionRes.MessageToSellerSettings>) {
	this.messageToSellerSettings = value
	return this
}
		/**
  * 
  * @type {string}
  **/
 #createdAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get createdAt () { return this.#createdAt }
/**
  * 
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
  * 
  * @type {string}
  **/
 #updatedAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get updatedAt () { return this.#updatedAt }
/**
  * 
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
  * @type {EditAnOfferActionRes.Description}
  **/
 #description ! : InstanceType<typeof EditAnOfferActionRes.Description>
		/**
  * 
  * @returns {EditAnOfferActionRes.Description}
  **/
get description () { return this.#description }
/**
  * 
  * @type {EditAnOfferActionRes.Description}
  **/
set description (value: InstanceType<typeof EditAnOfferActionRes.Description>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Description) {
			this.#description = value
		} else {
			this.#description = new EditAnOfferActionRes.Description(value)
		}
}
setDescription (value: InstanceType<typeof EditAnOfferActionRes.Description>) {
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
	* Creates an instance of EditAnOfferActionRes.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.CategoryType) {
		return new EditAnOfferActionRes.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.CategoryType>) {
		return new EditAnOfferActionRes.Category(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.CategoryType>): InstanceType<typeof EditAnOfferActionRes.Category> {
		return new EditAnOfferActionRes.Category ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Category> {
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
 #quantity ! : InstanceType<typeof EditAnOfferActionRes.ProductSet.Quantity>
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Quantity}
  **/
get quantity () { return this.#quantity }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Quantity}
  **/
set quantity (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Quantity>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.Quantity) {
			this.#quantity = value
		} else {
			this.#quantity = new EditAnOfferActionRes.ProductSet.Quantity(value)
		}
}
setQuantity (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Quantity>) {
	this.quantity = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product}
  **/
 #product ! : InstanceType<typeof EditAnOfferActionRes.ProductSet.Product>
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Product}
  **/
get product () { return this.#product }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product}
  **/
set product (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Product>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.Product) {
			this.#product = value
		} else {
			this.#product = new EditAnOfferActionRes.ProductSet.Product(value)
		}
}
setProduct (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Product>) {
	this.product = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.ResponsiblePerson}
  **/
 #responsiblePerson ! : InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsiblePerson>
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.ResponsiblePerson}
  **/
get responsiblePerson () { return this.#responsiblePerson }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.ResponsiblePerson}
  **/
set responsiblePerson (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsiblePerson>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.ResponsiblePerson) {
			this.#responsiblePerson = value
		} else {
			this.#responsiblePerson = new EditAnOfferActionRes.ProductSet.ResponsiblePerson(value)
		}
}
setResponsiblePerson (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsiblePerson>) {
	this.responsiblePerson = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.ResponsibleProducer}
  **/
 #responsibleProducer ! : InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsibleProducer>
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.ResponsibleProducer}
  **/
get responsibleProducer () { return this.#responsibleProducer }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.ResponsibleProducer}
  **/
set responsibleProducer (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsibleProducer>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.ResponsibleProducer) {
			this.#responsibleProducer = value
		} else {
			this.#responsibleProducer = new EditAnOfferActionRes.ProductSet.ResponsibleProducer(value)
		}
}
setResponsibleProducer (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsibleProducer>) {
	this.responsibleProducer = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.SafetyInformation}
  **/
 #safetyInformation ! : InstanceType<typeof EditAnOfferActionRes.ProductSet.SafetyInformation>
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.SafetyInformation}
  **/
get safetyInformation () { return this.#safetyInformation }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.SafetyInformation}
  **/
set safetyInformation (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.SafetyInformation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.SafetyInformation) {
			this.#safetyInformation = value
		} else {
			this.#safetyInformation = new EditAnOfferActionRes.ProductSet.SafetyInformation(value)
		}
}
setSafetyInformation (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.SafetyInformation>) {
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
  * @type {EditAnOfferActionRes.ProductSet.Deposits}
  **/
 #deposits : InstanceType<typeof EditAnOfferActionRes.ProductSet.Deposits>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Deposits}
  **/
get deposits () { return this.#deposits }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Deposits}
  **/
set deposits (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Deposits>[]) {
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
setDeposits (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Deposits>[]) {
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.Quantity, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType.QuantityType) {
		return new EditAnOfferActionRes.ProductSet.Quantity(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Quantity, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType.QuantityType>) {
		return new EditAnOfferActionRes.ProductSet.Quantity(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType.QuantityType>): InstanceType<typeof EditAnOfferActionRes.ProductSet.Quantity> {
		return new EditAnOfferActionRes.ProductSet.Quantity ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet.Quantity> {
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
  * @type {EditAnOfferActionRes.ProductSet.Product.Publication}
  **/
 #publication ! : InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Publication>
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Product.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product.Publication}
  **/
set publication (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Publication>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.Product.Publication) {
			this.#publication = value
		} else {
			this.#publication = new EditAnOfferActionRes.ProductSet.Product.Publication(value)
		}
}
setPublication (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Publication>) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product.Parameters}
  **/
 #parameters : InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Product.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product.Parameters}
  **/
set parameters (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters>[]) {
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
setParameters (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters>[]) {
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType.ProductType.PublicationType) {
		return new EditAnOfferActionRes.ProductSet.Product.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType.ProductType.PublicationType>) {
		return new EditAnOfferActionRes.ProductSet.Product.Publication(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType.ProductType.PublicationType>): InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Publication> {
		return new EditAnOfferActionRes.ProductSet.Product.Publication ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Publication> {
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
  * @type {EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
 #rangeValue ! : InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue>
		/**
  * 
  * @returns {EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
get rangeValue () { return this.#rangeValue }
/**
  * 
  * @type {EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue}
  **/
set rangeValue (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue) {
			this.#rangeValue = value
		} else {
			this.#rangeValue = new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue(value)
		}
}
setRangeValue (value: InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue>) {
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType>) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType>): InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue> {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue> {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters.RangeValue(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType.ProductType.ParametersType) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType.ProductType.ParametersType>) {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType.ProductType.ParametersType>): InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters> {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet.Product.Parameters> {
		return new EditAnOfferActionRes.ProductSet.Product.Parameters(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType.ProductType) {
		return new EditAnOfferActionRes.ProductSet.Product(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Product, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType.ProductType>) {
		return new EditAnOfferActionRes.ProductSet.Product(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType.ProductType>): InstanceType<typeof EditAnOfferActionRes.ProductSet.Product> {
		return new EditAnOfferActionRes.ProductSet.Product ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet.Product> {
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.ResponsiblePerson, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType.ResponsiblePersonType) {
		return new EditAnOfferActionRes.ProductSet.ResponsiblePerson(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.ResponsiblePerson, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType.ResponsiblePersonType>) {
		return new EditAnOfferActionRes.ProductSet.ResponsiblePerson(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType.ResponsiblePersonType>): InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsiblePerson> {
		return new EditAnOfferActionRes.ProductSet.ResponsiblePerson ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsiblePerson> {
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.ResponsibleProducer, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType.ResponsibleProducerType) {
		return new EditAnOfferActionRes.ProductSet.ResponsibleProducer(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.ResponsibleProducer, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType.ResponsibleProducerType>) {
		return new EditAnOfferActionRes.ProductSet.ResponsibleProducer(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType.ResponsibleProducerType>): InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsibleProducer> {
		return new EditAnOfferActionRes.ProductSet.ResponsibleProducer ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet.ResponsibleProducer> {
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.SafetyInformation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType.SafetyInformationType) {
		return new EditAnOfferActionRes.ProductSet.SafetyInformation(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.SafetyInformation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType.SafetyInformationType>) {
		return new EditAnOfferActionRes.ProductSet.SafetyInformation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType.SafetyInformationType>): InstanceType<typeof EditAnOfferActionRes.ProductSet.SafetyInformation> {
		return new EditAnOfferActionRes.ProductSet.SafetyInformation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet.SafetyInformation> {
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
	* Creates an instance of EditAnOfferActionRes.ProductSet.Deposits, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType.DepositsType) {
		return new EditAnOfferActionRes.ProductSet.Deposits(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet.Deposits, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType.DepositsType>) {
		return new EditAnOfferActionRes.ProductSet.Deposits(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType.DepositsType>): InstanceType<typeof EditAnOfferActionRes.ProductSet.Deposits> {
		return new EditAnOfferActionRes.ProductSet.Deposits ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet.Deposits> {
		return new EditAnOfferActionRes.ProductSet.Deposits(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.ProductSetType) {
		return new EditAnOfferActionRes.ProductSet(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.ProductSet, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ProductSetType>) {
		return new EditAnOfferActionRes.ProductSet(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ProductSetType>): InstanceType<typeof EditAnOfferActionRes.ProductSet> {
		return new EditAnOfferActionRes.ProductSet ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.ProductSet> {
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
	* Creates an instance of EditAnOfferActionRes.Stock, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.StockType) {
		return new EditAnOfferActionRes.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.StockType>) {
		return new EditAnOfferActionRes.Stock(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.StockType>): InstanceType<typeof EditAnOfferActionRes.Stock> {
		return new EditAnOfferActionRes.Stock ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Stock> {
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
	* Creates an instance of EditAnOfferActionRes.Payments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.PaymentsType) {
		return new EditAnOfferActionRes.Payments(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Payments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.PaymentsType>) {
		return new EditAnOfferActionRes.Payments(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.PaymentsType>): InstanceType<typeof EditAnOfferActionRes.Payments> {
		return new EditAnOfferActionRes.Payments ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Payments> {
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
  * @type {EditAnOfferActionRes.SellingMode.Price}
  **/
 #price ! : InstanceType<typeof EditAnOfferActionRes.SellingMode.Price>
		/**
  * 
  * @returns {EditAnOfferActionRes.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.Price}
  **/
set price (value: InstanceType<typeof EditAnOfferActionRes.SellingMode.Price>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new EditAnOfferActionRes.SellingMode.Price(value)
		}
}
setPrice (value: InstanceType<typeof EditAnOfferActionRes.SellingMode.Price>) {
	this.price = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.MinimalPrice}
  **/
 #minimalPrice ! : InstanceType<typeof EditAnOfferActionRes.SellingMode.MinimalPrice>
		/**
  * 
  * @returns {EditAnOfferActionRes.SellingMode.MinimalPrice}
  **/
get minimalPrice () { return this.#minimalPrice }
/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.MinimalPrice}
  **/
set minimalPrice (value: InstanceType<typeof EditAnOfferActionRes.SellingMode.MinimalPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SellingMode.MinimalPrice) {
			this.#minimalPrice = value
		} else {
			this.#minimalPrice = new EditAnOfferActionRes.SellingMode.MinimalPrice(value)
		}
}
setMinimalPrice (value: InstanceType<typeof EditAnOfferActionRes.SellingMode.MinimalPrice>) {
	this.minimalPrice = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.StartingPrice}
  **/
 #startingPrice ! : InstanceType<typeof EditAnOfferActionRes.SellingMode.StartingPrice>
		/**
  * 
  * @returns {EditAnOfferActionRes.SellingMode.StartingPrice}
  **/
get startingPrice () { return this.#startingPrice }
/**
  * 
  * @type {EditAnOfferActionRes.SellingMode.StartingPrice}
  **/
set startingPrice (value: InstanceType<typeof EditAnOfferActionRes.SellingMode.StartingPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.SellingMode.StartingPrice) {
			this.#startingPrice = value
		} else {
			this.#startingPrice = new EditAnOfferActionRes.SellingMode.StartingPrice(value)
		}
}
setStartingPrice (value: InstanceType<typeof EditAnOfferActionRes.SellingMode.StartingPrice>) {
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
	* Creates an instance of EditAnOfferActionRes.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.SellingModeType.PriceType) {
		return new EditAnOfferActionRes.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.SellingModeType.PriceType>) {
		return new EditAnOfferActionRes.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.SellingModeType.PriceType>): InstanceType<typeof EditAnOfferActionRes.SellingMode.Price> {
		return new EditAnOfferActionRes.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.SellingMode.Price> {
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
	* Creates an instance of EditAnOfferActionRes.SellingMode.MinimalPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.SellingModeType.MinimalPriceType) {
		return new EditAnOfferActionRes.SellingMode.MinimalPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SellingMode.MinimalPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.SellingModeType.MinimalPriceType>) {
		return new EditAnOfferActionRes.SellingMode.MinimalPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.SellingModeType.MinimalPriceType>): InstanceType<typeof EditAnOfferActionRes.SellingMode.MinimalPrice> {
		return new EditAnOfferActionRes.SellingMode.MinimalPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.SellingMode.MinimalPrice> {
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
	* Creates an instance of EditAnOfferActionRes.SellingMode.StartingPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.SellingModeType.StartingPriceType) {
		return new EditAnOfferActionRes.SellingMode.StartingPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SellingMode.StartingPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.SellingModeType.StartingPriceType>) {
		return new EditAnOfferActionRes.SellingMode.StartingPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.SellingModeType.StartingPriceType>): InstanceType<typeof EditAnOfferActionRes.SellingMode.StartingPrice> {
		return new EditAnOfferActionRes.SellingMode.StartingPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.SellingMode.StartingPrice> {
		return new EditAnOfferActionRes.SellingMode.StartingPrice(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.SellingModeType) {
		return new EditAnOfferActionRes.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.SellingModeType>) {
		return new EditAnOfferActionRes.SellingMode(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.SellingModeType>): InstanceType<typeof EditAnOfferActionRes.SellingMode> {
		return new EditAnOfferActionRes.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.SellingMode> {
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
  * @type {EditAnOfferActionRes.Delivery.ShippingRates}
  **/
 #shippingRates ! : InstanceType<typeof EditAnOfferActionRes.Delivery.ShippingRates>
		/**
  * 
  * @returns {EditAnOfferActionRes.Delivery.ShippingRates}
  **/
get shippingRates () { return this.#shippingRates }
/**
  * 
  * @type {EditAnOfferActionRes.Delivery.ShippingRates}
  **/
set shippingRates (value: InstanceType<typeof EditAnOfferActionRes.Delivery.ShippingRates>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Delivery.ShippingRates) {
			this.#shippingRates = value
		} else {
			this.#shippingRates = new EditAnOfferActionRes.Delivery.ShippingRates(value)
		}
}
setShippingRates (value: InstanceType<typeof EditAnOfferActionRes.Delivery.ShippingRates>) {
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
	* Creates an instance of EditAnOfferActionRes.Delivery.ShippingRates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.DeliveryType.ShippingRatesType) {
		return new EditAnOfferActionRes.Delivery.ShippingRates(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Delivery.ShippingRates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.DeliveryType.ShippingRatesType>) {
		return new EditAnOfferActionRes.Delivery.ShippingRates(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.DeliveryType.ShippingRatesType>): InstanceType<typeof EditAnOfferActionRes.Delivery.ShippingRates> {
		return new EditAnOfferActionRes.Delivery.ShippingRates ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Delivery.ShippingRates> {
		return new EditAnOfferActionRes.Delivery.ShippingRates(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.DeliveryType) {
		return new EditAnOfferActionRes.Delivery(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Delivery, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.DeliveryType>) {
		return new EditAnOfferActionRes.Delivery(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.DeliveryType>): InstanceType<typeof EditAnOfferActionRes.Delivery> {
		return new EditAnOfferActionRes.Delivery ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Delivery> {
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
  * @type {EditAnOfferActionRes.Publication.Marketplaces}
  **/
 #marketplaces ! : InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces>
		/**
  * 
  * @returns {EditAnOfferActionRes.Publication.Marketplaces}
  **/
get marketplaces () { return this.#marketplaces }
/**
  * 
  * @type {EditAnOfferActionRes.Publication.Marketplaces}
  **/
set marketplaces (value: InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Publication.Marketplaces) {
			this.#marketplaces = value
		} else {
			this.#marketplaces = new EditAnOfferActionRes.Publication.Marketplaces(value)
		}
}
setMarketplaces (value: InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces>) {
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
 #base ! : InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Base>
		/**
  * 
  * @returns {EditAnOfferActionRes.Publication.Marketplaces.Base}
  **/
get base () { return this.#base }
/**
  * 
  * @type {EditAnOfferActionRes.Publication.Marketplaces.Base}
  **/
set base (value: InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Base>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Publication.Marketplaces.Base) {
			this.#base = value
		} else {
			this.#base = new EditAnOfferActionRes.Publication.Marketplaces.Base(value)
		}
}
setBase (value: InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Base>) {
	this.base = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Publication.Marketplaces.Additional}
  **/
 #additional : InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Additional>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Publication.Marketplaces.Additional}
  **/
get additional () { return this.#additional }
/**
  * 
  * @type {EditAnOfferActionRes.Publication.Marketplaces.Additional}
  **/
set additional (value: InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Additional>[]) {
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
setAdditional (value: InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Additional>[]) {
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
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces.Base, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.PublicationType.MarketplacesType.BaseType) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Base(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces.Base, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.PublicationType.MarketplacesType.BaseType>) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Base(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.PublicationType.MarketplacesType.BaseType>): InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Base> {
		return new EditAnOfferActionRes.Publication.Marketplaces.Base ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Base> {
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
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces.Additional, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.PublicationType.MarketplacesType.AdditionalType) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Additional(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces.Additional, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.PublicationType.MarketplacesType.AdditionalType>) {
		return new EditAnOfferActionRes.Publication.Marketplaces.Additional(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.PublicationType.MarketplacesType.AdditionalType>): InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Additional> {
		return new EditAnOfferActionRes.Publication.Marketplaces.Additional ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces.Additional> {
		return new EditAnOfferActionRes.Publication.Marketplaces.Additional(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.PublicationType.MarketplacesType) {
		return new EditAnOfferActionRes.Publication.Marketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication.Marketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.PublicationType.MarketplacesType>) {
		return new EditAnOfferActionRes.Publication.Marketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.PublicationType.MarketplacesType>): InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces> {
		return new EditAnOfferActionRes.Publication.Marketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Publication.Marketplaces> {
		return new EditAnOfferActionRes.Publication.Marketplaces(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.PublicationType) {
		return new EditAnOfferActionRes.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.PublicationType>) {
		return new EditAnOfferActionRes.Publication(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.PublicationType>): InstanceType<typeof EditAnOfferActionRes.Publication> {
		return new EditAnOfferActionRes.Publication ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Publication> {
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
 #sellingMode ! : InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode>
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode}
  **/
set sellingMode (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode(value)
		}
}
setSellingMode (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode>) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication}
  **/
 #publication ! : InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication>
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication}
  **/
set publication (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces.Publication) {
			this.#publication = value
		} else {
			this.#publication = new EditAnOfferActionRes.AdditionalMarketplaces.Publication(value)
		}
}
setPublication (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication>) {
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
 #price ! : InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price>
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price}
  **/
set price (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price(value)
		}
}
setPrice (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price>) {
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
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType.PriceType) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType.PriceType>) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType.PriceType>): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price> {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price> {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode.Price(this.toJSON());
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
			if (d.price !== undefined) { this.price = d.price }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<SellingMode>;
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
	static from(possibleDtoObject: EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType>) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType>): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode> {
		return new EditAnOfferActionRes.AdditionalMarketplaces.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.SellingMode> {
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
 #state : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get state () { return this.#state }
/**
  * 
  * @type {string}
  **/
set state (value: string) {
		this.#state = String(value);
}
setState (value: string) {
	this.state = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons}
  **/
 #refusalReasons : InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons}
  **/
get refusalReasons () { return this.#refusalReasons }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons}
  **/
set refusalReasons (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons>[]) {
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
setRefusalReasons (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons>[]) {
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
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters}
  **/
 #parameters ! : InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters>
		/**
  * 
  * @returns {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters}
  **/
get parameters () { return this.#parameters }
/**
  * 
  * @type {EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters}
  **/
set parameters (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters) {
			this.#parameters = value
		} else {
			this.#parameters = new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters(value)
		}
}
setParameters (value: InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters>) {
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
 #maxAllowedPriceDecreasePercent : string[]  =  []
		/**
  * 
  * @returns {string[]}
  **/
get maxAllowedPriceDecreasePercent () { return this.#maxAllowedPriceDecreasePercent }
/**
  * 
  * @type {string[]}
  **/
set maxAllowedPriceDecreasePercent (value: string[]) {
}
setMaxAllowedPriceDecreasePercent (value: string[]) {
	this.maxAllowedPriceDecreasePercent = value
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
		const d = data as Partial<Parameters>;
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
	static from(possibleDtoObject: EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType.ParametersType) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType.ParametersType>) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType.ParametersType>): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters> {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters> {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons.Parameters(this.toJSON());
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
		const d = data as Partial<RefusalReasons>;
			if (d.code !== undefined) { this.code = d.code }
			if (d.userMessage !== undefined) { this.userMessage = d.userMessage }
			if (d.parameters !== undefined) { this.parameters = d.parameters }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<RefusalReasons>;
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
	static from(possibleDtoObject: EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType>) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType>): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons> {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons> {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication.RefusalReasons(this.toJSON());
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
		const d = data as Partial<Publication>;
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
	static from(possibleDtoObject: EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType>) {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType>): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication> {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces.Publication> {
		return new EditAnOfferActionRes.AdditionalMarketplaces.Publication(this.toJSON());
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
		const d = data as Partial<AdditionalMarketplaces>;
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
			if (d.publication !== undefined) { this.publication = d.publication }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<AdditionalMarketplaces>;
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
	static from(possibleDtoObject: EditAnOfferActionResType.AdditionalMarketplacesType) {
		return new EditAnOfferActionRes.AdditionalMarketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalMarketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType>) {
		return new EditAnOfferActionRes.AdditionalMarketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AdditionalMarketplacesType>): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces> {
		return new EditAnOfferActionRes.AdditionalMarketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AdditionalMarketplaces> {
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
	* Creates an instance of EditAnOfferActionRes.B2b, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.B2bType) {
		return new EditAnOfferActionRes.B2b(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.B2b, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.B2bType>) {
		return new EditAnOfferActionRes.B2b(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.B2bType>): InstanceType<typeof EditAnOfferActionRes.B2b> {
		return new EditAnOfferActionRes.B2b ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.B2b> {
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
	* Creates an instance of EditAnOfferActionRes.CompatibilityList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.CompatibilityListType) {
		return new EditAnOfferActionRes.CompatibilityList(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.CompatibilityList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.CompatibilityListType>) {
		return new EditAnOfferActionRes.CompatibilityList(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.CompatibilityListType>): InstanceType<typeof EditAnOfferActionRes.CompatibilityList> {
		return new EditAnOfferActionRes.CompatibilityList ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.CompatibilityList> {
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
  * @type {EditAnOfferActionRes.Validation.Errors}
  **/
 #errors : InstanceType<typeof EditAnOfferActionRes.Validation.Errors>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation.Errors}
  **/
get errors () { return this.#errors }
/**
  * 
  * @type {EditAnOfferActionRes.Validation.Errors}
  **/
set errors (value: InstanceType<typeof EditAnOfferActionRes.Validation.Errors>[]) {
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
setErrors (value: InstanceType<typeof EditAnOfferActionRes.Validation.Errors>[]) {
	this.errors = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.Validation.Warnings}
  **/
 #warnings : InstanceType<typeof EditAnOfferActionRes.Validation.Warnings>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation.Warnings}
  **/
get warnings () { return this.#warnings }
/**
  * 
  * @type {EditAnOfferActionRes.Validation.Warnings}
  **/
set warnings (value: InstanceType<typeof EditAnOfferActionRes.Validation.Warnings>[]) {
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
setWarnings (value: InstanceType<typeof EditAnOfferActionRes.Validation.Warnings>[]) {
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
  * @type {EditAnOfferActionRes.Validation.Errors.Metadata}
  **/
 #metadata ! : InstanceType<typeof EditAnOfferActionRes.Validation.Errors.Metadata>
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation.Errors.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {EditAnOfferActionRes.Validation.Errors.Metadata}
  **/
set metadata (value: InstanceType<typeof EditAnOfferActionRes.Validation.Errors.Metadata>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Validation.Errors.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new EditAnOfferActionRes.Validation.Errors.Metadata(value)
		}
}
setMetadata (value: InstanceType<typeof EditAnOfferActionRes.Validation.Errors.Metadata>) {
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
	* Creates an instance of EditAnOfferActionRes.Validation.Errors.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ValidationType.ErrorsType.MetadataType) {
		return new EditAnOfferActionRes.Validation.Errors.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Errors.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ValidationType.ErrorsType.MetadataType>) {
		return new EditAnOfferActionRes.Validation.Errors.Metadata(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ValidationType.ErrorsType.MetadataType>): InstanceType<typeof EditAnOfferActionRes.Validation.Errors.Metadata> {
		return new EditAnOfferActionRes.Validation.Errors.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Validation.Errors.Metadata> {
		return new EditAnOfferActionRes.Validation.Errors.Metadata(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.ValidationType.ErrorsType) {
		return new EditAnOfferActionRes.Validation.Errors(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Errors, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ValidationType.ErrorsType>) {
		return new EditAnOfferActionRes.Validation.Errors(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ValidationType.ErrorsType>): InstanceType<typeof EditAnOfferActionRes.Validation.Errors> {
		return new EditAnOfferActionRes.Validation.Errors ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Validation.Errors> {
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
  * @type {EditAnOfferActionRes.Validation.Warnings.Metadata}
  **/
 #metadata ! : InstanceType<typeof EditAnOfferActionRes.Validation.Warnings.Metadata>
		/**
  * 
  * @returns {EditAnOfferActionRes.Validation.Warnings.Metadata}
  **/
get metadata () { return this.#metadata }
/**
  * 
  * @type {EditAnOfferActionRes.Validation.Warnings.Metadata}
  **/
set metadata (value: InstanceType<typeof EditAnOfferActionRes.Validation.Warnings.Metadata>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Validation.Warnings.Metadata) {
			this.#metadata = value
		} else {
			this.#metadata = new EditAnOfferActionRes.Validation.Warnings.Metadata(value)
		}
}
setMetadata (value: InstanceType<typeof EditAnOfferActionRes.Validation.Warnings.Metadata>) {
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
	* Creates an instance of EditAnOfferActionRes.Validation.Warnings.Metadata, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ValidationType.WarningsType.MetadataType) {
		return new EditAnOfferActionRes.Validation.Warnings.Metadata(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Warnings.Metadata, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ValidationType.WarningsType.MetadataType>) {
		return new EditAnOfferActionRes.Validation.Warnings.Metadata(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ValidationType.WarningsType.MetadataType>): InstanceType<typeof EditAnOfferActionRes.Validation.Warnings.Metadata> {
		return new EditAnOfferActionRes.Validation.Warnings.Metadata ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Validation.Warnings.Metadata> {
		return new EditAnOfferActionRes.Validation.Warnings.Metadata(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.ValidationType.WarningsType) {
		return new EditAnOfferActionRes.Validation.Warnings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation.Warnings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ValidationType.WarningsType>) {
		return new EditAnOfferActionRes.Validation.Warnings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ValidationType.WarningsType>): InstanceType<typeof EditAnOfferActionRes.Validation.Warnings> {
		return new EditAnOfferActionRes.Validation.Warnings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Validation.Warnings> {
		return new EditAnOfferActionRes.Validation.Warnings(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.ValidationType) {
		return new EditAnOfferActionRes.Validation(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Validation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ValidationType>) {
		return new EditAnOfferActionRes.Validation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ValidationType>): InstanceType<typeof EditAnOfferActionRes.Validation> {
		return new EditAnOfferActionRes.Validation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Validation> {
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
 #impliedWarranty ! : InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty>
		/**
  * 
  * @returns {EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
get impliedWarranty () { return this.#impliedWarranty }
/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty}
  **/
set impliedWarranty (value: InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty) {
			this.#impliedWarranty = value
		} else {
			this.#impliedWarranty = new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty(value)
		}
}
setImpliedWarranty (value: InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty>) {
	this.impliedWarranty = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
 #returnPolicy ! : InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ReturnPolicy>
		/**
  * 
  * @returns {EditAnOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
get returnPolicy () { return this.#returnPolicy }
/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.ReturnPolicy}
  **/
set returnPolicy (value: InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ReturnPolicy>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AfterSalesServices.ReturnPolicy) {
			this.#returnPolicy = value
		} else {
			this.#returnPolicy = new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy(value)
		}
}
setReturnPolicy (value: InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ReturnPolicy>) {
	this.returnPolicy = value
	return this
}
		/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.Warranty}
  **/
 #warranty ! : InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.Warranty>
		/**
  * 
  * @returns {EditAnOfferActionRes.AfterSalesServices.Warranty}
  **/
get warranty () { return this.#warranty }
/**
  * 
  * @type {EditAnOfferActionRes.AfterSalesServices.Warranty}
  **/
set warranty (value: InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.Warranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.AfterSalesServices.Warranty) {
			this.#warranty = value
		} else {
			this.#warranty = new EditAnOfferActionRes.AfterSalesServices.Warranty(value)
		}
}
setWarranty (value: InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.Warranty>) {
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
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType) {
		return new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType>) {
		return new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType>): InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty> {
		return new EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ImpliedWarranty> {
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
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.ReturnPolicy, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.AfterSalesServicesType.ReturnPolicyType) {
		return new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.ReturnPolicy, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AfterSalesServicesType.ReturnPolicyType>) {
		return new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AfterSalesServicesType.ReturnPolicyType>): InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ReturnPolicy> {
		return new EditAnOfferActionRes.AfterSalesServices.ReturnPolicy ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.ReturnPolicy> {
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
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.Warranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.AfterSalesServicesType.WarrantyType) {
		return new EditAnOfferActionRes.AfterSalesServices.Warranty(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices.Warranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AfterSalesServicesType.WarrantyType>) {
		return new EditAnOfferActionRes.AfterSalesServices.Warranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AfterSalesServicesType.WarrantyType>): InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.Warranty> {
		return new EditAnOfferActionRes.AfterSalesServices.Warranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AfterSalesServices.Warranty> {
		return new EditAnOfferActionRes.AfterSalesServices.Warranty(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.AfterSalesServicesType) {
		return new EditAnOfferActionRes.AfterSalesServices(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AfterSalesServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AfterSalesServicesType>) {
		return new EditAnOfferActionRes.AfterSalesServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AfterSalesServicesType>): InstanceType<typeof EditAnOfferActionRes.AfterSalesServices> {
		return new EditAnOfferActionRes.AfterSalesServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AfterSalesServices> {
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
 #wholesalePriceList ! : InstanceType<typeof EditAnOfferActionRes.Discounts.WholesalePriceList>
		/**
  * 
  * @returns {EditAnOfferActionRes.Discounts.WholesalePriceList}
  **/
get wholesalePriceList () { return this.#wholesalePriceList }
/**
  * 
  * @type {EditAnOfferActionRes.Discounts.WholesalePriceList}
  **/
set wholesalePriceList (value: InstanceType<typeof EditAnOfferActionRes.Discounts.WholesalePriceList>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof EditAnOfferActionRes.Discounts.WholesalePriceList) {
			this.#wholesalePriceList = value
		} else {
			this.#wholesalePriceList = new EditAnOfferActionRes.Discounts.WholesalePriceList(value)
		}
}
setWholesalePriceList (value: InstanceType<typeof EditAnOfferActionRes.Discounts.WholesalePriceList>) {
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
	* Creates an instance of EditAnOfferActionRes.Discounts.WholesalePriceList, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.DiscountsType.WholesalePriceListType) {
		return new EditAnOfferActionRes.Discounts.WholesalePriceList(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Discounts.WholesalePriceList, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.DiscountsType.WholesalePriceListType>) {
		return new EditAnOfferActionRes.Discounts.WholesalePriceList(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.DiscountsType.WholesalePriceListType>): InstanceType<typeof EditAnOfferActionRes.Discounts.WholesalePriceList> {
		return new EditAnOfferActionRes.Discounts.WholesalePriceList ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Discounts.WholesalePriceList> {
		return new EditAnOfferActionRes.Discounts.WholesalePriceList(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.DiscountsType) {
		return new EditAnOfferActionRes.Discounts(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Discounts, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.DiscountsType>) {
		return new EditAnOfferActionRes.Discounts(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.DiscountsType>): InstanceType<typeof EditAnOfferActionRes.Discounts> {
		return new EditAnOfferActionRes.Discounts ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Discounts> {
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
	* Creates an instance of EditAnOfferActionRes.Contact, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ContactType) {
		return new EditAnOfferActionRes.Contact(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Contact, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ContactType>) {
		return new EditAnOfferActionRes.Contact(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ContactType>): InstanceType<typeof EditAnOfferActionRes.Contact> {
		return new EditAnOfferActionRes.Contact ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Contact> {
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
	* Creates an instance of EditAnOfferActionRes.Attachments, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.AttachmentsType) {
		return new EditAnOfferActionRes.Attachments(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Attachments, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AttachmentsType>) {
		return new EditAnOfferActionRes.Attachments(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AttachmentsType>): InstanceType<typeof EditAnOfferActionRes.Attachments> {
		return new EditAnOfferActionRes.Attachments ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Attachments> {
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
	* Creates an instance of EditAnOfferActionRes.FundraisingCampaign, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.FundraisingCampaignType) {
		return new EditAnOfferActionRes.FundraisingCampaign(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.FundraisingCampaign, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.FundraisingCampaignType>) {
		return new EditAnOfferActionRes.FundraisingCampaign(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.FundraisingCampaignType>): InstanceType<typeof EditAnOfferActionRes.FundraisingCampaign> {
		return new EditAnOfferActionRes.FundraisingCampaign ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.FundraisingCampaign> {
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
	* Creates an instance of EditAnOfferActionRes.AdditionalServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.AdditionalServicesType) {
		return new EditAnOfferActionRes.AdditionalServices(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.AdditionalServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.AdditionalServicesType>) {
		return new EditAnOfferActionRes.AdditionalServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.AdditionalServicesType>): InstanceType<typeof EditAnOfferActionRes.AdditionalServices> {
		return new EditAnOfferActionRes.AdditionalServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.AdditionalServices> {
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
	* Creates an instance of EditAnOfferActionRes.SizeTable, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.SizeTableType) {
		return new EditAnOfferActionRes.SizeTable(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.SizeTable, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.SizeTableType>) {
		return new EditAnOfferActionRes.SizeTable(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.SizeTableType>): InstanceType<typeof EditAnOfferActionRes.SizeTable> {
		return new EditAnOfferActionRes.SizeTable ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.SizeTable> {
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
	* Creates an instance of EditAnOfferActionRes.Location, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.LocationType) {
		return new EditAnOfferActionRes.Location(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Location, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.LocationType>) {
		return new EditAnOfferActionRes.Location(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.LocationType>): InstanceType<typeof EditAnOfferActionRes.Location> {
		return new EditAnOfferActionRes.Location ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Location> {
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
	* Creates an instance of EditAnOfferActionRes.External, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.ExternalType) {
		return new EditAnOfferActionRes.External(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.External, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.ExternalType>) {
		return new EditAnOfferActionRes.External(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.ExternalType>): InstanceType<typeof EditAnOfferActionRes.External> {
		return new EditAnOfferActionRes.External ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.External> {
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
  * @type {EditAnOfferActionRes.TaxSettings.Rates}
  **/
 #rates : InstanceType<typeof EditAnOfferActionRes.TaxSettings.Rates>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.TaxSettings.Rates}
  **/
get rates () { return this.#rates }
/**
  * 
  * @type {EditAnOfferActionRes.TaxSettings.Rates}
  **/
set rates (value: InstanceType<typeof EditAnOfferActionRes.TaxSettings.Rates>[]) {
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
setRates (value: InstanceType<typeof EditAnOfferActionRes.TaxSettings.Rates>[]) {
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
	* Creates an instance of EditAnOfferActionRes.TaxSettings.Rates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.TaxSettingsType.RatesType) {
		return new EditAnOfferActionRes.TaxSettings.Rates(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.TaxSettings.Rates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.TaxSettingsType.RatesType>) {
		return new EditAnOfferActionRes.TaxSettings.Rates(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.TaxSettingsType.RatesType>): InstanceType<typeof EditAnOfferActionRes.TaxSettings.Rates> {
		return new EditAnOfferActionRes.TaxSettings.Rates ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.TaxSettings.Rates> {
		return new EditAnOfferActionRes.TaxSettings.Rates(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.TaxSettingsType) {
		return new EditAnOfferActionRes.TaxSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.TaxSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.TaxSettingsType>) {
		return new EditAnOfferActionRes.TaxSettings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.TaxSettingsType>): InstanceType<typeof EditAnOfferActionRes.TaxSettings> {
		return new EditAnOfferActionRes.TaxSettings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.TaxSettings> {
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
	* Creates an instance of EditAnOfferActionRes.MessageToSellerSettings, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.MessageToSellerSettingsType) {
		return new EditAnOfferActionRes.MessageToSellerSettings(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.MessageToSellerSettings, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.MessageToSellerSettingsType>) {
		return new EditAnOfferActionRes.MessageToSellerSettings(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.MessageToSellerSettingsType>): InstanceType<typeof EditAnOfferActionRes.MessageToSellerSettings> {
		return new EditAnOfferActionRes.MessageToSellerSettings ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.MessageToSellerSettings> {
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
 #sections : InstanceType<typeof EditAnOfferActionRes.Description.Sections>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Description.Sections}
  **/
get sections () { return this.#sections }
/**
  * 
  * @type {EditAnOfferActionRes.Description.Sections}
  **/
set sections (value: InstanceType<typeof EditAnOfferActionRes.Description.Sections>[]) {
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
setSections (value: InstanceType<typeof EditAnOfferActionRes.Description.Sections>[]) {
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
 #items : InstanceType<typeof EditAnOfferActionRes.Description.Sections.Items>[]  =  []
		/**
  * 
  * @returns {EditAnOfferActionRes.Description.Sections.Items}
  **/
get items () { return this.#items }
/**
  * 
  * @type {EditAnOfferActionRes.Description.Sections.Items}
  **/
set items (value: InstanceType<typeof EditAnOfferActionRes.Description.Sections.Items>[]) {
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
setItems (value: InstanceType<typeof EditAnOfferActionRes.Description.Sections.Items>[]) {
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
	* Creates an instance of EditAnOfferActionRes.Description.Sections.Items, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: EditAnOfferActionResType.DescriptionType.SectionsType.ItemsType) {
		return new EditAnOfferActionRes.Description.Sections.Items(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Description.Sections.Items, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.DescriptionType.SectionsType.ItemsType>) {
		return new EditAnOfferActionRes.Description.Sections.Items(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.DescriptionType.SectionsType.ItemsType>): InstanceType<typeof EditAnOfferActionRes.Description.Sections.Items> {
		return new EditAnOfferActionRes.Description.Sections.Items ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Description.Sections.Items> {
		return new EditAnOfferActionRes.Description.Sections.Items(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.DescriptionType.SectionsType) {
		return new EditAnOfferActionRes.Description.Sections(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Description.Sections, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.DescriptionType.SectionsType>) {
		return new EditAnOfferActionRes.Description.Sections(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.DescriptionType.SectionsType>): InstanceType<typeof EditAnOfferActionRes.Description.Sections> {
		return new EditAnOfferActionRes.Description.Sections ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Description.Sections> {
		return new EditAnOfferActionRes.Description.Sections(this.toJSON());
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
	static from(possibleDtoObject: EditAnOfferActionResType.DescriptionType) {
		return new EditAnOfferActionRes.Description(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes.Description, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType.DescriptionType>) {
		return new EditAnOfferActionRes.Description(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType.DescriptionType>): InstanceType<typeof EditAnOfferActionRes.Description> {
		return new EditAnOfferActionRes.Description ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes.Description> {
		return new EditAnOfferActionRes.Description(this.toJSON());
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
		const d = data as Partial<EditAnOfferActionRes>;
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
		const d = data as Partial<EditAnOfferActionRes>;
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
	static from(possibleDtoObject: EditAnOfferActionResType) {
		return new EditAnOfferActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of EditAnOfferActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<EditAnOfferActionResType>) {
		return new EditAnOfferActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<EditAnOfferActionResType>): InstanceType<typeof EditAnOfferActionRes> {
		return new EditAnOfferActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof EditAnOfferActionRes> {
		return new EditAnOfferActionRes(this.toJSON());
	}
}
export abstract class EditAnOfferActionResFactory {
	abstract create(data: unknown): EditAnOfferActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for editAnOfferActionRes
  **/
	export type EditAnOfferActionResType =  {
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
  * @type {string}
  **/
 language : string;
			/**
  * 
  * @type {EditAnOfferActionResType.CategoryType}
  **/
 category : EditAnOfferActionResType.CategoryType;
			/**
  * 
  * @type {EditAnOfferActionResType.ProductSetType[]}
  **/
 productSet : EditAnOfferActionResType.ProductSetType[];
			/**
  * 
  * @type {EditAnOfferActionResType.StockType}
  **/
 stock : EditAnOfferActionResType.StockType;
			/**
  * 
  * @type {EditAnOfferActionResType.PaymentsType}
  **/
 payments : EditAnOfferActionResType.PaymentsType;
			/**
  * 
  * @type {EditAnOfferActionResType.SellingModeType}
  **/
 sellingMode : EditAnOfferActionResType.SellingModeType;
			/**
  * 
  * @type {EditAnOfferActionResType.DeliveryType}
  **/
 delivery : EditAnOfferActionResType.DeliveryType;
			/**
  * 
  * @type {EditAnOfferActionResType.PublicationType}
  **/
 publication : EditAnOfferActionResType.PublicationType;
			/**
  * 
  * @type {EditAnOfferActionResType.AdditionalMarketplacesType}
  **/
 additionalMarketplaces : EditAnOfferActionResType.AdditionalMarketplacesType;
			/**
  * 
  * @type {EditAnOfferActionResType.B2bType}
  **/
 b2b : EditAnOfferActionResType.B2bType;
			/**
  * 
  * @type {EditAnOfferActionResType.CompatibilityListType}
  **/
 compatibilityList : EditAnOfferActionResType.CompatibilityListType;
			/**
  * 
  * @type {EditAnOfferActionResType.ValidationType}
  **/
 validation : EditAnOfferActionResType.ValidationType;
			/**
  * 
  * @type {string[]}
  **/
 warnings : string[];
			/**
  * 
  * @type {EditAnOfferActionResType.AfterSalesServicesType}
  **/
 afterSalesServices : EditAnOfferActionResType.AfterSalesServicesType;
			/**
  * 
  * @type {EditAnOfferActionResType.DiscountsType}
  **/
 discounts : EditAnOfferActionResType.DiscountsType;
			/**
  * 
  * @type {EditAnOfferActionResType.ContactType}
  **/
 contact : EditAnOfferActionResType.ContactType;
			/**
  * 
  * @type {EditAnOfferActionResType.AttachmentsType[]}
  **/
 attachments : EditAnOfferActionResType.AttachmentsType[];
			/**
  * 
  * @type {EditAnOfferActionResType.FundraisingCampaignType}
  **/
 fundraisingCampaign : EditAnOfferActionResType.FundraisingCampaignType;
			/**
  * 
  * @type {EditAnOfferActionResType.AdditionalServicesType}
  **/
 additionalServices : EditAnOfferActionResType.AdditionalServicesType;
			/**
  * 
  * @type {EditAnOfferActionResType.SizeTableType}
  **/
 sizeTable : EditAnOfferActionResType.SizeTableType;
			/**
  * 
  * @type {EditAnOfferActionResType.LocationType}
  **/
 location : EditAnOfferActionResType.LocationType;
			/**
  * 
  * @type {EditAnOfferActionResType.ExternalType}
  **/
 external : EditAnOfferActionResType.ExternalType;
			/**
  * 
  * @type {EditAnOfferActionResType.TaxSettingsType}
  **/
 taxSettings : EditAnOfferActionResType.TaxSettingsType;
			/**
  * 
  * @type {EditAnOfferActionResType.MessageToSellerSettingsType}
  **/
 messageToSellerSettings : EditAnOfferActionResType.MessageToSellerSettingsType;
			/**
  * 
  * @type {string}
  **/
 createdAt : string;
			/**
  * 
  * @type {string}
  **/
 updatedAt : string;
			/**
  * 
  * @type {string[]}
  **/
 images : string[];
			/**
  * 
  * @type {EditAnOfferActionResType.DescriptionType}
  **/
 description : EditAnOfferActionResType.DescriptionType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace EditAnOfferActionResType {
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
  * @type {EditAnOfferActionResType.ProductSetType.QuantityType}
  **/
 quantity : EditAnOfferActionResType.ProductSetType.QuantityType;
			/**
  * 
  * @type {EditAnOfferActionResType.ProductSetType.ProductType}
  **/
 product : EditAnOfferActionResType.ProductSetType.ProductType;
			/**
  * 
  * @type {EditAnOfferActionResType.ProductSetType.ResponsiblePersonType}
  **/
 responsiblePerson : EditAnOfferActionResType.ProductSetType.ResponsiblePersonType;
			/**
  * 
  * @type {EditAnOfferActionResType.ProductSetType.ResponsibleProducerType}
  **/
 responsibleProducer : EditAnOfferActionResType.ProductSetType.ResponsibleProducerType;
			/**
  * 
  * @type {EditAnOfferActionResType.ProductSetType.SafetyInformationType}
  **/
 safetyInformation : EditAnOfferActionResType.ProductSetType.SafetyInformationType;
			/**
  * 
  * @type {boolean}
  **/
 marketedBeforeGPSRObligation : boolean;
			/**
  * 
  * @type {EditAnOfferActionResType.ProductSetType.DepositsType[]}
  **/
 deposits : EditAnOfferActionResType.ProductSetType.DepositsType[];
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
  * @type {EditAnOfferActionResType.ProductSetType.ProductType.PublicationType}
  **/
 publication : EditAnOfferActionResType.ProductSetType.ProductType.PublicationType;
			/**
  * 
  * @type {EditAnOfferActionResType.ProductSetType.ProductType.ParametersType[]}
  **/
 parameters : EditAnOfferActionResType.ProductSetType.ProductType.ParametersType[];
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
  * @type {EditAnOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType}
  **/
 rangeValue : EditAnOfferActionResType.ProductSetType.ProductType.ParametersType.RangeValueType;
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
  * @type {EditAnOfferActionResType.SellingModeType.PriceType}
  **/
 price : EditAnOfferActionResType.SellingModeType.PriceType;
			/**
  * 
  * @type {EditAnOfferActionResType.SellingModeType.MinimalPriceType}
  **/
 minimalPrice : EditAnOfferActionResType.SellingModeType.MinimalPriceType;
			/**
  * 
  * @type {EditAnOfferActionResType.SellingModeType.StartingPriceType}
  **/
 startingPrice : EditAnOfferActionResType.SellingModeType.StartingPriceType;
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
  * @type {EditAnOfferActionResType.DeliveryType.ShippingRatesType}
  **/
 shippingRates : EditAnOfferActionResType.DeliveryType.ShippingRatesType;
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
  * @type {EditAnOfferActionResType.PublicationType.MarketplacesType}
  **/
 marketplaces : EditAnOfferActionResType.PublicationType.MarketplacesType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PublicationType {
	/**
  * The base type definition for marketplacesType
  **/
	export type MarketplacesType =  {
			/**
  * 
  * @type {EditAnOfferActionResType.PublicationType.MarketplacesType.BaseType}
  **/
 base : EditAnOfferActionResType.PublicationType.MarketplacesType.BaseType;
			/**
  * 
  * @type {EditAnOfferActionResType.PublicationType.MarketplacesType.AdditionalType[]}
  **/
 additional : EditAnOfferActionResType.PublicationType.MarketplacesType.AdditionalType[];
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
  * The base type definition for additionalMarketplacesType
  **/
	export type AdditionalMarketplacesType =  {
			/**
  * 
  * @type {EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType}
  **/
 sellingMode : EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType;
			/**
  * 
  * @type {EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType}
  **/
 publication : EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdditionalMarketplacesType {
	/**
  * The base type definition for sellingModeType
  **/
	export type SellingModeType =  {
			/**
  * 
  * @type {EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType.PriceType}
  **/
 price : EditAnOfferActionResType.AdditionalMarketplacesType.SellingModeType.PriceType;
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
}
	/**
  * The base type definition for publicationType
  **/
	export type PublicationType =  {
			/**
  * 
  * @type {string}
  **/
 state : string;
			/**
  * 
  * @type {EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType[]}
  **/
 refusalReasons : EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PublicationType {
	/**
  * The base type definition for refusalReasonsType
  **/
	export type RefusalReasonsType =  {
			/**
  * 
  * @type {string}
  **/
 code : string;
			/**
  * 
  * @type {string}
  **/
 userMessage : string;
			/**
  * 
  * @type {EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType.ParametersType}
  **/
 parameters : EditAnOfferActionResType.AdditionalMarketplacesType.PublicationType.RefusalReasonsType.ParametersType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace RefusalReasonsType {
	/**
  * The base type definition for parametersType
  **/
	export type ParametersType =  {
			/**
  * 
  * @type {string[]}
  **/
 maxAllowedPriceDecreasePercent : string[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ParametersType {
}
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
  * @type {EditAnOfferActionResType.ValidationType.ErrorsType[]}
  **/
 errors : EditAnOfferActionResType.ValidationType.ErrorsType[];
			/**
  * 
  * @type {EditAnOfferActionResType.ValidationType.WarningsType[]}
  **/
 warnings : EditAnOfferActionResType.ValidationType.WarningsType[];
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
  * @type {EditAnOfferActionResType.ValidationType.ErrorsType.MetadataType}
  **/
 metadata : EditAnOfferActionResType.ValidationType.ErrorsType.MetadataType;
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
  * @type {EditAnOfferActionResType.ValidationType.WarningsType.MetadataType}
  **/
 metadata : EditAnOfferActionResType.ValidationType.WarningsType.MetadataType;
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
  * The base type definition for afterSalesServicesType
  **/
	export type AfterSalesServicesType =  {
			/**
  * 
  * @type {EditAnOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType}
  **/
 impliedWarranty : EditAnOfferActionResType.AfterSalesServicesType.ImpliedWarrantyType;
			/**
  * 
  * @type {EditAnOfferActionResType.AfterSalesServicesType.ReturnPolicyType}
  **/
 returnPolicy : EditAnOfferActionResType.AfterSalesServicesType.ReturnPolicyType;
			/**
  * 
  * @type {EditAnOfferActionResType.AfterSalesServicesType.WarrantyType}
  **/
 warranty : EditAnOfferActionResType.AfterSalesServicesType.WarrantyType;
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
  * @type {EditAnOfferActionResType.DiscountsType.WholesalePriceListType}
  **/
 wholesalePriceList : EditAnOfferActionResType.DiscountsType.WholesalePriceListType;
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
  * @type {EditAnOfferActionResType.TaxSettingsType.RatesType[]}
  **/
 rates : EditAnOfferActionResType.TaxSettingsType.RatesType[];
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
	/**
  * The base type definition for descriptionType
  **/
	export type DescriptionType =  {
			/**
  * 
  * @type {EditAnOfferActionResType.DescriptionType.SectionsType[]}
  **/
 sections : EditAnOfferActionResType.DescriptionType.SectionsType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DescriptionType {
	/**
  * The base type definition for sectionsType
  **/
	export type SectionsType =  {
			/**
  * 
  * @type {EditAnOfferActionResType.DescriptionType.SectionsType.ItemsType[]}
  **/
 items : EditAnOfferActionResType.DescriptionType.SectionsType.ItemsType[];
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
}