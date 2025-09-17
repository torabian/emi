import { AverageDto } from './AverageDto';
import { ComputeDto } from './ComputeDto';
import { buildUrl } from './sdk/common/buildUrl';
import { fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { type UseMutationOptions, useMutation } from '@tanstack/react-query';
import { useState } from 'react';
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
	ComputeAverageActionOptions
& {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  }
& Partial<{
	creatorFn: (item: unknown) => AverageDto
}>
export const useComputeAverageAction = (
	options?: ComputeAverageActionMutationOptions
) => {
 	const [isCompleted, setCompleteState] = useState(false);
	const mutationResult =  useMutation({
		mutationFn: (body: ComputeDto) =>
			ComputeAverageAction.Fetch(
				options?.creatorFn,
				options?.qs,
				{
					body,
					headers: options?.headers,
				},
				options?.onMessage,
				options?.overrideUrl,
			).then((x) => {
				x.done.then(() => {
					setCompleteState(true);
				});
				return x.response;
			}),
		...(options || {}),
	});
	return {
		...mutationResult,
		isCompleted
	}
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
				(item) => creatorFn(item),
				onMessage,
				init?.signal,
			);
	}
  static Definition = {
  "name": "computeAverage",
  "in": {
    "dto": "ComputeDto"
  },
  "out": {
    "dto": "AverageDto"
  }
}
}