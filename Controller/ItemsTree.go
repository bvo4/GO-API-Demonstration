package Controller

import (
	"API_DEMONSTRATION/Models"
	"API_DEMONSTRATION/Route"
)

func GetItemsTreeOrderID(API_Credentials Models.Credentials, OrderID string) Models.ItemsTreeResult {
	return Route.GetItemsTreeOrderID(API_Credentials, OrderID)
}

/* Query the API and get the Order Items */
func GetItemsOrderID(API_Credentials Models.Credentials, OrderID string) []Models.Items {
	return Route.GetItemsOrderID(API_Credentials, OrderID)
}
