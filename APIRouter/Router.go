package APIRouter

import (
	"API_DEMONSTRATION/Controller"
	"API_DEMONSTRATION/Models"
	"encoding/json"
	"fmt"
	"net/http"
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

	/* Upon receiving an Order ID, initiate the API to call the third-party API and return the results fetched as a JSON */
	sb := Controller.GetItemsTreeOrderID(HTTPAPI_CREDENTIALS, OrderId)

	/* Encrypt and Mask the information before sendin ganything out */
	ItemReturn := EncryptInformation(sb)

	fmt.Fprintf(w, string(ItemReturn)) //Write into the HTTP Response Body the JSON Contents
}

/* Opens the API Results, goes through every Item and Container and masks the information presented */
func EncryptInformation(ItemReturn Models.ItemsTreeResult) []byte {

	/* Applies a mask to the entire data set */
	Models.MaskItemsTreeResult(&ItemReturn)
	JSON_REPACKAGE, _ := json.Marshal(ItemReturn)
	return JSON_REPACKAGE
}
