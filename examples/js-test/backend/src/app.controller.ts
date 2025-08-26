import { Controller, Get, Post } from '@nestjs/common';
import { AppService } from './app.service';
import { GetSinglePostHeaders } from './generated/GetSinglePostAction';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Post()
  getHello(@GetSinglePostHeaders.Nest() headers: GetSinglePostHeaders): string {
    return '' + headers.get('accept-content');
  }

  @Get()
  home(@GetSinglePostHeaders.Nest() headers: GetSinglePostHeaders): string {
    return headers.get('accept-language') || '';
  }
}
