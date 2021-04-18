package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type User struct {
	Email  string `json:"email"`
	Phone  string `json:"phone_number"`
	Parcel string `json:"parcel_weight"`
}

//Download uploaded file and return file name
func downloadFile(r *http.Request) (string, error) {
	// FormFile returns the first file for the given key `csv`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("csv")
	if err != nil {
		return "", err
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
	return tempFile.Name(), nil
}

func procCSV(w http.ResponseWriter, r *http.Request) {
	fileName, err := downloadFile(r)
	if err != nil {
		log.Panic(err)
	}

	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(csvFile)

	var user []User
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		user = append(user, User{
			Email:  line[1],
			Phone:  line[2],
			Parcel: line[3],
		})
	}
	userJson, _ := json.Marshal(user)
	fmt.Fprintln(w, string(userJson))
}

func HandleRequest() {
	http.HandleFunc("/api", procCSV)

	log.Fatal(http.ListenAndServe(":7000", nil))

}

func main() {
	HandleRequest()
}
