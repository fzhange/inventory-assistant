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
  });

  // 输出数据
  return data;
}

export function matchOrderAndInventory(cookedOrderData, cookedInventoryData) {
  const matchResult = {};
  cookedOrderData.forEach((orderEle) => {
    // same order. different product
    orderEle.forEach((element) => {
      const { orderNo, customerName, couldStoreNum, barCode } = element;
      const productInfo = cookedInventoryData[barCode];

      if (!matchResult[orderNo]) matchResult[orderNo] = [];

      if (!productInfo) {
        matchResult[orderNo].push({
          barCode,
          customerName,
          info: '商品不存在',
        });
      } else if (productInfo.remainingQuantity - couldStoreNum >= 0) {
        matchResult[orderNo].push({
          barCode,
          customerName,
          info: `商品数量充足,库存数量: ${productInfo.remainingQuantity} 下单数量：${couldStoreNum}`,
        });
      } else if (productInfo.remainingQuantity - couldStoreNum < 0) {
        matchResult[orderNo].push({
          barCode,
          customerName,
          info: `商品数量不够差${Math.abs(couldStoreNum - productInfo.remainingQuantity)}个,库存数量: ${productInfo.remainingQuantity} 下单数量：${couldStoreNum}`,
        });
      }
    });
  });

  return matchResult;
}

export * from './inventory.helper';
export * from './order.helper';
