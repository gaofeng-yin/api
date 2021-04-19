package controller

import (
	"fmt"
	"log"
	"net/http"
	"quickOps_exercise/backendAPI/service"
)

func procCSV(w http.ResponseWriter, r *http.Request) {
	//download uploaded file
	fileName, err := service.DownloadFile(r)
	if err != nil {
		log.Panic(err)
	}

	//open file and convert it to json
	slice := service.CsvToJson(fileName)

	//convert []byte to string and write to the page
	fmt.Fprintln(w, string(slice))

	//inset into MongoDB, NoSQL database ?
	//...
}

func HandleRequest() {
	http.HandleFunc("/api", procCSV)

	log.Fatal(http.ListenAndServe(":7000", nil))
}
