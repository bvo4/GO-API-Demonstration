package Route

import (
	"API_DEMONSTRATION/HTTPServices"
	HTTPModels "API_DEMONSTRATION/HTTPServices/Models"
	"API_DEMONSTRATION/Models"
)

/* Queries to the given API and returns the Item Set from the HTTP Response */
func GetItemsTreeOrderID(API_Credentials Models.Credentials, OrderId string) []Models.Items {

	//Get the URI needed for the HTTP Request
	HTTP_CREDENTIALS := HTTPModels.ToApiCredentials(API_Credentials)
	FullCase := HTTPServices.Write_HTTP_GET(OrderId, &HTTP_CREDENTIALS)

	DEBUG_EXPOUND_API_RESULTS(FullCase)
	ItemList := ExtractAllItems(FullCase)

	return ItemList
}
