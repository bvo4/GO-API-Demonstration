package Controller

import (
	"API_DEMONSTRATION/Models"
	"API_DEMONSTRATION/Route"
)

func GetItemsTreeGTIN(GTIN string, SerialNumber string) {

}

func GetItemsTreeOrderID(API_Credentials Models.Credentials, OrderID string) Models.ItemsTreeResult {
	return Route.GetItemsTreeOrderID(API_Credentials, OrderID)
}
