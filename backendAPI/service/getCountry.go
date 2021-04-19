package service

import "regexp"

//get country by phone number
func getCountry(phone string) string {
	//regexp.Match return bool and error. If phone number match regular expression then return true
	cameroon, _ := regexp.Match(`237[2368]\d{7,8}$`, []byte(phone))
	ethiopia, _ := regexp.Match(`251[1-59]\d{8}$`, []byte(phone))
	morocco, _ := regexp.Match(`212[5-9]\d{8}$`, []byte(phone))
	mozambique, _ := regexp.Match(`258[28]\d{7,8}$`, []byte(phone))
	uganda, _ := regexp.Match(`256\d{9}$`, []byte(phone))

	//case true return country value
	switch {
	case cameroon:
		return "Cameroon"
	case ethiopia:
		return "Ethiopia"
	case morocco:
		return "Morocco"
	case mozambique:
		return "Mozambique"
	case uganda:
		return "Uganda"
	}

	//empty if no match
	return ""
}
