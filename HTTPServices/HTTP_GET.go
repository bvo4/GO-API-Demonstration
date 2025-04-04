package HTTPServices

import (
	HTTPServices "API_DEMONSTRATION/HTTPServices/Models"
	"fmt"
	"net/http"
)

/* Credit: https://github.com/bradfitz/exp-httpclient/blob/master/problems.md */
func Write_HTTP_GET(OrderId string, API_Credentials *HTTPServices.HEADER_VALUES) {
	URI := BuildOrderIdURI(OrderId)

	HEADER_PARAMS := HTTPServices.FormatAPIKey(API_Credentials)
	Send_HTTP_GET(URI, HEADER_PARAMS)
}

func Send_HTTP_GET(URI string, HEADER_PARAMS map[string]string) {
	//Get the Request URI using the OrderID

	r, err := http.NewRequest("GET", URI, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(r)

}
