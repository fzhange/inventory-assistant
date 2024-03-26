import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getHello(): string {
    console.log('coming');
    return this.appService.getHello();
  }

  @Get('/getMatchInfo')
  getMatchInfo() {
    return this.appService.getMatchInfo();
  }
}

// 185.151.146.252