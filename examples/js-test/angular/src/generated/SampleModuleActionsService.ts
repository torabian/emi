import {
  FetchGetSinglePostAction,
  FetchGetSinglePostActionPathParameter,
  GetSinglePostHeaders,
  GetSinglePostQueryParams,
  GetSinglePostRes,
  GetSinglePostResFactory,
} from './GetSinglePostAction';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Inject, Injectable, Optional } from '@angular/core';
import { Observable, map } from 'rxjs';
/**
 * Angular service for actions
 */
@Injectable({ providedIn: 'root' })
export class SampleModuleActionsService {
  constructor(
    private http: HttpClient,
    @Optional()
    @Inject(GetSinglePostResFactory)
    private GetSinglePostResFactory: GetSinglePostResFactory
  ) {}
  getSinglePost(
    params: FetchGetSinglePostActionPathParameter,
    options?: Parameters<HttpClient['get']>[1] & {
      params?: GetSinglePostQueryParams;
      headers?: GetSinglePostHeaders;
    }
  ) {
    return this.http
      .get<GetSinglePostRes>(FetchGetSinglePostAction.NewUrl(params), options)
      .pipe(
        map((res) => {
          return this.GetSinglePostResFactory
            ? this.GetSinglePostResFactory.create(res)
            : new GetSinglePostRes(res);
        })
      );
  }
}
