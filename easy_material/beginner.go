package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// }

// =========Number=============
// func main() {
// 	fmt.Println("Satu = ", 1)
// 	fmt.Println("Dua = ", 2)
// 	fmt.Println("Tiga Koma Lima = ", 3.5)
// }
// =========Number=============

// ============Boolean===========
// func main() {
// 	fmt.Println("Benar: ", true)
// 	fmt.Println("Salah: ", false)
// }
// ============Boolean===========

// ============String===========
// func main() {
// 	fmt.Println(len("Aku adalah Anak Gembala"))
// 	fmt.Println("Aku adalah Anak Gembala"[0]) // print byte dari huruf "A"
// 	fmt.Println("Aku adalah Anak Gembala"[1]) // print byte dari huruf "k"
// }

// ============String===========

// ============Variable===========
// func main() {
// 	var name string // declared type, forced to be string

// 	name = "Nopal"
// 	fmt.Println(name)

// 	name = "Buday"
// 	fmt.Println(name)

// 	var frinedName = "Totti" // undeclared type, smartly detect data type string
// 	fmt.Println(frinedName)

// 	var age = 25 // undeclared type, smartly detect data type int
// 	fmt.Println(age)

// 	pet := "iguana"
// 	fmt.Println(pet)

// 	pet = "cat"
// 	fmt.Println(pet)

// 	var (
// 		adik1 = "Bimo"
// 		adik2 = "Kemal"
// 	)

// 	fmt.Println(adik1, adik2)
// }

// ============Variable===========

// ============Constant===========
// func main() {
// 	const pet string = "Rabbit"
// 	const name = "Ayu"
// 	const country = "Indonesia"

// 	const (
// 		jumlahAdik int8   = 2
// 		adik1      string = "Diah"
// 		adik2             = "Marni"
// 	)

// 	// For constant (const), it does not matter if it not being used. Beacause, Golang will not compile it.
// 	// fmt.Println(pet)
// 	fmt.Println(name)
// 	fmt.Println(country)
// 	fmt.Println(jumlahAdik)
// 	fmt.Println(adik1)
// 	fmt.Println(adik2)
// }

// ============Constant===========

// ============Convert Data Type===========
// func main() {
// 	var nilai32 int32 = 128
// 	var nilai64 int64 = int64(nilai32)
// 	var nilai8 int8 = int8(nilai32)

// 	fmt.Println(nilai32)
// 	fmt.Println(nilai64)
// 	fmt.Println(nilai8) // After reaching the max value of 127, it would back to the lowest value and increasing the value again.

// 	var pet = "Cat"
// 	var e byte = pet[0]
// 	var eString string = string(e)

// 	fmt.Println(pet)
// 	fmt.Println(e)
// 	fmt.Println(eString)
// }

// ============Convert Data Type===========

// ============Type Declarations===========
// func main() {
// 	type NoKTP string
// 	type Status bool

// 	var noKTPKu NoKTP = "55656255451415"
// 	fmt.Println(noKTPKu)

// 	var statusKawin Status = false
// 	fmt.Println(statusKawin)

// }

// ============Type Declarations===========

// ============Math===========
// func main() {
// 	var result = 10 + 10
// 	fmt.Println(result)

// 	const (
// 		a = 10
// 		b = 10
// 		c = a * b
// 	)

// 	fmt.Println(c)

// 	// Below is augmented assignment
// 	var angka = 25
// 	angka += 25
// 	fmt.Println(angka)

// 	// Below is unary operator
// 	var num = 10
// 	num++ // num = 10 + 1
// 	num++ // num = 11 + 1
// 	fmt.Println(num)
// }

// ============Math===========

// ============Compare Data===========
// func main() {
// 	const (
// 		a = 22
// 		b = 14
// 		c = 26
// 		d = 26
// 		e = 77
// 	)

// 	var result bool = a == c
// 	fmt.Println(result)

// 	result = c == d
// 	fmt.Println(result)

// 	result = c > d
// 	fmt.Println(result)
// }

// ============Compare Data===========

// ============Boolean Operation===========
// func main() {
// 	const (
// 		a = true
// 		b = false
// 		c = false
// 		d = true
// 	)

// 	result := a && b
// 	fmt.Println(result)

// 	result = b && c
// 	fmt.Println(result)

// 	result = a || b
// 	fmt.Println(result)

// 	result = a && d
// 	fmt.Println(result)

// }

// ============Boolean Operation===========

