package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
	"github.com/xuri/excelize/v2"
)

type CookedInventoryData struct {
	StoreName         string `json:"storeName"`
	ProductNo         string `json:"productNo"`
	RemainingQuantity int    `json:"remainingQuantity"`
	ProductName       string `json:"productName"`
}

func CookInventory() {
	// read data from spreadsheet
	rawDataArr, _ := readInventorySpreadSheet()
	headRows := funk.Map((*rawDataArr)[0], func(ele string) string {
		return strings.Trim(ele, " ")
	})
	productNoIdx := funk.IndexOf(headRows, "商品")
	remainingQuantityIdx := funk.IndexOf(headRows, "可用量")
	productNameIdx := funk.IndexOf(headRows, "品名")

	// define variable
	// var result = make(map[int][]CookedInventoryData)
	// var productNoIdxMap = map[int]int{
	// 	0: -1,
	// }
	// var count = 0

	for idx, element := range *rawDataArr {
		if idx == 0 {
			continue
		}
		productNo := element[productNoIdx]
		remainingQuantity, _ := strconv.Atoi(element[remainingQuantityIdx])
		productName := element[productNameIdx]
		storeName := "可用库存"

		var cookedInventoryData = CookedInventoryData{
			ProductNo:         productNo,
			RemainingQuantity: remainingQuantity,
			StoreName:         storeName,
			ProductName:       productName,
		}
		fmt.Println("", cookedInventoryData)

	}
}
func readInventorySpreadSheet() (*[][]string, error) {
	wd, _ := os.Getwd()
	f, err := excelize.OpenFile(filepath.Join(wd, "./assets/inventory.xlsx"))

	if err == nil {
		rows, _ := f.GetRows(f.GetSheetName(0))
		return &rows, nil
	} else {
		return nil, err
	}
}
