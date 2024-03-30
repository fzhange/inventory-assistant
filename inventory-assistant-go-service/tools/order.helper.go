package tools

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/thoas/go-funk"
	"github.com/xuri/excelize/v2"
)

type CookedData struct {
	OrderNo       string `json:"orderNo"`
	CustomerName  string `json:"customerName"`
	CouldStoreNum string `json:"couldStoreNum"`
	BarCode       string `json:"barCode"`
	OrderTime     string `json:"orderTime"`
}

func CookOrderData() *map[int][]CookedData {
	// define variable
	var orderNoIdx, customerNameIdx, couldStoreNumIdx, barCodeIdx, orderTimeIdx int
	var result = make(map[int][]CookedData)
	var orderNoIdxMap = map[int]int{
		0: -1,
	}
	var count int = 0

	// step 1. read excel data
	rawDataArr, _ := readOrderSpreadSheet()

	// step2
	headRows := funk.Map((*rawDataArr)[0], func(x string) string {
		return strings.Trim(x, " ")
	})
	orderNoIdx = funk.IndexOf(headRows, "POS零售单号")
	customerNameIdx = funk.IndexOf(headRows, "会员卡号.顾客姓名")
	couldStoreNumIdx = funk.IndexOf(headRows, "云仓数量")
	barCodeIdx = funk.IndexOf(headRows, "条码")
	orderTimeIdx = funk.IndexOf(headRows, "单据日期")

	for rowIdx, element := range *rawDataArr {
		if rowIdx == 0 {
			continue
		} else {
			customerName := element[customerNameIdx]
			couldStoreNum := element[couldStoreNumIdx]
			barCode := element[barCodeIdx]
			orderTime := element[orderTimeIdx]
			orderNo := element[orderNoIdx]
			cookedData := CookedData{
				OrderNo:       orderNo,
				CustomerName:  customerName,
				CouldStoreNum: couldStoreNum,
				BarCode:       barCode,
				OrderTime:     orderTime,
			}

			if orderNoIdxMap[count] == 0 || orderNoIdxMap[count] == -1 {
				result[count] = []CookedData{cookedData}
			} else {
				result[count] = append(result[count], cookedData)
			}
			count = count + 1

		}
	}

	return &result
}

func readOrderSpreadSheet() (*[][]string, error) {
	wd, _ := os.Getwd()
	f, err := excelize.OpenFile(filepath.Join(wd, "./assets/order.xlsx"))

	if err == nil {
		rows, _ := f.GetRows(f.GetSheetName(0))
		return &rows, nil
	} else {
		return nil, err
	}
}
