package Route

import (
	"API_DEMONSTRATION/HTTPServices"
	HTTPModels "API_DEMONSTRATION/HTTPServices/Models"
	"API_DEMONSTRATION/Models"
)

func GetItemsTreeOrderID(API_Credentials Models.Credentials, OrderId string) Models.ItemsTreeResult {

	var Test Models.ItemsTreeResult
	Test.GTIN = "TEST"
	Test.LotNum = "TEST"
	Test.Qty = 2

	//Get the URI needed for the HTTP Request
	HTTP_CREDENTIALS := HTTPModels.ToApiCredentials(API_Credentials)
	HTTPServices.Write_HTTP_GET(OrderId, &HTTP_CREDENTIALS)

	return Test
}
