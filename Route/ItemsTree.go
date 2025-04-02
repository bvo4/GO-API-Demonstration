package Route

import (
	"API_DEMONSTRATION/HTTPServices"
	"API_DEMONSTRATION/Models"
)

func GetItemsTreeOrderID(OrderId string) Models.ItemsTreeResult {

	var Test Models.ItemsTreeResult
	Test.GTIN = "TEST"
	Test.LotNum = "TEST"
	Test.Qty = 2
	//Get the URI needed for the HTTP Request

	HTTPServices.Write_HTTP_GET(OrderId)

	return Test
}
