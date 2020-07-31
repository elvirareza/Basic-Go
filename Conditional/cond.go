package main

import "fmt"

func main() {

	angka := 12 // diubah cara deklarasinya

	if angka % 2 == 0 { // dihilangkan kurung buka tutupnya
		fmt.Printf("%d adalah genap", angka)
	} else {
		fmt.Printf("%d adalah ganjil", angka)
	}
	
	// ga ada yang salah, ini juga bukan cara yang terbaik.
	// Hanya memberi referensi saja kalau bisa juga seperti ini
}
