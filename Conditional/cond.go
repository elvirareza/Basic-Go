package main

import "fmt"

func main() {

	angka:= 12

	if (angka % 2) == 0 {
		fmt.Printf("%d adalah genap\n", angka)
	} else {
		fmt.Printf("%d adalah ganjil\n", angka)
	}

	angka= 25

	if (angka % 2) == 0 {
		fmt.Printf("%d adalah genap\n", angka)
	} else {
		fmt.Printf("%d adalah ganjil\n", angka)
	}
}
