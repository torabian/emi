import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
/**
* Angular service for actions
*/
@Injectable({ providedIn: 'root' })
export class SampleModuleActionsService {
	constructor(private http: HttpClient) {}
		<nil>(body: any, options: any, overrideUrl?: string) {
			/*
			// convert custom headers/query to Angular compatible objects
			const httpHeaders = new HttpHeaders(headers?.toObject() ?? {});
			const httpParams = new HttpParams({ fromObject: Object.fromEntries(query?.entries() ?? []) });
			return this.http.string<CreateWorkspaceRes>(
				this.baseUrl + 'create',
				body,
				{ headers: httpHeaders, params: httpParams }
			);
			*/
		}
}