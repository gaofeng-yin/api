package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"quickOps_exercise/backendAPI/model"
)

//user slice to store data
var user []model.User

//read csv file, append to user slice, turn to json and return slice of byte
func CsvToJson(filename string) []byte {
	//open csv file
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	//create new csv reader
	reader := csv.NewReader(csvFile)

	//loop to read
	for {
		col, err := reader.Read()
		//break cicle if hit end of file
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//append to user slice new User
		user = append(user, model.User{
			Email:  col[1],
			Phone:  col[2],
			Parcel: col[3],
			//country value is determined by phone number
			Country: getCountry(col[2]),
		})
	}

	//json.marshal return Json encoding of slice user
	userJson, _ := json.Marshal(user)
	return userJson
}
