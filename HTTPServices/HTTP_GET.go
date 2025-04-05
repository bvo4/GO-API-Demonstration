package HTTPServices

import (
	HTTPServices "API_DEMONSTRATION/HTTPServices/Models"
	"crypto/tls"
	"io"
	"net/http"
)

/* Credit: https://github.com/bradfitz/exp-httpclient/blob/master/problems.md */
func Write_HTTP_GET(OrderId string, API_Credentials *HTTPServices.HEADER_VALUES) {
	URI := BuildOrderIdURI(OrderId)

	HEADER_PARAMS := HTTPServices.FormatAPIKey(API_Credentials)
	ResponseBody := Send_HTTP_GET(URI, HEADER_PARAMS)

	//Decode the json response
}

/*
	Credit: https://blog.logrocket.com/making-http-requests-in-go/

// https://stackoverflow.com/questions/54401790/how-can-i-use-http-1-x-in-client-and-server
// https://stackoverflow.com/questions/57420833/tls-no-renegotiation-error-on-http-request
*/
func Send_HTTP_GET(URI string, HEADER_PARAMS map[string]string) string {
	//Get the Request URI using the OrderID
	r, err := http.NewRequest(http.MethodGet, URI, nil)
	if err != nil {
		panic(err)
	}

	/* HTTP CLIENT SETTINGS */
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Renegotiation:      tls.RenegotiateOnceAsClient,
			},
		},
	}

	AddHTTPHeaders(*r, HEADER_PARAMS)

	response, err := client.Do(r)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(response.Body)
	sb := string(body)
	//log.Printf(sb)

	return sb
}

/* Credit: https://stackoverflow.com/questions/8018719/iterating-through-a-golang-map */
func AddHTTPHeaders(r http.Request, HEADER_PARAMS map[string]string) {
	for key, value := range HEADER_PARAMS {
		r.Header.Add(key, value)
	}
}
