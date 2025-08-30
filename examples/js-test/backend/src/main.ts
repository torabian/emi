import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.enableCors({
    origin: 'http://localhost:5173', // frontend origin
    credentials: true, // allow cookies / credentials
  });
  await app.listen(process.env.PORT ?? 3000);
}
bootstrap();
