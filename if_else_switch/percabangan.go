package main

import "fmt"

// =========if-elseif-else=========
// func main() {
// 	name := "Tono"
// 	name = "Alkaff"
// 	// name = "Nouval"
// 	if name == "Nouval" {
// 		fmt.Println("Halo Nouval")
// 	} else if name == "Alkaff" {
// 		fmt.Println("Hai keluarga")
// 	} else {
// 		fmt.Println("Anda Siapa?")
// 	}
// }
// =========if-elseif-else=========

// =========if short statement=========
// func main() {
// 	const name = "Mir"

// 	if length := len(name); length <= 3 {
// 		fmt.Println("Nama terlalu pendek")
// 	} else {
// 		fmt.Println("Terima kasih telah mengisi form ini")
// 	}
// }

// =========if short statement=========

// =========switch statement=========

func main() {
	var num int32 = 22

	switch num {
	case 22:
		fmt.Println("Input Angka 22 (Dua Puluh Dua)")
	case 0:
		fmt.Println("Input Angka 0 (Nol)")
	default:
		fmt.Println("Input Angka Selain 22 dan 0")
	}

	// switch short statement
	switch number := 21; number % 2 {
	case 0:
		fmt.Println("Input adalah angka genap")
	default:
		fmt.Println("Input adalah angka ganjil")
	}

	// switch without condition
	var angka = 99
	switch {
	case angka%2 == 0:
		fmt.Println("Input bawah adalah angka genap")
	case angka%2 != 0:
		fmt.Println("Input bawah adalah angka ganjil")
	}

	// switch without condition
	pet := "kuda"
	length := len(pet)

	switch {
	case length <= 3:
		fmt.Println("nama jenis pet terlalu pendek")
	case length > 12:
		fmt.Println("nama jenis pet terlalu panjang")
	default:
		fmt.Println("nama jenis pet sudah cocok")
	}

}

// =========switch statement=========
