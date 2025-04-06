package HTTPServices

import (
	HTTPServices "API_DEMONSTRATION/HTTPServices/Models"
	"API_DEMONSTRATION/Models"
	"io"
	"net/http"
)

/* Credit: https://github.com/bradfitz/exp-httpclient/blob/master/problems.md */
/* Writes an HTTP GET Request to the given API URI and returns the results */
func Write_HTTP_GET(OrderId string, API_Credentials *HTTPServices.HEADER_VALUES) Models.ItemsTreeResult {
	/* Fetch the Order URI. */
	URI := BuildOrderIdURI(OrderId) //NOTE:  THE URI FUNCTION IS DELIBERATELY HIDDEN FROM THE REPO

	/* Format the HTTP Header, including the API Key credentials */
	HEADER_PARAMS := HTTPServices.FormatAPIKey(API_Credentials)

	/* Send the HTTP GET Request */
	ResponseBody := Send_HTTP_GET(URI, HEADER_PARAMS)

	//Decode the json response
	Results := DecodeItemsTreeResponse(ResponseBody)
	return Results
}

/*
	Credit: https://blog.logrocket.com/making-http-requests-in-go/
// https://stackoverflow.com/questions/54401790/how-can-i-use-http-1-x-in-client-and-server
// https://stackoverflow.com/questions/57420833/tls-no-renegotiation-error-on-http-request
*/
/* Send an HTTP GET Request and returns the Json Body */
func Send_HTTP_GET(URI string, HEADER_PARAMS map[string]string) string {
	//Get the Request URI using the OrderID
	r, err := http.NewRequest(http.MethodGet, URI, nil)
	Models.CheckError(err)

	/* Format and send an HTTP GET Request to the API */
	response, err := SendHTTPMessage(r, HEADER_PARAMS)

	/* Convert the HTTP Response into a json string */
	body, err := io.ReadAll(response.Body)
	sb := string(body)
	//log.Printf(sb)

	return sb
}
