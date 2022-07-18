package main

import "fmt"

type AnimalName interface {
	PetFamily() string
}

type Animal struct {
	Family string
}

// Any struct function that follow AnimalName interface can use sayFamily function
func (animal Animal) PetFamily() string {
	return animal.Family
}

func sayFamily(animalName AnimalName) {
	fmt.Println("My pet is from family", animalName.PetFamily())
}

func main() {
	var iggy = Animal{
		Family: "Iguanidae",
	}

	sayFamily(iggy)
}
