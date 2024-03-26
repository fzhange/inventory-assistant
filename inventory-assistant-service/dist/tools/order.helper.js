"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.cookOrderData = void 0;
function cookOrderData(rawOrderData) {
    const orderNoIdxMap = {};
    let count = 0;
    const result = [];
    const tableHeadInfos = rawOrderData[0].map((ele) => ele.trim());
    const orderNoIdx = tableHeadInfos.indexOf('POS零售单号');
    const customerNameIdx = tableHeadInfos.indexOf('会员卡号.顾客姓名');
    const couldStoreNumIdx = tableHeadInfos.indexOf('云仓数量');
    const barCodeIdx = tableHeadInfos.indexOf('条码');
    rawOrderData.forEach((element, idx) => {
        const orderNo = element[orderNoIdx];
        const customerName = element[customerNameIdx];
        const couldStoreNum = element[couldStoreNumIdx];
        const barCode = element[barCodeIdx];
        const mapIdx = orderNoIdxMap[orderNo];
        if (idx !== 0) {
            if (mapIdx !== 0 && !mapIdx) {
                orderNoIdxMap[orderNo] = count;
                count++;
                result.push([
                    {
                        orderNo,
                        customerName,
                        couldStoreNum,
                        barCode,
                    },
                ]);
            }
            else {
                result[mapIdx].push({
                    orderNo,
                    customerName,
                    couldStoreNum,
                    barCode,
                });
            }
        }
    });
    return result;
}
exports.cookOrderData = cookOrderData;
//# sourceMappingURL=order.helper.js.map