const XLSX = require('xlsx');
const {cookOrderData}  = require("./tools/order.helper")
const {cookInventoryData} = require("./tools/inventory.helper")
const fs = require("fs");










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



