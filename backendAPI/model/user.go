package model

//User structure
type User struct {
	Email   string `json:"email"`
	Phone   string `json:"phone_number"`
	Parcel  string `json:"parcel_weight"`
	Country string `json:"country"`
}
