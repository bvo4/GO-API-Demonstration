package main

import (
	"API_DEMONSTRATION/Controller"
	"API_DEMONSTRATION/Models"
	"encoding/json"
	"fmt"
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

func main() {
	fmt.Println("Beginning Web API Demonstration")

	/* Get the API Credentials */
	API_Credentials := GetConfig()

	/* Get Order Info */

	/* Input Order Info into API */
	Controller.GetItemsTreeOrderID()

}
