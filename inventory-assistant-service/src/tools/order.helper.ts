import * as moment from 'moment';

export function cookOrderData(rawOrderData) {
  // cache variable.
  const orderNoIdxMap = {};
  let count = 0;
  const result = [];
  // step1.
  const tableHeadInfos = rawOrderData[0].map((ele) => ele.trim());

  const orderNoIdx = tableHeadInfos.indexOf('POS零售单号');
  const customerNameIdx = tableHeadInfos.indexOf('会员卡号.顾客姓名');
  const couldStoreNumIdx = tableHeadInfos.indexOf('云仓数量');
  const barCodeIdx = tableHeadInfos.indexOf('条码');
  const orderTimeIdx = tableHeadInfos.indexOf('单据日期');

  // step2.
  rawOrderData.forEach((element, idx) => {
    const orderNo = element[orderNoIdx];
    const customerName = element[customerNameIdx];
    const couldStoreNum = element[couldStoreNumIdx];
    const barCode = element[barCodeIdx];
    const orderTime = element[orderTimeIdx];

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
            orderTimeMomentType: moment(orderTime, 'YYYY-MM-DD'),
            orderTime,
          },
        ]);
      } else {
        result[mapIdx].push({
          orderNo,
          customerName,
          couldStoreNum,
          barCode,
          orderTimeMomentType: moment(orderTime, 'YYYY-MM-DD'),
          orderTime,
        });
      }
    }
  });

  result.sort((a, b) => {
    if (!a?.orderTimeMomentType) return -1;
    if (!b?.b.orderTimeMomentType) return 1;
    return a.orderTimeMomentType.isBefore(b.orderTimeMomentType)
      ? -1
      : a.isAfter(b)
        ? 1
        : 0;
  });

  return result;
}
