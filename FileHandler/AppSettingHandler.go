package FileHandler

import (
	"API_DEMONSTRATION/Models"
	"encoding/json"
	"os"
)

func GetConfig() Models.Settings {

	//Get Current Directory Path
	DirectoryPath, err := os.Getwd()

	Models.CheckError(err)

	DirectoryPath += "\\appsetting.json"

	configFile, err := os.Open(DirectoryPath)
	defer configFile.Close()
	Models.CheckError(err)

	var API_Credentials Models.Settings
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&API_Credentials)
	return API_Credentials
}
