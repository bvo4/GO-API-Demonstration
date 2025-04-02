package HTTPServices

import (
	"API_DEMONSTRATION/Models"
)

var ISDEBUG bool = false

func FormatAPIKey(API_Credentials Models.Credentials, HEADER_VALUES map[string]string) {

	KeyValues := make(map[string]string)

	//Copy the HTTP Credentials as part of the header values
	KeyValues[API_Credentials.API_KEY] = API_Credentials.API_VALUE

	//Loop through the entire Dictionary and add the header parameters to the HTTP Header
	for i, j := range HEADER_VALUES {
		KeyValues[j] = HEADER_VALUES[i]
	}
}
