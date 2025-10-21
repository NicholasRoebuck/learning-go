package main

import (
	"learning-go/helpers"
	"log"
)

const numPool = 100

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}
func channels() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan

	log.Println(num)
}
