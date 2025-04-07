package main

import (
	"API_DEMONSTRATION/APIRouter"
	"API_DEMONSTRATION/Controller"
	"API_DEMONSTRATION/FileHandler"
	"API_DEMONSTRATION/MasterService"
	"API_DEMONSTRATION/Models"
	"fmt"
)

func main() {
	fmt.Println("Beginning Web API Demonstration")

	/* Get the API Credentials */
	API_Credentials := FileHandler.GetConfig()

	/* Get Order Info from local .CSV file */
	CSV_RESULTS := FileHandler.ReadCSV()

	//DEBUG:  PRINT OUT ORDERS
	PrintOrderID(CSV_RESULTS)

	/* Input Order Info into API */
	SEND_API_ORDERS(API_Credentials, CSV_RESULTS)

	/* Turn on HTTP Listener and act like an HTTP Server */
	APIRouter.InitiateRouter(API_Credentials.Credentials)
}

/* Debug:  Prints out the list of Order Ids acquired */
func PrintOrderID(CSV_RESULTS []Models.OrderContents) {
	for i := range len(CSV_RESULTS) {
		fmt.Printf("OrderID: %s \n", CSV_RESULTS[i].Order_ID)
	}
}

/* Takes the OrderContents array and inserts all items found into the SQL Database */
func SEND_API_ORDERS(API_CREDENTIALS Models.Settings, CSV_RESULTS []Models.OrderContents) {
	for i := range len(CSV_RESULTS) {
		EpcisDtl := Controller.GetItemsTreeOrderID(API_CREDENTIALS.Credentials, CSV_RESULTS[i].Order_ID)
		MasterService.InsertSSCC(API_CREDENTIALS, EpcisDtl)
	}
}
