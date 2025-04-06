package HTTPModels

import "API_DEMONSTRATION/Models"

type HEADER_VALUES struct {
	Credentials   Models.Credentials
	HEADER_VALUES map[string]string
}

/* Converts the API Credentials into a more flexible format to later inject into the HTTP Headers */
func ToApiCredentials(API_Credentials Models.Credentials) HEADER_VALUES {
	var NewHeader HEADER_VALUES
	NewHeader.Credentials = API_Credentials
	return NewHeader
}
