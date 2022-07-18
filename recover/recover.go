package main

import "fmt"

func appClosed() {
	message := recover()
	if message != nil {
		fmt.Println("KETANGKEP PESAN ERRORNYA:", message)
	}
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
	runApp(true)
	fmt.Println("APP LANJUT JALAN")
}
