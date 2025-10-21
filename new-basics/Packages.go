package main

import (
	"learning-go/helpers"
	"log"
)

func packages() {
	log.Println("Hello World")

	var myVar = helpers.NewType{
		TypeName:   "Name",
		TypeAction: "Action",
	}

	log.Println(myVar.TypeAction, myVar.TypeName)
}
