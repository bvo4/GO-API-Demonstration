package HTTPModels

var ISDEBUG bool = false

func FormatAPIKey(API_Credentials *HEADER_VALUES) map[string]string {

	KeyValues := make(map[string]string)

	//Copy the HTTP Credentials as part of the header values
	KeyValues[API_Credentials.Credentials.API_KEY] = API_Credentials.Credentials.API_VALUE

	//If API_CREDENTIALS has any additional header parameters, add them to the map
	if &API_Credentials.HEADER_VALUES != nil {
		for i, j := range API_Credentials.HEADER_VALUES {
			KeyValues[j] = API_Credentials.HEADER_VALUES[i]
		}
	}
	return KeyValues
}
