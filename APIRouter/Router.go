package APIRouter

import (
	"API_DEMONSTRATION/Controller"
	"API_DEMONSTRATION/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var HTTPAPI_CREDENTIALS Models.Credentials

// HTTP GET
// {localhost}/GetOrderDetails/{OrderId}
func SetRouterHeaders() {
	http.HandleFunc("/GetOrderDetails/{OrderId}", GetOrderDetails)
}

/* Turns on the HTTP Receiving Server:  Activates the list of HTTP Routes to listen for */
func InitiateRouter(API_CREDENTIALS Models.Credentials) {
	HTTPAPI_CREDENTIALS = API_CREDENTIALS //HTTP Credentials must be set globally
	SetRouterHeaders()                    //Set up list of HTTP Headers to listen to

	fmt.Println("HTTP Server turned on")
	http.ListenAndServe(":8090", nil) //Open on Port 8090
}

/* Credit: https://gobyexample.com/http-servers */
/* Upon receiving a valid HTTP Request, use the Order Id to make the API Query and return the results to the user */
func GetOrderDetails(w http.ResponseWriter, req *http.Request) {
	OrderId := req.PathValue("OrderId") //Extract the parameter from the GET Request
	//fmt.Fprintf(w, "Found: %s\n", OrderId)

	/* Upon receiving an Order ID, initiate the API to call the third-party API and return the results fetched as a JSON */
	sb := Controller.GetItemsTreeOrderID(HTTPAPI_CREDENTIALS, OrderId)
	JSONPARSE, _ := json.Marshal(sb) //Convert back to Json Payload
	ItemReturn := EncryptInformation(JSONPARSE)

	fmt.Fprintf(w, string(ItemReturn)) //Write into the HTTP Response Body the JSON Contents
}

func EncryptInformation(JSONPARSE []byte) []byte {
	var ItemReturn Models.ItemsTreeResult

	err := json.Unmarshal(JSONPARSE, &ItemReturn)
	Models.CheckError(err)

	MaskItemsTreeResult(ItemReturn)
	JSON_REPACKAGE, _ := json.Marshal(ItemReturn)
	return JSON_REPACKAGE
}

func MaskItemsTreeResult(ItemsTree Models.ItemsTreeResult) {

	ItemsTree.Containers = RecurseContainers(ItemsTree.Containers)

	/*
		if len(ItemsTree.Items) > 0 {
			for _, SelectedItem := range ItemsTree.Items {
				MaskItem(&SelectedItem)
			}
		}
	*/

	fmt.Println("Containers: ", len(ItemsTree.Containers))
}

func RecurseContainers(SelectedContainer []Models.Containers) []Models.Containers {
	if len(SelectedContainer) > 0 {
		for i, ChildContainer := range SelectedContainer {
			SelectedContainer[i] = MaskContainer(ChildContainer)
			RecurseContainers(ChildContainer.Containers)

			/*
				if len(ChildContainer.Items) > 0 {
					for _, ChildItem := range ChildContainer.Items {
						MaskItem(&ChildItem)
					}
				}
			*/
		}
	}
	return SelectedContainer
}

func MaskContainer(ChildContainer Models.Containers) Models.Containers {
	ChildContainer.SerialNumber = MaskText(ChildContainer.SerialNumber)
	ChildContainer.Number = MaskText(ChildContainer.Number)
	ChildContainer.LotNum = MaskText(ChildContainer.LotNum)
	ChildContainer.Expiry = MaskText(ChildContainer.Expiry)
	return ChildContainer
}

func MaskItem(ChildItem *Models.Items) {
	ChildItem.Gtin = MaskText(ChildItem.Gtin)
	ChildItem.LotNum = MaskText(ChildItem.LotNum)
	ChildItem.Expiry = MaskText(ChildItem.Expiry)
	ChildItem.Ndc = MaskText(ChildItem.Ndc)
	ChildItem.Upc = MaskText(ChildItem.Upc)
}

func MaskText(StringSet string) string {
	return strings.Repeat("*", len(StringSet))
}
