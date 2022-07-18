package main

import "fmt"

func appClosed() {
	fmt.Println("App Ditutup")
}

func runApp(error bool) {
	defer appClosed()
	if error {
		panic("APP ERROR MAS BRO")
	} else {
		fmt.Println("App Jalan aman...")
	}
}

func main() {
	runApp(false)
}
