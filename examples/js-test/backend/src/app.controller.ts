import { Controller, Get, Post } from '@nestjs/common';
import { AppService } from './app.service';
import { MyactionHeaders, MyactionReq } from './generated/MyactionAction';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Post()
  getHello(
    @MyactionReq.Nest() req: MyactionReq,
    @MyactionHeaders.Nest() headers: MyactionHeaders,
  ): string {
    return '' + req.getUser().firstName;
  }

  @Get()
  home(@MyactionHeaders.Nest() headers: MyactionHeaders): string {
    return headers.get('accept-language') || '';
  }
}
