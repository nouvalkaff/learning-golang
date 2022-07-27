package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", " Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1
	address4 := &address1

	address2.City = "Pangalengan"
	address3 = &Address{"Bandung", " Jawa Barat", "Indonesia"}

	//create "destruction" variable
	*address4 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
}