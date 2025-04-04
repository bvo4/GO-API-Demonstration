package main

import (
	"API_DEMONSTRATION/FileHandler"
	"fmt"
)

func main() {
	fmt.Println("Beginning Web API Demonstration")

	/* Get the API Credentials */
	//API_Credentials := FileHandler.GetConfig()

	/* Get Order Info from local .CSV file */
	CSV_RESULTS := FileHandler.ReadCSV()

	for i := range len(CSV_RESULTS) {
		fmt.Printf("OrderID: %s", CSV_RESULTS[i].Order_ID)
	}

	/* Extract Order IDs and use them to send an HTTP Query */

	/* Input Order Info into API */
	//Controller.GetItemsTreeOrderID()

}
