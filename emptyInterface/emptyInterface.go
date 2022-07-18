package main

import "fmt"

func tesInput(data int) interface{} {
	if data == 1 {
		return 1
	}

	if data == 2 {
		return "bukan 1 nih tapi 2"
	}

	if data != 1 && data != 2 {
		return "beda bos"
	}

	return ""
}

func main() {
	fmt.Println(tesInput(0))
}
