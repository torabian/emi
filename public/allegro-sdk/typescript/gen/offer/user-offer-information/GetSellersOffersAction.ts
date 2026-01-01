import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action Get sellers offers
*/
export type GetSellersOffersActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * GetSellersOffersAction
 */
export class GetSellersOffersAction { //
  static URL = 'https://api.{environment}/sale/offers';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		GetSellersOffersAction.URL,
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
		return fetchx<GetSellersOffersActionRes, unknown, unknown>(
			overrideUrl ?? GetSellersOffersAction.NewUrl(
				qs
			),
			{
				method: GetSellersOffersAction.Method,
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
				creatorFn?: ((item: unknown) => GetSellersOffersActionRes) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new GetSellersOffersActionRes(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GetSellersOffersActionRes(item))
		const res = await GetSellersOffersAction.Fetch$(
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
  "name": "Get sellers offers",
  "url": "https://api.{environment}/sale/offers",
  "method": "get",
  "description": "Use this resource to get the list of the seller's offers. You can use different query parameters to filter the list. Read more: PL / EN.",
  "out": {
    "fields": [
      {
        "name": "count",
        "description": "Number of offers in this page",
        "type": "int"
      },
      {
        "name": "totalCount",
        "description": "Total number of offers available",
        "type": "int"
      },
      {
        "name": "offers",
        "type": "array",
        "fields": [
          {
            "name": "id",
            "description": "Offer identifier",
            "type": "string"
          },
          {
            "name": "name",
            "description": "Offer name or title",
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
            "name": "primaryImage",
            "type": "object",
            "fields": [
              {
                "name": "url",
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
                "name": "priceAutomation",
                "type": "object",
                "fields": [
                  {
                    "name": "rule",
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
            "name": "saleInfo",
            "type": "object",
            "fields": [
              {
                "name": "currentPrice",
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
                "name": "biddersCount",
                "type": "int"
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
                "name": "sold",
                "type": "int"
              }
            ]
          },
          {
            "name": "stats",
            "type": "object",
            "fields": [
              {
                "name": "watchersCount",
                "type": "int"
              },
              {
                "name": "visitsCount",
                "type": "int"
              }
            ]
          },
          {
            "name": "publication",
            "type": "object",
            "fields": [
              {
                "name": "status",
                "type": "string"
              },
              {
                "name": "startingAt",
                "type": "string"
              },
              {
                "name": "startedAt",
                "type": "string"
              },
              {
                "name": "endingAt",
                "type": "string"
              },
              {
                "name": "endedAt",
                "type": "string"
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
            "name": "delivery",
            "type": "object",
            "fields": [
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
            "name": "additionalMarketplaces",
            "description": "Marketplace-specific extensions for offer",
            "type": "map?",
            "mapKey": "string",
            "fields": [
              {
                "name": "publication",
                "type": "object",
                "fields": [
                  {
                    "name": "state",
                    "type": "string"
                  }
                ]
              },
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
                  },
                  {
                    "name": "priceAutomation",
                    "type": "object",
                    "fields": [
                      {
                        "name": "rule",
                        "type": "object",
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
                "name": "stats",
                "type": "object",
                "fields": [
                  {
                    "name": "watchersCount",
                    "type": "int"
                  },
                  {
                    "name": "visitsCount",
                    "type": "int"
                  }
                ]
              },
              {
                "name": "stock",
                "type": "object",
                "fields": [
                  {
                    "name": "sold",
                    "type": "int"
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
  * The base class definition for getSellersOffersActionRes
  **/
export class GetSellersOffersActionRes {
		/**
  * Number of offers in this page
  * @type {number}
  **/
 #count : number  =  0
		/**
  * Number of offers in this page
  * @returns {number}
  **/
get count () { return this.#count }
/**
  * Number of offers in this page
  * @type {number}
  **/
set count (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#count = parsedValue;
		}
}
setCount (value: number) {
	this.count = value
	return this
}
		/**
  * Total number of offers available
  * @type {number}
  **/
 #totalCount : number  =  0
		/**
  * Total number of offers available
  * @returns {number}
  **/
get totalCount () { return this.#totalCount }
/**
  * Total number of offers available
  * @type {number}
  **/
set totalCount (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#totalCount = parsedValue;
		}
}
setTotalCount (value: number) {
	this.totalCount = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers}
  **/
 #offers : InstanceType<typeof GetSellersOffersActionRes.Offers>[]  =  []
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers}
  **/
get offers () { return this.#offers }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers}
  **/
set offers (value: InstanceType<typeof GetSellersOffersActionRes.Offers>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetSellersOffersActionRes.Offers) {
			this.#offers = value
		} else {
			this.#offers = value.map(item => new GetSellersOffersActionRes.Offers(item))
		}
}
setOffers (value: InstanceType<typeof GetSellersOffersActionRes.Offers>[]) {
	this.offers = value
	return this
}
/**
  * The base class definition for offers
  **/
static Offers = class Offers {
		/**
  * Offer identifier
  * @type {string}
  **/
 #id : string  =  ""
		/**
  * Offer identifier
  * @returns {string}
  **/
get id () { return this.#id }
/**
  * Offer identifier
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
  * Offer name or title
  * @type {string}
  **/
 #name : string  =  ""
		/**
  * Offer name or title
  * @returns {string}
  **/
get name () { return this.#name }
/**
  * Offer name or title
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
  * @type {GetSellersOffersActionRes.Offers.Category}
  **/
 #category ! : InstanceType<typeof GetSellersOffersActionRes.Offers.Category>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.Category}
  **/
get category () { return this.#category }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Category}
  **/
set category (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Category>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.Category) {
			this.#category = value
		} else {
			this.#category = new GetSellersOffersActionRes.Offers.Category(value)
		}
}
setCategory (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Category>) {
	this.category = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.PrimaryImage}
  **/
 #primaryImage ! : InstanceType<typeof GetSellersOffersActionRes.Offers.PrimaryImage>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.PrimaryImage}
  **/
get primaryImage () { return this.#primaryImage }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.PrimaryImage}
  **/
set primaryImage (value: InstanceType<typeof GetSellersOffersActionRes.Offers.PrimaryImage>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.PrimaryImage) {
			this.#primaryImage = value
		} else {
			this.#primaryImage = new GetSellersOffersActionRes.Offers.PrimaryImage(value)
		}
}
setPrimaryImage (value: InstanceType<typeof GetSellersOffersActionRes.Offers.PrimaryImage>) {
	this.primaryImage = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode}
  **/
 #sellingMode ! : InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.SellingMode}
  **/
get sellingMode () { return this.#sellingMode }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode}
  **/
set sellingMode (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.SellingMode) {
			this.#sellingMode = value
		} else {
			this.#sellingMode = new GetSellersOffersActionRes.Offers.SellingMode(value)
		}
}
setSellingMode (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode>) {
	this.sellingMode = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SaleInfo}
  **/
 #saleInfo ! : InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.SaleInfo}
  **/
get saleInfo () { return this.#saleInfo }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SaleInfo}
  **/
set saleInfo (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.SaleInfo) {
			this.#saleInfo = value
		} else {
			this.#saleInfo = new GetSellersOffersActionRes.Offers.SaleInfo(value)
		}
}
setSaleInfo (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo>) {
	this.saleInfo = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Stock}
  **/
 #stock ! : InstanceType<typeof GetSellersOffersActionRes.Offers.Stock>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.Stock}
  **/
get stock () { return this.#stock }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Stock}
  **/
set stock (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Stock>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.Stock) {
			this.#stock = value
		} else {
			this.#stock = new GetSellersOffersActionRes.Offers.Stock(value)
		}
}
setStock (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Stock>) {
	this.stock = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Stats}
  **/
 #stats ! : InstanceType<typeof GetSellersOffersActionRes.Offers.Stats>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.Stats}
  **/
get stats () { return this.#stats }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Stats}
  **/
set stats (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Stats>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.Stats) {
			this.#stats = value
		} else {
			this.#stats = new GetSellersOffersActionRes.Offers.Stats(value)
		}
}
setStats (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Stats>) {
	this.stats = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Publication}
  **/
 #publication ! : InstanceType<typeof GetSellersOffersActionRes.Offers.Publication>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.Publication}
  **/
get publication () { return this.#publication }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Publication}
  **/
set publication (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Publication>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.Publication) {
			this.#publication = value
		} else {
			this.#publication = new GetSellersOffersActionRes.Offers.Publication(value)
		}
}
setPublication (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Publication>) {
	this.publication = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AfterSalesServices}
  **/
 #afterSalesServices ! : InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.AfterSalesServices}
  **/
get afterSalesServices () { return this.#afterSalesServices }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AfterSalesServices}
  **/
set afterSalesServices (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.AfterSalesServices) {
			this.#afterSalesServices = value
		} else {
			this.#afterSalesServices = new GetSellersOffersActionRes.Offers.AfterSalesServices(value)
		}
}
setAfterSalesServices (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices>) {
	this.afterSalesServices = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AdditionalServices}
  **/
 #additionalServices ! : InstanceType<typeof GetSellersOffersActionRes.Offers.AdditionalServices>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.AdditionalServices}
  **/
get additionalServices () { return this.#additionalServices }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AdditionalServices}
  **/
set additionalServices (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AdditionalServices>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.AdditionalServices) {
			this.#additionalServices = value
		} else {
			this.#additionalServices = new GetSellersOffersActionRes.Offers.AdditionalServices(value)
		}
}
setAdditionalServices (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AdditionalServices>) {
	this.additionalServices = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.External}
  **/
 #external ! : InstanceType<typeof GetSellersOffersActionRes.Offers.External>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.External}
  **/
get external () { return this.#external }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.External}
  **/
set external (value: InstanceType<typeof GetSellersOffersActionRes.Offers.External>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.External) {
			this.#external = value
		} else {
			this.#external = new GetSellersOffersActionRes.Offers.External(value)
		}
}
setExternal (value: InstanceType<typeof GetSellersOffersActionRes.Offers.External>) {
	this.external = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Delivery}
  **/
 #delivery ! : InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.Delivery}
  **/
get delivery () { return this.#delivery }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Delivery}
  **/
set delivery (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.Delivery) {
			this.#delivery = value
		} else {
			this.#delivery = new GetSellersOffersActionRes.Offers.Delivery(value)
		}
}
setDelivery (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery>) {
	this.delivery = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.B2b}
  **/
 #b2b ! : InstanceType<typeof GetSellersOffersActionRes.Offers.B2b>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.B2b}
  **/
get b2b () { return this.#b2b }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.B2b}
  **/
set b2b (value: InstanceType<typeof GetSellersOffersActionRes.Offers.B2b>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.B2b) {
			this.#b2b = value
		} else {
			this.#b2b = new GetSellersOffersActionRes.Offers.B2b(value)
		}
}
setB2b (value: InstanceType<typeof GetSellersOffersActionRes.Offers.B2b>) {
	this.b2b = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.FundraisingCampaign}
  **/
 #fundraisingCampaign ! : InstanceType<typeof GetSellersOffersActionRes.Offers.FundraisingCampaign>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.FundraisingCampaign}
  **/
get fundraisingCampaign () { return this.#fundraisingCampaign }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.FundraisingCampaign}
  **/
set fundraisingCampaign (value: InstanceType<typeof GetSellersOffersActionRes.Offers.FundraisingCampaign>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.FundraisingCampaign) {
			this.#fundraisingCampaign = value
		} else {
			this.#fundraisingCampaign = new GetSellersOffersActionRes.Offers.FundraisingCampaign(value)
		}
}
setFundraisingCampaign (value: InstanceType<typeof GetSellersOffersActionRes.Offers.FundraisingCampaign>) {
	this.fundraisingCampaign = value
	return this
}
		/**
  * Marketplace-specific extensions for offer
  * @type {{[key: string]: any}}
  **/
 #additionalMarketplaces ? : {[key: string]: any}  | null  =  undefined
		/**
  * Marketplace-specific extensions for offer
  * @returns {{[key: string]: any}}
  **/
get additionalMarketplaces () { return this.#additionalMarketplaces }
/**
  * Marketplace-specific extensions for offer
  * @type {{[key: string]: any}}
  **/
set additionalMarketplaces (value: {[key: string]: any} | null | undefined) {
}
setAdditionalMarketplaces (value: {[key: string]: any} | null | undefined) {
	this.additionalMarketplaces = value
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
	* Creates an instance of GetSellersOffersActionRes.Offers.Category, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.CategoryType) {
		return new GetSellersOffersActionRes.Offers.Category(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Category, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.CategoryType>) {
		return new GetSellersOffersActionRes.Offers.Category(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.CategoryType>): InstanceType<typeof GetSellersOffersActionRes.Offers.Category> {
		return new GetSellersOffersActionRes.Offers.Category ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.Category> {
		return new GetSellersOffersActionRes.Offers.Category(this.toJSON());
	}
}
/**
  * The base class definition for primaryImage
  **/
static PrimaryImage = class PrimaryImage {
		/**
  * 
  * @type {string}
  **/
 #url : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get url () { return this.#url }
/**
  * 
  * @type {string}
  **/
set url (value: string) {
		this.#url = String(value);
}
setUrl (value: string) {
	this.url = value
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
		const d = data as Partial<PrimaryImage>;
			if (d.url !== undefined) { this.url = d.url }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				url: this.#url,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			url: 'url',
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.PrimaryImage, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.PrimaryImageType) {
		return new GetSellersOffersActionRes.Offers.PrimaryImage(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.PrimaryImage, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.PrimaryImageType>) {
		return new GetSellersOffersActionRes.Offers.PrimaryImage(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.PrimaryImageType>): InstanceType<typeof GetSellersOffersActionRes.Offers.PrimaryImage> {
		return new GetSellersOffersActionRes.Offers.PrimaryImage ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.PrimaryImage> {
		return new GetSellersOffersActionRes.Offers.PrimaryImage(this.toJSON());
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
  * @type {GetSellersOffersActionRes.Offers.SellingMode.Price}
  **/
 #price ! : InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.Price>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.SellingMode.Price}
  **/
get price () { return this.#price }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode.Price}
  **/
set price (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.Price>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.SellingMode.Price) {
			this.#price = value
		} else {
			this.#price = new GetSellersOffersActionRes.Offers.SellingMode.Price(value)
		}
}
setPrice (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.Price>) {
	this.price = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation}
  **/
 #priceAutomation ! : InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation}
  **/
get priceAutomation () { return this.#priceAutomation }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation}
  **/
set priceAutomation (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation) {
			this.#priceAutomation = value
		} else {
			this.#priceAutomation = new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation(value)
		}
}
setPriceAutomation (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation>) {
	this.priceAutomation = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice}
  **/
 #minimalPrice ! : InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice}
  **/
get minimalPrice () { return this.#minimalPrice }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice}
  **/
set minimalPrice (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice) {
			this.#minimalPrice = value
		} else {
			this.#minimalPrice = new GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice(value)
		}
}
setMinimalPrice (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice>) {
	this.minimalPrice = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode.StartingPrice}
  **/
 #startingPrice ! : InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.StartingPrice>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.SellingMode.StartingPrice}
  **/
get startingPrice () { return this.#startingPrice }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode.StartingPrice}
  **/
set startingPrice (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.StartingPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.SellingMode.StartingPrice) {
			this.#startingPrice = value
		} else {
			this.#startingPrice = new GetSellersOffersActionRes.Offers.SellingMode.StartingPrice(value)
		}
}
setStartingPrice (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.StartingPrice>) {
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
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.Price, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.SellingModeType.PriceType) {
		return new GetSellersOffersActionRes.Offers.SellingMode.Price(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.Price, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.PriceType>) {
		return new GetSellersOffersActionRes.Offers.SellingMode.Price(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.PriceType>): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.Price> {
		return new GetSellersOffersActionRes.Offers.SellingMode.Price ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.Price> {
		return new GetSellersOffersActionRes.Offers.SellingMode.Price(this.toJSON());
	}
}
/**
  * The base class definition for priceAutomation
  **/
static PriceAutomation = class PriceAutomation {
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule}
  **/
 #rule ! : InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule}
  **/
get rule () { return this.#rule }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule}
  **/
set rule (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule) {
			this.#rule = value
		} else {
			this.#rule = new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule(value)
		}
}
setRule (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule>) {
	this.rule = value
	return this
}
/**
  * The base class definition for rule
  **/
static Rule = class Rule {
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
		const d = data as Partial<Rule>;
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
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType.RuleType) {
		return new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType.RuleType>) {
		return new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType.RuleType>): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule> {
		return new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule> {
		return new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule(this.toJSON());
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
		const d = data as Partial<PriceAutomation>;
			if (d.rule !== undefined) { this.rule = d.rule }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<PriceAutomation>;
			if (!(d.rule instanceof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule)) { this.rule = new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule(d.rule || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				rule: this.#rule,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			rule$: 'rule',
get rule() {
					return withPrefix(
						"offers.sellingMode.priceAutomation.rule",
						GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Rule.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType) {
		return new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType>) {
		return new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType>): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation> {
		return new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation> {
		return new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation(this.toJSON());
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
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.SellingModeType.MinimalPriceType) {
		return new GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.MinimalPriceType>) {
		return new GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.MinimalPriceType>): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice> {
		return new GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice> {
		return new GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice(this.toJSON());
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
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.StartingPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.SellingModeType.StartingPriceType) {
		return new GetSellersOffersActionRes.Offers.SellingMode.StartingPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode.StartingPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.StartingPriceType>) {
		return new GetSellersOffersActionRes.Offers.SellingMode.StartingPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType.StartingPriceType>): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.StartingPrice> {
		return new GetSellersOffersActionRes.Offers.SellingMode.StartingPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode.StartingPrice> {
		return new GetSellersOffersActionRes.Offers.SellingMode.StartingPrice(this.toJSON());
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
			if (d.priceAutomation !== undefined) { this.priceAutomation = d.priceAutomation }
			if (d.minimalPrice !== undefined) { this.minimalPrice = d.minimalPrice }
			if (d.startingPrice !== undefined) { this.startingPrice = d.startingPrice }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<SellingMode>;
			if (!(d.price instanceof GetSellersOffersActionRes.Offers.SellingMode.Price)) { this.price = new GetSellersOffersActionRes.Offers.SellingMode.Price(d.price || {}) }	
			if (!(d.priceAutomation instanceof GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation)) { this.priceAutomation = new GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation(d.priceAutomation || {}) }	
			if (!(d.minimalPrice instanceof GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice)) { this.minimalPrice = new GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice(d.minimalPrice || {}) }	
			if (!(d.startingPrice instanceof GetSellersOffersActionRes.Offers.SellingMode.StartingPrice)) { this.startingPrice = new GetSellersOffersActionRes.Offers.SellingMode.StartingPrice(d.startingPrice || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				format: this.#format,
				price: this.#price,
				priceAutomation: this.#priceAutomation,
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
						"offers.sellingMode.price",
						GetSellersOffersActionRes.Offers.SellingMode.Price.Fields
						);
						},
			priceAutomation$: 'priceAutomation',
get priceAutomation() {
					return withPrefix(
						"offers.sellingMode.priceAutomation",
						GetSellersOffersActionRes.Offers.SellingMode.PriceAutomation.Fields
						);
						},
			minimalPrice$: 'minimalPrice',
get minimalPrice() {
					return withPrefix(
						"offers.sellingMode.minimalPrice",
						GetSellersOffersActionRes.Offers.SellingMode.MinimalPrice.Fields
						);
						},
			startingPrice$: 'startingPrice',
get startingPrice() {
					return withPrefix(
						"offers.sellingMode.startingPrice",
						GetSellersOffersActionRes.Offers.SellingMode.StartingPrice.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.SellingModeType) {
		return new GetSellersOffersActionRes.Offers.SellingMode(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SellingMode, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType>) {
		return new GetSellersOffersActionRes.Offers.SellingMode(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.SellingModeType>): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode> {
		return new GetSellersOffersActionRes.Offers.SellingMode ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.SellingMode> {
		return new GetSellersOffersActionRes.Offers.SellingMode(this.toJSON());
	}
}
/**
  * The base class definition for saleInfo
  **/
static SaleInfo = class SaleInfo {
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice}
  **/
 #currentPrice ! : InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice}
  **/
get currentPrice () { return this.#currentPrice }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice}
  **/
set currentPrice (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice) {
			this.#currentPrice = value
		} else {
			this.#currentPrice = new GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice(value)
		}
}
setCurrentPrice (value: InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice>) {
	this.currentPrice = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #biddersCount : number  =  0
		/**
  * 
  * @returns {number}
  **/
get biddersCount () { return this.#biddersCount }
/**
  * 
  * @type {number}
  **/
set biddersCount (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#biddersCount = parsedValue;
		}
}
setBiddersCount (value: number) {
	this.biddersCount = value
	return this
}
/**
  * The base class definition for currentPrice
  **/
static CurrentPrice = class CurrentPrice {
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
		const d = data as Partial<CurrentPrice>;
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
	* Creates an instance of GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.SaleInfoType.CurrentPriceType) {
		return new GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.SaleInfoType.CurrentPriceType>) {
		return new GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.SaleInfoType.CurrentPriceType>): InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice> {
		return new GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice> {
		return new GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice(this.toJSON());
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
		const d = data as Partial<SaleInfo>;
			if (d.currentPrice !== undefined) { this.currentPrice = d.currentPrice }
			if (d.biddersCount !== undefined) { this.biddersCount = d.biddersCount }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<SaleInfo>;
			if (!(d.currentPrice instanceof GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice)) { this.currentPrice = new GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice(d.currentPrice || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				currentPrice: this.#currentPrice,
				biddersCount: this.#biddersCount,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			currentPrice$: 'currentPrice',
get currentPrice() {
					return withPrefix(
						"offers.saleInfo.currentPrice",
						GetSellersOffersActionRes.Offers.SaleInfo.CurrentPrice.Fields
						);
						},
			biddersCount: 'biddersCount',
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SaleInfo, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.SaleInfoType) {
		return new GetSellersOffersActionRes.Offers.SaleInfo(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.SaleInfo, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.SaleInfoType>) {
		return new GetSellersOffersActionRes.Offers.SaleInfo(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.SaleInfoType>): InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo> {
		return new GetSellersOffersActionRes.Offers.SaleInfo ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.SaleInfo> {
		return new GetSellersOffersActionRes.Offers.SaleInfo(this.toJSON());
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
  * @type {number}
  **/
 #sold : number  =  0
		/**
  * 
  * @returns {number}
  **/
get sold () { return this.#sold }
/**
  * 
  * @type {number}
  **/
set sold (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#sold = parsedValue;
		}
}
setSold (value: number) {
	this.sold = value
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
			if (d.sold !== undefined) { this.sold = d.sold }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				available: this.#available,
				sold: this.#sold,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			available: 'available',
			sold: 'sold',
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Stock, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.StockType) {
		return new GetSellersOffersActionRes.Offers.Stock(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Stock, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.StockType>) {
		return new GetSellersOffersActionRes.Offers.Stock(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.StockType>): InstanceType<typeof GetSellersOffersActionRes.Offers.Stock> {
		return new GetSellersOffersActionRes.Offers.Stock ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.Stock> {
		return new GetSellersOffersActionRes.Offers.Stock(this.toJSON());
	}
}
/**
  * The base class definition for stats
  **/
static Stats = class Stats {
		/**
  * 
  * @type {number}
  **/
 #watchersCount : number  =  0
		/**
  * 
  * @returns {number}
  **/
get watchersCount () { return this.#watchersCount }
/**
  * 
  * @type {number}
  **/
set watchersCount (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#watchersCount = parsedValue;
		}
}
setWatchersCount (value: number) {
	this.watchersCount = value
	return this
}
		/**
  * 
  * @type {number}
  **/
 #visitsCount : number  =  0
		/**
  * 
  * @returns {number}
  **/
get visitsCount () { return this.#visitsCount }
/**
  * 
  * @type {number}
  **/
set visitsCount (value: number) {
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)
		if (!Number.isNaN(parsedValue)) {
			this.#visitsCount = parsedValue;
		}
}
setVisitsCount (value: number) {
	this.visitsCount = value
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
		const d = data as Partial<Stats>;
			if (d.watchersCount !== undefined) { this.watchersCount = d.watchersCount }
			if (d.visitsCount !== undefined) { this.visitsCount = d.visitsCount }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				watchersCount: this.#watchersCount,
				visitsCount: this.#visitsCount,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			watchersCount: 'watchersCount',
			visitsCount: 'visitsCount',
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Stats, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.StatsType) {
		return new GetSellersOffersActionRes.Offers.Stats(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Stats, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.StatsType>) {
		return new GetSellersOffersActionRes.Offers.Stats(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.StatsType>): InstanceType<typeof GetSellersOffersActionRes.Offers.Stats> {
		return new GetSellersOffersActionRes.Offers.Stats ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.Stats> {
		return new GetSellersOffersActionRes.Offers.Stats(this.toJSON());
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
 #startedAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get startedAt () { return this.#startedAt }
/**
  * 
  * @type {string}
  **/
set startedAt (value: string) {
		this.#startedAt = String(value);
}
setStartedAt (value: string) {
	this.startedAt = value
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
 #endedAt : string  =  ""
		/**
  * 
  * @returns {string}
  **/
get endedAt () { return this.#endedAt }
/**
  * 
  * @type {string}
  **/
set endedAt (value: string) {
		this.#endedAt = String(value);
}
setEndedAt (value: string) {
	this.endedAt = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Publication.Marketplaces}
  **/
 #marketplaces ! : InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.Publication.Marketplaces}
  **/
get marketplaces () { return this.#marketplaces }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Publication.Marketplaces}
  **/
set marketplaces (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.Publication.Marketplaces) {
			this.#marketplaces = value
		} else {
			this.#marketplaces = new GetSellersOffersActionRes.Offers.Publication.Marketplaces(value)
		}
}
setMarketplaces (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces>) {
	this.marketplaces = value
	return this
}
/**
  * The base class definition for marketplaces
  **/
static Marketplaces = class Marketplaces {
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base}
  **/
 #base ! : InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base}
  **/
get base () { return this.#base }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base}
  **/
set base (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base) {
			this.#base = value
		} else {
			this.#base = new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base(value)
		}
}
setBase (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base>) {
	this.base = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional}
  **/
 #additional : InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional>[]  =  []
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional}
  **/
get additional () { return this.#additional }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional}
  **/
set additional (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional>[]) {
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional) {
			this.#additional = value
		} else {
			this.#additional = value.map(item => new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional(item))
		}
}
setAdditional (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional>[]) {
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
	* Creates an instance of GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.BaseType) {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.BaseType>) {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.BaseType>): InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base> {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base> {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base(this.toJSON());
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
	* Creates an instance of GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.AdditionalType) {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.AdditionalType>) {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.AdditionalType>): InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional> {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional> {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional(this.toJSON());
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
			if (!(d.base instanceof GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base)) { this.base = new GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base(d.base || {}) }	
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
						"offers.publication.marketplaces.base",
						GetSellersOffersActionRes.Offers.Publication.Marketplaces.Base.Fields
						);
						},
			additional$: 'additional',
get additional() {
					return withPrefix(
						"offers.publication.marketplaces.additional[:i]",
						GetSellersOffersActionRes.Offers.Publication.Marketplaces.Additional.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Publication.Marketplaces, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType) {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Publication.Marketplaces, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType>) {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType>): InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces> {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.Publication.Marketplaces> {
		return new GetSellersOffersActionRes.Offers.Publication.Marketplaces(this.toJSON());
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
			if (d.status !== undefined) { this.status = d.status }
			if (d.startingAt !== undefined) { this.startingAt = d.startingAt }
			if (d.startedAt !== undefined) { this.startedAt = d.startedAt }
			if (d.endingAt !== undefined) { this.endingAt = d.endingAt }
			if (d.endedAt !== undefined) { this.endedAt = d.endedAt }
			if (d.marketplaces !== undefined) { this.marketplaces = d.marketplaces }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Publication>;
			if (!(d.marketplaces instanceof GetSellersOffersActionRes.Offers.Publication.Marketplaces)) { this.marketplaces = new GetSellersOffersActionRes.Offers.Publication.Marketplaces(d.marketplaces || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				status: this.#status,
				startingAt: this.#startingAt,
				startedAt: this.#startedAt,
				endingAt: this.#endingAt,
				endedAt: this.#endedAt,
				marketplaces: this.#marketplaces,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			status: 'status',
			startingAt: 'startingAt',
			startedAt: 'startedAt',
			endingAt: 'endingAt',
			endedAt: 'endedAt',
			marketplaces$: 'marketplaces',
get marketplaces() {
					return withPrefix(
						"offers.publication.marketplaces",
						GetSellersOffersActionRes.Offers.Publication.Marketplaces.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Publication, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.PublicationType) {
		return new GetSellersOffersActionRes.Offers.Publication(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Publication, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.PublicationType>) {
		return new GetSellersOffersActionRes.Offers.Publication(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.PublicationType>): InstanceType<typeof GetSellersOffersActionRes.Offers.Publication> {
		return new GetSellersOffersActionRes.Offers.Publication ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.Publication> {
		return new GetSellersOffersActionRes.Offers.Publication(this.toJSON());
	}
}
/**
  * The base class definition for afterSalesServices
  **/
static AfterSalesServices = class AfterSalesServices {
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty}
  **/
 #impliedWarranty ! : InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty}
  **/
get impliedWarranty () { return this.#impliedWarranty }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty}
  **/
set impliedWarranty (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty) {
			this.#impliedWarranty = value
		} else {
			this.#impliedWarranty = new GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty(value)
		}
}
setImpliedWarranty (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty>) {
	this.impliedWarranty = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy}
  **/
 #returnPolicy ! : InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy}
  **/
get returnPolicy () { return this.#returnPolicy }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy}
  **/
set returnPolicy (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy) {
			this.#returnPolicy = value
		} else {
			this.#returnPolicy = new GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy(value)
		}
}
setReturnPolicy (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy>) {
	this.returnPolicy = value
	return this
}
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty}
  **/
 #warranty ! : InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty}
  **/
get warranty () { return this.#warranty }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty}
  **/
set warranty (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty) {
			this.#warranty = value
		} else {
			this.#warranty = new GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty(value)
		}
}
setWarranty (value: InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty>) {
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
	* Creates an instance of GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ImpliedWarrantyType) {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ImpliedWarrantyType>) {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ImpliedWarrantyType>): InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty> {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty> {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty(this.toJSON());
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
	* Creates an instance of GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ReturnPolicyType) {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ReturnPolicyType>) {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ReturnPolicyType>): InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy> {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy> {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy(this.toJSON());
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
	* Creates an instance of GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.AfterSalesServicesType.WarrantyType) {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.AfterSalesServicesType.WarrantyType>) {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.AfterSalesServicesType.WarrantyType>): InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty> {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty> {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty(this.toJSON());
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
			if (!(d.impliedWarranty instanceof GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty)) { this.impliedWarranty = new GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty(d.impliedWarranty || {}) }	
			if (!(d.returnPolicy instanceof GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy)) { this.returnPolicy = new GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy(d.returnPolicy || {}) }	
			if (!(d.warranty instanceof GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty)) { this.warranty = new GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty(d.warranty || {}) }	
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
						"offers.afterSalesServices.impliedWarranty",
						GetSellersOffersActionRes.Offers.AfterSalesServices.ImpliedWarranty.Fields
						);
						},
			returnPolicy$: 'returnPolicy',
get returnPolicy() {
					return withPrefix(
						"offers.afterSalesServices.returnPolicy",
						GetSellersOffersActionRes.Offers.AfterSalesServices.ReturnPolicy.Fields
						);
						},
			warranty$: 'warranty',
get warranty() {
					return withPrefix(
						"offers.afterSalesServices.warranty",
						GetSellersOffersActionRes.Offers.AfterSalesServices.Warranty.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.AfterSalesServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.AfterSalesServicesType) {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.AfterSalesServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.AfterSalesServicesType>) {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.AfterSalesServicesType>): InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices> {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.AfterSalesServices> {
		return new GetSellersOffersActionRes.Offers.AfterSalesServices(this.toJSON());
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
	* Creates an instance of GetSellersOffersActionRes.Offers.AdditionalServices, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.AdditionalServicesType) {
		return new GetSellersOffersActionRes.Offers.AdditionalServices(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.AdditionalServices, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.AdditionalServicesType>) {
		return new GetSellersOffersActionRes.Offers.AdditionalServices(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.AdditionalServicesType>): InstanceType<typeof GetSellersOffersActionRes.Offers.AdditionalServices> {
		return new GetSellersOffersActionRes.Offers.AdditionalServices ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.AdditionalServices> {
		return new GetSellersOffersActionRes.Offers.AdditionalServices(this.toJSON());
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
	* Creates an instance of GetSellersOffersActionRes.Offers.External, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.ExternalType) {
		return new GetSellersOffersActionRes.Offers.External(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.External, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.ExternalType>) {
		return new GetSellersOffersActionRes.Offers.External(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.ExternalType>): InstanceType<typeof GetSellersOffersActionRes.Offers.External> {
		return new GetSellersOffersActionRes.Offers.External ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.External> {
		return new GetSellersOffersActionRes.Offers.External(this.toJSON());
	}
}
/**
  * The base class definition for delivery
  **/
static Delivery = class Delivery {
		/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Delivery.ShippingRates}
  **/
 #shippingRates ! : InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery.ShippingRates>
		/**
  * 
  * @returns {GetSellersOffersActionRes.Offers.Delivery.ShippingRates}
  **/
get shippingRates () { return this.#shippingRates }
/**
  * 
  * @type {GetSellersOffersActionRes.Offers.Delivery.ShippingRates}
  **/
set shippingRates (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery.ShippingRates>) {
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof GetSellersOffersActionRes.Offers.Delivery.ShippingRates) {
			this.#shippingRates = value
		} else {
			this.#shippingRates = new GetSellersOffersActionRes.Offers.Delivery.ShippingRates(value)
		}
}
setShippingRates (value: InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery.ShippingRates>) {
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
	* Creates an instance of GetSellersOffersActionRes.Offers.Delivery.ShippingRates, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.DeliveryType.ShippingRatesType) {
		return new GetSellersOffersActionRes.Offers.Delivery.ShippingRates(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Delivery.ShippingRates, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.DeliveryType.ShippingRatesType>) {
		return new GetSellersOffersActionRes.Offers.Delivery.ShippingRates(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.DeliveryType.ShippingRatesType>): InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery.ShippingRates> {
		return new GetSellersOffersActionRes.Offers.Delivery.ShippingRates ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery.ShippingRates> {
		return new GetSellersOffersActionRes.Offers.Delivery.ShippingRates(this.toJSON());
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
			if (d.shippingRates !== undefined) { this.shippingRates = d.shippingRates }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Delivery>;
			if (!(d.shippingRates instanceof GetSellersOffersActionRes.Offers.Delivery.ShippingRates)) { this.shippingRates = new GetSellersOffersActionRes.Offers.Delivery.ShippingRates(d.shippingRates || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				shippingRates: this.#shippingRates,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			shippingRates$: 'shippingRates',
get shippingRates() {
					return withPrefix(
						"offers.delivery.shippingRates",
						GetSellersOffersActionRes.Offers.Delivery.ShippingRates.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Delivery, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.DeliveryType) {
		return new GetSellersOffersActionRes.Offers.Delivery(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.Delivery, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.DeliveryType>) {
		return new GetSellersOffersActionRes.Offers.Delivery(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.DeliveryType>): InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery> {
		return new GetSellersOffersActionRes.Offers.Delivery ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.Delivery> {
		return new GetSellersOffersActionRes.Offers.Delivery(this.toJSON());
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
	* Creates an instance of GetSellersOffersActionRes.Offers.B2b, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.B2bType) {
		return new GetSellersOffersActionRes.Offers.B2b(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.B2b, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.B2bType>) {
		return new GetSellersOffersActionRes.Offers.B2b(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.B2bType>): InstanceType<typeof GetSellersOffersActionRes.Offers.B2b> {
		return new GetSellersOffersActionRes.Offers.B2b ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.B2b> {
		return new GetSellersOffersActionRes.Offers.B2b(this.toJSON());
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
	* Creates an instance of GetSellersOffersActionRes.Offers.FundraisingCampaign, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType.FundraisingCampaignType) {
		return new GetSellersOffersActionRes.Offers.FundraisingCampaign(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers.FundraisingCampaign, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType.FundraisingCampaignType>) {
		return new GetSellersOffersActionRes.Offers.FundraisingCampaign(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType.FundraisingCampaignType>): InstanceType<typeof GetSellersOffersActionRes.Offers.FundraisingCampaign> {
		return new GetSellersOffersActionRes.Offers.FundraisingCampaign ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers.FundraisingCampaign> {
		return new GetSellersOffersActionRes.Offers.FundraisingCampaign(this.toJSON());
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
		const d = data as Partial<Offers>;
			if (d.id !== undefined) { this.id = d.id }
			if (d.name !== undefined) { this.name = d.name }
			if (d.category !== undefined) { this.category = d.category }
			if (d.primaryImage !== undefined) { this.primaryImage = d.primaryImage }
			if (d.sellingMode !== undefined) { this.sellingMode = d.sellingMode }
			if (d.saleInfo !== undefined) { this.saleInfo = d.saleInfo }
			if (d.stock !== undefined) { this.stock = d.stock }
			if (d.stats !== undefined) { this.stats = d.stats }
			if (d.publication !== undefined) { this.publication = d.publication }
			if (d.afterSalesServices !== undefined) { this.afterSalesServices = d.afterSalesServices }
			if (d.additionalServices !== undefined) { this.additionalServices = d.additionalServices }
			if (d.external !== undefined) { this.external = d.external }
			if (d.delivery !== undefined) { this.delivery = d.delivery }
			if (d.b2b !== undefined) { this.b2b = d.b2b }
			if (d.fundraisingCampaign !== undefined) { this.fundraisingCampaign = d.fundraisingCampaign }
			if (d.additionalMarketplaces !== undefined) { this.additionalMarketplaces = d.additionalMarketplaces }
		this.#lateInitFields(data)
	}
	/**
	 * These are the class instances, which need to be initialised, regardless of the constructor incoming data
	**/
	#lateInitFields(data = {}) {
		const d = data as Partial<Offers>;
			if (!(d.category instanceof GetSellersOffersActionRes.Offers.Category)) { this.category = new GetSellersOffersActionRes.Offers.Category(d.category || {}) }	
			if (!(d.primaryImage instanceof GetSellersOffersActionRes.Offers.PrimaryImage)) { this.primaryImage = new GetSellersOffersActionRes.Offers.PrimaryImage(d.primaryImage || {}) }	
			if (!(d.sellingMode instanceof GetSellersOffersActionRes.Offers.SellingMode)) { this.sellingMode = new GetSellersOffersActionRes.Offers.SellingMode(d.sellingMode || {}) }	
			if (!(d.saleInfo instanceof GetSellersOffersActionRes.Offers.SaleInfo)) { this.saleInfo = new GetSellersOffersActionRes.Offers.SaleInfo(d.saleInfo || {}) }	
			if (!(d.stock instanceof GetSellersOffersActionRes.Offers.Stock)) { this.stock = new GetSellersOffersActionRes.Offers.Stock(d.stock || {}) }	
			if (!(d.stats instanceof GetSellersOffersActionRes.Offers.Stats)) { this.stats = new GetSellersOffersActionRes.Offers.Stats(d.stats || {}) }	
			if (!(d.publication instanceof GetSellersOffersActionRes.Offers.Publication)) { this.publication = new GetSellersOffersActionRes.Offers.Publication(d.publication || {}) }	
			if (!(d.afterSalesServices instanceof GetSellersOffersActionRes.Offers.AfterSalesServices)) { this.afterSalesServices = new GetSellersOffersActionRes.Offers.AfterSalesServices(d.afterSalesServices || {}) }	
			if (!(d.additionalServices instanceof GetSellersOffersActionRes.Offers.AdditionalServices)) { this.additionalServices = new GetSellersOffersActionRes.Offers.AdditionalServices(d.additionalServices || {}) }	
			if (!(d.external instanceof GetSellersOffersActionRes.Offers.External)) { this.external = new GetSellersOffersActionRes.Offers.External(d.external || {}) }	
			if (!(d.delivery instanceof GetSellersOffersActionRes.Offers.Delivery)) { this.delivery = new GetSellersOffersActionRes.Offers.Delivery(d.delivery || {}) }	
			if (!(d.b2b instanceof GetSellersOffersActionRes.Offers.B2b)) { this.b2b = new GetSellersOffersActionRes.Offers.B2b(d.b2b || {}) }	
			if (!(d.fundraisingCampaign instanceof GetSellersOffersActionRes.Offers.FundraisingCampaign)) { this.fundraisingCampaign = new GetSellersOffersActionRes.Offers.FundraisingCampaign(d.fundraisingCampaign || {}) }	
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				id: this.#id,
				name: this.#name,
				category: this.#category,
				primaryImage: this.#primaryImage,
				sellingMode: this.#sellingMode,
				saleInfo: this.#saleInfo,
				stock: this.#stock,
				stats: this.#stats,
				publication: this.#publication,
				afterSalesServices: this.#afterSalesServices,
				additionalServices: this.#additionalServices,
				external: this.#external,
				delivery: this.#delivery,
				b2b: this.#b2b,
				fundraisingCampaign: this.#fundraisingCampaign,
				additionalMarketplaces: this.#additionalMarketplaces,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			id: 'id',
			name: 'name',
			category$: 'category',
get category() {
					return withPrefix(
						"offers.category",
						GetSellersOffersActionRes.Offers.Category.Fields
						);
						},
			primaryImage$: 'primaryImage',
get primaryImage() {
					return withPrefix(
						"offers.primaryImage",
						GetSellersOffersActionRes.Offers.PrimaryImage.Fields
						);
						},
			sellingMode$: 'sellingMode',
get sellingMode() {
					return withPrefix(
						"offers.sellingMode",
						GetSellersOffersActionRes.Offers.SellingMode.Fields
						);
						},
			saleInfo$: 'saleInfo',
get saleInfo() {
					return withPrefix(
						"offers.saleInfo",
						GetSellersOffersActionRes.Offers.SaleInfo.Fields
						);
						},
			stock$: 'stock',
get stock() {
					return withPrefix(
						"offers.stock",
						GetSellersOffersActionRes.Offers.Stock.Fields
						);
						},
			stats$: 'stats',
get stats() {
					return withPrefix(
						"offers.stats",
						GetSellersOffersActionRes.Offers.Stats.Fields
						);
						},
			publication$: 'publication',
get publication() {
					return withPrefix(
						"offers.publication",
						GetSellersOffersActionRes.Offers.Publication.Fields
						);
						},
			afterSalesServices$: 'afterSalesServices',
get afterSalesServices() {
					return withPrefix(
						"offers.afterSalesServices",
						GetSellersOffersActionRes.Offers.AfterSalesServices.Fields
						);
						},
			additionalServices$: 'additionalServices',
get additionalServices() {
					return withPrefix(
						"offers.additionalServices",
						GetSellersOffersActionRes.Offers.AdditionalServices.Fields
						);
						},
			external$: 'external',
get external() {
					return withPrefix(
						"offers.external",
						GetSellersOffersActionRes.Offers.External.Fields
						);
						},
			delivery$: 'delivery',
get delivery() {
					return withPrefix(
						"offers.delivery",
						GetSellersOffersActionRes.Offers.Delivery.Fields
						);
						},
			b2b$: 'b2b',
get b2b() {
					return withPrefix(
						"offers.b2b",
						GetSellersOffersActionRes.Offers.B2b.Fields
						);
						},
			fundraisingCampaign$: 'fundraisingCampaign',
get fundraisingCampaign() {
					return withPrefix(
						"offers.fundraisingCampaign",
						GetSellersOffersActionRes.Offers.FundraisingCampaign.Fields
						);
						},
			additionalMarketplaces: 'additionalMarketplaces',
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType.OffersType) {
		return new GetSellersOffersActionRes.Offers(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes.Offers, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType.OffersType>) {
		return new GetSellersOffersActionRes.Offers(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType.OffersType>): InstanceType<typeof GetSellersOffersActionRes.Offers> {
		return new GetSellersOffersActionRes.Offers ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes.Offers> {
		return new GetSellersOffersActionRes.Offers(this.toJSON());
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
		const d = data as Partial<GetSellersOffersActionRes>;
			if (d.count !== undefined) { this.count = d.count }
			if (d.totalCount !== undefined) { this.totalCount = d.totalCount }
			if (d.offers !== undefined) { this.offers = d.offers }
	}
	/**
	*	Special toJSON override, since the field are private,
	*	Json stringify won't see them unless we mention it explicitly.
	**/
	toJSON() {
    	return { 
				count: this.#count,
				totalCount: this.#totalCount,
				offers: this.#offers,
		};
  	}
	toString() {
		return JSON.stringify(this);
	}
	static get Fields() {
      return {
			count: 'count',
			totalCount: 'totalCount',
			offers$: 'offers',
get offers() {
					return withPrefix(
						"offers[:i]",
						GetSellersOffersActionRes.Offers.Fields
						);
						},
	  }
	}
	/**
	* Creates an instance of GetSellersOffersActionRes, and possibleDtoObject
	* needs to satisfy the type requirement fully, otherwise typescript compile would
	* be complaining.
	**/
	static from(possibleDtoObject: GetSellersOffersActionResType) {
		return new GetSellersOffersActionRes(possibleDtoObject);
	}
	/**
	* Creates an instance of GetSellersOffersActionRes, and partialDtoObject
	* needs to satisfy the type, but partially, and rest of the content would
	* be constructed according to data types and nullability.
	**/
	static with(partialDtoObject: PartialDeep<GetSellersOffersActionResType>) {
		return new GetSellersOffersActionRes(partialDtoObject);
	}
	copyWith(partial: PartialDeep<GetSellersOffersActionResType>): InstanceType<typeof GetSellersOffersActionRes> {
		return new GetSellersOffersActionRes ({ ...this.toJSON(), ...partial });
	}
	clone(): InstanceType<typeof GetSellersOffersActionRes> {
		return new GetSellersOffersActionRes(this.toJSON());
	}
}
export abstract class GetSellersOffersActionResFactory {
	abstract create(data: unknown): GetSellersOffersActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
	/**
  * The base type definition for getSellersOffersActionRes
  **/
	export type GetSellersOffersActionResType =  {
			/**
  * Number of offers in this page
  * @type {number}
  **/
 count : number;
			/**
  * Total number of offers available
  * @type {number}
  **/
 totalCount : number;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType[]}
  **/
 offers : GetSellersOffersActionResType.OffersType[];
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GetSellersOffersActionResType {
	/**
  * The base type definition for offersType
  **/
	export type OffersType =  {
			/**
  * Offer identifier
  * @type {string}
  **/
 id : string;
			/**
  * Offer name or title
  * @type {string}
  **/
 name : string;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.CategoryType}
  **/
 category : GetSellersOffersActionResType.OffersType.CategoryType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.PrimaryImageType}
  **/
 primaryImage : GetSellersOffersActionResType.OffersType.PrimaryImageType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.SellingModeType}
  **/
 sellingMode : GetSellersOffersActionResType.OffersType.SellingModeType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.SaleInfoType}
  **/
 saleInfo : GetSellersOffersActionResType.OffersType.SaleInfoType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.StockType}
  **/
 stock : GetSellersOffersActionResType.OffersType.StockType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.StatsType}
  **/
 stats : GetSellersOffersActionResType.OffersType.StatsType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.PublicationType}
  **/
 publication : GetSellersOffersActionResType.OffersType.PublicationType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.AfterSalesServicesType}
  **/
 afterSalesServices : GetSellersOffersActionResType.OffersType.AfterSalesServicesType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.AdditionalServicesType}
  **/
 additionalServices : GetSellersOffersActionResType.OffersType.AdditionalServicesType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.ExternalType}
  **/
 external : GetSellersOffersActionResType.OffersType.ExternalType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.DeliveryType}
  **/
 delivery : GetSellersOffersActionResType.OffersType.DeliveryType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.B2bType}
  **/
 b2b : GetSellersOffersActionResType.OffersType.B2bType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.FundraisingCampaignType}
  **/
 fundraisingCampaign : GetSellersOffersActionResType.OffersType.FundraisingCampaignType;
			/**
  * Marketplace-specific extensions for offer
  * @type {{[key: string]: any}}
  **/
 additionalMarketplaces ?: {[key: string]: any};
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OffersType {
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
  * The base type definition for primaryImageType
  **/
	export type PrimaryImageType =  {
			/**
  * 
  * @type {string}
  **/
 url : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PrimaryImageType {
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
  * @type {GetSellersOffersActionResType.OffersType.SellingModeType.PriceType}
  **/
 price : GetSellersOffersActionResType.OffersType.SellingModeType.PriceType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType}
  **/
 priceAutomation : GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.SellingModeType.MinimalPriceType}
  **/
 minimalPrice : GetSellersOffersActionResType.OffersType.SellingModeType.MinimalPriceType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.SellingModeType.StartingPriceType}
  **/
 startingPrice : GetSellersOffersActionResType.OffersType.SellingModeType.StartingPriceType;
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
  * The base type definition for priceAutomationType
  **/
	export type PriceAutomationType =  {
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType.RuleType}
  **/
 rule : GetSellersOffersActionResType.OffersType.SellingModeType.PriceAutomationType.RuleType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PriceAutomationType {
	/**
  * The base type definition for ruleType
  **/
	export type RuleType =  {
			/**
  * 
  * @type {string}
  **/
 id : string;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace RuleType {
}
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
  * The base type definition for saleInfoType
  **/
	export type SaleInfoType =  {
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.SaleInfoType.CurrentPriceType}
  **/
 currentPrice : GetSellersOffersActionResType.OffersType.SaleInfoType.CurrentPriceType;
			/**
  * 
  * @type {number}
  **/
 biddersCount : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SaleInfoType {
	/**
  * The base type definition for currentPriceType
  **/
	export type CurrentPriceType =  {
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
export namespace CurrentPriceType {
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
  * @type {number}
  **/
 sold : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace StockType {
}
	/**
  * The base type definition for statsType
  **/
	export type StatsType =  {
			/**
  * 
  * @type {number}
  **/
 watchersCount : number;
			/**
  * 
  * @type {number}
  **/
 visitsCount : number;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace StatsType {
}
	/**
  * The base type definition for publicationType
  **/
	export type PublicationType =  {
			/**
  * 
  * @type {string}
  **/
 status : string;
			/**
  * 
  * @type {string}
  **/
 startingAt : string;
			/**
  * 
  * @type {string}
  **/
 startedAt : string;
			/**
  * 
  * @type {string}
  **/
 endingAt : string;
			/**
  * 
  * @type {string}
  **/
 endedAt : string;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType}
  **/
 marketplaces : GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType;
	}
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PublicationType {
	/**
  * The base type definition for marketplacesType
  **/
	export type MarketplacesType =  {
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.BaseType}
  **/
 base : GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.BaseType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.AdditionalType[]}
  **/
 additional : GetSellersOffersActionResType.OffersType.PublicationType.MarketplacesType.AdditionalType[];
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
  * The base type definition for afterSalesServicesType
  **/
	export type AfterSalesServicesType =  {
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ImpliedWarrantyType}
  **/
 impliedWarranty : GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ImpliedWarrantyType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ReturnPolicyType}
  **/
 returnPolicy : GetSellersOffersActionResType.OffersType.AfterSalesServicesType.ReturnPolicyType;
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.AfterSalesServicesType.WarrantyType}
  **/
 warranty : GetSellersOffersActionResType.OffersType.AfterSalesServicesType.WarrantyType;
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
  * The base type definition for deliveryType
  **/
	export type DeliveryType =  {
			/**
  * 
  * @type {GetSellersOffersActionResType.OffersType.DeliveryType.ShippingRatesType}
  **/
 shippingRates : GetSellersOffersActionResType.OffersType.DeliveryType.ShippingRatesType;
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
}
}