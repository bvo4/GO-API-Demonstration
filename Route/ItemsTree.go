package Route

import (
	"API_DEMONSTRATION/HTTPServices"
	HTTPModels "API_DEMONSTRATION/HTTPServices/Models"
	"API_DEMONSTRATION/Models"
	"fmt"
)

func GetItemsTreeOrderID(API_Credentials Models.Credentials, OrderId string) Models.ItemsTreeResult {

	//Get the URI needed for the HTTP Request
	HTTP_CREDENTIALS := HTTPModels.ToApiCredentials(API_Credentials)
	Results := HTTPServices.Write_HTTP_GET(OrderId, &HTTP_CREDENTIALS)

	DEBUG_EXPOUND_API_RESULTS(Results)

	return Results
}

func DEBUG_EXPOUND_API_RESULTS(Results Models.ItemsTreeResult) {
	fmt.Println("Item Len:", len(Results.Items))
	fmt.Println("Container Len:", len(Results.Containers))
}
