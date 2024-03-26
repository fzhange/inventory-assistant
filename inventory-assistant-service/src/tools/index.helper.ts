import { readFile, utils } from 'xlsx';

// 读取Excel文件
export function readExcelFile(filePath) {
  const workbook = readFile(filePath);

  // 获取第一个工作表
  const sheetName = workbook.SheetNames[0];
  const worksheet = workbook.Sheets[sheetName];

  // 将工作表转换为JSON对象数组
  const data = utils.sheet_to_json(worksheet, {
    header: 1 /* 表头在第一行 */,
    raw: false,
  });

  // 输出数据
  return data;
}

export function matchOrderAndInventory(cookedOrderData, cookedInventoryData) {
  const matchResult = {};
  cookedOrderData.forEach((orderEle) => {
    // same order. different product
    orderEle.forEach((element) => {
      const { orderNo, customerName, couldStoreNum, barCode, orderTime } =
        element;
      const productInfo = cookedInventoryData[barCode];

      if (!matchResult[orderNo]) matchResult[orderNo] = [];

      if (!productInfo) {
        matchResult[orderNo].push({
          barCode,
          orderTime,
          customerName,
          info: `商品不存在,下单数量：${couldStoreNum}`,
          state: 'danger',
        });
      } else if (productInfo.remainingQuantity - couldStoreNum >= 0) {
        matchResult[orderNo].push({
          barCode,
          orderTime,
          customerName,
          info: `商品数量充足,下单数量：${couldStoreNum} 库存数量: ${productInfo.remainingQuantity}`,
          state: 'success',
        });
      } else if (productInfo.remainingQuantity - couldStoreNum < 0) {
        matchResult[orderNo].push({
          barCode,
          orderTime,
          customerName,
          info: `商品数量不够差${Math.abs(couldStoreNum - productInfo.remainingQuantity)}个,下单数量：${couldStoreNum} 库存数量: ${productInfo.remainingQuantity}`,
          state: 'danger',
        });
      }
    });
  });

  return matchResult;
}

export * from './inventory.helper';
export * from './order.helper';
