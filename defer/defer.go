package main

import "fmt"

func logging() {
	fmt.Println("App ditutup")
}

func runApp(value int) {
	defer logging()
	fmt.Println("App jalan...")
	result := 10 / value
	fmt.Println("Hasil", result)
}
func main() {
	runApp(0)
}

// DEFER MEMBUAT FUNCTION TETAP JALAN DIAKHIR WALAUPUN APP ERROR
