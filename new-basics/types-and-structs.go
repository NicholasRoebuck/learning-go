package main

import "log"
import "time"

var s string = "seven" // module variable. can be used within entire module

type User struct {
	FirstName	string
	LastName	string
	PhoneNumber	string
	Age			int 
	BirthDate	time.Time	
}

var special string // private bc of lowercase
var Special string // public bc of uppercase

func typesAndStructs() {
	// var s2 = "six"

	// log.Println("s is",s)
	// log.Println("s2 is",s2)

	// s3, empty := something("one")

	// log.Println(s3,"is", empty)

	user := User {
		FirstName: "Nicholas",
		LastName: "Roebuck",
		PhoneNumber: "1 444 444 4444",
	}
	log.Println(user.FirstName, user.LastName, user.Age, user.BirthDate)

}

// func scope: funcs starting with lowercase are private (only visible within the module they are created) and uppercase funcs are visible to all modules
func something (){

}
func Something (){ // public
}

// func something(s3 string) (string, string) {
// 	log.Println("something func s is", s)
// 	return s3, "empty"
// }