package Route

import (
	"API_DEMONSTRATION/HTTPServices"
	HTTPModels "API_DEMONSTRATION/HTTPServices/Models"
	"API_DEMONSTRATION/Models"
	"fmt"
)

func GetItemsTreeOrderID(API_Credentials Models.Credentials, OrderId string) []Models.Items {

	//Get the URI needed for the HTTP Request
	HTTP_CREDENTIALS := HTTPModels.ToApiCredentials(API_Credentials)
	FullCase := HTTPServices.Write_HTTP_GET(OrderId, &HTTP_CREDENTIALS)

	DEBUG_EXPOUND_API_RESULTS(FullCase)
	ItemList := ExtractAllItems(FullCase)

	return ItemList
}

/* Recursively goes through every container in ItemsTreeResult and extracts all items */
func ExtractAllItems(ItemsTree Models.ItemsTreeResult) []Models.Items {

	if len(ItemsTree.Containers) < 1 {
		return ItemsTree.Items
	}
	var ItemList []Models.Items
	for _, ItemSet := range ItemsTree.Containers {
		ItemList = append(ItemList, ExtractFromContainer(ItemSet)...)
	}
	return ItemList
}

func ExtractFromContainer(ContainerList Models.Containers) []Models.Items {
	if len(ContainerList.Containers) < 1 {
		return ContainerList.Items
	}

	var ItemSet []Models.Items
	for _, SelectedContainer := range ContainerList.Containers {
		Temp := ExtractFromContainer(SelectedContainer) //Go through every container first then extract the items
		ItemSet = append(ItemSet, Temp...)              //Append the Items array
	}
	return ItemSet
}

func DEBUG_EXPOUND_API_RESULTS(Results Models.ItemsTreeResult) {
	fmt.Println("Item Len:", len(Results.Items))
	fmt.Println("Container Len:", len(Results.Containers))
}
