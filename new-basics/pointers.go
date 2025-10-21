package main

import "log"

func pointers() {
	var myCoin string
	myCoin = "Penny"
	log.Println("myCoin is set to",myCoin)
	changeUsingPointer(&myCoin)
	log.Println("After calling changeUsingPointer func, myCoin is set to", myCoin)
}

func changeUsingPointer (s *string){
	// *s = "Dime" //this first way works as well
	newValue := "Dime"
	*s = newValue
}