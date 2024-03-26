"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.matchOrderAndInventory = exports.readExcelFile = void 0;
const xlsx_1 = require("xlsx");
function readExcelFile(filePath) {
    const workbook = (0, xlsx_1.readFile)(filePath);
    const sheetName = workbook.SheetNames[0];
    const worksheet = workbook.Sheets[sheetName];
    const data = xlsx_1.utils.sheet_to_json(worksheet, {
        header: 1,
    });
    return data;
}
exports.readExcelFile = readExcelFile;
function matchOrderAndInventory(cookedOrderData, cookedInventoryData) {
    const matchResult = {};
    cookedOrderData.forEach((orderEle) => {
        orderEle.forEach((element) => {
            const { orderNo, customerName, couldStoreNum, barCode } = element;
            const productInfo = cookedInventoryData[barCode];
            if (!matchResult[orderNo])
                matchResult[orderNo] = [];
            if (!productInfo) {
                matchResult[orderNo].push({
                    barCode,
                    customerName,
                    info: `商品不存在,下单数量：${couldStoreNum}`,
                    state: 'danger',
                });
            }
            else if (productInfo.remainingQuantity - couldStoreNum >= 0) {
                matchResult[orderNo].push({
                    barCode,
                    customerName,
                    info: `商品数量充足,下单数量：${couldStoreNum} 库存数量: ${productInfo.remainingQuantity}`,
                    state: 'success',
                });
            }
            else if (productInfo.remainingQuantity - couldStoreNum < 0) {
                matchResult[orderNo].push({
                    barCode,
                    customerName,
                    info: `商品数量不够差${Math.abs(couldStoreNum - productInfo.remainingQuantity)}个,下单数量：${couldStoreNum} 库存数量: ${productInfo.remainingQuantity}`,
                    state: 'danger',
                });
            }
        });
    });
    return matchResult;
}
exports.matchOrderAndInventory = matchOrderAndInventory;
__exportStar(require("./inventory.helper"), exports);
__exportStar(require("./order.helper"), exports);
//# sourceMappingURL=index.helper.js.map