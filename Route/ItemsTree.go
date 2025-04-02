package Route

import (
	"API_DEMONSTRATION/Models"
)

func GetItemsTreeOrderID(GTIN string) Models.ItemsTreeResult {

	var Test Models.ItemsTreeResult

	Test.GTIN = "TEST"
	Test.LotNum = "TEST"
	Test.Qty = 2

	println("Placeholder")

	return Test
}
