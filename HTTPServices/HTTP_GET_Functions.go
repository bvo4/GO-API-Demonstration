package HTTPServices

import (
	"API_DEMONSTRATION/Models"
	"crypto/tls"
	"encoding/json"
	"net/http"
)

func LoadHTTP_ClientSettings() *http.Client {
	/* HTTP CLIENT SETTINGS */
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Renegotiation:      tls.RenegotiateOnceAsClient,
			},
		},
	}
	return client
}

/* Credit: https://stackoverflow.com/questions/8018719/iterating-through-a-golang-map */
func AddHTTPHeaders(r http.Request, HEADER_PARAMS map[string]string) {
	for key, value := range HEADER_PARAMS {
		r.Header.Add(key, value)
	}
}

func SendHTTPMessage(r *http.Request, HEADER_PARAMS map[string]string) (*http.Response, error) {
	/* HTTP CLIENT SETTINGS */
	client := LoadHTTP_ClientSettings()

	/* Add the header parameters and credentials */
	AddHTTPHeaders(*r, HEADER_PARAMS)

	/* Initiate the HTTP GET Request */
	response, err := client.Do(r)
	Models.CheckError(err)

	return response, err
}

/* https://stackoverflow.com/questions/47270595/how-to-parse-json-string-to-struct */
func DecodeItemsTreeResponse(JSONBODY string) Models.ItemsTreeResult {
	var API_Credentials Models.ItemsTreeResult

	json.Unmarshal([]byte(JSONBODY), &API_Credentials)
	return API_Credentials
}
