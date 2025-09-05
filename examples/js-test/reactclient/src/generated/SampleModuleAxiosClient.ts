import { Axios, type AxiosRequestConfig, type AxiosResponse } from 'axios';
import { GetSinglePostAction, GetSinglePostActionQueryParams, GetSinglePostActionReqHeaders, GetSinglePostActionRes, type GetSinglePostActionPathParameter } from './GetSinglePostAction';
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
		GetSinglePostAction(params: GetSinglePostActionPathParameter, config?: TypedAxiosRequestConfig<
			unknown,
			GetSinglePostActionQueryParams,
			GetSinglePostActionReqHeaders
		>) {
		 	const url = GetSinglePostAction.NewUrl(
				params
			)
			return this.request<GetSinglePostActionRes, TypedAxiosResponse<GetSinglePostActionRes, unknown, GetSinglePostActionReqHeaders>>(
				{
					url,
					method: GetSinglePostAction.Method,
					...(config || {})
				} as never
			)
		}
}