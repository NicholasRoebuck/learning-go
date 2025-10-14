package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	const MAXENTRIES = 10

	const employeeID int = 1001

	fmt.Println("Emploee ID:", employeeID)
}
