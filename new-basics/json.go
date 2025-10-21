package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	LivesInJP bool   `json:"lives_in_jp"`
}

func workingWithJson() {
	myJson := `
	[
		{
			"first_name": "Nick",
			"last_name": "Roebuck",
			"hair_color": "Black",
			"lives_in_jp": true
		},
		{
			"first_name": "Jimmy",
			"last_name": "Jones",
			"hair_color": "Red",
			"lives_in_jp": true
		}
	]`

	//read json into a struct

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("Unmarshaled: %v", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	var m2 Person

	m1.FirstName = "Pickle"
	m1.LastName = "Rick"
	m1.HairColor = "none"
	m1.LivesInJP = false

	m2.FirstName = "Sarah"
	m2.LastName = "Near"
	m2.HairColor = "White"
	m2.LivesInJP = true

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("Error marshalling json", err)
	}

	fmt.Println(string(newJson))
}
