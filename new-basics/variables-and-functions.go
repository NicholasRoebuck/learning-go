package main

import "fmt"

func variablesAndFunctions() {
	fmt.Println("Hello, world.")

	var whatToSay string = "Goodbye cruel world"
	var i int = 100

	fmt.Println(whatToSay)
	fmt.Println("The variable i is assigned:", i)
	somethingSaid, otherSomethingSaid := saySomething()
	fmt.Println("the func returned", somethingSaid, otherSomethingSaid)
}

func saySomething () (string, string) {
	return "something said", "other something"
}