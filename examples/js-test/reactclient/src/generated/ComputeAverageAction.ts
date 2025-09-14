import { AverageDto } from './AverageDto';
import { ComputeDto } from './ComputeDto';
import { buildUrl } from './sdk/common/buildUrl';
import { fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { type AxiosInstance, type AxiosRequestConfig, type AxiosResponse } from 'axios';
import { type UseMutationOptions, useMutation } from '@tanstack/react-query';
import { withPrefix } from './sdk/common/withPrefix';
/**
* Action to communicate with the action computeAverage
*/
export type ComputeAverageActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
export type ComputeAverageActionMutationOptions = Omit<
	UseMutationOptions<unknown, unknown, unknown, unknown>,
	"mutationFn"
> &
	ComputeAverageActionOptions;
export const useComputeAverageActionMutation = (
	options: ComputeAverageActionMutationOptions
) => {
	return useMutation({
		mutationFn: (vars: unknown) =>
			ComputeAverageAction.Fetch(
				options.qs,
				{
					body: vars,
					headers: options.headers,
				}
			),
		...(options || {}),
	});
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
			creatorFn: (item: unknown) => AverageDto = (item) => new AverageDto(item),
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