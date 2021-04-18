package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Email  string  `json:"email"`
	Phone  int     `json:"phone_number"`
	Parcel float32 `json:"parcel_weight"`
}

//Download uploaded file and return file name
func downloadFile(r *http.Request) string {
	// FormFile returns the first file for the given key `csv`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("csv")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	// Create a temporary file within our uploaded-file directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("uploaded-file", "upload-*.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	return tempFile.Name()
}

func procCSV(w http.ResponseWriter, r *http.Request) {
	fileName := downloadFile(r)

	fmt.Fprintln(w, fileName)
}

func HandleRequest() {
	http.HandleFunc("/api", procCSV)

	log.Fatal(http.ListenAndServe(":7000", nil))

}

func main() {
	HandleRequest()
}
