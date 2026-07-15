package main

import (
	"fmt"
)

func main() {

	var varString string = "Ini adalah string"

	fmt.Println(varString)

	var e = varString[1]

	fmt.Println(e) // Akan menampilkan nilai ASCII 110

	var eString = string(e) // Akan mengkonversi ASCII 110 menjadi huruf n
	fmt.Println(eString)

}