// ============Array===========
// func main() {
// 	var names [3]string

// 	names[0] = "Jamblang"
// 	names[1] = "Darius"
// 	names[2] = "Monolit"
// 	// names[3] = "Monolit" // in Go we cannot exceed the capacity of declared array

// 	fmt.Println(names)

// 	var primeNum = [3]int8{
// 		2, 3, 5,
// 	}

// 	fmt.Println(primeNum)

// 	var myPet = [2]string{
// 		"iguana", "cat",
// 	}
// 	fmt.Println(myPet)

// 	fmt.Println("Panjang array variable names:", len(names)) //len() to see the length of the array, not the amount of data within the array
// }

// ============Array===========

// ============Slice===========
/*
   SLICE Vs ARRAY
   In slices you dont have to mention the size for that particular slice. Since a slice doesnt have a specific size, we can append values to a slice.
   But in Arrays, we had to define its size when declaring.
*/
// Slice consists of 3 elements:
// 1. pointer (the first data in array of slice),
// 2. length (the length of the slice),
// 3. and capacity (the capacity of the slice, length cannot exceeds the capacity).
// func main() {
// 	var months = [...]string{
// 		"januari",
// 		"februari",
// 		"march",
// 		"april",
// 		"mei",
// 		"juni",
// 		"juli",
// 		"agustus",
// 		"september",
// 		"oktober",
// 		"november",
// 		"desember",
// 	}

// 	var slice1 = months[4:7]
// 	fmt.Println(slice1)
// 	fmt.Println(len(slice1)) // count mei, juni, juli = 3
// 	fmt.Println(cap(slice1)) // count from mei to desember = 8

// 	// months[5] = "diganti"
// 	// fmt.Println(slice1) // if array data is altered, slice data would also be altered, be careful!

// 	// slice1[0] = "meilani"
// 	// fmt.Println(months) // if slice data is altered, array data would also be altered, be careful!

// 	var slice2 = months[10:] // length = 2, capacity = 2 --> this slice if appended would exceed the capacity, so it would change the 'month' array
// 	// var slice2 = months[2:4] // length = 2, capacity = 10 --> this slice if appended would not exceed the capacity, so it would use the exisiting 'month' array
// 	fmt.Println(slice2)

// 	var slice3 = append(slice2, "Kuproy")
// 	// if the append is executed while it exceeds the capacity, it would be forming a new array
// 	// if the append is executed while the capacity is not exceeded, it would use the existing array
// 	fmt.Println(slice2, "slice2")
// 	fmt.Println(slice3, "slice3")
// 	/*
// 		--> get slice length: len(slice)
// 		--> get slice capacity: cap(slice)
// 	*/

// 	newSlice := make([]string, 2, 5)

// 	newSlice[0] = "Junaidi"
// 	newSlice[1] = "Latu"

// 	fmt.Println(newSlice, "newSlice")
// 	fmt.Println(len(newSlice))
// 	fmt.Println(cap(newSlice))

// 	copySlice := make([]string, len(newSlice), cap(newSlice))
// 	copy(copySlice, newSlice)
// 	fmt.Println(copySlice, "copySlice")
// }

// ============Slice===========

// ============Map===========
/*
   SLICE Vs ARRAY
   In slices you dont have to mention the size for that particular slice. Since a slice doesnt have a specific size, we can append values to a slice.
   But in Arrays, we had to define its size when declaring.
*/
// Slice consists of 3 elements:
// 1. pointer (the first data in array of slice),
// 2. length (the length of the slice),
// 3. and capacity (the capacity of the slice, length cannot exceeds the capacity).
// func main() {
// 	var animals map[string]string = map[string]string{
// 		"name": "cat",
// 		"leg":  "4",
// 	}
// 	animals["tail"] = "short"
// 	fmt.Println(animals)
// 	// fmt.Println(animals["name"])
// 	// fmt.Println(animals["leg"])
// 	// fmt.Println(animals["tail"])
// 	fmt.Println(len(animals))

// 	book := make(map[string]string)
// 	book["name"] = "Malin Kundang"
// 	book["price"] = "Rp 200.000"
// 	book["page"] = "90"
// 	book["unk"] = "-"

// 	fmt.Println(book, "atas")
// 	fmt.Println(len(book), "atas")
// 	delete(book, "unk")
// 	fmt.Println(book, "bawah")
// 	fmt.Println(len(book), "bawah")

// }

// ============Map===========
