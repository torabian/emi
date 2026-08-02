import { Entity2OptionalDto } from "./Entity2OptionalDto";
import { GResponse } from "./sdk/envelopes/index";
import { URLSearchParamsX } from "./sdk/common/URLSearchParamsX";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity2Browse
 */
/**
 * Entity2BrowseAction
 */
export class Entity2BrowseAction {
  //
  static URL = "/entity2/browse";
  static NewUrl = (qs) => buildUrl(Entity2BrowseAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity2BrowseAction.NewUrl(qs),
      {
        method: Entity2BrowseAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity2OptionalDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity2OptionalDto(item));
    const res = await Entity2BrowseAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse();
        if (creatorFn) {
          resp.setCreator(creatorFn);
        }
        resp.inject(data);
        return resp;
      },
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "entity2Browse",
    cliShort: "entity2-b",
    url: "/entity2/browse",
    method: "get",
    qs: [
      {
        name: "filter",
        type: "string",
      },
      {
        name: "sort",
        type: "string",
      },
      {
        name: "startIndex",
        type: "int",
      },
      {
        name: "itemsPerPage",
        type: "int",
      },
      {
        name: "cursor",
        type: "string",
      },
    ],
    description:
      'Returns "entity2" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).',
    out: {
      envelope: "GResponse",
      dto: "Entity2OptionalDto",
    },
  };
}
/**
 * Entity2BrowseActionQueryParams class
 * Auto-generated from EmiAction
 */
export class Entity2BrowseActionQueryParams extends URLSearchParamsX {
  /**
   *
   * @returns { string | null }
   */
  getFilter() {
    return this.getTyped("filter", "string | null");
  }
  /**
   *
   * @param { string | null } value
   */
  setFilter(value: string | null) {
    this.set("filter", value);
    return this;
  }
  /**
   *
   * @returns { string | null }
   */
  getSort() {
    return this.getTyped("sort", "string | null");
  }
  /**
   *
   * @param { string | null } value
   */
  setSort(value: string | null) {
    this.set("sort", value);
    return this;
  }
  /**
   *
   * @returns { number | null }
   */
  getStartIndex() {
    return this.getTyped("startIndex", "number | null");
  }
  /**
   *
   * @param { number | null } value
   */
  setStartIndex(value: number | null) {
    this.set("startIndex", value);
    return this;
  }
  /**
   *
   * @returns { number | null }
   */
  getItemsPerPage() {
    return this.getTyped("itemsPerPage", "number | null");
  }
  /**
   *
   * @param { number | null } value
   */
  setItemsPerPage(value: number | null) {
    this.set("itemsPerPage", value);
    return this;
  }
  /**
   *
   * @returns { string | null }
   */
  getCursor() {
    return this.getTyped("cursor", "string | null");
  }
  /**
   *
   * @param { string | null } value
   */
  setCursor(value: string | null) {
    this.set("cursor", value);
    return this;
  }
}
