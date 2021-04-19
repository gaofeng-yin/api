package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Download uploaded file and return file name and error
func DownloadFile(r *http.Request) (string, error) {
	// FormFile returns the first file for the given key `csv`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("csv")
	if err != nil {
		return "", err
	}
	//close file later in program
	defer file.Close()

	// Create a temporary file within our uploaded-file directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("uploaded-file", "upload-*.csv")
	if err != nil {
		fmt.Println(err)
	}
	//close file later in program
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return file name
	return tempFile.Name(), nil
}
