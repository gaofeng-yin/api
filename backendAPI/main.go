package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Email  string  `csv:"email"`
	Phone  int     `csv:"phone_number"`
	Parcel float32 `csv:"parcel_weight"`
}

func procCSV(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Processing CSV file")
	fmt.Println("Done")
}

func handleRequest() {
	http.HandleFunc("/api", procCSV)

	log.Fatal(http.ListenAndServe(":7000", nil))

}

func main() {
	handleRequest()
}
