package HTTPModels

import "API_DEMONSTRATION/Models"

type HEADER_VALUES struct {
	Credentials   Models.Credentials
	HEADER_VALUES map[string]string
}

func ToApiCredentials(API_Credentials Models.Credentials) HEADER_VALUES {
	var NewHeader HEADER_VALUES
	NewHeader.Credentials = API_Credentials
	return NewHeader
}
