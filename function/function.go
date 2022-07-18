package main

import "fmt"

func main() {
	var i int8 = 1
	for i <= 2 {
		sayHello()
		i++
	}

	helloPeople("Fatimah", 23)

	job := "Engineer"
	fmt.Println(myJob(job))

	warna1, _ := indonesia()
	fmt.Println(warna1)

	content1, _, content3 := threeStack()
	fmt.Println(content1, content3)

	myNumber := timesAll(2, 3, 10, 10)
	nums := []int32{2, 2, 2}

	fmt.Println(myNumber)
	fmt.Println(timesAll(nums...))

	byeBos := sayGoodBye

	fmt.Println(byeBos("Nopal"))

	sayHelloFiltered("Zetsuga", spamFilter)
	sayHelloFiltered("Anjing", spamFilter)

	blacklist := func(height uint8) bool {
		return height >= 160
	}

	heightFilter(170, blacklist)
}

// End of Main Function

func sayHello() {
	fmt.Println("Hello Moto!")
}

func helloPeople(name string, age int8) {
	fmt.Println("Hi", name, "are you", age, "years old", "?")
}

func myJob(job string) string {
	if job == "" {
		return "I am a deadwood"
	}
	return "My Job is a " + job
}

//Return multiple values
func indonesia() (string, string) {
	return "Merah", "Putih"
}

// Named return values
func threeStack() (firstStack, secondStack, thirdStack string) {
	firstStack = "Front-End Dev."
	secondStack = "Back-End Dev."
	thirdStack = "Mobile Dev."

	return
}

func timesAll(num ...int32) int32 {
	var total int32 = 1
	for _, nums := range num {
		total *= nums
	}
	return total
}

func sayGoodBye(name string) string {
	return "Bye " + name
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

// Type alias
type FilAlias func(string) string

func sayHelloFiltered(name string, filterFunction FilAlias) {
	dataFiltered := filterFunction(name)
	fmt.Println("Hello", dataFiltered)
}

type BlacklistAlias func(height uint8) bool

func heightFilter(height uint8, blacklist BlacklistAlias) {
	if blacklist(height) {
		fmt.Printf("You are %dcm, and passed the minimum height test of 160cm", height)
	} else {
		fmt.Printf("You are not passed the minimum height test of 160cm, you are %d", height)
	}
}
