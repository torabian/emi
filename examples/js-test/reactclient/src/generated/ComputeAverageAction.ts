import { AverageDto } from './AverageDto';
import { ComputeDto } from './ComputeDto';
import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { type UseMutationOptions, useMutation } from '@tanstack/react-query';
import { useFetchxContext } from './sdk/react/useFetchx';
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
	ctx?: FetchxContext<unknown, unknown>;
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
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;
	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState<TypedResponse<unknown>>();
	const fn = (
			body: ComputeDto,
	) =>
		{
			setCompleteState(false);
			ComputeAverageAction.Fetch(
				options?.creatorFn,
				options?.qs,
				{
						body,
					headers: options?.headers,
				},
				options?.onMessage,
				options?.overrideUrl,
				ctx,
			).then((x) => {
				x.done.then(() => {
					setCompleteState(true);
				});
				setResponse(x.response)
				return x.response.result;
			})
		}
	const result =  useMutation({
		mutationFn: fn,
		...(options || {}),
	});
	return {
		...result,
		isCompleted,
		response
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
		ctx?: FetchxContext<unknown, unknown>,
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
			},
			ctx
		)
	}
	static Fetch = async (
			creatorFn: (item: unknown) => AverageDto = (item) => new AverageDto(item),
		qs?: URLSearchParams,
		ctx?: FetchxContext<unknown, unknown>,
		init?: TypedRequestInit<ComputeDto, unknown>,
		onMessage?: (ev: MessageEvent) => void,
		overrideUrl?: string,
	) => {
		const res = await ComputeAverageAction.Fetch$(
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
  "name": "computeAverage",
  "in": {
    "dto": "ComputeDto"
  },
  "out": {
    "dto": "AverageDto"
  }
}
}