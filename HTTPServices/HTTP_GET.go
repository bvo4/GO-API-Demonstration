package HTTPServices

import (
	HTTPServices "API_DEMONSTRATION/HTTPServices/Models"
)

func Write_HTTP_GET(OrderId string, API_Credentials *HTTPServices.HEADER_VALUES) {
	URI := BuildOrderIdURI(OrderId)
	HEADER_PARAMS := HTTPServices.FormatAPIKey(API_Credentials, HEADER_VALUES)
	Send_HTTP_GET(URI, HEADER_PARAMS)
}

func Send_HTTP_GET(URI string, HEADER_PARAMS map[string]string) {
	//Get the Request URI using the OrderID

}
