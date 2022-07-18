package main

import "fmt"

func loop(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result
}

func recursive(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * recursive(num-1)
	}
}
func main() {
	fmt.Println(loop(5))
	fmt.Println(recursive(5))
}
