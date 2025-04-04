package FileHandler

import (
	"API_DEMONSTRATION/Models"
	"encoding/csv"
	"os"
	"path/filepath"
)

func CheckError(Exception error) {
	if Exception != nil {
		panic(Exception)
		os.Exit(3)
	}
}

/* Credit: https://pkg.go.dev/os#ReadDir */
func ReadCSV() []Models.OrderContents {
	DirectoryPath := GetDirectory() //Get base directory
	FilePaths := GetFileList(DirectoryPath)
	return ReadFile(FilePaths)
}

func GetDirectory() string {
	DirectoryPath, err := os.Getwd()
	DirectoryPath += "\\File_Examples\\"
	CheckError(err)

	return DirectoryPath
}

/* https://stackoverflow.com/questions/14668850/list-directory-in-go */
func GetFileList(DirectoryPath string) []string {
	_, err := os.ReadDir(DirectoryPath)
	CheckError(err)

	/* Go through every .csv file and extract the data */
	var AllFiles []string
	err = filepath.Walk(DirectoryPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			AllFiles = append(AllFiles, path)
		}
		return nil
	})
	return AllFiles
}

/* Credit:  https://stackoverflow.com/questions/24999079/reading-csv-file-in-go */
func ReadFile(FilePaths []string) []Models.OrderContents {

	var AllOrders []Models.OrderContents

	for i := range len(FilePaths) {
		FileContents, err := os.Open(FilePaths[i])
		CheckError(err)

		CSV_READER := csv.NewReader(FileContents)

		//Read the contents of the .csv file and save it
		records, err := CSV_READER.ReadAll()
		CheckError(err)

		//Convert the file into OrderContents and return it
		Temp := ProcessFile(records)
		AllOrders = append(AllOrders, Temp)
	}

	return AllOrders
}

/* Goes through the file contents and converts them into  */
func ProcessFile(Records [][]string) Models.OrderContents {
	var Temp Models.OrderContents
	for i := range len(Records) {
		Temp = Models.ToOrderContents(Records, i)
	}
	return Temp
}
