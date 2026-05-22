/**
 * Handy tool to create a final callable url, from query string, query params,
 * and the actual url.
 * @param template
 * @param params
 * @param qs
 * @returns
 */
export function buildUrl(url, params, qs) {
  // Replace :placeholders
  if (params) {
    Object.entries(params).forEach(([key, value]) => {
      url = url.replace(
        new RegExp(`:${key}`, "g"),
        encodeURIComponent(String(value))
      );
    });
  }
  if (qs && qs instanceof URLSearchParams) {
    url += `?${qs.toString()}`;
  } else if (qs && Object.keys(qs).length) {
    const query = new URLSearchParams(
      Object.entries(qs).map(([k, v]) => [k, String(v)])
    ).toString();
    url += `?${query}`;
  }
  return url;
}
