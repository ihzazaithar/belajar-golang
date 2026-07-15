package main

import (
	"fmt"
)

type NoKTP string

func main() {

	var NIK NoKTP = "00010100020"
	fmt.Println(NIK)

	/*
		// Untuk melihat tipe data aslinya gunakan
		// tempVar := reflect.TypeOf(namaVariable)
		// fmt.Println(tempVar.Kind())


		e := reflect.TypeOf(NIK) // Menghasilkan main.NoKTP
		fmt.Println(NIK.Kind()) // Menghasilkan string
		**/
}
