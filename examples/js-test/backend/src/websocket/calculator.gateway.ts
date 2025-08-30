// src/calculator/calculator.gateway.ts
import {
  WebSocketGateway,
  WebSocketServer,
  SubscribeMessage,
  MessageBody,
  OnGatewayConnection,
  OnGatewayDisconnect,
} from '@nestjs/websockets';
import { Server, Socket } from 'socket.io';
import { Injectable } from '@nestjs/common';

// DTO for request
export interface CalcRequest {
  op: '+' | '-' | '*' | '/';
  a: number;
  b: number;
}

// DTO for response
export interface CalcResponse {
  result: number;
}

@WebSocketGateway({ cors: true })
@Injectable()
export class CalculatorGateway
  implements OnGatewayConnection, OnGatewayDisconnect
{
  @WebSocketServer()
  server: Server;

  private lastResult = 0;
  private intervals = new Map<string, NodeJS.Timer>();

  handleConnection(client: Socket) {
    // send last result every second
    const timer = setInterval(() => {
      client.emit('lastResult', { result: this.lastResult });
    }, 1000);

    this.intervals.set(client.id, timer);
  }

  handleDisconnect(client: Socket) {
    const timer = this.intervals.get(client.id);
    if (timer) clearInterval(timer);
    this.intervals.delete(client.id);
  }

  @SubscribeMessage('calculate')
  handleCalculate(client: Socket, payload: CalcRequest): CalcResponse {
    const { a, b, op } = payload;
    let result = 0;

    switch (op) {
      case '+':
        result = a + b;
        break;
      case '-':
        result = a - b;
        break;
      case '*':
        result = a * b;
        break;
      case '/':
        result = b !== 0 ? a / b : NaN;
        break;
    }

    this.lastResult = result;
    // reply only to the sender
    client.emit('result', { result });

    return { result };
  }
}
