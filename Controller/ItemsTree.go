package Controller

import (
	"API_DEMONSTRATION/Models"
	"API_DEMONSTRATION/Route"
)

func GetItemsTreeGTIN(GTIN string, SerialNumber string) {

}

/* Query the API and get the Order Items */
func GetItemsTreeOrderID(API_Credentials Models.Credentials, OrderID string) []Models.Items {
	return Route.GetItemsTreeOrderID(API_Credentials, OrderID)
}
