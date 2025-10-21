package main

import "log"

func loopsAndRanging() {
	type User struct {
		FirstName   string
		LastName    string
		PhoneNumber int
		Email       string
	}

	//for i := 0; i <= 10; i++ {
	//	log.Println(i)
	//}

	//animals := []string{"dog", "cat", "fish", "horse", "cow"}
	//
	//for _, animal := range animals { // "_" is the blank identifier. meaning you don't want you use the index produced by the for loop
	//	log.Println(animal)
	//}
	//
	//for i, a := range animals {
	//	log.Println(i, a)
	//}

	animals := make(map[string]string)
	animals["dog"] = "Jimmy"
	animals["cat"] = "Timmy"
	animals["cow"] = "Himmy"
	//
	//for animalType, animal := range animals {
	//	log.Println(animalType, animal)
	//}
	//firsLine := "The hope and dream of your forefathers"
	//
	//for i, l := range firsLine {
	//	log.Println(i, ":", string(l))
	//}

	var users []User

	users = append(users, User{"Nick", "Roebuck", 9995553888, "nickroebuck@please.com"})
	users = append(users, User{"Steve", "Peterson", 9994443888, "steve123@please.com"})
	users = append(users, User{"Carl", "Sunny", 9993355888, "carlsunny@please.com"})

	for i, user := range users {
		log.Println(i, ":", user.FirstName, user.LastName, user.PhoneNumber, user.Email)
	}
}
