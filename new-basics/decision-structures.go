package main

import "log"

func decisionStructures() {
	myVar := "ca"
	switch myVar {
	case "cat":
		log.Println("Cat is a cat")

	case "dog":
		log.Println("Dog is a dog")

	default:
		log.Println("I don't know what to do")

	}

}
