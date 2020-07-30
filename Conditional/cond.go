package main

import "fmt"

func main() {

	var angka int = 12

	if (angka % 2) == 0 {
		fmt.Printf("%d adalah genap", angka)
	} else {
		fmt.Printf("%d adalah ganjil", angka)
	}
}
