package main

import (
	"API_DEMONSTRATION/Controller"
	"API_DEMONSTRATION/FileHandler"
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
	SEND_API_ORDERS(API_Credentials.Credentials, CSV_RESULTS)
}

func PrintOrderID(CSV_RESULTS []Models.OrderContents) {
	for i := range len(CSV_RESULTS) {
		fmt.Printf("OrderID: %s", CSV_RESULTS[i].Order_ID)
	}
}

func SEND_API_ORDERS(API_CREDENTIALS Models.Credentials, CSV_RESULTS []Models.OrderContents) {
	for i := range len(CSV_RESULTS) {
		Controller.GetItemsTreeOrderID(API_CREDENTIALS, CSV_RESULTS[i].Order_ID)
	}
}
