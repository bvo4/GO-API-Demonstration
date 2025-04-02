package Controller

import (
	"API_DEMONSTRATION/Models"
	"API_DEMONSTRATION/Route"
)

func GetItemsTreeGTIN(GTIN string, SerialNumber string) {

}

func GetItemsTreeOrderID(OrderID string) Models.ItemsTreeResult {
	return Route.GetItemsTreeOrderID(OrderID)
}
