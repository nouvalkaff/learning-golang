package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (animal Animal) animalType(tipe string) {
	fmt.Println("This is my", tipe, "its name is", animal.name, "now it is", animal.age, "years old")
}

func main() {
	myPet := Animal{
		name: "iguana",
		age:  3,
	}

	myPet.animalType("iguana")

}
