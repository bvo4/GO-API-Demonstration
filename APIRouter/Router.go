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
	http.ListenAndServe(":8090", nil)     //On Port 8090
}

/* Credit: https://gobyexample.com/http-servers */
func GetOrderDetails(w http.ResponseWriter, req *http.Request) {
	OrderId := req.PathValue("OrderId")
	//fmt.Fprintf(w, "Found: %s\n", OrderId)

	/* Upon receiving an Order ID, initiate the API to call the third-party API and return the results fetched as a JSON */
	sb := Controller.GetItemsTreeOrderID(HTTPAPI_CREDENTIALS, OrderId)
	JSONPARSE, _ := json.Marshal(sb)
	fmt.Fprintf(w, string(JSONPARSE))
}
