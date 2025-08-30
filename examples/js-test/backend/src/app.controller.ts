import { Controller, Get, MessageEvent, Post, Sse } from '@nestjs/common';
import { AppService } from './app.service';
import { GetSinglePostReqHeaders } from './generated/GetSinglePostAction';
import { interval, map, Observable, take } from 'rxjs';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Post()
  getHello(
    @GetSinglePostReqHeaders.Nest() headers: GetSinglePostReqHeaders,
  ): string {
    return '' + headers.get('accept-content');
  }

  @Get()
  home(
    @GetSinglePostReqHeaders.Nest() headers: GetSinglePostReqHeaders,
  ): string {
    return headers.get('accept-language') || '';
  }

  @Sse('stream')
  stream(): Observable<MessageEvent> {
    return interval(1000).pipe(
      map((i) => ({
        data: { message: `Mock event #${i}` }, // payload
        // event: 'custom',                    // optional custom event type
        // id: String(i),                      // optional event ID
      })),
    );
  }
}
