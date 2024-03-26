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
    console.log('__dirname: ', __dirname);

    const rootPath =
      '/Users/apple/fzhange/inventory-assistant/inventory-assistant-service';

    const rawInventoryData = readExcelFile(
      join(rootPath, 'src/assets/inventory.xlsx'),
    );

    const rawOrderData = readExcelFile(join(rootPath, 'src/assets/order.xlsx'));

    const cookedOrderData = cookOrderData(rawOrderData);
    const cookedInventoryData = cookInventoryData(rawInventoryData);

    return matchOrderAndInventory(cookedOrderData, cookedInventoryData);
  }
}