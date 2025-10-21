package main

import "log"

type Person1 struct {
	FirstName string
	LastName  string
}

func maps() {
	myMap := make(map[string]string)
	myMap["dog"] = "sammy"
	myMap["cat"] = "becky"

	log.Println(myMap["dog"], myMap["cat"])

	otherMap := make(map[string]int)
	otherMap["penny"] = 1
	otherMap["nickel"] = 5
	otherMap["dime"] = 10

	log.Println(otherMap["penny"], otherMap["dime"])

	userMap := make(map[string]Person1)

	me := Person1{
		FirstName: "nick",
		LastName:  "roebuck",
	}

	userMap["me"] = me

	log.Println(userMap["me"].FirstName, userMap["me"].LastName)
}
