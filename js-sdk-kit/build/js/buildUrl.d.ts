/**
 * Handy tool to create a final callable url, from query string, query params,
 * and the actual url.
 * @param template
 * @param params
 * @param qs
 * @returns
 */
export declare function buildUrl(url: string, params?: Record<string, unknown>, qs?: URLSearchParams): string;
