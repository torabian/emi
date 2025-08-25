import { Injectable } from '@angular/core';
import {
  GetSinglePostRes,
  GetSinglePostResFactory,
} from '../generated/GetSinglePostAction';

export class MyGetSinglePostRes extends GetSinglePostRes {
  override getId(): number | undefined {
    return (super.getId() || 0) + 100;
  }
  constructor(data?: Partial<GetSinglePostRes>) {
    super(data);
  }
}

@Injectable()
export class MyPostResFactory extends GetSinglePostResFactory {
  create(data: any) {
    return new MyGetSinglePostRes(data);
  }
}
