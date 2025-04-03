package FileHandler

import (
	"encoding/csv"
	"os"
)

func CheckError(Exception error) {
	if Exception != nill {
		panic(Exception)
	}
}

func GetDirectory() string {
	DirectoryPath, err := os.UserHomeDir()
	DirectoryPath += "\\Documents\\Go API Demonstration\\GO-API-Demonstration\\File_Examples\\"

	Data, err := os.ReadFile(DirectoryPath)
	CheckError(err)

	return DirectoryPath
}

/* Credit:  https://stackoverflow.com/questions/24999079/reading-csv-file-in-go */
func ReadFile() [][]string {
	DirectoryPath := GetDirectory()

	FileContents, err := os.Open(DirectoryPath)
	CheckError(err)

	CSV_READER := csv.NewReader(FileContents)

	records, err := CSV_READER.ReadAll()
	CheckError(err)

	return records
}
