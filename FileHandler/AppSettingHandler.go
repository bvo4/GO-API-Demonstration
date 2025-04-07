package FileHandler

import (
	"API_DEMONSTRATION/Models"
	"encoding/json"
	"os"
)

/* Load AppSetting.json to get the program settings */
func GetConfig() Models.Settings {

	//Get Current Directory
	DirectoryPath := GetAppSettingFile()

	/* Open AppSetting.Json */
	configFile, err := os.Open(DirectoryPath)
	defer configFile.Close()
	Models.CheckError(err)

	/* Decode the appsetting .json contents and convert them into Model.Settings */
	return DecodeAppSettings(configFile)
}

/* Get the AppSetting.json directory */
func GetAppSettingFile() string {
	//Get Current Directory
	DirectoryPath, err := os.Getwd()

	Models.CheckError(err)

	//Check for the appsetting.json file to load program settings
	DirectoryPath += "//appsetting.json"
	return DirectoryPath
}

/* Parse AppSetting.json contents */
func DecodeAppSettings(configFile *os.File) Models.Settings {
	/* Decode the appsetting .json contents and convert them into Model.Settings */
	var API_Credentials Models.Settings
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&API_Credentials)
	return API_Credentials
}
