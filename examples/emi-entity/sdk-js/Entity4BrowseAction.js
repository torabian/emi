import { Entity4OptionalDto } from "./Entity4OptionalDto";
import { GResponse } from "./sdk/envelopes/index";
import { URLSearchParamsX } from "./sdk/common/URLSearchParamsX";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity4Browse
 */
/**
 * Entity4BrowseAction
 */
export class Entity4BrowseAction {
  //
  static URL = "/entity4/browse";
  static NewUrl = (qs) => buildUrl(Entity4BrowseAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity4BrowseAction.NewUrl(qs),
      {
        method: Entity4BrowseAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity4OptionalDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity4OptionalDto(item));
    const res = await Entity4BrowseAction.Fetch$(qs, ctx, init, overrideUrl);
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
    name: "entity4Browse",
    cliShort: "entity4-b",
    url: "/entity4/browse",
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
      'Returns "entity4" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).',
    out: {
      envelope: "GResponse",
      dto: "Entity4OptionalDto",
    },
  };
}
/**
 * Entity4BrowseActionQueryParams class
 * Auto-generated from EmiAction
 */
export class Entity4BrowseActionQueryParams extends URLSearchParamsX {
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
