const XLSX = require('xlsx');
const {cookOrderData}  = require("./tools/order.helper")
const {cookInventoryData} = require("./tools/inventory.helper")
const fs = require("fs");

// 读取Excel文件
function readExcelFile(filePath) {
  const workbook = XLSX.readFile(filePath);

  // 获取第一个工作表
  const sheetName = workbook.SheetNames[0];
  const worksheet = workbook.Sheets[sheetName];

  // 将工作表转换为JSON对象数组
  const data = XLSX.utils.sheet_to_json(worksheet, { header: 1 /* 表头在第一行 */ });

  // 输出数据
  return data
}


const rawInventoryData = readExcelFile("./inventory.xlsx");
const rawOrderData = readExcelFile("./order.xlsx");





const cookedOrderData = cookOrderData(rawOrderData)
const cookedInventoryData = cookInventoryData(rawInventoryData)
// console.log('cookedInventoryData: ', cookedInventoryData);

// fs.writeFile("books.txt", `${JSON.stringify(cookedOrderData)}`, (err) => {
//   if (err)
//     console.log(err);
//   else {
//     console.log("File written successfully\n");
//     console.log("The written has the following contents:");
//     console.log(fs.readFileSync("books.txt", "utf8"));
//   }
// });

function matchOrderAndInventory(cookedOrderData,cookedInventoryData){
  const matchResult = {}
  cookedOrderData.forEach((orderEle)=>{
    // same order. different product
    orderEle.forEach(element => {
      const {orderNo,customerName,couldStoreNum,barCode} = element;
      const productInfo = cookedInventoryData[barCode];

      if(!matchResult[orderNo]) matchResult[orderNo] = []

      if(!productInfo){
        matchResult[orderNo].push({
          barCode,
          customerName,
          info: "商品不存在",
        })
      }else if(couldStoreNum - productInfo.remainingQuantity >=0){
        matchResult[orderNo].push({
          barCode,
          customerName,
          info: "商品数量充足",
        })
      }else if(couldStoreNum - productInfo.remainingQuantity < 0){
        matchResult[orderNo].push({
          barCode,
          customerName,
          info: `商品数量不够差${Math.abs(couldStoreNum - productInfo.remainingQuantity)}个,库存数量: ${productInfo.remainingQuantity} 下单数量：${couldStoreNum}`,
        })
      }
    });
  })
  console.log('matchResult: ', matchResult);
}

matchOrderAndInventory(cookedOrderData,cookInventoryData)
