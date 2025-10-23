import { AverageDto } from './AverageDto';
import { ComputeDto } from './ComputeDto';
import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit, type TypedResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
import { type UseMutationOptions, useMutation } from '@tanstack/react-query';
import { useFetchxContext } from './sdk/react/useFetchx';
import { useState } from 'react';
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
	ctx?: FetchxContext;
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
			body: ComputeDto
	) =>
		{
			setCompleteState(false);
			return ComputeAverageAction.Fetch(
				{
						body,
					headers: options?.headers,
				},
				{
					creatorFn: options?.creatorFn,
					qs: options?.qs,
					ctx,
					onMessage: options?.onMessage,
					overrideUrl: options?.overrideUrl,
				}
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
		ctx?: FetchxContext,
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
		init?: TypedRequestInit<ComputeDto, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => AverageDto) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new AverageDto(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new AverageDto(item))
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