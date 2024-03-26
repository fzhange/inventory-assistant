import { Injectable } from '@nestjs/common';
import { join } from 'path';

import {
  cookOrderData,
  cookInventoryData,
  readExcelFile,
  matchOrderAndInventory,
} from './tools/index.helper';

@Injectable()
export class AppService {
  getHello(): string {
    return 'Hello World!';
  }
  getMatchInfo() {
    const rawInventoryData = readExcelFile(
      join(__dirname, '../src/assets/inventory.xlsx'),
    );

    const rawOrderData = readExcelFile(
      join(__dirname, '../src/assets/order.xlsx'),
    );

    const cookedOrderData = cookOrderData(rawOrderData);
    const cookedInventoryData = cookInventoryData(rawInventoryData);

    return matchOrderAndInventory(cookedOrderData, cookedInventoryData);
  }
}
