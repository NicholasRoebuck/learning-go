package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string { // this function is a receiver of myStruct
	return m.FirstName
}

func receiverFunctionWithStructs() {
	var myVar myStruct
	myVar.FirstName = "Nick"

	myVar2 := myStruct{
		FirstName: "john",
	}

	log.Println("myVar is", myVar.printFirstName())
	log.Println("myVar is", myVar2.printFirstName())

}