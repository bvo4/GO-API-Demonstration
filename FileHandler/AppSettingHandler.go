package FileHandler

import (
	"API_DEMONSTRATION/Models"
	"encoding/json"
	"log"
	"os"
)

func GetConfig() Models.Settings {

	DirectoryPath, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	DirectoryPath += "\\Documents\\Go API Demonstration\\GO-API-Demonstration\\appsetting.json"

	configFile, err := os.Open(DirectoryPath)
	defer configFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	var API_Credentials Models.Settings
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&API_Credentials)
	return API_Credentials
}
