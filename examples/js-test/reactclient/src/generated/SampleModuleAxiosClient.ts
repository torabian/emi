import { Axios, type AxiosRequestConfig, type AxiosResponse } from 'axios';
import { FetchGetSinglePostAction, GetSinglePostQueryParams, GetSinglePostReqHeaders, GetSinglePostRes, type FetchGetSinglePostActionPathParameter } from './GetSinglePostAction';
import { FetchSampleSseAction, SampleSseQueryParams, SampleSseReqHeaders, SampleSseRes } from './SampleSseAction';
/**
* Axios bundle service
*/
import type {
  AxiosResponseHeaders,
  AxiosRequestConfig as BaseAxiosRequestConfig,
  RawAxiosResponseHeaders,
} from "axios";
export interface TypedAxiosRequestConfig<D = unknown, P = unknown, H = unknown>
  extends Omit<BaseAxiosRequestConfig, "data" | "params" | "headers"> {
  data?: D;
  params?: P;
  headers?: H;
}
export interface TypedAxiosResponse<
  T = unknown,
  D = unknown,
  H = RawAxiosResponseHeaders | AxiosResponseHeaders
> extends Omit<AxiosResponse<T, D>, "headers"> {
  headers: H;
}
export class SampleModuleAxiosClient extends Axios {
  static create(config?: AxiosRequestConfig) {
    return new SampleModuleAxiosClient(config);
  }
		getSinglePost(params: FetchGetSinglePostActionPathParameter, config?: TypedAxiosRequestConfig<
			unknown,
			GetSinglePostQueryParams,
			GetSinglePostReqHeaders
		>) {
		 	const url = FetchGetSinglePostAction.NewUrl(
				params
			)
			return this.request<GetSinglePostRes, TypedAxiosResponse<GetSinglePostRes, unknown, GetSinglePostReqHeaders>>(
				{
					url,
					method: FetchGetSinglePostAction.Method,
					...(config || {})
				} as any
			)
		}
		sampleSse(config?: TypedAxiosRequestConfig<
			unknown,
			SampleSseQueryParams,
			SampleSseReqHeaders
		>) {
		 	const url = FetchSampleSseAction.NewUrl(
			)
			return this.request<SampleSseRes, TypedAxiosResponse<SampleSseRes, unknown, SampleSseReqHeaders>>(
				{
					url,
					method: FetchSampleSseAction.Method,
					...(config || {})
				} as any
			)
		}
}