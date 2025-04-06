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
/* Opens the /File_Examples/ Directory and scan all .csv files */
func ReadCSV() []Models.OrderContents {
	DirectoryPath := GetDirectory() //Get base directory
	FilePaths := GetFileList(DirectoryPath)
	return ReadFile(FilePaths)
}

/* Credit: https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file */
func GetDirectory() string {
	DirectoryPath, err := os.Getwd()
	DirectoryPath += "\\File_Examples\\"
	CheckError(err)

	return DirectoryPath
}

/* https://stackoverflow.com/questions/14668850/list-directory-in-go */
/* Checks the given Directory Path and scans for every file found inside hte directory */
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
/* Goes through all .csv files and reads their contents, converting them into OrderContents array */
func ReadFile(FilePaths []string) []Models.OrderContents {

	var AllOrders []Models.OrderContents

	for i := range len(FilePaths) {
		//Opens the file
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

/* Goes through the file contents and converts them into OrderContents */
func ProcessFile(Records [][]string) Models.OrderContents {
	var Temp Models.OrderContents
	for i := range len(Records) {
		Temp = Models.ToOrderContents(Records, i)
	}
	return Temp
}
