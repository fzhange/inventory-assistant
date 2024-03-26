"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.cookInventoryData = void 0;
function cookInventoryData(rawInventoryData) {
    const productNoIdxMap = {};
    let count = 0;
    const result = [];
    const tableHeadInfos = rawInventoryData[0].map((ele) => ele.trim());
    const productNoIdx = tableHeadInfos.indexOf('商品');
    const remainingQuantityIdx = tableHeadInfos.indexOf('库存数量');
    const productNameIdx = tableHeadInfos.indexOf('品名');
    rawInventoryData.forEach((element, idx) => {
        const productNo = element[productNoIdx];
        const remainingQuantity = element[remainingQuantityIdx];
        const productName = element[productNameIdx];
        const storeName = '可用库存';
        const mapIdx = productNoIdxMap[productNo];
        if (idx !== 0) {
            if (mapIdx !== 0 && !mapIdx) {
                productNoIdxMap[productNo] = count;
                count++;
                result.push({
                    storeName,
                    productNo,
                    remainingQuantity,
                    productName,
                });
            }
            else {
                const { storeName, productNo, remainingQuantity, productName } = result[mapIdx];
                result[mapIdx] = {
                    storeName,
                    productNo,
                    remainingQuantity: result[mapIdx].remainingQuantity + remainingQuantity,
                    productName,
                };
            }
        }
    });
    const productNoMap = {};
    result.forEach((ele) => {
        productNoMap[ele.productNo] = {
            ...ele,
        };
    });
    return productNoMap;
}
exports.cookInventoryData = cookInventoryData;
//# sourceMappingURL=inventory.helper.js.map