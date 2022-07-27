package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	 result := random()
	//  resultString := result.(string)

	// fmt.Println(resultString)

	switch val := result.(type) {
	case string: fmt.Println(val, "is string")
	case int: fmt.Println(val, "is integer")
	case bool: fmt.Println(val, "is boolean")
	default: fmt.Print("Unknown type")
	}
}