import { AverageDto } from './AverageDto';
import { ComputeDto } from './ComputeDto';
import { SSEFetch, buildUrl, fetchx, handleFetchResponse, isPlausibleObject, type TypedRequestInit, withPrefix } from './sdk/js';
import { type AxiosInstance, type AxiosRequestConfig, type AxiosResponse } from 'axios';
/**
* Action to communicate with the action computeAverage
*/
export type ComputeAverageActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * ComputeAverageAction
 */
export class ComputeAverageAction {
  static URL = '';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		ComputeAverageAction.URL,
		 undefined,
		qs
	);
  static Method = '';
	static Fetch$ = async (
		qs?: URLSearchParams,
		init?: TypedRequestInit<ComputeDto, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<AverageDto, ComputeDto, unknown>(
			overrideUrl ?? ComputeAverageAction.NewUrl(
				qs
			),
			{
				method: ComputeAverageAction.Method,
				...(init || {})
			}
		)
	}
	static Fetch = async (
		qs?: URLSearchParams,
		init?: TypedRequestInit<ComputeDto, unknown>,
		onMessage?: (ev: MessageEvent) => void,
		overrideUrl?: string,
	) => {
		const res = await ComputeAverageAction.Fetch$(
			qs,
			init,
			overrideUrl,
		);
			return handleFetchResponse(
				res, 
				AverageDto,
				onMessage,
				init?.signal,
			);
	}
	static Axios : (
		clientInstance: AxiosInstance,
		config: AxiosRequestConfig<unknown>,
	)  => Promise<AxiosResponse<unknown>> = (
		clientInstance,
		config,
	) => 
		clientInstance
		.request<unknown, AxiosResponse<unknown>, unknown>(
			{
				method: ComputeAverageAction.Method,
				...(config || {})
			}
		)
		.then((res) => {
			return {
			...res,
			// if there is a output class, create instance out of it.
			data: new AverageDto(res.data),
			};
		});
}