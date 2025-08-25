import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';

import { provideHttpClient } from '@angular/common/http';
import { routes } from './app.routes';
import { MyPostResFactory } from './app.overrides';
import { GetSinglePostResFactory } from '../generated/GetSinglePostAction';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes),
    provideHttpClient(),
    [{ provide: GetSinglePostResFactory, useClass: MyPostResFactory }],
  ],
};
