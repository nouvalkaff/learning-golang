package main

import "fmt"

// ======For loop, range ======
// import "fmt"

// func main() {
// 	var hitung int8 = 0

// 	for hitung <= 3 {
// 		fmt.Println(hitung)
// 		hitung++
// 	}

// 	fmt.Println("Loop Selesai")

// 	fmt.Println("loop baru dimulai")

// 	for count := 0; count <= 3; count++ {
// 		fmt.Println("Aku berhitung", count)
// 	}
// 	fmt.Println("Aku selesai berhitung")

// 	fmt.Println("======================")
// 	pet := []string{"kuda", "macan", "harimau"}

// 	// Ganti nama variabel jadi underscore "_" agar bisa digunakan tanpa dipakai
// 	for _, animal := range pet {
// 		fmt.Println(animal)
// 	}

// 	fmt.Println("sudah semua binatang")

// 	mapData := make(map[string]string)
// 	mapData["name"] = "Nouval"
// 	mapData["domicile"] = "Indonesia"

// 	for key, value := range mapData {
// 		fmt.Println(key, "=", value)
// 	}
// }

// ======For loop, range ======

// ======Break & Continue======

func main() {
	// Break example
	for i := 0; i <= 7; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke-", i)
	}

	fmt.Println("=========================")

	// Continue example
	for i := 0; i <= 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke-", i)
	}
}

// ======Break & Continue======
