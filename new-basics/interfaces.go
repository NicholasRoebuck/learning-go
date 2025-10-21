package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Cat struct {
	Name  string
	Age   int
	Owner string
}

func interfaces() {
	dog := Dog{
		Name:  "Jimmy",
		Breed: "German Shepard",
	}

	PrintInfo(&dog)

	cat := Cat{
		Name:  "Smithy",
		Age:   4,
		Owner: "Himothy",
	}

	PrintInfo(&cat)

}
func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (c *Cat) Says() string {
	return "Meow"
}

func (c *Cat) NumberOfLegs() int {
	return 4
}
