package main

import (
	"API_DEMONSTRATION/Controller"
	"API_DEMONSTRATION/FileHandler"
	"fmt"
)

func main() {
	fmt.Println("Beginning Web API Demonstration")

	/* Get the API Credentials */
	API_Credentials := FileHandler.GetConfig()

	/* Get Order Info */

	/* Input Order Info into API */
	Controller.GetItemsTreeOrderID()

}
