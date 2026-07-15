package main

import "fmt"

func main() {

	// Number
	var var_intTanpaTipeData = 222

	fmt.Println(var_intTanpaTipeData)

	var var_int int = -230 // Bisa memakai nilai negatif hingga positif
	/*
		int, int8, int 16, int32 atau rune, int64
		**/

	var var_uint uint = 320 // Hanya bisa memakai nilai positif
	/*
		uint, uint8 atau byte, uint16, uint32, uint64
		**/

	var var_float32 float32 = 32.223231324123123
	var var_float64 float64 = 32.223231324123123
	/*
		float32, float64, complex64, complex128
		**/

	fmt.Println(var_int)
	fmt.Println(var_uint)
	fmt.Println(var_float32)
	fmt.Println(var_float64)
	// fmt.Println(var_byte)

	fmt.Println("=========================================")
	// Boolean

	var var_boolTrue bool = true
	var var_boolFalse bool = false

	fmt.Println(var_boolTrue)
	fmt.Println(var_boolFalse)

	fmt.Println("=========================================")
	// String

	var var_string = "Ini adalah String"
	// var var_string2 = 'ini juga string?' // Tidak bisa menggunakan kutip 1 ''

	fmt.Println(var_string)
	// fmt.Println(var_string2)

	fmt.Println(len(var_string)) // Mengetahui Panjang atau jumlah dari suatu string

	fmt.Println(var_string[10]) // Mengambil atau menampilkan index ke X dari suatu string. Nilai yang dikembalikan adalah nilai ASCII dari karakter di index tersebut, bukan karakter "aslinya".

}
